package parser

import "fmt"

// Type は yacc で定義
type Token struct {
	Type       int
	SubType    int
	Literal    string
	LineNumber int
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %s, SubType: %s, Literal: %s, LineNumber: %d}",
		yySymNames[yyXLAT[t.Type]], tokenLiteral(t.SubType), t.Literal, t.LineNumber)
}
