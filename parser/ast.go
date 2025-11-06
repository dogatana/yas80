package parser

import (
	"bytes"
	"fmt"
	"strings"
)

const (
	NODE_NODE = iota + 1

	// program
	NODE_PROGRAM

	// eror
	NODE_ERROR

	// statement
	NODE_STMT
	NODE_LABEL_STMT
	NODE_EXPR_STMT
	NODE_CONST_STMT
	NODE_VAR_STMT
	NODE_ASIGN_STMT
	NODE_ENUM_STMT
	NODE_ENUM_ELEMENTS_STMT
	NODE_REPEAT_STMT
	NODE_IF_STMT
	NODE_BLOCK_STMT
	NODE_FUNCTION_STMT

	// expression
	NODE_EXPR
	NODE_ENUM_ELEMENT
	NODE_NUMBER
	NODE_IDENT
	NODE_DOT_IDENT
	NODE_ARRAY
	NODE_INDEXED_EXPR
	NODE_INFIX_EXPR
	NODE_PREFIX_EXPR
	NODE_CALL
	NODE_EXPR_LIST
	NODE_LABEL
	NODE_LOCAL_LABEL
	NODE_INDIRECT // for Z80
)

// func NodeTypeNames(t NodeType) string {
// 	switch t {
// 	case NODE_NODE:
// 		return "NODE_NODE"
// 	case NODE_STMT:
// 		return "NODE_STMT"
// 	case NODE_EXPR_STMT:
// 		return "NODE_EXPR_STMT"
// 	case NODE_CONST_STMT:
// 		return "NODE_CONST_STMT"
// 	case NODE_VAR_STMT:
// 		return "NODE_VAR_STMT"
// 	case NODE_EXPR:
// 		return "NODE_EXPR"
// 	case NODE_NUMBER:
// 		return "NODE_NUMBER"
// 	case NODE_INDIRECT:
// 		return "NODE_INDIRECT"
// 	case NODE_INFIX_EXPR:
// 		return "NODE_INFIX_EXPR"
// 	case NODE_PREFIX_EXPR:
// 		return "NODE_PREFIX_EXPR"
// 	default:
// 		return yySymNames[yyXLAT[int(t)]]
// 	}
// }

type NodeType int
type NodeSubType int

// interface

// Node
type Node interface {
	NodeType() NodeType
	NodeSubType() NodeSubType
	String() string
}

// 文
type Statement interface {
	Node
	statementNode()
	LineNumber() int // エラー表示用
}

// 式
type Expression interface {
	Node
	expressionNode()
}

// 実装 (struct)

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

// Error(Expression, Statement)
type ParseError struct {
	Message    string
	lineNumber int
}

func (pe *ParseError) statementNode()           {}
func (pe *ParseError) expressionNode()          {}
func (pe *ParseError) NodeType() NodeType       { return NODE_ERROR }
func (pe *ParseError) NodeSubType() NodeSubType { return 0 }
func (pe *ParseError) LineNumber() int          { return pe.lineNumber }
func (pe *ParseError) String() string {
	return fmt.Sprintf("%s %d", pe.Message, pe.lineNumber)
}

// ラベル - 独立した文として生成
type LabelStatement struct {
	Value      Node
	lineNumber int
}

func (ls *LabelStatement) statementNode()           {}
func (ls *LabelStatement) NodeType() NodeType       { return NODE_LABEL_STMT }
func (ls *LabelStatement) NodeSubType() NodeSubType { return 0 }
func (ls *LabelStatement) LineNumber() int          { return ls.lineNumber }
func (ls *LabelStatement) String() string {
	out := ls.Value.(*Label).Name
	if out[0] != '.' {
		out += ":"
	}
	return out
}

// 式文 - Expression Statement
type ExpressionStatement struct {
	Value      Node
	lineNumber int
}

func (es *ExpressionStatement) statementNode()           {}
func (es *ExpressionStatement) NodeType() NodeType       { return NODE_EXPR_STMT }
func (es *ExpressionStatement) NodeSubType() NodeSubType { return 0 }
func (es *ExpressionStatement) LineNumber() int          { return es.lineNumber }
func (es *ExpressionStatement) String() string           { return es.Value.String() }

// enum 定義文
type EnumStatement struct {
	Name       string
	Elements   *EnumElements
	lineNumber int
}

