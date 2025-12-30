package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) evalDataStoreStatement(stmt *parser.DataStoreStatement, env object.Environment) object.Object {
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
		e.logger.Error(errcode.EDATA_COUNT, stmt.Context)
		return object.ERROR
	}
	if count <= 0 {
		e.logger.Error(errcode.EDATA_COUNT, stmt.Context)
		return object.ERROR
	}

	// default filler
	var filler int
	if stmt.Size == 1 {
		if obj, ok := env.Get("$FILL_BYTE"); !ok {
			panic("no $FILL_BYTE")
		} else {
			filler = obj.(*object.NumberObject).Value
		}
	} else {
		if obj, ok := env.Get("$FILL_WORD"); !ok {
			panic("no $FILL_BYTE")
		} else {
			filler = obj.(*object.NumberObject).Value
		}
	}

	if stmt.Filler != nil {
		// 文で指定した filler
		obj = e.evalExpression(stmt.Filler, env, stmt.Context)
		switch obj := obj.(type) {
		case *object.ErrorObject:
			return object.ERROR
		case *object.RefNotFoundObject:
			// do nothing デフォルトを使用
		case *object.NumberObject:
			filler = obj.Value
		default:
			e.logger.Error(errcode.EDATA_FILL, stmt.Context)
			return object.ERROR
		}
	}

	if stmt.Size == 1 && !e.isByteRange(filler) {
		e.logger.Warning(fmt.Sprintf(errcode.WROUND_BYTE, filler, filler), stmt.Context)
		filler &= 0xff
	}
	if stmt.Size == 2 && !e.isWordRange(filler) {
		e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, filler, filler), stmt.Context)
		filler &= 0xffff
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
	return &object.CodeObject{Code: data, Addr: addr, Line: stmt.Context.Line}
}
