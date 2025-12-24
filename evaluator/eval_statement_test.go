package evaluator

import (
	"testing"
	"yas80/logger"
	"yas80/object"
)

func TestAssemble(t *testing.T) {
	tests := []string{
		"label-backward",
		"equ-backward",
		"label-forward",
		"equ-forward",
		"forward",
		"forward_symbol",
		"forward_mix",
	}

	for tn, base := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		input := string(readTestDataFile(t, base+".asm"))
		expected := readTestDataFile(t, base+".bin")

		prog := evaluateInput(t, input, logger, env)
		logger.Print()
		result := CollectCode(prog)

		if !bytesEqual(result, expected) {
			t.Errorf("[%d] expected %d bytes. got %d bytes", tn, len(expected), len(result))
		}
	}
}

func TestConstStatement(t *testing.T) {
	tests := []struct {
		input    string
		name     string
		expected int
	}{
		{"const abc ## 123 = 456", "ABC123", 456},
		{"const abc ## (100 + 23) = 456", "ABC123", 456},
		{"const abc ## (100 + 23) = def ## 456 \\ const def ## (400 + 56) = 999", "ABC123", 999},
		{"abc ## 123 equ 456", "ABC123", 456},
		{"abc ## (100 + 23) equ 456", "ABC123", 456},
		{"abc ## (100 + 23) equ def ## 456 \\ const def ## (400 + 56) = 999", "ABC123", 999},
	}
	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		_ = evaluateInput(t, tt.input, logger, env)

		obj, ok := env.Get(tt.name)
		if !ok {
			t.Fatalf(`[%d] %q not in env`, tn, tt.name)
		}
		value := evalValue(obj)
		testNumberObject(t, tn, value, tt.expected)
	}
}

func TestProcStatement(t *testing.T) {
	tests := []struct {
		input    string
		name     string
		expected int
	}{
		{`test proc \ const abc ## (100 + 23) = 456 \ endp`, `ABC123`, 456},
		{`test proc \ const abc ## (100 + 23) = def ## 456 \ const def ## (400 + 56) = 999 \ endp`, `ABC123`, 999},
		{`test proc \ abc ## 123 equ 456 \ endp`, `ABC123`, 456},
		{`test proc \ abc ## (100 + 23) equ 456 \ endp`, `ABC123`, 456},
		{`test proc \ abc ## (100 + 23) equ def ## 456 \ const def ## (400 + 56) = 999  \ endp`, `ABC123`, 999},
	}
	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		_ = evaluateInput(t, tt.input, logger, env)

		obj, ok := env.Get(tt.name)
		if !ok {
			t.Fatalf(`[%d] %q not in env`, tn, tt.name)
		}
		value := evalValue(obj)
		testNumberObject(t, tn, value, tt.expected)
	}
}

func TestProcLabel(t *testing.T) {
	tests := []struct {
		input  string
		code   []byte
		names  []string
		values []int
	}{
		{`abc proc \ ld a, .abc \ .abc: nop \ endp \ xyz proc \ ld a, .abc \ .abc: ret \ endp`,
			[]byte{0x3e, 0x02, 0x00, 0x3e, 0x05, 0xc9}, // ld a,2 \ nop \ ld a,5 \ ret
			[]string{"ABC.ABC", "XYZ.ABC"},
			[]int{2, 5},
		},
		{`abc proc \ ld a, .abc \ nop \ const .abc = 0xa5 \ endp \ xyz proc \ ld a, .abc \ .abc: ret \ endp`,
			[]byte{0x3e, 0xa5, 0x00, 0x3e, 0x05, 0xc9}, // ld a,2 \ nop \ ld a,5 \ ret
			[]string{"ABC.ABC", "XYZ.ABC"},
			[]int{0xa5, 5},
		},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		prog := evaluateInput(t, tt.input, logger, env)

		code := CollectCode(prog)
		if len(code) != len(tt.code) && !bytesEqual(code, tt.code) {
			t.Errorf("[%d] generated code differ", tn)
		}
		for i, name := range tt.names {
			testDotIdentInEnv(t, tn, name, env, tt.values[i])
		}
	}
}
