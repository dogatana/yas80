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
	prog := parse(logger, input, file)
	logger.Print()

	// AST 表示
	fmt.Println("# ast")
	if len(prog.Statements) == 0 {
		os.Exit(0)
	}
	fmt.Println(prog.String())
	fmt.Println("")

	// env 作成
	env := object.NewEnvironment(nil)

	// pass1
	fmt.Println("# pass1")
	eval := evaluator.New(logger)
	eval.Debug = getDebugEnv("evaldebug")

	objects := eval.Eval(prog, env).(*object.ProgramObject)
	logger.Print()

	fmt.Printf("\n# %d objects\n", len(objects.Objects))
	for _, o := range objects.Objects {
		if o == nil {
			fmt.Println("<nil>")
		} else {
			fmt.Println(o.String())
		}
	}

	fmt.Println("\n# ast")
	fmt.Println(prog.String())

	fmt.Println("\n# env")
	env.Print()
	fmt.Println("\n# eval env")
	order, err := eval.EvalEnv(env)
	fmt.Println("order:", order)

	fmt.Println("\n# env after eval")
	if err != nil {
		fmt.Println("Error during EvalEnv:", err)
	}
	env.Print()

	fmt.Println("\n# pass2")
	eval.Pass1 = false
	objects = eval.Eval(prog, env).(*object.ProgramObject)
	logger.Print()

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
