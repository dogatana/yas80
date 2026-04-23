package parser

import (
	// "bytes"
	"bytes"
	"fmt"
	"strings"

	"github.com/dogatana/yas80/filecontent"
	"github.com/dogatana/yas80/intern"
	"github.com/dogatana/yas80/internal/util"
)

const (
	NODE_NODE = iota + 1

	// program
	NODE_PROGRAM
	NODE_FILE
	NODE_INCLUDE

	// eror
	NODE_ERROR

	// statement
	NODE_NULL // eval用: エラーが発生した文を置き換える
	NODE_STMT
	NODE_CHARMAP_STMT
	NODE_ORG_STMT
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
	NODE_DATA_DEF_STMT
	NODE_DATA_STORE_STMT
	NODE_COMMENT_STMT
	NODE_END_STMT

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
	NODE_ANON_LABEL
	NODE_ADDR_INDIRECT
	NODE_REG_INDIRECT
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
	GetContext() *filecontent.Context
	ReplaceContext(ctx filecontent.Context)
}

// 式
type Expression interface {
	Node
	expressionNode()
}

// 実装 (struct)

// Null
type NullStatement struct {
	Context *filecontent.Context
}

func (s *NullStatement) NodeType() NodeType               { return NODE_NULL }
func (s *NullStatement) GetContext() *filecontent.Context { return s.Context }
func (s *NullStatement) String() string                   { return "NullStatement" }
func (s *NullStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}

// Error(Expression, Statement)
type ParseError struct {
	Message string
	Context *filecontent.Context
}

func (s *ParseError) GetContext() *filecontent.Context { return s.Context }
func (s *ParseError) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
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

type AllocType int

const (
	ALLOC_ABS = iota
	ALLOC_REL
)

// file
type FileStatement struct {
	Filename string
	Line     int
}

func (s *FileStatement) NodeType() NodeType                   { return NODE_FILE }
func (s *FileStatement) GetContext() *filecontent.Context     { return nil }
func (s *FileStatement) ReplaceContext(_ filecontent.Context) {}
func (s *FileStatement) String() string {
	return fmt.Sprintf("FILE %q:%d", s.Filename, s.Line)
}

// include
type IncludeStatement struct {
	Filename string
	Context  *filecontent.Context
}

func (s *IncludeStatement) NodeType() NodeType               { return NODE_INCLUDE }
func (s *IncludeStatement) GetContext() *filecontent.Context { return s.Context }
func (s *IncludeStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *IncludeStatement) String() string {
	return "INCLUDE " + s.Filename
}

// charmap
type CharmapStatement struct {
	NameID   intern.SymbolID
	Filename Expression
	DefChar  Expression
	Context  *filecontent.Context
}

func (s *CharmapStatement) NodeType() NodeType               { return NODE_CHARMAP_STMT }
func (s *CharmapStatement) GetContext() *filecontent.Context { return s.Context }
func (s *CharmapStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *CharmapStatement) String() string {
	str := fmt.Sprintf("CHARMAP %s, %q", s.NameID, s.Filename)
	if s.DefChar != nil {
		str += ", " + s.DefChar.String()
	}
	return str
}

// org
type OrgStatement struct {
	Address   Expression
	AllocType AllocType
	Context   *filecontent.Context
}

func (s *OrgStatement) GetContext() *filecontent.Context { return s.Context }
func (s *OrgStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *OrgStatement) NodeType() NodeType { return NODE_ORG_STMT }
func (s *OrgStatement) String() string {
	ret := fmt.Sprintf("ORG %s", s.Address.String())
	if s.AllocType == ALLOC_REL {
		ret += ", REL"
	}
	return ret
}

// END
type EndStatement struct {
	Start   Expression
	Context *filecontent.Context
}

func (s *EndStatement) GetContext() *filecontent.Context { return s.Context }
func (s *EndStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *EndStatement) NodeType() NodeType { return NODE_END_STMT }
func (s *EndStatement) String() string {
	if s.Start == nil {
		return "END"
	} else {
		return "END " + s.Start.String()
	}
}

