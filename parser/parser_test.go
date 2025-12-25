package parser

import (
	"testing"
)

func TestParseNumberLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		// 0-
		{"a=123_456", 123456},
		{"a=0x1234", 0x1234},
		{"a=0X4567", 0x4567},
		{"a=$89ab", 0x89ab},
		{"a=1234h", 0x1234},
		// 5-
		{"a=5678H", 0x5678},
		{"a=$0123", 0x123},
		{"a=0o377", 0o377},
		{"a=0O123", 0o123},
		{"a=0b1111", 0b1111},
		// 10-
		{"a=0B1010", 0b1010},
		{"a=%1001", 0b1001},
		// 定数畳み込み
		{"a=2 + 3 * 4", 14},
		{"a=2 + 3 * (4 + 5)", 29},
		{"a=15 & %1010", 10},
		// 15-
		{"a=15 & 0b101", 5},
		{"a=%1010 | %0101", 15},
		{"a=%1010 ^ %0101", 15},
		{"a=15 ^ 15", 0},
		{"a=0b1111 ^ 0b0101", 0b1010},
		// 20-
		{"a=1 << 8", 256},
		{"a=8 >> 1", 4},

		{"a=+123", 123},
		{"a=-123", -123},
		{"a=~0", -1},
		// 25-
		{"a=~-1", 0},
		{"a=~0xa5 & 0xff", 0x5a},
		{"a=~5ah & 0xff", 0xa5},

		{"a=1 < 2", 1},
		{"a=1 <= 1", 1},
		// 30-
		{"a=1 <= 2", 1},

		{"a=1 > 2", 0},
		{"a=1 >= 1", 1},
		{"a=1 >= 2", 0},

		{"a=3 < 2", 0},
		// 35-
		{"a=3 <= 3", 1},
		{"a=3 <= 2", 0},

		{"a=3 > 2", 1},
		{"a=3 >= 3", 1},
		{"a=3 >= 2", 1},
		// 40-
		{"a=3 < 2", 0},
		{"a=3 <= 3", 1},
		{"a=3 <= 2", 0},

		{"a=2 == 2", 1},
		{"a=2 != 2", 0},
		// 45-
		{"a=!0", 1},
		{"a=!2", 0},
		{`a=!""`, 1},
		{`a=!"a"`, 0},
		{`a=!("" + "")`, 1},
		// 50-
		{`a=!("a" + "b")`, 0},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		if len(prog.Statements) == 0 {
			t.Fatalf("[%d] %d statements", tn, len(prog.Statements))
		}
		stmt := testAssignStatement(t, tn, prog.Statements[0])
		testNumberLiteral(t, tn, stmt.Value, tt.expected)
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

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		if len(prog.Statements) == 0 {
			t.Fatalf("[%d] %d statements", tn, len(prog.Statements))
		}
		stmt := testAssignStatement(t, tn, prog.Statements[0])
		testStringLiteral(t, tn, stmt.Value, tt.expected)
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
	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := testAssignStatement(t, tn, prog.Statements[0])

		array, ok := stmt.Value.(*ArrayLiteral)
		if !ok {
			t.Errorf("[%d] not *ArrayLiteral. got %T", tn, stmt)
		}
		eles := array.Elements
		if len(eles.Expressions) != tt.count {
			t.Errorf("[%d] must have %d elements. got %d", tn, tt.count, len(array.Elements.Expressions))
		}
		for i, n := range tt.expected {
			value := eles.Expressions[i].(*NumberLiteral).Value
			if value != n {
				t.Errorf("[%d] array[%d] not %d. got %d", tn, i, n, value)
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
	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		if len(prog.Statements) == 0 {
			t.Fatalf("[%d] statements empty", tn)
		}
		stmt, ok := prog.Statements[0].(*LabelStatement)
		if !ok {
			t.Errorf("[%d] prog.Statemtes[0] is not LabelStatement. got %T", tn, prog.Statements[0])
		}
		if stmt.Name.NodeType() != NODE_LABEL {
			t.Errorf("[%d] Value.Nodetyp() not NODE_LABEL. got %s", tn, nodeLiteral(int(stmt.Name.NodeType())))
		}
		if stmt.Name.NodeSubType() != tt.labelType {
			t.Errorf("[%d] Value.Nodetyp() not %s. got %s", tn, nodeLiteral(int(tt.labelType)), nodeLiteral(int(stmt.Name.NodeType())))
		}
		if stmt.Name.Name != tt.name {
			t.Errorf("[%d] Value.Name not %q. got %q", tn, tt.name, stmt.Name.Name)
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
		{".abc: ld a,a", DOT_IDENT, ".ABC"},
		// {".abc ", DOT_IDENT, ".ABC"}, // ラベルには : を必須としたので除外
	}
	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		if len(prog.Statements) == 0 {
			t.Fatalf("[%d] statements empty", tn)
		}
		switch stmt := prog.Statements[0].(type) {
		case *LabelStatement:
			name := stmt.Name.Name
			if name != tt.expected {
				t.Errorf("[%d] Label.Name is not %q. got %q", tn, tt.expected, name)
			}
		case *Z80Instruction:
			name := stmt.Label.Name
			if name != tt.expected {
				t.Errorf("[%d] Label.Name is not %q. got %q", tn, tt.expected, name)
			}
		default:
			t.Errorf("[%d] not *LabeStatement nor Z80Instruction", tn)
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

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)

		if len(prog.Statements) == 0 {
			t.Fatalf("[%d] statements empty", tn)
		}
		stmt := testAssignStatement(t, tn, prog.Statements[0])
		ident, ok := stmt.Value.(*DotIdent)
		if !ok {
			t.Errorf("[%d] not *DotIdent got %T", tn, stmt.Value)
		}
		if ident.Left != "ABC" || ident.Right != ".DEF" {
			t.Errorf("[%d] not %q. got %q", tn, tt.expected, ident.String())
		}
	}
}
