package parser

import "fmt"

// Type は yacc で定義
type Token struct {
	Type    int
	Literal string
	Op      int
}

func (t Token) String() string {
	return fmt.Sprintf("Token{Type: %s, Literal: %q}", tokenType(t.Type), t.Literal)
}

func tokenType(t int) string {
	switch {
	case t == 0:
		return "'EOF'"
	case t == '\n':
		return "'EOL'"
	case t < 0x20 || t == 0x7f:
		return fmt.Sprintf("'\\x%02x'", t)
	case t >= 0x20 && t < 0x100:
		return fmt.Sprintf("'%c'", t)
	default:
		return yySymNames[yyXLAT[t]]
		// return fmt.Sprintf("'TOKEN(%d)'", t)
	}
}
