package main

import (
	"yas80/assembler"
	"yas80/options"
)

func main() {
	opt := options.Parse()
	// opt.Print()

	as := assembler.New(opt)
	as.Run()
}
