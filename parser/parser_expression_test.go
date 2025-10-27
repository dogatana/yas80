package parser

import (
	"bufio"
	"strings"
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
		l := NewLexer(bufio.NewReader(strings.NewReader(tt.input)))
		_ = Parse(l)
		if len(Root.Statements) != 1 {
			t.Fatalf("parsing %s returns %d statements. not 1", tt.input, len(Root.Statements))
		}
		expr, ok := Root.Statements[0].(*ExpressionStatement)
		if !ok {
			t.Fatalf("parsing %s Statements[0] is not ExpressionStatement. got %T", tt.input, Root.Statements[0])
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
