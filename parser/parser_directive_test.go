package parser

import (
	"fmt"
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
	input := `test ENUM
abc
def = 1
xyz
END_ENUM`
	l := newLexerForTest(input)
	prog, ec, wc := Parse(l)
	if ec > 0 || wc > 0 {
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
	text := enum.String()
	if text != input {
		t.Errorf("expected %d chars. got %d chars", len(input), len(text))
		fmt.Printf("expected\n%s\n", input)
		fmt.Println([]byte(input))
		fmt.Printf("got\n%s\n", text)
		fmt.Println([]byte(text))
	}
}
