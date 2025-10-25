package ast

import (
	"bytes"
)

// interface
type Node interface {
	TokenType() int
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
	ExpressionNode()
}

// Program
type Program struct {
	Statements []Statement
}

func (p *Program) TokenType() int {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenType()
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
	Statement
	OpCode int
	Op1    Expression
	Op2    Expression
	Line   int
}

func (z *Z80Instruction) statementNode() {}
func (z *Z80Instruction) TokenType() int {
	return z.OpCode
}
func (z *Z80Instruction) Strig() string {
	var out bytes.Buffer

	// out.WriteString(parser.Z80Names(z.OpCode) + "\t")
	out.WriteString(z.Op1.String() + ", ")
	out.WriteString(z.Op2.String())

	return out.String()
}
