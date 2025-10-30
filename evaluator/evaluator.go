package evaluator

import (
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

func (e *Evaluator) Eval(node parser.Node) object.Object {
	switch node := node.(type) {
	case *parser.Program:
		return e.evalStatements(node.Statements)
	// Statement
	case *parser.ExpressionStatement:
		e.lineNumber = node.LineNumber
		return e.Eval(node.Value)
	case *parser.Z80Instruction:
		e.lineNumber = node.LineNumber
		return &object.StringObject{Value: node.String()}
	// Expression
	case *parser.NumberLiteral:
		return &object.NumberObject{Value: node.Value}
	case *parser.RegisterLiteral, *parser.FlagLiteral, *parser.IndirectExpression:
		return &object.StringObject{Value: node.String()}
	case *parser.InfixExpression:
		v1 := e.Eval(node.Op1)
		v2 := e.Eval(node.Op1)
		return e.evalInfixExpression(node.OpCode, v1, v2)
	}
	return nil
}

func (e *Evaluator) evalStatements(stmts []parser.Node) object.Object {
	var result object.Object
	for _, stmt := range stmts {
		result = e.Eval(stmt)
	}
	return result
}

func (e *Evaluator) evalInfixExpression(opCode int, op1, op2 object.Object) object.Object {
	switch {
	case op1.Type() == object.NUMBER_OBJ && op2.Type() == object.NUMBER_OBJ:
		return e.evalNumberInfixExpression(opCode, op1, op2)
	case opCode == '+' && op1.Type() == object.STRING_OBJ && op2.Type() == object.STRING_OBJ:
		s1 := op1.(*object.StringObject).Value
		s2 := op2.(*object.StringObject).Value
		return &object.StringObject{Value: s1 + " " + s2}
	}
	return object.NULL
}

func (e *Evaluator) evalNumberInfixExpression(opCode int, op1, op2 object.Object) object.Object {
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

func (e *Evaluator) toBool(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}
