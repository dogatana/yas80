package parser

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"yas80/fileblock"
	"yas80/logger"
)

func newLexerForTest(input string) *Lexer {
	file := "<string>"
	logger := logger.New(file)
	fb := fileblock.New(file, []byte(input))
	return NewLexer(fb, logger)
}

func ParseForTest(t *testing.T, lexer *Lexer, tn int) *Program {
	prog := Parse(lexer)
	prog = PreProrocess(lexer.logger, prog)
	ec, wc, _ := lexer.logger.Count()
	if ec > 0 || wc > 0 {
		lexer.logger.Print()
		if tn < 0 {
			t.Fatalf("%d errors and %d warnigs", ec, wc)
		} else {
			t.Fatalf("[%d] %d errors and %d warnigs", tn, ec, wc)
		}
	}
	return prog
}

func splitTrim(input string) string {
	strs := strings.Split(strings.ReplaceAll(input, "\\", "\n"), "\n")
	ret := []string{}
	for _, s := range strs {
		str := strings.Trim(s, " \n\t")
		if str != "" {
			ret = append(ret, str)
		}
	}
	return strings.Join(ret, "\n")
}

func testAssignStatement(t *testing.T, tn int, node Node) *AssignStatement {
	stmt, ok := node.(*AssignStatement)
	if !ok {
		t.Fatalf("[%d] not *AssignStatement. got %T", tn, node)
	}
	return stmt
}

func testNumberLiteral(t *testing.T, tn int, node Node, expected int) {
	literal, ok := node.(*NumberLiteral)
	if !ok {
		t.Errorf("[%d] not *NumberLiteral. got %T", tn, literal)
	}
	if literal.Value != expected {
		t.Errorf("[%d] not %d. got %d", tn, expected, literal.Value)
	}
}

func testStringLiteral(t *testing.T, tn int, node Node, expected string) {
	literal, ok := node.(*StringLiteral)
	if !ok {
		t.Errorf("[%d] not *StringLiteral. got %T", tn, literal)
	}
	if literal.Value != expected {
		t.Errorf("[%d] not %q. got %q", tn, expected, literal.Value)
	}
}
func readTestDataFile(t *testing.T, filename string) string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}
	path := filepath.Join(filepath.Dir(file), "testdata", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("could read %s", path)
	}
	return string(data)
}