func (es *EnumStatement) statementNode()           {}
func (es *EnumStatement) NodeType() NodeType       { return NODE_ENUM_STMT }
func (es *EnumStatement) NodeSubType() NodeSubType { return 0 }
func (es *EnumStatement) LineNumber() int          { return es.lineNumber }
func (es *EnumStatement) String() string {
	var out bytes.Buffer

	out.WriteString(es.Name + " ENUM\n")
	out.WriteString(es.Elements.String() + "\n")
	out.WriteString("END_ENUM")

	return out.String()
}

// enum 要素定義文
type EnumElements struct {
	Elements   []*EnumElement
	lineNumber int
}

func (ee *EnumElements) statementNode()           {}
func (ee *EnumElements) NodeType() NodeType       { return NODE_ENUM_ELEMENTS_STMT }
func (ee *EnumElements) NodeSubType() NodeSubType { return 0 }
func (ee *EnumElements) LineNumber() int          { return ee.lineNumber }
func (ee *EnumElements) String() string {
	stmts := []string{}
	for _, e := range ee.Elements {
		stmts = append(stmts, e.String())
	}
	return strings.Join(stmts, "\n")
}

// enum 要素
type EnumElement struct {
	Name  string
	Value Node
}

func (ee *EnumElement) expressionNode()          {}
func (ee *EnumElement) NodeType() NodeType       { return NODE_ENUM_ELEMENT }
func (ee *EnumElement) NodeSubType() NodeSubType { return 0 }
func (ee *EnumElement) String() string {
	if ee.Value == nil {
		return ee.Name
	} else {
		return ee.Name + " = " + ee.Value.String()
	}
}

// repeat statment
type RepeatStatement struct {
	MaxCount   Node
	Block      Node
	lineNumber int
}

func (rs *RepeatStatement) statementNode()           {}
func (rs *RepeatStatement) NodeType() NodeType       { return NODE_REPEAT_STMT }
func (rs *RepeatStatement) NodeSubType() NodeSubType { return 0 }
func (rs *RepeatStatement) LineNumber() int          { return rs.lineNumber }
func (rs *RepeatStatement) String() string {
	var out bytes.Buffer

	out.WriteString("REPEAT ")
	out.WriteString(rs.MaxCount.String() + "\n")
	block := rs.Block.String()
	if block != "" {
		out.WriteString(block + "\n")
	}
	out.WriteString("END_REPEAT")

	return out.String()
}

// if statement
type IfStatement struct {
	Condition   Node
	Consequence Node
	Alternative Node
	lineNumber  int
}

func (is *IfStatement) statementNode()           {}
func (is *IfStatement) NodeType() NodeType       { return NODE_IF_STMT }
func (is *IfStatement) NodeSubType() NodeSubType { return 0 }
func (is *IfStatement) LineNumber() int          { return is.lineNumber }
func (is *IfStatement) String() string {
	var out bytes.Buffer

	out.WriteString("IF " + is.Condition.String() + "\n")
	block := is.Consequence.String()
	if block != "" {
		out.WriteString(block + "\n")
	}
	if is.Alternative != nil {
		out.WriteString("ELSE\n")
		block = is.Alternative.String()
		if block != "" {
			out.WriteString(block + "\n")
		}
	}
	out.WriteString("END_IF")

	return out.String()
}

// function 文
type FunctionStatement struct {
	Name       string
	Params     []string
	Block      *BlockStatement
	lineNumber int
}

func (fs *FunctionStatement) statementNode()           {}
func (fs *FunctionStatement) NodeType() NodeType       { return NODE_FUNCTION_STMT }
func (fs *FunctionStatement) NodeSubType() NodeSubType { return 0 }
func (fs *FunctionStatement) LineNumber() int          { return fs.lineNumber }
func (fs *FunctionStatement) String() string {
	var out bytes.Buffer

	out.WriteString(fs.Name + " FUNCTION " + strings.Join(fs.Params, ", ") + "\n")
	out.WriteString(fs.Block.String() + "\n")
	out.WriteString("END_FUNCTION")

	return out.String()
}

// block statement
type BlockStatement struct {
	Block      []Node
	lineNumber int
}

func (bs *BlockStatement) statementNode()           {}
func (bs *BlockStatement) NodeType() NodeType       { return NODE_BLOCK_STMT }
func (bs *BlockStatement) NodeSubType() NodeSubType { return 0 }
func (bs *BlockStatement) LineNumber() int          { return bs.lineNumber }
func (bs *BlockStatement) String() string {
	stmts := []string{}

	for _, s := range bs.Block {
		stmts = append(stmts, s.String())
	}
	return strings.Join(stmts, "\n")
}

