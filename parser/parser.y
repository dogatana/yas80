%{
package parser

import (
	"fmt"
	"strings"
	"yas80/errcode"
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
	ident *Ident
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
%type<expr> operand
%type<ident> ident


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
%token<token> FUNCTION // 1行関数

%token<token> IF ELSE ELIF ENDIF
%token<token> MACRO ENDM EXITM
%token<token> REPT ENDR
%token<token> FUNC ENDF RETURN
%token<token> PROC ENDP
%token<token> ENUM ENDE
%token<token> BLOCK ENDB
%token<token> FOR ENDFOR

%token<token>  '(' ')' ',' '<' '>' '~' '!' '^' '|' '+' '-' '*' '/' '&' ':' '[' ']' '='

%token<token> INVALID 
%token<token> error

// 演算の優先度の指定
%left OR              // ||
%left AND             // &&
%left COMP            // == != < <= > >=
%left ADDSUB          // ADDSUB + ^ |
%left MULDIV SHIFT    // MULDIV * / SHIFT << >> 
%right UNARY          // ~ ! -
%right UMINUS
%nonassoc '(' '[' 


%%
// 文法規則を指定
program		: { }
			| program statement 
			{
				if $2 == nil {
					// do nothing
				} else if $2.NodeType() == NODE_ERROR {
					yylex.Error($2.(*ParseError).Message, $2.(*ParseError).LineNumber)
				} else {
					prog := yylex.(*Lexer).program
					prog.Statements = append(prog.Statements, $2)
				}
			}
			;

statement   : EOL { $$ = nil }
			| label EOL 
			{
				$$ = &LabelStatement{Value: $1,LineNumber: $1.LineNumber}
			}
			| label instruction EOL
			{ 
				prog := yylex.(*Lexer).program
				stmt := &LabelStatement{Value: $1,LineNumber: $1.LineNumber}
				prog.Statements = append(prog.Statements, stmt)
				$$ = $2 
			}
			| instruction EOL	{ $$ = $1 }
			| directive	 EOL	{ $$ = $1 }
//			| expr EOL			
//			{ 
//				$$ = &ParseError{Message: "式文は無効",LineNumber: $2.LineNumber}
//				if $1.NodeType() == NODE_ERROR {
//					$$ = $1
//				} else {
//					$$ = &ExpressionStatement{Value: $1,LineNumber: $2.LineNumber}
//				}
//			}
			| error EOL
			{
				yylex.Error(__yyfmt__.Sprintf("[statement error] %s", $1.String()), $2.Position.LineNumber)
			}
			;

directive	: CONST ident '=' expr
			{ 
				if $4.NodeType() == NODE_ERROR {
					$$ =  $4
				} else {
					$$ = &ConstStatement{Name: &Ident{Name: $2.Name}, Value: $4,LineNumber: $1.Position.LineNumber}
				}
			}
			| ident EQU expr		
			{ 
				if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else {
					$$ = &ConstStatement{Name: &Ident{Name: $1.Name}, Value: $3,LineNumber: $2.Position.LineNumber}
				}
			}
			| ident ENUM EOL enum_elements ENDE
			{
				$$ = &EnumStatement{Name: $1.Name, Elements: $4,LineNumber: $2.Position.LineNumber}
			}
			| VAR ident '=' expr
			{
				if $4.NodeType() == NODE_ERROR {
					$$ = $4
				} else {
					$$ = &VariableStatement{Name: &Ident{Name: $2.Name}, Value: $4, LineNumber: $1.Position.LineNumber}
				}
			}
			| expr '=' expr
			{
				if $1.NodeType() == NODE_ERROR && $3.NodeType() == NODE_ERROR {
					err := $1.(*ParseError)
					yylex.Error(err.Message, err.LineNumber)
					$$ = $3

				} else if $1.NodeType() == NODE_ERROR {
					$$ = $1
				} else if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else {
					$$ = &AsignStatement{Left: $1, Value: $3,LineNumber: $2.Position.LineNumber}
				}
			}
			| REPT expr EOL block_statement ENDR
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else {
					$$ = &ReptStatement{MaxCount: $2, Block: $4,LineNumber: $1.Position.LineNumber}
				}
			}
			| IF expr EOL block_statement elseifs ENDIF
			{
				if $2.NodeType() == NODE_ERROR && $5.NodeType() == NODE_ERROR{
					err := $2.(*ParseError)
					yylex.Error(err.Message, err.LineNumber)
					$$ = $5
				} else if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else if $5 == nil {
					$$ = &IfStatement{Condition: $2, Consequence: $4, Alternative: &BlockStatement{Block: []Node{}},LineNumber: $1.Position.LineNumber}
				} else if $5.NodeType() == NODE_ERROR {
					$$ = $5
				} else if $5.NodeType() == NODE_BLOCK_STMT {
					$$ = &IfStatement{Condition: $2, Consequence: $4, Alternative: $5,LineNumber: $1.Position.LineNumber}
				} else {
					$$ = &ParseError{Message: "IF error",LineNumber: $1.Position.LineNumber}
				} 
			}
			| IF expr EOL block_statement elseifs ELSE block_statement ENDIF
			{
				if $2.NodeType() == NODE_ERROR && $5 != nil && $5.NodeType() == NODE_ERROR {
					err := $2.(*ParseError)
					yylex.Error(err.Message, err.LineNumber)
					$$ = $5
				} else if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else if $5 == nil  && $7 == nil {
					$$ = &IfStatement{ Condition: $2, Consequence: $4, Alternative: &BlockStatement{Block: []Node{}},LineNumber: $1.Position.LineNumber}
				} else if $5 == nil {
					$$ = &IfStatement{Condition: $2, Consequence: $4, Alternative: $7,LineNumber: $1.Position.LineNumber}
				} else if $5.NodeType() == NODE_ERROR {
					$$ = $5
				}  else if block, ok := $5.(*BlockStatement); ok {
					if len(block.Block) != 1 || block.Block[0].NodeType() != NODE_IF_STMT {
						$$ = &ParseError{Message: "IF-ELSE error",LineNumber: $1.Position.LineNumber}
					} else {
						last := getLastIfStatement(block.Block[0].(*IfStatement))
						if last.NodeType() == NODE_ERROR {
							$$ = last
						} else {
							last.(*IfStatement).Alternative = $7
							$$ = &IfStatement{Condition: $2, Consequence: $4, Alternative: $5,LineNumber: $1.Position.LineNumber}
						}
					}
				} else {
					$$ = &ParseError{Message: "IF-ELSE error", LineNumber: $1.Position.LineNumber}
				}
			}
			| ident FUNC param_list EOL block_statement ENDF
			{
				$$ = &FuncStatement{Name: $1.Name, Params: $3, Block: $5, LineNumber: $2.Position.LineNumber}
			}
			| FUNCTION ident '(' param_list ')' expr
			{ 
				$$ = &FuncStatement{
					Name: $2.Name, Params: $4, 
					Block: &BlockStatement{
						Block: []Node {&ReturnStatement{Value: $6,LineNumber: $1.Position.LineNumber}}}, 
						LineNumber: $1.Position.LineNumber}
			}
			| EXITM			{ $$ = &ExitmStatement{LineNumber: $1.Position.LineNumber}}
			| RETURN		{ $$ = &ReturnStatement{Value: nil,LineNumber: $1.Position.LineNumber}} 
			| RETURN expr	
			{ 
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else {
					$$ = &ReturnStatement{Value: $2,LineNumber: $1.Position.LineNumber}} 
				}
			| ident PROC	{ $$ = &ProcStatement{Name:$1.Name, IsStart: true,LineNumber: $2.Position.LineNumber }}
			| ENDP 			{ $$ = &ProcStatement{IsStart: false,LineNumber: $1.Position.LineNumber}}
			| IDENT MACRO param_list EOL block_statement ENDM
			{
				// macro 定義は ident でなく IDENT 
				$$ = &MacroStatement{Name: strings.ToUpper($1.Literal), Params: $3, Body: $5,LineNumber: $1.Position.LineNumber}
			}
			| IDENT expr_list 
			{
				$$ = &MacroCallStatement{Name: strings.ToUpper($1.Literal), Args: $2,LineNumber: $1.Position.LineNumber}
			}
			;
	
ident		: IDENT		 	{ $$ = &Ident{Name: strings.ToUpper($1.Literal), IdentType: IDENT,LineNumber: $1.Position.LineNumber}}
			| LOCAL_IDENT	{ $$ = &Ident{Name: strings.ToUpper($1.Literal), IdentType: LOCAL_IDENT,LineNumber: $1.Position.LineNumber}}
			;
param_list	: 			{ $$ = []string{}}
			| IDENT		{ $$ = []string{strings.ToUpper($1.Literal)} }
			| param_list ',' IDENT
			{
				$1 = append($1, strings.ToUpper($3.Literal))
				$$ = $1
			}
			;

elseifs		: { $$ = nil }
			| elseifs ELIF expr EOL block_statement 
			{ 
				ifst := &IfStatement{Condition: $3, Consequence: $5, Alternative: &BlockStatement{Block:[]Node{}},LineNumber: $2.Position.LineNumber}
				if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else if $1 == nil {
					$$ = &BlockStatement{Block: []Node{ifst}}
				} else if block := $1.(*BlockStatement); len(block.Block) == 0 {
					block.Block = append(block.Block, ifst)
					$$ = $1
				} else if len(block.Block) == 1 && block.Block[0].NodeType() == NODE_IF_STMT {
					stmt := block.Block[0].(*IfStatement)
					for {
						if stmt.Alternative == nil {
							stmt.Alternative = &BlockStatement{Block:[]Node{ifst}}
							$$ = $1
							break
						} else if block := stmt.Alternative.(*BlockStatement); len(block.Block) == 0 {
							block.Block = append(block.Block, ifst)
							$$ = $1
							break
						} else if len(block.Block) == 1 && block.Block[0].NodeType() == NODE_IF_STMT {
							stmt = block.Block[0].(*IfStatement)
							continue
						} else {
							$$ = &ParseError{Message: fmt.Sprintf("elseif error %s", $1.String()),LineNumber: $2.Position.LineNumber}
							break
						}
					}
				}
			}
			;

	
// statement エラー検出時は yylex.Error() を呼んで伝播を止める
block_statement	: 	 				{ $$ = &BlockStatement{Block: []Node{}} }
			| block_statement statement 
			{ 
				if $2 == nil { // error
					// do nothing
				} else if $2.NodeType() == NODE_ERROR {
					err := $2.(*ParseError)
					yylex.Error(err.Message, err.LineNumber)
				} else {
					$1.Block = append($1.Block, $2.(Statement))
					$$ = $1
				}
			}
			;
	
// enum_element（実質 statement)エラー検出時は yylex.Error() を呼んで伝播を止める
enum_elements : 	 			{ $$ = &EnumElements{Elements: []*EnumElement{}} }
			| enum_elements EOL { $$ = $1 }
			| enum_elements enum_element EOL
			{
				if $2.NodeType() == NODE_ERROR {
					err := $2.(*ParseError)
					yylex.Error(err.Message, err.LineNumber)
				}
				$1.Elements = append($1.Elements, $2.(*EnumElement))
				$$ = $1
			}
			;

enum_element : IDENT 			{ $$ = &EnumElement{Name: strings.ToUpper($1.Literal), Value: nil,LineNumber: $1.Position.LineNumber} }
			| IDENT '=' expr	
			{ 
				if $3.NodeType() == NODE_ERROR {
					$$ = $3
				} else {
					stmt := &ExpressionStatement{Value:$3,LineNumber: $1.Position.LineNumber} 
					$$ = &EnumElement{Name: strings.ToUpper($1.Literal), Value: stmt, LineNumber: $1.Position.LineNumber }
				}
			}
			;

label		: IDENT ':'
			{
				$$ = &Label{LabelType: NODE_LABEL, Name: strings.ToUpper($1.Literal),LineNumber: $1.Position.LineNumber}
			}
			| LOCAL_IDENT ':'
			{
				$$ = &Label{LabelType: NODE_LOCAL_LABEL, Name: strings.ToUpper($1.Literal),LineNumber: $1.Position.LineNumber}
			}
			| AT_IDENT ':'
			{
				$$ = &Label{LabelType: NODE_AT_LABEL, Name: strings.ToUpper($1.Literal),LineNumber: $1.Position.LineNumber}
			}
//			| LOCAL_IDENT
//			{
//				$$ = &Label{nodeType: NODE_LOCAL_LABEL, Name: strings.ToUpper($1.Literal),LineNumber: $1.Position.LineNumber}
//			}
			;


instruction	: Z80_INST0
			{
				$$ = &Z80Instruction{
					InstType: Z80_INST0, Opcode: int($1.TokenSubType),LineNumber: $1.Position.LineNumber} 
			}
			| Z80_INST1
			{
				$$ = &Z80Instruction{
						InstType: Z80_INST1, Opcode: int($1.TokenSubType),LineNumber: $1.Position.LineNumber}
			}
			| Z80_INST1 operand
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else {
					$$ = &Z80Instruction{InstType: Z80_INST1, Opcode: int($1.TokenSubType), 
						Op1: $2, 
						LineNumber: $1.Position.LineNumber }
				}
			}
			| Z80_INST2 operand
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else {
					$$ = &Z80Instruction{InstType: Z80_INST2, Opcode: int($1.TokenSubType), 
						Op2: $2,
						LineNumber: $1.Position.LineNumber }
				}
			}
			| Z80_INST2 operand ',' operand
			{
				if $2.NodeType() == NODE_ERROR && $4.NodeType() == NODE_ERROR {
					err := $2.(*ParseError)
					yylex.Error(err.Message, err.LineNumber)
					$$ = $4
				} else if $2.NodeType() == NODE_ERROR {
					$$ = $2
				} else if $4.NodeType() == NODE_ERROR {
					$$ = $4
				} else {
					$$ = &Z80Instruction{InstType: Z80_INST2, Opcode: int($1.TokenSubType), 
							Op1: $2,
							Op2: $4,
							LineNumber: $1.Position.LineNumber }
				}
			}
			;
	

