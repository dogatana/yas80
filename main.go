package main

import (
	"yas80/assembler"
	"yas80/options"
)

func main() {
	opt := options.Parse()

	as := assembler.New(opt)
	as.Assemble()
}
