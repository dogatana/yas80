package parser

import (
	"fmt"
	"strings"
	"testing"
)

func TestConstStatement(t *testing.T) {
	input := `const abc = 123 \ def equ 456`
	expected := []struct {
		Name  string
		Value int
	}{
		{"abc", 123},
		{"def", 456},
	}
	l := newLexerForTest(input)
	prog := Parse(l)
	ec, wc, _ := l.logger.Count()

	if ec > 0 || wc > 0 {
		t.Fatalf("parsing %s returns %d errors and %d warnigs", input, ec, wc)
	}
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

func TestEnumStatement(t *testing.T) {
	input := ` test enum
 abc
def = 1  
 xyz  
  ende`
	l := newLexerForTest(input)
	prog := Parse(l)
	ec, wc, _ := l.logger.Count()
	if ec > 0 || wc > 0 {
		l.logger.Print()
		t.Fatalf("parsing %s returns %d errors and %d warnigs", input, ec, wc)
	}
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

func TestRepeatStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`repeat 1\ endr`, "REPEAT 1\nENDR"},
		{` Repeat 2 \\\\ endR`, "REPEAT 2\nENDR"},
		{` REPEAT 3 \\\\ EndR`, "REPEAT 3\nENDR"},
		{` REPEAT 4 \1\ 2\3 \ 4 \ EndR`, "REPEAT 4\n1\n2\n3\n4\nENDR"},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := Parse(l)
		ec, wc, _ := l.logger.Count()
		if ec > 0 || wc > 0 {
			l.logger.Print()
			t.Fatalf("parsing %s returns %d errors and %d warnigs", tt.input, ec, wc)
		}
		if len(prog.Statements) != 1 {
			t.Fatalf("expect 1 statements. got %d", len(prog.Statements))
		}
		stmt := prog.Statements[0]
		repeat, ok := stmt.(*RepeatStatement)
		if !ok {
			t.Errorf("prog.Statements[0] not *RepeatStatment. got %T", stmt)
		}
		text := repeat.String()
		if !strings.EqualFold(text, tt.expected) {
			t.Errorf("exptected len %d. got %d", len(tt.expected), len(text))
			fmt.Printf("expected\n%s\n", tt.expected)
			// fmt.Println([]byte(tt.expected))
			fmt.Printf("got\n%s\n", text)
			// fmt.Println([]byte(text))
		}
	}
}

func TestIfStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`if 1\ endif`,
			"IF 1\nENDIF",
		},
		{
			`if 1 \ else \ endif`,
			"IF 1\nENDIF",
		},
		{
			`if 1 \ 100 \ endif`,
			"IF 1\n100\nENDIF",
		},
		{
			`if 1 \ 100 \ else \ endif`,
			"IF 1\n100\nENDIF",
		},
		{
			`if 1 \ 100 \ else \ 200 \  endif`,
			"IF 1\n100\nELSE\n200\nENDIF",
		},
		{
			`if 1 \ 100 \ if 2 \ 200 \else \ 300 \ endif \ endif`,
			"IF 1\n100\nIF 2\n200\nELSE\n300\nENDIF\nENDIF",
		},
		{
			`if 1 \ 100 \ if 2 \ 200 \endif \ else \ 300 \ endif`,
			"IF 1\n100\nIF 2\n200\nENDIF\nELSE\n300\nENDIF",
		},
		{
			` if 1 \ 100 \ if 2 \ 200 \else \ 300 \endif \ else \ if 3 \ 400 \ else \ 500 \endif\endif`,
			"IF 1\n100\nIF 2\n200\nELSE\n300\nENDIF\nELSE\nIF 3\n400\nELSE\n500\nENDIF\nENDIF",
		},
		{
			`if 1 \ 200 \ elif 2 \ 300 \ endif`,
			"IF 1\n200\nELSE\nIF 2\n300\nENDif\nENDIF",
		},
		{
			`if 1 \ 100 \ elif 2 \ 200 \ elif 3 \ 300 \ elif 4 \ 400 \endif`,
			`IF 1\100\ELSE\IF 2\200\ELSE\IF 3\300\ELSE\IF 4\400\ENDIF\ENDIF\ENDIF\ENDIF`,
		},
		{
			`if 1 \ 100 \ elif 2 \ 200 \ elif 3 \ 300 \ elif 4 \ 400 \else\500\endif`,
			`IF 1\100\ELSE\IF 2\200\ELSE\IF 3\300\ELSE\IF 4\400\ELSE\500\ENDIF\ENDIF\ENDIF\ENDIF`,
		},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := Parse(l)
		ec, wc, _ := l.logger.Count()
		if ec > 0 || wc > 0 {
			fmt.Println("test:", tt.input)
			l.logger.Print()
			t.Fatalf("parsing %s returns %d errors and %d warnigs", tt.input, ec, wc)
		}
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

