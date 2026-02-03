package testutil

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"yas80/logging"
)

func TestLogMessage(t *testing.T, tn int, err string, logger *logging.Logger) {
	ename := ErrcodeNames[err]

	var msgs []logging.LogMessage
	switch ename[0] {
	case 'E':
		msgs = logger.Errors
		if len(msgs) == 0 {
			t.Fatalf("[%d] no error", tn)
			return
		}
	case 'W':
		msgs = logger.Warnings
		if len(msgs) == 0 {
			t.Fatalf("[%d] no warning", tn)
			return
		}
	case 'I':
		msgs = logger.Infomation
		if len(msgs) == 0 {
			t.Fatalf("[%d] no information", tn)
			return
		}
	}

	if !hasMessage(msgs, err) {
		t.Errorf("[%d] not [%s] \"%s\" => \"%s\"",
			tn,
			ename,
			err,
			msgs[0])
	}
}

func hasMessage(messages []logging.LogMessage, expected string) bool {
	re := regexp.MustCompile(`\.?%.\.?`)
	ss := re.Split(expected, -1)

	for _, emsg := range messages {
		result := true
		for _, s := range ss {
			if !strings.Contains(emsg.Message(), s) {
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
