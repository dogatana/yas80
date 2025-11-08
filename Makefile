SRC = main.go parser/token.go parser/lexer.go parser/z80.go \
	  parser/parser.go parser/ast.go parser/helper.go \
	  parser/modifyyaccerror.go \
	  object/object.go  evaluator/evaluator.go \
	  logger/logger.go

TEMP = parser/temp.go
YACC = parser/parser.y
PARSER = parser/parser.go
PATCH = parser/patch_parser.py
	  
main.exe: ${SRC}
	go build -o $@

${PARSER}: ${YACC} ${PATCH}
	goyacc -v parser/y.output -o $@ ${YACC}
	python parser/patch_parser.py $@ $@

yacc:
	goyacc -v parser/y.output -o ${PARSER} ${YACC}
	python parser/patch_parser.py ${PARSER} ${PARSER}

vet:
	go vet ./parser ./evaluator

check:
	staticcheck ./parser ./evaluator

test:
	go test ./parser ./evaluator

testv:
	go test -v ./parser ./evaluator
