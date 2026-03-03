package evaluator

import (
	"fmt"
	"slices"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) expandMacro(mcall *parser.MacroCallStatement, macro *object.MacroObject, checkExitM bool, ectx TContext, env TEnv) object.Object {
	stmts := []parser.Statement{}
	seq := e.Counter()

	// 仮引数と引数Node の紐づけ
	args := map[string]parser.Expression{}
	for i, param := range macro.Params {
		args[param] = mcall.Args.Expressions[i]
	}

	mfn := buildMangleNamesFunc(args, seq, mcall.Name)

	for _, stmt := range macro.Body.Block {
		news := e.mangleNamesInStatement(stmt, mfn, ectx)
		// news.ReplaceContext(*ectx)

		switch news := news.(type) {
		case *parser.MacroCallStatement:
			obj := e.evalMacroCallStatement(news, checkExitM, ectx, env)
			if isError(obj) {
				continue
			}
			mbs := obj.(*object.StatementObject).Statement.(*parser.MacroBlockStatement)
			// mbs.ReplaceContext(*ectx) // struct(not *struct)
			stmts = append(stmts, mbs)

		case *parser.ReptStatement:
			obj := e.evalReptStatement(news, checkExitM, ectx, env)
			if isError(obj) {
				continue
			}
			stmts = append(stmts, obj.(*object.StatementObject).Statement)

		default:
			stmts = append(stmts, news)
		}
	}

	// ENDM コメント追加
	cs := &parser.CommentStatement{Text: fmt.Sprintf("endm(%s)", mcall.Name), Context: ectx}
	stmts = append(stmts, cs)

	// トップレベルの mbc は Context 置き換えしない
	mbc := &parser.MacroBlockStatement{Label: mcall.Label, Name: mcall.Name, Block: stmts, Context: mcall.Context}
	for _, s := range stmts {
		e.replaceContext(s, ectx)
	}

	return &object.StatementObject{Statement: mbc}
}

// マクロ展開後の MacroBlockStatement 内の Statement.Context を置き換える
func (e *Evaluator) replaceContext(stmt parser.Statement, ectx TContext) {
	if stmt == nil {
		return
	}
	switch stmt := stmt.(type) {
	case *parser.IfStatement:
		stmt.ReplaceContext(*ectx)
		e.replaceContext(stmt.Consequence.(parser.Statement), ectx)
		e.replaceContext(stmt.Alternative.(parser.Statement), ectx)

	case *parser.MacroBlockStatement:
		stmt.ReplaceContext(*ectx)
		for _, s := range stmt.Block {
			e.replaceContext(s, ectx)
		}

	case *parser.BlockStatement:
		for _, s := range stmt.Block {
			e.replaceContext(s, ectx)
		}

	default:
		stmt.ReplaceContext(*ectx)
	}
}

func (e *Evaluator) expandReptBlock(rept *parser.ReptStatement, env TEnv, ectx TContext) object.Object {
	seq := e.Counter()

	args := map[string]parser.Expression{} // 空
	mangleFn := buildMangleNamesFunc(args, seq, "REPT")

	stmts := []parser.Statement{}
	for _, stmt := range rept.Block.Block {
		news := e.mangleNamesInStatement(stmt, mangleFn, ectx)
		switch news := news.(type) {
		case *parser.MacroCallStatement:
			obj := e.evalMacroCallStatement(news, true, ectx, env)
			if isError(obj) {
				return object.ERROR
			}
			sobj, ok := obj.(*object.StatementObject)
			if !ok {
				panic("not *object.StatementObject")
			}

			s := sobj.Statement
			s.ReplaceContext(*ectx)
			stmts = append(stmts, s)

		case *parser.ReptStatement:
			obj := e.evalReptStatement(news, true, ectx, env)
			if isError(obj) {
				return object.ERROR
			}
			rs, ok := obj.(*object.StatementObject)
			if !ok {
				panic("not *object.NodeObject")
			}
			rs.Statement.ReplaceContext(*ectx)
			stmts = append(stmts, rs.Statement)

		default:
			news.ReplaceContext(*ectx)
			stmts = append(stmts, news)
		}
	}
	return &object.StatemetnsObject{Statements: stmts}
}

