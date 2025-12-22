package parser

import (
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
	NODE_MACRO_STMT
	NODE_MACRO_CALL_STMT
	NODE_MACRO_BLOCK_STMT
	NODE_EXPANDED_MACRO_CALL_STMT
	NODE_SET_SYSVAR_STMT

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
	Message string
	Context *fileblock.Context
}

func (pe *ParseError) statementNode()           {}
func (pe *ParseError) expressionNode()          {}
func (pe *ParseError) NodeType() NodeType       { return NODE_ERROR }
func (pe *ParseError) NodeSubType() NodeSubType { return 0 }
func (pe *ParseError) String() string {
	if pe.Context == nil {
		return pe.Message
	} else {
		return fmt.Sprintf("%s %d", pe.Message, pe.Context.Line)
	}
}

// ラベル - 独立した文として生成
type LabelStatement struct {
	Name    *Label
	Context *fileblock.Context
}

func (ls *LabelStatement) statementNode()           {}
func (ls *LabelStatement) NodeType() NodeType       { return NODE_LABEL_STMT }
func (ls *LabelStatement) NodeSubType() NodeSubType { return ls.Name.LabelType }
func (ls *LabelStatement) String() string {
	out := ls.Name.Name
	if out[0] != '.' {
		out += ":"
	}
	return out
}

// TODO: 仮実装 - PROC/ENDPROC - それぞれ単一文として生成
type ProcStatement struct {
	Name    string
	IsStart bool
	Context *fileblock.Context
}

func (ps *ProcStatement) statementNode()           {}
func (ps *ProcStatement) NodeType() NodeType       { return NODE_PROC_STMT }
func (ps *ProcStatement) NodeSubType() NodeSubType { return 0 }
func (ps *ProcStatement) String() string {
	if ps.IsStart {
		return ps.Name + " PROC"
	} else {
		return "ENDP"
	}
}

// 式文 - Expression Statement
type ExpressionStatement struct {
	Value   Expression
	Context *fileblock.Context
}

func (es *ExpressionStatement) statementNode()           {}
func (es *ExpressionStatement) NodeType() NodeType       { return NODE_EXPR_STMT }
func (es *ExpressionStatement) NodeSubType() NodeSubType { return 0 }
func (es *ExpressionStatement) String() string           { return es.Value.String() }

// enum 定義文
type EnumStatement struct {
	Name     string
	Elements *EnumElements
	Context  *fileblock.Context
}

func (es *EnumStatement) statementNode()           {}
func (es *EnumStatement) NodeType() NodeType       { return NODE_ENUM_STMT }
func (es *EnumStatement) NodeSubType() NodeSubType { return 0 }
func (es *EnumStatement) String() string {
	var out bytes.Buffer

	out.WriteString(es.Name + " ENUM\n")
	out.WriteString(es.Elements.String() + "\n")
	out.WriteString("ENDE")

	return out.String()
}

// enum 要素定義文の集合
type EnumElements struct {
	Elements []*EnumElement
}

func (ee *EnumElements) statementNode()           {}
func (ee *EnumElements) NodeType() NodeType       { return NODE_ENUM_ELEMENTS_STMT }
func (ee *EnumElements) NodeSubType() NodeSubType { return 0 }
func (ee *EnumElements) String() string {
	stmts := []string{}
	for _, e := range ee.Elements {
		stmts = append(stmts, e.String())
	}
	return strings.Join(stmts, "\n")
}

// enum 要素定義文
type EnumElement struct {
	Name    string
	Value   Statement
	Context *fileblock.Context
}

func (ee *EnumElement) statementNode()           {}
func (ee *EnumElement) NodeType() NodeType       { return NODE_ENUM_ELEMENT }
func (ee *EnumElement) NodeSubType() NodeSubType { return 0 }
func (ee *EnumElement) String() string {
	if ee.Value == nil {
		return ee.Name
	} else {
		return ee.Name + " = " + ee.Value.String()
	}
}

// rept statment
type ReptStatement struct {
	MaxCount Expression
	Block    *BlockStatement
	Context  *fileblock.Context
}

