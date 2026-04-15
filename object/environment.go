package object

import (
	"fmt"

	"github.com/dogatana/yas80/intern"
)

const (
	ENV_GLOBAL = iota
	ENV_PROC
	ENV_MACRO
	envMapSize = 1000 // NewEnviroment で割り当てる map[inter.SymbolID]object.Object の cap
)

// interface
type Environment interface {
	EnvType() int
	Get(id intern.SymbolID) (Object, bool)
	Set(id intern.SymbolID, obj Object) Object
	Outer() Environment
	Store() map[intern.SymbolID]Object
}

// for Global
func NewEnvironment(outer Environment) Environment {
	env := &GlobalEnvironment{store: make(map[intern.SymbolID]Object, envMapSize), outer: outer}
	// 最上位の環境にはシステム変数を定義しておく
	if outer == nil {
		setupSystemVariables(env)
	}
	return env
}

func setupSystemVariables(env Environment) {
	name := "_"
	id := intern.Intern(name)
	env.Set(id, &SymbolObject{Name: name, NameID: id, SymType: SYM_VAR, Value: NULL})

	env.Set(intern.ID_LOC, &NumberObject{Value: 0})  // $
	env.Set(intern.ID_ALOC, &NumberObject{Value: 0}) // $$

	env.Set(intern.Intern("$FILL"), &NumberObject{Value: 0})
	env.Set(intern.Intern("$R800"), &NumberObject{Value: 0})
	env.Set(intern.Intern("$STAGE2"), &NumberObject{Value: 0})

	// 以下固定値
	v0 := &NumberObject{Value: 0}
	v1 := &NumberObject{Value: 1}
	env.Set(intern.Intern("$OFF"), v0)
	env.Set(intern.Intern("$ON"), v1)
	env.Set(intern.Intern("$FALSE"), v0)
	env.Set(intern.Intern("$TRUE"), v1)

	env.Set(intern.Intern("$CMAP_ERR"), &NumberObject{Value: -1})
	env.Set(intern.Intern("$CMAP_THRU"), &NumberObject{Value: -2})

}

// for Proc
func NewProcEnvironment(outer Environment) Environment {
	return &ProcEnvironment{store: make(map[intern.SymbolID]Object), outer: outer}
}

// for Macro
func NewMacroEnvironment(outer Environment) Environment {
	return &MacroEnvironment{store: make(map[intern.SymbolID]Object), outer: outer}
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
	store map[intern.SymbolID]Object
	outer Environment
}

func (e *GlobalEnvironment) EnvType() int { return ENV_GLOBAL }
func (e *GlobalEnvironment) Get(id intern.SymbolID) (Object, bool) {
	obj, ok := e.store[id]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(id)
	}
	return obj, ok
}
func (e *GlobalEnvironment) Set(id intern.SymbolID, obj Object) Object {
	e.store[id] = obj
	return obj
}
func (e *GlobalEnvironment) Outer() Environment                { return e.outer }
func (e *GlobalEnvironment) Store() map[intern.SymbolID]Object { return e.store }
func (e *GlobalEnvironment) Print() {
	for k, v := range e.store {
		fmt.Printf("env[%q] = %s\n", k, v.String())
	}
}

// Proc 環境
type ProcEnvironment struct {
	store map[intern.SymbolID]Object
	outer Environment
}

func (e *ProcEnvironment) EnvType() int { return ENV_PROC }
func (e *ProcEnvironment) Get(id intern.SymbolID) (Object, bool) {
	obj, ok := e.store[id]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(id)
	}
	return obj, ok
}
func (e *ProcEnvironment) Set(id intern.SymbolID, obj Object) Object {
	name := id.String()
	if name[0] == '.' || name == "@@" { // .local  @@
		e.store[id] = obj
	} else if len(name) == 2 && name[0] == '@' && '1' <= name[1] && name[1] <= '9' { // @1-@9
		e.store[id] = obj
	} else {
		e.outer.Set(id, obj)
	}
	return obj
}
func (e *ProcEnvironment) Outer() Environment                { return e.outer }
func (e *ProcEnvironment) Store() map[intern.SymbolID]Object { return e.store }

// Macro 環境
type MacroEnvironment struct {
	store map[intern.SymbolID]Object
	outer Environment
}

func (e *MacroEnvironment) EnvType() int { return ENV_MACRO }
func (e *MacroEnvironment) Get(id intern.SymbolID) (Object, bool) {
	obj, ok := e.store[id]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(id)
	}
	return obj, ok
}
func (e *MacroEnvironment) Set(id intern.SymbolID, obj Object) Object {
	name := id.String()
	if name[0] == '$' { // $ で始まるシステム変数は上位Envへ処理を移譲する
		e.store[id] = obj
	} else {
		e.outer.Set(id, obj)
	}
	return obj
}
func (e *MacroEnvironment) Outer() Environment                { return e.outer }
func (e *MacroEnvironment) Store() map[intern.SymbolID]Object { return e.store }

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
			name := k.String()
			fmt.Printf("%s[%d]%sENV[%s]=%s\n", prefix, i, envType, name, v.String())
			if pobj, ok := v.(*ProcObject); ok {
				for pk, pv := range pobj.Store() {
					fmt.Printf("%s%s%s=%s\n", prefix, name, pk, pv.String())
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
