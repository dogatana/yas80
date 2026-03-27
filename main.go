package main

import (
	"fmt"
	"os"
	"runtime/pprof"
	"yas80/assembler"
	"yas80/filecontent"
	"yas80/options"
)

func main() {
	opt := options.Parse()

	run(opt)
	os.Exit(0)
}

// main をテスト可能とするため、別関数へ
func run(opt options.Option) {
	if opt.Cpuprofile != "" {
		f, err := os.Create(opt.Cpuprofile)
		if err != nil {
			fmt.Printf("cannot create %s\n", opt.Cpuprofile)
			os.Exit(1)
		}
		defer f.Close()
		if err := pprof.StartCPUProfile(f); err != nil {
			fmt.Println("cannot start CPU profile", err)
			os.Exit(1)
		}
		defer pprof.StopCPUProfile()
	}
	fcs := makeContents(opt)
	iter := makeContentIterFunc(fcs)

	as := assembler.New(opt)
	as.Run(iter)

	if opt.Memprofile != "" {
		f, err := os.Create(opt.Memprofile)
		if err != nil {
			fmt.Printf("cannot create %s\n", opt.Memprofile)
			os.Exit(1)
		}
		defer f.Close()
		if err := pprof.WriteHeapProfile(f); err != nil {
			fmt.Println("cannot write memory profile", err)
		}
	}
}

// opt の内容に応じた []*filecontent.FileContent を生成
func makeContents(opt options.Option) []*filecontent.FileContent {
	fcs := []*filecontent.FileContent{}

	switch {
	case opt.AsmArg: // コマンドライン引数をアセンブル
		fc, _ := filecontent.NewFromString("arg", opt.Args[0])
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
