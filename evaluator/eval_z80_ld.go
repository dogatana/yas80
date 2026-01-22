package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

// LD - デスティネーションによって処理を分割
// r evalZ80_LD_REG8
// rr evalZ80_LD_REG16
// () evalZ80_LD_Indirect
func (e *Evaluator) evalZ80LD(stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0x3e, 0}, CZ80: 4, Context: stmt.Context} // LD A, 0

	if op1 == nil || op2 == nil {
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return code
	}

	if op2, ok := op2.(*object.RegisterObject); ok && op2.RegisterType == parser.Z80_REG16 {
		code.Code = []byte{0x21, 0x00, 0x00} // LD HL, 0
		code.CZ80 = 10
	}

	// op1 の型によって処理を分類
	switch op1 := op1.(type) {
	case *object.RefNotFoundObject:
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP1_NULL, stmt.Context)
		return code

	case *object.RegisterObject:
		if op1.RegisterType == parser.Z80_REG8 {
			// LD r, expr
			return e.evalZ80LD_REG8(stmt, op1, op2, env)
		} else {
			// LD rr, expr
			return e.evalZ80LD_REG16(stmt, op1, op2, env)
		}
	case *object.RegIndirectObject:
		// LD (expr), expr
		return e.evalZ80LD_RegIndirect(stmt, op1, op2, env)

	case *object.AddrIndirectObject:
		// LD (nn), A  LD (nn), rr
		return e.evalZ80LD_AddrIndirect(stmt, op1, op2, env)
	}
	e.logger.Error(errcode.EZ80_OP1, stmt.Context)
	return code
}

// 8 ビットレジスタへの LD
func (e *Evaluator) evalZ80LD_REG8(stmt *parser.Z80Instruction, op1 *object.RegisterObject, op2 object.Object, _ TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0x7f}, CZ80: 4, Context: stmt.Context} // LD A, A

	// r1 が A-L の場合
	r1, ok1 := Z80Reg8Index[op1.Register]

	switch op2 := op2.(type) {
	case *object.RefNotFoundObject:
		// 未確定の場合として LD A,0 を返す
		e.Resolved = false
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return code

	case *object.RegisterObject:
		// LD r, r
		r2, ok2 := Z80Reg8Index[op2.Register]
		if ok1 && ok2 {
			code.Code[0] = 0x40 | r1<<3 | r2
			return code
		}
		code.CZ80 = 9
		switch {
		case op1.Register == parser.Z80_REG_A && op2.Register == parser.Z80_REG_I:
			code.Code = []byte{0xed, 0x57}
			return code
		case op1.Register == parser.Z80_REG_A && op2.Register == parser.Z80_REG_R:
			code.Code = []byte{0xed, 0x5f}
			return code
		case op1.Register == parser.Z80_REG_I && op2.Register == parser.Z80_REG_A:
			code.Code = []byte{0xed, 0x47}
			return code
		case op1.Register == parser.Z80_REG_R && op2.Register == parser.Z80_REG_A:
			code.Code = []byte{0xed, 0x4f}
			return code
		}

		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op2.Register)), stmt.Context)
		return code

	case *object.NumberObject:
		code.Code = []byte{0x3e, 0x00} // LD A, 0
		code.CZ80 = 7
		if !ok1 {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op1.Register)), stmt.Context)
			return code
		}
		// LD r, n
		v, ok := e.intToByte(op2.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_BYTE, op2.Value, op2.Value), stmt.Context)
		}
		code.Code = []byte{0x06 | (r1 << 3), v}
		return code

	case *object.RegIndirectObject:
		// LD r, (rr)
		code.Code[0] = 0x46 | r1<<3 // LD r, (HL)
		code.CZ80 = 7
		switch op2.Register {
		case parser.Z80_REG_HL:
			return code
		case parser.Z80_REG_IX:
			code.Code = []byte{0xdd, code.Code[0], byte(op2.Displacement)}
			code.CZ80 = 19
			return code
		case parser.Z80_REG_IY:
			code.Code[0] = 0x46 | r1<<3
			code.Code = []byte{0xfd, code.Code[0], byte(op2.Displacement)}
			code.CZ80 = 19
			return code
		case parser.Z80_REG_BC, parser.Z80_REG_DE:
			if op1.Register != parser.Z80_REG_A {
				e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op1.Register)), stmt.Context)
				return code
			}
			if op2.Register == parser.Z80_REG_BC {
				code.Code[0] = 0x0a
			} else {
				code.Code[0] = 0x1a
			}
			return code

		default:
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(op2.Register)), stmt.Context)
			return code
		}

	case *object.AddrIndirectObject:
		// LD r, (nn)
		addr, ok := e.intToWord(op2.Address)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, op2.Address, op2.Address), stmt.Context)
		}
		code.Code = []byte{0x3a, byte(addr & 0xff), byte(addr >> 8)}
		code.CZ80 = 13
		return code

	default:
		e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		return object.ERROR
	}
}

