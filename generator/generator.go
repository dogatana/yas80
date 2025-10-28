package generator

import (
	"fmt"
	"yas80/object"
	"yas80/parser"
)

type Generator struct {
	objects []object.Object
}

func New() *Generator {
	return &Generator{}
}

func (g *Generator) Generate(prg *parser.Program) {
	for _, node := range prg.Statements {
		switch node.NodeType() {
		case parser.Z80_INST0:
			inst := node.(*parser.Z80Instruction)
			info := Z80CodeTable0[int(inst.OpCode)]
			obj := &object.FixedCode{Line: inst.LineNumber, Code: make([]byte, len(info.Bytes))}
			copy(obj.Code, info.Bytes)
			g.objects = append(g.objects, obj)
		default:
			fmt.Println(parser.NodeTypeNames(node.NodeType()))
		}
	}
}

func (g *Generator) Dump() {
	for _, o := range g.objects {
		fmt.Println(o.String())
	}
}
func (g *Generator) MergeCode() []byte {
	var out []byte

	for _, o := range g.objects {
		code, ok := o.(*object.FixedCode)
		if !ok {
			return []byte{}
		}
		out = append(out, code.Code...)
	}
	return out
}
