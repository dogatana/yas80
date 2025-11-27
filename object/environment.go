package object

import "fmt"

type Environment interface {
	Get(name string) (Object, bool)
	Set(name string, obj Object) Object
	Outer() Environment
	// Store() map[string]Object
}

func CollectNames(env Environment) []string {
	names := []string{}

	for {
		switch cenv := env.(type) {
		case *NormalEnvironment:
			for k := range cenv.store {
				names = append(names, k)
			}
			env = cenv.outer
		case *AtLocalEnvironment:
			for k := range cenv.store {
				names = append(names, k)
			}
			env = cenv.outer
		default:
			panic(fmt.Sprintf("unkown Environment %T", cenv))
		}
		if env == nil {
			break
		}
	}
	return names
}

func PrintEnv(env Environment) {
	prefix := ""
	for {
		switch cenv := env.(type) {
		case *NormalEnvironment:
			for k, v := range cenv.store {
				fmt.Printf("%sENV[%s]=%s\n", prefix, k, v.String())
			}
			env = cenv.outer
		case *AtLocalEnvironment:
			for k, v := range cenv.store {
				fmt.Printf("%s@ENV[%s]=%s\n", prefix, k, v.String())
			}
			env = cenv.outer
		default:
			panic(fmt.Sprintf("unkown Environment %T", cenv))
		}
		prefix += "  "
		if env == nil {
			break
		}
	}
}

type NormalEnvironment struct {
	store map[string]Object
	outer Environment
}

func NewEnvironment(outer Environment) Environment {
	env := &NormalEnvironment{make(map[string]Object), outer}
	if outer == nil {
		env.Set("$", &NumberObject{Value: 0})
	}
	return env
}

func (e *NormalEnvironment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *NormalEnvironment) Set(name string, obj Object) Object {
	e.store[name] = obj
	return obj
}

func (e *NormalEnvironment) Outer() Environment {
	return e.outer
}

func (e *NormalEnvironment) Store() map[string]Object {
	return e.store
}

func (e *NormalEnvironment) Print() {
	for k, v := range e.store {
		fmt.Printf("env[%q] = %s\n", k, v.String())
	}
}

type AtLocalEnvironment struct {
	store map[string]Object
	outer Environment
}

func NewAtLocalEnvironment(outer Environment) Environment {
	return &AtLocalEnvironment{make(map[string]Object), outer}
}

func (e *AtLocalEnvironment) Get(name string) (Object, bool) {
	var obj Object
	var ok bool

	if name[0] == '@' {
		obj, ok = e.store[name]
		return obj, ok
	} else if e.outer != nil {
		obj, ok = e.outer.Get(name)
		return obj, ok
	} else {
		panic("no outer Environment")
	}
}

func (e *AtLocalEnvironment) Set(name string, obj Object) Object {
	if name[0] == '@' {
		e.store[name] = obj
	} else {
		e.outer.Set(name, obj)
	}
	return obj
}

func (e *AtLocalEnvironment) Outer() Environment {
	return e.outer
}

func (e *AtLocalEnvironment) Store() map[string]Object {
	return e.store
}
