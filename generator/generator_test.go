package generator

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"yas80/errorstore"
	"yas80/parser"
)

func openTextFile(t *testing.T, name string) *os.File {
	path := filepath.Join("testdata", name)
	f, err := os.Open(path)
	if err != nil {
		t.Fatalf("cannot read testdata %s", err)
	}
	return f
}

func readBinaryData(t *testing.T, name string) []byte {
	path := filepath.Join("testdata", name)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("cannot read testdata %s", err)
	}
	return data
}

func TestZ80Inst0(t *testing.T) {
	input := openTextFile(t, "inst0.asm")
	defer input.Close()
	expected := readBinaryData(t, "inst0.bin")

	es := errorstore.New()
	l := parser.NewLexer(bufio.NewReader(input), "inst0.asm", es)
	root, ec, _ := parser.Parse(l)
	if ec != 0 {
		for _, e := range es.Errors {
			fmt.Println(e.String())
		}
		t.Fatalf("%d errors", ec)
	}

	g := New(root, es)
	g.Generate()

	result := g.MergeCode()

	if !bytes.Equal(g.MergeCode(), expected) {
		t.Errorf("invalid generated code. expected %d bytes. got %d bytes", len(expected), len(result))
	}
}
