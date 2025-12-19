package evaluator

import (
	"testing"
	"yas80/errcode"
	"yas80/logger"
	"yas80/object"
)

func TestError(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"const abc = 123 / 0", errcode.EBIN_OP_DIVZERO},
		{`const abc = 0 \ ld a, 1 / abc`, errcode.EBIN_OP_DIVZERO},
		{`test macro arg \ ld a, 1 / arg \ endm \ test 0`, errcode.EBIN_OP_DIVZERO},
		{`const abc = 0 \ const abc = 1`, errcode.ESYM_DUP},
		{`abc: nop \ abc: nop`, errcode.ELABEL_DUP},
		{`abc: \ nop \ abc: nop`, errcode.ELABEL_DUP},
		{`abc: nop \ abc:  \ nop`, errcode.ELABEL_DUP},
		{`abc: \  nop \ abc: \ nop`, errcode.ELABEL_DUP},
		{`abc: nop \ const abc = 123`, errcode.ESYM_DUP},
		{`const abc = 123 \ abc: nop`, errcode.ELABEL_USED},
	}
	for tn, tt := range tests {
		logger := logger.New("test")
		env := object.NewEnvironment(nil)
		evaluateErrorInput(tt.input, logger, env)
		if len(logger.Errors) == 0 {
			t.Fatalf("[%d] no error", tn)
		}
		if !hasError(logger, tt.expected) {
			t.Errorf("[%d] error dose not contains %s. got %s", tn, tt.expected, logger.Errors[0].Message)
		}
	}
}
