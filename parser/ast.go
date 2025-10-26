package parser

import (
	"bytes"
	"fmt"
	"strings"
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
	var lines []string
	for _, s := range p.Statements {
		lines = append(lines, s.String())
	}
	return strings.Join(lines, "\n")

}

// Z80Instruction
type Z80Instruction struct {
	OpCode int
	Op1    Node
	Op2    Node
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

// 数値リテラル
type NumberLiteral struct {
	TokenType int
	Value     int
}

func (n *NumberLiteral) expressionNode() {}
func (n *NumberLiteral) Type() int {
	return n.TokenType
}
func (n *NumberLiteral) String() string {
	return fmt.Sprintf("%d", n.Value)
}

// レジスタリテラル
type RegisterLiteral struct {
	TokenType int
}

func (r *RegisterLiteral) expressionNode() {}
func (r *RegisterLiteral) Type() int {
	return r.TokenType
}
func (r *RegisterLiteral) String() string {
	return Z80Names(r.TokenType)
}

// フラグリテラル
type FlagLiteral struct {
	TokenType int
}

func (f *FlagLiteral) expressionNode() {}
func (f *FlagLiteral) Type() int {
	return f.TokenType
}
func (f *FlagLiteral) String() string {
	return Z80Names(f.TokenType)
}
