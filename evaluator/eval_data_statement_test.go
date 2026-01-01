package evaluator

import (
	"testing"
	"yas80/errcode"
	"yas80/logging"
	"yas80/object"
)

func TestDataStoreStatement(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0- 単純ケース
		{input: `ds 1`, code: []byte{0}},
		{input: `ds 2`, code: []byte{0, 0}},
		{input: `ds 1, 255`, code: []byte{255}},
		{input: `ds 2, 255`, code: []byte{255, 255}},
		{input: `dsb 1`, code: []byte{0}},
		// 5-
		{input: `dsb 2`, code: []byte{0, 0}},
		{input: `dsb 1, 255`, code: []byte{255}},
		{input: `dsb 2, 255`, code: []byte{255, 255}},
		{input: `dsw 1`, code: []byte{0, 0}},
		{input: `dsw 2`, code: []byte{0, 0, 0, 0}},
		// 10-
		{input: `dsw 1, $1234`, code: []byte{0x34, 0x12}},
		{input: `dsw 2, $1234`, code: []byte{0x34, 0x12, 0x34, 0x12}},
		{input: `ds 0`, err: errcode.EDS_COUNT},
		{input: `dsb 0`, err: errcode.EDS_COUNT},
		{input: `dsw 0`, err: errcode.EDS_COUNT},
		// 15-
		{input: `ds -1`, err: errcode.EDS_COUNT},
		{input: `dsb -1`, err: errcode.EDS_COUNT},
		{input: `dsw -1`, err: errcode.EDS_COUNT},
		{input: "ds 1, -129", code: []byte{0x7f}, err: errcode.WROUND_BYTE},
		{input: "ds 1, 256", code: []byte{0}, err: errcode.WROUND_BYTE},
		// 20
		{input: "dsb 1, -129", code: []byte{0x7f}, err: errcode.WROUND_BYTE},
		{input: "dsb 1, 256", code: []byte{0}, err: errcode.WROUND_BYTE},
		{input: "dsw 1, -32769", code: []byte{0xff, 0x7f}, err: errcode.WROUND_WORD},
		{input: "dsw 1, 65536", code: []byte{0, 0}, err: errcode.WROUND_WORD},
		{input: `const size = def + 1 \ ds size \ const def = 1`, code: []byte{0, 0}},
		// 25-
		{input: `const size = def + 1 \ dsb size \ const def = 1`, code: []byte{0, 0}},
		{input: `const size = def + 1 \ dsw size \ const def = 1`, code: []byte{0, 0, 0, 0}},
		{input: `dummy func\endf\ const size=dummy() \ ds size`, err: errcode.EDS_COUNT},
		{input: `dummy func\endf\ const size=dummy() \ dsb size`, err: errcode.EDS_COUNT},
		{input: `dummy func\endf\ const size=dummy() \ dsw size`, err: errcode.EDS_COUNT},
		// 30-
		{input: "123 ds 1", err: errcode.ESYNTAX},
		{input: `aaa ds 1 \ aaa ds 1`, err: errcode.ELABEL_DUP},
		{input: `const aaa = 1 \ aaa ds 1`, err: errcode.ELABEL_USED},
		{input: `aaa ds 1, 1 \ bbb ds 2, 2 \ ccc ds 3, 3`,
			code: []byte{1, 2, 2, 3, 3, 3},
			syms: []symValue{
				{"AAA", 0},
				{"BBB", 1},
				{"CCC", 3},
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
func TestDataStatement(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		syms  []symValue
		err   string
	}{
		// 0- 単純ケース
		{input: `db 0`, code: []byte{0}},
		{input: `db 0,1`, code: []byte{0, 1}},
		{input: `dw $1234`, code: []byte{0x34, 0x12}},
		{input: `dw $1234, $5678`, code: []byte{0x34, 0x12, 0x78, 0x56}},
		{input: `dd 0, 0x1234, 5, 0x6789`, code: []byte{0, 0x34, 0x12, 5, 0x89, 0x67}},
		// 5-
		{input: `db -129`, code: []byte{0x7f}, err: errcode.WROUND_BYTE},
		{input: `db 256`, code: []byte{0}, err: errcode.WROUND_BYTE},
		{input: `dw -32769`, code: []byte{0xff, 0x7f}, err: errcode.WROUND_WORD},
		{input: `dw 65536`, code: []byte{0}, err: errcode.WROUND_WORD},
		{input: `dd -32769`, code: []byte{0xff, 0x7f}, err: errcode.WROUND_WORD},
		// 10-
		{input: `dd 65536`, code: []byte{0}, err: errcode.WROUND_WORD},
		{input: `db`, err: errcode.ESYNTAX},
		{input: `dw`, err: errcode.ESYNTAX},
		{input: `dd`, err: errcode.ESYNTAX},
		{input: `db "abcd"`, code: []byte{0x61, 0x62, 0x63, 0x64}},
		// 15-
		{input: `dd "abcd"`, code: []byte{0x61, 0x62, 0x63, 0x64}},
		{input: `db "あいう"`, code: []byte{0x82, 0xa0, 0x82, 0xa2, 0x82, 0xa4}},
		{input: `dd "あいう"`, code: []byte{0x82, 0xa0, 0x82, 0xa2, 0x82, 0xa4}},
		{input: `db "'\"\\\0\a\b\f\n\r\t\v"`,
			code: []byte{0x27, 0x22, 0x5c, 0, 0x07, 0x08, 0x0c, 0x0a, 0x0d, 0x09, 0x0b}},
		{input: `dd "'\"\\\0\a\b\f\n\r\t\v"`,
			code: []byte{0x27, 0x22, 0x5c, 0, 0x07, 0x08, 0x0c, 0x0a, 0x0d, 0x09, 0x0b}},
		// 20-
		// error
		{input: `
			db_data db 1, 2
			dw_data dw $1234, $5678
			dd_data dd $9a, $bcde, $f0
			data_end:`,
			code: []byte{1, 2, 0x34, 0x12, 0x78, 0x56, 0x9a, 0xde, 0xbc, 0xf0},
			syms: []symValue{
				{"DB_DATA", 0},
				{"DW_DATA", 2},
				{"DD_DATA", 6},
				{"DATA_END", 0x0a},
			},
		},
		{input: `dw "あいう"`, err: errcode.EDATA_DW_STR},
		{input: `db "` + string([]byte{0x80, 0xff}) + `"`, err: errcode.EDATA_ENCODE},
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
