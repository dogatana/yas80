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
		e.logger.Error(errcode.EZ80_NOT_FLAG, stmt.Context)
		return code
	}

	flag, ok := Z80FlagIndex[index]
	if !ok {
		e.logger.Error(fmt.Sprintf(errcode.EZ80_NOT_FLAG, stmt.Op1.String()), stmt.Context)
		return object.ERROR
	}
	b := byte(0xc0 | flag<<3)
	return &object.CodeObject{Code: []byte{b}, CZ80: 11, Context: stmt.Context}
}

func (e *Evaluator) evalZ80_CALL(stmt *parser.Z80Instruction, op1, op2 object.Object, _ TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0xcd, 0x00, 0x00}, CZ80: 17, Context: stmt.Context}

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
		return object.ERROR
	}

	// addr set
	code.Code[1] = byte(addr & 0xff)
	code.Code[2] = byte((addr >> 8) & 0xff)

	// CALL
	if op1 == nil {
		return code
	}
	// CALL cc
	index := -1
	switch op1 := op1.(type) {
	case *object.FlagObject:
		index = op1.Flag
	case *object.RegisterObject: // C レジスタを CY に読み替える
		index = op1.Register
	case *object.RefNotFoundObject:
		return code
	default:
		e.logger.Error(errcode.EZ80_NOT_FLAG, stmt.Context)
		return code
	}

	flag, ok := Z80FlagIndex[index]
	if !ok {
		e.logger.Error(fmt.Sprintf(errcode.EZ80_NOT_FLAG, stmt.Op1.String()), stmt.Context)
		return object.ERROR
	}
	code.Code[0] = byte(0xc4 | flag<<3)
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
	code.Code[0] |= byte(addr << 3)

	return code
}
