package parser

import (
	"fmt"
	"testing"
)

func TestNumberLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"123_456", 123456},
		{"0x1234", 0x1234},
		{"0X4567", 0x4567},
		{"$89ab", 0x89ab},
		{"1234h", 0x1234},
		{"5678H", 0x5678},
		{"$0123", 0x123},
		{"0o377", 0o377},
		{"0O123", 0o123},
		{"0b1111", 0b1111},
		{"0B1010", 0b1010},
		{"%1001", 0b1001},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := Parse(l)
		ec, wc, _ := l.logger.Count()
		if ec > 0 || wc > 0 {
			fmt.Println("parsing", tt.input)
			t.Fatalf("%d errors and %d warnigs", ec, wc)
		}
		if len(prog.Statements) == 0 {
			fmt.Println("parsing", tt.input)
			t.Fatalf("%d statements", len(prog.Statements))
		}
		stmt, ok := prog.Statements[0].(*ExpressionStatement)
		if !ok {
			fmt.Println("parsing", tt.input)
			t.Errorf("prog.Statemtes[0] is not ExpressionStatement. got %T", prog.Statements[0])
		}
		literal, ok := stmt.Value.(*NumberLiteral)
		if !ok {
			fmt.Println("parsing", tt.input)
			t.Errorf("Value is not *NumberLiteral. got %T", stmt.Value)
		}
		if literal.Value != tt.expected {
			fmt.Println("parsing", tt.input)
			t.Errorf("Value is not %d. got %d", tt.expected, stmt.Value)
		}
	}
}

func TestStringLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{` "abc" `, "abc"},
		{`"a\r\n" `, "a\\r\\n"},
		{` " abc" `, " abc"},
		{` "abc " `, "abc "},
		{` " abc " `, " abc "},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := Parse(l)
		ec, wc, _ := l.logger.Count()
		if ec > 0 || wc > 0 {
			fmt.Println("parsing", tt.input)
			t.Fatalf("%d errors and %d warnigs", ec, wc)
		}
		if len(prog.Statements) == 0 {
			fmt.Println("parsing", tt.input)
			t.Fatalf("%d statements", len(prog.Statements))
		}
		stmt, ok := prog.Statements[0].(*ExpressionStatement)
		if !ok {
			fmt.Println("parsing", tt.input)
			t.Errorf("prog.Statemtes[0] is not ExpressionStatement. got %T", prog.Statements[0])
		}
		literal, ok := stmt.Value.(*StringLiteral)
		if !ok {
			fmt.Println("parsing", tt.input)
			t.Errorf("Value is not *StringLiteral. got %T", stmt.Value)
		}
		if literal.Value != tt.expected {
			fmt.Println("parsing", tt.input)
			t.Errorf("Value is not %q. got %q", tt.expected, stmt.Value)
		}
	}
}
func TestLableStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" abc: ", "abc"},
		{" abc : ", "abc"},
		{"abc:", "abc"},
		{" abc:ld a,a ", "abc"},
		{" abc :ld a, a", "abc"},
		{"abc: ld a, a", "abc"},
		{" .def: ", ".def"},
		{" .def : ", ".def"},
		{".def:", ".def"},
		{" .def:ld a,a ", ".def"},
		{" .def :ld a, a", ".def"},
		{".def: ld a, a", ".def"},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := Parse(l)
		ec, wc, _ := l.logger.Count()
		if ec > 0 || wc > 0 {
			t.Fatalf("parsing %s. returns %d errors and %d warnigs", tt.input, ec, wc)
		}
		if len(prog.Statements) == 0 {
			t.Fatalf("parsing %s. statements empty", tt.input)
		}
		stmt, ok := prog.Statements[0].(*LabelStatement)
		if !ok {
			t.Errorf("parsing %s. prog.Statemtes[0] is not LabelStatement. got %T", tt.input, prog.Statements[0])
		}
		name := stmt.Value.Name
		if name != tt.expected {
			t.Errorf("parsing %s. Label.Name is not %q. got %q", tt.input, tt.expected, name)
		}
	}
}

func TestDotIdent(t *testing.T) {
	input := "abc.def"

	l := newLexerForTest(input)
	prog := Parse(l)
	ec, wc, _ := l.logger.Count()
	if ec > 0 || wc > 0 {
		t.Fatalf("parsing %s. returns %d errors and %d warnigs", input, ec, wc)
	}
	if len(prog.Statements) == 0 {
		t.Fatalf("parsing %s. statements empty", input)
	}
	stmt, ok := prog.Statements[0].(*ExpressionStatement)
	if !ok {
		t.Errorf("parsing %s. prog.Statemtes[0] is not ExpressionStatement. got %T", input, prog.Statements[0])
	}
	ident, ok := stmt.Value.(*DotIdent)
	if !ok {
		t.Errorf("parsing %s. not Expression got %T", input, stmt.Value)
	}
	if ident.Left != "ABC" || ident.Right != "DEF" {
		t.Errorf("parsing %s. not ABC.DEF. got %q", input, ident.String())
	}
}
