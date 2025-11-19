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
	REF_NOTFOUND_OBJ
	SYMBOL_OBJ
	SYMBOL_EXPR_OBJ
	DELETE_OBJ
)

type SymbolState int

const (
	SYMBOL_STATE_UNDEFINED SymbolState = iota
	SYMBOL_STATE_TENTATIVE
	SYMBOL_STATE_DEFINED
)

var (
	NULL  = &NullObject{}
	ERROR = &ErrorObject{}

	Z80_REG_A   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_A}
	Z80_REG_B   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_B}
	Z80_REG_C   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_C}
	Z80_REG_D   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_D}
	Z80_REG_E   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_E}
	Z80_REG_H   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_H}
	Z80_REG_L   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_L}
	Z80_REG_IXH = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_IXH}
	Z80_REG_IXL = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_IXL}
	Z80_REG_IYH = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_IYH}
	Z80_REG_IYL = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_IYL}
	Z80_REG_I   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_I}
	Z80_REG_R   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_R}

	Z80_REG_SP   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_SP}
	Z80_REG_IX   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_IX}
	Z80_REG_IY   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_IY}
	Z80_REG_AF   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_AF}
	Z80_REG_AFEX = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_AFEX}
	Z80_REG_BC   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_BC}
	Z80_REG_DE   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_DE}
	Z80_REG_HL   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_HL}

	Z80_FLAG_C  = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_C}
	Z80_FLAG_NC = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_NC}
	Z80_FLAG_Z  = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_Z}
	Z80_FLAG_NZ = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_NZ}
	Z80_FLAG_PO = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_PO}
	Z80_FLAG_PE = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_PE}
	Z80_FLAG_P  = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_P}
	Z80_FLAG_M  = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_M}
)

var Z80RgisterObjects map[int]Object = map[int]Object{
	parser.Z80_REG_A:   Z80_REG_A,
	parser.Z80_REG_B:   Z80_REG_B,
	parser.Z80_REG_C:   Z80_REG_C,
	parser.Z80_REG_D:   Z80_REG_D,
	parser.Z80_REG_E:   Z80_REG_E,
	parser.Z80_REG_H:   Z80_REG_H,
	parser.Z80_REG_L:   Z80_REG_L,
	parser.Z80_REG_IXH: Z80_REG_IXH,
	parser.Z80_REG_IXL: Z80_REG_IXL,
	parser.Z80_REG_IYH: Z80_REG_IYH,
	parser.Z80_REG_IYL: Z80_REG_IYL,
	parser.Z80_REG_I:   Z80_REG_IYL,
	parser.Z80_REG_R:   Z80_REG_IYL,

	parser.Z80_REG_SP:   Z80_REG_SP,
	parser.Z80_REG_IX:   Z80_REG_IX,
	parser.Z80_REG_IY:   Z80_REG_IY,
	parser.Z80_REG_AF:   Z80_REG_AF,
	parser.Z80_REG_AFEX: Z80_REG_AFEX,
	parser.Z80_REG_BC:   Z80_REG_BC,
	parser.Z80_REG_DE:   Z80_REG_DE,
	parser.Z80_REG_HL:   Z80_REG_HL,

	parser.Z80_FLAG_NC: Z80_FLAG_NC,
	parser.Z80_FLAG_Z:  Z80_FLAG_Z,
	parser.Z80_FLAG_NZ: Z80_FLAG_NZ,
	parser.Z80_FLAG_PO: Z80_FLAG_PO,
	parser.Z80_FLAG_PE: Z80_FLAG_PE,
	parser.Z80_FLAG_P:  Z80_FLAG_P,
	parser.Z80_FLAG_M:  Z80_FLAG_M,
}

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
	return "DELETED(" + d.Node.String() + ")"
}

// symbol
type SymbolObject struct {
	Name       string
	Node       parser.Node
	Value      Object
	State      SymbolState
	DependsOn  []string
	LineNumber int
}

func (s *SymbolObject) Type() ObjectType { return SYMBOL_OBJ }
func (s *SymbolObject) String() string {
	return fmt.Sprintf("SYMBOL{Name: %s, Node: %s, Value: %v, State: %d, Depends: [%s]}",
		s.Name, s.Node.String(), s.Value, s.State, strings.Join(s.DependsOn, ", "))
}

// symbol expressoin
type SymbolExprObject struct {
	Value      Object
	Names      []string
	LineNumber int
}

func (s *SymbolExprObject) Type() ObjectType { return SYMBOL_EXPR_OBJ }
func (s *SymbolExprObject) String() string {
	return fmt.Sprintf("SYMBOL_EXPR{Names: %s, Value: %s}",
		strings.Join(s.Names, ", "), s.Value.String())
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
	Env    *Environment
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
