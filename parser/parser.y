%{
package parser

import (
	"fmt"
)

var Root Program
var Result int
// goyacc が __yyfmt__ を勝手に import することの対策
var _ = __yyfmt__.Sprintf
%}
%union {
	token Token
	num Node
	node Node
	err any
}


// プログラムの構成要素を指定
%type<num> program
%type<node> expr reg_expr
%type<node> instruction

%token<token> NUMBER IDENT
%token<token> Z80_INST0 Z80_INST1 Z80_INST2 Z80_REG8 Z80_REG16 Z80_FLAG
%token '+' '-' '*' '/' '(' ')'
%token INVALID EOL
%token<token> error

// 演算の優先度の指定
%left '+','-'
%left '*','/'
%right UNARY_MINUS

%%
// 文法規則を指定
program		: { }
			| program EOL
//			| program expr EOL
//			{
//				Result = $2
//				__yyfmt__.Println("Result", $2)
//			}
			| program instruction EOL
			{
				Root.Statements = append(Root.Statements, $2)
			}
			| program error EOL
			{
				yylex.Error(
					__yyfmt__.Sprintf("%q の後に誤りがあります", $2.Literal))
				yyerrok()
			}
			;

instruction	: Z80_INST0
			{
				$$ = &Z80Instruction{OpCode: $1.Op, Line: $1.Line} 
			}
			| Z80_INST1 '(' reg_expr ')'
			{
				$$ = &Z80Instruction{OpCode: $1.Op, Line: $1.Line, Op1: &RegisterIndirectExpression{Expression: $3}}
				fmt.Println($$)
			}
			| Z80_INST1 expr
			{
				fmt.Printf("$2 %T(%#v)\n", $2, $2)
				$$ = &Z80Instruction{OpCode: $1.Op, Line: $1.Line, Op1: $2}
			}
			;

reg_expr	: Z80_REG16 { $$ = &RegisterLiteral{TokenType: $1.Op}}
			| reg_expr '+' expr
			{
				$$ = &InfixExpression{OpCode: '+', Op1: $1, Op2: $3}
			}

expr		: NUMBER
	 		{
				n, err := parseInt($1.Literal)
				if err == nil {
					$$ = &NumberLiteral{TokenType: NUMBER, Value: int(n)}
				} else {
					yylex.Error(fmt.Sprintf("invalid NUMBER literal '%s'", $1.Literal))
					$$ = nil
				}
			}
			| Z80_REG8 		{ $$ = &RegisterLiteral{TokenType: $1.Op}}
			| Z80_REG16 	{ $$ = &FlagLiteral{TokenType: $1.Op}}
			| Z80_FLAG 		{ $$ = &FlagLiteral{TokenType: $1.Op}}
			| '(' expr ')' 	{ $$ = $2 }
//			| expr '+' expr { $$ = $1 + $3 }
//			| expr '-' expr { $$ = $1 - $3 }
//			| expr '*' expr { $$ = $1 * $3 }
//			| expr '/' expr { $$ = $1 / $3 }
//			| '-' expr %prec UNARY_MINUS     { $$ = - $2 }
			| error { fmt.Println("error[expr]", $1)}
			;
%%

func Parse(l yyLexer) int {
	Root = Program{}
	return yyParse(l)
}