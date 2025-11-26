package parser

import (
	"fmt"
	"testing"
)

//lint:ignore U1000 非テスト：文字列に対してトークン列を返す
func testDisplayTokens(t *testing.T) {
	input := `()= + - * / & | ^ == != < <= > >= ! ~ << >> || &&`
	l := newLexerForTest(input)
	for {
		tok := l.NextToken()
		fmt.Println(tok.String())
		if tok.TokenType == 0 {
			break
		}
	}
}

func TestLexSymbols(t *testing.T) {
	input := " ( ) = - + | ^ * / & ! ~ << >> < <= == != >= > || && "
	expected := []struct {
		TokenType TokenType
		SubType   TokenSubType
		Literal   string
	}{
		{'(', 0, "("},
		{')', 0, ")"},
		{'=', 0, "="},

		{ADDSUB, '-', "-"},
		{ADDSUB, '+', "+"},
		{ADDSUB, '|', "|"},
		{ADDSUB, '^', "^"},

		{MULDIV, '*', "*"},
		{MULDIV, '/', "/"},
		{MULDIV, '&', "&"},

		{UNARY, '!', "!"},
		{UNARY, '~', "~"},

		{SHIFT, SL, "<<"},
		{SHIFT, SR, ">>"},

		{COMP, '<', "<"},
		{COMP, LE, "<="},
		{COMP, EQ, "=="},
		{COMP, NEQ, "!="},
		{COMP, GE, ">="},
		{COMP, '>', ">"},

		{OR, 0, "||"},
		{AND, 0, "&&"},
	}

	l := newLexerForTest(input)

	for _, e := range expected {
		tok := l.NextToken()
		if tok.LineNumber == 0 {
			t.Errorf("LineNumber not set. got %s", tok.String())
		}
		if tok.TokenType != e.TokenType {
			t.Errorf("expected Token.TokenType %s. got %s", TokenLiteral(int(e.TokenType)), tok.String())
		}
		if e.SubType != 0 && tok.TokenSubType != e.SubType {
			t.Errorf("expected Token.SubType %s. got %s", TokenLiteral(int(tok.TokenSubType)), tok.String())
		}
		if tok.Literal != e.Literal {
			t.Errorf("expected Token.Literal %q. got %s", tok.Literal, tok.String())
		}
	}
}

func TestLexBlankInput(t *testing.T) {
	tests := []struct {
		input           string
		expected_tokens []int
	}{
		{"", []int{EOL, EOF}},
		{";", []int{EOL, EOF}},
		{" ;", []int{EOL, EOF}},
		{"; ", []int{EOL, EOF}},
		{" ; comment", []int{EOL, EOF}},
		{"  ", []int{EOL, EOF}},
		{" \n ", []int{EOL, EOL, EOF}},
		{" \n \n", []int{EOL, EOL, EOF}},
		{" ; comment \n ;  comment \n", []int{EOL, EOL, EOF}},
		{" \\ ", []int{EOL, EOL, EOF}},
		{" ; comment \\ ; comment ", []int{EOL, EOF}},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		for _, expected := range tt.expected_tokens {
			tok := l.NextToken()
			if tt.input != "" && tok.LineNumber == 0 {
				fmt.Println("tokenize", tt.input)
				t.Errorf("LineNumber not set. got %s", tok.String())
			}
			if tok.TokenType != TokenType(expected) {
				t.Fatalf("tokenize %q, expected=%d, got=%s", tt.input, expected, tok.String())
			}
		}
	}
}

func TestLexInvalidCharacter(t *testing.T) {
	input := " あ "

	l := newLexerForTest(input)
	tok := l.NextToken()
	if tok.LineNumber == 0 {
		t.Errorf("LineNumber not set. got %s", tok.String())
	}
	if tok.TokenType != INVALID || tok.Literal != "あ" {
		t.Errorf("expected=INVALID literal 'あ', got=%#v", tok)
	}
}

func TestLexString(t *testing.T) {
	tests := []struct {
		input            string
		expected_literal string
	}{
		{`"abc"`, "abc"},
		{` "abc def" `, "abc def"},
		{`""`, ""},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.LineNumber == 0 {
			t.Errorf("LineNumber not set. got %s", tok.String())
		}
		if tok.TokenType != STRING || tok.Literal != tt.expected_literal {
			fmt.Printf("tokenize %q\n", tt.input)
			t.Errorf("expected=STRING with literal %q, got=%#v", tt.expected_literal, tok)
		}
	}
}

