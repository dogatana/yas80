package evaluator

import (
	"fmt"
	"yas80/object"
	"yas80/parser"
)

// マクロ Body 評価
func (e *Evaluator) evalExpandedMacroCallStatement(stmt *parser.ExpandedMacroCallStatement, env object.Environment) object.Object {
	// 引数を評価し、仮引数名で環境に設定
	newEnv := object.NewMacroEnvironment(env)
	for i, param := range stmt.Params {
		v := e.Eval(stmt.Args.Expressions[i], env)
		if isError(v) || isRefNotFound(v) {
			return v
		}
		newEnv.Set("@@"+param, v)
	}

	object.PrintEnv(newEnv)

	ret, ok := e.evalBlockStatement(stmt.Body, newEnv).(*object.BlockObject)
	if !ok {
		panic(fmt.Sprintf("call macro %s returns %T(%#v)", stmt.Name, ret, ret))
	}
	return ret
}
