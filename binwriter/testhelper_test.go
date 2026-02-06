package binwriter

import (
	"bytes"
	"yas80/evaluator"
	"yas80/filecontent"
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

func evalInput(input any, logger *logging.Logger, env object.Environment) (*object.BlockObject, *evaluator.Evaluator) {
	prog := parseTextForTest(input, logger)

	eval := evaluator.New(logger)

	var obj object.Object
	pass := 0

	eval.Resolved = true
	for i := 0; i < 256; i++ {
		pass++
		eval.Resolved = true
		obj = eval.EvalProgram(prog, pass, env)
		// eval.EvalEnv(env)
		eval.CheckCyclicError(env)
		if logger.ErrorCount() > 0 {
			return &object.BlockObject{}, eval
		}
		if eval.Resolved {
			break
		}
	}
	eval.CheckSymbolError(env)
	if logger.ErrorCount() > 0 || !eval.Resolved {
		return obj.(*object.BlockObject), eval
	}

	// finalize
	code := object.CollectCode(obj.(*object.BlockObject).Block)
	eval.CodeStable = false
	for i := 0; i < 256 && !eval.CodeStable; i++ {
		pass++
		obj = eval.EvalProgram(prog, pass, env)
		if logger.ErrorCount() > 0 {
			return obj.(*object.BlockObject), eval

		}
		newCode := object.CollectCode(obj.(*object.BlockObject).Block)
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

func parseTextForTest(input any, logger *logging.Logger) *parser.BlockStatement {
	var prog *parser.BlockStatement

	fcs := []*filecontent.FileContent{}
	switch input := input.(type) {
	case string:
		fc, _ := filecontent.NewFromString("text", input)
		fcs = append(fcs, fc)
	case []string:
		for _, s := range input {
			fc, _ := filecontent.NewFromString("text", s)
			fcs = append(fcs, fc)
		}
	}
	index := 0
	callback := func() *filecontent.FileContent {
		if index >= len(fcs) {
			return nil
		}
		fc := fcs[index]
		index++
		return fc
	}

	lex := parser.NewLexer(logger, callback)
	prog = parser.Parse(lex)
	if logger.ErrorCount() > 0 {
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
