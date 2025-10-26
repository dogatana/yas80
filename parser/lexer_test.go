package parser

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func testNextChars(t *testing.T) {
	tests := []string{
		"",
		" ",
		"\t",
		"\n",
		" + - * / ( ) ",
	}
	for _, input := range tests {
		l := NewLexer(bufio.NewReader(strings.NewReader(input)))
		for i := 0; i < 20; i++ {
			fmt.Printf("tokenize=%q, curChar=%q, index=%d\n", input, l.curChar, l.index)
			if l.curChar == EOF {
				break
			}
			l.nextChar()
		}
	}
}

func testNextTokens(t *testing.T) {
	input := " + - * / ( ) "
	l := NewLexer(bufio.NewReader(strings.NewReader(input)))

	for i := 0; i < 20; i++ {
		tok := l.NextToken()
		fmt.Println(tok.String())
	}
}

func TestLexerOneCharacter(t *testing.T) {
	input := " + - * / ( ) "
	expected_tokens := []int{
		'+',
		'-',
		'*',
		'/',
		'(',
		')',
		EOL,
		EOF,
	}

	l := NewLexer(bufio.NewReader(strings.NewReader(input)))

	for _, expected := range expected_tokens {
		tok := l.NextToken()
		if tok.Type == EOF {
			break
		}
		if tok.Type != expected {
			t.Errorf("expected=%d, got=%#v", expected, tok)
		}
	}
}

func TestBlankInput(t *testing.T) {
	tests := []struct {
		input           string
		expected_tokens []int
	}{
		{"", []int{EOL, EOF}},
		{"  ", []int{EOL, EOF}},
		{" \n ", []int{EOL, EOL, EOF}},
		{" \n \n", []int{EOL, EOL, EOF}},
	}

	for _, tt := range tests {
		l := NewLexer(bufio.NewReader(strings.NewReader(tt.input)))
		for _, expected := range tt.expected_tokens {
			tok := l.NextToken()
			if tok.Type != expected {
				t.Fatalf("tokenize %q, expected=%d, got=%#v", tt.input, expected, tok)
			}
		}
	}
}

func TestInvalidCharacter(t *testing.T) {
	input := " あ "

	l := NewLexer(bufio.NewReader(strings.NewReader(input)))
	tok := l.NextToken()
	if tok.Type != INVALID || tok.Literal != "あ" {
		t.Errorf("expected=INVALID literal 'あ', got=%#v", tok)
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
	}

	for _, tt := range tests {
		l := NewLexer(bufio.NewReader(strings.NewReader(tt.input)))
		tok := l.NextToken()
		if tok.Type != NUMBER || tok.Literal != tt.expected_literal {
			t.Errorf("ze %q, expected=NUMBER with literal %q, got=%#v",
				tt.input, tt.expected_literal, tok)
		}
	}
}

func TestHexNumbers(t *testing.T) {

	input := "  0x19af  0XABCD $12_34_56_78"
	expected_literals := []string{"0x19af", "0XABCD", "$12_34_56_78"}

	l := NewLexer(bufio.NewReader(strings.NewReader(input)))
	for _, expected := range expected_literals {
		tok := l.NextToken()
		if tok.Type != NUMBER {
			t.Errorf("expected=NUMBER, got=%#v", tok)
		}
		if tok.Type != NUMBER || tok.Literal != expected {
			t.Errorf("expected=%q, got=%q", expected, tok.Literal)
		}
	}
}

func TestLexInterface(t *testing.T) {
	input := " 123 + 456 \n"

	l := NewLexer(bufio.NewReader(strings.NewReader(input)))

	expected_tokens := []int{
		NUMBER,
		'+',
		NUMBER,
		EOL,
	}

	var lval yySymType
	for _, expected := range expected_tokens {
		ret := l.Lex(&lval)
		// fmt.Printf("ret=%d, lval=%#v\n", ret, lval)
		if ret != expected || lval.token.Type != expected {
			t.Errorf("expected=%d, got=%d", expected, lval.token.Type)
		}
	}
}

