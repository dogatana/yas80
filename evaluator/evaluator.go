package evaluator

import (
	"fmt"
	"strings"
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
func (e *Evaluator) Eval(node parser.Node, env *object.Environment) object.Object {
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
	case *parser.ExpressionStatement:
		e.lineNumber = node.LineNumber()
		return e.Eval(node.Value, env)
	case *parser.ReturnStatement:
		return e.evalReturnStatement(node, env)
	case *parser.IfStatement:
		return e.evalIfStatement(node, env)
	case *parser.FunctionStatement:
		return e.evalFunctionStatement(node, env)
	case *parser.Z80Instruction:
		printLocationCounter(env)
		e.lineNumber = node.LineNumber()
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
	case *parser.BlockStatement:
		return e.evalBlockStatement(node, env)
	case *parser.EnumStatement:
		name := strings.ToUpper(node.Name)
		_, ok := env.GlobalGet(name) // enum 定義は常にグローバルスコープ
		if ok {
			e.logger.Error(fmt.Sprintf(logger.E012, node.Name), node.LineNumber())
			return object.ERROR
		}
		v := e.evalEnumStatement(node, env)
		switch v.Type() {
		case object.ENUM_OBJ:
			env.GlobalSet(v.(*object.EnumObject).Name, v)
			return v
		case object.NULL_OBJ: // TODO
			return &object.NodeObject{Value: node}
		default:
			return object.ERROR
		}

	// Expression
	case *parser.CallExpression:
		return e.evalCallExpression(node, env)
	case *parser.NumberLiteral:
		return &object.NumberObject{Value: node.Value, LineNumber: node.LineNumber()}
	case *parser.StringLiteral:
		return &object.StringObject{Value: node.Value, LineNumber: node.LineNumber()}
	case *parser.FlagLiteral:
		return &object.StringObject{Value: node.String(), LineNumber: node.LineNumber()}
	case *parser.InfixExpression:
		return e.evalInfixExpression(node, env, node.LineNumber())
	case *parser.PrefixExpression:
		v := e.Eval(node.Op, env)
		if isError(v) || isRefNotFound(v) {
			return v
		}
		return e.evalPrefixExpression(node.Operator, v, node.LineNumber())
	case *parser.Ident:
		uname := strings.ToUpper(node.Name)
		obj, ok := env.Get(uname)
		if !ok && e.Pass1 {
			// Pass1 の場合は未定義識別子を UndefinedObject として返す
			return &object.RefNotFoundObject{Names: []string{uname}}
		} else if !ok {
			// Pass2 の場合は ERROR を返す
			e.logger.Error(fmt.Sprintf(logger.E009, uname), node.LineNumber())
			return object.ERROR
		}
		sym, ok := (obj).(*object.SymbolObject)
		if ok && sym.SymState == object.NOT_REGISTERED {
			e.logger.Error(fmt.Sprintf(logger.E009, uname), node.LineNumber())
			return object.ERROR
		}
		return obj
	case *parser.DotIdent:
		enum, ok := env.Get(node.Left)
		if !ok {
			e.logger.Error(fmt.Sprintf(logger.E010, node.Left), node.LineNumber())
			return object.ERROR
		}
		v, ok := enum.(*object.EnumObject).Get(node.Right)
		if !ok {
			e.logger.Error(fmt.Sprintf(logger.E011, node.Left, node.Right), node.LineNumber())
			return object.ERROR
		}
		return v
	case *parser.RegisterLiteral:
		return object.Z80RgisterObjects[int(node.NodeSubType())]
	default:
		fmt.Printf("default %T\n", node)
		e.logger.Error(fmt.Sprintf(logger.E999, node), node.LineNumber())
		return object.ERROR
	}
}

// Program 評価
func (e *Evaluator) evalProgram(prog *parser.Program, env *object.Environment) object.Object {
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
		case *parser.ConstStatement, *parser.FunctionStatement:
			prog.Statements[i] = &parser.DeletedStatement{Node: stmt}
		}
	}
	return results
}

