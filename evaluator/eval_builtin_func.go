package evaluator

import (
	"fmt"
	"slices"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

// TODO: map[e.Name] func としたいが、循環エラーになるので switch/case とする
func (e *Evaluator) evalBuiltinFunction(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
	switch expr.Name {
	case "$W", "$WORD":
		return e.ebfuncWord(expr, env, ctx)
	case "$H", "$HI", "$L", "$LO":
		return e.ebfuncHighLow(expr, env, ctx)
	case "$LEN", "$LENGTH":
		return e.ebfuncLength(expr, env, ctx)
	case "$REV", "$REVERSE":
		return e.ebfuncReverse(expr, env, ctx)
	case "$DEFINED":
		return e.ebfuncDefined(expr, env, ctx)
	case "$FMT", "$FORMAT":
		return e.ebfuncFormat(expr, env, ctx)
	case "$ISARY", "$ISARRAY":
		return e.ebfuncIsArray(expr, env, ctx)
	default:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_NOT_FOUND, expr.Name), ctx)
		return object.ERROR
	}
}

func (e *Evaluator) ebfuncWord(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
	args := expr.Arguments.Expressions

	if len(args) != 1 {
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_COUNT, expr.Name), ctx)
		return object.ERROR
	}
	v := e.evalExpression(args[0], env, ctx)

	switch v := v.(type) {
	case *object.ErrorObject:
		return v
	case *object.RefNotFoundObject:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_NULL, expr.Name), ctx)
		return object.ERROR
	case *object.NumberObject:
		w, ok := e.intToWord(v.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, v.Value, v.Value), ctx)
		}
		return &object.NumberObject{Value: w, ForceWord: true, Context: ctx}
	default:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_VALUE, expr.Name), ctx)
		return object.ERROR
	}
}

func (e *Evaluator) ebfuncHighLow(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
	args := expr.Arguments.Expressions

	if len(args) != 1 {
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_COUNT, expr.Name), ctx)
		return object.ERROR
	}
	v := e.evalExpression(args[0], env, ctx)

	switch v := v.(type) {
	case *object.ErrorObject:
		return v
	case *object.RefNotFoundObject:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_NULL, expr.Name), ctx)
		return object.ERROR
	case *object.NumberObject:
		var b int
		if expr.Name[1] == 'H' {
			b = (v.Value >> 8) & 0xff
		} else {
			b = v.Value & 0xff
		}
		return &object.NumberObject{Value: b, Context: ctx}
	default:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_VALUE, expr.Name), ctx)
		return object.ERROR
	}
}

func (e *Evaluator) ebfuncIsArray(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
	args := expr.Arguments.Expressions

	if len(args) != 1 {
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_COUNT, expr.Name), ctx)
		return object.ERROR
	}
	v := e.evalExpression(args[0], env, ctx)

	switch v := v.(type) {
	case *object.ErrorObject:
		return v
	case *object.RefNotFoundObject:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_NULL, expr.Name), ctx)
		return object.ERROR
	case *object.ArrayObject:
		return &object.NumberObject{Value: 1}
	default:
		return &object.NumberObject{Value: 0}
	}
}

func (e *Evaluator) ebfuncLength(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
	args := expr.Arguments.Expressions

	if len(args) != 1 {
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_COUNT, expr.Name), ctx)
		return object.ERROR
	}
	v := e.evalExpression(args[0], env, ctx)

	switch v := v.(type) {
	case *object.ErrorObject:
		return v
	case *object.RefNotFoundObject:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_NULL, expr.Name), ctx)
		return object.ERROR
	case *object.ArrayObject:
		return &object.NumberObject{Value: len(v.Values)}
	default:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_VALUE, expr.Name), ctx)
		return object.ERROR
	}
}

func (e *Evaluator) ebfuncReverse(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
	args := expr.Arguments.Expressions

	if len(args) != 1 {
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_COUNT, expr.Name), ctx)
		return object.ERROR
	}
	v := e.evalExpression(args[0], env, ctx)

	switch v := v.(type) {
	case *object.ErrorObject:
		return v
	case *object.RefNotFoundObject:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_NULL, expr.Name), ctx)
		return object.ERROR
	case *object.ArrayObject:
		nv := *v
		nv.Values = slices.Clone(nv.Values)
		slices.Reverse(nv.Values)
		return &nv
	default:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_VALUE, expr.Name), ctx)
		return object.ERROR
	}
}

func (e *Evaluator) ebfuncDefined(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
	args := expr.Arguments.Expressions

	if len(args) != 1 {
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_COUNT, expr.Name), ctx)
		return object.ERROR
	}
	if id, ok := args[0].(*parser.Ident); !ok {
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_VALUE, expr.Name), ctx)
		return object.ERROR
	} else {
		_, ok = env.Get(id.Name)
		return &object.NumberObject{Value: boolToInt(!ok), Context: ctx}
	}
}

func (e *Evaluator) ebfuncFormat(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
	args := expr.Arguments.Expressions

	if len(args) == 0 {
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_COUNT, expr.Name), ctx)
		return object.ERROR
	}

	var fmts string
	obj := e.evalExpression(args[0], env, ctx)
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_NULL, expr.Name), ctx)
		return object.ERROR
	case *object.StringObject:
		if len(args) == 1 {
			return obj
		}
		fmts = obj.Value
	default:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_VALUE, expr.Name), ctx)
		return object.ERROR
	}

	fargs := make([]any, 0, len(args)-1)
	for _, arg := range args[1:] {
		v := e.evalExpression(arg, env, ctx)
		switch v := v.(type) {
		case *object.NumberObject:
			fargs = append(fargs, v.Value)
		case *object.StringObject:
			fargs = append(fargs, v.Value)
		case *object.ErrorObject:
			return v
		default:
			e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_VALUE, expr.Name), ctx)
			return object.ERROR
		}
	}
	s := fmt.Sprintf(fmts, fargs...)
	return &object.StringObject{Value: s, Context: ctx}
}
