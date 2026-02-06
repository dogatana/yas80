package lister

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"yas80/filecontent"
	"yas80/internal/util"
	"yas80/object"
	"yas80/parser"
)

type Lister struct {
	pnodes *parser.BlockStatement
	pobj   *object.BlockObject
}

func New(pnode *parser.BlockStatement, pobj *object.BlockObject) *Lister {
	return &Lister{pnodes: pnode, pobj: pobj}
}

func (l *Lister) ProgramList(out io.Writer) {
	objs := util.FlattenObject(l.pobj)

	var (
		fc    *filecontent.FileContent
		sline string
		err   error
	)
	lnum := 1

	w := bufio.NewWriter(out)

	for _, obj := range objs {
		if obj.Type() == object.OBJ_FILE {
			file := obj.(*object.FileObject).Filename
			// (1)ファイル名出力
			fmt.Printf("%s:\n", file)
			continue
		}

		var co *object.CodeObject
		co, ok := obj.(*object.CodeObject)
		if !ok || len(co.Code) == 0 {
			continue
		}
		ctx := co.Context

		if fc == nil {
			fc = ctx.FileContent
		}

		// CodeObject の行までソース行を進める
		for lnum < ctx.Line {
			sline, err = fc.GetLine(lnum)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("%6d  %30s\t%s\n", lnum, "", sline)
			lnum++
		}

		sline, err = fc.GetLine(lnum)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		lines := l.codeToLines(co)
		fmt.Println(lines[0] + sline)
		for _, line := range lines[1:] {
			fmt.Println(line)
		}
		lnum++
	}
	w.Flush()
}

func (l *Lister) codeToLines(co *object.CodeObject) []string {
	lines := []string{}

	size := len(co.Code)
	addr := co.Addr

	for i := 0; i < size; i += 8 {
		var j int
		var buf bytes.Buffer
		for j = 0; j < 8 && i+j < size; j++ {
			buf.WriteString(fmt.Sprintf("%02x", co.Code[i+j]))
		}
		for j < 8 {
			buf.WriteString("  ")
			j++
		}
		if i == 0 {
			lines = append(lines, fmt.Sprintf("%6d  %04x  %-16s  [%2d]\t", co.Context.Line, addr, buf.String(), co.CZ80))
		} else {
			lines = append(lines, fmt.Sprintf("        %04x  %-16s  [%2d]\t", addr, buf.String(), co.CZ80))
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
