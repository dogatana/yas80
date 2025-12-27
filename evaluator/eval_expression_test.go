package evaluator

import (
	"testing"
	"yas80/logging"
	"yas80/object"
)

func TestEvalExpression(t *testing.T) {
	tests := []struct {
		input string
		syms  []SymValue
	}{
		// 0-
		{"const result = 1", []SymValue{{"RESULT", 1}}},
		{"const result = 1 + 2", []SymValue{{"RESULT", 3}}},
		{"const result = 1 * 2", []SymValue{{"RESULT", 2}}},
		{"const result = 1 + 2 * 3", []SymValue{{"RESULT", 7}}},
		{"const result = 10 % 2", []SymValue{{"RESULT", 0}}},
		//5-
		{"const result = 10 + 10 % 3", []SymValue{{"RESULT", 11}}},
		{`const val = 123 \ const result = val`, []SymValue{{"RESULT", 123}}},
		{`const val = 11 \ const result = val * val`, []SymValue{{"RESULT", 121}}},
		{`const val = 11 \ const val2 = val * val \ const result = val2`, []SymValue{{"RESULT", 121}}},
		{"const val = 11 \r\n const val2 = val * val \r\n const result = val2", []SymValue{{"RESULT", 121}}},
		// 10-
		{`nop \ abc proc \ ret \ endp \ const result = abc`, []SymValue{{"RESULT", 1}}},
		{`const result = abc \ nop \ abc proc \ ret \ endp`, []SymValue{{"RESULT", 1}}},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, "", e)

		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}
