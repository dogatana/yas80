package parser

import (
	"fmt"
	"yas80/fileblock"
)

// TokenType, TokenType は grammer.y の %token で定義される
type TokenType int
type TokenSubType int

type Token struct {
	TokenType    TokenType
	TokenSubType TokenSubType
	Literal      string
	TokenContext TokenContext
}

// ファイル内の位置
type TokenContext struct {
	FileBlock  *fileblock.FileBlock
	LineNumber int
	Index      int
}

func (t Token) String() string {
	tt := int(t.TokenType)
	tst := int(t.TokenSubType)
	tstName := TokenLiteral(tst)
	if tst == 0 {
		tstName = "0"
	}
	return fmt.Sprintf("Token{TokenType: %s(%d), SubType: %s(%d), Literal: %q, LineNumber: %d}",
		TokenLiteral(tt), tt, tstName, tst, t.Literal, t.TokenContext.LineNumber)
}

var reservedWords map[string]Token = map[string]Token{
	// 単一行構文
	"EQU":      {TokenType: EQU, Literal: "EQU"},
	"CONST":    {TokenType: CONST, Literal: "CONST"},
	"VAR":      {TokenType: VAR, Literal: "VAR"},
	"FUNCTION": {TokenType: FUNCTION, Literal: "FUNCTION"},
	"ORG":      {TokenType: ORG, Literal: "ORG"},

	// 複数行構文
	"IF":    {TokenType: IF, Literal: "IF"},
	"ELSE":  {TokenType: ELSE, Literal: "ELSE"},
	"ELIF":  {TokenType: ELIF, Literal: "ELIF"},
	"ENDIF": {TokenType: ENDIF, Literal: "ENDIF"},

	"MACRO": {TokenType: MACRO, Literal: "MACRO"},
	"ENDM":  {TokenType: ENDM, Literal: "ENDM"},
	"EXITM": {TokenType: EXITM, Literal: "EXITM"},

	"REPT": {TokenType: REPT, Literal: "REPT"},
	"ENDR": {TokenType: ENDR, Literal: "ENDR"},

	"PROC": {TokenType: PROC, Literal: "PROC"},
	"ENDP": {TokenType: ENDP, Literal: "ENDP"},

	"FUNC":   {TokenType: FUNC, Literal: "FUNC"},
	"ENDF":   {TokenType: ENDF, Literal: "ENDF"},
	"RETURN": {TokenType: RETURN, Literal: "RETURN"},

	"BLOCK": {TokenType: BLOCK, Literal: "BLOCK"},
	"ENDB":  {TokenType: ENDB, Literal: "ENDB"},

	"ENUM": {TokenType: ENUM, Literal: "ENDBLOCK"},
	"ENDE": {TokenType: ENDE, Literal: "ENDE"},

	"FOR": {TokenType: FOR, Literal: "FOR"},
	// "ENDFOR": {TokenType: ENDFOR, Literal: "ENDFOR"}, // ENDF とする
}
