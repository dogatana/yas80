package parser

import (
	"fmt"
	"yas80/filecontent"
)

// TokenType, TokenType は grammer.y の %token で定義される
type TokenType int
type TokenSubType int

type Token struct {
	TokenType    TokenType
	TokenSubType TokenSubType
	Literal      string
	Context      *filecontent.Context
}

func (t Token) String() string {
	tt := int(t.TokenType)
	tst := int(t.TokenSubType)
	tstName := ""
	if tst != 0 {
		tstName = fmt.Sprintf(", Sub: %s(%d)", TokenLiteral(tst), tst)
	}

	var ctx string
	if t.Context == nil {
		ctx = "<nil>"
	} else {
		ctx = t.Context.String()
	}
	return fmt.Sprintf("Token{%s(%d)%s, %q, %s}",
		TokenLiteral(tt), tt, tstName, t.Literal, ctx)
}

var reservedWords map[string]Token = map[string]Token{
	// 単一行構文
	"EQU":      {TokenType: EQU, Literal: "EQU"},
	"CONST":    {TokenType: CONST, Literal: "CONST"},
	"VAR":      {TokenType: VAR, Literal: "VAR"},
	"FUNCTION": {TokenType: FUNCTION, Literal: "FUNCTION"},
	"ORG":      {TokenType: ORG, Literal: "ORG"},
	"INCLUDE":  {TokenType: INCLUDE, Literal: "INCLUDE"},
	"CHARMAP":  {TokenType: CHARMAP, Literal: "CHARMAP"},

	// データ定義
	"DB":   {TokenType: DATA, TokenSubType: DB, Literal: "DB"},
	"DEFB": {TokenType: DATA, TokenSubType: DB, Literal: "DEFB"},
	"DW":   {TokenType: DATA, TokenSubType: DW, Literal: "DW"},
	"DEFW": {TokenType: DATA, TokenSubType: DW, Literal: "DEFW"},
	"DD":   {TokenType: DATA, TokenSubType: DD, Literal: "DD"},
	"DS":   {TokenType: DS, TokenSubType: DSB, Literal: "DS"},
	"DSB":  {TokenType: DS, TokenSubType: DSB, Literal: "DSB"},
	"DSW":  {TokenType: DS, TokenSubType: DSW, Literal: "DSW"},

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

	// "BLOCK": {TokenType: BLOCK, Literal: "BLOCK"}, // 予約
	// "ENDB":  {TokenType: ENDB, Literal: "ENDB"},   // 予約

	"ENUM": {TokenType: ENUM, Literal: "ENUM"},
	"ENDE": {TokenType: ENDE, Literal: "ENDE"},

	// "FOR":    {TokenType: FOR, Literal: "FOR"},       // 予約
	// "ENDFOR": {TokenType: ENDFOR, Literal: "ENDFOR"}, // 予約

	"END": {TokenType: END, Literal: "END"},
}
