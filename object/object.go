package object

import (
	"fmt"
	"yas80/parser"
)

const (
	FIXEDCODE_OBJ = iota + 1
	CODE_OBJ
)

type ObjectType int

type Object interface {
	Type() ObjectType
	String() string
}

var objectTypeNames map[ObjectType]string = map[ObjectType]string{
	FIXEDCODE_OBJ: "FIXEDCODE_OBJ",
	CODE_OBJ:      "CODE_OBJ",
}

func (o ObjectType) String() string {
	name, ok := objectTypeNames[o]
	if ok {
		return name
	}
	return "UNKNOWN_OBJ"
}

// 固定コード
type FixedCode struct {
	Line int
	Code []byte
}

func (f *FixedCode) Type() ObjectType { return FIXEDCODE_OBJ }
func (f *FixedCode) String() string {
	text := fmt.Sprintf("%d: ", f.Line)
	for _, b := range f.Code {
		text += fmt.Sprintf("%02x", b)
	}
	return text
}

// Fixed 付きコード
type FixUp struct {
	Offset     int
	Expression parser.Expression
}

type Code struct {
	LineNumber int
	Code       []byte
	Fixups     []FixUp
}

func (c *Code) Type() ObjectType { return CODE_OBJ }
