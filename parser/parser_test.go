package parser

import (
	"fmt"
	"strings"
	"testing"
)

func TestINST0(t *testing.T) {
	input := `EXX
LDI
LDIR
LDD
LDDR
CPI
CPIR
CPD
CPDR
DAA
CPL
NEG
CCF
SCF
NOP
HALT
DI
EI
RLCA
RLA
RRCA
RRA
RLD
RRD
RETI
RETN
INI
INIR
INDR
OUTI
OTIR
OUTD
OTDR
`
	l := newLexerForTest(input)
	prog, ec, wc := Parse(l)
	if ec > 0 || wc > 0 {
		for _, e := range l.logger.Errors {
			fmt.Println(e.String())
		}
		for _, e := range l.logger.Warnings {
			fmt.Println(e.String())
		}
		t.Fatalf("parsing %s returns %d errors and %d warnigs", input, ec, wc)
	}
	expected := strings.Trim(input, " \n\t")
	text := prog.String()
	if text != expected {
		t.Errorf("program differs. exptected %d chars. got %d chars",
			len(expected), len(text))
		fmt.Println(text)
	}
}

func TestINST1(t *testing.T) {
	input := `
PUSH 1
POP 2
SUB 3
AND 4
OR 5
CP 6
INC 7
DEC 8
IM 9
RLC 10
RL 11
RR 12
SLA 13
SRA 14
SRL 15
DJNZ 16
RST 17
RET
RET 18
`
	l := newLexerForTest(input)
	prog, ec, wc := Parse(l)
	if ec > 0 || wc > 0 {
		for _, e := range l.logger.Errors {
			fmt.Println(e.String())
		}
		for _, e := range l.logger.Warnings {
			fmt.Println(e.String())
		}
		t.Fatalf("Parser returns %d errors and %d warnigs", ec, wc)
	}
	expected := strings.Trim(input, " \n\t")
	text := strings.ReplaceAll(prog.String(), "\t", " ")
	if text != expected {
		t.Errorf("program differs. exptected %d chars. got %d chars",
			len(expected), len(text))
		fmt.Println(text)
	}
}

func TestINST2(t *testing.T) {
	input := `
EX AF, AF'
ADD A, 1
ADC A, 2
SBC A, 3
BIT 0, B
SET 1, C
RES 2, D
JP 100
JP C, 100
JR 200
JR NC, 200
CALL 300
CALL NC, 300
IN A, (1)
IN B, (C)
OUT (2), A
OUT (C), B
`
	l := newLexerForTest(input)
	prog, ec, wc := Parse(l)
	if ec > 0 || wc > 0 {
		t.Fatalf("parsing %s returns %d errors and %d warnigs", input, ec, wc)
	}
	expected := strings.Trim(input, " \n\t")
	text := strings.ReplaceAll(prog.String(), "\t", " ")
	if text != expected {
		t.Errorf("program differs. exptected %d chars. got %d chars",
			len(expected), len(text))
	}
}

func TestIndirect(t *testing.T) {
	input := `
LD A, (HL)
LD (HL), A
LD A, (IX + 1)
LD (IX + 1), A
`
	l := newLexerForTest(input)
	prog, ec, wc := Parse(l)
	if ec > 0 || wc > 0 {
		t.Fatalf("parsing %s returns %d errors and %d warnigs", input, ec, wc)
	}
	expected := strings.Trim(input, " \n\t")
	text := strings.ReplaceAll(prog.String(), "\t", " ")
	if text != expected {
		t.Errorf("program differs. exptected %d chars. got %d chars",
			len(expected), len(text))
		fmt.Println(text)
	}
}

func TestLableStatement(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{" abc: ", "abc"},
		{" abc : ", "abc"},
		{"abc:", "abc"},
		{" abc:ld a,a ", "abc"},
		{" abc :ld a, a", "abc"},
		{"abc: ld a, a", "abc"},
		{" .def: ", ".def"},
		{" .def : ", ".def"},
		{".def:", ".def"},
		{" .def:ld a,a ", ".def"},
		{" .def :ld a, a", ".def"},
		{".def: ld a, a", ".def"},
	}
	for _, tt := range tests {
		l := newLexerForTest(tt.input)
		prog, ec, wc := Parse(l)
		if ec > 0 || wc > 0 {
			t.Fatalf("parsing %s. returns %d errors and %d warnigs", tt.input, ec, wc)
		}
		if len(prog.Statements) == 0 {
			t.Fatalf("parsing %s. statements empty", tt.input)
		}
		stmt, ok := prog.Statements[0].(*LabelStatement)
		if !ok {
			t.Errorf("parsing %s. prog.Statemtes[0] is not LabelStatement. got %T", tt.input, prog.Statements[0])
		}
		name := stmt.Value.(*Label).Name
		if name != tt.expected {
			t.Errorf("parsing %s. Label.Name is not %q. got %q", tt.input, tt.expected, name)
		}
	}
}

func TestDotIdent(t *testing.T) {
	input := "abc.def"

	l := newLexerForTest(input)
	prog, ec, wc := Parse(l)
	if ec > 0 || wc > 0 {
		t.Fatalf("parsing %s. returns %d errors and %d warnigs", input, ec, wc)
	}
	if len(prog.Statements) == 0 {
		t.Fatalf("parsing %s. statements empty", input)
	}
	stmt, ok := prog.Statements[0].(*ExpressionStatement)
	if !ok {
		t.Errorf("parsing %s. prog.Statemtes[0] is not ExpressionStatement. got %T", input, prog.Statements[0])
	}
	ident, ok := stmt.Value.(*DotIdent)
	if !ok {
		t.Errorf("parsing %s. not Expression got %T", input, stmt.Value)
	}
	if ident.Left != "ABC" || ident.Right != "DEF" {
		t.Errorf("parsing %s. not ABC.DEF. got %q", input, ident.String())
	}
}
