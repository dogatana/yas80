package evaluator

import (
	"fmt"
	"yas80/errcode"
	"yas80/logger"
	"yas80/object"
	"yas80/parser"
)

type Evaluator struct {
	logger   *logger.Logger
	Debug    int
	Resolved bool
	Counter  func() int
}

func New(logger *logger.Logger) *Evaluator {
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
	objects := []object.Object{}
	stmts := []parser.Node{}

	var obj object.Object

	// 一旦 0 に初期化し ORG 他で上書きする
	initLocationCounter(env, 0)

	for i := 0; i < len(prog.Statements); i++ {
		if e.Debug > 0 {
			fmt.Printf("eval prog.Statements[%d]\n", i)
			addr, _ := env.Get("$")
			fmt.Printf("$ %s\n", addr.String())
		}

		node := prog.Statements[i]

		switch stmt := node.(type) {
		// 命令
		case *parser.Z80Instruction:
			obj = e.evalStatement(stmt, env)
			objects = append(objects, obj)
			stmts = append(stmts, node)

		// ラベル
		case *parser.LabelStatement:
			if stmt.Name.LabelType != parser.NODE_LABEL {
				// LOCAL/AT の場合は AST から LabelStatement を削除する
				e.logger.Error(fmt.Sprintf(errcode.EGLOBAL_NOT_ALLOWED, stmt.Name.Name), stmt.Context)
				continue
			}
			obj = e.evalStatement(stmt, env)
			// ValueObject にラップして返す
			objects = append(objects, &object.ValueObject{Value: obj, Context: stmt.Context})
			stmts = append(stmts, node)

		// const/equ
		case *parser.ConstStatement:
			e.concatenateSymbol(&stmt.Name, env, stmt.Context)
			e.concatenateSymbol(&stmt.Value, env, stmt.Context)

			ident, ok := stmt.Name.(*parser.Ident)
			if !ok {
				// シンボル結合演算式の右辺値でエラーの場合
				return object.ERROR
			}

			if ident.IdentType != parser.IDENT {
				e.logger.Error(fmt.Sprintf(errcode.EGLOBAL_NOT_ALLOWED, ident.Name), stmt.Context)
				continue
			}
			obj := e.evalStatement(stmt, env)
			objects = append(objects, &object.ValueObject{Value: obj, Context: stmt.Context})
			stmts = append(stmts, node)

		// 代入
		case *parser.AssignStatement:
			obj := e.evalStatement(stmt, env)
			objects = append(objects, obj)
			stmts = append(stmts, node)

		// マクロ定義
		case *parser.MacroStatement:
			e.evalStatement(stmt, env)
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
			stmts = append(stmts, &parser.MacroBlockStatement{
				Name:  stmt.Name,
				Block: obj.(*object.NodesObject).Nodes})
			e.Resolved = false

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

		default:
			e.logger.Info(fmt.Sprintf(errcode.ENOT_IMPL_STMT, node), nil)
			obj = e.evalStatement(node, env)
			if obj == object.ERROR {
				continue
			}
			objects = append(objects, obj)
			stmts = append(stmts, node)
		}
	}

	prog.Statements = stmts
	return &object.ProgramObject{Objects: objects}
}
