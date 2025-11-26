package parser

import (
	"fmt"
	"strings"
	"yas80/errcode"
	"yas80/logger"
)

type MacroTableType map[string]*MacroStatement

var macroTable MacroTableType = make(MacroTableType)

func PreProrocess(log *logger.Logger, prog *Program) *Program {

	// マクロ定義抽出
	macroTable = make(map[string]*MacroStatement)
	extractMacroDef(log, prog)

	// yacc でマクロ呼出し、関数呼出しを完全に区別できないので、
	// マクロ定義 0 件でも MacroCall の書き換えは必要なのに注意
	result := preprocessStatements(log, prog.Statements)
	return &Program{Statements: result}
}

// トップレベルからマクロ定義を抽出する
func extractMacroDef(log *logger.Logger, prog *Program) {
	for i, stmt := range prog.Statements {
		macro, ok := stmt.(*MacroStatement)
		if !ok {
			continue
		}
		name := macro.Name
		if _, ok := macroTable[name]; ok {
			// マクロ定義済み
			log.Error(fmt.Sprintf(errcode.EMACRO_DUP, name), macro.lineNumber)
		}
		macroTable[name] = macro
		// 登録後はNode削除
		prog.Statements[i] = &DeletedStatement{Node: stmt}
	}
	names := []string{}
	for name := range macroTable {
		names = append(names, name)
	}
	fmt.Println("regsitered macro: ", strings.Join(names, ","))
}

// Program.Statements, BlockStatement.Block の書き換え
func preprocessStatements(log *logger.Logger, stmts []Node) []Node {
	var result []Node

	for i := 0; i < len(stmts); i++ {
		node := stmts[i]

		fmt.Printf("ppStmt %#v\n", node)
		switch stmt := node.(type) {
		case *MacroStatement:
			// マクロ定義はネスト不可
			log.Error(fmt.Sprintf(errcode.EMACRO_NEST), stmt.lineNumber)
		case *MacroCallStatement:
			name := stmt.Name
			_, ok := macroTable[name]
			if !ok {
				// name マクロが登録されていない
				// 関数 call も macro call に構文解析される
				expStmt := replaceMacroCall(stmt)
				result = append(result, expStmt)
				continue
			}
			result = append(result, stmt)
			// if len(stmt.Args.Expressions) != len(macroDef.Params) {
			// 	// 仮引数、引数の数のチェック
			// 	log.Error(fmt.Sprintf(errcode.EMACRO_ARGS, name), stmt.lineNumber)
			// 	return nil
			// }
			// // マクロ展開
			// stmts, err := expandMacroBody(log, stmt, macroDef)
			// if err != nil {
			// 	log.Error(err.Error(), stmt.lineNumber)
			// } else {
			// 	result = append(result, stmts...)
			// }
		// case *IfStatement:
		// 	conseq := preprocessStatements(log, stmt.Consequence.(*BlockStatement).Block)
		// 	var alt []Node

		// 	switch altStmt := stmt.Alternative.(type) {
		// 	case *BlockStatement:
		// 		alt = preprocessStatements(log, stmt.Alternative.(*BlockStatement).Block)
		// 	case *IfStatement:
		// 		alt = preprocessStatements(log, []Node{altStmt})
		// 	default:
		// 		panic(fmt.Sprintf("cannot process %T", stmt))
		// 	}
		// 	stmt = &IfStatement{
		// 		Condition:   stmt.Condition,
		// 		Consequence: &BlockStatement{Block: conseq},
		// 		Alternative: &BlockStatement{Block: alt},
		// 		lineNumber:  stmt.lineNumber}
		// 	result = append(result, stmt)

		default:
			result = append(result, stmt)
		}
	}
	return result
}

