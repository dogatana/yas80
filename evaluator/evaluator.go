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
		printLocationCounter(env)
		uname := strings.ToUpper(node.Value.Name)
		// Get("$") で取得したものを利用すると、更新後の値が得られてしまうので新たに作成
		addr := &object.NumberObject{Value: getLocationCounter(env), LineNumber: node.LineNumber()}

		// TODO: 変数の場合、条件アセンブルによってに時的確定かどうかを判別する必要あり
		sym := &object.SymbolObject{
			Name: uname, Node: node.Value, Value: addr, State: object.SYMBOL_STATE_DEFINED, DependsOn: []string{}}
		env.Set(uname, sym)
		return addr
	case *parser.ConstStatement:
		// const/equ は参照内容によって NumberObject/StringObject/SymbolObject のいずれかになる
		uname := strings.ToUpper(node.Name.Name)
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
			sym := &object.SymbolObject{
				Name: uname, Node: node.Value, Value: nil, State: object.SYMBOL_STATE_UNDEFINED, DependsOn: v.Names}
			env.Set(uname, sym)
			return sym
		case *object.SymbolObject:
			// Symbo Object の場合は値を取得し新たに登録する
			sym := &object.SymbolObject{
				Name: uname, Node: node.Value, Value: v.Value, State: object.SYMBOL_STATE_UNDEFINED, DependsOn: []string{v.Name}}
			env.Set(uname, sym)
			return sym

		case *object.SymbolExprObject:
			// Symbo Object の場合は値を取得し新たに登録する
			sym := &object.SymbolObject{
				Name: uname, Node: node.Value, Value: nil, State: object.SYMBOL_STATE_UNDEFINED, DependsOn: v.Names}
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
			e.logger.Error(fmt.Sprintf(logger.E009, node.Name), node.LineNumber())
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

func (e *Evaluator) evalReturnStatement(stmt *parser.ReturnStatement, env *object.Environment) object.Object {
	var ret object.Object
	if stmt.Value == nil {
		ret = object.NULL
	} else {
		ret = e.Eval(stmt.Value, env)
	}
	return &object.ReturnObject{Value: ret, LineNumber: stmt.LineNumber()}
}

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

// 関数呼出し
func (e *Evaluator) evalCallExpression(expr *parser.CallExpression, env *object.Environment) object.Object {
	obj := e.Eval(expr.Function, env)
	if isError(obj) || isRefNotFound(obj) {
		return obj
	} else if obj == object.NULL {
		panic("object is NULL") // TODO
		// return &object.NodeObject{Value: expr, LineNumber: expr.LineNumber()}
	}

	fn, ok := obj.(*object.FunctionObject)
	if !ok {
		e.logger.Error(logger.E019, expr.LineNumber())
		return object.ERROR
	}
	if len(expr.Arguments.Expressions) != len(fn.Params) {
		e.logger.Error(fmt.Sprintf(logger.E021, fn.Name), expr.LineNumber())
		return object.ERROR
	}

	newEnv := object.NewEnvironment(fn.Env)
	for i, param := range fn.Params {
		p := strings.ToUpper(param)
		v := e.Eval(expr.Arguments.Expressions[i], env)
		if isError(v) || isRefNotFound(v) {
			return v
		}
		newEnv.Set(p, v)
	}

	// 関数本体の評価は Pass1 であっても未定義エラーを発生させる
	savePass := e.Pass1
	e.Pass1 = false

	ret, ok := e.evalBlockStatement(fn.Body.(*parser.BlockStatement), newEnv).(*object.BlockObject)
	if !ok {
		panic(fmt.Sprintf("call func %s returns %T(%#v)", fn.Name, ret, ret))
	}
	e.Pass1 = savePass

	if len(ret.Block) == 0 {
		return object.NULL
	}
	for _, obj := range ret.Block {
		if isError(obj) || isRefNotFound(obj) {
			return obj
		}
	}
	last := ret.Block[len(ret.Block)-1]
	if last.Type() == object.RETURN_OBJ {
		return last.(*object.ReturnObject).Value
	}
	return object.NULL
}

