package parser

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	NODE_NODE = iota + 1
	NODE_STMT
	NODE_EXPR_STMT
	NODE_EXPR
	NODE_NUMBER
	NODE_INDIRECT
	NODE_INFIX_EXPR
	NODE_PREFIX_EXPR
)

func NodeTypeNames(t NodeType) string {
	switch t {
	case NODE_NODE:
		{
			return "NODE_NODE"
		}
	case NODE_STMT:
		{
			return "NODE_STMT"
		}
	case NODE_EXPR_STMT:
		{
			return "NODE_EXPR_STMT"
		}
	case NODE_EXPR:
		{
			return "NODE_EXPR"
		}
	case NODE_NUMBER:
		{
			return "NODE_NUMBER"
		}
	case NODE_INDIRECT:
		{
			return "NODE_INDIRECT"
		}
	case NODE_INFIX_EXPR:
		{
			return "NODE_INFIX_EXPR"
		}
	case NODE_PREFIX_EXPR:
		{
			return "NODE_PREFIX_EXPR"
		}
	default:
		{
			return yySymNames[yyXLAT[int(t)]]
		}
	}
}

type NodeType int
type NodeSubType int

// interface
type Node interface {
	NodeType() NodeType
	NodeSubType() NodeSubType
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

func (e *ExpressionStatement) statementNode()           {}
func (e *ExpressionStatement) NodeType() NodeType       { return NODE_EXPR_STMT }
func (e *ExpressionStatement) NodeSubType() NodeSubType { return 0 }
func (e *ExpressionStatement) String() string           { return e.Value.String() }

// Z80Instruction Statement
type Z80Instruction struct {
	InstType   int
	OpCode     int
	Op1        Node
	Op2        Node
	LineNumber int
}

func (z *Z80Instruction) statementNode() {}
func (z *Z80Instruction) NodeType() NodeType {
	return NodeType(z.InstType)
}
func (z *Z80Instruction) NodeSubType() NodeSubType {
	return NodeSubType(z.OpCode)
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
	Value int
}

func (n *NumberLiteral) expressionNode()          {}
func (n *NumberLiteral) NodeType() NodeType       { return NUMBER }
func (n *NumberLiteral) NodeSubType() NodeSubType { return 0 }
func (n *NumberLiteral) String() string {
	return fmt.Sprintf("%d", n.Value)
}

// レジスタリテラル
type RegisterLiteral struct {
	RegisterType int
	Register     int
}

func (r *RegisterLiteral) expressionNode()          {}
func (r *RegisterLiteral) NodeType() NodeType       { return NodeType(r.RegisterType) }
func (r *RegisterLiteral) NodeSubType() NodeSubType { return NodeSubType(r.Register) }
func (r *RegisterLiteral) String() string {
	return z80OpCode2Name(r.Register)
}

// フラグリテラル
type FlagLiteral struct {
	Flag int
}

func (f *FlagLiteral) expressionNode()          {}
func (f *FlagLiteral) NodeType() NodeType       { return Z80_FLAG }
func (f *FlagLiteral) NodeSubType() NodeSubType { return NodeSubType(f.Flag) }
func (f *FlagLiteral) String() string {
	return z80OpCode2Name(f.Flag)
}

// 間接
type IndirectExpression struct {
	Expression Node
}

func (r *IndirectExpression) expressionNode()          {}
func (r *IndirectExpression) NodeType() NodeType       { return NODE_INDIRECT }
func (r *IndirectExpression) NodeSubType() NodeSubType { return 0 }
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

func (i *InfixExpression) expressionNode()          {}
func (i *InfixExpression) NodeType() NodeType       { return NODE_INFIX_EXPR }
func (i *InfixExpression) NodeSubType() NodeSubType { return NodeSubType(i.OpCode) }
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

func (p *PrefixExpression) expressionNode()          {}
func (p *PrefixExpression) NodeType() NodeType       { return NodeType(NODE_PREFIX_EXPR) }
func (p *PrefixExpression) NodeSubType() NodeSubType { return NodeSubType(p.OpCode) }
func (p *PrefixExpression) String() string {
	var op string
	if p.Op == nil {
		op = "<nil>"
	} else {
		op = p.Op.String()
	}

	return "(" + tokenLiteral(p.OpCode) + op + ")"
}
