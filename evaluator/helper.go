package evaluator

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"unicode/utf8"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/intern"
	"github.com/dogatana/yas80/internal/util"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
)

// 各種 object 判定

func isError(obj object.Object) bool {
	return obj.Type() == object.OBJ_ERROR
}

func isNumber(obj object.Object) bool {
	return obj.Type() == object.OBJ_NUMBER
}

func isString(obj object.Object) bool {
	return obj.Type() == object.OBJ_STRING
}

func isArray(obj object.Object) bool {
	return obj.Type() == object.OBJ_ARRAY
}

func isRefNotFound(obj object.Object) bool {
	return obj.Type() == object.OBJ_REF_NOTFOUND
}

func isNull(obj object.Object) bool {
	return obj.Type() == object.OBJ_NULL
}

// 依存先の識別子を抽出する
func mergeNameIDs(obj1, obj2 object.Object) []intern.SymbolID {
	ids1 := extractNameIDs(obj1)
	ids2 := extractNameIDs(obj2)

	uids := make(map[intern.SymbolID]struct{}, len(ids1)+len(ids2))
	for _, id := range ids1 {
		uids[id] = struct{}{}
	}
	for _, id := range ids2 {
		uids[id] = struct{}{}
	}

	ids := make([]intern.SymbolID, 0, len(uids))
	for id := range uids {
		ids = append(ids, id)
	}
	return ids
}

func extractNameIDs(obj object.Object) []intern.SymbolID {
	switch obj := obj.(type) {
	case *object.RefNotFoundObject:
		return obj.NameIDs
	case *object.SymbolObject:
		return []intern.SymbolID{obj.NameID}
	default:
		return []intern.SymbolID{}
	}
}

// $ location counter 更新
func setLocationCounter(env TEnv, addr int) {
	updateNumberInEnv(intern.ID_LOC, addr, env)
}

// $$ allocate location counter 更新
func setAllocLocationCounter(env TEnv, addr int) {
	updateNumberInEnv(intern.ID_ALOC, addr, env)
}

// $PASS 更新
func setPass(env TEnv, pass int) {
	updateNumberInEnv(intern.ID_PASS, pass, env)
}

// ENV[id] の NumberObject 更新
func updateNumberInEnv(id intern.SymbolID, value int, env TEnv) {
	if obj, ok := env.Get(id); !ok {
		panic(fmt.Sprintf("cannot get %s", id))
	} else if no, ok := obj.(*object.NumberObject); !ok {
		panic(fmt.Sprintf("ENV[%s] is not NumberObject", id))
	} else {
		no.Value = value
	}
}

// location counter 取得
func getLocationCounter(env TEnv) int {
	if addr, ok := getNumberFromEnv(intern.ID_LOC, env); !ok {
		panic("getLocationCounter failed")
	} else {
		return addr

	}
}

// location counter 更新
// set** を使用すると現階層の Env に設定するため、元のオブジェクトの値を直接書き換える
func advanceLocationCounters(env TEnv, n int) error {
	obj, ok := env.Get(intern.ID_LOC)
	if !ok {
		panic("getLocationCounter failed")
	}
	counter := obj.(*object.NumberObject)
	counter.Value += n

	// 64KB アドレス超過のチェック
	if counter.Value > 0x10000 {
		return fmt.Errorf(errcode.EADDR_OVERFLOW, counter.Value)
	}

	obj, ok = env.Get(intern.ID_ALOC)
	if !ok {
		panic("getAllocLocationCounter failed")
	}
	counter = obj.(*object.NumberObject)
	counter.Value += n

	// 64KB アドレス超過を許容するためチェックしない
	// if counter.Value > 0x10000 {
	// 	return fmt.Errorf(errcode.EALLOC_ADDR_OVERFLOW, counter.Value)
	// }

	return nil
}

// int -> byte へ。丸め発生した場合は false
func (e *Evaluator) intToByte(n int) (byte, bool) {
	return byte(n & 0xff), -128 <= n && n <= 255
}

// int -> word へ。丸めが発生した場合は false
func (e *Evaluator) intToWord(n int) (int, bool) {
	return (n & 0xffff), -32768 <= n && n <= 65535
}

// int ->  byte へ。丸めが発生した場合は false
func (e *Evaluator) intToPort(n int) (byte, bool) {
	return byte(n & 0xff), 0 <= n && n <= 255
}

// int ->  byte(0-7) へ。丸めが発生した場合は false
func (e *Evaluator) intToBit(n int) (byte, bool) {
	return byte(n & 0x7), 0 <= n && n <= 7
}

func boolToInt(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}

// 数値/文字列/レジスタ/フラグをTextObject に変換する
func toTextObject(obj object.Object, ctx TContext) object.Object {
	var text string
	switch v := obj.(type) {
	case *object.NumberObject:
		text = fmt.Sprintf("%04x(%d)", v.Value, v.Value)
	case *object.StringObject:
		text = fmt.Sprintf("%q", v.Value)
	case *object.RegisterObject:
		text = parser.TokenLiteral(v.Register)
	case *object.FlagObject:
		text = parser.TokenLiteral(v.Flag)
	}
	if text != "" {
		return &object.CodeTextObject{Text: text, Context: ctx}
	}

	// dummy リスト出力では無視する
	return &object.ValueObject{Value: obj, Context: ctx}
}

