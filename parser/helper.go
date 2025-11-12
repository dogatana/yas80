package parser

import (
	"fmt"
	"strconv"
	"strings"
	"yas80/logger"
)

func parseInt(s string) (int64, error) {
	// 数値リテラルの途中の _ を無視
	str := strings.ReplaceAll(s, "_", "")
	length := len(str)
	switch {
	case length >= 3 && str[0] == '0':
		// 0x 0o 0b
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
	if v > 0 {
		yyErrorVerbose = true
	}
}

var twoCharTokenNames map[int]string = map[int]string{
	LE:  "LE",  // "<=",
	GE:  "GE",  // ">=",
	EQ:  "EQ",  // "==",
	NEQ: "NEQ", // "!=",
	SL:  "SL",  // "<<",
	SR:  "SR",  // ">>",
	AND: "AND", // "&&",
	OR:  "OR",  // "||",
}

func tokenLiteral(token int) string {
	if token == '\n' {
		return "EOL"
	}
	name, ok := twoCharTokenNames[token]
	if ok {
		return name
	}
	pt := lexerTokenToParserToken((token))
	if 256 <= pt && pt < yyPrivate {
		return Z80OpCode2Name(pt)
	}
	name = yyTokname(pt)
	// 1 文字トークンは 'x' のように ' で囲まれているのでそれをはずす
	if name[0] == '\'' {
		return name[1 : len(name)-1]
	}
	return name
}

// yylex1 の冒頭処理
func lexerTokenToParserToken(char int) int {
	token := 0
	if char <= 0 {
		return int(yyTok1[0])
	}
	if char < len(yyTok1) {
		return int(yyTok1[char])
	}
	// Z80 Opcode
	if 256 <= char && char < yyPrivate {
		return char
	}
	if char >= yyPrivate && char < yyPrivate+len(yyTok2) {
		return int(yyTok2[char-yyPrivate])
	}
	for i := 0; i < len(yyTok3); i += 2 {
		token = int(yyTok3[i+0])
		if token == char {
			return int(yyTok3[i+1])
		}
	}
	return 0
}

// 数値リテラルの畳み込み(中置演算子)
type infixFuncType func(x, y int) int

var infixFuncs map[int]infixFuncType = map[int]infixFuncType{
	'+': func(x, y int) int { return x + y },
	'-': func(x, y int) int { return x - y },
	'*': func(x, y int) int { return x * y },
	'/': func(x, y int) int { return x / y },
	'&': func(x, y int) int { return x & y },
	'|': func(x, y int) int { return x | y },
	'^': func(x, y int) int { return x ^ y },
	SL:  func(x, y int) int { return x << y },
	SR:  func(x, y int) int { return x >> y },
	'<': func(x, y int) int {
		if x < y {
			return 1
		} else {
			return 0
		}
	},
	'>': func(x, y int) int {
		if x > y {
			return 1
		} else {
			return 0
		}
	},
	LE: func(x, y int) int {
		if x <= y {
			return 1
		} else {
			return 0
		}
	},
	GE: func(x, y int) int {
		if x >= y {
			return 1
		} else {
			return 0
		}
	},
	EQ: func(x, y int) int {
		if x == y {
			return 1
		} else {
			return 0
		}
	},
	NEQ: func(x, y int) int {
		if x != y {
			return 1
		} else {
			return 0
		}
	},
	OR: func(x, y int) int {
		if x != 0 || y != 0 {
			return 1
		} else {
			return 0
		}
	},
	AND: func(x, y int) int {
		if x != 0 && y != 0 {
			return 1
		} else {
			return 0
		}
	},
}

func buildInfixExpression(opcode int, op1, op2 Expression, lineNumber int) Expression {
	if op1.NodeType() == NODE_ERROR {
		return op1
	}
	if op2.NodeType() == NODE_ERROR {
		return op2
	}
	// 整数演算の畳み込み
	num1, ok1 := op1.(*NumberLiteral)
	num2, ok2 := op2.(*NumberLiteral)
	if ok1 && ok2 {
		if opcode == '/' && num2.Value == 0 {
			return &ParseError{Message: logger.E015, lineNumber: lineNumber}
		}

		fn, ok := infixFuncs[opcode]
		if ok {
			return &NumberLiteral{Value: fn(num1.Value, num2.Value), lineNumber: lineNumber}
		} else {
			return &ParseError{Message: fmt.Sprintf(logger.E016, tokenLiteral(opcode)), lineNumber: lineNumber}
		}
	}
	// 文字列演算(+)の畳み込み
	str1, ok1 := op1.(*StringLiteral)
	str2, ok2 := op2.(*StringLiteral)
	if ok1 && ok2 && opcode == '+' {
		return &StringLiteral{Value: str1.Value + str2.Value}
	}
	return &InfixExpression{OpCode: opcode, Op1: op1, Op2: op2}
}

// 数値リテラルの畳み込み(前置演算子)
type prefixFuncType func(x int) int

var prefixFuncs map[int]prefixFuncType = map[int]prefixFuncType{
	'+': func(x int) int { return x },
	'-': func(x int) int { return -x },
	'~': func(x int) int { return -1 ^ x },
	'!': func(x int) int {
		if x != 0 {
			return 0
		} else {
			return 1
		}
	},
}

func buildPrefixExpression(opcode int, op Expression, lineNumber int) Expression {
	if op.NodeType() == NODE_ERROR {
		return op
	}
	switch op := op.(type) {
	case *NumberLiteral:
		fn, ok := prefixFuncs[opcode]
		if ok {
			return &NumberLiteral{Value: fn(op.Value), lineNumber: lineNumber}
		} else {
			return &ParseError{Message: fmt.Sprintf(logger.E008, rune(opcode)), lineNumber: lineNumber}
		}
	case *StringLiteral:
		if opcode == '!' {
			var result int
			if op.Value == "" {
				result = 1
			} else {
				result = 0
			}
			return &NumberLiteral{Value: result, lineNumber: lineNumber}
		}
		return &ParseError{Message: fmt.Sprintf(logger.E007, rune(opcode)), lineNumber: lineNumber}
	}
	return &PrefixExpression{OpCode: opcode, Op: op, lineNumber: lineNumber}
}
