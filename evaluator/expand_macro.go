package evaluator

import (
	"fmt"
	"strings"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) expandMacro(macro *object.MacroObject) []parser.Node {
	nodes := []parser.Node{}
	seq := e.Counter()

	for _, stmt := range macro.Body.Block {
		nodes = append(nodes, e.replaceAtIdent(stmt, macro.Name, seq))
	}

	fmt.Println("expanded")
	for i, n := range nodes {
		fmt.Printf("[%d] %s\n", i, n.String())
	}
	return nodes
}

func (e *Evaluator) replaceAtIdent(node parser.Node, macroName string, seq int) parser.Node {
	if node == nil {
		return node
	}

	switch node := node.(type) {
	case *parser.LabelStatement:
		if !needReplace(node.Name.Name) {
			return node
		}
		label := *node.Name
		label.Name = replacedName(seq, macroName, label.Name)
		label.LabelType = parser.NODE_LABEL

		new := *node
		new.Name = &label
		fmt.Printf("new %#v(%s)\n", new, new.String())
		return &new

	case *parser.ConstStatement:
		if !needReplace(node.Name.Name) {
			return node
		}
		ident := *node.Name
		ident.Name = replacedName(seq, macroName, ident.Name)
		ident.IdentType = parser.IDENT

		new := *node
		new.Name = &ident
		fmt.Printf("new %#v(%s)\n", new, new.String())
		return &new

	case *parser.VariableStatement:
		if !needReplace(node.Name.Name) {
			return node
		}
		ident := *node.Name
		ident.Name = replacedName(seq, macroName, ident.Name)
		new := *node
		new.Name = &ident
		return &new

	case *parser.IfStatement:
		cond := e.replaceAtIdent(node.Condition, macroName, seq)
		conseq := e.replaceAtIdent(node.Consequence, macroName, seq)
		alt := e.replaceAtIdent(node.Alternative, macroName, seq)

		new := *node
		new.Condition = cond.(parser.Expression)
		new.Consequence = conseq
		new.Alternative = alt
		return &new

	case *parser.BlockStatement:
		nodes := []parser.Node{}
		for _, stmt := range node.Block {
			nodes = append(nodes, e.replaceAtIdent(stmt, macroName, seq))
		}
		return &parser.BlockStatement{Block: nodes}

	case *parser.Z80Instruction:
		new := *node
		if node.Op1 != nil {
			op1 := e.replaceAtIdent(node.Op1, macroName, seq)
			new.Op1 = op1.(parser.Expression)

		}
		if node.Op2 != nil {
			op2 := e.replaceAtIdent(node.Op2, macroName, seq)
			new.Op2 = op2.(parser.Expression)

		}
		return &new

	case *parser.Ident:
		if !needReplace(node.Name) {
			return node
		}
		new := *node
		new.Name = replacedName(seq, macroName, node.Name)
		return &new

	default:
		return node
	}
}

func needReplace(name string) bool {
	return name[0] == '@' && !strings.HasPrefix(name, "@@")
}

// @name => @<seq>_<macro>_name
func replacedName(seq int, macroName, name string) string {
	return fmt.Sprintf("__%d_%s_%s", seq, macroName, string(name[1:]))
}
