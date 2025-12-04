package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/logger"
	"yas80/object"
	"yas80/parser"
)

type Evaluator struct {
	logger     *logger.Logger
	Debug      int
	Resolved   bool
	Counter    func() int
	lineNumber int
}

func New(logger *logger.Logger) *Evaluator {
	return &Evaluator{logger: logger, Resolved: true, Counter: makeCounter(0)}
}

// start + 1 から順次生成するカウンタ
func makeCounter(start int) func() int {
	return func() int {
		start++
		return start
	}
}

// Eval
func (e *Evaluator) Eval(node parser.Node, env object.Environment) object.Object {
	if e.Debug > 0 {
		fmt.Printf("eval %#v)\n", node)
	}
	switch node := node.(type) {

	// Program
	case *parser.Program:
		// 一旦 0 に初期化し ORG 他で上書きする
		initLocationCounter(env, 0)
		return e.evalProgram(node, env)

	case *parser.Z80Instruction:
		return e.evalZ80Instruction(node, env)
	case *parser.LabelStatement:
		return e.evalLabelStatement(node, env)
	case *parser.ConstStatement:
		return e.evalConstStatement(node, env)

	case *parser.IfStatement:
		return e.evalIfStatement(node, env)
	case *parser.MacroStatement:
		return e.evalMacroStatement(node, env)

	case *parser.BlockStatement:
		return e.evalBlockStatement(node, env)
	case *parser.MacroCallStatement:
		return e.evalMacroCallStatement(node, env)

	case *parser.FuncStatement:
		return e.evalFuncStatement(node, env)
	case *parser.ReturnStatement:
		return e.evalReturnStatement(node, env)

	case *parser.EnumStatement:
		name := node.Name
		_, ok := env.Get(name) // TODO enum 定義は常にグローバルスコープ
		if ok {
			e.logger.Error(fmt.Sprintf(errcode.E012, name), node.Context)
			return object.ERROR
		}
		v := e.evalEnumStatement(node, env)
		switch v.Type() {
		case object.ENUM_OBJ:
			env.Set(v.(*object.EnumObject).Name, v)
			return v
		case object.NULL_OBJ: // TODO
			return &object.NodeObject{Value: node}
		default:
			return object.ERROR
		}
	case *parser.ExpressionStatement:
		e.lineNumber = node.Context.Line
		return e.Eval(node.Value, env)

	// Expression
	// 各種リテラル
	case *parser.NumberLiteral:
		return &object.NumberObject{Value: node.Value, Context: node.Context}
	case *parser.StringLiteral:
		return &object.StringObject{Value: node.Value, Context: node.Context}
	case *parser.RegisterLiteral:
		return object.Z80RegisterFlagObjects[int(node.NodeSubType())]
	case *parser.FlagLiteral:
		return object.Z80RegisterFlagObjects[int(node.NodeSubType())]

	// 識別子
	case *parser.Ident:
		name := node.Name
		obj, ok := env.Get(name)
		if !ok {
			// 未定義の場合
			obj = &object.RefNotFoundObject{Names: []string{name}}
			env.Set(name, obj)
			e.Resolved = false
		}
		return obj
	case *parser.DotIdent: // TODO enum か proc.local かの識別必要
		enum, ok := env.Get(node.Left)
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.E010, node.Left), node.Context)
			return object.ERROR
		}
		v, ok := enum.(*object.EnumObject).Get(node.Right)
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.E011, node.Left, node.Right), node.Context)
			return object.ERROR
		}
		return v

	// 関数呼出し
	case *parser.CallExpression:
		return e.evalCallExpression(node, env)

	// 式
	case *parser.InfixExpression:
		return e.evalInfixExpression(node, env, node.Context)
	case *parser.PrefixExpression:
		return e.evalPrefixExpression(node, env, node.Context)

	default:
		e.logger.Error(fmt.Sprintf(errcode.E999, node), nil) // TODO
		return object.ERROR
	}
}

