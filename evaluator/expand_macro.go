package evaluator

import (
	"fmt"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) expandMacro(mcall *parser.MacroCallStatement, macro *object.MacroObject, env object.Environment) []parser.Node {
	nodes := []parser.Node{}
	seq := e.Counter()

	// 仮引数と引数Node の紐づけ
	args := map[string]parser.Expression{}
	for i, param := range macro.Params {
		args[param] = mcall.Args.Expressions[i]
	}

	replace := replaceNameInMacro(args, seq, mcall.Name)

	for _, stmt := range macro.Body.Block {
		switch stmt := stmt.(type) {
		case *parser.LabelStatement:
			news := *stmt
			var expr parser.Expression = news.Name
			replace(&expr)
			news.Name = expr.(*parser.Label)
			nodes = append(nodes, &news)
		case *parser.ConstStatement:
			news := *stmt
			replace(&news.Name)
			replace(&news.Value)
			nodes = append(nodes, &news)
		case *parser.Z80Instruction:
			news := *stmt
			if news.Label != nil {
				var expr parser.Expression = news.Label
				replace(&expr)
				news.Label = expr.(*parser.Label)
			}
			replace(&news.Op1)
			replace(&news.Op2)
			fmt.Println(news.String())
			nodes = append(nodes, &news)
		default:
			nodes = append(nodes, stmt)
		}
	}
	return nodes
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
	return fmt.Sprintf("__%d_%s_%s", seq, macroName, string(name[1:]))
}
