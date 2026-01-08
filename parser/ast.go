package parser

import (
	// "bytes"
	"bytes"
	"fmt"
	"strings"
	"yas80/fileblock"
)

const (
	NODE_NODE = iota + 1

	// program
	NODE_PROGRAM

	// eror
	NODE_ERROR

	// statement
	NODE_STMT
	NODE_DELETED_STMT
	NODE_LABEL_STMT
	NODE_EXPR_STMT
	NODE_CONST_STMT
	NODE_VAR_STMT
	NODE_ASSIGN_STMT
	NODE_ENUM_STMT
	NODE_ENUM_ELEMENTS_STMT
	NODE_REPT_STMT
	NODE_IF_STMT
	NODE_BLOCK_STMT
	NODE_FUNC_STMT
	NODE_EXITM_STMT
	NODE_RETURN_STMT
	NODE_PROC_STMT
	NODE_PROC_BLOCK_STMT
	NODE_MACRO_STMT
	NODE_MACRO_CALL_STMT
	NODE_MACRO_BLOCK_STMT
	NODE_EXPANDED_MACRO_CALL_STMT
	NODE_SET_SYSVAR_STMT
	NODE_DATA_STMT
	NODE_DATA_STORE_STMT

	// expression
	NODE_EXPR
	NODE_ENUM_ELEMENT
	NODE_NUMBER
	NODE_STRING
	NODE_IDENT
	NODE_LOCAL_IDENT
	NODE_DOT_IDENT
	NODE_ARRAY
	NODE_INDEXED_EXPR
	NODE_INFIX_EXPR
	NODE_PREFIX_EXPR
	NODE_CALL
	NODE_EXPR_LIST
	NODE_LABEL
	NODE_LOCAL_LABEL
	NODE_AT_LABEL
	NODE_INDIRECT // for Z80
)

type NodeType int

// interface

// Node
type Node interface {
	NodeType() NodeType
	String() string
}

// 文
type Statement interface {
	Node
	ReplaceContext(ctx fileblock.Context)
	StatementNode() *fileblock.Context
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

func (p *Program) NodeType() NodeType { return NODE_PROGRAM }
func (p *Program) String() string {
	var lines []string
	for _, s := range p.Statements {

		lines = append(lines, s.String())
	}
	return strings.Join(lines, "\n")
}

// Error(Expression, Statement)
type ParseError struct {
	Message string
	Context *fileblock.Context
}

func (s *ParseError) StatementNode() *fileblock.Context { return s.Context }
func (s *ParseError) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}

func (s *ParseError) expressionNode()    {}
func (s *ParseError) NodeType() NodeType { return NODE_ERROR }
func (s *ParseError) String() string {
	if s.Context == nil {
		return s.Message
	} else {
		return fmt.Sprintf("%s %d", s.Message, s.Context.Line)
	}
}

// ラベル - 独立した文として生成
type LabelStatement struct {
	Name    Expression
	Context *fileblock.Context
}

func (s *LabelStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *LabelStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *LabelStatement) NodeType() NodeType { return NODE_LABEL_STMT }
func (s *LabelStatement) String() string {
	return s.Name.String() + ":"
}

// Proc
type ProcStatement struct {
	Name    string
	Block   *BlockStatement
	Context *fileblock.Context
}

func (s *ProcStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *ProcStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *ProcStatement) NodeType() NodeType { return NODE_PROC_STMT }
func (s *ProcStatement) String() string     { return fmt.Sprintf("PROC(%s)", s.Name) }

// Prock block statement
type ProcBlockStatement struct {
	Name    string
	Block   []Node
	Context *fileblock.Context
}

func (s *ProcBlockStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *ProcBlockStatement) NodeType() NodeType                { return NODE_PROC_BLOCK_STMT }
func (s *ProcBlockStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *ProcBlockStatement) String() string {
	var out bytes.Buffer

	out.WriteString("PROC BLOCK(" + s.Name + ") {\n")
	for _, n := range s.Block {
		out.WriteString(n.String() + "\n")
	}
	out.WriteString("}")

	return out.String()
}

