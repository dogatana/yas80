; LD (HL), r
ld (hl), a
ld (hl), b
ld (hl), c
ld (hl), d
ld (hl), e
ld (hl), h
ld (hl), l

; LD (IX), r
ld (ix + $12), a
ld (ix + $12), b
ld (ix + $12), c
ld (ix + $12), d
ld (ix + $12), e
ld (ix + $12), h
ld (ix + $12), l

; LD (IY), r
ld (iy - $34), a
ld (iy - $34), b
ld (iy - $34), c
ld (iy - $34), d
ld (iy - $34), e
ld (iy - $34), h
ld (iy - $34), l

ld (bc), a
ld (de), a

; LD (nn), A
ld ($1234), a

ld ($1234), hl
ld ($1234), ix
ld ($1234), iy

ld ($1234), bc
ld ($1234), de
ld ($1234), sp
