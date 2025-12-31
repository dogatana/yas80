package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/fileblock"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) ident2Label(id *parser.Ident) *parser.Label {
	l := &parser.Label{Name: id.Name, LabelType: parser.NODE_LABEL, Context: id.Context}
	switch id.Name[0] {
	case '.':
		l.LabelType = parser.NODE_LOCAL_LABEL
	case '@':
		l.LabelType = parser.NODE_AT_LABEL
	}
	return l
}

func (e *Evaluator) expr2Label(expr parser.Expression, env object.Environment, ctx *fileblock.Context) object.Object {
	id, ok := expr.(*parser.Ident)
	if !ok {
		e.logger.Error(errcode.ELABEL_EXPR, ctx)
		return object.ERROR
	}
	label := e.ident2Label(id)
	return e.evalLabel(label, env)
}

func (e *Evaluator) evalDataStoreStatement(stmt *parser.DataStoreStatement, env object.Environment) object.Object {
	e.concatenateSymbol(&stmt.Label, env, stmt.Context)
	e.concatenateSymbol(&stmt.Count, env, stmt.Context)
	e.concatenateSymbol(&stmt.FillValue, env, stmt.Context)

	// label
	if stmt.Label != nil {
		obj := e.expr2Label(stmt.Label, env, stmt.Context)
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

	if stmt.FillValue != nil {
		// 文で指定した filler
		obj = e.evalExpression(stmt.FillValue, env, stmt.Context)
		switch obj := obj.(type) {
		case *object.ErrorObject:
			return object.ERROR
		case *object.RefNotFoundObject:
			// do nothing デフォルトを使用
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
	return &object.CodeObject{Code: data, Addr: addr, Line: stmt.Context.Line}
}

func (e *Evaluator) evalDataStatement(stmt *parser.DataStatement, env object.Environment) object.Object {
	e.concatenateSymbol(&stmt.Label, env, stmt.Context)
	for i := range stmt.Values {
		e.concatenateSymbol(&stmt.Values[i], env, stmt.Context)
	}

	// label
	if stmt.Label != nil {
		obj := e.expr2Label(stmt.Label, env, stmt.Context)
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

	var value int
	for _, expr := range stmt.Values {
		obj := e.evalExpression(expr, env, stmt.Context)
		switch obj := obj.(type) {
		case *object.ErrorObject:
			return object.ERROR
		case *object.RefNotFoundObject:
			value = 0
			e.Resolved = false
		case *object.NumberObject:
			value = obj.Value
		case *object.StringObject:
			if stmt.Size == 2 {
				e.logger.Error(errcode.EDATA_DW_STR, stmt.Context)
				return object.ERROR
			}
			if s, err := e.utf8ToShiftJis(obj.Value); err != nil {
				e.logger.Error(fmt.Sprintf(errcode.EDATA_ENCODE, obj.Value), stmt.Context)
				return object.ERROR
			} else {
				code = append(code, s...)
			}
			continue
		}
		switch stmt.Size {
		case 0:
			if -128 <= value && value <= 255 {
				code = append(code, byte(value))
			} else {
				v, ok := e.intToWord(value)
				if !ok {
					e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, value, value), stmt.Context)
				}
				code = append(code, byte(v&0xff), byte(v>>8))
			}
		case 1:
			v, ok := e.intToByte(value)
			if !ok {
				e.logger.Warning(fmt.Sprintf(errcode.WROUND_BYTE, value, value), stmt.Context)
			}
			code = append(code, v)
		case 2:
			v, ok := e.intToWord(value)
			if !ok {
				e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, value, value), stmt.Context)
			}
			code = append(code, byte(v&0xff), byte(v>>8))
		}
	}
	return &object.CodeObject{Code: code, Addr: getLocationCounter(env), Line: stmt.Context.Line}
}
