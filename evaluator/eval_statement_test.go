package evaluator

import (
	"testing"
	"yas80/logger"
	"yas80/object"
)

func TestAssemble(t *testing.T) {
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

		prog := evaluateInput(t, input, logger, env)
		logger.Print()
		result := CollectCode(prog)

		if !bytesEqual(result, expected) {
			t.Errorf("[%d] expected %d bytes. got %d bytes", tn, len(expected), len(result))
		}
	}
}

func TestConstStatement(t *testing.T) {
	tests := []struct {
		input    string
		name     string
		expected int
	}{
		{"const abc ## 123 = 456", "ABC123", 456},
		{"const abc ## (100 + 23) = 456", "ABC123", 456},
		{"const abc ## (100 + 23) = def ## 456 \\ const def ## (400 + 56) = 999", "ABC123", 999},
	}
	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		_ = evaluateInput(t, tt.input, logger, env)

		obj, ok := env.Get(tt.name)
		if !ok {
			t.Fatalf(`[%d] %q not in env`, tn, tt.name)
		}
		value := evalValue(obj)
		testNumberObject(t, tn, value, tt.expected)
	}
}
