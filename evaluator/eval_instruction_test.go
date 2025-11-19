package evaluator

import (
	"testing"
	"yas80/logger"
	"yas80/object"

	"github.com/google/go-cmp/cmp"
)

func TestZ80Instruction(t *testing.T) {
	tests := []string{
		"inst0",
		"ld_r8_r8",
		"forward",
		"forward_symbol",
	}

	for _, base := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		input := string(readTestDataFile(t, base+".asm"))
		expected := readTestDataFile(t, base+".bin")

		prog := evaluateInput(t, input, logger, env)
		logger.Print()
		result := collectCode(prog)

		if !cmp.Equal(result, expected) {
			t.Errorf("output mismatch (-want +got):\n%s", cmp.Diff(expected, result))
		}
	}
}
