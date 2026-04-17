package parserbench

import (
	"bytes"
	"testing"

	"github.com/dogatana/yas80/evaluator"
	"github.com/dogatana/yas80/filecontent"
	"github.com/dogatana/yas80/logging"
	"github.com/dogatana/yas80/object"
	"github.com/dogatana/yas80/parser"
)

var src string

func init() {
	// src = GenerateLargeSource(100000)
	src = GenerateLargeSource(30000)
}

func BenchmarkParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		logger := logging.New()
		fc, _ := filecontent.NewFromString("exlarge.asm", src)

		lex := parser.NewLexer(logger, func() *filecontent.FileContent {
			ret := fc
			fc = nil
			return ret
		})
		prog := parser.Parse(lex, []string{})
		if logger.ErrorCount() > 0 {
			b.Fatalf("parsing failed with %d errors", logger.ErrorCount())
		}

		env := object.NewEnvironment(nil)
		eval := evaluator.New(logger, []string{})
		pass := 0
		var obj object.Object

		for i := 0; i < 256; i++ {
			pass++
			eval.Resolved = true
			obj = eval.EvalProgram(prog, pass, env)
			if logger.ErrorCount() > 0 {
				logger.Print()
				b.Fatalf("eval stage 1 failed with %d errors", logger.ErrorCount())
			}
			if eval.Resolved {
				break
			}
		}
		if !eval.Resolved {
			b.Fatalf("eval stage 1 failed Unresolved. with %d pass", pass)
		}
		code := object.CollectCode(obj.(*object.BlockObject).Block)
		for i := 0; i < 16; i++ {
			pass++
			obj = eval.EvalProgram(prog, pass, env)
			if logger.ErrorCount() > 0 {
				b.Fatalf("eval stage 2 failed with %d errors", logger.ErrorCount())
			}
			newCode := object.CollectCode(obj.(*object.BlockObject).Block)
			eval.CodeStable = bytes.Equal(code, newCode)
			if eval.CodeStable {
				break
			}
			code = newCode
		}
		if !eval.CodeStable {
			b.Fatalf("eval stage 2 failed CodeStable. with %d pass", pass)
		}
	}
}
