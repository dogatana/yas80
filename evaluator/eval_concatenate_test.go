package evaluator

import (
	"testing"
	"yas80/errcode"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
)

func TestConcatenateSymbol(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `const abc ## (100 + 23) = 456`, syms: []symValue{{"ABC123", 456}}},
		{input: `const abc ## ("_" + "123") = 456`, syms: []symValue{{"ABC_123", 456}}},
		{input: `abc ## (100 + 23) equ 456`, syms: []symValue{{"ABC123", 456}}},
		{input: `nop \ abc ## (100 + 23): ret`, syms: []symValue{{"ABC123", 1}}, code: []byte{0, 0xc9}},
		{input: `nop \ abc ## (100 + 23): \ ret`, syms: []symValue{{"ABC123", 1}}, code: []byte{0, 0xc9}},
		{input: `nop \ abc ## (100 + 23) ds 1, 0xc9`, syms: []symValue{{"ABC123", 1}}, code: []byte{0, 0xc9}},
		{input: `test proc \ nop \ abc ## (100 + 23) ds 1, 0xc9 \ endp`, syms: []symValue{{"ABC123", 1}}, code: []byte{0, 0xc9}},
		{input: `test macro arg\ abc ## arg: ret \ endm \ nop \ test 123`, syms: []symValue{{"ABC123", 1}}, code: []byte{0, 0xc9}},
		{input: `test macro arg\ abc ## arg ds arg, 255 \ endm \ nop \ test 2`, syms: []symValue{{"ABC2", 1}}, code: []byte{0, 255, 255}},
		{input: `rept 3\ abc ## $i: ld a, $i\ endr`,
			syms: []symValue{{"ABC0", 0}, {"ABC1", 2}, {"ABC2", 4}},
			code: []byte{0x3e, 0, 0x3e, 1, 0x3e, 2}},
		// 10
		{input: `const abc ## 123 = 1 \ ds abc ## 123, abc ## 123`, syms: []symValue{{"ABC123", 1}}, code: []byte{1}},
		{input: `nop \ abc ## 123 proc \ ret \ endp`, syms: []symValue{{"ABC123", 1}}, code: []byte{0, 0xc9}},
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
func TestConcatenateSymbolError(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		{input: `const 123 ## 123 = 1`, err: errcode.ESYNTAX},
		{input: `const abc ## a = 1`, err: errcode.ECONCAT_TYPE},
		{input: `const abc ## hl = 1`, err: errcode.ECONCAT_TYPE},
		{input: `const abc ## cy = 1`, err: errcode.ECONCAT_TYPE},
		{input: `function abc() 1 \ const aaa ## abc = 1`, err: errcode.ECONCAT_TYPE},
		{input: `const abc ## cy ## 1 = 1`, err: errcode.ESYNTAX}, // syntax error
		{input: `const abc ## 123 = 456 \ const abc123 = 123`, err: errcode.ECONST_DUP},
		{input: `const abc123 = 123 \ const abc ## 123 = 456`, err: errcode.ECONST_DUP},
		{input: `const abc ## def = 123`, err: errcode.ESYM_UNDEF},
		{input: `abc ## a proc \ endp`, err: errcode.ECONCAT_TYPE},
		{input: `abc ## cy proc \ endp`, err: errcode.ECONCAT_TYPE},
		{input: `abc ## def proc \ endp`, err: errcode.ESYM_UNDEF},
		{input: `const abc123 = 123 \ abc ## 123 proc \ endp`, err: errcode.EPROC_USED},
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
