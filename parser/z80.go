package parser

const (
	// 8bit
	Z80_REG_A   = 'A'
	Z80_REG_B   = 'B'
	Z80_REG_C   = 'C'
	Z80_REG_D   = 'D'
	Z80_REG_E   = 'E'
	Z80_REG_H   = 'H'
	Z80_REG_L   = 'L'
	Z80_REG_IXH = 'X'
	Z80_REG_IXL = 'x'
	Z80_REG_IYH = 'Y'
	Z80_REG_IYL = 'y'
	Z80_REG_I   = 'I'
	Z80_REG_R   = 'R'

	// 16bit
	Z80_REG_SP = 'S'
	Z80_REG_IX = 'X'
	Z80_REG_IY = 'Y'

	// register pair,
	Z80_REG_AF   = 'A'
	Z80_REG_AFEX = 'F' // AF'
	Z80_REG_BC   = 'B'
	Z80_REG_DE   = 'D'
	Z80_REG_HL   = 'H'

	// flag
	Z80_FLAG_C  = 'C'
	Z80_FLAG_NC = 'c'
	Z80_FLAG_Z  = 'Z'
	Z80_FLAG_NZ = 'z'
	Z80_FLAG_PO = 'O'
	Z80_FLAG_PE = 'E'
	Z80_FLAG_P  = 'P'
	Z80_FLAG_M  = 'M'
)

var Z80Registers map[string]Token = map[string]Token{
	"A":   {Type: Z80_REG8, Literal: "A", Op: Z80_REG_A},
	"B":   {Type: Z80_REG8, Literal: "B", Op: Z80_REG_B},
	"C":   {Type: Z80_REG8, Literal: "C", Op: Z80_REG_C},
	"D":   {Type: Z80_REG8, Literal: "D", Op: Z80_REG_D},
	"E":   {Type: Z80_REG8, Literal: "E", Op: Z80_REG_E},
	"H":   {Type: Z80_REG8, Literal: "H", Op: Z80_REG_H},
	"L":   {Type: Z80_REG8, Literal: "L", Op: Z80_REG_L},
	"IXH": {Type: Z80_REG8, Literal: "IXH", Op: Z80_REG_IXH},
	"IXL": {Type: Z80_REG8, Literal: "IXL", Op: Z80_REG_IXL},
	"IYH": {Type: Z80_REG8, Literal: "IYH", Op: Z80_REG_IYH},
	"IYL": {Type: Z80_REG8, Literal: "IYL", Op: Z80_REG_IYL},
	"I":   {Type: Z80_REG8, Literal: "I", Op: Z80_REG_I},
	"R":   {Type: Z80_REG8, Literal: "R", Op: Z80_REG_R},

	"SP":  {Type: Z80_REG16, Literal: "SP", Op: Z80_REG_SP},
	"IX":  {Type: Z80_REG16, Literal: "IX", Op: Z80_REG_IX},
	"IY":  {Type: Z80_REG16, Literal: "IY", Op: Z80_REG_IY},
	"AF":  {Type: Z80_REG16, Literal: "AF", Op: Z80_REG_AF},
	"AF'": {Type: Z80_REG16, Literal: "AF'", Op: Z80_REG_AFEX},
	"BC":  {Type: Z80_REG16, Literal: "BC", Op: Z80_REG_BC},
	"DE":  {Type: Z80_REG16, Literal: "DE", Op: Z80_REG_DE},
	"HL":  {Type: Z80_REG16, Literal: "HL", Op: Z80_REG_HL},

	// "C":  {Type: Z80_FLAG, Literal: "C", Op: Z80_FLAG_C},
	"NC": {Type: Z80_FLAG, Literal: "NC", Op: Z80_FLAG_NC},
	"Z":  {Type: Z80_FLAG, Literal: "Z", Op: Z80_FLAG_Z},
	"NZ": {Type: Z80_FLAG, Literal: "NZ", Op: Z80_FLAG_NZ},
	"PO": {Type: Z80_FLAG, Literal: "PO", Op: Z80_FLAG_PO},
	"PE": {Type: Z80_FLAG, Literal: "PE", Op: Z80_FLAG_PE},
	"P":  {Type: Z80_FLAG, Literal: "P", Op: Z80_FLAG_P},
	"M":  {Type: Z80_FLAG, Literal: "M", Op: Z80_FLAG_M},
}

var Z80Instructions map[string]int = map[string]int{
	"LD": Z80_INST0,

	"PUSH": Z80_INST1,
	"POP":  Z80_INST1,

	"EX":  Z80_INST2,
	"EXX": Z80_INST0,

	"LDI":  Z80_INST0,
	"LDIR": Z80_INST0,
	"LDDR": Z80_INST0,
	"CPI":  Z80_INST0,
	"CPIR": Z80_INST0,
	"CPDR": Z80_INST0,

	"ADD": Z80_INST2,
	"ADC": Z80_INST2,

	"SUB": Z80_INST1,
	"SBC": Z80_INST2,

	"AND": Z80_INST1,
	"OR":  Z80_INST1,
	"CP":  Z80_INST1,
	"INC": Z80_INST1,
	"DEC": Z80_INST1,

	"DAA":  Z80_INST0,
	"CPL":  Z80_INST0,
	"NEG":  Z80_INST0,
	"CCF":  Z80_INST0,
	"SCF":  Z80_INST0,
	"NOP":  Z80_INST0,
	"HALT": Z80_INST0,
	"DI":   Z80_INST0,
	"EI":   Z80_INST0,
	"IM":   Z80_INST1,

	"RLCA": Z80_INST0,
	"RLA":  Z80_INST0,
	"RRCA": Z80_INST0,
	"RRA":  Z80_INST0,

	"RLC": Z80_INST1,
	"RL":  Z80_INST1,
	"RR":  Z80_INST1,
	"SLA": Z80_INST1,
	"SRA": Z80_INST1,
	"SRL": Z80_INST1,

	"RLD": Z80_INST0,
	"RRD": Z80_INST0,

	"BIT": Z80_INST2,
	"SET": Z80_INST2,
	"RES": Z80_INST2,

	"JP":   Z80_INST2, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"JR":   Z80_INST2, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"DJNZ": Z80_INST1,

	"CALL": Z80_INST2, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"RET":  Z80_INST1, // cc 有無により OPCODE0, OPCODE1 両方あり
	"RETI": Z80_INST0,
	"RETN": Z80_INST0,
	"REST": Z80_INST1,

	"IN":   Z80_INST2,
	"INI":  Z80_INST0,
	"INIR": Z80_INST0,
	"INDR": Z80_INST0,
	"OUT":  Z80_INST2,
	"OUTI": Z80_INST0,
	"OTIR": Z80_INST0,
	"OUTD": Z80_INST0,
	"OTDR": Z80_INST0,
}
