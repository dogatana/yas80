package evaluator

import (
	"testing"
	"yas80/errcode"
	"yas80/internal/testutil"
	"yas80/logging"
	"yas80/object"
)

func TestFuncIfReturn(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
	}{
		{`function test() 1 \ const result = test()`, []symValue{{"RESULT", 1}}},
		{`test func\ const aa=1 \ endf \ const result = test()`, []symValue{{"RESULT", nil}}},
		{`test func\ const aa=1 \ return \ const aa=2\ endf \ const result = test()`, []symValue{{"RESULT", nil}}},
		{`test func\ const aa=1 \ return 99 \ const aa=2 \ endf \ const result = test()`, []symValue{{"RESULT", 99}}},
		{`test func\ if 1 \ if 2 \ const aa=3 \ return 99 \ \ endif \ return 98 \ const aa=5 \ endif \ endf \ const result = test()`, []symValue{{"RESULT", 99}}},
		{`test func\ if 1 \ if 0 \ const aa=3 \ return 99 \ \ endif \ return 98 \ const aa=5 \ endif \ endf \ const result = test()`, []symValue{{"RESULT", 98}}},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, "", e)

		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter) // sym.Expected < 0 なので testSymValuesEx を使用する
	}
}

func TestIf(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
	}{
		// 式文を除外したことで -1 は RESULT が未定義の意味に変更
		{`if 1 \ endif`, []symValue{{"RESULT", nil}}},
		{`if 0 \ endif`, []symValue{{"RESULT", nil}}},
		{`if 1 \ else \ endif`, []symValue{{"RESULT", nil}}},
		{`if 0 \ else \ endif`, []symValue{{"RESULT", nil}}},

		{`if 1 \ const result=100 \ endif`, []symValue{{"RESULT", 100}}},
		{`if 0 \ const result=100 \ endif`, []symValue{{"RESULT", nil}}},

		{`if 1 \ const result=100 \ else \ endif`, []symValue{{"RESULT", 100}}},
		{`if 0 \ const result=100 \ else \ endif`, []symValue{{"RESULT", nil}}},

		{`if 1 \ const result=100 \ else \ const result=200  \ endif`, []symValue{{"RESULT", 100}}},
		{`if 0 \ const result=100 \ else \ const result=200  \ endif`, []symValue{{"RESULT", 200}}},

		{`const val = 1 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200  \ endif`, []symValue{{"RESULT", 100}}},
		{`const val = 2 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200  \ endif`, []symValue{{"RESULT", 200}}},
		{`const val = 3 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200  \ endif`, []symValue{{"RESULT", nil}}},

		{`const val = 3 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200 \ else \ const result=300 \ endif`, []symValue{{"RESULT", 300}}},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, "", e)

		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter) // sym.Expected < 0 なので testSymValuesEx を使用する

	}
}

