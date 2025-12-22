package parser

import (
	"fmt"
	"strings"
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

// 各種記号のテスト
func TestLexSymbols(t *testing.T) {
	input := " ( ) = - + | ^ * / % & ! ~ << >> < <= == != >= > || && ## "
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
		{MULDIV, '%', "%"},
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

		{CONCAT, 0, "##"},
	}

	l := newLexerForTest(input)

	for _, e := range expected {
		tok := l.NextToken()
		if tok.Context.Line == 0 {
			t.Errorf("LineNumber not set. got %s", tok.String())
		}
		if tok.TokenType != e.TokenType {
			t.Errorf("expected Token.TokenType %s. got %s", TokenLiteral(int(e.TokenType)), tok.String())
		}
		if e.SubType != 0 && tok.TokenSubType != e.SubType {
			t.Errorf("expected Token.SubType %s. got %s", TokenLiteral(int(e.SubType)), tok.String())
		}
		if tok.Literal != e.Literal {
			t.Errorf("expected Token.Literal %q. got %s", e.Literal, tok.String())
		}
	}
}

// 空白入力、コメントのみの入力のテスト(最後に EOL, EOFが返ること)
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

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		for _, expected := range tt.expected_tokens {
			tok := l.NextToken()
			if tt.input != "" && tok.Context.Line == 0 {
				fmt.Println("tokenize", tt.input)
				t.Errorf("[%d] LineNumber not set. got %s", tn, tok.String())
			}
			if tok.TokenType != TokenType(expected) {
				t.Fatalf("[%d] expected=%d, got=%s", tn, expected, tok.String())
			}
		}
	}
}

// 不正な文字のテスト
func TestLexInvalidCharacter(t *testing.T) {
	input := " あ "

	l := newLexerForTest(input)
	tok := l.NextToken()
	if tok.Context.Line == 0 {
		t.Errorf("LineNumber not set. got %s", tok.String())
	}
	if tok.TokenType != INVALID || tok.Literal != "あ" {
		t.Errorf("expected=INVALID literal 'あ', got=%#v", tok)
	}
}

// 文字列リテラルのテスト
func TestLexString(t *testing.T) {
	tests := []struct {
		input            string
		expected_literal string
	}{
		{`"abc"`, "abc"},
		{` "abc def" `, "abc def"},
		{`""`, ""},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.Context.Line == 0 {
			t.Errorf("[%d] LineNumber not set. got %s", tn, tok.String())
		}
		if tok.TokenType != STRING || tok.Literal != tt.expected_literal {
			t.Errorf("[%d] expected=STRING with literal %q, got=%#v", tn, tt.expected_literal, tok)
		}
	}
}

// 数値リテラルのテスト
func TestLexNumber(t *testing.T) {
	tests := []struct {
		input            string
		expected_literal string
	}{
		// 0-
		{"0", "0"},
		{"  0", "0"},
		{"0  ", "0"},
		{"  0  ", "0"},
		{"12345", "12345"},
		// 5-
		{"12_34_5 ", "12_34_5"},

		{"0b111_000", "0b111_000"},
		{"0B_11_000_111", "0B_11_000_111"},
		{"%1111_0000", "%1111_0000"},

		{"0o_777", "0o_777"},
		// 10-
		{"0O_123", "0O_123"},

		{"0x19af", "0x19af"},
		{"0XABCD", "0XABCD"},
		{"$12_34_56_78", "$12_34_56_78"},
		{"1234_5467h", "1234_5467h"},
		// 15-
		{"89ab_cdefH", "89ab_cdefH"},

		{"0abc", "0abc"}, // 0-9 で始まる場合は16進数まで含めて受け付ける
		{"0xhoge", "0x"},
		{"0ohoge", "0o"},
		{"$0ggg", "$0"},
		// 20-
		{"$9ggg", "$9"},
		{"$aggg", "$a"},
		{"$fggg", "$f"},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.Context.Line == 0 {
			t.Errorf("[%d] LineNumber not set. got %s", tn, tok.String())
		}
		if tok.TokenType != NUMBER || tok.Literal != tt.expected_literal {
			t.Errorf("[%d] expected=NUMBER(%q), got=%s", tn, tt.expected_literal, tok.String())
		}
	}
}

// Lexインターフェースの基本動作確認テスト
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
		if ret != expected || lval.token.TokenType != TokenType(expected) {
			t.Errorf("expected=%d, got=%d", expected, lval.token.TokenType)
		}
	}
}

// 識別子のテスト
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

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.Context.Line == 0 {
			t.Errorf("[%d] LineNumber not set. got %s", tn, tok.String())
		}
		if tok.TokenType != tt.expectedType {
			t.Errorf("[%d] expected Type %s. got %s", tn, TokenLiteral(int(tt.expectedType)), tok.String())
		}
		if tok.Literal != tt.expectedLiteral {
			t.Errorf("[%d] expected Literal %q. got %s", tn, tt.expectedLiteral, tok.String())
		}
	}
}

// Z80 8ビットレジスタのテスト
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

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.Context.Line == 0 {
			t.Errorf("[%d] LineNumber not set. got %s", tn, tok.String())
		}
		if tok.TokenType != Z80_REG8 {
			t.Errorf("[%d] expected Type Z80_REG8. got %s", tn, tok.String())
		}
		if tok.TokenSubType != tt.expected {
			t.Errorf("[%d] expected TokenSubtype %q. got %s", tn, TokenLiteral(int(tt.expected)), tok.String())
		}
		if tok.Literal != tt.input {
			t.Errorf("[%d] expected Literal %q. got %s", tn, TokenLiteral(int(tt.expected)), tok.String())
		}
	}
}

