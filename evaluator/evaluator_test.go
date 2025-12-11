package evaluator

import (
	"testing"
	"yas80/logger"
	"yas80/object"
)

func TestEvalLabelStatement(t *testing.T) {
	tests := []struct {
		input string
		code  []byte
		names []string
		value []int
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
		logger := logger.New("<eval test>")
		prog := evaluateInput(t, tt.input, logger, env)

		code := collectCode(prog)
		if len(code) != len(tt.code) && !bytesEqual(code, tt.code) {
			t.Errorf("[%d] generated code differ", tn)
		}
		for i, name := range tt.names {
			if v, ok := env.Get(name); !ok {
				t.Errorf("[%d] no %q in env", tn, name)
			} else {
				testSymbolNumberObject(t, tn, v, tt.value[i])
			}
		}
	}
}

func TestEvalExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"const result = 1", 1},
		{"const result = 1 + 2", 3},
		{"const result = 1 * 2", 2},
		{"const result = 1 + 2 * 3", 7},
		{`const val = 123 \ const result = val`, 123},
		{`const val = 11 \ const result = val * val`, 121},
		{`const val = 11 \ const val2 = val * val \ const result = val2`, 121},
		{"const val = 11 \r\n const val2 = val * val \r\n const result = val2", 121},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		_ = evaluateInput(t, tt.input, logger, env)

		obj, ok := env.Get("RESULT")
		if !ok {
			t.Fatalf(`[%d] "RESULT" not in env`, tn)
		}

		value := evalValue(obj)
		testNumberObject(t, tn, value, tt.expected)
	}
}

func TestFuncIfReturn(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`function test() 1 \ const result = test()`, 1},
		{`test func\ const aa=1 \ endf \ const result = test()`, -1},
		{`test func\ const aa=1 \ return \ const aa=2\ endf \ const result = test()`, -1},
		{`test func\ const aa=1 \ return 99 \ const aa=2 \ endf \ const result = test()`, 99},
		{`test func\ if 1 \ if 2 \ const aa=3 \ return 99 \ \ endif \ return 98 \ const aa=5 \ endif \ endf \ const result = test()`, 99},
		{`test func\ if 1 \ if 0 \ const aa=3 \ return 99 \ \ endif \ return 98 \ const aa=5 \ endif \ endf \ const result = test()`, 98},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		_ = evaluateInput(t, tt.input, logger, env)

		obj, ok := env.Get("RESULT")
		if !ok {
			t.Fatalf(`[%d] "RESULT" not in env`, tn)
		}
		value := evalValue(obj)

		if tt.expected >= 0 {
			testNumberObject(t, tn, value, tt.expected)
			continue
		}
		if value != object.NULL {
			t.Errorf("[%d] should be NULL. got %T(%#v)", tn, value, value)
		}
	}
}

func testIf(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		// 式文を文法から削除したことで一部のケースを除外
		// {`if 1 \ endif`, -1}, // -1 は NullObject とする
		// {`if 0 \ endif`, -1},
		// {`if 1 \ else \ endif`, -1},
		// {`if 0 \ else \ endif`, -1},

		{`if 1 \ const result=100 \ endif`, 100},
		{`if 0 \ const result=100 \ endif`, -1},

		{`if 1 \ const result=100 \ else \ endif`, 100},
		{`if 0 \ const result=100 \ else \ endif`, -1},

		{`if 1 \ const result=100 \ else \ const result=200  \ endif`, 100},
		{`if 0 \ const result=100 \ else \ const result=200  \ endif`, 200},

		{`const val = 1 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200  \ endif`, 100},
		{`const val = 2 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200  \ endif`, 200},
		{`const val = 3 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200  \ endif`, -1},

		{`const val = 3 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200  \ else \ const result=300 \ endif`, 300},

		// {`const val = 2 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200 \ return 999 \ 250  \ else \ 300 \ endif`, 999},
		// {`const val = 2 \ if val == 1 \ const result=100 \ elif val == 2 \ const result=200 \ return 999 \ 250  \ else \ 300 \ endif`, 999},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		_ = evaluateInput(t, tt.input, logger, env)

		obj, ok := env.Get("RESULT")
		if !ok {
			t.Fatalf(`[%d] "RESULT" not in env`, tn)
		}
		value := evalValue(obj)

		if tt.expected >= 0 {
			testNumberObject(t, tn, value, tt.expected)
			continue
		} else {
			t.Errorf("[%d] should be NULL. got %T(%#v)", tn, obj, obj)
		}
	}
}

