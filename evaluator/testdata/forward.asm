; 前方参照（ラベル, 
fw_sym 	equ	end_addr

	ld	hl, end_addr
	ld	hl, fw_sym
	ld	hl, size
start:	ld	a, 1
	ld	a, 2
	ld	a, 3
	ret
end_addr:

size	equ	$ - start	
