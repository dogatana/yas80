package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

// 式評価
func (e *Evaluator) evalExpression(node parser.Node, env TEnv, ctx TContext) object.Object {

	switch node := node.(type) {

	// 各種リテラル
	case *parser.NumberLiteral:
		return &object.NumberObject{Value: node.Value, Context: ctx}
	case *parser.StringLiteral:
		return &object.StringObject{Value: node.Value, Context: ctx}
	case *parser.RegisterLiteral:
		return object.Z80RegisterFlagObjects[node.Register]
	case *parser.FlagLiteral:
		return object.Z80RegisterFlagObjects[node.Flag]

	// レジスタ間接 (HL),(BC),(DE),(IX+d),(IY+d),(C)
	case *parser.RegIndirectExpression:
		return e.evalRegIndirectExpression(node, env, ctx)

	// アドレス間接 (nn),(n)
	case *parser.AddrIndirectExpression:
		return e.evalAddrIndirectExpression(node, env, ctx)

	// 識別子
	case *parser.Ident:
		name := node.Name
		obj, ok := env.Get(name)
		if !ok && name[0] == '$' {
			// システム変数で未登録の場合はエラー
			e.logger.Error(fmt.Sprintf(errcode.ESYM_UNDEF, name), node.Context)
			return object.ERROR
		} else if !ok {
			// 未定義の場合、遅延評価するため RefNotFound を返す
			e.Resolved = false
			sym := object.NewUnknownSymbol(name, node.Context)
			env.Set(name, sym)
			return &object.RefNotFoundObject{Names: []string{name}}
		}

		switch obj := obj.(type) {
		case *object.ProcObject:
			return &object.NumberObject{Value: obj.Addr, Context: node.Context}
		case *object.SymbolObject:
			// _ ならそのまま返す
			if obj.Name == "_" {
				return obj
			}
			// 値が NULL なら RefNotFound にして返す
			if obj.Value == object.NULL {
				return &object.RefNotFoundObject{Names: []string{obj.Name}}
			}
			// 値が NULL でないなら Value を返す
			return obj.Value
		default:
			return obj
		}

	// enum or proc.local
	case *parser.DotIdent:
		obj, ok := env.Get(node.Left)
		// node.Left(PROC or ENUM) が未定義の場合、Name を SYM_UNKNOWN で登録
		if !ok {
			e.Resolved = false
			sym := object.NewUnknownSymbol(node.Name, node.Context)
			env.Set(node.Name, sym)
			return &object.RefNotFoundObject{Names: []string{node.Name}}
		}

		switch obj := obj.(type) {
		case *object.ProcObject:
			vobj, ok := obj.Get(node.Right)
			if !ok {
				// ローカルラベルが未定義の場合、PROC 環境に SYM_UNKNOWN で登録
				e.Resolved = false
				sym := object.NewUnknownSymbol(node.Right, node.Context)
				obj.Set(node.Right, sym) // env でなく ProcObject に登録
				return &object.RefNotFoundObject{Names: []string{node.Name}}
			}
			sym, ok := vobj.(*object.SymbolObject)
			// SymbolObject でないならそのまま返す
			if !ok {
				return vobj
			}
			// SymbolObject で値に応じた内容を返す
			if sym.Value == object.NULL {
				e.Resolved = false
				return &object.RefNotFoundObject{Names: []string{sym.Name}}
			} else {
				return sym.Value
			}

		case *object.EnumObject:
			if vobj, ok := obj.Get(node.Right); ok {
				return vobj.(*object.SymbolObject).Value
			}
			e.logger.Error(fmt.Sprintf(errcode.EENUM_ELE_UNDEF, node.Name), node.Context)
			return object.ERROR

		default:
			e.Resolved = false
			return &object.RefNotFoundObject{Names: []string{node.Name}}
		}

	// 配列リテラル
	case *parser.ArrayLiteral:
		return e.evalArrayLiteral(node, env, ctx)

	// 添え字式
	case *parser.IndexedExpression:
		return e.evalIndexedExpression(node, env, ctx)

	// 関数呼出し
	case *parser.FuncCallExpression:
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
func (e *Evaluator) evalCallExpression(expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
	if expr.Name[0] == '$' {
		return e.evalBuiltinFunction(expr, env, ctx)
	}
	obj, ok := env.Get(expr.Name)
	if !ok {
		e.logger.Error(fmt.Sprintf(errcode.EFUNC_UNDEF, expr.Name), expr.Context)
		return object.ERROR
	}

	// Symbol なら値を取り出す
	if sym, ok := obj.(*object.SymbolObject); ok {
		obj = sym.Value
	}

	fn, ok := obj.(*object.FunctionObject)
	if !ok {
		e.logger.Error(errcode.EFUNC_NOT_FUNC, ctx)
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
func (e *Evaluator) evalInfixExpression(node *parser.InfixExpression, env TEnv, ctx TContext) object.Object {
	op1 := e.evalExpression(node.Op1, env, ctx)
	op2 := e.evalExpression(node.Op2, env, ctx)

	switch {
	case isError(op1) || isError(op2):
		return object.ERROR
	case isRefNotFound(op1) || isRefNotFound(op2):
		e.Resolved = false
		return &object.RefNotFoundObject{Names: mergeNames(op1, op2)}

	case node.Operator == parser.OR || node.Operator == parser.AND:
		return e.evalLogicalInfixExpression(node.Operator, op1, op2, ctx)

	// 数値演算
	case isNumber(op1) && isNumber(op2):
		return e.evalNumberInfixExpression(node.Operator, op1, op2, ctx)

	// 文字列演算
	case isString(op1) && isString(op2):
		if node.Operator != '+' {
			e.logger.Error(errcode.EBIN_OP_TYPE, ctx)
			return object.ERROR
		}
		s1 := op1.(*object.StringObject).Value
		s2 := op2.(*object.StringObject).Value
		return &object.StringObject{Value: s1 + s2}

	default:
		if e.Debug > 0 {
			fmt.Printf("op1 %#v, op2 %#v", op1, op2)
		}
		e.logger.Error(fmt.Sprintf(errcode.EBIN_OP_TYPE, parser.TokenLiteral(node.Operator)), ctx)
		return object.ERROR
	}
}

// 論理演算 && ||
func (e *Evaluator) evalLogicalInfixExpression(opCode int, op1, op2 object.Object, ctx TContext) object.Object {
	var v1, v2 bool
	v1 = isTruthy(op1)
	v2 = isTruthy(op2)
	switch opCode {
	case parser.OR:
		return &object.NumberObject{Value: boolToInt(v1 || v2), Context: ctx}
	case parser.AND:
		return &object.NumberObject{Value: boolToInt(v1 && v2), Context: ctx}
	default:
		panic("invalid evalLogcalInfixExpression")
	}
}

// 中置演算子式（数値）
func (e *Evaluator) evalNumberInfixExpression(opCode int, op1, op2 object.Object, ctx TContext) object.Object {
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
			e.logger.Error(errcode.EBIN_OP_DIVZERO, ctx)
			return object.ERROR
		}
		return &object.NumberObject{Value: v1 / v2, Context: ctx}
	case '%':
		return &object.NumberObject{Value: v1 % v2, Context: ctx}
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
		return &object.NumberObject{Value: boolToInt(v1 != 0 && v2 != 0), Context: ctx}
	default:
		e.logger.Error(fmt.Sprintf(errcode.EBIN_OP_TYPE, parser.TokenLiteral(opCode)), nil)
		return object.ERROR
	}
}

// 前置演算子式
func (e *Evaluator) evalPrefixExpression(expr *parser.PrefixExpression, env TEnv, ctx TContext) object.Object {
	opcode := expr.Operator

	op := e.evalExpression(expr.Op, env, ctx)
	if isError(op) {
		return op
	}
	if isRefNotFound(op) {
		e.Resolved = false
		return op
	}

	// 論理否定は非演算子の Truthy を反転して返す
	if opcode == '!' {
		return &object.NumberObject{Value: boolToInt(!isTruthy(op)), Context: ctx}
	}

	// +, -, ~ は数値のみ利用可能
	if num, ok := op.(*object.NumberObject); ok {
		switch opcode {
		case '+':
			return &object.NumberObject{Value: num.Value, Context: ctx}
		case '-':
			return &object.NumberObject{Value: -num.Value, Context: ctx}
		case '~':
			return &object.NumberObject{Value: num.Value ^ -1, Context: ctx}
		}
	}

	e.logger.Error(fmt.Sprintf(errcode.EUNI_OP_TYPE, parser.TokenLiteral(opcode)), ctx)
	return object.ERROR
}

// 配列リテラル
func (e *Evaluator) evalArrayLiteral(expr *parser.ArrayLiteral, env TEnv, ctx TContext) object.Object {
	// ## の処理
	for i := range expr.Elements.Expressions {
		e.concatenateSymbol(&expr.Elements.Expressions[i], env, ctx)
	}

	values := []object.Object{}
	for _, ele := range expr.Elements.Expressions {
		obj := e.evalExpression(ele, env, ctx)
		if isError(obj) {
			continue
		}
		if isRefNotFound(obj) {
			e.Resolved = false
			return obj
		}
		values = append(values, obj)
	}

	return &object.ArrayObject{Values: values, Expressions: expr.Elements.Expressions}
}

// 添え字式
func (e *Evaluator) evalIndexedExpression(expr *parser.IndexedExpression, env TEnv, ctx TContext) object.Object {
	e.concatenateSymbol(&expr.Left, env, ctx)
	e.concatenateSymbol(&expr.Index, env, ctx)

	var array *object.ArrayObject
	var index int

	obj := e.evalExpression(expr.Left, env, ctx)
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
		return obj
	case *object.ArrayObject:
		array = obj
	default:
		e.logger.Error(errcode.EARRAY_NAME, ctx)
		return object.ERROR
	}

	obj = e.evalExpression(expr.Index, env, ctx)
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
		return obj
	case *object.NumberObject:
		index = obj.Value
		if index < 0 || index >= len(array.Values) {
			e.logger.Error(errcode.EARRAY_OUT_OF_INDEX, ctx)
			return object.ERROR
		}
	default:
		e.logger.Error(errcode.EARRAY_INDEX, ctx)
		return object.ERROR
	}

	return array.Values[index]
}

