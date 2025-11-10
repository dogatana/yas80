package parser

import (
	"fmt"
	"testing"
)

// 非テスト：文字列に対してトークン列を返す
func TestDisplayTokens(t *testing.T) {
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

func TestSymbols(t *testing.T) {
	input := " ( ) = - + | ^ * / & ! ~ << >> < <= == != >= > || && "
	expected := []struct {
		TokenType TokenType
		SubType   TokenSubType
		Literal   string
	}{
		{'(', 0, "("},
		{')', 0, ")"},
		{'=', 0, "="},
		{'-', 0, "-"},

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
		if tok.TokenType != e.TokenType {
			t.Errorf("expected Token.TokenType %s. got %s", tokenLiteral(int(e.TokenType)), tok.String())
		}
		if e.SubType != 0 && tok.TokenSubType != e.SubType {
			t.Errorf("expected Token.SubType %s. got %s", tokenLiteral(int(tok.TokenSubType)), tok.String())
		}
		if tok.Literal != e.Literal {
			t.Errorf("expected Token.Literal %s. got %s", tok.Literal, tok.String())
		}
	}
}

func TestBlankInput(t *testing.T) {
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
			if tok.TokenType != TokenType(expected) {
				t.Fatalf("tokenize %q, expected=%d, got=%s", tt.input, expected, tok.String())
			}
		}
	}
}

func TestInvalidCharacter(t *testing.T) {
	input := " あ "

	l := newLexerForTest(input)
	tok := l.NextToken()
	if tok.TokenType != INVALID || tok.Literal != "あ" {
		t.Errorf("expected=INVALID literal 'あ', got=%#v", tok)
	}
}

func TestString(t *testing.T) {
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
		if tok.TokenType != STRING || tok.Literal != tt.expected_literal {
			fmt.Printf("tokenize %q\n", tt.input)
			t.Errorf("expected=STRING with literal %q, got=%#v", tt.expected_literal, tok)
		}
	}
}

func TestNumber(t *testing.T) {
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
		{"$ggg", "$ggg"},
		{"%xyz", "%xyz"},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
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

func TestIDENT(t *testing.T) {
	tests := []struct {
		input             string
		expected_literals []string
	}{
		{" abc _abc ab_c abc_ ", []string{"abc", "_abc", "ab_c", "abc_"}},
		// testLABEL へ移動
		// {".abc @abc abc.def abc@def", []string{".abc", "@abc", "abc.def", "abc@def"}},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		for _, expected := range tt.expected_literals {
			tok := l.NextToken()
			if tok.TokenType != IDENT {
				t.Errorf("expected=IDENT, got=%#v", tok)
			}
			if tok.Literal != expected {
				t.Errorf("expected=%q, got=%q", expected, tok.Literal)
			}
		}
	}
}

func TestLabel(t *testing.T) {
	tests := []struct {
		input    string
		expected TokenType
	}{
		{" abc.def:", DOT_IDENT},
		{" .def:", LOCAL_IDENT},
		{" @def:", AT_IDENT},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.TokenType != tt.expected {
			t.Errorf("expected Type %s. got %#v", tokenLiteral(int(tt.expected)), tok)
		}
	}
}
func TestZ80REG8(t *testing.T) {
	tests := []struct {
		input string
		op    int
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
		if tok.TokenType != Z80_REG8 {
			t.Errorf("tokenize %q. expected Type Z80_REG8. got %#v", tt.input, tokenLiteral(int(tok.TokenType)))
		}
		if tok.Literal != tt.input {
			t.Errorf("tokenize %q. expected Literal %q. got %#v", tt.input, tt.input, tok)
		}
		if tok.TokenSubType != TokenSubType(tt.op) {
			t.Errorf("tokenize %q. expected Op '%c'. got '%c'", tt.input, tt.op, tok.TokenSubType)
		}
	}
}

