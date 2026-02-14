package lister

import (
	"path/filepath"
	"testing"
	"yas80/logging"
	"yas80/object"
)

func TestLister(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{input: "zilog.asm", expected: "zilog.txt"},
		{input: "ds32.asm", expected: "ds32.txt"},
		{input: "inc.asm", expected: "inc.txt"},
		{input: "macro-simple.asm", expected: "macro-simple.txt"},
		{input: "macro-if.asm", expected: "macro-if.txt"},
		{input: "rept.asm", expected: "rept.txt"},
		{input: "rept-exitm.asm", expected: "rept-exitm.txt"},
		{input: "macro-rept.asm", expected: "macro-rept.txt"},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()

		path := filepath.FromSlash("testdata/" + tt.input)
		list, err := evalFile(path, logger, env)
		if err != nil {
			t.Errorf("[%d] %s", tn, err)
			continue
		}
		if ec := logger.ErrorCount(); ec != 0 {
			t.Errorf("[%d] %d errors ", tn, ec)
			continue
		}
		if len(list) == 0 {
			logger.Print()
			t.Errorf("[%d] no list", tn)
			continue
		}
		expected, err := readListFile("testdata/" + tt.expected)
		if err != nil {
			logger.Print()
			t.Errorf("[%d] readListFile error %s", tn, err.Error())
			continue
		}
		if err := linesEqual(list, expected); err != nil {
			logger.Print()
			t.Errorf("[%d] %s", tn, err.Error())
		}

	}
}
