package parser

import (
	"bufio"
	"strings"
	"yas80/errorstore"
)

func newLexerForTest(input string) *Lexer {
	es := errorstore.New()
	return NewLexer(bufio.NewReader(strings.NewReader(input)), "<string>", es)
}
