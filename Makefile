SRC = main.go parser/token.go parser/lexer.go parser/z80.go \
	  parser/parser.go parser/ast.go parser/helper.go \
	  object/object.go  evaluator/evaluator.go

ERR = parser/error.txt
TEMP = parser/temp.go
YACC = parser/parser.y
PARSER = parser/parser.go
	  
main.exe: ${SRC}
	go build -o $@

${PARSER}: ${YACC} ${ERR}
	goyacc -xe ${ERR} -v parser/y.output -o $@ ${YACC}
	python parser/patch_parser.py $@ $@

xegen:
	goyacc -xegen ${ERR} -v parser/y.output -o ${PARSER}  ${YACC}
	python parser/patch_parser.py ${PARSER} ${PARSER}

test:
	go test ./parser ./evaluator

testv:
	go test -v ./parser ./evaluator
