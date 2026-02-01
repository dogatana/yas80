package binwriter

import (
	"bytes"
	"yas80/evaluator"
	"yas80/filecontent"
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

func evalInput(input string, logger *logging.Logger, env object.Environment) (*object.BlockObject, *evaluator.Evaluator) {
	prog := parseTextForTest(input, logger)

	eval := evaluator.New(logger)

	var obj object.Object
	var i int

	eval.Resolved = true
	for i = 0; i < 256; i++ {
		eval.Resolved = true
		obj = eval.EvalProgram(prog, env)
		// eval.EvalEnv(env)
		eval.CheckCyclicError(env)
		if len(logger.Errors) > 0 {
			return &object.BlockObject{}, eval
		}
		if eval.Resolved {
			break
		}
	}
	eval.CheckSymbolError(env)
	if len(logger.Errors) > 0 || !eval.Resolved {
		return obj.(*object.BlockObject), eval
	}

	// finalize
	code := evaluator.CollectCode(obj.(*object.BlockObject).Block)
	eval.CodeStable = false
	for i = 0; i < 256 && !eval.CodeStable; i++ {
		obj = eval.EvalProgram(prog, env)
		if len(logger.Errors) > 0 {
			return obj.(*object.BlockObject), eval

		}
		newCode := evaluator.CollectCode(obj.(*object.BlockObject).Block)
		eval.CodeStable = bytes.Equal(code, newCode)
		if !eval.CodeStable {
			code = newCode
		}
	}
	if !eval.CodeStable {
		return obj.(*object.BlockObject), eval
	}
	if prog, ok := obj.(*object.BlockObject); !ok {
		return prog, eval
	} else {
		return prog, eval
	}
}

func parseTextForTest(input string, logger *logging.Logger) *parser.BlockStatement {
	var prog *parser.BlockStatement

	fc := filecontent.New(logger.Filename, []byte(input))

	lex := parser.NewLexer(logger, func() *filecontent.FileContent {
		ret := fc
		fc = nil
		return ret
	})
	prog = parser.Parse(lex)
	if len(logger.Errors) > 0 {
		return prog
	}

	prog = parser.PreProrocess(logger, prog)
	return prog
}

func codeFromObj(obj object.Object, fill int, logger *logging.Logger) ([]byte, bool) {
	bw := New(obj, fill, logger)

	var buf bytes.Buffer
	ok := bw.Write(&buf)

	return buf.Bytes(), ok
}
