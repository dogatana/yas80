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
    rept [arg, 255]
      ld a, $v
    endr
  else
    rept [arg, $ffff]
      ld hl, $v
    endr
  endif
endm

t1 0
t1 1
