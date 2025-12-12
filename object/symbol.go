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
	SYM_UNKNOWN: "SYM_UNKNOWN",
	SYM_CONST:   "SYM_CONST",
	SYM_LABEL:   "SYM_LABEL",
	SYM_VAR:     "SYM_VAR",
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
	str := fmt.Sprintf("Symbol{%q, %s, %s",
		s.Name, symbolTypeNames[s.SymType], s.Value.String())
	if len(s.DependsOn) > 0 {
		str += ", [" + strings.Join(s.DependsOn, ",") + "]"
	}
	return str + "}"
}

func NewLabelSymbol(name string, addr int, ctx *fileblock.Context) *SymbolObject {
	return &SymbolObject{
		Name:    name,
		SymType: SYM_LABEL,
		Value:   &NumberObject{Value: addr, Context: ctx},
		Context: ctx,
	}
}

func NewConstSymbol(name string, node parser.Node, value Object, depends []string, ctx *fileblock.Context) *SymbolObject {
	return &SymbolObject{Name: name,
		SymType:   SYM_CONST,
		Node:      node,
		Value:     value,
		DependsOn: depends,
		Context:   ctx,
	}
}

func NewUnknownSymbol(name, depend string, ctx *fileblock.Context) *SymbolObject {
	return &SymbolObject{
		Name:      name,
		SymType:   SYM_UNKNOWN,
		Value:     NULL,
		DependsOn: []string{},
		Context:   ctx,
	}
}
