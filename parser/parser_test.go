package parser

import (
	"bufio"
	"fmt"
	"strings"
	"testing"
)

func TestINST0(t *testing.T) {
	input := `EXX
LDI
LDIR
LDDR
CPI
CPIR
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
RET
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
	l := NewLexer(bufio.NewReader(strings.NewReader(input)))
	ret := Parse(l)
	if ret != 0 {
		t.Errorf("Parse returns %d", ret)
	}
	expected := strings.Trim(input, " \n\t")
	text := Root.String()
	if text != expected {
		t.Errorf("program differs. exptected %d chars. got %d chars",
			len(expected), len(text))
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
`
	l := NewLexer(bufio.NewReader(strings.NewReader(input)))
	ret := Parse(l)
	if ret != 0 {
		t.Errorf("Parse returns %d", ret)
	}
	expected := strings.Trim(input, " \n\t")
	text := strings.ReplaceAll(Root.String(), "\t", " ")
	if text != expected {
		t.Errorf("program differs. exptected %d chars. got %d chars",
			len(expected), len(text))
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
	l := NewLexer(bufio.NewReader(strings.NewReader(input)))
	ret := Parse(l)
	if ret != 0 {
		t.Errorf("Parse returns %d", ret)
	}
	expected := strings.Trim(input, " \n\t")
	text := strings.ReplaceAll(Root.String(), "\t", " ")
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
	l := NewLexer(bufio.NewReader(strings.NewReader(input)))
	ret := Parse(l)
	if ret != 0 {
		t.Errorf("Parse returns %d", ret)
	}
	expected := strings.Trim(input, " \n\t")
	text := strings.ReplaceAll(Root.String(), "\t", " ")
	if text != expected {
		t.Errorf("program differs. exptected %d chars. got %d chars",
			len(expected), len(text))
		fmt.Println(text)
	}
}
