package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/fileblock"
	"yas80/object"
	"yas80/parser"
)

// 式評価
func (e *Evaluator) evalExpression(node parser.Node, env object.Environment, ctx *fileblock.Context) object.Object {

	switch node := node.(type) {

	// 各種リテラル
	case *parser.NumberLiteral:
		return &object.NumberObject{Value: node.Value, Context: ctx}
	case *parser.StringLiteral:
		return &object.StringObject{Value: node.Value, Context: ctx}
	case *parser.RegisterLiteral:
		return object.Z80RegisterFlagObjects[int(node.NodeSubType())]
	case *parser.FlagLiteral:
		return object.Z80RegisterFlagObjects[int(node.NodeSubType())]

	// 識別子
	case *parser.Ident:
		name := node.Name
		obj, ok := env.Get(name)
		if !ok {
			// 未定義の場合
			e.Resolved = false
			sym := object.NewUnknownSymbol(name, node.Context)
			env.Set(name, sym)
			return &object.RefNotFoundObject{Names: []string{name}}
		}

		sym, ok := obj.(*object.SymbolObject)
		if !ok {
			// SymbolObject 以外ならそのまま返す
			return obj
		}

		// _ ならそのまま返す
		if sym.Name == "_" {
			return sym
		}
		// 値が NULL でないなら value を返す
		if sym.Value != object.NULL {
			return sym.Value
		}
		// 値が NULL なら RefNotFound にして返す
		return &object.RefNotFoundObject{Names: []string{sym.Name}}
		// return &object.RefNotFoundObject{Names: append(sym.DependsOn, sym.Name)}

	// enum or proc.local
	case *parser.DotIdent: // TODO enum か proc.local かの識別必要
		enum, ok := env.Get(node.Left)
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.E010, node.Left), ctx)
			return object.ERROR
		}
		v, ok := enum.(*object.EnumObject).Get(node.Right)
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.E011, node.Left, node.Right), ctx)
			return object.ERROR
		}
		return v

	// 関数呼出し
	case *parser.CallExpression:
		return e.evalCallExpression(node, env, ctx)

	// 中置演算子
	case *parser.InfixExpression:
		return e.evalInfixExpression(node, env, ctx)

	// 前置演算子
	case *parser.PrefixExpression:
		return e.evalPrefixExpression(node, env, ctx)

	default:
		e.logger.Error(fmt.Sprintf(errcode.ENOT_IMPL_EXPR, node), ctx)
		return object.ERROR
	}
}

// 関数呼出し
func (e *Evaluator) evalCallExpression(expr *parser.CallExpression, env object.Environment, ctx *fileblock.Context) object.Object {
	obj := e.evalExpression(expr.Function, env, ctx)
	if isError(obj) || isRefNotFound(obj) { // TODO: エラーとRefNotFound を分ける
		e.Resolved = false
		return obj
	} else if obj == object.NULL {
		panic("object is NULL") // TODO
		// return &object.NodeObject{Value: expr, LineNumber: expr.ContextNumber()}
	}

	fn, ok := obj.(*object.FunctionObject)
	if !ok {
		e.logger.Error(errcode.E019, ctx)
		return object.ERROR
	}
	if len(expr.Arguments.Expressions) != len(fn.Params) {
		e.logger.Error(fmt.Sprintf(errcode.EFUNC_ARG_COUNT, fn.Name), ctx)
		return object.ERROR
	}

	newEnv := object.NewEnvironment(fn.Env)
	for i, param := range fn.Params {
		v := e.evalExpression(expr.Arguments.Expressions[i], env, nil) // TODO nil
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
	op1 := e.evalExpression(node.Op1, env, ctx)
	op2 := e.evalExpression(node.Op2, env, ctx)

	switch {
	case isError(op1) || isError(op2):
		return object.ERROR
	case isRefNotFound(op1) || isRefNotFound(op2):
		e.Resolved = false
		return &object.RefNotFoundObject{Names: mergeNames(op1, op2)}

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

	op := e.evalExpression(expr.Op, env, ctx)
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

	default:
		e.logger.Error(fmt.Sprintf(errcode.EUNARY_OP_TYPE, parser.TokenLiteral(opcode)), ctx)
		return object.ERROR
	}
}
