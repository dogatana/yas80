main.exe: main.go parser/parser.go parser/lexer.go parser/token.go
	go build -o $@

parser/parser.go: parser/parser.y
	goyacc -xegen parser/error.txt -o $@ $<
