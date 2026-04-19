package parser

import (
	"fmt"
	"testing"
)

//lint:ignore U1000 非テスト：文字列に対してトークン列を返す
func TestLexerDispToken(t *testing.T) {
	input := `a
	
	b
	
	c`
	l := newLexerForTest(input)
	tokens := []Token{}
	for {
		tok := l.NextToken()
		tokens = append(tokens, tok)
		if tok.TokenType == EOF {
			break
		}
		fmt.Println(tok)
	}
	fmt.Println("after")
	for _, tok := range tokens {
		fmt.Println(tok)
	}
}
