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

func (e *Evaluator) replaceAtIdent(node parser.Node, name string, seq int) parser.Node {
	if node == nil {
		return node
	}

	switch node := node.(type) {
	case *parser.LabelStatement:
		if !needReplace(node.Name.Name) {
			return node
		}
		label := *node.Name
		label.Name = replacedName(label.Name, name, seq)
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
		ident.Name = replacedName(ident.Name, name, seq)
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
		ident.Name = replacedName(ident.Name, name, seq)
		new := *node
		new.Name = &ident
		return &new

	case *parser.IfStatement:
		cond := e.replaceAtIdent(node.Condition, name, seq)
		conseq := e.replaceAtIdent(node.Consequence, name, seq)
		alt := e.replaceAtIdent(node.Alternative, name, seq)

		new := *node
		new.Condition = cond.(parser.Expression)
		new.Consequence = conseq
		new.Alternative = alt
		return &new

	case *parser.BlockStatement:
		nodes := []parser.Node{}
		for _, stmt := range node.Block {
			nodes = append(nodes, e.replaceAtIdent(stmt, name, seq))
		}
		return &parser.BlockStatement{Block: nodes}

	case *parser.Z80Instruction:
		new := *node
		if node.Op1 != nil {
			op1 := e.replaceAtIdent(node.Op1, name, seq)
			new.Op1 = op1.(parser.Expression)

		}
		if node.Op2 != nil {
			op2 := e.replaceAtIdent(node.Op2, name, seq)
			new.Op2 = op2.(parser.Expression)

		}
		return &new

	case *parser.Ident:
		if !needReplace(node.Name) {
			return node
		}
		new := *node
		new.Name = replacedName(node.Name, name, seq)
		return &new

	default:
		return node
	}
}

func needReplace(name string) bool {
	return name[0] == '@' && !strings.HasPrefix(name, "@@")
}

// @name => @<seq>_name
func replacedName(name, macName string, seq int) string {
	return fmt.Sprintf("__%s_%d_%s", macName, seq, string(name[1:]))
}
