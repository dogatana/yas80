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
		{`tm macro \ if 0 \ tm2 macro \ nop \ endm \ endif \ endm`, errcode.EMACRO_NEST},
		{`tm macro \ if 0 \ te enum \ ende \ endif \ endm`, errcode.WSCOPE_MACRO},
	}
	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
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
		logger := logging.New()
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
		logger := logging.New()
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
		{`
		rept 2
			ld a, $i + 1
			rept 2
				ld a, ($i + 1) * 16
				rept 3
				ld a, ($i + 1)* 32
				exitm if $i ==1
				ld hl, ($i + 1) * 32
				endr
				nop
			endr
			ret
		endr`, []byte{
			0x3e, 0x01,
			0x3e, 0x10,
			0x3e, 0x20, 0x21, 0x20, 0x00,
			0x3e, 0x40,
			0x3e, 0x60, 0x21, 0x60, 0x00,
			0x00,
			0x3e, 0x20,
			0x3e, 0x20, 0x21, 0x20, 0x00,
			0x3e, 0x40,
			0x3e, 0x60, 0x21, 0x60, 0x00,
			0x00,
			0xc9,

			0x3e, 0x02,
			0x3e, 0x10,
			0x3e, 0x20, 0x21, 0x20, 0x00,
			0x3e, 0x40,
			0x3e, 0x60, 0x21, 0x60, 0x00,
			0x00,
			0x3e, 0x20,
			0x3e, 0x20, 0x21, 0x20, 0x00,
			0x3e, 0x40,
			0x3e, 0x60, 0x21, 0x60, 0x00,
			0x00,
			0xc9},
		},
		{`
		tm0 macro arg
			ld a, 0
			tm1 arg
			ld hl, 0
		endm

		tm1 macro arg
			ld a, 1
			tm2 arg
			ld hl, 1
		endm

		tm2 macro arg
			ld a, 2
			exitm if arg == 2
			ld hl, 2
		endm

		tm0 2`, []byte{
			0x3e, 0x00,
			0x3e, 0x01,
			0x3e, 0x02,
			0x21, 0x01, 0x00,
			0x21, 0x00, 0x00,
		}},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
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
		logger := logging.New()
		input := tt.input

		prog, e := evalInput(input, logger, env)
		testEvalResult(t, tn, "", e)

		testCodeResult(t, tn, tt.code, prog)
	}
}

func TestMacroConstVar(t *testing.T) {
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
		{
			input: `test macro arg \ const @local = arg \ const @local2 = -@local \ ld a, @local2 \ endm \ test 1`,
			code:  []byte{0x3e, 0xff},
			syms: []symValue{
				{"__1_TEST@LOCAL", 1},
				{"__1_TEST@LOCAL2", -1},
			},
		},
		{
			input: `test macro \ var @aaa = 1 \ const @bbb = 2 \ endm\ test`,
			syms: []symValue{
				{"__1_TEST@AAA", 1},
				{"__1_TEST@BBB", 2},
			},
		},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
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
		logger := logging.New()
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
		logger := logging.New()
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

func TestReptArray(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		{
			input: `rept [1, 2, 3] \ ld a, $v \ endr`,
			code:  []byte{0x3e, 1, 0x3e, 2, 0x3e, 3},
		},
		{
			input: `const val = [1,2,3] \ rept val \ ld a, $v \ endr`,
			code:  []byte{0x3e, 1, 0x3e, 2, 0x3e, 3},
		},
		{
			input: `const val = [1,def,3] \ rept val \ ld a, $v \ endr \ const def = 2`,
			code:  []byte{0x3e, 1, 0x3e, 2, 0x3e, 3},
		},
		{
			input: `tm macro arg \ rept arg \ ld a, $v \ endr \ endm \ tm [1,2,3]`,
			code:  []byte{0x3e, 1, 0x3e, 2, 0x3e, 3},
		},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
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

func TestMacroArray(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		{
			input: `
			const ary = [1,2,3]
			tm macro arg
			ld a, arg[0]
			ld a, arg[1]
			ld a, arg[2]
			endm
			tm ary
			tm [4,5,6]`,

			code: []byte{0x3e, 1, 0x3e, 2, 0x3e, 3, 0x3e, 4, 0x3e, 5, 0x3e, 6},
		},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
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
func TestMacroReptCombination(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		{ // macro / macro
			input: `t1 macro arg\ld a,arg\endm\t2 macro arg\ld hl,arg\endm\t3 macro arg\t1 arg\t2 arg\endm\t3 255`,
			code:  []byte{0x3e, 0xff, 0x21, 0xff, 0},
		},
		{ // macro / rept
			input: `t1 macro arg\ rept arg \ ld a, $i \endr \ endm \ t1 3`,
			code:  []byte{0x3e, 0, 0x3e, 1, 0x3e, 2},
		},
		{ // macro / rept - error
			input: `t1 macro arg\ rept arg \ ld a, $v \endr \ endm \ t1 3`,
			err:   errcode.ESYM_UNDEF,
		},
		{ // macro / rept
			input: `t1 macro arg\ rept arg \ ld a, $v \endr \ endm \ t1 [16, 17, 18]`,
			code:  []byte{0x3e, 16, 0x3e, 17, 0x3e, 18},
		},
		{ // macro / rept
			input: `t1 macro arg\ rept arg \ ld a, $i \endr \ endm \ t1 [16, 17, 18]`,
			code:  []byte{0x3e, 0, 0x3e, 1, 0x3e, 2},
		},
		{ // rept / macro
			input: `t1 macro arg\ ld a, arg \ endm \ rept 3 \ t1 $i \ endr`,
			code:  []byte{0x3e, 0, 0x3e, 1, 0x3e, 2},
		},
		{ // rept / macro - error
			input: `t1 macro arg\ ld a, arg \ endm \ rept 3 \ t1 $v \ endr`,
			err:   errcode.ESYM_UNDEF,
		},
		{ // rept / macro
			input: `t1 macro arg\ ld a, arg \ endm \ rept [16, 17, 18] \ t1 $v \ endr`,
			code:  []byte{0x3e, 16, 0x3e, 17, 0x3e, 18},
		},
		{ // rept / macro
			input: `t1 macro arg\ ld a, arg \ endm \ rept [16, 17, 18] \ t1 $i \ endr`,
			code:  []byte{0x3e, 0, 0x3e, 1, 0x3e, 2},
		},
		{
			input: `
			push_regs macro arg
			var @arg1 = $isarray(arg)
			if $isarray(arg)
				rept arg
				push $v
				endr
			else
				push arg
			endif
			endm

			push_regs hl
			push_regs [bc, de]
			push_regs af
			`,
			code: []byte{0xe5, 0xc5, 0xd5, 0xf5},
		},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New()
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
