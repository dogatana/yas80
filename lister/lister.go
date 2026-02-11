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
	fmt.Println("-- FileBlocks start")
	fblocks := object.BuildFileBlock(l.objects)
	for _, fb := range fblocks {
		fb.Print()
	}
	fmt.Println("-- FileBlocks end")
	l.fblocks = fblocks

	return l
}

func (l *Lister) List(out io.Writer) {
	w := bufio.NewWriter(out)

	for _, fb := range l.fblocks {
		// (1)ファイル名出力
		fmt.Fprintf(w, "%s:\n", fb.Filename)
		lnum := fb.Line
		fc := l.fcMap[fb.Filename]
		var src string
		var err error

		for _, ln := range fb.LineObjects.Keys() {
			// ln までソース表示
			for lnum < ln {
				src, err = fc.GetLine(lnum)
				if err != nil {
					panic(fmt.Sprintf("GetLine(%d)", lnum))
				}
				fmt.Fprintf(w, "%5d %30s%s\n", lnum, "", src)
				lnum++
			}
			src, err = fc.GetLine(lnum)
			objs, _ := fb.LineObjects.Get(ln)
			for _, obj := range objs {
				switch obj := obj.(type) {
				case *object.CommentObject:
					fmt.Fprintf(w, "%5d %35s%s\n", lnum, "", src)
					// fmt.Fprintf(w, "%5d %30s%s\n", lnum, obj.Text, src)
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
	w.Flush()
}

// CodeObject を []string へ変換
func (l *Lister) codeToLines(co *object.CodeObject) []string {
	lines := []string{}

	size := len(co.Code)
	addr := co.Addr

	for i := 0; i < size; i += 8 {
		var j int
		var buf bytes.Buffer
		for j = 0; j < 8 && i+j < size; j++ {
			buf.WriteString(fmt.Sprintf("%02x ", co.Code[i+j]))
		}
		for j < 8 {
			buf.WriteString("   ")
			j++
		}
		if i == 0 {
			lines = append(lines, fmt.Sprintf("%5d  %04x %s %2d  ", co.Context.Line, addr, buf.String(), co.CZ80))
		} else {
			lines = append(lines, fmt.Sprintf("       %04x %s %2d  ", addr, buf.String(), co.CZ80))
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
