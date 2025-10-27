package parser

// Token.SubType 用
const (
	// 8bit Register
	Z80_REG_A = iota + 256
	Z80_REG_B
	Z80_REG_C
	Z80_REG_D
	Z80_REG_E
	Z80_REG_H
	Z80_REG_L
	Z80_REG_IXH
	Z80_REG_IXL
	Z80_REG_IYH
	Z80_REG_IYL
	Z80_REG_I
	Z80_REG_R

	// 16bit Register
	Z80_REG_SP
	Z80_REG_IX
	Z80_REG_IY

	// Register Pair,
	Z80_REG_AF
	Z80_REG_AFEX
	Z80_REG_BC
	Z80_REG_DE
	Z80_REG_HL

	// Flag
	Z80_FLAG_C
	Z80_FLAG_NC
	Z80_FLAG_Z
	Z80_FLAG_NZ
	Z80_FLAG_PO
	Z80_FLAG_PE
	Z80_FLAG_P
	Z80_FLAG_M

	// Instruction
	Z80_INST_LD
	Z80_INST_PUSH
	Z80_INST_POP
	Z80_INST_EX
	Z80_INST_EXX
	Z80_INST_LDI
	Z80_INST_LDIR
	Z80_INST_LDDR
	Z80_INST_CPI
	Z80_INST_CPIR
	Z80_INST_CPDR
	Z80_INST_ADD
	Z80_INST_ADC
	Z80_INST_SUB
	Z80_INST_SBC
	Z80_INST_AND
	Z80_INST_OR
	Z80_INST_CP
	Z80_INST_INC
	Z80_INST_DEC
	Z80_INST_DAA
	Z80_INST_CPL
	Z80_INST_NEG
	Z80_INST_CCF
	Z80_INST_SCF
	Z80_INST_NOP
	Z80_INST_HALT
	Z80_INST_DI
	Z80_INST_EI
	Z80_INST_IM
	Z80_INST_RLCA
	Z80_INST_RLA
	Z80_INST_RRCA
	Z80_INST_RRA
	Z80_INST_RLC
	Z80_INST_RL
	Z80_INST_RR
	Z80_INST_SLA
	Z80_INST_SRA
	Z80_INST_SRL
	Z80_INST_RLD
	Z80_INST_RRD
	Z80_INST_BIT
	Z80_INST_SET
	Z80_INST_RES
	Z80_INST_JP
	Z80_INST_JR
	Z80_INST_DJNZ
	Z80_INST_CALL
	Z80_INST_RET
	Z80_INST_RETI
	Z80_INST_RETN
	Z80_INST_RST
	Z80_INST_IN
	Z80_INST_INI
	Z80_INST_INIR
	Z80_INST_INDR
	Z80_INST_OUT
	Z80_INST_OUTI
	Z80_INST_OTIR
	Z80_INST_OUTD
	Z80_INST_OTDR
)

