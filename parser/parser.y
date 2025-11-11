%{
package parser

import (
	"fmt"
	"strings"
	"yas80/logger"
)

// goyacc が __yyfmt__ を勝手に import することの対策
var _ = __yyfmt__.Sprintf
%}
%union {
	token Token
	num Node
	node Node
	err any
	label *Label
	enum_element *EnumElement
	enum_elements *EnumElements
	block *BlockStatement
	params []string
	expr_list *ExpressionList
	expr Expression
}


// プログラムの構成要素を指定
%type<node> statement instruction directive elseifs enum_element
%type<label> label
%type<block> block_statement
%type<enum_elements> enum_elements
%type<params> param_list
%type<expr_list> expr_list
%type<expr> expr indexed_expr


%token<token> EOL
%token<token> NUMBER IDENT STRING
%token<token> AT_IDENT    // @def 
%token<token> LOCAL_IDENT // .def 
%token<token> DOT_IDENT   // abc.def ラベル, enum

%token<token> Z80_INST0 Z80_INST1 Z80_INST2 Z80_REG8 Z80_REG16 Z80_FLAG

%token<token> ADDSUB MULDIV COMP SHIFT UNARY
%token<token> SL SR EQ NEQ GE LE OR AND

%token<token> ORG
%token<token> CONST VAR EQU
%token<token> FN // 1行関数

%token<token> IF ELSE ELIF ENDIF
%token<token> MACRO ENDM
%token<token> REPEAT ENDR
%token<token> FUNC ENDF
%token<token> PROC ENDP
%token<token> ENUM ENDE
%token<token> BLOCK ENDB
%token<token> FOR ENDFOR

%token<token>  '(' ')' ',' '<' '>' '~' '!' '^' '|' '+' '-' '*' '/' '&' ':' '[' ']' '='

%token<token> INVALID 
%token<token> error

// 演算の優先度の指定
%left OR                       // ||
%left AND                      // &&
%left COMP                     // == != < <= > >=
%left ADDSUB '|' '^' '-'       // ADDSUB + ^ |
%left MULDIV SHIFT             // MULDIV * / SHIFT << >> 
%right UNARY                   // ~ ! -
%right UMINUS
%nonassoc '(' '[' 


%%
// 文法規則を指定
program		: { }
			| program EOL
			| program label EOL
			{
				stmt := &LabelStatement{Value: $2, lineNumber: $3.LineNumber}
				prog := yylex.(*Lexer).program
				prog.Statements = append(prog.Statements, stmt)
			}
			| program label statement
			{
				if $3.NodeType() == NODE_ERROR {
					yylex.Error($3.(*ParseError).Message, $3.LineNumber())
				} else {
					prog := yylex.(*Lexer).program
					stmt := &LabelStatement{Value: $2, lineNumber: $2.LineNumber()}
					prog.Statements = append(prog.Statements, stmt, $3)
				}
			}
			| program statement 
			{
				if $2.NodeType() == NODE_ERROR {
					yylex.Error($2.(*ParseError).Message, $2.LineNumber())
				} else {
					prog := yylex.(*Lexer).program
					prog.Statements = append(prog.Statements, $2)
				}
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
			| instruction EOL	{ $$ = $1 }
			| directive	 EOL	{ $$ = $1 }
			| error EOL
			{
				yylex.Error(__yyfmt__.Sprintf("[statement error] %s", $1.String()), $2.LineNumber)
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
			| IDENT ENUM EOL enum_elements ENDE
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
			| expr '=' expr
			{
				if $1.NodeType() == NODE_ERROR {
					$$ = $1
				} else if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else {
					$$ = &AsignStatement{Left: $1, Value: $3, lineNumber: $2.LineNumber}
				}
			}
			| indexed_expr '=' expr 
			{
				if $1.NodeType() == NODE_ERROR {
					$$ = $1
				} else if $3.NodeType() == NODE_ERROR {
					$3 = $3	
				} else {
					$$ = &AsignStatement{Left: $1, Value: $3, lineNumber: $1.(*IndexedExpression).lineNumber}
				}
			}
			| REPEAT expr EOL block_statement ENDR
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else if $4.NodeType() == NODE_ERROR {
					$$ = $4
				} else {
					$$ = &RepeatStatement{MaxCount: $2, Block: $4, lineNumber: $1.LineNumber}
				}
			}
			| IF expr EOL block_statement elseifs ENDIF
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
			| IF expr EOL block_statement elseifs ELSE block_statement ENDIF
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
			| IDENT FUNC param_list EOL block_statement ENDF
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
			}
			;

