package parser

import (
	"github.com/dogatana/yas80/intern"
)

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
	Z80_REG_F // IN F,(C) のみ

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
	Z80_INST_MUL
)

func Z80Opcode2Name(opcode int) string {
	for _, v := range z80ReservedWords {
		if int(v.TokenSubType) == opcode {
			return v.SymbolID.String()
		}
	}
	return "UNKNOWN"
}

func init() {
	z80ReservedWords = make(map[intern.SymbolID]Token, len(_z80ReservedWords))

	for s, tt := range _z80ReservedWords {
		id := intern.Intern(s)
		z80ReservedWords[id] = Token{TokenType: tt.Type, TokenSubType: tt.SubType, SymbolID: id}
	}

}

var z80ReservedWords map[intern.SymbolID]Token

var _z80ReservedWords = map[string]struct {
	Type    TokenType
	SubType TokenSubType
}{
	"A":   {Type: Z80_REG8, SubType: Z80_REG_A},
	"B":   {Type: Z80_REG8, SubType: Z80_REG_B},
	"C":   {Type: Z80_REG8, SubType: Z80_REG_C},
	"D":   {Type: Z80_REG8, SubType: Z80_REG_D},
	"E":   {Type: Z80_REG8, SubType: Z80_REG_E},
	"H":   {Type: Z80_REG8, SubType: Z80_REG_H},
	"L":   {Type: Z80_REG8, SubType: Z80_REG_L},
	"IXH": {Type: Z80_REG8, SubType: Z80_REG_IXH},
	"IXL": {Type: Z80_REG8, SubType: Z80_REG_IXL},
	"IYH": {Type: Z80_REG8, SubType: Z80_REG_IYH},
	"IYL": {Type: Z80_REG8, SubType: Z80_REG_IYL},
	"I":   {Type: Z80_REG8, SubType: Z80_REG_I},
	"R":   {Type: Z80_REG8, SubType: Z80_REG_R},
	"F":   {Type: Z80_REG8, SubType: Z80_REG_F},

	"SP":  {Type: Z80_REG16, SubType: Z80_REG_SP},
	"IX":  {Type: Z80_REG16, SubType: Z80_REG_IX},
	"IY":  {Type: Z80_REG16, SubType: Z80_REG_IY},
	"AF":  {Type: Z80_REG16, SubType: Z80_REG_AF},
	"AF'": {Type: Z80_REG16, SubType: Z80_REG_AFEX},
	"BC":  {Type: Z80_REG16, SubType: Z80_REG_BC},
	"DE":  {Type: Z80_REG16, SubType: Z80_REG_DE},
	"HL":  {Type: Z80_REG16, SubType: Z80_REG_HL},

	"CY": {Type: Z80_FLAG, SubType: Z80_FLAG_C}, // キャリーの別名を定義
	"NC": {Type: Z80_FLAG, SubType: Z80_FLAG_NC},
	"Z":  {Type: Z80_FLAG, SubType: Z80_FLAG_Z},
	"NZ": {Type: Z80_FLAG, SubType: Z80_FLAG_NZ},
	"PO": {Type: Z80_FLAG, SubType: Z80_FLAG_PO},
	"PE": {Type: Z80_FLAG, SubType: Z80_FLAG_PE},
	"P":  {Type: Z80_FLAG, SubType: Z80_FLAG_P},
	"M":  {Type: Z80_FLAG, SubType: Z80_FLAG_M},

	// Opcode
	"LD":   {Type: Z80_INST2, SubType: Z80_INST_LD},
	"PUSH": {Type: Z80_INST1, SubType: Z80_INST_PUSH},
	"POP":  {Type: Z80_INST1, SubType: Z80_INST_POP},
	"EX":   {Type: Z80_INST2, SubType: Z80_INST_EX},
	"EXX":  {Type: Z80_INST0, SubType: Z80_INST_EXX},
	"LDI":  {Type: Z80_INST0, SubType: Z80_INST_LDI},
	"LDIR": {Type: Z80_INST0, SubType: Z80_INST_LDIR},
	"LDD":  {Type: Z80_INST0, SubType: Z80_INST_LDD},
	"LDDR": {Type: Z80_INST0, SubType: Z80_INST_LDDR},
	"CPI":  {Type: Z80_INST0, SubType: Z80_INST_CPI},
	"CPIR": {Type: Z80_INST0, SubType: Z80_INST_CPIR},
	"CPD":  {Type: Z80_INST0, SubType: Z80_INST_CPD},
	"CPDR": {Type: Z80_INST0, SubType: Z80_INST_CPDR},
	"ADD":  {Type: Z80_INST2, SubType: Z80_INST_ADD},
	"ADC":  {Type: Z80_INST2, SubType: Z80_INST_ADC},
	"SUB":  {Type: Z80_INST2, SubType: Z80_INST_SUB},
	"SBC":  {Type: Z80_INST2, SubType: Z80_INST_SBC},
	"AND":  {Type: Z80_INST2, SubType: Z80_INST_AND},
	"OR":   {Type: Z80_INST2, SubType: Z80_INST_OR},
	"XOR":  {Type: Z80_INST2, SubType: Z80_INST_XOR},
	"CP":   {Type: Z80_INST2, SubType: Z80_INST_CP},
	"INC":  {Type: Z80_INST1, SubType: Z80_INST_INC},
	"DEC":  {Type: Z80_INST1, SubType: Z80_INST_DEC},
	"DAA":  {Type: Z80_INST0, SubType: Z80_INST_DAA},
	"CPL":  {Type: Z80_INST0, SubType: Z80_INST_CPL},
	"NEG":  {Type: Z80_INST0, SubType: Z80_INST_NEG},
	"CCF":  {Type: Z80_INST0, SubType: Z80_INST_CCF},
	"SCF":  {Type: Z80_INST0, SubType: Z80_INST_SCF},
	"NOP":  {Type: Z80_INST0, SubType: Z80_INST_NOP},
	"HALT": {Type: Z80_INST0, SubType: Z80_INST_HALT},
	"DI":   {Type: Z80_INST0, SubType: Z80_INST_DI},
	"EI":   {Type: Z80_INST0, SubType: Z80_INST_EI},
	"IM":   {Type: Z80_INST1, SubType: Z80_INST_IM},
	"RLCA": {Type: Z80_INST0, SubType: Z80_INST_RLCA},
	"RLA":  {Type: Z80_INST0, SubType: Z80_INST_RLA},
	"RRCA": {Type: Z80_INST0, SubType: Z80_INST_RRCA},
	"RRA":  {Type: Z80_INST0, SubType: Z80_INST_RRA},
	"RLC":  {Type: Z80_INST1, SubType: Z80_INST_RLC},
	"RL":   {Type: Z80_INST1, SubType: Z80_INST_RL},
	"RR":   {Type: Z80_INST1, SubType: Z80_INST_RR},
	"RRC":  {Type: Z80_INST1, SubType: Z80_INST_RRC},
	"SLA":  {Type: Z80_INST1, SubType: Z80_INST_SLA},
	"SRA":  {Type: Z80_INST1, SubType: Z80_INST_SRA},
	"SRL":  {Type: Z80_INST1, SubType: Z80_INST_SRL},
	"RLD":  {Type: Z80_INST0, SubType: Z80_INST_RLD},
	"RRD":  {Type: Z80_INST0, SubType: Z80_INST_RRD},
	"BIT":  {Type: Z80_INST2, SubType: Z80_INST_BIT},
	"SET":  {Type: Z80_INST2, SubType: Z80_INST_SET},
	"RES":  {Type: Z80_INST2, SubType: Z80_INST_RES},
	"JP":   {Type: Z80_INST2, SubType: Z80_INST_JP}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"JR":   {Type: Z80_INST2, SubType: Z80_INST_JR}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"DJNZ": {Type: Z80_INST1, SubType: Z80_INST_DJNZ},
	"CALL": {Type: Z80_INST2, SubType: Z80_INST_CALL}, // cc 有無によりOPCODE1, OPCODE2 両方あり
	"RET":  {Type: Z80_INST1, SubType: Z80_INST_RET},  // cc 有無により OPCODE0, OPCODE1 両方あり
	"RETI": {Type: Z80_INST0, SubType: Z80_INST_RETI},
	"RETN": {Type: Z80_INST0, SubType: Z80_INST_RETN},
	"RST":  {Type: Z80_INST1, SubType: Z80_INST_RST},
	"IN":   {Type: Z80_INST2, SubType: Z80_INST_IN},
	"INI":  {Type: Z80_INST0, SubType: Z80_INST_INI},
	"INIR": {Type: Z80_INST0, SubType: Z80_INST_INIR},
	"IND":  {Type: Z80_INST0, SubType: Z80_INST_IND},
	"INDR": {Type: Z80_INST0, SubType: Z80_INST_INDR},
	"OUT":  {Type: Z80_INST2, SubType: Z80_INST_OUT},
	"OUTI": {Type: Z80_INST0, SubType: Z80_INST_OUTI},
	"OTIR": {Type: Z80_INST0, SubType: Z80_INST_OTIR},
	"OUTD": {Type: Z80_INST0, SubType: Z80_INST_OUTD},
	"OTDR": {Type: Z80_INST0, SubType: Z80_INST_OTDR},
	"MUL":  {Type: Z80_INST2, SubType: Z80_INST_MUL},
}
