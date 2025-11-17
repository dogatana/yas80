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

	// 構文解析開始
	fmt.Println("-- parser")
	logger := logger.New(file)
	l := parser.NewLexer(bufio.NewReader(input), file, logger)
	prog := parser.Parse(l)
	fmt.Println("--")
	logger.Print()

	// 構文解析結果
	fmt.Printf("%d statements\n", len(prog.Statements))
	fmt.Println("--")
	if len(prog.Statements) == 0 {
		os.Exit(0)
	}
	fmt.Println(prog.String())

	// 評価開始
	fmt.Println("-- evaluator")
	eval := evaluator.New(logger)
	eval.Debug = getDebugEnv("evaldebug")

	env := object.NewEnvironment(nil)
	// eval.ResolveConst(prog, env)
	// fmt.Println("-- global env")
	// env.Print()

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
}
