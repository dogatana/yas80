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
		syms  []SymValue
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
			syms: []SymValue{
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
