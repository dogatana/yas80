main.exe: main.go parser/parser.go parser/lexer.go parser/token.go parser/z80.go parser/ast.go
	go build -o $@

parser/parser.go: parser/parser.y
	goyacc -xegen parser/error.txt -v parser/y.output -o $@ $<
