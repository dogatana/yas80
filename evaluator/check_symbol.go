package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
)

func (e *Evaluator) CheckSymbols(env object.Environment) {
	symbols := collectNullSymbols(env)

	_ = e.undefSymExists(symbols, env)
}

func collectNullSymbols(env object.Environment) []*object.SymbolObject {
	symbols := []*object.SymbolObject{}
	for _, v := range env.Store() {
		if sym, ok := v.(*object.SymbolObject); !ok {
			continue
		} else if sym.Value != object.NULL {
			continue
		} else {
			symbols = append(symbols, sym)
		}
	}
	return symbols
}

func (e *Evaluator) undefSymExists(syms []*object.SymbolObject, env object.Environment) bool {
	ret := false

	for _, sym := range syms {
		for _, name := range sym.DependsOn {
			if _, ok := env.Get(name); !ok {
				e.logger.Error(fmt.Sprintf(errcode.ESYM_NOT_FOUND, name), sym.Context)
				ret = true
			}
		}
	}
	return ret
}