operand	: '(' Z80_REG16 ')'
			{ 
				$$ = &IndirectExpression{Expression: 
					&RegisterLiteral{RegisterType: int($2.TokenType), Register: int($2.TokenSubType),LineNumber: $2.Position.LineNumber}}
			}
			| '(' Z80_REG16 ADDSUB expr ')' 
			{
				$$ = &IndirectExpression{Expression: 
						buildInfixExpression(
							int($3.TokenSubType), 
							&RegisterLiteral{RegisterType: int($2.TokenType), Register: int($2.TokenSubType),LineNumber: $2.Position.LineNumber},
							$4,
							$2.Position.LineNumber)}
			}
			| '(' Z80_REG8 ')'
			{ 
				$$ = &IndirectExpression{Expression: 
					&RegisterLiteral{RegisterType: int($2.TokenType), Register: int($2.TokenSubType),LineNumber: $2.Position.LineNumber}}
			}
			| '(' expr ')'		{ $$ = &IndirectExpression{Expression: $2} }
			| expr 				{ $$ = $1 }
			;
			

// expr エラー検出時は yylex.Error() を呼んで伝播を止める
expr_list	:   { $$ = &ExpressionList{Expressions: []Expression{}}}
			| expr		// 最低でも 1 個の引数で必要
			{ 
				if $1.NodeType() == NODE_ERROR {
					err := $1.(*ParseError)
					yylex.Error(err.Message, err.LineNumber)
				}
				$$ = &ExpressionList{Expressions: []Expression{$1}}
			}
			| expr_list ',' expr
			{
				if $3.NodeType() == NODE_ERROR {
					err := $3.(*ParseError)
					yylex.Error(err.Message, err.LineNumber)
				}
				$1.Expressions = append($1.Expressions, $3)
				$$ = $1
			}
			;
	
