package object

import (
	"bytes"
	"fmt"
	"strings"
	"yas80/fileblock"
	"yas80/parser"
)

const (
	OBJ_NULL = iota + 1
	OBJ_ERROR
	OBJ_NUMBER
	OBJ_STRING
	OBJ_ENUM
	OBJ_REGISTER
	OBJ_REG_INDIRECT
	OBJ_ADDR_INDIRECT
	OBJ_CODE
	OBJ_PROGRAM
	OBJ_NODE
	OBJ_NODES
	OBJ_RETURN
	OBJ_BLOCK
	OBJ_PROC
	OBJ_FUNC
	OBJ_MACRO
	OBJ_REF_NOTFOUND
	OBJ_SYMBOL
	OBJ_DELETE
	OBJ_EXITM
	OBJ_VALUE
	OBJ_COMMENT
	OBJ_ARRAY
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
	OBJ_CODE: "CODE_OBJ",
	OBJ_NODE: "NODE_OBJ",
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

func (p *ProgramObject) Type() ObjectType { return OBJ_PROGRAM }
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

func (v *ValueObject) Type() ObjectType { return OBJ_VALUE }
func (v *ValueObject) String() string {
	return fmt.Sprintf("VALUE(%s)", v.Value.String())
}

// value - list ファイル出力用
type CommentObject struct {
	Comments []string
	Context  *fileblock.Context
}

func (o *CommentObject) Type() ObjectType { return OBJ_COMMENT }
func (o *CommentObject) String() string {
	var out bytes.Buffer

	out.WriteString("comt ")
	if o.Context == nil {
		out.WriteString("--:--(  ) ")
	} else {
		out.WriteString(fmt.Sprintf("%2d:%2d", o.Context.Line, o.Context.Offset))
		if o.Context.Source == nil {
			out.WriteString("(  ) ")
		} else {
			out.WriteString(fmt.Sprintf("(%2d) ", o.Context.Source.Line))
		}
	}
	out.WriteString(strings.Join(o.Comments, "\n"))
	return out.String()
}

// code
type CodeObject struct {
	Addr    int
	Code    []byte
	CZ80    int
	CR800   int
	Context *fileblock.Context
}

func (f *CodeObject) Type() ObjectType { return OBJ_CODE }
func (f *CodeObject) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("code %2d:%2d", f.Context.Line, f.Context.Offset))
	if f.Context.Source == nil {
		out.WriteString("(  )")
	} else {
		out.WriteString(fmt.Sprintf("(%2d)", f.Context.Source.Line))
	}
	out.WriteString(fmt.Sprintf(" :%04x: ", f.Addr))
	for _, b := range f.Code {
		out.WriteString(fmt.Sprintf("%02x ", b))
	}
	return out.String()
}
func (f *CodeObject) Size() int {
	return len(f.Code)
}

// NULL
type NullObject struct{}

func (n *NullObject) Type() ObjectType { return OBJ_NULL }
func (n *NullObject) String() string   { return "NULL" }

// Error
type ErrorObject struct {
}

func (e *ErrorObject) Type() ObjectType { return OBJ_ERROR }
func (e *ErrorObject) String() string   { return "ERROR" }

// 右辺値で識別子が見つからない場合に使用
type RefNotFoundObject struct {
	Names []string
}

func (r *RefNotFoundObject) Type() ObjectType { return OBJ_REF_NOTFOUND }
func (r *RefNotFoundObject) String() string {
	var out bytes.Buffer

	out.WriteString("REF_NOTFOUND(")
	out.WriteString(strings.Join(r.Names, ", "))
	out.WriteRune(')')
	return out.String()
}

// proc
type ProcObject struct {
	Name string
	Addr int
	Env  Environment
}

// Object の実装
func (o *ProcObject) Type() ObjectType { return OBJ_PROC }
func (o *ProcObject) String() string   { return fmt.Sprintf("PROC(%s[0x%04x])", o.Name, o.Addr) }

// Environ interface の実装
func (o *ProcObject) EnvType() int                       { return o.Env.EnvType() }
func (o *ProcObject) Get(name string) (Object, bool)     { return o.Env.Get(name) }
func (o *ProcObject) Set(name string, obj Object) Object { return o.Env.Set(name, obj) }
func (o *ProcObject) Outer() Environment                 { return o.Env.Outer() }
func (o *ProcObject) Store() map[string]Object           { return o.Env.Store() }

// enum
type EnumObject struct {
	Name string
	Env  Environment
}

// Object の実装
func (o *EnumObject) Type() ObjectType { return OBJ_ENUM }
func (o *EnumObject) String() string {
	var out bytes.Buffer

	out.WriteString(fmt.Sprintf("ENUM(%s) {\n", o.Name))
	for k, v := range o.Env.Store() {
		out.WriteString(fmt.Sprintf("%s=%s\n", k, v.String()))
	}
	out.WriteString("}")

	return out.String()
}

