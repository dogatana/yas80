package object

import (
	"fmt"
	"strings"
	"yas80/fileblock"
	"yas80/parser"
)

type SymbolType int

const (
	SYM_UNKNOWN SymbolType = iota
	SYM_CONST
	SYM_LABEL
	SYM_VAR
)

var symbolTypeNames map[SymbolType]string = map[SymbolType]string{
	SYM_UNKNOWN: "Unknown",
	SYM_CONST:   "Const",
	SYM_LABEL:   "Label",
	SYM_VAR:     "Var",
}

// symbol
type SymbolObject struct {
	Name      string
	SymType   SymbolType
	Node      parser.Node
	Value     Object
	DependsOn []string
	Context   *fileblock.Context
}

func (s *SymbolObject) Type() ObjectType { return SYMBOL_OBJ }
func (s *SymbolObject) String() string {
	str := fmt.Sprintf("Symbol{Name:%q, SymType: %s, Value: %T",
		s.Name, symbolTypeNames[s.SymType], s.Value)
	if len(s.DependsOn) > 0 {
		str += ", [" + strings.Join(s.DependsOn, ",") + "]"
	}
	return str + "}"
}

func NewLabelSymbol(name string, addr int, ctx *fileblock.Context) *SymbolObject {
	return &SymbolObject{Name: name,
		SymType: SYM_LABEL,
		Value:   &NumberObject{Value: addr, Context: ctx},
		Context: ctx,
	}
}

func NewConstSymbol(name string, value Object, depends []string, ctx *fileblock.Context) *SymbolObject {
	return &SymbolObject{Name: name,
		SymType:   SYM_CONST,
		Value:     value,
		DependsOn: depends,
		Context:   ctx,
	}
}

func NewNullConstSymbol(name string, node parser.Node, depends []string, ctx *fileblock.Context) *SymbolObject {
	return &SymbolObject{Name: name,
		SymType:   SYM_CONST,
		Node:      node,
		Value:     NULL,
		DependsOn: depends,
		Context:   ctx,
	}
}

func NewUnknownSymbol(name, depend string, ctx *fileblock.Context) *SymbolObject {
	sym := &SymbolObject{Name: name,
		SymType:   SYM_UNKNOWN,
		DependsOn: []string{},
		Context:   ctx,
	}
	if depend != "" {
		sym.DependsOn = append(sym.DependsOn, depend)
	}
	return sym
}
