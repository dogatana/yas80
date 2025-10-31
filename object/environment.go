package object

import "fmt"

type Environment struct {
	Store map[string]Object
	outer *Environment
}

func NewEnvironment(outer *Environment) *Environment {
	return &Environment{make(map[string]Object), outer}
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.Store[name]
	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, obj Object) Object {
	e.Store[name] = obj
	return obj
}

func (e *Environment) GlobalGet(name string) (Object, bool) {
	obj, ok := e.GlobalEnv().Get(name)
	return obj, ok
}

func (e *Environment) GlobalSet(name string, obj Object) Object {
	e.GlobalEnv().Store[name] = obj
	return obj
}

func (e *Environment) GlobalEnv() *Environment {
	env := e
	for env.outer != nil {
		env = env.outer
	}
	return env
}

func (e *Environment) Print() {
	for k, v := range e.Store {
		fmt.Printf("env[%q] = %s\n", k, v.String())
	}
}
