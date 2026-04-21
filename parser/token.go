package parser

import (
	"fmt"

	"github.com/dogatana/yas80/filecontent"
	"github.com/dogatana/yas80/intern"
)

// TokenType, TokenType は grammer.y の %token で定義される
type TokenType int
type TokenSubType int

type Token struct {
	TokenType    TokenType
	TokenSubType TokenSubType
	SymbolID     intern.SymbolID
	Context      filecontent.Context
}

func (t Token) String() string {
	tt := int(t.TokenType)
	tst := int(t.TokenSubType)
	tstName := ""
	if tst != 0 {
		tstName = fmt.Sprintf(", Sub: %s(%d)", TokenLiteral(tst), tst)
	}

	ctx := t.Context.String()
	return fmt.Sprintf("Token{%s(%d)%s, %d, %q, %s}",
		TokenLiteral(tt), tt, tstName, t.SymbolID, t.SymbolID, ctx)
}

// parser からErrorで呼ばれる際、Token を文字列化するのに使用する
func (t Token) Error() string {
	if t.TokenSubType != 0 {
		return TokenLiteral(int(t.TokenSubType))
	}
	return TokenLiteral(int(t.TokenType))
}

// 予約語テーブルを初期化
func init() {
	reservedWords = make(map[intern.SymbolID]Token, len(_reservedWords))
	for s, tt := range _reservedWords {
		id := intern.Intern(s)
		reservedWords[id] = Token{TokenType: tt.Type, TokenSubType: tt.SubType, SymbolID: id}
	}
	_reservedWords = nil
}

// 予約語テーブル
var reservedWords map[intern.SymbolID]Token

// 予約語テーブル初期化用
var _reservedWords = map[string]struct {
	Type    TokenType
	SubType TokenSubType
}{
	// 単一行構文
	"EQU":      {Type: EQU},
	"CONST":    {Type: CONST},
	"VAR":      {Type: VAR},
	"FUNCTION": {Type: FUNCTION},
	"ORG":      {Type: ORG},
	"INCLUDE":  {Type: INCLUDE},
	"CHARMAP":  {Type: CHARMAP},

	// データ定義
	"DB":   {Type: DATA, SubType: DB},
	"DEFB": {Type: DATA, SubType: DB},
	"DW":   {Type: DATA, SubType: DW},
	"DEFW": {Type: DATA, SubType: DW},
	"DD":   {Type: DATA, SubType: DD},
	"DS":   {Type: DS, SubType: DSB},
	"DSB":  {Type: DS, SubType: DSB},
	"DSW":  {Type: DS, SubType: DSW},

	// 複数行構文
	"IF":    {Type: IF},
	"ELSE":  {Type: ELSE},
	"ELIF":  {Type: ELIF},
	"ENDIF": {Type: ENDIF},

	"MACRO": {Type: MACRO},
	"ENDM":  {Type: ENDM},
	"EXITM": {Type: EXITM},

	"REPT": {Type: REPT},
	"ENDR": {Type: ENDR},

	"PROC": {Type: PROC},
	"ENDP": {Type: ENDP},

	"FUNC":   {Type: FUNC},
	"ENDF":   {Type: ENDF},
	"RETURN": {Type: RETURN},

	// "BLOCK": {BLOCK},
	// "ENDB":  {ENDB},

	"ENUM": {Type: ENUM},
	"ENDE": {Type: ENDE},

	// "FOR":    {FOR},
	// "ENDFOR": {ENDFOR},

	"END": {Type: END},
}
