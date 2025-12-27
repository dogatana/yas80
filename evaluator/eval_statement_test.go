package evaluator

import (
	"testing"
	"yas80/logging"
	"yas80/object"
)

func TestLabelStatement(t *testing.T) {
	tests := []struct {
		input  string
		code   []byte
		names  []string
		values []int
	}{
		{`addr1: ld hl, $1234 \ addr2: ld a, a \ addr3: ld hl, $5678`,
			[]byte{0x21, 0x34, 0x12, 0x7f, 0x21, 0x78, 0x56},
			[]string{"ADDR1", "ADDR2", "ADDR3"},
			[]int{0, 3, 4},
		},
		{`addr1: ld a, a\ addr2: ld hl, $1234 \ addr3: ld hl, $5678`,
			[]byte{0x7f, 0x21, 0x34, 0x12, 0x21, 0x78, 0x56},
			[]string{"ADDR1", "ADDR2", "ADDR3"},
			[]int{0, 1, 4},
		},
		{`addr1: ld hl, $1234 \ addr2: ld hl, $5678 \ addr3: ld hl, $9abc`,
			[]byte{0x21, 0x34, 0x12, 0x21, 0x78, 0x56, 0x21, 0xbc, 0x9a},
			[]string{"ADDR1", "ADDR2", "ADDR3"},
			[]int{0, 3, 6},
		},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		prog, _ := evaluateInput(t, tt.input, logger, env)

		code := CollectCode(prog)
		if err := bytesEqual(code, tt.code); err != nil {
			t.Errorf("[%d] %s", tn, err.Error())
		}
		for i, name := range tt.names {
			if v, ok := env.Get(name); !ok {
				t.Errorf("[%d] no %q in env", tn, name)
			} else {
				testSymbolNumberObject(t, tn, v, tt.values[i])
			}
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
		logger := logging.New("<eval test>")
		_, _ = evaluateInput(t, tt.input, logger, env)

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
		logger := logging.New("<eval test>")
		_, _ = evaluateInput(t, tt.input, logger, env)

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
		// 0-
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
		{`abc proc \ const .xxx = zzz + 1 \nop \ endp \ const zzz = 1 \ ld a, abc.xxx`,
			[]byte{0, 0x3e, 0x02},
			[]string{"ABC.XXX", "ZZZ"},
			[]int{2, 1},
		},
		{`ld a, abc.def \ abc proc \ nop \ const .def = 4 \ nop \ endp`,
			[]byte{0x3e, 0x04, 0, 0},
			[]string{"ABC.DEF"},
			[]int{4},
		},
		{`ld a, abc.def \ abc proc \ nop \ .def: ret \ nop \ endp`,
			[]byte{0x3e, 0x03, 0, 0xc9, 0},
			[]string{"ABC.DEF"},
			[]int{3},
		},
		// 5-
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		prog, e := evaluateInput(t, tt.input, logger, env)

		code := CollectCode(prog)
		if err := bytesEqual(code, tt.code); err != nil {
			t.Errorf("[%d] generated code diff %s", tn, err.Error())
		}
		for i, name := range tt.names {
			sym, ok := e.getSymbolFromEnv(name, env)
			if !ok {
				t.Errorf("[%d] symbol %s not found", tn, name)
			} else {
				testSymbolNumberObject(t, tn, sym, tt.values[i])
			}
		}
	}
}
