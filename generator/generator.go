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
	env     *object.Environment
	objects []object.Object
}

func New(p *parser.Program, env *object.Environment, es *errorstore.ErrorStore) *Generator {
	return &Generator{program: p, env: env, es: es}
}

func (g *Generator) Generate() {
	for _, node := range g.program.Statements {
		switch node.NodeType() {
		case parser.Z80_INST0:
			node := node.(*parser.Z80Instruction)
			info := Z80CodeTable0[int(node.OpCode)]
			obj := &object.FixedCode{Line: node.LineNumber, Code: make([]byte, len(info.Bytes))}
			copy(obj.Code, info.Bytes)
			g.objects = append(g.objects, obj)
		case parser.Z80_INST1:
			node := node.(*parser.Z80Instruction)
			if node.OpCode == parser.Z80_INST_RET {
				g.generateRET(node)
				break
			}
			fmt.Println(parser.NodeTypeNames(node.NodeType()))
		default:
			fmt.Println(parser.NodeTypeNames(node.NodeType()))
		}
	}
}

func (g *Generator) generateRET(node *parser.Z80Instruction) {
	if node.Op1 == nil {
		g.objects = append(g.objects,
			&object.FixedCode{Line: node.LineNumber, Code: []byte{0xc9}})
		return
	}
	if node.Op1.NodeType() == parser.Z80_FLAG {
		flag := int(node.Op1.NodeSubType()) - parser.Z80_FLAG_NZ
		b := byte(0xc0 | flag<<3)
		g.objects = append(g.objects,
			&object.FixedCode{Line: node.LineNumber, Code: []byte{b}})
		return
	}
	g.es.AddError("", node.LineNumber,
		fmt.Sprintf("第1オペランドがフラグではありません '%s'", node.Op1.String()))
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
