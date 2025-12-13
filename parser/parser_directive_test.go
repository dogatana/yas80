package parser

import (
	"fmt"
	"strings"
	"testing"
)

func TestParseConstStatement(t *testing.T) {
	input := `const abc = 123 \ def equ 456`
	expected := []struct {
		Name  string
		Value int
	}{
		{"ABC", 123},
		{"DEF", 456},
	}
	l := newLexerForTest(input)
	prog := ParseForTest(t, l, input)

	if len(prog.Statements) != 2 {
		t.Fatalf("expect 2 statements. got %d", len(prog.Statements))
	}
	for i, stmt := range prog.Statements {
		cs := stmt.(*ConstStatement)
		if cs.Name.Name != expected[i].Name {
			t.Errorf("expected Name %q. got %q", expected[i].Name, cs.Name)
		}
		v, ok := cs.Value.(*NumberLiteral)
		if !ok {
			t.Errorf("Value is not NumberLiteral. got %t", cs.Value)
		}
		if v.Value != expected[i].Value {
			t.Errorf("Value is not %d. got %d", expected[i].Value, v.Value)
		}
	}
}

func TestParseEnumStatement(t *testing.T) {
	input := ` test enum
 abc
def = 1  
 xyz  
  ende`
	l := newLexerForTest(input)
	prog := ParseForTest(t, l, input)

	if len(prog.Statements) != 1 {
		t.Fatalf("expect 1 statements. got %d", len(prog.Statements))
	}
	stmt := prog.Statements[0]
	enum, ok := stmt.(*EnumStatement)
	if !ok {
		t.Errorf("prog.Statements[0] not *EnumStatement. got %T", stmt)
	}
	srcText := splitTrim(input)
	text := enum.String()
	if !strings.EqualFold(text, srcText) {
		t.Errorf("expected %d chars. got %d chars", len(srcText), len(text))
		fmt.Printf("expected\n%s\n", srcText)
		fmt.Println([]byte(srcText))
		fmt.Printf("got\n%s\n", text)
		fmt.Println([]byte(text))
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
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)

		if len(prog.Statements) != 1 {
			t.Fatalf("expect 1 statements. got %d", len(prog.Statements))
		}
		stmt := prog.Statements[0]
		repeat, ok := stmt.(*ReptStatement)
		if !ok {
			t.Errorf("prog.Statements[0] not *RepeatStatment. got %T", stmt)
		}
		count := repeat.MaxCount.(*NumberLiteral).Value
		if count != tt.count {
			t.Errorf("exptected count %d. got %d", tt.count, count)
		}
		if len(repeat.Block.Block) != tt.stmtCount {
			t.Errorf("must have %d statements. got %d", tt.stmtCount, len(repeat.Block.Block))
		}
	}
}

func TestParseIfStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
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
		{
			`if 1 \ a=100 \ elif 2 \ a=200 \ elif 3 \ a=300 \ elif 4 \ a=400 \else\a=500\endif`,
			`IF 1\A = 100\ELSE\IF 2\A = 200\ELSE\IF 3\A = 300\ELSE\IF 4\A = 400\ELSE\A = 500\ENDIF\ENDIF\ENDIF\ENDIF`,
		},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)

		if len(prog.Statements) != 1 {
			fmt.Println("test:", tt.input)
			t.Fatalf("expect 1 statements. got %d", len(prog.Statements))
		}
		stmt := prog.Statements[0]
		repeat, ok := stmt.(*IfStatement)
		if !ok {
			fmt.Println("test:", tt.input)
			t.Errorf("prog.Statements[0] not *IfStatment. got %T", stmt)
		}
		text := repeat.String()
		expected := splitTrim(tt.expected)
		if !strings.EqualFold(text, expected) {
			fmt.Println("test:", tt.input)
			t.Errorf("exptected len %d. got %d", len(expected), len(text))
			fmt.Printf("expected\n%s\n", expected)
			// fmt.Println([]byte(expected))
			fmt.Printf("got\n%s\n", text)
			// fmt.Println([]byte(text))
		}
	}
}