// シンボル結合処理
func (e *Evaluator) concatenateSymbol(ptr *parser.Expression, env TEnv, ctx TContext) bool {
	switch expr := (*ptr).(type) {
	case *parser.InfixExpression:
		if expr.Operator != parser.CONCAT {
			return e.concatenateSymbol(&expr.Op1, env, ctx) || e.concatenateSymbol(&expr.Op2, env, ctx)
		}
		ident, ok := expr.Op1.(*parser.Ident)
		if !ok {
			// parser 処理済みで段階でここには来ないはず
			panic("not ident")
		}
		{
			// copy &(*ident) では新しい値が生成されないため
			temp := *ident
			ident = &temp
		}

		suffix := ""
		op2 := e.evalExpression(expr.Op2, env, ctx)
		switch op2 := op2.(type) {
		case *object.ErrorObject:
			return false
		case *object.RefNotFoundObject:
			names := util.Map(op2.NameIDs, func(id intern.SymbolID) string { return id.String() })
			e.logger.Error(fmt.Sprintf(errcode.ESYM_UNDEF, strings.Join(names, ", ")), ctx)
			return false
		case *object.NumberObject:
			suffix = fmt.Sprintf("%d", op2.Value)
		case *object.StringObject:
			suffix = strings.ToUpper(op2.Value) // 文字列リテラルは大文字化して結合する
		default:
			e.logger.Error(errcode.ECONCAT_TYPE, ctx)
			return false
		}
		name := ident.NameID.String() + suffix
		ident.NameID = intern.InternString(name)
		*ptr = ident
		return true
	case *parser.PrefixExpression:
		return e.concatenateSymbol(&expr.Op, env, ctx)
	default:
		return false
	}
}

// ENV から SymbolObject を取得する
func (e *Evaluator) getSymbolFromEnv(name string, env TEnv) (*object.SymbolObject, bool) {
	names := strings.Split(name, ".")
	if len(names) == 1 {
		if obj, ok := env.Get(intern.InternString(name)); ok {
			switch obj := obj.(type) {
			case *object.SymbolObject:
				return obj, true
			case *object.ProcObject:
				return &object.SymbolObject{Name: name, Value: &object.NumberObject{Value: obj.Addr}}, true
			}
			return nil, false
		}
	}
	obj, ok := env.Get(intern.InternString(names[0]))
	if !ok {
		return nil, false
	}
	switch obj := obj.(type) {
	case *object.ProcObject:
		v, ok := obj.Get(intern.InternString("." + names[1]))
		if !ok {
			return nil, false
		}
		if sym, ok := v.(*object.SymbolObject); ok {
			return sym, ok
		}
		return nil, false

	case *object.EnumObject:
		v, ok := obj.Get(intern.InternString("." + names[1]))
		if !ok {
			return nil, false
		}
		if sym, ok := v.(*object.SymbolObject); ok {
			return sym, ok
		}
		return nil, false

	default:
		return nil, false
	}
}

// 参照用匿名ラベルを検索  呼ばれる前に isAnonRef なのはチェック済み
func (e *Evaluator) findAnonLabel(name string, env TEnv, ctx TContext) object.Object {
	var def string
	if len(name) == 2 {
		def = "@@" // @F, @B -> @@ を検索
	} else {
		def = name[:2] // @nF, @nB -> @n を検索
	}
	id := intern.InternString(def)
	obj, ok := env.Get(id)
	if !ok {
		if name[len(name)-1] == 'B' { // B の場合は定義済みのはず
			e.logger.Error(fmt.Sprintf(errcode.EANON_LABEL_NOT_FOUND, name), ctx)
			return object.ERROR
		}
		// 'F' で環境にないなら空で登録し下のチェック(stage 2)で失敗させるようにする
		obj := &object.AnonLabelsObject{Name: name, NameID: id, Labels: []*object.AnonLabel{}}
		env.Set(id, obj)

		// stage 2 のチェックへ移行させるため ture のままとする
		// e.Resolved = false
		return &object.RefNotFoundObject{}
	}

	labels := obj.(*object.AnonLabelsObject).Labels
	switch name[len(name)-1] {
	case 'B':
		// 逆順検索
		for i := len(labels) - 1; i >= 0; i-- {
			if labels[i].Filename == ctx.FileContent.Filename && labels[i].Line < int(ctx.Line) {
				return labels[i]
			}
		}
		// 逆順なので必ず見つかるはずで、そうでないならエラーとする
		e.logger.Error(fmt.Sprintf(errcode.EANON_LABEL_NOT_FOUND, name), ctx)
		return object.ERROR
	case 'F':
		// 順方向検索
		for i := 0; i < len(labels); i++ {
			if labels[i].Filename == ctx.FileContent.Filename && labels[i].Line > int(ctx.Line) {
				return labels[i]
			}
		}
		// Evaluator.Stage2 で見つからなければエラー
		if e.Stage2 {
			e.logger.Error(fmt.Sprintf(errcode.EANON_LABEL_NOT_FOUND, name), ctx)
			return object.ERROR
		}
		// stage 2 のチェックへ移行させるため ture のままとする
		// e.Resolved = false
		return &object.RefNotFoundObject{}
	default:
		panic(fmt.Sprintf("invalid anonymous label name: %s", name))
	}
}

