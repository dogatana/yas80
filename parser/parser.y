%{
package parser

import (
	"fmt"
	"strings"
	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/filecontent"
	"github.com/dogatana/yas80/intern"
)

// goyacc が __yyfmt__ を勝手に import することの対策
var _ = __yyfmt__.Sprintf
%}
%union {
	token Token
	statement Statement
	expr Expression
	err any
	params *[]string
}


// 非終端記号
%type<statement> statement instruction directive elseifs enum_element datadef datastore
%type<statement> block_statement
%type<statement> enum_elements
%type<params> param_list
%type<expr> expr_list
%type<expr> expr indexed_expr ident_expr
%type<expr> operand
%type<expr> ident
%type<token> string

// 終端記号
%token<token> EOL FILE INCLUDE CHARMAP
%token<token> NUMBER STRING
%token<token> IDENT 
%token<token> AT_IDENT    // @def 
%token<token> LOCAL_IDENT // .def 
%token<token> ANON_IDENT // @@ @f @b
%token<token> DOT_IDENT   // abc.def ラベル, enum

%token<token> Z80_INST0 Z80_INST1 Z80_INST2 Z80_REG8 Z80_REG16 Z80_FLAG

%token<token> ADDSUB MULDIV COMP SHIFT UNARY
%token<token> SL SR EQ NEQ GE LE OR AND
%token<token> CONCAT // シンボル結合演算子 ident ## expr

%token<token> ORG
%token<token> CONST EQU VAR
%token<token> FUNCTION // 1行関数

%token<token> DATA DB DEFB DW DEFW
%token<token> DS DSB DSW DD

%token<token> IF ELSE ELIF ENDIF
%token<token> MACRO ENDM EXITM
%token<token> REPT ENDR
%token<token> FUNC ENDF RETURN
%token<token> PROC ENDP
%token<token> ENUM ENDE
// %token<token> BLOCK ENDB // 予約
// %token<token> FOR ENDFOR // 予約
%token<token> END

%token<token>  '(' ')' ',' '<' '>' '~' '!' '^' '|' '+' '-' '*' '/' '&' ':' '[' ']' '=' '%'

%token<token> INVALID 

// 演算子優先度
%nonassoc CONCAT
%left OR              // ||
%left AND             // &&
%left COMP            // == != < <= > >=
%left ADDSUB          // ADDSUB + ^ |
%left MULDIV SHIFT    // MULDIV * / % SHIFT << >> 
%right UNARY          // ~ ! -
%nonassoc '(' '[' 

%%
// 文法規則を指定
program		: { }
			| program statement 
			{
				if $2 == nil {
					// do nothing
				} else if $2.NodeType() == NODE_ERROR {
					yylex.Error($2.(*ParseError).Message, $2.(*ParseError).Context)
				} else {
					prog := yylex.(*Lexer).program
					prog.Block = append(prog.Block, $2)
					// END なら構文解析を終了する
					if $2.NodeType() == NODE_END_STMT {
						return 0
						// YYACCEPT
					}
				}
			}
			;

statement   : EOL { $$ = nil }
			| FILE	{ $$ = &FileStatement{Filename: $1.SymbolID.String(), Line: int($1.TokenSubType)} }
			| ident_expr ':'  EOL 
			{
				if $1.NodeType() == NODE_ERROR {
					$$ =  $1.(*ParseError)
				} else {
					$$ = &LabelStatement{Name: $1, Context: newCtxFromTokenCtx($2.Context)}
				}
			}
			| ident_expr ':' instruction EOL
			{ 
				if $1.NodeType() == NODE_ERROR {
					$$ =  $1.(*ParseError)
				} else {
					$3.(*Z80Instruction).Label = $1
					$$ = $3 
				}
			}
			| instruction EOL	{ $$ = $1 }
			| directive	 EOL	{ 
				if inc, ok := $1.(*IncludeStatement); ok {
					fc, err := loadIncludeFile(inc.Context.FileContent.Filename, inc.Filename)
					if err != nil {
						$$ = &ParseError{Message: fmt.Sprintf(errcode.EFILE_NOT_FOUND, inc.Filename), Context: inc.Context}
					} else if err := yylex.Push(inc.Filename, fc, inc.Context); err != nil {
						$$ = &ParseError{Message: err.Error(), Context: inc.Context}
					} else {
						$$ = $1
					}
				} else {
					$$ = $1 
				}
			}
			| END EOL {
				$$ = &EndStatement{Start: nil, Context: newCtxFromTokenCtx($1.Context)}
			}
			| END expr EOL {
				if $2.NodeType() == NODE_ERROR {
					$$ = $2.(*ParseError)
				} else {
					$$ = &EndStatement{Start: $2, Context: newCtxFromTokenCtx($1.Context)}
				}
			}
			| error EOL
			{
				// ここで対処する前に syntax error が出力されているので改めてエラー出力はしない
				// lx.Error(__yyfmt__.Sprintf("[statement error] %s), $2.$1.String()), Context)
			}
			;

