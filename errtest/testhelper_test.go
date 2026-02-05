package errtest

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"testing"
	"yas80/evaluator"
	"yas80/filecontent"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

const (
	TEST_ERROR = iota
	TEST_WARNING
	TEST_INFORMATION
)

func evaluateInput(testType int, input string, logger *logging.Logger, env object.Environment) {
	prog := parseText(input, logger)

	if getCount(logger)[testType] > 0 {
		return
	}

	e := evaluator.New(logger)
	e.Resolved = true
	var obj object.Object
	pass := 0
	for i := 0; i < 256; i++ {
		pass++
		e.Resolved = true
		obj = e.EvalProgram(prog, pass, env)
		// e.EvalEnv(env)
		// e.CheckSymbols(env)
		e.CheckCyclicError(env)
		if getCount(logger)[testType] > 0 {
			return
		}
		if e.Resolved {
			break
		}
	}
	e.CheckSymbolError(env)
	if len(logger.Errors) > 0 || !e.Resolved {
		return
	}
	// finalize
	code := evaluator.CollectCode(obj.(*object.BlockObject).Block)
	codeStable := false
	for i := 0; i < 256 && !codeStable; i++ {
		pass++
		obj = e.EvalProgram(prog, pass, env)
		newCode := evaluator.CollectCode(obj.(*object.BlockObject).Block)
		codeStable = bytes.Equal(code, newCode)
		if !codeStable {
			code = newCode
		}
	}
}

func testMessage(t *testing.T, testType int, tn int, logger *logging.Logger, expected string) {

	//  logger.LogMessage is not a type エラー回避
	// var messages []logger.LogMessage
	messages := logger.Errors

	switch testType {
	case TEST_WARNING:
		messages = logger.Warnings
	case TEST_INFORMATION:
		messages = logger.Infomation
	}

	if len(messages) == 0 {
		t.Fatalf("[%d] no error", tn)
	}
	if !hasMessage(messages, expected) {
		logger.Print()
		t.Errorf("[%d] not [%s] \"%s\" => \"%s\"",
			tn,
			getErrcodeName(expected),
			expected,
			logger.Errors[0].Error())
	}
}

func parseText(input string, logger *logging.Logger) *parser.BlockStatement {
	file := "<string>"
	fc, _ := filecontent.NewFromString(file, input)
	lex := parser.NewLexer(logger, func() *filecontent.FileContent {
		ret := fc
		fc = nil
		return ret
	})
	return parser.Parse(lex)
}

func getCount(logger *logging.Logger) []int {
	e, w, i := logger.Count()
	return []int{e, w, i}
}

func hasMessage(messages []logging.LogMessage, expected string) bool {
	re := regexp.MustCompile(`\.?%.\.?`)
	ss := re.Split(expected, -1)

	for _, emsg := range messages {
		result := true
		for _, s := range ss {
			if !strings.Contains(emsg.Message(), s) {
				result = false
				break
			}
		}
		if result {
			return result
		}
	}
	return false
}

func getErrcodeName(msg string) string {
	if name, ok := testutil.ErrcodeNames[msg]; ok {
		return name
	}
	panic(fmt.Sprintf("not found '%s", msg))
}
