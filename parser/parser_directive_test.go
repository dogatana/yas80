package parser

import (
	"strings"
	"testing"
	"yas80/errcode"
	"yas80/internal/testutil"
)

func TestParseConstStatement(t *testing.T) {
	tests := []struct {
		input    string
		name     string
		expected int
	}{
		{`const abc = 123`, "ABC", 123},
		{`def equ 456`, "DEF", 456},
	}
	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		cs := stmt.(*ConstStatement)
		ident, ok := cs.Name.(*Ident)
		if !ok {
			t.Fatalf("[%d] Name must be Ident. got %T", tn, cs.Name)
		}
		if ident.Name != tt.name {
			t.Errorf("[%d] expected Name %q. got %q", tn, tt.name, ident.Name)
		}
		testNumberLiteral(t, tn, cs.Value, tt.expected)
	}
}

func TestParseEnumStatement(t *testing.T) {
	tests := []struct {
		input string
	}{
		{
			` test enum
			abc
			def = 1  
			xyz  
			ende`,
		},
	}
	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, -1)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		enum, ok := stmt.(*EnumStatement)
		if !ok {
			t.Errorf("[%d] prog.Block[0] not *EnumStatement. got %T", tn, stmt)
		}
		srcText := splitTrim(tt.input)
		text := enum.String()
		if !strings.EqualFold(text, srcText) {
			t.Errorf("[%d] expected %d chars. got %d chars", tn, len(srcText), len(text))
		}
	}
}

func TestParseReptStatement(t *testing.T) {
	tests := []struct {
		input     string
		count     int
		stmtCount int
	}{
		{`rept 1\ endr`, 1, 0},
		{` Rept 2 \ a = 1 \ endR`, 2, 1},
		{` REPT 3 \ a = 1 \ a = 2\ endr`, 3, 2},
	}
	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		repeat, ok := stmt.(*ReptStatement)
		if !ok {
			t.Errorf("[%d] prog.Block[0] not *RepeatStatment. got %T", tn, stmt)
		}
		count := repeat.MaxCount.(*NumberLiteral).Value
		if count != tt.count {
			t.Errorf("[%d] exptected count %d. got %d", tn, tt.count, count)
		}
		if len(repeat.Block.Block) != tt.stmtCount {
			t.Errorf("[%d] must have %d statements. got %d", tn, tt.stmtCount, len(repeat.Block.Block))
		}
	}
}

func TestParseIfStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// 0-
		{
			`if 1\ endif`,
			`IF 1\ELSE\ENDIF`,
		},
		{
			`if 1 \ else \ endif`,
			`IF 1\ELSE\ENDIF`,
		},
		{
			`if 1 \a = 100 \ endif`,
			`IF 1\A = 100\ELSE\ENDIF`,
		},
		{
			`if 1 \ a=100 \ else \ endif`,
			`IF 1\A = 100\ELSE\ENDIF`,
		},
		{
			`if 1 \ a=100 \ else \ a=200 \  endif`,
			`IF 1\A = 100\ELSE\A = 200\ENDIF`,
		},
		// 5-
		{
			`if 1 \ a=100 \ if 2 \ a=200 \else \ a=300 \ endif \ endif`,
			`IF 1\A = 100\IF 2\A = 200\ELSE\A = 300\ENDIF\ELSE\ENDIF`,
		},
		{
			`if 1 \ a=100 \ if 2 \ a=200 \endif \ else \ a=300 \ endif`,
			`IF 1\A = 100\IF 2\A = 200\ELSE\ENDIF\ELSE\A = 300\ENDIF`,
		},
		{
			` if 1 \ a=100 \ if 2 \ a=200 \else \ a=300 \endif \ else \ if 3 \ a=400 \ else \ a=500 \endif\endif`,
			`IF 1\A = 100\IF 2\A = 200\ELSE\A = 300\ENDIF\ELSE\IF 3\A = 400\ELSE\A = 500\ENDIF\ENDIF`,
		},
		{
			`if 1 \ a=200 \ elif 2 \ a=300 \ endif`,
			`IF 1\A = 200\ELSE\IF 2\A = 300\ELSE\ENDif\ENDIF`,
		},
		{
			`if 1 \ a=100 \ elif 2 \ a=200 \ elif 3 \ a=300 \ elif 4 \ a=400 \endif`,
			`IF 1\A = 100\ELSE\IF 2\A = 200\ELSE\IF 3\A = 300\ELSE\IF 4\A = 400\ELSE\ENDIF\ENDIF\ENDIF\ENDIF`,
		},
		// 10-
		{
			`if 1 \ a=100 \ elif 2 \ a=200 \ elif 3 \ a=300 \ elif 4 \ a=400 \else\a=500\endif`,
			`IF 1\A = 100\ELSE\IF 2\A = 200\ELSE\IF 3\A = 300\ELSE\IF 4\A = 400\ELSE\A = 500\ENDIF\ENDIF\ENDIF\ENDIF`,
		},
	}
	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		repeat, ok := stmt.(*IfStatement)
		if !ok {
			t.Errorf("[%d] prog.Block[0] not *IfStatment. got %T", tn, stmt)
		}
		text := repeat.String()
		expected := splitTrim(tt.expected)
		if !strings.EqualFold(text, expected) {
			t.Errorf("[%d] exptected len %d. got %d", tn, len(expected), len(text))
		}
	}
}

func TestParseFuncStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`abs func x
			if x > 0
			return x
			else
			return -x
			endif
			endf
		`,
			`ABS FUNC X
			IF (x > 0)
			RETURN X
			ELSE
			RETURN (-X)
			ENDIF
			ENDF
		`},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		fn, ok := stmt.(*FuncStatement)
		if !ok {
			t.Errorf("[%d] prog.Block[0] not *FunctionStatment. got %T", tn, stmt)
		}

		text := fn.String()
		expectedText := splitTrim(tt.expected)
		if !strings.EqualFold(text, expectedText) {
			t.Errorf("[%d] exptected len %d. got %d", tn, len(expectedText), len(text))
		}
	}
}

func TestParseVarStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"var abc = 1", "VAR ABC = 1"},
		{" var   xyz   =   123  + 456", "VAR XYZ = 579"},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		varStmt, ok := stmt.(*VariableStatement)
		if !ok {
			t.Errorf("[%d] prog.Block[0] not *VariableStatement. got %T", tn, stmt)
		}

		text := varStmt.String()
		expectedText := splitTrim(tt.expected)
		if !strings.EqualFold(text, expectedText) {
			t.Errorf("[%d] exptected len %d. got %d", tn, len(expectedText), len(text))
		}
	}
}

func TestParseAssignStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abc = 1", "ABC = 1"},
		{"def = 1 + 2 * 3", "DEF = 7"},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		varStmt, ok := stmt.(*AssignStatement)
		if !ok {
			t.Errorf("[%d] prog.Block[0] not *AssignStatement. got %T", tn, stmt)
		}

		text := varStmt.String()
		expectedText := splitTrim(tt.expected)
		if !strings.EqualFold(text, expectedText) {
			t.Errorf("[%d] exptected len %d. got %d", tn, len(expectedText), len(text))
		}
	}
}

func TestParseExitmStatement(t *testing.T) {
	tests := []struct {
		input    string
		NodeType NodeType
		literal  string
	}{
		{"exitm", NODE_EXITM_STMT, "EXITM"},
		{"ExitM", NODE_EXITM_STMT, "EXITM"},
		{"EXITM", NODE_EXITM_STMT, "EXITM"},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		if stmt.NodeType() != tt.NodeType {
			t.Errorf("[%d] NodeType not %s. got %s", tn, TokenLiteral(int(tt.NodeType)), TokenLiteral(int(stmt.NodeType())))
		}
	}
}

func TestParseReturnStatement(t *testing.T) {
	tests := []struct {
		input    string
		NodeType NodeType
		expected any
		literal  string
	}{
		{"return", NODE_RETURN_STMT, nil, "RETURN"},
		{"return 1", NODE_RETURN_STMT, 1, "RETURN"},
		{`RETURN "abc"`, NODE_RETURN_STMT, "abc", "RETURN"},
	}

	for tn, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tn)
		stmt := progHasOnlyOneStatement(t, tn, prog)

		if stmt.NodeType() != tt.NodeType {
			t.Errorf("[%d] NodeType not %s. got %s", tn, TokenLiteral(int(tt.NodeType)), TokenLiteral(int(stmt.NodeType())))
		}

		rs := stmt.(*ReturnStatement)

		switch v := tt.expected.(type) {
		case nil:
			if v != nil {
				t.Errorf("[%d] ReturnStatement.Value not %v. got %s", tn, v, rs.Value)
			}
		case int:
			result := rs.Value.(*NumberLiteral).Value
			if result != v {
				t.Errorf("[%d] ReturnStatement.Value not %d. got %d", tn, v, result)
			}
		case string:
			result := rs.Value.(*StringLiteral).Value
			if result != v {
				t.Errorf("[%d] ReturnStatement.Value not %s. got %s", tn, v, result)
			}
		default:
			t.Errorf("[%d] ReturnStatement.Value is unpexcted type %#v", tn, v)
		}
	}
}