func TestIDENT(t *testing.T) {
	tests := []struct {
		input             string
		expected_literals []string
	}{
		{" abc _abc ab_c abc_ ", []string{"abc", "_abc", "ab_c", "abc_"}},
		{".abc @abc abc.def abc@def", []string{".abc", "@abc", "abc.def", "abc@def"}},
	}

	for _, tt := range tests {
		l := NewLexer(bufio.NewReader(strings.NewReader(tt.input)))
		for _, expected := range tt.expected_literals {
			tok := l.NextToken()
			if tok.Type != IDENT {
				t.Errorf("expected=IDENT, got=%#v", tok)
			}
			if tok.Literal != expected {
				t.Errorf("expected=%q, got=%q", expected, tok.Literal)
			}
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
		l := NewLexer(bufio.NewReader(strings.NewReader(tt.input)))
		tok := l.NextToken()
		if tok.Type != Z80_REG8 {
			t.Errorf("tokenize %q. expected Type Z80_REG8. got %#v", tt.input, yySymNames[yyXLAT[tok.Type]])
		}
		if tok.Literal != tt.input {
			t.Errorf("tokenize %q. expected Literal %q. got %#v", tt.input, tt.input, tok)
		}
		if tok.Op != tt.op {
			t.Errorf("tokenize %q. expected Op '%c'. got '%c'", tt.input, tt.op, tok.Op)
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
		l := NewLexer(bufio.NewReader(strings.NewReader(tt.input)))
		tok := l.NextToken()
		if tok.Type != Z80_REG16 {
			t.Errorf("tokenize %q. expected Type Z80_REG16. got %#v", tt.input, yySymNames[yyXLAT[tok.Type]])
		}
		if tok.Literal != tt.input {
			t.Errorf("tokenize %q. expected Literal %q. got %#v", tt.input, tt.input, tok)
		}
		if tok.Op != tt.op {
			t.Errorf("tokenize %q. expected Op '%c'. got '%c'", tt.input, tt.op, tok.Op)
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
		l := NewLexer(bufio.NewReader(strings.NewReader(tt.input)))
		tok := l.NextToken()
		if tok.Type != Z80_FLAG {
			t.Errorf("tokenize %q. expected Type Z80_FLAG. got %#v", tt.input, yySymNames[yyXLAT[tok.Type]])
		}
		if tok.Literal != tt.input {
			t.Errorf("tokenize %q. expected Literal %q. got %#v", tt.input, tt.input, tok)
		}
		if tok.Op != tt.op {
			t.Errorf("tokenize %q. expected Op '%c'. got '%c'", tt.input, tt.op, tok.Op)
		}
	}
}

func TestZ80Instructions(t *testing.T) {
	input := "LD PUSH POP EX EXX LDI LDIR LDDR CPI CPIR CPDR ADD ADC SUB SBC AND OR " +
		"CP INC DEC DAA CPL NEG CCF SCF NOP HALT DI EI IM " +
		"RLCA RLA RRCA RRA RLC RL RR SLA SRA SRL RLD RRD " +
		"BIT SET RES JP JR DJNZ CALL RET RETI RETN RST " +
		"IN INI INIR INDR OUT OUTI OTIR OUTD OTDR"

	l := NewLexer(bufio.NewReader(strings.NewReader(input)))
	for {
		tok := l.NextToken()
		if tok.Type == EOL {
			break
		}

		expectedToken, ok := Z80OpCodes[tok.Literal]
		if !ok {
			t.Errorf("instruction %q not found", tok.Literal)
			continue
		}
		if tok.Type != expectedToken.Type {
			t.Errorf("expected Type %s. got %#v", yySymNames[yyXLAT[tok.Type]], tok)
		}
		if tok.Literal != expectedToken.Literal {
			t.Errorf("expected Literal %q. got %#v", expectedToken.Literal, tok)
		}
		if tok.Op != expectedToken.Op {
			t.Errorf("expected Op '%d'. got %#v", expectedToken.Op, tok)
		}

	}
}
