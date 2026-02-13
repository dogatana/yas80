rept 3
  ld a, $i
  exitm if $i == $count - 1
  ld hl, $i
endr
