package evaluator

import (
	"fmt"
	"testing"
	"yas80/logger"
	"yas80/object"
)

func TestEvalExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1", 1},
		{"1 + 2", 3},
		{"1 * 2", 2},
		{"1 + 2 * 3", 7},
		{`const val = 123 \ val`, 123},
		{`const val = 11 \ val * val`, 121},
		{`const val = 11 \ const val2 = val * val \ val2`, 121},
		{"const val = 11 \r\n const val2 = val * val \r\n val2", 121},
	}

	for _, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		prog := evaluateInput(t, tt.input, logger, env)

		last := prog.Objects[len(prog.Objects)-1]
		testNumberObject(t, last, tt.expected, tt.input)
	}
}

func TestReturn(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{"1", 1},
		{`1 \ 2 \ 3`, 3},
		{`1 \ 2 \ return \ 3`, -1},
		{`1 \ 2 \ return 99 \ 3`, 99},
		{`1 \ if 1 \ if 2 \ 3 \ return 99 \ 4 \ endif \ return 98 \ 5 \ endif`, 99},
		{`1 \ if 1 \ if 0 \ 3 \ return 99 \ 4 \ endif \ return 98 \ 5 \ endif`, 98},
	}

	for _, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		prog := evaluateInput(t, tt.input, logger, env)

		last := prog.Objects[len(prog.Objects)-1]
		if tt.expected >= 0 {
			testNumberObject(t, last, tt.expected, tt.input)
			continue
		}
		if last != object.NULL {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("should be NULL. got %T(%#v)", last, last)
		}
	}
}

func TestIf(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`if 1 \ endif`, -1}, // -1 は NullObject とする
		{`if 0 \ endif`, -1},
		{`if 1 \ else \ endif`, -1},
		{`if 0 \ else \ endif`, -1},

		{`if 1 \ 100 \ endif`, 100},
		{`if 0 \ 100 \ endif`, -1},

		{`if 1 \ 100 \ else \ endif`, 100},
		{`if 0 \ 100 \ else \ endif`, -1},

		{`if 1 \ 100 \ else \ 200  \ endif`, 100},
		{`if 0 \ 100 \ else \ 200  \ endif`, 200},

		{`const val = 1 \ if val == 1 \ 100 \ elif val == 2 \ 200  \ endif`, 100},
		{`const val = 2 \ if val == 1 \ 100 \ elif val == 2 \ 200  \ endif`, 200},
		{`const val = 3 \ if val == 1 \ 100 \ elif val == 2 \ 200  \ endif`, -1},

		{`const val = 3 \ if val == 1 \ 100 \ elif val == 2 \ 200  \ else \ 300 \ endif`, 300},

		{`const val = 2 \ if val == 1 \ 100 \ elif val == 2 \ 200 \ return 999 \ 250  \ else \ 300 \ endif`, 999},
		{`const val = 2 \ if val == 1 \ 100 \ elif val == 2 \ 200 \ return 999 \ 250  \ else \ 300 \ endif`, 999},
	}

	for _, tt := range tests {
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		prog := evaluateInput(t, tt.input, logger, env)

		last := prog.Objects[len(prog.Objects)-1]
		if tt.expected >= 0 {
			testNumberObject(t, last, tt.expected, tt.input)
			continue
		}
		if last != object.NULL {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("should be NULL. got %T(%#v)", last, last)
		}
	}
}

func TestFunc(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{`retNULL func \ 100 \ endf \ RETNULL()`, -1}, // -1 は NullObject とする
		{`ret100 func \ return 100 \ endf \ ret100()`, 100},
		{`ret100 func \ 1 \ 2 \ return 100 \ 4 \ endf \ ret100()`, 100},
		{`abs func arg \ 1 \ return arg \ 2 \ endf \  abs(123)`, 123},
		{`abs func arg \ if arg > 0 \ return arg \ else \ return -arg \ endif \ endf \ abs(100)`, 100},
		{`abs func arg \ if arg > 0 \ return arg \ else \ return -arg \  endif \endf \ abs(-100)`, 100},
		{`deep func arg \ if arg > 1 \ if arg > 2 \ if arg > 3 \ return 999 \ endif \ return 888 \ endif \  return 777 \ endif \ endf \ deep(1)`, -1},
		{`deep func arg \ if arg > 1 \ if arg > 2 \ if arg > 3 \ return 999 \ endif \ return 888 \ endif \  return 777 \ endif \ endf \ deep(2)`, 777},
		{`deep func arg \ if arg > 1 \ if arg > 2 \ if arg > 3 \ return 999 \ endif \ return 888 \ endif \  return 777 \ endif \ endf \ deep(3)`, 888},
		{`deep func arg \ if arg > 1 \ if arg > 2 \ if arg > 3 \ return 999 \ endif \ return 888 \ endif \  return 777 \ endif \ endf \ deep(4)`, 999},
		{`	fib func x
					if x < 2
						return 1
					else
						return fib(x - 1) + fib(x - 2)
					endif
				endf
			fib(5)
			`, 8},
	}

	for _, tt := range tests {
		fmt.Println("input", tt.input)
		env := object.NewEnvironment(nil)
		logger := logger.New("<eval test>")
		prog := evaluateInput(t, tt.input, logger, env)
		env.Print()

		last := prog.Objects[len(prog.Objects)-1]
		if tt.expected >= 0 {
			testNumberObject(t, last, tt.expected, tt.input)
			continue
		}
		if last != object.NULL {
			fmt.Printf("input %q\n", tt.input)
			t.Errorf("should be NULL. got %T(%#v)", last, last)
		}
	}
}
