package parser

import (
	"testing"
	"yas80/errcode"
	"yas80/internal/testutil"
)

// parse 時の定数式演算結果のテスト
func TestParsePrefixExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"_ = +123", 123},
		{"_ = -123", -123},
		{"_ = --123", 123},
		{"_ = ~0", -1},
		{"_ = ~1", -2},
		{"_ = ~-1", 0},
		{"_ = !0", 1},
		{"_ = !1", 0},
		{"_ = !2", 0},
		{"_ = !!0", 0},
		{"_ = !!1", 1},
		{`_ = !""`, 1},
		{`_ = !"a"`, 0},
		{`_ = !("" + "")`, 1},
		{`_ = !("a" + "b")`, 0},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)
		as := testAssignStatement(t, tn, stmt)
		testNumberLiteral(t, tn, as.Value, tt.expected)
	}
}

// parse 時の定数式演算結果のテスト
func TestParseInfixExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"_ = 2 + 3", 5},
		{"_ = 2 - 3", -1},
		{"_ = 2 * 3", 6},
		{"_ = 4 / 2", 2},
		{"_ = 5 % 2", 1},
		{"_ = 2 + 3 * 4", 14},
		{"_ = 2 * 3 + 4", 10},
		{"_ = 2 * (3 + 4)", 14},
		{"_ = 0x55 | 0xaa", 255},
		{"_ = 0x55 & 0xaa", 0},
		{"_ = 0xf ^ 1", 0xe},
		{"_ = 1 << 8", 256},
		{"_ = 256 >> 4", 16},
		{"_ = 1 > 0", 1},
		{"_ = 1 >= 0", 1},
		{"_ = 1 < 0", 0},
		{"_ = 1 <= 0", 0},
		{"_ = 2 > 3", 0},
		{"_ = 2 >= 3", 0},
		{"_ = 2 < 3", 1},
		{"_ = 2 <= 3", 1},
		{"_ = 2 == 2", 1},
		{"_ = 2 == 3", 0},
		{"_ = 2 != 3", 1},
		{"_ = 2 != 2", 0},
		{"_ = 0 || 0", 0},
		{"_ = 1 || 0", 1},
		{"_ = 0 || 1", 1},
		{"_ = 1 || 1", 1},
		{"_ = 0 && 0", 0},
		{"_ = 1 && 0", 0},
		{"_ = 0 && 1", 0},
		{"_ = 1 && 1", 1},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		as := testAssignStatement(t, tn, stmt)
		testNumberLiteral(t, tn, as.Value, tt.expected)
	}
}

func TestParseCallFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"_ = aFunc()", "AFUNC()"},
		{"_ = aFunc(   )  ", "AFUNC()"},
		{"_ = aFunc(1)", "AFUNC(1)"},
		{"_ = aFunc(1,2,3)", "AFUNC(1, 2, 3)"},
		{"_ = aFunc(1,2+3,4*5)", "AFUNC(1, 5, 20)"},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		as := testAssignStatement(t, tn, stmt)
		if as.Value.String() != tt.expected {
			t.Errorf("[%d] stmt.Value.String() is not %q. got %q", tn, tt.expected, as.Value.String())
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
	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		vs, ok := stmt.(*VariableStatement)
		if !ok {
			t.Fatalf("[%d] Statements[0] is not VariableStatement . got %T", tn, prog.Block[0])
		}
		if vs.String() != tt.expected {
			t.Errorf("[%d] stmt.String() is not %q. got %q", tn, tt.expected, vs.String())
		}
	}
}

func TestParseArrayElement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"_  = a0 [ 123 ]", "A0[123]"},
		{"_  = a1 [ 1 + 2 * 3 ]", "A1[7]"},
	}
	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		as := testAssignStatement(t, tn, stmt)
		if as.Value.String() != tt.expected {
			t.Errorf("[%d] stmt.String() is not %q. got %q", tn, tt.expected, as.Value.String())
		}
	}
}

func TestParseStringExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`_ = "abc" `, "abc"},
		{`_ = "abc" + "def" `, "abcdef"},
		{`_ = "abc" + "@" + "def" `, "abc@def"},
		// 連続する文字列リテラルの結合
		{`_ = "abc" "def"`, "abcdef"},
		{`_ = "abc" "123" "def"`, "abc123def"},
		{`_ = "abc" "123" "def"`, "abc123def"},
		{`_ = \
		"abc" \
		"123"\
		 "def"`, "abc123def"},
	}
	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		as := testAssignStatement(t, tn, stmt)
		testStringLiteral(t, tn, as.Value, tt.expected)
	}
}

func TestParseExpressionError(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		err      string
	}{
		// 0-
		{input: `_ = 1 / 0`, err: errcode.EBIN_OP_DIVZERO},
		// {input: `_ = 1 + "a"`, err: errcode.EBIN_OP_TYPE}, // 評価で検出
		// {input: `_ = "a" + 1`, err: errcode.EBIN_OP_TYPE}, // 評価で検出
		// {input: `_ = "a" * "b"`, err: errcode.EBIN_OP_TYPE}, // 評価で検出
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)

		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, l.logger)
			ename := testutil.ErrcodeNames[tt.err]
			if ename[0] == 'E' {
				continue
			}
		}

		if ec := l.logger.ErrorCount(); ec > 0 {
			t.Errorf("[%d] %d errors", tn, ec)
		}
		if len(prog.Block) == 0 {
			t.Errorf("[%d] %d statements", tn, len(prog.Block))
		}
		stmt := testAssignStatement(t, tn, prog.Block[0])
		testStringLiteral(t, tn, stmt.Value, tt.expected)
	}
}
