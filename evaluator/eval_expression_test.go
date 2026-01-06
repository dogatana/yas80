package evaluator

import (
	"testing"
	"yas80/errcode"
	"yas80/logging"
	"yas80/object"
)

func TestEvalExpression(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		{"const result = 1", []symValue{{"RESULT", 1}}, ""},
		{"const result = 1 + 2", []symValue{{"RESULT", 3}}, ""},
		{"const result = 1 * 2", []symValue{{"RESULT", 2}}, ""},
		{"const result = 1 - 2 * 3", []symValue{{"RESULT", -5}}, ""},
		{"const result = 10 % 2", []symValue{{"RESULT", 0}}, ""},
		//5-
		{"const result = 10 + 10 % 3", []symValue{{"RESULT", 11}}, ""},
		{`const val = 123 \ const result = val`, []symValue{{"RESULT", 123}}, ""},
		{`const val = 11 \ const result = val * val`, []symValue{{"RESULT", 121}}, ""},
		{`const val = 11 \ const val2 = val * val \ const result = val2`, []symValue{{"RESULT", 121}}, ""},
		{"const val = 11 \r\n const val2 = val * val \r\n const result = val2", []symValue{{"RESULT", 121}}, ""},
		// 10-
		{`nop \ abc proc \ ret \ endp \ const result = abc`, []symValue{{"RESULT", 1}}, ""},
		{`const result = abc \ nop \ abc proc \ ret \ endp`, []symValue{{"RESULT", 1}}, ""},
		{input: `_ = 0x`, err: errcode.ENUMBER},
		{input: `_ = 0o`, err: errcode.ENUMBER},
		{input: `_ = 0b`, err: errcode.ENUMBER},
		// 15-
		{input: `_ = abc[]`, err: errcode.EARRAY_INDEX},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testLogMessage(t, tn, tt.err, e.logger)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}
