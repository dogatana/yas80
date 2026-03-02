rem big.asm
py measure.py 10 "ailz80asm -bin -f -i big.asm"
py measure.py 10 "java -jar tools80.jar -bin -tgt=z80  big.asm"
py measure.py 10 "z80asm -b big.asm"
py measure.py 10 "..\yas80 big.asm"
py measure.py 10 "z80as big.asm"

rem label.asm
py measure.py 10 "ailz80asm -bin -f -i label.asm"
py measure.py 10 "java -jar tools80.jar -bin -tgt=z80  label.asm"
py measure.py 10 "z80asm -b label.asm"
py measure.py 10 "..\yas80 label.asm"
py measure.py 10 "z80as -bin label.asm"
