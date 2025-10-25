package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"yas80/parser"
)

// メイン関数
func main() {
	if len(os.Args) != 2 {
		return
	}
	input := os.Args[1]

	l := parser.NewLexer(bufio.NewReader(strings.NewReader(input)))
	ret := parser.Parse(l)

	// error トークンを使うと 0 が返ってくるので注意
	fmt.Printf("parser.Parse() returned: %d\n", ret)
	if ret != 0 {
		fmt.Printf("parse error: %d\n", ret)
	} else {
		fmt.Println("result: ", parser.Result)
	}
}
