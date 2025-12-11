package object

import (
	"fmt"
	"strings"
	"yas80/fileblock"
	"yas80/parser"
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
	SymState  SymbolState
	Node      parser.Node
	Value     Object
	DependsOn []string
	Context   *fileblock.Context
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

func NewLabelSymbol(name string, addr int, ctx *fileblock.Context) *SymbolObject {
	return &SymbolObject{Name: name,
		SymType:  SYM_LABEL,
		SymState: VALUE_TENTATIVE,
		Value:    &NumberObject{Value: addr, Context: ctx},
		Context:  ctx,
	}
}

func NewConstSymbol(name string, value Object, depends []string, ctx *fileblock.Context) *SymbolObject {
	return &SymbolObject{Name: name,
		SymType:   SYM_CONST,
		SymState:  VALUE_NULL,
		Value:     value,
		DependsOn: depends,
		Context:   ctx,
	}
}

func NewNullConstSymbol(name string, node parser.Node, depends []string, ctx *fileblock.Context) *SymbolObject {
	return &SymbolObject{Name: name,
		SymType:   SYM_CONST,
		SymState:  VALUE_NULL,
		Node:      node,
		Value:     NULL,
		DependsOn: depends,
		Context:   ctx,
	}
}

func NewUnknownSymbol(name, depend string, ctx *fileblock.Context) *SymbolObject {
	sym := &SymbolObject{Name: name,
		SymType:   SYM_UNKNOWN,
		SymState:  NOT_REGISTERED,
		DependsOn: []string{},
		Context:   ctx,
	}
	if depend != "" {
		sym.DependsOn = append(sym.DependsOn, depend)
	}
	return sym
}

// symbol expressoin
type SymbolExprObject struct {
	Names []string
}

func (s *SymbolExprObject) Type() ObjectType { return SYMBOL_EXPR_OBJ }
func (s *SymbolExprObject) String() string {
	return fmt.Sprintf("SYMBOL_EXPR{Names: [%s]}", strings.Join(s.Names, ", "))
}