// 式文 - Expression Statement
type ExpressionStatement struct {
	Value   Expression
	Context *fileblock.Context
}

func (s *ExpressionStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *ExpressionStatement) NodeType() NodeType                { return NODE_EXPR_STMT }
func (s *ExpressionStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *ExpressionStatement) String() string { return s.Value.String() }

// enum 定義文
type EnumStatement struct {
	Name     string
	Elements *EnumElements
	Context  *fileblock.Context
}

func (s *EnumStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *EnumStatement) NodeType() NodeType                { return NODE_ENUM_STMT }
func (s *EnumStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *EnumStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.Name + " ENUM\n")
	out.WriteString(s.Elements.String() + "\n")
	out.WriteString("ENDE")

	return out.String()
}

// enum 要素定義文の集合
type EnumElements struct {
	Elements []*EnumElement
}

func (s *EnumElements) StatementNode() *fileblock.Context    { return nil }
func (s *EnumElements) NodeType() NodeType                   { return NODE_ENUM_ELEMENTS_STMT }
func (s *EnumElements) ReplaceContext(ctx fileblock.Context) {}
func (s *EnumElements) String() string {
	stmts := []string{}
	for _, e := range s.Elements {
		stmts = append(stmts, e.String())
	}
	return strings.Join(stmts, "\n")
}

// enum 要素定義文
type EnumElement struct {
	Name    string
	Value   Expression
	Context *fileblock.Context
}

func (s *EnumElement) StatementNode() *fileblock.Context { return s.Context }
func (s *EnumElement) NodeType() NodeType                { return NODE_ENUM_ELEMENT }
func (s *EnumElement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *EnumElement) String() string {
	if s.Value == nil {
		return s.Name
	} else {
		return s.Name + " = " + s.Value.String()
	}
}

// rept statment
type ReptStatement struct {
	MaxCount Expression
	Block    *BlockStatement
	Context  *fileblock.Context
}

func (s *ReptStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *ReptStatement) NodeType() NodeType                { return NODE_REPT_STMT }
func (s *ReptStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *ReptStatement) String() string {
	var out bytes.Buffer

	out.WriteString("REPT ")
	out.WriteString(s.MaxCount.String() + "\n")
	block := s.Block.String()
	if block != "" {
		out.WriteString(block + "\n")
	}
	out.WriteString("ENDR")

	return out.String()
}

// システム変数設定 - REPT 展開時で使用する
type SetSysVarStatement struct {
	Name    string
	Value   Expression
	Context *fileblock.Context
}

func (s *SetSysVarStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *SetSysVarStatement) NodeType() NodeType                { return NODE_SET_SYSVAR_STMT }
func (s *SetSysVarStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *SetSysVarStatement) String() string {
	return fmt.Sprintf("SET_SYS_VAR(%s, %s)", s.Name, s.Value.String())
}

// if statement
type IfStatement struct {
	Condition   Expression
	Consequence Node
	Alternative Node
	Context     *fileblock.Context
}

func (s *IfStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *IfStatement) NodeType() NodeType                { return NODE_IF_STMT }
func (s *IfStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *IfStatement) String() string {
	var out bytes.Buffer

	out.WriteString("IF " + s.Condition.String() + "\n")
	if s.Consequence == nil {
		out.WriteString("<nil>\n")
	} else {
		block := s.Consequence.String()
		if block != "" {
			out.WriteString(block + "\n")
		}
	}
	out.WriteString("ELSE\n")
	if s.Alternative == nil {
		out.WriteString("<nil>\n")
	} else {
		block := s.Alternative.String()
		if block != "" {
			out.WriteString(block + "\n")
		}
	}
	out.WriteString("ENDIF")

	return out.String()
}

// func 文
type FuncStatement struct {
	Name    string
	Params  []string
	Block   *BlockStatement
	Context *fileblock.Context
}

func (s *FuncStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *FuncStatement) NodeType() NodeType                { return NODE_FUNC_STMT }
func (s *FuncStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *FuncStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.Name + " FUNC " + strings.Join(s.Params, ", ") + "\n")
	out.WriteString(s.Block.String() + "\n")
	out.WriteString("ENDF")

	return out.String()
}

