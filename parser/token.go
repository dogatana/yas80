package parser

import "fmt"

type TokenType int
type TokenSubType int

// Type は yacc で定義
type Token struct {
	TokenType    TokenType
	TokenSubType TokenSubType
	Literal      string
	LineNumber   int
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %s, SubType: %s, Literal: %s, LineNumber: %d}",
		yySymNames[yyXLAT[int(t.TokenType)]], tokenLiteral(int(t.TokenSubType)), t.Literal, t.LineNumber)
}
