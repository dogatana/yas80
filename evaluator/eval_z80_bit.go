package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

// BIT, SET RES
func (e *Evaluator) evalZ80_BIT(stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0xcb, 0x00}, Context: stmt.Context}
	switch stmt.Opcode {
	case parser.Z80_INST_BIT:
		code.Code[1] = 0x40
	case parser.Z80_INST_SET:
		code.Code[1] = 0xc0
	case parser.Z80_INST_RES:
		code.Code[1] = 0x80
	}

	// bit 番号
	bn := byte(0)
	switch op1 := op1.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP1_NULL, stmt.Context)
		return object.ERROR
	case *object.NumberObject:
		b, ok := e.intToBit(op1.Value)
		if !ok {
			e.logger.Error(errcode.EZ80_BIT_NUM_RANGE, stmt.Context)
			return code
		}
		bn = b
	default:
		e.logger.Error(errcode.EZ80_OP1, stmt.Context)
		return code
	}
	code.Code[1] |= byte(bn << 3)

	// レジスタ or (HL),(IX+d),(IY+d)
	switch op2 := op2.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return code
	case *object.RegisterObject:
		if op2.Register < parser.Z80_REG_B || op2.Register > parser.Z80_REG_A {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op2.Register)), stmt.Context)
			return code
		}
		if index, ok := Z80Reg8Index[op2.Register]; !ok {
			panic("unexpected register in BIT/SET/RES")
		} else {
			code.Code[1] |= byte(index)
			return code
		}
	case *object.RegIndirectObject:
		code.Code[1] |= 0x06
		switch op2.Register {
		case parser.Z80_REG_HL:
			return code
		case parser.Z80_REG_IX:
			code.Code[1] |= 0x06
			code.Code = []byte{0xdd, 0xcb, byte(op2.Displacement), code.Code[1]}
			return code
		case parser.Z80_REG_IY:
			code.Code[1] |= 0x06
			code.Code = []byte{0xfd, 0xcb, byte(op2.Displacement), code.Code[1]}
			return code
		default:
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(op2.Register)), stmt.Context)
			return code
		}
	default:
		e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		return code
	}
}
