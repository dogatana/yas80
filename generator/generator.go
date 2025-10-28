package generator

import (
	"fmt"
	"yas80/object"
	"yas80/parser"
)

type Generator struct {
	objects []object.Object
}

func (g *Generator) Generate(prg *parser.Program) {
	for _, node := range prg.Statements {
		fmt.Println(node.NodeType(), parser.NodeTypeNames(node.NodeType()))
	}
}
