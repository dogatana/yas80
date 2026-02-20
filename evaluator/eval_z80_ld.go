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

	if op1 == nil || op2 == nil {
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return object.ERROR
	}

	// op1 の型によって処理を分類
	switch op1 := op1.(type) {
	case *object.RefNotFoundObject:
		return op1
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP1_NULL, stmt.Context)
		return object.ERROR

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
	return object.ERROR
}

// 8 ビットレジスタへの LD
func (e *Evaluator) evalZ80LD_REG8(stmt *parser.Z80Instruction, op1 *object.RegisterObject, argOp2 object.Object, _ TEnv) object.Object {
	// r1 が A-L の場合
	r1, ok1 := Z80Reg8Index[op1.Register]

EVAL_AGAIN:
	switch op2 := argOp2.(type) {
	case *object.RefNotFoundObject:
		// 未確定の場合として LD A,0 を返す
		e.Resolved = false
		return op2
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return object.ERROR

	case *object.RegisterObject:
		code := &object.CodeObject{Code: []byte{0x40}, TStates: [2]byte{4, 1}, Context: stmt.Context} // LD A, A
		// LD r, r
		r2, ok2 := Z80Reg8Index[op2.Register]
		if ok1 && ok2 {
			code.Code[0] |= r1<<3 | r2
			return code
		}
		code.TStates = [2]byte{9, 2}
		switch {
		case op1.Register == parser.Z80_REG_A && op2.Register == parser.Z80_REG_I:
			code.Code = []byte{0xed, 0x57} // LD A, I
			return code
		case op1.Register == parser.Z80_REG_A && op2.Register == parser.Z80_REG_R:
			code.Code = []byte{0xed, 0x5f} // LD A, R
			return code
		case op1.Register == parser.Z80_REG_I && op2.Register == parser.Z80_REG_A:
			code.Code = []byte{0xed, 0x47} // LD I, A
			return code
		case op1.Register == parser.Z80_REG_R && op2.Register == parser.Z80_REG_A:
			code.Code = []byte{0xed, 0x4f} // LD I, R
			return code
		}

		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op2.Register)), stmt.Context)
		return object.ERROR

	case *object.StringObject:
		argOp2 = e.stringObjToOp2(op2, stmt.Context)
		goto EVAL_AGAIN

	case *object.NumberObject:
		code := &object.CodeObject{Code: []byte{0x06, 0}, TStates: [2]byte{7, 2}, Context: stmt.Context}
		if !ok1 {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op1.Register)), stmt.Context)
			return object.ERROR
		}
		// LD r, n
		v, ok := e.intToByte(op2.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_BYTE, op2.Value, op2.Value), stmt.Context)
		}
		code.Code[0] |= (r1 << 3)
		code.Code[1] = v
		return code

	case *object.RegIndirectObject:
		code := &object.CodeObject{Code: []byte{0x46}, TStates: [2]byte{7, 2}, Context: stmt.Context}
		code.Code[0] |= r1 << 3
		switch op2.Register {
		case parser.Z80_REG_HL: // LD r, (HL)
			return code
		case parser.Z80_REG_IX: // LD r, (IX + d)
			code.Code = []byte{0xdd, code.Code[0], byte(op2.Displacement)}
			code.TStates = [2]byte{19, 5}
			return code
		case parser.Z80_REG_IY: // LD r, (IY + d)
			code.Code = []byte{0xfd, code.Code[0], byte(op2.Displacement)}
			code.TStates = [2]byte{19, 5}
			return code
		case parser.Z80_REG_BC, parser.Z80_REG_DE:
			if op1.Register != parser.Z80_REG_A {
				e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op1.Register)), stmt.Context)
				return object.ERROR
			}
			if op2.Register == parser.Z80_REG_BC {
				code.Code[0] = 0x0a // LD A, (BC)
			} else {
				code.Code[0] = 0x1a // LD A, (DE)
			}
			return code

		default:
			e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(op2.Register)), stmt.Context)
			return object.ERROR
		}

	case *object.AddrIndirectObject:
		// LD r, (nn)
		code := &object.CodeObject{Code: []byte{0x3a, 0, 0}, TStates: [2]byte{13, 4}, Context: stmt.Context}
		addr, ok := e.intToWord(op2.Address)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, op2.Address, op2.Address), stmt.Context)
		}
		code.Code[1] = byte(addr & 0xff)
		code.Code[2] = byte(addr >> 8)
		return code

	default:
		e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		return object.ERROR
	}
}

