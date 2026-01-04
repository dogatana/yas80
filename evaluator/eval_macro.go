package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/fileblock"
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
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_USED, stmt.Name), stmt.Context)
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

	var ectx *fileblock.Context
	if stmt.Context.Offset == 0 {
		// トップレベルからのマクロ展開の場合 Offset は 1 から
		tmp := *stmt.Context
		ectx = &tmp
	} else {
		// マクロ内部からのマクロ展開の場合、Offset は前からの継続
		ectx = stmt.Context
	}

	expandingMacro[stmt.Name] = true
	nodes := e.expandMacro(stmt, macro, env, ectx)
	expandingMacro[stmt.Name] = false

	return nodes
}

// macro 用 BlockStatement 評価
func (e *Evaluator) evalMacroBlockStatement(node parser.Node, env object.Environment) object.Object {
	var block []parser.Node

	switch node := node.(type) {
	case *parser.MacroBlockStatement:
		// REPT の場合は $I, $COUNT 用の環境を作成する
		if node.Count != 0 {
			env = object.NewMacroEnvironment(env)
		}
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
		case *parser.MacroStatement:
			e.logger.Error(errcode.EMACRO_NEST, stmt.Context)
			continue

		case *parser.ReturnStatement:
			e.logger.Warning(fmt.Sprintf(errcode.WSCOPE_MACRO, "RETURN 文"), stmt.Context)
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
			// if len(objs) >= 1 && objs[0].Type() == object.EXITM_OBJ {
			// 	goto BREAK
			// }

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

func (e *Evaluator) evalReptStatement(stmt *parser.ReptStatement, env object.Environment) object.Object {
	obj := e.evalExpression(stmt.MaxCount, env, stmt.Context)
	if isError(obj) || isRefNotFound(obj) {
		return obj
	}
	num, ok := obj.(*object.NumberObject)
	if !ok {
		e.logger.Error(errcode.EREPT_COUNT, stmt.Context)
		return object.ERROR
	}

	var ectx *fileblock.Context
	if stmt.Context.Offset == 0 {
		// トップレベルからのマクロ展開の場合 Offset は 1 から
		tmp := *stmt.Context
		ectx = &tmp
	} else {
		// マクロ内部からのマクロ展開の場合、Offset は前からの継続
		ectx = stmt.Context
	}
	ectx.Offset += 1

	// 環境に $COUNT を設定
	s := &parser.SetSysVarStatement{
		Name:    "$COUNT",
		Value:   &parser.NumberLiteral{Value: num.Value, Context: stmt.Context},
		Context: stmt.Context}
	s.ReplaceContext(*ectx)
	nodes := []parser.Node{s}
	for i := 0; i < num.Value; i++ {
		ectx.Offset += 1
		rs := &parser.SetSysVarStatement{
			Name:    "$I",
			Value:   &parser.NumberLiteral{Value: i, Context: stmt.Context},
			Context: stmt.Context}
		rs.ReplaceContext(*ectx)
		nodes = append(nodes, rs)
		objs := e.expandReptBlock(stmt, env, ectx)
		nodes = append(nodes, objs.(*object.NodesObject).Nodes...)
	}
	fmt.Printf("expanded rept %d\n", len(nodes))
	mb := &parser.MacroBlockStatement{
		Name:  "REPT",
		Count: num.Value,
		Block: nodes}
	return &object.NodeObject{Node: mb}
}
