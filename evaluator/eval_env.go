package evaluator

import (
	"fmt"
	"strings"
	"yas80/object"
)

func (e *Evaluator) EvalEnv(env object.Environment) ([]string, error) {
	order, err := e.tSortEnv(env)
	// fmt.Println("order", order)
	if err != nil {
		return order, err
	}
	for _, name := range order {
		sym, ok := e.getSymbolFromEnv(name, env)
		if !ok {
			continue
		}
		// システム変数は除外
		if sym.Name[0] == '$' {
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
		sym, ok := e.getSymbolFromEnv(name, env)
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

	for _, name := range e.collectSymbolNames(env) {
		if err := visit(name); err != nil {
			return nil, err
		}
	}
	return order, nil
}

func (e *Evaluator) getSymbolFromEnv(name string, env object.Environment) (*object.SymbolObject, bool) {
	names := strings.Split(name, ".")
	if len(names) == 1 {
		if obj, ok := env.Get(name); ok && obj.Type() == object.SYMBOL_OBJ {
			return obj.(*object.SymbolObject), true
		}
		return nil, false
	}
	obj, ok := env.Get(names[0])
	if !ok {
		return nil, false
	}
	switch obj := obj.(type) {
	case *object.ProcObject:
		v, ok := obj.Get("." + names[1])
		if !ok {
			return nil, false
		}
		if sym, ok := v.(*object.SymbolObject); ok {
			return sym, ok
		}
		return nil, false
	default:
		panic(fmt.Sprintf("getSymbolFromEnv error %#v", obj))
	}
}

func (e *Evaluator) collectSymbolNames(env object.Environment) []string {
	names := []string{}

	for {
		for name, obj := range env.Store() {
			switch sym := obj.(type) {
			case *object.SymbolObject:
				if name != "_" && name[0] != '$' {
					names = append(names, name)
				}
			case *object.ProcObject:
				for k := range sym.Store() {
					names = append(names, name+k)
				}
			}
		}
		env = env.Outer()
		if env == nil {
			break
		}
	}
	// fmt.Println("names", names)
	return names
}
