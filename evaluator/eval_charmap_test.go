package evaluator

import (
	"testing"
	"yas80/errcode"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
)

func TestCharmapDef(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// json file
		{input: `charmap cmap, 'cmap.json' \ charmap cmap, 'cmap.json'`, err: errcode.ECHARMAP_DUP},
		{input: `const cmap = 1 \ charmap cmap,'cmap.json'`, err: errcode.ECHARMAP_USED},
		{input: `fn func\endf \ charmap cmap, fn()`, err: errcode.ECHARMAP_NULL},
		{input: `charmap cmap, 123`, err: errcode.ECHARMAP_NOT_STR},
		{input: `charmap cmap, HL`, err: errcode.ECHARMAP_NOT_STR},
		{input: `fn func\endf \ charmap cmap, 'cmap.json', fn()`, err: errcode.ECHARMAP_DEFCHAR_NULL},
		{input: `charmap cmap, 'cmap.json', 'x'`, err: errcode.ECHARMAP_DEFCHAR_NOT_INT},
		{input: `charmap cmap, 'cmap.json', HL`, err: errcode.ECHARMAP_DEFCHAR_NOT_INT},
		{input: `charmap cmap, 'no.json'`, err: errcode.ECHARMAP_READ},
		{input: `charmap cmap, 'zilog.bin`, err: errcode.ECHARMAP_JSON},
		{input: `charmap cmap, 'cmap-err1.json'`, err: errcode.ECHARMAP_JSON}, // [1]
		{input: `charmap cmap, 'cmap-err2.json'`, err: errcode.ECHARMAP_FMT},  // { "a": 1 }
		// json text
		{input: `charmap cm, '{"a":1}'`, err: errcode.ECHARMAP_FMT},
	}
	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New("testdata/test.asm") // testdata/ 下の json を読むため
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

func TestCharmapApply(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// cmap.json - { "a": [97], "b": [1, 2], "あ": [130, 160], "い": [66] }
		{
			input: `charmap cmap, 'cmap.json' \ db cmap('abあい')`,
			code:  []byte{97, 1, 2, 130, 160, 66},
		},
		{
			input: `charmap cmap, 'cmap.json', 255 \ db cmap('abxy')`,
			code:  []byte{97, 1, 2, 255, 255},
		},
		{
			input: `charmap cmap, 'cmap.json', 0x82a0\ db cmap('abxy')`,
			code:  []byte{97, 1, 2, 0x82, 0xa0, 0x82, 0xa0},
		},
		{
			input: `charmap cmap, 'cmap.json', 0x182a0\ db cmap('abxy')`,
			code:  []byte{97, 1, 2, 0x82, 0xa0, 0x82, 0xa0},
			err:   errcode.WROUND_WORD,
		},
		{
			input: `charmap cmap, '{"a":[97],"b":[1,2],"あ":[130,160],"い":[66]}' \ db cmap('abあい')`,
			code:  []byte{97, 1, 2, 130, 160, 66},
		},
		{
			input: `charmap cmap, '{"a":[97],"b":[1,2],"あ":[130,160],"い":[66]}', 255 \ db cmap('abxy')`,
			code:  []byte{97, 1, 2, 255, 255},
		},
		{
			input: `charmap cmap, '{"a":[97],"b":[1,2],"あ":[130,160],"い":[66]}', 0x82a0\ db cmap('abxy')`,
			code:  []byte{97, 1, 2, 0x82, 0xa0, 0x82, 0xa0},
		},
		{
			input: `charmap cmap, '{"a":[97],"b":[1,2],"あ":[130,160],"い":[66]}', 0x182a0\ db cmap('abxy')`,
			code:  []byte{97, 1, 2, 0x82, 0xa0, 0x82, 0xa0},
			err:   errcode.WROUND_WORD,
		},

		{input: `charmap cm, '{"a":1}'`, err: errcode.ECHARMAP_FMT},
		{input: `charmap cmap, 'cmap.json' \ db cmap('x')`, err: errcode.ECHARMAP_NOT_DEF},
		{input: `fn func \ endf \ charmap cmap, 'cmap.json' \ db cmap(fn())`, err: errcode.ECHARMAP_VALUE_NULL},
		{input: `fn func \ endf \ charmap cmap, 'cmap.json' \ db cmap(123)`, err: errcode.ECHARMAP_VALUE},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New("testdata/test.asm")
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
