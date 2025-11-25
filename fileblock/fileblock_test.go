package fileblock

import (
	"bytes"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestReadBlock(t *testing.T) {
	tests := []string{
		"utf8-bom.txt",
		"sjis.txt",
	}
	expected := bytes.ReplaceAll(readTestDataFile(t, "utf8.txt"), []byte{13, 10}, []byte{10})

	for _, filename := range tests {
		path := testDataFilePath(t, filename)
		fb, err := NewFromFile(path)
		if err != nil {
			t.Fatalf("NewFromFile(%q) returns %q", filename, err.Error())
		}
		if fb.Filename != path {
			t.Errorf("FileBlock.Filename not %s. got %s", filename, fb.Filename)
		}
		if !sameContent(fb.Content, expected) {
			t.Errorf("FileBlock(%s).Content is not expected. got %s", filename, fb.String())
		}
	}
}

func TestReadBlockError(t *testing.T) {
	tests := []struct {
		filename string
		expected string
	}{
		{"binary.bin", "unknown encoding"},
		{"hoge", "file not found"},
	}

	for _, tt := range tests {
		path := testDataFilePath(t, tt.filename)
		_, err := NewFromFile(path)
		if err == nil {
			t.Errorf("NewFromFile(%q) returns no error", tt.filename)
		}
		es := err.Error()
		if !strings.Contains(es, tt.expected) {
			t.Errorf("not %s. got %s", tt.expected, es)
		}
	}
}

func sameContent(b1, b2 []byte) bool {
	if len(b1) != len(b2) {
		return false
	}
	for i, b := range b1 {
		if b2[i] != b {
			return false
		}
	}
	return true
}

func readTestDataFile(t *testing.T, filename string) []byte {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}
	path := filepath.Join(filepath.Dir(file), "testdata", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%q) returns %s", path, err.Error())
	}
	return data
}

func testDataFilePath(t *testing.T, filename string) string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}
	return filepath.Join(filepath.Dir(file), "testdata", filename)
}
