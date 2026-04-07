rem min.asm
py measure.py 10 "yas80 min.asm"
py measure.py 10 "z80as min.asm"
py measure.py 10 "z80asm -b min.asm"
py measure.py 10 "java -jar tools80.jar -bin -tgt=z80  min.asm"
py measure.py 10 "ailz80asm -bin -f -i min.asm"

rem max.asm
py measure.py 10 "yas80 max.asm"
py measure.py 10 "z80as max.asm"
py measure.py 10 "z80asm -b max.asm"
py measure.py 10 "java -jar tools80.jar -bin -tgt=z80  max.asm"
py measure.py 10 "ailz80asm -bin -f -i max.asm"

rem label.asm
py measure.py 10 "yas80 label.asm"
py measure.py 10 "z80as label.asm"
py measure.py 10 "z80asm -b label.asm"
py measure.py 10 "java -jar tools80.jar -bin -tgt=z80  label.asm"
py measure.py 10 "ailz80asm -bin -f -i label.asm"
