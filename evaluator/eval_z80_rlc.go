package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) evalZ80_RLC(stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object {
	codeTable := map[int]byte{
		parser.Z80_INST_RLC: 0x00,
		parser.Z80_INST_RL:  0x10,
		parser.Z80_INST_RRC: 0x08,
		parser.Z80_INST_RR:  0x18,
		parser.Z80_INST_SLA: 0x20,
		parser.Z80_INST_SRA: 0x28,
		parser.Z80_INST_SRL: 0x38,
	}
	code := &object.CodeObject{Code: []byte{0xcb, 0x00}, CZ80: 8, Context: stmt.Context}
	if b, ok := codeTable[stmt.Opcode]; ok {
		code.Code[1] = b
	} else {
		panic("invalid opecode")
	}

	// レジスタ or (HL),(IX+d),(IY+d)
	switch op1 := op1.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP_NULL, stmt.Context)
		return code
	case *object.RegisterObject:
		if op1.Register < parser.Z80_REG_B || op1.Register > parser.Z80_REG_A {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op1.Register)), stmt.Context)
			return code
		}
		if index, ok := Z80Reg8Index[op1.Register]; !ok {
			panic("unexpected register in RLC...")
		} else {
			code.Code[1] |= byte(index)
			return code
		}
	case *object.RegIndirectObject:
		code.Code[1] |= 0x06
		code.CZ80 = 15
		switch op1.Register {
		case parser.Z80_REG_HL:
			return code
		case parser.Z80_REG_IX:
			code.Code[1] |= 0x06
			code.Code = []byte{0xdd, 0xcb, byte(op1.Displacement), code.Code[1]}
			code.CZ80 = 23
			return code
		case parser.Z80_REG_IY:
			code.Code[1] |= 0x06
			code.Code = []byte{0xfd, 0xcb, byte(op1.Displacement), code.Code[1]}
			code.CZ80 = 23
			return code
		default:
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(op1.Register)), stmt.Context)
			return code
		}
	default:
		e.logger.Error(errcode.EZ80_OP1, stmt.Context)
		return code
	}
}
