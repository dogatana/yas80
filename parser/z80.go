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
	Z80_INST_REST
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

var Z80OpCodes map[string]Token = map[string]Token{
	"LD":   {Type: Z80_INST2, Literal: "LD", Op: Z80_INST_LD},
	"PUSH": {Type: Z80_INST1, Literal: "PUSH", Op: Z80_INST_PUSH},
	"POP":  {Type: Z80_INST1, Literal: "POP", Op: Z80_INST_POP},
	"EX":   {Type: Z80_INST2, Literal: "EX", Op: Z80_INST_EX},
	"EXX":  {Type: Z80_INST0, Literal: "EXX", Op: Z80_INST_EXX},
	"LDI":  {Type: Z80_INST0, Literal: "LDI", Op: Z80_INST_LDI},
	"LDIR": {Type: Z80_INST0, Literal: "LDIR", Op: Z80_INST_LDIR},
	"LDDR": {Type: Z80_INST0, Literal: "LDDR", Op: Z80_INST_LDDR},
	"CPI":  {Type: Z80_INST0, Literal: "CPI", Op: Z80_INST_CPI},
	"CPIR": {Type: Z80_INST0, Literal: "CPIR", Op: Z80_INST_CPIR},
	"CPDR": {Type: Z80_INST0, Literal: "CPDR", Op: Z80_INST_CPDR},
	"ADD":  {Type: Z80_INST2, Literal: "ADD", Op: Z80_INST_ADD},
	"ADC":  {Type: Z80_INST2, Literal: "ADC", Op: Z80_INST_ADC},
	"SUB":  {Type: Z80_INST1, Literal: "SUB", Op: Z80_INST_SUB},
	"SBC":  {Type: Z80_INST2, Literal: "SBC", Op: Z80_INST_SBC},
	"AND":  {Type: Z80_INST1, Literal: "AND", Op: Z80_INST_AND},
	"OR":   {Type: Z80_INST1, Literal: "OR", Op: Z80_INST_OR},
	"CP":   {Type: Z80_INST1, Literal: "CP", Op: Z80_INST_CP},
	"INC":  {Type: Z80_INST1, Literal: "INC", Op: Z80_INST_INC},
	"DEC":  {Type: Z80_INST1, Literal: "DEC", Op: Z80_INST_DEC},
	"DAA":  {Type: Z80_INST0, Literal: "DAA", Op: Z80_INST_DAA},
	"CPL":  {Type: Z80_INST0, Literal: "CPL", Op: Z80_INST_CPL},
	"NEG":  {Type: Z80_INST0, Literal: "NEG", Op: Z80_INST_NEG},
	"CCF":  {Type: Z80_INST0, Literal: "CCF", Op: Z80_INST_CCF},
	"SCF":  {Type: Z80_INST0, Literal: "SCF", Op: Z80_INST_SCF},
	"NOP":  {Type: Z80_INST0, Literal: "NOP", Op: Z80_INST_NOP},
	"HALT": {Type: Z80_INST0, Literal: "HALT", Op: Z80_INST_HALT},
	"DI":   {Type: Z80_INST0, Literal: "DI", Op: Z80_INST_DI},
	"EI":   {Type: Z80_INST0, Literal: "EI", Op: Z80_INST_EI},
	"IM":   {Type: Z80_INST1, Literal: "IM", Op: Z80_INST_IM},
	"RLCA": {Type: Z80_INST0, Literal: "RLCA", Op: Z80_INST_RLCA},
	"RLA":  {Type: Z80_INST0, Literal: "RLA", Op: Z80_INST_RLA},
	"RRCA": {Type: Z80_INST0, Literal: "RRCA", Op: Z80_INST_RRCA},
	"RRA":  {Type: Z80_INST0, Literal: "RRA", Op: Z80_INST_RRA},
	"RLC":  {Type: Z80_INST1, Literal: "RLC", Op: Z80_INST_RLC},
	"RL":   {Type: Z80_INST1, Literal: "RL", Op: Z80_INST_RL},
	"RR":   {Type: Z80_INST1, Literal: "RR", Op: Z80_INST_RR},
	"SLA":  {Type: Z80_INST1, Literal: "SLA", Op: Z80_INST_SLA},
	"SRA":  {Type: Z80_INST1, Literal: "SRA", Op: Z80_INST_SRA},
	"SRL":  {Type: Z80_INST1, Literal: "SRL", Op: Z80_INST_SRL},
	"RLD":  {Type: Z80_INST0, Literal: "RLD", Op: Z80_INST_RLD},
	"RRD":  {Type: Z80_INST0, Literal: "RRD", Op: Z80_INST_RRD},
	"BIT":  {Type: Z80_INST2, Literal: "BIT", Op: Z80_INST_BIT},
	"SET":  {Type: Z80_INST2, Literal: "SET", Op: Z80_INST_SET},
	"RES":  {Type: Z80_INST2, Literal: "RES", Op: Z80_INST_RES},
	"JP":   {Type: Z80_INST2, Literal: "JP", Op: Z80_INST_JP}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"JR":   {Type: Z80_INST2, Literal: "JR", Op: Z80_INST_JR}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"DJNZ": {Type: Z80_INST1, Literal: "DJNZ", Op: Z80_INST_DJNZ},
	"CALL": {Type: Z80_INST2, Literal: "CALL", Op: Z80_INST_CALL}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"RET":  {Type: Z80_INST0, Literal: "RET", Op: Z80_INST_RET},   // cc 有無により OPCODE0, OPCODE1 両方あり
	"RETI": {Type: Z80_INST0, Literal: "RETI", Op: Z80_INST_RETI},
	"RETN": {Type: Z80_INST0, Literal: "RETN", Op: Z80_INST_RETN},
	"RST":  {Type: Z80_INST1, Literal: "RST", Op: Z80_INST_REST},
	"IN":   {Type: Z80_INST2, Literal: "IN", Op: Z80_INST_IN},
	"INI":  {Type: Z80_INST0, Literal: "INI", Op: Z80_INST_INI},
	"INIR": {Type: Z80_INST0, Literal: "INIR", Op: Z80_INST_INIR},
	"INDR": {Type: Z80_INST0, Literal: "INDR", Op: Z80_INST_INDR},
	"OUT":  {Type: Z80_INST2, Literal: "OUT", Op: Z80_INST_OUT},
	"OUTI": {Type: Z80_INST0, Literal: "OUTI", Op: Z80_INST_OUTI},
	"OTIR": {Type: Z80_INST0, Literal: "OTIR", Op: Z80_INST_OTIR},
	"OUTD": {Type: Z80_INST0, Literal: "OUTD", Op: Z80_INST_OUTD},
	"OTDR": {Type: Z80_INST0, Literal: "OTDR", Op: Z80_INST_OTDR},
}

func Z80Names(opcode int) string {
	for _, v := range Z80Registers {
		if v.Op == opcode {
			return v.Literal
		}
	}
	for _, v := range Z80OpCodes {
		if v.Op == opcode {
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
