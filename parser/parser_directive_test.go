package parser

import (
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
