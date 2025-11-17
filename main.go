package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"yas80/evaluator"
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
	l := parser.NewLexer(bufio.NewReader(input), filename, logger)
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
	fmt.Println("# parse")
	prog := parse(logger, input, file)
	logger.Print()
	fmt.Println("")

	// AST 表示
	fmt.Printf("# %d statements\n", len(prog.Statements))
	if len(prog.Statements) == 0 {
		os.Exit(0)
	}
	fmt.Println(prog.String())
	fmt.Println("")

	// pass1
	fmt.Println("-- evaluator")
	eval := evaluator.New(logger)
	eval.Debug = getDebugEnv("evaldebug")

	env := object.NewEnvironment(nil)

	result := eval.Eval(prog, env).(*object.ProgramObject)
	logger.Print()
	fmt.Println("-- after pass1")
	fmt.Println(prog.String())

	fmt.Println("--")
	fmt.Println(len(result.Objects), "objects")
	fmt.Println("--")
	for n, o := range result.Objects {
		if o == nil {
			fmt.Printf("%d: <nil>\n", n+1)
		} else {
			fmt.Printf("%d: %s\n", n+1, o.String())
		}
	}
	fmt.Println("-- env")
	env.Print()
	fmt.Println("-- env after EvalEnv")
	err := eval.EvalEnv(env)
	if err != nil {
		fmt.Println("Error during EvalEnv:", err)
	}
	env.Print()

	eval.Pass1 = true
	result = eval.Eval(prog, env).(*object.ProgramObject)
	logger.Print()
	fmt.Println("-- after pass2")
	fmt.Println(prog.String())
	fmt.Println("--")
	fmt.Println(len(result.Objects), "objects")
	fmt.Println("--")
	for n, o := range result.Objects {
		if o == nil {
			fmt.Printf("%d: <nil>\n", n+1)
		} else {
			fmt.Printf("%d: %s\n", n+1, o.String())
		}
	}

}