// 文の中のマクロローカル名をマングリングする
// 引数 ecx は展開後の Context
func (e *Evaluator) mangleNamesInStatement(stmt parser.Statement, replace func(ptr *parser.Expression), ectx TContext) parser.Statement {
	switch stmt := stmt.(type) {
	case *parser.LabelStatement:
		news := *stmt
		replace(&news.Name)
		return &news

	case *parser.ConstStatement:
		news := *stmt
		replace(&news.Name)
		replace(&news.Value)
		return &news

	case *parser.VariableStatement:
		news := *stmt
		replace(&news.Name)
		replace(&news.Value)
		return &news

	case *parser.Z80Instruction:
		news := *stmt
		if news.Label != nil {
			replace(&news.Label)
		}
		replace(&news.Op1)
		replace(&news.Op2)
		return &news

	case *parser.IfStatement:
		news := *stmt
		replace(&news.Condition)
		if news.Consequence != nil {
			news.Consequence = e.mangleNamesInStatement(news.Consequence.(parser.Statement), replace, ectx)
		}
		if news.Alternative != nil {
			news.Alternative = e.mangleNamesInStatement(news.Alternative.(parser.Statement), replace, ectx)
		}
		return &news

	case *parser.BlockStatement:
		news := *stmt
		// 元の文を壊さないよう Clone する
		blk := slices.Clone(stmt.Block)
		news.Block = blk
		for i, s := range news.Block {
			news.Block[i] = e.mangleNamesInStatement(s, replace, ectx)
		}
		return &news

	case *parser.DataStoreStatement:
		news := *stmt
		replace(&news.Label)
		replace(&news.Count)
		replace(&news.FillValue)
		return &news

	case *parser.DataDefineStatement:
		news := *stmt
		replace(&news.Label)
		for i := 0; i < len(news.Values); i++ {
			replace(&news.Values[i])
		}
		return &news

	case *parser.MacroCallStatement:
		news := *stmt
		args := *stmt.Args
		args.Expressions = slices.Clone(args.Expressions)
		for i := range args.Expressions {
			replace(&args.Expressions[i])
		}
		news.Args = &args
		return &news

	case *parser.ReptStatement:
		news := *stmt
		replace(&news.MaxCount)
		news.Block = e.mangleNamesInStatement(news.Block, replace, ectx).(*parser.BlockStatement)
		return &news

	default:
		return stmt
	}
}

// マクロローカルシンボルを置き換える関数を返す
func buildMangleNamesFunc(args map[string]parser.Expression, seq int, macroName string) func(ptr *parser.Expression) {
	// __<seq>_<macroName><local name>
	mangleName := func(seq int, macroName, name string) string { return fmt.Sprintf("__%d_%s%s", seq, macroName, name) }

	var fn func(ptr *parser.Expression)
	fn = func(ptr *parser.Expression) {
		switch expr := (*ptr).(type) {
		case *parser.Label:
			if expr.Name[0] == '@' {
				newLabel := *expr
				newLabel.Name = mangleName(seq, macroName, expr.Name)
				newLabel.LabelType = parser.NODE_LABEL
				*ptr = &newLabel
			} else if arg, ok := args[expr.Name]; ok {
				*ptr = arg
			}

		case *parser.Ident:
			if expr.Name[0] == '@' {
				newIdent := *expr
				newIdent.Name = mangleName(seq, macroName, expr.Name)
				newIdent.IdentType = parser.IDENT
				*ptr = &newIdent
			} else if arg, ok := args[expr.Name]; ok {
				*ptr = arg
			}

		case *parser.IndexedExpression:
			newe := *expr
			fn(&newe.Left)
			fn(&newe.Index)
			*ptr = &newe

		case *parser.InfixExpression:
			newe := *expr
			fn(&newe.Op1)
			fn(&newe.Op2)
			*ptr = &newe

		case *parser.PrefixExpression:
			newe := *expr
			fn(&newe.Op)
			*ptr = &newe

		case *parser.FuncCallExpression:
			newe := *expr
			args := *newe.Args
			args.Expressions = slices.Clone(args.Expressions)
			for i := range len(args.Expressions) {
				fn(&args.Expressions[i])
			}
			newe.Args = &args
			*ptr = &newe

		case *parser.RegIndirectExpression:
			newe := *expr
			fn(&newe.Displacement)
			*ptr = &newe

		case *parser.AddrIndirectExpression:
			newe := *expr
			fn(&newe.Address)
			*ptr = &newe

		case *parser.ArrayLiteral:
			newe := *expr
			items := *newe.Elements
			items.Expressions = slices.Clone(items.Expressions)
			for i := range len(items.Expressions) {
				fn(&items.Expressions[i])
			}
			newe.Elements = &items
			*ptr = &newe

		default:
		}
	}
	return fn
}
