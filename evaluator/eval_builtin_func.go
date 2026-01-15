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
	case "$LEN", "$LENGTH":
		return e.bfuncLength(expr, env, ctx)
	case "$REV", "$REVERSE":
		return e.bfuncReverse(expr, env, ctx)
	case "$DEFINED":
		return e.bfuncDefined(expr, env, ctx)
	default:
		e.logger.Error(fmt.Sprintf(errcode.EBFN_NOT_FOUND, expr.Name), ctx)
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
	if isError(v) {
		return object.ERROR
	}
	if v, ok := v.(*object.ArrayObject); !ok {
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_VALUE, expr.Name), ctx)
		return object.ERROR
	} else {
		return &object.NumberObject{Value: len(v.Values)}
	}
}

func (e *Evaluator) bfuncReverse(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
	args := expr.Arguments.Expressions

	if len(args) != 1 {
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_COUNT, expr.Name), ctx)
		return object.ERROR
	}
	v := e.evalExpression(args[0], env, ctx)
	if isError(v) {
		return object.ERROR
	}
	if v, ok := v.(*object.ArrayObject); !ok {
		e.logger.Error(fmt.Sprintf(errcode.EBFN_ARG_VALUE, expr.Name), ctx)
		return object.ERROR
	} else {
		nv := *v
		nv.Values = slices.Clone(nv.Values)
		slices.Reverse(nv.Values)
		return &nv
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
