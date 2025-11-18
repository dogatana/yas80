; 前方参照（ラベル, 
abc 	equ	end_addr

	ld	hl, end_addr
	ld	hl, abc
	ld	hl, size
start:	ld	a, 1
	ld	a, 2
	ld	a, 3
	ret
end_addr:

size	equ	$ - start	
