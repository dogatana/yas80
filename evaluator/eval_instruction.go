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
		info := Z80CodeTable0[int(node.OpCode)]
		obj := &object.CodeObject{Line: node.LineNumber(), Code: make([]byte, len(info.Bytes))}
		copy(obj.Code, info.Bytes)
		return obj
	case parser.Z80_INST1:
		if node.OpCode == parser.Z80_INST_RET {
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
	switch node.Op1.NodeType() {
	case parser.Z80_REG8:
		r1 := int(node.Op1.NodeSubType())
		if r1 > parser.Z80_REG_A {
			return object.NULL
		}
		switch node.Op2.NodeType() {
		case parser.Z80_REG8:
			// LD r,r'
			r2 := int(node.Op2.NodeSubType())
			if r2 > parser.Z80_REG_A {
				return object.NULL
			}
			b := 0xc0
			b |= ((r1 - parser.Z80_REG_B) << 3) | (r2 - parser.Z80_REG_B)
			return &object.CodeObject{Line: node.LineNumber(), Code: []byte{byte(b)}}
		default:
			op2 := e.Eval(node.Op2, env)
			if op2.Type() == object.NUMBER_OBJ {
				// LD r,n
				v := op2.(*object.NumberObject).Value
				bv, ok := e.intToByte(v)
				if !ok {
					e.logger.Warning(fmt.Sprintf(logger.W001, v, v), node.LineNumber())
					bv = byte(v & 0xff)
				}
				b := 0x06
				b |= (r1 - parser.Z80_REG_B) << 3
				return &object.CodeObject{Line: node.LineNumber(), Code: []byte{byte(b), bv}}
			} else if op2.Type() == object.NULL_OBJ {
				e.logger.Error(fmt.Sprintf("error expr: %s", node.Op2.String()), 0)
			}
		}
	case parser.Z80_REG16:
		r1 := int(node.Op1.NodeSubType())
		if r1 != parser.Z80_REG_HL {
			return object.NULL
		}
		op2 := e.Eval(node.Op2, env)
		num, ok := op2.(*object.NumberObject)
		if !ok {
			return object.NULL
		}
		return &object.CodeObject{Line: node.LineNumber(), Code: []byte{0x21, byte(num.Value & 0xff), byte((num.Value >> 8) & 0xff)}}

	default:
		return object.NULL
	}
	return object.NULL
}

func (e *Evaluator) intToByte(n int) (byte, bool) {
	switch {
	case 0 <= n && n <= 255:
		return byte(n), true
	case -128 <= n && n < 0:
		return byte(n & 0xff), true
	default:
		return 0, false
	}
}
