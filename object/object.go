package object

import (
	"bytes"
	"fmt"
	"strings"
	"yas80/parser"
)

const (
	NULL_OBJ = iota + 1
	ERROR_OBJ
	NUMBER_OBJ
	STRING_OBJ
	IDENT_OBJ
	ENUM_OBJ
	REGISTER_OBJ
	CODE_OBJ
	PROGRAM_OBJ
	NODE_OBJ
	RETURN_OBJ
	BLOCK_OBJ
	FUNC_OBJ
	MACRO_OBJ
	REF_NOTFOUND_OBJ
	SYMBOL_OBJ
	SYMBOL_EXPR_OBJ
	DELETE_OBJ
)

type SymbolState int

const (
	NOT_REGISTERED SymbolState = -1 + iota
	VALUE_NULL
	VALUE_TENTATIVE
	VALUE_DETERMINED
)

var symbolStateNames map[SymbolState]string = map[SymbolState]string{
	NOT_REGISTERED:   "NotRegistered",
	VALUE_NULL:       "NullValue",
	VALUE_TENTATIVE:  "TentativeValue",
	VALUE_DETERMINED: "Determined",
}

type SymbolType int

const (
	UNKNOWN SymbolType = iota
	CONST
	LABEL
	VAR
)

var symbolTypeNames map[SymbolType]string = map[SymbolType]string{
	UNKNOWN: "Unknown",
	CONST:   "Const",
	LABEL:   "Label",
	VAR:     "Var",
}

// 同一判定のため定数的に定義しておく
var (
	NULL  = &NullObject{}
	ERROR = &ErrorObject{}
)

type ObjectType int

type Object interface {
	Type() ObjectType
	String() string
}

var objectTypeNames map[ObjectType]string = map[ObjectType]string{
	CODE_OBJ: "CODE_OBJ",
	NODE_OBJ: "NODE_OBJ",
}

func (o ObjectType) String() string {
	name, ok := objectTypeNames[o]
	if ok {
		return name
	}
	return "UNKNOWN_OBJ"
}

// program
type ProgramObject struct {
	Objects []Object
}

func (p *ProgramObject) Type() ObjectType { return PROGRAM_OBJ }
func (p *ProgramObject) String() string {
	results := []string{}

	for _, result := range p.Objects {
		results = append(results, result.String())
	}
	return strings.Join(results, "\n")
}

// code
type CodeObject struct {
	Line int
	Addr int
	Code []byte
}

func (f *CodeObject) Type() ObjectType { return CODE_OBJ }
func (f *CodeObject) String() string {
	text := fmt.Sprintf("%d:%04x: ", f.Line, f.Addr)
	for _, b := range f.Code {
		text += fmt.Sprintf("%02x ", b)
	}
	return text
}
func (f *CodeObject) Size() int {
	return len(f.Code)
}

// NULL
type NullObject struct{}

func (n *NullObject) Type() ObjectType { return NULL_OBJ }
func (n *NullObject) String() string   { return "NULL" }

// Error
type ErrorObject struct {
}

func (e *ErrorObject) Type() ObjectType { return ERROR_OBJ }
func (e *ErrorObject) String() string   { return "ERROR" }

// pass1 で右辺式で識別が見つからない場合に使用
type RefNotFoundObject struct {
	Names []string
}

func (r *RefNotFoundObject) Type() ObjectType { return REF_NOTFOUND_OBJ }
func (r *RefNotFoundObject) String() string {
	var out bytes.Buffer

	out.WriteString("REF_NOTFOUND(")
	out.WriteString(strings.Join(r.Names, ", "))
	out.WriteRune(')')
	return out.String()
}

// deleted
type DeleltedObject struct {
	Node parser.Node
}

func (d *DeleltedObject) Type() ObjectType { return DELETE_OBJ }
func (d *DeleltedObject) String() string {
	body := strings.Join(strings.Split(d.Node.String(), "\n"), " \\ ")

	return "DELETED(" + body + ")"
}

// symbol
type SymbolObject struct {
	Name       string
	SymType    SymbolType
	SymState   SymbolState
	Node       parser.Node
	Value      Object
	DependsOn  []string
	LineNumber int
}

func (s *SymbolObject) Type() ObjectType { return SYMBOL_OBJ }
func (s *SymbolObject) String() string {
	str := fmt.Sprintf("Symbol{Name:%q, SymType: %s, SymState: %s, Value: %T",
		s.Name, symbolTypeNames[s.SymType], symbolStateNames[s.SymState], s.Value)
	if len(s.DependsOn) > 0 {
		str += ", [" + strings.Join(s.DependsOn, ",") + "]"
	}
	return str + "}"
}

func NewLabelSymbol(name string, addr int, lineNumber int) *SymbolObject {
	return &SymbolObject{Name: name,
		SymType:    LABEL,
		SymState:   VALUE_TENTATIVE,
		Value:      &NumberObject{Value: addr, LineNumber: lineNumber},
		LineNumber: lineNumber,
	}
}