// 複合文 BlockStatement
func (e *Evaluator) evalBlockStatement(stmt *parser.BlockStatement, env *object.Environment) object.Object {
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

func (e *Evaluator) evalLabelStatement(node *parser.LabelStatement, env *object.Environment) object.Object {
	uname := strings.ToUpper(node.Value.Name)
	// pass2 で定義済みならエラー
	if obj, ok := env.Get(uname); ok {
		switch obj := obj.(type) {
		case *object.SymbolObject:
			if obj.SymType != object.LABEL || obj.LineNumber != node.LineNumber() || obj.SymState == object.VALUE_DETERMINED {
				if !e.Pass1 {
					e.logger.Error(fmt.Sprintf(logger.E032, uname), node.LineNumber())
				}
				return object.ERROR
			}
			// fall through
		default:
			if !e.Pass1 {
				e.logger.Error(fmt.Sprintf(logger.E032, uname), node.LineNumber())
			}
			return object.ERROR
		}
	}
	// TODO: 変数の場合、条件アセンブルによってに時的確定かどうかを判別する必要あり
	// sym := &object.SymbolObject{
	// 	Name: uname, Node: node.Value, Value: addr, SymState: object.VALUE_DETERMINED, DependsOn: []string{}}
	sym := object.NewLabelSymbol(uname, getLocationCounter(env), node.LineNumber())
	env.Set(uname, sym)
	return sym
}

// const / equ 文
func (e *Evaluator) evalConstStatement(node *parser.ConstStatement, env *object.Environment) object.Object {
	uname := strings.ToUpper(node.Name.Name)

	// 定義済みならエラー
	if _, ok := env.Get(uname); ok {
		e.logger.Error(fmt.Sprintf(logger.E031, uname), node.LineNumber())
		return object.ERROR
	}
	v := e.Eval(node.Value, env)

	switch v := v.(type) {
	case *object.NumberObject, *object.StringObject:
		// 定数として確定
		env.Set(uname, v)
		return v
	case *object.RefNotFoundObject:
		// 階層でチェックが入っているはずだが念のため
		if !e.Pass1 {
			e.logger.Error(fmt.Sprintf(logger.E009, strings.Join(v.Names, ", ")), node.LineNumber())
			return object.ERROR
		}
		// 未定義定数として登録
		sym := object.NewNullConstSymbol(uname, node.Value, v.Names, node.LineNumber())
		env.Set(uname, sym)
		return sym
	case *object.SymbolObject:
		// Symbo Object の場合は値を取得し新たに登録する
		depends := make([]string, len(v.DependsOn)) // 他のシンボルの情報なので copy
		copy(depends, v.DependsOn)
		sym := object.NewNullConstSymbol(uname, node.Value, depends, node.LineNumber())
		env.Set(uname, sym)
		return sym

	case *object.SymbolExprObject:
		// Symbo Expression Object の場合は値を取得し新たに登録する
		sym := object.NewNullConstSymbol(uname, node.Value, v.Names, node.LineNumber())
		env.Set(uname, sym)
		return sym
	case *object.ErrorObject:
		return object.ERROR
	default:
		if e.Debug > 0 {
			fmt.Printf("const %s = %#v\n", uname, v)
		}
		env.Set(uname, v)
		return v
	}
}

// if 文
func (e *Evaluator) evalIfStatement(stmt *parser.IfStatement, env *object.Environment) object.Object {
	cond, ok := e.Eval(stmt.Condition, env).(*object.NumberObject)
	if !ok {
		return &object.NodeObject{Value: stmt, LineNumber: stmt.LineNumber()}
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
func (e *Evaluator) evalFunctionStatement(stmt *parser.FunctionStatement, env *object.Environment) object.Object {
	name := strings.ToUpper((stmt.Name))
	_, ok := env.Get(name)
	if ok {
		e.logger.Error(fmt.Sprintf(logger.E018, stmt.Name), stmt.LineNumber())
		return object.NULL
	}
	obj := &object.FunctionObject{Name: name, Params: stmt.Params, Body: stmt.Block, Env: env}
	env.Set(strings.ToUpper(stmt.Name), obj)
	return obj
}

// return 文
func (e *Evaluator) evalReturnStatement(stmt *parser.ReturnStatement, env *object.Environment) object.Object {
	var ret object.Object
	if stmt.Value == nil {
		ret = object.NULL
	} else {
		ret = e.Eval(stmt.Value, env)
	}
	return &object.ReturnObject{Value: ret, LineNumber: stmt.LineNumber()}
}

// enum 文
func (e *Evaluator) evalEnumStatement(node *parser.EnumStatement, env *object.Environment) object.Object {
	keys := []string{}
	enum := map[string]object.Object{}
	value := 0
	for _, ele := range node.Elements.Elements {
		eleName := strings.ToUpper(ele.Name)
		if _, ok := enum[eleName]; ok {
			e.logger.Error(fmt.Sprintf(logger.E013, node.Name, ele.Name), node.LineNumber())
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
			// e.logger.Error(fmt.Sprintf(logger.E014, v), ele.LineNumber())
			return object.ERROR
		}
	}
	return &object.EnumObject{Name: strings.ToUpper(node.Name), Value: enum, Keys: keys}
}