// 16 ビットレジスタへの LD
func (e *Evaluator) evalZ80LD_REG16(stmt *parser.Z80Instruction, op1 *object.RegisterObject, argOp2 object.Object, env TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0x01, 0x00, 0x00}, TStates: [2]byte{10, 0}, Context: stmt.Context} // LD BC, 0

EVAL_AGAIN:
	switch op2 := argOp2.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return op2
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return object.ERROR

	case *object.RegisterObject:
		// LD SP, rr
		if op1.Register != parser.Z80_REG_SP {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op1.Register)), stmt.Context)
			return object.ERROR
		}
		switch op2.Register {
		case parser.Z80_REG_HL: // LD SP, HL
			code = &object.CodeObject{Code: []byte{0xf9}, TStates: [2]byte{6, 1}, Context: stmt.Context}
			return code
		case parser.Z80_REG_IX: // LD SP, IX
			return &object.CodeObject{Code: []byte{0xdd, 0xf9}, TStates: [2]byte{10, 2}, Context: stmt.Context}
		case parser.Z80_REG_IY: // LD SP, IY
			return &object.CodeObject{Code: []byte{0xfd, 0xf9}, TStates: [2]byte{10, 2}, Context: stmt.Context}
		default:
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op2.Register)), stmt.Context)
			return object.ERROR
		}

	case *object.StringObject:
		argOp2 = e.stringObjToOp2(op2, stmt.Context)
		goto EVAL_AGAIN

	case *object.NumberObject:
		// LD rr, nn

		r1, ok := Z80Reg16IndexSPIXY[int(op1.Register)]
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op1.Register)), stmt.Context)
			return object.ERROR
		}

		v, ok := e.intToWord(op2.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, op2.Value, op2.Value), stmt.Context)
		}

		code = &object.CodeObject{Code: []byte{0x01, 0x00, 0x00}, TStates: [2]byte{10, 3}, Context: stmt.Context} // LD BC, 0
		code.Code[0] |= r1 << 4
		code.Code[1] = byte(v & 0xff)
		code.Code[2] = byte((v >> 8) & 0xff)

		switch op1.Register {
		case parser.Z80_REG_IX:
			ncode := []byte{0xdd}
			ncode = append(ncode, code.Code...)
			code.Code = ncode
			code.TStates = [2]byte{14, 4}
		case parser.Z80_REG_IY:
			ncode := []byte{0xfd}
			ncode = append(ncode, code.Code...)
			code.Code = ncode
			code.TStates = [2]byte{14, 4}
		default:
			// LD rr, nn
		}
		return code

	case *object.AddrIndirectObject:
		// LD rr, (nn)
		r1, ok := Z80Reg16IndexSPIXY[op1.Register]
		if !ok {
			e.logger.Error(errcode.EZ80_OP1, stmt.Context)
			return object.ERROR
		}
		v, ok := e.intToWord(op2.Address)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, op2.Address, op2.Address), stmt.Context)
		}

		switch op1.Register {
		case parser.Z80_REG_HL: // LD HL, (nn)
			return &object.CodeObject{Code: []byte{0x2a, byte(v & 0xff), byte((v >> 8) & 0xff)}, TStates: [2]byte{16, 5}, Context: stmt.Context}
		case parser.Z80_REG_IX: // LD IX, (nn)
			return &object.CodeObject{Code: []byte{0xdd, 0x2a, byte(v & 0xff), byte((v >> 8) & 0xff)}, TStates: [2]byte{20, 6}, Context: stmt.Context}
		case parser.Z80_REG_IY: // LD IY, (nn)
			return &object.CodeObject{Code: []byte{0xfd, 0x2a, byte(v & 0xff), byte((v >> 8) & 0xff)}, TStates: [2]byte{20, 6}, Context: stmt.Context}
		default: // LD rr, (nn)
			return &object.CodeObject{Code: []byte{0xed, 0x4b | (r1 << 4), byte(v & 0xff), byte((v >> 8) & 0xff)}, TStates: [2]byte{20, 6}, Context: stmt.Context}
		}

	default:
		e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		return object.ERROR
	}
}

