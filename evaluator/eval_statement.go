package evaluator

import (
	"fmt"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
)

// evalStatement Ex
func (e *Evaluator) evalStatement(stmt parser.Statement, checkExitM bool, ectx TContext, env TEnv) object.Object {

	switch stmt := stmt.(type) {

	// Z80 命令
	case *parser.Z80Instruction:
		obj := e.evalZ80Instruction(stmt, env)
		if obj.Type() == object.OBJ_CODE {
			code := obj.(*object.CodeObject)
			// アドレス設定はコード生成後
			code.Addr = getLocationCounter(env)
			// 生成コードのサイズ ロケーションカウンタを進める
			if err := advanceLocationCounters(env, code.Size()); err != nil {
				e.logger.Error(err.Error(), stmt.Context)
			}
		}
		return obj

	// // ラベル定義
	case *parser.LabelStatement:
		return e.evalLabelStatement(stmt, env)

	// // ORG
	// case *parser.OrgStatement:
	// 	return e.evalOrgStatement(stmt, env)

	// // INCLUDE
	// case *parser.IncludeStatement:
	// 	// ソースを表示するため nil とする
	// 	return &object.CommentObject{Text: nil, Context: stmt.Context}

	// // PROC
	// case *parser.ProcStatement:
	// 	if object.InProcEnv(env) {
	// 		e.logger.Error(fmt.Sprintf(errcode.ESCOPE_PROC, "PROC"), stmt.Context)
	// 		return object.ERROR
	// 	}
	// 	return e.evalProcStatement(stmt, env)

	// // PROC BLOCK
	// case *parser.ProcBlockStatement:
	// 	return e.evalProcBlockStatement(stmt, checkExitM, ectx, env)

	// // DS/DSB/DSW
	// case *parser.DataStoreStatement:
	// 	obj := e.evalDataStoreStatement(stmt, env)
	// 	if isError(obj) {
	// 		return object.ERROR
	// 	}
	// 	if err := advanceLocationCounters(env, len(obj.(*object.CodeObject).Code)); err != nil {
	// 		e.logger.Error(err.Error(), stmt.Context)
	// 	}
	// 	return obj

	// // DB/DW/DD
	// case *parser.DataDefineStatement:
	// 	obj := e.evalDataDefineStatement(stmt, env)
	// 	if isError(obj) {
	// 		return object.ERROR
	// 	}
	// 	if err := advanceLocationCounters(env, len(obj.(*object.CodeObject).Code)); err != nil {
	// 		e.logger.Error(err.Error(), stmt.Context)
	// 	}
	// 	return obj

	// // 定数定義
	// case *parser.ConstStatement:
	// 	return e.evalConstStatement(stmt, env)

	// // マクロ定義
	// case *parser.MacroStatement:
	// 	return e.evalMacroStatement(stmt, env)

	// // マクロ呼出し
	// case *parser.MacroCallStatement:
	// 	return e.evalMacroCallStatement(stmt, checkExitM, ectx, env)

	// // マクロ呼出し
	// case *parser.MacroBlockStatement:
	// 	return e.evalMacroBlockStatement(stmt, checkExitM, ectx, env)

	// // exitm
	// case *parser.ExitmStatement:
	// 	return &object.ExitmObject{}

	// // rept
	// case *parser.ReptStatement:
	// 	return e.evalReptStatement(stmt, checkExitM, ectx, env)

	// // var
	// case *parser.VariableStatement:
	// 	return e.evalVariableStatement(stmt, env)

	// // 代入文
	// case *parser.AssignStatement:
	// 	return e.evalAssignStatement(stmt, env)

	// // if
	// case *parser.IfStatement:
	// 	return e.evalIfStatement(stmt, checkExitM, ectx, env)

	// block
	case *parser.BlockStatement:
		return e.evalBlockStatement(stmt, checkExitM, ectx, env)

	// // func
	// case *parser.FuncStatement:
	// 	return e.evalFuncStatement(stmt, env)

	// // return
	// case *parser.ReturnStatement:
	// 	return e.evalReturnStatement(stmt, env)

	// // charmap
	// case *parser.CharmapStatement:
	// 	return e.evalCharmapStatement(stmt, env)

	// file
	case *parser.FileStatement:
		return &object.FileObject{Filename: stmt.Filename, Line: stmt.Line}

	// // comment
	// case *parser.CommentStatement:
	// 	return &object.CommentObject{Text: stmt.Text, Context: stmt.Context}

	// // システム変数設定
	// case *parser.SetSysVarStatement:
	// 	var v object.Object

	// 	if obj, ok := stmt.Value.(object.Object); ok {
	// 		v = obj
	// 	} else {
	// 		obj := e.evalExpression(stmt.Value.(parser.Expression), env, stmt.Context)
	// 		if isError(obj) {
	// 			return object.ERROR
	// 		}
	// 		v = obj
	// 	}

	// 	env.Set(stmt.Name, v)
	// 	comment := fmt.Sprintf("%s = %s", stmt.Name, v.String())
	// 	return &object.CommentObject{Text: comment, SetSysVar: true, Context: stmt.Context}

	// // enum
	// case *parser.EnumStatement:
	// 	obj := e.evalEnumStatement(stmt, env)
	// 	if isError(obj) {
	// 		return object.ERROR
	// 	}
	// 	return obj

	// // end
	// case *parser.EndStatement:
	// 	obj := e.evalEndStatement(stmt, env)
	// 	if isError(obj) {
	// 		return object.ERROR
	// 	}
	// 	return obj

	// Null
	case *parser.NullStatement:
		return object.NULL

	default:
		e.logger.Error(fmt.Sprintf(errcode.ENOT_IMPL_STMT, stmt), nil) // TODO
		return object.ERROR
	}
}

