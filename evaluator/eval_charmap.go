package evaluator

import (
	"bytes"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/internal/util"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
)

// charmap 定義
func (e *Evaluator) evalCharmapStatement(stmt *parser.CharmapStatement, env TEnv) object.Object {
	e.concatenateSymbol(&stmt.Filename, env, stmt.Context)
	e.concatenateSymbol(&stmt.DefChar, env, stmt.Context)

	// Name のチェック
	if obj, ok := env.Get(stmt.NameID); ok {
		if obj.Type() == object.OBJ_SYMBOL {
			obj = obj.(*object.SymbolObject).Value
		}
		switch obj := obj.(type) {
		case *object.CharamapObject:
			if !obj.Context.Equal(stmt.Context) {
				e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_DUP, stmt.NameID), stmt.Context)
				return object.ERROR
			}
			// 再度評価せずに CharmapObject を返す
			return obj
		default:
			e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_USED, stmt.NameID), stmt.Context)
			return object.ERROR
		}
	}

	obj := e.evalExpression(stmt.Filename, env, stmt.Context)

	// json ファイル名を評価
	var filename string
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
		return obj
	case *object.NullObject:
		e.logger.Error(errcode.ECHARMAP_NULL, stmt.Context)
		return object.ERROR

	case *object.StringObject:
		filename = obj.Value

	default:
		e.logger.Error(errcode.ECHARMAP_NOT_STR, stmt.Context)
		return object.ERROR
	}

	// default char
	defChar := -1 // 未定義の場合エラー
	if stmt.DefChar != nil {
		obj := e.evalExpression(stmt.DefChar, env, stmt.Context)
	EVAL_AGAIN:
		switch eobj := obj.(type) {
		case *object.ErrorObject:
			return eobj
		case *object.RefNotFoundObject:
			return eobj
		case *object.NullObject:
			e.logger.Error(errcode.ECHARMAP_DEFCHAR_NULL, stmt.Context)
			return object.ERROR

		case *object.NumberObject:
			defChar = eobj.Value
		case *object.StringObject:
			obj = e.evalOneCharStringAsNumber(eobj.Value, stmt.Context)
			goto EVAL_AGAIN
		case *object.ArrayObject:
			obj = e.evalArrayToInt(eobj.Values, stmt.Context)
			goto EVAL_AGAIN

		default:
			e.logger.Error(errcode.ECHARMAP_DEFCHAR_VALUE, stmt.Context)
			return object.ERROR
		}
	}

	// json ファイルを cmap へ読み込み
	cmap := e.loadCharmapJson(filename, stmt)
	if cmap == nil {
		return object.ERROR
	}

	// 環境へ設定
	charmap := &object.CharamapObject{NameID: stmt.NameID, DefChar: defChar, Cmap: cmap, Context: stmt.Context}
	env.Set(stmt.NameID, charmap)

	return charmap
}

// charmap json ファイル読み込み
func (e *Evaluator) loadCharmapJson(filename string, stmt *parser.CharmapStatement) map[string][]byte {
	var text []byte
	if filename[0] == '{' {
		// filenanme を json テキストとして扱う
		text = []byte(filename)
	} else {
		// json ファイル path をソースファイル相対で決定
		path := e.getRelativeFilepath(stmt.Context.FileContent.Filename, filename)

		// ファイル読み込み
		var err error
		// text, err = os.ReadFile(path)
		text, err = e.readFile(stmt.Context.FileContent.Filename, path)
		if err != nil {
			e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_READ, filename, err.Error()), stmt.Context)
			return nil
		}
		// BOM があれば削除
		if bytes.HasPrefix(text, []byte{0xef, 0xbb, 0xbf}) {
			text = text[3:]
		}
	}

	// map 読み込み
	var rawMap map[string]any
	if err := json.Unmarshal(text, &rawMap); err != nil {
		e.logger.Error(errcode.ECHARMAP_JSON, stmt.Context)
		return nil
	}

	// rawMap -> cmap 変換
	cmap := map[string][]byte{}
	for k, v := range rawMap {
		va, ok := v.([]any)
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_FMT, k), stmt.Context)
			continue
		}
		ary := make([]byte, 0, len(va))
		for _, v := range va {
			if num, ok := v.(float64); ok {
				ary = append(ary, byte(num))
			} else {
				e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_FMT, k), stmt.Context)
			}
		}
		cmap[k] = ary
	}
	return cmap
}

// ソースファイルの位置からの相対ファイルパスを取得
func (e *Evaluator) getRelativeFilepath(base, target string) string {
	dir := filepath.Dir(base)
	return filepath.Join(dir, target)
}

func (e *Evaluator) applyCharmap(cmap *object.CharamapObject, expr *parser.FuncCallExpression, env TEnv, ctx TContext) object.Object {

	// 対象文字列を取得
	var str string

	obj := e.evalExpression(expr.Args.Expressions[0], env, ctx)
	switch obj := obj.(type) {
	case *object.ErrorObject:
		return obj
	case *object.RefNotFoundObject:
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

	// default char int ->  []byte
	defChar := cmap.DefChar
	var defCode []byte
	if 0 <= defChar && defChar <= 255 { // 0-255
		defCode = []byte{byte(defChar)}
	} else if defChar > 255 { // 256-
		word, ok := e.intToWord(defChar)
		if !ok {
			e.logger.Warning(fmt.Sprintf(errcode.WROUND_WORD, defChar, defChar), ctx)
		}
		defCode = []byte{byte((word >> 8) & 0xff), byte(word & 0xff)}
	}

	// str に適用
	bstr := make([]byte, 0, len(str))
	for _, c := range str {
		if v, ok := cmap.Cmap[string(c)]; ok {
			bstr = append(bstr, v...)
			continue
		}
		switch {
		case defChar >= 0:
			bstr = append(bstr, defCode...)
		case defChar == -2:
			if s, err := util.Utf8ToShiftJis(string(c)); err == nil {
				bstr = append(bstr, s...)
			} else {
				e.logger.Error(fmt.Sprintf(errcode.EDATA_ENCODE, string(c)), ctx)
				bstr = append(bstr, '?')
			}
		default:
			// defChar < 0 && defChar != -2 の場合、未定義エラーにする
			e.logger.Error(fmt.Sprintf(errcode.ECHARMAP_NOT_DEF, c), ctx)
			bstr = append(bstr, '?') // ? とする
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