expr		: NUMBER
			{
				n, err := parseInt($1.Literal)
				if err == nil {
					$$ = &NumberLiteral{Value: int(n),LineNumber: $1.Position.LineNumber}
				} else {
					$$ = &ParseError{Message: fmt.Sprintf(errcode.E002, $1.Literal),LineNumber: $1.Position.LineNumber}
				}
			}
			| STRING 		{ $$ = &StringLiteral{Value: $1.Literal,LineNumber: $1.Position.LineNumber} }
			| Z80_REG8 		{ $$ = &RegisterLiteral{RegisterType: int($1.TokenType), Register:int($1.TokenSubType),LineNumber:$1.Position.LineNumber}}
			| Z80_REG16 	{ $$ = &RegisterLiteral{RegisterType: int($1.TokenType), Register:int($1.TokenSubType),LineNumber:$1.Position.LineNumber}}
			| Z80_FLAG 		{ $$ = &FlagLiteral{Flag: int($1.TokenSubType),LineNumber:$1.Position.LineNumber}}
			| IDENT 		{ $$ = &Ident{Name: strings.ToUpper($1.Literal), IdentType: IDENT,LineNumber: $1.Position.LineNumber} }
			| LOCAL_IDENT 	{ $$ = &Ident{Name: strings.ToUpper($1.Literal), IdentType: LOCAL_IDENT,LineNumber: $1.Position.LineNumber} }
			| DOT_IDENT
			{
				uname := strings.ToUpper($1.Literal)
				names := strings.Split(uname, ".")
				$$ = &DotIdent{Name: uname, Left: names[0], Right: names[1],LineNumber: $1.Position.LineNumber}
			}
