package lister

import (
	"bufio"
	"fmt"
	"io"
	"yas80/filecontent"
	"yas80/object"
	"yas80/parser"
)

type Lister struct {
	Nodes   *parser.BlockStatement
	Objects *object.BlockObject
}

func New(pnode *parser.BlockStatement, pobj *object.BlockObject) *Lister {
	return &Lister{Nodes: pnode, Objects: pobj}
}

func (l *Lister) ProgramList(out io.Writer) {
	w := bufio.NewWriter(out)
	printStatement(w, l.Nodes)
	w.Flush()
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
	}
	if ctx.Offset == 0 {
		fmt.Fprint(w, "   ")
	} else {
		fmt.Fprint(w, " + ")
	}
}

func printContext(w io.Writer, ctx *filecontent.Context) {
	fmt.Fprintf(w, "%2d:%2d ", ctx.Line, ctx.Offset)
	if ctx.Source == nil {
		fmt.Fprint(w, "(  ) ")
	} else {
		fmt.Fprintf(w, "(%2d) ", ctx.Source.Line)
	}
}
