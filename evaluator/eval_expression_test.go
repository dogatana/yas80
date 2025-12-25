package evaluator

import (
	"testing"
	"yas80/logger"
	"yas80/object"
)

func TestEvalExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		// 0-
		{"const result = 1", 1},
		{"const result = 1 + 2", 3},
		{"const result = 1 * 2", 2},
		{"const result = 1 + 2 * 3", 7},
		{"const result = 10 % 2", 0},
		//5-
		{"const result = 10 + 10 % 3", 11},
		{`const val = 123 \ const result = val`, 123},
		{`const val = 11 \ const result = val * val`, 121},
		{`const val = 11 \ const val2 = val * val \ const result = val2`, 121},
		{"const val = 11 \r\n const val2 = val * val \r\n const result = val2", 121},
		// 10-
		{`nop \ abc proc \ ret \ endp \ const result = abc`, 1},
		{`const result = abc \ nop \ abc proc \ ret \ endp`, 1},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		_, _ = evaluateInput(t, tt.input, logger, env)

		obj, ok := env.Get("RESULT")
		if !ok {
			t.Fatalf(`[%d] "RESULT" not in env`, tn)
		}

		value := evalValue(obj)
		testNumberObject(t, tn, value, tt.expected)
	}
}
