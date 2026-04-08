package evaluator

import (
	"testing"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/internal/testutil"
	"github.com/dogatana/yas80/logging"
	"github.com/dogatana/yas80/object"
)

func TestAssembleFile(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
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
		{input: "include-not-found", err: errcode.EFILE_NOT_FOUND},
		{input: "include-cyclic", err: errcode.EINCLUDE_CYCLIC},

		{input: "zilog"},   // 公開命令
		{input: "zilog-a"}, // SUB r 等を SUB A, r の揺らぎを許容
		{input: "ixy-hl"},  // IXH, IXL, IYH, IYL 非公開命令。R800 では公開命令

		{input: "anon-label"}, // 匿名シンボル
	}

	for tn, tt := range tests {
		path := testutil.GetTestFilePath(t, tt.input+".asm")
		env := object.NewEnvironment(nil)
		logger := logging.New()

		code, _ := evalFile(path, logger, env)
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, logger)
			continue
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
