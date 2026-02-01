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
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		input := string(testutil.ReadTestDataFile(t, tt.input+".asm"))

		code := tt.code
		if code == nil {
			code = testutil.ReadTestDataFile(t, tt.input+".bin")
		}

		prog, e := evalInput(input, logger, env)
		testEvalResult(t, tn, "", e)

		logger.Print()
		result := CollectCode(prog.Block)

		if err := testutil.BytesEqual(result, code); err != nil {
			t.Errorf("[%d] generated code diff %s", tn, err.Error())
		}
	}
}