// 複合文 BlockStatement
func (e *Evaluator) evalBlockStatement(bs *parser.BlockStatement, checkExitM bool, ectx TContext, env TEnv) object.Object {
	block := &object.BlockObject{Block: []object.Object{}}
	stmts := make([]parser.Statement, 0, len(bs.Block))

	// bs.Block の内容は書き換えるケースがあるので、インデックスでループする
LOOP:
	for i := range len(bs.Block) {
	EVAL_AGAIN:
		stmt := bs.Block[i]

		obj := e.evalStatement(stmt, checkExitM, ectx, env)
		if isError(obj) {
			continue
		}

		// bs.Block[i] を無効化
		switch stmt.NodeType() {
		// case parser.NODE_FUNC_STMT, parser.NODE_MACRO_STMT, parser.NODE_ENUM_STMT:
		case parser.NODE_MACRO_STMT, parser.NODE_ENUM_STMT:
			continue
		}

		switch obj := obj.(type) {
		case *object.ReturnObject: // TODO: Return
			block.Block = append(block.Block, obj)
			stmts = append(stmts, stmt)
			break LOOP

		case *object.ExitmObject:
			if checkExitM {
				block.Block = append(block.Block, obj)
				stmts = append(stmts, stmt)
				break LOOP
			}
			e.logger.Error(fmt.Sprintf(errcode.ESCOPE_MACRO, "EXITM"), stmt.GetContext())
			// stmt, obj とも append しない

		case *object.BlockObject:
			if len(obj.Block) == 0 {
				break
			}
			block.Block = append(block.Block, obj.Block...)
			if block.Block[len(block.Block)-1].Type() == object.OBJ_RETURN { // TODO: Return
				stmts = append(stmts, stmt)
				break LOOP
			}
			if checkExitM && block.Block[len(block.Block)-1].Type() == object.OBJ_EXITM {
				stmts = append(stmts, stmt)
				break LOOP
			}

		// ProcStatement => StatementObject(ProcBlockStatement)
		case *object.StatementObject:
			bs.Block[i] = obj.Statement
			goto EVAL_AGAIN

		default:
			block.Block = append(block.Block, obj)
		}
		stmts = append(stmts, stmt)
	}

	bs.Block = stmts
	return block
}

