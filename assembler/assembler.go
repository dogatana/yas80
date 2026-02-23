package assembler

import (
	"bytes"
	"fmt"
	"os"
	"yas80/binwriter"
	"yas80/evaluator"
	"yas80/filecontent"
	"yas80/lister"
	"yas80/logging"
	"yas80/object"
	"yas80/options"
	"yas80/parser"
)

type Assembler struct {
	options.Option
}

func New(option options.Option) *Assembler {
	return &Assembler{Option: option}
}

// アセンブル実行
func (as *Assembler) Run(fn func() *filecontent.FileContent) {
	// logger 作成
	logger := logging.New()

	// 構文解析
	prog, fcMap := as.parse(logger, fn)

	// エラーが 1件でもあれば終了
	if logger.ErrorCount() > 0 {
		logger.PrintSyntaxError()
		os.Exit(1)
	}

	// toplevel env 作成
	env := object.NewEnvironment(nil)

	// 環境初期化
	as.initEnvironment(env)

	eval := evaluator.New(logger)
	eval.Debug = as.Evaldebug

	// eval step 1
	// 評価後 eval.Resolved が true ならコード生成完了とみなす
	// true でないなら、規定回数（256) だけ eval を繰り返す
	var obj object.Object
	pass := 0 // 総評価回数

	for i := range 256 {
		pass++

		fmt.Printf("# [%d] EvalProgrram\n", i)
		eval.Resolved = true
		obj = eval.EvalProgram(prog, pass, env)
		if as.Evaldebug > 0 {
			logger.Print()
			object.PrintEnv(env)
		}

		if as.Evaldebug > 0 {
			showResult(i, prog, obj, env)
		}
		// $ の評価ができないので、EvalEnvは実行しない
		// eval.EvalEnv(env)
		// eval.CheckSymbols(env)
		// 循環参照のエラーチェックは実施
		eval.CheckCyclicError(env)
		if logger.ErrorCount() > 0 {
			break
		}
		if eval.Resolved {
			break
		}
	}
	// 未確定のシボルがないかどうかチェック
	eval.CheckSymbolError(env)

	// eval step 1 終了状況
	fmt.Printf("eval %d times, %d errors, eval.Resolved = %v\n", pass, logger.ErrorCount(), eval.Resolved)

	if logger.ErrorCount() == 0 && eval.Resolved {
		// eval step 2
		// 仮コード生成によってラベルアドレスが本来のものと異なる場合があるため
		// コードが安定するまで規定回数(16)評価を繰り返す
		// 例) const abc = xyz + 10 \ ld a, abc \  xyz: nop

		fmt.Println("\n# finalize")
		code := object.CollectCode(obj.(*object.BlockObject).Block)

		eval.CodeStable = false
		for i := 0; i < 16 && !eval.CodeStable; i++ {
			pass++
			env.Set("$PASS", &object.NumberObject{Value: pass})
			obj = eval.EvalProgram(prog, pass, env)
			if logger.ErrorCount() > 0 {
				break
			}
			newCode := object.CollectCode(obj.(*object.BlockObject).Block)
			eval.CodeStable = bytes.Equal(code, newCode)
			if !eval.CodeStable {
				code = newCode
			}
		}

		fmt.Printf("eval %d times, codeStable %v\n", pass, eval.CodeStable)
	}
	// 評価完了
	// if logger.ErrorCount() > 0 {
	// 	logger.Print()
	// 	fmt.Println(prog.String())
	// 	object.PrintEnv(env)
	// 	os.Exit(1)
	// }

	fmt.Println("-- ast")
	parser.PrintNode(prog, 0)

	if logger.ErrorCount() == 0 {
		// 出力ファイル生成
		fmt.Println("# 出力ファイル生成")
		fill, _ := getIntFromEnv(env, "$FILL")
		bw := binwriter.New(obj, fill, logger)

		var buf bytes.Buffer
		if bw.Write(&buf) {
			os.WriteFile("out.bin", buf.Bytes(), 0644)
		} else {
			logger.Print()
			os.Exit(1)
		}

		// マップファイル出力
		fmt.Println("# マップファイル出力")
		if err := bw.WriteMap(os.Stdout); err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}

	// リストファイル出力
	fmt.Println("-- log start")
	// LogMessage の重複削除
	logger.RemoveDupe()
	logger.Print()
	fmt.Println("-- log end")
	fmt.Println("-- message map start")
	mmap := logger.BuildMessageMap()
	mmap.Print()
	fmt.Println("-- message map end")

	fmt.Println("# リストファイル")
	for f := range fcMap {
		fmt.Printf("file %s\n", f)
	}
	lister := lister.New(as.R800, prog, obj.(*object.BlockObject), fcMap, mmap)
	lister.List(os.Stdout)
}

