package evaluator

import (
	"testing"
	"yas80/logger"
	"yas80/object"
)

func TestAssembleFile(t *testing.T) {
	tests := []string{
		"label-backward",
		"equ-backward",
		"label-forward",
		"equ-forward",
		"forward",
		"forward_symbol",
		"forward_mix",
	}

	for tn, base := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		input := string(readTestDataFile(t, base+".asm"))
		expected := readTestDataFile(t, base+".bin")

		prog, _ := evaluateInput(t, input, logger, env)
		logger.Print()
		result := CollectCode(prog)

		if err := bytesEqual(result, expected); err != nil {
			t.Errorf("[%d] generated code diff %s", tn, err.Error())
		}
	}
}
