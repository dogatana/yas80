package evaluator

import (
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

type Evaluator struct {
	logger     *logging.Logger
	Debug      int
	Resolved   bool
	CodeStable bool
	Counter    func() int
}

func New(logger *logging.Logger) *Evaluator {
	return &Evaluator{logger: logger, Resolved: true, Counter: makeCounter(0)}
}

// start + 1 から順次生成するカウンタ関数を返す
func makeCounter(start int) func() int {
	return func() int {
		start++
		return start
	}
}

// Program 評価
func (e *Evaluator) EvalProgram(prog *parser.BlockStatement, env TEnv) object.Object {
	// 一旦 0 に初期化し ORG 他で上書きする
	initLocationCounter(env, 0)
	// return e.evalBlockPtr(&prog.Statements, env)
	return e.evalStatementEx(prog, false, nil, env)
}
