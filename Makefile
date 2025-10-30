SRC = main.go parser/token.go parser/lexer.go parser/z80.go \
	  parser/parser.go parser/ast.go parser/helper.go \
	  object/object.go generator/generator.go evaluator/evaluator.go

	  
main.exe: ${SRC}
	go build -o $@

parser/parser.go: parser/parser.txt parser/patch_parser.py
	python parser/patch_parser.py $< $@

parser/parser.txt: parser/parser.y
	goyacc -xe parser/error.txt -v parser/y.output -o $@ $<

test:
	go test ./parser ./generator

testv:
	go test -v ./parser ./generator
