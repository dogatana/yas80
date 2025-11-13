package evaluator

import (
	"fmt"
	"testing"
	"yas80/logger"
	"yas80/object"
)

func TestEvalExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1", 1},
		{"1 + 2", 3},
		{"1 * 2", 2},
		{"1 + 2 * 3", 7},
		{`const val = 123 \ val`, 123},
		{`const val = 11 \ val * val`, 121},
		{`const val = 11 \ const val2 = val * val \ val2`, 121},
		{"const val = 11 \r\n const val2 = val * val \r\n val2", 121},
	}

	for _, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		prog := evaluateInput(t, tt.input, logger, env)

		last := prog.Objects[len(prog.Objects)-1]
		testNumberObject(t, last, tt.expected, tt.input)
	}
}

func TestReturn(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1", 1},
		{`1 \ 2 \ 3`, 3},
		{`1 \ 2 \ return \ 3`, 2},
		{`1 \ 2 \ return 99 \ 3`, 99},
	}

	for _, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		prog := evaluateInput(t, tt.input, logger, env)

		last := prog.Objects[len(prog.Objects)-1]
		testNumberObject(t, last, tt.expected, tt.input)
	}
}

func TestIf(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`if 1 \ endif`, -1}, // -1 は NullObject とする
		{`if 0 \ endif`, -1},
		{`if 1 \ 100 \ endif`, -1},
		{`if 0 \ 100 \ endif`, -1},

		{`if 1 \ 100 \ endif`, 100},
		{`if 0 \ 100 \ endif`, -1},
		{`if 1 \ 100 \ else \ 200  \ endif`, 100},
		{`if 1 \ 100 \ else \ 200  \ endif`, 200},
	}

	for _, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		prog := evaluateInput(t, tt.input, logger, env)

		last := prog.Objects[len(prog.Objects)-1]
		if tt.expected >= 0 {
			testNumberObject(t, last, tt.expected, tt.input)
			continue
		}
		if last != object.NULL {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("should be NULL. got %T(%#v)", last, last)
		}
	}
}
