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
func (e *Evaluator) EvalProgram(prog *parser.Program, env object.Environment) object.Object {
	// 一旦 0 に初期化し ORG 他で上書きする
	initLocationCounter(env, 0)
	return e.evalBlockPtr(&prog.Statements, env)
}

// Program.Statements, ProcBlockStatement.Block 評価
func (e *Evaluator) evalBlockPtr(ptr *[]parser.Node, env object.Environment) object.Object {
	statements := *ptr
	objects := []object.Object{}
	stmts := []parser.Node{}

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
				return object.ERROR
			}
			nobj, ok := obj.(*object.NodeObject)
			if !ok {
				panic("not NodeObject")
			}
			node = nobj.Node // ProcBlockStatement
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
				return object.ERROR
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

		// const/equ
		case *parser.ConstStatement:
			obj := e.evalStatement(stmt, env)
			objects = append(objects, &object.ValueObject{Value: obj, Context: stmt.Context})
			// ValueObject の場合は再度評価のため文を残す（FuncObject 等は文を削除し再評価しない）
			if obj.Type() == object.VALUE_OBJ {
				stmts = append(stmts, node)
			}

		// 代入
		case *parser.AssignStatement:
			obj := e.evalStatement(stmt, env)
			objects = append(objects, obj)
			stmts = append(stmts, node)

		// マクロ定義
		case *parser.MacroStatement:
			obj := e.evalStatement(stmt, env)
			if isError(obj) {
				return object.ERROR
			}
			continue

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
				Block:   obj.(*object.NodesObject).Nodes,
				Context: stmt.Context}
			// fmt.Println("-- expanded")
			// for _, n := range bs.Block {
			// 	fmt.Println(n.String())
			// }
			// fmt.Println("-- expanded")

			node = bs
			goto EVAL_AGAIN
			// stmts = append(stmts, bs)
			// e.Resolved = false

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
				objects = append(objects, objs[:len(objs)-1]...)
			default:
				objects = append(objects, objs...)
			}

		// rept
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
			nodeObj, ok := obj.(*object.NodeObject)
			if !ok {
				panic("not NodesObject")
			}
			stmts = append(stmts, nodeObj.Node)
			e.Resolved = false

		// 関数定義
		case *parser.FuncStatement:
			obj := e.evalStatement(stmt, env)
			if isError(obj) {
				return object.ERROR
			}
			// stmts, objects 両方に追加しない

		// システム変数設定
		case *parser.SetSysVarStatement:
			obj := e.evalStatement(stmt, env)
			if isError(obj) {
				return object.ERROR
			}

		case *parser.EnumStatement:
			obj := e.evalStatement(stmt, env)
			if isError(obj) {
				return object.ERROR
			}
		default:
			// e.logger.Error(fmt.Sprintf(errcode.ENOT_IMPL_STMT, node), nil)
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
