; LD r, r'
    LD A,A
    LD A,B
    LD A,C
    LD A,D
    LD A,E
    LD A,H
    LD A,L
    LD B,A
    LD B,B
    LD B,C
    LD B,D
    LD B,E
    LD B,H
    LD B,L
    LD C,A
    LD C,B
    LD C,C
    LD C,D
    LD C,E
    LD C,H
    LD C,L
    LD D,A
    LD D,B
    LD D,C
    LD D,D
    LD D,E
    LD D,H
    LD D,L
    LD E,A
    LD E,B
    LD E,C
    LD E,D
    LD E,E
    LD E,H
    LD E,L
    LD H,A
    LD H,B
    LD H,C
    LD H,D
    LD H,E
    LD H,H
    LD H,L
    LD L,A
    LD L,B
    LD L,C
    LD L,D
    LD L,E
    LD L,H
    LD L,L
    
; LD r, n
    LD A,$12
    LD B,$12
    LD C,$12
    LD D,$12
    LD E,$12
    LD H,$12
    LD L,$12
    
; LD r, (HL)
    LD A,(HL)
    LD B,(HL)
    LD C,(HL)
    LD D,(HL)
    LD E,(HL)
    LD H,(HL)
    LD L,(HL)
    
; LD r, (IX + d)
    LD A,(IX + 127)
    LD B,(IX + 127)
    LD C,(IX + 127)
    LD D,(IX + 127)
    LD E,(IX + 127)
    LD H,(IX + 127)
    LD L,(IX + 127)
    
; LD r, (IY + d)
    LD A,(IY - 128)
    LD B,(IY - 128)
    LD C,(IY - 128)
    LD D,(IY - 128)
    LD E,(IY - 128)
    LD H,(IY - 128)
    LD L,(IY - 128)
    
; LD (HL), r
    LD (HL),A
    LD (HL),B
    LD (HL),C
    LD (HL),D
    LD (HL),E
    LD (HL),H
    LD (HL),L
    
; LD (IX + d), r
    LD (IX + 127),A
    LD (IX + 127),B
    LD (IX + 127),C
    LD (IX + 127),D
    LD (IX + 127),E
    LD (IX + 127),H
    LD (IX + 127),L
    
; LD (IY + d), r
    LD (IY - 128),A
    LD (IY - 128),B
    LD (IY - 128),C
    LD (IY - 128),D
    LD (IY - 128),E
    LD (IY - 128),H
    LD (IY - 128),L
    
    LD (HL),$12
    LD (IX + 127),$12
    LD (IY - 128),$12
    
    LD A,(BC)
    LD A,(DE)
    LD A,($5678)
    
    LD (BC),A
    LD (DE),A
    LD ($5678),A
    
    LD A,I
    LD A,R
    LD I,A
    LD R,A
    
; LD dd, nn
    LD BC,$5678
    LD DE,$5678
    LD HL,$5678
    LD SP,$5678
    
    LD IX,$5678
    LD IY,$5678
    
    LD HL,($5678)
    
; LD dd, (nn)
    LD BC,($5678)
    LD DE,($5678)
    LD HL,($5678)
    LD SP,($5678)
    
    LD IX,($5678)
    LD IY,($5678)
    
    LD ($5678),HL
    
; LD (nn), dd
    LD ($5678),BC
    LD ($5678),DE
    LD ($5678),HL
    LD ($5678),SP
    
    LD ($5678),IX
    LD ($5678),IY
    
    LD SP,HL
    LD SP,IX
    LD SP,IY
    
; PUSH qq
    PUSH BC
    PUSH DE
    PUSH HL
    PUSH AF
    
    PUSH IX
    PUSH IY
    
; POP qq
    POP BC
    POP DE
    POP HL
    POP AF
    
    POP IX
    POP IY
    
    EX DE,HL
    EX AF,AF'
    EXX
    EX (SP),HL
    EX (SP),IX
    EX (SP),IY
    LDI
    LDIR
    LDD
    LDDR
    CPI
    CPIR
    CPD
    CPDR
    
; ADD A
    ADD A,A
    ADD A,B
    ADD A,C
    ADD A,D
    ADD A,E
    ADD A,H
    ADD A,L
    ADD A,$12
    ADD A,(HL)
    ADD A,(IX + 127)
    ADD A,(IY - 128)
    
