package evaluator

import (
	"fmt"
	"strings"
	"yas80/errcode"
	"yas80/fileblock"
	"yas80/object"
	"yas80/parser"
)

// 各種 object 判定

func isError(obj object.Object) bool {
	return obj.Type() == object.ERROR_OBJ
}

func isNumber(obj object.Object) bool {
	return obj.Type() == object.NUMBER_OBJ
}

func isString(obj object.Object) bool {
	return obj.Type() == object.STRING_OBJ
}

func isRefNotFound(obj object.Object) bool {
	return obj.Type() == object.REF_NOTFOUND_OBJ
}

// 依存先の識別子を抽出する: 重複する名は後段のソートでユニークになる
func mergeNames(obj1, obj2 object.Object) []string {
	names := []string{}

	names = append(names, extractNames(obj1)...)
	names = append(names, extractNames(obj2)...)

	return names
}

func extractNames(obj object.Object) []string {
	switch obj := obj.(type) {
	case *object.RefNotFoundObject:
		return obj.Names
	case *object.SymbolObject:
		return []string{obj.Name}
	default:
		return []string{}
	}
}

// location counter 初期化
func initLocationCounter(env object.Environment, addr int) {
	obj, ok := env.Get("$")
	if !ok {
		panic("no $ in Environment")
	}
	obj.(*object.NumberObject).Value = addr
}

// location counter 取得
func getLocationCounter(env object.Environment) int {
	counter, ok := env.Get("$")
	if !ok {
		panic("getLocationCounter failed")
	}
	return counter.(*object.NumberObject).Value

}

// location counter 表示
func printLocationCounter(env object.Environment) {
	fmt.Printf("$ %04x\n", getLocationCounter(env))
}

// location counter 更新
func advanceLocationCounter(env object.Environment, n int) {
	obj, ok := env.Get("$")
	if !ok {
		panic("getLocationCounter failed")
	}
	counter := obj.(*object.NumberObject)
	counter.Value += n
}

func boolToInt(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}

func isTruthy(obj object.Object) bool {
	switch obj := obj.(type) {
	case *object.NumberObject:
		return obj.Value != 0
	case *object.StringObject:
		return obj.Value != ""
	default:
		return false
	}
}

func CollectCode(prog *object.ProgramObject) []byte {
	var result []byte
	for _, obj := range prog.Objects {
		code, ok := obj.(*object.CodeObject)
		if !ok {
			continue
		}
		result = append(result, code.Code...)
	}
	return result
}

// シンボル結合処理
func (e *Evaluator) concatenateSymbol(ptr *parser.Expression, env object.Environment, ctx *fileblock.Context) bool {
	switch expr := (*ptr).(type) {
	case *parser.InfixExpression:
		if expr.Operator != parser.CONCAT {
			return e.concatenateSymbol(&expr.Op1, env, ctx) || e.concatenateSymbol(&expr.Op2, env, ctx)
		}
		ident, ok := expr.Op1.(*parser.Ident)
		if !ok {
			panic("not ident")
			// TODO: parser の段階でここには来ないはず
			// e.logger.Error(errcode.ESYM_CONCAT_NOTSYM, ctx)
			// return false
		}
		// copy &(*ident) では新しい値が生成されないため
		{
			temp := *ident
			ident = &temp
		}

		suffix := ""
		op2 := e.evalExpression(expr.Op2, env, ctx)
		switch op2 := op2.(type) {
		case *object.ErrorObject:
			return false
		case *object.RefNotFoundObject:
			names := strings.Join(op2.Names, ", ")
			e.logger.Error(fmt.Sprintf(errcode.E009, names), ctx)
			return false
		case *object.NumberObject:
			suffix = fmt.Sprintf("%d", op2.Value)
		case *object.StringObject:
			suffix = op2.Value
		default:
			e.logger.Error(errcode.ESYM_CONCAT_TYPE, ctx)
			return false
		}
		ident.Name += suffix
		*ptr = ident
		return true
	case *parser.PrefixExpression:
		return e.concatenateSymbol(&expr.Op, env, ctx)
	default:
		return false
	}
}
