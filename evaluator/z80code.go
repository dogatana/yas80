package evaluator

import "yas80/parser"

type Z80Code struct {
	Bytes  []byte
	Cycle  int
	Offset int
	Value  int
}

var Z80CodeTable0 map[int]*Z80Code = map[int]*Z80Code{
	parser.Z80_INST_EXX:  {Bytes: []byte{0xd9}, Cycle: 0},
	parser.Z80_INST_LDI:  {Bytes: []byte{0xed, 0xa0}, Cycle: 0},
	parser.Z80_INST_LDIR: {Bytes: []byte{0xed, 0xb0}, Cycle: 0},
	parser.Z80_INST_LDDR: {Bytes: []byte{0xed, 0xb8}, Cycle: 0},
	parser.Z80_INST_CPI:  {Bytes: []byte{0xed, 0xa1}, Cycle: 0},
	parser.Z80_INST_CPIR: {Bytes: []byte{0xed, 0xb1}, Cycle: 0},
	parser.Z80_INST_CPD:  {Bytes: []byte{0xed, 0xa9}, Cycle: 0},
	parser.Z80_INST_CPDR: {Bytes: []byte{0xed, 0xb9}, Cycle: 0},
	parser.Z80_INST_DAA:  {Bytes: []byte{0x27}, Cycle: 0},
	parser.Z80_INST_CPL:  {Bytes: []byte{0x2f}, Cycle: 0},
	parser.Z80_INST_NEG:  {Bytes: []byte{0xed, 0x44}, Cycle: 0},
	parser.Z80_INST_CCF:  {Bytes: []byte{0x3f}, Cycle: 0},
	parser.Z80_INST_SCF:  {Bytes: []byte{0x37}, Cycle: 0},
	parser.Z80_INST_NOP:  {Bytes: []byte{0x00}, Cycle: 0},
	parser.Z80_INST_HALT: {Bytes: []byte{0x76}, Cycle: 0},
	parser.Z80_INST_DI:   {Bytes: []byte{0xf3}, Cycle: 0},
	parser.Z80_INST_EI:   {Bytes: []byte{0xfb}, Cycle: 0},
	parser.Z80_INST_RLCA: {Bytes: []byte{0x07}, Cycle: 0},
	parser.Z80_INST_RLA:  {Bytes: []byte{0x17}, Cycle: 0},
	parser.Z80_INST_RRCA: {Bytes: []byte{0x0f}, Cycle: 0},
	parser.Z80_INST_RRA:  {Bytes: []byte{0x1f}, Cycle: 0},
	parser.Z80_INST_RLD:  {Bytes: []byte{0xed, 0x6f}, Cycle: 0},
	parser.Z80_INST_RRD:  {Bytes: []byte{0xed, 0x67}, Cycle: 0},
	parser.Z80_INST_RETI: {Bytes: []byte{0xed, 0x4d}, Cycle: 0},
	parser.Z80_INST_RETN: {Bytes: []byte{0xed, 0x45}, Cycle: 0},
	parser.Z80_INST_INI:  {Bytes: []byte{0xed, 0xa2}, Cycle: 0},
	parser.Z80_INST_INIR: {Bytes: []byte{0xed, 0xb2}, Cycle: 0},
	parser.Z80_INST_IND:  {Bytes: []byte{0xed, 0xaa}, Cycle: 0},
	parser.Z80_INST_INDR: {Bytes: []byte{0xed, 0xba}, Cycle: 0},
	parser.Z80_INST_OUTI: {Bytes: []byte{0xed, 0xa3}, Cycle: 0},
	parser.Z80_INST_OTIR: {Bytes: []byte{0xed, 0xb3}, Cycle: 0},
	parser.Z80_INST_OUTD: {Bytes: []byte{0xed, 0xab}, Cycle: 0},
	parser.Z80_INST_OTDR: {Bytes: []byte{0xed, 0xbb}, Cycle: 0},
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

var Z80Reg16Index map[int]byte = map[int]byte{
	parser.Z80_REG_BC: 0,
	parser.Z80_REG_DE: 1,
	parser.Z80_REG_HL: 2,
	parser.Z80_REG_SP: 3,
}

var Z80FlagIndex map[int]byte = map[int]byte{
	parser.Z80_FLAG_NZ: 0,
	parser.Z80_FLAG_Z:  1,
	parser.Z80_FLAG_NC: 2,
	parser.Z80_REG_C:   3, // lexer では C を Register トークンとして処理するため
	parser.Z80_FLAG_PO: 4,
	parser.Z80_FLAG_PE: 5,
	parser.Z80_FLAG_P:  6,
	parser.Z80_FLAG_M:  7,
}
