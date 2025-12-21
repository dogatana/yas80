package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/object"
	"yas80/parser"
)

// evalStatement
func (e *Evaluator) evalStatement(node parser.Node, env object.Environment) object.Object {

	switch node := node.(type) {

	// Z80 命令
	case *parser.Z80Instruction:
		obj := e.evalZ80Instruction(node, env)
		if obj.Type() == object.CODE_OBJ {
			code := obj.(*object.CodeObject)
			code.Addr = getLocationCounter(env)
			advanceLocationCounter(env, code.Size())
		}
		return obj

	// ラベル定義
	case *parser.LabelStatement:
		return e.evalLabelStatement(node, env)
	// 定数定義
	case *parser.ConstStatement:
		return e.evalConstStatement(node, env)

	// マクロ定義
	case *parser.MacroStatement:
		name := node.Name
		if name[0] == '@' || name[0] == '.' {
			e.logger.Error(fmt.Sprintf(errcode.EMACRO_NAME, name), node.Context)
			return object.ERROR
		}
		if obj, ok := env.Get(name); ok {
			if obj.Type() == object.MACRO_OBJ {
				e.logger.Error(fmt.Sprintf(errcode.EMACRO_DUP, name), node.Context)
			} else {
				e.logger.Error(fmt.Sprintf(errcode.EMACRO_USED, name), node.Context)
			}
			return object.ERROR
		}
		obj := &object.MacroObject{Name: name, Params: node.Params, Body: node.Body}
		env.Set(name, obj)
		return obj // 形式上必要

	// マクロ呼出し
	case *parser.MacroCallStatement:
		return e.evalMacroCallStatement(node, env)

	// 代入文
	case *parser.AssignStatement:
		target := e.evalExpression(node.Left, env, node.Context)

		if isError(target) {
			return object.ERROR
		}
		sym, ok := target.(*object.SymbolObject)
		if !ok {
			e.logger.Error(errcode.EASSIGN_INVALID_TAGET, node.Context)
			return object.ERROR
		}
		if sym.Name != "_" && sym.SymType != object.SYM_VAR {
			e.logger.Error(errcode.EASSIGN_INVALID_TAGET, node.Context)
			return object.ERROR
		}

		value := e.evalExpression(node.Value, env, node.Context)
		if isError(value) {
			return object.ERROR
		} else if isRefNotFound(value) {
			e.logger.Error(errcode.EASSIGN_INVALID_VALUE, node.Context)
			return object.ERROR
		}

		if sym.Name != "_" {
			e.logger.Error("_ 以外への代入は未実装", node.Context)
			return object.ERROR
		}
		return &object.ValueObject{Value: value, Context: node.Context}

	case *parser.IfStatement:
		return e.evalIfStatement(node, env)

	case *parser.BlockStatement:
		return e.evalBlockStatement(node, env)

	case *parser.FuncStatement:
		return e.evalFuncStatement(node, env)

	case *parser.ReturnStatement:
		return e.evalReturnStatement(node, env)

	case *parser.EnumStatement:
		name := node.Name
		_, ok := env.Get(name) // TODO enum 定義は常にグローバルスコープ
		if ok {
			e.logger.Error(fmt.Sprintf(errcode.EENUM_USED, name), node.Context)
			return object.ERROR
		}
		v := e.evalEnumStatement(node, env)
		switch v.Type() {
		case object.ENUM_OBJ:
			env.Set(v.(*object.EnumObject).Name, v)
			return v
		case object.NULL_OBJ: // TODO
			return &object.NodeObject{Node: node}
		default:
			return object.ERROR
		}

	default:
		e.logger.Error(fmt.Sprintf(errcode.ENOT_IMPL_STMT, node), nil) // TODO
		return object.ERROR
	}
}

