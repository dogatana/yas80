SRC = main.go \
	assembler/assembler.go \
	binwriter/binwriter.go \
	errcode/errcode.go \
	evaluator/check_symbol.go \
	evaluator/eval_builtin_func.go \
	evaluator/eval_builtin_macro.go \
	evaluator/eval_charmap.go \
	evaluator/eval_data_statement.go \
	evaluator/eval_expression.go \
	evaluator/eval_instruction.go \
	evaluator/eval_macro.go \
	evaluator/eval_statement.go \
	evaluator/eval_z80_add16.go \
	evaluator/eval_z80_add8.go \
	evaluator/eval_z80_bit.go \
	evaluator/eval_z80_ex_im.go \
	evaluator/eval_z80_inc.go \
	evaluator/eval_z80_io.go \
	evaluator/eval_z80_jp_call.go \
	evaluator/eval_z80_ld.go \
	evaluator/eval_z80_mul.go \
	evaluator/eval_z80_rlc.go \
	evaluator/evaluator.go \
	evaluator/expand_macro.go \
	evaluator/helper.go \
	evaluator/types.go \
	evaluator/z80code.go \
	filecontent/context.go \
	filecontent/filecontent.go \
	intern/intern.go \
	internal/util/util.go \
	lister/lister.go \
	logging/logging.go \
	object/environment.go \
	object/object.go \
	object/symbol.go \
	object/z80register_flag.go  \
	options/options.go \
	parser/ast.go \
	parser/ast_node_literal.go \
	parser/helper.go \
	parser/lexer.go \
	parser/parser.go \
	parser/preprocess.go \
	parser/rules.go \
	parser/token.go  \
	parser/z80.go 

TEMP = parser/temp.go
YACC = parser/parser.y
PARSER = parser/parser.go
PATCH = parser/patch_parser.py
YOUT = parser/y.output
TESTDIR = ./filecontent ./parser ./evaluator ./binwriter ./lister .

	  
yas80.exe: ${SRC} 
	go build

clean:
	rm yas80.exe
	rm parser/parser.go

${PARSER}: ${YACC} ${PATCH}
	goyacc -v ${YOUT} -o $@ ${YACC}
	python ${PATCH} ${PARSER} ${PARSER} ${YOUT}

yacc:
	goyacc -v ${YOUT} -o ${PARSER} ${YACC}
	python ${PATCH} ${PARSER} ${PARSER} ${YOUT}

vet: fmt
	go vet ./...

fmt:
	go fmt ./...

check:
	staticcheck ./...

tc:
	python internal/testutil/errcode_coverage.py

test: internal/errcodenames/errcode_names.go
	go test ${TESTDIR}

testv: internal/errcodenames/errcode_names.go
	go test -v ${TESTDIR}

internal/errcodenames/errcode_names.go: errcode/errcode.go
	python internal/errcodenames/errcode_names.py $< $@
	go fmt $@
	
bench:
	go test -bench Run

prof:
	main --cpuprofile cpu.prof performance\label.asm -o out.bin
	go tool pprof -top cpu.prof
	main --memprofile mem.prof performance\label.asm -o out.bin
	go tool pprof -top mem.prof
