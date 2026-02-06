package lister

import (
	"bytes"
	"fmt"
	"testing"
	"yas80/binwriter"
	"yas80/evaluator"
	"yas80/filecontent"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

func evalFile(filename string, logger *logging.Logger, env object.Environment) ([]byte, bool) {
	prog := parseFileForTest(filename, logger)
	obj, _ := evalProg(prog, logger, env)

	bw := binwriter.New(obj, 0, logger)
	var buf bytes.Buffer
	ok := bw.Write(&buf)
	return buf.Bytes(), ok
}

func evalProg(prog *parser.BlockStatement, logger *logging.Logger, env object.Environment) (*object.BlockObject, *evaluator.Evaluator) {

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
	code := evaluator.CollectCode(obj.(*object.BlockObject).Block)
	eval.CodeStable = false
	for i := 0; i < 256 && !eval.CodeStable; i++ {
		pass++
		obj = eval.EvalProgram(prog, pass, env)
		if logger.ErrorCount() > 0 {
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

func parseFileForTest(filename string, logger *logging.Logger) *parser.BlockStatement {
	var prog *parser.BlockStatement

	fc, err := filecontent.NewFromFile(filename)
	if err != nil {
		fmt.Println("parseFileForTest", err.Error())

	}

	lex := parser.NewLexer(logger, func() *filecontent.FileContent {
		ret := fc
		fc = nil
		return ret
	})
	prog = parser.Parse(lex)
	if logger.ErrorCount() > 0 {
		return prog
	}

	prog = parser.PreProrocess(logger, prog)
	return prog
}

func testCodeResult(t *testing.T, tn int, expected []byte, prog *object.BlockObject) {
	if len(expected) == 0 {
		return
	}
	code := evaluator.CollectCode(prog.Block)
	if err := testutil.BytesEqual(code, expected); err != nil {
		t.Errorf("[%d] generated code diff %s", tn, err.Error())
	}
}