// macro 定義文
type MacroStatement struct {
	Name    string
	Params  []string
	Body    *BlockStatement
	Context *fileblock.Context
}

func (s *MacroStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *MacroStatement) NodeType() NodeType                { return NODE_MACRO_STMT }
func (s *MacroStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *MacroStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.Name + " MACRO " + strings.Join(s.Params, ", ") + "\n")
	out.WriteString(s.Body.String() + "\n")
	out.WriteString("ENDM")

	return out.String()
}

// macro 呼出し Parse 後
type MacroCallStatement struct {
	Name    string
	Args    *ExpressionList
	Context *fileblock.Context
}

func (s *MacroCallStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *MacroCallStatement) NodeType() NodeType                { return NODE_MACRO_CALL_STMT }
func (s *MacroCallStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *MacroCallStatement) String() string {
	args := []string{}
	for _, arg := range s.Args.Expressions {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("MACRO %s CALL with %s", s.Name, strings.Join(args, ","))
}

// block statement
type BlockStatement struct {
	Block []Node
}

func (s *BlockStatement) StatementNode() *fileblock.Context    { return nil }
func (s *BlockStatement) NodeType() NodeType                   { return NODE_BLOCK_STMT }
func (s *BlockStatement) ReplaceContext(ctx fileblock.Context) {}
func (s *BlockStatement) String() string {
	stmts := []string{}

	for _, s := range s.Block {
		stmts = append(stmts, s.String())
	}
	return strings.Join(stmts, "\n")
}

// macro block statement
type MacroBlockStatement struct {
	Name    string
	Index   int // REPT 用
	Count   int // REPT 用
	Block   []Node
	Context *fileblock.Context
}

func (s *MacroBlockStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *MacroBlockStatement) NodeType() NodeType                { return NODE_MACRO_BLOCK_STMT }
func (s *MacroBlockStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *MacroBlockStatement) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("MACRO BLOCK(%s) {\n", s.Name))
	for _, s := range s.Block {
		out.WriteString(s.String() + "\n")
	}
	out.WriteString("}")

	return out.String()
}

// 定数定義文 - CONST, EQU Statement
type ConstStatement struct {
	Name    Expression
	Value   Expression
	Context *fileblock.Context
}

func (s *ConstStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *ConstStatement) NodeType() NodeType                { return NODE_CONST_STMT }
func (s *ConstStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *ConstStatement) String() string {
	var out bytes.Buffer

	out.WriteString("CONST ")
	out.WriteString(s.Name.String())
	out.WriteString(" = ")
	out.WriteString(s.Value.String())

	return out.String()
}

// 変数定義文 - VAR
type VariableStatement struct {
	Name    *Ident
	Value   Expression
	Context *fileblock.Context
}

func (s *VariableStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *VariableStatement) NodeType() NodeType                { return NODE_VAR_STMT }
func (s *VariableStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *VariableStatement) String() string {
	var out bytes.Buffer

	out.WriteString("VAR ")
	out.WriteString(s.Name.Name)
	out.WriteString(" = ")
	out.WriteString(s.Value.String())

	return out.String()
}

// 変数代入文
type AssignStatement struct {
	Left    Expression
	Value   Expression
	Context *fileblock.Context
}

func (s *AssignStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *AssignStatement) NodeType() NodeType                { return NODE_ASSIGN_STMT }
func (s *AssignStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *AssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.Left.String())
	out.WriteString(" = ")
	out.WriteString(s.Value.String())

	return out.String()
}

// Exitm 文
type ExitmStatement struct {
	Context *fileblock.Context
}

func (s *ExitmStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *ExitmStatement) NodeType() NodeType                { return NODE_EXITM_STMT }
func (s *ExitmStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *ExitmStatement) String() string { return "EXITM" }

// Return 文
type ReturnStatement struct {
	Value   Expression
	Context *fileblock.Context
}

func (s *ReturnStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *ReturnStatement) NodeType() NodeType                { return NODE_RETURN_STMT }
func (s *ReturnStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *ReturnStatement) String() string {
	str := "RETURN"
	if s.Value != nil {
		str += " " + s.Value.String()
	}
	return str
}