// 16 ビットレジスタへの LD
func (e *Evaluator) evalZ80LD_REG16(stmt *parser.Z80Instruction, op1 *object.RegisterObject, op2 object.Object, env TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0x01, 0x00, 0x00}, CZ80: 10, Context: stmt.Context} // LD BC, 0

	switch op2 := op2.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return code

	case *object.RegisterObject:
		// LD SP, rr
		code = &object.CodeObject{Code: []byte{0xf9}, CZ80: 6, Context: stmt.Context} // LD SP, HL

		if op1.Register != parser.Z80_REG_SP {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op1.Register)), stmt.Context)
			return code
		}
		switch op2.Register {
		case parser.Z80_REG_HL:
			return code
		case parser.Z80_REG_IX:
			return &object.CodeObject{Code: []byte{0xdd, 0xf9}, CZ80: 10, Context: stmt.Context} // LD SP, HL
		case parser.Z80_REG_IY:
			return &object.CodeObject{Code: []byte{0xfd, 0xf9}, CZ80: 10, Context: stmt.Context} // LD SP, HL
		default:
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op2.Register)), stmt.Context)
			return code
		}

	case *object.NumberObject:
		// LD rr, nn
		code = &object.CodeObject{Code: []byte{0x01, 0x00, 0x00}, CZ80: 10, Context: stmt.Context} // LD BC, 0

		r1, ok := Z80Reg16IndexSPIXY[int(op1.Register)]
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op1.Register)), stmt.Context)
			return code
		}

		v, ok := e.intToWord(op2.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, op2.Value, op2.Value), stmt.Context)
		}

		code.Code[0] |= r1 << 4
		code.Code[1] = byte(v & 0xff)
		code.Code[2] = byte((v >> 8) & 0xff)
		switch op1.Register {
		case parser.Z80_REG_IX:
			ncode := []byte{0xdd}
			ncode = append(ncode, code.Code...)
			code.Code = ncode
			code.CZ80 = 14
		case parser.Z80_REG_IY:
			ncode := []byte{0xfd}
			ncode = append(ncode, code.Code...)
			code.Code = ncode
			code.CZ80 = 14
		}
		return code

	case *object.AddrIndirectObject:
		code := &object.CodeObject{Code: []byte{0x2a, 0x00, 0x00}, CZ80: 16, Context: stmt.Context}
		r1, ok := Z80Reg16IndexSPIXY[op1.Register]
		if !ok {
			e.logger.Error(errcode.EZ80_OP1, stmt.Context)
			return code
		}
		v, ok := e.intToWord(op2.Address)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, op2.Address, op2.Address), stmt.Context)
		}

		code.Code[1] = byte(v & 0xff)
		code.Code[2] = byte((v >> 8) & 0xff)
		switch op1.Register {
		case parser.Z80_REG_HL:
			return code
		case parser.Z80_REG_IX:
			return &object.CodeObject{Code: []byte{0xdd, 0x2a, byte(v & 0xff), byte((v >> 8) & 0xff)}, CZ80: 20, Context: stmt.Context}
		case parser.Z80_REG_IY:
			return &object.CodeObject{Code: []byte{0xfd, 0x2a, byte(v & 0xff), byte((v >> 8) & 0xff)}, CZ80: 20, Context: stmt.Context}
		default:
			return &object.CodeObject{Code: []byte{0xed, 0x4b | (r1 << 4), byte(v & 0xff), byte((v >> 8) & 0xff)}, CZ80: 20, Context: stmt.Context}
		}

	default:
		e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		return object.ERROR
	}
}

