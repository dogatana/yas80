t1 macro arg 
 t2 arg
endm 

t2 macro arg
 t3 arg
endm

t3 macro arg
  t4 arg
endm

t4 macro arg
  if arg > 0
    ld a, arg
  else
    ld hl, arg
  endif
endm

t1 0
t1 1
