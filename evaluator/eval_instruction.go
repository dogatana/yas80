package evaluator

import (
	"fmt"
	"yas80/logger"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) evalZ80Instruction(node *parser.Z80Instruction, env *object.Environment) object.Object {
	switch node.NodeType() {
	case parser.Z80_INST0:
		info := Z80CodeTable0[int(node.Opcode)]
		obj := &object.CodeObject{Line: node.LineNumber(), Code: make([]byte, len(info.Bytes))}
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

func (e *Evaluator) generateRET(node *parser.Z80Instruction, _ *object.Environment) object.Object {
	if node.Op1 == nil {
		return &object.CodeObject{Line: node.LineNumber(), Code: []byte{0xc9}}
	}
	if node.Op1.NodeType() == parser.Z80_FLAG {
		flag := int(node.Op1.NodeSubType()) - parser.Z80_FLAG_NZ
		b := byte(0xc0 | flag<<3)
		return &object.CodeObject{Line: node.LineNumber(), Code: []byte{b}}
	}
	e.logger.Error(
		fmt.Sprintf(logger.E017, node.Op1.String()), node.LineNumber())
	return object.ERROR
}

func (e *Evaluator) evalZ80Instruction2(node *parser.Z80Instruction, env *object.Environment) object.Object {
	switch node.NodeSubType() {
	case parser.Z80_INST_LD:
		return e.evalZ80LD(node, env)
	default:
		e.logger.Error(fmt.Sprintf(logger.E999, node), node.LineNumber())
		return object.ERROR
	}
}

func (e *Evaluator) evalZ80LD(node *parser.Z80Instruction, env *object.Environment) object.Object {
	op1 := e.Eval(node.Op1, env)

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
		e.logger.Error(logger.E024, node.LineNumber())
		return object.ERROR
	// case *object.IndirectExpression:
	// 	e.logger.Error(fmt.Sprintf(logger.E999, node), node.LineNumber())
	// 	return object.ERROR
	default:
		if e.Pass1 {
			// pass1 の場合はダミーとして LD A, A を返す
			return &object.CodeObject{Line: node.LineNumber(), Code: []byte{0x7f}}
		}
		e.logger.Error(logger.E024, node.LineNumber())
		return object.ERROR
	}
}

func (e *Evaluator) evalZ80LD_reg8(node *parser.Z80Instruction, op1 *object.RegisterObject, env *object.Environment) object.Object {
	op2 := e.Eval(node.Op2, env)

	fmt.Printf("LD r8 op2 %s => %#v", node.Op2.String(), op2)
	switch op2 := op2.(type) {
	case *object.RegisterObject:
		// LD r, r'
		if op2.RegisterType != parser.Z80_REG8 {
			e.logger.Error(logger.E025, node.LineNumber())
			return object.ERROR
		}
		r1, ok := Z80Reg8Index[int(op1.Register)]
		if !ok {
			e.logger.Error(fmt.Sprintf(logger.E028, parser.TokenLiteral(op1.Register)), node.LineNumber())
			return object.ERROR
		}
		r2, ok := Z80Reg8Index[int(op2.Register)]
		if !ok {
			e.logger.Error(fmt.Sprintf(logger.E028, parser.TokenLiteral(op2.Register)), node.LineNumber())
			return object.ERROR
		}

		b := 0x40 | r1<<3 | r2
		return &object.CodeObject{Line: node.LineNumber(), Code: []byte{byte(b)}}
	case *object.NumberObject:
		// LD r, n
		v, ok := e.intToByte(op2.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(logger.W001, op2.Value, op2.Value), node.LineNumber())
		}
		r1 := Z80Reg8Index[int(op1.Register)]
		b := byte(0x06 | (r1 << 3))
		return &object.CodeObject{Line: node.LineNumber(), Code: []byte{b, v}}
	// case *object.IndirectExpression:
	// 	e.logger.Error(fmt.Sprintf(logger.E999, node), node.LineNumber())
	// 	return object.ERROR
	case *object.SymbolObject:
		fmt.Printf("op2 %#v", op2)
		op2v, ok := op2.Value.(*object.NumberObject)
		if ok {
			v, ok := e.intToByte(op2v.Value)
			if !ok {
				e.logger.Warning(fmt.Sprintf(logger.W001, op2.Value, op2.Value), node.LineNumber())
			}
			r1 := Z80Reg8Index[int(op1.Register)]
			b := byte(0x06 | (r1 << 3))
			return &object.CodeObject{Line: node.LineNumber(), Code: []byte{b, v}}
		}
		if e.Pass1 {
			// pass1 の場合はダミーとして LD A, A を返す
			return &object.CodeObject{Line: node.LineNumber(), Code: []byte{0x7f}}
		}
		e.logger.Error(logger.E025, node.LineNumber())
		return object.ERROR
	default:
		fmt.Printf("default op2 %#v", op2)
		if e.Pass1 {
			// pass1 の場合はダミーとして LD A, A を返す
			return &object.CodeObject{Line: node.LineNumber(), Code: []byte{0x7f}}
		}
		e.logger.Error(logger.E025, node.LineNumber())
		return object.ERROR
	}
}

