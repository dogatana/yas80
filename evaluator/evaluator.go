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
	debug      int
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
	if e.debug > 0 {
		fmt.Printf("eval %T(%#v)\n", node, node)
	}
	switch node := node.(type) {
	// Program
	case *parser.Program:
		return e.evalProgram(node, env)

	// Statement
	case *parser.ExpressionStatement:
		e.lineNumber = node.LineNumber()
		return e.Eval(node.Value, env)
	case *parser.ReturnStatement:
		return e.evalReturnStatement(node, env)
	case *parser.IfStatement:
		return e.evalIfStatement(node, env)
	case *parser.Z80Instruction:
		e.lineNumber = node.LineNumber()
		return e.evalZ80Instruction(node, env)
	case *parser.ConstStatement:
		v := e.Eval(node.Value, env)
		env.Set(node.Name.Name, v)
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
	case *parser.NumberLiteral:
		return &object.NumberObject{Value: node.Value, LineNumber: node.LineNumber()}
	case *parser.StringLiteral:
		return &object.StringObject{Value: node.Value, LineNumber: node.LineNumber()}
	case *parser.FlagLiteral:
		return &object.StringObject{Value: node.String(), LineNumber: node.LineNumber()}
	case *parser.InfixExpression:
		v1 := e.Eval(node.Op1, env)
		v2 := e.Eval(node.Op2, env)
		return e.evalInfixExpression(node.OpCode, v1, v2, node.LineNumber())
	case *parser.Ident:
		obj, ok := env.Get(node.Name)
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

func (e *Evaluator) evalProgram(prog *parser.Program, env *object.Environment) object.Object {
	results := &object.ProgramObject{}

	for _, stmt := range prog.Statements {
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
		return &object.NumberObject{Value: e.toBool(v1 == v2)}
	case parser.NEQ:
		return &object.NumberObject{Value: e.toBool(v1 != v2)}
	case '<':
		return &object.NumberObject{Value: e.toBool(v1 < v2)}
	case parser.LE:
		return &object.NumberObject{Value: e.toBool(v1 <= v2)}
	case '>':
		return &object.NumberObject{Value: e.toBool(v1 > v2)}
	case parser.GE:
		return &object.NumberObject{Value: e.toBool(v1 >= v2)}
	case parser.OR:
		return &object.NumberObject{Value: e.toBool(v1 != 0 || v2 != 0)}
	case parser.AND:
		return &object.NumberObject{Value: e.toBool(v1 != 1 && v2 != 1)}
	default:
		e.logger.Error(fmt.Sprintf(logger.E016, string(rune(opCode))), lineNumber)
		return object.ERROR
	}
}

func (e *Evaluator) evalZ80Instruction(node *parser.Z80Instruction, env *object.Environment) object.Object {
	switch node.NodeType() {
	case parser.Z80_INST0:
		info := Z80CodeTable0[int(node.OpCode)]
		obj := &object.Code{Line: node.LineNumber(), Code: make([]byte, len(info.Bytes))}
		copy(obj.Code, info.Bytes)
		return obj
	case parser.Z80_INST1:
		if node.OpCode == parser.Z80_INST_RET {
			return e.generateRET(node, env)
		}
		return object.NULL
	case parser.Z80_INST2:
		return e.evalZ80Instruction2(node, env)
	default:
		return object.NULL
	}
}

func (e *Evaluator) generateRET(node *parser.Z80Instruction, env *object.Environment) object.Object {
	if node.Op1 == nil {
		return &object.Code{Line: node.LineNumber(), Code: []byte{0xc9}}
	}
	if node.Op1.NodeType() == parser.Z80_FLAG {
		flag := int(node.Op1.NodeSubType()) - parser.Z80_FLAG_NZ
		b := byte(0xc0 | flag<<3)
		return &object.Code{Line: node.LineNumber(), Code: []byte{b}}
	}
	e.logger.Error(
		fmt.Sprintf(logger.E017, node.Op1.String()), node.LineNumber())
	return object.ERROR
}

func (e *Evaluator) evalZ80Instruction2(node *parser.Z80Instruction, env *object.Environment) object.Object {
	switch node.NodeSubType() {
	case parser.Z80_INST_LD:
		return e.evalZ80LD(node, env)
	default:
		e.logger.Error(fmt.Sprintf(logger.E999, node), node.LineNumber())
		return object.ERROR
	}
}

func (e *Evaluator) evalZ80LD(node *parser.Z80Instruction, env *object.Environment) object.Object {
	switch node.Op1.NodeType() {
	case parser.Z80_REG8:
		r1 := int(node.Op1.NodeSubType())
		if r1 > parser.Z80_REG_A {
			return object.NULL
		}
		switch node.Op2.NodeType() {
		case parser.Z80_REG8:
			// LD r,r'
			r2 := int(node.Op2.NodeSubType())
			if r2 > parser.Z80_REG_A {
				return object.NULL
			}
			b := 0xc0
			b |= ((r1 - parser.Z80_REG_B) << 3) | (r2 - parser.Z80_REG_B)
			return &object.Code{Line: node.LineNumber(), Code: []byte{byte(b)}}
		default:
			op2 := e.Eval(node.Op2, env)
			if op2.Type() == object.NUMBER_OBJ {
				// LD r,n
				v := op2.(*object.NumberObject).Value
				bv, ok := e.toByte(v)
				if !ok {
					e.logger.Warning(fmt.Sprintf(logger.W001, v, v), node.LineNumber())
					bv = byte(v & 0xff)
				}
				b := 0x06
				b |= (r1 - parser.Z80_REG_B) << 3
				return &object.Code{Line: node.LineNumber(), Code: []byte{byte(b), bv}}
			} else if op2.Type() == object.NULL_OBJ {
				e.logger.Error(fmt.Sprintf("error expr: %s", node.Op2.String()), 0)
			}
		}
	default:
		return object.NULL
	}
	return object.NULL
}

func (e *Evaluator) toByte(n int) (byte, bool) {
	switch {
	case 0 <= n && n <= 255:
		return byte(n), true
	case -128 <= n && n < 0:
		return byte(n & 0xff), true
	default:
		return 0, false
	}
}

func (e *Evaluator) toBool(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
