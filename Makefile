SRC = main.go parser/token.go parser/lexer.go parser/z80.go \
	  parser/parser.go parser/ast.go parser/helper.go parser/rules.go \
	  parser/modifyyaccerror.go \
	  object/object.go  \
	  evaluator/evaluator.go evaluator/eval_instruction.go evaluator\helper.go \
	  logger/logger.go logger/error_messages.go

TEMP = parser/temp.go
YACC = parser/parser.y
PARSER = parser/parser.go
PATCH = parser/patch_parser.py
YOUT = parser/y.output
	  
main.exe: ${SRC}
	go build -o $@

${PARSER}: ${YACC} ${PATCH}
	goyacc -v ${YOUT} -o $@ ${YACC}
	python ${PATCH} ${PARSER} ${PARSER} ${YOUT}

yacc:
	goyacc -v ${YOUT} -o ${PARSER} ${YACC}
	python ${PATCH} ${PARSER} ${PARSER} ${YOUT}

vet:
	go vet ./parser ./evaluator

check:
	staticcheck ./parser ./evaluator

test:
	go test ./parser ./evaluator

testv:
	go test -v ./parser ./evaluator
