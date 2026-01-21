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
		"ld_r8_r8",
		"inst-ex-im",
		"inst-rlc",
		"inst-bit",
		"inst-call-ret",
		"inst-io",
	}

	for tn, base := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		input := string(testutil.ReadTestDataFile(t, base+".asm"))
		expected := testutil.ReadTestDataFile(t, base+".bin")

		prog, e := evalInput(input, logger, env)
		testEvalResult(t, tn, "", e)

		logger.Print()
		result := CollectCode(prog)

		if err := bytesEqual(result, expected); err != nil {
			t.Errorf("[%d] generated code diff %s", tn, err.Error())
		}
	}
}

func TestInstructionDefault(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `ld a, VAL`, code: []byte{0x3e, 0}},
		{input: `ld VAL, a`, code: []byte{0x3e, 0}},
		{input: `ld hl, VAL`, code: []byte{0x21, 0, 0}},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}

		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		prog, _ := evalInput(tt.input, logger, env)

		// エラー発生時のデフォルトコードのチェックのため、発生エラーは無視する
		// testEvalResult(t, tn, tt.err, e)
		// // error, warning, information
		// if tt.err != "" {
		// 	testutil.TestLogMessage(t, tn, tt.err, e.logger)
		// 	continue
		// }

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
	}
}

func TestInstructionError_LD(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		err   string
	}{
		// 0-
		{input: `ld 123, 456`, err: errcode.EZ80_OP1},
		{input: `ld a, nc`, err: errcode.EZ80_OP2},
		{input: `ld a, hl`, err: errcode.EZ80_OP2},
		{input: `ld hl, nc`, err: errcode.EZ80_OP2},
		{input: `ld hl, a`, err: errcode.EZ80_OP2},
		{input: `ld hl, hl`, err: errcode.EZ80_OP1_SP},
		{input: `ld sp, de`, err: errcode.EZ80_OP2_HL_IXY},
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
	}
}