// レジスタ間接
func (e *Evaluator) evalZ80LD_RegIndirect(stmt *parser.Z80Instruction, op1 *object.RegIndirectObject, op2 object.Object, env TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0x70}, CZ80: 7, Context: stmt.Context} // LD (HL),B

	switch op2.(type) {
	case *object.RefNotFoundObject:
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return code
	}

	r2, ok := op2.(*object.RegisterObject)
	if ok {
		index, ok := Z80Reg8Index[r2.Register]
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(r2.Register)), stmt.Context)
			return code
		}
		switch {
		case op1.Register == parser.Z80_REG_HL:
			// LD (HL), r
			code.Code[0] |= index
			return code
		case op1.Register == parser.Z80_REG_IX:
			// LD (IX + d)
			code.Code = []byte{0xdd, code.Code[0] | index, byte(op1.Displacement)}
			code.CZ80 = 19
			return code
		case op1.Register == parser.Z80_REG_IY:
			// LD (IY + d)
			code.Code = []byte{0xfd, code.Code[0] | index, byte(op1.Displacement)}
			code.CZ80 = 19
			return code
		case op1.Register == parser.Z80_REG_BC && r2.Register == parser.Z80_REG_A:
			// LD (BC), A
			code.Code[0] = 0x02
			return code
		case op1.Register == parser.Z80_REG_DE && r2.Register == parser.Z80_REG_A:
			// LD (DE), A
			code.Code[0] = 0x12
			return code
		}
		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op1.Register)), stmt.Context)
		return code
	}

	num, ok := op2.(*object.NumberObject)
	if ok {
		b, ok := e.intToByte(num.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_BYTE, num.Value, num.Value), stmt.Context)
		}
		switch op1.Register {
		case parser.Z80_REG_HL:
			return &object.CodeObject{Code: []byte{0x36, b}, CZ80: 10, Context: stmt.Context}
		case parser.Z80_REG_IX:
			return &object.CodeObject{Code: []byte{0xdd, 0x36, b}, CZ80: 19, Context: stmt.Context}
		case parser.Z80_REG_IY:
			return &object.CodeObject{Code: []byte{0xfd, 0x36, b}, CZ80: 19, Context: stmt.Context}
		}
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return code
	}

	e.logger.Error(errcode.EZ80_OP2, stmt.Context)
	return code
}

// アドレ間接（メモリ）
func (e *Evaluator) evalZ80LD_AddrIndirect(stmt *parser.Z80Instruction, op1 *object.AddrIndirectObject, op2 object.Object, env TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0}, CZ80: 4, Context: stmt.Context}

	switch op2.(type) {
	case *object.RefNotFoundObject:
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return code
	}

	r2, ok := op2.(*object.RegisterObject)
	if !ok {
		e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		return code
	}

	addr, ok := e.intToWord(op1.Address)
	if !ok {
		e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, op1.Address, op1.Address), stmt.Context)
	}

	switch r2.Register {
	case parser.Z80_REG_A:
		// LD (nn), A
		return &object.CodeObject{Code: []byte{0x32, byte(addr & 0xff), byte((addr >> 8) & 0xff)}, CZ80: 13, Context: stmt.Context}

	case parser.Z80_REG_HL:
		// LD (nn), HL
		return &object.CodeObject{Code: []byte{0x22, byte(addr & 0xff), byte((addr >> 8) & 0xff)}, CZ80: 16, Context: stmt.Context}
	case parser.Z80_REG_IX:
		// LD (nn), IX
		return &object.CodeObject{Code: []byte{0xdd, 0x22, byte(addr & 0xff), byte((addr >> 8) & 0xff)}, CZ80: 20, Context: stmt.Context}
	case parser.Z80_REG_IY:
		// LD (nn), IY
		return &object.CodeObject{Code: []byte{0xfd, 0x22, byte(addr & 0xff), byte((addr >> 8) & 0xff)}, CZ80: 20, Context: stmt.Context}

	default:
		if index, ok := Z80Reg16IndexSP[r2.Register]; ok {
			return &object.CodeObject{Code: []byte{0xed, 0x43 | (index << 4), byte(addr & 0xff), byte((addr >> 8) & 0xff)}, CZ80: 20, Context: stmt.Context}
		}
		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(r2.Register)), stmt.Context)
		return code
	}
}
