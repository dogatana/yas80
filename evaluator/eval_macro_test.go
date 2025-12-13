package evaluator

import (
	"testing"
	"yas80/logger"
	"yas80/object"
)

func TestMacro(t *testing.T) {
	tests := []string{
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
			t.Errorf("[%d] expected %d bytes. got %d bytes", tn, len(expected), len(result))
		}
	}
}
