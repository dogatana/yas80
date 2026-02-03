package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/filecontent"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) evalDataStoreStatement(stmt *parser.DataStoreStatement, env TEnv) object.Object {
	e.concatenateSymbol(&stmt.Label, env, stmt.Context)
	e.concatenateSymbol(&stmt.Count, env, stmt.Context)
	e.concatenateSymbol(&stmt.FillValue, env, stmt.Context)

	// label
	if stmt.Label != nil {
		obj := e.exprToLabel(stmt.Label, env, stmt.Context)
		if isError(obj) {
			return object.ERROR
		}
	}

	// count
	obj := e.evalExpression(stmt.Count, env, stmt.Context)
	var count int
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return object.ERROR
	case *object.RefNotFoundObject:
		count = 1
	case *object.NumberObject:
		count = obj.Value
	default:
		e.logger.Error(errcode.EDS_COUNT, stmt.Context)
		return object.ERROR
	}
	if count <= 0 {
		e.logger.Error(errcode.EDS_COUNT, stmt.Context)
		return object.ERROR
	}

	// default filler
	var filler int

	if stmt.FillValue == nil {
		// 指定がなければ $FILL を使用
		if obj, ok := env.Get("$FILL"); !ok {
			panic("no $FILL")
		} else {
			filler = obj.(*object.NumberObject).Value
		}
		filler &= 0xff
		if stmt.Size == 2 { // WORD とする
			filler = filler<<8 + filler
		}
	} else {
		// 文で指定した filler
		obj = e.evalExpression(stmt.FillValue, env, stmt.Context)
		switch obj := obj.(type) {
		case *object.ErrorObject:
			return object.ERROR
		case *object.RefNotFoundObject:
			// $FILL ベースの値を利用
		case *object.NumberObject:
			filler = obj.Value
		default:
			e.logger.Error(errcode.EDS_FILL, stmt.Context)
			return object.ERROR
		}
	}

	switch stmt.Size {
	case 1:
		filler, ok := e.intToByte(filler)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_BYTE, filler, filler), stmt.Context)
		}
	case 2:
		filler, ok := e.intToWord(filler)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, filler, filler), stmt.Context)
		}
	}

	data := make([]byte, count*stmt.Size)

	for i := 0; i < count; i++ {
		if stmt.Size == 1 {
			data[i] = byte(filler)
		} else {
			data[i*2] = byte(filler & 0xff)
			data[i*2+1] = byte(filler >> 8)
		}
	}
	addr := getLocationCounter(env)
	return &object.CodeObject{Code: data, Addr: addr, Context: stmt.Context}
}

func (e *Evaluator) evalDataStatement(stmt *parser.DataStatement, env TEnv) object.Object {
	e.concatenateSymbol(&stmt.Label, env, stmt.Context)
	for i := range stmt.Values {
		e.concatenateSymbol(&stmt.Values[i], env, stmt.Context)
	}

	// label
	if stmt.Label != nil {
		obj := e.exprToLabel(stmt.Label, env, stmt.Context)
		if isError(obj) {
			return object.ERROR
		}
	}

	count := len(stmt.Values)
	var code []byte
	if stmt.Size == 2 {
		code = make([]byte, 0, count*2)
	} else {
		code = make([]byte, 0, count)
	}

	if e.test(code, stmt.Size, stmt.Values, env, stmt.Context) == nil {
		return object.ERROR
	}
	fmt.Printf("code: %v\n", code)
	return &object.CodeObject{Code: code, Addr: getLocationCounter(env), Context: stmt.Context}
}

func (e *Evaluator) test(code []byte, size int, exprs []parser.Expression, env object.Environment, ctx *filecontent.Context) []byte {

	for _, expr := range exprs {
		obj := e.evalExpression(expr, env, ctx)

		switch obj := obj.(type) {
		case *object.ErrorObject:
			return nil

		case *object.RefNotFoundObject:
			code = append(code, 0)

		case *object.NumberObject:
			e.numberToCode(obj, code, size, ctx)

		case *object.StringObject:
			if size == 2 {
				e.logger.Error(errcode.EDATA_DW_STR, ctx)
				return nil
			}
			if e.stringToCode(obj, code, size, ctx) == nil {
				return nil
			}

		case *object.ArrayObject:
			for _, v := range obj.Values {
				switch v := v.(type) {
				case *object.NumberObject:
					code = e.numberToCode(v, code, size, ctx)
				case *object.StringObject:
					if size == 2 {
						e.logger.Error(errcode.EDATA_DW_STR, ctx)
						return nil
					}
					code = e.stringToCode(v, code, size, ctx)
					if code == nil {
						return nil
					}
				default:
					e.logger.Error(errcode.EDATA_VALUE, ctx)
					return nil
				}
			}
		default:
			e.logger.Error(errcode.EDATA_VALUE, ctx)
			return nil
		}
	}
	return code
}

func (e *Evaluator) numberToCode(obj *object.NumberObject, code []byte, size int, ctx TContext) []byte {
	v := obj.Value
	switch {
	case size == 1 || size == 0 && -128 <= v && v <= 127:
		v, ok := e.intToByte(v)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_BYTE, v, v), ctx)
		}
		code = append(code, v)
	case size == 2 || size == 0 || obj.ForceWord:
		v, ok := e.intToWord(obj.Value)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, v, v), ctx)
		}
		code = append(code, byte(v&0xff), byte((v>>8)&0xff))
	}
	return code
}

func (e *Evaluator) stringToCode(obj *object.StringObject, code []byte, size int, ctx TContext) []byte {
	if e.stringToCode(obj, code, size, ctx) == nil {
		return nil
	}
	if s, err := e.utf8ToShiftJis(obj.Value); err != nil {
		e.logger.Error(fmt.Sprintf(errcode.EDATA_ENCODE, obj.Value), ctx)
		return nil
	} else {
		code = append(code, s...)
	}
	return code
}
