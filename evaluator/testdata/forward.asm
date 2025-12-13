; 前方参照（ラベル, 
fw_sym 	equ	end_addr

	ld	a, end_addr
	ld	a, fw_sym
	ld	a, size
start:	ld	a, 1
	ld	a, 2
	ld	a, 3
	ret
end_addr:

size	equ	$ - start	
