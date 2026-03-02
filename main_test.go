package main

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"yas80/internal/testutil"
	"yas80/options"
)

func TestOutputBIN(t *testing.T) {
	// 入力ファイルを一時フォルダに移して実行
	outdir := t.TempDir()

	in := "zilog.asm"
	out := "zilog.bin"
	testsrc := filepath.Join(outdir, in)

	testutil.CopyFile(testsrc, filepath.Join("testdata", in))

	result := filepath.Join(outdir, out)
	opt := options.Option{
		Args:   []string{testsrc},
		OutBIN: true,
		Output: result,
	}
	run(opt)

	expected := filepath.Join("testdata", out)
	if err := testutil.FileEqual(result, expected); err != nil {
		fmt.Println(err)
		t.Error(err)
	}
}

func TestOutputMZT(t *testing.T) {
	// 入力ファイルを一時フォルダに移して実行
	outdir := t.TempDir()

	in := "zilog.asm"
	out := "zilog.mzt"
	testsrc := filepath.Join(outdir, in)
	testutil.CopyFile(testsrc, filepath.Join("testdata", in))

	result := filepath.Join(outdir, out)
	opt := options.Option{
		Args:   []string{filepath.Join(outdir, in)},
		OutMZT: true,
		Output: result,
	}
	run(opt)

	expected := filepath.Join("testdata", out)
	if err := testutil.FileEqual(result, expected); err != nil {
		fmt.Println(err)
		t.Error(err)
	}
}

func TestOutputMZTwithLongName(t *testing.T) {
	// 入力ファイルを一時フォルダに移して実行
	outdir := t.TempDir()

	in := "zilog.asm"
	out := "zilooooooooooooooooooooooooooooog.mzt"
	testsrc := filepath.Join(outdir, in)
	testutil.CopyFile(testsrc, filepath.Join("testdata", in))

	result := filepath.Join(outdir, out)
	opt := options.Option{
		Args:   []string{filepath.Join(outdir, in)},
		OutMZT: true,
		Output: result,
	}
	run(opt)

	name, _ := getMztNameBody(result)
	if name == "" {
		t.Error("no valid MZT")
	}
	if name != out[:16] {
		t.Errorf("name in MZT is not expected %q\ngot %q", out[:16], name)
	}
}

func getMztNameBody(filename string) (string, []byte) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", nil
	}

	ofs := 1
	for ofs < 0x12 && data[ofs] != 0x0d {
		ofs++
	}
	name := string(data[1:ofs])

	return name, data[0x80:]
}