func NewConstSymbol(name string, node parser.Node, value Object, depends []string, lineNumber int) *SymbolObject {
	return &SymbolObject{Name: name,
		SymType:    CONST,
		SymState:   VALUE_NULL,
		Node:       node,
		Value:      value,
		DependsOn:  depends,
		LineNumber: lineNumber,
	}
}

func NewNullConstSymbol(name string, node parser.Node, depends []string, lineNumber int) *SymbolObject {
	return &SymbolObject{Name: name,
		SymType:    CONST,
		SymState:   VALUE_NULL,
		Node:       node,
		Value:      NULL,
		DependsOn:  depends,
		LineNumber: lineNumber,
	}
}

func NewUnknownSymbol(name, depend string, lineNumber int) *SymbolObject {
	sym := &SymbolObject{Name: name,
		SymType:    UNKNOWN,
		SymState:   NOT_REGISTERED,
		DependsOn:  []string{},
		LineNumber: lineNumber,
	}
	if depend != "" {
		sym.DependsOn = append(sym.DependsOn, depend)
	}
	return sym
}

// symbol expressoin
type SymbolExprObject struct {
	Names      []string
	LineNumber int
}

func (s *SymbolExprObject) Type() ObjectType { return SYMBOL_EXPR_OBJ }
func (s *SymbolExprObject) String() string {
	return fmt.Sprintf("SYMBOL_EXPR{Names: [%s]}", strings.Join(s.Names, ", "))
}

// return
type ReturnObject struct {
	Value      Object
	LineNumber int
}

func (r *ReturnObject) Type() ObjectType { return RETURN_OBJ }
func (r *ReturnObject) String() string {
	if r.Value == NULL {
		return "RETURN"
	}
	return "RETURN " + r.Value.String()
}

// 数値
type NumberObject struct {
	Value      int
	LineNumber int
}

func (n *NumberObject) Type() ObjectType { return NUMBER_OBJ }
func (n *NumberObject) String() string   { return fmt.Sprintf("%d(0x%x)", n.Value, n.Value) }

// 文字列
type StringObject struct {
	Value      string
	LineNumber int
}

func (s *StringObject) Type() ObjectType { return STRING_OBJ }
func (s *StringObject) String() string   { return fmt.Sprintf("%q", s.Value) }

// 識別子
type IdentObject struct {
	Name       string
	Value      Object
	LineNumber int
}

// レジスタ
type RegisterObject struct {
	RegisterType int
	Register     int
}

func (r *RegisterObject) Type() ObjectType { return REGISTER_OBJ }
func (r *RegisterObject) String() string   { return parser.Z80Opcode2Name(r.Register) }

// Node
type NodeObject struct {
	Value      parser.Node
	LineNumber int
}

func (n *NodeObject) Type() ObjectType { return NODE_OBJ }
func (n *NodeObject) String() string   { return n.Value.String() }

// ENUM
type EnumObject struct {
	Name  string
	Value map[string]Object
	Keys  []string
}

func (e *EnumObject) Type() ObjectType { return ENUM_OBJ }
func (e *EnumObject) String() string {

	stmts := []string{"ENUM " + e.Name}

	for k, v := range e.Value {
		s := fmt.Sprintf("%s = %s", k, v.String())
		stmts = append(stmts, s)
	}
	stmts = append(stmts, "END_ENUM")

	return strings.Join(stmts, "\n")
}

func (e *EnumObject) Get(key string) (Object, bool) {
	v, ok := e.Value[key]
	return v, ok
}

// Block
type BlockObject struct {
	Block []Object
}

func (b *BlockObject) Type() ObjectType { return BLOCK_OBJ }
func (b *BlockObject) String() string {
	strs := []string{}

	for _, o := range b.Block {
		strs = append(strs, o.String())
	}
	return strings.Join(strs, "\n")
}

// Function
type FunctionObject struct {
	Name   string
	Params []string
	Body   parser.Node
	Env    Environment
}

func (f *FunctionObject) Type() ObjectType { return FUNC_OBJ }
func (f *FunctionObject) String() string {
	var out bytes.Buffer

	out.WriteString(f.Name + " FUNC")
	if len(f.Params) > 0 {
		out.WriteRune(' ')
		out.WriteString(strings.Join(f.Params, ", "))
	}
	out.WriteRune('\n')
	out.WriteString(f.Body.String() + "\n")
	out.WriteString("ENDF")

	return out.String()
}

// Macro
type MacroObject struct {
	Name   string
	Params []string
	Body   *parser.BlockStatement
}

func (mo *MacroObject) Type() ObjectType { return MACRO_OBJ }
func (mo *MacroObject) String() string {
	var out bytes.Buffer

	out.WriteString(mo.Name + " MACRO")
	if len(mo.Params) > 0 {
		out.WriteRune(' ')
		out.WriteString(strings.Join(mo.Params, ", "))
	}
	out.WriteString(" \\ " + mo.Body.String() + " \\ ENDM")

	return out.String()
}