// 複合文 BlockStatement
func (e *Evaluator) evalBlockStatement(stmt *parser.BlockStatement, env object.Environment) object.Object {
	block := &object.BlockObject{Block: []object.Object{}}

	for i, node := range stmt.Block {
		obj := e.evalStatement(node, env)
		switch obj := obj.(type) {
		case *object.EnumObject:
			for _, k := range obj.Keys {
				block.Block = append(block.Block, obj.Value[k])
			}
		case *object.ReturnObject:
			block.Block = append(block.Block, obj)
			return block
		case *object.BlockObject:
			if len(obj.Block) == 0 {
				block.Block = append(block.Block, object.NULL)
				continue
			}
			block.Block = append(block.Block, obj.Block...)
			if block.Block[len(block.Block)-1].Type() == object.RETURN_OBJ {
				return block
			}
		case *object.NodeObject:
			stmt.Block[i] = obj.Node
		default:
			block.Block = append(block.Block, obj)
		}
	}
	return block
}

// ラベル定義文
func (e *Evaluator) evalLabelStatement(node *parser.LabelStatement, env object.Environment) object.Object {
	return e.evalLabel(node.Name, env)
}

func (e *Evaluator) evalLabel(label *parser.Label, env object.Environment) object.Object {
	name := label.Name

	obj, ok := env.Get(name)
	if !ok {
		// 環境にないなら新規登録
		sym := object.NewLabelSymbol(name, getLocationCounter(env), label.Context)
		env.Set(name, sym)
		return sym
	}

	sym, ok := obj.(*object.SymbolObject)
	if !ok || sym.SymType != object.SYM_LABEL && sym.SymType != object.SYM_UNKNOWN {
		// Symbol 以外か、SYM_LABEL でない場合
		e.logger.Error(fmt.Sprintf(errcode.ELABEL_USED, name), label.Context)
		return object.ERROR
	}
	if sym.SymType == object.SYM_LABEL && !sym.Context.Equal(label.Context) {
		// ラベル 二重定義
		e.logger.Error(fmt.Sprintf(errcode.ELABEL_DUP, name), label.Context)
		return object.ERROR
	}
	if sym.SymType == object.SYM_UNKNOWN {
		sym = object.NewLabelSymbol(name, 0, label.Context)
		env.Set(name, sym)
	}
	// SYM_UNKNOWN の場合も上書き
	// 同じラベルなら値を更新
	sym.Value.(*object.NumberObject).Value = getLocationCounter(env)
	return sym
}

// const / equ 文
func (e *Evaluator) evalConstStatement(node *parser.ConstStatement, env object.Environment) object.Object {
	// e.concatenateSymbol(&node.Name, env, node.Context)
	// e.concatenateSymbol(&node.Value, env, node.Context)

	id, ok := node.Name.(*parser.Ident)
	if !ok {
		return object.ERROR
	}
	name := id.Name

	// 定義済みならエラー
	obj, ok := env.Get(name)
	if ok {
		switch obj := obj.(type) {
		case *object.SymbolObject:
			if obj.SymType == object.SYM_UNKNOWN {
				// 不明シンボンルなら更新
			} else if obj.Name != id.Name || obj.Context != node.Context {
				// 別シンボルなら二重定義エラー
				e.logger.Error(fmt.Sprintf(errcode.ESYM_DUP, name), node.Context)
				return object.ERROR
			}
			// 同一シンボルなら更新
		case *object.RefNotFoundObject:
			// 未定で登録済なら更新
		default:
			e.logger.Error(fmt.Sprintf(errcode.ESYM_USED, name), node.Context)
			return object.ERROR
		}
	}

	v := e.evalExpression(node.Value, env, node.Context)

	switch v := v.(type) {
	case *object.ErrorObject:
		return object.ERROR

	case *object.RefNotFoundObject:
		names := removeSelfName(v.Names, name)
		sym := object.NewConstSymbol(name, node.Value, object.NULL, names, node.Context)
		// sym := object.NewConstSymbol(name, node.Value, object.NULL, v.Names, node.Context)
		env.Set(name, sym)
		return &object.ValueObject{Value: object.NULL, Context: node.Context}

	case *object.NumberObject, *object.StringObject, *object.RegisterObject, *object.FlagObject:
		// リテラルを値とする Symbol を作成し環境へ登録
		sym := object.NewConstSymbol(name, node.Value, v, []string{}, node.Context)
		env.Set(name, sym)
		return &object.ValueObject{Value: v, Context: node.Context}
	default:
		if e.Debug > 0 {
			fmt.Printf("const %s = %#v\n", name, v)
		}
		env.Set(name, v)
		return v
	}
}