//			| IDENT '(' ')'
//			{
//				$$ = &CallExpression{
//					Function: &Ident{Name: strings.ToUpper($1.Literal), IdentType: IDENT,LineNumber: $1.Position.LineNumber}, 
//					Arguments: &ExpressionList{Expressions: []Expression{}},
//					LineNumber: $1.Position.LineNumber}
//			}
			| IDENT '(' expr_list ')'
			{
				$$ = &CallExpression{
					Function: &Ident{Name: strings.ToUpper($1.Literal), IdentType: IDENT,LineNumber: $1.Position.LineNumber}, 
					Arguments: $3, 
					LineNumber: $1.Position.LineNumber}
			}
//			| '[' ']'			{ $$ = &ArrayLiteral{Elements: &ExpressionList{Expressions: []Expression{}},LineNumber: $1.Position.LineNumber}}
			| '[' expr_list ']' { $$ = &ArrayLiteral{Elements: $2,LineNumber: $1.Position.LineNumber} }
			| indexed_expr 			{ $$ = $1}
			| '(' expr ')'			{ $$ = $2}
			| expr ADDSUB expr		{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3, $2.Position.LineNumber) }
			| expr MULDIV expr		{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3, $2.Position.LineNumber) }
			| expr COMP expr 		{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3, $2.Position.LineNumber) }
			| expr SHIFT expr		{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3, $2.Position.LineNumber) }
			| expr OR expr			{ $$ = buildInfixExpression(OR, $1, $3, $2.Position.LineNumber) }
			| expr AND expr			{ $$ = buildInfixExpression(AND, $1, $3, $2.Position.LineNumber) }
			| ADDSUB expr %prec UNARY	{ $$ = buildPrefixExpression(int($1.TokenSubType), $2, $1.Position.LineNumber) }
			| UNARY expr 			{ $$ = buildPrefixExpression(int($1.TokenSubType), $2, $1.Position.LineNumber) }
			;

