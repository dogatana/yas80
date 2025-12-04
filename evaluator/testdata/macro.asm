ld_arg	macro	arg1, arg2
	ld arg1, arg2
	endm
	
	ld_arg	a, a
	ld_arg	a, 255
	ld_arg	hl, $1234

cp_call	macro	value, addr
@next:	
	ld	a, value
	ld	hl, addr
	ld	hl, @next
	nop
	endm
	
	cp_call	1, $1100
	cp_call 2, $2200
	cp_call 3, $3300
	
	