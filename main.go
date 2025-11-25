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
	l := parser.NewLexer(filename, fb, logger)
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

	// log 作成
	log := logger.New(file)

	// 構文解析開始
	prog := parse(log, input, file)

	// 構文解析直後の AST 表示
	if getDebugEnv("astdebug") > 0 {
		log.Print()
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
	prog = parser.PreProrocess(log, prog)
	log.Print()

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

	// pass1 実行
	fmt.Println("# pass1")
	eval := evaluator.New(log)
	eval.Debug = getDebugEnv("evaldebug")

	objects := eval.Eval(prog, env).(*object.ProgramObject)

	fmt.Printf("\n# %d objects\n", len(objects.Objects))
	for _, o := range objects.Objects {
		if o == nil {
			fmt.Println("<nil>")
		} else {
			fmt.Println(o.String())
		}
	}

	fmt.Println("\n# ast after pass1")
	fmt.Println(prog.String())

	fmt.Println("\n# env")
	env.Print()

	// eval env
	fmt.Println("\n# eval env")
	order, err := eval.EvalEnv(env)
	if err != nil {
		log.Error(err.Error(), 0)
	}
	fmt.Println("order:", order)

	fmt.Println("\n# env after eval")
	env.Print()

	if len(log.Errors) > 0 {
		fmt.Println("\n*** abort ***")
		log.Print()
		os.Exit(1)
	}

	fmt.Println("\n# pass2")
	eval.Pass1 = false
	objects = eval.Eval(prog, env).(*object.ProgramObject)
	log.Print()

	fmt.Println("\n# ast")
	fmt.Println(prog.String())

	fmt.Printf("\n# %d objects\n", len(objects.Objects))
	for _, o := range objects.Objects {
		if o == nil {
			fmt.Println("<nil>")
			continue
		}
		if o.Type() == object.CODE_OBJ {
			fmt.Print("code: ")
		}
		fmt.Println(o.String())
	}
	fmt.Println("\n# final env")
	env.Print()
}
