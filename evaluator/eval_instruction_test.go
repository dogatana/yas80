package evaluator

import (
	"testing"
	"yas80/logging"
	"yas80/object"
)

func TestZ80Instruction(t *testing.T) {
	tests := []string{
		"inst0",
		"ld_r8_r8",
		"ret-cc",
	}

	for tn, base := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		input := string(readTestDataFile(t, base+".asm"))
		expected := readTestDataFile(t, base+".bin")

		prog, e := evalInput(input, logger, env)
		testEvalResult(t, tn, "", e)

		logger.Print()
		result := CollectCode(prog)

		if err := bytesEqual(result, expected); err != nil {
			t.Errorf("[%d] generated code diff %s", tn, err.Error())
		}
	}
}
