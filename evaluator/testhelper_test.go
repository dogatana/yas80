package evaluator

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
	"yas80/logger"
	"yas80/object"
	"yas80/parser"
)

func evaluateInput(t *testing.T, input string, logger *logger.Logger, env *object.Environment) *object.ProgramObject {
	progNode := parseTextForTest(t, input)

	evaluator := New(logger)
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