// ラベル - 独立した文として生成
type LabelStatement struct {
	Name    Expression
	Context *filecontent.Context
}

func (s *LabelStatement) GetContext() *filecontent.Context { return s.Context }
func (s *LabelStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *LabelStatement) NodeType() NodeType { return NODE_LABEL_STMT }
func (s *LabelStatement) String() string {
	return s.Name.String() + ":"
}

// Proc
type ProcStatement struct {
	Name    Expression
	Block   *BlockStatement
	Context *filecontent.Context
}

func (s *ProcStatement) GetContext() *filecontent.Context { return s.Context }
func (s *ProcStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *ProcStatement) NodeType() NodeType { return NODE_PROC_STMT }
func (s *ProcStatement) String() string     { return fmt.Sprintf("PROC(%s)", s.Name) }

// Prock block statement
type ProcBlockStatement struct {
	Name    string
	Block   []Statement
	Context *filecontent.Context
}

func (s *ProcBlockStatement) GetContext() *filecontent.Context { return s.Context }
func (s *ProcBlockStatement) NodeType() NodeType               { return NODE_PROC_BLOCK_STMT }
func (s *ProcBlockStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *ProcBlockStatement) String() string {
	var out bytes.Buffer

	out.WriteString("PROC BLOCK(" + s.Name + ") {\n")
	block := util.Map(s.Block, func(s Statement) string { return s.String() })
	out.WriteString(strings.Join(block, "\n"))
	out.WriteString("}")

	return out.String()
}

// enum 定義文
type EnumStatement struct {
	NameID   intern.SymbolID
	Elements *EnumElements
	Context  *filecontent.Context
}

func (s *EnumStatement) GetContext() *filecontent.Context { return s.Context }
func (s *EnumStatement) NodeType() NodeType               { return NODE_ENUM_STMT }
func (s *EnumStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *EnumStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.NameID.String() + " ENUM\n")
	out.WriteString(s.Elements.String() + "\n")
	out.WriteString("ENDE")

	return out.String()
}

// enum 要素定義文の集合
type EnumElements struct {
	Elements []*EnumElement
}

func (s *EnumElements) GetContext() *filecontent.Context       { return nil }
func (s *EnumElements) NodeType() NodeType                     { return NODE_ENUM_ELEMENTS_STMT }
func (s *EnumElements) ReplaceContext(ctx filecontent.Context) {}
func (s *EnumElements) String() string {
	stmts := util.Map(s.Elements, func(e *EnumElement) string { return e.String() })
	return strings.Join(stmts, "\n")
}

// enum 要素定義文
type EnumElement struct {
	NameID  intern.SymbolID
	Value   Expression
	Context *filecontent.Context
}

func (s *EnumElement) GetContext() *filecontent.Context { return s.Context }
func (s *EnumElement) NodeType() NodeType               { return NODE_ENUM_ELEMENT }
func (s *EnumElement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *EnumElement) String() string {
	name := s.NameID.String()
	if s.Value == nil {
		return name
	} else {
		return name + " = " + s.Value.String()
	}
}

// rept statment
type ReptStatement struct {
	Label    Expression
	MaxCount Expression
	Block    *BlockStatement
	Start    int // REPT 行
	Context  *filecontent.Context
}

func (s *ReptStatement) GetContext() *filecontent.Context { return s.Context }
func (s *ReptStatement) NodeType() NodeType               { return NODE_REPT_STMT }
func (s *ReptStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
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
	Value   any // Expression or Object
	Context *filecontent.Context
}

func (s *SetSysVarStatement) GetContext() *filecontent.Context { return s.Context }
func (s *SetSysVarStatement) NodeType() NodeType               { return NODE_SET_SYSVAR_STMT }
func (s *SetSysVarStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *SetSysVarStatement) String() string {
	return fmt.Sprintf("SET_SYS_VAR(%s, %v)", s.Name, s.Value)
}

