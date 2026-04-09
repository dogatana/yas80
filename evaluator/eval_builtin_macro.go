package evaluator

import (
	"fmt"
	"strings"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/logging"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
)

// 組み込みマクロかどうか
func isBuiltinMacroName(name string) bool {
	builtinMacroNames := map[string]bool{
		"ENTRY":    true,
		"ALIGN":    true,
		"ERROR":    true,
		"WARN":     true,
		"INFO":     true,
		"INCBIN":   true,
		"BINCLUDE": true,
		"SETMAP":   true,
		"LIST":     true,
		"CHECK256": true,
	}
	return builtinMacroNames[name]
}

// 組み込みマクロの実行
func (e *Evaluator) evalBuiltinMacro(stmt *parser.MacroCallStatement, env TEnv) (object.Object, bool) {
	switch stmt.Name {
	case "ENTRY":
		return e.ebMacroEntry(stmt, env), true
	case "ALIGN":
		return e.ebMacroAlign(stmt, env), true
	case "ERROR":
		return e.ebMacroLogMessage(logging.Err, stmt, env), true
	case "WARN":
		return e.ebMacroLogMessage(logging.Warn, stmt, env), true
	case "INFO":
		return e.ebMacroLogMessage(logging.Info, stmt, env), true
	case "INCBIN", "BINCLUDE":
		return e.ebMacroIncBin(stmt, env), true
	case "SETMAP":
		return e.ebMacroSetMap(stmt, env), true
	case "LIST":
		return e.ebMacroList(stmt, env), true
	case "CHECK256":
		return e.ebMacroCheck256(stmt, env), true
	}
	return object.NULL, false

}

// entry
func (e *Evaluator) ebMacroEntry(stmt *parser.MacroCallStatement, env TEnv) object.Object {
	args := stmt.Args.Expressions

	if len(args) != 1 {
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_COUNT, stmt.Name), stmt.Context)
		return object.ERROR
	}

	obj := e.evalExpression(args[0], env, stmt.Context)
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
		return obj
	case *object.NullObject:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_NULL, stmt.Name), stmt.Context)
		return object.ERROR

	case *object.NumberObject:
		addr, ok := e.intToWord(obj.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, obj.Value, obj.Value), stmt.Context)
		}
		return &object.EntryObject{StartAddr: addr}

	default:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_VALUE, stmt.Name), stmt.Context)
		return object.ERROR
	}
}

// align n, fill
func (e *Evaluator) ebMacroAlign(stmt *parser.MacroCallStatement, env TEnv) object.Object {
	args := stmt.Args.Expressions

	if len(args) < 1 || len(args) > 2 {
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_COUNT, stmt.Name), stmt.Context)
		return object.ERROR
	}

	// 第1オペランド
	var num int
	obj := e.evalExpression(args[0], env, stmt.Context)
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
		return obj
	case *object.NullObject:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_NULL, stmt.Name), stmt.Context)
		return object.ERROR

	case *object.NumberObject:
		num = obj.Value

	default:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_VALUE, stmt.Name), stmt.Context)
		return object.ERROR
	}

	var fill byte

	if len(args) == 2 {
		// あれば第2オペランド
		obj = e.evalExpression(args[1], env, stmt.Context)
	} else {
		// なければ $FILL
		obj, _ = env.Get("$FILL")
	}
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
		return obj
	case *object.NullObject:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_NULL, stmt.Name), stmt.Context)
		return object.ERROR

	case *object.NumberObject:
		b, ok := e.intToByte(obj.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_BYTE, obj.Value, obj.Value), stmt.Context)
			args[1] = &parser.NumberLiteral{Value: int(b), Context: stmt.Context}
		}
		fill = b

	default:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_VALUE, stmt.Name), stmt.Context)
	}

	addr := getLocationCounter(env)
	mod := addr % num
	if mod == 0 {
		return &object.CodeObject{Code: []byte{}, Addr: addr, Context: stmt.Context}
	}

	size := num - mod
	code := make([]byte, size)
	for i := 0; i < size; i++ {
		code[i] = fill
	}

	if err := advanceLocationCounters(env, size); err != nil {
		e.logger.Error(err.Error(), stmt.Context)
	}
	return &object.CodeObject{Addr: addr, Code: code, Filled: true, Context: stmt.Context}
}

