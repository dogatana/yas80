package parser

import (
	"fmt"
	"yas80/logger"
)

var macroTable map[string]*MacroStatement

func PreProrocess(log *logger.Logger, prog *Program) *Program {

	// マクロ定義抽出
	macroTable = make(map[string]*MacroStatement)
	extractMacroDef(log, prog)

	// マクロ定義 0 件でも MacroCall の書き換えは必要
	// if len(macroTable) == 0 {
	// 	return prog
	// }

	var result []Node
	for i := 0; i < len(prog.Statements); i++ {
		node := prog.Statements[i]

		switch stmt := node.(type) {
		case *MacroStatement:
			// マクロ定義はネスト不可
			log.Error(fmt.Sprintf(logger.EMACRO_NEST), stmt.lineNumber)
		case *MacroCallStatement:
			name := stmt.Name
			macroDef, ok := macroTable[name]
			if !ok {
				// 関数 call も macro call に構文解析される
				// マクロでなければ、関数コール式文に変換
				expStmt := &ExpressionStatement{
					Value: &CallExpression{
						Function: &Ident{Name: name, IdentType: IDENT}, Arguments: stmt.Args, lineNumber: stmt.LineNumber()},
					lineNumber: stmt.LineNumber()}
				result = append(result, expStmt)
				continue
			}
			if len(stmt.Args.Expressions) != len(macroDef.Params) {
				log.Error(fmt.Sprintf(logger.EMACRO_ARGS, name), stmt.lineNumber)
				return nil
			}
			// マクロ適用。複数文になりうる
			stmts, err := applyMacro(log, stmt, macroDef)
			if err != nil {
				log.Error(err.Error(), stmt.lineNumber)
			} else {
				result = append(result, stmts...)
			}
		default:
			result = append(result, stmt)
		}
	}
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
			log.Error(fmt.Sprintf(logger.EMACRO_DUP, name), macro.lineNumber)
		}
		macroTable[name] = macro
		// 登録後はNode削除
		prog.Statements[i] = &DeletedStatement{Node: stmt}
	}
}

// マクロ適用
func applyMacro(log *logger.Logger, call *MacroCallStatement, def *MacroStatement) ([]Node, error) {

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
			log.Error(fmt.Sprintf(logger.EMACRO_NEST), stmt.LineNumber())
			continue
		}
		fmt.Println("before", stmt.String())
		node, _ := modifyNode(stmt, paramTable)
		fmt.Println("after", node.String())
		result = append(result, node.(Statement))
	}
	return result, nil
}

// 仮引数を実引数で置き換え
func modifyNode(node Node, paramTable map[string]Expression) (Node, bool) {
	switch node := node.(type) {
	case *LabelStatement:
		// TODO ラベル結合演算対応の際に要修正
		return node, false
	case *Z80Instruction:
		op1, mod1 := modifyNode(node.Op1, paramTable)
		op2, mod2 := modifyNode(node.Op2, paramTable)
		if !mod1 && !mod2 {
			return node, false
		}
		return &Z80Instruction{
			InstType: node.InstType, Opcode: node.Opcode,
			Op1: op1.(Expression), Op2: op2.(Expression),
			lineNumber: node.lineNumber}, true
	case *ConstStatement:
		v, mod := modifyNode(node.Value, paramTable)
		if !mod {
			return node, false
		}
		return &ConstStatement{Name: node.Name, Value: v.(Expression), lineNumber: node.lineNumber}, false

	case *InfixExpression:
		op1, mod1 := modifyNode(node.Op1, paramTable)
		op2, mod2 := modifyNode(node.Op2, paramTable)
		if !mod1 && !mod2 {
			return node, false
		}
		return &InfixExpression{Operator: node.Operator,
			Op1: op1.(Expression), Op2: op2.(Expression), lineNumber: node.lineNumber}, true
	case *PrefixExpression:
		op, mod := modifyNode(node.Op, paramTable)
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
