package errtest

import (
	"testing"
	"yas80/errcode"
	"yas80/logging"
	"yas80/object"
)

func TestErrorExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// 0-
		{"const abc = 123 / 0", errcode.EBIN_OP_DIVZERO},
		{`const abc = 0 \ ld a, 1 / abc`, errcode.EBIN_OP_DIVZERO},
		{`test macro arg \ ld a, 1 / arg \ endm \ test 0`, errcode.EBIN_OP_DIVZERO},
	}
	for tn, tt := range tests {
		logger := logging.New("test")
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testMessage(t, TEST_ERROR, tn, logger, tt.expected)
	}
}

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
		logger := logging.New("test")
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testMessage(t, TEST_ERROR, tn, logger, tt.expected)
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
		logger := logging.New("test")
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testMessage(t, TEST_ERROR, tn, logger, tt.expected)
	}
}

func TestErrorFuncDef(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// tn: 0-
		{`@abc func \ endf`, errcode.EFUNC_NAME},
		{`.abc func \ endf`, errcode.EFUNC_NAME},
		{`const abc = 0 \ abc func \ endf`, errcode.EFUNC_USED},
		{`abc: nop \ abc func \ endf`, errcode.EFUNC_USED},
		{`abc func \ endf \ abc func \ endf`, errcode.EFUNC_DUP},
		// tn: 5-
		{`function @abc() 1`, errcode.EFUNC_NAME},
		{`function .abc() 1`, errcode.EFUNC_NAME},
	}
	for tn, tt := range tests {
		logger := logging.New("test")
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testMessage(t, TEST_ERROR, tn, logger, tt.expected)
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
		logger := logging.New("test")
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testMessage(t, TEST_ERROR, tn, logger, tt.expected)
	}

}
func TestErrorMacroDef(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// 0-
		{`@abc macro \ endm`, errcode.EMACRO_NAME},
		{`.abc macro \ endm`, errcode.EMACRO_NAME},
		{`const abc = 1 \ abc macro \ endm`, errcode.EMACRO_USED},
		{`abc macro \ endm \ abc macro \ endm`, errcode.EMACRO_DUP},
	}
	for tn, tt := range tests {
		logger := logging.New("test")
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testMessage(t, TEST_ERROR, tn, logger, tt.expected)
	}
}

func TestErrorMacroCall(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// 0-
		{`abc`, errcode.EMACRO_UNDEF},
		{`abc 1`, errcode.EMACRO_UNDEF},
		{`abc macro \ endm \ abc 1`, errcode.EMACRO_ARG_COUNT},
		{`abc macro arg \ endm \ abc`, errcode.EMACRO_ARG_COUNT},
		{`abc macro arg \ endm \ abc 1, 2`, errcode.EMACRO_ARG_COUNT},
		// 5-
		{`aaa macro \ nop \ bbb macro \ nop \ endm \ endm \ aaa`, errcode.EMACRO_NEST},
		{`aaa macro \ bbb \ endm \ bbb macro \ aaa \ endm \ aaa`, errcode.EMACRO_CYCLIC},
	}
	for tn, tt := range tests {
		logger := logging.New("test")
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testMessage(t, TEST_ERROR, tn, logger, tt.expected)
	}
}

func TestWarningMacroCall(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// 0-
		{`abc macro \ return \ endm \ abc`, errcode.WSCOPE_MACRO},
		{`abc macro \ fn func \ return 1 \ endf \ endm \ abc`, errcode.WSCOPE_MACRO},
		// TODO enum
		// TODO proc
	}
	for tn, tt := range tests {
		logger := logging.New("test")
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_WARNING, tt.input, logger, env)
		testMessage(t, TEST_WARNING, tn, logger, tt.expected)
	}
}

func TestErrorSymbol(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// 0-
		{`ld hl, abc.xxx`, errcode.ESYM_UNDEF},
		{`ld hl, abc.xxx \abc proc \ nop \ .def: ret \ nop \ endp`, errcode.ESYM_UNDEF},
	}
	for tn, tt := range tests {
		logger := logging.New("test")
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testMessage(t, TEST_ERROR, tn, logger, tt.expected)
	}
}

func TestErrorRept(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// 0-
		{`rept a \ nop \ endr`, errcode.EREPT_COUNT},
	}
	for tn, tt := range tests {
		logger := logging.New("test")
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testMessage(t, TEST_ERROR, tn, logger, tt.expected)
	}
}

func TestErrorEnumDef(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		// 0-
		{`abc enum \ ende \ abc enum \ ende`, errcode.EENUM_DUP},
		{`const abc = 1 \ abc enum \ ende`, errcode.EENUM_USED},
		{`abc: nop \ abc enum \ ende`, errcode.EENUM_USED},
		{`abc enum \ aaa = 1 \ aaa = 1 \ ende`, errcode.EENUM_ELE_DUP},
		{`abc enum \ aaa = zzz \ ende \ const zzz = 1`, errcode.EENUM_ELE_FWD},
		// 5-
		{`abc enum \ aaa = 1 \ ende \ ld a, abc.xyz`, errcode.EENUM_ELE_UNDEF},
	}
	for tn, tt := range tests {
		logger := logging.New("test")
		env := object.NewEnvironment(nil)
		evaluateInput(TEST_ERROR, tt.input, logger, env)
		testMessage(t, TEST_ERROR, tn, logger, tt.expected)
	}
}
