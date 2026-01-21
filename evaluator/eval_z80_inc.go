package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) evalZ80_INC_DEC(stmt *parser.Z80Instruction, op, _ object.Object, env TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0x04}, CZ80: 4, Context: stmt.Context} // INC B
	if stmt.Opcode == parser.Z80_INST_DEC {
		code.Code[0] = 0x05 // DEC B
	}

	if op == nil {
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return code
	}

	switch op := op.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP_NULL, stmt.Context)
		return code

	case *object.RegisterObject:
		if op.RegisterType == parser.Z80_REG8 {
			// INC/DEC r
			index, ok := Z80Reg8Index[op.Register]
			if !ok {
				e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op.Register)), stmt.Context)
				return code
			}
			code.Code[0] |= index << 3
			return code
		}
		// INC rr
		if stmt.Opcode == parser.Z80_INST_INC {
			code.Code[0] = 0x03
		} else {
			code.Code[0] = 0x0b
		}
		if index, ok := Z80Reg16IndexSP[op.Register]; ok {
			code.Code[0] |= index << 4
			code.CZ80 = 6
			return code
		}

		code.CZ80 = 10
		switch op.Register {

		// INC/DEC IX
		case parser.Z80_REG_IX:
			code.Code = []byte{0xdd, code.Code[0] | 0x20}
			return code

		// INC/DEC IY
		case parser.Z80_REG_IY:
			code.Code = []byte{0xfd, code.Code[0] | 0x20}
			return code
		}
		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op.Register)), stmt.Context)
		return code

	case *object.RegIndirectObject:
		code.Code[0] |= 0x30 // INC/DEC (HL)
		code.CZ80 = 11

		switch op.Register {
		// INC/DEC (HL)
		case parser.Z80_REG_HL:
			if op.Displacement != 0 {
				e.logger.Error(errcode.EINDIRECT_DISP_REG, stmt.Context)
			}
			return code
		// INC/DEC (IX + d)
		case parser.Z80_REG_IX:
			code.Code = []byte{0xdd, code.Code[0], byte(op.Displacement)}
			code.CZ80 = 23
			return code

		// INC/DEC (IY + d)
		case parser.Z80_REG_IY:
			code.Code = []byte{0xfd, code.Code[0], byte(op.Displacement)}
			code.CZ80 = 23
			return code

		default:
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(op.Register)), stmt.Context)
			return code
		}
	}
	e.logger.Error(errcode.EZ80_OP, stmt.Context)
	return code
}

func (e *Evaluator) evalZ80_PUSH_POP(stmt *parser.Z80Instruction, op, _ object.Object, env TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0xc5}, CZ80: 11, Context: stmt.Context} // PUSH BC
	if stmt.Opcode == parser.Z80_INST_POP {
		code.Code[0] = 0xc1 // POP BC
		code.CZ80 = 10
	}

	if op == nil {
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return code
	}

	switch op := op.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP_NULL, stmt.Context)
		return code

	case *object.RegisterObject:
		index, ok := Z80Reg16IndexAF[op.Register]
		if ok {
			code.Code[0] |= index << 4
			return code
		}

		code.Code[0] |= 0x20 // HL
		code.CZ80 += 4
		switch op.Register {
		case parser.Z80_REG_IX:
			code.Code = []byte{0xDD, code.Code[0]}
			return code
		case parser.Z80_REG_IY:
			code.Code = []byte{0xFD, code.Code[0]}
			return code
		}
		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op.Register)), stmt.Context)
		return code
	default:
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return code
	}
}
