rem min.asm
python measure.py 10 "yas80 min.asm"
python measure.py 10 "z80as min.asm"
python measure.py 10 "z80asm -b min.asm"
python measure.py 10 "java -jar tools80.jar -bin -tgt=z80  min.asm"
python measure.py 10 "ailz80asm -bin -f -i min.asm"

rem max.asm
python measure.py 10 "yas80 max.asm"
python measure.py 10 "z80as max.asm"
python measure.py 10 "z80asm -b max.asm"
python measure.py 10 "java -jar tools80.jar -bin -tgt=z80  max.asm"
python measure.py 10 "ailz80asm -bin -f -i max.asm"

rem label.asm
python measure.py 10 "yas80 label.asm"
python measure.py 10 "z80as label.asm"
python measure.py 10 "z80asm -b label.asm"
python measure.py 10 "java -jar tools80.jar -bin -tgt=z80  label.asm"
python measure.py 10 "ailz80asm -bin -f -i label.asm"
