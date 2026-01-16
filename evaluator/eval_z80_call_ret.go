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
		return &object.CodeObject{Code: []byte{0xc9}, CZ80: 10, Context: stmt.Context}
	}
	// RET cc
	index := -1
	switch op1 := op1.(type) {
	case *object.FlagObject:
		index = op1.Flag
	case *object.RegisterObject: // C レジスタを CY に読み替える
		index = op1.Register
	}

	flag, ok := Z80FlagIndex[index]
	if !ok {
		e.logger.Error(fmt.Sprintf(errcode.EZ80_FLAG, stmt.Op1.String()), stmt.Context)
		return object.ERROR
	}
	b := byte(0xc0 | flag<<3)
	return &object.CodeObject{Code: []byte{b}, CZ80: 11, Context: stmt.Context}
}