func TestParseFuncStatement(t *testing.T) {
	input := `abs func x
	if x > 0
	  return x
	else
	  return -x
	endif
	endf
	`
	expected := `
	ABS FUNC X
	IF (x > 0)
	RETURN X
	ELSE
	RETURN (-X)
	ENDIF
	ENDF`

	l := newLexerForTest(input)
	prog := ParseForTest(t, l, input)

	if len(prog.Statements) != 1 {
		t.Fatalf("expect 1 statements. got %d", len(prog.Statements))
	}
	stmt := prog.Statements[0]
	fn, ok := stmt.(*FuncStatement)
	if !ok {
		t.Errorf("prog.Statements[0] not *FunctionStatment. got %T", stmt)
	}

	text := fn.String()
	expectedText := splitTrim(expected)
	if !strings.EqualFold(text, expectedText) {
		t.Errorf("exptected len %d. got %d", len(expectedText), len(text))
		fmt.Printf("expected\n%s\n", expectedText)
		fmt.Println([]byte(expectedText))
		fmt.Printf("got\n%s\n", text)
		fmt.Println([]byte(text))
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

	for _, tt := range tests {

		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)

		if len(prog.Statements) != 1 {
			t.Fatalf("expect 1 statements. got %d", len(prog.Statements))
		}
		stmt := prog.Statements[0]
		varStmt, ok := stmt.(*VariableStatement)
		if !ok {
			t.Errorf("prog.Statements[0] not *VariableStatement. got %T", stmt)
		}

		text := varStmt.String()
		expectedText := splitTrim(tt.expected)
		if !strings.EqualFold(text, expectedText) {
			t.Errorf("exptected len %d. got %d", len(expectedText), len(text))
			fmt.Printf("expected\n%s\n", expectedText)
			fmt.Println([]byte(expectedText))
			fmt.Printf("got\n%s\n", text)
			fmt.Println([]byte(text))
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

	for _, tt := range tests {

		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)

		if len(prog.Statements) != 1 {
			t.Fatalf("expect 1 statements. got %d", len(prog.Statements))
		}
		stmt := prog.Statements[0]
		varStmt, ok := stmt.(*AssignStatement)
		if !ok {
			t.Errorf("prog.Statements[0] not *AssignStatement. got %T", stmt)
		}

		text := varStmt.String()
		expectedText := splitTrim(tt.expected)
		if !strings.EqualFold(text, expectedText) {
			t.Errorf("exptected len %d. got %d", len(expectedText), len(text))
			fmt.Printf("expected\n%s\n", expectedText)
			fmt.Println([]byte(expectedText))
			fmt.Printf("got\n%s\n", text)
			fmt.Println([]byte(text))
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

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)

		if len(prog.Statements) != 1 {
			fmt.Printf("input %q\n", tt.input)
			t.Fatalf("expect 1 statements. got %d", len(prog.Statements))
		}
		stmt, ok := prog.Statements[0].(*ExitmStatement)
		if !ok {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("not ExitmStatement. got %T", stmt)
		}

		if stmt.NodeType() != tt.NodeType {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("NodeType not %s. got %s", TokenLiteral(int(tt.NodeType)), TokenLiteral(int(stmt.NodeType())))
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

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)

		if len(prog.Statements) != 1 {
			fmt.Printf("input %q\n", tt.input)
			t.Fatalf("expect 1 statements. got %d", len(prog.Statements))
		}
		stmt, ok := prog.Statements[0].(*ReturnStatement)
		if !ok {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("not ExitmStatement. got %T", stmt)
		}
		if stmt.NodeType() != tt.NodeType {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("NodeType not %s. got %s", TokenLiteral(int(tt.NodeType)), TokenLiteral(int(stmt.NodeType())))
		}
		switch v := tt.expected.(type) {
		case nil:
			if v != nil {
				fmt.Printf("input %q\n", tt.input)
				t.Errorf("ReturnStatement.Value not %v. got %s", v, stmt.Value)
			}
		case int:
			result := stmt.Value.(*NumberLiteral).Value
			if result != v {
				fmt.Printf("input %q\n", tt.input)
				t.Errorf("ReturnStatement.Value not %d. got %d", v, result)
			}
		case string:
			result := stmt.Value.(*StringLiteral).Value
			if result != v {
				fmt.Printf("input %q\n", tt.input)
				t.Errorf("ReturnStatement.Value not %s. got %s", v, result)
			}
		default:
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("ReturnStatement.Value is unpexcted type %#v", v)
		}
	}
}