func (e *Evaluator) evalZ80LD_reg16(node *parser.Z80Instruction, op1 *object.RegisterObject, env *object.Environment) object.Object {
	op2 := e.Eval(node.Op2, env)

	fmt.Printf("LD r16 op2 %s => %#v", node.Op2.String(), op2)

	switch op2 := op2.(type) {
	case *object.RegisterObject:
		// LD rr, rr'
		if op1.Register != parser.Z80_REG_SP {
			e.logger.Error(logger.E026, node.LineNumber())
			return object.ERROR
		}
		if op2.RegisterType != parser.Z80_REG16 {
			e.logger.Error(logger.E025, node.LineNumber())
			return object.ERROR
		}
		switch op2.Register {
		case parser.Z80_REG_HL:
			return &object.CodeObject{Line: node.LineNumber(), Code: []byte{0xf9}}
		case parser.Z80_REG_IX:
			return &object.CodeObject{Line: node.LineNumber(), Code: []byte{0xdd, 0xf9}}
		case parser.Z80_REG_IY:
			return &object.CodeObject{Line: node.LineNumber(), Code: []byte{0xfd, 0xf9}}
		default:
			e.logger.Error(logger.E027, node.LineNumber())
			return object.ERROR
		}
	case *object.NumberObject:
		// LD rr, nn
		v, ok := e.intToWord(op2.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(logger.W002, op2.Value, op2.Value), node.LineNumber())
		}
		r1 := Z80Reg16Index[int(op1.Register)]
		b := byte(0x01 | (r1 << 4))
		return &object.CodeObject{Line: node.LineNumber(), Code: []byte{b, byte(v & 0xff), byte((v >> 8) & 0xff)}}
	// case *object.IndirectExpression:
	// 	e.logger.Error(fmt.Sprintf(logger.E999, node), node.LineNumber())
	// 	return object.ERROR
	case *object.SymbolObject:
		fmt.Printf("op2 %#v", op2)
		op2v, ok := op2.Value.(*object.NumberObject)
		if ok {
			v, ok := e.intToWord(op2v.Value)
			if !ok {
				e.logger.Warning(fmt.Sprintf(logger.W002, op2v.Value, op2v.Value), node.LineNumber())
			}
			r1 := Z80Reg16Index[int(op1.Register)]
			b := byte(0x01 | (r1 << 4))
			return &object.CodeObject{Line: node.LineNumber(), Code: []byte{b, byte(v & 0xff), byte((v >> 8) & 0xff)}}
		}
		if e.Pass1 {
			// pass1 の場合はダミーとして LD A, A を返す
			return &object.CodeObject{Line: node.LineNumber(), Code: []byte{0x21, 0, 0}}
		}
		e.logger.Error(logger.E025, node.LineNumber())
		return object.ERROR
	default:
		if e.Pass1 {
			// pass1 の場合はダミーとして LD HL, 0 を返す
			return &object.CodeObject{Line: node.LineNumber(), Code: []byte{0x21, 0, 0}}
		}
		e.logger.Error(logger.E025, node.LineNumber())
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
