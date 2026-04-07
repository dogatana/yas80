rem min.asm
py measure.py 100 "ailz80asm -bin -f -i min.asm"
py measure.py 100 "java -jar tools80.jar -bin -tgt=z80  min.asm"
py measure.py 100 "z80asm -b min.asm"
py measure.py 100 "yas80 min.asm"
py measure.py 100 "z80as min.asm"

rem max.asm
py measure.py 100 "ailz80asm -bin -f -i max.asm"
py measure.py 100 "java -jar tools80.jar -bin -tgt=z80  max.asm"
py measure.py 100 "z80asm -b max.asm"
py measure.py 100 "yas80 max.asm"
py measure.py 100 "z80as max.asm"

rem label.asm
py measure.py 100 "ailz80asm -bin -f -i label.asm"
py measure.py 100 "java -jar tools80.jar -bin -tgt=z80  label.asm"
py measure.py 100 "z80asm -b label.asm"
py measure.py 100 "yas80 label.asm"
py measure.py 100 "z80as label.asm"