directive	: CONST ident_expr '=' expr
			{ 
				if $2.NodeType() == NODE_ERROR {
					$$ =  $2.(*ParseError)
				} else if $4.NodeType() == NODE_ERROR {
					$$ =  $4.(*ParseError)
				} else {
					$$ = &ConstStatement{Name: $2, Value: $4, Context: newCtxFromTokenCtx($1.Context)}
				}
			}
			| ident_expr EQU expr		
			{ 
				if $1.NodeType() == NODE_ERROR {
					$$ = $1.(*ParseError)
				} else if $3.NodeType() == NODE_ERROR {
					$$ = $3.(*ParseError)
				} else {
					$$ = &ConstStatement{Name: $1, Value: $3, Context: newCtxFromTokenCtx($2.Context)}
				}
			}
			| ident_expr PROC EOL block_statement ENDP
			{
				if $1.NodeType() == NODE_ERROR {
					$$ = $1.(*ParseError)
				} else {
					$$ = &ProcStatement{Name: $1, Block: $4.(*BlockStatement), Context: newCtxFromTokenCtx($2.Context)}
				}
			}
			| ident ENUM EOL enum_elements ENDE
			{
				id := $1.(*Ident)
				e := $4.(*EnumElements)
				$$ = &EnumStatement{NameID: id.NameID, Elements: e, Context: newCtxFromTokenCtx($2.Context)}
			}
			| VAR ident '=' expr
			{
				if $4.NodeType() == NODE_ERROR {
					$$ = $4.(*ParseError)
				} else {
					id := $2.(*Ident)
					$$ = &VariableStatement{Name: &Ident{Name: id.Name, NameID: id.NameID}, Value: $4, Context: newCtxFromTokenCtx($1.Context)}
				}
			}
			| expr '=' expr
			{
				if $1.NodeType() == NODE_ERROR && $3.NodeType() == NODE_ERROR {
					// 左辺のエラーを出力し、右辺のエラーを返す
					err := $1.(*ParseError)
					yylex.Error(err.Message, err.Context)
					$$ = $3.(*ParseError)

				} else if $1.NodeType() == NODE_ERROR {
					$$ = $1.(*ParseError)
				} else if $3.NodeType() == NODE_ERROR {
					$$ = $3.(*ParseError)
				} else {
					$$ = &AssignStatement{Left: $1, Value: $3, Context: newCtxFromTokenCtx($2.Context)}
				}
			}
			| REPT expr EOL block_statement ENDR
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2.(*ParseError)
				} else {
					$$ = &ReptStatement{MaxCount: $2, Block: $4.(*BlockStatement), Start: int($1.Context.Line), Context: newCtxFromTokenCtx($5.Context)}
				}
			}
			| ident_expr ':' REPT expr EOL block_statement ENDR
			{
				if $4.NodeType() == NODE_ERROR {
					$$ = $4.(*ParseError)
				} else {
					$$ = &ReptStatement{Label: $1, MaxCount: $4, Block: $6.(*BlockStatement), Start: int($3.Context.Line), Context: newCtxFromTokenCtx($7.Context)}
				}
			}
			| IF expr EOL block_statement elseifs ENDIF
			{
				if $2.NodeType() == NODE_ERROR && $5.NodeType() == NODE_ERROR{
					// 条件式のエラーを出力し、elseifs のエラーを返す
					err := $2.(*ParseError)
					yylex.Error(err.Message, err.Context)
					$$ = $5.(*ParseError)
				} else if $2.NodeType() == NODE_ERROR {
					$$ = $2.(*ParseError)
				} else if $5 == nil {
					$$ = &IfStatement{Condition: $2, Consequence: $4.(*BlockStatement), Alternative: &BlockStatement{Block: []Statement{}}, Context: newCtxFromTokenCtx($1.Context)}
				} else if $5.NodeType() == NODE_ERROR {
					$$ = $5.(*ParseError)
				} else if $5.NodeType() == NODE_BLOCK_STMT {
					$$ = &IfStatement{Condition: $2, Consequence: $4.(*BlockStatement), Alternative: $5, Context: newCtxFromTokenCtx($1.Context)}
				} else {
					$$ = &ParseError{Message: fmt.Sprintf(errcode.EINTERNAL, "IF"), Context: newCtxFromTokenCtx($1.Context)}
				} 
			}
			| IF expr EOL block_statement elseifs ELSE block_statement ENDIF
			{
				if $2.NodeType() == NODE_ERROR && $5 != nil && $5.NodeType() == NODE_ERROR {
					// 条件式のエラーを出力し、elseifs のエラーを返す
					err := $2.(*ParseError)
					yylex.Error(err.Message, err.Context)
					$$ = $5
				} else if $2.NodeType() == NODE_ERROR {
					$$ = $2.(*ParseError)
				} else if $5 == nil  && $7 == nil {
					$$ = &IfStatement{ Condition: $2, Consequence: $4.(*BlockStatement), Alternative: &BlockStatement{Block: []Statement{}}, Context: newCtxFromTokenCtx($1.Context)}
				} else if $5 == nil {
					$$ = &IfStatement{Condition: $2, Consequence: $4.(*BlockStatement), Alternative: $7, Context: newCtxFromTokenCtx($1.Context)}
				} else if $5.NodeType() == NODE_ERROR {
					$$ = $5
				}  else if block, ok := $5.(*BlockStatement); ok {
					if len(block.Block) != 1 || block.Block[0].NodeType() != NODE_IF_STMT {
						$$ = &ParseError{Message: fmt.Sprintf(errcode.EINTERNAL, "IF-ELSE"), Context: newCtxFromTokenCtx($1.Context)}
					} else {
						last := getLastIfStatement(block.Block[0].(*IfStatement))
						if last.NodeType() == NODE_ERROR {
							$$ = last.(*ParseError)
						} else {
							last.(*IfStatement).Alternative = $7
							$$ = &IfStatement{Condition: $2, Consequence: $4, Alternative: $5, Context: newCtxFromTokenCtx($1.Context)}
						}
					}
				} else {
					$$ = &ParseError{Message: fmt.Sprintf(errcode.EINTERNAL, "IF-ELSE"), Context: newCtxFromTokenCtx($1.Context)}
				}
			}
			| ident FUNC param_list EOL block_statement ENDF
			{
				id := $1.(*Ident)
				$$ = &FuncStatement{NameID: id.NameID, Params: *$3, Block: $5.(*BlockStatement), Context: newCtxFromTokenCtx($2.Context)}
			}
			| FUNCTION ident '(' param_list ')' expr
			{ 
				if $6.NodeType() == NODE_ERROR {
					$$ = $6.(*ParseError)
				} else {
					ctx := newCtxFromTokenCtx($1.Context)
					$$ = &FuncStatement{
						NameID: $2.(*Ident).NameID, 
						Params: *$4, 
						Block: &BlockStatement{Block: []Statement {&ReturnStatement{Value: $6, Context: ctx}}},
						Context: ctx}
				}
			}
			| RETURN		{ $$ = &ReturnStatement{Value: nil, Context: newCtxFromTokenCtx($1.Context)}} 
			| RETURN expr	
			{ 
				if $2.NodeType() == NODE_ERROR {
					$$ = $2.(*ParseError)
				} else {
					$$ = &ReturnStatement{Value: $2, Context: newCtxFromTokenCtx($1.Context)} 
				}
			}
			| ident MACRO param_list EOL block_statement ENDM
			{
				id := $1.(*Ident)
				$$ = &MacroStatement{NameID: id.NameID, Params: *$3, Body: $5.(*BlockStatement), End: int($6.Context.Line), Context: id.Context}
			}
			| IDENT expr_list 
			{
				e := $2.(*ExpressionList)
				$$ = &MacroCallStatement{NameID: $1.SymbolID, Args: e, Context: newCtxFromTokenCtx($1.Context)}
			}
			| ident_expr ':' IDENT expr_list
			{
				e := $4.(*ExpressionList)
				$$ = &MacroCallStatement{Label: $1, NameID: $3.SymbolID, Args: e, Context: newCtxFromTokenCtx($3.Context)}
			}
			| EXITM			{ $$ = &ExitmStatement{Context: newCtxFromTokenCtx($1.Context)}}
			| EXITM IF expr
			{
				if $3.NodeType() == NODE_ERROR {
					$$ = $3.(*ParseError)
				} else {
					ctx := newCtxFromTokenCtx($1.Context)
					$$ = &IfStatement{
						Condition: $3,
						Consequence: &BlockStatement{Block: []Statement{&ExitmStatement{Context: ctx}}},
						Alternative:  &BlockStatement{Block: []Statement{}},
						Context: newCtxFromTokenCtx($1.Context)}
				}
			}
			| datadef	{ $$ = $1 }
			| ident_expr datadef	
			{ 
				data := $2.(*DataDefineStatement)
				data.Label = $1
				$$ = data
			}
			| datastore	{ $$ = $1 }
			| ident_expr datastore	
			{ 
				data := $2.(*DataStoreStatement)
				data.Label = $1
				$$ = $2 
			}
			| ORG expr	{ $$ = &OrgStatement{Address: $2, AllocType: ALLOC_ABS, Context: newCtxFromTokenCtx($1.Context) }}
			| ORG expr ',' ident	
			{ 
				switch $4.(*Ident).NameID {
				case intern.InternString("ABS"):
					$$ = &OrgStatement{Address: $2, AllocType: ALLOC_ABS, Context: newCtxFromTokenCtx($1.Context) }
				case intern.InternString("REL"):
					$$ = &OrgStatement{Address: $2, AllocType: ALLOC_REL, Context: newCtxFromTokenCtx($1.Context) }
				default:
					$$ = &ParseError{Message: errcode.EORG_ALLOC, Context: newCtxFromTokenCtx($1.Context)}
				}
			}
			| INCLUDE STRING { $$ = &IncludeStatement{Filename: $2.SymbolID.String(), Context: newCtxFromTokenCtx($1.Context)} }
			| CHARMAP IDENT ',' expr
			{ 
				$$ = &CharmapStatement{NameID: $2.SymbolID, Filename: $4, Context: newCtxFromTokenCtx($1.Context)} 
			}
			| CHARMAP IDENT ',' expr ',' expr
			{ 
				$$ = &CharmapStatement{NameID: $2.SymbolID, Filename: $4, DefChar: $6, Context: newCtxFromTokenCtx($1.Context)} 
			}
			;


