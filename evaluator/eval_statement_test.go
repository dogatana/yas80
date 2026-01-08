package evaluator

import (
	"testing"
	"yas80/errcode"
	"yas80/logging"
	"yas80/object"
)

func TestLabelStatement(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
	}{
		{`addr1: ld hl, $1234 \ addr2: ld a, a \ addr3: ld hl, $5678`,
			[]byte{0x21, 0x34, 0x12, 0x7f, 0x21, 0x78, 0x56},
			[]symValue{
				{"ADDR1", 0},
				{"ADDR2", 3},
				{"ADDR3", 4},
			},
		},
		{`addr1: ld a, a\ addr2: ld hl, $1234 \ addr3: ld hl, $5678`,
			[]byte{0x7f, 0x21, 0x34, 0x12, 0x21, 0x78, 0x56},
			[]symValue{
				{"ADDR1", 0},
				{"ADDR2", 1},
				{"ADDR3", 4},
			},
		},
		{`addr1: ld hl, $1234 \ addr2: ld hl, $5678 \ addr3: ld hl, $9abc`,
			[]byte{0x21, 0x34, 0x12, 0x21, 0x78, 0x56, 0x21, 0xbc, 0x9a},
			[]symValue{
				{"ADDR1", 0},
				{"ADDR2", 3},
				{"ADDR3", 6},
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
		testEvalResult(t, tn, "", e)

		testCodeResult(t, tn, tt.code, prog)

		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestConstStatement(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `const abc = 123`, syms: []symValue{{"ABC", 123}}},
		{input: `const abc = 1 + 2 * 3`, syms: []symValue{{"ABC", 7}}},
		{input: `const abc = xyz + 1 \ const xyz = 9`, syms: []symValue{{"ABC", 10}, {"XYZ", 9}}},
		{input: `const abc = xyz`, err: errcode.ESYM_UNDEF},
		{input: `const abc = xyz`, err: errcode.ESYM_NULL},
	}
	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testLogMessage(t, tn, tt.err, e.logger)
			continue
		}

		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestProcStatement(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0- 単純ケース
		{input: `abc proc \ nop \ .ret: ret \ nop \ const .xyz = 123 \ endp`,
			code: []byte{0, 0xc9, 0},
			syms: []symValue{{"ABC.RET", 1}, {"ABC.XYZ", 123}},
		},
		{},
		{},
		{},
		{},
		// 5-
		{`abc proc \ ld a, .abc \ .abc: nop \ endp \ xyz proc \ ld a, .abc \ .abc: ret \ endp`,
			[]byte{0x3e, 0x02, 0x00, 0x3e, 0x05, 0xc9}, // ld a,2 \ nop \ ld a,5 \ ret
			[]symValue{
				{"ABC.ABC", 2},
				{"XYZ.ABC", 5},
			},
			"",
		},
		{`abc proc \ ld a, .abc \ nop \ const .abc = 0xa5 \ endp \ xyz proc \ ld a, .abc \ .abc: ret \ endp`,
			[]byte{0x3e, 0xa5, 0x00, 0x3e, 0x05, 0xc9}, // ld a,2 \ nop \ ld a,5 \ ret
			[]symValue{
				{"ABC.ABC", 0xa5},
				{"XYZ.ABC", 5},
			},
			"",
		},
		{`abc proc \ const .xxx = zzz + 1 \nop \ endp \ const zzz = 1 \ ld a, abc.xxx`,
			[]byte{0, 0x3e, 0x02},
			[]symValue{
				{"ABC.XXX", 2},
				{"ZZZ", 1},
			},
			"",
		},
		{`ld a, abc.def \ abc proc \ nop \ const .def = 4 \ nop \ endp`,
			[]byte{0x3e, 0x04, 0, 0},
			[]symValue{
				{"ABC.DEF", 4},
			},
			"",
		},
		{`ld a, abc.def \ abc proc \ nop \ .def: ret \ nop \ endp`,
			[]byte{0x3e, 0x03, 0, 0xc9, 0},
			[]symValue{
				{"ABC.DEF", 3},
			},
			"",
		},

		// 10- error case
		{input: `abc proc \ nop \ endp \ abc proc \ ret \ endp`, err: errcode.EPROC_DUP},
		{input: `const abc = 1 \ abc proc \ ret \ endp`, err: errcode.EPROC_USED},
		{input: `abc: nop \ abc proc \ ret \ endp`, err: errcode.EPROC_USED},
		{input: `abc proc \ nop \ endp \ ld a, abc.xxx`, err: errcode.ESYM_UNDEF},
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
			testLogMessage(t, tn, tt.err, e.logger)
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

func TestEnumStatement(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `test enum \ aaa = 123 \ bbb = "value" \ ccc \ ende`,
			syms: []symValue{
				{"TEST.AAA", 123},
				{"TEST.BBB", "value"},
				{"TEST.CCC", 124},
				{"TEST.DDD", nil},
			},
		},
		{input: `test enum \ aaa = 1 \ bbb = 10 \ ccc = 100\ ende`,
			syms: []symValue{
				{"TEST.AAA", 1},
				{"TEST.BBB", 10},
				{"TEST.CCC", 100},
			},
		},
		{input: `test enum \ aaa = 1 \ bbb \ ccc\ ende`,
			syms: []symValue{
				{"TEST.AAA", 1},
				{"TEST.BBB", 2},
				{"TEST.CCC", 3},
			},
		},
		{input: `test enum \ aaa = 1 \ bbb = .aaa * 2\ ccc = .bbb * 3\ ende`,
			syms: []symValue{
				{"TEST.AAA", 1},
				{"TEST.BBB", 2},
				{"TEST.CCC", 6},
			},
		},
		{input: `test enum \ aaa \ ende \ test enum \ aaa \ende`, err: errcode.EENUM_DUP},
		{input: `const test = 1 \ test enum \ aaa \ ende`, err: errcode.EENUM_USED},
		{input: `test: nop \ test enum \ aaa \ ende`, err: errcode.EENUM_USED},
		{input: `test enum \ aaa \ aaa \ ende`, err: errcode.EENUM_ELE_DUP},
		{input: `test enum \ aaa = hl \ ende`, err: errcode.EENUM_ELE_VALUE},
		{input: `test enum \ aaa = .bbb \ bbb = 1 \ ende`, err: errcode.EENUM_ELE_FWD},
		{input: `test enum \ aaa = outer_value \ ende`, err: errcode.EENUM_ELE_FWD},
		{input: `abc enum \ aaa = 1 \ ende \ ld a, abc.xyz`, err: errcode.EENUM_ELE_UNDEF},
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
			testLogMessage(t, tn, tt.err, e.logger)
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

func TestVarStatement(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `var v = 1 \ v = v + 1`, syms: []symValue{{"V", 2}}},
		{input: `var .test = 1`, err: errcode.ESCOPE_PROC},
		{input: `var @test = 1`, err: errcode.ESCOPE_MACRO},
		{input: `var _ = 1`, err: errcode.EVAR_SYS},
		{input: `const abc = 1 \ var abc = 1`, err: errcode.EVAR_USED},
		// 5-
		{input: `var abc = 1 \ var abc = 2`, err: errcode.EVAR_USED},
		{input: `var abc = def \ const def = 1`, err: errcode.EVAR_VALUE},
		{input: `const def = 1\ var abc = def`, syms: []symValue{{"ABC", 1}}},
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
			testLogMessage(t, tn, tt.err, e.logger)
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
