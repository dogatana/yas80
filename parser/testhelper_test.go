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
	ret := []string{}
	for _, s := range strs {
		str := strings.Trim(s, " \n\t")
		if str != "" {
			ret = append(ret, str)
		}
	}
	return strings.Join(ret, "\n")
}