elseifs		: { $$ = nil }
			| elseifs ELIF expr EOL block_statement 
			{ 
				if $3.NodeType() == NODE_ERROR {
					$$ = $3
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

	
// statement エラー検出時は yylex.Err() を呼んで伝播を止める
block_statement	: 	 				{ $$ = &BlockStatement{Block: []Statement{}} }
			| block_statement EOL 	{ $$ = $1}
			| block_statement statement 
			{ 
				if $2.NodeType() == NODE_ERROR {
					err := $2.(*ParseError)
					yylex.Error(err.Message, err.LineNumber())
				}
				$1.Block = append($1.Block, $2.(Statement))
				$$ = $1
			}
			;
	
// enum_element（実質 statement)エラー検出時は yylex.Err() を呼んで伝播を止める
enum_elements : 	 			{ $$ = &EnumElements{Elements: []*EnumElement{}} }
			| enum_elements EOL { $$ = $1 }
			| enum_elements enum_element EOL
			{
				if $2.NodeType() == NODE_ERROR {
					err := $2.(*ParseError)
					yylex.Error(err.Message, err.LineNumber())
				}
				$1.Elements = append($1.Elements, $2.(*EnumElement))
				$$ = $1
			}
			;

enum_element : IDENT 			{ $$ = &EnumElement{Name: $1.Literal, Value: nil, lineNumber: $1.LineNumber} }
			| IDENT '=' expr	
			{ 
				if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else {
					stmt := &ExpressionStatement{Value:$3, lineNumber: $3.LineNumber()} 
					$$ = &EnumElement{Name: $1.Literal, Value: stmt, lineNumber: $1.LineNumber }
				}
			}
			;

label		: IDENT ':'
			{
				$$ = &Label{nodeType: NODE_LABEL, Name: $1.Literal, lineNumber: $1.LineNumber}
			}
			| LOCAL_IDENT ':'
			{
				// info のみ表示し、処理継続
				yylex.Error(logger.I001, $1.LineNumber)
				$$ = &Label{nodeType: NODE_LOCAL_LABEL, Name: $1.Literal, lineNumber: $1.LineNumber}
			}
			| LOCAL_IDENT
			{
				$$ = &Label{nodeType: NODE_LOCAL_LABEL, Name: $1.Literal, lineNumber: $1.LineNumber}
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
				if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else {
					$$ = &Z80Instruction{
							InstType: Z80_INST1, OpCode: int($1.TokenSubType), 
							Op1: &IndirectExpression{Expression: $3},
							lineNumber: $1.LineNumber }
				}
			}
			| Z80_INST1 expr
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else {
					$$ = &Z80Instruction{
							InstType: Z80_INST1, OpCode: int($1.TokenSubType), 
							Op1: $2,
							lineNumber: $1.LineNumber }
				}
			}
			| Z80_INST2 '(' expr ')'
			{
				if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else {
					$$ = &Z80Instruction{
							InstType: Z80_INST2, OpCode: int($1.TokenSubType), 
							Op2: &IndirectExpression{Expression: $3},
							lineNumber: $1.LineNumber }
				}
			}
			| Z80_INST2 expr
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else {
					$$ = &Z80Instruction{
							InstType: Z80_INST2, OpCode: int($1.TokenSubType), 
							Op2: $2,
							lineNumber: $1.LineNumber }
				}
			}
			| Z80_INST2 '(' expr ')' ',' expr
			{
				if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else if $6.NodeType() == NODE_ERROR {
					$$ = $6
				} else {

					$$ = &Z80Instruction{
							InstType: Z80_INST2, OpCode: int($1.TokenSubType), 
							Op1: &IndirectExpression{Expression: $3},
							Op2: $6,
							lineNumber: $1.LineNumber }
				}
			}
			| Z80_INST2 expr ',' '(' expr ')'
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else if $5.NodeType() == NODE_ERROR {
					$$ = $5
				} else {
					$$ = &Z80Instruction{
							InstType: Z80_INST2, OpCode: int($1.TokenSubType), 
							Op1: $2,
							Op2: &IndirectExpression{Expression: $5},
							lineNumber: $1.LineNumber }
				}
			}
			| Z80_INST2 '(' expr ')' ',' '(' expr ')'
			{
				$$ = &ParseError{Message: logger.E006, lineNumber: $1.LineNumber}
			}
			| Z80_INST2 expr ',' expr
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else if $4.NodeType() == NODE_ERROR {
					$$ = $4
				} else {
					$$ = &Z80Instruction{
							InstType: Z80_INST2, OpCode: int($1.TokenSubType),
							Op1: $2,
							Op2: $4,
							lineNumber: $1.LineNumber}
				}
			}
			;
	
