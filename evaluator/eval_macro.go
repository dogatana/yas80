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
		if obj.Type() == object.OBJ_MACRO {
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
	stmts := make([]parser.Statement, 0, len(bs.Block))

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
func (e *Evaluator) evalMacroCallStatementEx(stmt *parser.MacroCallStatement, checkExitM bool, ectx TContext, env TEnv) object.Object {
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

	if ectx == nil {
		// トップレベルからのマクロ展開の場合 Offset は 1 からに変更する
		tmp := *stmt.Context
		ectx = &tmp
	}
	ectx.Offset++

	expandingMacro[stmt.Name] = true
	obj = e.expandMacroEx(stmt, macro, checkExitM, ectx, env)
	expandingMacro[stmt.Name] = false

	return obj
}

// 変更版 evalMacroBlockStatement
func (e *Evaluator) evalMacroBlockStatementEx(node parser.Statement, checkExitM bool, ectx TContext, env TEnv) object.Object {
	objects := []object.Object{}
	stmts := []parser.Statement{}

	stmt, ok := node.(*parser.MacroBlockStatement)
	if !ok {
		panic("invalid node type in evalMacroBlockStatement")
	}

	if e.Debug == 6 {
		fmt.Println("----- mbs start")
		fmt.Printf("name %s\n", stmt.Name)
		for i, s := range stmt.Block {
			fmt.Printf("%d: %s\n", i, s.String())
		}
		fmt.Println("----- mbs end")
	}
	// REPT の場合は $I, $COUNT 用の環境を作成する
	if stmt.Count != 0 {
		env = object.NewMacroEnvironment(env)
	}

	block := stmt.Block
	comment := stmt.Name
	if comment == "REPT" {
		comment = fmt.Sprintf("REPT %d/%d", stmt.Index, stmt.Count)
	}
	co := &object.CommentObject{Comments: []string{comment}, Context: stmt.Context}
	objects = append(objects, co)

	for _, node := range block {
	EVAL_AGAIN:
		switch stmt := node.(type) {
		case *parser.MacroStatement:
			// マクロ定義時除外されているはず
			e.logger.Error(errcode.EMACRO_NEST, stmt.Context)
			continue

		case *parser.IfStatement:
			obj := e.evalStatementEx(stmt, true, ectx, env)
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
			if len(bo.Block) > 0 && bo.Block[0].Type() == object.OBJ_EXITM {
				goto BREAK
			}

		case *parser.ExitmStatement:
			stmts = append(stmts, stmt)
			objects = append(objects, &object.ExitmObject{})
			goto BREAK

		case *parser.MacroCallStatement:
			obj := e.evalStatementEx(node, true, ectx, env)
			if isError(obj) {
				continue
			}
			if obj.Type() != object.OBJ_NODES {
				panic("not nodes object")
			}
			bs := &parser.MacroBlockStatement{Name: stmt.Name, Block: obj.(*object.StatemetnsObject).Statements, Context: stmt.Context}

			if e.Debug == 5 {
				fmt.Println("-- expanded start")
				for _, n := range bs.Block {
					fmt.Println(n.String())
				}
				fmt.Println("-- expanded end")
			}

			node = bs
			goto EVAL_AGAIN // 戻らずに自己ループする
			// stmts = append(stmts, bs)
			// e.Resolved = false

		// マクロ ブロック (展開済み)
		case *parser.MacroBlockStatement:
			stmts = append(stmts, stmt)
			obj := e.evalMacroBlockStatementEx(stmt, true, ectx, env)
			bo, ok := obj.(*object.BlockObject)
			if !ok {
				panic("not block object")
			}
			objs := bo.Block
			objects = append(objects, objs...)

		default:
			obj := e.evalStatementEx(node, true, ectx, env)
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

// REPT 展開
func (e *Evaluator) evalReptStatementEx(stmt *parser.ReptStatement, _ bool, ectx TContext, env TEnv) object.Object {
	obj := e.evalExpression(stmt.MaxCount, env, stmt.Context)
	if isError(obj) || isRefNotFound(obj) {
		return obj
	}

	var num int
	var values []any

	switch obj := obj.(type) {
	case *object.NumberObject:
		num = obj.Value
	case *object.ArrayObject:
		num = len(obj.Values)
		values = make([]any, len(obj.Values))
		for i, o := range obj.Values {
			values[i] = o
		}
	default:
		e.logger.Error(errcode.EREPT_COUNT, stmt.Context)
		return object.ERROR
	}

	if ectx == nil {
		// トップレベルからのマクロ展開の場合 Offset は 1 から
		tmp := *stmt.Context
		ectx = &tmp
	}
	ectx.Offset++

	mb := &parser.MacroBlockStatement{
		Name:    "REPT",
		Count:   num,
		Context: stmt.Context,
	}

	for i := 0; i < num; i++ {
		bs := &parser.BlockStatement{}

		var s parser.Statement
		s = &parser.SetSysVarStatement{
			Name:    "$COUNT",
			Value:   &object.NumberObject{Value: num},
			Context: stmt.Context}
		s.ReplaceContext(*ectx)
		bs.Block = append(bs.Block, s)

		s = &parser.SetSysVarStatement{
			Name:    "$I",
			Value:   &object.NumberObject{Value: i},
			Context: stmt.Context}
		s.ReplaceContext(*ectx)
		bs.Block = append(bs.Block, s)

		if values != nil {
			s = &parser.SetSysVarStatement{
				Name:    "$V",
				Value:   values[i],
				Context: stmt.Context}
			s.ReplaceContext(*ectx)
			bs.Block = append(bs.Block, s)
		}

		objs := e.expandReptBlock(stmt, env, ectx)
		if isError(objs) {
			continue
		}
		bs.Block = append(bs.Block, objs.(*object.StatemetnsObject).Statements...)
		mb.Block = append(mb.Block, bs)
	}

	mb.Block = append(mb.Block, &parser.CommentStatement{Text: "ENDR", Context: ectx})

	return &object.StatementObject{Statement: mb}
}
func (e *Evaluator) evalReptStatement(stmt *parser.ReptStatement, env TEnv) object.Object {
	obj := e.evalExpression(stmt.MaxCount, env, stmt.Context)
	if isError(obj) || isRefNotFound(obj) {
		return obj
	}

	var num int
	var values []any

	switch obj := obj.(type) {
	case *object.NumberObject:
		num = obj.Value
	case *object.ArrayObject:
		num = len(obj.Values)
		values = make([]any, len(obj.Values))
		for i, o := range obj.Values {
			values[i] = o
		}
	default:
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
	ectx.Offset++

	nodes := []parser.Statement{}
	for i := 0; i < num; i++ {
		mb := &parser.MacroBlockStatement{
			Name:    "REPT",
			Count:   num,
			Index:   i,
			Context: stmt.Context,
		}
		mb.ReplaceContext(*ectx)

		var s parser.Statement
		// ectx.Offset++
		s = &parser.SetSysVarStatement{
			Name:    "$COUNT",
			Value:   &object.NumberObject{Value: num},
			Context: stmt.Context}
		s.ReplaceContext(*ectx)
		mb.Block = append(mb.Block, s)

		// ectx.Offset++
		s = &parser.SetSysVarStatement{
			Name:    "$I",
			Value:   &object.NumberObject{Value: i},
			Context: stmt.Context}
		s.ReplaceContext(*ectx)
		mb.Block = append(mb.Block, s)

		if values != nil {
			// ectx.Offset++
			s = &parser.SetSysVarStatement{
				Name:    "$V",
				Value:   values[i],
				Context: stmt.Context}
			s.ReplaceContext(*ectx)
			mb.Block = append(mb.Block, s)
		}

		// ectx.Offset++
		objs := e.expandReptBlock(stmt, env, ectx)
		if isError(objs) {
			continue
		}
		mb.Block = append(mb.Block, objs.(*object.StatemetnsObject).Statements...)
		nodes = append(nodes, mb)
	}

	s := &parser.CommentStatement{Text: "ENDR", Context: stmt.Context}
	s.ReplaceContext(*ectx)
	nodes = append(nodes, s)

	bs := &parser.BlockStatement{Block: nodes}
	return &object.StatementObject{Statement: bs}
}
