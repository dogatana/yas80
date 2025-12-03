SRC = main.go \
	parser/token.go parser/lexer.go parser/z80.go \
	parser/parser.go parser/ast.go parser/ast_node_literal.go \
	parser/helper.go parser/rules.go \
	parser/preprocess.go \
	object/object.go  object/symbol.go object/z80register_flag.go object/environment.go \
	evaluator/evaluator.go evaluator/eval_instruction.go evaluator/eval_env.go \
	evaluator/z80code.go evaluator/helper.go \
	logger/logger.go \
	errcode/errcode.go \
	fileblock/fileblock.go

TEMP = parser/temp.go
YACC = parser/parser.y
PARSER = parser/parser.go
PATCH = parser/patch_parser.py
YOUT = parser/y.output
	  
main.exe: ${SRC}
	go build -o $@

clean:
	rm main.exe
	rm parser/parser.go

${PARSER}: ${YACC} ${PATCH}
	goyacc -v ${YOUT} -o $@ ${YACC}
	python ${PATCH} ${PARSER} ${PARSER} ${YOUT}

yacc:
	goyacc -v ${YOUT} -o ${PARSER} ${YACC}
	python ${PATCH} ${PARSER} ${PARSER} ${YOUT}

vet:
	go vet ./parser ./evaluator ./fileblock

check:
	staticcheck ./parser ./evaluator

test:
	go test ./parser ./evaluator ./fileblock

testv:
	go test -v ./parser ./evaluator ./fileblock
