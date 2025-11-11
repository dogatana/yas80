package parser

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
	"yas80/logger"
)

func newLexerForTest(input string) *Lexer {
	file := "<string>"
	es := logger.New(file)
	return NewLexer(bufio.NewReader(strings.NewReader(input)), file, es)
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
