package parser

import (
	"strings"
	"testing"
)

func TestParseINST0(t *testing.T) {
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
	prog := ParseForTest(t, l, -1)

	expected := strings.Split(strings.Trim(input, " \n"), "\n")
	for i, node := range prog.Statements {
		inst, ok := node.(*Z80Instruction)
		if !ok {
			t.Errorf("not Z80Instuction. got %T", node)
		}
		if inst.InstType != Z80_INST0 {
			t.Errorf("not Z80_INST0. got %s", TokenLiteral(inst.InstType))
		}
		if TokenLiteral(inst.Opcode) != expected[i] {
			t.Errorf("not %s. got %s", expected[i], TokenLiteral(inst.Opcode))
		}
	}
}

func TestParseINST1(t *testing.T) {
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
	prog := ParseForTest(t, l, -1)

	expected := []string{}
	for _, line := range strings.Split(strings.Trim(input, " \n"), "\n") {
		words := strings.Split(line, " ")
		expected = append(expected, words[0])
	}

	for i, node := range prog.Statements {
		inst, ok := node.(*Z80Instruction)
		if !ok {
			t.Errorf("not Z80Instuction. got %T", node)
		}
		if inst.InstType != Z80_INST1 {
			t.Errorf("not Z80_INST1. got %s", TokenLiteral(inst.InstType))
		}
		if TokenLiteral(inst.Opcode) != expected[i] {
			t.Errorf("not %s. got %s", expected[i], TokenLiteral(inst.Opcode))
		}
	}
}

func TestParseINST2(t *testing.T) {
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
	prog := ParseForTest(t, l, -1)

	expected := []string{}
	for _, line := range strings.Split(strings.Trim(input, " \n"), "\n") {
		words := strings.Split(line, " ")
		expected = append(expected, words[0])
	}
	for i, node := range prog.Statements {
		inst, ok := node.(*Z80Instruction)
		if !ok {
			t.Errorf("not Z80Instuction. got %T", node)
		}
		if inst.InstType != Z80_INST2 {
			t.Errorf("not Z80_INST2. got %s", TokenLiteral(inst.InstType))
		}
		if TokenLiteral(inst.Opcode) != expected[i] {
			t.Errorf("not %s. got %s", expected[i], TokenLiteral(inst.Opcode))
		}
	}
}

func TestParseIndirect(t *testing.T) {
	input := `
LD A, (HL)
LD (HL), A
LD A, (IX-1)
LD (IX+1), A
LD A, (100)
LD (100), A
`
	l := newLexerForTest(input)
	prog := ParseForTest(t, l, -1)

	expected := strings.Trim(input, " \n\t")
	text := strings.ReplaceAll(prog.String(), "\t", " ")
	if text != expected {
		t.Errorf("program differs. exptected %d chars. got %d chars",
			len(expected), len(text))
	}
}
