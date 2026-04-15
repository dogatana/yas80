package evaluator

import (
	"slices"
	"testing"

	"github.com/dogatana/yas80/logging"
	"github.com/dogatana/yas80/object"
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
		{`a1: nop \ .l1: nop \ .l2: nop \ @@: nop \ @1: nop \ a2: nop`,
			[]byte{0, 0, 0, 0, 0, 0},
			[]string{"A1"},
			[]symValue{
				{"A1", 0},
				{"A1.L1", 1},
				{"A1.L2", 2},
				{"A2", 5},
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

		proc_names, procs := collectProc(env)
		if len(procs) != len(tt.names) {
			t.Errorf("[%d] expect %d procs. got %d", tn, len(tt.names), len(procs))
			continue
		}
		names := slices.Clone(tt.names)
		slices.Sort(names)
		slices.Sort(proc_names)
		if !slices.Equal(names, proc_names) {
			t.Errorf("[%d] expect %v procs. got %v", tn, names, proc_names)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func collectProc(env object.Environment) ([]string, map[string]*object.ProcObject) {
	procs := map[string]*object.ProcObject{}
	keys := []string{}
	for id, obj := range env.Store() {
		name := id.String()
		if obj, ok := obj.(*object.ProcObject); ok {
			procs[name] = obj
			keys = append(keys, name)
		}
	}
	return keys, procs
}
