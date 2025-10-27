package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"yas80/errorstore"
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

	es := errorstore.New()
	l := parser.NewLexer(bufio.NewReader(strings.NewReader(input)), "<string>", es)
	ret := parser.Parse(l)

	// error トークンを使うと 0 が返ってくる
	fmt.Printf("Parse returns %d\n", ret)

	fmt.Printf("parser.Parse() returned: %d\n", ret)
	ec, wc := es.Count()
	fmt.Printf("%d errors\n", ec)
	for _, e := range es.Errors {
		fmt.Println(e.String())
	}
	fmt.Printf("%d warnings\n", wc)
	for _, e := range es.Warnings {
		fmt.Println(e.String())
	}
}
