package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"yas80/errorstore"
	"yas80/evaluator"
	"yas80/object"
	"yas80/parser"
)

func getYYDebugEnv() int {
	v := os.Getenv("YYDEBUG")
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

	parser.SetYYDebug(getYYDebugEnv())

	fmt.Println("-- parser")
	es := errorstore.New()
	l := parser.NewLexer(bufio.NewReader(input), file, es)
	prog, ec, wc := parser.Parse(l)
	es.Print()
	if ec != 0 || wc != 0 {
		os.Exit(1)
	}

	fmt.Printf("%d statements\n", len(prog.Statements))
	if len(prog.Statements) == 0 {
		os.Exit(0)
	}
	fmt.Println(prog.String())

	es = errorstore.New()
	eval := evaluator.New(es)
	env := object.NewEnvironment(nil)

	eval.ResolveConst(prog, env)
	fmt.Println("-- global env")
	env.Print()

	fmt.Println("-- evaluator")
	result := eval.Eval(prog, env).(*object.Program)
	for _, o := range result.Objects {
		if o == nil {
			fmt.Println("<nil>")
		} else {
			fmt.Println(o.String())
		}
	}
	env.Print()
	es.Print()

}