func TestFunctionStatement(t *testing.T) {
	input := `abs func x
	if x > 0
	  x
	else
	  -x
	endif
	endf
	`
	expected := `
	abs FUNC x
	IF (x > 0)
	X
	ELSE
	(-x)
	ENDIF
	ENDF`

	l := newLexerForTest(input)
	prog := Parse(l)
	ec, wc, _ := l.logger.Count()
	if ec > 0 || wc > 0 {
		l.logger.Print()
		t.Fatalf("parsing %s returns %d errors and %d warnigs", input, ec, wc)
	}
	if len(prog.Statements) != 1 {
		t.Fatalf("expect 1 statements. got %d", len(prog.Statements))
	}
	stmt := prog.Statements[0]
	fn, ok := stmt.(*FunctionStatement)
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

func TestVarStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"var abc = 1", "VAR ABC = 1"},
		{" var   xyz   =   123  + 456", "VAR XYZ = 579"},
	}

	for _, tt := range tests {

		l := newLexerForTest(tt.input)
		prog := Parse(l)
		ec, wc, _ := l.logger.Count()
		if ec > 0 || wc > 0 {
			l.logger.Print()
			t.Fatalf("parsing %s returns %d errors and %d warnigs", tt.input, ec, wc)
		}
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

func TestAsignStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"abc = 1", "ABC = 1"},
		{"def = 1 + 2 * 3", "DEF = 7"},
	}

	for _, tt := range tests {

		l := newLexerForTest(tt.input)
		prog := Parse(l)
		ec, wc, _ := l.logger.Count()
		if ec > 0 || wc > 0 {
			l.logger.Print()
			t.Fatalf("parsing %s returns %d errors and %d warnigs", tt.input, ec, wc)
		}
		if len(prog.Statements) != 1 {
			t.Fatalf("expect 1 statements. got %d", len(prog.Statements))
		}
		stmt := prog.Statements[0]
		varStmt, ok := stmt.(*AsignStatement)
		if !ok {
			t.Errorf("prog.Statements[0] not *AsignStatement. got %T", stmt)
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

func TestExitmStatement(t *testing.T) {
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
		prog := Parse(l)
		ec, wc, _ := l.logger.Count()
		if ec > 0 || wc > 0 {
			fmt.Printf("input %q\n", tt.input)
			l.logger.Print()
			t.Fatalf("%d errors and %d warnigs", ec, wc)
		}
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
			t.Errorf("NodeType not %s. got %s", tokenLiteral(int(tt.NodeType)), tokenLiteral(int(stmt.NodeType())))
		}
	}
}

func TestReturnStatement(t *testing.T) {
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
		prog := Parse(l)
		ec, wc, _ := l.logger.Count()
		if ec > 0 || wc > 0 {
			fmt.Printf("input %q\n", tt.input)
			l.logger.Print()
			t.Fatalf("%d errors and %d warnigs", ec, wc)
		}
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
			t.Errorf("NodeType not %s. got %s", tokenLiteral(int(tt.NodeType)), tokenLiteral(int(stmt.NodeType())))
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