func (rs *ReptStatement) statementNode()           {}
func (rs *ReptStatement) NodeType() NodeType       { return NODE_REPT_STMT }
func (rs *ReptStatement) NodeSubType() NodeSubType { return 0 }
func (rs *ReptStatement) String() string {
	var out bytes.Buffer

	out.WriteString("REPT ")
	out.WriteString(rs.MaxCount.String() + "\n")
	block := rs.Block.String()
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

func (s *SetSysVarStatement) statementNode()           {}
func (s *SetSysVarStatement) NodeType() NodeType       { return NODE_SET_SYSVAR_STMT }
func (s *SetSysVarStatement) NodeSubType() NodeSubType { return 0 }
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

func (is *IfStatement) statementNode()           {}
func (is *IfStatement) NodeType() NodeType       { return NODE_IF_STMT }
func (is *IfStatement) NodeSubType() NodeSubType { return 0 }
func (is *IfStatement) String() string {
	var out bytes.Buffer

	out.WriteString("IF " + is.Condition.String() + "\n")
	if is.Consequence == nil {
		out.WriteString("<nil>\n")
	} else {
		block := is.Consequence.String()
		if block != "" {
			out.WriteString(block + "\n")
		}
	}
	out.WriteString("ELSE\n")
	if is.Alternative == nil {
		out.WriteString("<nil>\n")
	} else {
		block := is.Alternative.String()
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

func (fs *FuncStatement) statementNode()           {}
func (fs *FuncStatement) NodeType() NodeType       { return NODE_FUNC_STMT }
func (fs *FuncStatement) NodeSubType() NodeSubType { return 0 }
func (fs *FuncStatement) String() string {
	var out bytes.Buffer

	out.WriteString(fs.Name + " FUNC " + strings.Join(fs.Params, ", ") + "\n")
	out.WriteString(fs.Block.String() + "\n")
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

func (ms *MacroStatement) statementNode()           {}
func (ms *MacroStatement) NodeType() NodeType       { return NODE_MACRO_STMT }
func (ms *MacroStatement) NodeSubType() NodeSubType { return 0 }
func (ms *MacroStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ms.Name + " MACRO " + strings.Join(ms.Params, ", ") + "\n")
	out.WriteString(ms.Body.String() + "\n")
	out.WriteString("ENDM")

	return out.String()
}

// macro 呼出し Parse 後
type MacroCallStatement struct {
	Name    string
	Args    *ExpressionList
	Context *fileblock.Context
}

func (mc *MacroCallStatement) statementNode()           {}
func (mc *MacroCallStatement) NodeType() NodeType       { return NODE_MACRO_CALL_STMT }
func (mc *MacroCallStatement) NodeSubType() NodeSubType { return 0 }
func (mc *MacroCallStatement) String() string {
	args := []string{}
	for _, arg := range mc.Args.Expressions {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("MACRO %s CALL with %s", mc.Name, strings.Join(args, ","))
}

// macro 呼出し Evaluator 内で変換
type ExpandedMacroCallStatement struct {
	Name    string
	Params  []string
	Args    *ExpressionList
	Body    *BlockStatement
	Context *fileblock.Context
}

func (em *ExpandedMacroCallStatement) statementNode()           {}
func (em *ExpandedMacroCallStatement) NodeType() NodeType       { return NODE_MACRO_CALL_STMT }
func (em *ExpandedMacroCallStatement) NodeSubType() NodeSubType { return 0 }
func (em *ExpandedMacroCallStatement) String() string {
	args := []string{}
	for _, arg := range em.Args.Expressions {
		args = append(args, arg.String())
	}
	return fmt.Sprintf("EXPANDED MACRO %s CALL with %s", em.Name, strings.Join(args, ","))
}

// block statement
type BlockStatement struct {
	Block []Node
}

func (bs *BlockStatement) statementNode()           {}
func (bs *BlockStatement) NodeType() NodeType       { return NODE_BLOCK_STMT }
func (bs *BlockStatement) NodeSubType() NodeSubType { return 0 }
func (bs *BlockStatement) String() string {
	stmts := []string{}

	for _, s := range bs.Block {
		stmts = append(stmts, s.String())
	}
	return strings.Join(stmts, "\n")
}

// macro block statement
type MacroBlockStatement struct {
	Name  string
	Index int // REPT 用
	Count int // REPT 用
	Block []Node
}

func (mb *MacroBlockStatement) statementNode()           {}
func (mb *MacroBlockStatement) NodeType() NodeType       { return NODE_MACRO_BLOCK_STMT }
func (mb *MacroBlockStatement) NodeSubType() NodeSubType { return 0 }
func (mb *MacroBlockStatement) String() string           { return fmt.Sprintf("MACRO BLOCK(%s)", mb.Name) }

// 定数定義文 - CONST, EQU Statement
type ConstStatement struct {
	Name    Expression
	Value   Expression
	Context *fileblock.Context
}

func (cs *ConstStatement) statementNode()           {}
func (cs *ConstStatement) NodeType() NodeType       { return NODE_CONST_STMT }
func (cs *ConstStatement) NodeSubType() NodeSubType { return 0 }
func (cs *ConstStatement) String() string {
	var out bytes.Buffer

	out.WriteString("CONST ")
	out.WriteString(cs.Name.String())
	out.WriteString(" = ")
	out.WriteString(cs.Value.String())

	return out.String()
}

// 変数定義文 - VAR
type VariableStatement struct {
	Name    *Ident
	Value   Expression
	Context *fileblock.Context
}

func (vs *VariableStatement) statementNode()           {}
func (vs *VariableStatement) NodeType() NodeType       { return NODE_VAR_STMT }
func (vs *VariableStatement) NodeSubType() NodeSubType { return 0 }
func (vs *VariableStatement) String() string {
	var out bytes.Buffer

	out.WriteString("VAR ")
	out.WriteString(vs.Name.Name)
	out.WriteString(" = ")
	out.WriteString(vs.Value.String())

	return out.String()
}

// 変数代入文
type AssignStatement struct {
	Left    Expression
	Value   Expression
	Context *fileblock.Context
}

func (as *AssignStatement) statementNode()           {}
func (as *AssignStatement) NodeType() NodeType       { return NODE_ASSIGN_STMT }
func (as *AssignStatement) NodeSubType() NodeSubType { return 0 }
func (as *AssignStatement) String() string {
	var out bytes.Buffer

	out.WriteString(as.Left.String())
	out.WriteString(" = ")
	out.WriteString(as.Value.String())

	return out.String()
}

// Exitm 文
type ExitmStatement struct {
	Context *fileblock.Context
}

func (es *ExitmStatement) statementNode()           {}
func (es *ExitmStatement) NodeType() NodeType       { return NODE_EXITM_STMT }
func (ee *ExitmStatement) NodeSubType() NodeSubType { return 0 }
func (es *ExitmStatement) String() string           { return "EXITM" }

// Exitm 文
type ReturnStatement struct {
	Value   Expression
	Context *fileblock.Context
}

func (rs *ReturnStatement) statementNode()           {}
func (rs *ReturnStatement) NodeType() NodeType       { return NODE_RETURN_STMT }
func (rs *ReturnStatement) NodeSubType() NodeSubType { return 0 }
func (rs *ReturnStatement) String() string {
	s := "RETURN"
	if rs.Value != nil {
		s += " " + rs.Value.String()
	}
	return s
}

// Z80 命令文 - Z80Instruction Statement
type Z80Instruction struct {
	Label    *Label
	InstType int
	Opcode   int
	Op1      Expression
	Op2      Expression
	Context  *fileblock.Context
}

func (zi *Z80Instruction) statementNode() {}
func (zi *Z80Instruction) NodeType() NodeType {
	return NodeType(zi.InstType)
}
func (zi *Z80Instruction) NodeSubType() NodeSubType {
	return NodeSubType(zi.Opcode)
}
func (zi *Z80Instruction) String() string {
	var out bytes.Buffer

	if zi.Label != nil {
		out.WriteString(zi.Label.Name + ": ")
	}
	out.WriteString(Z80Opcode2Name(zi.Opcode))
	switch {
	case zi.Op1 == nil && zi.Op2 == nil:
		break
	case zi.Op1 != nil && zi.Op2 != nil:
		out.WriteString(" " + opString(zi.Op1))
		out.WriteString(", " + opString(zi.Op2))
	case zi.Op1 != nil:
		out.WriteString(" " + opString(zi.Op1))
	default:
		out.WriteString(" " + opString(zi.Op2))
	}

	return out.String()
}

// これ以降は 式 (Exspression)

// ラベル
type Label struct {
	LabelType NodeSubType
	Name      string
	Context   *fileblock.Context
}

func (le *Label) expressionNode()          {}
func (le *Label) NodeType() NodeType       { return NODE_LABEL }
func (le *Label) NodeSubType() NodeSubType { return le.LabelType }
func (le *Label) String() string           { return le.Name }

// 数値
type NumberLiteral struct {
	Value   int
	Context *fileblock.Context
}

func (nl *NumberLiteral) expressionNode()          {}
func (nl *NumberLiteral) NodeType() NodeType       { return NODE_NUMBER }
func (nl *NumberLiteral) NodeSubType() NodeSubType { return 0 }
func (nl *NumberLiteral) String() string {
	return fmt.Sprintf("%d", nl.Value)
}

// 文字列
type StringLiteral struct {
	Value   string
	Context *fileblock.Context
}

func (sl *StringLiteral) expressionNode()          {}
func (sl *StringLiteral) NodeType() NodeType       { return NODE_STRING }
func (sl *StringLiteral) NodeSubType() NodeSubType { return 0 }
func (sl *StringLiteral) String() string {
	return fmt.Sprintf("%q", sl.Value)
}

// 配列
type ArrayLiteral struct {
	Elements *ExpressionList
	Context  *fileblock.Context
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
	Left    Expression
	Index   Expression
	Context *fileblock.Context
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
	Context      *fileblock.Context
}

func (rl *RegisterLiteral) expressionNode()          {}
func (rl *RegisterLiteral) NodeType() NodeType       { return NodeType(rl.RegisterType) }
func (rl *RegisterLiteral) NodeSubType() NodeSubType { return NodeSubType(rl.Register) }
func (rl *RegisterLiteral) String() string {
	return Z80Opcode2Name(rl.Register)
}

// フラグ
type FlagLiteral struct {
	Flag    int
	Context *fileblock.Context
}

func (fl *FlagLiteral) expressionNode()          {}
func (fl *FlagLiteral) NodeType() NodeType       { return Z80_FLAG }
func (fl *FlagLiteral) NodeSubType() NodeSubType { return NodeSubType(fl.Flag) }
func (fl *FlagLiteral) String() string {
	return Z80Opcode2Name(fl.Flag)
}

// 識別子
type Ident struct {
	Name      string
	IdentType int
	Value     Expression
	Context   *fileblock.Context
}

func (i *Ident) expressionNode()          {}
func (i *Ident) NodeType() NodeType       { return NODE_IDENT }
func (i *Ident) NodeSubType() NodeSubType { return NodeSubType(i.IdentType) }
func (i *Ident) String() string           { return i.Name }

// ドット識別子
type DotIdent struct {
	Name    string
	Left    string
	Right   string
	Value   Expression
	Context *fileblock.Context
}

func (di *DotIdent) expressionNode()          {}
func (di *DotIdent) NodeType() NodeType       { return NODE_DOT_IDENT }
func (di *DotIdent) NodeSubType() NodeSubType { return 0 }
func (di *DotIdent) String() string           { return di.Left + "." + di.Right }

// 間接指定
type IndirectExpression struct {
	Expression Expression
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
	Operator int
	Op1      Expression
	Op2      Expression
	Context  *fileblock.Context
}

func (ie *InfixExpression) expressionNode()          {}
func (ie *InfixExpression) NodeType() NodeType       { return NODE_INFIX_EXPR }
func (ie *InfixExpression) NodeSubType() NodeSubType { return NodeSubType(ie.Operator) }
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
	out.WriteString(TokenLiteral(ie.Operator))
	out.WriteString(" " + op2 + ")")

	return out.String()
}

// 前置演算子式
type PrefixExpression struct {
	Operator int
	Op       Expression
	Context  *fileblock.Context
}

func (pe *PrefixExpression) expressionNode()          {}
func (pe *PrefixExpression) NodeType() NodeType       { return NodeType(NODE_PREFIX_EXPR) }
func (pe *PrefixExpression) NodeSubType() NodeSubType { return NodeSubType(pe.Operator) }
func (pe *PrefixExpression) String() string {
	var op string
	if pe.Op == nil {
		op = "<nil>"
	} else {
		op = pe.Op.String()
	}

	return "(" + TokenLiteral(pe.Operator) + op + ")"
}

// 関数呼出し
type FuncCallExpression struct {
	Name      string
	Arguments *ExpressionList
	Context   *fileblock.Context
}

func (fc *FuncCallExpression) expressionNode()          {}
func (fc *FuncCallExpression) NodeType() NodeType       { return NODE_CALL }
func (fc *FuncCallExpression) NodeSubType() NodeSubType { return 0 }
func (fc *FuncCallExpression) String() string {
	var out bytes.Buffer

	out.WriteString(fc.Name + "(")
	out.WriteString(fc.Arguments.String())
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