indexed_expr: expr '[' ']'
			{
				if $1.NodeType() == NODE_ERROR {
					err := $1.(*ParseError)
					yylex.Error(err.Message, err.LineNumber)
					yylex.Error(errcode.E003, $2.Position.LineNumber)
				} 
				$$ = &ParseError{Message: errcode.E004,LineNumber: $2.Position.LineNumber}
			}
			| expr '[' expr ']'
			{
				if $1.NodeType() == NODE_ERROR && $3.NodeType() == NODE_ERROR {
					yylex.Error(errcode.E003, $2.Position.LineNumber)
					err := $1.(*ParseError)
					yylex.Error(err.Message, err.LineNumber)

					yylex.Error(errcode.E005, $2.Position.LineNumber)
					$$ = $3
				} else if $1.NodeType() == NODE_ERROR {
					yylex.Error(errcode.E003, $2.Position.LineNumber)
					$$ = $1
				} else if $3.NodeType() == NODE_ERROR {
					yylex.Error(errcode.E005, $2.Position.LineNumber)
					$$ = $3
				} else {
					$$ = &IndexedExpression{Left: $1, Index: $3,LineNumber: $2.Position.LineNumber}
				}
			}
			;

%%


func Parse(l *Lexer) (*Program) {
	// 常に有効
	yyErrorVerbose = true
	// error トークンでリカバリすると yyParse() は 0 を返すため、戻り値には意味がない
	yyParse(l)
	return l.program
}