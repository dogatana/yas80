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
		syms  []SymValue
	}{
		{`addr1: ld hl, $1234 \ addr2: ld a, a \ addr3: ld hl, $5678`,
			[]byte{0x21, 0x34, 0x12, 0x7f, 0x21, 0x78, 0x56},
			[]SymValue{
				{"ADDR1", 0},
				{"ADDR2", 3},
				{"ADDR3", 4},
			},
		},
		{`addr1: ld a, a\ addr2: ld hl, $1234 \ addr3: ld hl, $5678`,
			[]byte{0x7f, 0x21, 0x34, 0x12, 0x21, 0x78, 0x56},
			[]SymValue{
				{"ADDR1", 0},
				{"ADDR2", 1},
				{"ADDR3", 4},
			},
		},
		{`addr1: ld hl, $1234 \ addr2: ld hl, $5678 \ addr3: ld hl, $9abc`,
			[]byte{0x21, 0x34, 0x12, 0x21, 0x78, 0x56, 0x21, 0xbc, 0x9a},
			[]SymValue{
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
		syms  []SymValue
		err   string
	}{
		// 0-
		{input: `const abc = 123`, syms: []SymValue{{"ABC", 123}}},
		{input: `const abc = 1 + 2 * 3`, syms: []SymValue{{"ABC", 7}}},
		{input: `const abc = xyz + 1 \ const xyz = 9`, syms: []SymValue{{"ABC", 10}, {"XYZ", 9}}},
		{input: `const abc = xyz`, err: errcode.ESYM_UNDEF},
		{input: `const abc = xyz`, err: errcode.ESYM_NULL},

		// 5-
		{input: "const abc ## 123 = 456", syms: []SymValue{{"ABC123", 456}}},
		{input: "const abc ## (100 + 23) = 456", syms: []SymValue{{"ABC123", 456}}},
		{input: "const abc ## (100 + 23) = def ## 456 \\ const def ## (400 + 56) = 999", syms: []SymValue{{"ABC123", 999}}},
		{input: `const abc ## "_XYZ" = 123`, syms: []SymValue{{"ABC_XYZ", 123}}},
		{input: `const abc ## "_xyz" = 123`, syms: []SymValue{{"ABC_XYZ", 123}}},

		// 10-
		{input: "abc ## (100 + 23) equ def ## 456 \\ const def ## (400 + 56) = 999", syms: []SymValue{{"ABC123", 999}}},
		{input: `const abc ## a = 1`, err: errcode.ESYM_CONCAT_TYPE},
		{input: `const abc ## hl = 1`, err: errcode.ESYM_CONCAT_TYPE},
		{input: `const abc ## cy = 1`, err: errcode.ESYM_CONCAT_TYPE},
		{input: `const abc ## cy ## 1 = 1`, err: errcode.ESYNTAX}, // syntax error
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
		syms  []SymValue
		err   string
	}{
		// 0- 単純ケース
		{input: `abc proc \ nop \ .ret: ret \ nop \ const .xyz = 123 \ endp`,
			code: []byte{0, 0xc9, 0},
			syms: []SymValue{{"ABC.RET", 1}, {"ABC.XYZ", 123}},
		},
		{},
		{},
		{},
		{},
		// 5-
		{`abc proc \ ld a, .abc \ .abc: nop \ endp \ xyz proc \ ld a, .abc \ .abc: ret \ endp`,
			[]byte{0x3e, 0x02, 0x00, 0x3e, 0x05, 0xc9}, // ld a,2 \ nop \ ld a,5 \ ret
			[]SymValue{
				{"ABC.ABC", 2},
				{"XYZ.ABC", 5},
			},
			"",
		},
		{`abc proc \ ld a, .abc \ nop \ const .abc = 0xa5 \ endp \ xyz proc \ ld a, .abc \ .abc: ret \ endp`,
			[]byte{0x3e, 0xa5, 0x00, 0x3e, 0x05, 0xc9}, // ld a,2 \ nop \ ld a,5 \ ret
			[]SymValue{
				{"ABC.ABC", 0xa5},
				{"XYZ.ABC", 5},
			},
			"",
		},
		{`abc proc \ const .xxx = zzz + 1 \nop \ endp \ const zzz = 1 \ ld a, abc.xxx`,
			[]byte{0, 0x3e, 0x02},
			[]SymValue{
				{"ABC.XXX", 2},
				{"ZZZ", 1},
			},
			"",
		},
		{`ld a, abc.def \ abc proc \ nop \ const .def = 4 \ nop \ endp`,
			[]byte{0x3e, 0x04, 0, 0},
			[]SymValue{
				{"ABC.DEF", 4},
			},
			"",
		},
		{`ld a, abc.def \ abc proc \ nop \ .def: ret \ nop \ endp`,
			[]byte{0x3e, 0x03, 0, 0xc9, 0},
			[]SymValue{
				{"ABC.DEF", 3},
			},
			"",
		},

		// 10- シンボル結合
		{input: `test proc \ const abc ## (100 + 23) = 456 \ endp`, syms: []SymValue{{"ABC123", 456}}},
		{input: `test proc \ const abc ## (100 + 23) = def ## 456 \ const def ## (400 + 56) = 999 \ endp`, syms: []SymValue{{"ABC123", 999}}},
		{input: `test proc \ abc ## 123 equ 456 \ endp`, syms: []SymValue{{"ABC123", 456}}},
		{input: `test proc \ abc ## (100 + 23) equ 456 \ endp`, syms: []SymValue{{`ABC123`, 456}}},
		{input: `test proc \ abc ## (100 + 23) equ def ## 456 \ const def ## (400 + 56) = 999  \ endp`, syms: []SymValue{{"ABC123", 999}}},

		// 15- error case
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
