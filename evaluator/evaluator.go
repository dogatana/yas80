package evaluator

import (
	"fmt"
	"strings"
	"yas80/errcode"
	"yas80/logger"
	"yas80/object"
	"yas80/parser"
)

type Evaluator struct {
	logger     *logger.Logger
	lineNumber int
	Pass1      bool
	Debug      int
}

func New(logger *logger.Logger) *Evaluator {
	return &Evaluator{logger: logger, Pass1: true}
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

	// Statement
	case *parser.Z80Instruction:
		e.lineNumber = node.TokenContext.Line
		obj := e.evalZ80Instruction(node, env)
		if obj.Type() == object.CODE_OBJ {
			code := obj.(*object.CodeObject)
			code.Addr = getLocationCounter(env)
			advanceLocationCounter(env, code.Size())
		}
		return obj

	case *parser.LabelStatement:
		return e.evalLabelStatement(node, env)
	case *parser.ConstStatement:
		// const/equ は参照内容によって NumberObject/StringObject/SymbolObject/SymbolExprObject のいずれかになる
		return e.evalConstStatement(node, env)

	case *parser.IfStatement:
		return e.evalIfStatement(node, env)
	case *parser.BlockStatement:
		return e.evalBlockStatement(node, env)

	case *parser.MacroStatement:
		return e.evalMacroStatement(node, env)
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
			e.logger.Error(fmt.Sprintf(errcode.E012, name), node.TokenContext.Line)
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
		e.lineNumber = node.TokenContext.Line
		return e.Eval(node.Value, env)

	// Expression
	case *parser.CallExpression:
		return e.evalCallExpression(node, env)
	case *parser.NumberLiteral:
		return &object.NumberObject{Value: node.Value, LineNumber: node.TokenContext.Line}
	case *parser.StringLiteral:
		return &object.StringObject{Value: node.Value, LineNumber: node.TokenContext.Line}
	case *parser.FlagLiteral:
		return &object.StringObject{Value: node.String(), LineNumber: node.TokenContext.Line}
	case *parser.InfixExpression:
		return e.evalInfixExpression(node, env, node.TokenContext.Line)
	case *parser.PrefixExpression:
		v := e.Eval(node.Op, env)
		if isError(v) || isRefNotFound(v) {
			return v
		}
		return e.evalPrefixExpression(node.Operator, v, node.TokenContext.Line)
	case *parser.Ident:
		uname := node.Name
		obj, ok := env.Get(uname)
		if !ok && e.Pass1 {
			// Pass1 の場合は未定義識別子を UndefinedObject として返す
			return &object.RefNotFoundObject{Names: []string{uname}}
		} else if !ok {
			// Pass2 の場合は ERROR を返す
			e.logger.Error(fmt.Sprintf(errcode.E009, uname), node.TokenContext.Line)
			return object.ERROR
		}
		sym, ok := (obj).(*object.SymbolObject)
		if ok && sym.SymState == object.NOT_REGISTERED {
			e.logger.Error(fmt.Sprintf(errcode.E009, uname), node.TokenContext.Line)
			return object.ERROR
		}
		return obj
	case *parser.DotIdent:
		enum, ok := env.Get(node.Left)
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.E010, node.Left), node.TokenContext.Line)
			return object.ERROR
		}
		v, ok := enum.(*object.EnumObject).Get(node.Right)
		if !ok {
			e.logger.Error(fmt.Sprintf(errcode.E011, node.Left, node.Right), node.TokenContext.Line)
			return object.ERROR
		}
		return v
	case *parser.RegisterLiteral:
		return object.Z80RgisterObjects[int(node.NodeSubType())]
	default:
		e.logger.Error(fmt.Sprintf(errcode.E999, node), 0) // TODO
		return object.ERROR
	}
}

// Program 評価
func (e *Evaluator) evalProgram(prog *parser.Program, env object.Environment) object.Object {
	results := &object.ProgramObject{}

	for i, stmt := range prog.Statements {
		if e.Debug > 1 {
			fmt.Println("eval stmt", stmt.String())
		}
		if stmt.NodeType() == parser.NODE_DELETED_STMT {
			continue
		}
		obj := e.Eval(stmt, env)
		switch obj := obj.(type) {
		case *object.EnumObject:
			for _, k := range obj.Keys {
				results.Objects = append(results.Objects, obj.Value[k])
			}
		case *object.ReturnObject:
			results.Objects = append(results.Objects, obj.Value)
			return results
		case *object.BlockObject:
			if len(obj.Block) == 0 {
				results.Objects = append(results.Objects, object.NULL)
				continue
			}
			ret, ok := obj.Block[len(obj.Block)-1].(*object.ReturnObject)
			if ok {
				obj.Block[len(obj.Block)-1] = ret.Value
			}
			results.Objects = append(results.Objects, obj.Block...)
		// case *object.SymbolObject:
		// 	prog.Statements[i] = &parser.DeletedStatement{Node: stmt}
		// 	results.Objects = append(results.Objects, obj)
		default:
			results.Objects = append(results.Objects, obj)
		}

		// pass1 で無効化するステートメント
		switch stmt := stmt.(type) {
		case *parser.ConstStatement, *parser.FuncStatement, *parser.MacroStatement:
			prog.Statements[i] = &parser.DeletedStatement{Node: stmt}
		}
	}
	return results
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
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_DEF, node.Name), node.TokenContext.Line)
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
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_FUNC_NOT_FOUND, node.Name), node.TokenContext.Line)
		return object.ERROR
	} else if !ok {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_NOT_FOUND, node.Name), node.TokenContext.Line)
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
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_NOT_MACRO, node.Name), node.TokenContext.Line)
		return object.ERROR
	}
}

