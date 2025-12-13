package evaluator

import (
	"testing"
	"yas80/logger"
	"yas80/object"
)

func TestZ80Instruction(t *testing.T) {
	tests := []string{
		"inst0",
		"ld_r8_r8",
		"label-backward",
		"equ-backward",
		"label-forward",
		"equ-forward",
		"ret-cc",
		"forward",
		"forward_symbol",
		"forward_mix",
		"macro",
	}

	for tn, base := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		input := string(readTestDataFile(t, base+".asm"))
		expected := readTestDataFile(t, base+".bin")

		prog := evaluateInput(t, input, logger, env)
		logger.Print()
		result := CollectCode(prog)

		if !bytesEqual(result, expected) {
			// t.Errorf("output mismatch (-want +got):\n%s", cmp.Diff(expected, result))
			t.Errorf("[%d] expected %d bytes. got %d bytes", tn, len(expected), len(result))
		}
	}
}
