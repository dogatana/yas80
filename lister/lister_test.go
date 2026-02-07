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
		// {input: "inc.asm", expected: "ds32.txt"},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()

		path := filepath.FromSlash("testdata/" + tt.input)
		list := evalFile(path, logger, env)
		if len(list) == 0 {
			t.Errorf("[%d] no list", tn)
			continue
		}
		expected, err := readListFile("testdata/" + tt.expected)
		if err != nil {
			t.Errorf("[%d] readListFile error %s", tn, err.Error())
		}
		if err := linesEqual(list, expected); err != nil {
			t.Errorf("[%d] %s", tn, err.Error())
		}

	}
}
