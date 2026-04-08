package evaluator

import (
	"fmt"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
)

func (e *Evaluator) evalZ80_INC_DEC(stmt *parser.Z80Instruction, op, _ object.Object, env TEnv) object.Object {

	if op == nil {
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return object.ERROR
	}

	switch op := op.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return op
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP_NULL, stmt.Context)
		return object.ERROR

	case *object.RegisterObject:

		if op.RegisterType == parser.Z80_REG8 {
			// 8 bit register
			code := &object.CodeObject{Code: []byte{0x04}, TStates: [2]byte{4, 1}, Context: stmt.Context}
			if stmt.Opcode == parser.Z80_INST_DEC {
				code.Code[0] = 0x05 // DEC
			}
			if index, ok := Z80Reg8Index[op.Register]; ok {
				code.Code[0] |= index << 3
				return code
			}

			code.TStates = [2]byte{8, 2}
			if index, ok := Z80Reg8IndexIX[op.Register]; ok {
				code.Code = []byte{0xdd, code.Code[0] | index<<3}
				return code
			}
			if index, ok := Z80Reg8IndexIY[op.Register]; ok {
				code.Code = []byte{0xfd, code.Code[0] | index<<3}
				return code
			}
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op.Register)), stmt.Context)
			return object.ERROR
		}
		// 16 bit register
		code := &object.CodeObject{Code: []byte{0x03}, TStates: [2]byte{6, 1}, Context: stmt.Context}
		if stmt.Opcode == parser.Z80_INST_DEC {
			code.Code[0] = 0x0b // DEC
		}
		if index, ok := Z80Reg16IndexSP[op.Register]; ok {
			code.Code[0] |= index << 4
			return code // INC rr / DEC rr
		}

		code.TStates = [2]byte{10, 2}
		code.Code[0] |= 0x20
		switch op.Register {

		case parser.Z80_REG_IX: // INC/DEC IX
			code.Code = []byte{0xdd, code.Code[0]}
			return code

		case parser.Z80_REG_IY: // INC/DEC IY
			code.Code = []byte{0xfd, code.Code[0]}
			return code
		}
		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op.Register)), stmt.Context)
		return object.ERROR

	case *object.RegIndirectObject:
		code := &object.CodeObject{Code: []byte{0x034}, TStates: [2]byte{11, 4}, Context: stmt.Context}
		if stmt.Opcode == parser.Z80_INST_DEC {
			code.Code[0] = 0x35 // DEC
		}

		switch op.Register {
		case parser.Z80_REG_HL: // INC/DEC (HL)
			if op.Displacement != 0 {
				e.logger.Error(errcode.EINDIRECT_DISP_REG, stmt.Context)
				return object.ERROR
			}
			return code
		case parser.Z80_REG_IX: // INC/DEC (IX + d)
			code.Code = []byte{0xdd, code.Code[0], byte(op.Displacement)}
			code.TStates = [2]byte{23, 7}
			return code

		case parser.Z80_REG_IY: // INC/DEC (IY + d)
			code.Code = []byte{0xfd, code.Code[0], byte(op.Displacement)}
			code.TStates = [2]byte{23, 7}
			return code

		default:
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(op.Register)), stmt.Context)
			return object.ERROR
		}
	}
	e.logger.Error(errcode.EZ80_OP, stmt.Context)
	return object.ERROR
}

func (e *Evaluator) evalZ80_PUSH_POP(stmt *parser.Z80Instruction, op, _ object.Object, env TEnv) object.Object {

	if op == nil {
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return object.ERROR
	}

	switch op := op.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return op
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP_NULL, stmt.Context)
		return object.ERROR

	case *object.RegisterObject:
		// PUSH rr / POP rr
		code := &object.CodeObject{Code: []byte{0xc5}, TStates: [2]byte{11, 4}, Context: stmt.Context} // PUSH BC
		if stmt.Opcode == parser.Z80_INST_POP {
			code.Code[0] = 0xc1 // POP BC
			code.TStates = [2]byte{10, 3}
		}
		index, ok := Z80Reg16IndexAF[op.Register]
		if ok {
			code.Code[0] |= index << 4
			return code
		}

		// PUSH IX, IY / POP IX, IY
		if stmt.Opcode == parser.Z80_INST_PUSH {
			code.Code = []byte{0xdd, 0xe5}
			code.TStates = [2]byte{15, 5}
		} else {
			code.Code = []byte{0xdd, 0xe1}
			code.TStates = [2]byte{14, 4}
		}
		switch op.Register {
		case parser.Z80_REG_IX:
			return code
		case parser.Z80_REG_IY:
			code.Code[0] = 0xfd
			return code
		}
		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op.Register)), stmt.Context)
		return object.ERROR

	default:
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return object.ERROR
	}
}
