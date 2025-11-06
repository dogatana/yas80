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
	prog, ec, wc := Parse(l)
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
  end_enum `
	l := newLexerForTest(input)
	prog, ec, wc := Parse(l)
	if ec > 0 || wc > 0 {
		l.ErrorStore.Print()
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
		{`repeat 1\ end_repeat`, "REPEAT 1\nEND_REPEAT"},
		{` Repeat 2 \\\\ endRepeat`, "REPEAT 2\nEND_REPEAT"},
		{` REPEAT 3 \\\\ EndR`, "REPEAT 3\nEND_REPEAT"},
		{` REPEAT 4 \1\ 2\3 \ 4 \ EndR`, "REPEAT 4\n1\n2\n3\n4\nEND_REPEAT"},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog, ec, wc := Parse(l)
		if ec > 0 || wc > 0 {
			l.ErrorStore.Print()
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
		{`if 1\ end_if`, "IF 1\nEND_IF"},
		{`if 1 \ else \ end_if`, "IF 1\nELSE\nEND_IF"},
		{` if 1 \ 100 \ end_if`, "IF 1\n100\nEND_IF"},
		{` if 1 \ 100 \ else \ end_if`, "IF 1\n100\nELSE\nEND_IF"},
		{` if 1 \ 100 \ else \ 200 \  end_if`, "IF 1\n100\nELSE\n200\nEND_IF"},
		{` if 1 \ 100 \ if 2 \ 200 \else \ 300 \ end_if \ endif`, "IF 1\n100\nIF 2\n200\nELSE\n300\nEND_IF\nEND_IF"},
		{` if 1 \ 100 \ if 2 \ 200 \endif \ else \ 300 \ end_if`, "IF 1\n100\nIF 2\n200\nEND_IF\nELSE\n300\nEND_IF"},
		{
			` if 1 \ 100 \ if 2 \ 200 \else \ 300 \endif \ else \ if 3 \ 400 \ else \ 500 \endif\endif`,
			"IF 1\n100\nIF 2\n200\nELSE\n300\nEND_IF\nELSE\nIF 3\n400\nELSE\n500\nEND_IF\nEND_IF",
		},
		{`	if 1
			  200
			elif 2
			  300
			endif`,
			`IF 1
			   200
			 ELSE
			   IF 2
			     300
			   END_IF
			 END_IF`,
		},
		{
			`if 1 \ 100 \ elif 2 \ 200 \ elif 3 \ 300 \ elif 4 \ 400 \endif`,
			`IF 1\100\ELSE\IF 2\200\ELSE\IF 3\300\ELSE\IF 4\400\END_IF\END_IF\END_IF\END_IF`,
		},
		{
			`if 1 \ 100 \ elif 2 \ 200 \ elif 3 \ 300 \ elif 4 \ 400 \else\500\endif`,
			`IF 1\100\ELSE\IF 2\200\ELSE\IF 3\300\ELSE\IF 4\400\ELSE\500\END_IF\END_IF\END_IF\END_IF`,
		},
	}
	for _, tt := range tests {
		fmt.Println("test:", tt.input)
		l := newLexerForTest(tt.input)
		prog, ec, wc := Parse(l)
		if ec > 0 || wc > 0 {
			l.ErrorStore.Print()
			t.Fatalf("parsing %s returns %d errors and %d warnigs", tt.input, ec, wc)
		}
		if len(prog.Statements) != 1 {
			t.Fatalf("expect 1 statements. got %d", len(prog.Statements))
		}
		stmt := prog.Statements[0]
		repeat, ok := stmt.(*IfStatement)
		if !ok {
			t.Errorf("prog.Statements[0] not *IfStatment. got %T", stmt)
		}
		text := repeat.String()
		expected := splitTrim(tt.expected)
		if !strings.EqualFold(text, expected) {
			t.Errorf("exptected len %d. got %d", len(expected), len(text))
			fmt.Printf("expected\n%s\n", expected)
			// fmt.Println([]byte(expected))
			fmt.Printf("got\n%s\n", text)
			// fmt.Println([]byte(text))
		}
	}
}

func TestFunctionStatement(t *testing.T) {
	input := `abs function x
	if x > 0
	  x
	else
	  -x
	endif
	end_function
	`
	expected := `
	abs FUNCTION x
	IF (x > 0)
	X
	ELSE
	(-x)
	END_IF
	END_FUNCTION`

	l := newLexerForTest(input)
	prog, ec, wc := Parse(l)
	if ec > 0 || wc > 0 {
		l.ErrorStore.Print()
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
		prog, ec, wc := Parse(l)
		if ec > 0 || wc > 0 {
			l.ErrorStore.Print()
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
		prog, ec, wc := Parse(l)
		if ec > 0 || wc > 0 {
			l.ErrorStore.Print()
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