// 定数定義文 - CONST, EQU Statement
type ConstStatement struct {
	Name       *Ident
	Value      Node
	lineNumber int
}

func (cs *ConstStatement) statementNode()           {}
func (cs *ConstStatement) NodeType() NodeType       { return NODE_CONST_STMT }
func (cs *ConstStatement) NodeSubType() NodeSubType { return 0 }
func (cs *ConstStatement) LineNumber() int          { return cs.lineNumber }
func (cs *ConstStatement) String() string {
	var out bytes.Buffer

	out.WriteString("CONST ")
	out.WriteString(cs.Name.Name)
	out.WriteString(" = ")
	out.WriteString(cs.Value.String())

	return out.String()
}

// 変数定義文 - VAR
type VariableStatement struct {
	Name       *Ident
	Value      Node
	lineNumber int
}

func (vs *VariableStatement) statementNode()           {}
func (vs *VariableStatement) NodeType() NodeType       { return NODE_VAR_STMT }
func (vs *VariableStatement) NodeSubType() NodeSubType { return 0 }
func (vs *VariableStatement) LineNumber() int          { return vs.lineNumber }
func (vs *VariableStatement) String() string {
	var out bytes.Buffer

	out.WriteString("VAR ")
	out.WriteString(vs.Name.Name)
	out.WriteString(" = ")
	out.WriteString(vs.Value.String())

	return out.String()
}

// 変数代入文
type AsignStatement struct {
	Left       Expression
	Value      Node
	lineNumber int
}

func (as *AsignStatement) statementNode()           {}
func (as *AsignStatement) NodeType() NodeType       { return NODE_ASIGN_STMT }
func (as *AsignStatement) NodeSubType() NodeSubType { return 0 }
func (as *AsignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.Left.String())
	out.WriteString(" = ")
	out.WriteString(as.Value.String())

	return out.String()
}

// Z80 命令文 - Z80Instruction Statement
type Z80Instruction struct {
	InstType   int
	OpCode     int
	Op1        Node
	Op2        Node
	lineNumber int
}

func (zi *Z80Instruction) statementNode() {}
func (zi *Z80Instruction) NodeType() NodeType {
	return NodeType(zi.InstType)
}
func (zi *Z80Instruction) NodeSubType() NodeSubType {
	return NodeSubType(zi.OpCode)
}
func (zi *Z80Instruction) LineNumber() int { return zi.lineNumber }
func (zi *Z80Instruction) String() string {
	var out bytes.Buffer

	out.WriteString(Z80OpCode2Name(zi.OpCode))
	switch {
	case zi.Op1 == nil && zi.Op2 == nil:
		break
	case zi.Op1 != nil && zi.Op2 != nil:
		out.WriteString("\t" + opString(zi.Op1))
		out.WriteString(", " + opString(zi.Op2))
	case zi.Op1 != nil:
		out.WriteString("\t" + opString(zi.Op1))
	default:
		out.WriteString("\t" + opString(zi.Op2))
	}

	return out.String()
}

// これ以降は 式 (Exspression)

// ラベル
type Label struct {
	nodeType   NodeType
	Name       string
	LineNumber int
}

func (le *Label) expressionNode()          {}
func (le *Label) NodeType() NodeType       { return le.nodeType }
func (le *Label) NodeSubType() NodeSubType { return 0 }
func (le *Label) String() string           { return le.Name }

// 数値
type NumberLiteral struct {
	Value int
}

func (nl *NumberLiteral) expressionNode()          {}
func (nl *NumberLiteral) NodeType() NodeType       { return NODE_NUMBER }
func (nl *NumberLiteral) NodeSubType() NodeSubType { return 0 }
func (nl *NumberLiteral) String() string {
	return fmt.Sprintf("%d", nl.Value)
}

// 配列
type ArrayLiteral struct {
	Elements *ExpressionList
}

func (al *ArrayLiteral) expressionNode()          {}
func (al *ArrayLiteral) NodeType() NodeType       { return NODE_ARRAY }
func (al *ArrayLiteral) NodeSubType() NodeSubType { return 0 }
func (al *ArrayLiteral) String() string {
	elems := []string{}

	for _, e := range al.Elements.Expressions {
		elems = append(elems, e.String())
	}
	return "[" + strings.Join(elems, ", ") + "]"
}

