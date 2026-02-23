package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

// MUL
func (e *Evaluator) evalZ80_MUL(stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object {
	// $R800 を確認
	if obj, ok := env.Get("$R800"); !ok {
		panic("no $R800")
	} else if obj.(*object.NumberObject).Value == 0 {
		e.logger.Error(errcode.ER800, stmt.Context)
		return object.ERROR
	}

	if op1 == nil || op2 == nil {
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return object.ERROR
	}
	if isRefNotFound(op1) {
		e.Resolved = false
		return op1
	}
	if isRefNotFound(op2) {
		e.Resolved = false
		return op2
	}
	if op1 == object.NULL {
		e.logger.Error(errcode.EZ80_OP1_NULL, stmt.Context)
		return object.ERROR
	}
	if op2 == object.NULL {
		e.logger.Error(errcode.EZ80_OP2_NULL, stmt.Context)
		return object.ERROR
	}

	r1, ok1 := op1.(*object.RegisterObject)
	r2, ok2 := op2.(*object.RegisterObject)

	if !ok1 || !ok2 {
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return object.ERROR
	}

	switch r1.Register {
	case parser.Z80_REG_A:
		if r2.Register < parser.Z80_REG_B || r2.Register > parser.Z80_REG_E {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(r2.Register)), stmt.Context)
			return object.ERROR
		}
		// MUL A, r (B,C,D, E)
		return &object.CodeObject{Code: []byte{0xed, 0xc1 | Z80Reg8Index[r2.Register]<<3}, TStates: [2]byte{0, 14}, Context: stmt.Context}

	case parser.Z80_REG_HL:
		if r2.Register != parser.Z80_REG_BC && r2.Register != parser.Z80_REG_SP {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(r2.Register)), stmt.Context)
			return object.ERROR
		}
		// MUL HL, rr (BC, SP)
		return &object.CodeObject{Code: []byte{0xed, 0xc3 | Z80Reg16IndexSP[r2.Register]<<4}, TStates: [2]byte{0, 36}, Context: stmt.Context}

	default:
		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(r1.Register)), stmt.Context)
		return object.ERROR

	}
}
