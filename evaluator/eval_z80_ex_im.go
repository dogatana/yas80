package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

// BIT, SET RES
func (e *Evaluator) evalZ80_IM(stmt *parser.Z80Instruction, op, _ object.Object, env TEnv) object.Object {
	code := &object.CodeObject{Code: []byte{0xed, 0x46}, CZ80: 8, Context: stmt.Context}

	// IM n
	mode := 0
	switch op := op.(type) {
	case *object.RefNotFoundObject:
		e.Resolved = false
		return code
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP_NULL, stmt.Context)
		return object.ERROR
	case *object.NumberObject:
		mode = op.Value
		if mode < 0 || mode > 2 {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_IM_RANGE, mode, mode), stmt.Context)
			return code
		}
		switch mode {
		case 1:
			code.Code[1] = 0x56
		case 2:
			code.Code[1] = 0x5E
		}
		return code
	default:
		// op = nil
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return code
	}
}