// macro call でないものを適切な Node に置き換える
func replaceMacroCall(stmt *MacroCallStatement) Statement {
	var expr Expression
	ln := stmt.lineNumber
	switch len(stmt.Args.Expressions) {
	case 0:
		// 引数が 0 なら Ident の式文とする
		expr = &Ident{Name: stmt.Name, lineNumber: ln}
	case 1:
		if len(stmt.Args.Expressions) == 1 {
			array, ok := stmt.Args.Expressions[0].(*ArrayLiteral)
			if ok && len(array.Elements.Expressions) == 1 {
				expr = &IndexedExpression{
					Left:       &Ident{Name: stmt.Name, lineNumber: ln},
					Index:      array.Elements.Expressions[0],
					lineNumber: ln}
				break
			}
		}
		fallthrough
	default:
		expr = &CallExpression{
			Function:   &Ident{Name: stmt.Name, IdentType: IDENT},
			Arguments:  stmt.Args,
			lineNumber: ln}
	}
	return &ExpressionStatement{Value: expr, lineNumber: ln}
}

// マクロ適用
func expandMacroBody(log *logger.Logger, call *MacroCallStatement, def *MacroStatement) ([]Node, error) {

	// 仮引数-実引数テーブル
	var paramTable map[string]Expression = map[string]Expression{}
	for i, param := range def.Params {
		paramTable[param] = call.Args.Expressions[i]
	}

	var result []Node
	for _, stmt := range def.Body.Block {
		if stmt == nil {
			continue
		}
		if stmt.NodeType() == NODE_MACRO_STMT {
			// マクロ定義はネスト不可
			log.Error(fmt.Sprintf(errcode.EMACRO_NEST), stmt.LineNumber())
			continue
		}
		node, _ := modifyNode(log, stmt, paramTable)
		result = append(result, node.(Statement))
	}
	return result, nil
}

// 仮引数を実引数で置き換え
func modifyNode(log *logger.Logger, node Node, paramTable map[string]Expression) (Node, bool) {
	switch node := node.(type) {
	case *LabelStatement:
		// TODO ラベル結合演算対応の際に要修正
		return node, false

	case *Z80Instruction:
		op1, mod1 := modifyNode(log, node.Op1, paramTable)
		op2, mod2 := modifyNode(log, node.Op2, paramTable)
		if !mod1 && !mod2 {
			return node, false
		}
		return &Z80Instruction{
			InstType: node.InstType, Opcode: node.Opcode,
			Op1: op1.(Expression), Op2: op2.(Expression),
			lineNumber: node.lineNumber}, true

	case *ConstStatement:
		v, mod := modifyNode(log, node.Value, paramTable)
		if !mod {
			return node, false
		}
		return &ConstStatement{Name: node.Name, Value: v.(Expression), lineNumber: node.lineNumber}, false

	case *InfixExpression:
		op1, mod1 := modifyNode(log, node.Op1, paramTable)
		op2, mod2 := modifyNode(log, node.Op2, paramTable)
		if !mod1 && !mod2 {
			return node, false
		}
		return &InfixExpression{Operator: node.Operator,
			Op1: op1.(Expression), Op2: op2.(Expression), lineNumber: node.lineNumber}, true
	case *PrefixExpression:
		op, mod := modifyNode(log, node.Op, paramTable)
		if !mod {
			return node, false
		}
		return &PrefixExpression{Operator: node.Operator, Op: op.(Expression), lineNumber: node.lineNumber}, true
	case *Ident:
		if node.IdentType != IDENT {
			return node, false
		}
		expr, ok := paramTable[node.Name]
		if !ok {
			return node, false
		}
		return expr, true
		// case *CallExpression:
		// 	fn, mod1 := modifyNode(node.Function, paramTable)
		// 	args, mod2 := modifyNode(node.Arguments, paramTable)
		// 	if !mod1 && !mod2 {
		// 		return node, false
		// 	}
		// 	return &CallExpression{Function: fn, Arguments: args, lineNumber: node.lineNumber}, true
	default:
		return node, false
	}
}
