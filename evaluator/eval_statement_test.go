package evaluator

import (
	"testing"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/internal/testutil"
	"github.com/dogatana/yas80/logging"
	"github.com/dogatana/yas80/object"
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
		logger := logging.New()
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
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
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

// エラーのみ正常系は evaluator_test.TestAssembleFile で確認
func TestProcAnonLabelError(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		{input: `
			name proc
				jp @f
			endp
			`,
			err: errcode.EANON_LABEL_NOT_FOUND,
		},
		{input: `
			name proc
				jp @1f
			endp
			`,
			err: errcode.EANON_LABEL_NOT_FOUND,
		},
		{input: `
			name proc
				jp @9f
			endp
			`,
			err: errcode.EANON_LABEL_NOT_FOUND,
		},
		{input: `
			name proc
				jp @b
			endp
			`,
			err: errcode.EANON_LABEL_NOT_FOUND,
		},
		{input: `
			name proc
				jp @2b
			endp
			`,
			err: errcode.EANON_LABEL_NOT_FOUND,
		},
		{input: `
			name proc
				jp @8b
			endp
			`,
			err: errcode.EANON_LABEL_NOT_FOUND,
		},
		{input: `
			name proc
				jp @@
			endp
			`,
			err: errcode.EANON_LABEL_DEF_ONLY,
		},
		{input: `
			name proc
				jp @1
			endp
			`,
			err: errcode.EANON_LABEL_DEF_ONLY,
		},
		{input: `
			name proc
			@f: nop
			endp
			`,
			err: errcode.EANON_LABEL_REF_ONLY,
		},
		{input: `
			name proc
			@1b: nop
			endp
			`,
			err: errcode.EANON_LABEL_REF_ONLY,
		},

		{input: `@@: nop`, err: errcode.ESCOPE_PROC},
		{input: `@f: nop`, err: errcode.ESCOPE_PROC},
		{input: `@b: nop`, err: errcode.ESCOPE_PROC},
		{input: `@1: nop`, err: errcode.ESCOPE_PROC},
		{input: `@9: nop`, err: errcode.ESCOPE_PROC},
		{input: `@2f: nop`, err: errcode.ESCOPE_PROC},
		{input: `@8b: nop`, err: errcode.ESCOPE_PROC},
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

func TestVariableStatement(t *testing.T) {
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
		{input: `var abc = def \ const def = 1`, err: errcode.EVAR_VALUE_FWD},
		{input: `const def = 1\ var abc = def`, syms: []symValue{{"ABC", 1}}},
		{input: `abc = 123`, err: errcode.EVAR_UNDEF},
		{input: `hl = 123`, err: errcode.EASSIGN_LEFT},
		// 10-
		{input: `fn func \ endf \ var abc = 1 \ abc = fn()`, err: errcode.EASSIGN_VALUE},
		{input: `fn func \ endf \ var abc = 1 \ abc = fn()`, err: errcode.EASSIGN_VALUE},
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

func TestIfStatement(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
	}{
		// 式文を除外したことで -1 は RESULT が未定義の意味に変更
		{`if 1 \ endif`, []symValue{{"RESULT", nil}}},
		{`if 0 \ endif`, []symValue{{"RESULT", nil}}},
		{`if 1 \ else \ endif`, []symValue{{"RESULT", nil}}},
		{`if 0 \ else \ endif`, []symValue{{"RESULT", nil}}},

		{`if 1 \ const result=100 \ endif`, []symValue{{"RESULT", 100}}},
		{`if 0 \ const result=100 \ endif`, []symValue{{"RESULT", nil}}},

		{`if 1 \ const result=100 \ else \ endif`, []symValue{{"RESULT", 100}}},
		{`if 0 \ const result=100 \ else \ endif`, []symValue{{"RESULT", nil}}},

		{`if 1 \ const result=100 \ else \ const result=200  \ endif`, []symValue{{"RESULT", 100}}},
		{`if 0 \ const result=100 \ else \ const result=200  \ endif`, []symValue{{"RESULT", 200}}},

		{`const val = 1 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200  \ endif`, []symValue{{"RESULT", 100}}},
		{`const val = 2 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200  \ endif`, []symValue{{"RESULT", 200}}},
		{`const val = 3 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200  \ endif`, []symValue{{"RESULT", nil}}},

		{`const val = 3 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200 \ else \ const result=300 \ endif`, []symValue{{"RESULT", 300}}},

		{`if '' \ const val = 1 \ else \ const val = 0 \ endif`, []symValue{{"VAL", 0}}},
		{`if HL \ const val = 1 \ else \ const val = 0 \ endif`, []symValue{{"VAL", 0}}},
		{`if CY \ const val = 1 \ else \ const val = 0 \ endif`, []symValue{{"VAL", 0}}},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, "", e)

		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter) // sym.Expected < 0 なので testSymValuesEx を使用する
	}
}