func TestZ80REG16(t *testing.T) {
	tests := []struct {
		input string
		op    int
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
		if tok.TokenType != Z80_REG16 {
			t.Errorf("tokenize %q. expected Type Z80_REG16. got %#v", tt.input, tokenLiteral(int(tok.TokenType)))
		}
		if tok.Literal != tt.input {
			t.Errorf("tokenize %q. expected Literal %q. got %#v", tt.input, tt.input, tok)
		}
		if tok.TokenSubType != TokenSubType(tt.op) {
			t.Errorf("tokenize %q. expected Op '%c'. got '%c'", tt.input, tt.op, tok.TokenSubType)
		}
	}
}

func TestZ80FLAG(t *testing.T) {
	tests := []struct {
		input string
		op    int
	}{
		// C は Z80_REG8 トークンとなる
		{"NC", Z80_FLAG_NC},
		{"Z", Z80_FLAG_Z},
		{"NZ", Z80_FLAG_NZ},
		{"PO", Z80_FLAG_PO},
		{"PE", Z80_FLAG_PE},
		{"P", Z80_FLAG_P},
		{"M", Z80_FLAG_M},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		tok := l.NextToken()
		if tok.TokenType != Z80_FLAG {
			t.Errorf("tokenize %q. expected Type Z80_FLAG. got %#v", tt.input, tokenLiteral(int(tok.TokenType)))
		}
		if tok.Literal != tt.input {
			t.Errorf("tokenize %q. expected Literal %q. got %#v", tt.input, tt.input, tok)
		}
		if tok.TokenSubType != TokenSubType(tt.op) {
			t.Errorf("tokenize %q. expected Op '%c'. got '%c'", tt.input, tt.op, tok.TokenSubType)
		}
	}
}

func TestZ80Instructions(t *testing.T) {
	input := "LD PUSH POP EX EXX LDI LDIR LDDR CPI CPIR CPDR ADD ADC SUB SBC AND OR " +
		"CP INC DEC DAA CPL NEG CCF SCF NOP HALT DI EI IM " +
		"RLCA RLA RRCA RRA RLC RL RR SLA SRA SRL RLD RRD " +
		"BIT SET RES JP JR DJNZ CALL RET RETI RETN RST " +
		"IN INI INIR INDR OUT OUTI OTIR OUTD OTDR"

	l := newLexerForTest(input)
	for {
		tok := l.NextToken()
		if tok.TokenType == EOL {
			break
		}

		expectedToken, ok := z80ReservedWords[tok.Literal]
		if !ok {
			t.Errorf("instruction %q not found", tok.Literal)
			continue
		}
		if tok.TokenType != expectedToken.TokenType {
			t.Errorf("expected Type %s. got %#v", tokenLiteral(int(tok.TokenType)), tok)
		}
		if tok.Literal != expectedToken.Literal {
			t.Errorf("expected Literal %q. got %#v", expectedToken.Literal, tok)
		}
		if tok.TokenSubType != expectedToken.TokenSubType {
			t.Errorf("expected Op '%d'. got %#v", expectedToken.TokenSubType, tok)
		}

	}
}

func TestReservedWords(t *testing.T) {
	input := "CONST VAR EQU FN ORG " +
		"IF ELSE ELIF ENDIF " +
		"MACRO ENDM " +
		"REPEAT ENDR " +
		"FUNC ENDF " +
		"PROC ENDP " +
		"BLOCK ENDB " +
		"FOR ENDF "
	l := newLexerForTest(input)

	for {
		tok := l.NextToken()
		if tok.TokenType == EOL {
			break
		}
		expected, ok := reservedWords[tok.Literal]
		if !ok {
			t.Fatalf("word %q is not registered", tok.Literal)
		}
		if tok.TokenType != expected.TokenType {
			t.Errorf("expected Type %s. got %#v", tokenLiteral(int(tok.TokenType)), tok)
		}
		if tok.Literal != expected.Literal {
			t.Errorf("expected Literal %q. got %#v", expected.Literal, tok)
		}
		if tok.TokenSubType != expected.TokenSubType {
			t.Errorf("expected Op '%d'. got %#v", expected.TokenSubType, tok)
		}
	}
}
