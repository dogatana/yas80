package evaluator

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

func (e *Evaluator) evalCharmapStatement(stmt *parser.CharmapStatement, env TEnv) object.Object {
	obj := e.evalExpression(stmt.Filename, env, stmt.Context)

	// json ファイル名を評価
	var filename string
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
		e.Resolved = false
		return obj
	case *object.NullObject:
		e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_NULL, stmt.Name), stmt.Context)
		return object.ERROR

	case *object.StringObject:
		filename = obj.Value

	default:
		e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_FILENAME, stmt.Name), stmt.Context)
		return object.ERROR
	}

	// default char
	defChar := 0
	if stmt.DefChar != nil {
		obj := e.evalExpression(stmt.DefChar, env, stmt.Context)
		switch obj := obj.(type) {
		case *object.ErrorObject:
			return obj
		case *object.RefNotFoundObject:
			e.Resolved = false
			return obj
		case *object.NullObject:
			e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_NULL, stmt.Name), stmt.Context)
			return object.ERROR

		case *object.NumberObject:
			defChar = obj.Value

		default:
			e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_FILENAME, stmt.Name), stmt.Context)
			return object.ERROR
		}
	}

	// json ファイル path をソースファイル相対で決定
	path := e.getRelativeFilepath(stmt.Context.FileContent.Filename, filename)

	// json ファイルを cmap へ読み込み
	cmap := e.loadCharmapJson(filename, path, stmt)
	if cmap == nil {
		return object.ERROR
	}

	// 環境へ設定
	charmap := &object.CharamapObject{Name: stmt.Name, DefChar: defChar, Cmap: cmap, Context: stmt.Context}
	env.Set(stmt.Name, charmap)

	return charmap
}

func (e *Evaluator) loadCharmapJson(filename, absPath string, stmt *parser.CharmapStatement) map[string][]byte {
	// ファイル読み込み
	text, err := os.ReadFile(absPath)
	if err != nil {
		e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_FILE, stmt.Name, filename, err.Error()), stmt.Context)
		return nil
	}
	// map 読み込み
	var rawMap map[string]any
	if err := json.Unmarshal(text, &rawMap); err != nil {
		e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_FILE, stmt.Name, filename, err.Error()), stmt.Context)
		return nil
	}

	// rawMap -> cmap 変換
	cmap := map[string][]byte{}
	for k, v := range rawMap {
		va, ok := v.([]any)
		if !ok {
			kverr := fmt.Sprintf("%q = %v", k, v)
			e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_JSON, stmt.Name, filename, kverr), stmt.Context)
			continue
		}
		ary := make([]byte, 0, len(va))
		for _, v := range va {
			if num, ok := v.(float64); ok {
				ary = append(ary, byte(num))
			} else {
				kverr := fmt.Sprintf("%q = %v", k, v)
				e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_JSON, stmt.Name, filename, kverr), stmt.Context)
			}
		}
		cmap[k] = ary
	}
	return cmap
}

func (e *Evaluator) getRelativeFilepath(base, target string) string {
	dir := filepath.Dir(base)
	return filepath.Join(dir, target)
}

func (e *Evaluator) applyCharmap(cmap *object.CharamapObject, expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {
	// len(expr.Args.Expression) は 1 or 2

	var str string // charmap の適用対象

	obj := e.evalExpression(expr.Args.Expressions[0], env, ctx)
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
		e.Resolved = false
		return obj
	case *object.NullObject:
		e.logger.Error(errcode.ECHARMAP_VALUE_NULL, ctx)
		return object.ERROR

	case *object.StringObject:
		str = obj.Value

	default:
		e.logger.Error(errcode.ECHARMAP_VALUE, ctx)
		return object.ERROR
	}

	defChar := 0 // default character
	if len(expr.Args.Expressions) == 2 {
		obj := e.evalExpression(expr.Args.Expressions[1], env, ctx)
		switch obj := obj.(type) {
		case *object.ErrorObject:
			return obj
		case *object.RefNotFoundObject:
			e.Resolved = false
			return obj
		case *object.NullObject:
			e.logger.Error(errcode.ECHARMAP_DEFCHAR_VALUE_NULL, ctx)
			return object.ERROR

		case *object.NumberObject:
			defChar = obj.Value

		default:
			e.logger.Error(errcode.ECHARMAP_DEFCHAR_VALUE, ctx)
			return object.ERROR
		}
	}

	// default char int ->  []byte
	var defCode []byte
	if -128 <= defChar && defChar <= 255 {
		defCode = []byte{byte(defChar)}
	} else {
		word, ok := e.intToWord(defChar)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, defChar, defChar), ctx)
		}
		defCode = []byte{byte(word & 0xff), byte((word >> 8) & 0xff)}
	}

	// str に適用
	bstr := make([]byte, 0, len(str))
	for _, c := range str {
		if v, ok := cmap.Cmap[string(c)]; ok {
			bstr = append(bstr, v...)
		} else {
			bstr = append(bstr, defCode...)
		}
	}

	// bstr を ArrayObject として返す
	ret := &object.ArrayObject{}
	ret.Values = make([]object.Object, len(bstr))
	for i, b := range bstr {
		ret.Values[i] = &object.NumberObject{Value: int(b)}
	}
	return ret
}