// Environ interface の実装
func (o *EnumObject) EnvType() int                       { return o.Env.EnvType() }
func (o *EnumObject) Get(name string) (Object, bool)     { return o.Env.Get(name) }
func (o *EnumObject) Set(name string, obj Object) Object { return o.Env.Set(name, obj) }
func (o *EnumObject) Outer() Environment                 { return o.Env.Outer() }
func (o *EnumObject) Store() map[string]Object           { return o.Env.Store() }

// deleted
type DeleltedObject struct {
	Node parser.Node
}

func (d *DeleltedObject) Type() ObjectType { return OBJ_DELETE }
func (d *DeleltedObject) String() string {
	body := strings.Join(strings.Split(d.Node.String(), "\n"), " \\ ")

	return "DELETED(" + body + ")"
}

// exitm
type ExitmObject struct {
}

func (e *ExitmObject) Type() ObjectType { return OBJ_EXITM }
func (e *ExitmObject) String() string   { return "EXITM" }

// return
type ReturnObject struct {
	Value      Object
	LineNumber int
}

func (r *ReturnObject) Type() ObjectType { return OBJ_RETURN }
func (r *ReturnObject) String() string {
	if r.Value == NULL {
		return "RETURN"
	}
	return "RETURN " + r.Value.String()
}

// 数値
type NumberObject struct {
	Value     int
	ForceWord bool // true: word 強制
	Context   *fileblock.Context
}

func (n *NumberObject) Type() ObjectType { return OBJ_NUMBER }
func (n *NumberObject) String() string   { return fmt.Sprintf("%d(0x%x)", n.Value, n.Value) }

// 文字列
type StringObject struct {
	Value   string
	Context *fileblock.Context
}

func (s *StringObject) Type() ObjectType { return OBJ_STRING }
func (s *StringObject) String() string   { return fmt.Sprintf("%q", s.Value) }

// レジスタ間接
type RegIndirectObject struct {
	Register     int
	Displacement int
}

func (o *RegIndirectObject) Type() ObjectType { return OBJ_REG_INDIRECT }
func (o *RegIndirectObject) String() string {
	if o.Displacement != 0 {
		return fmt.Sprintf("(%s%+d)", parser.TokenLiteral(o.Register), o.Displacement)
	} else {
		return fmt.Sprintf("(%s)", parser.TokenLiteral(o.Register))
	}
}

// アドレス間接
type AddrIndirectObject struct {
	Address int
}

func (o *AddrIndirectObject) Type() ObjectType { return OBJ_ADDR_INDIRECT }
func (o *AddrIndirectObject) String() string {
	return fmt.Sprintf("($%x)", o.Address)
}

// Node
type StatementObject struct {
	Statement parser.Statement
}

func (n *StatementObject) Type() ObjectType { return OBJ_NODE }
func (n *StatementObject) String() string   { return n.Statement.String() }

// Nodes
type StatemetnsObject struct {
	Statements []parser.Statement
}

func (n *StatemetnsObject) Type() ObjectType { return OBJ_NODES }
func (n *StatemetnsObject) String() string   { return fmt.Sprintf("NODES(%v)", n.Statements) }

// // ENUM
// type EnumObject struct {
// 	Name  string
// 	Value map[string]Object
// 	Keys  []string
// }

// func (e *EnumObject) Type() ObjectType { return ENUM_OBJ }
// func (e *EnumObject) String() string {

// 	stmts := []string{"ENUM " + e.Name}

// 	for k, v := range e.Value {
// 		s := fmt.Sprintf("%s = %s", k, v.String())
// 		stmts = append(stmts, s)
// 	}
// 	stmts = append(stmts, "END_ENUM")

// 	return strings.Join(stmts, "\n")
// }

// func (e *EnumObject) Get(key string) (Object, bool) {
// 	v, ok := e.Value[key]
// 	return v, ok
// }

// Block
type BlockObject struct {
	Block []Object
}

func (b *BlockObject) Type() ObjectType { return OBJ_BLOCK }
func (b *BlockObject) String() string {
	strs := []string{}

	for _, o := range b.Block {
		strs = append(strs, o.String())
	}
	return strings.Join(strs, "\n")
}

// Function
type FunctionObject struct {
	Name    string
	Params  []string
	Body    parser.Node
	Env     Environment
	Context *fileblock.Context
}

func (f *FunctionObject) Type() ObjectType { return OBJ_FUNC }
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

func (mo *MacroObject) Type() ObjectType { return OBJ_MACRO }
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

// 配列
type ArrayObject struct {
	Values      []Object
	Expressions []parser.Expression
}

func (o *ArrayObject) Type() ObjectType { return OBJ_ARRAY }
func (o *ArrayObject) String() string {

	values := []string{}
	for _, v := range o.Values {
		values = append(values, v.String())
	}

	var out bytes.Buffer
	out.WriteRune('[')
	out.WriteString(strings.Join(values, ","))
	out.WriteRune(']')

	return out.String()
}
