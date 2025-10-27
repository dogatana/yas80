main.exe: main.go parser/parser.go parser/lexer.go parser/token.go parser/z80.go parser/ast.go parser/helper.go
	go build -o $@

parser/parser.go: parser/parser.txt parser/patch_parser.py
	python parser/patch_parser.py $< $@

parser/parser.txt: parser/parser.y
	goyacc -xegen parser/error.txt -v parser/y.output -o $@ $<
