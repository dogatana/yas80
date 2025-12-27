package evaluator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	"yas80/fileblock"
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

func evaluateInput(t *testing.T, input string, logger *logging.Logger, env object.Environment) (*object.ProgramObject, *Evaluator) {
	progNode := parseTextForTest(t, input)

	evaluator := New(logger)
	checkDebug(evaluator)

	evaluator.Resolved = true
	var obj object.Object
	var i int
	for i = 0; i < 256; i++ {
		evaluator.Resolved = true
		obj = evaluator.EvalProgram(progNode, env)
		evaluator.EvalEnv(env)
		evaluator.CheckSymbols(env)
		ec, wc, _ := logger.Count()
		if ec > 0 || wc > 0 {
			fmt.Printf("input %q\n", input)
			logger.Print()
			t.Fatalf("EvalProgram() %d errors and %d warnigs", ec, wc)
		}
		if evaluator.Resolved {
			break
		}
	}
	// finalize
	code := CollectCode(obj.(*object.ProgramObject))
	codeStable := false
	for i = 0; i < 256 && !codeStable; i++ {
		obj = evaluator.EvalProgram(progNode, env)
		newCode := CollectCode(obj.(*object.ProgramObject))
		codeStable = bytes.Equal(code, newCode)
		if !codeStable {
			code = newCode
		}
	}

	programObject, ok := obj.(*object.ProgramObject)
	if !ok {
		fmt.Printf("input %q\n", input)
		t.Fatalf("not ProgramObject. got %T", obj)
	}
	if len(programObject.Objects) == 0 {
		fmt.Printf("input %q\n", input)
		t.Fatal("Eval() return 0 Objects")
	}
	return programObject, evaluator
}

func parseTextForTest(t *testing.T, input string) *parser.Program {
	file := "<string>"
	logger := logging.New(file)
	fb := fileblock.New(file, []byte(input))
	l := parser.NewLexer(fb, logger)
	prog := parser.Parse(l)
	ec, wc, _ := l.Logger().Count()
	if ec > 0 || wc > 0 {
		fmt.Printf("input %q\n", input)
		l.Logger().Print()
		t.Fatalf("Parse() %d errors and %d warnigs", ec, wc)
	}

	if len(prog.Statements) == 0 {
		fmt.Printf("input %q\n", input)
		t.Fatal("Parse() returns 0 statements")
	}

	prog = parser.PreProrocess(logger, prog)
	ec, wc, _ = l.Logger().Count()
	if ec > 0 || wc > 0 {
		fmt.Printf("input %q\n", input)
		l.Logger().Print()
		t.Fatalf("PreProcess() %d errors and %d warnigs", ec, wc)
	}

	return prog
}

func checkDebug(e *Evaluator) {
	val := os.Getenv("evaldebug")
	if val == "" {
		e.Debug = 0
		return
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		e.Debug = 0
	}
	e.Debug = n
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
