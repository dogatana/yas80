package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"yas80/errorstore"
	"yas80/generator"
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

	es := errorstore.New()
	l := parser.NewLexer(bufio.NewReader(input), file, es)
	ec, wc := parser.Parse(l)
	es.Print()
	if ec != 0 || wc != 0 {
		os.Exit(1)
	}

	prog := &parser.Root
	fmt.Printf("%d statements\n", len(prog.Statements))
	if len(prog.Statements) == 0 {
		os.Exit(0)
	}
	fmt.Println(prog.String())

	g := generator.New(prog, es)
	g.Generate()
	ec, wc = es.Count()
	es.Print()
	if ec != 0 || wc != 0 {
		os.Exit(1)
	}

	fmt.Println("generated code")
	g.Dump()

}
