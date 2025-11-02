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

var reservedWords map[string]Token = map[string]Token{
	"EQU":          {TokenType: EQU, Literal: "EQU"},
	"CONST":        {TokenType: CONST, Literal: "CONST"},
	"VAR":          {TokenType: VAR, Literal: "VAR"},
	"FUNC":         {TokenType: FUNC, Literal: "FUNC"},
	"IF":           {TokenType: IF, Literal: "IF"},
	"ELSE":         {TokenType: ELSE, Literal: "ELSE"},
	"ELIF":         {TokenType: ELIF, Literal: "ELIF"},
	"ENDIF":        {TokenType: END_IF, Literal: "ENDIF"},
	"END_IF":       {TokenType: END_IF, Literal: "END_IF"},
	"MACRO":        {TokenType: MACRO, Literal: "MACRO"},
	"ENDMACRO":     {TokenType: END_MACRO, Literal: "ENDMACRO"},
	"END_MACRO":    {TokenType: END_MACRO, Literal: "END_MACRO"},
	"ENDM":         {TokenType: END_MACRO, Literal: "ENDM"},
	"REPEAT":       {TokenType: REPEAT, Literal: "REPEAT"},
	"ENDREPEAT":    {TokenType: END_REPEAT, Literal: "ENDREPEAT"},
	"END_REPEAT":   {TokenType: END_REPEAT, Literal: "END_REPEAT"},
	"ENDR":         {TokenType: END_REPEAT, Literal: "ENDR"},
	"PROC":         {TokenType: PROC, Literal: "PROC"},
	"ENDPROC":      {TokenType: END_PROC, Literal: "ENDPROC"},
	"END_PROC":     {TokenType: END_PROC, Literal: "END_PROC"},
	"ENDP":         {TokenType: END_PROC, Literal: "ENDP"},
	"FUNCTION":     {TokenType: FUNCTION, Literal: "FUNCTION"},
	"ENDFUNCTION":  {TokenType: END_FUNCTION, Literal: "ENDFUNCTION"},
	"END_FUNCTION": {TokenType: END_FUNCTION, Literal: "END_FUNCTION"},
	"ENDF":         {TokenType: END_FUNCTION, Literal: "ENDF"},
	"BLOCK":        {TokenType: BLOCK, Literal: "BLOCK"},
	"ENDBLOCK":     {TokenType: END_BLOCK, Literal: "ENDBLOCK"},
	"END_BLOCK":    {TokenType: END_BLOCK, Literal: "END_BLOCK"},
	"ENDB":         {TokenType: END_BLOCK, Literal: "ENDB"},
	"ENUM":         {TokenType: ENUM, Literal: "ENDBLOCK"},
	"END_ENUM":     {TokenType: END_ENUM, Literal: "END_ENUM"},
	"ENDE":         {TokenType: END_ENUM, Literal: "ENDE"},
}
