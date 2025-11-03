package object

import (
	"fmt"
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
type Program struct {
	Objects []Object
}

func (p *Program) Type() ObjectType { return PROGRAM_OBJ }
func (p *Program) String() string   { return "PROGRAM_OBJ" }

// code
type Code struct {
	Line int
	Code []byte
}

func (f *Code) Type() ObjectType { return CODE_OBJ }
func (f *Code) String() string {
	text := fmt.Sprintf("%d: ", f.Line)
	for _, b := range f.Code {
		text += fmt.Sprintf("%02x", b)
	}
	return text
}

// NULL
type NullObject struct{}

func (n *NullObject) Type() ObjectType { return NULL_OBJ }
func (n *NullObject) String() string   { return "NULL" }

// Error
type ErrorObject struct {
	Message string
}

func (e *ErrorObject) Type() ObjectType { return ERROR_OBJ }
func (e *ErrorObject) String() string   { return "ERROR " + e.Message }

// 数値
type NumberObject struct {
	Value int
}

func (n *NumberObject) Type() ObjectType { return NUMBER_OBJ }
func (n *NumberObject) String() string   { return fmt.Sprintf("%d", n.Value) }

// 文字列
type StringObject struct {
	Value string
}

func (s *StringObject) Type() ObjectType { return STRING_OBJ }
func (s *StringObject) String() string   { return s.Value }

// 識別子
type IdentObject struct {
	Name  string
	Value Object
}

// レジスタ
type RegisterObject struct {
	RegisterType int
	Register     int
}

func (r *RegisterObject) Type() ObjectType { return REGISTER_OBJ }
func (r *RegisterObject) String() string   { return parser.Z80OpCode2Name(r.Register) }

// Node
type NodeObject struct {
	Value parser.Node
}

func (n *NodeObject) Type() ObjectType { return NODE_OBJ }
func (n *NodeObject) String() string   { return n.Value.String() }

// ENUM
type EnumObject struct {
	Name  string
	Value map[string]Object
}

func (e *EnumObject) Type() ObjectType { return ENUM_OBJ }
func (e *EnumObject) String() string   { return "ENUM " + e.Name }
func (e *EnumObject) Get(key string) (Object, bool) {
	v, ok := e.Value[key]
	return v, ok
}
