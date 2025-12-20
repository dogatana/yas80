package errtest

import (
	"bytes"
	"fmt"
	"regexp"
	"strings"
	"testing"
	"yas80/evaluator"
	"yas80/fileblock"
	"yas80/logger"
	"yas80/object"
	"yas80/parser"
)

func evaluateErrorInput(input string, logger *logger.Logger, env object.Environment) {
	prog := parseErrorTextForTest(input, logger)
	if len(logger.Errors) > 0 {
		return
	}

	e := evaluator.New(logger)
	e.Resolved = true
	var obj object.Object
	var i int
	for i = 0; i < 256; i++ {
		e.Resolved = true
		obj = e.EvalProgram(prog, env)
		e.EvalEnv(env)
		e.CheckSymbols(env)
		if len(logger.Errors) > 0 {
			return
		}
		if e.Resolved {
			break
		}
	}
	// finalize
	code := evaluator.CollectCode(obj.(*object.ProgramObject))
	codeStable := false
	for i = 0; i < 256 && !codeStable; i++ {
		obj = e.EvalProgram(prog, env)
		newCode := evaluator.CollectCode(obj.(*object.ProgramObject))
		codeStable = bytes.Equal(code, newCode)
		if !codeStable {
			code = newCode
		}
	}
}

func parseErrorTextForTest(input string, logger *logger.Logger) *parser.Program {
	file := "<string>"
	fb := fileblock.New(file, []byte(input))
	l := parser.NewLexer(fb, logger)
	return parser.Parse(l)
}

func hasError(logger *logger.Logger, expected string) bool {
	re := regexp.MustCompile(`\.?%.\.?`)
	ss := re.Split(expected, -1)

	for _, emsg := range logger.Errors {
		result := true
		for _, s := range ss {
			if !strings.Contains(emsg.Message, s) {
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
	if name, ok := errcodeNames[msg]; ok {
		return name
	}
	panic(fmt.Sprintf("not found '%s", msg))
}

func testError(t *testing.T, tn int, logger *logger.Logger, expected string) {
	if len(logger.Errors) == 0 {
		t.Fatalf("[%d] no error", tn)
	}
	if !hasError(logger, expected) {
		t.Errorf("[%d] not [%s] \"%s\" but \"%s\"",
			tn,
			getErrcodeName(expected),
			expected,
			logger.Errors[0].Message)
	}
}
