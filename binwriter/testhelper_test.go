package binwriter

import (
	"bytes"
	"errors"
	"fmt"
	"slices"
	"yas80/evaluator"
	"yas80/filecontent"
	"yas80/logging"
	"yas80/object"
	"yas80/parser"
)

func evalInput(input any, logger *logging.Logger, env object.Environment) (*object.BlockObject, *evaluator.Evaluator) {
	prog := parseTextForTest(input, logger)

	eval := evaluator.New(logger, []string{})

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

func parseTextForTest(input any, logger *logging.Logger) *parser.BlockStatement {
	var prog *parser.BlockStatement

	fcs := []*filecontent.FileContent{}
	switch input := input.(type) {
	case string:
		fc, _ := filecontent.NewFromString("text", input)
		fcs = append(fcs, fc)
	case []string:
		for _, s := range input {
			fc, _ := filecontent.NewFromString("text", s)
			fcs = append(fcs, fc)
		}
	}
	index := 0
	callback := func() *filecontent.FileContent {
		if index >= len(fcs) {
			return nil
		}
		fc := fcs[index]
		index++
		return fc
	}

	lex := parser.NewLexer(logger, callback)
	prog = parser.Parse(lex, []string{})
	if logger.ErrorCount() > 0 {
		return prog
	}

	// prog = parser.PreProrocess(logger, prog)
	return prog
}

func binFromObj(obj object.Object, fill int, logger *logging.Logger) ([]byte, bool) {
	bw := New(obj, fill, logger)

	var buf bytes.Buffer
	result := bw.WriteBin(&buf)

	return buf.Bytes(), result
}

func mztFromObj(obj object.Object, fill int, logger *logging.Logger, name string, start int) ([]byte, bool) {
	bw := New(obj, fill, logger)

	var buf bytes.Buffer
	result := bw.WriteMzt(&buf, name, start)
	return buf.Bytes(), result
}

func t88FromObj(obj object.Object, fill int, logger *logging.Logger, name string) (string, int, error) {
	bw := New(obj, fill, logger)

	var buf bytes.Buffer
	ok := bw.WriteT88(&buf, name)
	if !ok {
		return "", -1, errors.New("WriteT88 error")
	}
	data := buf.Bytes()

	var t88name string
	if string(data[:24]) != "PC-8801 Tape Image(T88)\x00" {
		return "", -1, errors.New("no T88 magic")
	}
	ofs := 24 // skip magic
	for ofs < len(data) {
		id, err := bytesToInt(data[ofs:ofs+2], 2)
		if err != nil {
			goto ERROR
		}
		ofs += 2
		size, err := bytesToInt(data[ofs:ofs+2], 2)
		if err != nil {
			goto ERROR
		}
		ofs += 2
		tagData := data[ofs : ofs+size]
		ofs += size
		if id != 0x0101 {
			continue
		}
		t88name = string(tagData[12:21])
		break
	}
	for ofs < len(data) {
		id, err := bytesToInt(data[ofs:ofs+2], 2)
		if err != nil {
			goto ERROR
		}
		ofs += 2
		size, err := bytesToInt(data[ofs:ofs+2], 2)
		if err != nil {
			goto ERROR
		}
		ofs += 2
		tagData := data[ofs : ofs+size]
		ofs += size
		if id != 0x0101 {
			continue
		}
		if tagData[12] != ':' {
			return "", -1, errors.New("invalid Load Address Record")
		}
		addr := int(tagData[13])*256 + int(tagData[14])
		return t88name, addr, nil
	}
ERROR:
	return "", -1, errors.New("invalid T88")
}

func bytesToInt(b []byte, size int) (int, error) {
	if len(b) != size {
		return 0, fmt.Errorf("len([]byte) must be %d. got %d", size, len(b))
	}
	data := slices.Clone(b)
	slices.Reverse(data)

	out := 0
	for _, b := range data {
		out = out*256 + int(b)
	}
	return out, nil
}
