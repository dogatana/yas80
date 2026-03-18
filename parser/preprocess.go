package parser

import (
	"fmt"
)

func PreProrocess(prog *BlockStatement) *BlockStatement {
	check := map[NodeType]bool{ // inProc 処理を終了する NodeType
		NODE_FILE:      true,
		NODE_ENUM_STMT: true,
		NODE_PROC_STMT: true,
		NODE_FUNC_STMT: true,
	}

	block := []Statement{}

	var pblock []Statement

	inProc := false   // auto-proc 処理中
	hasLocal := false // local label あり

	for _, stmt := range prog.Block {
		label := getStatementLabel(stmt)

		// label なし or auto-proc 終了文
		if !inProc && label == nil {
			block = append(block, stmt)
			continue
		}
		if !inProc && check[stmt.NodeType()] {
			block = append(block, stmt)
			continue
		}
		if inProc && check[stmt.NodeType()] {
			if !hasLocal { // local がないので元のまま
				block = append(block, pblock[1:]...)
			} else { // PROC 作成
				block = append(block, buildProc(pblock))
			}
			inProc = false
			hasLocal = false
			pblock = []Statement{}
			block = append(block, stmt)
			continue
		}
		if inProc && label == nil {
			pblock = append(pblock, stmt)
			continue
		}
		// label あり
		name := getName(label)
		if !inProc && name[0] != '.' {
			// proc + stmt
			pblock = []Statement{
				&ProcStatement{Name: label, Block: &BlockStatement{Block: []Statement{}}, Context: stmt.GetContext()},
				stmt}
			inProc = true
			continue
		}
		if !inProc {
			block = append(block, stmt)
			continue
		}
		if inProc && name[0] == '.' {
			hasLocal = true
			pblock = append(pblock, stmt)
			continue
		}
		if inProc {
			if !hasLocal {
				block = append(block, pblock[1:]...)
			} else {
				block = append(block, buildProc(pblock))
			}
			pblock = []Statement{
				&ProcStatement{Name: label, Block: &BlockStatement{Block: []Statement{}}, Context: stmt.GetContext()},
				stmt}
			hasLocal = false
			continue
		}
		pblock = append(pblock, stmt)
	}

	if len(pblock) != 0 {
		if !hasLocal {
			block = append(block, pblock[1:]...)
		} else {
			block = append(block, buildProc(pblock))
		}
	}
	prog.Block = block
	fmt.Printf("--prog\n%s\n", prog.String())
	return prog
}

func getStatementLabel(stmt Statement) Expression {
	switch stmt := stmt.(type) {
	case *LabelStatement:
		return stmt.Name
	case *Z80Instruction:
		return stmt.Label
	case *ConstStatement:
		return stmt.Name
	case *ProcStatement:
		return stmt.Name
	case *EnumStatement:
		return &StringLiteral{Value: stmt.Name}
	case *VariableStatement:
		return stmt.Name
	case *ReptStatement:
		return stmt.Label
	case *FuncStatement:
		return &StringLiteral{Value: stmt.Name}
	case *MacroStatement:
		return &StringLiteral{Value: stmt.Name}
	case *MacroCallStatement:
		return stmt.Label
	case *DataDefineStatement:
		return stmt.Label
	case *DataStoreStatement:
		return stmt.Label
	default:
		return nil
	}
}

func buildProc(pb []Statement) Statement {
	proc := pb[0].(*ProcStatement)
	if pb[1].NodeType() == NODE_LABEL_STMT {
		proc.Block.Block = pb[2:]
	} else {
		removeLabel(pb[1])
		proc.Block.Block = pb[1:]
	}
	return proc
}

func removeLabel(stmt Statement) {
	switch stmt := stmt.(type) {
	case *LabelStatement:
		stmt.Name = nil
	case *Z80Instruction:
		stmt.Label = nil
	case *ConstStatement:
		stmt.Name = nil
	// case *ProcStatement:
	// case *EnumStatement:
	case *VariableStatement:
		stmt.Name = nil
	case *ReptStatement:
		stmt.Label = nil
	// case *FuncStatement:
	// case *MacroStatement:
	case *MacroCallStatement:
		stmt.Label = nil
	case *DataDefineStatement:
		stmt.Label = nil
	case *DataStoreStatement:
		stmt.Label = nil
	}
}

// 名前を string として取得
func getName(expr Expression) string {
	switch expr := expr.(type) {
	case *StringLiteral:
		return expr.Value
	case *Ident:
		if expr.IdentType == IDENT || expr.IdentType == LOCAL_IDENT {
			return expr.Name
		}
		panic(fmt.Sprintf("unexpected Ident %s", expr.Name))
	case *InfixExpression:
		if expr.Operator == CONCAT {
			return getName(expr.Op1)
		}
		panic(fmt.Sprintf("unexpected InfixExpression %s", TokenLiteral(expr.Operator)))
	}
	panic(fmt.Sprintf("unexpected Expression %s", expr.String()))
}
