package evaluator

import (
	"bytes"
	"fmt"
	"testing"
	"yas80/filecontent"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

type symValue struct {
	name     string
	expected any
}

func evalInput(input string, logger *logging.Logger, env TEnv) (*object.BlockObject, *Evaluator) {
	prog := parseTextForTest(input, logger)

	eval := New(logger)

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
	code := CollectCode(obj.(*object.BlockObject).Block)
	eval.CodeStable = false
	for i = 0; i < 256 && !eval.CodeStable; i++ {
		obj = eval.EvalProgram(prog, env)
		if len(logger.Errors) > 0 {
			return obj.(*object.BlockObject), eval

		}
		newCode := CollectCode(obj.(*object.BlockObject).Block)
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

func testEvalResult(t *testing.T, tn int, err string, eval *Evaluator) {
	if err == "" {
		if len(eval.logger.Errors) > 0 {
			eval.logger.Print()
			t.Fatalf("[%d] EvalProgram() %d errors", tn, len(eval.logger.Errors))
		}
		if !eval.Resolved {
			t.Fatalf("[%d] eval.Resolved false", tn)
		}
		if !eval.CodeStable {
			t.Fatalf("[%d] eval.CodeStable false", tn)
		}
	} else {
		e, w, i := eval.logger.Count()
		if e == 0 && w == 0 && i == 0 {
			fmt.Printf("[%d] logger.Print()\n", tn)
			eval.logger.Print()
			t.Fatalf("[%d] no logmessages", tn)
		}
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

func testCodeResult(t *testing.T, tn int, expected []byte, prog *object.BlockObject) {
	if len(expected) == 0 {
		return
	}
	code := CollectCode(prog.Block)
	if err := testutil.BytesEqual(code, expected); err != nil {
		t.Errorf("[%d] generated code diff %s", tn, err.Error())
	}
}

// sym.Expected nil : 存在するとエラー
// sym.Expected int, string : 存在しないとエラー
func testSymValues(t *testing.T, tn int, syms []symValue, getter func(name string) (*object.SymbolObject, bool)) {
	for _, sym := range syms {
		s, ok := getter(sym.name)
		switch {
		case !ok && sym.expected != nil:
			t.Errorf("[%d] symbol %s not found", tn, sym.name)
		case !ok:
			// do nothing
		case ok && sym.expected == nil:
			t.Errorf("[%d] symbol %s exists", tn, sym.name)
		case ok:
			testSymbolValue(t, tn, s, sym.expected)
		}
	}
}

func testSymbolValue(t *testing.T, tn int, obj object.Object, expected any) bool {
	sym, ok := obj.(*object.SymbolObject)
	if !ok {
		t.Errorf("[%d] Object not SymbolObject. got %T", tn, obj)
		return false
	}
	switch expected := expected.(type) {
	case int:
		return testNumberObject(t, tn, sym.Value, expected)
	case string:
		return testStringObject(t, tn, sym.Value, expected)
	default:
		t.Fatalf("[%d] expected: unknown type %T", tn, expected)
		return false
	}
}

func testNumberObject(t *testing.T, tn int, obj object.Object, expected int) bool {
	number, ok := obj.(*object.NumberObject)
	if !ok {
		t.Errorf("[%d] Object not NumberObject. got %T", tn, obj)
		return false
	}
	if number.Value != expected {
		t.Errorf("[%d] object is not %d. got %d", tn, expected, number.Value)
		return false
	}
	return true
}

func testStringObject(t *testing.T, tn int, obj object.Object, expected string) bool {
	s, ok := obj.(*object.StringObject)
	if !ok {
		t.Errorf("[%d] Object not StringObject. got %T", tn, obj)
		return false
	}
	if s.Value != expected {
		t.Errorf("[%d] object is not %s. got %s", tn, expected, s.Value)
		return false
	}
	return true
}
