package evaluator

import (
	"testing"
	"yas80/logging"
	"yas80/object"
)

func TestAutoProc(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		names []string // 自動生成した proc 名
		syms  []symValue
	}{
		{`a1: nop \ .l1: nop \ .l2: nop \ a2: nop`,
			[]byte{0, 0, 0, 0},
			[]string{"A1"},
			[]symValue{
				{"A1", 0},
				{"A1.L1", 1},
				{"A1.L2", 2},
				{"A2", 3},
			},
		},
		{`a1: \ nop \ .l1: \ nop \ .l2: \ nop \ a2: \ nop`,
			[]byte{0, 0, 0, 0},
			[]string{"A1"},
			[]symValue{
				{"A1", 0},
				{"A1.L1", 1},
				{"A1.L2", 2},
				{"A2", 3},
			},
		},
		{`a1: nop \ fn func \ endf \ a2: nop`,
			[]byte{0, 0},
			[]string{},
			[]symValue{
				{"A1", 0},
				{"A2", 1},
			},
		},
		{`a1: nop \ pr proc \ endp \ a2: nop`,
			[]byte{0, 0},
			[]string{"PR"},
			[]symValue{
				{"A1", 0},
				{"A2", 1},
			},
		},
		{`a1: nop \ .l1: \ nop \ pr proc \ nop \ endp \ a2: nop`,
			[]byte{0, 0, 0, 0},
			[]string{"A1", "PR"},
			[]symValue{
				{"A1", 0},
				{"A1.L1", 1},
				{"PR", 2},
				{"A2", 3},
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

		for _, name := range tt.names {
			obj, ok := env.Get(name)
			if !ok {
				t.Errorf("[%d] not %s in env", tn, name)
			} else if obj, ok := obj.(*object.ProcObject); !ok {
				t.Errorf("[%d] not ProcObject. got %T", tn, obj)
			}
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}
