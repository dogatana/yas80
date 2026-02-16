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
	option options.Option
}

func New(option options.Option) *Assembler {
	return &Assembler{option: option}
}

func (a *Assembler) Assemble() {
	fcs := []*filecontent.FileContent{}

	switch {
	case a.option.Line != "":
		fc, _ := filecontent.NewFromString("line", a.option.Line)
		fcs = append(fcs, fc)

	case len(a.option.Args) == 0:
		fc, err := filecontent.NewFromReader("stdin", os.Stdin)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fcs = append(fcs, fc)
	case len(a.option.Args) > 0:
		for _, arg := range a.option.Args {
			fc, err := filecontent.NewFromFile(arg)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fcs = append(fcs, fc)
		}
	}

	index := 0
	logger := logging.New()
	l := parser.NewLexer(logger, func() *filecontent.FileContent {
		if index < len(fcs) {
			fc := fcs[index]
			index++
			return fc
		}
		return nil
	})

	// go-yacc debug flag
	parser.SetYYDebug(a.option.YYdebug)

	// 構文解析開始
	fmt.Println("# parse")
	prog := parser.Parse(l)

	// プリプロセス
	fmt.Println("# preprocess")
	prog = parser.PreProrocess(logger, prog)

	// AST 表示
	fmt.Println("# ast")
	if a.option.Astdebug != 0 {
		if len(prog.Block) == 0 {
			fmt.Print("no statements detected")
		} else {
			fmt.Println(prog.String())
		}
		fmt.Println("")
	}

	// env 作成
	env := object.NewEnvironment(nil)

	// システム変数初期値上書き
	env.Set("$FILL", &object.NumberObject{Value: a.option.Fill})

	eval := evaluator.New(logger)
	eval.Debug = a.option.Evaldebug

	// eval 戦略
	// 評価後 eval.Resolved が true ならコード生成完了とみなす
	// true でないなら、規定回数（例: 256 とか 1,024) だけ eval を繰り返す
	var obj object.Object
	pass := 0 // 総評価回数
	for i := range 256 {
		pass++

		fmt.Printf("# [%d] EvalProgrram\n", i)
		eval.Resolved = true
		obj = eval.EvalProgram(prog, pass, env)
		if a.option.Evaldebug > 0 {
			logger.Print()
			object.PrintEnv(env)
		}

		if obj == object.ERROR {
			fmt.Printf("*** evaluate program returns ERROR")
			os.Exit(1)
		}

		ec, _, _ := logger.Count()
		if ec != 0 {
			logger.Print()
			fmt.Println("*** Error")
			os.Exit(1)

		}
		if a.option.Evaldebug > 0 {
			showResult(i, prog, obj, env)
		}
		// $ の評価ができないので、EvalEnvは実行しない
		// eval.EvalEnv(env)
		// eval.CheckSymbols(env)
		// 循環参照のエラーチェックは実施
		eval.CheckCyclicError(env)
		if logger.ErrorCount() > 0 {
			fmt.Println("*** abort")
			logger.Print()
			os.Exit(1)
		}

		if eval.Resolved {
			break
		}
	}
	eval.CheckSymbolError(env)
	fmt.Printf("eval %d times, %d errors, eval.Resolved = %v\n", pass, logger.ErrorCount(), eval.Resolved)
	if logger.ErrorCount() > 0 || !eval.Resolved {
		fmt.Print("\n** error or  not resolved")
		logger.Print()
		os.Exit(1)
	}

	// eval 戦略
	// 仮コード生成によってラベルアドレスが本来のものと異なる場合があるため
	// コードが安定するまで規定回数、評価を繰り返す
	// 例) const abc = xyz + 10 \ ld a, abc \  xyz: nop
	fmt.Println("\n# finalize")
	code := object.CollectCode(obj.(*object.BlockObject).Block)

	eval.CodeStable = false
	for i := 0; i < 256 && !eval.CodeStable; i++ {
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
	if logger.ErrorCount() > 0 {
		logger.Print()
		fmt.Println(prog.String())
		object.PrintEnv(env)
		os.Exit(1)
	}

	fmt.Println("-- ast")
	parser.PrintNode(prog, 0)

	// fmt.Println("-- objects")
	// printObjects(obj.(*object.BlockObject).Block)

	fmt.Println("-- binwriter")

	fill, _ := getIntFromEnv(env, "$FILL")

	bw := binwriter.New(obj, fill, logger)

	var buf bytes.Buffer
	if bw.Write(&buf) {
		os.WriteFile("out.bin", buf.Bytes(), 0644)
	}

	fmt.Println("-- log")
	// LogMessage の重複削除
	logger.RemoveDupe()
	logger.Print()

	fmt.Println("-- map")
	if err := bw.WriteMap(os.Stdout); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println("-- list")
	for f := range l.FcMap {
		fmt.Printf("file %s\n", f)
	}
	lister := lister.New(prog, obj.(*object.BlockObject), l.FcMap)
	lister.List(os.Stdout)
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
