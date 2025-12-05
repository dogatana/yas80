package evaluator

import (
	"fmt"
	"yas80/object"
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

func isSymolOrSymbolExpr(obj object.Object) bool {
	return obj.Type() == object.SYMBOL_OBJ || obj.Type() == object.SYMBOL_EXPR_OBJ
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
	case *object.SymbolExprObject:
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

// Symbol.SymState に応じて値を unwrap して返す
func unwrapSymbol(obj object.Object) object.Object {
	sym, ok := obj.(*object.SymbolObject)
	if !ok {
		return obj
	}
	if sym.SymState == object.VALUE_DETERMINED || sym.SymState == object.VALUE_TENTATIVE {
		// unwrap して Value を返す
		return sym.Value
	}
	return obj
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