// // PROC
//
//	func (e *Evaluator) evalProcStatement(node *parser.ProcStatement, env TEnv) object.Object {
//		e.concatenateSymbol(&node.Name, env, node.Context)
//
//		id, ok := node.Name.(*parser.Ident)
//		if !ok {
//			e.logger.Error(errcode.ECONCAT_TYPE, node.Context)
//			return object.ERROR
//		}
//		name := id.Name
//		obj, ok := env.Get(name)
//		if ok {
//			switch obj := obj.(type) {
//			case *object.ProcObject:
//				e.logger.Error(fmt.Sprintf(errcode.EPROC_DUP, name), node.Context)
//				return object.ERROR
//			case *object.SymbolObject:
//				if obj.SymType != object.SYM_UNKNOWN {
//					e.logger.Error(fmt.Sprintf(errcode.EPROC_USED, name), node.Context)
//					return object.ERROR
//				}
//				// SYM_UNKNOWN（前方参照）なら proc として登録
//			default:
//				e.logger.Error(fmt.Sprintf(errcode.EPROC_USED, name), node.Context)
//				return object.ERROR
//			}
//		}
//		penv := object.NewProcEnvironment(env)
//		env.Set(name, &object.ProcObject{Name: name, Addr: getLocationCounter(env), Env: penv})
//
//		e.filterValidStatementForProc(node.Block)
//
//		pbs := &parser.ProcBlockStatement{Name: name, Block: node.Block.Block, Context: node.Context}
//		return &object.StatementObject{Statement: pbs}
//	}
//
// // PROC 内で有効な文をフィルタ
//
//	func (e *Evaluator) filterValidStatementForProc(bs *parser.BlockStatement) {
//		stmts := make([]parser.Statement, 0, len(bs.Block))
//
//		for _, stmt := range bs.Block {
//			switch stmt := stmt.(type) {
//			// error
//			case *parser.ProcStatement:
//				e.logger.Error(errcode.EPROC_NEST, stmt.GetContext())
//
//			// warning
//			case *parser.ReturnStatement:
//				e.logger.Warning(errcode.WSCOPE_PROC, stmt.Context)
//
//			// 要再帰チェック
//			case *parser.IfStatement:
//				if bs, ok := stmt.Consequence.(*parser.BlockStatement); ok {
//					e.filterValidStatementForProc(bs)
//				}
//				if bs, ok := stmt.Alternative.(*parser.BlockStatement); ok {
//					e.filterValidStatementForProc(bs)
//				}
//				stmts = append(stmts, stmt)
//			case *parser.BlockStatement:
//				e.filterValidStatementForProc(stmt)
//				stmts = append(stmts, stmt)
//
//			default:
//				// 利用可能
//				stmts = append(stmts, stmt)
//			}
//		}
//		bs.Block = stmts
//
// }
//
// // PROC BLOCK
//
//	func (e *Evaluator) evalProcBlockStatement(stmt *parser.ProcBlockStatement, checkExitM bool, ectx TContext, env TEnv) object.Object {
//		obj, ok := env.Get(stmt.Name)
//		if !ok {
//			panic(fmt.Sprintf("no ProcEnv(%s)", stmt.Name))
//		}
//		if po, ok := obj.(*object.ProcObject); ok {
//			po.Addr = getLocationCounter(env) // proc アドレスを更新
//		} else {
//			panic(fmt.Sprintf("invalid ProcjObject(%s)", stmt.Name))
//		}
//
//		// ProcObject は Environment intterface を実装
//		bs := &parser.BlockStatement{Block: stmt.Block}
//		return e.evalStatement(bs, checkExitM, ectx, obj.(object.Environment))
//	}
//
// // ORG
//
//	func (e *Evaluator) evalOrgStatement(stmt *parser.OrgStatement, env TEnv) object.Object {
//		e.concatenateSymbol(&stmt.Address, env, stmt.Context)
//		obj := e.evalExpression(stmt.Address, env, stmt.Context)
//
//		var value int
//		switch obj := obj.(type) {
//		case *object.ErrorObject:
//			return obj
//		case *object.RefNotFoundObject:
//			return obj
//		case *object.NullObject:
//			e.logger.Error(errcode.EORG_NULL, stmt.Context)
//			return object.ERROR
//
//		case *object.NumberObject:
//			value = obj.Value
//
//		default:
//			e.logger.Error(errcode.EORG_VALUE, stmt.Context)
//			return object.ERROR
//		}
//
//		addr, ok := e.intToWord(value)
//		if !ok {
//			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, value, value), stmt.Context)
//		}
//
//		// ABS は $, $$ REL は $ のみ変更
//		setLocationCounter(env, addr)
//		if stmt.AllocType == parser.ALLOC_ABS {
//			setAllocLocationCounter(env, addr)
//		}
//		return &object.OrgObject{Addr: addr, AllocType: int(stmt.AllocType)}
//	}
//
// // END
//
//	func (e *Evaluator) evalEndStatement(stmt *parser.EndStatement, env TEnv) object.Object {
//		if stmt.Start == nil {
//			return &object.EntryObject{StartAddr: -1}
//		}
//
//		e.concatenateSymbol(&stmt.Start, env, stmt.Context)
//		obj := e.evalExpression(stmt.Start, env, stmt.Context)
//		switch obj := obj.(type) {
//		case *object.ErrorObject:
//			return obj
//		case *object.RefNotFoundObject, *object.NullObject:
//			e.logger.Error(errcode.EEND_NULL, stmt.Context)
//			return object.ERROR
//
//		case *object.NumberObject:
//			return &object.EntryObject{StartAddr: obj.Value}
//
//		default:
//			e.logger.Error(errcode.EEND_VALUE, stmt.Context)
//			return object.ERROR
//		}
//	}

