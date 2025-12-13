package parser

import (
	"testing"
)

func TestParseExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"a=1", 1},
		{"a=-1", -1},
		{"a=~0", -1},
		{"a=~1", -2},
		{"a=~-1", 0},
		{"a=!0", 1},
		{"a=!1", 0},
		{"a=!-1", 0},
		{"a=2 + 3", 5},
		{"a=2 - 3", -1},
		{"a=2 * 3", 6},
		{"a=4 / 2", 2},
		{"a=2 + 3 * 4", 14},
		{"a=2 * 3 + 4", 10},
		{"a=2 * (3 + 4)", 14},
		{"a=0x55 | 0xaa", 255},
		{"a=0x55 & 0xaa", 0},
		{"a=0xf ^ 1", 14},
		{"a=1 << 8", 256},
		{"a=256 >> 4", 16},
		{"a=1 > 0", 1},
		{"a=1 >= 0", 1},
		{"a=1 < 0", 0},
		{"a=1 <= 0", 0},
		{"a=2 > 3", 0},
		{"a=2 >= 3", 0},
		{"a=2 < 3", 1},
		{"a=2 <= 3", 1},
		{"a=2 != 3", 1},
		{"a=2 == 3", 0},
		{"a=100 || 0", 1},
		{"a=100 && 0", 0},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)

		if len(prog.Statements) != 1 {
			t.Fatalf("parsing %s returns %d statements. not 1", tt.input, len(prog.Statements))
		}
		stmt := testAssignStatement(t, tt.input, prog.Statements[0])
		testNumberLiteral(t, tt.input, stmt.Value, tt.expected)
	}
}

func TestParseCallFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a=aFunc()", "AFUNC()"},
		{"  a=aFunc(   )  ", "AFUNC()"},
		{"a = aFunc(1)", "AFUNC(1)"},
		{"a=aFunc(1,2,3)", "AFUNC(1, 2, 3)"},
		{"a=aFunc(1,2+3,4*5)", "AFUNC(1, 5, 20)"},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)

		if len(prog.Statements) != 1 {
			t.Fatalf("parsing %s returns %d statements. not 1", tt.input, len(prog.Statements))
		}
		stmt := testAssignStatement(t, tt.input, prog.Statements[0])
		if stmt.Value.String() != tt.expected {
			t.Errorf("stmt.Value.String() is not %q. got %q", tt.expected, stmt.Value.String())
		}
	}
}

func TestParseArrayVariable(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"var a0 = []", "VAR A0 = []"},
		{"var array = [1,2,3]", "VAR ARRAY = [1, 2, 3]"},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)

		if len(prog.Statements) != 1 {
			t.Fatalf("parsing %s returns %d statements. not 1", tt.input, len(prog.Statements))
		}
		stmt, ok := prog.Statements[0].(*VariableStatement)
		if !ok {
			t.Fatalf("parsing %s Statements[0] is not VariableStatement . got %T", tt.input, prog.Statements[0])
		}
		if stmt.String() != tt.expected {
			t.Errorf("stmt.String() is not %q. got %q", tt.expected, stmt.String())
		}
	}
}

func TestParseArrayElement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a = a0 [ 123 ]", "A0[123]"},
		{"a = a1 [ 1 + 2 * 3 ]", "A1[7]"},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)

		if len(prog.Statements) != 1 {
			t.Fatalf("parsing %s returns %d statements. not 1", tt.input, len(prog.Statements))
		}
		stmt := testAssignStatement(t, tt.input, prog.Statements[0])
		if stmt.Value.String() != tt.expected {
			t.Errorf("stmt.String() is not %q. got %q", tt.expected, stmt.Value.String())
		}
	}
}

func TestParseStringExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`a= "abc" `, "abc"},
		{`a= "abc" + "def" `, "abcdef"},
		{`a= "abc" + "@" + "def" `, "abc@def"},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)

		if len(prog.Statements) != 1 {
			t.Fatalf("parsing %s returns %d statements. not 1", tt.input, len(prog.Statements))
		}
		stmt := testAssignStatement(t, tt.input, prog.Statements[0])
		testStringLiteral(t, tt.input, stmt.Value, tt.expected)
	}
}
