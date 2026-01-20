package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) evalZ80_IN(stmt *parser.Z80Instruction, op1, op2 object.Object, _ TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0xdb, 0x00}, CZ80: 11, Context: stmt.Context}

	var reg *object.RegisterObject
	switch obj := op1.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP1_NULL, stmt.Context)
		return code
	case *object.RegisterObject:
		reg = obj
	default:
		e.logger.Error(errcode.EZ80_OP1, stmt.Context)
		return code
	}
	reg, ok := op1.(*object.RegisterObject)
	if !ok {
		e.logger.Error(errcode.EZ80_OP1, stmt.Context)
		return code
	}

	switch port := op2.(type) {
	case *object.RegIndirectObject:
		if port.Register != parser.Z80_REG_C {
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(port.Register)), stmt.Context)
			return code
		}
		code.Code[0] = 0xed
		if index, ok := Z80Reg8Index[reg.Register]; !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(reg.Register)), stmt.Context)
			return code
		} else {
			code.Code[1] = 0x40 | index<<3
			return code
		}
	default:
		e.logger.Error(errcode.ENOT_IMPL_STMT, stmt.Context)
		return code
	}
}
