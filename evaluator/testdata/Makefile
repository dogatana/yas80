ASM = ailz80asm
ASMOPT = -lst -bin -f -i
SRCS = inst0.asm ld_r8_r8.asm forward.asm forward_symbol.asm
BINS = $(SRCS:.asm=.bin)

ALL: $(BINS)

clean:
	rm *.bin

%.bin: %.asm
	$(ASM) $(ASMOPT) $<
