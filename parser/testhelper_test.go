package parser

import (
	"bufio"
	"strings"
	"yas80/errorstore"
)

func newLexerForTest(input string) *Lexer {
	file := "<string>"
	es := errorstore.New(file)
	return NewLexer(bufio.NewReader(strings.NewReader(input)), file, es)
}

func splitTrim(input string) string {
	strs := strings.Split(strings.ReplaceAll(input, "\\", "\n"), "\n")
	for i, s := range strs {
		strs[i] = strings.Trim(s, " ")
	}
	return strings.Join(strs, "\n")
}
