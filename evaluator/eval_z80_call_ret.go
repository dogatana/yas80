package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) evalZ80_RET(stmt *parser.Z80Instruction, op1, _ object.Object, _ TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0xc9}, CZ80: 10, Context: stmt.Context}

	// RET
	if op1 == nil {
		return code
	}
	// RET cc
	index := -1
	switch op1 := op1.(type) {
	case *object.FlagObject:
		index = op1.Flag
	case *object.RegisterObject: // C レジスタを CY に読み替える
		index = op1.Register
	case *object.RefNotFoundObject:
		return code
	default:
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return code
	}

	flag, ok := Z80FlagIndex[index]
	if !ok {
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return code
	}
	b := byte(0xc0 | flag<<3)
	return &object.CodeObject{Code: []byte{b}, CZ80: 11, Context: stmt.Context}
}

func (e *Evaluator) evalZ80_CALL(stmt *parser.Z80Instruction, op1, op2 object.Object, _ TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0xcd, 0x00, 0x00}, CZ80: 17, Context: stmt.Context}

	value := 0
	switch op2 := op2.(type) {
	case *object.NumberObject:
		value = op2.Value
	case *object.RefNotFoundObject:
		return code
	default:
		if op1 == nil {
			e.logger.Error(errcode.EZ80_OP, stmt.Context)
		} else {
			e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		}
		return code
	}

	// addr check
	addr, ok := e.intToAddr(value)
	if !ok {
		e.logger.Warning(fmt.Sprintf(errcode.WROUND_ADDR, value, value), stmt.Context)
	}

	// addr set
	code.Code[1] = byte(addr & 0xff)
	code.Code[2] = byte((addr >> 8) & 0xff)

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
		return code
	default:
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return code
	}

	index, ok := Z80FlagIndex[flag]
	if !ok {
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return code
	}
	code.Code[0] = 0xc4 | index<<3
	return code
}

func (e *Evaluator) evalZ80_RST(stmt *parser.Z80Instruction, op1, op2 object.Object, _ TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0xc7}, CZ80: 11, Context: stmt.Context}

	addr := 0
	switch op1 := op1.(type) {
	case *object.NumberObject:
		addr = op1.Value
		if addr&0xc7 != 0 {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_RST, addr, addr), stmt.Context)
			return object.ERROR
		}
	case *object.RefNotFoundObject:
		return code
	default:
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return object.ERROR
	}

	// addr set
	code.Code[0] |= byte(addr)

	return code
}

func (e *Evaluator) evalZ80_JP(stmt *parser.Z80Instruction, op1, op2 object.Object, _ TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0xc3, 0x00, 0x00}, CZ80: 10, Context: stmt.Context}

	value := 0
	switch op2 := op2.(type) {
	case *object.NumberObject:
		value = op2.Value
	case *object.RefNotFoundObject:
		return code
	default:
		if op1 == nil {
			e.logger.Error(errcode.EZ80_OP, stmt.Context)
		} else {
			e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		}
		return code
	}

	// addr check
	addr, ok := e.intToAddr(value)
	if !ok {
		e.logger.Warning(fmt.Sprintf(errcode.WROUND_ADDR, value, value), stmt.Context)
	}

	// addr set
	code.Code[1] = byte(addr & 0xff)
	code.Code[2] = byte((addr >> 8) & 0xff)

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
		return code
	default:
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return code
	}

	index, ok := Z80FlagIndex[flag]
	if !ok {
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return code
	}
	code.Code[0] = 0xc4 | index<<3
	return code
}

func (e *Evaluator) evalZ80_JR(stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0x18, 0x00}, CZ80: 12, Context: stmt.Context}

	addr := 0
	switch op2 := op2.(type) {
	case *object.NumberObject:
		addr = op2.Value
	case *object.RefNotFoundObject:
		return code
	default:
		if op1 == nil {
			e.logger.Error(errcode.EZ80_OP, stmt.Context)
		} else {
			e.logger.Error(errcode.EZ80_OP2, stmt.Context)
		}
		return code
	}

	// addr set
	ofs := addr - getLocationCounter(env) - 2
	if ofs < -128 || ofs > 127 {
		e.logger.Error(fmt.Sprintf(errcode.EZ80_JR_RANGE, ofs, ofs), stmt.Context)
		return code
	}

	code.Code[1] = byte(ofs)

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
		return code
	default:
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return code
	}
	if flag > parser.Z80_FLAG_C {
		e.logger.Error(fmt.Sprintf(errcode.EZ80_JR_FLAG, parser.TokenLiteral(flag)), stmt.Context)
		return code
	}

	index, ok := Z80FlagIndex[flag]
	if !ok {
		e.logger.Error(errcode.EZ80_FLAG, stmt.Context)
		return code
	}
	code.Code[0] |= index << 3
	return code
}

func (e *Evaluator) evalZ80_DJNZ(stmt *parser.Z80Instruction, op1, _ object.Object, env TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0x10, 0x00}, CZ80: 13, Context: stmt.Context}

	addr := 0
	switch op := op1.(type) {
	case *object.NumberObject:
		addr = op.Value
	case *object.RefNotFoundObject:
		return code
	default:
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return code
	}

	// addr set
	ofs := addr - getLocationCounter(env) - 2
	if ofs < -128 || ofs > 127 {
		e.logger.Error(fmt.Sprintf(errcode.EZ80_JR_RANGE, ofs, ofs), stmt.Context)
		return code
	}

	code.Code[1] = byte(ofs)

	return code
}
