package evaluator

import (
	"fmt"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/intern"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
)

// macro 定義
func (e *Evaluator) evalMacroStatement(stmt *parser.MacroStatement, env TEnv) object.Object {
	name := stmt.NameID.String()
	if name[0] == '@' || name[0] == '.' {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_NAME, name), stmt.Context)
		return object.ERROR
	}
	if obj, ok := env.Get(stmt.NameID); ok {
		mo, ok := obj.(*object.MacroObject)
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.EMACRO_USED, name), stmt.Context)
			return object.ERROR
		}
		if !stmt.Context.Equal(mo.Context) { // Context が同一でない場合は重複定義エラー
			e.logger.Error(fmt.Sprintf(errcode.EMACRO_DUP, name), stmt.Context)
		}
		return mo
	}
	// 無効な文をチェック
	e.filterValidStatementForMacro(stmt.Body)

	obj := &object.MacroObject{NameID: stmt.NameID, Params: stmt.Params, End: stmt.End, Body: stmt.Body, Context: stmt.Context}
	env.Set(stmt.NameID, obj)
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
var expandingMacro = map[intern.SymbolID]bool{}

// マクロ評価（展開のみで引数は評価しない）
func (e *Evaluator) evalMacroCallStatement(stmt *parser.MacroCallStatement, checkExitM bool, ectx TContext, env TEnv) object.Object {
	// 組み込みマクロ
	if obj, ok := e.evalBuiltinMacro(stmt, env); ok {
		return obj
	}

	// ユーザ定義マクロ
	obj, ok := env.Get(stmt.NameID)
	if !ok {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_UNDEF, stmt.NameID), stmt.Context)
		return object.ERROR
	}
	macro, ok := obj.(*object.MacroObject)
	if !ok {
		// rule 上発生しない
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_NOT_MACRO, stmt.NameID), stmt.Context)
		return object.ERROR
	}
	if len(stmt.Args.Expressions) != len(macro.Params) {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_ARG_COUNT, stmt.NameID), stmt.Context)
		return object.ERROR
	}

	if expandingMacro[stmt.NameID] {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_CYCLIC, stmt.NameID), stmt.Context)
		return object.ERROR
	}

	if ectx == nil {
		// トップレベルからのマクロ展開の場合 Offset は 1 からに変更する
		tmp := *stmt.Context
		ectx = &tmp
		ectx.Offset++
	}

	expandingMacro[stmt.NameID] = true
	obj = e.expandMacro(stmt, macro, checkExitM, ectx, env)
	expandingMacro[stmt.NameID] = false

	return obj
}

// MacroBlockStatement 評価時に、Context.Offset を連番に設定する
func replaceOffset(stmt parser.Statement, fn func() int) {
	if stmt == nil {
		return
	}
	switch stmt := stmt.(type) {
	case *parser.MacroBlockStatement:
		if stmt.Context.Offset != 0 {
			stmt.Context.Offset = uint32(fn())
		}
		for _, s := range stmt.Block {
			replaceOffset(s, fn)
		}
	case *parser.BlockStatement:
		for _, s := range stmt.Block {
			replaceOffset(s, fn)
		}
	case *parser.IfStatement:
		stmt.Context.Offset = uint32(fn())
		replaceOffset(stmt.Consequence.(parser.Statement), fn)
		replaceOffset(stmt.Alternative.(parser.Statement), fn)

	case *parser.FileStatement:
		// do nothing
	default:
		stmt.GetContext().Offset = uint32(fn())
	}
}

