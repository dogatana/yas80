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
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_UNDEF, stmt.Name), stmt.Context)
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
		e.logger.Error(fmt.Sprintf(errcode.EMACROCALL_CYCLIC, stmt.Name), stmt.Context)
		return object.ERROR
	}
	expandingMacro[stmt.Name] = true
	nodes := e.expandMacro(stmt, macro)
	expandingMacro[stmt.Name] = false

	return nodes
}

// macro 用 BlockStatement 評価
func (e *Evaluator) evalMacroBlockStatement(node parser.Node, env object.Environment) object.Object {
	var block []parser.Node

	switch node := node.(type) {
	case *parser.MacroBlockStatement:
		block = node.Block
	case *parser.BlockStatement:
		block = node.Block
	default:
		panic("invalid node type in evalMacroBlockStatement")
	}

	objects := []object.Object{}
	stmts := []parser.Node{}

	for _, node := range block {
		switch stmt := node.(type) {
		case *parser.ReturnStatement:
			e.logger.Warning(fmt.Sprintf(errcode.WSCOPE_MACRO, "RETURN 文"), stmt.Context)
			continue
		case *parser.MacroStatement:
			e.logger.Warning(fmt.Sprintf(errcode.WSCOPE_MACRO, "MACRO 文"), stmt.Context)
			continue
		case *parser.ProcStatement:
			e.logger.Warning(fmt.Sprintf(errcode.WSCOPE_MACRO, "PROC 文"), stmt.Context)
			continue
		case *parser.FuncStatement:
			e.logger.Warning(fmt.Sprintf(errcode.WSCOPE_MACRO, "FUNC 文"), stmt.Context)
			continue
		case *parser.EnumStatement:
			e.logger.Warning(fmt.Sprintf(errcode.WSCOPE_MACRO, "ENUM 文"), stmt.Context)
			continue

		case *parser.IfStatement:
			obj := e.evalIfStatementWithFunc(stmt, env, e.evalMacroBlockStatement)
			if isError(obj) {
				continue
			}
			stmts = append(stmts, stmt)
			if isRefNotFound(obj) {
				continue
			}
			bo, ok := obj.(*object.BlockObject)
			if !ok {
				panic("not block object")
			}
			objects = append(objects, bo.Block...)
			if len(bo.Block) > 0 && bo.Block[0].Type() == object.EXITM_OBJ {
				goto BREAK
			}

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
			bs := &parser.MacroBlockStatement{Name: stmt.Name, Block: obj.(*object.NodesObject).Nodes}
			// fmt.Println("-- expanded")
			// for _, n := range bs.Block {
			// 	fmt.Println(n.String())
			// }
			// fmt.Println("-- expanded")

			stmts = append(stmts, bs)
			e.Resolved = false

		// マクロ ブロック (展開済み)
		case *parser.MacroBlockStatement:
			stmts = append(stmts, stmt)
			obj := e.evalMacroBlockStatement(stmt, env)
			bo, ok := obj.(*object.BlockObject)
			if !ok {
				panic("not block object")
			}
			objs := bo.Block
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
	switch node := node.(type) {
	case *parser.MacroBlockStatement:
		node.Block = stmts
	case *parser.BlockStatement:
		node.Block = stmts
	default:
		panic("invalid node type in evalMacroBlockStatement")
	}
	return &object.BlockObject{Block: objects}
}
