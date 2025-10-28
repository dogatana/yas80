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
%token<token> ADD MULDIV COMP SHIFT UNARY
%token SL SR EQ NEQ GE LE OR AND
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
			| program instruction EOL
			{
				if $2 != nil {
					Root.Statements = append(Root.Statements, $2)
				}
			}
			| program expr EOL
			{
				if $2 != nil {
					Root.Statements = append(Root.Statements, &ExpressionStatement{Value: $2})
				}
			}
			| program error EOL
			{
				yylex.Error(__yyfmt__.Sprintf("[program error] %#v", $2))
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
			| expr ADD expr
			{
				$$ = buildInfixExpression($2.SubType, $1, $3, yylex.Error)
			}
			| expr '-' expr
			{
				$$ = buildInfixExpression('-', $1, $3, yylex.Error)
			}
			| expr MULDIV expr
			{ 	
				$$ = buildInfixExpression($2.SubType, $1, $3, yylex.Error)
			}
			| expr COMP expr
			{ 	
				$$ = buildInfixExpression($2.SubType, $1, $3, yylex.Error)
			}
			| expr SHIFT expr
			{ 	
				$$ = buildInfixExpression($2.SubType, $1, $3, yylex.Error)
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
				$$ = buildPrefixExpression($1.SubType, $2, yylex.Error)
			}
			| error 
			{ 
				yylex.Error(__yyfmt__.Sprintf("[expr error] %s", $1.String()))
				$$ = nil
			}
			;
%%

func Parse(l *Lexer) (int, int) {
	lexerInstance = l
	Root = Program{}
	// error トークンでリカバリすると yyParse() は 0 を返す
	yyParse(l)
	return l.Errors.Count()
}