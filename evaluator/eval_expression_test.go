package evaluator

import (
	"testing"

	"github.com/dogatana/yas80/errcode"
	"github.com/dogatana/yas80/internal/testutil"
	"github.com/dogatana/yas80/logging"
	"github.com/dogatana/yas80/object"
)

func TestEvalExpressionError(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `_ = abc[]`, err: errcode.EARRAY_INDEX},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestEvalEqualityInnfixExpression(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		{`const val = 1 \ const result = val == 1`, []symValue{{"RESULT", 1}}, ""},
		{`const val = 'a' \ const result = val == 'a'`, []symValue{{"RESULT", 1}}, ""},
		{`const val = [] \ const result = val == []`, []symValue{{"RESULT", 1}}, ""},
		{`const val = [1] \ const result = val == [1]`, []symValue{{"RESULT", 1}}, ""},
		{`const val = [1,2] \ const result = val == [1,2]`, []symValue{{"RESULT", 1}}, ""},
		{`const val = a \ const result = val == a`, []symValue{{"RESULT", 1}}, ""},
		{`const val = z \ const result = val == z`, []symValue{{"RESULT", 1}}, ""},
		{`const val = cy \ const result = val == cy`, []symValue{{"RESULT", 1}}, ""},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestEvalInfixExpression(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		{`const val = 1 \ const result = val`, []symValue{{"RESULT", 1}}, ""},
		{`const val = 1 \ const result = val + 2`, []symValue{{"RESULT", 3}}, ""},
		{`const val = 1 \ const result = val * 2`, []symValue{{"RESULT", 2}}, ""},
		{`const val = 1 \ const result = val - 2 * 3`, []symValue{{"RESULT", -5}}, ""},
		{`const val = 1 \ const result = val % 2`, []symValue{{"RESULT", 1}}, ""},
		// 5-
		{`const val = 1 \ const result = val << 4`, []symValue{{"RESULT", 16}}, ""},
		{`const val = 16 \ const result = val >> 2`, []symValue{{"RESULT", 4}}, ""},
		{`const val = 1 \ const result = val | 2`, []symValue{{"RESULT", 3}}, ""},
		{`const val = 1 \ const result = val & 2`, []symValue{{"RESULT", 0}}, ""},
		{`const val = 1 \ const result = val & 3`, []symValue{{"RESULT", 1}}, ""},
		// 10-
		{`const val = 1 \ const result = val ^ 1`, []symValue{{"RESULT", 0}}, ""},
		{`const val = 1 \ const result = val ^ 0`, []symValue{{"RESULT", 1}}, ""},
		{`const val = 1 \ const result = val == 1`, []symValue{{"RESULT", 1}}, ""},
		{`const val = 1 \ const result = val != 1`, []symValue{{"RESULT", 0}}, ""},
		{`const val = 1 \ const result = val == 2`, []symValue{{"RESULT", 0}}, ""},
		// 15-
		{`const val = 1 \ const result = val != 2`, []symValue{{"RESULT", 1}}, ""},
		{`const val = 1 \ const result = val > 1`, []symValue{{"RESULT", 0}}, ""},
		{`const val = 1 \ const result = val >= 1`, []symValue{{"RESULT", 1}}, ""},
		{`const val = 1 \ const result = val < 0`, []symValue{{"RESULT", 0}}, ""},
		{`const val = 1 \ const result = val <= 1`, []symValue{{"RESULT", 1}}, ""},
		// 20-
		{`const val = 1 \ const result = val &&  1`, []symValue{{"RESULT", 1}}, ""},
		{`const val = 1 \ const result = val &&  0`, []symValue{{"RESULT", 0}}, ""},
		{`const val = 0 \ const result = val &&  0`, []symValue{{"RESULT", 0}}, ""},
		{`const val = 1 \ const result = val &&  "a"`, []symValue{{"RESULT", 1}}, ""},
		{`const val = 1 \ const result = val &&  ""`, []symValue{{"RESULT", 0}}, ""},
		// 25-
		{`const val = 0 \ const result = val &&  ""`, []symValue{{"RESULT", 0}}, ""},
		{`const val = 1 \ const result = val ||  1`, []symValue{{"RESULT", 1}}, ""},
		{`const val = 1 \ const result = val ||  0`, []symValue{{"RESULT", 1}}, ""},
		{`const val = 1 \ const result = val ||  "a"`, []symValue{{"RESULT", 1}}, ""},
		{`const val = 1 \ const result = val ||  ""`, []symValue{{"RESULT", 1}}, ""},
		// 30-
		{`const val = 0 \ const result = val ||  ""`, []symValue{{"RESULT", 0}}, ""},
		{`const val = 1 \ const result = val ||  ""`, []symValue{{"RESULT", 1}}, ""},
		{`const val = "a" \ const result = val + "b"`, []symValue{{"RESULT", "ab"}}, ""},
		// 25-
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestEvalInfixExpressionNumberString(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		{`const val = 1 \ const RESULT = 'a' +  val`, []symValue{{"RESULT", 0x61 + 1}}, ""},
		{`const val = 1 \ const RESULT = 'a' +  val  + 2`, []symValue{{"RESULT", 0x61 + 1 + 2}}, ""},
		{`const val = 1 \ const RESULT = 'a' +  val  + 2 + 'b'`, []symValue{{"RESULT", 0x61 + 1 + 2 + 0x62}}, ""},
		{`const val = 1 \ const RESULT = val + 'a'`, []symValue{{"RESULT", 0x61 + 1}}, ""},
		{`const val = 1 \ const RESULT = val + 'a' +  2`, []symValue{{"RESULT", 0x61 + 1 + 2}}, ""},
		{`const val = 1 \ const RESULT = val + 'a' +  2 + 'b'`, []symValue{{"RESULT", 0x61 + 1 + 2 + 0x62}}, ""},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}
func TestEvalInfixExpressionNumberStringError(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `const val = 1 \ const RESULT = 'ab' +  val`, err: errcode.ESTR_TO_INT_LEN},
		{input: `const val = 1 \ const RESULT = '' +  val`, err: errcode.ESTR_TO_INT_LEN},
		{input: `const val = 1 \ const RESULT = val + 'ab'`, err: errcode.ESTR_TO_INT_LEN},
		{input: `const val = 1 \ const RESULT = val + ''`, err: errcode.ESTR_TO_INT_LEN},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}
func TestEvalInfixExpressionNumberArray(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		{`const val = 2 \ const RESULT = [1] +  val`, []symValue{{"RESULT", 1 + 2}}, ""},
		{`const val = 2 \ const RESULT = [1] +  val  + 3`, []symValue{{"RESULT", 1 + 2 + 3}}, ""},
		{`const val = 2 \ const RESULT = [1] +  val  + 3 + [4]`, []symValue{{"RESULT", 1 + 2 + 3 + 4}}, ""},
		{`const val = 2 \ const RESULT = val + [1]`, []symValue{{"RESULT", 2 + 1}}, ""},
		{`const val = 2 \ const RESULT = val + [1] +  3`, []symValue{{"RESULT", 2 + 1 + 3}}, ""},
		{`const val = 2 \ const RESULT = val + [1] +  3 + [4]`, []symValue{{"RESULT", 2 + 1 + 3 + 4}}, ""},
		{`const val = 2 \ const RESULT = [1,3] +  val`, []symValue{{"RESULT", 0x0103 + 2}}, ""},
		{`const val = 2 \ const RESULT = val + [1,3]`, []symValue{{"RESULT", 2 + 0x0103}}, ""},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestEvalInfixExpressionNumberArrayError(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `const val = 1 \ const RESULT = [] +  val`, err: errcode.EARRAY_TO_INT_LEN},
		{input: `const val = 1 \ const RESULT = [1,2,3] +  val`, err: errcode.EARRAY_TO_INT_LEN},
		{input: `const val = 1 \ const RESULT = val + []`, err: errcode.EARRAY_TO_INT_LEN},
		{input: `const val = 1 \ const RESULT = val + [1,2,3]`, err: errcode.EARRAY_TO_INT_LEN},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestEvalInfixExpressionError(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		// {input: `const val = 1 \ _ = val + "a"`, err: errcode.EBIN_OP_TYPE},
		// {input: `const val = 1 \ _ = "a" - val`, err: errcode.EBIN_OP_TYPE},
		// {input: `const val = 1 \ _ = val * a`, err: errcode.EBIN_OP_TYPE},
		// {input: `const val = 1 \ _ = val * nc`, err: errcode.EBIN_OP_TYPE},
		// {input: `const val = "a" \  _ = 1 / val`, err: errcode.EBIN_OP_TYPE},
		// 5-
		// 10-
		{input: `const val = 0 \  _ = 1 / val`, err: errcode.EBIN_OP_DIVZERO},
		{input: `const abc = 123 / 0`, err: errcode.EBIN_OP_DIVZERO}, // これは parser で出力される
		{input: `test macro arg \ ld a, 1 / arg \ endm \ test 0`, err: errcode.EBIN_OP_DIVZERO},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestEvalPrefixExpression(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		{`const val = +1 \ const result = val`, []symValue{{"RESULT", 1}}, ""},
		{`const val = ++1 \ const result = val`, []symValue{{"RESULT", 1}}, ""},
		{`const val = -1 \ const result = val`, []symValue{{"RESULT", -1}}, ""},
		{`const val = --1 \ const result = val`, []symValue{{"RESULT", 1}}, ""},
		{`const val = 1 \ const result = ~val`, []symValue{{"RESULT", -2}}, ""},
		// 5-
		{`const val = 1 \ const result = !val`, []symValue{{"RESULT", 0}}, ""},
		{`const val = 2 \ const result = !val`, []symValue{{"RESULT", 0}}, ""},
		{`const val = 0 \ const result = !val`, []symValue{{"RESULT", 1}}, ""},
		{`const val = 2 \ const result = !!val`, []symValue{{"RESULT", 1}}, ""},
		{`const val = "" \ const result = !val`, []symValue{{"RESULT", 1}}, ""},
		// 10-
		{`const val = "a" \ const result = !val`, []symValue{{"RESULT", 0}}, ""},
		{`const val = hl \ const result = !val`, []symValue{{"RESULT", 1}}, ""},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestEvalPrefixExpressionError(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `const val = "a" \ const result = +val`, err: errcode.EUNI_OP_TYPE},
		{input: `const val = "a" \ const result = -val`, err: errcode.EUNI_OP_TYPE},
		{input: `const val = "a" \ const result = ~val`, err: errcode.EUNI_OP_TYPE},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestIndexExpression(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `const v = [1,2,3][0]`, syms: []symValue{{"V", 1}}},
		{input: `const v = [1,2,3][1]`, syms: []symValue{{"V", 2}}},
		{input: `const v = [1,2,3][2]`, syms: []symValue{{"V", 3}}},
		{input: `const v = [1,2,3][-1]`, err: errcode.EARRAY_OUT_OF_INDEX},
		{input: `const v = [1,2,3][3]`, err: errcode.EARRAY_OUT_OF_INDEX},
		{input: `const ary = [1,2,3] \ const v = ary[0]`, syms: []symValue{{"V", 1}}},
		{input: `const ary = [1,2,3] \ const v = ary[1]`, syms: []symValue{{"V", 2}}},
		{input: `const ary = [1,2,3] \ const v = ary[2]`, syms: []symValue{{"V", 3}}},
		{input: `const ary = [1,2,3] \ const v = ary[-1]`, err: errcode.EARRAY_OUT_OF_INDEX},
		{input: `const ary = [1,2,3] \ const v = ary[3]`, err: errcode.EARRAY_OUT_OF_INDEX},
		{input: `const v = [1,2,3][a]`, err: errcode.EARRAY_INDEX},
		{input: `const v = [1,2,3]["a"]`, err: errcode.EARRAY_INDEX},
		{input: `const v = ary[1] \ const ary = [1,2,3]`, syms: []symValue{{"V", 2}}},
		{input: `const v = xyz[1] \ const ary = [1,2,3]`, err: errcode.ESYM_UNDEF},
		{input: `const v = [1,2,3][xyz] \ const ary = [2,2,3]`, err: errcode.ESYM_UNDEF},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}

}

func TestRegIndirectExpression(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
		err   string
	}{
		// 0-
		{input: `jp (de)`, err: errcode.EZ80_JP_INDIRECT_REG},
		{input: `jp (de+1)`, err: errcode.EINDIRECT_DISP_REG},
		{input: `jp (ix+1)`, err: errcode.EZ80_JP_INDIRECT_DISP},
		{input: `jp (iy-1)`, err: errcode.EZ80_JP_INDIRECT_DISP},
		{input: `jp (ix+128)`, err: errcode.EINDIRECT_DISP_RANGE},
		// 5-
		{input: `jp (sp)`, err: errcode.EZ80_JP_INDIRECT_REG},
		{input: `jp (a)`, err: errcode.EINDIRECT_REG},
		{input: `jp ('a')`, err: errcode.EINDIRECT_VALUE},
		{input: `jp (123)`, err: errcode.EZ80_JP_INDIRECT_REG},
		// 10-
		{input: `in 1, (c)`, err: errcode.EZ80_OP1},
		{input: `in hl, (c)`, err: errcode.EZ80_OP_REG},
		{input: `in ixl, (c)`, err: errcode.EZ80_OP_REG},
		{input: `in a, (hl)`, err: errcode.EINDIRECT_REG},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New()
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}
		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}

}
