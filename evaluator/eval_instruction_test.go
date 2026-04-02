package evaluator

import (
	"testing"
	"yas80/errcode"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
)

func TestZ80Instruction(t *testing.T) {
	tests := []string{
		"inst0",
		"inst-ld8",
		"inst-ld16",
		"inst-ldind",
		"inst-add8",
		"inst-add8-alt", // or a,a のように A レジスタを指定したもの
		"inst-add16",
		"inst-inc-push",
		"inst-ex-im",
		"inst-rlc",
		"inst-bit",
		"inst-call-ret",
		"inst-io",
		"zilog",
	}

	for tn, base := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		input := string(testutil.ReadTestDataFile(t, base+".asm"))
		expected := testutil.ReadTestDataFile(t, base+".bin")

		prog, e := evalInput(input, logger, env)
		testEvalResult(t, tn, "", e)

		logger.Print()
		result := object.CollectCode(prog.Block)

		if err := testutil.BytesEqual(result, expected); err != nil {
			t.Errorf("[%d] generated code diff %s", tn, err.Error())
		}
	}
}

func TestInstructionAmbiguous(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `ex hl, de`, code: []byte{0xeb}},
		{input: `ex af', af`, code: []byte{0x08}},
		{input: `ex hl, (sp)`, code: []byte{0xe3}},
		{input: `ex ix, (sp)`, code: []byte{0xdd, 0xe3}},
		{input: `ex iy, (sp)`, code: []byte{0xfd, 0xe3}},
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
	}
}

func TestInstruction_LD_STR(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `ld a, 'a'`, code: []byte{0x3e, 'a'}},
		{input: `ld a, 'あ'`, code: []byte{0x3e, 0xa0}, err: errcode.WROUND_BYTE},
		{input: `ld hl, 'a'`, code: []byte{0x21, 'a', 0}},
		{input: `ld hl, 'あ'`, code: []byte{0x21, 0xa0, 0x82}},
		{input: `ld (hl), 'a'`, code: []byte{0x36, 0x61}},
		{input: `ld (ix), 'b'`, code: []byte{0xdd, 0x36, 0, 0x62}},
		{input: `ld (iy+1), 'c'`, code: []byte{0xfd, 0x36, 1, 0x63}},
		{input: `ld a, 'ab'`, err: errcode.ESTR_TO_INT_LEN},
		{input: `ld hl, 'ab'`, err: errcode.ESTR_TO_INT_LEN},
		{input: `ld hl, 'abc'`, err: errcode.ESTR_TO_INT_LEN},
		{input: `ld hl, 'あc'`, err: errcode.ESTR_TO_INT_LEN},
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
	}
}

func TestInstruction_LD_Array(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `ld a, [1]`, code: []byte{0x3e, 1}},
		{input: `ld hl, [0x82, 0xa0]`, code: []byte{0x21, 0xa0, 0x82}},
		{input: `ld (hl), [2]`, code: []byte{0x36, 2}},
		{input: `ld (ix), [3]`, code: []byte{0xdd, 0x36, 0, 03}},
		{input: `ld a, []`, err: errcode.EARRAY_TO_INT_LEN},
		{input: `ld a, [1,2,3]`, err: errcode.EARRAY_TO_INT_LEN},
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
	}
}

func TestInstructionError_LD16(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `ld`, err: errcode.ESYNTAX},
		// {input: `ld hl, 'a'`, err: errcode.EZ80_OP2}, // 有効オペランド
		{input: `ld hl, cy`, err: errcode.EZ80_OP2},
		{input: `fn func\endf\ld hl, fn()`, err: errcode.EZ80_OP2_NULL},
		{input: `ld hl, sp`, err: errcode.EZ80_OP_REG},
		{input: `ld ix, sp`, err: errcode.EZ80_OP_REG},
		{input: `ld sp, de`, err: errcode.EZ80_OP_REG},
		{input: `ld sp, bc`, err: errcode.EZ80_OP_REG},
		{input: `ld hl, 65536`, err: errcode.WROUND_WORD},
		{input: `ld hl, -32769`, err: errcode.WROUND_WORD},
		{input: `ld hl, (65536)`, err: errcode.WROUND_WORD},
		{input: `ld hl, (-32769)`, err: errcode.WROUND_WORD},
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
	}
}

