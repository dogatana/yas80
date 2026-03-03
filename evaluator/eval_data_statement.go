package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/filecontent"
	"yas80/internal/util"
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
	return &object.CodeObject{Addr: addr, Code: data, Filled: true, Context: stmt.Context}
}

func (e *Evaluator) evalDataDefineStatement(stmt *parser.DataDefineStatement, env TEnv) object.Object {
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

	code := e.evalDataStatementSub(stmt.Size, stmt.Values, env, stmt.Context)
	if code == nil {
		return object.ERROR
	}
	return &object.CodeObject{Code: code, Addr: getLocationCounter(env), Context: stmt.Context}
}

func (e *Evaluator) evalDataStatementSub(size int, exprs []parser.Expression, env object.Environment, ctx *filecontent.Context) []byte {
	var code []byte

	for _, expr := range exprs {
		obj := e.evalExpression(expr, env, ctx)

		switch obj := obj.(type) {
		case *object.ErrorObject:
			return nil

		case *object.RefNotFoundObject:
			code = append(code, 0)

		case *object.NumberObject:
			if c := e.numberToCode(obj, size, ctx); c == nil {
				return nil
			} else {
				code = append(code, c...)
			}

		case *object.StringObject:
			if c := e.stringToCode(obj, size, ctx); c == nil {
				return nil
			} else {
				code = append(code, c...)
			}

		case *object.ArrayObject:
			if len(obj.Values) == 0 {
				e.logger.Error(errcode.EARRAY_EMPTY, ctx)
				return nil
			}
			for _, v := range obj.Values {
				switch v := v.(type) {
				case *object.NumberObject:
					if c := e.numberToCode(v, size, ctx); c == nil {
						return nil
					} else {
						code = append(code, c...)
					}
				case *object.StringObject:
					if c := e.stringToCode(v, size, ctx); c == nil {
						return nil
					} else {
						code = append(code, c...)
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

func (e *Evaluator) numberToCode(obj *object.NumberObject, size int, ctx TContext) []byte {
	v := obj.Value
	switch {
	case size == 1 && !obj.ForceWord || size == 0 && !obj.ForceWord && -128 <= v && v <= 255:
		if b, ok := e.intToByte(v); ok {
			return []byte{b}
		} else {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_BYTE, v, v), ctx)
			return []byte{b}
		}

	default:
		if w, ok := e.intToWord(obj.Value); ok {
			return []byte{byte(w & 0xff), byte((w >> 8) & 0xff)}
		} else {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, v, v), ctx)
			return []byte{byte(w & 0xff), byte((w >> 8) & 0xff)}
		}
	}
}

func (e *Evaluator) stringToCode(obj *object.StringObject, size int, ctx TContext) []byte {
	// sjis 1バイト目か？
	var isJis1st = func(c byte) bool {
		return (0x81 <= c && c <= 0x9f) || (0xe0 <= c && c <= 0xfc)
	}

	// utf8 -> sjis
	s, err := util.Utf8ToShiftJis(obj.Value)
	if err != nil {
		e.logger.Error(fmt.Sprintf(errcode.EDATA_ENCODE, obj.Value), ctx)
		return nil
	}
	// db, dd ならそのまま返す
	if size != 2 {
		return s
	}

	// dw なら文字コードをwordとして返す
	word := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		if isJis1st(s[i]) && i+1 < len(s) {
			word = append(word, s[i+1], s[i])
			i++
		} else {
			word = append(word, s[i], 0)
		}
	}
	return word
}
