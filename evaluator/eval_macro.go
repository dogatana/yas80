package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

// マクロ展開が再帰しているかのチェック用
var expandingMacro map[string]bool = map[string]bool{}

// マクロ評価（展開のみで引数は評価しない）
func (e *Evaluator) evalMacroCallStatement(stmt *parser.MacroCallStatement, env object.Environment) object.Object {
	obj, ok := env.Get(stmt.Name)
	if !ok {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_NOT_FOUND, stmt.Name), stmt.Context)
		return object.ERROR
	}
	macro, ok := obj.(*object.MacroObject)
	if !ok {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_NOT_MACRO, stmt.Name), stmt.Context)
		return object.ERROR
	}
	if len(stmt.Args.Expressions) != len(macro.Params) {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_ARG_COUNT, stmt.Name), stmt.Context)
		return object.ERROR
	}

	if expandingMacro[stmt.Name] {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_CYCLIC, stmt.Name), stmt.Context)
		return object.ERROR
	}
	expandingMacro[stmt.Name] = true
	nodes := e.expandMacro(stmt, macro)
	expandingMacro[stmt.Name] = false

	return nodes
}

// // マクロ Body 評価
// func (e *Evaluator) evalExpandedMacroCallStatement(stmt *parser.ExpandedMacroCallStatement, env object.Environment) object.Object {
// 	if expandingMacro[stmt.Name] {
// 		e.logger.Error(fmt.Sprintf(errcode.EMACRO_CYCLIC, stmt.Name), stmt.Context)
// 		return object.ERROR
// 	}
// 	expandingMacro[stmt.Name] = true
// 	// 引数を評価し、仮引数名で環境に設定
// 	newEnv := object.NewMacroEnvironment(env)
// 	for i, param := range stmt.Params {
// 		v := e.evalExpression(stmt.Args.Expressions[i], env, stmt.Context)
// 		if isError(v) || isRefNotFound(v) {
// 			return v
// 		}
// 		newEnv.Set("@@"+param, v)
// 	}

// 	// object.PrintEnv(newEnv)

// 	ret, ok := e.evalMacroBlockStatement(stmt.Body, newEnv).(*object.BlockObject)
// 	if !ok {
// 		panic(fmt.Sprintf("call macro %s returns %T(%#v)", stmt.Name, ret, ret))
// 	}
// 	expandingMacro[stmt.Name] = false
// 	return ret
// }

// macro 用 BlockStatement 評価
func (e *Evaluator) evalMacroBlockStatement(block *parser.MacroBlockStatement, env object.Environment) []object.Object {
	objects := []object.Object{}
	stmts := []parser.Node{}

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

		// case *parser.IfStatement:
		// 	obj := e.evalIfStatementWithFunc(stmt, env, e.evalMacroBlockStatement)
		// 	if isError(obj) {
		// 		continue
		// 	}
		// 	nodes = append(nodes, stmt)
		// 	if isRefNotFound(obj) {
		// 		continue
		// 	}
		// 	if obj, ok := obj.(*object.BlockObject); ok {
		// 		ret.Block = append(ret.Block, obj.Block...)
		// 		if ret.Block[len(ret.Block)-1].Type() == object.EXITM_OBJ {
		// 			goto BREAK
		// 		}
		// 	}
		// 	ret.Block = append(ret.Block, obj)

		case *parser.ExitmStatement:
			stmts = append(stmts, stmt)
			objects = append(objects, &object.ExitmObject{})
			goto BREAK

		case *parser.MacroCallStatement:
			obj := e.evalStatement(node, env)
			if isError(obj) {
				continue
			}
			if obj.Type() != object.NODES_OBJ {
				panic("not nodes object")
			}
			stmts = append(stmts, &parser.MacroBlockStatement{
				Name:  stmt.Name,
				Block: obj.(*object.NodesObject).Nodes})
			e.Resolved = false

		// マクロ ブロック (展開済み)
		case *parser.MacroBlockStatement:
			stmts = append(stmts, stmt)
			objs := e.evalMacroBlockStatement(stmt, env)
			objects = append(objects, objs...)
			// 評価結果の末尾が EXITM なら評価を終了し戻る
			if len(objs) >= 1 && objs[0].Type() == object.EXITM_OBJ {
				goto BREAK
			}

		default:
			obj := e.evalStatement(node, env)
			if isError(obj) {
				continue
			}
			stmts = append(stmts, stmt)
			objects = append(objects, obj)
		}
	}
