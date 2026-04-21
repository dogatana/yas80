package evaluator

import (
	"fmt"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
)

var _add8_opcodes = map[int]byte{
	parser.Z80_INST_ADD: 0x80,
	parser.Z80_INST_ADC: 0x88,
	parser.Z80_INST_SUB: 0x90,
	parser.Z80_INST_SBC: 0x98,
	parser.Z80_INST_AND: 0xa0,
	parser.Z80_INST_OR:  0xb0,
	parser.Z80_INST_XOR: 0xa8,
	parser.Z80_INST_CP:  0xb8,
}

// 8 ビット演算命令
// ADD, ADC, SBC 以外は 1 オペランド命令だが、A の指定も可能とする
// ADD, ADC, SBC は evalZ80_ADD16 から呼び出される
func (e *Evaluator) evalZ80_ADD8(stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object {

	// op1 は省略（nil) か A のみ有効
	if op1 != nil {
		switch op1 := op1.(type) {
		case *object.RefNotFoundObject:
			return op1
		case *object.NullObject:
			e.logger.Error(errcode.EZ80_OP1_NULL, stmt.Context)
			return object.ERROR
		case *object.RegisterObject:
			if op1.Register != parser.Z80_REG_A {
				e.logger.Error(errcode.EZ80_OP1_REG_A, stmt.Context)
				return object.ERROR
			}
		default:
			e.logger.Error(errcode.EZ80_OP1, stmt.Context)
			return object.ERROR
		}
	}

	code := &object.CodeObject{Code: []byte{_add8_opcodes[stmt.Opcode]}, TStates: [2]byte{4, 1}, Context: stmt.Context}

EVAL_AGAIN:
	switch eop2 := op2.(type) {
	case *object.ErrorObject:
		return eop2
	case *object.RefNotFoundObject:
		return eop2
	case *object.NullObject:
		if op1 == nil {
			e.logger.Error(errcode.EZ80_OP_NULL, stmt.Context)
		} else {
			e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		}
		return object.ERROR

	case *object.RegisterObject:
		// OP r
		if index, ok := Z80Reg8Index[eop2.Register]; ok {
			code.Code[0] |= index
			return code // OP r
		}

		code.TStates = [2]byte{8, 2}
		if index, ok := Z80Reg8IndexIX[eop2.Register]; ok {
			code.Code = []byte{0xdd, code.Code[0] | index}
			return code // OP p - IXH/IXL
		}
		if index, ok := Z80Reg8IndexIY[eop2.Register]; ok {
			code.Code = []byte{0xfd, code.Code[0] | index}
			return code // OP q - IYH/IYL
		}
		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(eop2.Register)), stmt.Context)
		return object.ERROR

	// OP n
	case *object.NumberObject:
		code.TStates = [2]byte{7, 2}
		code.Code[0] |= 0x46 // n を利用
		n, ok := e.intToByte(eop2.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_BYTE, eop2.Value, eop2.Value), stmt.Context)
		}
		code.Code = []byte{code.Code[0], n}
		return code

	case *object.StringObject:
		op2 = e.evalOneCharStringAsNumber(eop2.Value, stmt.Context)
		goto EVAL_AGAIN

	case *object.ArrayObject:
		op2 = e.evalArrayToInt(eop2.Values, stmt.Context)
		goto EVAL_AGAIN

	// OP (HL),(IX+d),(IY+d)
	case *object.RegIndirectObject:
		code.Code[0] |= 0x06
		switch eop2.Register {
		case parser.Z80_REG_HL:
			code.TStates = [2]byte{7, 2}
			return code
		case parser.Z80_REG_IX:
			code.Code = []byte{0xdd, code.Code[0], byte(eop2.Displacement)}
			code.TStates = [2]byte{19, 5}
			return code
		case parser.Z80_REG_IY:
			code.Code = []byte{0xfd, code.Code[0], byte(eop2.Displacement)}
			code.TStates = [2]byte{19, 5}
			return code
		default:
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(eop2.Register)), stmt.Context)
			return object.ERROR
		}

	default:
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return object.ERROR
	}
}
