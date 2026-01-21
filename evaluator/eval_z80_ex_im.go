package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

// EX
func (e *Evaluator) evalZ80_EX(stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0xeb}, CZ80: 4, Context: stmt.Context} // EX DE, HL
	if op1 == nil || op2 == nil {
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return code
	}
	if isRefNotFound(op1) || isRefNotFound(op2) {
		e.Resolved = false
		return code
	}
	if op1 == object.NULL {
		e.logger.Error(errcode.EZ80_OP1_NULL, stmt.Context)
		return code
	}
	if op2 == object.NULL {
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return code
	}

	switch {
	// EX rr, rr'
	case op1.Type() == object.REGISTER_OBJ && op2.Type() == object.REGISTER_OBJ:
		reg1 := op1.(*object.RegisterObject)
		reg2 := op2.(*object.RegisterObject)
		switch {
		case reg1.Register == parser.Z80_REG_DE && reg2.Register == parser.Z80_REG_HL, reg1.Register == parser.Z80_REG_HL && reg2.Register == parser.Z80_REG_DE:
			return code
		case reg1.Register == parser.Z80_REG_AF && reg2.Register == parser.Z80_REG_AFEX, reg1.Register == parser.Z80_REG_AFEX && reg2.Register == parser.Z80_REG_AF:
			code.Code[0] = 0x08
			return code
		default:
			e.logger.Error(errcode.EZ80_OP, stmt.Context)
		}
	// EX (SP), rr
	case op1.Type() == object.REG_INDIRECT_OBJ && op2.Type() == object.REGISTER_OBJ:
		return e.evalExSpReg16(code, op1, op2, stmt.Context)
	// EX rr, (SP), => EX (SP), rr
	case op1.Type() == object.REGISTER_OBJ && op2.Type() == object.REG_INDIRECT_OBJ:
		return e.evalExSpReg16(code, op2, op1, stmt.Context)
	}

	e.logger.Error(errcode.EZ80_OP, stmt.Context)
	return code
}

func (e *Evaluator) evalExSpReg16(code *object.CodeObject, op1, op2 object.Object, ctx TContext) object.Object {
	code.Code[0] = 0xe3
	code.CZ80 = 19

	reg1 := op1.(*object.RegIndirectObject)
	if reg1.Register != parser.Z80_REG_SP {
		e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(reg1.Register)), ctx)
		return code // EX (SP), HL
	}
	reg2 := op2.(*object.RegisterObject)
	switch reg2.Register {
	case parser.Z80_REG_HL:
		return code // EX (SP), HL
	case parser.Z80_REG_IX:
		code.Code = []byte{0xdd, 0xe3}
		code.CZ80 = 23
		return code
	case parser.Z80_REG_IY:
		code.Code = []byte{0xfd, 0xe3}
		code.CZ80 = 23
		return code
	default:
		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(reg2.Register)), ctx)
		return code
	}
}

// IM
func (e *Evaluator) evalZ80_IM(stmt *parser.Z80Instruction, op, _ object.Object, env TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0xed, 0x46}, CZ80: 8, Context: stmt.Context}

	// IM n
	mode := 0
	switch op := op.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP_NULL, stmt.Context)
		return object.ERROR
	case *object.NumberObject:
		mode = op.Value
		if mode < 0 || mode > 2 {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_IM_RANGE, mode, mode), stmt.Context)
			return code
		}
		switch mode {
		case 1:
			code.Code[1] = 0x56
		case 2:
			code.Code[1] = 0x5E
		}
		return code
	default:
		// op = nil
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return code
	}
}