; ADC A
    ADC A,A
    ADC A,B
    ADC A,C
    ADC A,D
    ADC A,E
    ADC A,H
    ADC A,L
    ADC A,$12
    ADC A,(HL)
    ADC A,(IX + 127)
    ADC A,(IY - 128)
    
; SUB
    SUB A
    SUB B
    SUB C
    SUB D
    SUB E
    SUB H
    SUB L
    SUB $12
    SUB (HL)
    SUB (IX + 127)
    SUB (IY - 128)
    
; SBC A
    SBC A,A
    SBC A,B
    SBC A,C
    SBC A,D
    SBC A,E
    SBC A,H
    SBC A,L
    SBC A,$12
    SBC A,(HL)
    SBC A,(IX + 127)
    SBC A,(IY - 128)
    
; AND
    AND A
    AND B
    AND C
    AND D
    AND E
    AND H
    AND L
    AND $12
    AND (HL)
    AND (IX + 127)
    AND (IY - 128)
    
; AND A - 本来は指定不要
    AND A
    AND B
    AND C
    AND D
    AND E
    AND H
    AND L
    AND $12
    AND (HL)
    AND (IX + 127)
    AND (IY - 128)
    
; OR
    OR A
    OR B
    OR C
    OR D
    OR E
    OR H
    OR L
    OR $12
    OR (HL)
    OR (IX + 127)
    OR (IY - 128)
    
; XOR
    XOR A
    XOR B
    XOR C
    XOR D
    XOR E
    XOR H
    XOR L
    XOR $12
    XOR (HL)
    XOR (IX + 127)
    XOR (IY - 128)
    
; CP
    CP A
    CP B
    CP C
    CP D
    CP E
    CP H
    CP L
    CP $12
    CP (HL)
    CP (IX + 127)
    CP (IY - 128)
    
; INC
    INC A
    INC B
    INC C
    INC D
    INC E
    INC H
    INC L
    
    INC (HL)
    INC (IX + 127)
    INC (IY - 128)
    
; DEC
    DEC A
    DEC B
    DEC C
    DEC D
    DEC E
    DEC H
    DEC L
    
    DEC (HL)
    DEC (IX + 127)
    DEC (IY - 128)
    
    DAA
    CPL
    NEG
    CCF
    SCF
    NOP
    HALT
    DI
    EI
    IM 0
    IM 1
    IM 2
    
; ADD HL, ss
    ADD HL,BC
    ADD HL,DE
    ADD HL,HL
    ADD HL,SP
    
; ADC HL, ss
    ADC HL,BC
    ADC HL,DE
    ADC HL,HL
    ADC HL,SP
    
; ABC HL, ss
    SBC HL,BC
    SBC HL,DE
    SBC HL,HL
    SBC HL,SP
    
; ABC IX, pp
    ADD IX,BC
    ADD IX,DE
    ADD IX,IX
    ADD IX,SP
    
; ABC IY, rr
    ADD IY,BC
    ADD IY,DE
    ADD IY,IY
    ADD IY,SP
    
; INC ss
    INC BC
    INC DE
    INC HL
    INC SP
    
    INC IX
    INC IY
    
; DEC ss
    DEC BC
    DEC DE
    DEC HL
    DEC SP
    
    DEC IX
    DEC IY
    
; 
    RLCA
    RLA
    RRCA
    RRA
    
; RLC
    RLC A
    RLC B
    RLC C
    RLC D
    RLC E
    RLC H
    RLC L
    RLC (HL)
    RLC (IX + 127)
    RLC (IY - 128)
    
; RL
    RL A
    RL B
    RL C
    RL D
    RL E
    RL H
    RL L
    RL (HL)
    RL (IX + 127)
    RL (IY - 128)
    
; RRC
    RRC A
    RRC B
    RRC C
    RRC D
    RRC E
    RRC H
    RRC L
    RRC (HL)
    RRC (IX + 127)
    RRC (IY - 128)
    
; RR
    RR A
    RR B
    RR C
    RR D
    RR E
    RR H
    RR L
    RR (HL)
    RR (IX + 127)
    RR (IY - 128)
    
