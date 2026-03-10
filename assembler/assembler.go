package assembler

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"yas80/binwriter"
	"yas80/errcode"
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

	// 構文解析でエラーがあれば終了
	if logger.ErrorCount() > 0 {
		logger.PrintSyntaxError()
		os.Exit(1)
	}

	// toplevel env 作成
	env := object.NewEnvironment(nil)

	// 環境初期化
	as.initEnvironment(env)

	eval := evaluator.New(logger, as.IncDirs)
	eval.Debug = as.EvalDebug

	// 総評価回数 1...
	pass := 0
	// eval step 1
	obj, pass := as.evalStage1(eval, pass, logger, prog, env)

	// 未確定のシボルがないかどうかチェック
	eval.CheckSymbolError(env)

	// eval stage 1 終了状況
	fmt.Printf("eval %d times, %d errors, eval.Resolved = %v\n", pass, logger.ErrorCount(), eval.Resolved)

	if logger.ErrorCount() == 0 && eval.Resolved {
		obj, pass = as.evalStage2(eval, pass, logger, prog, env, obj)
		fmt.Printf("eval %d times, codeStable %v\n", pass, eval.CodeStable)
		if !eval.CodeStable {
			logger.Warning(errcode.WEVAL_CODE_STABLE, nil)
		}
	}

	// LogMessage の重複削除
	logger.RemoveDupe()

	if as.ListDebug > 0 {
		fmt.Println("-- ast")
		parser.PrintNode(prog, 0)
	}

	logger.Print()
	if logger.ErrorCount() == 0 {
		// 出力ファイル生成
		fmt.Println("# 出力ファイル生成")
		fill := as.Fill
		bw := binwriter.New(obj, fill, logger)

		var buf bytes.Buffer
		var ok bool
		switch {
		case as.OutBIN:
			ok = bw.WriteBin(&buf)
		case as.OutMZT:
			var name string
			if as.LoadName == "" {
				name = filepath.Base(as.Output)
			} else {
				name = as.LoadName
			}
			ok = bw.WriteMzt(&buf, name, as.EntryAddr)
		case as.OutT88:
			var name string
			if as.LoadName == "" {
				name = filepath.Base(as.Output)
			} else {
				name = as.LoadName
			}
			ok = bw.WriteT88(&buf, name)
		default:
			panic("no output format specified")
		}

		if !ok {
			logger.Print()
			os.Exit(1)
		}
		if err := os.WriteFile(as.Output, buf.Bytes(), 0644); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// マップファイル出力
		if as.MapFile != "" {
			fmt.Println("# マップファイル出力")
			buf.Reset()
			if err := bw.WriteMap(&buf); err != nil {
				fmt.Println(err.Error())
			} else if err := os.WriteFile(as.MapFile, buf.Bytes(), 0644); err != nil {
				fmt.Println(err)
			}
		}

		// シンボルファイル出力
		if as.SymFile != "" {
			fmt.Printf("# シンボルファイル出力")
			buf.Reset()
			if err := bw.WriteSym(&buf, env); err != nil {
				fmt.Println(err.Error())
			} else if err := os.WriteFile(as.SymFile, buf.Bytes(), 0644); err != nil {
				fmt.Println(err)
			}
		}
	}

	// リストファイル出力
	if as.ListFile == "" {
		return
	}

	fmt.Println("-- log start")
	logger.Print()
	fmt.Println("-- log end")
	mmap := logger.BuildMessageMap()
	if as.ListDebug > 0 {
		fmt.Println("-- message map start")
		mmap.Print()
		fmt.Println("-- message map end")

		for f := range fcMap {
			fmt.Printf("file %s\n", f)
		}
	}
	lister := lister.New(as.R800, prog, obj.(*object.BlockObject), fcMap, mmap)

	var buf bytes.Buffer
	lister.List(&buf)
	if err := os.WriteFile(as.ListFile, buf.Bytes(), 0644); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// 構文解析とプリプロセス
func (as *Assembler) parse(logger *logging.Logger, fn func() *filecontent.FileContent) (*parser.BlockStatement, map[string]*filecontent.FileContent) {
	// lexer 作成
	lexer := parser.NewLexer(logger, fn)

	// 構文解析
	fmt.Println("# parse")
	parser.SetYYDebug(as.YYdebug) // go-yacc debug flag
	prog := parser.Parse(lexer, as.IncDirs)

	if as.AutoProc {
		// AutoProc プリプロセス
		fmt.Println("# preprocess")
		if as.AstDebug == 2 {
			fmt.Println("before\n", prog.String())
		}
		prog = parser.PreProrocess(prog)
	}

	// 構文解析後 AST 表示
	fmt.Println("# ast")
	if as.AstDebug != 0 {
		if len(prog.Block) == 0 {
			fmt.Print("no statements detected")
		} else {
			fmt.Println(prog.String())
		}
		fmt.Println("")
	}
	return prog, lexer.FcMap
}

// eval step 1
// 評価後 eval.Resolved が true ならコード生成完了とみなす
// true でないなら、規定回数（256) だけ eval を繰り返す
func (as *Assembler) evalStage1(eval *evaluator.Evaluator, pass int, logger *logging.Logger, prog *parser.BlockStatement, env object.Environment) (object.Object, int) {
	fmt.Println("\n# eval stage 1")
	var obj object.Object

	for i := range 256 {
		pass++

		eval.Resolved = true
		obj = eval.EvalProgram(prog, pass, env)
		if as.EvalDebug > 0 {
			logger.Print()
			object.PrintEnv(env)
		}

		if as.EvalDebug > 0 {
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
	return obj, pass
}

// eval step 2
// 仮コード生成によってラベルアドレスが本来のものと異なる場合があるため
// コードが安定するまで規定回数(16)評価を繰り返す
// 例) const abc = xyz + 10 \ ld a, abc \  xyz: nop
func (as *Assembler) evalStage2(eval *evaluator.Evaluator, pass int, logger *logging.Logger, prog *parser.BlockStatement, env object.Environment, obj object.Object) (object.Object, int) {
	fmt.Println("\n# eval stage 2")
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
	return obj, pass
}

// 環境初期化
func (as *Assembler) initEnvironment(env object.Environment) {
	// システム変数初期値設定
	env.Set("$FILL", &object.NumberObject{Value: as.Fill})
	env.Set("$CMAP_ERR", &object.NumberObject{Value: -1})
	env.Set("$CMAP_THRU", &object.NumberObject{Value: -2})

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
