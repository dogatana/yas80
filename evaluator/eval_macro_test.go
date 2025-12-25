package evaluator

import (
	"testing"
	"yas80/logger"
	"yas80/object"
)

func TestMacro(t *testing.T) {
	tests := []string{
		"macro",
	}

	for tn, base := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		input := string(readTestDataFile(t, base+".asm"))
		expected := readTestDataFile(t, base+".bin")

		prog := evaluateInput(t, input, logger, env)
		logger.Print()
		result := CollectCode(prog)

		if err := bytesEqual(result, expected); err != nil {
			t.Errorf("[%d] generated code diff %s", tn, err.Error())
		}
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
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		input := tt.input
		expected := tt.code

		prog := evaluateInput(t, input, logger, env)
		logger.Print()
		result := CollectCode(prog)

		if err := bytesEqual(result, expected); err != nil {
			t.Errorf("[%d] generated code diff %s", tn, err.Error())
		}
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
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		input := tt.input
		expected := tt.code

		prog := evaluateInput(t, input, logger, env)
		logger.Print()
		result := CollectCode(prog)

		if err := bytesEqual(result, expected); err != nil {
			t.Errorf("[%d] generated code diff %s", tn, err.Error())
		}
	}
}

func TestMacroIfExitmLocal(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
	}{
		{
			`test macro arg \ if arg == 0 \ const @abc = arg \ elif arg == 1 \ const @abc = arg * 16 \ else \ const @abc = arg * 256\ endif \ ld hl, @abc \ endm \ test 0`,
			[]byte{0x21, 0x00, 0x00},
		},
		{
			`test macro arg \ if arg == 0 \ const @abc = arg \ elif arg == 1 \ const @abc = arg * 16 \ else \ const @abc = arg * 256\ endif \ ld hl, @abc \ endm \ test 1`,
			[]byte{0x21, 0x10, 0x00}, // ld hl, 16
		},
		{
			`test macro arg \ if arg == 0 \ const @abc = arg \ elif arg == 1 \ const @abc = arg * 16 \ else \ const @abc = arg * 256\ endif \ ld hl, @abc \ endm \ test 2`,
			[]byte{0x21, 0x00, 0x02}, // ld hl, 512
		},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		input := tt.input
		expected := tt.code

		prog := evaluateInput(t, input, logger, env)
		logger.Print()
		result := CollectCode(prog)

		if err := bytesEqual(result, expected); err != nil {
			t.Errorf("[%d] generated code diff %s", tn, err.Error())
		}
	}
}