// マクロ Body 評価
func (e *Evaluator) evalMacroBody(node *parser.MacroCallStatement, macro *object.MacroObject, env object.Environment) object.Object {
	// 仮引数、引数の数のチェック
	if len(node.Args.Expressions) != len(macro.Params) {
		e.logger.Error(fmt.Sprintf(errcode.EMACRO_ARG_COUNT, macro.Name), node.TokenContext.Line)
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
	savePass := e.Pass1
	e.Pass1 = false
	ret, ok := e.evalBlockStatement(macro.Body, newEnv).(*object.BlockObject)
	if !ok {
		panic(fmt.Sprintf("call macro %s returns %T(%#v)", macro.Name, ret, ret))
	}

	e.Pass1 = savePass
	return ret
}

// ラベル定義文
func (e *Evaluator) evalLabelStatement(node *parser.LabelStatement, env object.Environment) object.Object {
	name := node.Value.Name
	// pass2 で定義済みならエラー
	if obj, ok := env.Get(name); ok {
		switch obj := obj.(type) {
		case *object.SymbolObject:
			if obj.SymType != object.LABEL || obj.LineNumber != node.TokenContext.Line || obj.SymState == object.VALUE_DETERMINED {
				if !e.Pass1 {
					e.logger.Error(fmt.Sprintf(errcode.E032, name), node.TokenContext.Line)
				}
				return object.ERROR
			}
			// fall through
		default:
			if !e.Pass1 {
				e.logger.Error(fmt.Sprintf(errcode.E032, name), node.TokenContext.Line)
			}
			return object.ERROR
		}
	}
	// TODO: 変数の場合、条件アセンブルによってに時的確定かどうかを判別する必要あり
	// sym := &object.SymbolObject{
	// 	Name: uname, Node: node.Value, Value: addr, SymState: object.VALUE_DETERMINED, DependsOn: []string{}}
	sym := object.NewLabelSymbol(name, getLocationCounter(env), node.TokenContext.Line)
	env.Set(name, sym)
	return sym
}

// const / equ 文
func (e *Evaluator) evalConstStatement(node *parser.ConstStatement, env object.Environment) object.Object {
	name := node.Name.Name

	// 定義済みならエラー
	if _, ok := env.Get(name); ok {
		e.logger.Error(fmt.Sprintf(errcode.E031, name), node.TokenContext.Line)
		return object.ERROR
	}
	v := e.Eval(node.Value, env)

	switch v := v.(type) {
	case *object.NumberObject, *object.StringObject:
		// 定数として確定
		env.Set(name, v)
		return v
	case *object.RefNotFoundObject:
		// 階層でチェックが入っているはずだが念のため
		if !e.Pass1 {
			e.logger.Error(fmt.Sprintf(errcode.E009, strings.Join(v.Names, ", ")), node.TokenContext.Line)
			return object.ERROR
		}
		// 未定義定数として登録
		sym := object.NewNullConstSymbol(name, node.Value, v.Names, node.TokenContext.Line)
		env.Set(name, sym)
		return sym
	case *object.SymbolObject:
		// Symbo Object の場合は値を取得し新たに登録する
		depends := make([]string, len(v.DependsOn)) // 他のシンボルの情報なので copy
		copy(depends, v.DependsOn)
		sym := object.NewNullConstSymbol(name, node.Value, depends, node.TokenContext.Line)
		env.Set(name, sym)
		return sym

	case *object.SymbolExprObject:
		// Symbo Expression Object の場合は値を取得し新たに登録する
		sym := object.NewNullConstSymbol(name, node.Value, v.Names, node.TokenContext.Line)
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
		return &object.NodeObject{Value: stmt, LineNumber: stmt.TokenContext.Line}
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
		e.logger.Error(fmt.Sprintf(errcode.E018, name), stmt.TokenContext.Line)
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
	return &object.ReturnObject{Value: ret, LineNumber: stmt.TokenContext.Line}
}

// enum 文
func (e *Evaluator) evalEnumStatement(node *parser.EnumStatement, env object.Environment) object.Object {
	keys := []string{}
	enum := map[string]object.Object{}
	value := 0
	for _, ele := range node.Elements.Elements {
		eleName := ele.Name
		if _, ok := enum[eleName]; ok {
			e.logger.Error(fmt.Sprintf(errcode.E013, node.Name, ele.Name), node.TokenContext.Line)
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
