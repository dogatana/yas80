package lister

import (
	"bytes"
	"fmt"
	"os"
	"strings"
	"yas80/evaluator"
	"yas80/filecontent"
	"yas80/internal/util"
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

func evalFile(filename string, logger *logging.Logger, env object.Environment) []byte {
	prog := parseFileForTest(filename, logger)
	obj, _ := evalProg(prog, logger, env)

	lister := New(prog, obj)

	var buf bytes.Buffer
	lister.ProgramList(&buf)
	return buf.Bytes()

	// bw := binwriter.New(obj, 0, logger)
	// var buf bytes.Buffer
	// ok := bw.Write(&buf)
	// return buf.Bytes(), ok
}

func evalProg(prog *parser.BlockStatement, logger *logging.Logger, env object.Environment) (*object.BlockObject, *evaluator.Evaluator) {

	eval := evaluator.New(logger)

	var obj object.Object
	pass := 0

	eval.Resolved = true
	for i := 0; i < 256; i++ {
		pass++
		eval.Resolved = true
		obj = eval.EvalProgram(prog, pass, env)
		// eval.EvalEnv(env)
		eval.CheckCyclicError(env)
		if logger.ErrorCount() > 0 {
			return &object.BlockObject{}, eval
		}
		if eval.Resolved {
			break
		}
	}
	eval.CheckSymbolError(env)
	if logger.ErrorCount() > 0 || !eval.Resolved {
		return obj.(*object.BlockObject), eval
	}

	// finalize
	code := object.CollectCode(obj.(*object.BlockObject).Block)
	eval.CodeStable = false
	for i := 0; i < 256 && !eval.CodeStable; i++ {
		pass++
		obj = eval.EvalProgram(prog, pass, env)
		if logger.ErrorCount() > 0 {
			return obj.(*object.BlockObject), eval

		}
		newCode := object.CollectCode(obj.(*object.BlockObject).Block)
		eval.CodeStable = bytes.Equal(code, newCode)
		if !eval.CodeStable {
			code = newCode
		}
	}
	if !eval.CodeStable {
		return obj.(*object.BlockObject), eval
	}
	if prog, ok := obj.(*object.BlockObject); !ok {
		return prog, eval
	} else {
		return prog, eval
	}
}

func parseFileForTest(filename string, logger *logging.Logger) *parser.BlockStatement {
	var prog *parser.BlockStatement

	fc, err := filecontent.NewFromFile(filename)
	if err != nil {
		fmt.Println("parseFileForTest", err.Error())
	}

	lex := parser.NewLexer(logger, func() *filecontent.FileContent {
		ret := fc
		fc = nil
		return ret
	})
	prog = parser.Parse(lex)
	if logger.ErrorCount() > 0 {
		return prog
	}

	prog = parser.PreProrocess(logger, prog)
	return prog
}

// リストファイル読み込み
func readListFile(filename string) ([]byte, error) {
	data, err := os.ReadFile(filename)
	return data, err
}

// []byte 内の CR/LF を LF に置換
func replaceCRLF(in []byte) []byte {
	return bytes.ReplaceAll(in, []byte{0x0d, 0x0a}, []byte{0xa})
}

func linesEqual(a, b []byte) error {
	var err error
	al := strings.Split(string(replaceCRLF(a)), "\n")
	bl := strings.Split(string(replaceCRLF(b)), "\n")

	// 文字列末尾の空白（タブ、スペース）を削除
	trim := func(s string) string { return strings.TrimRight(s, "\t ") }
	al = util.Map(al, trim)
	bl = util.Map(al, trim)

	if len(al) != len(bl) {
		err = fmt.Errorf("line count %d %d", len(al), len(bl))
	}
	for i := range min(len(al), len(bl)) {
		if al[i] != bl[i] {
			fmt.Printf("a[%d] %s\n", i, al[i])
			fmt.Printf("b[%d] %s\n", i, al[i])
			err = fmt.Errorf("line %d\n%s\n%s", i, al[i], bl[i])
		}
	}
	if err != nil {
		os.WriteFile("out.txt", a, 0o644)
	}
	return err
}
