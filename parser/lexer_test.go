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
