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
		obj, ok := last.(*object.NumberObject)
		if !ok {
			fmt.Printf("input %q\n", tt.input)
			t.Fatalf("prog.Objects[-1] not NumberObject. got %T", last)
		}
		if obj.Value != tt.expected {
			fmt.Printf("input %q\n", tt.input)
			fmt.Println(prog.String())
			t.Errorf("object is not %d. got %d", tt.expected, obj.Value)
		}
	}
}
