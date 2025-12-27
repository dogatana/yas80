package evaluator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"yas80/fileblock"
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

func evalInput(input string, logger *logging.Logger, env object.Environment) (*object.ProgramObject, *Evaluator) {
	progNode := parseTextForTest(input)

	eval := New(logger)

	var obj object.Object
	var i int

	eval.Resolved = true
	for i = 0; i < 256; i++ {
		eval.Resolved = true
		obj = eval.EvalProgram(progNode, env)
		eval.EvalEnv(env)
		eval.CheckSymbols(env)
		if len(logger.Errors) > 0 {
			return &object.ProgramObject{}, eval
		}
		if eval.Resolved {
			break
		}
	}
	if !eval.Resolved {
		return &object.ProgramObject{}, eval
	}

	// finalize
	code := CollectCode(obj.(*object.ProgramObject))
	eval.CodeStable = false
	for i = 0; i < 256 && !eval.CodeStable; i++ {
		obj = eval.EvalProgram(progNode, env)
		if len(logger.Errors) > 0 {
			return &object.ProgramObject{}, eval

		}
		newCode := CollectCode(obj.(*object.ProgramObject))
		eval.CodeStable = bytes.Equal(code, newCode)
		if !eval.CodeStable {
			code = newCode
		}
	}
	if !eval.CodeStable {
		return &object.ProgramObject{}, eval
	}
	if prog, ok := obj.(*object.ProgramObject); !ok {
		return &object.ProgramObject{}, eval
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
			t.Fatalf("[%d] eval.Resolved fasle", tn)
		}
		if !eval.CodeStable {
			t.Fatalf("[%d] eval.CodeStable fasle", tn)
		}
	} else {
		e, w, i := eval.logger.Count()
		if e == 0 && w == 0 && i == 0 {
			t.Fatalf("[%d] no logmessages", tn)
		}
	}
}

func parseTextForTest(input string) *parser.Program {
	var prog *parser.Program

	file := "<string>"
	logger := logging.New(file)
	fb := fileblock.New(file, []byte(input))

	l := parser.NewLexer(fb, logger)
	prog = parser.Parse(l)
	if len(logger.Errors) > 0 {
		return prog
	}

	prog = parser.PreProrocess(logger, prog)
	return prog
}

func evalValue(obj object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.SymbolObject:
		return obj.Value
	default:
		return obj
	}
}

func testSymbolNumberObject(t *testing.T, tn int, obj object.Object, expected int) bool {
	sym, ok := obj.(*object.SymbolObject)
	if !ok {
		t.Errorf("[%d] Object not SymbolObject. got %T", tn, obj)
		return false
	}
	return testNumberObject(t, tn, sym.Value, expected)
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

func readTestDataFile(t *testing.T, filename string) []byte {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}
	path := filepath.Join(filepath.Dir(file), "testdata", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%q) returns %s", path, err.Error())
	}
	return data
}

func collectValue(prog *object.ProgramObject) []*object.ValueObject {
	var result []*object.ValueObject
	for _, obj := range prog.Objects {
		value, ok := obj.(*object.ValueObject)
		if !ok {
			continue
		}
		result = append(result, value)
	}
	return result
}

func bytesEqual(v1, v2 []byte) error {
	if len(v1) != len(v2) {
		return fmt.Errorf("size diff %d %d\n", len(v1), len(v2))
	}
	for i, v := range v1 {
		if v != v2[i] {
			return fmt.Errorf("contentis diff [%d] %02x %02x\n", i, v, v2[i])
		}
	}
	return nil
}