datadef		: DATA expr
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2.(*ParseError)
				} else {
					size := 0
					switch $1.TokenSubType {
					case DB:
						size = 1
					case DW:
						size = 2
					case DD:
						size = 0
					}
					$$ = &DataDefineStatement{Size: size, Values: []Expression{$2}, Context: newCtxFromTokenCtx($1.Context)}
				}
			}
			| datadef ',' expr
			{
				if $3.NodeType() == NODE_ERROR {
					$$ = $3.(*ParseError)
				} else {
					s := $1.(*DataDefineStatement)
					s.Values = append(s.Values, $3)
					$$ = s
				}
			}
			;

datastore	: DS expr
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2.(*ParseError)
				} else {
					var size int
					switch $1.TokenSubType {
					case DSB:
						size = 1
					case DSW:
						size = 2
					}
					$$ = &DataStoreStatement{Size: size, Count: $2, FillValue: nil, Context: newCtxFromTokenCtx($1.Context)}
				}
			}
			| DS expr ',' expr
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2.(*ParseError)
				} else if $4.NodeType() == NODE_ERROR {
					$$ = $4.(*ParseError)
				} else {
					size := 0
					switch $1.TokenSubType {
					case DSB:
						size = 1
					case DSW:
						size = 2
					}
					$$ = &DataStoreStatement{Size: size, Count: $2, FillValue: $4, Context: newCtxFromTokenCtx($1.Context)}
				}
			}
			;
	
