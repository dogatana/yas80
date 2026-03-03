package main

import (
	"fmt"
	"path/filepath"
	"testing"
	"yas80/internal/testutil"
	"yas80/options"
)

func TestOutputBIN(t *testing.T) {
	// 入力ファイルを一時フォルダに移して実行
	outdir := t.TempDir()

	tests := []struct {
		in, out string
	}{
		{"zilog.asm", "zilog.bin"},
		{"zilog-a.asm", "zilog-a.bin"},
		{"ixy-hl.asm", "ixy-hl.bin"},
	}

	for _, tt := range tests {

		in := tt.in
		out := tt.out
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
}

func TestOutputMZT(t *testing.T) {
	// 入力ファイルを一時フォルダに移して実行
	outdir := t.TempDir()

	tests := []struct {
		in, out string
	}{
		{"name.asm", "name.mzt"},
		{"a_very_long_filename.asm", "a_very_long_filename.mzt"},
	}

	for _, tt := range tests {
		in := tt.in
		out := tt.out
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
}

func TestOutputT88(t *testing.T) {
	tests := []struct {
		in, out, ref string
	}{
		{"inc.asm", "incdat.t88", "dle.t88"},
	}

	for _, tt := range tests {
		in := tt.in
		out := tt.out
		ref := tt.ref

		result := filepath.Join("testdata", out)
		opt := options.Option{
			Args:   []string{filepath.Join("testdata", in)},
			OutT88: true,
			Output: result,
		}
		run(opt)

		expected := filepath.Join("testdata", ref)
		if err := testutil.FileEqual(result, expected); err != nil {
			fmt.Println(err)
			t.Error(err)
		}
	}
}
