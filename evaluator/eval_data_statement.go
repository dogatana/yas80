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

func (e *Evaluator) expr2Label(ptr *parser.Expression, env object.Environment, ctx *fileblock.Context) object.Object {
	e.concatenateSymbol(ptr, env, ctx)
	expr := *ptr
	id, ok := expr.(*parser.Ident)
	if !ok {
		e.logger.Error(errcode.ELABEL_EXPR, ctx)
		return object.ERROR
	}
	label := e.ident2Label(id)
	return e.evalLabel(label, env)
}

func (e *Evaluator) evalDataStoreStatement(stmt *parser.DataStoreStatement, env object.Environment) object.Object {
	// label
	if stmt.Label != nil {
		obj := e.expr2Label(&stmt.Label, env, stmt.Context)
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
