package lister

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
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
	nodes       *parser.BlockStatement
	objects     []object.Object
	fcMap       map[string]*filecontent.FileContent
	fcProcessed map[string]int
	fMap        object.FileMap
}

func New(pnode *parser.BlockStatement, pobj *object.BlockObject, fcmap map[string]*filecontent.FileContent) *Lister {
	l := &Lister{nodes: pnode, fcMap: fcmap}
	l.objects = object.FlattenObject(pobj)
	// l.fMap = object.BuildGroupMap(object.FlattenObject(pobj))

	// fmt.Println("-- FileMap start")
	// l.fMap.Print()
	// fmt.Println("-- FileMap end")

	// fcProcessed の初期化
	l.fcProcessed = map[string]int{}
	for f := range l.fcMap {
		l.fcProcessed[f] = 0
	}
	return l
}

func (l *Lister) List(out io.Writer) {
	// l.objects を iterate する関数
	objIter := func() func() object.Object {
		index := 0
		return func() object.Object {
			if index < len(l.objects) {
				v := l.objects[index]
				index++
				return v
			} else {
				return nil
			}
		}
	}
	iter := objIter()

	var (
		fc    *filecontent.FileContent
		sline string
		err   error
	)

	var lnum int // 行番号

	w := bufio.NewWriter(out)

	var obj object.Object

	state := 0

LOOP:
	for {
		switch state {
		case 0: // FileObject 取得
			for {
				if obj = iter(); obj == nil {
					state = 9
				}
				if o, ok := obj.(*object.FileObject); ok {
					fc, _ = l.fcMap[o.Filename]
					lnum = 1
					state++
					break
				}
			}
		case 1: // Object.取得＆解析
			if obj = iter(); obj == nil {
				state = 9
				break
			}
			switch obj := obj.(type) {
			case *object.CodeObject:
				// ctx.Line までソース表示
				for lnum < obj.Context.Line {
					sline, _ = fc.GetLine(lnum)
					fmt.Fprintf(w, "%5d %30s%s\n", lnum, "", sline)
					lnum++
				}
				// ソース取得
				sline, _ = fc.GetLine(lnum)
				// code 行取得
				lines := l.codeToLines(obj)
				fmt.Fprintln(w, lines[0]+sline)
				for _, line := range lines[1:] {
					fmt.Fprintln(w, line)
				}
			case *object.CommentObject:
				// ctx.Line までソース表示
				for lnum < obj.Context.Line {
					sline, _ = fc.GetLine(lnum)
					fmt.Fprintf(w, "%5d %30s%s\n", lnum, "", sline)
					lnum++
				}
				// ソース取得
				sline, _ = fc.GetLine(lnum)
				fmt.Fprintf(w, "%5d %30s%s\n", lnum, obj.Text, sline)

			case *object.FileObject:
				if fc != nil && obj.Included {
					// 処理中をファイルを push
					fmt.Printf("push %s\n", fc.Filename)
				}
				fc, _ = l.fcMap[obj.Filename]
				lnum = 1
				state = 1
			}

		case 9: // ソースの残りを表示
			for {
				sline, err = fc.GetLine(lnum)
				if err != nil {
					break LOOP
				}
				fmt.Fprintf(w, "%5d %30s%s\n", lnum, "", sline)
				lnum++
			}

		}
	}
	for _, obj := range l.objects {
		if obj.Type() == object.OBJ_FILE {
			file := obj.(*object.FileObject).Filename
			if fc != nil && lnum < fc.LineCount() {
				// リスト処理中なら処理済みの行番号を保存しておく
				l.fcProcessed[fc.Filename] = lnum
			}
			// (1)ファイル名出力
			fmt.Fprintf(w, "%s:\n", file)
			var ok bool
			fc, ok = l.fcMap[file]
			if !ok {
				panic(fmt.Sprintf("filecontent not found for %s", file))
			}
			if fc.Filename != file {
				panic(fmt.Sprintf("FILE %s, fc.Filename %s", file, fc.Filename))
			}
			lnum = l.fcProcessed[file] + 1
			// if lnum > fc.LineCount() {
			// lnum = 1
			// }
			continue
		}
		// fmt.Printf("fc %#v\n", fc)

		var co *object.CodeObject
		co, ok := obj.(*object.CodeObject)
		if !ok || len(co.Code) == 0 {
			continue
		}
		ctx := co.Context

		if fc == nil {
			fc = ctx.FileContent
		}

		// CodeObject の行までソース行のみリスト出力
		for lnum < ctx.Line {
			sline, err = fc.GetLine(lnum)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Fprintf(w, "%5d %30s%s\n", lnum, "", sline)
			lnum++
		}

		// ソース行取得
		sline, err = fc.GetLine(lnum)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// リスト出力
		lines := l.codeToLines(co)
		fmt.Fprintln(w, lines[0]+sline)
		for _, line := range lines[1:] {
			fmt.Fprintln(w, line)
		}
		lnum++
	}

	if fc == nil {
		fmt.Println("f nil")
		w.Flush()
		return
	}
	for lnum <= fc.LineCount() {
		sline, _ = fc.GetLine(lnum)
		fmt.Fprintf(w, "%5d %30s%s\n", lnum, "", sline)
		lnum++
	}
	w.Flush()
}

func (l *Lister) ProgramList(out io.Writer) {

	var (
		fc    *filecontent.FileContent
		sline string
		err   error
	)

	var lnum int // 行番号

	w := bufio.NewWriter(out)

	for _, obj := range l.objects {
		if obj.Type() == object.OBJ_FILE {
			file := obj.(*object.FileObject).Filename
			if fc != nil && lnum < fc.LineCount() {
				// リスト処理中なら処理済みの行番号を保存しておく
				l.fcProcessed[fc.Filename] = lnum
			}
			// (1)ファイル名出力
			fmt.Fprintf(w, "%s:\n", file)
			var ok bool
			fc, ok = l.fcMap[file]
			if !ok {
				panic(fmt.Sprintf("filecontent not found for %s", file))
			}
			if fc.Filename != file {
				panic(fmt.Sprintf("FILE %s, fc.Filename %s", file, fc.Filename))
			}
			lnum = l.fcProcessed[file] + 1
			// if lnum > fc.LineCount() {
			// lnum = 1
			// }
			continue
		}
		// fmt.Printf("fc %#v\n", fc)

		var co *object.CodeObject
		co, ok := obj.(*object.CodeObject)
		if !ok || len(co.Code) == 0 {
			continue
		}
		ctx := co.Context

		if fc == nil {
			fc = ctx.FileContent
		}

		// CodeObject の行までソース行のみリスト出力
		for lnum < ctx.Line {
			sline, err = fc.GetLine(lnum)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Fprintf(w, "%5d %30s%s\n", lnum, "", sline)
			lnum++
		}

		// ソース行取得
		sline, err = fc.GetLine(lnum)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// リスト出力
		lines := l.codeToLines(co)
		fmt.Fprintln(w, lines[0]+sline)
		for _, line := range lines[1:] {
			fmt.Fprintln(w, line)
		}
		lnum++
	}

	if fc == nil {
		fmt.Println("f nil")
		w.Flush()
		return
	}
	for lnum <= fc.LineCount() {
		sline, _ = fc.GetLine(lnum)
		fmt.Fprintf(w, "%5d %30s%s\n", lnum, "", sline)
		lnum++
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
