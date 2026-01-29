package evaluator

import (
	"testing"
	"yas80/errcode"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
)

func TestBuiltinFuncLength(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `const v = $len([1,2,3])`, syms: []symValue{{"V", 3}}},
		{input: `const array = [1,2,3] \ var v = $length(array)`, syms: []symValue{{"V", 3}}},
		{input: `const v = $len()`, err: errcode.EEBFN_ARG_COUNT},
		{input: `const v = $len(1,2)`, err: errcode.EEBFN_ARG_COUNT},
		{input: `const v = $len(1)`, err: errcode.EEBFN_ARG_VALUE},
		{input: `const v = $len(xyz)`, err: errcode.EEBFN_ARG_NULL},
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

func TestBuiltinFuncIsArray(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `const v = $isary([1,2,3])`, syms: []symValue{{"V", 1}}},
		{input: `const v = $isarray([1,2,3])`, syms: []symValue{{"V", 1}}},
		{input: `const v = $isary(1)`, syms: []symValue{{"V", 0}}},
		{input: `const v = $isary("a")`, syms: []symValue{{"V", 0}}},
		{input: `const v = $isary(hl)`, syms: []symValue{{"V", 0}}},
		{input: `const v = $isary()`, err: errcode.EEBFN_ARG_COUNT},
		{input: `const v = $isary(1,2)`, err: errcode.EEBFN_ARG_COUNT},
		{input: `const v = $isary(xyz)`, err: errcode.EEBFN_ARG_NULL},
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

func TestBuiltinFuncReverse(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `rept [1,2,3] \ ld a, $v \endr\ rept $rev([1,2,3]) \ ld a, $v \endr`, code: []byte{0x3e, 1, 0x3e, 2, 0x3e, 3, 0x3e, 3, 0x3e, 2, 0x3e, 1}},
		{input: `rept [1,2,3] \ ld a, $v \endr\ rept $reverse($rev([1,2,3])) \ ld a, $v \endr`, code: []byte{0x3e, 1, 0x3e, 2, 0x3e, 3, 0x3e, 1, 0x3e, 2, 0x3e, 3}},
		{input: `const abc = $rev()`, err: errcode.EEBFN_ARG_COUNT},
		{input: `const abc = $rev(1, 2)`, err: errcode.EEBFN_ARG_COUNT},
		{input: `const abc = $rev(1)`, err: errcode.EEBFN_ARG_VALUE},
		{input: `const abc = $rev(def)`, err: errcode.EEBFN_ARG_NULL},
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

func TestBuiltinFuncFormat(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `db $fmt("ABC")`, code: []byte("ABC")},
		{input: `db $fmt("ABC_%d", 1)`, code: []byte("ABC_1")},
		{input: `db $fmt("ABC_%03d", 1)`, code: []byte("ABC_001")},
		{
			input: `ret \ var num = 123 \ data ## $fmt("_%03d", num) db num`,
			code:  []byte{0xc9, 123},
			syms:  []symValue{{"DATA_123", 1}},
		},
		{input: `db $fmt()`, err: errcode.EEBFN_ARG_COUNT},
		{input: `db $fmt(1,2)`, err: errcode.EEBFN_ARG_VALUE},
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

func TestBuiltinFuncHighLow(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `const v = $h(0x1234)`, syms: []symValue{{"V", 0x12}}},
		{input: `const v = $hi(0x1234)`, syms: []symValue{{"V", 0x12}}},
		{input: `const v = $l(0x1234)`, syms: []symValue{{"V", 0x34}}},
		{input: `const v = $lo(0x1234)`, syms: []symValue{{"V", 0x34}}},
		{input: `const v = $h()`, err: errcode.EEBFN_ARG_COUNT},
		{input: `const v = $h(1,2)`, err: errcode.EEBFN_ARG_COUNT},
		{input: `const v = $h("a")`, err: errcode.EEBFN_ARG_VALUE},
		{input: `const v = $h(xyz)`, err: errcode.EEBFN_ARG_NULL},
		{input: `const v = $l()`, err: errcode.EEBFN_ARG_COUNT},
		{input: `const v = $l(1,2)`, err: errcode.EEBFN_ARG_COUNT},
		{input: `const v = $l("a")`, err: errcode.EEBFN_ARG_VALUE},
		{input: `const v = $l(xyz)`, err: errcode.EEBFN_ARG_NULL},
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

func TestBuiltinFuncWord(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `dd $w(0x12)`, code: []byte{0x12, 0}},
		{input: `db $w(0x12)`, code: []byte{0x12, 0}},
		{input: `db $w()`, err: errcode.EEBFN_ARG_COUNT},
		{input: `db $w(1,2)`, err: errcode.EEBFN_ARG_COUNT},
		{input: `db $w("a")`, err: errcode.EEBFN_ARG_VALUE},
		{input: `const v \ db $w(v)`, err: errcode.EEBFN_ARG_NULL},
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
