package object

import (
	"fmt"
)

const (
	ENV_GLOBAL = iota
	ENV_PROC
	ENV_MACRO
)

// interface
type Environment interface {
	EnvType() int
	Get(name string) (Object, bool)
	Set(name string, obj Object) Object
	Outer() Environment
	Store() map[string]Object
}

// for Global
func NewEnvironment(outer Environment) Environment {
	env := &GlobalEnvironment{store: make(map[string]Object), outer: outer}
	// 最上位の環境にはシステム変数を定義しておく
	if outer == nil {
		setupSystemVariables(env)
	}
	return env
}

func setupSystemVariables(env Environment) {
	env.Set("_", &SymbolObject{Name: "_", SymType: SYM_VAR, Value: NULL})

	env.Set("$", &NumberObject{Value: 0})
	env.Set("$FILL", &NumberObject{Value: 0})
}

// for Proc
func NewProcEnvironment(outer Environment) Environment {
	return &ProcEnvironment{store: make(map[string]Object), outer: outer}
}

// for Macro
func NewMacroEnvironment(outer Environment) Environment {
	return &MacroEnvironment{store: make(map[string]Object), outer: outer}
}

// 引数の環境の上位が ENV_GLOBAL か ENV_PROC かを返す
func OuterEnvType(env Environment) int {
	for env.EnvType() == ENV_MACRO {
		env = env.Outer()
	}
	return env.EnvType()
}

// 引数の環境が ProcEnvironment に含まれるかどうか
func InProcEnv(env Environment) bool {
	for {
		if env.EnvType() == ENV_PROC {
			return true
		}
		env = env.Outer()
		if env == nil {
			return false
		}
	}
}

// グローバル環境
type GlobalEnvironment struct {
	store map[string]Object
	outer Environment
}

func (e *GlobalEnvironment) EnvType() int { return ENV_GLOBAL }
func (e *GlobalEnvironment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}
func (e *GlobalEnvironment) Set(name string, obj Object) Object {
	e.store[name] = obj
	return obj
}
func (e *GlobalEnvironment) Outer() Environment       { return e.outer }
func (e *GlobalEnvironment) Store() map[string]Object { return e.store }
func (e *GlobalEnvironment) Print() {
	for k, v := range e.store {
		fmt.Printf("env[%q] = %s\n", k, v.String())
	}
}

// Proc 環境
type ProcEnvironment struct {
	store map[string]Object
	outer Environment
}

func (e *ProcEnvironment) EnvType() int { return ENV_PROC }
func (e *ProcEnvironment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}
func (e *ProcEnvironment) Set(name string, obj Object) Object {
	if name[0] == '.' { // Proc Local
		e.store[name] = obj
	} else {
		e.outer.Set(name, obj)
	}
	return obj
}
func (e *ProcEnvironment) Outer() Environment       { return e.outer }
func (e *ProcEnvironment) Store() map[string]Object { return e.store }

// Macro 環境
type MacroEnvironment struct {
	store map[string]Object
	outer Environment
}

func (e *MacroEnvironment) EnvType() int { return ENV_MACRO }
func (e *MacroEnvironment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}
func (e *MacroEnvironment) Set(name string, obj Object) Object {
	// if strings.HasPrefix(name, "@@") { // Macro Parameter
	// 	e.store[name[2:]] = obj
	// } else if name[0] == '@' { // Macro Local
	// 	e.store[name] = obj
	// } else {
	// 	e.outer.Set(name, obj)
	// }
	if name[0] == '$' {
		e.store[name] = obj
	} else {
		e.outer.Set(name, obj)
	}
	return obj
}
func (e *MacroEnvironment) Outer() Environment       { return e.outer }
func (e *MacroEnvironment) Store() map[string]Object { return e.store }

// debug
func PrintEnv(env Environment) {
	prefix := ""
	for i := 0; ; i++ {
		var envType string
		switch env.(type) {
		case *GlobalEnvironment:
			envType = ""
		case *MacroEnvironment:
			envType = "@"
		case *ProcEnvironment:
			envType = "P"
		default:
			envType = "?"
		}
		for k, v := range env.Store() {
			fmt.Printf("%s[%d]%sENV[%s]=%s\n", prefix, i, envType, k, v.String())
			if pobj, ok := v.(*ProcObject); ok {
				for pk, pv := range pobj.Store() {
					fmt.Printf("%s%s%s=%s\n", prefix, k, pk, pv.String())
				}
			}
		}
		prefix += "  "
		env = env.Outer()
		if env == nil {
			break
		}
		fmt.Println("---")
	}
	fmt.Println("")
}
