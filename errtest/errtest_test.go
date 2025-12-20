package errtest

import (
	"testing"
	"yas80/errcode"
	"yas80/logger"
	"yas80/object"
)

func TestErrorExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"const abc = 123 / 0", errcode.EBIN_OP_DIVZERO},
		{`const abc = 0 \ ld a, 1 / abc`, errcode.EBIN_OP_DIVZERO},
		{`test macro arg \ ld a, 1 / arg \ endm \ test 0`, errcode.EBIN_OP_DIVZERO},
	}
	for tn, tt := range tests {
		logger := logger.New("test")
		env := object.NewEnvironment(nil)
		evaluateErrorInput(tt.input, logger, env)
		testError(t, tn, logger, tt.expected)
	}
}

func TestErrorConstLabel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
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
		testError(t, tn, logger, tt.expected)
	}
}

func TestErrorScrope(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`const @abc = 0 \ ld a, @abc`, errcode.ESCOPE_GLOBAL},
		{`const .abc = 0 \ ld a, .abc`, errcode.ESCOPE_GLOBAL},
		{`if 1 \ const @abc = 1 \ ld a, @abc \ endif`, errcode.ESCOPE},
		{`if 1 \ const .abc = 1 \ ld a, .abc \ endif`, errcode.ESCOPE},
		{`test func arg \ if arg \ const @abc = 1 \ endif\ ld a, @abc \ endf \ _ = test(1)`, errcode.ESCOPE},
	}
	for tn, tt := range tests {
		logger := logger.New("test")
		env := object.NewEnvironment(nil)
		evaluateErrorInput(tt.input, logger, env)
		testError(t, tn, logger, tt.expected)
	}
}

func TestErrorFunc(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{`@abc func \ endf`, errcode.EFUNC_NAME},
		{`.abc func \ endf`, errcode.EFUNC_NAME},
		{`const abc = 0 \ abc func \ endf`, errcode.EFUNC_USED},
		{`abc: nop \ abc func \ endf`, errcode.EFUNC_USED},
		{`abc func \ endf \ abc func \ endf`, errcode.EFUNC_DUP},
	}
	for tn, tt := range tests {
		logger := logger.New("test")
		env := object.NewEnvironment(nil)
		evaluateErrorInput(tt.input, logger, env)
		testError(t, tn, logger, tt.expected)
	}

}
