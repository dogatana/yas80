%{
package parser

import (
	"fmt"
	"strings"
)

// goyacc が __yyfmt__ を勝手に import することの対策
var _ = __yyfmt__.Sprintf
%}
%union {
	token Token
	num Node
	node Node
	err any
	enum_element *EnumElement
	enum_elements *EnumElements
	block *BlockStatement
	params []string
	expr_list *ExpressionList
	expr Expression
}


// プログラムの構成要素を指定
%type<node> statement instruction directive label elseifs
%type<block> block_statement
%type<enum_elements> enum_elements
%type<enum_element> enum_element
%type<params> param_list
%type<expr_list> expr_list
%type<expr> expr indexed_expr


%token<token> EOL
%token<token> NUMBER IDENT
%token<token> AT_IDENT    // @def 
%token<token> LOCAL_IDENT // .def 
%token<token> DOT_IDENT   // abc.def ラベル, enum
%token<token> Z80_INST0 Z80_INST1 Z80_INST2 Z80_REG8 Z80_REG16 Z80_FLAG
%token<token> ADDSUB MULDIV COMP SHIFT UNARY
%token SL SR EQ NEQ GE LE OR AND

%token CONST VAR EQU FUNC ORG

%token<token> IF 
%token ELSE ELIF END_IF
%token<token>ELIF
%token MACRO END_MACRO
%token<token> REPEAT
%token END_REPEAT
%token FUNCTION END_FUNCTION
%token PROC END_PROC
%token BLOCK END_BLOCK
%token ENUM END_ENUM
%token  '(' ')' ',' '<' '>' '~' '!' '^' '|' '+' '-' '*' '/' '&' ':' '[' ']'
%token INVALID 
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
			| program label EOL
			{
				stmt := &LabelStatement{Value: $2, lineNumber: $2.(*Label).LineNumber}
				prog := yylex.(*Lexer).program
				prog.Statements = append(prog.Statements, stmt)
			}
			| program label statement
			{
				if $3 == nil {
					// do nothing
				} else if $3.NodeType() == NODE_ERROR {
					yylex.Error($3.(*ParseError).Message, $3.(Statement).LineNumber())
				} else {
					prog := yylex.(*Lexer).program
					stmt := &LabelStatement{Value: $2, lineNumber: $2.(*Label).LineNumber}
					prog.Statements = append(prog.Statements, stmt, $3)
				}
			}
			| program statement 
			{
				if $2 == nil {
					// do nothing
				} else if $2.NodeType() == NODE_ERROR {
					yylex.Error($2.(*ParseError).Message, $2.(Statement).LineNumber())
				} else {
					prog := yylex.(*Lexer).program
					prog.Statements = append(prog.Statements, $2)
				}
			}
			| program error EOL
			{
				yylex.Error(__yyfmt__.Sprintf("[program error] %#v", $2), $3.LineNumber)
				yyerrok()
			}
			;

statement   : expr EOL			
			{ 
				if $1.NodeType() == NODE_ERROR {
					$$ = $1
				} else {
					$$ = &ExpressionStatement{Value: $1, lineNumber: $2.LineNumber}
				}
			}
			| instruction EOL	{ $$ = $1}
			| directive	 EOL	{ $$ = $1}
			| error EOL
			{
				yylex.Error(__yyfmt__.Sprintf("[statement error] %#v", $1), $2.LineNumber)
				yyerrok()

			}
			;

