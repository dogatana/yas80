package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) evalZ80_RET(stmt *parser.Z80Instruction, op1, _ object.Object, _ TEnv) object.Object {

	// RET
	if op1 == nil {
		return &object.CodeObject{Code: []byte{0xc9}, TStates: [2]byte{10, 3}, Context: stmt.Context}
	}

	// RET cc
	index := -1
	switch op1 := op1.(type) {
	case *object.FlagObject:
		index = op1.Flag
	case *object.RegisterObject: // C レジスタを CY に読み替える
		index = op1.Register
	case *object.RefNotFoundObject:
		return op1
	default:
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return object.ERROR
	}

	flag, ok := Z80FlagIndex[index]
	if !ok {
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return object.ERROR
	}
	b := byte(0xc0 | flag<<3)
	return &object.CodeObject{Code: []byte{b}, TStates: [2]byte{11, 3}, Context: stmt.Context}
}

func (e *Evaluator) evalZ80_CALL(stmt *parser.Z80Instruction, op1, op2 object.Object, _ TEnv) object.Object {

	value := 0
	switch op2 := op2.(type) {
	case *object.NumberObject:
		value = op2.Value
	case *object.RefNotFoundObject:
		return op2
	default:
		if op1 == nil {
			e.logger.Error(errcode.EZ80_OP, stmt.Context)
		} else {
			e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		}
		return object.ERROR
	}

	// addr check
	addr, ok := e.intToWord(value)
	if !ok {
		e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, value, value), stmt.Context)
	}

	// addr set
	code := &object.CodeObject{Code: []byte{0xcd, byte(addr & 0xff), byte((addr >> 8) & 0xff)}, TStates: [2]byte{17, 5}, Context: stmt.Context}
	// CALL
	if op1 == nil {
		return code
	}
	// CALL cc
	flag := -1
	switch op1 := op1.(type) {
	case *object.FlagObject:
		flag = op1.Flag
	case *object.RegisterObject: // C レジスタを CY に読み替える
		flag = op1.Register
	case *object.RefNotFoundObject:
		return op1
	default:
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return object.ERROR
	}

	index, ok := Z80FlagIndex[flag]
	if !ok {
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return object.ERROR
	}
	code.Code[0] = 0xc4 | index<<3
	code.TStates = [2]byte{17, 5}
	return code
}

func (e *Evaluator) evalZ80_RST(stmt *parser.Z80Instruction, op1, op2 object.Object, _ TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0xc7}, TStates: [2]byte{11, 4}, Context: stmt.Context}

	addr := 0
	switch op1 := op1.(type) {
	case *object.NumberObject:
		addr = op1.Value
		if addr&0xc7 != 0 {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_RST, addr, addr), stmt.Context)
			return object.ERROR
		}
	case *object.RefNotFoundObject:
		return op1
	default:
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return object.ERROR
	}

	// addr set
	code.Code[0] |= byte(addr)

	return code
}