func TestLexNumber(t *testing.T) {
	tests := []struct {
		input            string
		expected_literal string
	}{
		{"0", "0"},
		{"  0", "0"},
		{"0  ", "0"},
		{"  0  ", "0"},
		{"12345", "12345"},
		{"12_34_5 ", "12_34_5"},
		{"0x19af", "0x19af"},
		{"0XABCD", "0XABCD"},
		{"0o_777", "0o_777"},
		{"$12_34_56_78", "$12_34_56_78"},
		{"%1111_0000", "%1111_0000"},
		// 不正な数値でも受け付ける
		{"0abc", "0abc"},
		{"0xhoge", "0xhoge"},
		{"0ohoge", "0ohoge"},
		{"$0ggg", "$0ggg"},
		{"$9ggg", "$9ggg"},
		{"$aggg", "$aggg"},
		{"$fggg", "$fggg"},
		{"%xyz", "%xyz"},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.LineNumber == 0 {
			t.Errorf("LineNumber not set. got %s", tok.String())
		}
		if tok.TokenType != NUMBER || tok.Literal != tt.expected_literal {
			t.Errorf("expected=NUMBER(%q), got=%s", tt.expected_literal, tok.String())
		}
	}
}

func TestLexInterface(t *testing.T) {
	input := " 123 + 456 \n"

	l := newLexerForTest(input)

	expected_tokens := []int{
		NUMBER,
		ADDSUB,
		NUMBER,
		EOL,
	}

	var lval yySymType
	for _, expected := range expected_tokens {
		ret := l.Lex(&lval)
		// fmt.Printf("ret=%d, lval=%#v\n", ret, lval)
		if ret != expected || lval.token.TokenType != TokenType(expected) {
			t.Errorf("expected=%d, got=%d", expected, lval.token.TokenType)
		}
	}
}

func TestLexIdent(t *testing.T) {
	tests := []struct {
		input           string
		expectedType    TokenType
		expectedLiteral string
	}{
		{" abc     ", IDENT, "abc"},
		{" _abc    ", IDENT, "_abc"},
		{" ab_c    ", IDENT, "ab_c"},
		{" abc_    ", IDENT, "abc_"},
		{" abc.def:", DOT_IDENT, "abc.def"},
		{" .def:   ", LOCAL_IDENT, ".def"},
		{" .def    ", LOCAL_IDENT, ".def"},
		{" @def    ", AT_IDENT, "@def"},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.LineNumber == 0 {
			t.Errorf("LineNumber not set. got %s", tok.String())
		}
		if tok.TokenType != tt.expectedType {
			t.Errorf("expected Type %s. got %s", TokenLiteral(int(tt.expectedType)), tok.String())
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("expected Literal %q. got %s", tt.expectedLiteral, tok.String())
		}
	}
}
func TestLexZ80REG8(t *testing.T) {
	tests := []struct {
		input    string
		expected TokenSubType
	}{
		{"A", Z80_REG_A},
		{"B", Z80_REG_B},
		{"C", Z80_REG_C},
		{"D", Z80_REG_D},
		{"E", Z80_REG_E},
		{"H", Z80_REG_H},
		{"L", Z80_REG_L},
		{"IXH", Z80_REG_IXH},
		{"IXL", Z80_REG_IXL},
		{"IYH", Z80_REG_IYH},
		{"IYL", Z80_REG_IYL},
		{"I", Z80_REG_I},
		{"R", Z80_REG_R},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.LineNumber == 0 {
			t.Errorf("LineNumber not set. got %s", tok.String())
		}
		if tok.TokenType != Z80_REG8 {
			t.Errorf("expected Type Z80_REG8. got %s", tok.String())
		}
		if tok.TokenSubType != tt.expected {
			t.Errorf("expected TokenSubtype %q. got %s", TokenLiteral(int(tt.expected)), tok.String())
		}
		if tok.Literal != tt.input {
			t.Errorf("expected Literal %q. got %s", tt.input, tok.String())
		}
	}
}

func TestLexZ80REG16(t *testing.T) {
	tests := []struct {
		input    string
		expected TokenSubType
	}{
		{"SP", Z80_REG_SP},
		{"IX", Z80_REG_IX},
		{"IY", Z80_REG_IY},
		{"AF", Z80_REG_AF},
		{"AF'", Z80_REG_AFEX},
		{"BC", Z80_REG_BC},
		{"DE", Z80_REG_DE},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.LineNumber == 0 {
			t.Errorf("LineNumber not set. got %s", tok.String())
		}
		if tok.TokenType != Z80_REG16 {
			t.Errorf("expected Type Z80_REG16. got %s", tok.String())
		}
		if tok.TokenSubType != TokenSubType(tt.expected) {
			t.Errorf("expected TokenSubType %q. got %s", TokenLiteral(int(tt.expected)), tok.String())
		}
		if tok.Literal != TokenLiteral(int(tt.expected)) {
			t.Errorf("expected Literal %q. got %s", tt.input, tok.String())
		}
	}
}

