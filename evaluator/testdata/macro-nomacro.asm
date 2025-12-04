; ld_arg	macro	arg1, arg2
; 	ld arg1, arg2
; 	endm
; 	
; 	ld_arg	a, a
; 	ld_arg	a, 255
; 	ld_arg	hl, $1234
; 
; cp_call	macro	value, addr
; @next:	
; 	cp	value
; 	ld	hl, addr
; 	ld	hl, @next
; 	nop
; 	endm
; 	
; 	cp_call	1, $1100
; 	cp_call 2, $2200
; 	cp_call 3, $3300
; 	
	ld a, a
	ld a, 255
	ld hl, $1234
	
next1:	ld a, 1
	ld hl, $1100
	ld hl, next1
	nop

next2:	ld a, 2
	ld hl, $2200
	ld hl, next2
	nop

next3:	ld a, 3
	ld hl, $3300
	ld hl, next3
	nop

	