BREAK:
	block.Block = stmts
	return objects
}

// マクロ Body 評価
// func (e *Evaluator) evalExpandedMacroCallStatement(stmt *parser.ExpandedMacroCallStatement, env object.Environment) object.Object {
// 	if expandingMacro[stmt.Name] {
// 		e.logger.Error(fmt.Sprintf(errcode.EMACRO_CYCLIC, stmt.Name), stmt.Context)
// 		return object.ERROR
// 	}
// 	expandingMacro[stmt.Name] = true
// 	// 引数を評価し、仮引数名で環境に設定
// 	newEnv := object.NewMacroEnvironment(env)
// 	for i, param := range stmt.Params {
// 		v := e.evalExpression(stmt.Args.Expressions[i], env, stmt.Context)
// 		if isError(v) || isRefNotFound(v) {
// 			return v
// 		}
// 		newEnv.Set("@@"+param, v)
// 	}

// 	// object.PrintEnv(newEnv)

// 	ret, ok := e.evalMacroBlockStatement(stmt.Body, newEnv).(*object.BlockObject)
// 	if !ok {
// 		panic(fmt.Sprintf("call macro %s returns %T(%#v)", stmt.Name, ret, ret))
// 	}
// 	expandingMacro[stmt.Name] = false
// 	return ret
// }

// // macro 用 BlockStatement 評価
// func (e *Evaluator) evalMacroBlockStatement(block *parser.BlockStatement, env object.Environment) object.Object {
// 	ret := &object.BlockObject{Block: []object.Object{}}
// 	nodes := []parser.Node{}

// 	for _, node := range block.Block {
// 		switch stmt := node.(type) {
// 		case *parser.ReturnStatement:
// 			e.logger.Warning(fmt.Sprintf(errcode.WMACRO_NOT_ALLOWED, "RETURN 文"), stmt.Context)
// 			continue
// 		case *parser.MacroStatement:
// 			e.logger.Warning(fmt.Sprintf(errcode.WMACRO_NOT_ALLOWED, "MACRO 文"), stmt.Context)
// 			continue
// 		case *parser.ProcStatement:
// 			e.logger.Warning(fmt.Sprintf(errcode.WMACRO_NOT_ALLOWED, "PROC 文"), stmt.Context)
// 			continue
// 		case *parser.FuncStatement:
// 			e.logger.Warning(fmt.Sprintf(errcode.WMACRO_NOT_ALLOWED, "FUNC 文"), stmt.Context)
// 			continue
// 		case *parser.EnumStatement:
// 			e.logger.Warning(fmt.Sprintf(errcode.WMACRO_NOT_ALLOWED, "ENUM 文"), stmt.Context)
// 			continue

// 		case *parser.IfStatement:
// 			obj := e.evalIfStatementWithFunc(stmt, env, e.evalMacroBlockStatement)
// 			if isError(obj) {
// 				continue
// 			}
// 			nodes = append(nodes, stmt)
// 			if isRefNotFound(obj) {
// 				continue
// 			}
// 			if obj, ok := obj.(*object.BlockObject); ok {
// 				ret.Block = append(ret.Block, obj.Block...)
// 				if ret.Block[len(ret.Block)-1].Type() == object.EXITM_OBJ {
// 					goto BREAK
// 				}
// 			}
// 			ret.Block = append(ret.Block, obj)

// 		case *parser.ExitmStatement:
// 			// exitm で BlockStatement を書き換えて戻る
// 			nodes = append(nodes, stmt)
// 			ret.Block = append(ret.Block, &object.ExitmObject{})
// 			goto BREAK

// 		case *parser.MacroCallStatement:
// 			obj := e.evalStatement(node, env)
// 			if isError(obj) {
// 				continue
// 			}
// 			expanded, ok := obj.(*object.NodeObject)
// 			if !ok {
// 				panic(fmt.Sprintf("not *object.NodeObject. got %T", obj))
// 			}
// 			nodes = append(nodes, expanded.Node)
// 			e.Resolved = false

// 		default:
// 			obj := e.evalStatement(node, env)
// 			if isError(obj) {
// 				continue
// 			}
// 			nodes = append(nodes, stmt)
// 			ret.Block = append(ret.Block, obj)
// 		}
// 	}
// BREAK:
// 	block.Block = nodes
// 	return ret
// }