func TestLexZ80FLAG(t *testing.T) {
	tests := []struct {
		input    string
		expected TokenSubType
	}{
		// C は Z80_REG8 トークンとなる
		// {"C", Z80_FLAG_C},
		{"NC", Z80_FLAG_NC},
		{"Z", Z80_FLAG_Z},
		{"NZ", Z80_FLAG_NZ},
		{"PO", Z80_FLAG_PO},
		{"PE", Z80_FLAG_PE},
		{"p", Z80_FLAG_P},
		{"M", Z80_FLAG_M},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.LineNumber == 0 {
			t.Errorf("LineNumber not set. got %s", tok.String())
		}
		if tok.TokenType != Z80_FLAG {
			t.Errorf("tokenize %q. expected Type Z80_FLAG. got %#v", tt.input, TokenLiteral(int(tok.TokenType)))
		}
		if tok.TokenSubType != TokenSubType(tt.expected) {
			t.Errorf("expected TokenSubType %q. got %s", TokenLiteral(int(tt.expected)), tok.String())
		}
		if tok.Literal != TokenLiteral(int(tt.expected)) {
			t.Errorf("expected Literal %q. got %s", TokenLiteral(int(tt.expected)), tok.String())
		}
	}
}

func TestLexZ80Instructions(t *testing.T) {
	input := readTestDataFile(t, "z80instruction.txt")

	l := newLexerForTest(input)
	for {
		tok := l.NextToken()
		if tok.TokenType == EOL {
			continue
		}
		if tok.TokenType == EOF {
			break
		}

		if tok.LineNumber == 0 {
			t.Errorf("LineNumber not set. got %s", tok.String())
		}
		expectedToken, ok := z80ReservedWords[tok.Literal]
		if !ok {
			t.Errorf("instruction %q not found", tok.Literal)
			continue
		}
		if tok.TokenType != expectedToken.TokenType {
			t.Errorf("expected Type %s. got %s", TokenLiteral(int(tok.TokenType)), tok.String())
		}
		if tok.Literal != expectedToken.Literal {
			t.Errorf("expected Literal %q. got %s", expectedToken.Literal, tok.String())
		}
		if tok.TokenSubType != expectedToken.TokenSubType {
			t.Errorf("expected Op '%d'. got %s", expectedToken.TokenSubType, tok.String())
		}

	}
}

func TestLexReservedWords(t *testing.T) {
	input := "CONST VAR EQU FUNCTION ORG " +
		"IF ELSE ELIF ENDIF " +
		"MACRO ENDM EXITM " +
		"REPT ENDR " +
		"FUNC ENDF RETURN " +
		"PROC ENDP " +
		"BLOCK ENDB " +
		"FOR ENDF "
	l := newLexerForTest(input)

	for {
		tok := l.NextToken()
		if tok.LineNumber == 0 {
			t.Errorf("LineNumber not set. got %s", tok.String())
		}
		if tok.TokenType == EOL {
			break
		}
		expected, ok := reservedWords[tok.Literal]
		if !ok {
			t.Fatalf("word %q is not registered", tok.Literal)
		}
		if tok.TokenType != expected.TokenType {
			t.Errorf("expected Type %s. got %s", TokenLiteral(int(tok.TokenType)), tok.String())
		}
		if tok.Literal != expected.Literal {
			t.Errorf("expected Literal %q. got %s", expected.Literal, tok.String())
		}
		if tok.TokenSubType != expected.TokenSubType {
			t.Errorf("expected Op '%d'. got %s", expected.TokenSubType, tok.String())
		}
	}
}

func TestLexDoller(t *testing.T) {
	tests := []struct {
		input    string
		expected []TokenType
	}{
		// C は Z80_REG8 トークンとなる
		// {"C", Z80_FLAG_C},
		{"$", []TokenType{IDENT, EOL}},
		{"$ ", []TokenType{IDENT, EOL}},
		{"$0", []TokenType{NUMBER, EOL}},
		{"$a", []TokenType{NUMBER, EOL}},
		{"$f", []TokenType{NUMBER, EOL}},
		{"$A", []TokenType{NUMBER, EOL}},
		{"$F", []TokenType{NUMBER, EOL}},
		{"$_", []TokenType{IDENT, IDENT, EOL}},
		{"$g", []TokenType{IDENT, IDENT, EOL}},
		{"$$", []TokenType{IDENT, IDENT, EOL}},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		for _, tokenType := range tt.expected {
			tok := l.NextToken()
			if tok.LineNumber == 0 {
				fmt.Printf("input %q\n", tt.input)
				t.Errorf("LineNumber not set. got %s", tok.String())
			}
			if tok.TokenType != tokenType {
				fmt.Printf("input %q\n", tt.input)
				t.Errorf("expected TokenType %s. got %#v", TokenLiteral(int(tokenType)), tok)
			}
			if tok.TokenType == EOL {
				break
			}
		}
	}
}
