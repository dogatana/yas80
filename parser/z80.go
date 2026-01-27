package parser

// Token.SubType 用
const (
	// 8bit Register
	Z80_REG_B = iota + 256
	Z80_REG_C
	Z80_REG_D
	Z80_REG_E
	Z80_REG_H
	Z80_REG_L

	Z80_REG_A
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
	Z80_FLAG_NZ
	Z80_FLAG_Z
	Z80_FLAG_NC
	Z80_FLAG_C
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
	Z80_INST_LDD
	Z80_INST_LDDR
	Z80_INST_CPI
	Z80_INST_CPIR
	Z80_INST_CPD
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
	Z80_INST_RRC
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
	Z80_INST_IND
	Z80_INST_INDR
	Z80_INST_OUT
	Z80_INST_OUTI
	Z80_INST_OTIR
	Z80_INST_OUTD
	Z80_INST_OTDR
	Z80_INST_XOR
)

var z80ReservedWords map[string]Token = map[string]Token{
	"A":   {TokenType: Z80_REG8, TokenSubType: Z80_REG_A, Literal: "A"},
	"B":   {TokenType: Z80_REG8, TokenSubType: Z80_REG_B, Literal: "B"},
	"C":   {TokenType: Z80_REG8, TokenSubType: Z80_REG_C, Literal: "C"},
	"D":   {TokenType: Z80_REG8, TokenSubType: Z80_REG_D, Literal: "D"},
	"E":   {TokenType: Z80_REG8, TokenSubType: Z80_REG_E, Literal: "E"},
	"H":   {TokenType: Z80_REG8, TokenSubType: Z80_REG_H, Literal: "H"},
	"L":   {TokenType: Z80_REG8, TokenSubType: Z80_REG_L, Literal: "L"},
	"IXH": {TokenType: Z80_REG8, TokenSubType: Z80_REG_IXH, Literal: "IXH"},
	"IXL": {TokenType: Z80_REG8, TokenSubType: Z80_REG_IXL, Literal: "IXL"},
	"IYH": {TokenType: Z80_REG8, TokenSubType: Z80_REG_IYH, Literal: "IYH"},
	"IYL": {TokenType: Z80_REG8, TokenSubType: Z80_REG_IYL, Literal: "IYL"},
	"I":   {TokenType: Z80_REG8, TokenSubType: Z80_REG_I, Literal: "I"},
	"R":   {TokenType: Z80_REG8, TokenSubType: Z80_REG_R, Literal: "R"},

	"SP":  {TokenType: Z80_REG16, TokenSubType: Z80_REG_SP, Literal: "SP"},
	"IX":  {TokenType: Z80_REG16, TokenSubType: Z80_REG_IX, Literal: "IX"},
	"IY":  {TokenType: Z80_REG16, TokenSubType: Z80_REG_IY, Literal: "IY"},
	"AF":  {TokenType: Z80_REG16, TokenSubType: Z80_REG_AF, Literal: "AF"},
	"AF'": {TokenType: Z80_REG16, TokenSubType: Z80_REG_AFEX, Literal: "AF'"},
	"BC":  {TokenType: Z80_REG16, TokenSubType: Z80_REG_BC, Literal: "BC"},
	"DE":  {TokenType: Z80_REG16, TokenSubType: Z80_REG_DE, Literal: "DE"},
	"HL":  {TokenType: Z80_REG16, TokenSubType: Z80_REG_HL, Literal: "HL"},

	"CY": {TokenType: Z80_FLAG, TokenSubType: Z80_FLAG_C, Literal: "CY"}, // キャリーの別名を定義
	"NC": {TokenType: Z80_FLAG, TokenSubType: Z80_FLAG_NC, Literal: "NC"},
	"Z":  {TokenType: Z80_FLAG, TokenSubType: Z80_FLAG_Z, Literal: "Z"},
	"NZ": {TokenType: Z80_FLAG, TokenSubType: Z80_FLAG_NZ, Literal: "NZ"},
	"PO": {TokenType: Z80_FLAG, TokenSubType: Z80_FLAG_PO, Literal: "PO"},
	"PE": {TokenType: Z80_FLAG, TokenSubType: Z80_FLAG_PE, Literal: "PE"},
	"P":  {TokenType: Z80_FLAG, TokenSubType: Z80_FLAG_P, Literal: "P"},
	"M":  {TokenType: Z80_FLAG, TokenSubType: Z80_FLAG_M, Literal: "M"},

	// Opcode
	"LD":   {TokenType: Z80_INST2, TokenSubType: Z80_INST_LD, Literal: "LD"},
	"PUSH": {TokenType: Z80_INST1, TokenSubType: Z80_INST_PUSH, Literal: "PUSH"},
	"POP":  {TokenType: Z80_INST1, TokenSubType: Z80_INST_POP, Literal: "POP"},
	"EX":   {TokenType: Z80_INST2, TokenSubType: Z80_INST_EX, Literal: "EX"},
	"EXX":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_EXX, Literal: "EXX"},
	"LDI":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_LDI, Literal: "LDI"},
	"LDIR": {TokenType: Z80_INST0, TokenSubType: Z80_INST_LDIR, Literal: "LDIR"},
	"LDD":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_LDD, Literal: "LDD"},
	"LDDR": {TokenType: Z80_INST0, TokenSubType: Z80_INST_LDDR, Literal: "LDDR"},
	"CPI":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_CPI, Literal: "CPI"},
	"CPIR": {TokenType: Z80_INST0, TokenSubType: Z80_INST_CPIR, Literal: "CPIR"},
	"CPD":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_CPD, Literal: "CPD"},
	"CPDR": {TokenType: Z80_INST0, TokenSubType: Z80_INST_CPDR, Literal: "CPDR"},
	"ADD":  {TokenType: Z80_INST2, TokenSubType: Z80_INST_ADD, Literal: "ADD"},
	"ADC":  {TokenType: Z80_INST2, TokenSubType: Z80_INST_ADC, Literal: "ADC"},
	"SUB":  {TokenType: Z80_INST2, TokenSubType: Z80_INST_SUB, Literal: "SUB"},
	"SBC":  {TokenType: Z80_INST2, TokenSubType: Z80_INST_SBC, Literal: "SBC"},
	"AND":  {TokenType: Z80_INST2, TokenSubType: Z80_INST_AND, Literal: "AND"},
	"OR":   {TokenType: Z80_INST2, TokenSubType: Z80_INST_OR, Literal: "OR"},
	"XOR":  {TokenType: Z80_INST2, TokenSubType: Z80_INST_XOR, Literal: "XOR"},
	"CP":   {TokenType: Z80_INST2, TokenSubType: Z80_INST_CP, Literal: "CP"},
	"INC":  {TokenType: Z80_INST1, TokenSubType: Z80_INST_INC, Literal: "INC"},
	"DEC":  {TokenType: Z80_INST1, TokenSubType: Z80_INST_DEC, Literal: "DEC"},
	"DAA":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_DAA, Literal: "DAA"},
	"CPL":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_CPL, Literal: "CPL"},
	"NEG":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_NEG, Literal: "NEG"},
	"CCF":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_CCF, Literal: "CCF"},
	"SCF":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_SCF, Literal: "SCF"},
	"NOP":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_NOP, Literal: "NOP"},
	"HALT": {TokenType: Z80_INST0, TokenSubType: Z80_INST_HALT, Literal: "HALT"},
	"DI":   {TokenType: Z80_INST0, TokenSubType: Z80_INST_DI, Literal: "DI"},
	"EI":   {TokenType: Z80_INST0, TokenSubType: Z80_INST_EI, Literal: "EI"},
	"IM":   {TokenType: Z80_INST1, TokenSubType: Z80_INST_IM, Literal: "IM"},
	"RLCA": {TokenType: Z80_INST0, TokenSubType: Z80_INST_RLCA, Literal: "RLCA"},
	"RLA":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_RLA, Literal: "RLA"},
	"RRCA": {TokenType: Z80_INST0, TokenSubType: Z80_INST_RRCA, Literal: "RRCA"},
	"RRA":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_RRA, Literal: "RRA"},
	"RLC":  {TokenType: Z80_INST1, TokenSubType: Z80_INST_RLC, Literal: "RLC"},
	"RL":   {TokenType: Z80_INST1, TokenSubType: Z80_INST_RL, Literal: "RL"},
	"RR":   {TokenType: Z80_INST1, TokenSubType: Z80_INST_RR, Literal: "RR"},
	"RRC":  {TokenType: Z80_INST1, TokenSubType: Z80_INST_RRC, Literal: "RRC"},
	"SLA":  {TokenType: Z80_INST1, TokenSubType: Z80_INST_SLA, Literal: "SLA"},
	"SRA":  {TokenType: Z80_INST1, TokenSubType: Z80_INST_SRA, Literal: "SRA"},
	"SRL":  {TokenType: Z80_INST1, TokenSubType: Z80_INST_SRL, Literal: "SRL"},
	"RLD":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_RLD, Literal: "RLD"},
	"RRD":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_RRD, Literal: "RRD"},
	"BIT":  {TokenType: Z80_INST2, TokenSubType: Z80_INST_BIT, Literal: "BIT"},
	"SET":  {TokenType: Z80_INST2, TokenSubType: Z80_INST_SET, Literal: "SET"},
	"RES":  {TokenType: Z80_INST2, TokenSubType: Z80_INST_RES, Literal: "RES"},
	"JP":   {TokenType: Z80_INST2, TokenSubType: Z80_INST_JP, Literal: "JP"}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"JR":   {TokenType: Z80_INST2, TokenSubType: Z80_INST_JR, Literal: "JR"}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"DJNZ": {TokenType: Z80_INST1, TokenSubType: Z80_INST_DJNZ, Literal: "DJNZ"},
	"CALL": {TokenType: Z80_INST2, TokenSubType: Z80_INST_CALL, Literal: "CALL"}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"RET":  {TokenType: Z80_INST1, TokenSubType: Z80_INST_RET, Literal: "RET"},   // cc 有無により OPCODE0, OPCODE1 両方あり
	"RETI": {TokenType: Z80_INST0, TokenSubType: Z80_INST_RETI, Literal: "RETI"},
	"RETN": {TokenType: Z80_INST0, TokenSubType: Z80_INST_RETN, Literal: "RETN"},
	"RST":  {TokenType: Z80_INST1, TokenSubType: Z80_INST_RST, Literal: "RST"},
	"IN":   {TokenType: Z80_INST2, TokenSubType: Z80_INST_IN, Literal: "IN"},
	"INI":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_INI, Literal: "INI"},
	"INIR": {TokenType: Z80_INST0, TokenSubType: Z80_INST_INIR, Literal: "INIR"},
	"IND":  {TokenType: Z80_INST0, TokenSubType: Z80_INST_IND, Literal: "IND"},
	"INDR": {TokenType: Z80_INST0, TokenSubType: Z80_INST_INDR, Literal: "INDR"},
	"OUT":  {TokenType: Z80_INST2, TokenSubType: Z80_INST_OUT, Literal: "OUT"},
	"OUTI": {TokenType: Z80_INST0, TokenSubType: Z80_INST_OUTI, Literal: "OUTI"},
	"OTIR": {TokenType: Z80_INST0, TokenSubType: Z80_INST_OTIR, Literal: "OTIR"},
	"OUTD": {TokenType: Z80_INST0, TokenSubType: Z80_INST_OUTD, Literal: "OUTD"},
	"OTDR": {TokenType: Z80_INST0, TokenSubType: Z80_INST_OTDR, Literal: "OTDR"},
}

func Z80Opcode2Name(opcode int) string {
	for _, v := range z80ReservedWords {
		if int(v.TokenSubType) == opcode {
			return v.Literal
		}
	}
	return "UNKNOWN"
}
