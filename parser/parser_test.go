package parser

import (
	"bufio"
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