func TestInstructionError_LD8(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `ld`, err: errcode.ESYNTAX},
		// {input: `ld a, 'a'`, err: errcode.EZ80_OP2}, // 有効オペランド
		{input: `ld a, cy`, err: errcode.EZ80_OP2},
		{input: `fn func\endf\ld a, fn()`, err: errcode.EZ80_OP2_NULL},
		{input: `ld i, i`, err: errcode.EZ80_OP_REG},
		{input: `ld r, r`, err: errcode.EZ80_OP_REG},
		{input: `ld i, b`, err: errcode.EZ80_OP_REG},
		{input: `ld r, b`, err: errcode.EZ80_OP_REG},
		{input: `ld c, i`, err: errcode.EZ80_OP_REG},
		{input: `ld c, r`, err: errcode.EZ80_OP_REG},
		{input: `ld i, 1`, err: errcode.EZ80_OP_REG},
		{input: `ld r, 1`, err: errcode.EZ80_OP_REG},
		{input: `ld a, -129`, err: errcode.WROUND_BYTE},
		{input: `ld a, 256`, err: errcode.WROUND_BYTE},
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
	}
}

func TestInstructionError_LDRegIndirect(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `ld (hl)`, err: errcode.EZ80_OP},
		{input: `fn func\endf\ld (hl), fn()`, err: errcode.EZ80_OP2_NULL},
		{input: `ld (hl), i`, err: errcode.EZ80_OP_REG},
		{input: `ld (hl), r`, err: errcode.EZ80_OP_REG},
		{input: `ld (sp), a`, err: errcode.EZ80_OP_REG},
		{input: `ld (sp), 1`, err: errcode.EZ80_OP},
		// {input: `ld (hl), 'a'`, err: errcode.EZ80_OP2},
		{input: `ld (hl), cy`, err: errcode.EZ80_OP2},
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
	}
}

func TestInstructionError_LDAddrIndirect(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `ld (0)`, err: errcode.EZ80_OP},
		{input: `fn func\endf\ld (0), fn()`, err: errcode.EZ80_OP2_NULL},
		{input: `ld (0), i`, err: errcode.EZ80_OP_REG},
		{input: `ld (0), r`, err: errcode.EZ80_OP_REG},
		{input: `ld (0), 1`, err: errcode.EZ80_OP2},
		{input: `ld (0), 'a'`, err: errcode.EZ80_OP2},
		{input: `ld (0), cy`, err: errcode.EZ80_OP2},
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
	}
}

func testInstructionAdd8(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		{input: `
add a, 'A'
adc a, 'A'
sub    'A'
sbc a, 'A'
and    'A'
or     'A'
xor    'A'
cp     'A'
add a, '@' + 1
adc a, '@' + 1
sub    '@' + 1
sbc a, '@' + 1
and    '@' + 1
or     '@' + 1
xor    '@' + 1
cp     '@' + 1
add a, [0x41]
adc a, [0x41]
sub    [0x41]
sbc a, [0x41]
and    [0x41]
or     [0x41]
xor    [0x41]
cp     [0x41]
add a, [0, 0x41]
adc a, [0, 0x41]
sub    [0, 0x41]
sbc a, [0, 0x41]
and    [0, 0x41]
or     [0, 0x41]
xor    [0, 0x41]
cp     [0, 0x41]`, code: []byte{
			0xc6, 0x41, 0xce, 0x41, 0xd6, 0x41, 0xde, 0x41, 0xe6, 0x41, 0xf6, 0x41, 0xee, 0x41, 0xfe, 0x41, // 'A'
			0xc6, 0x41, 0xce, 0x41, 0xd6, 0x41, 0xde, 0x41, 0xe6, 0x41, 0xf6, 0x41, 0xee, 0x41, 0xfe, 0x41, // '@' + 1
			0xc6, 0x41, 0xce, 0x41, 0xd6, 0x41, 0xde, 0x41, 0xe6, 0x41, 0xf6, 0x41, 0xee, 0x41, 0xfe, 0x41, // [0x41]
			0xc6, 0x41, 0xce, 0x41, 0xd6, 0x41, 0xde, 0x41, 0xe6, 0x41, 0xf6, 0x41, 0xee, 0x41, 0xfe, 0x41, // [0, 0x41]
		}},
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
	}
}

