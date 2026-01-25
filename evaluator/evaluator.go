package evaluator

import (
	"fmt"
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

type Evaluator struct {
	logger     *logging.Logger
	Debug      int
	Resolved   bool
	CodeStable bool
	Counter    func() int
}

func New(logger *logging.Logger) *Evaluator {
	return &Evaluator{logger: logger, Resolved: true, Counter: makeCounter(0)}
}

// start + 1 から順次生成するカウンタ関数を返す
func makeCounter(start int) func() int {
	return func() int {
		start++
		return start
	}
}

// Program 評価
func (e *Evaluator) EvalProgram(prog *parser.Program, env TEnv) object.Object {
	// 一旦 0 に初期化し ORG 他で上書きする
	initLocationCounter(env, 0)
	// return e.evalBlockPtr(&prog.Statements, env)
	bs := &parser.BlockStatement{Block: prog.Statements}
	return e.evalStatementEx(bs, false, nil, env)
}

// Program.Statements, ProcBlockStatement.Block 評価
func (e *Evaluator) evalBlockPtr(ptr *[]parser.Statement, env TEnv) object.Object {
	statements := *ptr
	objects := []object.Object{}
	stmts := []parser.Statement{}

	var obj object.Object

	for i, node := range statements {

	EVAL_AGAIN:
		if e.Debug > 0 {
			addr, _ := env.Get("$")
			fmt.Printf("eval BlockPtr.satements[%d/%d] %T. $%s\n", i, len(statements), node, addr.String())
		}

		switch stmt := node.(type) {

		// PROC
		case *parser.ProcStatement:
			obj := e.evalStatement(stmt, env)
			if isError(obj) {
				continue
			}
			nobj, ok := obj.(*object.StatementObject)
			if !ok {
				panic("not NodeObject")
			}
			node = nobj.Statement // ProcBlockStatement
			goto EVAL_AGAIN

		// PROC BLOCK
		case *parser.ProcBlockStatement:
			pobj, ok := env.Get(stmt.Name)
			if !ok {
				panic(fmt.Sprintf("no ProcEnv(%s)", stmt.Name))
			}
			// ProcObject は Environment intterface を実装
			obj = e.evalBlockPtr(&stmt.Block, pobj.(*object.ProcObject))
			prog, ok := obj.(*object.ProgramObject)
			if !ok {
				continue
			}
			objects = append(objects, prog.Objects...)
			stmts = append(stmts, stmt)
			continue

		// Z80 命令
		case *parser.Z80Instruction:
			obj = e.evalStatement(stmt, env)
			if isError(obj) {
				continue
			}
			objects = append(objects, obj)
			stmts = append(stmts, node)

		// ラベル
		case *parser.LabelStatement:
			obj = e.evalStatement(stmt, env)
			if isError(obj) {
				continue
			}
			// ValueObject にラップして返す
			objects = append(objects, &object.ValueObject{Value: obj, Context: stmt.Context})
			stmts = append(stmts, node)

		// const/equ 定数定義
		case *parser.ConstStatement:
			obj := e.evalStatement(stmt, env)
			if isError(obj) {
				continue
			}
			objects = append(objects, &object.ValueObject{Value: obj, Context: stmt.Context})
			// 常に再評価するよう変更
			// if obj.Type() == object.VALUE_OBJ {
			stmts = append(stmts, node)
			// }

		// var 変数定義
		case *parser.VariableStatement:
			obj := e.evalStatement(stmt, env)
			if isError(obj) {
				continue
			}
			objects = append(objects, &object.ValueObject{Value: obj, Context: stmt.Context})
			// 変数は毎回評価する
			stmts = append(stmts, node)

		// 代入
		case *parser.AssignStatement:
			obj := e.evalStatement(stmt, env)
			if isError(obj) {
				continue
			}
			objects = append(objects, obj)
			stmts = append(stmts, node)

		// マクロ定義
		case *parser.MacroStatement:
			_ = e.evalStatement(stmt, env)

		// マクロ呼出し
		case *parser.MacroCallStatement:
			obj := e.evalStatement(stmt, env)
			if isError(obj) {
				return object.ERROR
			}
			if obj.Type() != object.NODES_OBJ {
				panic("not nodes object")
			}
			bs := &parser.MacroBlockStatement{
				Name:    stmt.Name,
				Block:   obj.(*object.StatemetnsObject).Statements,
				Context: stmt.Context}
			// fmt.Println("-- expanded")
			// for _, n := range bs.Block {
			// 	fmt.Println(n.String())
			// }
			// fmt.Println("-- expanded")

			node = bs
			goto EVAL_AGAIN

		// マクロブロック (展開済みマクロ)
		case *parser.MacroBlockStatement:
			stmts = append(stmts, stmt)
			obj := e.evalMacroBlockStatement(stmt, env)
			if isError(obj) || isRefNotFound(obj) {
				continue
			}
			bo, ok := obj.(*object.BlockObject)
			if !ok {
				panic("not block object")
			}
			objs := bo.Block
			switch {
			case len(objs) == 0:
				// do nothing
			case len(objs) == 1 && objs[0].Type() == object.EXITM_OBJ:
				// do nothing
			case len(objs) > 1 && objs[len(objs)-1].Type() == object.EXITM_OBJ:
				// exitm までの obj を append
				objects = append(objects, objs[:len(objs)-1]...)
			default:
				objects = append(objects, objs...)
			}

		// REPT
		case *parser.ReptStatement:
			obj := e.evalStatement(node, env)
			if isError(obj) {
				continue
			}
			if isRefNotFound(obj) {
				stmts = append(stmts, stmt)
				e.Resolved = false
				continue
			}
			nodeObj, ok := obj.(*object.StatementObject)
			if !ok {
				panic("not NodesObject")
			}
			// stmts = append(stmts, nodeObj.Node)
			// e.Resolved = false
			node = nodeObj.Statement
			goto EVAL_AGAIN

		// REPT 展開結果
		case *parser.BlockStatement:
			obj := e.evalStatement(stmt, env)
			if isError(obj) {
				continue
			}
			objects = append(objects, obj.(*object.BlockObject).Block...)
			stmts = append(stmts, node)

		// func
		case *parser.FuncStatement:
			_ = e.evalStatement(stmt, env)

		// enum
		case *parser.EnumStatement:
			_ = e.evalStatement(stmt, env)

		// システム変数設定
		case *parser.SetSysVarStatement:
			obj := e.evalStatement(stmt, env)
			if isError(obj) {
				continue
			}
			objects = append(objects, obj)
			stmts = append(stmts, node)

		default:
			obj = e.evalStatement(node, env)
			if isError(obj) {
				continue
			}
			objects = append(objects, obj)
			stmts = append(stmts, node)
		}
	}

	*ptr = stmts
	return &object.ProgramObject{Objects: objects}
}