// 中置演算子式
func (e *Evaluator) evalInfixExpression(node *parser.InfixExpression, env *object.Environment, lineNumber int) object.Object {
	op1 := e.Eval(node.Op1, env)
	op2 := e.Eval(node.Op2, env)

	switch {
	case isError(op1) || isError(op2):
		return object.ERROR
	case isNumber(op1) && isNumber(op2):
		return e.evalNumberInfixExpression(node.Operator, op1, op2, lineNumber)
	case isString(op1) && isString(op2):
		if node.Operator != '+' {
			if !e.Pass1 {
				e.logger.Error(logger.E029, lineNumber)
			}
			return object.ERROR
		}
		s1 := op1.(*object.StringObject).Value
		s2 := op2.(*object.StringObject).Value
		return &object.StringObject{Value: s1 + " " + s2}
	case isRefNotFound(op1) || isRefNotFound(op2):
		return &object.RefNotFoundObject{Names: mergeNames(op1, op2)}
	case isSymolOrSymbolExpr(op1) || isSymolOrSymbolExpr(op2):
		return &object.SymbolExprObject{Names: mergeNames(op1, op2)}
	default:
		if e.Debug > 0 {
			fmt.Printf("op1 %#v, op2 %#v", op1, op2)
		}
		e.logger.Error(fmt.Sprintf(logger.E023, parser.TokenLiteral(node.Operator)), lineNumber)
		return object.ERROR
	}
}
func (e *Evaluator) evalNumberInfixExpression(opCode int, op1, op2 object.Object, lineNumber int) object.Object {
	v1 := op1.(*object.NumberObject).Value
	v2 := op2.(*object.NumberObject).Value
	switch opCode {
	case '+':
		return &object.NumberObject{Value: v1 + v2}
	case '-':
		return &object.NumberObject{Value: v1 - v2}
	case '*':
		return &object.NumberObject{Value: v1 * v2}
	case '/':
		if v2 == 0 {
			e.logger.Error(logger.E015, lineNumber)
			return object.ERROR
		}
		return &object.NumberObject{Value: v1 / v2}
	case parser.SL:
		return &object.NumberObject{Value: v1 << v2}
	case parser.SR:
		return &object.NumberObject{Value: v1 >> v2}
	case '&':
		return &object.NumberObject{Value: v1 & v2}
	case '|':
		return &object.NumberObject{Value: v1 | v2}
	case '^':
		return &object.NumberObject{Value: v1 ^ v2}
	case parser.EQ:
		return &object.NumberObject{Value: e.boolToInt(v1 == v2)}
	case parser.NEQ:
		return &object.NumberObject{Value: e.boolToInt(v1 != v2)}
	case '<':
		return &object.NumberObject{Value: e.boolToInt(v1 < v2)}
	case parser.LE:
		return &object.NumberObject{Value: e.boolToInt(v1 <= v2)}
	case '>':
		return &object.NumberObject{Value: e.boolToInt(v1 > v2)}
	case parser.GE:
		return &object.NumberObject{Value: e.boolToInt(v1 >= v2)}
	case parser.OR:
		return &object.NumberObject{Value: e.boolToInt(v1 != 0 || v2 != 0)}
	case parser.AND:
		return &object.NumberObject{Value: e.boolToInt(v1 != 1 && v2 != 1)}
	default:
		e.logger.Error(fmt.Sprintf(logger.E016, string(rune(opCode))), lineNumber)
		return object.ERROR
	}
}

// 前置演算子式
func (e *Evaluator) evalPrefixExpression(opCode int, op object.Object, lineNumber int) object.Object {
	switch op := op.(type) {
	case *object.NumberObject:
		switch opCode {
		case '+':
			return &object.NumberObject{Value: op.Value, LineNumber: lineNumber}
		case '-':
			return &object.NumberObject{Value: -op.Value, LineNumber: lineNumber}
		case '~':
			return &object.NumberObject{Value: op.Value ^ -1, LineNumber: lineNumber}
		case '!':
			return &object.NumberObject{Value: e.boolToInt(op.Value == 0), LineNumber: lineNumber}
		default:
			e.logger.Error(fmt.Sprintf(logger.E008, rune(opCode)), lineNumber)
			return object.ERROR
		}
	case *object.StringObject:
		if opCode == '!' {
			return &object.NumberObject{Value: e.boolToInt(op.Value == ""), LineNumber: lineNumber}
		}
		e.logger.Error(fmt.Sprintf(logger.E007, rune(opCode)), lineNumber)
		return object.ERROR
	}
	e.logger.Error(logger.E022, lineNumber)
	return object.ERROR
}

func (e *Evaluator) boolToInt(value bool) int {
	if value {
		return 1
	} else {
		return 0
	}
}

func (e *Evaluator) EvalEnv(env *object.Environment) ([]string, error) {
	order, err := e.tSortEnv(env)
	if err != nil {
		return order, err
	}
	for _, name := range order {
		obj, ok := env.Get(name)
		if !ok {
			return order, fmt.Errorf("internal error: could not get %s", name)
		}
		sym, ok := obj.(*object.SymbolObject)
		if !ok {
			continue
		}
		if sym.State == object.SYMBOL_STATE_DEFINED {
			env.Set(name, sym.Value)
			continue
		}
		value := e.Eval(sym.Node, env)
		if isError(value) || isRefNotFound(value) {
			return order, fmt.Errorf("could not eval symbol %s", name)
		}
		env.Set(name, value)
	}
	return order, nil
}

// 環境をトポロジカルソート
func (e *Evaluator) tSortEnv(env *object.Environment) ([]string, error) {
	visited := map[string]bool{}
	visiting := map[string]bool{}
	order := []string{}

	var visit func(string) error
	visit = func(name string) error {
		if visiting[name] {
			return fmt.Errorf("循環参照: %s", name)
		}
		if visited[name] {
			return nil
		}
		visiting[name] = true
		obj, ok := env.Get(name)
		if !ok {
			return fmt.Errorf("未定義シンボル: %s", name)
		}
		sym, ok := obj.(*object.SymbolObject)
		if ok {
			for _, dep := range sym.DependsOn {
				if err := visit(dep); err != nil {
					return err
				}
			}
		}
		visited[name] = true
		visiting[name] = false
		order = append(order, name)
		return nil
	}

	for name := range env.Store {
		if err := visit(name); err != nil {
			return nil, err
		}
	}
	return order, nil
}
