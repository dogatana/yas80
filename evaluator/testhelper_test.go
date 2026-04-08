package evaluator

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/dogatana/yas80/binwriter"
	"github.com/dogatana/yas80/filecontent"
	"github.com/dogatana/yas80/internal/testutil"
	"github.com/dogatana/yas80/logging"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
)

type symValue struct {
	name     string
	expected any
}

func evalInput(input string, logger *logging.Logger, env TEnv) (*object.BlockObject, *Evaluator) {
	prog := parseTextForTest(input, "testdata/text", logger) // testdata/*.json を読み込むため
	return evalProg(prog, logger, env)
}

func evalFile(filename string, logger *logging.Logger, env TEnv) ([]byte, bool) {
	prog := parseFileForTest(filename, logger)
	obj, _ := evalProg(prog, logger, env)

	bw := binwriter.New(obj, 0, logger)
	var buf bytes.Buffer
	ok := bw.WriteBin(&buf)
	return buf.Bytes(), ok
}

func evalProg(prog *parser.BlockStatement, logger *logging.Logger, env TEnv) (*object.BlockObject, *Evaluator) {

	// -I testdata の設定
	eval := New(logger, []string{"testdata"})

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
	eval.Stage2 = true
	env.Set("$STAGE2", &object.NumberObject{Value: 1})
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

func testEvalResult(t *testing.T, tn int, err string, eval *Evaluator) {
	if err == "" {
		if ec := eval.logger.ErrorCount(); ec > 0 {
			eval.logger.Print()
			t.Fatalf("[%d] EvalProgram() %d errors", tn, ec)
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

func parseTextForTest(input string, filename string, logger *logging.Logger) *parser.BlockStatement {
	var prog *parser.BlockStatement

	fc, _ := filecontent.NewFromString(filename, input)

	lex := parser.NewLexer(logger, func() *filecontent.FileContent {
		ret := fc
		fc = nil
		return ret
	})
	prog = parser.Parse(lex, []string{})
	if logger.ErrorCount() > 0 {
		return prog
	}

	prog = parser.PreProrocess(prog)
	return prog
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
	prog = parser.Parse(lex, []string{})
	if logger.ErrorCount() > 0 {
		return prog
	}

	// prog = parser.PreProrocess(logger, prog)
	return prog
}

func testCodeResult(t *testing.T, tn int, expected []byte, prog *object.BlockObject) {
	if len(expected) == 0 {
		return
	}
	code := object.CollectCode(prog.Block)
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