// if statement
type IfStatement struct {
	Condition   Expression
	Consequence Node
	Alternative Node
	Context     *filecontent.Context
}

func (s *IfStatement) GetContext() *filecontent.Context { return s.Context }
func (s *IfStatement) NodeType() NodeType               { return NODE_IF_STMT }
func (s *IfStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
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
	NameID  intern.SymbolID
	Params  []string
	Block   *BlockStatement
	Context *filecontent.Context
}

func (s *FuncStatement) GetContext() *filecontent.Context { return s.Context }
func (s *FuncStatement) NodeType() NodeType               { return NODE_FUNC_STMT }
func (s *FuncStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *FuncStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.NameID.String() + " FUNC " + strings.Join(s.Params, ", ") + "\n")
	out.WriteString(s.Block.String() + "\n")
	out.WriteString("ENDF")

	return out.String()
}

// macro 定義文
type MacroStatement struct {
	NameID  intern.SymbolID
	Params  []string
	Body    *BlockStatement
	End     int // ENDM 行
	Context *filecontent.Context
}

func (s *MacroStatement) GetContext() *filecontent.Context { return s.Context }
func (s *MacroStatement) NodeType() NodeType               { return NODE_MACRO_STMT }
func (s *MacroStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *MacroStatement) String() string {
	var out bytes.Buffer

	out.WriteString(s.NameID.String() + " MACRO " + strings.Join(s.Params, ", ") + "\n")
	out.WriteString(s.Body.String() + "\n")
	out.WriteString("ENDM")

	return out.String()
}

// macro 呼出し Parse 後
type MacroCallStatement struct {
	Label   Expression
	NameID  intern.SymbolID
	Args    *ExpressionList
	Context *filecontent.Context
}

func (s *MacroCallStatement) GetContext() *filecontent.Context { return s.Context }
func (s *MacroCallStatement) NodeType() NodeType               { return NODE_MACRO_CALL_STMT }
func (s *MacroCallStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *MacroCallStatement) String() string {
	args := util.Map(s.Args.Expressions, func(s Expression) string { return s.String() })
	return fmt.Sprintf("MACRO %s CALL with %s", s.NameID, strings.Join(args, ","))
}

// block statement
type BlockStatement struct {
	Block []Statement
}

func (s *BlockStatement) GetContext() *filecontent.Context       { return &filecontent.Context{} }
func (s *BlockStatement) NodeType() NodeType                     { return NODE_BLOCK_STMT }
func (s *BlockStatement) ReplaceContext(ctx filecontent.Context) {}
func (s *BlockStatement) String() string {

	stmts := util.Map(s.Block, func(s Statement) string { return s.String() })
	return strings.Join(stmts, "\n")
}

// macro block statement
type MacroBlockStatement struct {
	Label   Expression
	NameID  intern.SymbolID // マクロ名 もしくは "REPT"
	Index   int             // REPT 用
	Count   int             // REPT 用
	Start   int             // REPT 行
	Value   any             // REPT 用 Expression/Object
	Block   []Statement
	Context *filecontent.Context
}

func (s *MacroBlockStatement) GetContext() *filecontent.Context { return s.Context }
func (s *MacroBlockStatement) NodeType() NodeType               { return NODE_MACRO_BLOCK_STMT }
func (s *MacroBlockStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *MacroBlockStatement) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("MACRO BLOCK(%s) {\n", s.NameID))
	block := util.Map(s.Block, func(s Statement) string { return s.String() })
	out.WriteString(strings.Join(block, "\n"))
	out.WriteString("}")

	return out.String()
}

// 定数定義文 - CONST, EQU Statement
type ConstStatement struct {
	Name    Expression
	Value   Expression
	Context *filecontent.Context
}

func (s *ConstStatement) GetContext() *filecontent.Context { return s.Context }
func (s *ConstStatement) NodeType() NodeType               { return NODE_CONST_STMT }
func (s *ConstStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
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
	Name    Expression
	Value   Expression
	Context *filecontent.Context
}

