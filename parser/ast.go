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

// Program - これは Node interface を実装していない
type Program struct {
	Statements []Node
}

func (p *Program) String() string {
	var lines []string
	for _, s := range p.Statements {
		lines = append(lines, s.String())
	}
	return strings.Join(lines, "\n")
}

// Expression Statement
type ExpressionStatement struct {
	Value Node
}

func (e *ExpressionStatement) statementNode() {}
func (e *ExpressionStatement) Type() int      { return -1 } // ダミー
func (e *ExpressionStatement) String() string { return e.Value.String() }

// Z80Instruction Statement
type Z80Instruction struct {
	OpCode     int
	Op1        Node
	Op2        Node
	lineNumber int
}

func (z *Z80Instruction) statementNode() {}
func (z *Z80Instruction) Type() int {
	return z.OpCode
}
func (z *Z80Instruction) String() string {
	var out bytes.Buffer

	out.WriteString(z80OpCode2Name(z.OpCode))
	switch {
	case z.Op1 == nil && z.Op2 == nil:
		break
	case z.Op1 != nil && z.Op2 != nil:
		out.WriteString("\t" + opString(z.Op1))
		out.WriteString(", " + opString(z.Op2))
	case z.Op1 != nil:
		out.WriteString("\t" + opString(z.Op1))
	default:
		out.WriteString("\t" + opString(z.Op2))
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
	return z80OpCode2Name(r.TokenType)
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
	return z80OpCode2Name(f.TokenType)
}

// 間接
type IndirectExpression struct {
	Expression Node
}

func (r *IndirectExpression) expressionNode() {}
func (r *IndirectExpression) Type() int {
	return r.Expression.Type()
}
func (r *IndirectExpression) String() string {
	expr := trimParen(r.Expression.String())
	return "(" + expr + ")"
}

// 中置演算子式
type InfixExpression struct {
	OpCode int
	Op1    Node
	Op2    Node
}

func (i *InfixExpression) expressionNode() {}
func (i *InfixExpression) Type() int {
	return i.OpCode
}
func (i *InfixExpression) String() string {
	var op1, op2 string
	if i.Op1 == nil {
		op1 = "<nil>"
	} else {
		op1 = i.Op1.String()
	}
	if i.Op2 == nil {
		op2 = "<nil>"
	} else {
		op2 = i.Op2.String()
	}
	var out bytes.Buffer

	out.WriteString("(" + op1 + " ")
	out.WriteString(tokenLiteral(i.OpCode))
	out.WriteString(" " + op2 + ")")

	return out.String()
}

// 前置演算子式
type PrefixExpression struct {
	OpCode int
	Op     Node
}

func (p *PrefixExpression) expressionNode() {}
func (p *PrefixExpression) Type() int {
	return p.OpCode
}
func (p *PrefixExpression) String() string {
	var op string
	if p.Op == nil {
		op = "<nil>"
	} else {
		op = p.Op.String()
	}

	return "(" + tokenLiteral(p.OpCode) + op + ")"
}
