    LD A, IXH
    LD A, IXL
    LD B, IXH
    LD B, IXL
    LD C, IXH
    LD C, IXL
    LD D, IXH
    LD D, IXL
    LD E, IXH
    LD E, IXL

    LD IXH, A
    LD IXH, B
    LD IXH, C
    LD IXH, D
    LD IXH, E
    LD IXH, IXH
    LD IXH, IXL

    LD IXL, A
    LD IXL, B
    LD IXL, C
    LD IXL, D
    LD IXL, E
    LD IXL, IXH
    LD IXL, IXL

    LD IXH, $12
    LD IXL, $12

    LD A, IYH
    LD A, IYL
    LD B, IYH
    LD B, IYL
    LD C, IYH
    LD C, IYL
    LD D, IYH
    LD D, IYL
    LD E, IYH
    LD E, IYL

    LD IYH, A
    LD IYH, B
    LD IYH, C
    LD IYH, D
    LD IYH, E
    LD IYH, IYH
    LD IYH, IYL

    LD IYL, A
    LD IYL, B
    LD IYL, C
    LD IYL, D
    LD IYL, E
    LD IYL, IYH
    LD IYL, IYL

    LD IYH, $12
    LD IYL, $12

    ADD A, IXH
    ADD A, IXL
    ADD A, IYH
    ADD A, IYL

    ADC A, IXH
    ADC A, IXL
    ADC A, IYH
    ADC A, IYL

    SUB IXH
    SUB IXL
    SUB IYH
    SUB IYL

    SBC A, IXH
    SBC A, IXL
    SBC A, IYH
    SBC A, IYL

    AND IXH
    AND IXL
    AND IYH
    AND IYL

    OR IXH
    OR IXL
    OR IYH
    OR IYL

    XOR IXH
    XOR IXL
    XOR IYH
    XOR IYL

    CP IXH
    CP IXL
    CP IYH
    CP IYL

    INC IXH
    INC IXL
    INC IYH
    INC IYL

    DEC IXH
    DEC IXL
    DEC IYH
    DEC IYL

; with A
    SUB A, IXH
    SUB A, IXL
    SUB A, IYH
    SUB A, IYL

; ailz80asm はアセンブルエラーになる
;    AND A, IXH
;    AND A, IXL
;    AND A, IYH
;    AND A, IYL
;
;    OR A, IXH
;    OR A, IXL
;    OR A, IYH
;    OR A, IYL
;
;    XOR A, IXH
;    XOR A, IXL
;    XOR A, IYH
;    XOR A, IYL
;
;    CP A, IXH
;    CP A, IXL
;    CP A, IYH
;    CP A, IYL

