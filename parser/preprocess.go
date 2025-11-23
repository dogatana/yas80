package parser

import (
	"fmt"
	"strings"
)

var macroTable map[string]*MacroStatement = map[string]*MacroStatement{}

func PreProrocess(prog *Program) *Program {
	var result []Node
	var stmt Node

	for i := 0; i < len(prog.Statements); i++ {
		stmt = prog.Statements[i]

		switch stmt := stmt.(type) {
		case *MacroStatement:
			uname := strings.ToUpper(stmt.Name)
			fmt.Println("register macro def", uname)
			macroTable[uname] = stmt
		case *MacroCallStatement:
			uname := strings.ToUpper(stmt.Name)
			fmt.Printf("macroDef[%s] = %v\n", uname, macroTable[uname])
			if _, ok := macroTable[uname]; !ok {
				// macro に登録がなければ、関数コール式文に変換
				fmt.Println("replace macro call to IDENT")
				expStmt := &ExpressionStatement{
					Value:      &CallExpression{Function: &Ident{Name: uname, IdentType: IDENT}, Arguments: stmt.Args, lineNumber: stmt.LineNumber()},
					lineNumber: stmt.LineNumber()}
				result = append(result, expStmt)
				continue
			}
		default:
			result = append(result, stmt)
		}
	}
	return &Program{Statements: result}
}
