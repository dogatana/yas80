package parser

import (
	"fmt"
	"testing"
)

func TestParseNumberLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"a=123_456", 123456},
		{"a=0x1234", 0x1234},
		{"a=0X4567", 0x4567},
		{"a=$89ab", 0x89ab},
		{"a=1234h", 0x1234},
		{"a=5678H", 0x5678},
		{"a=$0123", 0x123},
		{"a=0o377", 0o377},
		{"a=0O123", 0o123},
		{"a=0b1111", 0b1111},
		{"a=0B1010", 0b1010},
		{"a=%1001", 0b1001},
		// 定数畳み込み
		{"a=2 + 3 * 4", 14},
		{"a=2 + 3 * (4 + 5)", 29},
		{"a=15 & %1010", 10},
		{"a=15 & 0b101", 5},
		{"a=%1010 | %0101", 15},
		{"a=%1010 ^ %0101", 15},
		{"a=15 ^ 15", 0},
		{"a=0b1111 ^ 0b0101", 0b1010},
		{"a=1 << 8", 256},
		{"a=8 >> 1", 4},

		{"a=+123", 123},
		{"a=-123", -123},
		{"a=~0", -1},
		{"a=~-1", 0},
		{"a=~0xa5 & 0xff", 0x5a},
		{"a=~5ah & 0xff", 0xa5},

		{"a=1 < 2", 1},
		{"a=1 <= 1", 1},
		{"a=1 <= 2", 1},

		{"a=1 > 2", 0},
		{"a=1 >= 1", 1},
		{"a=1 >= 2", 0},

		{"a=3 < 2", 0},
		{"a=3 <= 3", 1},
		{"a=3 <= 2", 0},

		{"a=3 > 2", 1},
		{"a=3 >= 3", 1},
		{"a=3 >= 2", 1},

		{"a=3 < 2", 0},
		{"a=3 <= 3", 1},
		{"a=3 <= 2", 0},

		{"a=2 == 2", 1},
		{"a=2 != 2", 0},

		{"a=!0", 1},
		{"a=!2", 0},
		{`a=!""`, 1},
		{`a=!"a"`, 0},
		{`a=!("" + "")`, 1},
		{`a=!("a" + "b")`, 0},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)
		if len(prog.Statements) == 0 {
			fmt.Printf("input %q\n", tt.input)
			t.Fatalf("%d statements", len(prog.Statements))
		}
		stmt := testAsignStatement(t, tt.input, prog.Statements[0])
		testNumberLiteral(t, tt.input, stmt.Value, tt.expected)
	}
}

func TestParseStringLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`a= "abc" `, "abc"},
		{`a="a\r\n" `, "a\\r\\n"},
		{`a= " abc" `, " abc"},
		{`a= "abc " `, "abc "},
		{`a= " abc " `, " abc "},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)
		if len(prog.Statements) == 0 {
			fmt.Println("parsing", tt.input)
			t.Fatalf("%d statements", len(prog.Statements))
		}
		stmt := testAsignStatement(t, tt.input, prog.Statements[0])
		testStringLiteral(t, tt.input, stmt.Value, tt.expected)
	}
}

func TestParseArrayLiteral(t *testing.T) {
	tests := []struct {
		input    string
		count    int
		expected []int
	}{
		{"a=[]", 0, []int{}},
		{"a=[1]", 1, []int{1}},
		{"a=[1, 2]", 2, []int{1, 2}},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)
		stmt := testAsignStatement(t, tt.input, prog.Statements[0])

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

func TestParseLabel(t *testing.T) {
	tests := []struct {
		input     string
		labelType NodeSubType
		name      string
	}{
		{"abc:", NODE_LABEL, "ABC"},
		{".abc:", NODE_LOCAL_LABEL, ".ABC"},
		{"@abc:", NODE_AT_LABEL, "@ABC"},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)
		if len(prog.Statements) == 0 {
			t.Fatal("statements empty")
		}
		stmt, ok := prog.Statements[0].(*LabelStatement)
		if !ok {
			t.Errorf("prog.Statemtes[0] is not LabelStatement. got %T", prog.Statements[0])
		}
		if stmt.Value.NodeType() != NODE_LABEL {
			t.Errorf("Value.Nodetyp() not NODE_LABEL. got %s", nodeLiteral(int(stmt.Value.NodeType())))
		}
		if stmt.Value.NodeSubType() != tt.labelType {
			t.Errorf("Value.Nodetyp() not %s. got %s", nodeLiteral(int(tt.labelType)), nodeLiteral(int(stmt.Value.NodeType())))
		}
		if stmt.Value.Name != tt.name {
			t.Errorf("Value.Name not %q. got %q", tt.name, stmt.Value.Name)
		}
	}
}
func TestParseLabelStatement(t *testing.T) {
	tests := []struct {
		input     string
		TokenType int
		expected  string
	}{
		{"abc:", IDENT, "ABC"},
		{"abc :ld a, a", IDENT, "ABC"},
		// {".abc ", DOT_IDENT, ".ABC"}, // ラベルには : を必須としたので除外
		{".abc: ld a,a", DOT_IDENT, ".ABC"},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)
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

func TestParseDotIdent(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"a = abc.def", "ABC.DEF"},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)

		if len(prog.Statements) == 0 {
			fmt.Printf("input %q\n", tt.input)
			t.Fatal("statements empty")
		}
		stmt := testAsignStatement(t, tt.input, prog.Statements[0])
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
