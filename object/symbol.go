package object

import (
	"fmt"
	"strings"

	"github.com/dogatana/yas80/filecontent"
	"github.com/dogatana/yas80/intern"
	"github.com/dogatana/yas80/internal/util"
	"github.com/dogatana/yas80/parser"
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
	NameID    intern.SymbolID
	SymType   SymbolType
	Node      parser.Node
	Value     Object
	DependsOn []intern.SymbolID
	Context   *filecontent.Context
}

func (s *SymbolObject) Type() ObjectType { return OBJ_SYMBOL }
func (s *SymbolObject) String() string {
	str := fmt.Sprintf("Symbol{%q, %s, %s",
		s.Name, symbolTypeNames[s.SymType], s.Value.String())
	if len(s.DependsOn) > 0 {
		deps := util.Map(s.DependsOn, intern.Lookup)
		str += ", [" + strings.Join(deps, ",") + "]"
	}
	return str + "}"
}

func NewLabelSymbol(id intern.SymbolID, name string, addr int, ctx *filecontent.Context) *SymbolObject {
	return &SymbolObject{
		NameID:  id,
		Name:    name,
		SymType: SYM_LABEL,
		Value:   &NumberObject{Value: addr, Context: ctx},
		Context: ctx,
	}
}

func NewConstSymbol(id intern.SymbolID, name string, node parser.Node, value Object, depends []intern.SymbolID, ctx *filecontent.Context) *SymbolObject {
	return &SymbolObject{
		NameID:    id,
		Name:      name,
		SymType:   SYM_CONST,
		Node:      node,
		Value:     value,
		DependsOn: depends,
		Context:   ctx,
	}
}

func NewVarSymbol(id intern.SymbolID, name string, node parser.Node, value Object, depends []intern.SymbolID, ctx *filecontent.Context) *SymbolObject {
	return &SymbolObject{Name: name,
		NameID:    id,
		SymType:   SYM_VAR,
		Node:      node,
		Value:     value,
		DependsOn: depends,
		Context:   ctx,
	}
}

func NewUnknownSymbol(id intern.SymbolID, name string, ctx *filecontent.Context) *SymbolObject {
	return &SymbolObject{
		NameID:    id,
		Name:      name,
		SymType:   SYM_UNKNOWN,
		Value:     NULL,
		DependsOn: []intern.SymbolID{},
		Context:   ctx,
	}
}
