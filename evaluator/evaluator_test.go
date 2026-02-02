package evaluator

import (
	"testing"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
)

func TestAssembleFile(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
	}{
		{input: "label-backward"},
		{input: "equ-backward"},
		{input: "label-forward"},
		{input: "equ-forward"},
		{input: "forward"},
		{input: "forward_symbol"},
		{input: "forward_mix"},
		{input: "macro"},
		{input: "var-macro", code: []byte{1, 0, 0x10, 2, 0, 0x20, 3, 0, 0x30}},

		{input: "include"},
	}

	for tn, tt := range tests {
		path := testutil.GetTestFilePath(t, tt.input+".asm")

		env := object.NewEnvironment(nil)
		logger := logging.New(path)

		code, ok := evalFile(logger, env)
		if !ok {
			t.Errorf("[%d] BinWrier.Write() failed. got %d", tn, len(code))
			return
		}

		expected := tt.code
		if expected == nil {
			expected = testutil.ReadTestDataFile(t, tt.input+".bin")
		}

		if err := testutil.BytesEqual(code, expected); err != nil {
			t.Errorf("[%d] generated code diff %s", tn, err.Error())
		}
	}
}
