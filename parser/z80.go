package parser

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

var Z80Registers map[string]Token = map[string]Token{
	"A":   {Type: Z80_REG8, Literal: "A", SubType: Z80_REG_A},
	"B":   {Type: Z80_REG8, Literal: "B", SubType: Z80_REG_B},
	"C":   {Type: Z80_REG8, Literal: "C", SubType: Z80_REG_C},
	"D":   {Type: Z80_REG8, Literal: "D", SubType: Z80_REG_D},
	"E":   {Type: Z80_REG8, Literal: "E", SubType: Z80_REG_E},
	"H":   {Type: Z80_REG8, Literal: "H", SubType: Z80_REG_H},
	"L":   {Type: Z80_REG8, Literal: "L", SubType: Z80_REG_L},
	"IXH": {Type: Z80_REG8, Literal: "IXH", SubType: Z80_REG_IXH},
	"IXL": {Type: Z80_REG8, Literal: "IXL", SubType: Z80_REG_IXL},
	"IYH": {Type: Z80_REG8, Literal: "IYH", SubType: Z80_REG_IYH},
	"IYL": {Type: Z80_REG8, Literal: "IYL", SubType: Z80_REG_IYL},
	"I":   {Type: Z80_REG8, Literal: "I", SubType: Z80_REG_I},
	"R":   {Type: Z80_REG8, Literal: "R", SubType: Z80_REG_R},

	"SP":  {Type: Z80_REG16, Literal: "SP", SubType: Z80_REG_SP},
	"IX":  {Type: Z80_REG16, Literal: "IX", SubType: Z80_REG_IX},
	"IY":  {Type: Z80_REG16, Literal: "IY", SubType: Z80_REG_IY},
	"AF":  {Type: Z80_REG16, Literal: "AF", SubType: Z80_REG_AF},
	"AF'": {Type: Z80_REG16, Literal: "AF'", SubType: Z80_REG_AFEX},
	"BC":  {Type: Z80_REG16, Literal: "BC", SubType: Z80_REG_BC},
	"DE":  {Type: Z80_REG16, Literal: "DE", SubType: Z80_REG_DE},
	"HL":  {Type: Z80_REG16, Literal: "HL", SubType: Z80_REG_HL},

	// "C":  {Type: Z80_FLAG, Literal: "C", Op: Z80_FLAG_C},
	"NC": {Type: Z80_FLAG, Literal: "NC", SubType: Z80_FLAG_NC},
	"Z":  {Type: Z80_FLAG, Literal: "Z", SubType: Z80_FLAG_Z},
	"NZ": {Type: Z80_FLAG, Literal: "NZ", SubType: Z80_FLAG_NZ},
	"PO": {Type: Z80_FLAG, Literal: "PO", SubType: Z80_FLAG_PO},
	"PE": {Type: Z80_FLAG, Literal: "PE", SubType: Z80_FLAG_PE},
	"P":  {Type: Z80_FLAG, Literal: "P", SubType: Z80_FLAG_P},
	"M":  {Type: Z80_FLAG, Literal: "M", SubType: Z80_FLAG_M},
}