func testFunc(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`ret100 func \ return 100 \ endf \ const result = ret100()`, 100},
		{`ret100 func \ 1 \ 2 \ return 100 \ 4 \ endf \ const result = ret100()`, 100},
		{`abs func arg \ 1 \ return arg \ 2 \ endf \  const result = abs(123)`, 123},
		{`abs func arg \ if arg > 0 \ return arg \ else \ return -arg \ endif \ endf \ const result = abs(100)`, 100},
		{`abs func arg \ if arg > 0 \ return arg \ else \ return -arg \  endif \endf \ const result = abs(-100)`, 100},
		{`deep func arg \ if arg > 1 \ if arg > 2 \ if arg > 3 \ return 999 \ endif \ return 888 \ endif \  return 777 \ endif \ endf \ const result = deep(1)`, -1},
		{`deep func arg \ if arg > 1 \ if arg > 2 \ if arg > 3 \ return 999 \ endif \ return 888 \ endif \  return 777 \ endif \ endf \ const result = deep(2)`, 777},
		{`deep func arg \ if arg > 1 \ if arg > 2 \ if arg > 3 \ return 999 \ endif \ return 888 \ endif \  return 777 \ endif \ endf \ const result = deep(3)`, 888},
		{`deep func arg \ if arg > 1 \ if arg > 2 \ if arg > 3 \ return 999 \ endif \ return 888 \ endif \  return 777 \ endif \ endf \ const result = deep(4)`, 999},
		{`	fib func x
					if x < 2
						return 1
					else
						return fib(x - 1) + fib(x - 2)
					endif
				endf
			const result = fib(5)
			`, 8},
	}
	// t.Fatal("const 再定義要修正")

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		_ = evaluateInput(t, tt.input, logger, env)

		obj, ok := env.Get("RESULT")
		if !ok {
			t.Fatalf(`[%d] "RESULT" not in env`, tn)
		}
		value := evalValue(obj)

		if tt.expected >= 0 {
			testNumberObject(t, tn, value, tt.expected)
		} else if value != object.NULL {
			t.Errorf("[%d] should be NULL. got %T(%#v)", tn, obj, obj)
		}
	}
}

func TestClosure(t *testing.T) {
	tests := []struct {
		input    string
		expected int
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
		`, 13},
	}

	for tn, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		_ = evaluateInput(t, tt.input, logger, env)

		obj, ok := env.Get("RESULT")
		if !ok {
			t.Fatalf(`[%d]"RESULT" not in env`, tn)
		}

		value := evalValue(obj)

		if tt.expected >= 0 {
			testNumberObject(t, tn, value, tt.expected)
			continue
		}
	}
}

func TestFibFunc(t *testing.T) {
	tests := []struct {
		input    string
		expected int
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
		`, 89},
	}

	for testnum, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		_ = evaluateInput(t, tt.input, logger, env)

		obj, ok := env.Get("RESULT")
		if !ok {
			t.Fatalf(`[%d] "RESULT" not in env`, testnum)
		}
		value := evalValue(obj)

		if tt.expected >= 0 {
			testNumberObject(t, testnum, value, tt.expected)
			continue
		}
	}
}

func TestFunction(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`function vram(x, y) 0xd000 + y * 40 + x \ const result = vram(0, 0)`, 0xd000},
		{`function vram(x, y) 0xd000 + y * 40 + x \ const result = vram(39, 24)`, 0xd000 + 40*24 + 39},
		{`function add1(x) x + 1 \ const fn = add1 \ const result = fn(99)`, 100},
		{`	function add1(x) x + 1
			function double(fn, x) fn(fn(x))
			const result = double(add1, 98)`, 100},
	}

	for testnum, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		_ = evaluateInput(t, tt.input, logger, env)

		obj, ok := env.Get("RESULT")
		if !ok {
			t.Fatalf(`[%d] "RESULT" not in env`, testnum)
		}
		value := evalValue(obj)

		if tt.expected >= 0 {
			testNumberObject(t, testnum, value, tt.expected)
			continue
		}
	}
}
