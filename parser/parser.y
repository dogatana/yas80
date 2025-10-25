%{
package parser

import (
	"fmt"
)

var Result int
// goyacc が __yyfmt__ を勝手に import することの対策
var _ = __yyfmt__.Sprintf
%}
%union {
	token Token
	num int
	err any
}
// プログラムの構成要素を指定
%type<num> expr program
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
program		: { $$ = 0}
			| program EOL
			| program expr EOL
			{
				Result = $2
				__yyfmt__.Println("Result", $2)
			}
			| program error EOL
			{
				yylex.Error(
					__yyfmt__.Sprintf("%q の後に誤りがあります", $2.Literal))
				yyerrok()
			}
			;

expr		: NUMBER
	 		{
				n, err := parseInt($1.Literal)
				if err == nil {
					$$ = int(n)
				} else {
					yylex.Error(fmt.Sprintf("invalid NUMBER literal '%s'", $1.Literal))
					$$ = 0
				}
			}
			| '(' expr ')' 	{ $$ = $2 }
			| expr '+' expr { $$ = $1 + $3 }
			| expr '-' expr { $$ = $1 - $3 }
			| expr '*' expr { $$ = $1 * $3 }
			| expr '/' expr { $$ = $1 / $3 }
			| '-' expr %prec UNARY_MINUS     { $$ = - $2 }
			| error { fmt.Println("error[expr]", $1)}
			;
%%

func Parse(l yyLexer) int {
	yyDebug = 3
	return yyParse(l)
}