directive	: CONST IDENT '=' expr
			{ 
				if $4.NodeType() == NODE_ERROR {
					$$ =  $4
				} else {
					$$ = &ConstStatement{Name: &Ident{Name: $2.Literal}, Value: $4, lineNumber: $2.LineNumber}
				}
			}
			| IDENT EQU expr		
			{ 
				if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else {
					$$ = &ConstStatement{Name: &Ident{Name: $1.Literal}, Value: $3, lineNumber: $1.LineNumber}
				}
			}
			| IDENT ENUM EOL enum_elements END_ENUM
			{
				if $4.NodeType() == NODE_ERROR {
					$$ =  $4
				} else {
					$$ = &EnumStatement{Name: $1.Literal, Elements: $4, lineNumber: $1.LineNumber}
				}
			}
			| VAR IDENT '=' expr
			{
				if $4.NodeType() == NODE_ERROR {
					$$ = $4
				} else {
					$$ = &VariableStatement{Name: &Ident{Name: $2.Literal}, Value: $4, lineNumber: $2.LineNumber}
				}
			}
			| IDENT '=' expr
			{
				if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else {
					$$ = &AsignStatement{Left: &Ident{Name: $1.Literal}, Value: $3, lineNumber: $1.LineNumber}
				}
			}
			| indexed_expr '=' expr
			{
				if $1.NodeType() == NODE_ERROR {
					$$ = $1
				} else {
					$$ = &AsignStatement{Left: $1, Value: $3, lineNumber: $1.(*IndexedExpression).lineNumber}
				}
			}
			| REPEAT expr EOL block_statement END_REPEAT
			{
				__yyfmt__.Println("block_statement", $4.String())
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else if $4.NodeType() == NODE_ERROR {
					$$ = $4
				} else {
					$$ = &RepeatStatement{MaxCount: $2, Block: $4, lineNumber: $1.LineNumber}
				}
			}
			| IF expr EOL block_statement elseifs END_IF
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else if $4.NodeType() == NODE_ERROR {
					$$ = $4
				} else if $5 == nil {
					$$ = &IfStatement{Condition: $2, Consequence: $4, Alternative: $5, lineNumber: $1.LineNumber}
				} else if $5.NodeType() == NODE_ERROR {
					$$ = $5
				} else {
					$$ = &IfStatement{Condition: $2, Consequence: $4, Alternative: $5, lineNumber: $1.LineNumber}
				} 
			}
			| IF expr EOL block_statement elseifs ELSE block_statement END_IF
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else if $4.NodeType() == NODE_ERROR {
					$$ = $4
				} else if $7.NodeType() == NODE_ERROR {
					$$ = $7
				} else if $5 == nil {
					$$ = &IfStatement{Condition: $2, Consequence: $4, Alternative: $7, lineNumber: $1.LineNumber}
				} else if $5.NodeType() == NODE_ERROR {
					$$ = $5
				}  else {
					s := $5.(*IfStatement)
					for s.Alternative != nil {
						s = s.Alternative.(*IfStatement)
					}
					s.Alternative = $7

					$$ = &IfStatement{Condition: $2, Consequence: $4, Alternative: $5, lineNumber: $1.LineNumber}
				}
			}
			| IDENT FUNCTION param_list EOL block_statement END_FUNCTION
			{
				$$ = &FunctionStatement{Name: $1.Literal, Params: $3, Block: $5, lineNumber: $1.LineNumber}
			}
			;
	
param_list	: 			{ $$ = []string{}}
			| IDENT		{ $$ = []string{$1.Literal} }
			| param_list ',' IDENT
			{
				$1 = append($1, $3.Literal)
				$$ = $1
				fmt.Printf("params: %T(%#v)\n", $$, $$)
			}
			;

elseifs		: { $$ = nil }
			| elseifs ELIF expr EOL block_statement 
			{ 
				if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else if $5.NodeType() == NODE_ERROR {
					$$ = $5
				} else if $1 == nil {
					$$ = &IfStatement{Condition: $3, Consequence: $5, lineNumber: $2.LineNumber}
				} else {
					s := $1.(*IfStatement)
					for s.Alternative != nil {
						s = s.Alternative.(*IfStatement)
					}
					s.Alternative = &IfStatement{Condition: $3, Consequence: $5, lineNumber: $2.LineNumber}

					$$ = $1
				}
			}
			;
			

block_statement	: 	 				{ $$ = &BlockStatement{Block: []Node{}} }
			| block_statement EOL 	{ $$ = $1}
			| block_statement statement 
			{ 
				if $2.NodeType() == NODE_ERROR {
					$$ = $2.(*BlockStatement)
				} else {
					$1.Block = append($1.Block, $2)
					$$ = $1
				}
			}
			;
	
enum_elements : 	 			{ $$ = &EnumElements{Elements: []*EnumElement{}} }
			| enum_elements EOL { $$ = $1 }
			| enum_elements enum_element EOL
			{
				$1.Elements = append($1.Elements, $2)
				$$ = $1
			}
			;

enum_element : IDENT 			{ $$ = &EnumElement{Name: $1.Literal, Value: nil} }
			| IDENT '=' expr	{ $$ = &EnumElement{Name: $1.Literal, Value: $3} }
			;

label		: IDENT ':'
			{
				$$ = &Label{nodeType: NODE_LABEL, Name: $1.Literal, LineNumber: $1.LineNumber}
			}
			| LOCAL_IDENT ':'
			{
				$$ = &Label{nodeType: NODE_LOCAL_LABEL, Name: $1.Literal, LineNumber: $1.LineNumber}
			}
			| LOCAL_IDENT
			{
				$$ = &Label{nodeType: NODE_LOCAL_LABEL, Name: $1.Literal, LineNumber: $1.LineNumber}
			}
			;


