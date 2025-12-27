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
	}{
		{"const abc ## 123 = 456",
			[]SymValue{{"ABC123", 456}}},
		{"const abc ## (100 + 23) = 456",
			[]SymValue{{"ABC123", 456}}},
		{"const abc ## (100 + 23) = def ## 456 \\ const def ## (400 + 56) = 999",
			[]SymValue{{"ABC123", 999}}},
		{"abc ## 123 equ 456",
			[]SymValue{{"ABC123", 456}}},
		{"abc ## (100 + 23) equ 456",
			[]SymValue{{"ABC123", 456}}},
		{"abc ## (100 + 23) equ def ## 456 \\ const def ## (400 + 56) = 999",
			[]SymValue{{"ABC123", 999}}},
	}
	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, "", e)

		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestProcStatement(t *testing.T) {
	tests := []struct {
		input string
		syms  []SymValue
	}{
		{`test proc \ const abc ## (100 + 23) = 456 \ endp`,
			[]SymValue{{"ABC123", 456}}},
		{`test proc \ const abc ## (100 + 23) = def ## 456 \ const def ## (400 + 56) = 999 \ endp`,
			[]SymValue{{"ABC123", 999}}},
		{`test proc \ abc ## 123 equ 456 \ endp`,
			[]SymValue{{"ABC123", 456}}},
		{`test proc \ abc ## (100 + 23) equ 456 \ endp`,
			[]SymValue{{`ABC123`, 456}}},
		{`test proc \ abc ## (100 + 23) equ def ## 456 \ const def ## (400 + 56) = 999  \ endp`,
			[]SymValue{{"ABC123", 999}}},
	}
	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, "", e)

		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestProcLabel(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []SymValue
		err   string
	}{
		// 0-
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
		// 5-
		{input: `const abc = 1 \ abc proc \ endp`, err: errcode.EPROC_USED},
		{},
		{},
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
