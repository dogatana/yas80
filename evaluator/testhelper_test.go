package evaluator

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"testing"
	"yas80/logger"
	"yas80/object"
	"yas80/parser"
)

func evaluateInput(t *testing.T, input string, logger *logger.Logger, env *object.Environment) *object.ProgramObject {
	progNode := parseTextForTest(t, input)

	evaluator := New(logger)
	checkDebug(evaluator)

	obj := evaluator.Eval(progNode, env)
	ec, wc, _ := logger.Count()
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
	if len(programObject.Objects) == 0 {
		fmt.Printf("input %q\n", input)
		t.Fatal("Eval() return 0 Objects")
	}
	return programObject
}

func checkDebug(e *Evaluator) {
	val := os.Getenv("yydebug")
	if val == "" {
		e.debug = 0
		return
	}
	n, err := strconv.Atoi(val)
	if err != nil {
		e.debug = 0
	}
	e.debug = n
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

func parseTextForTest(t *testing.T, input string) *parser.Program {
	file := "<string>"
	es := logger.New(file)
	l := parser.NewLexer(bufio.NewReader(strings.NewReader(input)), file, es)
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
	return prog
}
