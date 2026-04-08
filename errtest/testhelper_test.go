package errtest

import (
	"bytes"

	"github.com/dogatana/yas80/evaluator"
	"github.com/dogatana/yas80/filecontent"
	"github.com/dogatana/yas80/logging"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
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

	e := evaluator.New(logger, []string{})
	e.Resolved = true
	var obj object.Object
	pass := 0
	for range 256 {
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
	if logger.ErrorCount() > 0 || !e.Resolved {
		return
	}
	// finalize
	code := object.CollectCode(obj.(*object.BlockObject).Block)
	codeStable := false
	for i := 0; i < 256 && !codeStable; i++ {
		pass++
		obj = e.EvalProgram(prog, pass, env)
		newCode := object.CollectCode(obj.(*object.BlockObject).Block)
		codeStable = bytes.Equal(code, newCode)
		if !codeStable {
			code = newCode
		}
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
	return parser.Parse(lex, []string{})
}

func getCount(logger *logging.Logger) []int {
	e, w, i := logger.Count()
	return []int{e, w, i}
}