func (e *Evaluator) evalZ80_JP(stmt *parser.Z80Instruction, op1, op2 object.Object, _ TEnv) object.Object {

	value := 0
	switch op2 := op2.(type) {
	case *object.NumberObject:
		value = op2.Value

	case *object.RefNotFoundObject:
		return op2
	case *object.RegIndirectObject:
		// レジスタ間接
		if op2.Displacement != 0 {
			e.logger.Error(errcode.EZ80_JP_INDIRECT_DISP, stmt.Context)
			return object.ERROR
		}
		switch op2.Register {
		case parser.Z80_REG_HL:
			return &object.CodeObject{Code: []byte{0xe9}, TStates: [2]byte{4, 1}, Context: stmt.Context}
		case parser.Z80_REG_IX:
			return &object.CodeObject{Code: []byte{0xdd, 0xe9}, TStates: [2]byte{8, 2}, Context: stmt.Context}
		case parser.Z80_REG_IY:
			return &object.CodeObject{Code: []byte{0xfd, 0xe9}, TStates: [2]byte{8, 2}, Context: stmt.Context}
		default:
			e.logger.Error(errcode.EZ80_JP_INDIRECT_REG, stmt.Context)
			return object.ERROR
		}
	case *object.AddrIndirectObject:
		e.logger.Error(errcode.EZ80_JP_INDIRECT_REG, stmt.Context)
		return object.ERROR
	default:
		if op1 == nil {
			e.logger.Error(errcode.EZ80_OP, stmt.Context)
		} else {
			e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		}
		return object.ERROR
	}

	// addr check
	addr, ok := e.intToWord(value)
	if !ok {
		e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, value, value), stmt.Context)
	}

	code := &object.CodeObject{Code: []byte{0xc3, byte(addr & 0xff), byte((addr >> 8) & 0xff)}, TStates: [2]byte{10, 3}, Context: stmt.Context}
	// JP nn
	if op1 == nil {
		return code
	}
	// JP cc, nn
	flag := -1
	switch op1 := op1.(type) {
	case *object.FlagObject:
		flag = op1.Flag
	case *object.RegisterObject: // C レジスタを CY に読み替える
		flag = op1.Register
	case *object.RefNotFoundObject:
		return op1
	default:
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return object.ERROR
	}

	index, ok := Z80FlagIndex[flag]
	if !ok {
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return object.ERROR
	}
	code.Code[0] = 0xc2 | index<<3
	return code
}

func (e *Evaluator) evalZ80_JR(stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object {

	addr := 0
	switch op2 := op2.(type) {
	case *object.NumberObject:
		addr = op2.Value
	case *object.RefNotFoundObject:
		return op2
	default:
		if op1 == nil {
			e.logger.Error(errcode.EZ80_OP, stmt.Context)
		} else {
			e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		}
		return object.ERROR
	}

	// addr set
	ofs := addr - getLocationCounter(env) - 2
	// ofs 検査は Evaluator.Satge2 で行う。Stage1 ではラベルのアドレスが確定していないため
	if e.Stage2 && (ofs < -128 || ofs > 127) {
		e.logger.Error(fmt.Sprintf(errcode.EZ80_JR_RANGE, ofs, ofs), stmt.Context)
		return object.ERROR
	}

	code := &object.CodeObject{Code: []byte{0x18, byte(ofs)}, TStates: [2]byte{12, 3}, Context: stmt.Context}

	// JR e
	if op1 == nil {
		return code
	}

	code.Code[0] = 0x20
	// JR cc, e
	flag := -1
	switch op1 := op1.(type) {
	case *object.FlagObject:
		flag = op1.Flag
	case *object.RegisterObject: // C レジスタを CY に読み替える
		flag = op1.Register
	case *object.RefNotFoundObject:
		return op1
	default:
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return object.ERROR
	}
	if flag > parser.Z80_FLAG_C {
		e.logger.Error(fmt.Sprintf(errcode.EZ80_JR_FLAG, parser.TokenLiteral(flag)), stmt.Context)
		return object.ERROR
	}

	index, ok := Z80FlagIndex[flag]
	if !ok {
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return object.ERROR
	}
	code.Code[0] |= index << 3
	return code
}

func (e *Evaluator) evalZ80_DJNZ(stmt *parser.Z80Instruction, op1, _ object.Object, env TEnv) object.Object {
	addr := 0
	switch op := op1.(type) {
	case *object.NumberObject:
		addr = op.Value
	case *object.RefNotFoundObject:
		return op
	default:
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return object.ERROR
	}

	// addr set
	ofs := addr - getLocationCounter(env) - 2
	if ofs < -128 || ofs > 127 {
		e.logger.Error(fmt.Sprintf(errcode.EZ80_JR_RANGE, ofs, ofs), stmt.Context)
		return object.ERROR
	}

	return &object.CodeObject{Code: []byte{0x10, byte(ofs)}, TStates: [2]byte{13, 2}, Context: stmt.Context}
}