// レジスタ間接
func (e *Evaluator) evalZ80LD_RegIndirect(stmt *parser.Z80Instruction, op1 *object.RegIndirectObject, op2 object.Object, env TEnv) object.Object {

	switch op2.(type) {
	case *object.RefNotFoundObject:
		return op2
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return object.ERROR
	}

	r2, ok := op2.(*object.RegisterObject)
	if ok {
		index, ok := Z80Reg8Index[r2.Register]
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(r2.Register)), stmt.Context)
			return object.ERROR
		}
		code := &object.CodeObject{Code: []byte{0x70}, TStates: [2]byte{7, 2}, Context: stmt.Context} // LD (HL),B
		switch {
		case op1.Register == parser.Z80_REG_HL: // LD (HL), r
			code.Code[0] |= index
			return code
		case op1.Register == parser.Z80_REG_IX: // LD (IX + d)
			code.Code = []byte{0xdd, code.Code[0] | index, byte(op1.Displacement)}
			code.TStates = [2]byte{19, 5}
			return code
		case op1.Register == parser.Z80_REG_IY: // LD (IY + d)
			code.Code = []byte{0xfd, code.Code[0] | index, byte(op1.Displacement)}
			code.TStates = [2]byte{19, 5}
			return code
		case op1.Register == parser.Z80_REG_BC && r2.Register == parser.Z80_REG_A: // LD (BC), A
			code.Code[0] = 0x02
			return code
		case op1.Register == parser.Z80_REG_DE && r2.Register == parser.Z80_REG_A: // LD (DE), A
			code.Code[0] = 0x12
			return code
		}
		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(op1.Register)), stmt.Context)
		return object.ERROR
	}

	num, ok := op2.(*object.NumberObject)
	if ok {
		b, ok := e.intToByte(num.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_BYTE, num.Value, num.Value), stmt.Context)
		}
		switch op1.Register {
		case parser.Z80_REG_HL: // LD (HL), n
			return &object.CodeObject{Code: []byte{0x36, b}, TStates: [2]byte{10, 3}, Context: stmt.Context}
		case parser.Z80_REG_IX: // LD (IX + d), n
			return &object.CodeObject{Code: []byte{0xdd, 0x36, byte(op1.Displacement), b}, TStates: [2]byte{19, 5}, Context: stmt.Context}
		case parser.Z80_REG_IY: // LD (IY + d), n
			return &object.CodeObject{Code: []byte{0xfd, 0x36, byte(op1.Displacement), b}, TStates: [2]byte{19, 5}, Context: stmt.Context}
		}
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return object.ERROR
	}

	e.logger.Error(errcode.EZ80_OP2, stmt.Context)
	return object.ERROR
}

// アドレ間接（メモリ）
func (e *Evaluator) evalZ80LD_AddrIndirect(stmt *parser.Z80Instruction, op1 *object.AddrIndirectObject, op2 object.Object, env TEnv) object.Object {

	switch op2.(type) {
	case *object.RefNotFoundObject:
		return op2
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return object.ERROR
	}

	r2, ok := op2.(*object.RegisterObject)
	if !ok {
		e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		return object.ERROR
	}

	addr, ok := e.intToWord(op1.Address)
	if !ok {
		e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, op1.Address, op1.Address), stmt.Context)
	}

	switch r2.Register {
	case parser.Z80_REG_A: // LD (nn), A
		return &object.CodeObject{Code: []byte{0x32, byte(addr & 0xff), byte((addr >> 8) & 0xff)}, TStates: [2]byte{13, 4}, Context: stmt.Context}

	case parser.Z80_REG_HL: // LD (nn), HL
		return &object.CodeObject{Code: []byte{0x22, byte(addr & 0xff), byte((addr >> 8) & 0xff)}, TStates: [2]byte{16, 5}, Context: stmt.Context}

	case parser.Z80_REG_IX: // LD (nn), IX
		return &object.CodeObject{Code: []byte{0xdd, 0x22, byte(addr & 0xff), byte((addr >> 8) & 0xff)}, TStates: [2]byte{20, 6}, Context: stmt.Context}

	case parser.Z80_REG_IY: // LD (nn), IY
		return &object.CodeObject{Code: []byte{0xfd, 0x22, byte(addr & 0xff), byte((addr >> 8) & 0xff)}, TStates: [2]byte{20, 6}, Context: stmt.Context}

	default: // LD (nn), rr
		if index, ok := Z80Reg16IndexSP[r2.Register]; ok {
			return &object.CodeObject{Code: []byte{0xed, 0x43 | (index << 4), byte(addr & 0xff), byte((addr >> 8) & 0xff)}, TStates: [2]byte{20, 6}, Context: stmt.Context}
		}
		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(r2.Register)), stmt.Context)
		return object.ERROR
	}
}
