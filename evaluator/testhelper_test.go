package evaluator

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"testing"
	"yas80/errtest"
	"yas80/fileblock"
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

type SymValue struct {
	name     string
	Expected int
}

func evalInput(input string, logger *logging.Logger, env object.Environment) (*object.ProgramObject, *Evaluator) {
	progNode := parseTextForTest(input, logger)

	eval := New(logger)

	var obj object.Object
	var i int

	eval.Resolved = true
	for i = 0; i < 256; i++ {
		eval.Resolved = true
		obj = eval.EvalProgram(progNode, env)
		// eval.EvalEnv(env)
		eval.CheckCyclicError(env)
		if len(logger.Errors) > 0 {
			return &object.ProgramObject{}, eval
		}
		if eval.Resolved {
			break
		}
	}
	eval.CheckSymbolError(env)
	if len(logger.Errors) > 0 || !eval.Resolved {
		return &object.ProgramObject{}, eval
	}

	// finalize
	code := CollectCode(obj.(*object.ProgramObject))
	eval.CodeStable = false
	for i = 0; i < 256 && !eval.CodeStable; i++ {
		obj = eval.EvalProgram(progNode, env)
		if len(logger.Errors) > 0 {
			return &object.ProgramObject{}, eval

		}
		newCode := CollectCode(obj.(*object.ProgramObject))
		eval.CodeStable = bytes.Equal(code, newCode)
		if !eval.CodeStable {
			code = newCode
		}
	}
	if !eval.CodeStable {
		return &object.ProgramObject{}, eval
	}
	if prog, ok := obj.(*object.ProgramObject); !ok {
		return &object.ProgramObject{}, eval
	} else {
		return prog, eval
	}
}

func testEvalResult(t *testing.T, tn int, err string, eval *Evaluator) {
	if err == "" {
		if len(eval.logger.Errors) > 0 {
			eval.logger.Print()
			t.Fatalf("[%d] EvalProgram() %d errors", tn, len(eval.logger.Errors))
		}
		if !eval.Resolved {
			t.Fatalf("[%d] eval.Resolved fasle", tn)
		}
		if !eval.CodeStable {
			t.Fatalf("[%d] eval.CodeStable fasle", tn)
		}
	} else {
		e, w, i := eval.logger.Count()
		if e == 0 && w == 0 && i == 0 {
			fmt.Printf("[%d] logger.Print()\n", tn)
			eval.logger.Print()
			t.Fatalf("[%d] no logmessages", tn)
		}
	}
}

func parseTextForTest(input string, logger *logging.Logger) *parser.Program {
	var prog *parser.Program

	fb := fileblock.New(logger.Filename, []byte(input))

	l := parser.NewLexer(fb, logger)
	prog = parser.Parse(l)
	if len(logger.Errors) > 0 {
		return prog
	}

	prog = parser.PreProrocess(logger, prog)
	return prog
}

func testCodeResult(t *testing.T, tn int, expected []byte, prog *object.ProgramObject) {
	if len(expected) == 0 {
		return
	}
	code := CollectCode(prog)
	if err := bytesEqual(code, expected); err != nil {
		t.Errorf("[%d] generated code diff %s", tn, err.Error())
	}
}

func testSymValues(t *testing.T, tn int, syms []SymValue, getter func(name string) (*object.SymbolObject, bool)) {
	for _, sym := range syms {
		s, ok := getter(sym.name)
		if !ok {
			t.Errorf("[%d] symbol %s not found", tn, sym.name)
		} else {
			testSymbolNumberObject(t, tn, s, sym.Expected)
		}
	}

}

// sym.Expected < 0: 存在するとエラー
// sym.Expected >= 0 : 存在しないとエラー
func testSymValuesEx(t *testing.T, tn int, syms []SymValue, getter func(name string) (*object.SymbolObject, bool)) {
	for _, sym := range syms {
		s, ok := getter(sym.name)
		if sym.Expected >= 0 && !ok {
			t.Errorf("[%d] symbol %s not found", tn, sym.name)
		} else if sym.Expected >= 0 {
			testSymbolNumberObject(t, tn, s, sym.Expected)
		} else if ok {
			t.Errorf("[%d] symbol %s exists", tn, sym.name)
		}
	}
}

func testLogMessage(t *testing.T, tn int, err string, logger *logging.Logger) {
	ename := errtest.ErrcodeNames[err]

	var msgs []logging.LogMessage
	switch ename[0] {
	case 'E':
		msgs = logger.Errors
	case 'W':
		msgs = logger.Warnings
	case 'I':
		msgs = logger.Infomation
	}

	if !hasMessage(msgs, err) {
		t.Errorf("[%d] not [%s] \"%s\" => \"%s\"",
			tn,
			ename,
			err,
			msgs[0])
	}
}

func hasMessage(messages []logging.LogMessage, expected string) bool {
	re := regexp.MustCompile(`\.?%.\.?`)
	ss := re.Split(expected, -1)

	for _, emsg := range messages {
		result := true
		for _, s := range ss {
			if !strings.Contains(emsg.Message(), s) {
				result = false
				break
			}
		}
		if result {
			return result
		}
	}
	return false
}

func evalValue(obj object.Object) object.Object {
	switch obj := obj.(type) {
	case *object.SymbolObject:
		return obj.Value
	default:
		return obj
	}
}

func testSymbolNumberObject(t *testing.T, tn int, obj object.Object, expected int) bool {
	sym, ok := obj.(*object.SymbolObject)
	if !ok {
		t.Errorf("[%d] Object not SymbolObject. got %T", tn, obj)
		return false
	}
	return testNumberObject(t, tn, sym.Value, expected)
}

func testNumberObject(t *testing.T, tn int, obj object.Object, expected int) bool {
	number, ok := obj.(*object.NumberObject)
	if !ok {
		t.Errorf("[%d] Object not NumberObject. got %T", tn, obj)
		return false
	}
	if number.Value != expected {
		t.Errorf("[%d] object is not %d. got %d", tn, expected, number.Value)
		return false
	}
	return true
}

func readTestDataFile(t *testing.T, filename string) []byte {
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatal("runtime.Caller(0) failed")
	}
	path := filepath.Join(filepath.Dir(file), "testdata", filename)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatalf("ReadFile(%q) returns %s", path, err.Error())
	}
	return data
}

func collectValue(prog *object.ProgramObject) []*object.ValueObject {
	var result []*object.ValueObject
	for _, obj := range prog.Objects {
		value, ok := obj.(*object.ValueObject)
		if !ok {
			continue
		}
		result = append(result, value)
	}
	return result
}

func bytesEqual(v1, v2 []byte) error {
	if len(v1) != len(v2) {
		return fmt.Errorf("size diff %d %d\n", len(v1), len(v2))
	}
	for i, v := range v1 {
		if v != v2[i] {
			return fmt.Errorf("contentis diff [%d] %02x %02x\n", i, v, v2[i])
		}
	}
	return nil
}
