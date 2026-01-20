ASM = ailz80asm
ASMOPT = -lst -bin -f -i
SRCS = inst0.asm ld_r8_r8.asm \
		label-backward.asm equ-backward.asm \
		label-forward.asm equ-forward.asm \
		call-ret.asm io.asm \
		forward.asm forward_symbol.asm forward_mix.asm
BINS = $(SRCS:.asm=.bin)

ALL: $(BINS) macro.bin

clean:
	rm *.bin

%.bin: %.asm
	$(ASM) $(ASMOPT) $<

macro.bin : macro-nomacro.asm
	$(ASM) -bin $@ -lst -f -i $<