type DataStatement struct {
	Label   Expression
	Size    int
	Values  []Expression
	Context *fileblock.Context
}

func (s *DataStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *DataStatement) NodeType() NodeType                { return NODE_DATA_STMT }
func (s *DataStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *DataStatement) String() string {
	return fmt.Sprintf("DATA [size: %d, len: %d]", s.Size, len(s.Values))
}

type DataStoreStatement struct {
	Label     Expression
	Size      int
	Count     Expression
	FillValue Expression
	Context   *fileblock.Context
}

func (s *DataStoreStatement) StatementNode() *fileblock.Context { return s.Context }
func (s *DataStoreStatement) NodeType() NodeType                { return NODE_DATA_STORE_STMT }
func (s *DataStoreStatement) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *DataStoreStatement) String() string {
	var f string
	if s.FillValue == nil {
		f = "nil"
	} else {
		f = s.FillValue.String()
	}
	return fmt.Sprintf("DS [size: %d, count: %s, fill: %s]", s.Size, s.Count.String(), f)
}

// Z80 命令文 - Z80Instruction Statement
type Z80Instruction struct {
	Label    Expression
	InstType int
	Opcode   int
	Op1      Expression
	Op2      Expression
	Context  *fileblock.Context
}

func (s *Z80Instruction) StatementNode() *fileblock.Context { return s.Context }
func (s *Z80Instruction) NodeType() NodeType {
	return NodeType(s.InstType)
}
func (s *Z80Instruction) ReplaceContext(ctx fileblock.Context) {
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *Z80Instruction) String() string {
	var out bytes.Buffer

	if s.Label != nil {
		out.WriteString(s.Label.String() + ": ")
	}
	out.WriteString(Z80Opcode2Name(s.Opcode))
	switch {
	case s.Op1 == nil && s.Op2 == nil:
		break
	case s.Op1 != nil && s.Op2 != nil:
		out.WriteString(" " + opString(s.Op1))
		out.WriteString(", " + opString(s.Op2))
	case s.Op1 != nil:
		out.WriteString(" " + opString(s.Op1))
	default:
		out.WriteString(" " + opString(s.Op2))
	}

	return out.String()
}

// これ以降は 式 (Exspression)

// ラベル
type Label struct {
	LabelType int
	Name      string
	Context   *fileblock.Context
}

func (e *Label) expressionNode()    {}
func (e *Label) NodeType() NodeType { return NODE_LABEL }
func (e *Label) String() string     { return e.Name }

// 数値
type NumberLiteral struct {
	Value   int
	Context *fileblock.Context
}

func (e *NumberLiteral) expressionNode()    {}
func (e *NumberLiteral) NodeType() NodeType { return NODE_NUMBER }
func (e *NumberLiteral) String() string {
	return fmt.Sprintf("%d", e.Value)
}

// 文字列
type StringLiteral struct {
	Value   string
	Context *fileblock.Context
}

func (e *StringLiteral) expressionNode()    {}
func (e *StringLiteral) NodeType() NodeType { return NODE_STRING }
func (e *StringLiteral) String() string {
	return fmt.Sprintf("%q", e.Value)
}

// 配列
type ArrayLiteral struct {
	Elements *ExpressionList
	Context  *fileblock.Context
}

func (e *ArrayLiteral) expressionNode()    {}
func (e *ArrayLiteral) NodeType() NodeType { return NODE_ARRAY }
func (e *ArrayLiteral) String() string {
	elems := []string{}

	for _, e := range e.Elements.Expressions {
		elems = append(elems, e.String())
	}
	return "[" + strings.Join(elems, ", ") + "]"
}

// 添え字参照
type IndexedExpression struct {
	Left    Expression
	Index   Expression
	Context *fileblock.Context
}

func (e *IndexedExpression) expressionNode()    {}
func (e *IndexedExpression) NodeType() NodeType { return NODE_INDEXED_EXPR }
func (e *IndexedExpression) String() string {
	var out bytes.Buffer

	out.WriteString(e.Left.String())
	out.WriteRune('[')
	if e.Index != nil {
		out.WriteString(e.Index.String())
	}
	out.WriteRune(']')

	return out.String()
}

