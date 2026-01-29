package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

// 16 ビット演算命令
// ADD A, ADC A, SBC A, は evalZ80_ADD8 に処理を移譲する
func (e *Evaluator) evalZ80_ADD16(stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object {
	var code *object.CodeObject
	switch stmt.Opcode {
	case parser.Z80_INST_ADD:
		code = &object.CodeObject{Code: []byte{0x09}, CZ80: 11, Context: stmt.Context}
	case parser.Z80_INST_ADC:
		code = &object.CodeObject{Code: []byte{0xed, 0x4a}, CZ80: 15, Context: stmt.Context}
	case parser.Z80_INST_SBC:
		code = &object.CodeObject{Code: []byte{0xed, 0x42}, CZ80: 15, Context: stmt.Context}
	}

	if op1 == nil {
		e.logger.Error(errcode.EZ80_OP1, stmt.Context)
		return code
	}

	reg1 := 0
	switch op1 := op1.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP1_NULL, stmt.Context)
		return code

	case *object.RegisterObject:
		if op1.RegisterType == parser.Z80_REG8 {
			// 第1オペランドが8ビットレジスタの場合は evalZ80_ADD8 を呼び出す
			return e.evalZ80_ADD8(stmt, op1, op2, env)
		}
		if stmt.Opcode == parser.Z80_INST_ADD && op1.Register != parser.Z80_REG_HL &&
			// ADD 第1オペランドは HL/IX/IY
			op1.Register != parser.Z80_REG_IX && op1.Register != parser.Z80_REG_IY {
			e.logger.Error(errcode.EZ80_OP1_REG_HL_IXY, stmt.Context)
		} else if stmt.Opcode != parser.Z80_INST_ADD && op1.Register != parser.Z80_REG_HL {
			// ADC, SBC 第1オペランドは HL
			e.logger.Error(errcode.EZ80_OP1_REG_HL, stmt.Context)
			return code
		}
		reg1 = op1.Register

	default:
		e.logger.Error(errcode.EZ80_OP1, stmt.Context)
	}

	switch op2 := op2.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return code

	case *object.RegisterObject:
		index, ok := Z80Reg16IndexSPIXY[op2.Register]
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op2.Register)), stmt.Context)
			return code
		}
		// アドレス情報設定
		code.Code[len(code.Code)-1] |= index << 4

		switch reg1 {
		case parser.Z80_REG_HL:
			return code
		case parser.Z80_REG_IX: // ADD のみ
			code.Code = []byte{0xdd, code.Code[0]}
			code.CZ80 = 15
			return code
		case parser.Z80_REG_IY: // ADD のみ
			code.Code = []byte{0xfd, code.Code[0]}
			code.CZ80 = 15
			return code
		}
	}
	e.logger.Error(errcode.EZ80_OP2, stmt.Context)
	return code
}
