package evaluator

import (
	"yas80/object"
)

func (e *Evaluator) EvalEnv(env object.Environment) ([]string, error) {
	order, err := e.tSortEnv(env)
	if err != nil {
		return order, err
	}
	for _, name := range order {
		obj, ok := env.Get(name)
		if !ok {
			continue
		}
		sym, ok := obj.(*object.SymbolObject)
		if !ok {
			continue
		}
		if sym.Value != object.NULL {
			continue
		}
		if sym.Node != nil {
			value := e.evalExpression(sym.Node, env, sym.Context)
			if !isError(value) && !isRefNotFound(value) {
				sym.Value = value
			}
		}
	}
	return order, nil
}

// 環境をトポロジカルソート
func (e *Evaluator) tSortEnv(env object.Environment) ([]string, error) {
	visited := map[string]bool{}
	visiting := map[string]bool{}
	order := []string{}

	var visit func(string) error
	visit = func(name string) error {
		if visiting[name] {
			return nil
		}
		if visited[name] {
			return nil
		}
		visiting[name] = true
		obj, _ := env.Get(name)
		sym, ok := obj.(*object.SymbolObject)
		if !ok {
			visited[name] = true
			visiting[name] = false
			return nil
		}
		for _, dep := range sym.DependsOn {
			if err := visit(dep); err != nil {
				return err
			}
		}
		visited[name] = true
		visiting[name] = false
		order = append(order, name)
		return nil
	}

	for _, name := range object.CollectNames(env) {
		if err := visit(name); err != nil {
			return nil, err
		}
	}
	return order, nil
}
