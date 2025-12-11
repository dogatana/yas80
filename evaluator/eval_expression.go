package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/fileblock"
	"yas80/object"
	"yas80/parser"
)

// 関数呼出し
func (e *Evaluator) evalCallExpression(expr *parser.CallExpression, env object.Environment) object.Object {
	obj := e.Eval(expr.Function, env)
	if isError(obj) || isRefNotFound(obj) {
		e.Resolved = false
		return obj
	} else if obj == object.NULL {
		panic("object is NULL") // TODO
		// return &object.NodeObject{Value: expr, LineNumber: expr.ContextNumber()}
	}

	fn, ok := obj.(*object.FunctionObject)
	if !ok {
		e.logger.Error(errcode.E019, expr.Context)
		return object.ERROR
	}
	if len(expr.Arguments.Expressions) != len(fn.Params) {
		e.logger.Error(fmt.Sprintf(errcode.EFUNC_ARG_COUNT, fn.Name), expr.Context)
		return object.ERROR
	}

	newEnv := object.NewEnvironment(fn.Env)
	for i, param := range fn.Params {
		v := e.Eval(expr.Arguments.Expressions[i], env)
		if isError(v) || isRefNotFound(v) {
			return v
		}
		newEnv.Set(param, v)
	}

	ret, ok := e.evalBlockStatement(fn.Body.(*parser.BlockStatement), newEnv).(*object.BlockObject)
	if !ok {
		panic(fmt.Sprintf("call func %s returns %T(%#v)", fn.Name, ret, ret))
	}

	if len(ret.Block) == 0 {
		return object.NULL
	}
	for _, obj := range ret.Block {
		if isError(obj) || isRefNotFound(obj) {
			return obj
		}
	}
	last := ret.Block[len(ret.Block)-1]
	if last.Type() == object.RETURN_OBJ {
		return last.(*object.ReturnObject).Value
	}
	return object.NULL
}

// 中置演算子式
func (e *Evaluator) evalInfixExpression(node *parser.InfixExpression, env object.Environment, ctx *fileblock.Context) object.Object {
	op1 := unwrapSymbol(e.Eval(node.Op1, env))
	op2 := unwrapSymbol(e.Eval(node.Op2, env))

	switch {
	case isError(op1) || isError(op2):
		return object.ERROR
	case isNumber(op1) && isNumber(op2):
		return e.evalNumberInfixExpression(node.Operator, op1, op2, ctx)
	case isString(op1) && isString(op2):
		if node.Operator != '+' {
			e.logger.Error(errcode.EBIN_OP_STRING, ctx)
			return object.ERROR
		}
		s1 := op1.(*object.StringObject).Value
		s2 := op2.(*object.StringObject).Value
		return &object.StringObject{Value: s1 + " " + s2}
	case isRefNotFound(op1) || isRefNotFound(op2):
		e.Resolved = false
		return &object.RefNotFoundObject{Names: mergeNames(op1, op2)}
	case isSymolOrSymbolExpr(op1) || isSymolOrSymbolExpr(op2):
		return &object.SymbolExprObject{Names: mergeNames(op1, op2)}
	default:
		if e.Debug > 0 {
			fmt.Printf("op1 %#v, op2 %#v", op1, op2)
		}
		e.logger.Error(fmt.Sprintf(errcode.EBIN_OP_TYPE, parser.TokenLiteral(node.Operator)), ctx)
		return object.ERROR
	}
}

func (e *Evaluator) evalNumberInfixExpression(opCode int, op1, op2 object.Object, ctx *fileblock.Context) object.Object {
	v1 := op1.(*object.NumberObject).Value
	v2 := op2.(*object.NumberObject).Value
	switch opCode {
	case '+':
		return &object.NumberObject{Value: v1 + v2, Context: ctx}
	case '-':
		return &object.NumberObject{Value: v1 - v2, Context: ctx}
	case '*':
		return &object.NumberObject{Value: v1 * v2, Context: ctx}
	case '/':
		if v2 == 0 {
			e.logger.Error(errcode.EBIN_OP_DIVZERO, nil)
			return object.ERROR
		}
		return &object.NumberObject{Value: v1 / v2, Context: ctx}
	case parser.SL:
		return &object.NumberObject{Value: v1 << v2, Context: ctx}
	case parser.SR:
		return &object.NumberObject{Value: v1 >> v2, Context: ctx}
	case '&':
		return &object.NumberObject{Value: v1 & v2, Context: ctx}
	case '|':
		return &object.NumberObject{Value: v1 | v2, Context: ctx}
	case '^':
		return &object.NumberObject{Value: v1 ^ v2, Context: ctx}
	case parser.EQ:
		return &object.NumberObject{Value: boolToInt(v1 == v2), Context: ctx}
	case parser.NEQ:
		return &object.NumberObject{Value: boolToInt(v1 != v2), Context: ctx}
	case '<':
		return &object.NumberObject{Value: boolToInt(v1 < v2), Context: ctx}
	case parser.LE:
		return &object.NumberObject{Value: boolToInt(v1 <= v2), Context: ctx}
	case '>':
		return &object.NumberObject{Value: boolToInt(v1 > v2), Context: ctx}
	case parser.GE:
		return &object.NumberObject{Value: boolToInt(v1 >= v2), Context: ctx}
	case parser.OR:
		return &object.NumberObject{Value: boolToInt(v1 != 0 || v2 != 0), Context: ctx}
	case parser.AND:
		return &object.NumberObject{Value: boolToInt(v1 != 1 && v2 != 1), Context: ctx}
	default:
		e.logger.Error(fmt.Sprintf(errcode.EBIN_OP_NUMBER, string(rune(opCode))), nil)
		return object.ERROR
	}
}

// 前置演算子式
func (e *Evaluator) evalPrefixExpression(expr *parser.PrefixExpression, env object.Environment, ctx *fileblock.Context) object.Object {
	opcode := expr.Operator

	op := unwrapSymbol(e.Eval(expr.Op, env))
	if isError(op) {
		return op
	}
	if isRefNotFound(op) {
		e.Resolved = false
		return op
	}

	switch op := op.(type) {
	case *object.NumberObject:
		switch opcode {
		case '+':
			return &object.NumberObject{Value: op.Value, Context: ctx}
		case '-':
			return &object.NumberObject{Value: -op.Value, Context: ctx}
		case '~':
			return &object.NumberObject{Value: op.Value ^ -1, Context: ctx}
		case '!':
			return &object.NumberObject{Value: boolToInt(op.Value == 0), Context: ctx}
		default:
			e.logger.Error(fmt.Sprintf(errcode.EUNARY_OP_NUMBER, rune(opcode)), ctx)
			return object.ERROR
		}
	case *object.StringObject:
		if opcode == '!' {
			return &object.NumberObject{Value: boolToInt(op.Value == ""), Context: ctx}
		}
		e.logger.Error(fmt.Sprintf(errcode.EUNARY_OP_STRING, rune(opcode)), ctx)
		return object.ERROR
	case *object.SymbolExprObject:
		return op // TODO
	default:
		e.logger.Error(fmt.Sprintf(errcode.EUNARY_OP_TYPE, parser.TokenLiteral(opcode)), ctx)
		return object.ERROR
	}
}