; SLA
    SLA A
    SLA B
    SLA C
    SLA D
    SLA E
    SLA H
    SLA L
    SLA (HL)
    SLA (IX + 127)
    SLA (IY - 128)
    
; SRA
    SRA A
    SRA B
    SRA C
    SRA D
    SRA E
    SRA H
    SRA L
    SRA (HL)
    SRA (IX + 127)
    SRA (IY - 128)
    
; SRL
    SRL A
    SRL B
    SRL C
    SRL D
    SRL E
    SRL H
    SRL L
    SRL (HL)
    SRL (IX + 127)
    SRL (IY - 128)
    
    RLD
    RRD
    
; BIT
    BIT 0,A
    BIT 1,A
    BIT 2,A
    BIT 3,A
    BIT 4,A
    BIT 5,A
    BIT 6,A
    BIT 7,A
    BIT 0,B
    BIT 1,B
    BIT 2,B
    BIT 3,B
    BIT 4,B
    BIT 5,B
    BIT 6,B
    BIT 7,B
    BIT 0,C
    BIT 1,C
    BIT 2,C
    BIT 3,C
    BIT 4,C
    BIT 5,C
    BIT 6,C
    BIT 7,C
    BIT 0,D
    BIT 1,D
    BIT 2,D
    BIT 3,D
    BIT 4,D
    BIT 5,D
    BIT 6,D
    BIT 7,D
    BIT 0,E
    BIT 1,E
    BIT 2,E
    BIT 3,E
    BIT 4,E
    BIT 5,E
    BIT 6,E
    BIT 7,E
    BIT 0,H
    BIT 1,H
    BIT 2,H
    BIT 3,H
    BIT 4,H
    BIT 5,H
    BIT 6,H
    BIT 7,H
    BIT 0,L
    BIT 1,L
    BIT 2,L
    BIT 3,L
    BIT 4,L
    BIT 5,L
    BIT 6,L
    BIT 7,L
    BIT 0,(HL)
    BIT 1,(HL)
    BIT 2,(HL)
    BIT 3,(HL)
    BIT 4,(HL)
    BIT 5,(HL)
    BIT 6,(HL)
    BIT 7,(HL)
    BIT 0,(IX + 127)
    BIT 1,(IX + 127)
    BIT 2,(IX + 127)
    BIT 3,(IX + 127)
    BIT 4,(IX + 127)
    BIT 5,(IX + 127)
    BIT 6,(IX + 127)
    BIT 7,(IX + 127)
    BIT 0,(IY - 128)
    BIT 1,(IY - 128)
    BIT 2,(IY - 128)
    BIT 3,(IY - 128)
    BIT 4,(IY - 128)
    BIT 5,(IY - 128)
    BIT 6,(IY - 128)
    BIT 7,(IY - 128)
    
; SET
    SET 0,A
    SET 1,A
    SET 2,A
    SET 3,A
    SET 4,A
    SET 5,A
    SET 6,A
    SET 7,A
    SET 0,B
    SET 1,B
    SET 2,B
    SET 3,B
    SET 4,B
    SET 5,B
    SET 6,B
    SET 7,B
    SET 0,C
    SET 1,C
    SET 2,C
    SET 3,C
    SET 4,C
    SET 5,C
    SET 6,C
    SET 7,C
    SET 0,D
    SET 1,D
    SET 2,D
    SET 3,D
    SET 4,D
    SET 5,D
    SET 6,D
    SET 7,D
    SET 0,E
    SET 1,E
    SET 2,E
    SET 3,E
    SET 4,E
    SET 5,E
    SET 6,E
    SET 7,E
    SET 0,H
    SET 1,H
    SET 2,H
    SET 3,H
    SET 4,H
    SET 5,H
    SET 6,H
    SET 7,H
    SET 0,L
    SET 1,L
    SET 2,L
    SET 3,L
    SET 4,L
    SET 5,L
    SET 6,L
    SET 7,L
    SET 0,(HL)
    SET 1,(HL)
    SET 2,(HL)
    SET 3,(HL)
    SET 4,(HL)
    SET 5,(HL)
    SET 6,(HL)
    SET 7,(HL)
    SET 0,(IX + 127)
    SET 1,(IX + 127)
    SET 2,(IX + 127)
    SET 3,(IX + 127)
    SET 4,(IX + 127)
    SET 5,(IX + 127)
    SET 6,(IX + 127)
    SET 7,(IX + 127)
    SET 0,(IY - 128)
    SET 1,(IY - 128)
    SET 2,(IY - 128)
    SET 3,(IY - 128)
    SET 4,(IY - 128)
    SET 5,(IY - 128)
    SET 6,(IY - 128)
    SET 7,(IY - 128)
    
