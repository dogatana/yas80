package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"yas80/evaluator"
	"yas80/fileblock"
	"yas80/lister"
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

func getDebugEnv(name string) int {
	v := os.Getenv(name)
	if v == "" {
		return 0
	}
	num, err := strconv.Atoi(v)
	if err == nil {
		return num
	}
	return 0
}

func parse(logger *logging.Logger, input io.Reader, filename string) *parser.BlockStatement {
	// l := parser.NewLexer(bufio.NewReader(input), filename, logger)
	fb, err := fileblock.NewFromReader(filename, input)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	l := parser.NewLexer(fb, logger)
	return parser.Parse(l)
}

// メイン関数
func main() {
	var (
		file  string
		input io.Reader
	)

	switch len(os.Args) {
	case 1:
		file = "stdin"
		input = os.Stdin
	case 2:
		file = "arg"
		input = strings.NewReader(os.Args[1])
	case 3:
		fmt.Println("usage: main | main text")
		os.Exit(1)
	}

	input = strings.NewReader(`
	rept 3
	nop
	exitm if $i == 1
	ret
	endr
	`)
	// logger 作成
	logger := logging.New(file)

	// lexer debug
	if getDebugEnv("lexdebug") != 0 {
		fb, err := fileblock.NewFromReader(file, input)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		l := parser.NewLexer(fb, logger)
		for {
			tok := l.NextToken()
			fmt.Println(tok.String())
			if tok.TokenType == 0 {
				logger.Print()
				os.Exit(0)
			}
		}
	}

	parser.SetYYDebug(getDebugEnv("yydebug"))

	// 構文解析開始
	fmt.Println("# parse")
	prog := parse(logger, input, file)

	// 構文解析直後の AST 表示
	if getDebugEnv("astdebug") > 0 {
		logger.Print()
		fmt.Println("--")
		for i, s := range prog.Block {
			fmt.Printf("%d: %#v\n", i, s)
		}
		fmt.Println("--")
		fmt.Println(prog.String())
		if getDebugEnv("astdebug") == 1 {
			os.Exit(0)
		}
	}

	// プリプロセス
	fmt.Println("# preprocess")
	prog = parser.PreProrocess(logger, prog)

	// プリプロセス直後の AST 表示
	if getDebugEnv("astdebug") > 1 {
		fmt.Println("--")
		for i, s := range prog.Block {
			fmt.Printf("%d: %#v\n", i, s)
		}
		fmt.Println("--")
		fmt.Println(prog.String())
		os.Exit(0)
	}

	// AST 表示
	fmt.Println("# ast")
	if len(prog.Block) == 0 {
		fmt.Print("no statements detected")
	} else {
		fmt.Println(prog.String())
	}
	fmt.Println("")

	// env 作成
	env := object.NewEnvironment(nil)

	eval := evaluator.New(logger)
	eval.Debug = getDebugEnv("evaldebug")

	// eval 戦略
	// 評価後 eval.Resolved が true ならコード生成完了とみなす
	// true でないなら、規定回数（例: 256 とか 1,024) だけ eval を繰り返す
	var i int
	var obj object.Object
	for i = 0; i < 256; i++ {
		fmt.Printf("# [#%d] EvalProgrram\n", i)
		eval.Resolved = true
		obj = eval.EvalProgram(prog, env)
		logger.Print()
		object.PrintEnv(env)

		if obj == object.ERROR {
			fmt.Printf("*** evaluate program returns ERROR")
			os.Exit(1)
		}

		ec, _, _ := logger.Count()
		if ec != 0 {
			fmt.Println("*** Error")
			os.Exit(1)

		}
		showResult(i, prog, obj, env)
		// $ の評価ができないので、EvalEnvは実行しない
		// eval.EvalEnv(env)
		// eval.CheckSymbols(env)
		// 循環参照のエラーチェックは実施
		eval.CheckCyclicError(env)
		if len(logger.Errors) > 0 {
			fmt.Println("*** abort")
			logger.Print()
			os.Exit(1)
		}

		if eval.Resolved {
			break
		}
	}
	eval.CheckSymbolError(env)
	fmt.Printf("eval %d times, %d errors, eval.Resolved = %v\n", i, len(logger.Errors), eval.Resolved)
	if len(logger.Errors) > 0 || !eval.Resolved {
		fmt.Print("\n** error or  not resolved")
		logger.Print()
		os.Exit(1)
	}

	// eval 戦略
	// 仮コード生成によってラベルアドレスが本来のものと異なる場合があるため
	// コードが安定するまで規定回数、評価を繰り返す
	// 例) const abc = xyz + 10 \ ld a, abc \  xyz: nop
	fmt.Println("\n# finalize")
	code := evaluator.CollectCode(obj.(*object.BlockObject).Block)

	eval.CodeStable = false
	for i = 0; i < 256 && !eval.CodeStable; i++ {
		obj = eval.EvalProgram(prog, env)
		if len(logger.Errors) > 0 {
			break
		}
		newCode := evaluator.CollectCode(obj.(*object.BlockObject).Block)
		eval.CodeStable = bytes.Equal(code, newCode)
		if !eval.CodeStable {
			code = newCode
		}
	}

	fmt.Printf("finalize %d times, codeStable %v\n", i, eval.CodeStable)
	if len(logger.Errors) > 0 {
		logger.Print()
		fmt.Println(prog.String())
		object.PrintEnv(env)
		os.Exit(1)
	}

	lister := lister.New(prog, obj.(*object.BlockObject))
	lister.ProgramList(os.Stdout)

	printObjects(obj.(*object.BlockObject).Block)
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
			printContext(s.(parser.Statement).GetContext())
			fmt.Println(s.String())
		}
	default:
		printContext(stmt.GetContext())
		fmt.Println(stmt.String())
	}
}
func printContext(ctx *fileblock.Context) {
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
