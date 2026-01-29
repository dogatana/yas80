package options

import "flag"

type Option struct {
	Lexdebug  bool
	Astdebug  int
	YYdebug   int
	Evaldebug int
	Line      string
	Args      []string
	R800      bool
}

func Parse() Option {
	var opt Option

	// debug flag
	flag.BoolVar(&opt.Lexdebug, "lexdebug", false, "lexer debug(bool)")
	flag.IntVar(&opt.YYdebug, "yydebug", 0, "YYDebug(int)")
	flag.IntVar(&opt.Astdebug, "astdebug", 0, "parser debug(int)")
	flag.IntVar(&opt.Evaldebug, "evaldebug", 0, "evaluator debug(int)")

	// -l string をアセンブル
	flag.StringVar(&opt.Line, "l", "", "assmble arg")

	// R800 CPU をターゲットにする
	flag.BoolVar(&opt.R800, "R800", false, "assemble for R800(bool)")
	flag.Parse()
	opt.Args = flag.Args()
	return opt
}
