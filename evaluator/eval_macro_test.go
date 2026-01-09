package evaluator

import (
	"testing"
	"yas80/errcode"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
)

func TestErrorMacroDef(t *testing.T) {
	tests := []struct {
		input string
		err   string
	}{
		// 0-
		{`@abc macro \ endm`, errcode.EMACRO_NAME},
		{`.abc macro \ endm`, errcode.EMACRO_NAME},
		{`const abc = 1 \ abc macro \ endm`, errcode.EMACRO_USED},
		{`abc macro \ endm \ abc macro \ endm`, errcode.EMACRO_DUP},
	}
	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New("<test>")
		_, e := evalInput(tt.input, logger, env)

		// error, warning, information
		testutil.TestLogMessage(t, tn, tt.err, e.logger)
	}
}

func TestErrorMacroCall(t *testing.T) {
	tests := []struct {
		input string
		err   string
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
		{`abc macro \ return \ endm \ abc`, errcode.WSCOPE_MACRO},
		{`abc macro \ fn func \ return 1 \ endf \ endm \ abc`, errcode.WSCOPE_MACRO},
	}
	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New("<test>")
		_, e := evalInput(tt.input, logger, env)

		// error, warning, information
		testutil.TestLogMessage(t, tn, tt.err, e.logger)
	}
}

func TestWarningMacroCall(t *testing.T) {
	tests := []struct {
		input string
		err   string
	}{
		// 0-
		// TODO enum
		// TODO proc
	}
	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New("<test>")
		_, e := evalInput(tt.input, logger, env)

		// error, warning, information
		testutil.TestLogMessage(t, tn, tt.err, e.logger)
	}
}

func TestExitm(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
	}{
		{`
			test1 macro
			ld a, 0
			exitm
			ld a, 1
			endm

			test2 macro
			ld hl, 0
			test1
			ld hl, 1
			endm

			test1
			test2
		`, []byte{
			0x3E, 0x00, // ld a, 0
			0x21, 0x00, 0x00, // ld hl, 0
			0x3E, 0x00, // ld a, 0
			0x21, 0x01, 0x00, // ld hl, 1
		},
		},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		input := tt.input

		prog, e := evalInput(input, logger, env)
		testEvalResult(t, tn, "", e)

		testCodeResult(t, tn, tt.code, prog)
	}
}

func TestMacroIf(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
	}{
		{
			`test macro arg \ if arg > 0 \ ld a, 1 \ else \ ld a, $ff \ endif \ endm \ test 1`,
			[]byte{0x3E, 0x01},
		},
		{
			`test macro arg \ if arg > 0 \ ld a, 1 \ else \ ld a, $ff \ endif \ endm \ test -1`,
			[]byte{0x3E, 0xFF},
		},
		{
			`test macro arg \ ld a, 0 \ exitm if arg == 0 \ ld a, 2 \ exitm if arg == 1 \ ld a, $ff \endm \ test 0`,
			[]byte{0x3E, 0x00},
		},
		{
			`test macro arg \ ld a, 0 \ exitm if arg == 0 \ ld a, 1 \ exitm if arg == 1 \ ld a, $ff \endm \ test 1`,
			[]byte{0x3E, 0x00, 0x3E, 0x01},
		},
		{
			`test macro arg \ ld a, 0 \ exitm if arg == 0 \ ld a, 1 \ exitm if arg == 1 \ ld a, $ff \endm \ test 2`,
			[]byte{0x3E, 0x00, 0x3E, 0x01, 0x3E, 0xFF},
		},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		input := tt.input

		prog, e := evalInput(input, logger, env)
		testEvalResult(t, tn, "", e)

		testCodeResult(t, tn, tt.code, prog)
	}
}

func TestMacroConst(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		{
			input: `test macro\ const abc = 123 \ endm \ test`,
			syms:  []symValue{{"ABC", 123}},
		},
		{
			input: `test macro\ const abc = 123 \ endm \ test \ test`,
			err:   errcode.ECONST_DUP,
		},
		{
			input: `test macro\ const @abc = 123 \ endm \ test \ test`,
			syms:  []symValue{{"__1_TEST@ABC", 123}, {"__2_TEST@ABC", 123}},
		},
		{
			input: `test macro arg \ if arg == 0 \ const @abc = arg \ elif arg == 1 \ const @abc = arg * 16 \ else \ const @abc = arg * 256\ endif \ ld hl, @abc \ endm \ test 0 \ test 1 \ test 2`,
			code:  []byte{0x21, 0x00, 0x00, 0x21, 0x10, 0x00, 0x21, 0x00, 0x02},
		},
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

func TestMacroLabel(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		{
			input: `test macro\ nop \ abc: \ ret \ endm \ test`,
			code:  []byte{0, 0xc9},
			syms:  []symValue{{"ABC", 1}},
		},
		{
			input: `test macro\ nop \ abc: \ ret \ endm \ test \ test`,
			err:   errcode.ELABEL_DUP,
		},
		{
			input: `test macro\ nop \ abc: ret \ endm \ test \ test`,
			err:   errcode.ELABEL_DUP,
		},
		{
			input: `test macro arg\if arg==0\const abc=0\else\const abc=$1234\endif\ld hl, abc\endm\ test 0`,
			code:  []byte{0x21, 0x00, 0x00},
			syms:  []symValue{{"ABC", 0}},
		},
		{
			input: `test macro arg\if arg==0\const abc=0\else\const abc=$1234\endif\ld hl, abc\endm\ test 1`,
			code:  []byte{0x21, 0x34, 0x12},
			syms:  []symValue{{"ABC", 0x1234}},
		},
		{
			input: `test macro\ nop \ @abc: \ ret \ endm \ test \ test`,
			code:  []byte{0, 0xc9, 0, 0xc9},
			syms: []symValue{
				{"__1_TEST@ABC", 1},
				{"__2_TEST@ABC", 3},
			},
		},
		{
			input: `test macro\ nop \ @abc: ret \ endm \ test \ test`,
			code:  []byte{0, 0xc9, 0, 0xc9},
			syms: []symValue{
				{"__1_TEST@ABC", 1},
				{"__2_TEST@ABC", 3},
			},
		},
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

func TestMacroData(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		{
			input: `test macro\ nop \ abc db 1 \ endm \ test`,
			code:  []byte{0, 1},
			syms:  []symValue{{"ABC", 1}},
		},
		{
			input: `test macro\ nop \ abc db 1 \ endm \ test \ test`,
			err:   errcode.ELABEL_DUP,
		},
		{
			input: `test macro\ nop \ abc db 1 \ endm \ test \ test`,
			err:   errcode.ELABEL_DUP,
		},
		{
			input: `test macro\ nop \ @abc ds 1, 255 \ endm \ test \ test`,
			code:  []byte{0, 0xff, 0, 0xff},
			syms: []symValue{
				{"__1_TEST@ABC", 1},
				{"__2_TEST@ABC", 3},
			},
		},
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
