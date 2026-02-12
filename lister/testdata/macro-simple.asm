tm1 macro arg
  ld a, arg
    tm2 arg
  ld a, arg + 1
endm

tm2 macro arg
  ld hl, arg
  ret
endm

tm1 1

tm1 10
