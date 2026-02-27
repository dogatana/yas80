package lister

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"yas80/filecontent"
	"yas80/internal/util"
	"yas80/logging"
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
	fmtSrc = "%5d%41s%c %s\n" // 行番号、展開マーク、ソース行

	fmtText = "%5d       %s   %c %s\n" // 行番号、TextObject.Text、展開マーク、ソース行

	fmtCode0 = "%5d  %04x %-25s[  ]     %c %s\n" // len(Code) == 0 の場合
	fmtCodeF = "%5d  %04x %-25s[  ]     %c "     // DS 等 fill で埋められた場合で len(Code) > 8 の場合
	fmtCode1 = "%5d  %04x %-25s[%2s]     %c "    // 行番号、コード（コード出力の最初の行は行番号あり）
	fmtCode2 = "       %04x %-25s[%2s]     %c "  // 行番号なし、コード（コード出力2行目以降）
)

type Lister struct {
	tsIndex int
	nodes   *parser.BlockStatement
	objects []object.Object
	fcMap   map[string]*filecontent.FileContent
	fblocks []*object.FileBlock
}

func New(r800 bool, pnode *parser.BlockStatement, pobj *object.BlockObject, fcmap map[string]*filecontent.FileContent, mmap logging.MessageMap) *Lister {
	l := &Lister{nodes: pnode, fcMap: fcmap}
	if r800 {
		l.tsIndex = 1
	}
	l.objects = object.FlattenObject(pobj)

	fmt.Println("-- objects start")
	for _, o := range l.objects {
		fmt.Println(o.String())
	}
	fmt.Println("-- objects end")

	// []*FileBlock の収集
	fblocks := object.BuildFileBlock(l.objects)
	// MessageMap を FileBlock に挿入
	fblocks = object.InsertMessages(fblocks, mmap)
	l.fblocks = fblocks

	fmt.Println("-- FileBlocks start")
	for _, fb := range fblocks {
		fb.Print()
	}
	fmt.Println("-- FileBlocks end")

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
				fmt.Fprintf(w, fmtSrc, lnum, "", ' ', src)
				lnum++
			}
			// src, err = fc.GetLine(lnum)
			objs, _ := fb.LineObjects.Get(ln)
			for _, obj := range objs {
				switch obj := obj.(type) {
				case *object.CommentObject:
					text := obj.Text
					if text == nil {
						// Context の指すソース行を取得
						text = obj.Context.GetLine()
					}
					if obj.Context.Offset == 0 {
						fmt.Fprintf(w, fmtSrc, lnum, "", ' ', text)
					} else {
						fmt.Fprintf(w, fmtSrc, lnum, "", '+', text)
					}

				case *object.TextObject:
					src = obj.Context.GetLine()
					text := util.TruncateWithEllipsis(obj.Text, 31)
					if obj.Context.Offset == 0 {
						fmt.Fprintf(w, fmtText, lnum, text, ' ', src)
					} else {
						fmt.Fprintf(w, fmtText, lnum, text, '+', src)
					}

				case *object.ErrorObject:
					fmt.Fprintln(w, "    *  "+obj.Message)

				case *object.CodeObject:
					src = obj.Context.GetLine()
					if len(obj.Code) == 0 {
						// コードがない場合は addr, soruce ソース行だけ表示
						ext := ' '
						if obj.Context.Offset != 0 {
							ext = '+'
						}
						fmt.Fprintf(w, fmtCode0, obj.Context.Line, obj.Addr, "", ext, src)
						continue
					}
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
		fmt.Fprintf(w, fmtSrc, lnum, "", ' ', src)
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
	if co.TStates[l.tsIndex] != 0 {
		cycle = fmt.Sprintf("%2d", co.TStates[l.tsIndex])
	}
	exp := ' '
	if co.Context.Offset != 0 {
		exp = '+'
	}

	// DS 等で確保された領域が 8 バイト超なら省略形で表示
	if co.Filled && size > 8 {
		var data string
		if co.Code[0] == co.Code[1] {
			data = fmt.Sprintf("%02x ... (%04x bytes)", co.Code[0], len(co.Code))
		} else {
			data = fmt.Sprintf("%02x %02x ... (%04x bytes)", co.Code[0], co.Code[1], len(co.Code))
		}
		return []string{fmt.Sprintf(fmtCodeF, co.Context.Line, addr, data, exp)}
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

	// INCBIN のデータで3行以上なら省略形で表示
	if co.IncBin && len(lines) <= 2 {
		return lines
	} else if co.IncBin {
		return []string{
			lines[0],
			fmt.Sprintf("%12s... (%04x bytes total)", "", co.Size()),
			lines[len(lines)-1]}
	}
	return lines
}
