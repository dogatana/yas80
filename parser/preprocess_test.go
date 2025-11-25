package parser

import (
	"fmt"
	"strings"
	"testing"
)

func TestMacroExpand(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			`
		ld_macro macro left, right
		ld left, right
		endm
		
		ld_macro a, 1
		ld_macro a, 1 + 3
		ld_macro a, b
		ld_macro hl, $100`,
			`LD A, 1
LD A, 4
LD A, B
LD HL, 256`},
	}

	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog := ParseForTest(t, l, tt.input)

		if len(prog.Statements) != 5 {
			t.Fatalf("%d statements. not 5", len(prog.Statements))
		}
		if !strings.HasSuffix(prog.String(), tt.expected) {
			l.logger.Print()
			fmt.Println("got     ", prog.String())
			fmt.Println("expected", tt.expected)
			t.Errorf("prog.String() is not expecged one")
		}
	}
}