// 変更版 evalMacroBlockStatement
func (e *Evaluator) evalMacroBlockStatement(node parser.Statement, checkExitM bool, ectx TContext, env TEnv) object.Object {
	objects := []object.Object{}
	stmts := []parser.Statement{}

	if !checkExitM {
		// TOP レベル
		fn := makeCounter(0)
		replaceOffset(node, fn)
	}

	stmt, ok := node.(*parser.MacroBlockStatement)
	if !ok {
		panic("invalid node type in evalMacroBlockStatement")
	}

	name := stmt.NameID.String()

	// Label
	if stmt.Label != nil {
		e.concatenateSymbol(&stmt.Label, env, stmt.Context)
		obj := e.exprToLabel(stmt.Label, env, stmt.Context)
		if isError(obj) {
			return object.ERROR
		}
	}

	if e.Debug == 6 {
		fmt.Println("----- mbs start")
		fmt.Printf("%s: ", stmt.Context.String())
		fmt.Printf("name %s\n", name)
		for i, s := range stmt.Block {
			fmt.Printf("%s: ", stmt.GetContext().String())
			fmt.Printf("%d: %s\n", i, s.String())
		}
		fmt.Println("----- mbs end")
	}
	// REPT の場合は $I, $COUNT 用の環境を作成する
	if name == "REPT" {
		env = object.NewMacroEnvironment(env)
	}

	if name != "REPT" {
		co := &object.SourceTextObject{Text: nil, Context: stmt.Context} // MACRO 呼出し行を表示
		objects = append(objects, co)
	} else if stmt.Context.Source == nil { // トップレベルの ENDR 行のみ表示する
		ctx := *stmt.Context
		ctx.Line = uint32(stmt.Start) // トップレベルの場合、Offset が0のため、この代入は利用されない
		ec := *stmt.Context
		ec.Source = &ctx
		co := &object.SourceTextObject{Text: nil, Context: &ec} // ソース行を表示のため Text は nil
		objects = append(objects, co)
	}

	block := stmt.Block
	for _, node := range block {
	EVAL_AGAIN:
		switch stmt := node.(type) {
		case *parser.MacroStatement:
			// マクロ定義時除外されているはず
			e.logger.Error(errcode.EMACRO_NEST, stmt.Context)
			continue

		case *parser.BlockStatement:
			obj := e.evalStatement(stmt, true, ectx, env)
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

		case *parser.IfStatement:
			obj := e.evalStatement(stmt, true, ectx, env)
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
			// 組み込みマクロは直接実行
			// ユーザ定義マクロは展開しMacroBlockStatementとしてから実行
			obj := e.evalStatement(node, true, ectx, env)
			if isError(obj) {
				continue
			}
			if obj.Type() != object.OBJ_NODES {
				// 組み込みマクロの結果は単一 Object
				stmts = append(stmts, stmt)
				objects = append(objects, obj)
				continue
			}
			bs := &parser.MacroBlockStatement{NameID: stmt.NameID, Block: obj.(*object.StatemetnsObject).Statements, Context: stmt.Context}

			if e.Debug == 5 {
				fmt.Println("-- expanded start")
				for _, n := range bs.Block {
					fmt.Println(n.String())
				}
				fmt.Println("-- expanded end")
			}

			node = bs
			goto EVAL_AGAIN // 戻らずに自己ループする

		// マクロ ブロック (展開済み)
		case *parser.MacroBlockStatement:
			stmts = append(stmts, stmt)
			obj := e.evalMacroBlockStatement(stmt, true, ectx, env)
			bo, ok := obj.(*object.BlockObject)
			if !ok {
				panic("not block object")
			}
			objs := bo.Block
			objects = append(objects, objs...)

		default:
			obj := e.evalStatement(node, true, ectx, env)
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

// // REPT 展開
// func (e *Evaluator) evalReptStatement(stmt *parser.ReptStatement, _ bool, ectx TContext, env TEnv) object.Object {
// 	obj := e.evalExpression(stmt.MaxCount, env, stmt.Context)
// 	if isError(obj) || isRefNotFound(obj) {
// 		return obj
// 	}

// 	var num int
// 	var values []any

// 	switch obj := obj.(type) {
// 	case *object.NumberObject:
// 		num = obj.Value
// 	case *object.ArrayObject:
// 		num = len(obj.Values)
// 		values = make([]any, len(obj.Values))
// 		for i, o := range obj.Values {
// 			values[i] = o
// 		}
// 	default:
// 		e.logger.Error(errcode.EREPT_COUNT, stmt.Context)
// 		return object.ERROR
// 	}

// 	if ectx == nil {
// 		// トップレベルからのマクロ展開の場合 Offset は 1 から
// 		tmp := *stmt.Context
// 		ectx = &tmp
// 		ectx.Offset++
// 	}

// 	mb := &parser.MacroBlockStatement{
// 		Label:   stmt.Label,
// 		Name:    "REPT",
// 		Count:   num,
// 		Start:   stmt.Start,
// 		Context: stmt.Context,
// 	}

// 	for i := 0; i < num; i++ {
// 		bs := &parser.BlockStatement{}

// 		var s parser.Statement
// 		s = &parser.SetSysVarStatement{
// 			Name:    "$COUNT",
// 			Value:   &object.NumberObject{Value: num},
// 			Context: stmt.Context}
// 		s.ReplaceContext(*ectx)
// 		bs.Block = append(bs.Block, s)

// 		s = &parser.SetSysVarStatement{
// 			Name:    "$I",
// 			Value:   &object.NumberObject{Value: i},
// 			Context: stmt.Context}
// 		s.ReplaceContext(*ectx)
// 		bs.Block = append(bs.Block, s)

// 		if values != nil {
// 			s = &parser.SetSysVarStatement{
// 				Name:    "$V",
// 				Value:   values[i],
// 				Context: stmt.Context}
// 			s.ReplaceContext(*ectx)
// 			bs.Block = append(bs.Block, s)
// 		}

// 		objs := e.expandReptBlock(stmt, env, ectx)
// 		if isError(objs) {
// 			continue
// 		}
// 		bs.Block = append(bs.Block, objs.(*object.StatemetnsObject).Statements...)
// 		mb.Block = append(mb.Block, bs)
// 	}

// 	return &object.StatementObject{Statement: mb}
// }
