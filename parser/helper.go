package parser

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInt(s string) (int64, error) {
	str := strings.ReplaceAll(s, "_", "")
	length := len(str)
	switch {
	case length >= 3 && str[0] == '0':
		return strconv.ParseInt(str, 0, 0)
	case length >= 2 && str[0] == '$':
		return strconv.ParseInt("0x"+str[1:length], 0, 0)
	case length >= 2 && (str[length-1] == 'h' || str[length-1] == 'H'):
		return strconv.ParseInt(str[0:length-1], 16, 0)
	case length >= 2 && str[0] == '%':
		return strconv.ParseInt(str[1:length], 2, 0)
	default:
		return strconv.ParseInt(str, 0, 0)
	}
}

func opString(e Node) string {
	indirect, ok := e.(*IndirectExpression)
	if ok {
		return indirect.String()
	}
	s := e.String()
	return trimParen(s)
}
func trimParen(s string) string {
	if s[0] == '(' {
		s = s[1 : len(s)-1]
	}
	return s
}

func SetYYDebug(v int) {
	yyDebug = v
}

var twoCharTokenNames map[int]string = map[int]string{
	LE:  "<=",
	GE:  ">=",
	EQ:  "==",
	NEQ: "!=",
	SL:  "<<",
	SR:  ">>",
	AND: "&&",
	OR:  "||",
}

func tokenLiteral(t int) string {
	if t == '\n' {
		return "EOL"
	}
	name, ok := twoCharTokenNames[t]
	if ok {
		return name
	}
	name = yySymNames[yyXLAT[t]]
	// 1 文字トークンは 'x' のように ' で囲まれているのでそれをはずす
	if name[0] == '\'' {
		return name[1 : len(name)-1]
	}
	return name
}

type errorFunc func(s string, args ...any)

// 数値リテラルの畳み込み(中置演算子)
type infixFuncType func(x, y int, errFn errorFunc) int

var infixFuncs map[int]infixFuncType = map[int]infixFuncType{
	'+': func(x, y int, fn errorFunc) int { return x + y },
	'-': func(x, y int, fn errorFunc) int { return x - y },
	'*': func(x, y int, fn errorFunc) int { return x * y },
	'/': func(x, y int, fn errorFunc) int {
		if y != 0 {
			return x / y
		} else {
			fn("division by 0")
			return 0
		}
	},
	'&': func(x, y int, fn errorFunc) int { return x & y },
	'|': func(x, y int, fn errorFunc) int { return x | y },
	'^': func(x, y int, fn errorFunc) int { return x ^ y },
	SL:  func(x, y int, fn errorFunc) int { return x << y },
	SR:  func(x, y int, fn errorFunc) int { return x >> y },
	'<': func(x, y int, fn errorFunc) int {
		if x < y {
			return 1
		} else {
			return 0
		}
	},
	'>': func(x, y int, fn errorFunc) int {
		if x > y {
			return 1
		} else {
			return 0
		}
	},
	LE: func(x, y int, fn errorFunc) int {
		if x <= y {
			return 1
		} else {
			return 0
		}
	},
	GE: func(x, y int, fn errorFunc) int {
		if x >= y {
			return 1
		} else {
			return 0
		}
	},
	EQ: func(x, y int, fn errorFunc) int {
		if x == y {
			return 1
		} else {
			return 0
		}
	},
	NEQ: func(x, y int, fn errorFunc) int {
		if x != y {
			return 1
		} else {
			return 0
		}
	},
	OR: func(x, y int, fn errorFunc) int {
		if x != 0 || y != 0 {
			return 1
		} else {
			return 0
		}
	},
	AND: func(x, y int, fn errorFunc) int {
		if x != 0 && y != 0 {
			return 1
		} else {
			return 0
		}
	},
}

func buildInfixExpression(opcode int, op1, op2 Node, errFn errorFunc) Expression {
	num1, ok1 := op1.(*NumberLiteral)
	num2, ok2 := op2.(*NumberLiteral)
	if ok1 && ok2 {
		fn, ok := infixFuncs[opcode]
		var v int
		if ok {
			v = fn(num1.Value, num2.Value, errFn)
		} else {
			errFn(fmt.Sprintf("UNKNOWN infix %s", yySymNames[yyXLAT[opcode]]))
			v = 0
		}
		return &NumberLiteral{TokenType: NUMBER, Value: v}
	}
	return &InfixExpression{OpCode: opcode, Op1: op1, Op2: op2}
}

// 数値リテラルの畳み込み(前置演算子)
type prefixFuncType func(x int, fn errorFunc) int

var prefixFuncs map[int]prefixFuncType = map[int]prefixFuncType{
	'-': func(x int, fn errorFunc) int { return -x },
	'~': func(x int, fn errorFunc) int { return -1 ^ x },
	'!': func(x int, fn errorFunc) int {
		if x != 0 {
			return 0
		} else {
			return 1
		}
	},
}

func buildPrefixExpression(opcode int, op Node, errFn errorFunc) Expression {
	num, ok := op.(*NumberLiteral)
	if ok {
		fn, ok := prefixFuncs[opcode]
		var v int
		if ok {
			v = fn(num.Value, errFn)
		} else {
			errFn(fmt.Sprintf("UNKNOWN prefix %s", yySymNames[yyXLAT[opcode]]))
			v = 0
		}
		return &NumberLiteral{TokenType: NUMBER, Value: v}
	}
	return &PrefixExpression{OpCode: opcode, Op: op}
}
