package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

// 8 ビット演算命令
// 本来 1 オペランド命令だが、演算対象が A レジスタなので、A を含む 2 オペランドの形式も許容する
func (e *Evaluator) evalZ80_ADD8(stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object {
	var opcodes = map[int]byte{
		parser.Z80_INST_SUB: 0x90,
		parser.Z80_INST_AND: 0xa0,
		parser.Z80_INST_OR:  0xb0,
		parser.Z80_INST_XOR: 0xa8,
		parser.Z80_INST_CP:  0xb8,
	}
	code := &object.CodeObject{Code: []byte{opcodes[stmt.Opcode]}, CZ80: 4, Context: stmt.Context} // PUSH BC

	// op1 は省略（nil) か A のみ有効
	if op1 != nil {
		switch op1 := op1.(type) {
		case *object.RefNotFoundObject:
			e.Resolved = false
			return code
		case *object.NullObject:
			e.logger.Error(errcode.EZ80_OP1_NULL, stmt.Context)
			return code
		case *object.RegisterObject:
			if op1.Register != parser.Z80_REG_A {
				e.logger.Error(errcode.EZ80_OP1_REG_A, stmt.Context)
				return code
			}
		default:
			e.logger.Error(errcode.EZ80_OP1, stmt.Context)
		}
	}

	switch op2 := op2.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return code
	case *object.NullObject:
		if op1 == nil {
			e.logger.Error(errcode.EZ80_OP_NULL, stmt.Context)
		} else {
			e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		}
		return code

	// OP r
	case *object.RegisterObject:
		index, ok := Z80Reg8Index[op2.Register]
		if ok {
			code.Code[0] |= index
			return code
		}
		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op2.Register)), stmt.Context)
		return code

	// OP n
	case *object.NumberObject:
		code.CZ80 = 7
		code.Code[0] |= 0x46 // n を利用
		n, ok := e.intToByte(op2.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_BYTE, op2.Value, op2.Value), stmt.Context)
		}
		code.Code = []byte{code.Code[0], n}
		return code

	// OP (HL),(IX+d),(IY+d)
	case *object.RegIndirectObject:
		code.Code[0] |= 0x06
		code.CZ80 = 7
		switch op2.Register {
		case parser.Z80_REG_HL:
			return code
		case parser.Z80_REG_IX:
			code.Code = []byte{0xdd, code.Code[0], byte(op2.Displacement)}
			code.CZ80 = 19
			return code
		case parser.Z80_REG_IY:
			code.Code = []byte{0xfd, code.Code[0], byte(op2.Displacement)}
			code.CZ80 = 19
			return code
		default:
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(op2.Register)), stmt.Context)
			return code

		}

	default:
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return code
	}
}
