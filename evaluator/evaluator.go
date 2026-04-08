package evaluator

import (
	"github.com/dogatana/yas80/logging"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
)

type Evaluator struct {
	logger         *logging.Logger
	Debug          int
	Resolved       bool
	CodeStable     bool
	Stage1MaxCount int  // eval stage 1 の最大評価回数 256回
	Stage2MaxCount int  // eval stage 2 の最大評価回数 16回
	Stage2         bool // eval stage 2 評価中
	Counter        func() int
	incDirs        []string
}

func New(logger *logging.Logger, incDirs []string) *Evaluator {
	e := &Evaluator{logger: logger, Resolved: true, Counter: makeCounter(0), incDirs: incDirs}
	e.Stage1MaxCount = 256
	e.Stage2MaxCount = 16
	return e
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
