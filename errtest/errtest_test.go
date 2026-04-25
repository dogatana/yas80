package errtest

import (
	"testing"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/internal/testutil"
	"github.com/dogatana/yas80/logging"
	"github.com/dogatana/yas80/object"
)

func TestErrorConstLabel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// 0-
		{`const abc = 0 \ const abc = 1`, errcode.ECONST_DUP},
		{`abc: nop \ abc: nop`, errcode.ELABEL_DUP},
		{`abc: \ nop \ abc: nop`, errcode.ELABEL_DUP},
		{`abc: nop \ abc:  \ nop`, errcode.ELABEL_DUP},
		{`abc: \  nop \ abc: \ nop`, errcode.ELABEL_DUP},
		// 5-
		{`abc: nop \ const abc = 123`, errcode.ECONST_DUP},
		{`const abc = 123 \ abc: nop`, errcode.ELABEL_USED},
		{`function abc() x \ const abc = 1`, errcode.ECONST_USED},
		{`const abc = def \ const def = abc`, errcode.ESYM_CYCLIC},
		{`const abc = def + 1 \ const def = xyz + 2 \ const xyz = abc + 3`, errcode.ESYM_CYCLIC},
	}
	for tn, tt := range tests {
		logger := logging.New()
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testutil.TestLogMessage(t, tn, tt.expected, logger)
	}
}

func TestErrorScrope(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// 0-
		{`const @abc = 0 \ ld a, @abc`, errcode.ESCOPE_MACRO},
		{`const .abc = 0 \ ld a, .abc`, errcode.ESCOPE_PROC},
		{`if 1 \ const @abc = 1 \ ld a, @abc \ endif`, errcode.ESCOPE_MACRO},
		{`if 1 \ const .abc = 1 \ ld a, .abc \ endif`, errcode.ESCOPE_PROC},
		{`test func arg \ if arg \ const @abc = 1 \ endif\ ld a, @abc \ endf \ _ = test(1)`, errcode.ESCOPE_MACRO},
	}
	for tn, tt := range tests {
		logger := logging.New()
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testutil.TestLogMessage(t, tn, tt.expected, logger)
	}
}

func TestErrorFuncCall(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// 0-
		{`abc func \ return 1 \ endf \ _=abc(0)`, errcode.EFUNC_ARG_COUNT},
		{`abc func arg \ return 1 \ endf \ _=abc()`, errcode.EFUNC_ARG_COUNT},
		{`abc func arg \ return 1 \ endf \ _=abc(0, 1)`, errcode.EFUNC_ARG_COUNT},
		{`_=abc()`, errcode.EFUNC_UNDEF},
		{`_=abc(0)`, errcode.EFUNC_UNDEF},
		// 5-
		{`const abc = 1 \ _=abc(0)`, errcode.EFUNC_NOT_FUNC},
		{`const abc = 1 \ _=abc()`, errcode.EFUNC_NOT_FUNC},
	}
	for tn, tt := range tests {
		logger := logging.New()
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testutil.TestLogMessage(t, tn, tt.expected, logger)
	}

}

func TestErrorRept(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// 0-
		{`rept a \ nop \ endr`, errcode.EREPT_ARG},
	}
	for tn, tt := range tests {
		logger := logging.New()
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testutil.TestLogMessage(t, tn, tt.expected, logger)
	}
}
