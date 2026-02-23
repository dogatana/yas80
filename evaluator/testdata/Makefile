ASM = ailz80asm
ASMOPT = -lst -bin -f -i
SRCS = inst0.asm \
		label-backward.asm equ-backward.asm \
		label-forward.asm equ-forward.asm \
		inst-ld8.asm \
		inst-ld16.asm \
		inst-ldind.asm \
		inst-add8.asm \
		inst-add16.asm \
		inst-inc-push.asm \
		inst-ex-im.asm \
		inst-rlc.asm \
		inst-bit.asm \
		inst-call-ret.asm \
		inst-io.asm \
		forward.asm forward_symbol.asm forward_mix.asm \
		include.asm \
		zilog.asm \
		zilog-a.asm \
		ixy-hl.asm

BINS = $(SRCS:.asm=.bin)

ALL: $(BINS) macro.bin

clean:
	rm $(BINS) $(SRCS:.asm=.lst)

%.bin: %.asm
	$(ASM) $(ASMOPT) $<

macro.bin : macro-nomacro.asm
	$(ASM) -bin $@ -lst -f -i $<

inst-add16.bin: inst-add16.asm
	z80asm -l -b $<

zilog.bin: zilog.asm
	z80asm -l -b $<

zilog-a.bin: zilog-a.asm
	z80asm -l -b $<

ixy-hl.bin: ixy-hl.asm
	z80asm -l -b $<