; RES
    RES 0,A
    RES 1,A
    RES 2,A
    RES 3,A
    RES 4,A
    RES 5,A
    RES 6,A
    RES 7,A
    RES 0,B
    RES 1,B
    RES 2,B
    RES 3,B
    RES 4,B
    RES 5,B
    RES 6,B
    RES 7,B
    RES 0,C
    RES 1,C
    RES 2,C
    RES 3,C
    RES 4,C
    RES 5,C
    RES 6,C
    RES 7,C
    RES 0,D
    RES 1,D
    RES 2,D
    RES 3,D
    RES 4,D
    RES 5,D
    RES 6,D
    RES 7,D
    RES 0,E
    RES 1,E
    RES 2,E
    RES 3,E
    RES 4,E
    RES 5,E
    RES 6,E
    RES 7,E
    RES 0,H
    RES 1,H
    RES 2,H
    RES 3,H
    RES 4,H
    RES 5,H
    RES 6,H
    RES 7,H
    RES 0,L
    RES 1,L
    RES 2,L
    RES 3,L
    RES 4,L
    RES 5,L
    RES 6,L
    RES 7,L
    RES 0,(HL)
    RES 1,(HL)
    RES 2,(HL)
    RES 3,(HL)
    RES 4,(HL)
    RES 5,(HL)
    RES 6,(HL)
    RES 7,(HL)
    RES 0,(IX + 127)
    RES 1,(IX + 127)
    RES 2,(IX + 127)
    RES 3,(IX + 127)
    RES 4,(IX + 127)
    RES 5,(IX + 127)
    RES 6,(IX + 127)
    RES 7,(IX + 127)
    RES 0,(IY - 128)
    RES 1,(IY - 128)
    RES 2,(IY - 128)
    RES 3,(IY - 128)
    RES 4,(IY - 128)
    RES 5,(IY - 128)
    RES 6,(IY - 128)
    RES 7,(IY - 128)
    
; JP nn
    JP $5678
    
; JP cc, nn
    JP NZ,$5678
    JP Z,$5678
    JP NC,$5678
    JP C,$5678
    JP PO,$5678
    JP PE,$5678
    JP P,$5678
    JP M,$5678
    
; JR e
    JR $ + 2
    
; JR cc, e
    JR NZ,$ + 2
    JR Z,$ + 2
    JR NC,$ + 2
    JR C,$ + 2
    
    JP (HL)
    JP (IX)
    JP (IY)
    
    DJNZ $ + 2
    
; CALL nn
    CALL $5678
    
; CALL cc, nn
    CALL NZ,$5678
    CALL Z,$5678
    CALL NC,$5678
    CALL C,$5678
    CALL PO,$5678
    CALL PE,$5678
    CALL P,$5678
    CALL M,$5678
    
; RET
    RET
    
; RET cc
    RET NZ
    RET Z
    RET NC
    RET C
    RET PO
    RET PE
    RET P
    RET M
    
    RETI
    RETN
    
; RST
    RST $00
    RST $08
    RST $10
    RST $18
    RST $20
    RST $28
    RST $30
    RST $38
    
; IN A, (n)
    IN A,($12)
    
; IN r, (C)
    IN A,(C)
    IN B,(C)
    IN C,(C)
    IN D,(C)
    IN E,(C)
    IN H,(C)
    IN L,(C)
    IN F,(C)
    
    INI
    INIR
    IND
    INDR
    
; OUT (n), A
    OUT ($12),A
    
; OUT (C), r
    OUT (C),A
    OUT (C),B
    OUT (C),C
    OUT (C),D
    OUT (C),E
    OUT (C),H
    OUT (C),L
    
    OUTI
    OTIR
    OUTD
    OTDR
