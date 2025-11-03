package parser

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	NODE_NODE = iota + 1
	// eror
	NODE_ERROR
	// program
	NODE_PROGRAM
	// statement
	NODE_STMT
	NODE_LABEL_STMT
	NODE_EXPR_STMT
	NODE_CONST_STMT
	NODE_VAR_STMT
	NODE_ENUM_STMT
	NODE_ENUM_ELEMENTS_STMT

	// expression
	NODE_EXPR
	NODE_ENUM_ELEMENT
	NODE_NUMBER
	NODE_IDENT
	NODE_DOT_IDENT
	NODE_INDIRECT
	NODE_INFIX_EXPR
	NODE_PREFIX_EXPR
	NODE_LABEL
	NODE_LOCAL_LABEL
)

func NodeTypeNames(t NodeType) string {
	switch t {
	case NODE_NODE:
		return "NODE_NODE"
	case NODE_STMT:
		return "NODE_STMT"
	case NODE_EXPR_STMT:
		return "NODE_EXPR_STMT"
	case NODE_CONST_STMT:
		return "NODE_CONST_STMT"
	case NODE_VAR_STMT:
		return "NODE_VAR_STMT"
	case NODE_EXPR:
		return "NODE_EXPR"
	case NODE_NUMBER:
		return "NODE_NUMBER"
	case NODE_INDIRECT:
		return "NODE_INDIRECT"
	case NODE_INFIX_EXPR:
		return "NODE_INFIX_EXPR"
	case NODE_PREFIX_EXPR:
		return "NODE_PREFIX_EXPR"
	default:
		return yySymNames[yyXLAT[int(t)]]
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

// これ以降実装 (struct)

// Program
type Program struct {
	Statements []Node
}

func (p *Program) NodeType() NodeType       { return NODE_PROGRAM }
func (p *Program) NodeSubType() NodeSubType { return 0 }
func (p *Program) String() string {
	var lines []string
	for _, s := range p.Statements {
		lines = append(lines, s.String())
	}
	return strings.Join(lines, "\n")
}

// Error
type ParseError struct {
	Message string
}

func (e *ParseError) NodeType() NodeType       { return NODE_ERROR }
func (e *ParseError) NodeSubType() NodeSubType { return 0 }
func (e *ParseError) String() string           { return e.Message }

// ラベルは独立した文として生成
type LabelStatement struct {
	Value      Node
	LineNumber int
}

func (l *LabelStatement) statementNode()           {}
func (l *LabelStatement) NodeType() NodeType       { return NODE_LABEL_STMT }
func (l *LabelStatement) NodeSubType() NodeSubType { return 0 }
func (l *LabelStatement) String() string {
	out := l.Value.(*Label).Name
	if out[0] != '.' {
		out += ":"
	}
	return out
}

// 式文 - Expression Statement
type ExpressionStatement struct {
	Value      Node
	LineNumber int
}

func (e *ExpressionStatement) statementNode()           {}
func (e *ExpressionStatement) NodeType() NodeType       { return NODE_EXPR_STMT }
func (e *ExpressionStatement) NodeSubType() NodeSubType { return 0 }
func (e *ExpressionStatement) String() string           { return e.Value.String() }

// enum 定義
type EnumStatement struct {
	Name       string
	Elements   *EnumElements
	LineNumber int
}

func (e *EnumStatement) statementNode()           {}
func (e *EnumStatement) NodeType() NodeType       { return NODE_ENUM_STMT }
func (e *EnumStatement) NodeSubType() NodeSubType { return 0 }
func (e *EnumStatement) String() string {
	var out bytes.Buffer

	out.WriteString(e.Name + " ENUM\n")
	out.WriteString(e.Elements.String() + "\n")
	out.WriteString("END_ENUM")

	return out.String()
}

// enum 要素定義文
type EnumElements struct {
	Elements []*EnumElement
}

func (e *EnumElements) statementNode()        {}
func (e *EnumElements) NodeType() NodeType    { return NODE_ENUM_ELEMENTS_STMT }
func (e *EnumElements) NodeSubType() NodeType { return 0 }
func (e *EnumElements) String() string {
	stmts := []string{}
	for _, e := range e.Elements {
		stmts = append(stmts, e.String())
	}
	return strings.Join(stmts, "\n")
}

// enum 要素
type EnumElement struct {
	Name  string
	Value Node
}

func (e *EnumElement) expressionNode()       {}
func (e *EnumElement) NodeType() NodeType    { return NODE_ENUM_ELEMENT }
func (e *EnumElement) NodeSubType() NodeType { return 0 }
func (e *EnumElement) String() string {
	if e.Value == nil {
		return e.Name
	} else {
		return e.Name + " = " + e.Value.String()
	}
}

// 定数定義文 - CONST, EQU Statement
type ConstStatement struct {
	Name       *Ident
	Value      Node
	LineNumber int
}

func (c *ConstStatement) statementNode()           {}
func (c *ConstStatement) NodeType() NodeType       { return NODE_CONST_STMT }
func (c *ConstStatement) NodeSubType() NodeSubType { return 0 }
func (c *ConstStatement) String() string {
	return "CONST " + c.Name.String() + " = " + c.Value.String()
}

// Z80 命令文 - Z80Instruction Statement
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

	out.WriteString(Z80OpCode2Name(z.OpCode))
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

// これ以降は 式 (Exspression)

// ラベル
type Label struct {
	nodeType NodeType
	Name     string
}

func (l *Label) expressionNode()          {}
func (l *Label) NodeType() NodeType       { return l.nodeType }
func (l *Label) NodeSubType() NodeSubType { return 0 }
func (l *Label) String() string           { return l.Name }

// 数値
type NumberLiteral struct {
	Value int
}

func (n *NumberLiteral) expressionNode()          {}
func (n *NumberLiteral) NodeType() NodeType       { return NODE_NUMBER }
func (n *NumberLiteral) NodeSubType() NodeSubType { return 0 }
func (n *NumberLiteral) String() string {
	return fmt.Sprintf("%d", n.Value)
}

// レジスタ
type RegisterLiteral struct {
	RegisterType int
	Register     int
}

func (r *RegisterLiteral) expressionNode()          {}
func (r *RegisterLiteral) NodeType() NodeType       { return NodeType(r.RegisterType) }
func (r *RegisterLiteral) NodeSubType() NodeSubType { return NodeSubType(r.Register) }
func (r *RegisterLiteral) String() string {
	return Z80OpCode2Name(r.Register)
}

// フラグ
type FlagLiteral struct {
	Flag int
}

func (f *FlagLiteral) expressionNode()          {}
func (f *FlagLiteral) NodeType() NodeType       { return Z80_FLAG }
func (f *FlagLiteral) NodeSubType() NodeSubType { return NodeSubType(f.Flag) }
func (f *FlagLiteral) String() string {
	return Z80OpCode2Name(f.Flag)
}

// 識別子
type Ident struct {
	Name  string
	Value Node
}

func (i *Ident) expressionNode()          {}
func (i *Ident) NodeType() NodeType       { return NODE_IDENT }
func (i *Ident) NodeSubType() NodeSubType { return 0 }
func (i *Ident) String() string           { return i.Name }

// ドット識別子
type DotIdent struct {
	Left  string
	Right string
	Value Node
}

func (di *DotIdent) expressionNode()          {}
func (di *DotIdent) NodeType() NodeType       { return NODE_DOT_IDENT }
func (di *DotIdent) NodeSubType() NodeSubType { return 0 }
func (di *DotIdent) String() string           { return di.Left + "." + di.Right }

// 間接指定
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
