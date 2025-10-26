package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]

	parser.SetYYDebug(getYYDebugEnv())
	l := parser.NewLexer(bufio.NewReader(strings.NewReader(input)))
	ret := parser.Parse(l)

	// error トークンを使うと 0 が返ってくるので注意
	fmt.Printf("parser.Parse() returned: %d\n", ret)
	if ret != 0 {
		fmt.Printf("parse error: %d\n", ret)
	} else {
		fmt.Println("result: ", parser.Result)
		fmt.Printf("%d Statements\n", len(parser.Root.Statements))
		fmt.Println("----------")
		fmt.Println(parser.Root.String())
	}
}
