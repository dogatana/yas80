package evaluator

import (
	"strings"
	"yas80/object"
)

func (e *Evaluator) EvalEnv(env TEnv) ([]string, error) {
	env.Set("$", object.NULL)

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
		// DotIdent が取得できたので、現在の環境からは削除しておく
		if strings.Contains(name, ".") {
			delete(env.Store(), name)
		}
		// システム変数は除外
		if sym.Name[0] == '$' {
			continue
		}
		// 値が NULL で Node が登録されている場合、値を更新する
		if sym.Value == object.NULL && sym.Node != nil {
			value := e.evalExpression(sym.Node, env, sym.Context)
			if !isError(value) && !isRefNotFound(value) {
				sym.Value = value
			}
		}
	}
	return order, nil
}

// 環境をトポロジカルソート
func (e *Evaluator) tSortEnv(env TEnv) ([]string, error) {
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

func (e *Evaluator) collectSymbolNames(env TEnv) []string {
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