// Program 評価
func (e *Evaluator) evalProgram(prog *parser.Program, env object.Environment) object.Object {
	objects := []object.Object{}
	stmts := []parser.Node{}

	var obj object.Object

	for i := 0; i < len(prog.Statements); i++ {
		if e.Debug > 0 {
			fmt.Printf("eval prog.Statements[%d]\n", i)
			addr, _ := env.Get("$")
			fmt.Printf("$ %s\n", addr.String())
		}

		node := prog.Statements[i]

		switch stmt := node.(type) {
		// 命令
		case *parser.Z80Instruction:
			obj = e.Eval(stmt, env)
			if obj.Type() == object.CODE_OBJ {
				code := obj.(*object.CodeObject)
				code.Addr = getLocationCounter(env)
				advanceLocationCounter(env, code.Size())
			}
			objects = append(objects, obj)
			stmts = append(stmts, node)

		// ラベル
		case *parser.LabelStatement:
			if stmt.Name.LabelType != parser.NODE_LABEL {
				// LOCAL/AT の場合は AST から LabelStatement を削除する
				e.logger.Error(fmt.Sprintf(errcode.EGLOBAL_NOT_ALLOWED, stmt.Name.Name), stmt.Context)
				continue
			}
			obj = e.Eval(stmt, env)
			// ValueObject にラップして返す
			objects = append(objects, &object.ValueObject{Value: obj, Context: stmt.Context})
			stmts = append(stmts, node)

		// const/equ
		case *parser.ConstStatement:
			if stmt.Name.IdentType != parser.IDENT {
				e.logger.Error(fmt.Sprintf(errcode.EGLOBAL_NOT_ALLOWED, stmt.Name.Name), stmt.Context)
				continue
			}
			obj := e.Eval(stmt, env)
			objects = append(objects, &object.ValueObject{Value: obj, Context: stmt.Context})
			stmts = append(stmts, node)

		case *parser.MacroStatement:
			if _, ok := env.Get(stmt.Name); ok {
				e.logger.Error(fmt.Sprintf(errcode.EMACRO_DEF, stmt.Name), stmt.Context)
				continue
			}
			obj := &object.MacroObject{Name: stmt.Name, Params: stmt.Params, Body: stmt.Body}
			env.Set(stmt.Name, obj)
			continue

		case *parser.MacroCallStatement:
			obj, ok := env.Get(stmt.Name)
			if !ok {
				e.logger.Error(fmt.Sprintf(errcode.EMACRO_NOT_FOUND, stmt.Name), stmt.Context)
				continue
			}
			macro, ok := obj.(*object.MacroObject)
			if !ok {
				e.logger.Error(fmt.Sprintf(errcode.EMACRO_NOT_MACRO, stmt.Name), stmt.Context)
				continue
			}
			expanded := e.expandMacro(macro)
			stmts = append(stmts, expanded...)
			e.Resolved = false

		default:
			e.logger.Info(fmt.Sprintf(errcode.E999, node), nil)
			obj = e.Eval(node, env)
			if obj == object.ERROR {
				continue
			}
			objects = append(objects, obj)
			stmts = append(stmts, node)
		}
	}

	prog.Statements = stmts
	return &object.ProgramObject{Objects: objects}
}

// 複合文 BlockStatement
func (e *Evaluator) evalBlockStatement(stmt *parser.BlockStatement, env object.Environment) object.Object {
	block := &object.BlockObject{Block: []object.Object{}}

	for _, stmt := range stmt.Block {
		obj := e.Eval(stmt, env)
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
		default:
			block.Block = append(block.Block, obj)
		}
	}
	return block
}

// マクロ定義文
func (e *Evaluator) evalMacroStatement(node *parser.MacroStatement, env object.Environment) object.Object {
	if _, ok := env.Get(node.Name); ok {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_DEF, node.Name), node.Context)
		return object.ERROR
	}
	obj := &object.MacroObject{Name: node.Name, Params: node.Params, Body: node.Body}
	env.Set(node.Name, obj)
	return obj
}

// マクロ定義文
func (e *Evaluator) evalMacroCallStatement(node *parser.MacroCallStatement, env object.Environment) object.Object {
	obj, ok := env.Get(node.Name)
	if !ok && len(node.Args.Expressions) == 1 {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_FUNC_NOT_FOUND, node.Name), node.Context)
		return object.ERROR
	} else if !ok {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_NOT_FOUND, node.Name), node.Context)
		return object.ERROR
	}
	// 関数オブジェクトなら関数呼出し評価へ回す
	// 文法定義上 1引数の関数呼出し式文は MacroCallStatement となるので、ここで置き換える

	switch obj := obj.(type) {
	case *object.FunctionObject:
		funcall := &parser.CallExpression{
			Function:  &parser.Ident{Name: node.Name},
			Arguments: node.Args}
		return e.evalCallExpression(funcall, env) // TODO linenumber はパッケージ外から設定できない
	case *object.MacroObject:
		return e.evalMacroBody(node, obj, env)
	default:
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_NOT_MACRO, node.Name), node.Context)
		return object.ERROR
	}
}

// マクロ Body 評価
func (e *Evaluator) evalMacroBody(node *parser.MacroCallStatement, macro *object.MacroObject, env object.Environment) object.Object {
	// 仮引数、引数の数のチェック
	if len(node.Args.Expressions) != len(macro.Params) {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_ARG_COUNT, macro.Name), node.Context)
		return object.ERROR
	}

	// 引数を評価し、仮引数名で環境に設定
	newEnv := object.NewAtLocalEnvironment(env)
	for i, param := range macro.Params {
		v := e.Eval(node.Args.Expressions[i], env)
		if isError(v) || isRefNotFound(v) {
			return v
		}
		newEnv.Set("@@"+param, v)
	}

	object.PrintEnv(newEnv)

	// TODO 引数評価の前に設定しなくても良いか？
	// マクロ展開の評価は Pass1 であっても未定義エラーを発生させる
	ret, ok := e.evalBlockStatement(macro.Body, newEnv).(*object.BlockObject)
	if !ok {
		panic(fmt.Sprintf("call macro %s returns %T(%#v)", macro.Name, ret, ret))
	}
	return ret
}

