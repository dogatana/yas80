package testutil

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"yas80/internal/errcodenames"
	"yas80/logging"
)

func TestLogMessage(t *testing.T, tn int, err string, logger *logging.Logger) {
	ename, ok := errcodenames.ErrcodeNames[err]
	if !ok {
		t.Fatalf("[%d] errcode の定義が見つからない %q", tn, err)
		return
	}

	ec, wc, ic := logger.Count()
	var mt logging.MessageType
	switch ename[0] {
	case 'E':
		mt = logging.Err
		if ec == 0 {
			t.Fatalf("[%d] no error", tn)
			return
		}
	case 'W':
		mt = logging.Warn
		if wc == 0 {
			t.Fatalf("[%d] no warning", tn)
			return
		}
	case 'I':
		mt = logging.Info
		if ic == 0 {
			t.Fatalf("[%d] no information", tn)
			return
		}
	}

	if !hasMessage(logger, mt, err) {
		t.Errorf("[%d] not [%s] \"%s\" => \"%s\"",
			tn,
			ename,
			err,
			logger.GetMessages()[0])
	}
}

func hasMessage(logger *logging.Logger, mt logging.MessageType, expected string) bool {
	re := regexp.MustCompile(`\.?%.\.?`)
	ss := re.Split(expected, -1)

	for _, m := range logger.GetMessages() {
		if m.Type != mt {
			continue
		}
		result := true
		for _, s := range ss {
			if !strings.Contains(m.Text, s) {
				result = false
				break
			}
		}
		if result {
			return result
		}
	}
	return false
}

func CopyFile(dst, src string) error {
	data, err := os.ReadFile(src)
	if err != err {
		return err
	}
	err = os.WriteFile(dst, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

// go.mod のあるディレクトリを返す
func ProjectRoot() string {
	dir, _ := os.Getwd()
	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			panic("go.mod not found")
		}
		dir = parent
	}
}

// testdata フォルダ内のファイル内容を返す
func ReadTestDataFile(t *testing.T, filename string) []byte {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		t.Fatal("runtime.Caller(1) failed")
	}
	path := filepath.Join(filepath.Dir(file), "testdata", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%q) returns %s", path, err.Error())
	}
	return data
}

// testdata フォルダ内のファイルパスを返す
func GetTestFilePath(t *testing.T, filename string) string {
	_, file, _, ok := runtime.Caller(1)
	if !ok {
		t.Fatal("runtime.Caller(1) failed")
	}
	return filepath.Join(filepath.Dir(file), "testdata", filename)
}

// ファイル比較
func FileEqual(f1, f2 string) error {
	b1, err := os.ReadFile(f1)
	if err != nil {
		return err
	}
	b2, err := os.ReadFile(f2)
	if err != nil {
		return err
	}
	return BytesEqual(b1, b2)
}

// []byte 比較
func BytesEqual(v1, v2 []byte) error {
	if len(v1) != len(v2) {
		return fmt.Errorf("size diff got 0x%x. expected 0x%x", len(v1), len(v2))
	}
	for i, v := range v1 {
		if v != v2[i] {
			return fmt.Errorf("contentis diff [0x%x] 0x%02x 0x%02x", i, v, v2[i])
		}
	}
	return nil
}
