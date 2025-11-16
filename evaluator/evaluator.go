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
	Debug      int
}

func New(logger *logger.Logger) *Evaluator {
	return &Evaluator{logger: logger}
}

func (e *Evaluator) ResolveConst(prog *parser.Program, env *object.Environment) {
	e.scanConst(prog, env)
	e.updateEnv(env)
}

func (e *Evaluator) scanConst(prog *parser.Program, env *object.Environment) {
	for _, stmt := range prog.Statements {
		cs, ok := stmt.(*parser.ConstStatement)
		if !ok {
			continue
		}
		o := e.Eval(cs.Value, env)
		switch o.Type() {
		case object.NUMBER_OBJ, object.STRING_OBJ:
			env.GlobalSet(cs.Name.Name, o)
		case object.NULL_OBJ:
			env.GlobalSet(cs.Name.Name, &object.NodeObject{Value: cs.Value})
		default:
			env.GlobalSet(cs.Name.Name, object.NULL)
		}
	}
}

func (e *Evaluator) updateEnv(env *object.Environment) {
	genv := env.GlobalEnv()
	for k, v := range genv.Store {
		if v.Type() == object.NODE_OBJ {
			o := e.Eval(v.(*object.NodeObject).Value, env)
			if o.Type() == object.NUMBER_OBJ || o.Type() == object.STRING_OBJ {
				genv.Set(k, o)
			} else {
				genv.Set(k, object.NULL)
			}
		}
	}
}

// Eval
func (e *Evaluator) Eval(node parser.Node, env *object.Environment) object.Object {
	if e.Debug > 0 {
		fmt.Printf("eval %#v)\n", node)
	}
	switch node := node.(type) {
	// Program
	case *parser.Program:
		e.initLocationCounter(env, 0)
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
		e.printLocationCounter(env)
		e.lineNumber = node.LineNumber()
		obj := e.evalZ80Instruction(node, env)
		if obj.Type() == object.CODE_OBJ {
			code := obj.(*object.CodeObject)
			code.Addr = e.getLocationCounter(env)
			e.advanceLocationCounter(env, code.Size())
		}
		e.printLocationCounter(env)
		return obj
	case *parser.LabelStatement:
		obj, ok := env.Get("$")
		if !ok {
			panic(fmt.Sprintf("could not get $ at %s", node.String()))
		}
		addr := obj.(*object.NumberObject)
		return &object.NumberObject{Value: addr.Value, LineNumber: node.LineNumber()}
	case *parser.ConstStatement:
		v := e.Eval(node.Value, env)
		env.Set(strings.ToUpper(node.Name.Name), v)
		switch v := v.(type) {
		case *object.NumberObject, *object.StringObject:
			return v
		default:
			return &object.NodeObject{Value: node, LineNumber: node.LineNumber()}

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
		case object.NULL_OBJ:
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
		v1 := e.Eval(node.Op1, env)
		v2 := e.Eval(node.Op2, env)
		if v1.Type() == object.ERROR_OBJ || v2.Type() == object.ERROR_OBJ {
			return object.ERROR
		}
		return e.evalInfixExpression(node.Operator, v1, v2, node.LineNumber())
	case *parser.PrefixExpression:
		v := e.Eval(node.Op, env)
		if v.Type() == object.ERROR_OBJ {
			return object.ERROR
		}
		return e.evalPrefixExpression(node.Operator, v, node.LineNumber())
	case *parser.Ident:
		obj, ok := env.Get(strings.ToUpper(node.Name))
		if !ok {
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

// location counter 初期化
func (e *Evaluator) initLocationCounter(env *object.Environment, addr int) {
	env.GlobalSet("$", &object.NumberObject{Value: addr})
}

// location counter 取得
func (e *Evaluator) getLocationCounter(env *object.Environment) int {
	counter, ok := env.GlobalGet("$")
	if !ok {
		panic("getLocationCounter failed")
	}
	return counter.(*object.NumberObject).Value

}

// location counter 表示
func (e *Evaluator) printLocationCounter(env *object.Environment) {
	fmt.Printf("$ %04x\n", e.getLocationCounter(env))
}

// location counter 更新
func (e *Evaluator) advanceLocationCounter(env *object.Environment, n int) {
	obj, ok := env.GlobalGet("$")
	if !ok {
		panic("getLocationCounter failed")
	}
	counter := obj.(*object.NumberObject)
	counter.Value += n
	fmt.Println(counter.String())
}

func (e *Evaluator) evalProgram(prog *parser.Program, env *object.Environment) object.Object {
	results := &object.ProgramObject{}

	for _, stmt := range prog.Statements {
		if e.Debug > 1 {
			fmt.Println("eval stmt", stmt.String())
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
		default:
			results.Objects = append(results.Objects, obj)
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
	if obj == object.ERROR {
		return object.ERROR
	} else if obj == object.NULL {
		return &object.NodeObject{Value: expr, LineNumber: expr.LineNumber()}
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
		if v == object.ERROR {
			return object.ERROR
		} else if v == object.NULL {
			v = &object.NodeObject{Value: expr.Arguments.Expressions[i]}
		}
		newEnv.Set(p, v)
	}

	ret, ok := e.evalBlockStatement(fn.Body.(*parser.BlockStatement), newEnv).(*object.BlockObject)
	if !ok {
		panic(fmt.Sprintf("call func %s returns %T(%#v)", fn.Name, ret, ret))
	}
	last := ret.Block[len(ret.Block)-1]
	if len(ret.Block) == 0 || last.Type() != object.RETURN_OBJ {
		return object.NULL
	}
	return last.(*object.ReturnObject).Value
}

// 中置演算子式
func (e *Evaluator) evalInfixExpression(opCode int, op1, op2 object.Object, lineNumber int) object.Object {
	switch {
	case op1.Type() == object.NUMBER_OBJ && op2.Type() == object.NUMBER_OBJ:
		return e.evalNumberInfixExpression(opCode, op1, op2, lineNumber)
	case opCode == '+' && op1.Type() == object.STRING_OBJ && op2.Type() == object.STRING_OBJ:
		s1 := op1.(*object.StringObject).Value
		s2 := op2.(*object.StringObject).Value
		return &object.StringObject{Value: s1 + " " + s2}
	}
	return object.NULL
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
		return &object.NumberObject{Value: v1 + v2}
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
