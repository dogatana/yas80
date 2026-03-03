package parser

import (
	"fmt"
)

func PreProrocess(prog *BlockStatement) *BlockStatement {

	block := []Statement{}

	var proc *ProcStatement
	var label Expression

	in_proc := false
	for _, stmt := range prog.Block {

		switch stmt := stmt.(type) {
		case *LabelStatement:
			label = stmt.Name
		case *Z80Instruction:
			label = stmt.Label
		case *ConstStatement:
			label = stmt.Name
		case *ProcStatement:
			label = stmt.Name
		case *EnumStatement:
			label = &StringLiteral{Value: stmt.Name}
		case *VariableStatement:
			label = stmt.Name
		case *ReptStatement:
			label = stmt.Label
		case *FuncStatement:
			label = &StringLiteral{Value: stmt.Name}
		case *MacroStatement:
			label = &StringLiteral{Value: stmt.Name}
		case *MacroCallStatement:
			label = stmt.Label
		case *DataStatement:
			label = stmt.Label
		case *DataStoreStatement:
			label = stmt.Label
		default:
			label = nil
		}
		if !in_proc && label == nil {
			block = append(block, stmt)
			continue
		}
		if !in_proc && (stmt.NodeType() == NODE_ENUM_STMT || stmt.NodeType() == NODE_PROC_STMT || stmt.NodeType() == NODE_FUNC_STMT) {
			block = append(block, stmt)
			continue
		}
		if in_proc && label == nil {
			proc.Block.Block = append(proc.Block.Block, stmt)
			continue
		}
		if in_proc && (stmt.NodeType() == NODE_ENUM_STMT || stmt.NodeType() == NODE_PROC_STMT || stmt.NodeType() == NODE_FUNC_STMT) {
			block = append(block, proc, stmt)
			proc = nil
			in_proc = false
			continue
		}
		name := getName(label)
		if !in_proc && name[0] != '.' {
			// create PROC
			if stmt.NodeType() == NODE_LABEL_STMT {
				proc = &ProcStatement{Name: label, Block: &BlockStatement{Block: []Statement{}}, Context: stmt.GetContext()}
			} else {
				removeLabel(stmt)
				proc = &ProcStatement{Name: label, Block: &BlockStatement{Block: []Statement{stmt}}, Context: stmt.GetContext()}
			}
			in_proc = true
			continue
		}
		if in_proc && name[0] == '.' {
			proc.Block.Block = append(proc.Block.Block, stmt)
			continue
		}
		if in_proc && name[0] != '.' {
			block = append(block, proc)
			if stmt.NodeType() == NODE_LABEL_STMT {
				proc = &ProcStatement{Name: label, Block: &BlockStatement{Block: []Statement{}}, Context: stmt.GetContext()}
			} else {
				removeLabel(stmt)
				proc = &ProcStatement{Name: label, Block: &BlockStatement{Block: []Statement{stmt}}, Context: stmt.GetContext()}
			}
			continue
		}
		proc.Block.Block = append(proc.Block.Block, stmt)

	}
	if proc != nil {
		block = append(block, proc)
	}
	prog.Block = block
	fmt.Println("--prog\n", prog.String())
	return prog
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
	case *DataStatement:
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
