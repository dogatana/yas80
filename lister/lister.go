package lister

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"yas80/filecontent"
	"yas80/object"
	"yas80/parser"
)

// レイアウト  36
// [ 6] 行番号
// [ 2]
// [ 4] アドレス
// [ 2]
// [16] 8バイトデデータ
// [ 2]
// [ 2] サイクル
// [ 1]
// [ 1] 展開マーク(+)

// レイアウト 40
// [ 5] 行番号
// [ 2]
// [ 4] アドレス
// [ 1]
// [23] 8バイトデデータ
// [ 1]
// [ 2] サイクル
// [ 1]
// [ 1] 展開マーク(+)

// [5:行番号][2][4:addr][1][23:code][2][4:CZ80][2][3][1:+][1][srouce]  full
// [5:行番号][30                   ][11             ][1:+][1][source]  code1
// [5       ][30                   ][11             ][1:+][1][source]  code2
// [5:行番号][41                                    ][1:+][1][source]  value
// [5:行番号][41                                    ][1:+][1][comment]

// リスト出力の書式指定文字列
const (
	// 行
	fmtSrcOnly = "%5d%41s%c %s\n" // 行番号、展開マーク、ソース行

	fmtCode1 = "%5d  %04x %-25s[%2s]     %c "   // 行番号、コード（コード出力の最初の行は行番号あり）
	fmtCode2 = "       %04x %-25s[%2s]     %c " // 行番号なし、コード（コード出力2行目以降）
)

type Lister struct {
	nodes   *parser.BlockStatement
	objects []object.Object
	fcMap   map[string]*filecontent.FileContent
	fMap    object.FileMap
	fblocks []*object.FileBlock
}

func New(pnode *parser.BlockStatement, pobj *object.BlockObject, fcmap map[string]*filecontent.FileContent) *Lister {
	l := &Lister{nodes: pnode, fcMap: fcmap}
	l.objects = object.FlattenObject(pobj)

	// filemap の作成
	// l.fMap = object.BuildGroupMap(object.FlattenObject(pobj))
	// fmt.Println("-- FileMap start")
	// l.fMap.Print()
	// fmt.Println("-- FileMap end")

	// []*FileBlock の収集
	fblocks := object.BuildFileBlock(l.objects)
	fmt.Println("-- FileBlocks start")
	for _, fb := range fblocks {
		fb.Print()
	}
	fmt.Println("-- FileBlocks end")
	l.fblocks = fblocks

	return l
}

func (l *Lister) List(out io.Writer) {
	w := bufio.NewWriter(out)

	var lnum int // ソース行番号
	var fc *filecontent.FileContent

	for _, fb := range l.fblocks {
		// (1)ファイル名出力
		fmt.Fprintf(w, "%s:\n", fb.Filename)
		lnum = fb.Line
		fc = l.fcMap[fb.Filename]
		var src string
		var err error

		for _, ln := range fb.LineObjects.Keys() {
			// ln までソース表示
			for lnum < ln {
				src, err = fc.GetLine(lnum)
				if err != nil {
					panic(fmt.Sprintf("GetLine(%d)", lnum))
				}
				fmt.Fprintf(w, fmtSrcOnly, lnum, "", ' ', src)
				lnum++
			}
			src, err = fc.GetLine(lnum)
			objs, _ := fb.LineObjects.Get(ln)
			for _, obj := range objs {
				switch obj := obj.(type) {
				case *object.CommentObject:
					if obj.Context.Offset == 0 {
						fmt.Fprintf(w, fmtSrcOnly, lnum, "", ' ', obj.Text)
					} else {
						fmt.Fprintf(w, fmtSrcOnly, lnum, "", '+', obj.Text)
					}
				case *object.CodeObject:
					lines := l.codeToLines(obj)
					fmt.Fprintln(w, lines[0]+src)
					for _, line := range lines[1:] {
						fmt.Fprintln(w, line)
					}
				}
			}
			lnum++
		}
	}
	// ソースの残りを表示 TODO:コード生成ないので不要か？
	for lnum <= fc.LineCount() {
		src, err := fc.GetLine(lnum)
		if err != nil {
			panic(fmt.Sprintf("GetLine(%d/%d)", lnum, fc.LineCount()))
		}
		fmt.Fprintf(w, "%5d %30s%s\n", lnum, "", src)
		lnum++
	}
	w.Flush()
}

// CodeObject を []string へ変換
func (l *Lister) codeToLines(co *object.CodeObject) []string {
	lines := []string{}

	size := len(co.Code)
	addr := co.Addr
	cycle := "  "
	if co.CZ80 != 0 {
		cycle = fmt.Sprintf("%2d", co.CZ80)
	}
	exp := ' '
	if co.Context.Offset != 0 {
		exp = '+'
	}

	for i := 0; i < size; i += 8 {
		var j int
		var buf bytes.Buffer
		for j = 0; j < 8 && i+j < size; j++ {
			buf.WriteString(fmt.Sprintf("%02x ", co.Code[i+j]))
		}
		if i == 0 {
			lines = append(lines, fmt.Sprintf(fmtCode1, co.Context.Line, addr, buf.String(), cycle, exp))
		} else {
			lines = append(lines, fmt.Sprintf(fmtCode2, addr, buf.String(), cycle, exp))
		}
		addr += 8
	}
	return lines
}

func printStatement(w io.Writer, stmt parser.Statement) {
	// fmt.Fprintf(w, "%T\n", stmt)

	switch stmt := stmt.(type) {
	case *parser.MacroBlockStatement:
		printLineHead(w, stmt.Context)
		fmt.Fprintln(w, stmt.Name)

		for _, s := range stmt.Block {
			printStatement(w, s)
		}
	case *parser.BlockStatement:
		for _, s := range stmt.Block {
			printStatement(w, s)
		}
	default:
		ctx := stmt.GetContext()
		printLineHead(w, ctx)
		fmt.Fprintln(w, stmt.String())
	}
}

func printLineHead(w io.Writer, ctx *filecontent.Context) {
	printExpand(w, ctx)
	printContext(w, ctx)
}

func printExpand(w io.Writer, ctx *filecontent.Context) {
	if ctx == nil {
		fmt.Fprint(w, " ? ")
		return
	}
	if ctx.Offset == 0 {
		fmt.Fprint(w, "   ")
	} else {
		fmt.Fprint(w, " + ")
	}
}

func printContext(w io.Writer, ctx *filecontent.Context) {
	if ctx == nil {
		fmt.Fprint(w, "--:-- (  ) ")
		return
	}
	fmt.Fprintf(w, "%2d:%2d ", ctx.Line, ctx.Offset)
	if ctx.Source == nil {
		fmt.Fprint(w, "(  ) ")
	} else {
		fmt.Fprintf(w, "(%2d) ", ctx.Source.Line)
	}
}
