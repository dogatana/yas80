package evaluator

import (
	"fmt"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
)

var _ex_regs = map[int]bool{
	parser.Z80_REG_HL:   true,
	parser.Z80_REG_DE:   true,
	parser.Z80_REG_AF:   true,
	parser.Z80_REG_AFEX: true,
}

// EX
func (e *Evaluator) evalZ80_EX(stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object {
	if op1 == nil || op2 == nil {
		e.logger.Error(errcode.EZ80_OP_LESS, stmt.Context)
		return object.ERROR
	}
	if isRefNotFound(op1) {
		return op1
	}
	if isRefNotFound(op2) {
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

	code := &object.CodeObject{Code: []byte{0xeb}, TStates: [2]byte{4, 1}, Context: stmt.Context} // EX DE, HL
	switch {
	// EX rr, rr'
	case op1.Type() == object.OBJ_REGISTER && op2.Type() == object.OBJ_REGISTER:
		reg1 := op1.(*object.RegisterObject)
		if !_ex_regs[reg1.Register] {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(reg1.Register)), stmt.Context)
			return object.ERROR
		}
		reg2 := op2.(*object.RegisterObject)
		if !_ex_regs[reg2.Register] {
			e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(reg2.Register)), stmt.Context)
			return object.ERROR
		}
		switch {
		case reg1.Register == parser.Z80_REG_DE && reg2.Register == parser.Z80_REG_HL, reg1.Register == parser.Z80_REG_HL && reg2.Register == parser.Z80_REG_DE:
			return code // EX DE, HL / EX HL, DE
		case reg1.Register == parser.Z80_REG_AF && reg2.Register == parser.Z80_REG_AFEX, reg1.Register == parser.Z80_REG_AFEX && reg2.Register == parser.Z80_REG_AF:
			code.Code[0] = 0x08
			return code // EX AF, AF' / EX AF', AF
		default:
			e.logger.Error(fmt.Sprintf(errcode.EZ80_EX_REG, parser.TokenLiteral(reg1.Register), parser.TokenLiteral(reg2.Register)), stmt.Context)
			return object.ERROR
		}

	// EX (SP), rr
	case op1.Type() == object.OBJ_REG_INDIRECT && op2.Type() == object.OBJ_REGISTER:
		return e.evalExSpReg16(op1, op2, stmt.Context)

	// EX rr, (SP), => EX (SP), rr
	case op1.Type() == object.OBJ_REGISTER && op2.Type() == object.OBJ_REG_INDIRECT:
		return e.evalExSpReg16(op2, op1, stmt.Context)
	}

	e.logger.Error(errcode.EZ80_OP, stmt.Context)
	return object.ERROR
}

func (e *Evaluator) evalExSpReg16(op1, op2 object.Object, ctx TContext) object.Object {

	reg1 := op1.(*object.RegIndirectObject)
	if reg1.Register != parser.Z80_REG_SP {
		e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(reg1.Register)), ctx)
		return object.ERROR
	}

	code := &object.CodeObject{Code: []byte{0xe3}, TStates: [2]byte{19, 5}, Context: ctx} // EX (SP), HL

	reg2 := op2.(*object.RegisterObject)
	switch reg2.Register {
	case parser.Z80_REG_HL:
		return code // EX (SP), HL
	case parser.Z80_REG_IX:
		code.Code = []byte{0xdd, 0xe3}
		code.TStates = [2]byte{23, 6}
		return code
	case parser.Z80_REG_IY:
		code.Code = []byte{0xfd, 0xe3}
		code.TStates = [2]byte{23, 6}
		return code
	default:
		e.logger.Error(fmt.Sprintf(errcode.EZ80_OP_REG, parser.TokenLiteral(reg2.Register)), ctx)
		return object.ERROR
	}
}

// IM n
func (e *Evaluator) evalZ80_IM(stmt *parser.Z80Instruction, op, _ object.Object, env TEnv) object.Object {
	if op == nil {
		e.logger.Error(errcode.EZ80_OP_LESS, stmt.Context)
		return object.ERROR
	}

	// IM n
	mode := 0
	switch op := op.(type) {
	case *object.RefNotFoundObject:
		return op
	case *object.NullObject:
		e.logger.Error(errcode.EZ80_OP_NULL, stmt.Context)
		return object.ERROR

	case *object.NumberObject:
		mode = op.Value
		if mode < 0 || mode > 2 {
			e.logger.Error(errcode.EZ80_IM_RANGE, stmt.Context)
			return object.ERROR
		}

		code := &object.CodeObject{Code: []byte{0xed, 0x46}, TStates: [2]byte{8, 3}, Context: stmt.Context}
		switch mode {
		case 1:
			code.Code[1] = 0x56
		case 2:
			code.Code[1] = 0x5E
		default:
			// IM 0
		}
		return code

	default:
		// op = nil
		e.logger.Error(errcode.EZ80_OP, stmt.Context)
		return object.ERROR
	}
}