// レジスタ
type RegisterLiteral struct {
	RegisterType int
	Register     int
	Context      *fileblock.Context
}

func (e *RegisterLiteral) expressionNode()    {}
func (e *RegisterLiteral) NodeType() NodeType { return NodeType(e.RegisterType) }
func (e *RegisterLiteral) String() string {
	return Z80Opcode2Name(e.Register)
}

// フラグ
type FlagLiteral struct {
	Flag    int
	Context *fileblock.Context
}

func (e *FlagLiteral) expressionNode()    {}
func (e *FlagLiteral) NodeType() NodeType { return Z80_FLAG }
func (e *FlagLiteral) String() string {
	return Z80Opcode2Name(e.Flag)
}

// 識別子
type Ident struct {
	Name      string
	IdentType int
	Value     Expression
	Context   *fileblock.Context
}

func (e *Ident) expressionNode()    {}
func (e *Ident) NodeType() NodeType { return NODE_IDENT }
func (e *Ident) String() string     { return e.Name }

// ドット識別子
type DotIdent struct {
	Name    string
	Left    string
	Right   string
	Value   Expression
	Context *fileblock.Context
}

func (e *DotIdent) expressionNode()    {}
func (e *DotIdent) NodeType() NodeType { return NODE_DOT_IDENT }
func (e *DotIdent) String() string     { return e.Name }

// 間接指定
type IndirectExpression struct {
	Expression Expression
}

func (e *IndirectExpression) expressionNode()    {}
func (e *IndirectExpression) NodeType() NodeType { return NODE_INDIRECT }
func (e *IndirectExpression) String() string {
	expr := trimParen(e.Expression.String())
	return "(" + expr + ")"
}

// 中置演算子式
type InfixExpression struct {
	Operator int
	Op1      Expression
	Op2      Expression
	Context  *fileblock.Context
}

func (e *InfixExpression) expressionNode()    {}
func (e *InfixExpression) NodeType() NodeType { return NODE_INFIX_EXPR }
func (e *InfixExpression) String() string {
	var op1, op2 string
	if e.Op1 == nil {
		op1 = "<nil>"
	} else {
		op1 = e.Op1.String()
	}
	if e.Op2 == nil {
		op2 = "<nil>"
	} else {
		op2 = e.Op2.String()
	}
	var out bytes.Buffer

	out.WriteString("(" + op1 + " ")
	out.WriteString(TokenLiteral(e.Operator))
	out.WriteString(" " + op2 + ")")

	return out.String()
}

// 前置演算子式
type PrefixExpression struct {
	Operator int
	Op       Expression
	Context  *fileblock.Context
}

func (e *PrefixExpression) expressionNode()    {}
func (e *PrefixExpression) NodeType() NodeType { return NodeType(NODE_PREFIX_EXPR) }
func (e *PrefixExpression) String() string {
	var op string
	if e.Op == nil {
		op = "<nil>"
	} else {
		op = e.Op.String()
	}

	return "(" + TokenLiteral(e.Operator) + op + ")"
}

// 関数呼出し
type FuncCallExpression struct {
	Name      string
	Arguments *ExpressionList
	Context   *fileblock.Context
}

func (e *FuncCallExpression) expressionNode()    {}
func (e *FuncCallExpression) NodeType() NodeType { return NODE_CALL }
func (e *FuncCallExpression) String() string {
	var out bytes.Buffer

	out.WriteString(e.Name + "(")
	out.WriteString(e.Arguments.String())
	out.WriteRune(')')

	return out.String()
}

// 式リスト
type ExpressionList struct {
	Expressions []Expression
}

func (e *ExpressionList) expressionNode()    {}
func (e *ExpressionList) NodeType() NodeType { return NODE_EXPR_LIST }
func (e *ExpressionList) String() string {
	list := []string{}

	for _, e := range e.Expressions {
		list = append(list, e.String())
	}

	return strings.Join(list, ", ")
}