func (s *VariableStatement) GetContext() *filecontent.Context { return s.Context }
func (s *VariableStatement) NodeType() NodeType               { return NODE_VAR_STMT }
func (s *VariableStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *VariableStatement) String() string {
	var out bytes.Buffer

	out.WriteString("VAR ")
	out.WriteString(s.Name.(*Ident).NameID.String())
	out.WriteString(" = ")
	out.WriteString(s.Value.String())

	return out.String()
}

// 変数代入文
type AssignStatement struct {
	Left    Expression
	Value   Expression
	Context *filecontent.Context
}

func (s *AssignStatement) GetContext() *filecontent.Context { return s.Context }
func (s *AssignStatement) NodeType() NodeType               { return NODE_ASSIGN_STMT }
func (s *AssignStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
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
	Context *filecontent.Context
}

func (s *ExitmStatement) GetContext() *filecontent.Context { return s.Context }
func (s *ExitmStatement) NodeType() NodeType               { return NODE_EXITM_STMT }
func (s *ExitmStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *ExitmStatement) String() string { return "EXITM" }

// Return 文
type ReturnStatement struct {
	Value   Expression
	Context *filecontent.Context
}

func (s *ReturnStatement) GetContext() *filecontent.Context { return s.Context }
func (s *ReturnStatement) NodeType() NodeType               { return NODE_RETURN_STMT }
func (s *ReturnStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
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

type DataDefineStatement struct {
	Label   Expression
	Size    int
	Values  []Expression
	Context *filecontent.Context
}

func (s *DataDefineStatement) GetContext() *filecontent.Context { return s.Context }
func (s *DataDefineStatement) NodeType() NodeType               { return NODE_DATA_DEF_STMT }
func (s *DataDefineStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *DataDefineStatement) String() string {
	return fmt.Sprintf("DATA [size: %d, len: %d]", s.Size, len(s.Values))
}

type DataStoreStatement struct {
	Label     Expression
	Size      int
	Count     Expression
	FillValue Expression
	Context   *filecontent.Context
}

func (s *DataStoreStatement) GetContext() *filecontent.Context { return s.Context }
func (s *DataStoreStatement) NodeType() NodeType               { return NODE_DATA_STORE_STMT }
func (s *DataStoreStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
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
	Context  *filecontent.Context
	Code     any // *object.CodeObject
}

func (s *Z80Instruction) GetContext() *filecontent.Context { return s.Context }
func (s *Z80Instruction) NodeType() NodeType {
	return NodeType(s.InstType)
}
func (s *Z80Instruction) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
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

// Return 文
type CommentStatement struct {
	Text    any
	Context *filecontent.Context
}

func (s *CommentStatement) GetContext() *filecontent.Context { return s.Context }
func (s *CommentStatement) NodeType() NodeType               { return NODE_COMMENT_STMT }
func (s *CommentStatement) ReplaceContext(ctx filecontent.Context) {
	if s.Context.Source != nil {
		return
	}
	ctx.Source = s.Context
	s.Context = &ctx
}
func (s *CommentStatement) String() string {
	if s.Text == nil {
		return "<source>"
	} else {
		t, _ := s.Text.(string)
		return t
	}
}

// これ以降は 式 (Exspression)

// ラベル
type Label struct {
	LabelType int
	NameID    intern.SymbolID
	Context   *filecontent.Context
}

func (e *Label) expressionNode()    {}
func (e *Label) NodeType() NodeType { return NODE_LABEL }
func (e *Label) String() string     { return e.NameID.String() }

// 数値
type NumberLiteral struct {
	Value   int
	Context *filecontent.Context
}

func (e *NumberLiteral) expressionNode()    {}
func (e *NumberLiteral) NodeType() NodeType { return NODE_NUMBER }
func (e *NumberLiteral) String() string {
	return fmt.Sprintf("%d", e.Value)
}

// 文字列
type StringLiteral struct {
	Value   string
	Context *filecontent.Context
}

func (e *StringLiteral) expressionNode()    {}
func (e *StringLiteral) NodeType() NodeType { return NODE_STRING }
func (e *StringLiteral) String() string {
	return fmt.Sprintf("%q", e.Value)
}

// 配列
type ArrayLiteral struct {
	Elements *ExpressionList
	Context  *filecontent.Context
}

func (e *ArrayLiteral) expressionNode()    {}
func (e *ArrayLiteral) NodeType() NodeType { return NODE_ARRAY }
func (e *ArrayLiteral) String() string {

	elems := util.Map(e.Elements.Expressions, func(e Expression) string { return e.String() })
	return "[" + strings.Join(elems, ", ") + "]"
}

// 添え字参照
type IndexedExpression struct {
	Left    Expression
	Index   Expression
	Context *filecontent.Context
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
	Context      *filecontent.Context
}

func (e *RegisterLiteral) expressionNode()    {}
func (e *RegisterLiteral) NodeType() NodeType { return NodeType(e.RegisterType) }
func (e *RegisterLiteral) String() string {
	return Z80Opcode2Name(e.Register)
}

// フラグ
type FlagLiteral struct {
	Flag    int
	Context *filecontent.Context
}

func (e *FlagLiteral) expressionNode()    {}
func (e *FlagLiteral) NodeType() NodeType { return Z80_FLAG }
func (e *FlagLiteral) String() string {
	return Z80Opcode2Name(e.Flag)
}

// 識別子
type Ident struct {
	NameID    intern.SymbolID
	IdentType int
	Value     Expression
	Context   *filecontent.Context
}

func (e *Ident) expressionNode()    {}
func (e *Ident) NodeType() NodeType { return NODE_IDENT }
func (e *Ident) String() string     { return e.NameID.String() }

// ドット識別子
type DotIdent struct {
	NameID  intern.SymbolID
	Left    intern.SymbolID
	Right   intern.SymbolID
	Value   Expression
	Context *filecontent.Context
}

func (e *DotIdent) expressionNode()    {}
func (e *DotIdent) NodeType() NodeType { return NODE_DOT_IDENT }
func (e *DotIdent) String() string     { return e.NameID.String() }

// レジスタ間接指定
type RegIndirectExpression struct {
	Register     *RegisterLiteral
	Displacement Expression
	Context      *filecontent.Context
}

func (e *RegIndirectExpression) expressionNode()    {}
func (e *RegIndirectExpression) NodeType() NodeType { return NODE_REG_INDIRECT }
func (e *RegIndirectExpression) String() string {
	if e.Displacement == nil {
		return "(" + e.Register.String() + ")"
	} else {
		expr := trimParen(e.Displacement.String())
		if expr[0] != '-' {
			expr = "+" + expr
		}
		return "(" + e.Register.String() + expr + ")"
	}
}

// アドレス間接指定
type AddrIndirectExpression struct {
	Address Expression
	Context *filecontent.Context
}

func (e *AddrIndirectExpression) expressionNode()    {}
func (e *AddrIndirectExpression) NodeType() NodeType { return NODE_ADDR_INDIRECT }
func (e *AddrIndirectExpression) String() string {
	return "((" + e.Address.String() + "))"
}

// 中置演算子式
type InfixExpression struct {
	Operator int
	Op1      Expression
	Op2      Expression
	Context  *filecontent.Context
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
	Context  *filecontent.Context
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
	NameID  intern.SymbolID
	Args    *ExpressionList
	Context *filecontent.Context
}

func (e *FuncCallExpression) expressionNode()    {}
func (e *FuncCallExpression) NodeType() NodeType { return NODE_CALL }
func (e *FuncCallExpression) String() string {
	var out bytes.Buffer

	out.WriteString(e.NameID.String() + "(")
	out.WriteString(e.Args.String())
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

	list := util.Map(e.Expressions, func(e Expression) string { return e.String() })
	return strings.Join(list, ", ")
}
