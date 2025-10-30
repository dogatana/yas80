%{
package parser

import (
	"fmt"
)

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
%type<node> expr 
%type<node> instruction statement

%token<token> NUMBER IDENT
%token<token> Z80_INST0 Z80_INST1 Z80_INST2 Z80_REG8 Z80_REG16 Z80_FLAG
%token<token> ADD MULDIV COMP SHIFT UNARY
%token SL SR EQ NEQ GE LE OR AND
%token CONST VAR EQU FUNC
%token IF ELSE ELIF END_IF
%token MACRO END_MACRO
%token REPEAT END_REPEAT
%token FUNCTION END_FUNCTION
%token PROC END_PROC
%token BLOCK END_BLOCK
%token  '(' ')' ',' '<' '>' '~' '!' '^' '|' '+' '-' '*' '/' '&' ':'
%token INVALID EOL 
%token<token> error

// 演算の優先度の指定
%left OR
%left AND
%left COMP
%left ADD '|' '^' '-'
%left MULDIV SHIFT
%right UNARY 

%%
// 文法規則を指定
program		: { }
			| program EOL
			| program statement EOL
			{
				if $2 != nil {
					prog := yylex.(*Lexer).program
					prog.Statements = append(prog.Statements, $2)
				}
			}
			| program expr EOL
			{
				if $2 != nil {
					prog := yylex.(*Lexer).program
					prog.Statements = append(prog.Statements, &ExpressionStatement{Value: $2})
				}
			}
			| program error EOL
			{
				yylex.Error(__yyfmt__.Sprintf("[program error] %#v", $2))
				yyerrok()
			}
			;

statement   : instruction			{ $$ = $1}
			| CONST IDENT '=' expr
			{ 
				$$ = &ConstStatement{Name: &Ident{Name: $2.Literal}, Value: $4}
			}
			| IDENT EQU expr		
			{ 
				$$ = &ConstStatement{Name: &Ident{Name: $1.Literal}, Value: $3}
			}
			;
instruction	: Z80_INST0
			{
				$$ = &Z80Instruction{
					InstType: Z80_INST0, OpCode: int($1.TokenSubType), LineNumber: $1.LineNumber} 
			}
			| Z80_INST1
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST1, OpCode: int($1.TokenSubType), LineNumber: $1.LineNumber}
			}
			| Z80_INST1 '(' expr ')'
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST1, OpCode: int($1.TokenSubType), LineNumber: $1.LineNumber,
						Op1: &IndirectExpression{Expression: $3}}
			}
			| Z80_INST1 expr
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST1, OpCode: int($1.TokenSubType), LineNumber: $1.LineNumber,
						Op1: $2}
			}
			| Z80_INST2 '(' expr ')'
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST2, OpCode: int($1.TokenSubType), LineNumber: $1.LineNumber,
						Op2: &IndirectExpression{Expression: $3}}
			}
			| Z80_INST2 expr
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST2, OpCode: int($1.TokenSubType), LineNumber: $1.LineNumber,
						Op2: $2}
			}
			| Z80_INST2 '(' expr ')' ',' expr
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST2, OpCode: int($1.TokenSubType), LineNumber: $1.LineNumber,
						Op1: &IndirectExpression{Expression: $3},
						Op2: $6}
			}
			| Z80_INST2 expr ',' '(' expr ')'
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST2, OpCode: int($1.TokenSubType), LineNumber: $1.LineNumber,
						Op1: $2,
						Op2: &IndirectExpression{Expression: $5}}
			}
			| Z80_INST2 '(' expr ')' ',' '(' expr ')'
			{
				yylex.Error("両方のオペランドを間接指定にすることはできません", $1.LineNumber)
				$$ = nil
			}
			| Z80_INST2 expr ',' expr
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST2, OpCode: int($1.TokenSubType), LineNumber: $1.LineNumber,
						Op1: $2,
						Op2: $4}
			}
			;

expr		: NUMBER
	 		{
				n, err := parseInt($1.Literal)
				if err == nil {
					$$ = &NumberLiteral{Value: int(n)}
				} else {
					yylex.Error(fmt.Sprintf("invalid NUMBER literal '%s'", $1.Literal))
					$$ = nil
				}
			}
			| IDENT
			{
				$$ = &Ident{Name: $1.Literal}
			}
			| Z80_REG8 			{ $$ = &RegisterLiteral{RegisterType: int($1.TokenType), Register:int($1.TokenSubType)}}
			| Z80_REG16 		{ $$ = &RegisterLiteral{RegisterType: int($1.TokenType), Register:int($1.TokenSubType)}}
			| Z80_FLAG 			{ $$ = &FlagLiteral{Flag: int($1.TokenSubType)}}
			| '(' expr ')'		{ $$ = $2}
			| expr ADD expr
			{
				$$ = buildInfixExpression(int($2.TokenSubType), $1, $3, yylex.Error)
			}
			| expr '-' expr
			{
				$$ = buildInfixExpression('-', $1, $3, yylex.Error)
			}
			| expr MULDIV expr
			{ 	
				$$ = buildInfixExpression(int($2.TokenSubType), $1, $3, yylex.Error)
			}
			| expr COMP expr
			{ 	
				$$ = buildInfixExpression(int($2.TokenSubType), $1, $3, yylex.Error)
			}
			| expr SHIFT expr
			{ 	
				$$ = buildInfixExpression(int($2.TokenSubType), $1, $3, yylex.Error)
			}
			| expr OR expr
			{ 	
				$$ = buildInfixExpression(OR, $1, $3, yylex.Error)
			}
			| expr AND expr
			{ 	
				$$ = buildInfixExpression(AND, $1, $3, yylex.Error)
			}
			| '-' expr %prec UNARY
			{
				$$ = buildPrefixExpression('-', $2, yylex.Error)
			}
			| UNARY expr
			{
				$$ = buildPrefixExpression(int($1.TokenSubType), $2, yylex.Error)
			}
			| error 
			{ 
				yylex.Error(__yyfmt__.Sprintf("[expr error] %s", $1.String()))
				$$ = nil
			}
			;
%%

func Parse(l *Lexer) (*Program, int, int) {
	// error トークンでリカバリすると yyParse() は 0 を返す
	yyParse(l)
	ec, wc := l.ErrorStore.Count()
	return l.program, ec, wc
}