func TestInstructionError_JP_JR_DJNZ(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// {input: `jp -1`, err: errcode.WROUND_ADDR},
		{input: `jp hl`, err: errcode.EZ80_OP},
		{input: `jp c,hl`, err: errcode.EZ80_OP2},
		{input: `jp a,$1234`, err: errcode.EZ80_FLAG},
		{input: `jr c,hl`, err: errcode.EZ80_OP2},
		{input: `jr a,$12`, err: errcode.EZ80_FLAG},
		{input: `jr $1000`, err: errcode.EZ80_JR_RANGE},
		{input: `jr c,hl`, err: errcode.EZ80_OP2},
		{input: `jr a,$12`, err: errcode.EZ80_FLAG},
		{input: `jr po,$12`, err: errcode.EZ80_JR_FLAG},
		{input: `djnz hl`, err: errcode.EZ80_OP},
		{input: `djnz $1000`, err: errcode.EZ80_JR_RANGE},
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

		// sym
		obj, ok := env.Get("VAL")
		if !ok {
			t.Errorf("[%d] VAL not in env", tn)
			continue
		}
		sym, ok := obj.(*object.SymbolObject)
		if !ok {
			t.Errorf("[%d] env[\"VAL\" not SymbolObject", tn)
			continue
		}
		if sym.SymType != object.SYM_UNKNOWN {
			t.Errorf("[%d] SymType not SYM_UNNOWN. got %d", tn, sym.SymType)
		}
	}
}
func TestInstructionError_CALL_RET_RST(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// {input: `call -1`, err: errcode.WROUND_ADDR}, // intToAddr から intToWord に変更のため
		{input: `call hl`, err: errcode.EZ80_OP},
		{input: `call hl,1234`, err: errcode.EZ80_FLAG},
		{input: `call 123,1234`, err: errcode.EZ80_FLAG},
		{input: `ret VAL`, err: errcode.ESYM_UNDEF},
		{input: `ret hl`, err: errcode.EZ80_FLAG},
		{input: `ret 123`, err: errcode.EZ80_FLAG},
		{input: `rst`, err: errcode.EZ80_OP},
		{input: `rst 1`, err: errcode.EZ80_RST},
		{input: `rst 40h`, err: errcode.EZ80_RST},
		{input: `rst hl`, err: errcode.EZ80_OP},
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
	}
}

func TestInstructionError_IN_OUT(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `in hl, (c)`, err: errcode.EZ80_OP_REG},
		{input: `in i, (c)`, err: errcode.EZ80_OP_REG},
		{input: `in r, (c)`, err: errcode.EZ80_OP_REG},
		{input: `in 123, (c)`, err: errcode.EZ80_OP1},
		{input: `in a, (hl)`, err: errcode.EINDIRECT_REG},
		{input: `fn func \ return \ endf \ in fn(), (hl)`, err: errcode.EZ80_OP1_NULL},
		{input: `in a, (-1)`, err: errcode.EZ80_PORT_RANGE},
		{input: `in a, (256)`, err: errcode.EZ80_PORT_RANGE},
		{input: `fn func \ return \ endf \ in a, (fn())`, err: errcode.EINDIRECT_NULL},
		{input: `in i, (0)`, err: errcode.EZ80_OP_REG},
		{input: `in r, (0)`, err: errcode.EZ80_OP_REG},

		{input: `out (c), hl`, err: errcode.EZ80_OP_REG},
		{input: `out (c), i`, err: errcode.EZ80_OP_REG},
		{input: `out (c), r`, err: errcode.EZ80_OP_REG},
		{input: `out (c), 123`, err: errcode.EZ80_OP1},
		{input: `out (hl), a`, err: errcode.EINDIRECT_REG},
		{input: `fn func \ return \ endf \ out (hl), fn()`, err: errcode.EZ80_OP2_NULL},
		{input: `out (-1), a`, err: errcode.EZ80_PORT_RANGE},
		{input: `out (256), a`, err: errcode.EZ80_PORT_RANGE},
		{input: `fn func \ return \ endf \ out (fn()), a`, err: errcode.EINDIRECT_NULL},
		{input: `out (0), i`, err: errcode.EZ80_OP_REG},
		{input: `out (0), r`, err: errcode.EZ80_OP_REG},
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
	}
}

