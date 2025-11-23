package parser

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"yas80/logger"
)

func newLexerForTest(input string) *Lexer {
	file := "<string>"
	es := logger.New(file)
	return NewLexer(bufio.NewReader(strings.NewReader(input)), file, es)
}

func ParseForTest(t *testing.T, lexer *Lexer, input string) *Program {
	prog := Parse(lexer)
	ec, wc, _ := lexer.logger.Count()
	if ec > 0 || wc > 0 {
		fmt.Printf("input %q\n", input)
		t.Fatalf("%d errors and %d warnigs", ec, wc)
	}
	prog = PreProrocess(prog)
	ec, wc, _ = lexer.logger.Count()
	if ec > 0 || wc > 0 {
		fmt.Printf("input %q\n", input)
		t.Fatalf("%d errors and %d warnigs", ec, wc)
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

func testExpressionStatement(t *testing.T, input string, node Node) *ExpressionStatement {
	stmt, ok := node.(*ExpressionStatement)
	if !ok {
		fmt.Printf("input %q\n", input)
		t.Fatalf("not *ExpressionStatement. got %T", node)
		return nil // ここへは到達しないが形式として必要
	}
	return stmt
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
