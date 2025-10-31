package evaluator

import (
	"fmt"
	"yas80/errorstore"
	"yas80/object"
	"yas80/parser"
)

type Evaluator struct {
	es         *errorstore.ErrorStore
	lineNumber int
}

func New(es *errorstore.ErrorStore) *Evaluator {
	return &Evaluator{es: es}
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

func (e *Evaluator) Eval(node parser.Node, env *object.Environment) object.Object {
	switch node := node.(type) {
	case *parser.Program:
		return e.evalStatements(node.Statements, env)
	// Statement
	case *parser.ExpressionStatement:
		e.lineNumber = node.LineNumber
		return e.Eval(node.Value, env)
	case *parser.Z80Instruction:
		e.lineNumber = node.LineNumber
		return e.evalZ80Instruction(node, env)
	// Expression
	case *parser.NumberLiteral:
		return &object.NumberObject{Value: node.Value}
	case *parser.RegisterLiteral, *parser.FlagLiteral, *parser.IndirectExpression:
		return &object.StringObject{Value: node.String()}
	case *parser.InfixExpression:
		v1 := e.Eval(node.Op1, env)
		v2 := e.Eval(node.Op2, env)
		return e.evalInfixExpression(node.OpCode, v1, v2, env)
	case *parser.Ident:
		obj, ok := env.Get(node.Name)
		if ok {
			return obj
		}
		return object.NULL
	case *parser.ConstStatement:
		v := e.Eval(node.Value, env)
		env.Set(node.Name.Name, v)
		return object.NULL
	}
	return nil
}

func (e *Evaluator) evalStatements(stmts []parser.Node, env *object.Environment) object.Object {
	p := &object.Program{}
	for _, stmt := range stmts {
		obj := e.Eval(stmt, env)
		// if obj != object.NULL {
		p.Objects = append(p.Objects, obj)
		// }
	}
	return p
}

func (e *Evaluator) evalInfixExpression(opCode int, op1, op2 object.Object, env *object.Environment) object.Object {
	switch {
	case op1.Type() == object.NUMBER_OBJ && op2.Type() == object.NUMBER_OBJ:
		return e.evalNumberInfixExpression(opCode, op1, op2, env)
	case opCode == '+' && op1.Type() == object.STRING_OBJ && op2.Type() == object.STRING_OBJ:
		s1 := op1.(*object.StringObject).Value
		s2 := op2.(*object.StringObject).Value
		return &object.StringObject{Value: s1 + " " + s2}
	}
	return object.NULL
}

func (e *Evaluator) evalNumberInfixExpression(opCode int, op1, op2 object.Object, env *object.Environment) object.Object {
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
			e.es.AddError("", e.lineNumber, "division by 0")
			return object.NULL
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
		return object.NULL
	}
}

func (e *Evaluator) evalZ80Instruction(node *parser.Z80Instruction, env *object.Environment) object.Object {
	switch node.NodeType() {
	case parser.Z80_INST0:
		info := Z80CodeTable0[int(node.OpCode)]
		obj := &object.FixedCode{Line: node.LineNumber, Code: make([]byte, len(info.Bytes))}
		copy(obj.Code, info.Bytes)
		return obj
	case parser.Z80_INST1:
		if node.OpCode == parser.Z80_INST_RET {
			return e.generateRET(node, env)
		}
		return object.NULL
	default:
		return object.NULL
	}
}

func (e *Evaluator) generateRET(node *parser.Z80Instruction, env *object.Environment) object.Object {
	if node.Op1 == nil {
		return &object.FixedCode{Line: node.LineNumber, Code: []byte{0xc9}}
	}
	if node.Op1.NodeType() == parser.Z80_FLAG {
		flag := int(node.Op1.NodeSubType()) - parser.Z80_FLAG_NZ
		b := byte(0xc0 | flag<<3)
		return &object.FixedCode{Line: node.LineNumber, Code: []byte{b}}
	}
	e.es.AddError("", node.LineNumber,
		fmt.Sprintf("第1オペランドがフラグではありません '%s'", node.Op1.String()))
	return object.NULL
}

func (e *Evaluator) toBool(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