func TestInstructionError_BIT_SET_RES(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `bit 10, a`, err: errcode.EZ80_BIT_NUM_RANGE},
		{input: `set hl, a`, err: errcode.EZ80_OP1},
		{input: `fn func \ return \ endf \ res fn(), a`, err: errcode.EZ80_OP1_NULL},
		{input: `fn func \ return \ endf \ set 0, fn()`, err: errcode.EZ80_OP2_NULL},
		{input: `bit 0, i`, err: errcode.EZ80_OP_REG},
		{input: `bit 0, r`, err: errcode.EZ80_OP_REG},
		{input: `bit 0, (de)`, err: errcode.EINDIRECT_REG},
		{input: `bit 0, (c)`, err: errcode.EINDIRECT_REG},
		{input: `set 0, hl`, err: errcode.EZ80_OP_REG},
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
	}
}

func TestInstructionError_RLC(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `fn func\endf\ rlc fn()`, err: errcode.EZ80_OP_NULL},
		{input: `rlc i`, err: errcode.EZ80_OP_REG},
		{input: `rlc r`, err: errcode.EZ80_OP_REG},
		{input: `rlc (de)`, err: errcode.EINDIRECT_REG},
		{input: `rlc 1`, err: errcode.EZ80_OP},
		{input: `rlc 'a'`, err: errcode.EZ80_OP},
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
	}
}

func TestInstructionError_IM(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `im`, err: errcode.EZ80_OP},
		{input: `im hl`, err: errcode.EZ80_OP},
		{input: `im 'a'`, err: errcode.EZ80_OP},
		{input: `im -1`, err: errcode.EZ80_IM_RANGE},
		{input: `im 3`, err: errcode.EZ80_IM_RANGE},
		{input: `fn func\endf\ im fn()`, err: errcode.EZ80_OP_NULL},
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
	}
}

func TestInstructionError_EX(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `ex`, err: errcode.ESYNTAX},
		{input: `ex hl`, err: errcode.EZ80_OP},
		{input: `ex hl, 1`, err: errcode.EZ80_OP},
		{input: `ex hl, sp`, err: errcode.EZ80_OP},
		{input: `ex ix, iy`, err: errcode.EZ80_OP},
		{input: `ex hl, iy`, err: errcode.EZ80_OP},
		{input: `ex (hl), iy`, err: errcode.EINDIRECT_REG},
		{input: `ex af, (sp), `, err: errcode.EZ80_OP_REG},
		{input: `ex (sp), (hl), `, err: errcode.EZ80_OP},
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
	}
}

func TestInstructionError_INC_DEC(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `inc`, err: errcode.EZ80_OP},
		{input: `inc 1`, err: errcode.EZ80_OP},
		{input: `inc 'a'`, err: errcode.EZ80_OP},
		{input: `fn func\endf \ inc fn()`, err: errcode.EZ80_OP_NULL},
		{input: `inc i`, err: errcode.EZ80_OP_REG},
		{input: `dec r`, err: errcode.EZ80_OP_REG},
		{input: `inc (hl + 1)`, err: errcode.EINDIRECT_DISP_REG},
		{input: `inc (ix + 128)`, err: errcode.EINDIRECT_DISP_RANGE},
		{input: `dec (iy - 129)`, err: errcode.EINDIRECT_DISP_RANGE},
		{input: `dec (af)`, err: errcode.EINDIRECT_REG},
		{input: `dec (a)`, err: errcode.EINDIRECT_REG},
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
	}
}

func TestInstructionError_PUSH_POP(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `push`, err: errcode.EZ80_OP},
		{input: `push 1`, err: errcode.EZ80_OP},
		{input: `push 'a'`, err: errcode.EZ80_OP},
		{input: `push a`, err: errcode.EZ80_OP_REG},
		{input: `fn func\endf\ push fn()`, err: errcode.EZ80_OP_NULL},
		{input: `push sp`, err: errcode.EZ80_OP_REG},
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
	}
}