// 添え字参照
type IndexedExpression struct {
	Left       Expression
	Index      Expression
	lineNumber int
}

func (ie *IndexedExpression) expressionNode()          {}
func (ie *IndexedExpression) NodeType() NodeType       { return NODE_INDEXED_EXPR }
func (ie *IndexedExpression) NodeSubType() NodeSubType { return 0 }
func (ie *IndexedExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ie.Left.String())
	out.WriteRune('[')
	if ie.Index != nil {
		out.WriteString(ie.Index.String())
	}
	out.WriteRune(']')

	return out.String()
}

// レジスタ
type RegisterLiteral struct {
	RegisterType int
	Register     int
}

func (rl *RegisterLiteral) expressionNode()          {}
func (rl *RegisterLiteral) NodeType() NodeType       { return NodeType(rl.RegisterType) }
func (rl *RegisterLiteral) NodeSubType() NodeSubType { return NodeSubType(rl.Register) }
func (rl *RegisterLiteral) String() string {
	return Z80OpCode2Name(rl.Register)
}

// フラグ
type FlagLiteral struct {
	Flag int
}

func (fl *FlagLiteral) expressionNode()          {}
func (fl *FlagLiteral) NodeType() NodeType       { return Z80_FLAG }
func (fl *FlagLiteral) NodeSubType() NodeSubType { return NodeSubType(fl.Flag) }
func (fl *FlagLiteral) String() string {
	return Z80OpCode2Name(fl.Flag)
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

func (ie *IndirectExpression) expressionNode()          {}
func (ie *IndirectExpression) NodeType() NodeType       { return NODE_INDIRECT }
func (ie *IndirectExpression) NodeSubType() NodeSubType { return 0 }
func (ie *IndirectExpression) String() string {
	expr := trimParen(ie.Expression.String())
	return "(" + expr + ")"
}

// 中置演算子式
type InfixExpression struct {
	OpCode int
	Op1    Node
	Op2    Node
}

func (ie *InfixExpression) expressionNode()          {}
func (ie *InfixExpression) NodeType() NodeType       { return NODE_INFIX_EXPR }
func (ie *InfixExpression) NodeSubType() NodeSubType { return NodeSubType(ie.OpCode) }
func (ie *InfixExpression) String() string {
	var op1, op2 string
	if ie.Op1 == nil {
		op1 = "<nil>"
	} else {
		op1 = ie.Op1.String()
	}
	if ie.Op2 == nil {
		op2 = "<nil>"
	} else {
		op2 = ie.Op2.String()
	}
	var out bytes.Buffer

	out.WriteString("(" + op1 + " ")
	out.WriteString(tokenLiteral(ie.OpCode))
	out.WriteString(" " + op2 + ")")

	return out.String()
}

// 前置演算子式
type PrefixExpression struct {
	OpCode int
	Op     Node
}

func (pe *PrefixExpression) expressionNode()          {}
func (pe *PrefixExpression) NodeType() NodeType       { return NodeType(NODE_PREFIX_EXPR) }
func (pe *PrefixExpression) NodeSubType() NodeSubType { return NodeSubType(pe.OpCode) }
func (pe *PrefixExpression) String() string {
	var op string
	if pe.Op == nil {
		op = "<nil>"
	} else {
		op = pe.Op.String()
	}

	return "(" + tokenLiteral(pe.OpCode) + op + ")"
}

// 関数呼出し
type CallExpression struct {
	Function  Node
	Arguments *ExpressionList
}

func (ce *CallExpression) expressionNode()          {}
func (ce *CallExpression) NodeType() NodeType       { return NODE_CALL }
func (ce *CallExpression) NodeSubType() NodeSubType { return 0 }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	out.WriteString(ce.Function.String())
	out.WriteRune('(')
	out.WriteString(ce.Arguments.String())
	out.WriteRune(')')

	return out.String()
}

// 式リスト
type ExpressionList struct {
	Expressions []Expression
}

func (el *ExpressionList) expressionNode()          {}
func (el *ExpressionList) NodeType() NodeType       { return NODE_EXPR_LIST }
func (el *ExpressionList) NodeSubType() NodeSubType { return 0 }
func (el *ExpressionList) String() string {
	list := []string{}

	for _, e := range el.Expressions {
		list = append(list, e.String())
	}

	return strings.Join(list, ", ")
}
