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

	switch port := op2.(type) {
	case *object.RegIndirectObject:
		if port.Register != parser.Z80_REG_C {
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(port.Register)), stmt.Context)
			return code
		}
		code.Code[0] = 0xed
		code.CZ80 = 12
		if index, ok := Z80Reg8Index[reg.Register]; !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(reg.Register)), stmt.Context)
			return code
		} else {
			code.Code[1] = 0x40 | index<<3
			return code
		}
	case *object.AddrIndirectObject:
		if reg.Register != parser.Z80_REG_A {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(reg.Register)), stmt.Context)
			return code
		}
		addr, ok := e.intToPort(port.Address)
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_PORT_RANGE, port.Address, port.Address), stmt.Context)
		}
		code.Code[1] = addr
		return code

	default:
		e.logger.Error(errcode.ENOT_IMPL_STMT, stmt.Context)
		return code
	}
}

func (e *Evaluator) evalZ80_OUT(stmt *parser.Z80Instruction, op1, op2 object.Object, _ TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0xd3, 0x00}, CZ80: 11, Context: stmt.Context}

	var reg *object.RegisterObject
	switch obj := op2.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return code
	case *object.RegisterObject:
		reg = obj
	default:
		e.logger.Error(errcode.EZ80_OP1, stmt.Context)
		return code
	}

	switch port := op1.(type) {
	case *object.RegIndirectObject:
		if port.Register != parser.Z80_REG_C {
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(port.Register)), stmt.Context)
			return code
		}
		code.Code[0] = 0xed
		code.CZ80 = 12
		if index, ok := Z80Reg8Index[reg.Register]; !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(reg.Register)), stmt.Context)
			return code
		} else {
			code.Code[1] = 0x41 | index<<3
			return code
		}
	case *object.AddrIndirectObject:
		if reg.Register != parser.Z80_REG_A {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(reg.Register)), stmt.Context)
			return code
		}
		addr, ok := e.intToPort(port.Address)
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_PORT_RANGE, port.Address, port.Address), stmt.Context)
		}
		code.Code[1] = addr
		return code

	default:
		e.logger.Error(errcode.ENOT_IMPL_STMT, stmt.Context)
		return code
	}
}
