package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

type evalZ80InstructionFunc func(e *Evaluator, stmt *parser.Z80Instruction, op1, op2 object.Object, env TEnv) object.Object

var evalZ80InstructionFuncs = map[int]evalZ80InstructionFunc{
	parser.Z80_INST_LD:   (*Evaluator).evalZ80LD,
	parser.Z80_INST_RET:  (*Evaluator).evalZ80_RET,
	parser.Z80_INST_CALL: (*Evaluator).evalZ80_CALL,
	parser.Z80_INST_RST:  (*Evaluator).evalZ80_RST,
	parser.Z80_INST_JP:   (*Evaluator).evalZ80_JP,
	parser.Z80_INST_JR:   (*Evaluator).evalZ80_JR,
	parser.Z80_INST_DJNZ: (*Evaluator).evalZ80_DJNZ,
	parser.Z80_INST_IN:   (*Evaluator).evalZ80_IN,
}

func (e *Evaluator) evalZ80Instruction(stmt *parser.Z80Instruction, env TEnv) object.Object {
	e.concatenateSymbol(&stmt.Label, env, stmt.Context)
	if stmt.Label != nil {
		e.exprToLabel(stmt.Label, env, stmt.Context)
	}

	// オペランドなし
	if stmt.NodeType() == parser.Z80_INST0 {
		code := Z80CodeTable0[stmt.Opcode]
		code.Context = stmt.Context
		return &code
	}

	// 1 or 2 オペランド
	e.concatenateSymbol(&stmt.Op1, env, stmt.Context)
	e.concatenateSymbol(&stmt.Op2, env, stmt.Context)

	// 命令毎の評価関数から evalExpression を呼び出すと循環参照エラーになるので事前に評価しておく
	var op1, op2 object.Object
	if stmt.Op1 != nil {
		op1 = e.evalExpression(stmt.Op1, env, stmt.Context)
		if isError(op1) {
			return op1
		}
	}
	if stmt.Op2 != nil {
		op2 = e.evalExpression(stmt.Op2, env, stmt.Context)
		if isError(op2) {
			return op2

		}
	}
	if fn, ok := evalZ80InstructionFuncs[stmt.Opcode]; ok {
		return fn(e, stmt, op1, op2, env)
	}

	// 未実装
	e.logger.Error(fmt.Sprintf(errcode.EZ80_NOT_IMPL, parser.TokenLiteral(stmt.Opcode)), stmt.Context)
	return object.ERROR
}