ident		: IDENT 		{ $$ = &Ident{Name: $1.SymbolID.String(), NameID: $1.SymbolID, IdentType: IDENT, Context: newCtxFromTokenCtx($1.Context)} }
			| LOCAL_IDENT	{ $$ = &Ident{Name: $1.SymbolID.String(), NameID: $1.SymbolID, IdentType: LOCAL_IDENT, Context: newCtxFromTokenCtx($1.Context)}}
			| AT_IDENT		{ $$ = &Ident{Name: $1.SymbolID.String(), NameID: $1.SymbolID, IdentType: AT_IDENT, Context: newCtxFromTokenCtx($1.Context)}}
			| ANON_IDENT	{ $$ = &Ident{Name: $1.SymbolID.String(), NameID: $1.SymbolID, IdentType: ANON_IDENT, Context: newCtxFromTokenCtx($1.Context)}}
			;

ident_expr	: ident			{ $$ = $1 }
			| ident CONCAT expr	
			{ 
				if $3.NodeType() == NODE_ERROR {
					$$ = $3.(*ParseError)
				} else {
					$$ = buildInfixExpression(CONCAT, $1, $3, newCtxFromTokenCtx($2.Context)) 
				}
			}
			;

param_list	: 			{ $$ = &[]string{}}
			| IDENT		{ $$ = &[]string{($1.SymbolID.String())} }
			| param_list ',' IDENT
			{
				params := *$1
				params =append(params, $3.SymbolID.String())
				$$ = &params
			}
			;