func removeSelfName(names []string, name string) []string {
	nmap := map[string]bool{}
	for _, n := range names {
		nmap[n] = true
	}
	nmap[name] = false

	result := []string{}
	for k, v := range nmap {
		if v {
			result = append(result, k)
		}
	}
	return result
}

// if 文
func (e *Evaluator) evalIfStatement(stmt *parser.IfStatement, env object.Environment) object.Object {
	cond, ok := e.evalExpression(stmt.Condition, env, stmt.Context).(*object.NumberObject)
	if !ok {
		return &object.NodeObject{Node: stmt}
	}
	if cond.Value != 0 {
		if stmt.Consequence == nil {
			return object.NULL
		}
		return e.evalStatement(stmt.Consequence, env)
	} else if stmt.Alternative == nil {
		return object.NULL
	} else {
		return e.evalStatement(stmt.Alternative, env)
	}
}

type evalBlockStatementFunc func(block parser.Node, env object.Environment) object.Object

// block 評価関数指定の If 文評価
func (e *Evaluator) evalIfStatementWithFunc(
	stmt *parser.IfStatement,
	env object.Environment,
	fn evalBlockStatementFunc) object.Object {

	cond := e.evalExpression(stmt.Condition, env, stmt.Context)
	if isError(cond) || isRefNotFound(cond) {
		return cond
	}

	if isTruthy(cond) {
		return fn(stmt.Consequence, env)
	} else if stmt.Alternative != nil {
		return fn(stmt.Alternative, env)
	} else {
		return object.NULL
	}
}

// function 文
func (e *Evaluator) evalFuncStatement(stmt *parser.FuncStatement, env object.Environment) object.Object {
	name := stmt.Name
	if name[0] == '@' || name[0] == '.' {
		e.logger.Error(fmt.Sprintf(errcode.EFUNC_NAME, name), stmt.Context)
		return object.ERROR
	}
	if obj, ok := env.Get(name); ok {
		if obj.Type() == object.FUNC_OBJ {
			e.logger.Error(fmt.Sprintf(errcode.EFUNC_DUP, name), stmt.Context)
		} else {
			e.logger.Error(fmt.Sprintf(errcode.EFUNC_USED, name), stmt.Context)

		}
		return object.NULL
	}
	obj := &object.FunctionObject{Name: name, Params: stmt.Params, Body: stmt.Block, Env: env}
	env.Set(name, obj)
	return obj
}

// return 文
func (e *Evaluator) evalReturnStatement(stmt *parser.ReturnStatement, env object.Environment) object.Object {
	var ret object.Object
	if stmt.Value == nil {
		ret = object.NULL
	} else {
		ret = e.evalExpression(stmt.Value, env, stmt.Context)
	}
	return &object.ReturnObject{Value: ret, LineNumber: stmt.Context.Line}
}

// enum 文
func (e *Evaluator) evalEnumStatement(node *parser.EnumStatement, env object.Environment) object.Object {
	keys := []string{}
	enum := map[string]object.Object{}
	value := 0
	for _, ele := range node.Elements.Elements {
		eleName := ele.Name
		if _, ok := enum[eleName]; ok {
			e.logger.Error(fmt.Sprintf(errcode.EENUM_USED_ELE, node.Name, ele.Name), node.Context)
			return object.ERROR
		}
		keys = append(keys, eleName)
		if ele.Value == nil {
			enum[eleName] = &object.NumberObject{Value: value}
			value += 1
			continue
		}
		v := e.evalExpression(ele.Value, env, node.Context)
		switch v.Type() {
		case object.NULL_OBJ:
			enum[eleName] = &object.NodeObject{Node: ele.Value}
		case object.NUMBER_OBJ:
			enum[eleName] = v
			value = v.(*object.NumberObject).Value + 1
		case object.STRING_OBJ:
			enum[eleName] = v
		default:
			// e.logger.Error(fmt.Sprintf(errcode.E014, v), ele.LineNumber())
			return object.ERROR
		}
	}
	return &object.EnumObject{Name: node.Name, Value: enum, Keys: keys}
}
