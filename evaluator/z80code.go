package evaluator

import (
	"yas80/object"
	"yas80/parser"
)

var Z80CodeTable0 = map[int]object.CodeObject{
	parser.Z80_INST_EXX:  {Code: []byte{0xd9}, CZ80: 4},
	parser.Z80_INST_LDI:  {Code: []byte{0xed, 0xa0}, CZ80: 16},
	parser.Z80_INST_LDIR: {Code: []byte{0xed, 0xb0}, CZ80: 21},
	parser.Z80_INST_LDD:  {Code: []byte{0xed, 0xa8}, CZ80: 16},
	parser.Z80_INST_LDDR: {Code: []byte{0xed, 0xb8}, CZ80: 21},
	parser.Z80_INST_CPI:  {Code: []byte{0xed, 0xa1}, CZ80: 16},
	parser.Z80_INST_CPIR: {Code: []byte{0xed, 0xb1}, CZ80: 21},
	parser.Z80_INST_CPD:  {Code: []byte{0xed, 0xa9}, CZ80: 16},
	parser.Z80_INST_CPDR: {Code: []byte{0xed, 0xb9}, CZ80: 21},
	parser.Z80_INST_DAA:  {Code: []byte{0x27}, CZ80: 4},
	parser.Z80_INST_CPL:  {Code: []byte{0x2f}, CZ80: 4},
	parser.Z80_INST_NEG:  {Code: []byte{0xed, 0x44}, CZ80: 8},
	parser.Z80_INST_CCF:  {Code: []byte{0x3f}, CZ80: 4},
	parser.Z80_INST_SCF:  {Code: []byte{0x37}, CZ80: 4},
	parser.Z80_INST_NOP:  {Code: []byte{0x00}, CZ80: 4},
	parser.Z80_INST_HALT: {Code: []byte{0x76}, CZ80: 4},
	parser.Z80_INST_DI:   {Code: []byte{0xf3}, CZ80: 4},
	parser.Z80_INST_EI:   {Code: []byte{0xfb}, CZ80: 4},
	parser.Z80_INST_RLCA: {Code: []byte{0x07}, CZ80: 4},
	parser.Z80_INST_RLA:  {Code: []byte{0x17}, CZ80: 4},
	parser.Z80_INST_RRCA: {Code: []byte{0x0f}, CZ80: 4},
	parser.Z80_INST_RRA:  {Code: []byte{0x1f}, CZ80: 4},
	parser.Z80_INST_RLD:  {Code: []byte{0xed, 0x6f}, CZ80: 18},
	parser.Z80_INST_RRD:  {Code: []byte{0xed, 0x67}, CZ80: 18},
	parser.Z80_INST_RETI: {Code: []byte{0xed, 0x4d}, CZ80: 14},
	parser.Z80_INST_RETN: {Code: []byte{0xed, 0x45}, CZ80: 14},
	parser.Z80_INST_INI:  {Code: []byte{0xed, 0xa2}, CZ80: 16},
	parser.Z80_INST_INIR: {Code: []byte{0xed, 0xb2}, CZ80: 21},
	parser.Z80_INST_IND:  {Code: []byte{0xed, 0xaa}, CZ80: 16},
	parser.Z80_INST_INDR: {Code: []byte{0xed, 0xba}, CZ80: 21},
	parser.Z80_INST_OUTI: {Code: []byte{0xed, 0xa3}, CZ80: 16},
	parser.Z80_INST_OTIR: {Code: []byte{0xed, 0xb3}, CZ80: 21},
	parser.Z80_INST_OUTD: {Code: []byte{0xed, 0xab}, CZ80: 16},
	parser.Z80_INST_OTDR: {Code: []byte{0xed, 0xbb}, CZ80: 21},
}

var Z80Reg8Index map[int]byte = map[int]byte{
	parser.Z80_REG_B: 0,
	parser.Z80_REG_C: 1,
	parser.Z80_REG_D: 2,
	parser.Z80_REG_E: 3,
	parser.Z80_REG_H: 4,
	parser.Z80_REG_L: 5,
	parser.Z80_REG_A: 7,
}

var Z80Reg16IndexSP map[int]byte = map[int]byte{
	parser.Z80_REG_BC: 0,
	parser.Z80_REG_DE: 1,
	parser.Z80_REG_HL: 2,
	parser.Z80_REG_SP: 3,
}

var Z80Reg16IndexSPIXY map[int]byte = map[int]byte{
	parser.Z80_REG_BC: 0,
	parser.Z80_REG_DE: 1,
	parser.Z80_REG_HL: 2,
	parser.Z80_REG_IX: 2,
	parser.Z80_REG_IY: 2,
	parser.Z80_REG_SP: 3,
}

var Z80Reg16IndexAF map[int]byte = map[int]byte{
	parser.Z80_REG_BC: 0,
	parser.Z80_REG_DE: 1,
	parser.Z80_REG_HL: 2,
	parser.Z80_REG_AF: 3,
}

var Z80FlagIndex map[int]byte = map[int]byte{
	parser.Z80_FLAG_NZ: 0,
	parser.Z80_FLAG_Z:  1,
	parser.Z80_FLAG_NC: 2,
	parser.Z80_REG_C:   3, // lexer では C を Register トークンとして処理するため
	parser.Z80_FLAG_C:  3, // リテラル CY を指定した場合
	parser.Z80_FLAG_PO: 4,
	parser.Z80_FLAG_PE: 5,
	parser.Z80_FLAG_P:  6,
	parser.Z80_FLAG_M:  7,
}
