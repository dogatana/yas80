package evaluator

import (
	"fmt"
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

	replace := replaceNameInMacro(args, seq, mcall.Name)

	for _, stmt := range macro.Body.Block {
		ectx.Offset += 1
		c := *ectx
		news := e.replaceStatement(stmt.(parser.Statement), replace, &c)
		if news.NodeType() == parser.NODE_MACRO_CALL_STMT {
			subcall := news.(*parser.MacroCallStatement)
			sub := e.evalMacroCallStatement(subcall, env)
			if isError(sub) {
				return object.ERROR
			}
			bs := &parser.MacroBlockStatement{Name: subcall.Name, Block: sub.(*object.NodesObject).Nodes, Context: mcall.Context}
			bs.ReplaceContext(*ectx) // struct(not *struct)
			nodes = append(nodes, bs)
		} else {
			news.ReplaceContext(*ectx)
			nodes = append(nodes, news)
		}
	}
	return &object.NodesObject{Nodes: nodes}
}

func (e *Evaluator) expandReptBlock(rept *parser.ReptStatement, env TEnv, ectx TContext) object.Object {
	seq := e.Counter()
	args := map[string]parser.Expression{}
	replace := replaceNameInMacro(args, seq, "REPT")
	nodes := []parser.Node{}
	for _, stmt := range rept.Block.Block {
		ectx.Offset += 1
		c := *ectx
		news := e.replaceStatement(stmt.(parser.Statement), replace, &c)
		switch news := news.(type) {
		case *parser.MacroCallStatement:
			sub := e.evalMacroCallStatement(news, env)
			if isError(sub) {
				return object.ERROR
			}
			bs := &parser.MacroBlockStatement{Name: news.Name, Block: sub.(*object.NodesObject).Nodes, Context: rept.Context}
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

// func (e *Evaluator) expandReptBlock(rept *parser.ReptStatement, count int, env TEnv) object.Object {
// 	seq := e.Counter()
// 	args := map[string]parser.Expression{}
// 	replace := replaceNameInMacro(args, seq, "REPT")
// 	nodes := []parser.Node{}
// 	for _, stmt := range rept.Block.Block {
// 		news := e.replaceStatement(stmt.(parser.Statement), replace)
// 		if news.NodeType() == parser.NODE_MACRO_CALL_STMT {
// 			mcall := news.(*parser.MacroCallStatement)
// 			sub := e.evalMacroCallStatement(mcall, env)
// 			if isError(sub) {
// 				return object.ERROR
// 			}
// 			bs := &parser.MacroBlockStatement{Name: mcall.Name, Block: sub.(*object.NodesObject).Nodes}
// 			nodes = append(nodes, bs)
// 		} else {
// 			nodes = append(nodes, news)
// 		}
// 	}
// 	return &object.NodesObject{Nodes: nodes}
// }

// ectx 展開後 Context
func (e *Evaluator) replaceStatement(stmt parser.Statement, replace func(ptr *parser.Expression), ectx TContext) parser.Statement {
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
			news.Consequence = e.replaceStatement(news.Consequence.(parser.Statement), replace, ectx)
		}
		if news.Alternative != nil {
			news.Alternative = e.replaceStatement(news.Alternative.(parser.Statement), replace, ectx)
		}
		return &news
	case *parser.BlockStatement:
		news := *stmt
		// 元の文を壊さないよう、コピーしたスライスへ置き換え
		blk := make([]parser.Node, len(stmt.Block))
		copy(blk, stmt.Block)
		news.Block = blk
		for i, s := range news.Block {
			ectx.Offset += 1
			news.Block[i] = e.replaceStatement(s.(parser.Statement), replace, ectx)
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
	default:
		return stmt
	}
}

func replaceNameInMacro(args map[string]parser.Expression, seq int, macroName string) func(ptr *parser.Expression) {
	var replacer func(ptr *parser.Expression)
	replacer = func(ptr *parser.Expression) {
		switch expr := (*ptr).(type) {
		case *parser.Label:
			if expr.Name[0] == '@' {
				newLabel := *expr
				newLabel.Name = replacedName(seq, macroName, expr.Name)
				newLabel.LabelType = parser.NODE_LABEL
				*ptr = &newLabel
			} else if arg, ok := args[expr.Name]; ok {
				*ptr = arg
			}

		case *parser.Ident:
			if expr.Name[0] == '@' {
				newIdent := *expr
				newIdent.Name = replacedName(seq, macroName, expr.Name)
				newIdent.IdentType = parser.IDENT
				*ptr = &newIdent
			} else if arg, ok := args[expr.Name]; ok {
				*ptr = arg
			}

		case *parser.InfixExpression:
			newe := *expr
			replacer(&newe.Op1)
			replacer(&newe.Op2)
			*ptr = &newe
		default:
		}
	}
	return replacer
}

// @name => @<seq>_<macro>_name
func replacedName(seq int, macroName, name string) string {
	return fmt.Sprintf("__%d_%s@%s", seq, macroName, string(name[1:]))
}