// レジスタ間接 (HL),(IX+d),(IY+d),(C)
func (e *Evaluator) evalRegIndirectExpression(expr *parser.RegIndirectExpression, env TEnv, ctx TContext) object.Object {
	e.concatenateSymbol(&expr.Displacement, env, ctx)

	reg := expr.Register
	if reg.RegisterType == parser.Z80_REG8 && reg.Register != parser.Z80_REG_C {
		e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(reg.Register)), ctx)
		return object.ERROR
	}
	if reg.Register == parser.Z80_REG_SP || reg.Register == parser.Z80_REG_AF || reg.Register == parser.Z80_REG_AFEX {
		e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_REG, parser.TokenLiteral(reg.Register)), ctx)
		return object.ERROR
	}

	if expr.Displacement == nil {
		return &object.RegIndirectObject{Register: reg.Register}
	}

	// オフセットあり
	if reg.Register != parser.Z80_REG_IX && reg.Register != parser.Z80_REG_IY {
		e.logger.Error(errcode.EINDIRECT_DISP_REG, ctx)
		return object.ERROR
	}

	obj := e.evalExpression(expr.Displacement, env, ctx)
	if isError(obj) {
		return obj
	}
	if isRefNotFound(obj) {
		e.Resolved = false
		return obj
	}

	num, ok := obj.(*object.NumberObject)
	if !ok {
		e.logger.Error(errcode.EINDIRECT_DISP, ctx)
		return object.ERROR
	}
	if num.Value < -128 || num.Value > 127 {
		e.logger.Error(fmt.Sprintf(errcode.EINDIRECT_DISP_RANGE, num.Value, num.Value), ctx)
		return object.ERROR
	}

	return &object.RegIndirectObject{Register: reg.Register, Displacement: num.Value}
}

// アドレス間接 (nn), (n)
func (e *Evaluator) evalAddrIndirectExpression(expr *parser.AddrIndirectExpression, env TEnv, ctx TContext) object.Object {
	e.concatenateSymbol(&expr.Address, env, ctx)

	obj := e.evalExpression(expr.Address, env, ctx)

	switch addr := obj.(type) {
	case *object.ErrorObject:
		return addr
	case *object.RefNotFoundObject:
		e.Resolved = false
		return addr
	case *object.NumberObject:
		return &object.AddrIndirectObject{Address: addr.Value}
	case *object.NullObject:
		e.logger.Error(errcode.EINDIRECT_NULL, ctx)
		return object.ERROR
	default:
		e.logger.Error(errcode.EINDIRECT_VALUE, ctx)
		return object.ERROR
	}
}