// ラベル定義文
func (e *Evaluator) evalLabelStatement(stmt *parser.LabelStatement, env TEnv) object.Object {
	e.concatenateSymbol(&stmt.Name, env, stmt.Context)
	addr := getLocationCounter(env)
	obj := e.exprToLabel(stmt.Name, env, stmt.Context)
	if isError(obj) {
		return obj
	}
	return &object.CodeObject{Addr: addr, Code: []byte{}, Context: stmt.Context}
}

// parser.Label 評価&環境登録
func (e *Evaluator) evalLabel(label *parser.Label, env TEnv) object.Object {
	name := label.Name

	// // 匿名ラベル処理
	// if label.LabelType == parser.NODE_ANON_LABEL {
	// 	if object.OuterEnvType(env) != object.ENV_PROC {
	// 		e.logger.Error(fmt.Sprintf(errcode.ESCOPE_PROC, name), label.Context)
	// 		return object.ERROR
	// 	}
	// 	name = label.Name
	// 	if util.IsAnonDef(name) {
	// 		return e.evalAnonymouseLable(label, env)
	// 	}
	// 	// @F @B @nF @nB
	// 	e.logger.Error(fmt.Sprintf(errcode.EANON_LABEL_REF_ONLY, name), label.Context)
	// 	return object.ERROR
	// }

	// // . @ 1 文字のラベルは利用不可
	// if name == "." || name == "@" {
	// 	e.logger.Error(fmt.Sprintf(errcode.ESYM_INVALID, name), label.Context)
	// 	return object.ERROR
	// }

	// switch {
	// case name[0] == '.' && object.OuterEnvType(env) != object.ENV_PROC:
	// 	e.logger.Error(fmt.Sprintf(errcode.ESCOPE_PROC, name), label.Context)
	// 	return object.ERROR
	// case name[0] == '@' && env.EnvType() != object.ENV_MACRO:
	// 	e.logger.Error(fmt.Sprintf(errcode.ESCOPE_MACRO, name), label.Context)
	// 	return object.ERROR
	// }

	obj, ok := env.Get(name)
	if !ok {
		// 環境にないなら新規登録
		sym := object.NewLabelSymbol(name, getLocationCounter(env), label.Context)
		env.Set(name, sym)
		return sym
	}

	sym, ok := obj.(*object.SymbolObject)
	if !ok || sym.SymType != object.SYM_LABEL && sym.SymType != object.SYM_UNKNOWN {
		// Symbol 以外か、SYM_LABEL でない場合
		e.logger.Error(fmt.Sprintf(errcode.ELABEL_USED, name), label.Context)
		return object.ERROR
	}
	if sym.SymType == object.SYM_LABEL && !sym.Context.Equal(label.Context) {
		// ラベル 二重定義
		e.logger.Error(fmt.Sprintf(errcode.ELABEL_DUP, name), label.Context)
		return object.ERROR
	}

	// SYM_UNKNOWN の場合 SYM_LABEL として登録後値を更新
	if sym.SymType == object.SYM_UNKNOWN {
		sym = object.NewLabelSymbol(name, 0, label.Context)
		env.Set(name, sym)
	}
	// 値を更新
	sym.Value.(*object.NumberObject).Value = getLocationCounter(env)
	return sym
}