// 構文解析とプリプロセス
func (as *Assembler) parse(logger *logging.Logger, fn func() *filecontent.FileContent) (*parser.BlockStatement, map[string]*filecontent.FileContent) {
	// lexer 作成
	lexer := parser.NewLexer(logger, fn)

	// 構文解析
	fmt.Println("# parse")
	parser.SetYYDebug(as.YYdebug) // go-yacc debug flag
	prog := parser.Parse(lexer, as.IncDirs)

	// プリプロセス
	fmt.Println("# preprocess")
	prog = parser.PreProrocess(logger, prog)

	// 構文解析後 AST 表示
	fmt.Println("# ast")
	if as.Astdebug != 0 {
		if len(prog.Block) == 0 {
			fmt.Print("no statements detected")
		} else {
			fmt.Println(prog.String())
		}
		fmt.Println("")
	}
	return prog, lexer.FcMap
}

// 環境初期化
func (as *Assembler) initEnvironment(env object.Environment) {
	// システム変数初期値設定
	env.Set("$FILL", &object.NumberObject{Value: as.Fill})

	v := 0
	if as.R800 {
		v = 1
	}
	env.Set("$R800", &object.NumberObject{Value: v})

	// -D オプションで定義した定数を NumberObject として環境に登録
	for name, value := range as.Constants {
		env.Set(name, &object.NumberObject{Value: value})
	}
}

func printObjects(objs []object.Object) {
	for _, o := range objs {
		if o == nil {
			fmt.Println("<nil>")
			continue
		}
		if o == object.NULL {
			fmt.Println("NULL")
			continue
		}
		if bo, ok := o.(*object.BlockObject); ok {
			printObjects(bo.Block)
			continue
		}
		fmt.Println(o.String())
	}
}

func getIntFromEnv(env object.Environment, name string) (int, bool) {
	obj, ok := env.Get(name)
	if !ok {
		return 0, false
	}
	num, ok := obj.(*object.NumberObject)
	if !ok {
		return 0, false
	}
	return num.Value, true
}

func showResult(count int, prog *parser.BlockStatement, obj object.Object, env object.Environment) {
	path := ""
	if count >= 0 {
		path = fmt.Sprintf("[%d] ", count)
	}

	bo := obj.(*object.BlockObject)
	fmt.Printf("\n# %s%d objects\n", path, len(bo.Block))
	printObjects(bo.Block)

	fmt.Printf("\n# %sast\n", path)
	for _, node := range prog.Block {
		switch stmt := node.(type) {
		case *parser.MacroBlockStatement:
			printContext(stmt.GetContext())
			if stmt.Name == "REPT" {
				fmt.Printf("REPT %d\n", stmt.Count)
			} else {
				fmt.Println("MACRO", stmt.Name)
			}
			for _, s := range stmt.Block {
				printStatement(s)
			}
		default:
			printStatement(node)
		}
	}
	// fmt.Println(prog.String())

	fmt.Printf("\n# %senv\n", path)
	object.PrintEnv(env)

}

func printStatement(stmt parser.Statement) {
	switch stmt := stmt.(type) {
	case *parser.MacroBlockStatement:
		printContext(stmt.GetContext())
		if stmt.Name == "REPT" {
			fmt.Printf("REPT %d\n", stmt.Count)
		} else {
			fmt.Println("MACRO", stmt.Name)
		}
		for _, s := range stmt.Block {
			printContext(s.GetContext())
			fmt.Println(s.String())
		}
	default:
		printContext(stmt.GetContext())
		fmt.Println(stmt.String())
	}
}

func printContext(ctx *filecontent.Context) {
	if ctx == nil {
		fmt.Println("--:--")
		return
	}
	fmt.Printf("%2d:%2d ", ctx.Line, ctx.Offset)
	if ctx.Source == nil {
		fmt.Print("(  ) ")
	} else {
		fmt.Printf("(%2d) ", ctx.Source.Line)
	}
}