func TestFunc(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
	}{
		{`ret100 func \ return 100 \ endf \ const result = ret100()`, []symValue{{"RESULT", 100}}},
		{`ret100 func \ _=1 \ _=2 \ return 100 \ _=4 \ endf \ const result = ret100()`, []symValue{{"RESULT", 100}}},
		{`abs func arg \ _=1 \ return arg \ _=2 \ endf \  const result = abs(123)`, []symValue{{"RESULT", 123}}},
		{`abs func arg \ if arg > 0 \ return arg \ else \ return -arg \ endif \ endf \ const result = abs(100)`, []symValue{{"RESULT", 100}}},
		{`abs func arg \ if arg > 0 \ return arg \ else \ return -arg \  endif \endf \ const result = abs(-100)`, []symValue{{"RESULT", 100}}},
		{`deep func arg \ if arg > 1 \ if arg > 2 \ if arg > 3 \ return 999 \ endif \ return 888 \ endif \  return 777 \ endif \ endf \ const result = deep(1)`, []symValue{{"RESULT", nil}}},
		{`deep func arg \ if arg > 1 \ if arg > 2 \ if arg > 3 \ return 999 \ endif \ return 888 \ endif \  return 777 \ endif \ endf \ const result = deep(2)`, []symValue{{"RESULT", 777}}},
		{`deep func arg \ if arg > 1 \ if arg > 2 \ if arg > 3 \ return 999 \ endif \ return 888 \ endif \  return 777 \ endif \ endf \ const result = deep(3)`, []symValue{{"RESULT", 888}}},
		{`deep func arg \ if arg > 1 \ if arg > 2 \ if arg > 3 \ return 999 \ endif \ return 888 \ endif \  return 777 \ endif \ endf \ const result = deep(4)`, []symValue{{"RESULT", 999}}},
		{`	fib func x
					if x < 2
						return 1
					else
						return fib(x - 1) + fib(x - 2)
					endif
				endf
			const result = fib(5)
			`, []symValue{{"RESULT", 8}}},
	}
	// t.Fatal("const 再定義要修正")

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, "", e)

		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestClosure(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
	}{

		{`
		adder func x
			inner func y
				return x + y
			endf
			return inner
		endf
		const add3 = adder(3)
		const result = add3(10)
		`,
			[]symValue{{"RESULT", 13}},
		},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, "", e)

		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestFibFunc(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
	}{
		{`
		fib func x
			if x == 0
				return 1
			elif x == 1
				return 1
			else
				return fib(x - 1) + fib(x - 2)
			endif
		endf
		const result = fib(10)
		`,
			[]symValue{{"RESULT", 89}}},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, "", e)

		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestFunction(t *testing.T) {
	tests := []struct {
		input string
		syms  []symValue
	}{
		{`function vram(x, y) 0xd000 + y * 40 + x \ const result = vram(0, 0)`, []symValue{{"RESULT", 0xd000}}},
		{`function vram(x, y) 0xd000 + y * 40 + x \ const result = vram(39, 24)`, []symValue{{"RESULT", 0xd000 + 40*24 + 39}}},
		{`function add1(x) x + 1 \ const fn = add1 \ const result = fn(99)`, []symValue{{"RESULT", 100}}},
		{`	function add1(x) x + 1
			function double(fn, x) fn(fn(x))
			const result = double(add1, 98)`, []symValue{{"RESULT", 100}}},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		_, e := evalInput(tt.input, logger, env)
		testEvalResult(t, tn, "", e)

		getter := func(name string) (*object.SymbolObject, bool) {
			return e.getSymbolFromEnv(name, env)
		}
		testSymValues(t, tn, tt.syms, getter)
	}
}

func TestFuncErrorWarning(t *testing.T) {
	tests := []struct {
		input string
		err   string
	}{
		{`t func \ ld a,a \endf`, errcode.WSCOPE_FUNC},
		{`t func \ abc: \endf`, errcode.WSCOPE_FUNC},
		{`t func \ rept 3 \ nop \ endr \endf`, errcode.WSCOPE_FUNC},
		{`t func \ tm macro \ nop \ endm \endf`, errcode.WSCOPE_FUNC},
		{`tm macro \ _ = 1 \ endm \ t func \ tm \endf`, errcode.WSCOPE_FUNC},
		// 5-
		{`tm macro \ exitm \ endm \ t func \ tm \ endf`, errcode.WSCOPE_FUNC},
		{`f1 func \ endf \ f1 func \ endf`, errcode.EFUNC_DUP},
		{`const f1 = 1 \ f1 func \ endf`, errcode.EFUNC_USED},
		{`.t func \ endf`, errcode.EFUNC_NAME},
		{`@t func \ endf`, errcode.EFUNC_NAME},
		{`test func \ if 0 \ elif 1 \ elif 2 \elif 3 \ elif 4 \ aaa enum \ende \endif \endf`, errcode.WSCOPE_FUNC},
	}

	for tn, tt := range tests {
		if tt.input == "" {
			continue
		}
		env := object.NewEnvironment(nil)
		logger := logging.New("<eval test>")
		_, e := evalInput(tt.input, logger, env)

		testEvalResult(t, tn, tt.err, e)

		// error, warning, information
		if tt.err != "" {
			testutil.TestLogMessage(t, tn, tt.err, e.logger)
			continue
		}
	}
}
