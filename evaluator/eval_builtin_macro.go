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

	var str string
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
		str = fmt.Sprintf("%d", obj.Value)

	case *object.StringObject:
		str = obj.Value

	default:
		e.logger.Error(fmt.Sprintf(errcode.EEBMAC_ARG_VALUE, stmt.Name), stmt.Context)
		return object.ERROR
	}

	switch msgType {
	case logging.Err:
		e.logger.Error(str, stmt.Context)
	case logging.Warn:
		e.logger.Warning(str, stmt.Context)
	case logging.Info:
		e.logger.Info(str, stmt.Context)
	}
	return object.NULL // エラーではないので NULL を返す
}
