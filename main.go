package main

import (
	"fmt"
	"os"
	"yas80/assembler"
	"yas80/filecontent"
	"yas80/options"
)

func main() {
	opt := options.Parse()

	if !opt.Stdin && len(opt.Args) == 0 {
		fmt.Println("no input files")
		os.Exit(1)
	}

	fcs := makeContents(opt)
	iter := makeContentIterFunc(fcs)

	as := assembler.New(opt)
	as.Run(iter)
	os.Exit(0)
}

// opt の内容に応じた []*filecontent.FileContent を生成
func makeContents(opt options.Option) []*filecontent.FileContent {
	fcs := []*filecontent.FileContent{}

	switch {
	case opt.AsmArg: // コマンドライン引数をアセンブル
		fc, _ := filecontent.NewFromString("line", opt.Args[0])
		fcs = append(fcs, fc)

	case len(opt.Args) == 0: // 標準入力をアセンブル
		fc, err := filecontent.NewFromReader("stdin", os.Stdin)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
		fcs = append(fcs, fc)

	case len(opt.Args) > 0: // コマンドライン引数を入力ファイルとしてアセンブル
		for _, arg := range opt.Args {
			fc, err := filecontent.NewFromFile(arg)
			if err != nil {
				fmt.Println(err.Error())
				os.Exit(1)
			}
			fcs = append(fcs, fc)
		}
	}
	return fcs
}

// FileContent ファイルコンテンツイテレータ関数を生成
func makeContentIterFunc(fcs []*filecontent.FileContent) func() *filecontent.FileContent {
	index := 0
	// closure を返す
	return func() *filecontent.FileContent {
		if index < len(fcs) {
			fc := fcs[index]
			index++
			return fc
		}
		return nil
	}
}