//
//// 匿名ラベル処理
//func (e *Evaluator) evalAnonymouseLable(label *parser.Label, env TEnv) object.Object {
//	// 匿名ラベル情報
//	pos := &object.AnonLabel{
//		Addr:     getLocationCounter(env),
//		Filename: label.Context.FileContent.Filename,
//		Line:     label.Context.Line}
//
//	obj, ok := env.Get(label.Name)
//	if !ok {
//		// 環境にないなら新規登録
//		obj := &object.AnonLabelsObject{Name: label.Name, Labels: []*object.AnonLabel{pos}}
//		env.Set(label.Name, obj)
//		return obj
//	}
//	// 追加
//	lo, ok := obj.(*object.AnonLabelsObject)
//	if !ok {
//		panic(fmt.Sprintf("invalid AnonLabelsObject: %#v", obj))
//	}
//	lo.Add(pos)
//	return lo
//}
//
//// const / equ 文
//func (e *Evaluator) evalConstStatement(node *parser.ConstStatement, env TEnv) object.Object {
//	e.concatenateSymbol(&node.Name, env, node.Context)
//	e.concatenateSymbol(&node.Value, env, node.Context)
//
//	id, ok := node.Name.(*parser.Ident)
//	if !ok {
//		// rule で回避されているため発生しない
//		return object.ERROR
//	}
//	name := id.Name
//
//	switch {
//	case name[0] == '.' && object.OuterEnvType(env) != object.ENV_PROC:
//		e.logger.Error(fmt.Sprintf(errcode.ESCOPE_PROC, name), node.Context)
//		return object.ERROR
//	case name[0] == '@' && env.EnvType() != object.ENV_MACRO:
//		e.logger.Error(fmt.Sprintf(errcode.ESCOPE_MACRO, name), node.Context)
//		return object.ERROR
//	}
//
//	// 定義済みならエラー
//	obj, ok := env.Get(name)
//	if ok {
//		switch obj := obj.(type) {
//		case *object.SymbolObject:
//			if obj.SymType == object.SYM_UNKNOWN {
//				// 不明シンボンルなら更新
//			} else if obj.Name != id.Name || !obj.Context.Equal(node.Context) {
//				// 別シンボルなら二重定義エラー
//				e.logger.Error(fmt.Sprintf(errcode.ECONST_DUP, name), node.Context)
//				return object.ERROR
//			}
//			// fmt.Printf("obj.Context %s\n", obj.Context.String())
//			// fmt.Printf("node.Context %s\n", node.Context.String())
//			// 同一シンボルなら更新
//		case *object.RefNotFoundObject:
//			// 未定で登録済なら更新
//		case *object.NullObject:
//			// NULL ならエラー
//			e.logger.Error(fmt.Sprintf(errcode.ECONST_NULL, name), node.Context)
//			return object.ERROR
//
//		default:
//			e.logger.Error(fmt.Sprintf(errcode.ECONST_USED, name), node.Context)
//			return object.ERROR
//		}
//	}
//
//	v := e.evalExpression(node.Value, env, node.Context)
//
//	switch v := v.(type) {
//	case *object.ErrorObject:
//		return object.ERROR
//
//	case *object.RefNotFoundObject:
//		names := removeSelfName(v.Names, name)
//		sym := object.NewConstSymbol(name, node.Value, object.NULL, names, node.Context)
//		env.Set(name, sym)
//		return object.NULL
//
//	case *object.NumberObject:
//		// NumberObject の copy を値とする Symbol を作成し環境へ登録
//		val := *v // copy
//		sym := object.NewConstSymbol(name, node.Value, &val, []string{}, node.Context)
//		env.Set(name, sym)
//		text := fmt.Sprintf("%04x(%d)", v.Value, v.Value)
//		return &object.TextObject{Text: text, Context: node.Context}
//
//	case *object.StringObject:
//		// StringObject の copy を値とする Symbol を作成し環境へ登録
//		val := *v // copy
//		sym := object.NewConstSymbol(name, node.Value, &val, []string{}, node.Context)
//		env.Set(name, sym)
//		text := fmt.Sprintf("%q", v.Value)
//		return &object.TextObject{Text: text, Context: node.Context}
//
//	case *object.RegisterObject:
//		// 値を SymbolObject として環境へ登録
//		sym := object.NewConstSymbol(name, node.Value, v, []string{}, node.Context)
//		env.Set(name, sym)
//		text := parser.TokenLiteral(v.Register)
//		return &object.TextObject{Text: text, Context: node.Context}
//
//	case *object.FlagObject:
//		// 値を SymbolObject として環境へ登録
//		sym := object.NewConstSymbol(name, node.Value, v, []string{}, node.Context)
//		env.Set(name, sym)
//		text := parser.TokenLiteral(v.Flag)
//		return &object.TextObject{Text: text, Context: node.Context}
//
//	case *object.FunctionObject, *object.ArrayObject:
//		// 値を SymbolObject として環境へ登録
//		sym := object.NewConstSymbol(name, node.Value, v, []string{}, node.Context)
//		env.Set(name, sym)
//		return &object.ValueObject{Value: v, Context: node.Context}
//
//	default:
//		if e.Debug > 0 {
//			fmt.Printf("const %s = %#v\n", name, v)
//		}
//		env.Set(name, v)
//		return v
//	}
//}
//
//// RefNotFoundObjectの依存リストから名前を削除
//func removeSelfName(names []string, name string) []string {
//	// slices パッケージ利用へ変更
//	s := slices.Clone(names)
//	return slices.DeleteFunc(s, func(v string) bool { return v == name })
//}
//
//// var 文
//func (e *Evaluator) evalVariableStatement(stmt *parser.VariableStatement, env TEnv) object.Object {
//	e.concatenateSymbol(&stmt.Value, env, stmt.Context)
//
//	id := stmt.Name.(*parser.Ident)
//	name := id.Name
//
//	if name == "_" {
//		e.logger.Error(errcode.EVAR_SYS, stmt.Context)
//		return object.ERROR
//	}
//
//	switch {
//	case name[0] == '.' && object.OuterEnvType(env) != object.ENV_PROC:
//		e.logger.Error(fmt.Sprintf(errcode.ESCOPE_PROC, name), stmt.Context)
//		return object.ERROR
//	case name[0] == '@' && env.EnvType() != object.ENV_MACRO:
//		e.logger.Error(fmt.Sprintf(errcode.ESCOPE_MACRO, name), stmt.Context)
//		return object.ERROR
//	}
//
//	// 定義済みで同じ Symbol でないならエラー
//	obj, ok := env.Get(name)
//	if ok {
//		sym, ok := obj.(*object.SymbolObject)
//		if !ok || sym.Name != name || sym.SymType != object.SYM_VAR || !sym.Context.Equal(stmt.Context) {
//			e.logger.Error(fmt.Sprintf(errcode.EVAR_USED, name), stmt.Context)
//		}
//	}
//
//	v := e.evalExpression(stmt.Value, env, stmt.Context)
//
//	switch v := v.(type) {
//	case *object.ErrorObject:
//		return object.ERROR
//
//	case *object.NumberObject:
//		// NumberObject の copy を値とする Symbol を作成し環境へ登録
//		val := *v // copy
//		sym := object.NewVarSymbol(name, stmt.Value, &val, []string{}, stmt.Context)
//		env.Set(name, sym)
//		text := fmt.Sprintf("%04x(%d)", v.Value, v.Value)
//		return &object.TextObject{Text: text, Context: stmt.Context}
//
//	case *object.StringObject:
//		// StringObject の copy を値とする Symbol を作成し環境へ登録
//		val := *v // copy
//		sym := object.NewVarSymbol(name, stmt.Value, &val, []string{}, stmt.Context)
//		env.Set(name, sym)
//		text := fmt.Sprintf("%q", v.Value)
//		return &object.TextObject{Text: text, Context: stmt.Context}
//
//	case *object.RegisterObject:
//		// 値を持つ Symbol を作成し環境へ登録
//		sym := object.NewVarSymbol(name, stmt.Value, v, []string{}, stmt.Context)
//		env.Set(name, sym)
//		text := parser.TokenLiteral(v.Register)
//		return &object.TextObject{Text: text, Context: stmt.Context}
//
//	case *object.FlagObject:
//		// 値を持つ Symbol を作成し環境へ登録
//		sym := object.NewVarSymbol(name, stmt.Value, v, []string{}, stmt.Context)
//		env.Set(name, sym)
//		text := parser.TokenLiteral(v.Flag)
//		return &object.TextObject{Text: text, Context: stmt.Context}
//
//	case *object.FunctionObject, *object.ArrayObject:
//		// 値を持つ Symbol を作成し環境へ登録
//		sym := object.NewVarSymbol(name, stmt.Value, v, []string{}, stmt.Context)
//		env.Set(name, sym)
//		return &object.ValueObject{Value: v, Context: stmt.Context}
//
//	default:
//		e.logger.Error(fmt.Sprintf(errcode.EVAR_VALUE, name), stmt.Context)
//		return object.ERROR
//	}
//}
//
//// 代入
//func (e *Evaluator) evalAssignStatement(stmt *parser.AssignStatement, env TEnv) object.Object {
//	e.concatenateSymbol(&stmt.Left, env, stmt.Context)
//	e.concatenateSymbol(&stmt.Value, env, stmt.Context)
//
//	id, ok := stmt.Left.(*parser.Ident)
//	if !ok {
//		e.logger.Error(errcode.EASSIGN_LEFT, stmt.Context)
//		return object.ERROR
//	}
//	name := id.Name
//
//	switch {
//	case name[0] == '.' && object.OuterEnvType(env) != object.ENV_PROC:
//		e.logger.Error(fmt.Sprintf(errcode.ESCOPE_PROC, name), stmt.Context)
//		return object.ERROR
//	case name[0] == '@' && env.EnvType() != object.ENV_MACRO:
//		e.logger.Error(fmt.Sprintf(errcode.ESCOPE_MACRO, name), stmt.Context)
//		return object.ERROR
//	}
//
//	obj, ok := env.Get(name)
//	if !ok {
//		// 未定義ならエラー
//		e.logger.Error(fmt.Sprintf(errcode.EVAR_UNDEF, name), stmt.Context)
//		return object.ERROR
//
//	}
//	sym, ok := obj.(*object.SymbolObject)
//	if !ok || sym.SymType != object.SYM_VAR {
//		// SymbolObject でないか変数でない
//		e.logger.Error(errcode.EASSIGN_LEFT, stmt.Context)
//		return object.ERROR
//	}
//
//	value := e.evalExpression(stmt.Value, env, stmt.Context)
//
//	switch value.(type) {
//	case *object.ErrorObject:
//		return object.ERROR
//	case *object.RefNotFoundObject:
//		e.logger.Error(errcode.EASSIGN_VALUE, stmt.Context)
//		return object.ERROR
//	case *object.NullObject:
//		e.logger.Error(errcode.EASSIGN_VALUE, stmt.Context)
//		return object.ERROR
//	}
//	sym.Value = value
//
//	return e.valueToTextObject(value, stmt.Context)
//}
//
//// if 文
//func (e *Evaluator) evalIfStatement(stmt *parser.IfStatement, checkExitM bool, ectx TContext, env TEnv) object.Object {
//	obj := e.evalExpression(stmt.Condition, env, stmt.Context)
//	if isError(obj) {
//		return object.ERROR
//	}
//
//	if object.IsTruthy(obj) {
//		if stmt.Consequence == nil {
//			return object.NULL
//		}
//		return e.evalStatement(stmt.Consequence.(parser.Statement), checkExitM, ectx, env)
//	} else {
//		if stmt.Alternative == nil {
//			return object.NULL
//		}
//		return e.evalStatement(stmt.Alternative.(parser.Statement), checkExitM, ectx, env)
//	}
//}
//
//// func 文
//func (e *Evaluator) evalFuncStatement(stmt *parser.FuncStatement, env TEnv) object.Object {
//	name := stmt.Name
//	if name[0] == '@' || name[0] == '.' {
//		e.logger.Error(fmt.Sprintf(errcode.EFUNC_NAME, name), stmt.Context)
//		return object.ERROR
//	}
//	if obj, ok := env.Get(name); ok {
//		fobj, ok := obj.(*object.FunctionObject)
//		if !ok {
//			e.logger.Error(fmt.Sprintf(errcode.EFUNC_USED, name), stmt.Context)
//			return object.ERROR
//		}
//		if !fobj.Context.Equal(stmt.Context) {
//			e.logger.Error(fmt.Sprintf(errcode.EFUNC_DUP, name), stmt.Context)
//			return object.ERROR
//		}
//	} else {
//		// 無効な文をチェック
//		e.filterValidStatementForFunc(stmt.Block)
//	}
//
//	obj := &object.FunctionObject{Name: name, Params: stmt.Params, Body: stmt.Block, Env: env, Context: stmt.Context}
//	env.Set(name, obj)
//	return obj
//}
//
//// func 内で利用可能な文を抽出するフィルタ
//func (e *Evaluator) filterValidStatementForFunc(bs *parser.BlockStatement) {
//	stmts := make([]parser.Statement, 0, len(bs.Block))
//
//	for _, stmt := range bs.Block {
//		switch stmt := stmt.(type) {
//		// 利用可能
//		case *parser.ConstStatement, *parser.VariableStatement, *parser.AssignStatement, *parser.ReturnStatement:
//			stmts = append(stmts, stmt)
//
//		// 要再帰チェック
//		case *parser.FuncStatement:
//			e.filterValidStatementForFunc(stmt.Block)
//			stmts = append(stmts, stmt)
//
//		case *parser.IfStatement:
//			if bs, ok := stmt.Consequence.(*parser.BlockStatement); ok {
//				e.filterValidStatementForFunc(bs)
//			}
//			if bs, ok := stmt.Alternative.(*parser.BlockStatement); ok {
//				e.filterValidStatementForFunc(bs)
//			}
//			stmts = append(stmts, stmt)
//
//		case *parser.BlockStatement:
//			e.filterValidStatementForFunc(stmt)
//			stmts = append(stmts, stmt)
//
//		default:
//			// 利用不可
//			e.logger.Warning(errcode.WSCOPE_FUNC, stmt.GetContext())
//		}
//	}
//	bs.Block = stmts
//}
//
//// return 文
//func (e *Evaluator) evalReturnStatement(stmt *parser.ReturnStatement, env TEnv) object.Object {
//	var ret object.Object
//	if stmt.Value == nil {
//		ret = object.NULL
//	} else {
//		ret = e.evalExpression(stmt.Value, env, stmt.Context)
//	}
//	return &object.ReturnObject{Value: ret, LineNumber: stmt.Context.Line}
//}
//
//// enum 文
//func (e *Evaluator) evalEnumStatement(stmt *parser.EnumStatement, env TEnv) object.Object {
//	name := stmt.Name
//	obj, ok := env.Get(name)
//	if ok {
//		if obj.Type() == object.OBJ_ENUM {
//			e.logger.Error(fmt.Sprintf(errcode.EENUM_DUP, name), stmt.Context)
//		} else {
//			e.logger.Error(fmt.Sprintf(errcode.EENUM_USED, name), stmt.Context)
//		}
//		return object.ERROR
//	}
//
//	// EnumObject は Enviromnet interface を実装している
//	enum := &object.EnumObject{Name: name, Env: object.NewEnvironment(env)}
//	env.Set(name, enum)
//
//	// enum element の評価でエラーが発生した場合、単に無効とする
//	value := 0 // 初期値
//	for _, ele := range stmt.Elements.Elements {
//		ename := "." + ele.Name // . を先頭に付けたものを要素に内部名
//		if _, ok := enum.Get(ename); ok {
//			e.logger.Error(fmt.Sprintf(errcode.EENUM_ELE_DUP, name, ename), stmt.Context)
//			// 定義済みなら無効（無視）
//			continue
//		}
//		if ele.Value == nil {
//			esym := &object.SymbolObject{
//				Name:    ename,
//				SymType: object.SYM_CONST,
//				Value:   &object.NumberObject{Value: value}}
//			enum.Set(ename, esym)
//			value++
//			continue
//		}
//		v := e.evalExpression(ele.Value, enum, stmt.Context)
//		if isError(v) {
//			// 値がエラーなら無効（無視）
//			continue
//		}
//		sym := &object.SymbolObject{Name: ename, SymType: object.SYM_CONST}
//		switch v := v.(type) {
//		case *object.RefNotFoundObject:
//			e.logger.Error(errcode.EENUM_ELE_FWD, ele.Context)
//		case *object.NumberObject:
//			sym.Value = v
//			enum.Set(ename, sym)
//			value = v.Value + 1
//		case *object.StringObject:
//			sym.Value = v
//			enum.Set(ename, sym)
//		default:
//			e.logger.Error(errcode.EENUM_ELE_VALUE, ele.Context)
//		}
//	}
//	return enum
//}
//
