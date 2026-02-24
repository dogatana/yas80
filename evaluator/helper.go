package evaluator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"yas80/errcode"
	"yas80/internal/util"
	"yas80/object"
	"yas80/parser"
)

// 各種 object 判定

func isError(obj object.Object) bool {
	return obj.Type() == object.OBJ_ERROR
}

func isNumber(obj object.Object) bool {
	return obj.Type() == object.OBJ_NUMBER
}

func isString(obj object.Object) bool {
	return obj.Type() == object.OBJ_STRING
}

func isRefNotFound(obj object.Object) bool {
	return obj.Type() == object.OBJ_REF_NOTFOUND
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
func initLocationCounter(env TEnv, addr int) {
	env.Set("$", &object.NumberObject{Value: addr})
}

// location counter 取得
func getLocationCounter(env TEnv) int {
	counter, ok := env.Get("$")
	if !ok {
		panic("getLocationCounter failed")
	}
	return counter.(*object.NumberObject).Value

}

// location counter 更新
func advanceLocationCounter(env TEnv, n int) error {
	obj, ok := env.Get("$")
	if !ok {
		panic("getLocationCounter failed")
	}
	counter := obj.(*object.NumberObject)
	counter.Value += n

	// 64KB アドレス超過のチェック
	if counter.Value > 0x10000 {
		return fmt.Errorf(errcode.EADDRESS_OVERFLOW, counter.Value)
	}
	return nil
}

// int -> byte へ。丸め発生した場合は false
func (e *Evaluator) intToByte(n int) (byte, bool) {
	return byte(n & 0xff), -128 <= n && n <= 255
}

// int -> word へ。丸めが発生した場合は false
func (e *Evaluator) intToWord(n int) (int, bool) {
	return (n & 0xffff), -32768 <= n && n <= 65535
}

// int ->  byte へ。丸めが発生した場合は false
func (e *Evaluator) intToPort(n int) (byte, bool) {
	return byte(n & 0xff), 0 <= n && n <= 255
}

// int ->  byte(0-7) へ。丸めが発生した場合は false
func (e *Evaluator) intToBit(n int) (byte, bool) {
	return byte(n & 0x7), 0 <= n && n <= 7
}

func boolToInt(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}

// シンボル結合処理
func (e *Evaluator) concatenateSymbol(ptr *parser.Expression, env TEnv, ctx TContext) bool {
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
			e.logger.Error(fmt.Sprintf(errcode.ESYM_UNDEF, names), ctx)
			return false
		case *object.NumberObject:
			suffix = fmt.Sprintf("%d", op2.Value)
		case *object.StringObject:
			suffix = strings.ToUpper(op2.Value) // 文字列リテラルは大文字化して結合する
		default:
			e.logger.Error(errcode.ECONCAT_TYPE, ctx)
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

func (e *Evaluator) getSymbolFromEnv(name string, env TEnv) (*object.SymbolObject, bool) {
	names := strings.Split(name, ".")
	if len(names) == 1 {
		if obj, ok := env.Get(name); ok {
			switch obj := obj.(type) {
			case *object.SymbolObject:
				return obj, true
			case *object.ProcObject:
				return &object.SymbolObject{Name: name, Value: &object.NumberObject{Value: obj.Addr}}, true
			}
			return nil, false
		}
	}
	obj, ok := env.Get(names[0])
	if !ok {
		return nil, false
	}
	switch obj := obj.(type) {
	case *object.ProcObject:
		v, ok := obj.Get("." + names[1])
		if !ok {
			return nil, false
		}
		if sym, ok := v.(*object.SymbolObject); ok {
			return sym, ok
		}
		return nil, false

	case *object.EnumObject:
		v, ok := obj.Get("." + names[1])
		if !ok {
			return nil, false
		}
		if sym, ok := v.(*object.SymbolObject); ok {
			return sym, ok
		}
		return nil, false

	default:
		panic(fmt.Sprintf("getSymbolFromEnv error %#v", obj))
	}
}

// parser.Expression -> parser.Ident - >parser.Label を評価・環境登録し object.SymbolObject を返す
func (e *Evaluator) exprToLabel(expr parser.Expression, env TEnv, ctx TContext) object.Object {
	id, ok := expr.(*parser.Ident)
	if !ok {
		return object.ERROR
	}
	label := e.identToLabel(id)
	label.Context = ctx // Ident の Context でなく、引数（文）の Conext を設定する

	return e.evalLabel(label, env)
}

// parser.Ident -> parser.Label 変換(exprToLabel から呼ばれる)
func (e *Evaluator) identToLabel(id *parser.Ident) *parser.Label {
	l := &parser.Label{Name: id.Name, LabelType: parser.NODE_LABEL, Context: id.Context}
	switch id.Name[0] {
	case '.':
		l.LabelType = parser.NODE_LOCAL_LABEL
	case '@':
		l.LabelType = parser.NODE_AT_LABEL
	}
	return l
}

// ld r, n / ld rr, nn の OP2 に指定された文字列を NumberObjectに変換する
func (e *Evaluator) stringObjToOp2(so *object.StringObject, ctx TContext) object.Object {
	str, err := util.Utf8ToShiftJis(so.Value)
	if err != nil {
		e.logger.Error(fmt.Sprintf(errcode.EDATA_ENCODE, so.Value), ctx)
		return object.ERROR
	}
	switch len(str) {
	case 1:
		return &object.NumberObject{Value: int(str[0])}
	case 2:
		return &object.NumberObject{Value: int(str[0])<<8 + int(str[1])}
	default:
		e.logger.Error(errcode.EZ80_OP2_STR, ctx)
		return object.ERROR
	}
}

// incbin charmap ファイル読み込み
func (e *Evaluator) readFile(name string) ([]byte, error) {
	dirs := make([]string, 0, len(e.incDirs)+1)

	dirs = append(dirs, filepath.Dir(name))
	dirs = append(dirs, e.incDirs...)

	base := filepath.Base(name)
	for _, dir := range dirs {
		path := filepath.Join(dir, base)
		if content, err := os.ReadFile(path); err == nil {
			return content, nil
		}
	}
	return nil, fmt.Errorf("cannot read %s", name)
}
