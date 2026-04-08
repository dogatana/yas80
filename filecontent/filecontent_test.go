package filecontent

import (
	"bytes"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/internal/errcodenames"
)

func TestReadBlock(t *testing.T) {
	tests := []string{
		"utf8-bom.txt",
		"sjis.txt",
	}
	expected := bytes.ReplaceAll(readTestDataFile(t, "utf8.txt"), []byte{13, 10}, []byte{10})

	for _, filename := range tests {
		path := testGetDataFilePath(t, filename)
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
func TestReadBlockError(t *testing.T) {
	tests := []struct {
		filename string
		expected string
	}{
		{"binary.bin", errcode.EFILE_ENCODING},
		{"hoge", errcode.EFILE_NOT_FOUND},
	}

	for tn, tt := range tests {
		path := testGetDataFilePath(t, tt.filename)
		_, err := NewFromFile(path)
		if err == nil {
			t.Errorf("[%d] NewFromFile(%q) returns no error", tn, tt.filename)
			continue
		}

		ename, ok := errcodenames.ErrcodeNames[tt.expected]
		if !ok {
			t.Errorf("[%d] errcode の定義が見つからない %q", tn, tt.expected)
			continue
		}

		if !hasMessage(err.Error(), tt.expected) {
			t.Errorf("[%d] not [%s] \"%s\" => \"%s\"",
				tn,
				ename,
				tt.expected,
				err.Error(),
			)
		}
	}
}

func hasMessage(msg, err string) bool {

	re := regexp.MustCompile(`\.?%.\.?`)
	for _, s := range re.Split(err, -1) {
		if !strings.Contains(msg, s) {
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

func testGetDataFilePath(t *testing.T, filename string) string {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}
	return filepath.Join(filepath.Dir(file), "testdata", filename)
}
