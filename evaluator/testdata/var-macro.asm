t1 macro arg
    data ## arg:
        db arg
        dw arg * $1000
endm

var seq = 1
t1 seq \ seq = seq + 1
t1 seq \ seq = seq + 1
t1 seq \ seq = seq + 1