elseifs		: { $$ = nil }
			| elseifs ELIF expr EOL block_statement 
			{ 
				ifst := &IfStatement{Condition: $3, Consequence: $5.(*BlockStatement), Alternative: &BlockStatement{Block:[]Statement{}}, Context: newCtxFromTokenCtx($2.Context)}
				if $3.NodeType() == NODE_ERROR {
					$$ = $3.(*ParseError)
				} else if $1 == nil {
					$$ = &BlockStatement{Block: []Statement{ifst}}
				} else if block := $1.(*BlockStatement); len(block.Block) == 0 {
					block.Block = append(block.Block, ifst)
					$$ = $1
				} else if len(block.Block) == 1 && block.Block[0].NodeType() == NODE_IF_STMT {
					stmt := block.Block[0].(*IfStatement)
					for {
						if stmt.Alternative == nil {
							stmt.Alternative = &BlockStatement{Block:[]Statement{ifst}}
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
							$$ = &ParseError{Message: fmt.Sprintf(errcode.EINTERNAL, "ELIF"), Context: newCtxFromTokenCtx($2.Context)}
							break
						}
					}
				}
			}
			;

	
// statement エラー検出時は yylex.Error() を呼んで伝播を止める
block_statement	: 	 				{ $$ = &BlockStatement{Block: []Statement{}} }
			| block_statement statement 
			{ 
				if $2 == nil { // error
					// do nothing
				} else if $2.NodeType() == NODE_ERROR {
					err := $2.(*ParseError)
					yylex.Error(err.Message, err.Context)
				} else {
					block := $1.(*BlockStatement)
					block.Block = append(block.Block, $2.(Statement))
					$$ = block
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
					yylex.Error(err.Message, err.Context)
				}
				e := $1.(*EnumElements)
				e.Elements = append(e.Elements, $2.(*EnumElement))
				$$ = e
			}
			;

enum_element : IDENT 			{ $$ = &EnumElement{NameID: $1.SymbolID, Value: nil, Context: newCtxFromTokenCtx($1.Context)} }
			| IDENT '=' expr	
			{ 
				if $3.NodeType() == NODE_ERROR {
					$$ = $3.(*ParseError)
				} else {
					$$ = &EnumElement{NameID: $1.SymbolID, Value: $3, Context: newCtxFromTokenCtx($1.Context) }
				}
			}
			;

