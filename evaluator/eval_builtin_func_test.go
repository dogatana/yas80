package evaluator

import (
	"testing"
	"yas80/errcode"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
)

// Len, Length
func TestBuiltinFuncLen(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `const v = $len([1,2,3])`, syms: []symValue{{"V", 3}}},
		{input: `const array = [1,2,3] \ var v = $length(array)`, syms: []symValue{{"V", 3}}},
		{input: `const v = $len()`, err: errcode.EBFN_ARG_COUNT},
		{input: `const v = $len(1)`, err: errcode.EBFN_ARG_VALUE},
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

// Len, Length
func TestBuiltinFuncRev(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `rept [1,2,3] \ ld a, $v \endr\ rept $rev([1,2,3]) \ ld a, $v \endr`, code: []byte{0x3e, 1, 0x3e, 2, 0x3e, 3, 0x3e, 3, 0x3e, 2, 0x3e, 1}},
		{input: `rept [1,2,3] \ ld a, $v \endr\ rept $reverse($rev([1,2,3])) \ ld a, $v \endr`, code: []byte{0x3e, 1, 0x3e, 2, 0x3e, 3, 0x3e, 1, 0x3e, 2, 0x3e, 3}},
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
