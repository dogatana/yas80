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
	case "$H", "$HI", "$L", "$LO":
		return e.bfuncHighLow(expr, env, ctx)
	case "$LEN", "$LENGTH":
		return e.bfuncLength(expr, env, ctx)
	case "$REV", "$REVERSE":
		return e.bfuncReverse(expr, env, ctx)
	case "$DEFINED":
		return e.bfuncDefined(expr, env, ctx)
	case "$FMT", "$FORMAT":
		return e.bfuncFormat(expr, env, ctx)
	default:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_NOT_FOUND, expr.Name), ctx)
		return object.ERROR
	}
}

func (e *Evaluator) bfuncHighLow(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
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

func (e *Evaluator) bfuncLength(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
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

func (e *Evaluator) bfuncReverse(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
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

func (e *Evaluator) bfuncDefined(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
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

func (e *Evaluator) bfuncFormat(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
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
