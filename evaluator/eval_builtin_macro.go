package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

// 組み込みマクロの実行
func (e *Evaluator) evalBuiltinMacro(stmt *parser.MacroCallStatement, env TEnv) (object.Object, bool) {
	switch stmt.Name {
	case "ALIGN":
		return e.ebMacroAlign(stmt, env), true
	case "ERROR":
		return e.ebMacroLogMessage(logging.Err, stmt, env), true
	case "WARN":
		return e.ebMacroLogMessage(logging.Warn, stmt, env), true
	case "INFO":
		return e.ebMacroLogMessage(logging.Info, stmt, env), true
	case "INCBIN":
		return e.ebMacroIncBin(stmt, env), true
	}
	return object.NULL, false

}

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

	if err := advanceLocationCounter(env, size); err != nil {
		e.logger.Error(err.Error(), stmt.Context)
	}
	return &object.CodeObject{Addr: addr, Code: code, Filled: true, Context: stmt.Context}
}

func (e *Evaluator) ebMacroLogMessage(msgType int, stmt *parser.MacroCallStatement, env TEnv) object.Object {

	args := stmt.Args.Expressions

	if len(args) != 1 {
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_COUNT, stmt.Name), stmt.Context)
		return object.ERROR
	}

	var msg string
	obj := e.evalMacroStringArg(stmt.Name, args[0], env, stmt.Context)
	if obj, ok := obj.(*object.StringObject); !ok {
		return obj
	} else {
		msg = obj.Value
	}

	switch msgType {
	case logging.Err:
		e.logger.Error(msg, stmt.Context)
	case logging.Warn:
		e.logger.Warning(msg, stmt.Context)
	case logging.Info:
		e.logger.Info(msg, stmt.Context)
	}
	return object.NULL // エラーではないので NULL を返す
}

// incbin "file", [start, [length]]
func (e *Evaluator) ebMacroIncBin(stmt *parser.MacroCallStatement, env TEnv) object.Object {
	env.Set("$RSIZE", &object.NumberObject{Value: -1})

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

	data, err := e.readFile(file)
	if err != nil {
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
	env.Set("$RSIZE", &object.NumberObject{Value: co.Size()})
	if err := advanceLocationCounter(env, co.Size()); err != nil {
		e.logger.Error(err.Error(), stmt.Context)
	}

	return co
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
