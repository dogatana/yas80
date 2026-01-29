package evaluator

import (
	"testing"
	"yas80/errcode"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
)

func TestAlign(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{
			input: `align 4 \ ret`,
			code:  []byte{0xc9},
		},
		{
			input: `nop \ align 4 \ ret`,
			code:  []byte{0x00, 0x00, 0x00, 0x00, 0xc9},
		},
		{
			input: `nop \ align 4,255 \ ret`,
			code:  []byte{0x00, 0xff, 0xff, 0xff, 0xc9},
		},
		{
			input: `ds 4 \ align 4 \ ret`,
			code:  []byte{0x00, 0x00, 0x00, 0x00, 0xc9},
		},
		{input: `nop \ align 4, 257 \ ret`, code: []byte{0x00, 0x01, 0x01, 0x01, 0xc9}, err: errcode.WROUND_BYTE},
		{input: `align`, err: errcode.EEBMAC_ARG_COUNT},
		{input: `align 1, 2, 3`, err: errcode.EEBMAC_ARG_COUNT},
		{input: `align 'a'`, err: errcode.EEBMAC_ARG_VALUE},
		{input: `align HL`, err: errcode.EEBMAC_ARG_VALUE},
		{input: `fn func\endf\ align fn()`, err: errcode.EEBMAC_ARG_NULL},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		prog, e := evalInput(tt.input, logger, env)

		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}

		// code
		testCodeResult(t, tn, tt.code, prog)

		// syms
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}
