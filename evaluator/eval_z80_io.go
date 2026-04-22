package evaluator

import (
	"fmt"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
)

func (e *Evaluator) evalZ80_IN(stmt *parser.Z80Instruction, op1, op2 object.Object, _ TEnv) object.Object {
	if op1 == nil || op2 == nil {
		e.logger.Error(errcode.EZ80_OP_LESS, stmt.Context)
		return object.ERROR
	}

	var reg *object.RegisterObject
	switch obj := op1.(type) {
	case *object.RefNotFoundObject:
		return obj
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP1_NULL, stmt.Context)
		return object.ERROR
	case *object.RegisterObject:
		reg = obj
	default:
		e.logger.Error(errcode.EZ80_OP1, stmt.Context)
		return object.ERROR
	}

	switch port := op2.(type) {
	case *object.RegIndirectObject: // IN r, (C)
		if port.Register != parser.Z80_REG_C {
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(port.Register)), stmt.Context)
			return object.ERROR
		}
		if index, ok := Z80Reg8IndexIN[reg.Register]; !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(reg.Register)), stmt.Context)
			return object.ERROR
		} else {
			return &object.CodeObject{Code: []byte{0xed, 0x40 | index<<3}, TStates: [2]byte{12, 3}, Context: stmt.Context}
		}

	case *object.AddrIndirectObject: // IN A, (n)
		if reg.Register != parser.Z80_REG_A {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(reg.Register)), stmt.Context)
			return object.ERROR
		}
		addr, ok := e.intToPort(port.Address)
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_PORT_RANGE, port.Address, port.Address), stmt.Context)
			return object.ERROR
		}
		return &object.CodeObject{Code: []byte{0xdb, addr}, TStates: [2]byte{11, 3}, Context: stmt.Context}

	default:
		e.logger.Error(errcode.ENOT_IMPL_STMT, stmt.Context)
		return object.ERROR
	}
}

func (e *Evaluator) evalZ80_OUT(stmt *parser.Z80Instruction, op1, op2 object.Object, _ TEnv) object.Object {
	if op1 == nil || op2 == nil {
		e.logger.Error(errcode.EZ80_OP_LESS, stmt.Context)
		return object.ERROR
	}

	var reg *object.RegisterObject
	switch obj := op2.(type) {
	case *object.RefNotFoundObject:
		return obj
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return object.ERROR
	case *object.RegisterObject:
		reg = obj
	default:
		e.logger.Error(errcode.EZ80_OP1, stmt.Context)
		return object.ERROR
	}

	switch port := op1.(type) {
	case *object.RegIndirectObject: // OUT (C), r
		if port.Register != parser.Z80_REG_C {
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(port.Register)), stmt.Context)
			return object.ERROR
		}
		if index, ok := Z80Reg8Index[reg.Register]; !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(reg.Register)), stmt.Context)
			return object.ERROR
		} else {
			return &object.CodeObject{Code: []byte{0xed, 0x41 | index<<3}, TStates: [2]byte{12, 3}, Context: stmt.Context}
		}
	case *object.AddrIndirectObject: // OUT (n), r
		if reg.Register != parser.Z80_REG_A {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(reg.Register)), stmt.Context)
			return object.ERROR
		}
		addr, ok := e.intToPort(port.Address)
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_PORT_RANGE, port.Address, port.Address), stmt.Context)
			return object.ERROR
		}
		return &object.CodeObject{Code: []byte{0xd3, addr}, TStates: [2]byte{11, 3}, Context: stmt.Context}

	default:
		e.logger.Error(errcode.ENOT_IMPL_STMT, stmt.Context)
		return object.ERROR
	}
}
