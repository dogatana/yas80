package evaluator

import (
	"fmt"
	"slices"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) expandMacro(mcall *parser.MacroCallStatement, macro *object.MacroObject, env TEnv, ectx TContext) object.Object {
	nodes := []parser.Node{}
	seq := e.Counter()

	// 仮引数と引数Node の紐づけ
	args := map[string]parser.Expression{}
	for i, param := range macro.Params {
		args[param] = mcall.Args.Expressions[i]
	}

	mfn := buildMangleNamesFunc(args, seq, mcall.Name)

	for _, stmt := range macro.Body.Block {
		ectx.Offset++
		// 引数のContextの内容を壊さないよう Clone してから使用する
		c := *ectx
		news := e.mangleNamesInStatement(stmt.(parser.Statement), mfn, &c)
		news.ReplaceContext(*ectx)

		switch news := news.(type) {
		case *parser.MacroCallStatement:
			obj := e.evalMacroCallStatement(news, env)
			if isError(obj) {
				continue
			}
			bs := &parser.MacroBlockStatement{Name: news.Name, Block: obj.(*object.NodesObject).Nodes, Context: news.Context}
			bs.ReplaceContext(*ectx) // struct(not *struct)
			nodes = append(nodes, bs)
			if len(bs.Block) > 0 {
				ectx.Offset = bs.Block[len(bs.Block)-1].(parser.Statement).GetContext().Offset
			}

		case *parser.ReptStatement:
			obj := e.evalReptStatement(news, env)
			nodes = append(nodes, obj.(*object.NodeObject).Node)

		default:
			nodes = append(nodes, news)
		}
	}
	return &object.NodesObject{Nodes: nodes}
}

func (e *Evaluator) expandReptBlock(rept *parser.ReptStatement, env TEnv, ectx TContext) object.Object {
	seq := e.Counter()

	args := map[string]parser.Expression{} // 空
	mangleFn := buildMangleNamesFunc(args, seq, "REPT")

	nodes := []parser.Node{}
	for _, stmt := range rept.Block.Block {
		ectx.Offset += 1
		c := *ectx
		news := e.mangleNamesInStatement(stmt.(parser.Statement), mangleFn, &c)
		switch news := news.(type) {
		case *parser.MacroCallStatement:
			obj := e.evalMacroCallStatement(news, env)
			if isError(obj) {
				return object.ERROR
			}
			bs := &parser.MacroBlockStatement{Name: news.Name, Block: obj.(*object.NodesObject).Nodes, Context: rept.Context}
			bs.ReplaceContext(*ectx)
			nodes = append(nodes, bs)

		case *parser.ReptStatement:
			obj := e.evalReptStatement(news, env)
			if isError(obj) {
				return object.ERROR
			}
			rs, ok := obj.(*object.NodeObject)
			if !ok {
				panic("not *object.NodeObject")
			}
			rs.Node.(parser.Statement).ReplaceContext(*ectx)
			nodes = append(nodes, rs.Node)

		default:
			news.ReplaceContext(*ectx)
			nodes = append(nodes, news)
		}
	}
	return &object.NodesObject{Nodes: nodes}
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
			ectx.Offset += 1
			news.Block[i] = e.mangleNamesInStatement(s.(parser.Statement), replace, ectx)
		}
		return &news

	case *parser.DataStoreStatement:
		news := *stmt
		replace(&news.Label)
		replace(&news.Count)
		replace(&news.FillValue)
		return &news

	case *parser.DataStatement:
		news := *stmt
		replace(&news.Label)
		for i := 0; i < len(news.Values); i++ {
			replace(&news.Values[i])
		}
		return &news

	case *parser.MacroCallStatement:
		news := *stmt
		args := *stmt.Args
		args.Expressions = make([]parser.Expression, len(args.Expressions))
		copy(args.Expressions, stmt.Args.Expressions)
		for i := range args.Expressions {
			replace(&args.Expressions[i])
		}
		news.Args = &args
		return &news

	case *parser.ReptStatement:
		news := *stmt
		replace(&news.MaxCount)
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

		case *parser.InfixExpression:
			newe := *expr
			fn(&newe.Op1)
			fn(&newe.Op2)
			*ptr = &newe

		case *parser.PrefixExpression:
			newe := *expr
			fn(&newe.Op)
			*ptr = &newe

		default:
		}
	}
	return fn
}
