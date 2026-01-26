package lister

import (
	"bufio"
	"fmt"
	"io"
	"yas80/fileblock"
	"yas80/object"
	"yas80/parser"
)

type Lister struct {
	Nodes   *parser.Program
	Objects *object.BlockObject
}

func New(pnode *parser.Program, pobj *object.BlockObject) *Lister {
	return &Lister{Nodes: pnode, Objects: pobj}
}

func (l *Lister) ProgramList(out io.Writer) {
	w := bufio.NewWriter(out)
	for _, s := range l.Nodes.Statements {
		printStatement(w, s)
	}
	w.Flush()
}

func printStatement(w io.Writer, stmt parser.Statement) {
	fmt.Printf("%+v\n", stmt)
	ctx := stmt.GetContext()
	if ctx == nil {
		return
	}
	if ctx.Offset == 0 {
		fmt.Fprint(w, "  ")
	} else {
		fmt.Fprint(w, "+ ")
	}
	printContext(w, ctx)
	switch stmt := stmt.(type) {
	case *parser.MacroBlockStatement:
		if stmt.Name == "REPT" {
			fmt.Fprintf(w, "REPT %d\n", stmt.Count)
		} else {
			fmt.Fprintln(w, "MACRO", stmt.Name)
		}
		for _, s := range stmt.Block {
			printStatement(w, s.(parser.Statement))
		}
	default:
		fmt.Fprintln(w, stmt.String())
	}
}

func printContext(w io.Writer, ctx *fileblock.Context) {
	fmt.Fprintf(w, "%2d:%2d ", ctx.Line, ctx.Offset)
	if ctx.Source == nil {
		fmt.Fprint(w, "(  ) ")
	} else {
		fmt.Fprintf(w, "(%2d) ", ctx.Source.Line)
	}
}