instruction	: Z80_INST0
			{
				$$ = &Z80Instruction{InstType: Z80_INST0, Opcode: int($1.TokenSubType), Context: newCtxFromTokenCtx($1.Context)} 
			}
			| Z80_INST1
			{
				$$ = &Z80Instruction{InstType: Z80_INST1, Opcode: int($1.TokenSubType), Context: newCtxFromTokenCtx($1.Context)}
			}
			| Z80_INST1 operand
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2.(*ParseError)
				} else {
					$$ = &Z80Instruction{InstType: Z80_INST1, Opcode: int($1.TokenSubType), 
						Op1: $2, 
						Context: newCtxFromTokenCtx($1.Context) }
				}
			}
			| Z80_INST2 operand
			{
				if $2.NodeType() == NODE_ERROR {
					$$ = $2.(*ParseError)
				} else {
					$$ = &Z80Instruction{InstType: Z80_INST2, Opcode: int($1.TokenSubType), 
						Op2: $2,
						Context: newCtxFromTokenCtx($1.Context) }
				}
			}
			| Z80_INST2 operand ',' operand
			{
				if $2.NodeType() == NODE_ERROR && $4.NodeType() == NODE_ERROR {
					err := $2.(*ParseError)
					yylex.Error(err.Message, err.Context)
					$$ = $4.(*ParseError)
				} else if $2.NodeType() == NODE_ERROR {
					$$ = $2.(*ParseError)
				} else if $4.NodeType() == NODE_ERROR {
					$$ = $4.(*ParseError)
				} else {
					$$ = &Z80Instruction{InstType: Z80_INST2, Opcode: int($1.TokenSubType), 
							Op1: $2,
							Op2: $4,
							Context: newCtxFromTokenCtx($1.Context) }
				}
			}
			;
	

operand		: '(' Z80_REG16 ')'
			{ 
				reg := &RegisterLiteral{RegisterType: int($2.TokenType), Register:int($2.TokenSubType), Context: newCtxFromTokenCtx($2.Context)}
				$$ = &RegIndirectExpression{
						Register: reg,
						Context: newCtxFromTokenCtx($1.Context)}
			}
			| '(' Z80_REG16 ADDSUB expr ')' 
			{
				reg := &RegisterLiteral{RegisterType: int($2.TokenType), Register:int($2.TokenSubType), Context: newCtxFromTokenCtx($2.Context)}
				ctx := newCtxFromTokenCtx($1.Context)
				$$ = &RegIndirectExpression{
						Register: reg,
						Displacement: buildPrefixExpression(int($3.TokenSubType), $4, ctx),
						Context: ctx}
			}
			| '(' Z80_REG8 ')'
			{ 
				reg := &RegisterLiteral{RegisterType: int($2.TokenType), Register:int($2.TokenSubType), Context: newCtxFromTokenCtx($2.Context)}
				$$ = &RegIndirectExpression{
						Register: reg,
						Context: newCtxFromTokenCtx($1.Context)}
			}
			| '(' expr ')'		
			{ 
				$$ = &AddrIndirectExpression{Address: $2, Context: newCtxFromTokenCtx($1.Context)} 
			}
			| expr 				{ $$ = $1 }
			;
			

// expr エラー検出時は yylex.Error() を呼んで伝播を止める
expr_list	:   { $$ = &ExpressionList{Expressions: []Expression{}}}
			| expr		
			{ 
				if $1.NodeType() == NODE_ERROR {
					err := $1.(*ParseError)
					yylex.Error(err.Message, err.Context)
				}
				$$ = &ExpressionList{Expressions: []Expression{$1}}
			}
			| expr_list ',' expr
			{
				if $3.NodeType() == NODE_ERROR {
					err := $3.(*ParseError)
					yylex.Error(err.Message, err.Context)
				}
				e := $1.(*ExpressionList)
				e.Expressions = append(e.Expressions, $3)
				$$ = e
			}
			;

string		: STRING 	{ $$ = $1 }
			| string STRING
			{
				tok := $$
				tok.SymbolID = intern.InternString(tok.SymbolID.String() + $2.SymbolID.String())
				$$ = tok
			}
			;
	
