package evaluator

import (
	"fmt"
	"strings"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/intern"
	"github.com/dogatana/yas80/object"
)

func (e *Evaluator) CheckSymbolError(env TEnv) {
	for id, obj := range env.Store() {
		if obj == nil {
			continue
		}
		name := intern.SymbolID(id).String()
		switch obj := obj.(type) {
		case *object.SymbolObject:
			if name == "_" || name[0] == '$' {
				continue
			}
			if name[0] == '@' && env.EnvType() != object.ENV_MACRO {
				e.logger.Error(fmt.Sprintf(errcode.ESCOPE_MACRO, name), obj.Context)
				continue
			}
			if name[0] == '.' && object.OuterEnvType(env) != object.ENV_PROC {
				e.logger.Error(fmt.Sprintf(errcode.ESCOPE_PROC, name), obj.Context)
				continue
			}
			if obj.SymType == object.SYM_UNKNOWN && strings.Contains(name[1:], ".") {
				// dotident が仮登録されている場合で、本登録されていれば削除しておく
				if _, ok := e.getSymbolFromEnv(name, env); ok {
					env.Store()[id] = nil // map->slice に変更したので、削除は nil にする
					continue
				}
				e.logger.Error(fmt.Sprintf(errcode.ESYM_UNDEF, name), obj.Context)
				continue
			}
			if obj.SymType == object.SYM_UNKNOWN {
				e.logger.Error(fmt.Sprintf(errcode.ESYM_UNDEF, name), obj.Context)
				continue
			}
			if obj.Value == object.NULL {
				e.logger.Error(fmt.Sprintf(errcode.ESYM_NULL, name), obj.Context)
			}

		case *object.ProcObject:
			e.CheckSymbolError(obj)
		}
	}
}

func (e *Evaluator) CheckCyclicError(env TEnv) {
	visited := map[intern.SymbolID]bool{}
	visiting := map[intern.SymbolID]bool{}

	var visit func(sym *object.SymbolObject, id intern.SymbolID)
	visit = func(sym *object.SymbolObject, id intern.SymbolID) {
		if visiting[id] {
			e.logger.Error(fmt.Sprintf(errcode.ESYM_CYCLIC, id), sym.Context)
			return
		}
		if visited[id] {
			return
		}
		visiting[id] = true
		if obj, ok := env.Get(id); ok {
			if newSym, ok := obj.(*object.SymbolObject); ok {
				for _, dep := range newSym.DependsOn {
					visit(newSym, dep)
				}
			}
		}
		visited[id] = true
		visiting[id] = false
	}

	for id, obj := range env.Store() {
		sym, ok := obj.(*object.SymbolObject)
		if !ok || sym.Name != "_" && sym.Name[0] != '$' && sym.Value != object.NULL {
			continue
		}
		visit(sym, intern.SymbolID(id))
	}
}