var z80ReservedWords map[string]Token = map[string]Token{
	"A":   {Type: Z80_REG8, SubType: Z80_REG_A, Literal: "A"},
	"B":   {Type: Z80_REG8, SubType: Z80_REG_B, Literal: "B"},
	"C":   {Type: Z80_REG8, SubType: Z80_REG_C, Literal: "C"},
	"D":   {Type: Z80_REG8, SubType: Z80_REG_D, Literal: "D"},
	"E":   {Type: Z80_REG8, SubType: Z80_REG_E, Literal: "E"},
	"H":   {Type: Z80_REG8, SubType: Z80_REG_H, Literal: "H"},
	"L":   {Type: Z80_REG8, SubType: Z80_REG_L, Literal: "L"},
	"IXH": {Type: Z80_REG8, SubType: Z80_REG_IXH, Literal: "IXH"},
	"IXL": {Type: Z80_REG8, SubType: Z80_REG_IXL, Literal: "IXL"},
	"IYH": {Type: Z80_REG8, SubType: Z80_REG_IYH, Literal: "IYH"},
	"IYL": {Type: Z80_REG8, SubType: Z80_REG_IYL, Literal: "IYL"},
	"I":   {Type: Z80_REG8, SubType: Z80_REG_I, Literal: "I"},
	"R":   {Type: Z80_REG8, SubType: Z80_REG_R, Literal: "R"},

	"SP":  {Type: Z80_REG16, SubType: Z80_REG_SP, Literal: "SP"},
	"IX":  {Type: Z80_REG16, SubType: Z80_REG_IX, Literal: "IX"},
	"IY":  {Type: Z80_REG16, SubType: Z80_REG_IY, Literal: "IY"},
	"AF":  {Type: Z80_REG16, SubType: Z80_REG_AF, Literal: "AF"},
	"AF'": {Type: Z80_REG16, SubType: Z80_REG_AFEX, Literal: "AF'"},
	"BC":  {Type: Z80_REG16, SubType: Z80_REG_BC, Literal: "BC"},
	"DE":  {Type: Z80_REG16, SubType: Z80_REG_DE, Literal: "DE"},
	"HL":  {Type: Z80_REG16, SubType: Z80_REG_HL, Literal: "HL"},

	//	"C":   {Type: Z80_FLAG, SubType: Z80_FLAG_C, Literal: "C"}, レジスタ C と重複
	"NC": {Type: Z80_FLAG, SubType: Z80_FLAG_NC, Literal: "NC"},
	"Z":  {Type: Z80_FLAG, SubType: Z80_FLAG_Z, Literal: "Z"},
	"NZ": {Type: Z80_FLAG, SubType: Z80_FLAG_NZ, Literal: "NZ"},
	"PO": {Type: Z80_FLAG, SubType: Z80_FLAG_PO, Literal: "PO"},
	"PE": {Type: Z80_FLAG, SubType: Z80_FLAG_PE, Literal: "PE"},
	"P":  {Type: Z80_FLAG, SubType: Z80_FLAG_P, Literal: "P"},
	"M":  {Type: Z80_FLAG, SubType: Z80_FLAG_M, Literal: "M"},

	// OpCode
	"LD":   {Type: Z80_INST2, SubType: Z80_INST_LD, Literal: "LD"},
	"PUSH": {Type: Z80_INST1, SubType: Z80_INST_PUSH, Literal: "PUSH"},
	"POP":  {Type: Z80_INST1, SubType: Z80_INST_POP, Literal: "POP"},
	"EX":   {Type: Z80_INST2, SubType: Z80_INST_EX, Literal: "EX"},
	"EXX":  {Type: Z80_INST0, SubType: Z80_INST_EXX, Literal: "EXX"},
	"LDI":  {Type: Z80_INST0, SubType: Z80_INST_LDI, Literal: "LDI"},
	"LDIR": {Type: Z80_INST0, SubType: Z80_INST_LDIR, Literal: "LDIR"},
	"LDDR": {Type: Z80_INST0, SubType: Z80_INST_LDDR, Literal: "LDDR"},
	"CPI":  {Type: Z80_INST0, SubType: Z80_INST_CPI, Literal: "CPI"},
	"CPIR": {Type: Z80_INST0, SubType: Z80_INST_CPIR, Literal: "CPIR"},
	"CPDR": {Type: Z80_INST0, SubType: Z80_INST_CPDR, Literal: "CPDR"},
	"ADD":  {Type: Z80_INST2, SubType: Z80_INST_ADD, Literal: "ADD"},
	"ADC":  {Type: Z80_INST2, SubType: Z80_INST_ADC, Literal: "ADC"},
	"SUB":  {Type: Z80_INST1, SubType: Z80_INST_SUB, Literal: "SUB"},
	"SBC":  {Type: Z80_INST2, SubType: Z80_INST_SBC, Literal: "SBC"},
	"AND":  {Type: Z80_INST1, SubType: Z80_INST_AND, Literal: "AND"},
	"OR":   {Type: Z80_INST1, SubType: Z80_INST_OR, Literal: "OR"},
	"CP":   {Type: Z80_INST1, SubType: Z80_INST_CP, Literal: "CP"},
	"INC":  {Type: Z80_INST1, SubType: Z80_INST_INC, Literal: "INC"},
	"DEC":  {Type: Z80_INST1, SubType: Z80_INST_DEC, Literal: "DEC"},
	"DAA":  {Type: Z80_INST0, SubType: Z80_INST_DAA, Literal: "DAA"},
	"CPL":  {Type: Z80_INST0, SubType: Z80_INST_CPL, Literal: "CPL"},
	"NEG":  {Type: Z80_INST0, SubType: Z80_INST_NEG, Literal: "NEG"},
	"CCF":  {Type: Z80_INST0, SubType: Z80_INST_CCF, Literal: "CCF"},
	"SCF":  {Type: Z80_INST0, SubType: Z80_INST_SCF, Literal: "SCF"},
	"NOP":  {Type: Z80_INST0, SubType: Z80_INST_NOP, Literal: "NOP"},
	"HALT": {Type: Z80_INST0, SubType: Z80_INST_HALT, Literal: "HALT"},
	"DI":   {Type: Z80_INST0, SubType: Z80_INST_DI, Literal: "DI"},
	"EI":   {Type: Z80_INST0, SubType: Z80_INST_EI, Literal: "EI"},
	"IM":   {Type: Z80_INST1, SubType: Z80_INST_IM, Literal: "IM"},
	"RLCA": {Type: Z80_INST0, SubType: Z80_INST_RLCA, Literal: "RLCA"},
	"RLA":  {Type: Z80_INST0, SubType: Z80_INST_RLA, Literal: "RLA"},
	"RRCA": {Type: Z80_INST0, SubType: Z80_INST_RRCA, Literal: "RRCA"},
	"RRA":  {Type: Z80_INST0, SubType: Z80_INST_RRA, Literal: "RRA"},
	"RLC":  {Type: Z80_INST1, SubType: Z80_INST_RLC, Literal: "RLC"},
	"RL":   {Type: Z80_INST1, SubType: Z80_INST_RL, Literal: "RL"},
	"RR":   {Type: Z80_INST1, SubType: Z80_INST_RR, Literal: "RR"},
	"SLA":  {Type: Z80_INST1, SubType: Z80_INST_SLA, Literal: "SLA"},
	"SRA":  {Type: Z80_INST1, SubType: Z80_INST_SRA, Literal: "SRA"},
	"SRL":  {Type: Z80_INST1, SubType: Z80_INST_SRL, Literal: "SRL"},
	"RLD":  {Type: Z80_INST0, SubType: Z80_INST_RLD, Literal: "RLD"},
	"RRD":  {Type: Z80_INST0, SubType: Z80_INST_RRD, Literal: "RRD"},
	"BIT":  {Type: Z80_INST2, SubType: Z80_INST_BIT, Literal: "BIT"},
	"SET":  {Type: Z80_INST2, SubType: Z80_INST_SET, Literal: "SET"},
	"RES":  {Type: Z80_INST2, SubType: Z80_INST_RES, Literal: "RES"},
	"JP":   {Type: Z80_INST2, SubType: Z80_INST_JP, Literal: "JP"}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"JR":   {Type: Z80_INST2, SubType: Z80_INST_JR, Literal: "JR"}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"DJNZ": {Type: Z80_INST1, SubType: Z80_INST_DJNZ, Literal: "DJNZ"},
	"CALL": {Type: Z80_INST2, SubType: Z80_INST_CALL, Literal: "CALL"}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"RET":  {Type: Z80_INST0, SubType: Z80_INST_RET, Literal: "RET"},   // cc 有無により OPCODE0, OPCODE1 両方あり
	"RETI": {Type: Z80_INST0, SubType: Z80_INST_RETI, Literal: "RETI"},
	"RETN": {Type: Z80_INST0, SubType: Z80_INST_RETN, Literal: "RETN"},
	"RST":  {Type: Z80_INST1, SubType: Z80_INST_RST, Literal: "RST"},
	"IN":   {Type: Z80_INST2, SubType: Z80_INST_IN, Literal: "IN"},
	"INI":  {Type: Z80_INST0, SubType: Z80_INST_INI, Literal: "INI"},
	"INIR": {Type: Z80_INST0, SubType: Z80_INST_INIR, Literal: "INIR"},
	"INDR": {Type: Z80_INST0, SubType: Z80_INST_INDR, Literal: "INDR"},
	"OUT":  {Type: Z80_INST2, SubType: Z80_INST_OUT, Literal: "OUT"},
	"OUTI": {Type: Z80_INST0, SubType: Z80_INST_OUTI, Literal: "OUTI"},
	"OTIR": {Type: Z80_INST0, SubType: Z80_INST_OTIR, Literal: "OTIR"},
	"OUTD": {Type: Z80_INST0, SubType: Z80_INST_OUTD, Literal: "OUTD"},
	"OTDR": {Type: Z80_INST0, SubType: Z80_INST_OTDR, Literal: "OTDR"},
}

func z80OpCode2Name(opcode int) string {
	for _, v := range z80ReservedWords {
		if v.SubType == opcode {
			return v.Literal
		}
	}
	return "UNKNOWN"
}
