package evaluator

import (
	"fmt"
	"yas80/errcode"
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

	ret, ok := e.evalMacroBlockStatement(stmt.Body, newEnv).(*object.BlockObject)
	if !ok {
		panic(fmt.Sprintf("call macro %s returns %T(%#v)", stmt.Name, ret, ret))
	}
	return ret
}

// 複合文 BlockStatement
func (e *Evaluator) evalMacroBlockStatement(block *parser.BlockStatement, env object.Environment) object.Object {
	ret := &object.BlockObject{Block: []object.Object{}}
	nodes := []parser.Node{}

	for _, node := range block.Block {
		switch stmt := node.(type) {
		case *parser.ReturnStatement:
			e.logger.Warning(fmt.Sprintf(errcode.WMACRO_NOT_ALLOWED, "RETURN 文"), stmt.Context)
			continue
		case *parser.MacroStatement:
			e.logger.Warning(fmt.Sprintf(errcode.WMACRO_NOT_ALLOWED, "MACRO 文"), stmt.Context)
			continue
		case *parser.ProcStatement:
			e.logger.Warning(fmt.Sprintf(errcode.WMACRO_NOT_ALLOWED, "PROC 文"), stmt.Context)
			continue
		case *parser.FuncStatement:
			e.logger.Warning(fmt.Sprintf(errcode.WMACRO_NOT_ALLOWED, "FUNC 文"), stmt.Context)
			continue
		case *parser.EnumStatement:
			e.logger.Warning(fmt.Sprintf(errcode.WMACRO_NOT_ALLOWED, "ENUM 文"), stmt.Context)
			continue

		case *parser.ExitmStatement:
			block.Block = nodes
			return ret

			//ネストした IF の中の EXITM を動作させるには？

		case *parser.MacroCallStatement:
			obj := e.Eval(node, env)
			if isError(obj) {
				continue
			}
			expanded, ok := obj.(*object.NodeObject)
			if !ok {
				panic(fmt.Sprintf("not *object.NodeObject. got %T", obj))
			}
			nodes = append(nodes, expanded.Node)
			e.Resolved = false
		default:
			obj := e.Eval(node, env)
			if isError(obj) {
				continue
			}
			nodes = append(nodes, stmt)
			ret.Block = append(ret.Block, obj)
		}
	}
	block.Block = nodes
	return ret
}