var Z80OpCodes map[string]Token = map[string]Token{
	"LD":   {Type: Z80_INST2, Literal: "LD", SubType: Z80_INST_LD},
	"PUSH": {Type: Z80_INST1, Literal: "PUSH", SubType: Z80_INST_PUSH},
	"POP":  {Type: Z80_INST1, Literal: "POP", SubType: Z80_INST_POP},
	"EX":   {Type: Z80_INST2, Literal: "EX", SubType: Z80_INST_EX},
	"EXX":  {Type: Z80_INST0, Literal: "EXX", SubType: Z80_INST_EXX},
	"LDI":  {Type: Z80_INST0, Literal: "LDI", SubType: Z80_INST_LDI},
	"LDIR": {Type: Z80_INST0, Literal: "LDIR", SubType: Z80_INST_LDIR},
	"LDDR": {Type: Z80_INST0, Literal: "LDDR", SubType: Z80_INST_LDDR},
	"CPI":  {Type: Z80_INST0, Literal: "CPI", SubType: Z80_INST_CPI},
	"CPIR": {Type: Z80_INST0, Literal: "CPIR", SubType: Z80_INST_CPIR},
	"CPDR": {Type: Z80_INST0, Literal: "CPDR", SubType: Z80_INST_CPDR},
	"ADD":  {Type: Z80_INST2, Literal: "ADD", SubType: Z80_INST_ADD},
	"ADC":  {Type: Z80_INST2, Literal: "ADC", SubType: Z80_INST_ADC},
	"SUB":  {Type: Z80_INST1, Literal: "SUB", SubType: Z80_INST_SUB},
	"SBC":  {Type: Z80_INST2, Literal: "SBC", SubType: Z80_INST_SBC},
	"AND":  {Type: Z80_INST1, Literal: "AND", SubType: Z80_INST_AND},
	"OR":   {Type: Z80_INST1, Literal: "OR", SubType: Z80_INST_OR},
	"CP":   {Type: Z80_INST1, Literal: "CP", SubType: Z80_INST_CP},
	"INC":  {Type: Z80_INST1, Literal: "INC", SubType: Z80_INST_INC},
	"DEC":  {Type: Z80_INST1, Literal: "DEC", SubType: Z80_INST_DEC},
	"DAA":  {Type: Z80_INST0, Literal: "DAA", SubType: Z80_INST_DAA},
	"CPL":  {Type: Z80_INST0, Literal: "CPL", SubType: Z80_INST_CPL},
	"NEG":  {Type: Z80_INST0, Literal: "NEG", SubType: Z80_INST_NEG},
	"CCF":  {Type: Z80_INST0, Literal: "CCF", SubType: Z80_INST_CCF},
	"SCF":  {Type: Z80_INST0, Literal: "SCF", SubType: Z80_INST_SCF},
	"NOP":  {Type: Z80_INST0, Literal: "NOP", SubType: Z80_INST_NOP},
	"HALT": {Type: Z80_INST0, Literal: "HALT", SubType: Z80_INST_HALT},
	"DI":   {Type: Z80_INST0, Literal: "DI", SubType: Z80_INST_DI},
	"EI":   {Type: Z80_INST0, Literal: "EI", SubType: Z80_INST_EI},
	"IM":   {Type: Z80_INST1, Literal: "IM", SubType: Z80_INST_IM},
	"RLCA": {Type: Z80_INST0, Literal: "RLCA", SubType: Z80_INST_RLCA},
	"RLA":  {Type: Z80_INST0, Literal: "RLA", SubType: Z80_INST_RLA},
	"RRCA": {Type: Z80_INST0, Literal: "RRCA", SubType: Z80_INST_RRCA},
	"RRA":  {Type: Z80_INST0, Literal: "RRA", SubType: Z80_INST_RRA},
	"RLC":  {Type: Z80_INST1, Literal: "RLC", SubType: Z80_INST_RLC},
	"RL":   {Type: Z80_INST1, Literal: "RL", SubType: Z80_INST_RL},
	"RR":   {Type: Z80_INST1, Literal: "RR", SubType: Z80_INST_RR},
	"SLA":  {Type: Z80_INST1, Literal: "SLA", SubType: Z80_INST_SLA},
	"SRA":  {Type: Z80_INST1, Literal: "SRA", SubType: Z80_INST_SRA},
	"SRL":  {Type: Z80_INST1, Literal: "SRL", SubType: Z80_INST_SRL},
	"RLD":  {Type: Z80_INST0, Literal: "RLD", SubType: Z80_INST_RLD},
	"RRD":  {Type: Z80_INST0, Literal: "RRD", SubType: Z80_INST_RRD},
	"BIT":  {Type: Z80_INST2, Literal: "BIT", SubType: Z80_INST_BIT},
	"SET":  {Type: Z80_INST2, Literal: "SET", SubType: Z80_INST_SET},
	"RES":  {Type: Z80_INST2, Literal: "RES", SubType: Z80_INST_RES},
	"JP":   {Type: Z80_INST2, Literal: "JP", SubType: Z80_INST_JP}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"JR":   {Type: Z80_INST2, Literal: "JR", SubType: Z80_INST_JR}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"DJNZ": {Type: Z80_INST1, Literal: "DJNZ", SubType: Z80_INST_DJNZ},
	"CALL": {Type: Z80_INST2, Literal: "CALL", SubType: Z80_INST_CALL}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"RET":  {Type: Z80_INST0, Literal: "RET", SubType: Z80_INST_RET},   // cc 有無により OPCODE0, OPCODE1 両方あり
	"RETI": {Type: Z80_INST0, Literal: "RETI", SubType: Z80_INST_RETI},
	"RETN": {Type: Z80_INST0, Literal: "RETN", SubType: Z80_INST_RETN},
	"RST":  {Type: Z80_INST1, Literal: "RST", SubType: Z80_INST_RST},
	"IN":   {Type: Z80_INST2, Literal: "IN", SubType: Z80_INST_IN},
	"INI":  {Type: Z80_INST0, Literal: "INI", SubType: Z80_INST_INI},
	"INIR": {Type: Z80_INST0, Literal: "INIR", SubType: Z80_INST_INIR},
	"INDR": {Type: Z80_INST0, Literal: "INDR", SubType: Z80_INST_INDR},
	"OUT":  {Type: Z80_INST2, Literal: "OUT", SubType: Z80_INST_OUT},
	"OUTI": {Type: Z80_INST0, Literal: "OUTI", SubType: Z80_INST_OUTI},
	"OTIR": {Type: Z80_INST0, Literal: "OTIR", SubType: Z80_INST_OTIR},
	"OUTD": {Type: Z80_INST0, Literal: "OUTD", SubType: Z80_INST_OUTD},
	"OTDR": {Type: Z80_INST0, Literal: "OTDR", SubType: Z80_INST_OTDR},
}

func Z80Names(opcode int) string {
	for _, v := range Z80Registers {
		if v.SubType == opcode {
			return v.Literal
		}
	}
	for _, v := range Z80OpCodes {
		if v.SubType == opcode {
			return v.Literal
		}
	}
	switch opcode {
	case Z80_FLAG_NC:
		{
			return "NC"
		}
	case Z80_FLAG_Z:
		{
			return "Z"
		}
	case Z80_FLAG_NZ:
		{
			return "NZ"
		}
	case Z80_FLAG_PO:
		{
			return "PO"
		}
	case Z80_FLAG_PE:
		{
			return "PE"
		}
	case Z80_FLAG_P:
		{
			return "P"
		}
	case Z80_FLAG_M:
		{
			return "M"
		}
	}
	return "UNKNOWN"
}
