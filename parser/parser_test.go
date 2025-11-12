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
		// 定数畳み込み
		{"2 + 3 * 4", 14},
		{"2 + 3 * (4 + 5)", 29},
		{"15 & %1010", 10},
		{"15 & 0b101", 5},
		{"%1010 | %0101", 15},
		{"%1010 ^ %0101", 15},
		{"15 ^ 15", 0},
		{"0b1111 ^ 0b0101", 0b1010},
		{"1 << 8", 256},
		{"8 >> 1", 4},

		{"+123", 123},
		{"-123", -123},
		{"~0", -1},
		{"~-1", 0},
		{"~0xa5 & 0xff", 0x5a},
		{"~5ah & 0xff", 0xa5},

		{"1 < 2", 1},
		{"1 <= 1", 1},
		{"1 <= 2", 1},

		{"1 > 2", 0},
		{"1 >= 1", 1},
		{"1 >= 2", 0},

		{"3 < 2", 0},
		{"3 <= 3", 1},
		{"3 <= 2", 0},

		{"3 > 2", 1},
		{"3 >= 3", 1},
		{"3 >= 2", 1},

		{"3 < 2", 0},
		{"3 <= 3", 1},
		{"3 <= 2", 0},

		{"2 == 2", 1},
		{"2 != 2", 0},

		{"!0", 1},
		{"!2", 0},
		{`!""`, 1},
		{`!"a"`, 0},
		{`!("" + "")`, 1},
		{`!("a" + "b")`, 0},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := Parse(l)
		ec, wc, _ := l.logger.Count()
		if ec > 0 || wc > 0 {
			fmt.Printf("input %q\n", tt.input)
			t.Fatalf("%d errors and %d warnigs", ec, wc)
		}
		if len(prog.Statements) == 0 {
			fmt.Printf("input %q\n", tt.input)
			t.Fatalf("%d statements", len(prog.Statements))
		}
		testNumberLiteralStatement(t, tt.input, prog.Statements[0], tt.expected)
	}
}

func testNumberLiteralStatement(t *testing.T, input string, node Node, expected int) {
	stmt := testExpressionStatement(t, input, node)

	literal, ok := stmt.Value.(*NumberLiteral)
	if !ok {
		fmt.Printf("input %q\n", input)
		t.Errorf("not *NumberLiteral. got %T", literal)
	}
	if literal.Value != expected {
		fmt.Printf("input %q\n", input)
		t.Errorf("not %d. got %d", expected, literal.Value)
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
			fmt.Printf("input %q\n", tt.input)
			t.Fatalf("%d errors and %d warnigs", ec, wc)
		}
		if len(prog.Statements) == 0 {
			fmt.Println("parsing", tt.input)
			t.Fatalf("%d statements", len(prog.Statements))
		}
		testStringLiteralStatement(t, tt.input, prog.Statements[0], tt.expected)
	}
}

func testStringLiteralStatement(t *testing.T, input string, node Node, expected string) {
	stmt := testExpressionStatement(t, input, node)

	literal, ok := stmt.Value.(*StringLiteral)
	if !ok {
		fmt.Printf("input %q\n", input)
		t.Errorf("not *StringLiteral. got %T", literal)
	}
	if literal.Value != expected {
		fmt.Printf("input %q\n", input)
		t.Errorf("not %q. got %q", expected, literal.Value)
	}
}

func TestArrayLiteral(t *testing.T) {
	tests := []struct {
		input    string
		count    int
		expected []int
	}{
		{"[]", 0, []int{}},
		{"[1]", 1, []int{1}},
		{"[1, 2]", 2, []int{1, 2}},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := Parse(l)
		ec, wc, _ := l.logger.Count()
		if ec > 0 || wc > 0 {
			fmt.Printf("input %q\n", tt.input)
			t.Fatalf("%d errors and %d warnigs", ec, wc)
		}
		stmt := testExpressionStatement(t, tt.input, prog.Statements[0])

		array, ok := stmt.Value.(*ArrayLiteral)
		if !ok {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("not *ArrayLiteral. got %T", stmt)
		}
		eles := array.Elements
		if len(eles.Expressions) != tt.count {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("must have %d elements. got %d", tt.count, len(array.Elements.Expressions))
		}
		for i, n := range tt.expected {
			value := eles.Expressions[i].(*NumberLiteral).Value
			if value != n {
				fmt.Printf("input %q\n", tt.input)
				t.Errorf("array[%d] not %d. got %d", i, n, value)
			}
		}
	}
}

func TestLableStatement(t *testing.T) {
	tests := []struct {
		input     string
		TokenType int
		expected  string
	}{
		{"abc:", IDENT, "abc"},
		{"abc :ld a, a", IDENT, "abc"},
		{".abc ", DOT_IDENT, ".abc"},
		{".abc: ld a,a", DOT_IDENT, ".abc"},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := Parse(l)
		ec, wc, _ := l.logger.Count()
		if ec > 0 || wc > 0 {
			fmt.Printf("input %q\n", tt.input)
			t.Fatalf("%d errors and %d warnigs", ec, wc)
		}
		if len(prog.Statements) == 0 {
			fmt.Printf("input %q\n", tt.input)
			t.Fatal("statements empty")
		}
		stmt, ok := prog.Statements[0].(*LabelStatement)
		if !ok {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("prog.Statemtes[0] is not LabelStatement. got %T", prog.Statements[0])
		}
		name := stmt.Value.Name
		if name != tt.expected {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("Label.Name is not %q. got %q", tt.expected, name)
		}
	}
}

func TestDotIdent(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abc.def", "ABC.DEF"},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := Parse(l)
		ec, wc, _ := l.logger.Count()
		if ec > 0 || wc > 0 {
			fmt.Printf("input %q\n", tt.input)
			t.Fatalf("%d errors and %d warnigs", ec, wc)
		}
		if len(prog.Statements) == 0 {
			fmt.Printf("input %q\n", tt.input)
			t.Fatal("statements empty")
		}
		stmt := testExpressionStatement(t, tt.input, prog.Statements[0])
		ident, ok := stmt.Value.(*DotIdent)
		if !ok {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("not *DotIdent got %T", stmt.Value)
		}
		if ident.Left != "ABC" || ident.Right != "DEF" {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("not %q. got %q", tt.expected, ident.String())
		}
	}
}