// error/warn/info
func (e *Evaluator) ebMacroLogMessage(msgType int, stmt *parser.MacroCallStatement, env TEnv) object.Object {

	args := stmt.Args.Expressions

	if len(args) != 1 {
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_COUNT, stmt.Name), stmt.Context)
		return object.ERROR
	}

	var msg string

	obj := e.evalExpression(args[0], env, stmt.Context)
	switch obj := obj.(type) {
	case *object.StringObject:
		msg = obj.Value // 文字列引数
	case *object.NumberObject:
		msg = fmt.Sprintf("%d", obj.Value) // 数値引数を文字列化
	case *object.RegisterObject:
		msg = parser.TokenLiteral(obj.Register) // レジスタ引数を文字列化
	case *object.FlagObject:
		msg = parser.TokenLiteral(obj.Flag) // フラグ引数を文字列化

	case *object.RefNotFoundObject:
		return obj // そのまま返す

	case *object.NullObject:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_NULL, stmt.Name), stmt.Context)
		return object.ERROR
	default:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_VALUE, stmt.Name), stmt.Context)
		return object.ERROR
	}

	switch msgType {
	case logging.Err:
		e.logger.Error(msg, stmt.Context)
	case logging.Warn:
		e.logger.Warning(msg, stmt.Context)
	case logging.Info:
		e.logger.Info(msg, stmt.Context)
	}
	return object.NULL
}

// incbin "file", [start, [length]]
func (e *Evaluator) ebMacroIncBin(stmt *parser.MacroCallStatement, env TEnv) object.Object {
	// $RSIZE 初期化
	if obj, ok := env.Get("$RSIZE"); !ok {
		env.Set("$RSIZE", &object.NumberObject{Value: -1})
	} else {
		obj.(*object.NumberObject).Value = -1
	}

	args := stmt.Args.Expressions
	if len(args) < 1 || len(args) > 3 {
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_COUNT, stmt.Name), stmt.Context)
		return object.ERROR
	}

	var file string // 読み込みファイル名
	obj := e.evalMacroStringArg(stmt.Name, args[0], env, stmt.Context)
	if obj, ok := obj.(*object.StringObject); !ok {
		return obj
	} else {
		file = obj.Value
	}

	start := 0 // 読み込み開始位置
	if len(args) > 1 {
		obj := e.evalMacroNumberArg(stmt.Name, args[1], env, stmt.Context)
		if obj, ok := obj.(*object.NumberObject); !ok {
			return obj
		} else {
			start = obj.Value
		}
	}
	size := 0 // 読み込みサイズ
	if len(args) > 2 {
		obj := e.evalMacroNumberArg(stmt.Name, args[2], env, stmt.Context)
		if obj, ok := obj.(*object.NumberObject); !ok {
			return obj
		} else {
			size = obj.Value
		}
	}

	data, err := e.readFile(stmt.Context.FileContent.Filename, file)
	if err != nil {
		e.logger.Error(err.Error(), stmt.Context)
		return object.ERROR
	}

	// start, size に合わせて data から code を作成
	code := []byte{}
	switch {
	case start == 0 && size == 0:
		code = data
	case start != 0 && size == 0:
		if start < len(data) {
			code = data[start:]
		}
	default:
		if start < len(data) {
			code = data[start:]
			if len(code) > size {
				code = code[:size]
			}
		}
	}
	co := &object.CodeObject{Addr: getLocationCounter(env), Code: code, IncBin: true, Context: stmt.Context}

	// $RSIZE 更新
	o, _ := env.Get("$RSIZE")
	o.(*object.NumberObject).Value = co.Size()

	if err := advanceLocationCounters(env, co.Size()); err != nil {
		e.logger.Error(err.Error(), stmt.Context)
	}

	return co
}

