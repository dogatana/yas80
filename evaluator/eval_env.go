package evaluator

import (
	"fmt"
	"yas80/logger"
	"yas80/object"
)

func (e *Evaluator) EvalEnv(env *object.Environment) ([]string, error) {
	order, err := e.tSortEnv(env)
	if err != nil {
		return order, err
	}
	for _, name := range order {
		obj, ok := env.Get(name)
		if !ok {
			return order, fmt.Errorf(logger.E900, fmt.Sprintf(": could not get %s", name))
		}
		sym, ok := obj.(*object.SymbolObject)
		if !ok {
			continue
		}
		if sym.State == object.SYMBOL_STATE_DEFINED {
			env.Set(name, sym.Value)
			continue
		}
		value := e.Eval(sym.Node, env)
		if isError(value) || isRefNotFound(value) {
			return order, fmt.Errorf(logger.E900, fmt.Sprintf(": could not eval '%s'", name))
		}
		env.Set(name, value)
	}
	return order, nil
}

// 環境をトポロジカルソート
func (e *Evaluator) tSortEnv(env *object.Environment) ([]string, error) {
	visited := map[string]bool{}
	visiting := map[string]bool{}
	order := []string{}

	var visit func(string) error
	visit = func(name string) error {
		if visiting[name] {
			return fmt.Errorf(logger.E030, name)
		}
		if visited[name] {
			return nil
		}
		visiting[name] = true
		obj, ok := env.Get(name)
		if !ok {
			return fmt.Errorf(logger.E009, name)
		}
		sym, ok := obj.(*object.SymbolObject)
		if ok {
			for _, dep := range sym.DependsOn {
				if err := visit(dep); err != nil {
					return err
				}
			}
		}
		visited[name] = true
		visiting[name] = false
		order = append(order, name)
		return nil
	}

	for name := range env.Store {
		if err := visit(name); err != nil {
			return nil, err
		}
	}
	return order, nil
}
