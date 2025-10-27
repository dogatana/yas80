%{
package parser

import (
	"fmt"
)

var Root Program
var lexerInstance *Lexer

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
%type<node> expr 
%type<node> instruction

%token<token> NUMBER IDENT
%token<token> Z80_INST0 Z80_INST1 Z80_INST2 Z80_REG8 Z80_REG16 Z80_FLAG
%token<token> ADDSUB MULDIV COMP SHIFT UNARY
%token SL SR EQ NEQ GE LE OR AND
%token  '(' ')' ',' '<' '>' '~' '!' '^' '|' '+' '-' '*' '/' '&' ':'
%token INVALID EOL 
%token<token> error

// 演算の優先度の指定
%left OR
%left AND
%left COMP
%left ADDSUB '|' '^' '-'
%left MULDIV SHIFT
%right UNARY 

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
				if $2 != nil {
					Root.Statements = append(Root.Statements, $2)
				}
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
				$$ = &Z80Instruction{OpCode: $1.SubType, lineNumber: lexerInstance.lineNumber} 
			}
			| Z80_INST1 '(' expr ')'
			{
				$$ = &Z80Instruction{OpCode: $1.SubType, lineNumber: lexerInstance.lineNumber, Op1: &IndirectExpression{Expression: $3}}
			}
			| Z80_INST1 expr
			{
				$$ = &Z80Instruction{OpCode: $1.SubType, lineNumber: lexerInstance.lineNumber, Op1: $2}
			}
			| Z80_INST2 '(' expr ')'
			{
				$$ = &Z80Instruction{
					OpCode: $1.SubType, lineNumber: lexerInstance.lineNumber, Op1: nil, Op2: &IndirectExpression{Expression: $3}}
			}
			| Z80_INST2 expr
			{
				$$ = &Z80Instruction{OpCode: $1.SubType, lineNumber: lexerInstance.lineNumber, Op1: nil, Op2: $2}
			}
			| Z80_INST2 '(' expr ')' ',' expr
			{
				$$ = &Z80Instruction{
					OpCode: $1.SubType, lineNumber: lexerInstance.lineNumber, Op1: &IndirectExpression{Expression: $3}, Op2: $6}
			}
			| Z80_INST2 expr ',' '(' expr ')'
			{
				$$ = &Z80Instruction{
					OpCode: $1.SubType, lineNumber: lexerInstance.lineNumber, Op1: $2, Op2: &IndirectExpression{Expression: $5}}
			}
			| Z80_INST2 '(' expr ')' ',' '(' expr ')'
			{
				yylex.Error("両方のオペランドを間接指定にすることはできません")
				$$ = nil
			}
			| Z80_INST2 expr ',' expr
			{
				$$ = &Z80Instruction{OpCode: $1.SubType, lineNumber: lexerInstance.lineNumber, Op1: $2, Op2: $4}
			}
			;

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
			| Z80_REG8 			{ $$ = &RegisterLiteral{TokenType: $1.SubType}}
			| Z80_REG16 		{ $$ = &FlagLiteral{TokenType: $1.SubType}}
			| Z80_FLAG 			{ $$ = &FlagLiteral{TokenType: $1.SubType}}
			| '(' expr ')'	{ $$ = $2}
			| expr ADDSUB expr
			{
				$$ = &InfixExpression{OpCode: $2.SubType, Op1: $1, Op2: $3}
			}
			| expr '-' expr
			{
				$$ = &InfixExpression{OpCode: '-', Op1: $1, Op2: $3}
			}
			| expr MULDIV expr
			{ 	
				$$ = &InfixExpression{OpCode: $2.SubType, Op1: $1, Op2: $3}
			}
			| expr COMP expr
			{ 	
				$$ = &InfixExpression{OpCode: $2.SubType, Op1: $1, Op2: $3}
			}
			| expr SHIFT expr
			{ 	
				$$ = &InfixExpression{OpCode: $2.SubType, Op1: $1, Op2: $3}
			}
			| expr OR expr
			{ 	
				$$ = &InfixExpression{OpCode: OR, Op1: $1, Op2: $3}
			}
			| expr AND expr
			{ 	
				$$ = &InfixExpression{OpCode: AND, Op1: $1, Op2: $3}
			}
			| '-' expr %prec UNARY
			{
				$$ = &PrefixExpression{OpCode: '-', Op: $2}
			}
			| UNARY expr
			{
				$$ = &PrefixExpression{OpCode: $1.SubType, Op: $2}
			}
			| error 
			{ 
				fmt.Println("error[expr]", $1)
				$$ = nil
			}
			;
%%

func Parse(l *Lexer) int {
	lexerInstance = l
	Root = Program{}
	return yyParse(l)
}