// setmap "c", [b1, b2]
func (e *Evaluator) ebMacroSetMap(stmt *parser.MacroCallStatement, env TEnv) object.Object {

	args := stmt.Args.Expressions
	if len(args) != 3 {
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_COUNT, stmt.Name), stmt.Context)
		return object.ERROR
	}

	// charmap の取得
	var cmap *object.CharamapObject
	obj := e.evalExpression(args[0], env, stmt.Context)
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
		return obj
	case *object.NullObject:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_NULL, stmt.Name), stmt.Context)
		return object.ERROR

	case *object.CharamapObject:
		cmap = obj

	default:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_VALUE, stmt.Name), stmt.Context)
		return object.ERROR
	}

	var str string
	obj = e.evalMacroStringArg(stmt.Name, args[1], env, stmt.Context)
	if obj, ok := obj.(*object.StringObject); !ok {
		return obj
	} else {
		str = obj.Value
	}

	// 文字数チェック
	runes := []rune(str)
	if len(runes) != 1 {
		e.logger.Error(errcode.ESETMAP_NOT_A_CHAR, stmt.Context)
		return object.ERROR
	}

	// 配列取得
	var values []byte

	obj = e.evalExpression(args[2], env, stmt.Context)
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
		return obj
	case *object.NullObject:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_NULL, stmt.Name), stmt.Context)
		return object.ERROR

	case *object.ArrayObject:
		// 配列が空
		if len(obj.Values) == 0 {
			e.logger.Error(errcode.ESETMAP_ARRAY_EMPTY, stmt.Context)
			return object.ERROR
		}
		values = make([]byte, len(obj.Values))
		for i, v := range obj.Values {
			if v, ok := v.(*object.NumberObject); ok {
				values[i] = byte(v.Value & 0xff)
			} else {
				e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_VALUE, stmt.Name), stmt.Context)
				return object.ERROR
			}
		}

	default:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_VALUE, stmt.Name), stmt.Context)
		return object.ERROR
	}

	// 置き換え実施
	cmap.Cmap[str] = values

	atext := strings.ReplaceAll(fmt.Sprint(values), " ", ", ")
	return &object.TextObject{Text: fmt.Sprintf("%q:%s", str, atext), Context: stmt.Context}
}

// list 制御 1/0
func (e *Evaluator) ebMacroList(stmt *parser.MacroCallStatement, env TEnv) object.Object {
	args := stmt.Args.Expressions

	if len(args) != 1 {
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_COUNT, stmt.Name), stmt.Context)
		return object.ERROR
	}

	obj := e.evalMacroNumberArg(stmt.Name, args[0], env, stmt.Context)
	no, ok := obj.(*object.NumberObject)
	if !ok {
		return obj
	}
	v := no.Value
	if v != 0 && v != 2 {
		v = 1
	}
	return &object.ListControl{Value: v, Context: stmt.Context}
}

// check256 base
func (e *Evaluator) ebMacroCheck256(stmt *parser.MacroCallStatement, env TEnv) object.Object {
	args := stmt.Args.Expressions

	if len(args) != 1 {
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_COUNT, stmt.Name), stmt.Context)
		return object.ERROR
	}

	// アドレスに関係するので Stage2 で実行する
	if !e.Stage2 {
		return object.NULL
	}

	// base
	obj := e.evalMacroNumberArg(stmt.Name, args[0], env, stmt.Context)
	no, ok := obj.(*object.NumberObject)
	if !ok {
		return obj
	}
	base := no.Value

	// $
	addr := getLocationCounter(env)

	if addr > (base&0xff00)+0x0100 {
		e.logger.Error(fmt.Sprintf(errcode.ECHECK256_ERROR, base, addr), stmt.Context)
	}
	return object.NULL
}

// マクロ引数を StringObject として評価
func (e *Evaluator) evalMacroStringArg(name string, expr parser.Expression, env TEnv, ctx TContext) object.Object {
	obj := e.evalExpression(expr, env, ctx)
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
		return obj
	case *object.NullObject:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_NULL, name), ctx)
		return object.ERROR

	case *object.StringObject:
		return obj

	default:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_VALUE, name), ctx)
		return object.ERROR
	}
}

// マクロ引数を NumberObject として評価
func (e *Evaluator) evalMacroNumberArg(name string, expr parser.Expression, env TEnv, ctx TContext) object.Object {
	obj := e.evalExpression(expr, env, ctx)
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
		return obj
	case *object.NullObject:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_NULL, name), ctx)
		return object.ERROR

	case *object.NumberObject:
		return obj

	default:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_VALUE, name), ctx)
		return object.ERROR
	}
}
