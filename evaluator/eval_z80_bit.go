package evaluator

import (
	"fmt"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
)

// BIT, SET RES
func (e *Evaluator) evalZ80_BIT(stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object {
	if op1 == nil || op2 == nil {
		e.logger.Error(errcode.EZ80_OP_LESS, stmt.Context)
		return object.ERROR
	}

	code := &object.CodeObject{Code: []byte{0xcb, 0x00}, TStates: [2]byte{8, 2}, Context: stmt.Context}
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
		return op1
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP1_NULL, stmt.Context)
		return object.ERROR
	case *object.NumberObject:
		b, ok := e.intToBit(op1.Value)
		if !ok {
			e.logger.Error(errcode.EZ80_BIT_NUM_RANGE, stmt.Context)
			return object.ERROR
		}
		bn = b
	default:
		e.logger.Error(errcode.EZ80_OP1, stmt.Context)
		return object.ERROR
	}
	code.Code[1] |= byte(bn << 3)

	// レジスタ or (HL),(IX+d),(IY+d)
	switch op2 := op2.(type) {
	case *object.RefNotFoundObject:
		return op2
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return object.ERROR

	case *object.RegisterObject: // BIT/SET/RES n, r
		if op2.Register < parser.Z80_REG_B || op2.Register > parser.Z80_REG_A {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op2.Register)), stmt.Context)
			return object.ERROR
		}
		if index, ok := Z80Reg8Index[op2.Register]; !ok {
			panic("unexpected register in BIT/SET/RES")
		} else {
			code.Code[1] |= byte(index)
			return code
		}

	case *object.RegIndirectObject: // BIT/SET/RES n, (HL)(IX+d)(IY+d)
		code.Code[1] |= 0x06
		switch op2.Register {
		case parser.Z80_REG_HL:
			if stmt.Opcode == parser.Z80_INST_BIT {
				code.TStates = [2]byte{12, 3}
			} else {
				code.TStates = [2]byte{15, 5}
			}
			return code
		case parser.Z80_REG_IX:
			code.Code = []byte{0xdd, 0xcb, byte(op2.Displacement), code.Code[1]}
			if stmt.Opcode == parser.Z80_INST_BIT {
				code.TStates = [2]byte{20, 5}
			} else {
				code.TStates = [2]byte{23, 7}
			}
			return code
		case parser.Z80_REG_IY:
			code.Code = []byte{0xfd, 0xcb, byte(op2.Displacement), code.Code[1]}
			if stmt.Opcode == parser.Z80_INST_BIT {
				code.TStates = [2]byte{20, 5}
			} else {
				code.TStates = [2]byte{23, 7}
			}
			return code
		default:
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(op2.Register)), stmt.Context)
			return object.ERROR
		}
	default:
		e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		return object.ERROR
	}
}