// expr エラー検出時は yylex.Error() を呼んで伝播を止める
expr_list	: 			{ $$ = &ExpressionList{Expressions: []Expression{}} }
			| expr
			{ 
				if $1.NodeType() == NODE_ERROR {
					err := $1.(*ParseError)
					yylex.Error("[E]" + err.Message, err.LineNumber())
				}
				$$ = &ExpressionList{Expressions: []Expression{$1}}
			}
			| expr_list ',' expr
			{
				if $3.NodeType() == NODE_ERROR {
					err := $3.(*ParseError)
					yylex.Error("[E]" + err.Message, err.LineNumber())
				}
				$1.Expressions = append($1.Expressions, $3)
				$$ = $1
			}
			;

expr		: NUMBER
	 		{
				n, err := parseInt($1.Literal)
				if err == nil {
					$$ = &NumberLiteral{Value: int(n), lineNumber: $1.LineNumber}
				} else {
					$$ = &ParseError{Message: fmt.Sprintf(logger.E002, $1.Literal), lineNumber: $1.LineNumber}
				}
			}
			| STRING 		{ $$ = &StringLiteral{Value: $1.Literal, lineNumber: $1.LineNumber} }
			| IDENT 		{ $$ = &Ident{Name: $1.Literal, lineNumber: $1.LineNumber} }
			| DOT_IDENT
			{
				names := strings.Split(strings.ToUpper($1.Literal), ".")
				$$ = &DotIdent{Left: names[0], Right: names[1], lineNumber: $1.LineNumber}
			}
			| expr '(' expr_list ')'
			{
				if $1.NodeType() == NODE_ERROR {
					$$ = $1
				} else if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else {
					$$ = &CallExpression{Function: $1, Arguments: $3, lineNumber: $1.LineNumber()}
				}
			}
			| '[' expr_list ']'
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else {
					$$ = &ArrayLiteral{Elements: $2, lineNumber: $1.LineNumber}
				}
			}
			| indexed_expr 			{ $$ = $1}
			| Z80_REG8 				{ $$ = &RegisterLiteral{RegisterType: int($1.TokenType), Register:int($1.TokenSubType), lineNumber:$1.LineNumber}}
			| Z80_REG16 			{ $$ = &RegisterLiteral{RegisterType: int($1.TokenType), Register:int($1.TokenSubType), lineNumber:$1.LineNumber}}
			| Z80_FLAG 				{ $$ = &FlagLiteral{Flag: int($1.TokenSubType), lineNumber:$1.LineNumber}}
			| '(' expr ')'			{ $$ = $2}
			| expr ADDSUB expr		{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3, $2.LineNumber) }
			| expr '-' expr		 	{ $$ = buildInfixExpression('-', $1, $3, $2.LineNumber) }
			| expr MULDIV expr		{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3, $2.LineNumber) }
			| expr COMP expr 		{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3, $2.LineNumber) }
			| expr SHIFT expr		{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3, $2.LineNumber) }
			| expr OR expr			{ $$ = buildInfixExpression(OR, $1, $3, $2.LineNumber) }
			| expr AND expr			{ $$ = buildInfixExpression(AND, $1, $3, $2.LineNumber) }
			| '-' expr %prec UNARY	{ $$ = buildPrefixExpression('-', $2, $1.LineNumber) }
			| UNARY expr 			{ $$ = buildPrefixExpression(int($1.TokenSubType), $2, $1.LineNumber) }
			;

indexed_expr: expr '[' ']'
			{
				if $1.NodeType() == NODE_ERROR {
					$$ = $1
					yylex.Error(logger.E003, $2.LineNumber)
				} else {
					$$ = &ParseError{Message: logger.E004, lineNumber: $2.LineNumber}
				}
			}
			| expr '[' expr ']'
			{
				if $1.NodeType() == NODE_ERROR {
					$$ = $1
					yylex.Error(logger.E003, $2.LineNumber)
				} else if $3.NodeType() == NODE_ERROR {
					$$ = $3
					yylex.Error(logger.E005, $2.LineNumber)
				} else {
					$$ = &IndexedExpression{Left: $1, Index: $3, lineNumber: $2.LineNumber}
				}
			}
			;

%%

func Parse(l *Lexer) (*Program) {
	// error トークンでリカバリすると yyParse() は 0 を返すため、戻り値には意味がない
	yyParse(l)
	return l.program
}