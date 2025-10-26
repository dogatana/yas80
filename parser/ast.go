package parser

import (
	"bytes"
)

// interface
type Node interface {
	Type() int
	String() string
}

// 文
type Statement interface {
	Node
	statementNode()
}

// 式
type Expression interface {
	Node
	expressionNode()
}

// Program
type Program struct {
	Statements []Node
}

func (p *Program) TokenType() int {
	if len(p.Statements) > 0 {
		return p.Statements[0].Type()
	}
	return 0
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String() + "\n")
	}
	return out.String()
}

// Z80Instruction
type Z80Instruction struct {
	OpCode int
	Op1    Expression
	Op2    Expression
	Line   int
}

func (z *Z80Instruction) statementNode() {}
func (z *Z80Instruction) Type() int {
	return z.OpCode
}
func (z *Z80Instruction) String() string {
	var out bytes.Buffer

	out.WriteString(Z80Names(z.OpCode))
	if z.Op1 != nil {
		out.WriteString("\t" + z.Op1.String())
	}
	if z.Op2 != nil {
		out.WriteString(", " + z.Op2.String())
	}

	return out.String()
}
