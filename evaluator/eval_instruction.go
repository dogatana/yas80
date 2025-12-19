package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) evalZ80Instruction(node *parser.Z80Instruction, env object.Environment) object.Object {
	if node.Label != nil {
		e.evalLabel(node.Label, env)
	}
	switch node.NodeType() {
	case parser.Z80_INST0:
		info := Z80CodeTable0[int(node.Opcode)]
		obj := &object.CodeObject{Line: node.Context.Line, Code: make([]byte, len(info.Bytes))}
		copy(obj.Code, info.Bytes)
		return obj
	case parser.Z80_INST1:
		if node.Opcode == parser.Z80_INST_RET {
			return e.generateRET(node, env)
		}
		return object.NULL
	case parser.Z80_INST2:
		return e.evalZ80Instruction2(node, env)
	default:
		return object.NULL
	}
}

func (e *Evaluator) generateRET(node *parser.Z80Instruction, _ object.Environment) object.Object {
	op1 := node.Op1
	// RET
	if op1 == nil {
		return &object.CodeObject{Line: node.Context.Line, Code: []byte{0xc9}}
	}
	// RET cc
	index := -1
	if op1.NodeType() == parser.Z80_FLAG || op1.NodeType() == parser.Z80_REG8 {
		index = int(op1.NodeSubType())
	}
	flag, ok := Z80FlagIndex[index]
	if !ok {
		e.logger.Error(fmt.Sprintf(errcode.EZ80_FLAG, node.Op1.String()), node.Context)
		return object.ERROR
	}
	b := byte(0xc0 | flag<<3)
	return &object.CodeObject{Line: node.Context.Line, Code: []byte{b}}
}

func (e *Evaluator) evalZ80Instruction2(node *parser.Z80Instruction, env object.Environment) object.Object {
	switch node.NodeSubType() {
	case parser.Z80_INST_LD:
		return e.evalZ80LD(node, env)
	default:
		e.logger.Error(fmt.Sprintf(errcode.ENOT_IMPL_STMT, node), node.Context)
		return object.ERROR
	}
}

func (e *Evaluator) evalZ80LD(node *parser.Z80Instruction, env object.Environment) object.Object {
	op1 := e.evalExpression(node.Op1, env, node.Context)

	switch op1 := op1.(type) {
	case *object.RegisterObject:
		switch op1.RegisterType {
		case parser.Z80_REG8:
			// LD r, x
			return e.evalZ80LD_reg8(node, op1, env)
		case parser.Z80_REG16:
			// LD rr, x
			return e.evalZ80LD_reg16(node, op1, env)
		}
		e.logger.Error(errcode.EZ80_OP1, node.Context)
		return object.ERROR
	// case *object.IndirectExpression:
	// 	e.logger.Error(fmt.Sprintf(errcode.E999, node), node.Context.LineNumber)
	// 	return object.ERROR
	default:
		return &object.CodeObject{Line: node.Context.Line, Code: []byte{0x7f}}
	}
}

func (e *Evaluator) evalZ80LD_reg8(node *parser.Z80Instruction, op1 *object.RegisterObject, env object.Environment) object.Object {
	op2 := e.evalExpression(node.Op2, env, node.Context)

	switch op2 := op2.(type) {
	case *object.RegisterObject:
		// LD r, r'
		if op2.RegisterType != parser.Z80_REG8 {
			e.logger.Error(errcode.EZ80_OP2, node.Context)
			return object.ERROR
		}
		r1, ok := Z80Reg8Index[int(op1.Register)]
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.E028, parser.TokenLiteral(op1.Register)), node.Context)
			return object.ERROR
		}
		r2, ok := Z80Reg8Index[int(op2.Register)]
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.E028, parser.TokenLiteral(op2.Register)), node.Context)
			return object.ERROR
		}

		b := 0x40 | r1<<3 | r2
		return &object.CodeObject{Line: node.Context.Line, Code: []byte{byte(b)}}
	case *object.NumberObject:
		// LD r, n
		v, ok := e.intToByte(op2.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WEXPR_BYTE, op2.Value, op2.Value), node.Context)
		}
		r1 := Z80Reg8Index[int(op1.Register)]
		b := byte(0x06 | (r1 << 3))
		return &object.CodeObject{Line: node.Context.Line, Code: []byte{b, v}}
	// case *object.IndirectExpression:
	// 	e.logger.Error(fmt.Sprintf(errcode.E999, node), node.Context.LineNumber)
	// 	return object.ERROR
	case *object.RefNotFoundObject:
		e.Resolved = false
		return &object.CodeObject{Line: node.Context.Line, Code: []byte{0x7f}}

	default:

		e.logger.Error(errcode.EZ80_OP2, node.Context)
		return object.ERROR
	}
}

func (e *Evaluator) evalZ80LD_reg16(node *parser.Z80Instruction, op1 *object.RegisterObject, env object.Environment) object.Object {
	op2 := e.evalExpression(node.Op2, env, node.Context)

	switch op2 := op2.(type) {
	case *object.RegisterObject:
		// LD rr, rr'
		if op1.Register != parser.Z80_REG_SP {
			e.logger.Error(errcode.EZ80_OP1_SP, node.Context)
			return object.ERROR
		}
		if op2.RegisterType != parser.Z80_REG16 {
			e.logger.Error(errcode.EZ80_OP2, node.Context)
			return object.ERROR
		}
		switch op2.Register {
		case parser.Z80_REG_HL:
			return &object.CodeObject{Line: node.Context.Line, Code: []byte{0xf9}}
		case parser.Z80_REG_IX:
			return &object.CodeObject{Line: node.Context.Line, Code: []byte{0xdd, 0xf9}}
		case parser.Z80_REG_IY:
			return &object.CodeObject{Line: node.Context.Line, Code: []byte{0xfd, 0xf9}}
		default:
			e.logger.Error(errcode.EZ80_OP2_HL_IXY, node.Context)
			return object.ERROR
		}
	case *object.NumberObject:
		// LD rr, nn
		v, ok := e.intToWord(op2.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WEXPR_WORD, op2.Value, op2.Value), node.Context)
		}
		r1 := Z80Reg16Index[int(op1.Register)]
		b := byte(0x01 | (r1 << 4))
		return &object.CodeObject{Line: node.Context.Line, Code: []byte{b, byte(v & 0xff), byte((v >> 8) & 0xff)}}
	// case *object.IndirectExpression:
	// 	e.logger.Error(fmt.Sprintf(errcode.E999, node), node.Context.LineNumber)
	// 	return object.ERROR
	case *object.RefNotFoundObject:
		e.Resolved = false
		return &object.CodeObject{Line: node.Context.Line, Code: []byte{0x21, 0, 0}}

	default:
		fmt.Printf("%#v\n", op2)
		e.logger.Error(errcode.EZ80_OP2, node.Context)
		return object.ERROR
	}
}

func (e *Evaluator) intToByte(n int) (byte, bool) {
	switch {
	case 0 <= n && n <= 255:
		return byte(n), true
	case -128 <= n && n < 0:
		return byte(n & 0xff), true
	default:
		return byte(n & 0xff), false
	}
}

func (e *Evaluator) intToWord(n int) (int, bool) {
	switch {
	case 0 <= n && n <= 65535:
		return n, true
	case -32768 <= n && n < 0:
		return (n & 0xffff), true
	default:
		return (n & 0xffff), false
	}
}