// ラベル定義文
func (e *Evaluator) evalLabelStatement(node *parser.LabelStatement, env object.Environment) object.Object {
	return e.evalLabel(node.Name, env)
}

func (e *Evaluator) evalLabel(label *parser.Label, env object.Environment) object.Object {
	name := label.Name

	obj, ok := env.Get(name)
	if !ok || obj.Type() == object.REF_NOTFOUND_OBJ {
		// 環境にないか、RefNotFoundObject なら新規登録
		sym := object.NewLabelSymbol(name, getLocationCounter(env), label.Context)
		env.Set(name, sym)
		return sym
	}
	sym, ok := obj.(*object.SymbolObject)
	if !ok || sym.SymType != object.SYM_LABEL {
		// Symbol で || LABEL でなけれがエラー
		e.logger.Error(fmt.Sprintf(errcode.ELABEL_USED_NAME, name), label.Context)
		return object.ERROR
	}
	if !sym.Context.Equal(label.Context) {
		// ラベル 二重定義
		e.logger.Error(fmt.Sprintf(errcode.ELABEL_DUP, name), label.Context)
		return object.ERROR
	}
	// 同じラベルなら値を更新
	sym.Value.(*object.NumberObject).Value = getLocationCounter(env)
	return sym
}

// const / equ 文
func (e *Evaluator) evalConstStatement(node *parser.ConstStatement, env object.Environment) object.Object {
	name := node.Name.Name

	// 定義済みならエラー
	obj, ok := env.Get(name)
	if ok {
		switch obj := obj.(type) {
		case *object.NumberObject, *object.StringObject, *object.RegisterObject:
			// 定数として確定済
			return &object.ValueObject{Value: obj, Context: node.Context}
		case *object.SymbolObject:
			if obj.Name != node.Name.Name || obj.Context != node.Context {
				// 別シンボルなら二重定義エラー
				e.logger.Error(fmt.Sprintf(errcode.ESYM_DUP, name), node.Context)
				return object.ERROR
			}
			// 同一シンボルなら更新
		case *object.RefNotFoundObject:
			// 未定で登録済なら更新
		default:
			e.logger.Error(fmt.Sprintf(errcode.ESYM_USED_NAME, name), node.Context)
			return object.ERROR
		}
	}

	v := e.Eval(node.Value, env)

	switch v := v.(type) {
	case *object.NumberObject, *object.StringObject:
		// 定数として確定
		env.Set(name, v)
		return &object.ValueObject{Value: v, Context: node.Context}
	case *object.RefNotFoundObject:
		sym := object.NewNullConstSymbol(name, node.Value, v.Names, node.Context)
		env.Set(name, sym)
		return &object.ValueObject{Value: object.NULL, Context: node.Context}
	case *object.SymbolObject:
		// 他のシンボルの場合は値をコピーして新規に登録
		depends := make([]string, len(v.DependsOn)+1) // 他のシンボルの情報なので copy
		copy(depends, v.DependsOn)
		depends = append(depends, v.Name) // 参照シンボルの名前も追加
		sym := object.NewConstSymbol(name, v.Value, depends, node.Context)
		env.Set(name, sym)
		return sym
	case *object.SymbolExprObject:
		// Symbo Expression Object の場合は値を取得し新たに登録する
		sym := object.NewNullConstSymbol(name, node.Value, v.Names, node.Context)
		env.Set(name, sym)
		return sym
	case *object.ErrorObject:
		return object.ERROR
	default:
		if e.Debug > 0 {
			fmt.Printf("const %s = %#v\n", name, v)
		}
		env.Set(name, v)
		return v
	}
}

// if 文
func (e *Evaluator) evalIfStatement(stmt *parser.IfStatement, env object.Environment) object.Object {
	cond, ok := e.Eval(stmt.Condition, env).(*object.NumberObject)
	if !ok {
		return &object.NodeObject{Value: stmt, LineNumber: stmt.Context.Line}
	}
	if cond.Value != 0 {
		if stmt.Consequence == nil {
			return object.NULL
		}
		return e.Eval(stmt.Consequence, env)
	} else if stmt.Alternative == nil {
		return object.NULL
	} else {
		return e.Eval(stmt.Alternative, env)
	}
}

// function 文
func (e *Evaluator) evalFuncStatement(stmt *parser.FuncStatement, env object.Environment) object.Object {
	name := stmt.Name
	_, ok := env.Get(name)
	if ok {
		e.logger.Error(fmt.Sprintf(errcode.E018, name), stmt.Context)
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
		ret = e.Eval(stmt.Value, env)
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
			e.logger.Error(fmt.Sprintf(errcode.E013, node.Name, ele.Name), node.Context)
			return object.ERROR
		}
		keys = append(keys, eleName)
		if ele.Value == nil {
			enum[eleName] = &object.NumberObject{Value: value}
			value += 1
			continue
		}
		v := e.Eval(ele.Value, env)
		switch v.Type() {
		case object.NULL_OBJ:
			enum[eleName] = &object.NodeObject{Value: ele.Value}
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