// parser.Expression -> parser.Ident - >parser.Label を評価・環境登録し object.SymbolObject を返す
func (e *Evaluator) exprToLabel(expr parser.Expression, env TEnv, ctx TContext) object.Object {
	id, ok := expr.(*parser.Ident)
	if !ok {
		return object.ERROR
	}
	label := e.identToLabel(id)
	label.Context = ctx // Ident の Context でなく、引数（文）の Conext を設定する

	return e.evalLabel(label, env)
}

// parser.Ident -> parser.Label 変換(exprToLabel から呼ばれる)
func (e *Evaluator) identToLabel(id *parser.Ident) *parser.Label {
	l := &parser.Label{NameID: id.NameID, LabelType: parser.NODE_LABEL, Context: id.Context}
	name := id.NameID.String()
	switch {
	case util.IsAnonDef(name) || util.IsAnonRef(name):
		l.LabelType = parser.NODE_ANON_LABEL
	case name[0] == '.':
		l.LabelType = parser.NODE_LOCAL_LABEL
	case name[0] == '@':
		l.LabelType = parser.NODE_AT_LABEL
	}
	return l
}

// 1文字の文字列を数値に変換する
func (e *Evaluator) evalOneCharStringAsNumber(s string, ctx TContext) object.Object {
	r, size := utf8.DecodeLastRuneInString(s)
	if size == 0 || size != len(s) {
		e.logger.Error(errcode.ESTR_TO_INT_LEN, ctx)
		return object.ERROR
	} else if r == utf8.RuneError {
		e.logger.Error(fmt.Sprintf(errcode.EDATA_ENCODE, s), ctx)
		return object.ERROR
	}
	b, ok := util.TransformBytes(s)
	if !ok {
		e.logger.Error(fmt.Sprintf(errcode.EDATA_ENCODE, string(r)), ctx)
		return object.ERROR
	}

	var code int
	if len(b) == 1 {
		code = int(b[0])
	} else {
		code = int(b[0])*256 + int(b[1])
	}
	return &object.NumberObject{Value: code, Context: ctx}
}

// 1/2 要素の数値配列を数値に変換する
func (e *Evaluator) evalArrayToInt(values []object.Object, ctx TContext) object.Object {
	if len(values) < 1 || len(values) > 2 {
		e.logger.Error(fmt.Sprintf(errcode.EARRAY_TO_INT_LEN, len(values)), ctx)
		return object.ERROR

	}
	// op1 をByte値に変換
	v, ok := values[0].(*object.NumberObject)
	if !ok {
		e.logger.Error(errcode.EARRAY_TO_INT_TYPE, ctx)
		return object.ERROR
	}
	v1, ok := e.intToByte(v.Value)
	if !ok {
		e.logger.Error(fmt.Sprintf(errcode.WROUND_BYTE, v.Value, v.Value), ctx)
	}

	if len(values) == 1 {
		return &object.NumberObject{Value: int(v1), Context: ctx}
	}

	// op2 をByte値に変換
	v, ok = values[1].(*object.NumberObject)
	if !ok {
		e.logger.Error(errcode.EARRAY_TO_INT_TYPE, ctx)
		return object.ERROR
	}
	v2, ok := e.intToByte(v.Value)
	if !ok {
		e.logger.Error(fmt.Sprintf(errcode.WROUND_BYTE, v.Value, v.Value), ctx)
		return object.ERROR
	}
	return &object.NumberObject{Value: int(v1)*256 + int(v2), Context: ctx}
}

// incbin charmap ファイル読み込み
func (e *Evaluator) readFile(from, name string) ([]byte, error) {
	dirs := make([]string, 0, len(e.incDirs)+1)

	dirs = append(dirs, filepath.Dir(from))
	dirs = append(dirs, e.incDirs...)

	base := filepath.Base(name)
	for _, dir := range dirs {
		path := filepath.Join(dir, base)
		if content, err := os.ReadFile(path); err == nil {
			return content, nil
		}
	}
	return nil, fmt.Errorf(errcode.EFILE_NOT_FOUND, name)
}

// ENV から NumberObject を取得する（システム変数用）
func getNumberFromEnv(id intern.SymbolID, env TEnv) (int, bool) {
	obj, ok := env.Get(id)
	if !ok {
		return 0, false
	}
EVAL_AGAIN:
	switch no := obj.(type) {
	case *object.NumberObject:
		return no.Value, true
	case *object.SymbolObject:
		obj = no.Value
		goto EVAL_AGAIN
	default:
		return 0, false
	}
}
