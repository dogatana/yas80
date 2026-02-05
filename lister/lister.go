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
	line := 1

	w := bufio.NewWriter(out)

	for _, obj := range objs {
		var code *object.CodeObject
		if obj.Type() == object.OBJ_FILE {
			file := obj.(*object.FileObject).Filename
			// (1)ファイル名出力
			fmt.Printf("%s:\n", file)
			continue
		}

		code, ok := obj.(*object.CodeObject)
		if !ok || len(code.Code) == 0 {
			continue
		}

		ctx := code.Context

		if fc == nil {
			fc = ctx.FileContent
		}

		for line < ctx.Line {
			sline, err := fc.GetLine(line)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
			fmt.Printf("%6d: %30s\t%s\n", line, "", sline)
			line++
		}
		sline, err = fc.GetLine(line)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println(l.codeToLine(code) + sline)
		line++
	}
	w.Flush()

	for i := 148; i <= 151; i++ {
		s, _ := fc.GetLine(i)
		fmt.Printf("%d: %s\n", i, s)
	}
}

func (l *Lister) codeToLine(code *object.CodeObject) string {
	var buf bytes.Buffer
	for _, b := range code.Code {
		buf.WriteString(fmt.Sprintf("%02x", b))
	}
	return fmt.Sprintf("%6d  %04x  %-16s  [%2d]\t", code.Context.Line, code.Addr, string(buf.String()), code.CZ80)
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
