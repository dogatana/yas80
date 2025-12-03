package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"yas80/evaluator"
	"yas80/fileblock"
	"yas80/logger"
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

func parse(logger *logger.Logger, input io.Reader, filename string) *parser.Program {
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

	parser.SetYYDebug(getDebugEnv("yydebug"))

	// logger 作成
	logger := logger.New(file)

	// 構文解析開始
	prog := parse(logger, input, file)

	// 構文解析直後の AST 表示
	if getDebugEnv("astdebug") > 0 {
		logger.Print()
		fmt.Println("--")
		for i, s := range prog.Statements {
			fmt.Printf("%d: %#v\n", i, s)
		}
		fmt.Println("--")
		fmt.Println(prog.String())
		if getDebugEnv("astdebug") == 1 {
			os.Exit(0)
		}
	}
	// プリプロセス
	fmt.Println("\n# preprocess")
	prog = parser.PreProrocess(logger, prog)
	logger.Print()

	// プリプロセス直後の AST 表示
	if getDebugEnv("astdebug") > 1 {
		fmt.Println("--")
		for i, s := range prog.Statements {
			fmt.Printf("%d: %#v\n", i, s)
		}
		fmt.Println("--")
		fmt.Println(prog.String())
		os.Exit(0)
	}
	// AST 表示
	fmt.Println("# ast")
	if len(prog.Statements) == 0 {
		fmt.Print("no statements detected")
	} else {
		fmt.Println(prog.String())
	}
	fmt.Println("")

	// env 作成
	env := object.NewEnvironment(nil)

	eval := evaluator.New(logger)
	eval.Debug = getDebugEnv("evaldebug")

	//
	// eval 戦略
	// 評価後 eval.Resolved が true ならコード生成完了とみなす
	// true でないなら、規定回数（例: 256 とか 1,024) だけ eval を繰り返す
	var i int
	for i = 0; i < 256; i++ {
		fmt.Printf("# eval [#%d]\n", i)
		eval.Resolved = true
		obj := eval.Eval(prog, env)
		logger.Print()

		if obj == object.ERROR {
			fmt.Printf("*** evaluate program returns ERROR")
			os.Exit(1)
		}

		ec, _, _ := logger.Count()
		if ec != 0 {
			fmt.Println("*** Error")
			os.Exit(1)

		}
		progObj := obj.(*object.ProgramObject)
		fmt.Printf("\n# %d objects\n", len(progObj.Objects))
		for _, o := range progObj.Objects {
			if o == nil {
				fmt.Println("<nil>")
			} else {
				fmt.Println(o.String())
			}
		}
		fmt.Println("\n# ast after pass1")
		fmt.Println(prog.String())

		fmt.Println("\n# env")
		object.PrintEnv(env)

		if eval.Resolved {
			break
		}
	}
	fmt.Printf("eval %d times, eval.Resolved = %v\n", i, eval.Resolved)
}
