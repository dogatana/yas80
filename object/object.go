package object

import (
	"bytes"
	"fmt"
	"strings"
	"yas80/fileblock"
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
	DELETE_OBJ
	VALUE_OBJ
	EXITM_OBJ
)

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

// value - list ファイル出力用
type ValueObject struct {
	Value   Object
	Context *fileblock.Context
}

func (v *ValueObject) Type() ObjectType { return VALUE_OBJ }
func (v *ValueObject) String() string {
	return fmt.Sprintf("VALUE(%s)", v.Value.String())
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

// 右辺値で識別子が見つからない場合に使用
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

// exitm
type ExitmObject struct {
}

func (e *ExitmObject) Type() ObjectType { return EXITM_OBJ }
func (e *ExitmObject) String() string   { return "EXITM" }

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
	Value   int
	Context *fileblock.Context
}

func (n *NumberObject) Type() ObjectType { return NUMBER_OBJ }
func (n *NumberObject) String() string   { return fmt.Sprintf("%d(0x%x)", n.Value, n.Value) }

// 文字列
type StringObject struct {
	Value   string
	Context *fileblock.Context
}

func (s *StringObject) Type() ObjectType { return STRING_OBJ }
func (s *StringObject) String() string   { return fmt.Sprintf("%q", s.Value) }

// 識別子
type IdentObject struct {
	Name       string
	Value      Object
	LineNumber int
}

// Node
type NodeObject struct {
	Node parser.Node
}

func (n *NodeObject) Type() ObjectType { return NODE_OBJ }
func (n *NodeObject) String() string   { return n.Node.String() }

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
	// out.WriteString(" \\ " + mo.Body.String() + " \\ ENDM")

	return out.String()
}
