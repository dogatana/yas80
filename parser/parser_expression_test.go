package parser

import (
	"testing"
)

func TestExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1", 1},
		{"-1", -1},
		{"~0", -1},
		{"~1", -2},
		{"~-1", 0},
		{"!0", 1},
		{"!1", 0},
		{"!-1", 0},
		{"2 + 3", 5},
		{"2 - 3", -1},
		{"2 * 3", 6},
		{"4 / 2", 2},
		{"2 + 3 * 4", 14},
		{"2 * 3 + 4", 10},
		{"2 * (3 + 4)", 14},
		{"0x55 | 0xaa", 255},
		{"0x55 & 0xaa", 0},
		{"0xf ^ 1", 14},
		{"1 << 8", 256},
		{"256 >> 4", 16},
		{"1 > 0", 1},
		{"1 >= 0", 1},
		{"1 < 0", 0},
		{"1 <= 0", 0},
		{"2 > 3", 0},
		{"2 >= 3", 0},
		{"2 < 3", 1},
		{"2 <= 3", 1},
		{"2 != 3", 1},
		{"2 == 3", 0},
		{"100 || 0", 1},
		{"100 && 0", 0},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog, ec, wc := Parse(l)
		if ec > 0 || wc > 0 {
			t.Errorf("parsing %s returns %d errors and %d warnigs", tt.input, ec, wc)
		}
		if len(prog.Statements) != 1 {
			t.Fatalf("parsing %s returns %d statements. not 1", tt.input, len(prog.Statements))
		}
		expr, ok := prog.Statements[0].(*ExpressionStatement)
		if !ok {
			t.Fatalf("parsing %s Statements[0] is not ExpressionStatement. got %T", tt.input, prog.Statements[0])
		}
		literal, ok := expr.Value.(*NumberLiteral)
		if !ok {
			t.Fatalf("parsing %s Statements[0].Expression is not Expression. got %T", tt.input, expr)
		}
		if literal.Value != tt.expected {
			t.Errorf("parsing %s NumberLiteral.Value is not %d. got %#v", tt.input, tt.expected, literal.Value)
		}
	}
}

func TestCallFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"fn()", "fn()"},
		{"  fn(   )  ", "fn()"},
		{"fn(1)", "fn(1)"},
		{"fn(1,2,3)", "fn(1, 2, 3)"},
		{"fn(1,2+3,4*5)", "fn(1, 5, 20)"},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog, ec, wc := Parse(l)
		if ec > 0 || wc > 0 {
			l.ErrorStore.Print()
			t.Fatalf("parsing %s returns %d errors and %d warnigs", tt.input, ec, wc)
		}
		if len(prog.Statements) != 1 {
			t.Fatalf("parsing %s returns %d statements. not 1", tt.input, len(prog.Statements))
		}
		stmt, ok := prog.Statements[0].(*ExpressionStatement)
		if !ok {
			t.Fatalf("parsing %s Statements[0] is not ExpressionStatement . got %T", tt.input, prog.Statements[0])
		}
		call, ok := stmt.Value.(*CallExpression)
		if !ok {
			t.Fatalf("parsing %s Expression.Value is not CallExpression . got %T", tt.input, stmt.Value)
		}
		if call.String() != tt.expected {
			t.Errorf("CallExpression.String() is not %q. got %q", tt.expected, call.String())
		}
	}
}

func TestArrayVariable(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"var a0 = []", "VAR a0 = []"},
		{"var array = [1,2,3]", "VAR array = [1, 2, 3]"},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog, ec, wc := Parse(l)
		if ec > 0 || wc > 0 {
			l.ErrorStore.Print()
			t.Fatalf("parsing %s returns %d errors and %d warnigs", tt.input, ec, wc)
		}
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

func TestArrayElement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a0 [ ]", "a0[]"},
		{"a1 [ 123 ]", "a1[123]"},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog, ec, wc := Parse(l)
		if ec > 0 || wc > 0 {
			l.ErrorStore.Print()
			t.Fatalf("parsing %s returns %d errors and %d warnigs", tt.input, ec, wc)
		}
		if len(prog.Statements) != 1 {
			t.Fatalf("parsing %s returns %d statements. not 1", tt.input, len(prog.Statements))
		}
		stmt, ok := prog.Statements[0].(*ExpressionStatement)
		if !ok {
			t.Fatalf("parsing %s Statements[0] is not ExpressionStatement . got %T", tt.input, prog.Statements[0])
		}
		if stmt.String() != tt.expected {
			t.Errorf("stmt.String() is not %q. got %q", tt.expected, stmt.String())
		}
	}
}
