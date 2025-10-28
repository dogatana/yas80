package generator

import (
	"fmt"
	"yas80/errorstore"
	"yas80/object"
	"yas80/parser"
)

type Generator struct {
	program *parser.Program
	es      *errorstore.ErrorStore
	objects []object.Object
}

func New(p *parser.Program, es *errorstore.ErrorStore) *Generator {
	return &Generator{program: p, es: es}
}

func (g *Generator) Generate() {
	for _, node := range g.program.Statements {
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