expr		: NUMBER
			{
				s := $1.SymbolID.String()
				n, err := parseInt(s)
				if err == nil {
					$$ = &NumberLiteral{Value: int(n), Context: newCtxFromTokenCtx($1.Context)}
				} else {
					$$ = &ParseError{Message: fmt.Sprintf(errcode.ENUMBER_LITERAL, s), Context: newCtxFromTokenCtx($1.Context)}
				}
			}
			| string 		{ $$ = &StringLiteral{Value: $1.SymbolID.String(), Context: newCtxFromTokenCtx($1.Context)}}
			| Z80_REG8 		{ $$ = &RegisterLiteral{RegisterType: int($1.TokenType), Register:int($1.TokenSubType), Context: newCtxFromTokenCtx($1.Context)}}
			| Z80_REG16 	{ $$ = &RegisterLiteral{RegisterType: int($1.TokenType), Register:int($1.TokenSubType), Context: newCtxFromTokenCtx($1.Context)}}
			| Z80_FLAG		{ $$ = &FlagLiteral{Flag: int($1.TokenSubType), Context: newCtxFromTokenCtx($1.Context)}}
			| ident_expr	{ $$ = $1 }
			| DOT_IDENT
			{
				name := $1.SymbolID.String()
				names := strings.Split(name, ".")
				$$ = &DotIdent{Name: name, NameID: $1.SymbolID, Left: intern.InternString(names[0]), Right: intern.InternString("." + names[1]), Context: newCtxFromTokenCtx($1.Context)}
			}
			| IDENT '(' expr_list ')' 	
			{ 
				e := $3.(*ExpressionList)
				$$ = &FuncCallExpression{ Name: $1.SymbolID.String(), NameID: $1.SymbolID, Args: e, Context: newCtxFromTokenCtx($1.Context)} 
			}
			| '[' expr_list ']' 
			{ 
				$$ = &ArrayLiteral{Elements: $2.(*ExpressionList), Context: newCtxFromTokenCtx($1.Context)} 
			}
			| indexed_expr 				{ $$ = $1}
			| '(' expr ')'				{ $$ = $2}
			| expr ADDSUB expr			{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3, newCtxFromTokenCtx($2.Context)) }
			| expr MULDIV expr			{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3, newCtxFromTokenCtx($2.Context)) }
			| expr COMP expr 			{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3, newCtxFromTokenCtx($2.Context)) }
			| expr SHIFT expr			{ $$ = buildInfixExpression(int($2.TokenSubType), $1, $3, newCtxFromTokenCtx($2.Context)) }
			| expr OR expr				{ $$ = buildInfixExpression(OR, $1, $3, newCtxFromTokenCtx($2.Context)) }
			| expr AND expr				{ $$ = buildInfixExpression(AND, $1, $3, newCtxFromTokenCtx($2.Context)) }
			| ADDSUB expr %prec UNARY	{ $$ = buildPrefixExpression(int($1.TokenSubType), $2, newCtxFromTokenCtx($1.Context)) }
			| UNARY expr 				{ $$ = buildPrefixExpression(int($1.TokenSubType), $2, newCtxFromTokenCtx($1.Context)) }
			;

indexed_expr: expr '[' ']'
			{
				if $1.NodeType() == NODE_ERROR {
					err := $1.(*ParseError)
					yylex.Error(err.Message, err.Context)
					yylex.Error(errcode.EARRAY_NAME, newCtxFromTokenCtx($2.Context))
				} 
				$$ = &ParseError{Message: errcode.EARRAY_INDEX, Context: newCtxFromTokenCtx($2.Context)}
			}
			| expr '[' expr ']'
			{
				if $1.NodeType() == NODE_ERROR && $3.NodeType() == NODE_ERROR {
					yylex.Error(errcode.EARRAY_NAME, newCtxFromTokenCtx($2.Context))
					err := $1.(*ParseError)
					yylex.Error(err.Message, err.Context)

					yylex.Error(errcode.EARRAY_INDEX, newCtxFromTokenCtx($2.Context))
					$$ = $3
				} else if $1.NodeType() == NODE_ERROR {
					yylex.Error(errcode.EARRAY_NAME, newCtxFromTokenCtx($2.Context))
					$$ = $1
				} else if $3.NodeType() == NODE_ERROR {
					yylex.Error(errcode.EARRAY_INDEX, newCtxFromTokenCtx($2.Context))
					$$ = $3
				} else {
					$$ = &IndexedExpression{Left: $1, Index: $3, Context: newCtxFromTokenCtx($2.Context)}
				}
			}
			;

%%

var includeDirs []string

func Parse(l *Lexer, dirs []string) (*BlockStatement) {
	// 常に有効
	yyErrorVerbose = true

	// loadIncludeFile で参照
	includeDirs = dirs
	
	// error トークンでリカバリすると yyParse() は 0 を返すため、戻り値には意味がない
	yyParse(l)
	return l.program
}