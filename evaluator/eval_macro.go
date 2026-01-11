package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

// macro 定義
func (e *Evaluator) evalMacroStatement(stmt *parser.MacroStatement, env TEnv) object.Object {
	name := stmt.Name
	if name[0] == '@' || name[0] == '.' {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_NAME, name), stmt.Context)
		return object.ERROR
	}
	if obj, ok := env.Get(name); ok {
		if obj.Type() == object.MACRO_OBJ {
			e.logger.Error(fmt.Sprintf(errcode.EMACRO_DUP, name), stmt.Context)
		} else {
			e.logger.Error(fmt.Sprintf(errcode.EMACRO_USED, name), stmt.Context)
		}
		return object.ERROR
	}
	// 無効な文をチェック
	e.filterValidStatementForMacro(stmt.Body)

	obj := &object.MacroObject{Name: name, Params: stmt.Params, Body: stmt.Body}
	env.Set(name, obj)
	return obj // 形式上必要
}

// macro 内で利用可能な文を抽出するフィルタ
func (e *Evaluator) filterValidStatementForMacro(bs *parser.BlockStatement) {
	stmts := make([]parser.Node, 0, len(bs.Block))

	for _, stmt := range bs.Block {
		switch stmt := stmt.(type) {
		// error
		case *parser.MacroStatement:
			e.logger.Error(errcode.EMACRO_NEST, stmt.GetContext())

		// warning
		case *parser.FuncStatement:
			e.logger.Warning(errcode.WSCOPE_MACRO, stmt.Context)
		case *parser.ReturnStatement:
			e.logger.Warning(errcode.WSCOPE_MACRO, stmt.Context)
		case *parser.ProcStatement:
			e.logger.Warning(errcode.WSCOPE_MACRO, stmt.Context)
		case *parser.ProcBlockStatement:
			e.logger.Warning(errcode.WSCOPE_MACRO, stmt.Context)
		case *parser.EnumStatement:
			e.logger.Warning(errcode.WSCOPE_MACRO, stmt.Context)

		// 要再帰チェック
		case *parser.IfStatement:
			if bs, ok := stmt.Consequence.(*parser.BlockStatement); ok {
				e.filterValidStatementForMacro(bs)
			}
			if bs, ok := stmt.Alternative.(*parser.BlockStatement); ok {
				e.filterValidStatementForMacro(bs)
			}
			stmts = append(stmts, stmt)
		case *parser.BlockStatement:
			e.filterValidStatementForMacro(stmt)
			stmts = append(stmts, stmt)

		default:
			// 利用可能
			stmts = append(stmts, stmt)
		}
	}
	bs.Block = stmts
}

// マクロ展開が再帰しているかのチェック用
var expandingMacro map[string]bool = map[string]bool{}

// マクロ評価（展開のみで引数は評価しない）
func (e *Evaluator) evalMacroCallStatement(stmt *parser.MacroCallStatement, env TEnv) object.Object {
	obj, ok := env.Get(stmt.Name)
	if !ok {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_UNDEF, stmt.Name), stmt.Context)
		return object.ERROR
	}
	macro, ok := obj.(*object.MacroObject)
	if !ok {
		// rule 上発生しない
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

	var ectx TContext
	if stmt.Context.Offset == 0 {
		// トップレベルからのマクロ展開の場合 Offset は 1 からに変更する
		tmp := *stmt.Context
		ectx = &tmp
	} else {
		// マクロ内部からのマクロ展開の場合、Offset は前から継続する
		ectx = stmt.Context
	}

	expandingMacro[stmt.Name] = true
	nodes := e.expandMacro(stmt, macro, env, ectx)
	expandingMacro[stmt.Name] = false

	return nodes
}

// macro 用 BlockStatement 評価
func (e *Evaluator) evalMacroBlockStatement(node parser.Node, env TEnv) object.Object {
	objects := []object.Object{}
	stmts := []parser.Node{}

	var block []parser.Node

	switch node := node.(type) {
	case *parser.MacroBlockStatement:
		// REPT の場合は $I, $COUNT 用の環境を作成する
		if node.Count != 0 {
			env = object.NewMacroEnvironment(env)
		}
		block = node.Block
		comment := node.Name
		if comment == "REPT" {
			comment = fmt.Sprintf("REPT %d", node.Count)
		}
		co := &object.CommentObject{Comments: []string{comment}, Context: node.Context}
		objects = append(objects, co)
	case *parser.BlockStatement:
		block = node.Block
	default:
		panic("invalid node type in evalMacroBlockStatement")
	}

	for _, node := range block {
	EVAL_AGAIN:
		switch stmt := node.(type) {
		case *parser.MacroStatement:
			// マクロ定義時除外されているはず
			e.logger.Error(errcode.EMACRO_NEST, stmt.Context)
			continue

		case *parser.IfStatement:
			// exitm の評価をするため、評価関数を渡す
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
			bs := &parser.MacroBlockStatement{Name: stmt.Name, Block: obj.(*object.NodesObject).Nodes, Context: stmt.Context}
			// fmt.Println("-- expanded")
			// for _, n := range bs.Block {
			// 	fmt.Println(n.String())
			// }
			// fmt.Println("-- expanded")

			node = bs
			goto EVAL_AGAIN // 戻らずに自己ループする
			// stmts = append(stmts, bs)
			// e.Resolved = false

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

func (e *Evaluator) evalReptStatement(stmt *parser.ReptStatement, env TEnv) object.Object {
	obj := e.evalExpression(stmt.MaxCount, env, stmt.Context)
	if isError(obj) || isRefNotFound(obj) {
		return obj
	}
	num, ok := obj.(*object.NumberObject)
	if !ok {
		e.logger.Error(errcode.EREPT_COUNT, stmt.Context)
		return object.ERROR
	}

	var ectx TContext
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
		if isError(objs) {
			continue
		}
		nodes = append(nodes, objs.(*object.NodesObject).Nodes...)
	}

	mb := &parser.MacroBlockStatement{
		Name:    "REPT",
		Count:   num.Value,
		Block:   nodes,
		Context: stmt.Context}
	return &object.NodeObject{Node: mb}
}
