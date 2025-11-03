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
