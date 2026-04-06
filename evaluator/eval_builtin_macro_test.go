package evaluator

import (
	"bytes"
	"os"
	"strings"
	"testing"
	"yas80/errcode"
	"yas80/internal/testutil"
	"yas80/internal/util"
	"yas80/logging"
	"yas80/object"
)

func TestAlign(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{
			input: `align 4 \ ret`,
			code:  []byte{0xc9},
		},
		{
			input: `nop \ align 4 \ ret`,
			code:  []byte{0x00, 0x00, 0x00, 0x00, 0xc9},
		},
		{
			input: `nop \ align 4,255 \ ret`,
			code:  []byte{0x00, 0xff, 0xff, 0xff, 0xc9},
		},
		{
			input: `ds 4 \ align 4 \ ret`,
			code:  []byte{0x00, 0x00, 0x00, 0x00, 0xc9},
		},
		{input: `nop \ align 4, 257 \ ret`, code: []byte{0x00, 0x01, 0x01, 0x01, 0xc9}, err: errcode.WROUND_BYTE},
		{input: `align`, err: errcode.EEBMAC_ARG_COUNT},
		{input: `align 1, 2, 3`, err: errcode.EEBMAC_ARG_COUNT},
		{input: `align 'a'`, err: errcode.EEBMAC_ARG_VALUE},
		{input: `align HL`, err: errcode.EEBMAC_ARG_VALUE},
		{input: `fn func\endf\ align fn()`, err: errcode.EEBMAC_ARG_NULL},
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

func TestLogMessage(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{
			input: `error "this is error"`,
			err:   "this is error",
		},
		{
			input: `warn "this is warning"`,
			err:   "this is warning",
		},
		{
			input: `info "this is information"`,
			err:   "this is information",
		},
		{
			input: `info 123`,
			err:   "123",
		},
		{
			input: `warn 123 + 456`,
			err:   "579",
		},
		{
			input: `error 'A' + 1`,
			err:   "66", // 0x41 + 1
		},
		{
			input: `info hl`,
			err:   "HL",
		},
		{
			input: `info nz`,
			err:   "NZ",
		},
		{
			input: `error`,
			err:   errcode.EEBMAC_ARG_COUNT,
		},
		{
			input: `error 1, 2`,
			err:   errcode.EEBMAC_ARG_COUNT,
		},
		{input: `fn func\endf\ error fn()`, err: errcode.EEBMAC_ARG_NULL},
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
			// errcode.*
			if strings.Contains(tt.err, "%") {
				testutil.TestLogMessage(t, tn, tt.err, e.logger)
				continue
			}
			// error, warn, info マクロ
			msgs := util.Map(logger.GetMessages(), func(m *logging.Message) string { return m.String() })
			text := strings.Join(msgs, "\n")
			if !strings.Contains(text, tt.err) {
				t.Errorf("[%d] no %q", tn, tt.err)
			}
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

func TestIncBin(t *testing.T) {
	bin, err := os.ReadFile("testdata/inc.bin")
	if err != nil {
		t.Fatalf("cannot load 'testdata/inc.bin' for test")
	}
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		{
			input: `incbin "inc.bin"`,
			code:  bin,
			syms:  []symValue{{"$RSIZE", len(bin)}},
		},
		{
			input: `binclude "inc.bin"`,
			code:  bin,
			syms:  []symValue{{"$RSIZE", len(bin)}},
		},
		{
			input: `incbin "inc.bin", 0x80`,
			code:  bin[0x80:],
			syms:  []symValue{{"$RSIZE", 0x80}},
		},
		{
			input: `incbin "inc.bin", 0, 0x80`,
			code:  bin[:0x80],
			syms:  []symValue{{"$RSIZE", 0x80}},
		},
		{
			input: `incbin "inc.bin", 0x40, 0x40`,
			code:  bin[0x40:0x80],
			syms:  []symValue{{"$RSIZE", 0x40}},
		},
		{
			input: `incbin "inc.bin", 0x200`,
			code:  []byte{},
			syms:  []symValue{{"$RSIZE", 0}},
		},
		{
			input: `incbin "inc.bin", 0x80, 0x100`,
			code:  bin[0x80:],
			syms:  []symValue{{"$RSIZE", 0x80}},
		},
		{
			input: `incbin 123`,
			err:   errcode.EEBMAC_ARG_VALUE,
			syms:  []symValue{{"$RSIZE", -1}},
		},
		{
			input: `fn func\endf\incbin fn()`,
			err:   errcode.EEBMAC_ARG_NULL,
			syms:  []symValue{{"$RSIZE", -1}},
		},
		{
			input: `incbin`,
			err:   errcode.EEBMAC_ARG_COUNT,
			syms:  []symValue{{"$RSIZE", -1}},
		},
		{
			input: `incbin "inc.bin", 1, 2, 3`,
			err:   errcode.EEBMAC_ARG_COUNT,
			syms:  []symValue{{"$RSIZE", -1}},
		},
		{
			input: `incbin "inc-not-found.bin"`,
			err:   errcode.EFILE_NOT_FOUND,
			syms:  []symValue{{"$RSIZE", -1}},
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

		// check system variables
		for _, s := range tt.syms {
			num, ok := e.getNumberFromEnv(s.name, env)
			if !ok {
				t.Errorf("[%d] %s not found in env", tn, s.name)
			}
			if num.Value != s.expected {
				t.Errorf("[%d] %s is not %v. bot %v", tn, s.name, s.expected, num.Value)
			}
		}

		// error, warning, information
		if tt.err != "" {
			// errcode.*
			if strings.Contains(tt.err, "%") {
				testutil.TestLogMessage(t, tn, tt.err, e.logger)
				continue
			}
			// error, warn, info マクロ
			msgs := util.Map(logger.GetMessages(), func(m *logging.Message) string { return m.String() })
			text := strings.Join(msgs, "\n")
			if !strings.Contains(text, tt.err) {
				t.Errorf("[%d] no %q", tn, tt.err)
			}
			continue
		}

		// code
		testCodeResult(t, tn, tt.code, prog)

	}
}

func TestSetMap(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0-
		{
			input: `charmap cmap, '{ "a": [1] }', -2 \ setmap cmap, "a", [2] \ db cmap('abc')`,
			code:  []byte{2, 0x62, 0x63},
		},
		{
			input: `charmap cmap, '{ "a": [1] }', -2 \ setmap cmap, "b", [2] \ db cmap('abc')`,
			code:  []byte{1, 2, 0x63},
		},
		{
			input: `charmap cmap, '{ "a": [1] }', -2 \ setmap cmap, "b", [2, 3, 4] \ db cmap('abc')`,
			code:  []byte{1, 2, 3, 4, 0x63},
		},
		{
			input: `charmap cmap, '{ "a": [1] }', -2 \ setmap cmap, "あ", [2] \ db cmap('aあc')`,
			code:  []byte{1, 2, 0x63},
		},
		{
			input: `charmap cmap, '{ "a": [1] }', -2 \ setmap cmap, "ab", [2] \ db cmap('abc')`,
			err:   errcode.ESETMAP_NOT_A_CHAR,
		},
		{
			input: `charmap cmap, '{ "a": [1] }', -2 \ setmap cmap, "", [2] \ db cmap('abc')`,
			err:   errcode.ESETMAP_NOT_A_CHAR,
		},
		{
			input: `charmap cmap, '{ "a": [1] }', -2 \ setmap cmap, "a", [] \ db cmap('abc')`,
			err:   errcode.ESETMAP_ARRAY_EMPTY,
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

		// check system variables
		for _, s := range tt.syms {
			num, ok := e.getNumberFromEnv(s.name, env)
			if !ok {
				t.Errorf("[%d] %s not found in env", tn, s.name)
			}
			if num.Value != s.expected {
				t.Errorf("[%d] %s is not %v. bot %v", tn, s.name, s.expected, num.Value)
			}
		}

		// error, warning, information
		if tt.err != "" {
			// errcode.*
			if strings.Contains(tt.err, "%") {
				testutil.TestLogMessage(t, tn, tt.err, e.logger)
				continue
			}
			// error, warn, info マクロ
			msgs := util.Map(logger.GetMessages(), func(m *logging.Message) string { return m.String() })
			text := strings.Join(msgs, "\n")
			if !strings.Contains(text, tt.err) {
				t.Errorf("[%d] no %q", tn, tt.err)
			}
			continue
		}

		// code
		testCodeResult(t, tn, tt.code, prog)
	}
}

func TestBuiltinMacroInUserMacro(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		{
			input: `
				tm macro arg
				  var @value = arg
				  info @value
				endm
				tm "this is a message"
				`,
			err: "this is a message",
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

		// check system variables
		for _, s := range tt.syms {
			num, ok := e.getNumberFromEnv(s.name, env)
			if !ok {
				t.Errorf("[%d] %s not found in env", tn, s.name)
			}
			if num.Value != s.expected {
				t.Errorf("[%d] %s is not %v. bot %v", tn, s.name, s.expected, num.Value)
			}
		}

		// error, warning, information
		if tt.err != "" {
			// errcode.*
			if strings.Contains(tt.err, "%") {
				testutil.TestLogMessage(t, tn, tt.err, e.logger)
				continue
			}
			// error, warn, info マクロ
			msgs := util.Map(logger.GetMessages(), func(m *logging.Message) string { return m.String() })
			text := strings.Join(msgs, "\n")
			if !strings.Contains(text, tt.err) {
				t.Errorf("[%d] no %q", tn, tt.err)
			}
			continue
		}

		// code
		testCodeResult(t, tn, tt.code, prog)
	}
}
func TestBuiltinMacroCheck256(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		{
			input: `base ds 256, 0 \ check256 base`,
			code:  bytes.Repeat([]byte{0}, 256),
		},
		{
			input: `base ds 257, 0 \ check256 base`,
			code:  bytes.Repeat([]byte{0}, 257),
			err:   errcode.ECHECK256_ERROR,
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
		if tt.input == "" {
			continue
		}
	}
}