func TestInstructionError_ADD8(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `and`, err: errcode.ESYNTAX},
		{input: `fn func\endf\ or fn(), a`, err: errcode.EZ80_OP1_NULL},
		{input: `xor 'a', a`, err: errcode.EZ80_OP1},
		{input: `sub i, a`, err: errcode.EZ80_OP1_REG_A},
		{input: `sub r, a`, err: errcode.EZ80_OP1_REG_A},
		{input: `sub 1, a`, err: errcode.EZ80_OP1},
		{input: `fn func\endf\ or a, fn()`, err: errcode.EZ80_OP2_NULL},
		{input: `fn func\endf\ or fn()`, err: errcode.EZ80_OP_NULL},
		{input: `cp i`, err: errcode.EZ80_OP_REG},
		{input: `and -129`, err: errcode.WROUND_BYTE},
		{input: `and 256`, err: errcode.WROUND_BYTE},
		{input: `and (DE)`, err: errcode.EINDIRECT_REG},
		{input: `and (IX+128)`, err: errcode.EINDIRECT_DISP_RANGE},
		{input: `and (IY-129)`, err: errcode.EINDIRECT_DISP_RANGE},
		{input: `and 'abc'`, err: errcode.ESTR_TO_INT_LEN},
		{input: `and ''`, err: errcode.ESTR_TO_INT_LEN},
		{input: `and []`, err: errcode.EARRAY_TO_INT_LEN},
		{input: `and [1,2,3]`, err: errcode.EARRAY_TO_INT_LEN},
		{input: `and ['a']`, err: errcode.EARRAY_TO_INT_TYPE},
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
	}
}

func TestInstructionError_ADD16(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `add`, err: errcode.ESYNTAX},
		{input: `fn func\endf\ add fn(), a`, err: errcode.EZ80_OP1_NULL},
		{input: `add 'a', a`, err: errcode.EZ80_OP1},
		{input: `add bc, hl`, err: errcode.EZ80_OP1_REG_HL_IXY},
		{input: `adc bc, hl`, err: errcode.EZ80_OP1_REG_HL},
		{input: `sbc bc, hl`, err: errcode.EZ80_OP1_REG_HL},
		{input: `fn func\endf\ add hl,fn()`, err: errcode.EZ80_OP2_NULL},
		{input: `add hl, af`, err: errcode.EZ80_OP_REG},
		{input: `adc hl, 1`, err: errcode.EZ80_OP2},
		{input: `adc hl, 'a'`, err: errcode.EZ80_OP2},
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
	}
}

func TestInstruction_MUL(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `mul a, b`, code: []byte{0xed, 0xc1}},
		{input: `mul a, c`, code: []byte{0xed, 0xc9}},
		{input: `mul a, d`, code: []byte{0xed, 0xd1}},
		{input: `mul a, e`, code: []byte{0xed, 0xd9}},
		{input: `mul a, h`, err: errcode.EZ80_OP_REG},
		{input: `mul l, h`, err: errcode.EZ80_OP_REG},

		{input: `mul hl, bc`, code: []byte{0xed, 0xc3}},
		{input: `mul hl, sp`, code: []byte{0xed, 0xf3}},
		{input: `mul hl, de`, err: errcode.EZ80_OP_REG},
		{input: `mul bc, sp`, err: errcode.EZ80_OP_REG},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		env.Set("$R800", &object.NumberObject{Value: 1}) // -R800 有効化

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
	}
}

func TestInstruction_MUL_on_Z80(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `mul a, b`, err: errcode.ER800},
		{input: `mul a, c`, err: errcode.ER800},
		{input: `mul a, d`, err: errcode.ER800},
		{input: `mul a, e`, err: errcode.ER800},
		{input: `mul a, h`, err: errcode.ER800},
		{input: `mul l, h`, err: errcode.ER800},

		{input: `mul hl, bc`, err: errcode.ER800},
		{input: `mul hl, sp`, err: errcode.ER800},
		{input: `mul hl, de`, err: errcode.ER800},
		{input: `mul bc, sp`, err: errcode.ER800},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		env.Set("$R800", &object.NumberObject{Value: 0}) // -R800 無効化

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
	}
}
