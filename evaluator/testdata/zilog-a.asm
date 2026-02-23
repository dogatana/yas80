; A は本来は指定不要だが許容
; SUB
    SUB A,A
    SUB A,B
    SUB A,C
    SUB A,D
    SUB A,E
    SUB A,H
    SUB A,L
    SUB A,$12
    SUB A,(HL)
    SUB A,(IX + 127)
    SUB A,(IY - 128)
    
; AND
    AND A,A
    AND A,B
    AND A,C
    AND A,D
    AND A,E
    AND A,H
    AND A,L
    AND A,$12
    AND A,(HL)
    AND A,(IX + 127)
    AND A,(IY - 128)
    
; AND
    AND A,A
    AND A,B
    AND A,C
    AND A,D
    AND A,E
    AND A,H
    AND A,L
    AND A,$12
    AND A,(HL)
    AND A,(IX + 127)
    AND A,(IY - 128)
    
; OR
    OR A,A
    OR A,B
    OR A,C
    OR A,D
    OR A,E
    OR A,H
    OR A,L
    OR A,$12
    OR A,(HL)
    OR A,(IX + 127)
    OR A,(IY - 128)
    
; XOR
    XOR A,A
    XOR A,B
    XOR A,C
    XOR A,D
    XOR A,E
    XOR A,H
    XOR A,L
    XOR A,$12
    XOR A,(HL)
    XOR A,(IX + 127)
    XOR A,(IY - 128)
    
; CP
    CP A,A
    CP A,B
    CP A,C
    CP A,D
    CP A,E
    CP A,H
    CP A,L
    CP A,$12
    CP A,(HL)

; IX + d, IY + d の d 省略を許容
    CP A,(IX + 127)
    CP A,(IY - 128)
    LD A,(IX)
    LD B,(IX)
    LD C,(IX)
    LD D,(IX)
    LD E,(IX)
    LD H,(IX)
    LD L,(IX)
    
    LD A,(IY)
    LD B,(IY)
    LD C,(IY)
    LD D,(IY)
    LD E,(IY)
    LD H,(IY)
    LD L,(IY)
    
    LD (IX),A
    LD (IX),B
    LD (IX),C
    LD (IX),D
    LD (IX),E
    LD (IX),H
    LD (IX),L
    
    LD (IY),A
    LD (IY),B
    LD (IY),C
    LD (IY),D
    LD (IY),E
    LD (IY),H
    LD (IY),L
    LD (IX),$12
    LD (IY),$12
    ADD A,(IX)
    ADD A,(IY)
    ADC A,(IX)
    ADC A,(IY)
    SUB (IX)
    SUB (IY)
    SBC A,(IX)
    SBC A,(IY)
    AND (IX)
    AND (IY)
    AND (IX)
    AND (IY)
    OR (IX)
    OR (IY)
    XOR (IX)
    XOR (IY)
    CP (IX)
    CP (IY)
    INC (IX)
    INC (IY)
    DEC (IX)
    DEC (IY)
    RLC (IX)
    RLC (IY)
    RL (IX)
    RL (IY)
    RRC (IX)
    RRC (IY)
    RR (IX)
    RR (IY)
    SLA (IX)
    SLA (IY)
    SRA (IX)
    SRA (IY)
    SRL (IX)
    SRL (IY)
    BIT 0,(IX)
    BIT 1,(IX)
    BIT 2,(IX)
    BIT 3,(IX)
    BIT 4,(IX)
    BIT 5,(IX)
    BIT 6,(IX)
    BIT 7,(IX)
    BIT 0,(IY)
    BIT 1,(IY)
    BIT 2,(IY)
    BIT 3,(IY)
    BIT 4,(IY)
    BIT 5,(IY)
    BIT 6,(IY)
    BIT 7,(IY)
    SET 0,(IX)
    SET 1,(IX)
    SET 2,(IX)
    SET 3,(IX)
    SET 4,(IX)
    SET 5,(IX)
    SET 6,(IX)
    SET 7,(IX)
    SET 0,(IY)
    SET 1,(IY)
    SET 2,(IY)
    SET 3,(IY)
    SET 4,(IY)
    SET 5,(IY)
    SET 6,(IY)
    SET 7,(IY)
    RES 0,(IX)
    RES 1,(IX)
    RES 2,(IX)
    RES 3,(IX)
    RES 4,(IX)
    RES 5,(IX)
    RES 6,(IX)
    RES 7,(IX)
    RES 0,(IY)
    RES 1,(IY)
    RES 2,(IY)
    RES 3,(IY)
    RES 4,(IY)
    RES 5,(IY)
    RES 6,(IY)
    RES 7,(IY)
    JP (IX)
    JP (IY)
