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
	Stage2     bool // eval staget 2
	Counter    func() int
	incDirs    []string
}

func New(logger *logging.Logger, incDirs []string) *Evaluator {
	return &Evaluator{logger: logger, Resolved: true, Counter: makeCounter(0), incDirs: incDirs}
}

// start + 1 から順次生成するカウンタ関数を返す
func makeCounter(start int) func() int {
	return func() int {
		start++
		return start
	}
}

// Program 評価
func (e *Evaluator) EvalProgram(prog *parser.BlockStatement, pass int, env TEnv) object.Object {
	// 一旦 0 に初期化し ORG 他で上書きする
	initLocationCounter(env, 0)
	env.Set("$PASS", &object.NumberObject{Value: pass})
	// return e.evalBlockPtr(&prog.Statements, env)
	return e.evalStatement(prog, false, nil, env)
}
