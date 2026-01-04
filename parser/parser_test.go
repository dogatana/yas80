package parser

import (
	"testing"
	"yas80/errcode"
	"yas80/errtest"
)

func TestParseNumberLiteral(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		// 0-
		{"_ = 123_456", 123456},
		{"_ = 0x1234", 0x1234},
		{"_ = 0X4567", 0x4567},
		{"_ = $89ab", 0x89ab},
		{"_ = 1234h", 0x1234},
		// 5-
		{"_ = 5678H", 0x5678},
		{"_ = $0123", 0x123},
		{"_ = 0o377", 0o377},
		{"_ = 0O123", 0o123},
		{"_ = 0b1111", 0b1111},
		// 10-
		{"_ = 0B1010", 0b1010},
		{"_ = %1001", 0b1001},
		// 定数畳み込み
		{"_ = 2 + 3 * 4", 14},
		{"_ = 2 + 3 * (4 + 5)", 29},
		{"_ = 15 & %1010", 10},
		// 15-
		{"_ = 15 & 0b101", 5},
		{"_ = %1010 | %0101", 15},
		{"_ = %1010 ^ %0101", 15},
		{"_ = 15 ^ 15", 0},
		{"_ = 0b1111 ^ 0b0101", 0b1010},
		// 20-
		{"_ = 1 << 8", 256},
		{"_ = 8 >> 1", 4},

		{"_ = +123", 123},
		{"_ = -123", -123},
		{"_ = ~0", -1},
		// 25-
		{"_ = ~-1", 0},
		{"_ = ~0xa5 & 0xff", 0x5a},
		{"_ = ~5ah & 0xff", 0xa5},

		{"_ = 1 < 2", 1},
		{"_ = 1 <= 1", 1},
		// 30-
		{"_ = 1 <= 2", 1},

		{"_ = 1 > 2", 0},
		{"_ = 1 >= 1", 1},
		{"_ = 1 >= 2", 0},

		{"_ = 3 < 2", 0},
		// 35-
		{"_ = 3 <= 3", 1},
		{"_ = 3 <= 2", 0},

		{"_ = 3 > 2", 1},
		{"_ = 3 >= 3", 1},
		{"_ = 3 >= 2", 1},
		// 40-
		{"_ = 3 < 2", 0},
		{"_ = 3 <= 3", 1},
		{"_ = 3 <= 2", 0},

		{"_ = 2 == 2", 1},
		{"_ = 2 != 2", 0},
		// 45-
		{"_ = !0", 1},
		{"_ = !2", 0},
		{`_ = !""`, 1},
		{`_ = !"a"`, 0},
		{`_ = !("" + "")`, 1},
		// 50-
		{`_ = !("a" + "b")`, 0},
		{`_ = 1111b`, 15},
		{`_ = 0101B`, 5},
		{`_ = 1010_0101B`, 0xa5},
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
		err      string
	}{
		// 0-
		{`_ =  "abc" `, "abc", ""},
		{`_ = "a\r\n" `, "a\r\n", ""},
		{`_ =  " abc" `, " abc", ""},
		{`_ =  "abc " `, "abc ", ""},
		{`_ =  " abc " `, " abc ", ""},
		// 5-
		{`_ = "xxx '\"\\\0\a\b\f\n\r\t\v xxx"`, "xxx '\"\\\x00\a\b\f\n\r\t\v xxx", ""},
		{input: "_ = 'abc\ndef'", err: errcode.ESTR_END_QUOTE},
		{input: "_ = \"abc\ndef\"", err: errcode.ESTR_END_QUOTE},
		{input: "_ = 'abc\tndef'", err: errcode.ESTR_CTRL},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)

		if tt.err != "" {
			testLogMessage(t, tn, tt.err, l.logger)
			ename := errtest.ErrcodeNames[tt.err]
			if ename[0] == 'E' {
				continue
			}
		}

		if len(l.logger.Errors) > 0 {
			t.Fatalf("[%d] %d errors", tn, len(l.logger.Errors))
		}
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
		{"_ = []", 0, []int{}},
		{"_ = [1]", 1, []int{1}},
		{"_ = [1, 2]", 2, []int{1, 2}},
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

func TestParseLabelStatement(t *testing.T) {
	tests := []struct {
		input     string
		identType int
		name      string
	}{
		{"abc:", IDENT, "ABC"},
		{"abc :ld a, a", IDENT, "ABC"},
		{".abc:", LOCAL_IDENT, ".ABC"},
		{".abc: ld a,a", LOCAL_IDENT, ".ABC"},
		{"@abc:", AT_IDENT, "@ABC"},
		{"@abc: ld a,a", AT_IDENT, "@ABC"},
	}
	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		if len(prog.Statements) == 0 {
			t.Fatalf("[%d] statements empty", tn)
		}
		var node Node
		switch stmt := prog.Statements[0].(type) {
		case *LabelStatement:
			node = stmt.Name
		case *Z80Instruction:
			node = stmt.Label
		default:
			t.Errorf("[%d] not *LabeStatement nor *Z80Instruction", tn)
		}
		id, ok := node.(*Ident)
		if !ok {
			t.Errorf("[%d] Name not *Ident. got %T", tn, node)
		}
		if id.IdentType != tt.identType {
			t.Errorf("[%d] IdentType not %s. got %s", tn, nodeLiteral(tt.identType), nodeLiteral(id.IdentType))
		}
		if id.Name != tt.name {
			t.Errorf("[%d] Name not %s. got %s", tn, tt.name, id.Name)
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
