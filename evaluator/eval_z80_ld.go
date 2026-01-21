package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) evalZ80LD(stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object {
	switch op1 := op1.(type) {
	case *object.RegisterObject:
		switch op1.RegisterType {
		case parser.Z80_REG8:
			// LD r, x
			return e.evalZ80LD_reg8(stmt, op1, op2, env)
		case parser.Z80_REG16:
			// LD rr, x
			return e.evalZ80LD_reg16(stmt, op1, op2, env)
		}
	// case *object.IndirectExpression:
	// 	e.logger.Error(fmt.Sprintf(errcode.E999, node), node.Context.LineNumber)
	// 	return object.ERROR
	case *object.RefNotFoundObject:
		// 仮として LD A,0 を返す
		return &object.CodeObject{Code: []byte{0x3e, 0}, Context: stmt.Context}
	}
	e.logger.Error(errcode.EZ80_OP1, stmt.Context)
	return object.ERROR
}

func (e *Evaluator) evalZ80LD_reg8(node *parser.Z80Instruction, op1 *object.RegisterObject, op2 object.Object, env TEnv) object.Object {

	switch op2 := op2.(type) {
	case *object.RegisterObject:
		// LD r, r'
		if op2.RegisterType != parser.Z80_REG8 {
			e.logger.Error(errcode.EZ80_OP2, node.Context)
			return object.ERROR
		}
		r1, ok := Z80Reg8Index[int(op1.Register)]
		if !ok {
			e.logger.Error(errcode.EZ80_OP1, node.Context)
			return object.ERROR
		}
		r2, ok := Z80Reg8Index[int(op2.Register)]
		if !ok {
			e.logger.Error(errcode.EZ80_OP2, node.Context)
			return object.ERROR
		}

		b := 0x40 | r1<<3 | r2
		return &object.CodeObject{Code: []byte{byte(b)}, Context: node.Context}
	case *object.NumberObject:
		// LD r, n
		v, ok := e.intToByte(op2.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_BYTE, op2.Value, op2.Value), node.Context)
		}
		r1 := Z80Reg8Index[int(op1.Register)]
		b := byte(0x06 | (r1 << 3))
		return &object.CodeObject{Code: []byte{b, v}, Context: node.Context}
	// case *object.IndirectExpression:
	// 	e.logger.Error(fmt.Sprintf(errcode.E999, node), node.Context.LineNumber)
	// 	return object.ERROR
	case *object.RefNotFoundObject:
		// 未確定の場合として LD A,0 を返す
		e.Resolved = false
		return &object.CodeObject{Code: []byte{0x3e, 00}, Context: node.Context}

	default:
		e.logger.Error(errcode.EZ80_OP2, node.Context)
		return object.ERROR
	}
}

func (e *Evaluator) evalZ80LD_reg16(node *parser.Z80Instruction, op1 *object.RegisterObject, op2 object.Object, env TEnv) object.Object {

	switch op2 := op2.(type) {
	case *object.RegisterObject:
		// LD rr, rr'
		if op2.RegisterType != parser.Z80_REG16 {
			e.logger.Error(errcode.EZ80_OP2, node.Context)
			return object.ERROR
		}
		if op1.Register != parser.Z80_REG_SP {
			e.logger.Error(errcode.EZ80_OP1_SP, node.Context)
			return object.ERROR
		}
		switch op2.Register {
		case parser.Z80_REG_HL:
			return &object.CodeObject{Code: []byte{0xf9}, Context: node.Context}
		case parser.Z80_REG_IX:
			return &object.CodeObject{Code: []byte{0xdd, 0xf9}, Context: node.Context}
		case parser.Z80_REG_IY:
			return &object.CodeObject{Code: []byte{0xfd, 0xf9}, Context: node.Context}
		default:
			e.logger.Error(errcode.EZ80_OP2_HL_IXY, node.Context)
			return object.ERROR
		}
	case *object.NumberObject:
		// LD rr, nn
		v, ok := e.intToWord(op2.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, op2.Value, op2.Value), node.Context)
		}
		r1 := Z80Reg16IndexSP[int(op1.Register)]
		b := byte(0x01 | (r1 << 4))
		return &object.CodeObject{Code: []byte{b, byte(v & 0xff), byte((v >> 8) & 0xff)}, Context: node.Context}
	// case *object.IndirectExpression:
	// 	e.logger.Error(fmt.Sprintf(errcode.E999, node), node.Context.LineNumber)
	// 	return object.ERROR
	case *object.RefNotFoundObject:
		// 未確定として LD HL,0 を返す
		e.Resolved = false
		return &object.CodeObject{Code: []byte{0x21, 0, 0}, Context: node.Context}

	default:
		e.logger.Error(errcode.EZ80_OP2, node.Context)
		return object.ERROR
	}
}
