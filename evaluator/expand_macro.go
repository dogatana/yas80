package evaluator

import (
	"fmt"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) expandMacro(mcall *parser.MacroCallStatement, macro *object.MacroObject, env object.Environment) object.Object {
	nodes := []parser.Node{}
	seq := e.Counter()

	// 仮引数と引数Node の紐づけ
	args := map[string]parser.Expression{}
	for i, param := range macro.Params {
		args[param] = mcall.Args.Expressions[i]
	}

	replace := replaceNameInMacro(args, seq, mcall.Name)

	for _, stmt := range macro.Body.Block {
		news := e.replaceStatement(stmt.(parser.Statement), replace)
		if news.NodeType() == parser.NODE_MACRO_CALL_STMT {
			mcall := news.(*parser.MacroCallStatement)
			sub := e.evalMacroCallStatement(mcall, env)
			if isError(sub) {
				return object.ERROR
			}
			bs := &parser.MacroBlockStatement{Name: mcall.Name, Block: sub.(*object.NodesObject).Nodes}
			nodes = append(nodes, bs)
		} else {
			nodes = append(nodes, news)
		}
	}
	return &object.NodesObject{Nodes: nodes}
}

func (e *Evaluator) expandReptBlock(rept *parser.ReptStatement, count int, env object.Environment) object.Object {
	seq := e.Counter()
	args := map[string]parser.Expression{}
	replace := replaceNameInMacro(args, seq, "REPT")
	nodes := []parser.Node{}
	for _, stmt := range rept.Block.Block {
		news := e.replaceStatement(stmt.(parser.Statement), replace)
		if news.NodeType() == parser.NODE_MACRO_CALL_STMT {
			mcall := news.(*parser.MacroCallStatement)
			sub := e.evalMacroCallStatement(mcall, env)
			if isError(sub) {
				return object.ERROR
			}
			bs := &parser.MacroBlockStatement{Name: mcall.Name, Block: sub.(*object.NodesObject).Nodes}
			nodes = append(nodes, bs)
		} else {
			nodes = append(nodes, news)
		}
	}
	return &object.NodesObject{Nodes: nodes}
}

func (e *Evaluator) replaceStatement(stmt parser.Statement, replace func(ptr *parser.Expression)) parser.Statement {
	switch stmt := stmt.(type) {
	case *parser.LabelStatement:
		news := *stmt
		var expr parser.Expression = news.Name
		replace(&expr)
		news.Name = expr.(*parser.Label)
		return &news
	case *parser.ConstStatement:
		news := *stmt
		replace(&news.Name)
		replace(&news.Value)
		return &news
	case *parser.Z80Instruction:
		news := *stmt
		if news.Label != nil {
			var expr parser.Expression = news.Label
			replace(&expr)
			news.Label = expr.(*parser.Label)
		}
		replace(&news.Op1)
		replace(&news.Op2)
		return &news
	case *parser.IfStatement:
		news := *stmt
		replace(&news.Condition)
		if news.Consequence != nil {
			e.replaceStatement(news.Consequence.(parser.Statement), replace)
		}
		if news.Alternative != nil {
			e.replaceStatement(news.Alternative.(parser.Statement), replace)
		}
		return &news
	case *parser.BlockStatement:
		news := *stmt
		for i, s := range news.Block {
			news.Block[i] = e.replaceStatement(s.(parser.Statement), replace)
		}
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
