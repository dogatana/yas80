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
	Store() []Object
}

// for Global
func NewEnvironment(outer Environment) Environment {
	env := &GlobalEnvironment{store: make([]Object, 0, envMapSize), outer: outer}
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
	env.Set(intern.ID_PASS, &NumberObject{Value: 0}) // $PASS

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
	return &ProcEnvironment{store: []Object{}, outer: outer}
}

// for Macro
func NewMacroEnvironment(outer Environment) Environment {
	return &MacroEnvironment{store: []Object{}, outer: outer}
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

// store を適切に拡大する
func ensureSize(store []Object, id int) []Object {
	if id < len(store) {
		return store
	}
	newSize := len(store)
	if newSize == 0 {
		newSize = 8
	}
	for newSize <= id {
		newSize *= 2
	}
	newStore := make([]Object, newSize)
	copy(newStore, store)
	return newStore
}

// グローバル環境
type GlobalEnvironment struct {
	store []Object
	outer Environment
}

func (e *GlobalEnvironment) EnvType() int { return ENV_GLOBAL }
func (e *GlobalEnvironment) Get(id intern.SymbolID) (Object, bool) {
	idx := int(id)
	if idx < len(e.store) {
		obj := e.store[idx]
		if obj != nil {
			return obj, true
		}
	}
	if e.outer != nil {
		return e.outer.Get(id)
	}
	return nil, false
}
func (e *GlobalEnvironment) Set(id intern.SymbolID, obj Object) Object {
	idx := int(id)
	e.store = ensureSize(e.store, idx)
	e.store[idx] = obj
	return obj
}
func (e *GlobalEnvironment) Outer() Environment { return e.outer }
func (e *GlobalEnvironment) Store() []Object    { return e.store }
func (e *GlobalEnvironment) Print() {
	for k, v := range e.store {
		fmt.Printf("env[%q] = %s\n", k, v.String())
	}
}

// Proc 環境
type ProcEnvironment struct {
	store []Object
	outer Environment
}

func (e *ProcEnvironment) EnvType() int { return ENV_PROC }
func (e *ProcEnvironment) Get(id intern.SymbolID) (Object, bool) {
	idx := int(id)
	if idx < len(e.store) {
		obj := e.store[idx]
		if obj != nil {
			return obj, true
		}
	}
	if e.outer != nil {
		return e.outer.Get(id)
	}
	return nil, false
}
func (e *ProcEnvironment) Set(id intern.SymbolID, obj Object) Object {
	name := id.String()
	if name[0] == '.' || name == "@@" || (len(name) == 2 && name[0] == '@' && '1' <= name[1] && name[1] <= '9') {
		idx := int(id)
		e.store = ensureSize(e.store, idx)
		e.store[idx] = obj
		return obj
	}
	e.outer.Set(id, obj)
	return obj
}
func (e *ProcEnvironment) Outer() Environment { return e.outer }
func (e *ProcEnvironment) Store() []Object    { return e.store }

// Macro 環境
type MacroEnvironment struct {
	store []Object
	outer Environment
}

func (e *MacroEnvironment) EnvType() int { return ENV_MACRO }
func (e *MacroEnvironment) Get(id intern.SymbolID) (Object, bool) {
	idx := int(id)
	if idx < len(e.store) {
		obj := e.store[idx]
		if obj != nil {
			return obj, true
		}
	}
	if e.outer != nil {
		return e.outer.Get(id)
	}
	return nil, false
}
func (e *MacroEnvironment) Set(id intern.SymbolID, obj Object) Object {
	name := id.String()
	if name[0] == '$' { // $ で始まるシステム変数は上位Envへ処理を移譲する
		e.outer.Set(id, obj)
		return obj
	}
	// それ以外は現在の環境に登録
	idx := int(id)
	e.store = ensureSize(e.store, idx)
	e.store[idx] = obj
	return obj
}
func (e *MacroEnvironment) Outer() Environment { return e.outer }
func (e *MacroEnvironment) Store() []Object    { return e.store }

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
			if v == nil {
				continue // 削除済みエントリ
			}
			name := intern.SymbolID(k).String()
			fmt.Printf("%s[%d]%sENV[%s]=%s\n", prefix, i, envType, name, v.String())
			if pobj, ok := v.(*ProcObject); ok {
				for pk, pv := range pobj.Store() {
					if pv == nil {
						continue // 削除済みエントリ
					}
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
