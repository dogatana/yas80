# name.mzt と a_very_long_filename.mzt は別扱い
# 名前の格納形式が ailz80asm と異なるため、手作業で作成
ASM = z80asm
ASMOPT = -b
SRCS = zilog.asm zilog-a.asm ixy-hl.asm

BINS = $(SRCS:.asm=.bin)

ALL: $(BINS)

clean:
	rm $(BINS) $(SRCS:.asm=.o) 

%.bin: %.asm
	$(ASM) $(ASMOPT) $<