func TestOrgStatement(t *testing.T) {
	tests := []struct {
		input string
		addr  int
		err   string
	}{
		// 0-
		{input: `nop `, addr: 0},
		{input: `org $1000 \ nop `, addr: 0x1000},
		{input: `org $ffff \ nop `, addr: 0xffff},
		{input: `org -1 \ nop `, addr: 0xffff},
		{input: `org $ffff \ ld a, 1`, err: errcode.EADDR_OVERFLOW},
		{input: `org hl \ nop `, err: errcode.EORG_VALUE},
		{input: `org 'hl' \ nop `, err: errcode.EORG_VALUE},
		{input: `org abc \ nop `, err: errcode.ESYM_UNDEF},
		{input: `fn func\endf\ org fn()`, err: errcode.EORG_NULL},
		{input: `org 0, aaa`, err: errcode.EORG_ALLOC},
		// 10-
		{input: `org $ffff \ dsw 1`, err: errcode.EADDR_OVERFLOW},
		{input: `org $ffff \ ds 2`, err: errcode.EADDR_OVERFLOW},
		{input: `org $ffff \ db $ff, $ff`, err: errcode.EADDR_OVERFLOW},
		{input: `org $ffff \ dw $ffff`, err: errcode.EADDR_OVERFLOW},
		{input: `org $ffff \ dd $w(0)`, err: errcode.EADDR_OVERFLOW},
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

		if logger.ErrorCount() != 0 {
			logger.Print()

		}

		// code address
		var code *object.CodeObject
		for _, obj := range prog.Block {
			obj, ok := obj.(*object.CodeObject)
			if ok {
				code = obj
				break
			}
		}
		if code == nil {
			t.Errorf("[%d] not CodeObject. got %T", tn, prog.Block[0])
			continue
		}
		if code.Addr != tt.addr {
			t.Errorf("[%d] address is not $%04x. %s", tn, tt.addr, code.String())
		}
	}
}

func TestOrgAbsRel(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		{input: `a1:nop `, syms: []symValue{{"A1", 0}}},
		{input: `org $1000 \ a1:nop `, syms: []symValue{{"A1", 0x1000}}},
		{input: `org $ffff \ a1:nop `, syms: []symValue{{"A1", 0xffff}}},
		{input: `org -1 \ a1:nop `, syms: []symValue{{"A1", 0xffff}}},
		{
			input: `	org $1000
					a1r	equ $$
					a1: nop

					org $2000
					a2r	equ $$
					a2: nop

					org $3000, rel
					a3r	equ $$
					a3: nop

					org $4000, rel
					a4r	equ $$
					a4: nop

					org $5000
					a5r	equ $$
					a5: nop
			`, syms: []symValue{
				{"A1", 0x1000},
				{"A2", 0x2000},
				{"A3", 0x3000},
				{"A4", 0x4000},
				{"A5", 0x5000},
				{"A1R", 0x1000},
				{"A2R", 0x2000},
				{"A3R", 0x2001},
				{"A4R", 0x2002},
				{"A5R", 0x5000},
			},
		},
		{input: `org $ffff \ ld a, 1`, err: errcode.EADDR_OVERFLOW},
		// {input: `org 0xffff \ org 0, rel \ call 0`, err: errcode.EALLOC_ADDR_OVERFLOW}, // 配置アドレスは64k超を許容するためテスト対象外
		{input: `org hl \ nop `, err: errcode.EORG_VALUE},
		{input: `org 'hl' \ nop `, err: errcode.EORG_VALUE},
		{input: `org abc \ nop `, err: errcode.ESYM_UNDEF},
		{input: `fn func\endf\ org fn()`, err: errcode.EORG_NULL},
		{input: `org 0, aaa`, err: errcode.EORG_ALLOC},
		// 10-
		{input: `org $ffff \ dsw 1`, err: errcode.EADDR_OVERFLOW},
		{input: `org $ffff \ ds 2`, err: errcode.EADDR_OVERFLOW},
		{input: `org $ffff \ db $ff, $ff`, err: errcode.EADDR_OVERFLOW},
		{input: `org $ffff \ dw $ffff`, err: errcode.EADDR_OVERFLOW},
		{input: `org $ffff \ dd $w(0)`, err: errcode.EADDR_OVERFLOW},
		// 0-
		{input: `var v = 1 \ v = v + 1`, syms: []symValue{{"V", 2}}},
		{input: `var .test = 1`, err: errcode.ESCOPE_PROC},
		{input: `var @test = 1`, err: errcode.ESCOPE_MACRO},
		{input: `var _ = 1`, err: errcode.EVAR_SYS},
		{input: `const abc = 1 \ var abc = 1`, err: errcode.EVAR_USED},
		// 5-
		{input: `var abc = 1 \ var abc = 2`, err: errcode.EVAR_USED},
		{input: `var abc = def \ const def = 1`, err: errcode.EVAR_VALUE_FWD},
		{input: `const def = 1\ var abc = def`, syms: []symValue{{"ABC", 1}}},
		{input: `abc = 123`, err: errcode.EVAR_UNDEF},
		{input: `hl = 123`, err: errcode.EASSIGN_LEFT},
		// 10-
		{input: `fn func \ endf \ var abc = 1 \ abc = fn()`, err: errcode.EASSIGN_VALUE},
		{input: `fn func \ endf \ var abc = 1 \ abc = fn()`, err: errcode.EASSIGN_VALUE},
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