instruction	: Z80_INST0
			{
				$$ = &Z80Instruction{
					InstType: Z80_INST0, OpCode: int($1.TokenSubType), lineNumber: $1.LineNumber} 
			}
			| Z80_INST1
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST1, OpCode: int($1.TokenSubType), lineNumber: $1.LineNumber}
			}
			| Z80_INST1 '(' expr ')'
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST1, OpCode: int($1.TokenSubType), lineNumber: $1.LineNumber,
						Op1: &IndirectExpression{Expression: $3}}
			}
			| Z80_INST1 expr
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST1, OpCode: int($1.TokenSubType), lineNumber: $1.LineNumber,
						Op1: $2}
			}
			| Z80_INST2 '(' expr ')'
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST2, OpCode: int($1.TokenSubType), lineNumber: $1.LineNumber,
						Op2: &IndirectExpression{Expression: $3}}
			}
			| Z80_INST2 expr
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST2, OpCode: int($1.TokenSubType), lineNumber: $1.LineNumber,
						Op2: $2}
			}
			| Z80_INST2 '(' expr ')' ',' expr
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST2, OpCode: int($1.TokenSubType), lineNumber: $1.LineNumber,
						Op1: &IndirectExpression{Expression: $3},
						Op2: $6}
			}
			| Z80_INST2 expr ',' '(' expr ')'
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST2, OpCode: int($1.TokenSubType), lineNumber: $1.LineNumber,
						Op1: $2,
						Op2: &IndirectExpression{Expression: $5}}
			}
			| Z80_INST2 '(' expr ')' ',' '(' expr ')'
			{
		
				$$ = &ParseError{Message: "両方のオペランドを間接指定にすることはできません"}
			}
			| Z80_INST2 expr ',' expr
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST2, OpCode: int($1.TokenSubType), lineNumber: $1.LineNumber,
						Op1: $2,
						Op2: $4}
			}
			;
	
expr_list	: 			{ $$ = &ExpressionList{Expressions: []Expression{}} }
			| expr		{ $$ = &ExpressionList{Expressions: []Expression{$1}} }
			| expr_list ',' expr
			{
				if $1.NodeType() == NODE_ERROR {
					$$ = $1
				} else {
					$1.Expressions = append($1.Expressions, $3)
					$$ = $1
				}
			}
			;

expr		: NUMBER
	 		{
				n, err := parseInt($1.Literal)
				if err == nil {
					$$ = &NumberLiteral{Value: int(n)}
				} else {
					$$ = &ParseError{Message: fmt.Sprintf("数値リテラル誤り: '%s'", $1.Literal)}
				}
			}
			| IDENT 		{ $$ = &Ident{Name: $1.Literal} }
			| DOT_IDENT
			{
				names := strings.Split(strings.ToUpper($1.Literal), ".")
				$$ = &DotIdent{Left: names[0], Right: names[1]}
			}
			| expr '(' expr_list ')'
			{
				if $1.NodeType() == NODE_ERROR {
					$$ = $1
				} else if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else {
					$$ = &CallExpression{Function: $1, Arguments: $3}
				}
			}
			| '[' expr_list ']'
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else {
					$$ = &ArrayLiteral{Elements: $2}
				}
			}
			| indexed_expr 			{ $$ = $1}
			| Z80_REG8 				{ $$ = &RegisterLiteral{RegisterType: int($1.TokenType), Register:int($1.TokenSubType)}}
			| Z80_REG16 			{ $$ = &RegisterLiteral{RegisterType: int($1.TokenType), Register:int($1.TokenSubType)}}
			| Z80_FLAG 				{ $$ = &FlagLiteral{Flag: int($1.TokenSubType)}}
			| '(' expr ')'			{ $$ = $2}
			| expr ADDSUB expr		{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3) }
			| expr '-' expr		 	{ $$ = buildInfixExpression('-', $1, $3) }
			| expr MULDIV expr		{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3) }
			| expr COMP expr 		{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3) }
			| expr SHIFT expr		{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3) }
			| expr OR expr			{ $$ = buildInfixExpression(OR, $1, $3) }
			| expr AND expr			{ $$ = buildInfixExpression(AND, $1, $3) }
			| '-' expr %prec UNARY	{ $$ = buildPrefixExpression('-', $2) }
			| UNARY expr 			{ $$ = buildPrefixExpression(int($1.TokenSubType), $2) }
			| error 
			{ 
				$$ = &ParseError{Message: __yyfmt__.Sprintf("[expr error] %s", $1.String())} 
				yyerrok()
			}
			;

indexed_expr: IDENT '[' ']'
			{
				$$ = &ParseError{Message: fmt.Sprintf("配列 %s のインデックス未指定", $1.Literal), lineNumber: $1.LineNumber}
			}
			| IDENT '[' expr ']'
			{
				if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else {
					$$ = &IndexedExpression{Ident: &Ident{Name: $1.Literal}, Index: $3, lineNumber: $1.LineNumber}
				}
			}
			;

%%

func Parse(l *Lexer) (*Program, int, int) {
	// error トークンでリカバリすると yyParse() は 0 を返す
	yyParse(l)
	ec, wc := l.ErrorStore.Count()
	return l.program, ec, wc
}