func TestParseDataStatement(t *testing.T) {
	tests := []struct {
		input    string
		NodeType NodeType
		size     int
		length   int
		err      string
	}{
		{"db 1", NODE_DATA_STMT, 1, 1, ""},
		{"db 1,2,3", NODE_DATA_STMT, 1, 3, ""},
		{"dw 1", NODE_DATA_STMT, 2, 1, ""},
		{"dw 1,2,3", NODE_DATA_STMT, 2, 3, ""},
		{"dd 1", NODE_DATA_STMT, 0, 1, ""},
		{"dd 1,2,3", NODE_DATA_STMT, 0, 3, ""},
		{input: "db", err: errcode.ESYNTAX},
		{input: "dw", err: errcode.ESYNTAX},
		{input: "dd", err: errcode.ESYNTAX},
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
		if len(l.logger.Errors) > 0 {
			t.Fatalf("[%d] %d errors", tn, len(l.logger.Errors))
		}

		stmt := progHasOnlyOneStatement(t, tn, prog)

		if stmt.NodeType() != tt.NodeType {
			t.Errorf("[%d] NodeType not %s. got %s", tn, TokenLiteral(int(tt.NodeType)), TokenLiteral(int(stmt.NodeType())))
		}

		ds := stmt.(*DataStatement)
		if ds.Size != tt.size {
			t.Errorf("[%d] Size not %d. got %d", tn, tt.size, ds.Size)
		}
		if len(ds.Values) != tt.length {
			t.Errorf("[%d] Data Length not %d. got %d", tn, tt.length, len(ds.Values))
		}
	}
}

func TestParseDataStoreStatement(t *testing.T) {
	tests := []struct {
		input    string
		NodeType NodeType
		size     int
		count    int
		filler   int
		err      string
	}{
		// 0-
		{"ds 0", NODE_DATA_STORE_STMT, 1, 0, -1, ""},
		{"ds 1", NODE_DATA_STORE_STMT, 1, 1, -1, ""},
		{"ds 1, 255", NODE_DATA_STORE_STMT, 1, 1, 255, ""},
		{"dsb 0", NODE_DATA_STORE_STMT, 1, 0, -1, ""},
		{"dsb 1", NODE_DATA_STORE_STMT, 1, 1, -1, ""},
		// 5-
		{"dsb 1, 255", NODE_DATA_STORE_STMT, 1, 1, 255, ""},
		{"dsw 0", NODE_DATA_STORE_STMT, 2, 0, -1, ""},
		{"dsw 1", NODE_DATA_STORE_STMT, 2, 1, -1, ""},
		{"dsw 1, 65535", NODE_DATA_STORE_STMT, 2, 1, 65535, ""},
		{input: "ds", err: errcode.ESYNTAX},
		// 10-
		{input: "ds 1,2,3", err: errcode.ESYNTAX},
		{input: "dsb", err: errcode.ESYNTAX},
		{input: "dsb 1,2,3", err: errcode.ESYNTAX},
		{input: "dsw", err: errcode.ESYNTAX},
		{input: "dsw 1,2,3", err: errcode.ESYNTAX},
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
		if len(l.logger.Errors) > 0 {
			t.Fatalf("[%d] %d errors", tn, len(l.logger.Errors))
		}

		stmt := progHasOnlyOneStatement(t, tn, prog)

		ds, ok := stmt.(*DataStoreStatement)
		if !ok {
			t.Errorf("[%d] not DataStoreStatement. got %T", tn, ds)
		}
		if ds.NodeType() != tt.NodeType {
			t.Errorf("[%d] NodeType not %s. got %s", tn, TokenLiteral(int(tt.NodeType)), TokenLiteral(int(ds.NodeType())))
		}
		if ds.Size != tt.size {
			t.Errorf("[%d] Size not %d. got %d", tn, tt.size, ds.Size)
		}
		testNumberLiteral(t, tn, ds.Count, tt.count)
		if tt.filler >= 0 {
			testNumberLiteral(t, tn, ds.FillValue, tt.filler)
		}
	}
}
