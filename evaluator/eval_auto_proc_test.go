package evaluator

import (
	"slices"
	"testing"

	"github.com/dogatana/yas80/intern"
	"github.com/dogatana/yas80/internal/util"
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

		proc_ids, procs := collectProc(env)
		if len(procs) != len(tt.names) {
			t.Errorf("[%d] expect %d procs. got %d", tn, len(tt.names), len(procs))
			continue
		}
		ids := util.Map(tt.names, func(s string) intern.SymbolID { return intern.InternString(s) })
		slices.Sort(ids)
		slices.Sort(proc_ids)
		if !slices.Equal(ids, proc_ids) {
			t.Errorf("[%d] expect %v procs. got %v", tn, ids, proc_ids)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func collectProc(env object.Environment) ([]intern.SymbolID, map[intern.SymbolID]*object.ProcObject) {
	procs := map[intern.SymbolID]*object.ProcObject{}
	keys := []intern.SymbolID{}
	for i, obj := range env.Store() {
		id := intern.SymbolID(i)
		if obj, ok := obj.(*object.ProcObject); ok {
			procs[id] = obj
			keys = append(keys, id)
		}
	}
	return keys, procs
}
