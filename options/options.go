package options

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"yas80/internal/util"

	flag "github.com/spf13/pflag"
)

const (
	progName    = "yas80"
	progVersion = "0.1.0"
)

type Option struct {
	IncDirs   []string
	Constants []string
	AutoProc  bool
	Fill      int
	Output    string
	OutMZT    bool
	OutT88    bool
	OutBIN    bool
	List      string
	Map       string
	Sym       string
	Args      []string
	Version   bool
	// for debug
	Astdebug  int
	YYdebug   int
	Evaldebug int
	Listebug  int
	AsmArg    bool
}

func (opt Option) Print() {
	fmt.Printf("IncDirs: %v\n", opt.IncDirs)
	fmt.Printf("Constants: %v\n", opt.Constants)
	fmt.Printf("AutoProc: %v\n", opt.AutoProc)
	fmt.Printf("Fill: %d\n", opt.Fill)
	fmt.Printf("Output: %q\n", opt.Output)
	fmt.Printf("OutMZT: %v\n", opt.OutMZT)
	fmt.Printf("OutT88: %v\n", opt.OutT88)
	fmt.Printf("OutBIN: %v\n", opt.OutBIN)
	fmt.Printf("List: %q\n", opt.List)
	fmt.Printf("Map: %q\n", opt.Map)
	fmt.Printf("Sym: %q\n", opt.Sym)
	fmt.Printf("Args: %v\n", opt.Args)
	fmt.Printf("AstDebug: %d\n", opt.Astdebug)
	fmt.Printf("YYDebug: %d\n", opt.YYdebug)
	fmt.Printf("EvalDebug: %d\n", opt.Evaldebug)
	fmt.Printf("ListDebug: %d\n", opt.Listebug)
}

func Parse() Option {

	opt := Option{}

	flag.StringSliceVarP(&opt.IncDirs, "I", "I", []string{}, "directories to search for iclude")
	flag.StringSliceVarP(&opt.Constants, "D", "D", []string{}, "define constants")
	flag.StringVarP(&opt.Output, "output", "o", "", "output file name")
	flag.IntVarP(&opt.Fill, "fill", "f", 0xff, "filler for DS and Segment Gap")
	flag.BoolVar(&opt.OutMZT, "mzt", false, "output with MZT format")
	flag.BoolVar(&opt.OutT88, "t88", false, "output with T88 format")
	flag.StringVarP(&opt.List, "list", "l", "", "list file name")
	flag.Lookup("list").NoOptDefVal = "<<input>>.lst"
	flag.StringVarP(&opt.Map, "map", "m", "", "map file name")
	flag.Lookup("map").NoOptDefVal = "<<input>>.map"
	flag.StringVarP(&opt.Sym, "sym", "s", "", "symbol file name")
	flag.Lookup("sym").NoOptDefVal = "<<input>>.sym"
	flag.BoolVar(&opt.AutoProc, "auto-proc", false, "generate PROC from normal label")
	flag.BoolVarP(&opt.Version, "version", "v", false, "print version")
	// 以下デバッグ用非表示オプション
	flag.IntVar(&opt.Astdebug, "ast-debug", 0, "debug level for Parse")
	flag.IntVar(&opt.YYdebug, "yy-debug", 0, "debug level for go-yacc")
	flag.IntVar(&opt.Evaldebug, "eval-debug", 0, "debug level for Evaluator")
	flag.IntVar(&opt.Listebug, "list-debug", 0, "debug level for Lister")
	// flag.CommandLine.MarkHidden("ast-debug")
	// flag.CommandLine.MarkHidden("yy-debug")
	// flag.CommandLine.MarkHidden("eval-debug")
	// flag.CommandLine.MarkHidden("list-debug")
	flag.BoolVarP(&opt.AsmArg, "arg", "a", false, "assemble Args[0]")
	// usage をカスタマイズ
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [options] <input file>\n", progName)
		fmt.Fprintf(os.Stderr, "Options:\n")
		flag.PrintDefaults()
	}
	flag.CommandLine.SortFlags = false
	flag.Parse()
	opt.Args = flag.Args()

	if opt.Version {
		fmt.Printf("%s version %s\n", progName, progVersion)
		os.Exit(0)
	}

	// -I の値の整形(trim)
	opt.IncDirs = trimStringSlice(opt.IncDirs)

	return opt
}

// []string の各要素に対し trim 処理を実行
func trimStringSlice(strs []string) []string {
	if len(strs) == 0 {
		return strs
	}
	return util.Map(strs, func(s string) string { return strings.Trim(s, " \t") })
}

// 出力ファイル名の決定
func setOutput(opt *Option) {
	// opt.Output が指定されていない場合は、出力形式のフラグから出力ファイル名を決定する
	if opt.Output == "" {
		count := 0
		if opt.OutBIN {
			count++
		}
		if opt.OutMZT {
			count++
		}
		if opt.OutT88 {
			count++
		}

		switch count {
		case 0:
			// 出力形式のフラグが指定されていない場合は、OutBIN をデフォルトにする
			opt.OutBIN = true
		case 1:
			// break
		default:
			// 出力形式のフラグが複数指定されている場合はエラー
			fmt.Println("multiple output format is specified")
			os.Exit(1)
		}
		switch {
		case opt.OutBIN:
			opt.Output = replaceExt(opt.Args[0], ".bin")
		case opt.OutMZT:
			opt.Output = replaceExt(opt.Args[0], ".mzt")
		case opt.OutT88:
			opt.Output = replaceExt(opt.Args[0], ".t88")
		default:
			panic("unreachable")
		}
		return
	}

	// opt.Output が指定されている場合は、拡張子から出力形式を決定する
	ext := strings.ToLower(filepath.Ext(opt.Output))
	switch ext {
	case ".mzt":
		if opt.OutBIN || opt.OutT88 {
			fmt.Println("--bin or --t88 is specified")
			os.Exit(1)
		}
		opt.OutMZT = true
	case ".t88":
		if opt.OutMZT || opt.OutBIN {
			fmt.Println("--mzt or --bin is specified")
			os.Exit(1)
		}
		opt.OutT88 = true
	case ".bin":
		if opt.OutMZT || opt.OutT88 {
			fmt.Println("--mzt or --t88 is specified")
			os.Exit(1)
		}
		opt.OutBIN = true
	default:
		// 拡張子がない場合は、出力形式のフラグから出力形式を決定する
		count := 0
		if opt.OutBIN {
			count++
		}
		if opt.OutMZT {
			count++
		}
		if opt.OutT88 {
			count++
		}

		switch count {
		case 0:
			// 出力形式のフラグが指定されていない場合は、OutBIN をデフォルトにする
			opt.OutBIN = true
		case 1:
			// break
		default:
			// 出力形式のフラグが複数指定されている場合はエラー
			fmt.Println("multiple output format is specified")
			os.Exit(1)
		}
	}
}

// 拡張子を置き換える
func replaceExt(filename, newExt string) string {
	ext := filepath.Ext(filename)
	return filename[:len(filename)-len(ext)] + newExt
}