// Z80 16ビットレジスタ、レジスタペアのテスト
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

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.Context.Line == 0 {
			t.Errorf("[%d] LineNumber not set. got %s", tn, tok.String())
		}
		if tok.TokenType != Z80_REG16 {
			t.Errorf("[%d] expected Type Z80_REG16. got %s", tn, tok.String())
		}
		if tok.TokenSubType != TokenSubType(tt.expected) {
			t.Errorf("[%d] expected TokenSubType %q. got %s", tn, TokenLiteral(int(tt.expected)), tok.String())
		}
		if tok.Literal != TokenLiteral(int(tt.expected)) {
			t.Errorf("[%d] expected Literal %q. got %s", tn, TokenLiteral(int(tt.expected)), tok.String())
		}
	}
}

// Z80フラグのテスト
func TestLexZ80FLAG(t *testing.T) {
	tests := []struct {
		input    string
		expected TokenSubType
	}{
		{"CY", Z80_FLAG_C}, // CY はキャリフラグの別名
		{"NC", Z80_FLAG_NC},
		{"Z", Z80_FLAG_Z},
		{"NZ", Z80_FLAG_NZ},
		{"PO", Z80_FLAG_PO},
		{"PE", Z80_FLAG_PE},
		{"p", Z80_FLAG_P},
		{"M", Z80_FLAG_M},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.Context.Line == 0 {
			t.Errorf("[%d] LineNumber not set. got %s", tn, tok.String())
		}
		if tok.TokenType != Z80_FLAG {
			t.Errorf("[%d] expected Type Z80_FLAG. got %#v", tn, TokenLiteral(int(tok.TokenType)))
		}
		if tok.TokenSubType != TokenSubType(tt.expected) {
			t.Errorf("[%d] expected TokenSubType %q. got %s", tn, TokenLiteral(int(tt.expected)), tok.String())
		}
		if tok.Literal != TokenLiteral(int(tt.expected)) {
			t.Errorf("[%d] expected Literal %q. got %s", tn, TokenLiteral(int(tt.expected)), tok.String())
		}
	}
}

// Z80命令のテスト
func TestLexZ80Instructions(t *testing.T) {
	input := strings.ToLower(readTestDataFile(t, "z80instruction.txt"))

	l := newLexerForTest(input)
	for {
		tok := l.NextToken()
		if tok.TokenType == EOL {
			continue
		}
		if tok.TokenType == EOF {
			break
		}

		if tok.Context.Line == 0 {
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

// 予約語のテスト
func TestLexReservedWords(t *testing.T) {
	input := "const var equ function org " +
		"if else elif endif " +
		"macro endm exitm " +
		"rept endr " +
		"func endf return " +
		"proc endp " +
		"block endb " +
		"for endfor "
	l := newLexerForTest(input)

	for {
		tok := l.NextToken()
		if tok.Context.Line == 0 {
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

// $で始まるトークンのテスト
func TestLexDoller(t *testing.T) {
	tests := []struct {
		input    string
		expected []TokenType
	}{
		// 0-
		{"$", []TokenType{IDENT, EOL}},
		{"$ ", []TokenType{IDENT, EOL}},
		{"$0", []TokenType{NUMBER, EOL}},
		{"$a", []TokenType{NUMBER, EOL}},
		{"$f", []TokenType{NUMBER, EOL}},
		// 5-
		{"$A", []TokenType{NUMBER, EOL}},
		{"$F", []TokenType{NUMBER, EOL}},
		{"$_", []TokenType{IDENT, EOL}},
		{"$g", []TokenType{IDENT, EOL}},
		// {"$$", []TokenType{IDENT, EOL}}, // $$ は2つの $ になる
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		for _, tokenType := range tt.expected {
			tok := l.NextToken()
			if tok.Context.Line == 0 {
				t.Errorf("[%d] LineNumber not set. got %s", tn, tok.String())
			}
			if tok.TokenType != tokenType {
				t.Errorf("[%d] expected TokenType %s. got %#v", tn, TokenLiteral(int(tokenType)), tok)
			}
			if tok.TokenType == EOL {
				break
			}
		}
	}
}

// 行継続のテスト
func TestLexLineContinuation(t *testing.T) {
	tests := []struct {
		input    string
		expected []TokenType
	}{
		{`ld \
		a \
		,\
		b`, []TokenType{Z80_INST2, Z80_REG8, ',', Z80_REG8, EOL}},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		for _, tokenType := range tt.expected {
			tok := l.NextToken()
			if tok.Context.Line == 0 {
				t.Errorf("[%d] LineNumber not set. got %s", tn, tok.String())
			}
			if tok.TokenType != tokenType {
				t.Errorf("[%d] expected TokenType %s. got %#v", tn, TokenLiteral(int(tokenType)), tok)
			}
			if tok.TokenType == EOL {
				break
			}
		}
	}
}

// 行番号のテスト
func TestLexerContextLineNumber(t *testing.T) {
	tests := []struct {
		input string
	}{
		{"abc"},               // Context.Line == 1
		{"abc \n def"},        // Context.Line == 1,2
		{"abc \n def \n xyz"}, // Context.Line == 1,2,3
		{"abc \\ def \\ xyz"}, // Context.Line == 1,1,1
	}
	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		ln := 1
		for {
			tok := l.NextToken()
			if tok.TokenType == EOL {
				continue
			} else if tok.TokenType == EOF {
				break
			}
			if tok.Context.Line != ln {
				t.Errorf("[%d] LineNumber not %d. got %d", tn, ln, tok.Context.Line)
			}
			if strings.Contains(tt.input, "\\") {
				continue // マルチステートメントは同一行
			}
			ln++
		}
	}
}
