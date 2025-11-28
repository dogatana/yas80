package evaluator

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"testing"
	"yas80/fileblock"
	"yas80/logger"
	"yas80/object"
	"yas80/parser"
)

func evaluateInput(t *testing.T, input string, logger *logger.Logger, env object.Environment) *object.ProgramObject {
	progNode := parseTextForTest(t, input)

	evaluator := New(logger)
	checkDebug(evaluator)

	// pass1
	_ = evaluator.Eval(progNode, env)
	ec, wc, _ := logger.Count()
	if ec > 0 || wc > 0 {
		fmt.Printf("input %q\n", input)
		logger.Print()
		t.Fatalf("Eval() %d errors and %d warnigs", ec, wc)
	}

	// resolve forward reference
	_, err := evaluator.EvalEnv(env)
	if err != nil {
		fmt.Println("input")
		fmt.Println(input)
		t.Fatalf("EvalEnv() error: %v", err)
	}

	// pass2
	evaluator.Pass1 = false
	obj := evaluator.Eval(progNode, env)
	ec, wc, _ = logger.Count()
	if ec > 0 || wc > 0 {
		fmt.Printf("input %q\n", input)
		logger.Print()
		t.Fatalf("Eval() %d errors and %d warnigs", ec, wc)
	}

	programObject, ok := obj.(*object.ProgramObject)
	if !ok {
		fmt.Printf("input %q\n", input)
		t.Fatalf("not ProgramObject. got %T", obj)
	}
	// if len(programObject.Objects) == 0 {
	// 	fmt.Printf("input %q\n", input)
	// 	t.Fatal("Eval() return 0 Objects")
	// }
	return programObject
}

func parseTextForTest(t *testing.T, input string) *parser.Program {
	file := "<string>"
	logger := logger.New(file)
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

func testNumberObject(t *testing.T, obj object.Object, expected int, input string) bool {
	number, ok := obj.(*object.NumberObject)
	if !ok {
		fmt.Printf("input %q\n", input)
		t.Errorf("Object not NumberObject. got %T", obj)
		return false
	}
	if number.Value != expected {
		fmt.Printf("input %q\n", input)
		t.Errorf("object is not %d. got %d", expected, number.Value)
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

func collectCode(prog *object.ProgramObject) []byte {
	var result []byte
	for _, obj := range prog.Objects {
		code, ok := obj.(*object.CodeObject)
		if !ok {
			continue
		}
		result = append(result, code.Code...)
	}
	return result
}

func bytesEqual(sa, sb []byte) bool {
	if len(sa) != len(sb) {
		return false
	}
	for i, b := range sa {
		if b != sb[i] {
			return false
		}
	}
	return true
}
