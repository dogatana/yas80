package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
)

func (e *Evaluator) CheckSymbols(env object.Environment) {

	e.checkUnknwonAndNullSymbol(env)
	e.checkCyclic(env)
}

func (e *Evaluator) checkUnknwonAndNullSymbol(env object.Environment) {
	for name, obj := range env.Store() {
		if sym, ok := obj.(*object.SymbolObject); !ok {
			continue
		} else if sym.Name == "_" || sym.Name[0] == '$' {
			continue
		} else if sym.Name[0] == '@' && env.EnvType() != object.ENV_MACRO {
			e.logger.Error(fmt.Sprintf(errcode.ESCOPE_MACRO, name), sym.Context)
		} else if sym.Name[0] == '.' && object.OuterEnvType(env) != object.ENV_PROC {
			e.logger.Error(fmt.Sprintf(errcode.ESCOPE_PROC, name), sym.Context)
		} else if sym.SymType == object.SYM_UNKNOWN {
			e.logger.Error(fmt.Sprintf(errcode.ESYM_UNDEF, name), sym.Context)
		} else if sym.Value == object.NULL {
			e.logger.Error(fmt.Sprintf(errcode.ESYM_NOT_DETERMINED, name), sym.Context)
		}
	}
}

func (e *Evaluator) checkCyclic(env object.Environment) {
	visited := map[string]bool{}
	visiting := map[string]bool{}

	var visit func(sym *object.SymbolObject, name string)
	visit = func(sym *object.SymbolObject, name string) {
		if visiting[name] {
			e.logger.Error(fmt.Sprintf(errcode.ESYM_CYCLIC, name), sym.Context)
			return
		}
		if visited[name] {
			return
		}
		visiting[name] = true
		if obj, ok := env.Get(name); ok {
			if newSym, ok := obj.(*object.SymbolObject); ok {
				for _, dep := range newSym.DependsOn {
					visit(newSym, dep)
				}
			}
		}
		visited[name] = true
		visiting[name] = false
	}

	for name, obj := range env.Store() {
		sym, ok := obj.(*object.SymbolObject)
		if !ok || sym.Name != "_" && sym.Name[0] != '$' && sym.Value != object.NULL {
			continue
		}
		visit(sym, name)
	}
}
