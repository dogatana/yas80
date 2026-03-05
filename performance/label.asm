label_0: LD A,A
label_1: LD A,B
label_2: LD A,C
label_3: LD A,D
label_4: LD A,E
label_5: LD A,H
label_6: LD A,L
label_7: LD B,A
label_8: LD B,B
label_9: LD B,C
label_10: LD B,D
label_11: LD B,E
label_12: LD B,H
label_13: LD B,L
label_14: LD C,A
label_15: LD C,B
label_16: LD C,C
label_17: LD C,D
label_18: LD C,E
label_19: LD C,H
label_20: LD C,L
label_21: LD D,A
label_22: LD D,B
label_23: LD D,C
label_24: LD D,D
label_25: LD D,E
label_26: LD D,H
label_27: LD D,L
label_28: LD E,A
label_29: LD E,B
label_30: LD E,C
label_31: LD E,D
label_32: LD E,E
label_33: LD E,H
label_34: LD E,L
label_35: LD H,A
label_36: LD H,B
label_37: LD H,C
label_38: LD H,D
label_39: LD H,E
label_40: LD H,H
label_41: LD H,L
label_42: LD L,A
label_43: LD L,B
label_44: LD L,C
label_45: LD L,D
label_46: LD L,E
label_47: LD L,H
label_48: LD L,L
label_49: LD A,$12
label_50: LD B,$12
label_51: LD C,$12
label_52: LD D,$12
label_53: LD E,$12
label_54: LD H,$12
label_55: LD L,$12
label_56: LD A,(HL)
label_57: LD B,(HL)
label_58: LD C,(HL)
label_59: LD D,(HL)
label_60: LD E,(HL)
label_61: LD H,(HL)
label_62: LD L,(HL)
label_63: LD A,(IX + 127)
label_64: LD B,(IX + 127)
label_65: LD C,(IX + 127)
label_66: LD D,(IX + 127)
label_67: LD E,(IX + 127)
label_68: LD H,(IX + 127)
label_69: LD L,(IX + 127)
label_70: LD A,(IY - 128)
label_71: LD B,(IY - 128)
label_72: LD C,(IY - 128)
label_73: LD D,(IY - 128)
label_74: LD E,(IY - 128)
label_75: LD H,(IY - 128)
label_76: LD L,(IY - 128)
label_77: LD (HL),A
label_78: LD (HL),B
label_79: LD (HL),C
label_80: LD (HL),D
label_81: LD (HL),E
label_82: LD (HL),H
label_83: LD (HL),L
label_84: LD (IX + 127),A
label_85: LD (IX + 127),B
label_86: LD (IX + 127),C
label_87: LD (IX + 127),D
label_88: LD (IX + 127),E
label_89: LD (IX + 127),H
label_90: LD (IX + 127),L
label_91: LD (IY - 128),A
label_92: LD (IY - 128),B
label_93: LD (IY - 128),C
label_94: LD (IY - 128),D
label_95: LD (IY - 128),E
label_96: LD (IY - 128),H
label_97: LD (IY - 128),L
label_98: LD (HL),$12
label_99: LD (IX + 127),$12
label_100: LD (IY - 128),$12
label_101: LD A,(BC)
label_102: LD A,(DE)
label_103: LD A,($5678)
label_104: LD (BC),A
label_105: LD (DE),A
label_106: LD ($5678),A
label_107: LD A,I
label_108: LD A,R
label_109: LD I,A
label_110: LD R,A
label_111: LD BC,$5678
label_112: LD DE,$5678
label_113: LD HL,$5678
label_114: LD SP,$5678
label_115: LD IX,$5678
label_116: LD IY,$5678
label_117: LD HL,($5678)
label_118: LD BC,($5678)
label_119: LD DE,($5678)
label_120: LD HL,($5678)
label_121: LD SP,($5678)
label_122: LD IX,($5678)
label_123: LD IY,($5678)
label_124: LD ($5678),HL
label_125: LD ($5678),BC
label_126: LD ($5678),DE
label_127: LD ($5678),HL
label_128: LD ($5678),SP
label_129: LD ($5678),IX
label_130: LD ($5678),IY
label_131: LD SP,HL
label_132: LD SP,IX
label_133: LD SP,IY
label_134: PUSH BC
label_135: PUSH DE
label_136: PUSH HL
label_137: PUSH AF
label_138: PUSH IX
label_139: PUSH IY
label_140: POP BC
label_141: POP DE
label_142: POP HL
label_143: POP AF
label_144: POP IX
label_145: POP IY
label_146: EX DE,HL
label_147: EX AF,AF'
label_148: EXX
label_149: EX (SP),HL
label_150: EX (SP),IX
label_151: EX (SP),IY
label_152: LDI
label_153: LDIR
label_154: LDD
label_155: LDDR
label_156: CPI
label_157: CPIR
label_158: CPD
label_159: CPDR
label_160: ADD A,A
label_161: ADD A,B
label_162: ADD A,C
label_163: ADD A,D
label_164: ADD A,E
label_165: ADD A,H
label_166: ADD A,L
label_167: ADD A,$12
label_168: ADD A,(HL)
label_169: ADD A,(IX + 127)
label_170: ADD A,(IY - 128)
label_171: ADC A,A
label_172: ADC A,B
label_173: ADC A,C
label_174: ADC A,D
label_175: ADC A,E
label_176: ADC A,H
label_177: ADC A,L
label_178: ADC A,$12
label_179: ADC A,(HL)
label_180: ADC A,(IX + 127)
label_181: ADC A,(IY - 128)
label_182: SUB A
label_183: SUB B
label_184: SUB C
label_185: SUB D
label_186: SUB E
label_187: SUB H
label_188: SUB L
label_189: SUB $12
label_190: SUB (HL)
label_191: SUB (IX + 127)
label_192: SUB (IY - 128)
label_193: SBC A,A
label_194: SBC A,B
label_195: SBC A,C
label_196: SBC A,D
label_197: SBC A,E
label_198: SBC A,H
label_199: SBC A,L
label_200: SBC A,$12
label_201: SBC A,(HL)
label_202: SBC A,(IX + 127)
label_203: SBC A,(IY - 128)
label_204: AND A
label_205: AND B
label_206: AND C
label_207: AND D
label_208: AND E
label_209: AND H
label_210: AND L
label_211: AND $12
label_212: AND (HL)
label_213: AND (IX + 127)
label_214: AND (IY - 128)
label_215: AND A
label_216: AND B
label_217: AND C
label_218: AND D
label_219: AND E
label_220: AND H
label_221: AND L
label_222: AND $12
label_223: AND (HL)
label_224: AND (IX + 127)
label_225: AND (IY - 128)
label_226: OR A
label_227: OR B
label_228: OR C
label_229: OR D
label_230: OR E
label_231: OR H
label_232: OR L
label_233: OR $12
label_234: OR (HL)
label_235: OR (IX + 127)
label_236: OR (IY - 128)
label_237: XOR A
label_238: XOR B
label_239: XOR C
label_240: XOR D
label_241: XOR E
label_242: XOR H
label_243: XOR L
label_244: XOR $12
label_245: XOR (HL)
label_246: XOR (IX + 127)
label_247: XOR (IY - 128)
label_248: CP A
label_249: CP B
label_250: CP C
label_251: CP D
label_252: CP E
label_253: CP H
label_254: CP L
label_255: CP $12
label_256: CP (HL)
label_257: CP (IX + 127)
label_258: CP (IY - 128)
label_259: INC A
label_260: INC B
label_261: INC C
label_262: INC D
label_263: INC E
label_264: INC H
label_265: INC L
label_266: INC (HL)
label_267: INC (IX + 127)
label_268: INC (IY - 128)
label_269: DEC A
label_270: DEC B
label_271: DEC C
label_272: DEC D
label_273: DEC E
label_274: DEC H
label_275: DEC L
label_276: DEC (HL)
label_277: DEC (IX + 127)
label_278: DEC (IY - 128)
label_279: DAA
label_280: CPL
label_281: NEG
label_282: CCF
label_283: SCF
label_284: NOP
label_285: HALT
label_286: DI
label_287: EI
label_288: IM 0
label_289: IM 1
label_290: IM 2
label_291: ADD HL,BC
label_292: ADD HL,DE
label_293: ADD HL,HL
label_294: ADD HL,SP
label_295: ADC HL,BC
label_296: ADC HL,DE
label_297: ADC HL,HL
label_298: ADC HL,SP
label_299: SBC HL,BC
label_300: SBC HL,DE
label_301: SBC HL,HL
label_302: SBC HL,SP
label_303: ADD IX,BC
label_304: ADD IX,DE
label_305: ADD IX,SP
label_306: ADD IY,BC
label_307: ADD IY,DE
label_308: ADD IY,SP
label_309: INC BC
label_310: INC DE
label_311: INC HL
label_312: INC SP
label_313: INC IX
label_314: INC IY
label_315: DEC BC
label_316: DEC DE
label_317: DEC HL
label_318: DEC SP
label_319: DEC IX
label_320: DEC IY
label_321: RLCA
label_322: RLA
label_323: RRCA
label_324: RRA
label_325: RLC A
label_326: RLC B
label_327: RLC C
label_328: RLC D
label_329: RLC E
label_330: RLC H
label_331: RLC L
label_332: RLC (HL)
label_333: RLC (IX + 127)
label_334: RLC (IY - 128)
label_335: RL A
label_336: RL B
label_337: RL C
label_338: RL D
label_339: RL E
label_340: RL H
label_341: RL L
label_342: RL (HL)
label_343: RL (IX + 127)
label_344: RL (IY - 128)
label_345: RRC A
label_346: RRC B
label_347: RRC C
label_348: RRC D
label_349: RRC E
label_350: RRC H
label_351: RRC L
label_352: RRC (HL)
label_353: RRC (IX + 127)
label_354: RRC (IY - 128)
label_355: RR A
label_356: RR B
label_357: RR C
label_358: RR D
label_359: RR E
label_360: RR H
label_361: RR L
label_362: RR (HL)
label_363: RR (IX + 127)
label_364: RR (IY - 128)
label_365: SLA A
label_366: SLA B
label_367: SLA C
label_368: SLA D
label_369: SLA E
label_370: SLA H
label_371: SLA L
label_372: SLA (HL)
label_373: SLA (IX + 127)
label_374: SLA (IY - 128)
label_375: SRA A
label_376: SRA B
label_377: SRA C
label_378: SRA D
label_379: SRA E
label_380: SRA H
label_381: SRA L
label_382: SRA (HL)
label_383: SRA (IX + 127)
label_384: SRA (IY - 128)
label_385: SRL A
label_386: SRL B
label_387: SRL C
label_388: SRL D
label_389: SRL E
label_390: SRL H
label_391: SRL L
label_392: SRL (HL)
label_393: SRL (IX + 127)
label_394: SRL (IY - 128)
label_395: RLD
label_396: RRD
label_397: BIT 0,A
label_398: BIT 1,A
label_399: BIT 2,A
label_400: BIT 3,A
label_401: BIT 4,A
label_402: BIT 5,A
label_403: BIT 6,A
label_404: BIT 7,A
label_405: BIT 0,B
label_406: BIT 1,B
label_407: BIT 2,B
label_408: BIT 3,B
label_409: BIT 4,B
label_410: BIT 5,B
label_411: BIT 6,B
label_412: BIT 7,B
label_413: BIT 0,C
label_414: BIT 1,C
label_415: BIT 2,C
label_416: BIT 3,C
label_417: BIT 4,C
label_418: BIT 5,C
label_419: BIT 6,C
label_420: BIT 7,C
label_421: BIT 0,D
label_422: BIT 1,D
label_423: BIT 2,D
label_424: BIT 3,D
label_425: BIT 4,D
label_426: BIT 5,D
label_427: BIT 6,D
label_428: BIT 7,D
label_429: BIT 0,E
label_430: BIT 1,E
label_431: BIT 2,E
label_432: BIT 3,E
label_433: BIT 4,E
label_434: BIT 5,E
label_435: BIT 6,E
label_436: BIT 7,E
label_437: BIT 0,H
label_438: BIT 1,H
label_439: BIT 2,H
label_440: BIT 3,H
label_441: BIT 4,H
label_442: BIT 5,H
label_443: BIT 6,H
label_444: BIT 7,H
label_445: BIT 0,L
label_446: BIT 1,L
label_447: BIT 2,L
label_448: BIT 3,L
label_449: BIT 4,L
label_450: BIT 5,L
label_451: BIT 6,L
label_452: BIT 7,L
label_453: BIT 0,(HL)
label_454: BIT 1,(HL)
label_455: BIT 2,(HL)
label_456: BIT 3,(HL)
label_457: BIT 4,(HL)
label_458: BIT 5,(HL)
label_459: BIT 6,(HL)
label_460: BIT 7,(HL)
label_461: BIT 0,(IX + 127)
label_462: BIT 1,(IX + 127)
label_463: BIT 2,(IX + 127)
label_464: BIT 3,(IX + 127)
label_465: BIT 4,(IX + 127)
label_466: BIT 5,(IX + 127)
label_467: BIT 6,(IX + 127)
label_468: BIT 7,(IX + 127)
label_469: BIT 0,(IY - 128)
label_470: BIT 1,(IY - 128)
label_471: BIT 2,(IY - 128)
label_472: BIT 3,(IY - 128)
label_473: BIT 4,(IY - 128)
label_474: BIT 5,(IY - 128)
label_475: BIT 6,(IY - 128)
label_476: BIT 7,(IY - 128)
label_477: SET 0,A
label_478: SET 1,A
label_479: SET 2,A
label_480: SET 3,A
label_481: SET 4,A
label_482: SET 5,A
label_483: SET 6,A
label_484: SET 7,A
label_485: SET 0,B
label_486: SET 1,B
label_487: SET 2,B
label_488: SET 3,B
label_489: SET 4,B
label_490: SET 5,B
label_491: SET 6,B
label_492: SET 7,B
label_493: SET 0,C
label_494: SET 1,C
label_495: SET 2,C
label_496: SET 3,C
label_497: SET 4,C
label_498: SET 5,C
label_499: SET 6,C
label_500: SET 7,C
label_501: SET 0,D
label_502: SET 1,D
label_503: SET 2,D
label_504: SET 3,D
label_505: SET 4,D
label_506: SET 5,D
label_507: SET 6,D
label_508: SET 7,D
label_509: SET 0,E
label_510: SET 1,E
label_511: SET 2,E
label_512: SET 3,E
label_513: SET 4,E
label_514: SET 5,E
label_515: SET 6,E
label_516: SET 7,E
label_517: SET 0,H
label_518: SET 1,H
label_519: SET 2,H
label_520: SET 3,H
label_521: SET 4,H
label_522: SET 5,H
label_523: SET 6,H
label_524: SET 7,H
label_525: SET 0,L
label_526: SET 1,L
label_527: SET 2,L
label_528: SET 3,L
label_529: SET 4,L
label_530: SET 5,L
label_531: SET 6,L
label_532: SET 7,L
label_533: SET 0,(HL)
label_534: SET 1,(HL)
label_535: SET 2,(HL)
label_536: SET 3,(HL)
label_537: SET 4,(HL)
label_538: SET 5,(HL)
label_539: SET 6,(HL)
label_540: SET 7,(HL)
label_541: SET 0,(IX + 127)
label_542: SET 1,(IX + 127)
label_543: SET 2,(IX + 127)
label_544: SET 3,(IX + 127)
label_545: SET 4,(IX + 127)
label_546: SET 5,(IX + 127)
label_547: SET 6,(IX + 127)
label_548: SET 7,(IX + 127)
label_549: SET 0,(IY - 128)
label_550: SET 1,(IY - 128)
label_551: SET 2,(IY - 128)
label_552: SET 3,(IY - 128)
label_553: SET 4,(IY - 128)
label_554: SET 5,(IY - 128)
label_555: SET 6,(IY - 128)
label_556: SET 7,(IY - 128)
label_557: RES 0,A
label_558: RES 1,A
label_559: RES 2,A
label_560: RES 3,A
label_561: RES 4,A
label_562: RES 5,A
label_563: RES 6,A
label_564: RES 7,A
label_565: RES 0,B
label_566: RES 1,B
label_567: RES 2,B
label_568: RES 3,B
label_569: RES 4,B
label_570: RES 5,B
label_571: RES 6,B
label_572: RES 7,B
label_573: RES 0,C
label_574: RES 1,C
label_575: RES 2,C
label_576: RES 3,C
label_577: RES 4,C
label_578: RES 5,C
label_579: RES 6,C
label_580: RES 7,C
label_581: RES 0,D
label_582: RES 1,D
label_583: RES 2,D
label_584: RES 3,D
label_585: RES 4,D
label_586: RES 5,D
label_587: RES 6,D
label_588: RES 7,D
label_589: RES 0,E
label_590: RES 1,E
label_591: RES 2,E
label_592: RES 3,E
label_593: RES 4,E
label_594: RES 5,E
label_595: RES 6,E
label_596: RES 7,E
label_597: RES 0,H
label_598: RES 1,H
label_599: RES 2,H
label_600: RES 3,H
label_601: RES 4,H
label_602: RES 5,H
label_603: RES 6,H
label_604: RES 7,H
label_605: RES 0,L
label_606: RES 1,L
label_607: RES 2,L
label_608: RES 3,L
label_609: RES 4,L
label_610: RES 5,L
label_611: RES 6,L
label_612: RES 7,L
label_613: RES 0,(HL)
label_614: RES 1,(HL)
label_615: RES 2,(HL)
label_616: RES 3,(HL)
label_617: RES 4,(HL)
label_618: RES 5,(HL)
label_619: RES 6,(HL)
label_620: RES 7,(HL)
label_621: RES 0,(IX + 127)
label_622: RES 1,(IX + 127)
label_623: RES 2,(IX + 127)
label_624: RES 3,(IX + 127)
label_625: RES 4,(IX + 127)
label_626: RES 5,(IX + 127)
label_627: RES 6,(IX + 127)
label_628: RES 7,(IX + 127)
label_629: RES 0,(IY - 128)
label_630: RES 1,(IY - 128)
label_631: RES 2,(IY - 128)
label_632: RES 3,(IY - 128)
label_633: RES 4,(IY - 128)
label_634: RES 5,(IY - 128)
label_635: RES 6,(IY - 128)
label_636: RES 7,(IY - 128)
label_637: JP $5678
label_638: JP NZ,$5678
label_639: JP Z,$5678
label_640: JP NC,$5678
label_641: JP C,$5678
label_642: JP PO,$5678
label_643: JP PE,$5678
label_644: JP P,$5678
label_645: JP M,$5678
label_646: JR $ + 2
label_647: JR NZ,$ + 2
label_648: JR Z,$ + 2
label_649: JR NC,$ + 2
label_650: JR C,$ + 2
label_651: JP (HL)
label_652: JP (IX)
label_653: JP (IY)
label_654: DJNZ $ + 2
label_655: CALL $5678
label_656: CALL NZ,$5678
label_657: CALL Z,$5678
label_658: CALL NC,$5678
label_659: CALL C,$5678
label_660: CALL PO,$5678
label_661: CALL PE,$5678
label_662: CALL P,$5678
label_663: CALL M,$5678
label_664: RET
label_665: RET NZ
label_666: RET Z
label_667: RET NC
label_668: RET C
label_669: RET PO
label_670: RET PE
label_671: RET P
label_672: RET M
label_673: RETI
label_674: RETN
label_675: RST $00
label_676: RST $08
label_677: RST $10
label_678: RST $18
label_679: RST $20
label_680: RST $28
label_681: RST $30
label_682: RST $38
label_683: IN A,($12)
label_684: IN A,(C)
label_685: IN B,(C)
label_686: IN C,(C)
label_687: IN D,(C)
label_688: IN E,(C)
label_689: IN H,(C)
label_690: IN L,(C)
label_691: IN F,(C)
label_692: INI
label_693: INIR
label_694: IND
label_695: INDR
label_696: OUT ($12),A
label_697: OUT (C),A
label_698: OUT (C),B
label_699: OUT (C),C
label_700: OUT (C),D
label_701: OUT (C),E
label_702: OUT (C),H
label_703: OUT (C),L
label_704: OUTI
label_705: OTIR
label_706: OUTD
label_707: OTDR
label_708: LD A,A
label_709: LD A,B
label_710: LD A,C
label_711: LD A,D
label_712: LD A,E
label_713: LD A,H
label_714: LD A,L
label_715: LD B,A
label_716: LD B,B
label_717: LD B,C
label_718: LD B,D
label_719: LD B,E
label_720: LD B,H
label_721: LD B,L
label_722: LD C,A
label_723: LD C,B
label_724: LD C,C
label_725: LD C,D
label_726: LD C,E
label_727: LD C,H
label_728: LD C,L
label_729: LD D,A
label_730: LD D,B
label_731: LD D,C
label_732: LD D,D
label_733: LD D,E
label_734: LD D,H
label_735: LD D,L
label_736: LD E,A
label_737: LD E,B
label_738: LD E,C
label_739: LD E,D
label_740: LD E,E
label_741: LD E,H
label_742: LD E,L
label_743: LD H,A
label_744: LD H,B
label_745: LD H,C
label_746: LD H,D
label_747: LD H,E
label_748: LD H,H
label_749: LD H,L
label_750: LD L,A
label_751: LD L,B
label_752: LD L,C
label_753: LD L,D
label_754: LD L,E
label_755: LD L,H
label_756: LD L,L
label_757: LD A,$12
label_758: LD B,$12
label_759: LD C,$12
label_760: LD D,$12
label_761: LD E,$12
label_762: LD H,$12
label_763: LD L,$12
label_764: LD A,(HL)
label_765: LD B,(HL)
label_766: LD C,(HL)
label_767: LD D,(HL)
label_768: LD E,(HL)
label_769: LD H,(HL)
label_770: LD L,(HL)
label_771: LD A,(IX + 127)
label_772: LD B,(IX + 127)
label_773: LD C,(IX + 127)
label_774: LD D,(IX + 127)
label_775: LD E,(IX + 127)
label_776: LD H,(IX + 127)
label_777: LD L,(IX + 127)
label_778: LD A,(IY - 128)
label_779: LD B,(IY - 128)
label_780: LD C,(IY - 128)
label_781: LD D,(IY - 128)
label_782: LD E,(IY - 128)
label_783: LD H,(IY - 128)
label_784: LD L,(IY - 128)
label_785: LD (HL),A
label_786: LD (HL),B
label_787: LD (HL),C
label_788: LD (HL),D
label_789: LD (HL),E
label_790: LD (HL),H
label_791: LD (HL),L
label_792: LD (IX + 127),A
label_793: LD (IX + 127),B
label_794: LD (IX + 127),C
label_795: LD (IX + 127),D
label_796: LD (IX + 127),E
label_797: LD (IX + 127),H
label_798: LD (IX + 127),L
label_799: LD (IY - 128),A
label_800: LD (IY - 128),B
label_801: LD (IY - 128),C
label_802: LD (IY - 128),D
label_803: LD (IY - 128),E
label_804: LD (IY - 128),H
label_805: LD (IY - 128),L
label_806: LD (HL),$12
label_807: LD (IX + 127),$12
label_808: LD (IY - 128),$12
label_809: LD A,(BC)
label_810: LD A,(DE)
label_811: LD A,($5678)
label_812: LD (BC),A
label_813: LD (DE),A
label_814: LD ($5678),A
label_815: LD A,I
label_816: LD A,R
label_817: LD I,A
label_818: LD R,A
label_819: LD BC,$5678
label_820: LD DE,$5678
label_821: LD HL,$5678
label_822: LD SP,$5678
label_823: LD IX,$5678
label_824: LD IY,$5678
label_825: LD HL,($5678)
label_826: LD BC,($5678)
label_827: LD DE,($5678)
label_828: LD HL,($5678)
label_829: LD SP,($5678)
label_830: LD IX,($5678)
label_831: LD IY,($5678)
label_832: LD ($5678),HL
label_833: LD ($5678),BC
label_834: LD ($5678),DE
label_835: LD ($5678),HL
label_836: LD ($5678),SP
label_837: LD ($5678),IX
label_838: LD ($5678),IY
label_839: LD SP,HL
label_840: LD SP,IX
label_841: LD SP,IY
label_842: PUSH BC
label_843: PUSH DE
label_844: PUSH HL
label_845: PUSH AF
label_846: PUSH IX
label_847: PUSH IY
label_848: POP BC
label_849: POP DE
label_850: POP HL
label_851: POP AF
label_852: POP IX
label_853: POP IY
label_854: EX DE,HL
label_855: EX AF,AF'
label_856: EXX
label_857: EX (SP),HL
label_858: EX (SP),IX
label_859: EX (SP),IY
label_860: LDI
label_861: LDIR
label_862: LDD
label_863: LDDR
label_864: CPI
label_865: CPIR
label_866: CPD
label_867: CPDR
label_868: ADD A,A
label_869: ADD A,B
label_870: ADD A,C
label_871: ADD A,D
label_872: ADD A,E
label_873: ADD A,H
label_874: ADD A,L
label_875: ADD A,$12
label_876: ADD A,(HL)
label_877: ADD A,(IX + 127)
label_878: ADD A,(IY - 128)
label_879: ADC A,A
label_880: ADC A,B
label_881: ADC A,C
label_882: ADC A,D
label_883: ADC A,E
label_884: ADC A,H
label_885: ADC A,L
label_886: ADC A,$12
label_887: ADC A,(HL)
label_888: ADC A,(IX + 127)
label_889: ADC A,(IY - 128)
label_890: SUB A
label_891: SUB B
label_892: SUB C
label_893: SUB D
label_894: SUB E
label_895: SUB H
label_896: SUB L
label_897: SUB $12
label_898: SUB (HL)
label_899: SUB (IX + 127)
label_900: SUB (IY - 128)
label_901: SBC A,A
label_902: SBC A,B
label_903: SBC A,C
label_904: SBC A,D
label_905: SBC A,E
label_906: SBC A,H
label_907: SBC A,L
label_908: SBC A,$12
label_909: SBC A,(HL)
label_910: SBC A,(IX + 127)
label_911: SBC A,(IY - 128)
label_912: AND A
label_913: AND B
label_914: AND C
label_915: AND D
label_916: AND E
label_917: AND H
label_918: AND L
label_919: AND $12
label_920: AND (HL)
label_921: AND (IX + 127)
label_922: AND (IY - 128)
label_923: AND A
label_924: AND B
label_925: AND C
label_926: AND D
label_927: AND E
label_928: AND H
label_929: AND L
label_930: AND $12
label_931: AND (HL)
label_932: AND (IX + 127)
label_933: AND (IY - 128)
label_934: OR A
label_935: OR B
label_936: OR C
label_937: OR D
label_938: OR E
label_939: OR H
label_940: OR L
label_941: OR $12
label_942: OR (HL)
label_943: OR (IX + 127)
label_944: OR (IY - 128)
label_945: XOR A
label_946: XOR B
label_947: XOR C
label_948: XOR D
label_949: XOR E
label_950: XOR H
label_951: XOR L
label_952: XOR $12
label_953: XOR (HL)
label_954: XOR (IX + 127)
label_955: XOR (IY - 128)
label_956: CP A
label_957: CP B
label_958: CP C
label_959: CP D
label_960: CP E
label_961: CP H
label_962: CP L
label_963: CP $12
label_964: CP (HL)
label_965: CP (IX + 127)
label_966: CP (IY - 128)
label_967: INC A
label_968: INC B
label_969: INC C
label_970: INC D
label_971: INC E
label_972: INC H
label_973: INC L
label_974: INC (HL)
label_975: INC (IX + 127)
label_976: INC (IY - 128)
label_977: DEC A
label_978: DEC B
label_979: DEC C
label_980: DEC D
label_981: DEC E
label_982: DEC H
label_983: DEC L
label_984: DEC (HL)
label_985: DEC (IX + 127)
label_986: DEC (IY - 128)
label_987: DAA
label_988: CPL
label_989: NEG
label_990: CCF
label_991: SCF
label_992: NOP
label_993: HALT
label_994: DI
label_995: EI
label_996: IM 0
label_997: IM 1
label_998: IM 2
label_999: ADD HL,BC
label_1000: ADD HL,DE
label_1001: ADD HL,HL
label_1002: ADD HL,SP
label_1003: ADC HL,BC
label_1004: ADC HL,DE
label_1005: ADC HL,HL
label_1006: ADC HL,SP
label_1007: SBC HL,BC
label_1008: SBC HL,DE
label_1009: SBC HL,HL
label_1010: SBC HL,SP
label_1011: ADD IX,BC
label_1012: ADD IX,DE
label_1013: ADD IX,SP
label_1014: ADD IY,BC
label_1015: ADD IY,DE
label_1016: ADD IY,SP
label_1017: INC BC
label_1018: INC DE
label_1019: INC HL
label_1020: INC SP
label_1021: INC IX
label_1022: INC IY
label_1023: DEC BC
label_1024: DEC DE
label_1025: DEC HL
label_1026: DEC SP
label_1027: DEC IX
label_1028: DEC IY
label_1029: RLCA
label_1030: RLA
label_1031: RRCA
label_1032: RRA
label_1033: RLC A
label_1034: RLC B
label_1035: RLC C
label_1036: RLC D
label_1037: RLC E
label_1038: RLC H
label_1039: RLC L
label_1040: RLC (HL)
label_1041: RLC (IX + 127)
label_1042: RLC (IY - 128)
label_1043: RL A
label_1044: RL B
label_1045: RL C
label_1046: RL D
label_1047: RL E
label_1048: RL H
label_1049: RL L
label_1050: RL (HL)
label_1051: RL (IX + 127)
label_1052: RL (IY - 128)
label_1053: RRC A
label_1054: RRC B
label_1055: RRC C
label_1056: RRC D
label_1057: RRC E
label_1058: RRC H
label_1059: RRC L
label_1060: RRC (HL)
label_1061: RRC (IX + 127)
label_1062: RRC (IY - 128)
label_1063: RR A
label_1064: RR B
label_1065: RR C
label_1066: RR D
label_1067: RR E
label_1068: RR H
label_1069: RR L
label_1070: RR (HL)
label_1071: RR (IX + 127)
label_1072: RR (IY - 128)
label_1073: SLA A
label_1074: SLA B
label_1075: SLA C
label_1076: SLA D
label_1077: SLA E
label_1078: SLA H
label_1079: SLA L
label_1080: SLA (HL)
label_1081: SLA (IX + 127)
label_1082: SLA (IY - 128)
label_1083: SRA A
label_1084: SRA B
label_1085: SRA C
label_1086: SRA D
label_1087: SRA E
label_1088: SRA H
label_1089: SRA L
label_1090: SRA (HL)
label_1091: SRA (IX + 127)
label_1092: SRA (IY - 128)
label_1093: SRL A
label_1094: SRL B
label_1095: SRL C
label_1096: SRL D
label_1097: SRL E
label_1098: SRL H
label_1099: SRL L
label_1100: SRL (HL)
label_1101: SRL (IX + 127)
label_1102: SRL (IY - 128)
label_1103: RLD
label_1104: RRD
label_1105: BIT 0,A
label_1106: BIT 1,A
label_1107: BIT 2,A
label_1108: BIT 3,A
label_1109: BIT 4,A
label_1110: BIT 5,A
label_1111: BIT 6,A
label_1112: BIT 7,A
label_1113: BIT 0,B
label_1114: BIT 1,B
label_1115: BIT 2,B
label_1116: BIT 3,B
label_1117: BIT 4,B
label_1118: BIT 5,B
label_1119: BIT 6,B
label_1120: BIT 7,B
label_1121: BIT 0,C
label_1122: BIT 1,C
label_1123: BIT 2,C
label_1124: BIT 3,C
label_1125: BIT 4,C
label_1126: BIT 5,C
label_1127: BIT 6,C
label_1128: BIT 7,C
label_1129: BIT 0,D
label_1130: BIT 1,D
label_1131: BIT 2,D
label_1132: BIT 3,D
label_1133: BIT 4,D
label_1134: BIT 5,D
label_1135: BIT 6,D
label_1136: BIT 7,D
label_1137: BIT 0,E
label_1138: BIT 1,E
label_1139: BIT 2,E
label_1140: BIT 3,E
label_1141: BIT 4,E
label_1142: BIT 5,E
label_1143: BIT 6,E
label_1144: BIT 7,E
label_1145: BIT 0,H
label_1146: BIT 1,H
label_1147: BIT 2,H
label_1148: BIT 3,H
label_1149: BIT 4,H
label_1150: BIT 5,H
label_1151: BIT 6,H
label_1152: BIT 7,H
label_1153: BIT 0,L
label_1154: BIT 1,L
label_1155: BIT 2,L
label_1156: BIT 3,L
label_1157: BIT 4,L
label_1158: BIT 5,L
label_1159: BIT 6,L
label_1160: BIT 7,L
label_1161: BIT 0,(HL)
label_1162: BIT 1,(HL)
label_1163: BIT 2,(HL)
label_1164: BIT 3,(HL)
label_1165: BIT 4,(HL)
label_1166: BIT 5,(HL)
label_1167: BIT 6,(HL)
label_1168: BIT 7,(HL)
label_1169: BIT 0,(IX + 127)
label_1170: BIT 1,(IX + 127)
label_1171: BIT 2,(IX + 127)
label_1172: BIT 3,(IX + 127)
label_1173: BIT 4,(IX + 127)
label_1174: BIT 5,(IX + 127)
label_1175: BIT 6,(IX + 127)
label_1176: BIT 7,(IX + 127)
label_1177: BIT 0,(IY - 128)
label_1178: BIT 1,(IY - 128)
label_1179: BIT 2,(IY - 128)
label_1180: BIT 3,(IY - 128)
label_1181: BIT 4,(IY - 128)
label_1182: BIT 5,(IY - 128)
label_1183: BIT 6,(IY - 128)
label_1184: BIT 7,(IY - 128)
label_1185: SET 0,A
label_1186: SET 1,A
label_1187: SET 2,A
label_1188: SET 3,A
label_1189: SET 4,A
label_1190: SET 5,A
label_1191: SET 6,A
label_1192: SET 7,A
label_1193: SET 0,B
label_1194: SET 1,B
label_1195: SET 2,B
label_1196: SET 3,B
label_1197: SET 4,B
label_1198: SET 5,B
label_1199: SET 6,B
label_1200: SET 7,B
label_1201: SET 0,C
label_1202: SET 1,C
label_1203: SET 2,C
label_1204: SET 3,C
label_1205: SET 4,C
label_1206: SET 5,C
label_1207: SET 6,C
label_1208: SET 7,C
label_1209: SET 0,D
label_1210: SET 1,D
label_1211: SET 2,D
label_1212: SET 3,D
label_1213: SET 4,D
label_1214: SET 5,D
label_1215: SET 6,D
label_1216: SET 7,D
label_1217: SET 0,E
label_1218: SET 1,E
label_1219: SET 2,E
label_1220: SET 3,E
label_1221: SET 4,E
label_1222: SET 5,E
label_1223: SET 6,E
label_1224: SET 7,E
label_1225: SET 0,H
label_1226: SET 1,H
label_1227: SET 2,H
label_1228: SET 3,H
label_1229: SET 4,H
label_1230: SET 5,H
label_1231: SET 6,H
label_1232: SET 7,H
label_1233: SET 0,L
label_1234: SET 1,L
label_1235: SET 2,L
label_1236: SET 3,L
label_1237: SET 4,L
label_1238: SET 5,L
label_1239: SET 6,L
label_1240: SET 7,L
label_1241: SET 0,(HL)
label_1242: SET 1,(HL)
label_1243: SET 2,(HL)
label_1244: SET 3,(HL)
label_1245: SET 4,(HL)
label_1246: SET 5,(HL)
label_1247: SET 6,(HL)
label_1248: SET 7,(HL)
label_1249: SET 0,(IX + 127)
label_1250: SET 1,(IX + 127)
label_1251: SET 2,(IX + 127)
label_1252: SET 3,(IX + 127)
label_1253: SET 4,(IX + 127)
label_1254: SET 5,(IX + 127)
label_1255: SET 6,(IX + 127)
label_1256: SET 7,(IX + 127)
label_1257: SET 0,(IY - 128)
label_1258: SET 1,(IY - 128)
label_1259: SET 2,(IY - 128)
label_1260: SET 3,(IY - 128)
label_1261: SET 4,(IY - 128)
label_1262: SET 5,(IY - 128)
label_1263: SET 6,(IY - 128)
label_1264: SET 7,(IY - 128)
label_1265: RES 0,A
label_1266: RES 1,A
label_1267: RES 2,A
label_1268: RES 3,A
label_1269: RES 4,A
label_1270: RES 5,A
label_1271: RES 6,A
label_1272: RES 7,A
label_1273: RES 0,B
label_1274: RES 1,B
label_1275: RES 2,B
label_1276: RES 3,B
label_1277: RES 4,B
label_1278: RES 5,B
label_1279: RES 6,B
label_1280: RES 7,B
label_1281: RES 0,C
label_1282: RES 1,C
label_1283: RES 2,C
label_1284: RES 3,C
label_1285: RES 4,C
label_1286: RES 5,C
label_1287: RES 6,C
label_1288: RES 7,C
label_1289: RES 0,D
label_1290: RES 1,D
label_1291: RES 2,D
label_1292: RES 3,D
label_1293: RES 4,D
label_1294: RES 5,D
label_1295: RES 6,D
label_1296: RES 7,D
label_1297: RES 0,E
label_1298: RES 1,E
label_1299: RES 2,E
label_1300: RES 3,E
label_1301: RES 4,E
label_1302: RES 5,E
label_1303: RES 6,E
label_1304: RES 7,E
label_1305: RES 0,H
label_1306: RES 1,H
label_1307: RES 2,H
label_1308: RES 3,H
label_1309: RES 4,H
label_1310: RES 5,H
label_1311: RES 6,H
label_1312: RES 7,H
label_1313: RES 0,L
label_1314: RES 1,L
label_1315: RES 2,L
label_1316: RES 3,L
label_1317: RES 4,L
label_1318: RES 5,L
label_1319: RES 6,L
label_1320: RES 7,L
label_1321: RES 0,(HL)
label_1322: RES 1,(HL)
label_1323: RES 2,(HL)
label_1324: RES 3,(HL)
label_1325: RES 4,(HL)
label_1326: RES 5,(HL)
label_1327: RES 6,(HL)
label_1328: RES 7,(HL)
label_1329: RES 0,(IX + 127)
label_1330: RES 1,(IX + 127)
label_1331: RES 2,(IX + 127)
label_1332: RES 3,(IX + 127)
label_1333: RES 4,(IX + 127)
label_1334: RES 5,(IX + 127)
label_1335: RES 6,(IX + 127)
label_1336: RES 7,(IX + 127)
label_1337: RES 0,(IY - 128)
label_1338: RES 1,(IY - 128)
label_1339: RES 2,(IY - 128)
label_1340: RES 3,(IY - 128)
label_1341: RES 4,(IY - 128)
label_1342: RES 5,(IY - 128)
label_1343: RES 6,(IY - 128)
label_1344: RES 7,(IY - 128)
label_1345: JP $5678
label_1346: JP NZ,$5678
label_1347: JP Z,$5678
label_1348: JP NC,$5678
label_1349: JP C,$5678
label_1350: JP PO,$5678
label_1351: JP PE,$5678
label_1352: JP P,$5678
label_1353: JP M,$5678
label_1354: JR $ + 2
label_1355: JR NZ,$ + 2
label_1356: JR Z,$ + 2
label_1357: JR NC,$ + 2
label_1358: JR C,$ + 2
label_1359: JP (HL)
label_1360: JP (IX)
label_1361: JP (IY)
label_1362: DJNZ $ + 2
label_1363: CALL $5678
label_1364: CALL NZ,$5678
label_1365: CALL Z,$5678
label_1366: CALL NC,$5678
label_1367: CALL C,$5678
label_1368: CALL PO,$5678
label_1369: CALL PE,$5678
label_1370: CALL P,$5678
label_1371: CALL M,$5678
label_1372: RET
label_1373: RET NZ
label_1374: RET Z
label_1375: RET NC
label_1376: RET C
label_1377: RET PO
label_1378: RET PE
label_1379: RET P
label_1380: RET M
label_1381: RETI
label_1382: RETN
label_1383: RST $00
label_1384: RST $08
label_1385: RST $10
label_1386: RST $18
label_1387: RST $20
label_1388: RST $28
label_1389: RST $30
label_1390: RST $38
label_1391: IN A,($12)
label_1392: IN A,(C)
label_1393: IN B,(C)
label_1394: IN C,(C)
label_1395: IN D,(C)
label_1396: IN E,(C)
label_1397: IN H,(C)
label_1398: IN L,(C)
label_1399: IN F,(C)
label_1400: INI
label_1401: INIR
label_1402: IND
label_1403: INDR
label_1404: OUT ($12),A
label_1405: OUT (C),A
label_1406: OUT (C),B
label_1407: OUT (C),C
label_1408: OUT (C),D
label_1409: OUT (C),E
label_1410: OUT (C),H
label_1411: OUT (C),L
label_1412: OUTI
label_1413: OTIR
label_1414: OUTD
label_1415: OTDR
label_1416: LD A,A
label_1417: LD A,B
label_1418: LD A,C
label_1419: LD A,D
label_1420: LD A,E
label_1421: LD A,H
label_1422: LD A,L
label_1423: LD B,A
label_1424: LD B,B
label_1425: LD B,C
label_1426: LD B,D
label_1427: LD B,E
label_1428: LD B,H
label_1429: LD B,L
label_1430: LD C,A
label_1431: LD C,B
label_1432: LD C,C
label_1433: LD C,D
label_1434: LD C,E
label_1435: LD C,H
label_1436: LD C,L
label_1437: LD D,A
label_1438: LD D,B
label_1439: LD D,C
label_1440: LD D,D
label_1441: LD D,E
label_1442: LD D,H
label_1443: LD D,L
label_1444: LD E,A
label_1445: LD E,B
label_1446: LD E,C
label_1447: LD E,D
label_1448: LD E,E
label_1449: LD E,H
label_1450: LD E,L
label_1451: LD H,A
label_1452: LD H,B
label_1453: LD H,C
label_1454: LD H,D
label_1455: LD H,E
label_1456: LD H,H
label_1457: LD H,L
label_1458: LD L,A
label_1459: LD L,B
label_1460: LD L,C
label_1461: LD L,D
label_1462: LD L,E
label_1463: LD L,H
label_1464: LD L,L
label_1465: LD A,$12
label_1466: LD B,$12
label_1467: LD C,$12
label_1468: LD D,$12
label_1469: LD E,$12
label_1470: LD H,$12
label_1471: LD L,$12
label_1472: LD A,(HL)
label_1473: LD B,(HL)
label_1474: LD C,(HL)
label_1475: LD D,(HL)
label_1476: LD E,(HL)
label_1477: LD H,(HL)
label_1478: LD L,(HL)
label_1479: LD A,(IX + 127)
label_1480: LD B,(IX + 127)
label_1481: LD C,(IX + 127)
label_1482: LD D,(IX + 127)
label_1483: LD E,(IX + 127)
label_1484: LD H,(IX + 127)
label_1485: LD L,(IX + 127)
label_1486: LD A,(IY - 128)
label_1487: LD B,(IY - 128)
label_1488: LD C,(IY - 128)
label_1489: LD D,(IY - 128)
label_1490: LD E,(IY - 128)
label_1491: LD H,(IY - 128)
label_1492: LD L,(IY - 128)
label_1493: LD (HL),A
label_1494: LD (HL),B
label_1495: LD (HL),C
label_1496: LD (HL),D
label_1497: LD (HL),E
label_1498: LD (HL),H
label_1499: LD (HL),L
label_1500: LD (IX + 127),A
label_1501: LD (IX + 127),B
label_1502: LD (IX + 127),C
label_1503: LD (IX + 127),D
label_1504: LD (IX + 127),E
label_1505: LD (IX + 127),H
label_1506: LD (IX + 127),L
label_1507: LD (IY - 128),A
label_1508: LD (IY - 128),B
label_1509: LD (IY - 128),C
label_1510: LD (IY - 128),D
label_1511: LD (IY - 128),E
label_1512: LD (IY - 128),H
label_1513: LD (IY - 128),L
label_1514: LD (HL),$12
label_1515: LD (IX + 127),$12
label_1516: LD (IY - 128),$12
label_1517: LD A,(BC)
label_1518: LD A,(DE)
label_1519: LD A,($5678)
label_1520: LD (BC),A
label_1521: LD (DE),A
label_1522: LD ($5678),A
label_1523: LD A,I
label_1524: LD A,R
label_1525: LD I,A
label_1526: LD R,A
label_1527: LD BC,$5678
label_1528: LD DE,$5678
label_1529: LD HL,$5678
label_1530: LD SP,$5678
label_1531: LD IX,$5678
label_1532: LD IY,$5678
label_1533: LD HL,($5678)
label_1534: LD BC,($5678)
label_1535: LD DE,($5678)
label_1536: LD HL,($5678)
label_1537: LD SP,($5678)
label_1538: LD IX,($5678)
label_1539: LD IY,($5678)
label_1540: LD ($5678),HL
label_1541: LD ($5678),BC
label_1542: LD ($5678),DE
label_1543: LD ($5678),HL
label_1544: LD ($5678),SP
label_1545: LD ($5678),IX
label_1546: LD ($5678),IY
label_1547: LD SP,HL
label_1548: LD SP,IX
label_1549: LD SP,IY
label_1550: PUSH BC
label_1551: PUSH DE
label_1552: PUSH HL
label_1553: PUSH AF
label_1554: PUSH IX
label_1555: PUSH IY
label_1556: POP BC
label_1557: POP DE
label_1558: POP HL
label_1559: POP AF
label_1560: POP IX
label_1561: POP IY
label_1562: EX DE,HL
label_1563: EX AF,AF'
label_1564: EXX
label_1565: EX (SP),HL
label_1566: EX (SP),IX
label_1567: EX (SP),IY
label_1568: LDI
label_1569: LDIR
label_1570: LDD
label_1571: LDDR
label_1572: CPI
label_1573: CPIR
label_1574: CPD
label_1575: CPDR
label_1576: ADD A,A
label_1577: ADD A,B
label_1578: ADD A,C
label_1579: ADD A,D
label_1580: ADD A,E
label_1581: ADD A,H
label_1582: ADD A,L
label_1583: ADD A,$12
label_1584: ADD A,(HL)
label_1585: ADD A,(IX + 127)
label_1586: ADD A,(IY - 128)
label_1587: ADC A,A
label_1588: ADC A,B
label_1589: ADC A,C
label_1590: ADC A,D
label_1591: ADC A,E
label_1592: ADC A,H
label_1593: ADC A,L
label_1594: ADC A,$12
label_1595: ADC A,(HL)
label_1596: ADC A,(IX + 127)
label_1597: ADC A,(IY - 128)
label_1598: SUB A
label_1599: SUB B
label_1600: SUB C
label_1601: SUB D
label_1602: SUB E
label_1603: SUB H
label_1604: SUB L
label_1605: SUB $12
label_1606: SUB (HL)
label_1607: SUB (IX + 127)
label_1608: SUB (IY - 128)
label_1609: SBC A,A
label_1610: SBC A,B
label_1611: SBC A,C
label_1612: SBC A,D
label_1613: SBC A,E
label_1614: SBC A,H
label_1615: SBC A,L
label_1616: SBC A,$12
label_1617: SBC A,(HL)
label_1618: SBC A,(IX + 127)
label_1619: SBC A,(IY - 128)
label_1620: AND A
label_1621: AND B
label_1622: AND C
label_1623: AND D
label_1624: AND E
label_1625: AND H
label_1626: AND L
label_1627: AND $12
label_1628: AND (HL)
label_1629: AND (IX + 127)
label_1630: AND (IY - 128)
label_1631: AND A
label_1632: AND B
label_1633: AND C
label_1634: AND D
label_1635: AND E
label_1636: AND H
label_1637: AND L
label_1638: AND $12
label_1639: AND (HL)
label_1640: AND (IX + 127)
label_1641: AND (IY - 128)
label_1642: OR A
label_1643: OR B
label_1644: OR C
label_1645: OR D
label_1646: OR E
label_1647: OR H
label_1648: OR L
label_1649: OR $12
label_1650: OR (HL)
label_1651: OR (IX + 127)
label_1652: OR (IY - 128)
label_1653: XOR A
label_1654: XOR B
label_1655: XOR C
label_1656: XOR D
label_1657: XOR E
label_1658: XOR H
label_1659: XOR L
label_1660: XOR $12
label_1661: XOR (HL)
label_1662: XOR (IX + 127)
label_1663: XOR (IY - 128)
label_1664: CP A
label_1665: CP B
label_1666: CP C
label_1667: CP D
label_1668: CP E
label_1669: CP H
label_1670: CP L
label_1671: CP $12
label_1672: CP (HL)
label_1673: CP (IX + 127)
label_1674: CP (IY - 128)
label_1675: INC A
label_1676: INC B
label_1677: INC C
label_1678: INC D
label_1679: INC E
label_1680: INC H
label_1681: INC L
label_1682: INC (HL)
label_1683: INC (IX + 127)
label_1684: INC (IY - 128)
label_1685: DEC A
label_1686: DEC B
label_1687: DEC C
label_1688: DEC D
label_1689: DEC E
label_1690: DEC H
label_1691: DEC L
label_1692: DEC (HL)
label_1693: DEC (IX + 127)
label_1694: DEC (IY - 128)
label_1695: DAA
label_1696: CPL
label_1697: NEG
label_1698: CCF
label_1699: SCF
label_1700: NOP
label_1701: HALT
label_1702: DI
label_1703: EI
label_1704: IM 0
label_1705: IM 1
label_1706: IM 2
label_1707: ADD HL,BC
label_1708: ADD HL,DE
label_1709: ADD HL,HL
label_1710: ADD HL,SP
label_1711: ADC HL,BC
label_1712: ADC HL,DE
label_1713: ADC HL,HL
label_1714: ADC HL,SP
label_1715: SBC HL,BC
label_1716: SBC HL,DE
label_1717: SBC HL,HL
label_1718: SBC HL,SP
label_1719: ADD IX,BC
label_1720: ADD IX,DE
label_1721: ADD IX,SP
label_1722: ADD IY,BC
label_1723: ADD IY,DE
label_1724: ADD IY,SP
label_1725: INC BC
label_1726: INC DE
label_1727: INC HL
label_1728: INC SP
label_1729: INC IX
label_1730: INC IY
label_1731: DEC BC
label_1732: DEC DE
label_1733: DEC HL
label_1734: DEC SP
label_1735: DEC IX
label_1736: DEC IY
label_1737: RLCA
label_1738: RLA
label_1739: RRCA
label_1740: RRA
label_1741: RLC A
label_1742: RLC B
label_1743: RLC C
label_1744: RLC D
label_1745: RLC E
label_1746: RLC H
label_1747: RLC L
label_1748: RLC (HL)
label_1749: RLC (IX + 127)
label_1750: RLC (IY - 128)
label_1751: RL A
label_1752: RL B
label_1753: RL C
label_1754: RL D
label_1755: RL E
label_1756: RL H
label_1757: RL L
label_1758: RL (HL)
label_1759: RL (IX + 127)
label_1760: RL (IY - 128)
label_1761: RRC A
label_1762: RRC B
label_1763: RRC C
label_1764: RRC D
label_1765: RRC E
label_1766: RRC H
label_1767: RRC L
label_1768: RRC (HL)
label_1769: RRC (IX + 127)
label_1770: RRC (IY - 128)
label_1771: RR A
label_1772: RR B
label_1773: RR C
label_1774: RR D
label_1775: RR E
label_1776: RR H
label_1777: RR L
label_1778: RR (HL)
label_1779: RR (IX + 127)
label_1780: RR (IY - 128)
label_1781: SLA A
label_1782: SLA B
label_1783: SLA C
label_1784: SLA D
label_1785: SLA E
label_1786: SLA H
label_1787: SLA L
label_1788: SLA (HL)
label_1789: SLA (IX + 127)
label_1790: SLA (IY - 128)
label_1791: SRA A
label_1792: SRA B
label_1793: SRA C
label_1794: SRA D
label_1795: SRA E
label_1796: SRA H
label_1797: SRA L
label_1798: SRA (HL)
label_1799: SRA (IX + 127)
label_1800: SRA (IY - 128)
label_1801: SRL A
label_1802: SRL B
label_1803: SRL C
label_1804: SRL D
label_1805: SRL E
label_1806: SRL H
label_1807: SRL L
label_1808: SRL (HL)
label_1809: SRL (IX + 127)
label_1810: SRL (IY - 128)
label_1811: RLD
label_1812: RRD
label_1813: BIT 0,A
label_1814: BIT 1,A
label_1815: BIT 2,A
label_1816: BIT 3,A
label_1817: BIT 4,A
label_1818: BIT 5,A
label_1819: BIT 6,A
label_1820: BIT 7,A
label_1821: BIT 0,B
label_1822: BIT 1,B
label_1823: BIT 2,B
label_1824: BIT 3,B
label_1825: BIT 4,B
label_1826: BIT 5,B
label_1827: BIT 6,B
label_1828: BIT 7,B
label_1829: BIT 0,C
label_1830: BIT 1,C
label_1831: BIT 2,C
label_1832: BIT 3,C
label_1833: BIT 4,C
label_1834: BIT 5,C
label_1835: BIT 6,C
label_1836: BIT 7,C
label_1837: BIT 0,D
label_1838: BIT 1,D
label_1839: BIT 2,D
label_1840: BIT 3,D
label_1841: BIT 4,D
label_1842: BIT 5,D
label_1843: BIT 6,D
label_1844: BIT 7,D
label_1845: BIT 0,E
label_1846: BIT 1,E
label_1847: BIT 2,E
label_1848: BIT 3,E
label_1849: BIT 4,E
label_1850: BIT 5,E
label_1851: BIT 6,E
label_1852: BIT 7,E
label_1853: BIT 0,H
label_1854: BIT 1,H
label_1855: BIT 2,H
label_1856: BIT 3,H
label_1857: BIT 4,H
label_1858: BIT 5,H
label_1859: BIT 6,H
label_1860: BIT 7,H
label_1861: BIT 0,L
label_1862: BIT 1,L
label_1863: BIT 2,L
label_1864: BIT 3,L
label_1865: BIT 4,L
label_1866: BIT 5,L
label_1867: BIT 6,L
label_1868: BIT 7,L
label_1869: BIT 0,(HL)
label_1870: BIT 1,(HL)
label_1871: BIT 2,(HL)
label_1872: BIT 3,(HL)
label_1873: BIT 4,(HL)
label_1874: BIT 5,(HL)
label_1875: BIT 6,(HL)
label_1876: BIT 7,(HL)
label_1877: BIT 0,(IX + 127)
label_1878: BIT 1,(IX + 127)
label_1879: BIT 2,(IX + 127)
label_1880: BIT 3,(IX + 127)
label_1881: BIT 4,(IX + 127)
label_1882: BIT 5,(IX + 127)
label_1883: BIT 6,(IX + 127)
label_1884: BIT 7,(IX + 127)
label_1885: BIT 0,(IY - 128)
label_1886: BIT 1,(IY - 128)
label_1887: BIT 2,(IY - 128)
label_1888: BIT 3,(IY - 128)
label_1889: BIT 4,(IY - 128)
label_1890: BIT 5,(IY - 128)
label_1891: BIT 6,(IY - 128)
label_1892: BIT 7,(IY - 128)
label_1893: SET 0,A
label_1894: SET 1,A
label_1895: SET 2,A
label_1896: SET 3,A
label_1897: SET 4,A
label_1898: SET 5,A
label_1899: SET 6,A
label_1900: SET 7,A
label_1901: SET 0,B
label_1902: SET 1,B
label_1903: SET 2,B
label_1904: SET 3,B
label_1905: SET 4,B
label_1906: SET 5,B
label_1907: SET 6,B
label_1908: SET 7,B
label_1909: SET 0,C
label_1910: SET 1,C
label_1911: SET 2,C
label_1912: SET 3,C
label_1913: SET 4,C
label_1914: SET 5,C
label_1915: SET 6,C
label_1916: SET 7,C
label_1917: SET 0,D
label_1918: SET 1,D
label_1919: SET 2,D
label_1920: SET 3,D
label_1921: SET 4,D
label_1922: SET 5,D
label_1923: SET 6,D
label_1924: SET 7,D
label_1925: SET 0,E
label_1926: SET 1,E
label_1927: SET 2,E
label_1928: SET 3,E
label_1929: SET 4,E
label_1930: SET 5,E
label_1931: SET 6,E
label_1932: SET 7,E
label_1933: SET 0,H
label_1934: SET 1,H
label_1935: SET 2,H
label_1936: SET 3,H
label_1937: SET 4,H
label_1938: SET 5,H
label_1939: SET 6,H
label_1940: SET 7,H
label_1941: SET 0,L
label_1942: SET 1,L
label_1943: SET 2,L
label_1944: SET 3,L
label_1945: SET 4,L
label_1946: SET 5,L
label_1947: SET 6,L
label_1948: SET 7,L
label_1949: SET 0,(HL)
label_1950: SET 1,(HL)
label_1951: SET 2,(HL)
label_1952: SET 3,(HL)
label_1953: SET 4,(HL)
label_1954: SET 5,(HL)
label_1955: SET 6,(HL)
label_1956: SET 7,(HL)
label_1957: SET 0,(IX + 127)
label_1958: SET 1,(IX + 127)
label_1959: SET 2,(IX + 127)
label_1960: SET 3,(IX + 127)
label_1961: SET 4,(IX + 127)
label_1962: SET 5,(IX + 127)
label_1963: SET 6,(IX + 127)
label_1964: SET 7,(IX + 127)
label_1965: SET 0,(IY - 128)
label_1966: SET 1,(IY - 128)
label_1967: SET 2,(IY - 128)
label_1968: SET 3,(IY - 128)
label_1969: SET 4,(IY - 128)
label_1970: SET 5,(IY - 128)
label_1971: SET 6,(IY - 128)
label_1972: SET 7,(IY - 128)
label_1973: RES 0,A
label_1974: RES 1,A
label_1975: RES 2,A
label_1976: RES 3,A
label_1977: RES 4,A
label_1978: RES 5,A
label_1979: RES 6,A
label_1980: RES 7,A
label_1981: RES 0,B
label_1982: RES 1,B
label_1983: RES 2,B
label_1984: RES 3,B
label_1985: RES 4,B
label_1986: RES 5,B
label_1987: RES 6,B
label_1988: RES 7,B
label_1989: RES 0,C
label_1990: RES 1,C
label_1991: RES 2,C
label_1992: RES 3,C
label_1993: RES 4,C
label_1994: RES 5,C
label_1995: RES 6,C
label_1996: RES 7,C
label_1997: RES 0,D
label_1998: RES 1,D
label_1999: RES 2,D
label_2000: RES 3,D
label_2001: RES 4,D
label_2002: RES 5,D
label_2003: RES 6,D
label_2004: RES 7,D
label_2005: RES 0,E
label_2006: RES 1,E
label_2007: RES 2,E
label_2008: RES 3,E
label_2009: RES 4,E
label_2010: RES 5,E
label_2011: RES 6,E
label_2012: RES 7,E
label_2013: RES 0,H
label_2014: RES 1,H
label_2015: RES 2,H
label_2016: RES 3,H
label_2017: RES 4,H
label_2018: RES 5,H
label_2019: RES 6,H
label_2020: RES 7,H
label_2021: RES 0,L
label_2022: RES 1,L
label_2023: RES 2,L
label_2024: RES 3,L
label_2025: RES 4,L
label_2026: RES 5,L
label_2027: RES 6,L
label_2028: RES 7,L
label_2029: RES 0,(HL)
label_2030: RES 1,(HL)
label_2031: RES 2,(HL)
label_2032: RES 3,(HL)
label_2033: RES 4,(HL)
label_2034: RES 5,(HL)
label_2035: RES 6,(HL)
label_2036: RES 7,(HL)
label_2037: RES 0,(IX + 127)
label_2038: RES 1,(IX + 127)
label_2039: RES 2,(IX + 127)
label_2040: RES 3,(IX + 127)
label_2041: RES 4,(IX + 127)
label_2042: RES 5,(IX + 127)
label_2043: RES 6,(IX + 127)
label_2044: RES 7,(IX + 127)
label_2045: RES 0,(IY - 128)
label_2046: RES 1,(IY - 128)
label_2047: RES 2,(IY - 128)
label_2048: RES 3,(IY - 128)
label_2049: RES 4,(IY - 128)
label_2050: RES 5,(IY - 128)
label_2051: RES 6,(IY - 128)
label_2052: RES 7,(IY - 128)
label_2053: JP $5678
label_2054: JP NZ,$5678
label_2055: JP Z,$5678
label_2056: JP NC,$5678
label_2057: JP C,$5678
label_2058: JP PO,$5678
label_2059: JP PE,$5678
label_2060: JP P,$5678
label_2061: JP M,$5678
label_2062: JR $ + 2
label_2063: JR NZ,$ + 2
label_2064: JR Z,$ + 2
label_2065: JR NC,$ + 2
label_2066: JR C,$ + 2
label_2067: JP (HL)
label_2068: JP (IX)
label_2069: JP (IY)
label_2070: DJNZ $ + 2
label_2071: CALL $5678
label_2072: CALL NZ,$5678
label_2073: CALL Z,$5678
label_2074: CALL NC,$5678
label_2075: CALL C,$5678
label_2076: CALL PO,$5678
label_2077: CALL PE,$5678
label_2078: CALL P,$5678
label_2079: CALL M,$5678
label_2080: RET
label_2081: RET NZ
label_2082: RET Z
label_2083: RET NC
label_2084: RET C
label_2085: RET PO
label_2086: RET PE
label_2087: RET P
label_2088: RET M
label_2089: RETI
label_2090: RETN
label_2091: RST $00
label_2092: RST $08
label_2093: RST $10
label_2094: RST $18
label_2095: RST $20
label_2096: RST $28
label_2097: RST $30
label_2098: RST $38
label_2099: IN A,($12)
label_2100: IN A,(C)
label_2101: IN B,(C)
label_2102: IN C,(C)
label_2103: IN D,(C)
label_2104: IN E,(C)
label_2105: IN H,(C)
label_2106: IN L,(C)
label_2107: IN F,(C)
label_2108: INI
label_2109: INIR
label_2110: IND
label_2111: INDR
label_2112: OUT ($12),A
label_2113: OUT (C),A
label_2114: OUT (C),B
label_2115: OUT (C),C
label_2116: OUT (C),D
label_2117: OUT (C),E
label_2118: OUT (C),H
label_2119: OUT (C),L
label_2120: OUTI
label_2121: OTIR
label_2122: OUTD
label_2123: OTDR
label_2124: LD A,A
label_2125: LD A,B
label_2126: LD A,C
label_2127: LD A,D
label_2128: LD A,E
label_2129: LD A,H
label_2130: LD A,L
label_2131: LD B,A
label_2132: LD B,B
label_2133: LD B,C
label_2134: LD B,D
label_2135: LD B,E
label_2136: LD B,H
label_2137: LD B,L
label_2138: LD C,A
label_2139: LD C,B
label_2140: LD C,C
label_2141: LD C,D
label_2142: LD C,E
label_2143: LD C,H
label_2144: LD C,L
label_2145: LD D,A
label_2146: LD D,B
label_2147: LD D,C
label_2148: LD D,D
label_2149: LD D,E
label_2150: LD D,H
label_2151: LD D,L
label_2152: LD E,A
label_2153: LD E,B
label_2154: LD E,C
label_2155: LD E,D
label_2156: LD E,E
label_2157: LD E,H
label_2158: LD E,L
label_2159: LD H,A
label_2160: LD H,B
label_2161: LD H,C
label_2162: LD H,D
label_2163: LD H,E
label_2164: LD H,H
label_2165: LD H,L
label_2166: LD L,A
label_2167: LD L,B
label_2168: LD L,C
label_2169: LD L,D
label_2170: LD L,E
label_2171: LD L,H
label_2172: LD L,L
label_2173: LD A,$12
label_2174: LD B,$12
label_2175: LD C,$12
label_2176: LD D,$12
label_2177: LD E,$12
label_2178: LD H,$12
label_2179: LD L,$12
label_2180: LD A,(HL)
label_2181: LD B,(HL)
label_2182: LD C,(HL)
label_2183: LD D,(HL)
label_2184: LD E,(HL)
label_2185: LD H,(HL)
label_2186: LD L,(HL)
label_2187: LD A,(IX + 127)
label_2188: LD B,(IX + 127)
label_2189: LD C,(IX + 127)
label_2190: LD D,(IX + 127)
label_2191: LD E,(IX + 127)
label_2192: LD H,(IX + 127)
label_2193: LD L,(IX + 127)
label_2194: LD A,(IY - 128)
label_2195: LD B,(IY - 128)
label_2196: LD C,(IY - 128)
label_2197: LD D,(IY - 128)
label_2198: LD E,(IY - 128)
label_2199: LD H,(IY - 128)
label_2200: LD L,(IY - 128)
label_2201: LD (HL),A
label_2202: LD (HL),B
label_2203: LD (HL),C
label_2204: LD (HL),D
label_2205: LD (HL),E
label_2206: LD (HL),H
label_2207: LD (HL),L
label_2208: LD (IX + 127),A
label_2209: LD (IX + 127),B
label_2210: LD (IX + 127),C
label_2211: LD (IX + 127),D
label_2212: LD (IX + 127),E
label_2213: LD (IX + 127),H
label_2214: LD (IX + 127),L
label_2215: LD (IY - 128),A
label_2216: LD (IY - 128),B
label_2217: LD (IY - 128),C
label_2218: LD (IY - 128),D
label_2219: LD (IY - 128),E
label_2220: LD (IY - 128),H
label_2221: LD (IY - 128),L
label_2222: LD (HL),$12
label_2223: LD (IX + 127),$12
label_2224: LD (IY - 128),$12
label_2225: LD A,(BC)
label_2226: LD A,(DE)
label_2227: LD A,($5678)
label_2228: LD (BC),A
label_2229: LD (DE),A
label_2230: LD ($5678),A
label_2231: LD A,I
label_2232: LD A,R
label_2233: LD I,A
label_2234: LD R,A
label_2235: LD BC,$5678
label_2236: LD DE,$5678
label_2237: LD HL,$5678
label_2238: LD SP,$5678
label_2239: LD IX,$5678
label_2240: LD IY,$5678
label_2241: LD HL,($5678)
label_2242: LD BC,($5678)
label_2243: LD DE,($5678)
label_2244: LD HL,($5678)
label_2245: LD SP,($5678)
label_2246: LD IX,($5678)
label_2247: LD IY,($5678)
label_2248: LD ($5678),HL
label_2249: LD ($5678),BC
label_2250: LD ($5678),DE
label_2251: LD ($5678),HL
label_2252: LD ($5678),SP
label_2253: LD ($5678),IX
label_2254: LD ($5678),IY
label_2255: LD SP,HL
label_2256: LD SP,IX
label_2257: LD SP,IY
label_2258: PUSH BC
label_2259: PUSH DE
label_2260: PUSH HL
label_2261: PUSH AF
label_2262: PUSH IX
label_2263: PUSH IY
label_2264: POP BC
label_2265: POP DE
label_2266: POP HL
label_2267: POP AF
label_2268: POP IX
label_2269: POP IY
label_2270: EX DE,HL
label_2271: EX AF,AF'
label_2272: EXX
label_2273: EX (SP),HL
label_2274: EX (SP),IX
label_2275: EX (SP),IY
label_2276: LDI
label_2277: LDIR
label_2278: LDD
label_2279: LDDR
label_2280: CPI
label_2281: CPIR
label_2282: CPD
label_2283: CPDR
label_2284: ADD A,A
label_2285: ADD A,B
label_2286: ADD A,C
label_2287: ADD A,D
label_2288: ADD A,E
label_2289: ADD A,H
label_2290: ADD A,L
label_2291: ADD A,$12
label_2292: ADD A,(HL)
label_2293: ADD A,(IX + 127)
label_2294: ADD A,(IY - 128)
label_2295: ADC A,A
label_2296: ADC A,B
label_2297: ADC A,C
label_2298: ADC A,D
label_2299: ADC A,E
label_2300: ADC A,H
label_2301: ADC A,L
label_2302: ADC A,$12
label_2303: ADC A,(HL)
label_2304: ADC A,(IX + 127)
label_2305: ADC A,(IY - 128)
label_2306: SUB A
label_2307: SUB B
label_2308: SUB C
label_2309: SUB D
label_2310: SUB E
label_2311: SUB H
label_2312: SUB L
label_2313: SUB $12
label_2314: SUB (HL)
label_2315: SUB (IX + 127)
label_2316: SUB (IY - 128)
label_2317: SBC A,A
label_2318: SBC A,B
label_2319: SBC A,C
label_2320: SBC A,D
label_2321: SBC A,E
label_2322: SBC A,H
label_2323: SBC A,L
label_2324: SBC A,$12
label_2325: SBC A,(HL)
label_2326: SBC A,(IX + 127)
label_2327: SBC A,(IY - 128)
label_2328: AND A
label_2329: AND B
label_2330: AND C
label_2331: AND D
label_2332: AND E
label_2333: AND H
label_2334: AND L
label_2335: AND $12
label_2336: AND (HL)
label_2337: AND (IX + 127)
label_2338: AND (IY - 128)
label_2339: AND A
label_2340: AND B
label_2341: AND C
label_2342: AND D
label_2343: AND E
label_2344: AND H
label_2345: AND L
label_2346: AND $12
label_2347: AND (HL)
label_2348: AND (IX + 127)
label_2349: AND (IY - 128)
label_2350: OR A
label_2351: OR B
label_2352: OR C
label_2353: OR D
label_2354: OR E
label_2355: OR H
label_2356: OR L
label_2357: OR $12
label_2358: OR (HL)
label_2359: OR (IX + 127)
label_2360: OR (IY - 128)
label_2361: XOR A
label_2362: XOR B
label_2363: XOR C
label_2364: XOR D
label_2365: XOR E
label_2366: XOR H
label_2367: XOR L
label_2368: XOR $12
label_2369: XOR (HL)
label_2370: XOR (IX + 127)
label_2371: XOR (IY - 128)
label_2372: CP A
label_2373: CP B
label_2374: CP C
label_2375: CP D
label_2376: CP E
label_2377: CP H
label_2378: CP L
label_2379: CP $12
label_2380: CP (HL)
label_2381: CP (IX + 127)
label_2382: CP (IY - 128)
label_2383: INC A
label_2384: INC B
label_2385: INC C
label_2386: INC D
label_2387: INC E
label_2388: INC H
label_2389: INC L
label_2390: INC (HL)
label_2391: INC (IX + 127)
label_2392: INC (IY - 128)
label_2393: DEC A
label_2394: DEC B
label_2395: DEC C
label_2396: DEC D
label_2397: DEC E
label_2398: DEC H
label_2399: DEC L
label_2400: DEC (HL)
label_2401: DEC (IX + 127)
label_2402: DEC (IY - 128)
label_2403: DAA
label_2404: CPL
label_2405: NEG
label_2406: CCF
label_2407: SCF
label_2408: NOP
label_2409: HALT
label_2410: DI
label_2411: EI
label_2412: IM 0
label_2413: IM 1
label_2414: IM 2
label_2415: ADD HL,BC
label_2416: ADD HL,DE
label_2417: ADD HL,HL
label_2418: ADD HL,SP
label_2419: ADC HL,BC
label_2420: ADC HL,DE
label_2421: ADC HL,HL
label_2422: ADC HL,SP
label_2423: SBC HL,BC
label_2424: SBC HL,DE
label_2425: SBC HL,HL
label_2426: SBC HL,SP
label_2427: ADD IX,BC
label_2428: ADD IX,DE
label_2429: ADD IX,SP
label_2430: ADD IY,BC
label_2431: ADD IY,DE
label_2432: ADD IY,SP
label_2433: INC BC
label_2434: INC DE
label_2435: INC HL
label_2436: INC SP
label_2437: INC IX
label_2438: INC IY
label_2439: DEC BC
label_2440: DEC DE
label_2441: DEC HL
label_2442: DEC SP
label_2443: DEC IX
label_2444: DEC IY
label_2445: RLCA
label_2446: RLA
label_2447: RRCA
label_2448: RRA
label_2449: RLC A
label_2450: RLC B
label_2451: RLC C
label_2452: RLC D
label_2453: RLC E
label_2454: RLC H
label_2455: RLC L
label_2456: RLC (HL)
label_2457: RLC (IX + 127)
label_2458: RLC (IY - 128)
label_2459: RL A
label_2460: RL B
label_2461: RL C
label_2462: RL D
label_2463: RL E
label_2464: RL H
label_2465: RL L
label_2466: RL (HL)
label_2467: RL (IX + 127)
label_2468: RL (IY - 128)
label_2469: RRC A
label_2470: RRC B
label_2471: RRC C
label_2472: RRC D
label_2473: RRC E
label_2474: RRC H
label_2475: RRC L
label_2476: RRC (HL)
label_2477: RRC (IX + 127)
label_2478: RRC (IY - 128)
label_2479: RR A
label_2480: RR B
label_2481: RR C
label_2482: RR D
label_2483: RR E
label_2484: RR H
label_2485: RR L
label_2486: RR (HL)
label_2487: RR (IX + 127)
label_2488: RR (IY - 128)
label_2489: SLA A
label_2490: SLA B
label_2491: SLA C
label_2492: SLA D
label_2493: SLA E
label_2494: SLA H
label_2495: SLA L
label_2496: SLA (HL)
label_2497: SLA (IX + 127)
label_2498: SLA (IY - 128)
label_2499: SRA A
label_2500: SRA B
label_2501: SRA C
label_2502: SRA D
label_2503: SRA E
label_2504: SRA H
label_2505: SRA L
label_2506: SRA (HL)
label_2507: SRA (IX + 127)
label_2508: SRA (IY - 128)
label_2509: SRL A
label_2510: SRL B
label_2511: SRL C
label_2512: SRL D
label_2513: SRL E
label_2514: SRL H
label_2515: SRL L
label_2516: SRL (HL)
label_2517: SRL (IX + 127)
label_2518: SRL (IY - 128)
label_2519: RLD
label_2520: RRD
label_2521: BIT 0,A
label_2522: BIT 1,A
label_2523: BIT 2,A
label_2524: BIT 3,A
label_2525: BIT 4,A
label_2526: BIT 5,A
label_2527: BIT 6,A
label_2528: BIT 7,A
label_2529: BIT 0,B
label_2530: BIT 1,B
label_2531: BIT 2,B
label_2532: BIT 3,B
label_2533: BIT 4,B
label_2534: BIT 5,B
label_2535: BIT 6,B
label_2536: BIT 7,B
label_2537: BIT 0,C
label_2538: BIT 1,C
label_2539: BIT 2,C
label_2540: BIT 3,C
label_2541: BIT 4,C
label_2542: BIT 5,C
label_2543: BIT 6,C
label_2544: BIT 7,C
label_2545: BIT 0,D
label_2546: BIT 1,D
label_2547: BIT 2,D
label_2548: BIT 3,D
label_2549: BIT 4,D
label_2550: BIT 5,D
label_2551: BIT 6,D
label_2552: BIT 7,D
label_2553: BIT 0,E
label_2554: BIT 1,E
label_2555: BIT 2,E
label_2556: BIT 3,E
label_2557: BIT 4,E
label_2558: BIT 5,E
label_2559: BIT 6,E
label_2560: BIT 7,E
label_2561: BIT 0,H
label_2562: BIT 1,H
label_2563: BIT 2,H
label_2564: BIT 3,H
label_2565: BIT 4,H
label_2566: BIT 5,H
label_2567: BIT 6,H
label_2568: BIT 7,H
label_2569: BIT 0,L
label_2570: BIT 1,L
label_2571: BIT 2,L
label_2572: BIT 3,L
label_2573: BIT 4,L
label_2574: BIT 5,L
label_2575: BIT 6,L
label_2576: BIT 7,L
label_2577: BIT 0,(HL)
label_2578: BIT 1,(HL)
label_2579: BIT 2,(HL)
label_2580: BIT 3,(HL)
label_2581: BIT 4,(HL)
label_2582: BIT 5,(HL)
label_2583: BIT 6,(HL)
label_2584: BIT 7,(HL)
label_2585: BIT 0,(IX + 127)
label_2586: BIT 1,(IX + 127)
label_2587: BIT 2,(IX + 127)
label_2588: BIT 3,(IX + 127)
label_2589: BIT 4,(IX + 127)
label_2590: BIT 5,(IX + 127)
label_2591: BIT 6,(IX + 127)
label_2592: BIT 7,(IX + 127)
label_2593: BIT 0,(IY - 128)
label_2594: BIT 1,(IY - 128)
label_2595: BIT 2,(IY - 128)
label_2596: BIT 3,(IY - 128)
label_2597: BIT 4,(IY - 128)
label_2598: BIT 5,(IY - 128)
label_2599: BIT 6,(IY - 128)
label_2600: BIT 7,(IY - 128)
label_2601: SET 0,A
label_2602: SET 1,A
label_2603: SET 2,A
label_2604: SET 3,A
label_2605: SET 4,A
label_2606: SET 5,A
label_2607: SET 6,A
label_2608: SET 7,A
label_2609: SET 0,B
label_2610: SET 1,B
label_2611: SET 2,B
label_2612: SET 3,B
label_2613: SET 4,B
label_2614: SET 5,B
label_2615: SET 6,B
label_2616: SET 7,B
label_2617: SET 0,C
label_2618: SET 1,C
label_2619: SET 2,C
label_2620: SET 3,C
label_2621: SET 4,C
label_2622: SET 5,C
label_2623: SET 6,C
label_2624: SET 7,C
label_2625: SET 0,D
label_2626: SET 1,D
label_2627: SET 2,D
label_2628: SET 3,D
label_2629: SET 4,D
label_2630: SET 5,D
label_2631: SET 6,D
label_2632: SET 7,D
label_2633: SET 0,E
label_2634: SET 1,E
label_2635: SET 2,E
label_2636: SET 3,E
label_2637: SET 4,E
label_2638: SET 5,E
label_2639: SET 6,E
label_2640: SET 7,E
label_2641: SET 0,H
label_2642: SET 1,H
label_2643: SET 2,H
label_2644: SET 3,H
label_2645: SET 4,H
label_2646: SET 5,H
label_2647: SET 6,H
label_2648: SET 7,H
label_2649: SET 0,L
label_2650: SET 1,L
label_2651: SET 2,L
label_2652: SET 3,L
label_2653: SET 4,L
label_2654: SET 5,L
label_2655: SET 6,L
label_2656: SET 7,L
label_2657: SET 0,(HL)
label_2658: SET 1,(HL)
label_2659: SET 2,(HL)
label_2660: SET 3,(HL)
label_2661: SET 4,(HL)
label_2662: SET 5,(HL)
label_2663: SET 6,(HL)
label_2664: SET 7,(HL)
label_2665: SET 0,(IX + 127)
label_2666: SET 1,(IX + 127)
label_2667: SET 2,(IX + 127)
label_2668: SET 3,(IX + 127)
label_2669: SET 4,(IX + 127)
label_2670: SET 5,(IX + 127)
label_2671: SET 6,(IX + 127)
label_2672: SET 7,(IX + 127)
label_2673: SET 0,(IY - 128)
label_2674: SET 1,(IY - 128)
label_2675: SET 2,(IY - 128)
label_2676: SET 3,(IY - 128)
label_2677: SET 4,(IY - 128)
label_2678: SET 5,(IY - 128)
label_2679: SET 6,(IY - 128)
label_2680: SET 7,(IY - 128)
label_2681: RES 0,A
label_2682: RES 1,A
label_2683: RES 2,A
label_2684: RES 3,A
label_2685: RES 4,A
label_2686: RES 5,A
label_2687: RES 6,A
label_2688: RES 7,A
label_2689: RES 0,B
label_2690: RES 1,B
label_2691: RES 2,B
label_2692: RES 3,B
label_2693: RES 4,B
label_2694: RES 5,B
label_2695: RES 6,B
label_2696: RES 7,B
label_2697: RES 0,C
label_2698: RES 1,C
label_2699: RES 2,C
label_2700: RES 3,C
label_2701: RES 4,C
label_2702: RES 5,C
label_2703: RES 6,C
label_2704: RES 7,C
label_2705: RES 0,D
label_2706: RES 1,D
label_2707: RES 2,D
label_2708: RES 3,D
label_2709: RES 4,D
label_2710: RES 5,D
label_2711: RES 6,D
label_2712: RES 7,D
label_2713: RES 0,E
label_2714: RES 1,E
label_2715: RES 2,E
label_2716: RES 3,E
label_2717: RES 4,E
label_2718: RES 5,E
label_2719: RES 6,E
label_2720: RES 7,E
label_2721: RES 0,H
label_2722: RES 1,H
label_2723: RES 2,H
label_2724: RES 3,H
label_2725: RES 4,H
label_2726: RES 5,H
label_2727: RES 6,H
label_2728: RES 7,H
label_2729: RES 0,L
label_2730: RES 1,L
label_2731: RES 2,L
label_2732: RES 3,L
label_2733: RES 4,L
label_2734: RES 5,L
label_2735: RES 6,L
label_2736: RES 7,L
label_2737: RES 0,(HL)
label_2738: RES 1,(HL)
label_2739: RES 2,(HL)
label_2740: RES 3,(HL)
label_2741: RES 4,(HL)
label_2742: RES 5,(HL)
label_2743: RES 6,(HL)
label_2744: RES 7,(HL)
label_2745: RES 0,(IX + 127)
label_2746: RES 1,(IX + 127)
label_2747: RES 2,(IX + 127)
label_2748: RES 3,(IX + 127)
label_2749: RES 4,(IX + 127)
label_2750: RES 5,(IX + 127)
label_2751: RES 6,(IX + 127)
label_2752: RES 7,(IX + 127)
label_2753: RES 0,(IY - 128)
label_2754: RES 1,(IY - 128)
label_2755: RES 2,(IY - 128)
label_2756: RES 3,(IY - 128)
label_2757: RES 4,(IY - 128)
label_2758: RES 5,(IY - 128)
label_2759: RES 6,(IY - 128)
label_2760: RES 7,(IY - 128)
label_2761: JP $5678
label_2762: JP NZ,$5678
label_2763: JP Z,$5678
label_2764: JP NC,$5678
label_2765: JP C,$5678
label_2766: JP PO,$5678
label_2767: JP PE,$5678
label_2768: JP P,$5678
label_2769: JP M,$5678
label_2770: JR $ + 2
label_2771: JR NZ,$ + 2
label_2772: JR Z,$ + 2
label_2773: JR NC,$ + 2
label_2774: JR C,$ + 2
label_2775: JP (HL)
label_2776: JP (IX)
label_2777: JP (IY)
label_2778: DJNZ $ + 2
label_2779: CALL $5678
label_2780: CALL NZ,$5678
label_2781: CALL Z,$5678
label_2782: CALL NC,$5678
label_2783: CALL C,$5678
label_2784: CALL PO,$5678
label_2785: CALL PE,$5678
label_2786: CALL P,$5678
label_2787: CALL M,$5678
label_2788: RET
label_2789: RET NZ
label_2790: RET Z
label_2791: RET NC
label_2792: RET C
label_2793: RET PO
label_2794: RET PE
label_2795: RET P
label_2796: RET M
label_2797: RETI
label_2798: RETN
label_2799: RST $00
label_2800: RST $08
label_2801: RST $10
label_2802: RST $18
label_2803: RST $20
label_2804: RST $28
label_2805: RST $30
label_2806: RST $38
label_2807: IN A,($12)
label_2808: IN A,(C)
label_2809: IN B,(C)
label_2810: IN C,(C)
label_2811: IN D,(C)
label_2812: IN E,(C)
label_2813: IN H,(C)
label_2814: IN L,(C)
label_2815: IN F,(C)
label_2816: INI
label_2817: INIR
label_2818: IND
label_2819: INDR
label_2820: OUT ($12),A
label_2821: OUT (C),A
label_2822: OUT (C),B
label_2823: OUT (C),C
label_2824: OUT (C),D
label_2825: OUT (C),E
label_2826: OUT (C),H
label_2827: OUT (C),L
label_2828: OUTI
label_2829: OTIR
label_2830: OUTD
label_2831: OTDR
label_2832: LD A,A
label_2833: LD A,B
label_2834: LD A,C
label_2835: LD A,D
label_2836: LD A,E
label_2837: LD A,H
label_2838: LD A,L
label_2839: LD B,A
label_2840: LD B,B
label_2841: LD B,C
label_2842: LD B,D
label_2843: LD B,E
label_2844: LD B,H
label_2845: LD B,L
label_2846: LD C,A
label_2847: LD C,B
label_2848: LD C,C
label_2849: LD C,D
label_2850: LD C,E
label_2851: LD C,H
label_2852: LD C,L
label_2853: LD D,A
label_2854: LD D,B
label_2855: LD D,C
label_2856: LD D,D
label_2857: LD D,E
label_2858: LD D,H
label_2859: LD D,L
label_2860: LD E,A
label_2861: LD E,B
label_2862: LD E,C
label_2863: LD E,D
label_2864: LD E,E
label_2865: LD E,H
label_2866: LD E,L
label_2867: LD H,A
label_2868: LD H,B
label_2869: LD H,C
label_2870: LD H,D
label_2871: LD H,E
label_2872: LD H,H
label_2873: LD H,L
label_2874: LD L,A
label_2875: LD L,B
label_2876: LD L,C
label_2877: LD L,D
label_2878: LD L,E
label_2879: LD L,H
label_2880: LD L,L
label_2881: LD A,$12
label_2882: LD B,$12
label_2883: LD C,$12
label_2884: LD D,$12
label_2885: LD E,$12
label_2886: LD H,$12
label_2887: LD L,$12
label_2888: LD A,(HL)
label_2889: LD B,(HL)
label_2890: LD C,(HL)
label_2891: LD D,(HL)
label_2892: LD E,(HL)
label_2893: LD H,(HL)
label_2894: LD L,(HL)
label_2895: LD A,(IX + 127)
label_2896: LD B,(IX + 127)
label_2897: LD C,(IX + 127)
label_2898: LD D,(IX + 127)
label_2899: LD E,(IX + 127)
label_2900: LD H,(IX + 127)
label_2901: LD L,(IX + 127)
label_2902: LD A,(IY - 128)
label_2903: LD B,(IY - 128)
label_2904: LD C,(IY - 128)
label_2905: LD D,(IY - 128)
label_2906: LD E,(IY - 128)
label_2907: LD H,(IY - 128)
label_2908: LD L,(IY - 128)
label_2909: LD (HL),A
label_2910: LD (HL),B
label_2911: LD (HL),C
label_2912: LD (HL),D
label_2913: LD (HL),E
label_2914: LD (HL),H
label_2915: LD (HL),L
label_2916: LD (IX + 127),A
label_2917: LD (IX + 127),B
label_2918: LD (IX + 127),C
label_2919: LD (IX + 127),D
label_2920: LD (IX + 127),E
label_2921: LD (IX + 127),H
label_2922: LD (IX + 127),L
label_2923: LD (IY - 128),A
label_2924: LD (IY - 128),B
label_2925: LD (IY - 128),C
label_2926: LD (IY - 128),D
label_2927: LD (IY - 128),E
label_2928: LD (IY - 128),H
label_2929: LD (IY - 128),L
label_2930: LD (HL),$12
label_2931: LD (IX + 127),$12
label_2932: LD (IY - 128),$12
label_2933: LD A,(BC)
label_2934: LD A,(DE)
label_2935: LD A,($5678)
label_2936: LD (BC),A
label_2937: LD (DE),A
label_2938: LD ($5678),A
label_2939: LD A,I
label_2940: LD A,R
label_2941: LD I,A
label_2942: LD R,A
label_2943: LD BC,$5678
label_2944: LD DE,$5678
label_2945: LD HL,$5678
label_2946: LD SP,$5678
label_2947: LD IX,$5678
label_2948: LD IY,$5678
label_2949: LD HL,($5678)
label_2950: LD BC,($5678)
label_2951: LD DE,($5678)
label_2952: LD HL,($5678)
label_2953: LD SP,($5678)
label_2954: LD IX,($5678)
label_2955: LD IY,($5678)
label_2956: LD ($5678),HL
label_2957: LD ($5678),BC
label_2958: LD ($5678),DE
label_2959: LD ($5678),HL
label_2960: LD ($5678),SP
label_2961: LD ($5678),IX
label_2962: LD ($5678),IY
label_2963: LD SP,HL
label_2964: LD SP,IX
label_2965: LD SP,IY
label_2966: PUSH BC
label_2967: PUSH DE
label_2968: PUSH HL
label_2969: PUSH AF
label_2970: PUSH IX
label_2971: PUSH IY
label_2972: POP BC
label_2973: POP DE
label_2974: POP HL
label_2975: POP AF
label_2976: POP IX
label_2977: POP IY
label_2978: EX DE,HL
label_2979: EX AF,AF'
label_2980: EXX
label_2981: EX (SP),HL
label_2982: EX (SP),IX
label_2983: EX (SP),IY
label_2984: LDI
label_2985: LDIR
label_2986: LDD
label_2987: LDDR
label_2988: CPI
label_2989: CPIR
label_2990: CPD
label_2991: CPDR
label_2992: ADD A,A
label_2993: ADD A,B
label_2994: ADD A,C
label_2995: ADD A,D
label_2996: ADD A,E
label_2997: ADD A,H
label_2998: ADD A,L
label_2999: ADD A,$12
label_3000: ADD A,(HL)
label_3001: ADD A,(IX + 127)
label_3002: ADD A,(IY - 128)
label_3003: ADC A,A
label_3004: ADC A,B
label_3005: ADC A,C
label_3006: ADC A,D
label_3007: ADC A,E
label_3008: ADC A,H
label_3009: ADC A,L
label_3010: ADC A,$12
label_3011: ADC A,(HL)
label_3012: ADC A,(IX + 127)
label_3013: ADC A,(IY - 128)
label_3014: SUB A
label_3015: SUB B
label_3016: SUB C
label_3017: SUB D
label_3018: SUB E
label_3019: SUB H
label_3020: SUB L
label_3021: SUB $12
label_3022: SUB (HL)
label_3023: SUB (IX + 127)
label_3024: SUB (IY - 128)
label_3025: SBC A,A
label_3026: SBC A,B
label_3027: SBC A,C
label_3028: SBC A,D
label_3029: SBC A,E
label_3030: SBC A,H
label_3031: SBC A,L
label_3032: SBC A,$12
label_3033: SBC A,(HL)
label_3034: SBC A,(IX + 127)
label_3035: SBC A,(IY - 128)
label_3036: AND A
label_3037: AND B
label_3038: AND C
label_3039: AND D
label_3040: AND E
label_3041: AND H
label_3042: AND L
label_3043: AND $12
label_3044: AND (HL)
label_3045: AND (IX + 127)
label_3046: AND (IY - 128)
label_3047: AND A
label_3048: AND B
label_3049: AND C
label_3050: AND D
label_3051: AND E
label_3052: AND H
label_3053: AND L
label_3054: AND $12
label_3055: AND (HL)
label_3056: AND (IX + 127)
label_3057: AND (IY - 128)
label_3058: OR A
label_3059: OR B
label_3060: OR C
label_3061: OR D
label_3062: OR E
label_3063: OR H
label_3064: OR L
label_3065: OR $12
label_3066: OR (HL)
label_3067: OR (IX + 127)
label_3068: OR (IY - 128)
label_3069: XOR A
label_3070: XOR B
label_3071: XOR C
label_3072: XOR D
label_3073: XOR E
label_3074: XOR H
label_3075: XOR L
label_3076: XOR $12
label_3077: XOR (HL)
label_3078: XOR (IX + 127)
label_3079: XOR (IY - 128)
label_3080: CP A
label_3081: CP B
label_3082: CP C
label_3083: CP D
label_3084: CP E
label_3085: CP H
label_3086: CP L
label_3087: CP $12
label_3088: CP (HL)
label_3089: CP (IX + 127)
label_3090: CP (IY - 128)
label_3091: INC A
label_3092: INC B
label_3093: INC C
label_3094: INC D
label_3095: INC E
label_3096: INC H
label_3097: INC L
label_3098: INC (HL)
label_3099: INC (IX + 127)
label_3100: INC (IY - 128)
label_3101: DEC A
label_3102: DEC B
label_3103: DEC C
label_3104: DEC D
label_3105: DEC E
label_3106: DEC H
label_3107: DEC L
label_3108: DEC (HL)
label_3109: DEC (IX + 127)
label_3110: DEC (IY - 128)
label_3111: DAA
label_3112: CPL
label_3113: NEG
label_3114: CCF
label_3115: SCF
label_3116: NOP
label_3117: HALT
label_3118: DI
label_3119: EI
label_3120: IM 0
label_3121: IM 1
label_3122: IM 2
label_3123: ADD HL,BC
label_3124: ADD HL,DE
label_3125: ADD HL,HL
label_3126: ADD HL,SP
label_3127: ADC HL,BC
label_3128: ADC HL,DE
label_3129: ADC HL,HL
label_3130: ADC HL,SP
label_3131: SBC HL,BC
label_3132: SBC HL,DE
label_3133: SBC HL,HL
label_3134: SBC HL,SP
label_3135: ADD IX,BC
label_3136: ADD IX,DE
label_3137: ADD IX,SP
label_3138: ADD IY,BC
label_3139: ADD IY,DE
label_3140: ADD IY,SP
label_3141: INC BC
label_3142: INC DE
label_3143: INC HL
label_3144: INC SP
label_3145: INC IX
label_3146: INC IY
label_3147: DEC BC
label_3148: DEC DE
label_3149: DEC HL
label_3150: DEC SP
label_3151: DEC IX
label_3152: DEC IY
label_3153: RLCA
label_3154: RLA
label_3155: RRCA
label_3156: RRA
label_3157: RLC A
label_3158: RLC B
label_3159: RLC C
label_3160: RLC D
label_3161: RLC E
label_3162: RLC H
label_3163: RLC L
label_3164: RLC (HL)
label_3165: RLC (IX + 127)
label_3166: RLC (IY - 128)
label_3167: RL A
label_3168: RL B
label_3169: RL C
label_3170: RL D
label_3171: RL E
label_3172: RL H
label_3173: RL L
label_3174: RL (HL)
label_3175: RL (IX + 127)
label_3176: RL (IY - 128)
label_3177: RRC A
label_3178: RRC B
label_3179: RRC C
label_3180: RRC D
label_3181: RRC E
label_3182: RRC H
label_3183: RRC L
label_3184: RRC (HL)
label_3185: RRC (IX + 127)
label_3186: RRC (IY - 128)
label_3187: RR A
label_3188: RR B
label_3189: RR C
label_3190: RR D
label_3191: RR E
label_3192: RR H
label_3193: RR L
label_3194: RR (HL)
label_3195: RR (IX + 127)
label_3196: RR (IY - 128)
label_3197: SLA A
label_3198: SLA B
label_3199: SLA C
label_3200: SLA D
label_3201: SLA E
label_3202: SLA H
label_3203: SLA L
label_3204: SLA (HL)
label_3205: SLA (IX + 127)
label_3206: SLA (IY - 128)
label_3207: SRA A
label_3208: SRA B
label_3209: SRA C
label_3210: SRA D
label_3211: SRA E
label_3212: SRA H
label_3213: SRA L
label_3214: SRA (HL)
label_3215: SRA (IX + 127)
label_3216: SRA (IY - 128)
label_3217: SRL A
label_3218: SRL B
label_3219: SRL C
label_3220: SRL D
label_3221: SRL E
label_3222: SRL H
label_3223: SRL L
label_3224: SRL (HL)
label_3225: SRL (IX + 127)
label_3226: SRL (IY - 128)
label_3227: RLD
label_3228: RRD
label_3229: BIT 0,A
label_3230: BIT 1,A
label_3231: BIT 2,A
label_3232: BIT 3,A
label_3233: BIT 4,A
label_3234: BIT 5,A
label_3235: BIT 6,A
label_3236: BIT 7,A
label_3237: BIT 0,B
label_3238: BIT 1,B
label_3239: BIT 2,B
label_3240: BIT 3,B
label_3241: BIT 4,B
label_3242: BIT 5,B
label_3243: BIT 6,B
label_3244: BIT 7,B
label_3245: BIT 0,C
label_3246: BIT 1,C
label_3247: BIT 2,C
label_3248: BIT 3,C
label_3249: BIT 4,C
label_3250: BIT 5,C
label_3251: BIT 6,C
label_3252: BIT 7,C
label_3253: BIT 0,D
label_3254: BIT 1,D
label_3255: BIT 2,D
label_3256: BIT 3,D
label_3257: BIT 4,D
label_3258: BIT 5,D
label_3259: BIT 6,D
label_3260: BIT 7,D
label_3261: BIT 0,E
label_3262: BIT 1,E
label_3263: BIT 2,E
label_3264: BIT 3,E
label_3265: BIT 4,E
label_3266: BIT 5,E
label_3267: BIT 6,E
label_3268: BIT 7,E
label_3269: BIT 0,H
label_3270: BIT 1,H
label_3271: BIT 2,H
label_3272: BIT 3,H
label_3273: BIT 4,H
label_3274: BIT 5,H
label_3275: BIT 6,H
label_3276: BIT 7,H
label_3277: BIT 0,L
label_3278: BIT 1,L
label_3279: BIT 2,L
label_3280: BIT 3,L
label_3281: BIT 4,L
label_3282: BIT 5,L
label_3283: BIT 6,L
label_3284: BIT 7,L
label_3285: BIT 0,(HL)
label_3286: BIT 1,(HL)
label_3287: BIT 2,(HL)
label_3288: BIT 3,(HL)
label_3289: BIT 4,(HL)
label_3290: BIT 5,(HL)
label_3291: BIT 6,(HL)
label_3292: BIT 7,(HL)
label_3293: BIT 0,(IX + 127)
label_3294: BIT 1,(IX + 127)
label_3295: BIT 2,(IX + 127)
label_3296: BIT 3,(IX + 127)
label_3297: BIT 4,(IX + 127)
label_3298: BIT 5,(IX + 127)
label_3299: BIT 6,(IX + 127)
label_3300: BIT 7,(IX + 127)
label_3301: BIT 0,(IY - 128)
label_3302: BIT 1,(IY - 128)
label_3303: BIT 2,(IY - 128)
label_3304: BIT 3,(IY - 128)
label_3305: BIT 4,(IY - 128)
label_3306: BIT 5,(IY - 128)
label_3307: BIT 6,(IY - 128)
label_3308: BIT 7,(IY - 128)
label_3309: SET 0,A
label_3310: SET 1,A
label_3311: SET 2,A
label_3312: SET 3,A
label_3313: SET 4,A
label_3314: SET 5,A
label_3315: SET 6,A
label_3316: SET 7,A
label_3317: SET 0,B
label_3318: SET 1,B
label_3319: SET 2,B
label_3320: SET 3,B
label_3321: SET 4,B
label_3322: SET 5,B
label_3323: SET 6,B
label_3324: SET 7,B
label_3325: SET 0,C
label_3326: SET 1,C
label_3327: SET 2,C
label_3328: SET 3,C
label_3329: SET 4,C
label_3330: SET 5,C
label_3331: SET 6,C
label_3332: SET 7,C
label_3333: SET 0,D
label_3334: SET 1,D
label_3335: SET 2,D
label_3336: SET 3,D
label_3337: SET 4,D
label_3338: SET 5,D
label_3339: SET 6,D
label_3340: SET 7,D
label_3341: SET 0,E
label_3342: SET 1,E
label_3343: SET 2,E
label_3344: SET 3,E
label_3345: SET 4,E
label_3346: SET 5,E
label_3347: SET 6,E
label_3348: SET 7,E
label_3349: SET 0,H
label_3350: SET 1,H
label_3351: SET 2,H
label_3352: SET 3,H
label_3353: SET 4,H
label_3354: SET 5,H
label_3355: SET 6,H
label_3356: SET 7,H
label_3357: SET 0,L
label_3358: SET 1,L
label_3359: SET 2,L
label_3360: SET 3,L
label_3361: SET 4,L
label_3362: SET 5,L
label_3363: SET 6,L
label_3364: SET 7,L
label_3365: SET 0,(HL)
label_3366: SET 1,(HL)
label_3367: SET 2,(HL)
label_3368: SET 3,(HL)
label_3369: SET 4,(HL)
label_3370: SET 5,(HL)
label_3371: SET 6,(HL)
label_3372: SET 7,(HL)
label_3373: SET 0,(IX + 127)
label_3374: SET 1,(IX + 127)
label_3375: SET 2,(IX + 127)
label_3376: SET 3,(IX + 127)
label_3377: SET 4,(IX + 127)
label_3378: SET 5,(IX + 127)
label_3379: SET 6,(IX + 127)
label_3380: SET 7,(IX + 127)
label_3381: SET 0,(IY - 128)
label_3382: SET 1,(IY - 128)
label_3383: SET 2,(IY - 128)
label_3384: SET 3,(IY - 128)
label_3385: SET 4,(IY - 128)
label_3386: SET 5,(IY - 128)
label_3387: SET 6,(IY - 128)
label_3388: SET 7,(IY - 128)
label_3389: RES 0,A
label_3390: RES 1,A
label_3391: RES 2,A
label_3392: RES 3,A
label_3393: RES 4,A
label_3394: RES 5,A
label_3395: RES 6,A
label_3396: RES 7,A
label_3397: RES 0,B
label_3398: RES 1,B
label_3399: RES 2,B
label_3400: RES 3,B
label_3401: RES 4,B
label_3402: RES 5,B
label_3403: RES 6,B
label_3404: RES 7,B
label_3405: RES 0,C
label_3406: RES 1,C
label_3407: RES 2,C
label_3408: RES 3,C
label_3409: RES 4,C
label_3410: RES 5,C
label_3411: RES 6,C
label_3412: RES 7,C
label_3413: RES 0,D
label_3414: RES 1,D
label_3415: RES 2,D
label_3416: RES 3,D
label_3417: RES 4,D
label_3418: RES 5,D
label_3419: RES 6,D
label_3420: RES 7,D
label_3421: RES 0,E
label_3422: RES 1,E
label_3423: RES 2,E
label_3424: RES 3,E
label_3425: RES 4,E
label_3426: RES 5,E
label_3427: RES 6,E
label_3428: RES 7,E
label_3429: RES 0,H
label_3430: RES 1,H
label_3431: RES 2,H
label_3432: RES 3,H
label_3433: RES 4,H
label_3434: RES 5,H
label_3435: RES 6,H
label_3436: RES 7,H
label_3437: RES 0,L
label_3438: RES 1,L
label_3439: RES 2,L
label_3440: RES 3,L
label_3441: RES 4,L
label_3442: RES 5,L
label_3443: RES 6,L
label_3444: RES 7,L
label_3445: RES 0,(HL)
label_3446: RES 1,(HL)
label_3447: RES 2,(HL)
label_3448: RES 3,(HL)
label_3449: RES 4,(HL)
label_3450: RES 5,(HL)
label_3451: RES 6,(HL)
label_3452: RES 7,(HL)
label_3453: RES 0,(IX + 127)
label_3454: RES 1,(IX + 127)
label_3455: RES 2,(IX + 127)
label_3456: RES 3,(IX + 127)
label_3457: RES 4,(IX + 127)
label_3458: RES 5,(IX + 127)
label_3459: RES 6,(IX + 127)
label_3460: RES 7,(IX + 127)
label_3461: RES 0,(IY - 128)
label_3462: RES 1,(IY - 128)
label_3463: RES 2,(IY - 128)
label_3464: RES 3,(IY - 128)
label_3465: RES 4,(IY - 128)
label_3466: RES 5,(IY - 128)
label_3467: RES 6,(IY - 128)
label_3468: RES 7,(IY - 128)
label_3469: JP $5678
label_3470: JP NZ,$5678
label_3471: JP Z,$5678
label_3472: JP NC,$5678
label_3473: JP C,$5678
label_3474: JP PO,$5678
label_3475: JP PE,$5678
label_3476: JP P,$5678
label_3477: JP M,$5678
label_3478: JR $ + 2
label_3479: JR NZ,$ + 2
label_3480: JR Z,$ + 2
label_3481: JR NC,$ + 2
label_3482: JR C,$ + 2
label_3483: JP (HL)
label_3484: JP (IX)
label_3485: JP (IY)
label_3486: DJNZ $ + 2
label_3487: CALL $5678
label_3488: CALL NZ,$5678
label_3489: CALL Z,$5678
label_3490: CALL NC,$5678
label_3491: CALL C,$5678
label_3492: CALL PO,$5678
label_3493: CALL PE,$5678
label_3494: CALL P,$5678
label_3495: CALL M,$5678
label_3496: RET
label_3497: RET NZ
label_3498: RET Z
label_3499: RET NC
label_3500: RET C
label_3501: RET PO
label_3502: RET PE
label_3503: RET P
label_3504: RET M
label_3505: RETI
label_3506: RETN
label_3507: RST $00
label_3508: RST $08
label_3509: RST $10
label_3510: RST $18
label_3511: RST $20
label_3512: RST $28
label_3513: RST $30
label_3514: RST $38
label_3515: IN A,($12)
label_3516: IN A,(C)
label_3517: IN B,(C)
label_3518: IN C,(C)
label_3519: IN D,(C)
label_3520: IN E,(C)
label_3521: IN H,(C)
label_3522: IN L,(C)
label_3523: IN F,(C)
label_3524: INI
label_3525: INIR
label_3526: IND
label_3527: INDR
label_3528: OUT ($12),A
label_3529: OUT (C),A
label_3530: OUT (C),B
label_3531: OUT (C),C
label_3532: OUT (C),D
label_3533: OUT (C),E
label_3534: OUT (C),H
label_3535: OUT (C),L
label_3536: OUTI
label_3537: OTIR
label_3538: OUTD
label_3539: OTDR
label_3540: LD A,A
label_3541: LD A,B
label_3542: LD A,C
label_3543: LD A,D
label_3544: LD A,E
label_3545: LD A,H
label_3546: LD A,L
label_3547: LD B,A
label_3548: LD B,B
label_3549: LD B,C
label_3550: LD B,D
label_3551: LD B,E
label_3552: LD B,H
label_3553: LD B,L
label_3554: LD C,A
label_3555: LD C,B
label_3556: LD C,C
label_3557: LD C,D
label_3558: LD C,E
label_3559: LD C,H
label_3560: LD C,L
label_3561: LD D,A
label_3562: LD D,B
label_3563: LD D,C
label_3564: LD D,D
label_3565: LD D,E
label_3566: LD D,H
label_3567: LD D,L
label_3568: LD E,A
label_3569: LD E,B
label_3570: LD E,C
label_3571: LD E,D
label_3572: LD E,E
label_3573: LD E,H
label_3574: LD E,L
label_3575: LD H,A
label_3576: LD H,B
label_3577: LD H,C
label_3578: LD H,D
label_3579: LD H,E
label_3580: LD H,H
label_3581: LD H,L
label_3582: LD L,A
label_3583: LD L,B
label_3584: LD L,C
label_3585: LD L,D
label_3586: LD L,E
label_3587: LD L,H
label_3588: LD L,L
label_3589: LD A,$12
label_3590: LD B,$12
label_3591: LD C,$12
label_3592: LD D,$12
label_3593: LD E,$12
label_3594: LD H,$12
label_3595: LD L,$12
label_3596: LD A,(HL)
label_3597: LD B,(HL)
label_3598: LD C,(HL)
label_3599: LD D,(HL)
label_3600: LD E,(HL)
label_3601: LD H,(HL)
label_3602: LD L,(HL)
label_3603: LD A,(IX + 127)
label_3604: LD B,(IX + 127)
label_3605: LD C,(IX + 127)
label_3606: LD D,(IX + 127)
label_3607: LD E,(IX + 127)
label_3608: LD H,(IX + 127)
label_3609: LD L,(IX + 127)
label_3610: LD A,(IY - 128)
label_3611: LD B,(IY - 128)
label_3612: LD C,(IY - 128)
label_3613: LD D,(IY - 128)
label_3614: LD E,(IY - 128)
label_3615: LD H,(IY - 128)
label_3616: LD L,(IY - 128)
label_3617: LD (HL),A
label_3618: LD (HL),B
label_3619: LD (HL),C
label_3620: LD (HL),D
label_3621: LD (HL),E
label_3622: LD (HL),H
label_3623: LD (HL),L
label_3624: LD (IX + 127),A
label_3625: LD (IX + 127),B
label_3626: LD (IX + 127),C
label_3627: LD (IX + 127),D
label_3628: LD (IX + 127),E
label_3629: LD (IX + 127),H
label_3630: LD (IX + 127),L
label_3631: LD (IY - 128),A
label_3632: LD (IY - 128),B
label_3633: LD (IY - 128),C
label_3634: LD (IY - 128),D
label_3635: LD (IY - 128),E
label_3636: LD (IY - 128),H
label_3637: LD (IY - 128),L
label_3638: LD (HL),$12
label_3639: LD (IX + 127),$12
label_3640: LD (IY - 128),$12
label_3641: LD A,(BC)
label_3642: LD A,(DE)
label_3643: LD A,($5678)
label_3644: LD (BC),A
label_3645: LD (DE),A
label_3646: LD ($5678),A
label_3647: LD A,I
label_3648: LD A,R
label_3649: LD I,A
label_3650: LD R,A
label_3651: LD BC,$5678
label_3652: LD DE,$5678
label_3653: LD HL,$5678
label_3654: LD SP,$5678
label_3655: LD IX,$5678
label_3656: LD IY,$5678
label_3657: LD HL,($5678)
label_3658: LD BC,($5678)
label_3659: LD DE,($5678)
label_3660: LD HL,($5678)
label_3661: LD SP,($5678)
label_3662: LD IX,($5678)
label_3663: LD IY,($5678)
label_3664: LD ($5678),HL
label_3665: LD ($5678),BC
label_3666: LD ($5678),DE
label_3667: LD ($5678),HL
label_3668: LD ($5678),SP
label_3669: LD ($5678),IX
label_3670: LD ($5678),IY
label_3671: LD SP,HL
label_3672: LD SP,IX
label_3673: LD SP,IY
label_3674: PUSH BC
label_3675: PUSH DE
label_3676: PUSH HL
label_3677: PUSH AF
label_3678: PUSH IX
label_3679: PUSH IY
label_3680: POP BC
label_3681: POP DE
label_3682: POP HL
label_3683: POP AF
label_3684: POP IX
label_3685: POP IY
label_3686: EX DE,HL
label_3687: EX AF,AF'
label_3688: EXX
label_3689: EX (SP),HL
label_3690: EX (SP),IX
label_3691: EX (SP),IY
label_3692: LDI
label_3693: LDIR
label_3694: LDD
label_3695: LDDR
label_3696: CPI
label_3697: CPIR
label_3698: CPD
label_3699: CPDR
label_3700: ADD A,A
label_3701: ADD A,B
label_3702: ADD A,C
label_3703: ADD A,D
label_3704: ADD A,E
label_3705: ADD A,H
label_3706: ADD A,L
label_3707: ADD A,$12
label_3708: ADD A,(HL)
label_3709: ADD A,(IX + 127)
label_3710: ADD A,(IY - 128)
label_3711: ADC A,A
label_3712: ADC A,B
label_3713: ADC A,C
label_3714: ADC A,D
label_3715: ADC A,E
label_3716: ADC A,H
label_3717: ADC A,L
label_3718: ADC A,$12
label_3719: ADC A,(HL)
label_3720: ADC A,(IX + 127)
label_3721: ADC A,(IY - 128)
label_3722: SUB A
label_3723: SUB B
label_3724: SUB C
label_3725: SUB D
label_3726: SUB E
label_3727: SUB H
label_3728: SUB L
label_3729: SUB $12
label_3730: SUB (HL)
label_3731: SUB (IX + 127)
label_3732: SUB (IY - 128)
label_3733: SBC A,A
label_3734: SBC A,B
label_3735: SBC A,C
label_3736: SBC A,D
label_3737: SBC A,E
label_3738: SBC A,H
label_3739: SBC A,L
label_3740: SBC A,$12
label_3741: SBC A,(HL)
label_3742: SBC A,(IX + 127)
label_3743: SBC A,(IY - 128)
label_3744: AND A
label_3745: AND B
label_3746: AND C
label_3747: AND D
label_3748: AND E
label_3749: AND H
label_3750: AND L
label_3751: AND $12
label_3752: AND (HL)
label_3753: AND (IX + 127)
label_3754: AND (IY - 128)
label_3755: AND A
label_3756: AND B
label_3757: AND C
label_3758: AND D
label_3759: AND E
label_3760: AND H
label_3761: AND L
label_3762: AND $12
label_3763: AND (HL)
label_3764: AND (IX + 127)
label_3765: AND (IY - 128)
label_3766: OR A
label_3767: OR B
label_3768: OR C
label_3769: OR D
label_3770: OR E
label_3771: OR H
label_3772: OR L
label_3773: OR $12
label_3774: OR (HL)
label_3775: OR (IX + 127)
label_3776: OR (IY - 128)
label_3777: XOR A
label_3778: XOR B
label_3779: XOR C
label_3780: XOR D
label_3781: XOR E
label_3782: XOR H
label_3783: XOR L
label_3784: XOR $12
label_3785: XOR (HL)
label_3786: XOR (IX + 127)
label_3787: XOR (IY - 128)
label_3788: CP A
label_3789: CP B
label_3790: CP C
label_3791: CP D
label_3792: CP E
label_3793: CP H
label_3794: CP L
label_3795: CP $12
label_3796: CP (HL)
label_3797: CP (IX + 127)
label_3798: CP (IY - 128)
label_3799: INC A
label_3800: INC B
label_3801: INC C
label_3802: INC D
label_3803: INC E
label_3804: INC H
label_3805: INC L
label_3806: INC (HL)
label_3807: INC (IX + 127)
label_3808: INC (IY - 128)
label_3809: DEC A
label_3810: DEC B
label_3811: DEC C
label_3812: DEC D
label_3813: DEC E
label_3814: DEC H
label_3815: DEC L
label_3816: DEC (HL)
label_3817: DEC (IX + 127)
label_3818: DEC (IY - 128)
label_3819: DAA
label_3820: CPL
label_3821: NEG
label_3822: CCF
label_3823: SCF
label_3824: NOP
label_3825: HALT
label_3826: DI
label_3827: EI
label_3828: IM 0
label_3829: IM 1
label_3830: IM 2
label_3831: ADD HL,BC
label_3832: ADD HL,DE
label_3833: ADD HL,HL
label_3834: ADD HL,SP
label_3835: ADC HL,BC
label_3836: ADC HL,DE
label_3837: ADC HL,HL
label_3838: ADC HL,SP
label_3839: SBC HL,BC
label_3840: SBC HL,DE
label_3841: SBC HL,HL
label_3842: SBC HL,SP
label_3843: ADD IX,BC
label_3844: ADD IX,DE
label_3845: ADD IX,SP
label_3846: ADD IY,BC
label_3847: ADD IY,DE
label_3848: ADD IY,SP
label_3849: INC BC
label_3850: INC DE
label_3851: INC HL
label_3852: INC SP
label_3853: INC IX
label_3854: INC IY
label_3855: DEC BC
label_3856: DEC DE
label_3857: DEC HL
label_3858: DEC SP
label_3859: DEC IX
label_3860: DEC IY
label_3861: RLCA
label_3862: RLA
label_3863: RRCA
label_3864: RRA
label_3865: RLC A
label_3866: RLC B
label_3867: RLC C
label_3868: RLC D
label_3869: RLC E
label_3870: RLC H
label_3871: RLC L
label_3872: RLC (HL)
label_3873: RLC (IX + 127)
label_3874: RLC (IY - 128)
label_3875: RL A
label_3876: RL B
label_3877: RL C
label_3878: RL D
label_3879: RL E
label_3880: RL H
label_3881: RL L
label_3882: RL (HL)
label_3883: RL (IX + 127)
label_3884: RL (IY - 128)
label_3885: RRC A
label_3886: RRC B
label_3887: RRC C
label_3888: RRC D
label_3889: RRC E
label_3890: RRC H
label_3891: RRC L
label_3892: RRC (HL)
label_3893: RRC (IX + 127)
label_3894: RRC (IY - 128)
label_3895: RR A
label_3896: RR B
label_3897: RR C
label_3898: RR D
label_3899: RR E
label_3900: RR H
label_3901: RR L
label_3902: RR (HL)
label_3903: RR (IX + 127)
label_3904: RR (IY - 128)
label_3905: SLA A
label_3906: SLA B
label_3907: SLA C
label_3908: SLA D
label_3909: SLA E
label_3910: SLA H
label_3911: SLA L
label_3912: SLA (HL)
label_3913: SLA (IX + 127)
label_3914: SLA (IY - 128)
label_3915: SRA A
label_3916: SRA B
label_3917: SRA C
label_3918: SRA D
label_3919: SRA E
label_3920: SRA H
label_3921: SRA L
label_3922: SRA (HL)
label_3923: SRA (IX + 127)
label_3924: SRA (IY - 128)
label_3925: SRL A
label_3926: SRL B
label_3927: SRL C
label_3928: SRL D
label_3929: SRL E
label_3930: SRL H
label_3931: SRL L
label_3932: SRL (HL)
label_3933: SRL (IX + 127)
label_3934: SRL (IY - 128)
label_3935: RLD
label_3936: RRD
label_3937: BIT 0,A
label_3938: BIT 1,A
label_3939: BIT 2,A
label_3940: BIT 3,A
label_3941: BIT 4,A
label_3942: BIT 5,A
label_3943: BIT 6,A
label_3944: BIT 7,A
label_3945: BIT 0,B
label_3946: BIT 1,B
label_3947: BIT 2,B
label_3948: BIT 3,B
label_3949: BIT 4,B
label_3950: BIT 5,B
label_3951: BIT 6,B
label_3952: BIT 7,B
label_3953: BIT 0,C
label_3954: BIT 1,C
label_3955: BIT 2,C
label_3956: BIT 3,C
label_3957: BIT 4,C
label_3958: BIT 5,C
label_3959: BIT 6,C
label_3960: BIT 7,C
label_3961: BIT 0,D
label_3962: BIT 1,D
label_3963: BIT 2,D
label_3964: BIT 3,D
label_3965: BIT 4,D
label_3966: BIT 5,D
label_3967: BIT 6,D
label_3968: BIT 7,D
label_3969: BIT 0,E
label_3970: BIT 1,E
label_3971: BIT 2,E
label_3972: BIT 3,E
label_3973: BIT 4,E
label_3974: BIT 5,E
label_3975: BIT 6,E
label_3976: BIT 7,E
label_3977: BIT 0,H
label_3978: BIT 1,H
label_3979: BIT 2,H
label_3980: BIT 3,H
label_3981: BIT 4,H
label_3982: BIT 5,H
label_3983: BIT 6,H
label_3984: BIT 7,H
label_3985: BIT 0,L
label_3986: BIT 1,L
label_3987: BIT 2,L
label_3988: BIT 3,L
label_3989: BIT 4,L
label_3990: BIT 5,L
label_3991: BIT 6,L
label_3992: BIT 7,L
label_3993: BIT 0,(HL)
label_3994: BIT 1,(HL)
label_3995: BIT 2,(HL)
label_3996: BIT 3,(HL)
label_3997: BIT 4,(HL)
label_3998: BIT 5,(HL)
label_3999: BIT 6,(HL)
label_4000: BIT 7,(HL)
label_4001: BIT 0,(IX + 127)
label_4002: BIT 1,(IX + 127)
label_4003: BIT 2,(IX + 127)
label_4004: BIT 3,(IX + 127)
label_4005: BIT 4,(IX + 127)
label_4006: BIT 5,(IX + 127)
label_4007: BIT 6,(IX + 127)
label_4008: BIT 7,(IX + 127)
label_4009: BIT 0,(IY - 128)
label_4010: BIT 1,(IY - 128)
label_4011: BIT 2,(IY - 128)
label_4012: BIT 3,(IY - 128)
label_4013: BIT 4,(IY - 128)
label_4014: BIT 5,(IY - 128)
label_4015: BIT 6,(IY - 128)
label_4016: BIT 7,(IY - 128)
label_4017: SET 0,A
label_4018: SET 1,A
label_4019: SET 2,A
label_4020: SET 3,A
label_4021: SET 4,A
label_4022: SET 5,A
label_4023: SET 6,A
label_4024: SET 7,A
label_4025: SET 0,B
label_4026: SET 1,B
label_4027: SET 2,B
label_4028: SET 3,B
label_4029: SET 4,B
label_4030: SET 5,B
label_4031: SET 6,B
label_4032: SET 7,B
label_4033: SET 0,C
label_4034: SET 1,C
label_4035: SET 2,C
label_4036: SET 3,C
label_4037: SET 4,C
label_4038: SET 5,C
label_4039: SET 6,C
label_4040: SET 7,C
label_4041: SET 0,D
label_4042: SET 1,D
label_4043: SET 2,D
label_4044: SET 3,D
label_4045: SET 4,D
label_4046: SET 5,D
label_4047: SET 6,D
label_4048: SET 7,D
label_4049: SET 0,E
label_4050: SET 1,E
label_4051: SET 2,E
label_4052: SET 3,E
label_4053: SET 4,E
label_4054: SET 5,E
label_4055: SET 6,E
label_4056: SET 7,E
label_4057: SET 0,H
label_4058: SET 1,H
label_4059: SET 2,H
label_4060: SET 3,H
label_4061: SET 4,H
label_4062: SET 5,H
label_4063: SET 6,H
label_4064: SET 7,H
label_4065: SET 0,L
label_4066: SET 1,L
label_4067: SET 2,L
label_4068: SET 3,L
label_4069: SET 4,L
label_4070: SET 5,L
label_4071: SET 6,L
label_4072: SET 7,L
label_4073: SET 0,(HL)
label_4074: SET 1,(HL)
label_4075: SET 2,(HL)
label_4076: SET 3,(HL)
label_4077: SET 4,(HL)
label_4078: SET 5,(HL)
label_4079: SET 6,(HL)
label_4080: SET 7,(HL)
label_4081: SET 0,(IX + 127)
label_4082: SET 1,(IX + 127)
label_4083: SET 2,(IX + 127)
label_4084: SET 3,(IX + 127)
label_4085: SET 4,(IX + 127)
label_4086: SET 5,(IX + 127)
label_4087: SET 6,(IX + 127)
label_4088: SET 7,(IX + 127)
label_4089: SET 0,(IY - 128)
label_4090: SET 1,(IY - 128)
label_4091: SET 2,(IY - 128)
label_4092: SET 3,(IY - 128)
label_4093: SET 4,(IY - 128)
label_4094: SET 5,(IY - 128)
label_4095: SET 6,(IY - 128)
label_4096: SET 7,(IY - 128)
label_4097: RES 0,A
label_4098: RES 1,A
label_4099: RES 2,A
label_4100: RES 3,A
label_4101: RES 4,A
label_4102: RES 5,A
label_4103: RES 6,A
label_4104: RES 7,A
label_4105: RES 0,B
label_4106: RES 1,B
label_4107: RES 2,B
label_4108: RES 3,B
label_4109: RES 4,B
label_4110: RES 5,B
label_4111: RES 6,B
label_4112: RES 7,B
label_4113: RES 0,C
label_4114: RES 1,C
label_4115: RES 2,C
label_4116: RES 3,C
label_4117: RES 4,C
label_4118: RES 5,C
label_4119: RES 6,C
label_4120: RES 7,C
label_4121: RES 0,D
label_4122: RES 1,D
label_4123: RES 2,D
label_4124: RES 3,D
label_4125: RES 4,D
label_4126: RES 5,D
label_4127: RES 6,D
label_4128: RES 7,D
label_4129: RES 0,E
label_4130: RES 1,E
label_4131: RES 2,E
label_4132: RES 3,E
label_4133: RES 4,E
label_4134: RES 5,E
label_4135: RES 6,E
label_4136: RES 7,E
label_4137: RES 0,H
label_4138: RES 1,H
label_4139: RES 2,H
label_4140: RES 3,H
label_4141: RES 4,H
label_4142: RES 5,H
label_4143: RES 6,H
label_4144: RES 7,H
label_4145: RES 0,L
label_4146: RES 1,L
label_4147: RES 2,L
label_4148: RES 3,L
label_4149: RES 4,L
label_4150: RES 5,L
label_4151: RES 6,L
label_4152: RES 7,L
label_4153: RES 0,(HL)
label_4154: RES 1,(HL)
label_4155: RES 2,(HL)
label_4156: RES 3,(HL)
label_4157: RES 4,(HL)
label_4158: RES 5,(HL)
label_4159: RES 6,(HL)
label_4160: RES 7,(HL)
label_4161: RES 0,(IX + 127)
label_4162: RES 1,(IX + 127)
label_4163: RES 2,(IX + 127)
label_4164: RES 3,(IX + 127)
label_4165: RES 4,(IX + 127)
label_4166: RES 5,(IX + 127)
label_4167: RES 6,(IX + 127)
label_4168: RES 7,(IX + 127)
label_4169: RES 0,(IY - 128)
label_4170: RES 1,(IY - 128)
label_4171: RES 2,(IY - 128)
label_4172: RES 3,(IY - 128)
label_4173: RES 4,(IY - 128)
label_4174: RES 5,(IY - 128)
label_4175: RES 6,(IY - 128)
label_4176: RES 7,(IY - 128)
label_4177: JP $5678
label_4178: JP NZ,$5678
label_4179: JP Z,$5678
label_4180: JP NC,$5678
label_4181: JP C,$5678
label_4182: JP PO,$5678
label_4183: JP PE,$5678
label_4184: JP P,$5678
label_4185: JP M,$5678
label_4186: JR $ + 2
label_4187: JR NZ,$ + 2
label_4188: JR Z,$ + 2
label_4189: JR NC,$ + 2
label_4190: JR C,$ + 2
label_4191: JP (HL)
label_4192: JP (IX)
label_4193: JP (IY)
label_4194: DJNZ $ + 2
label_4195: CALL $5678
label_4196: CALL NZ,$5678
label_4197: CALL Z,$5678
label_4198: CALL NC,$5678
label_4199: CALL C,$5678
label_4200: CALL PO,$5678
label_4201: CALL PE,$5678
label_4202: CALL P,$5678
label_4203: CALL M,$5678
label_4204: RET
label_4205: RET NZ
label_4206: RET Z
label_4207: RET NC
label_4208: RET C
label_4209: RET PO
label_4210: RET PE
label_4211: RET P
label_4212: RET M
label_4213: RETI
label_4214: RETN
label_4215: RST $00
label_4216: RST $08
label_4217: RST $10
label_4218: RST $18
label_4219: RST $20
label_4220: RST $28
label_4221: RST $30
label_4222: RST $38
label_4223: IN A,($12)
label_4224: IN A,(C)
label_4225: IN B,(C)
label_4226: IN C,(C)
label_4227: IN D,(C)
label_4228: IN E,(C)
label_4229: IN H,(C)
label_4230: IN L,(C)
label_4231: IN F,(C)
label_4232: INI
label_4233: INIR
label_4234: IND
label_4235: INDR
label_4236: OUT ($12),A
label_4237: OUT (C),A
label_4238: OUT (C),B
label_4239: OUT (C),C
label_4240: OUT (C),D
label_4241: OUT (C),E
label_4242: OUT (C),H
label_4243: OUT (C),L
label_4244: OUTI
label_4245: OTIR
label_4246: OUTD
label_4247: OTDR
label_4248: LD A,A
label_4249: LD A,B
label_4250: LD A,C
label_4251: LD A,D
label_4252: LD A,E
label_4253: LD A,H
label_4254: LD A,L
label_4255: LD B,A
label_4256: LD B,B
label_4257: LD B,C
label_4258: LD B,D
label_4259: LD B,E
label_4260: LD B,H
label_4261: LD B,L
label_4262: LD C,A
label_4263: LD C,B
label_4264: LD C,C
label_4265: LD C,D
label_4266: LD C,E
label_4267: LD C,H
label_4268: LD C,L
label_4269: LD D,A
label_4270: LD D,B
label_4271: LD D,C
label_4272: LD D,D
label_4273: LD D,E
label_4274: LD D,H
label_4275: LD D,L
label_4276: LD E,A
label_4277: LD E,B
label_4278: LD E,C
label_4279: LD E,D
label_4280: LD E,E
label_4281: LD E,H
label_4282: LD E,L
label_4283: LD H,A
label_4284: LD H,B
label_4285: LD H,C
label_4286: LD H,D
label_4287: LD H,E
label_4288: LD H,H
label_4289: LD H,L
label_4290: LD L,A
label_4291: LD L,B
label_4292: LD L,C
label_4293: LD L,D
label_4294: LD L,E
label_4295: LD L,H
label_4296: LD L,L
label_4297: LD A,$12
label_4298: LD B,$12
label_4299: LD C,$12
label_4300: LD D,$12
label_4301: LD E,$12
label_4302: LD H,$12
label_4303: LD L,$12
label_4304: LD A,(HL)
label_4305: LD B,(HL)
label_4306: LD C,(HL)
label_4307: LD D,(HL)
label_4308: LD E,(HL)
label_4309: LD H,(HL)
label_4310: LD L,(HL)
label_4311: LD A,(IX + 127)
label_4312: LD B,(IX + 127)
label_4313: LD C,(IX + 127)
label_4314: LD D,(IX + 127)
label_4315: LD E,(IX + 127)
label_4316: LD H,(IX + 127)
label_4317: LD L,(IX + 127)
label_4318: LD A,(IY - 128)
label_4319: LD B,(IY - 128)
label_4320: LD C,(IY - 128)
label_4321: LD D,(IY - 128)
label_4322: LD E,(IY - 128)
label_4323: LD H,(IY - 128)
label_4324: LD L,(IY - 128)
label_4325: LD (HL),A
label_4326: LD (HL),B
label_4327: LD (HL),C
label_4328: LD (HL),D
label_4329: LD (HL),E
label_4330: LD (HL),H
label_4331: LD (HL),L
label_4332: LD (IX + 127),A
label_4333: LD (IX + 127),B
label_4334: LD (IX + 127),C
label_4335: LD (IX + 127),D
label_4336: LD (IX + 127),E
label_4337: LD (IX + 127),H
label_4338: LD (IX + 127),L
label_4339: LD (IY - 128),A
label_4340: LD (IY - 128),B
label_4341: LD (IY - 128),C
label_4342: LD (IY - 128),D
label_4343: LD (IY - 128),E
label_4344: LD (IY - 128),H
label_4345: LD (IY - 128),L
label_4346: LD (HL),$12
label_4347: LD (IX + 127),$12
label_4348: LD (IY - 128),$12
label_4349: LD A,(BC)
label_4350: LD A,(DE)
label_4351: LD A,($5678)
label_4352: LD (BC),A
label_4353: LD (DE),A
label_4354: LD ($5678),A
label_4355: LD A,I
label_4356: LD A,R
label_4357: LD I,A
label_4358: LD R,A
label_4359: LD BC,$5678
label_4360: LD DE,$5678
label_4361: LD HL,$5678
label_4362: LD SP,$5678
label_4363: LD IX,$5678
label_4364: LD IY,$5678
label_4365: LD HL,($5678)
label_4366: LD BC,($5678)
label_4367: LD DE,($5678)
label_4368: LD HL,($5678)
label_4369: LD SP,($5678)
label_4370: LD IX,($5678)
label_4371: LD IY,($5678)
label_4372: LD ($5678),HL
label_4373: LD ($5678),BC
label_4374: LD ($5678),DE
label_4375: LD ($5678),HL
label_4376: LD ($5678),SP
label_4377: LD ($5678),IX
label_4378: LD ($5678),IY
label_4379: LD SP,HL
label_4380: LD SP,IX
label_4381: LD SP,IY
label_4382: PUSH BC
label_4383: PUSH DE
label_4384: PUSH HL
label_4385: PUSH AF
label_4386: PUSH IX
label_4387: PUSH IY
label_4388: POP BC
label_4389: POP DE
label_4390: POP HL
label_4391: POP AF
label_4392: POP IX
label_4393: POP IY
label_4394: EX DE,HL
label_4395: EX AF,AF'
label_4396: EXX
label_4397: EX (SP),HL
label_4398: EX (SP),IX
label_4399: EX (SP),IY
label_4400: LDI
label_4401: LDIR
label_4402: LDD
label_4403: LDDR
label_4404: CPI
label_4405: CPIR
label_4406: CPD
label_4407: CPDR
label_4408: ADD A,A
label_4409: ADD A,B
label_4410: ADD A,C
label_4411: ADD A,D
label_4412: ADD A,E
label_4413: ADD A,H
label_4414: ADD A,L
label_4415: ADD A,$12
label_4416: ADD A,(HL)
label_4417: ADD A,(IX + 127)
label_4418: ADD A,(IY - 128)
label_4419: ADC A,A
label_4420: ADC A,B
label_4421: ADC A,C
label_4422: ADC A,D
label_4423: ADC A,E
label_4424: ADC A,H
label_4425: ADC A,L
label_4426: ADC A,$12
label_4427: ADC A,(HL)
label_4428: ADC A,(IX + 127)
label_4429: ADC A,(IY - 128)
label_4430: SUB A
label_4431: SUB B
label_4432: SUB C
label_4433: SUB D
label_4434: SUB E
label_4435: SUB H
label_4436: SUB L
label_4437: SUB $12
label_4438: SUB (HL)
label_4439: SUB (IX + 127)
label_4440: SUB (IY - 128)
label_4441: SBC A,A
label_4442: SBC A,B
label_4443: SBC A,C
label_4444: SBC A,D
label_4445: SBC A,E
label_4446: SBC A,H
label_4447: SBC A,L
label_4448: SBC A,$12
label_4449: SBC A,(HL)
label_4450: SBC A,(IX + 127)
label_4451: SBC A,(IY - 128)
label_4452: AND A
label_4453: AND B
label_4454: AND C
label_4455: AND D
label_4456: AND E
label_4457: AND H
label_4458: AND L
label_4459: AND $12
label_4460: AND (HL)
label_4461: AND (IX + 127)
label_4462: AND (IY - 128)
label_4463: AND A
label_4464: AND B
label_4465: AND C
label_4466: AND D
label_4467: AND E
label_4468: AND H
label_4469: AND L
label_4470: AND $12
label_4471: AND (HL)
label_4472: AND (IX + 127)
label_4473: AND (IY - 128)
label_4474: OR A
label_4475: OR B
label_4476: OR C
label_4477: OR D
label_4478: OR E
label_4479: OR H
label_4480: OR L
label_4481: OR $12
label_4482: OR (HL)
label_4483: OR (IX + 127)
label_4484: OR (IY - 128)
label_4485: XOR A
label_4486: XOR B
label_4487: XOR C
label_4488: XOR D
label_4489: XOR E
label_4490: XOR H
label_4491: XOR L
label_4492: XOR $12
label_4493: XOR (HL)
label_4494: XOR (IX + 127)
label_4495: XOR (IY - 128)
label_4496: CP A
label_4497: CP B
label_4498: CP C
label_4499: CP D
label_4500: CP E
label_4501: CP H
label_4502: CP L
label_4503: CP $12
label_4504: CP (HL)
label_4505: CP (IX + 127)
label_4506: CP (IY - 128)
label_4507: INC A
label_4508: INC B
label_4509: INC C
label_4510: INC D
label_4511: INC E
label_4512: INC H
label_4513: INC L
label_4514: INC (HL)
label_4515: INC (IX + 127)
label_4516: INC (IY - 128)
label_4517: DEC A
label_4518: DEC B
label_4519: DEC C
label_4520: DEC D
label_4521: DEC E
label_4522: DEC H
label_4523: DEC L
label_4524: DEC (HL)
label_4525: DEC (IX + 127)
label_4526: DEC (IY - 128)
label_4527: DAA
label_4528: CPL
label_4529: NEG
label_4530: CCF
label_4531: SCF
label_4532: NOP
label_4533: HALT
label_4534: DI
label_4535: EI
label_4536: IM 0
label_4537: IM 1
label_4538: IM 2
label_4539: ADD HL,BC
label_4540: ADD HL,DE
label_4541: ADD HL,HL
label_4542: ADD HL,SP
label_4543: ADC HL,BC
label_4544: ADC HL,DE
label_4545: ADC HL,HL
label_4546: ADC HL,SP
label_4547: SBC HL,BC
label_4548: SBC HL,DE
label_4549: SBC HL,HL
label_4550: SBC HL,SP
label_4551: ADD IX,BC
label_4552: ADD IX,DE
label_4553: ADD IX,SP
label_4554: ADD IY,BC
label_4555: ADD IY,DE
label_4556: ADD IY,SP
label_4557: INC BC
label_4558: INC DE
label_4559: INC HL
label_4560: INC SP
label_4561: INC IX
label_4562: INC IY
label_4563: DEC BC
label_4564: DEC DE
label_4565: DEC HL
label_4566: DEC SP
label_4567: DEC IX
label_4568: DEC IY
label_4569: RLCA
label_4570: RLA
label_4571: RRCA
label_4572: RRA
label_4573: RLC A
label_4574: RLC B
label_4575: RLC C
label_4576: RLC D
label_4577: RLC E
label_4578: RLC H
label_4579: RLC L
label_4580: RLC (HL)
label_4581: RLC (IX + 127)
label_4582: RLC (IY - 128)
label_4583: RL A
label_4584: RL B
label_4585: RL C
label_4586: RL D
label_4587: RL E
label_4588: RL H
label_4589: RL L
label_4590: RL (HL)
label_4591: RL (IX + 127)
label_4592: RL (IY - 128)
label_4593: RRC A
label_4594: RRC B
label_4595: RRC C
label_4596: RRC D
label_4597: RRC E
label_4598: RRC H
label_4599: RRC L
label_4600: RRC (HL)
label_4601: RRC (IX + 127)
label_4602: RRC (IY - 128)
label_4603: RR A
label_4604: RR B
label_4605: RR C
label_4606: RR D
label_4607: RR E
label_4608: RR H
label_4609: RR L
label_4610: RR (HL)
label_4611: RR (IX + 127)
label_4612: RR (IY - 128)
label_4613: SLA A
label_4614: SLA B
label_4615: SLA C
label_4616: SLA D
label_4617: SLA E
label_4618: SLA H
label_4619: SLA L
label_4620: SLA (HL)
label_4621: SLA (IX + 127)
label_4622: SLA (IY - 128)
label_4623: SRA A
label_4624: SRA B
label_4625: SRA C
label_4626: SRA D
label_4627: SRA E
label_4628: SRA H
label_4629: SRA L
label_4630: SRA (HL)
label_4631: SRA (IX + 127)
label_4632: SRA (IY - 128)
label_4633: SRL A
label_4634: SRL B
label_4635: SRL C
label_4636: SRL D
label_4637: SRL E
label_4638: SRL H
label_4639: SRL L
label_4640: SRL (HL)
label_4641: SRL (IX + 127)
label_4642: SRL (IY - 128)
label_4643: RLD
label_4644: RRD
label_4645: BIT 0,A
label_4646: BIT 1,A
label_4647: BIT 2,A
label_4648: BIT 3,A
label_4649: BIT 4,A
label_4650: BIT 5,A
label_4651: BIT 6,A
label_4652: BIT 7,A
label_4653: BIT 0,B
label_4654: BIT 1,B
label_4655: BIT 2,B
label_4656: BIT 3,B
label_4657: BIT 4,B
label_4658: BIT 5,B
label_4659: BIT 6,B
label_4660: BIT 7,B
label_4661: BIT 0,C
label_4662: BIT 1,C
label_4663: BIT 2,C
label_4664: BIT 3,C
label_4665: BIT 4,C
label_4666: BIT 5,C
label_4667: BIT 6,C
label_4668: BIT 7,C
label_4669: BIT 0,D
label_4670: BIT 1,D
label_4671: BIT 2,D
label_4672: BIT 3,D
label_4673: BIT 4,D
label_4674: BIT 5,D
label_4675: BIT 6,D
label_4676: BIT 7,D
label_4677: BIT 0,E
label_4678: BIT 1,E
label_4679: BIT 2,E
label_4680: BIT 3,E
label_4681: BIT 4,E
label_4682: BIT 5,E
label_4683: BIT 6,E
label_4684: BIT 7,E
label_4685: BIT 0,H
label_4686: BIT 1,H
label_4687: BIT 2,H
label_4688: BIT 3,H
label_4689: BIT 4,H
label_4690: BIT 5,H
label_4691: BIT 6,H
label_4692: BIT 7,H
label_4693: BIT 0,L
label_4694: BIT 1,L
label_4695: BIT 2,L
label_4696: BIT 3,L
label_4697: BIT 4,L
label_4698: BIT 5,L
label_4699: BIT 6,L
label_4700: BIT 7,L
label_4701: BIT 0,(HL)
label_4702: BIT 1,(HL)
label_4703: BIT 2,(HL)
label_4704: BIT 3,(HL)
label_4705: BIT 4,(HL)
label_4706: BIT 5,(HL)
label_4707: BIT 6,(HL)
label_4708: BIT 7,(HL)
label_4709: BIT 0,(IX + 127)
label_4710: BIT 1,(IX + 127)
label_4711: BIT 2,(IX + 127)
label_4712: BIT 3,(IX + 127)
label_4713: BIT 4,(IX + 127)
label_4714: BIT 5,(IX + 127)
label_4715: BIT 6,(IX + 127)
label_4716: BIT 7,(IX + 127)
label_4717: BIT 0,(IY - 128)
label_4718: BIT 1,(IY - 128)
label_4719: BIT 2,(IY - 128)
label_4720: BIT 3,(IY - 128)
label_4721: BIT 4,(IY - 128)
label_4722: BIT 5,(IY - 128)
label_4723: BIT 6,(IY - 128)
label_4724: BIT 7,(IY - 128)
label_4725: SET 0,A
label_4726: SET 1,A
label_4727: SET 2,A
label_4728: SET 3,A
label_4729: SET 4,A
label_4730: SET 5,A
label_4731: SET 6,A
label_4732: SET 7,A
label_4733: SET 0,B
label_4734: SET 1,B
label_4735: SET 2,B
label_4736: SET 3,B
label_4737: SET 4,B
label_4738: SET 5,B
label_4739: SET 6,B
label_4740: SET 7,B
label_4741: SET 0,C
label_4742: SET 1,C
label_4743: SET 2,C
label_4744: SET 3,C
label_4745: SET 4,C
label_4746: SET 5,C
label_4747: SET 6,C
label_4748: SET 7,C
label_4749: SET 0,D
label_4750: SET 1,D
label_4751: SET 2,D
label_4752: SET 3,D
label_4753: SET 4,D
label_4754: SET 5,D
label_4755: SET 6,D
label_4756: SET 7,D
label_4757: SET 0,E
label_4758: SET 1,E
label_4759: SET 2,E
label_4760: SET 3,E
label_4761: SET 4,E
label_4762: SET 5,E
label_4763: SET 6,E
label_4764: SET 7,E
label_4765: SET 0,H
label_4766: SET 1,H
label_4767: SET 2,H
label_4768: SET 3,H
label_4769: SET 4,H
label_4770: SET 5,H
label_4771: SET 6,H
label_4772: SET 7,H
label_4773: SET 0,L
label_4774: SET 1,L
label_4775: SET 2,L
label_4776: SET 3,L
label_4777: SET 4,L
label_4778: SET 5,L
label_4779: SET 6,L
label_4780: SET 7,L
label_4781: SET 0,(HL)
label_4782: SET 1,(HL)
label_4783: SET 2,(HL)
label_4784: SET 3,(HL)
label_4785: SET 4,(HL)
label_4786: SET 5,(HL)
label_4787: SET 6,(HL)
label_4788: SET 7,(HL)
label_4789: SET 0,(IX + 127)
label_4790: SET 1,(IX + 127)
label_4791: SET 2,(IX + 127)
label_4792: SET 3,(IX + 127)
label_4793: SET 4,(IX + 127)
label_4794: SET 5,(IX + 127)
label_4795: SET 6,(IX + 127)
label_4796: SET 7,(IX + 127)
label_4797: SET 0,(IY - 128)
label_4798: SET 1,(IY - 128)
label_4799: SET 2,(IY - 128)
label_4800: SET 3,(IY - 128)
label_4801: SET 4,(IY - 128)
label_4802: SET 5,(IY - 128)
label_4803: SET 6,(IY - 128)
label_4804: SET 7,(IY - 128)
label_4805: RES 0,A
label_4806: RES 1,A
label_4807: RES 2,A
label_4808: RES 3,A
label_4809: RES 4,A
label_4810: RES 5,A
label_4811: RES 6,A
label_4812: RES 7,A
label_4813: RES 0,B
label_4814: RES 1,B
label_4815: RES 2,B
label_4816: RES 3,B
label_4817: RES 4,B
label_4818: RES 5,B
label_4819: RES 6,B
label_4820: RES 7,B
label_4821: RES 0,C
label_4822: RES 1,C
label_4823: RES 2,C
label_4824: RES 3,C
label_4825: RES 4,C
label_4826: RES 5,C
label_4827: RES 6,C
label_4828: RES 7,C
label_4829: RES 0,D
label_4830: RES 1,D
label_4831: RES 2,D
label_4832: RES 3,D
label_4833: RES 4,D
label_4834: RES 5,D
label_4835: RES 6,D
label_4836: RES 7,D
label_4837: RES 0,E
label_4838: RES 1,E
label_4839: RES 2,E
label_4840: RES 3,E
label_4841: RES 4,E
label_4842: RES 5,E
label_4843: RES 6,E
label_4844: RES 7,E
label_4845: RES 0,H
label_4846: RES 1,H
label_4847: RES 2,H
label_4848: RES 3,H
label_4849: RES 4,H
label_4850: RES 5,H
label_4851: RES 6,H
label_4852: RES 7,H
label_4853: RES 0,L
label_4854: RES 1,L
label_4855: RES 2,L
label_4856: RES 3,L
label_4857: RES 4,L
label_4858: RES 5,L
label_4859: RES 6,L
label_4860: RES 7,L
label_4861: RES 0,(HL)
label_4862: RES 1,(HL)
label_4863: RES 2,(HL)
label_4864: RES 3,(HL)
label_4865: RES 4,(HL)
label_4866: RES 5,(HL)
label_4867: RES 6,(HL)
label_4868: RES 7,(HL)
label_4869: RES 0,(IX + 127)
label_4870: RES 1,(IX + 127)
label_4871: RES 2,(IX + 127)
label_4872: RES 3,(IX + 127)
label_4873: RES 4,(IX + 127)
label_4874: RES 5,(IX + 127)
label_4875: RES 6,(IX + 127)
label_4876: RES 7,(IX + 127)
label_4877: RES 0,(IY - 128)
label_4878: RES 1,(IY - 128)
label_4879: RES 2,(IY - 128)
label_4880: RES 3,(IY - 128)
label_4881: RES 4,(IY - 128)
label_4882: RES 5,(IY - 128)
label_4883: RES 6,(IY - 128)
label_4884: RES 7,(IY - 128)
label_4885: JP $5678
label_4886: JP NZ,$5678
label_4887: JP Z,$5678
label_4888: JP NC,$5678
label_4889: JP C,$5678
label_4890: JP PO,$5678
label_4891: JP PE,$5678
label_4892: JP P,$5678
label_4893: JP M,$5678
label_4894: JR $ + 2
label_4895: JR NZ,$ + 2
label_4896: JR Z,$ + 2
label_4897: JR NC,$ + 2
label_4898: JR C,$ + 2
label_4899: JP (HL)
label_4900: JP (IX)
label_4901: JP (IY)
label_4902: DJNZ $ + 2
label_4903: CALL $5678
label_4904: CALL NZ,$5678
label_4905: CALL Z,$5678
label_4906: CALL NC,$5678
label_4907: CALL C,$5678
label_4908: CALL PO,$5678
label_4909: CALL PE,$5678
label_4910: CALL P,$5678
label_4911: CALL M,$5678
label_4912: RET
label_4913: RET NZ
label_4914: RET Z
label_4915: RET NC
label_4916: RET C
label_4917: RET PO
label_4918: RET PE
label_4919: RET P
label_4920: RET M
label_4921: RETI
label_4922: RETN
label_4923: RST $00
label_4924: RST $08
label_4925: RST $10
label_4926: RST $18
label_4927: RST $20
label_4928: RST $28
label_4929: RST $30
label_4930: RST $38
label_4931: IN A,($12)
label_4932: IN A,(C)
label_4933: IN B,(C)
label_4934: IN C,(C)
label_4935: IN D,(C)
label_4936: IN E,(C)
label_4937: IN H,(C)
label_4938: IN L,(C)
label_4939: IN F,(C)
label_4940: INI
label_4941: INIR
label_4942: IND
label_4943: INDR
label_4944: OUT ($12),A
label_4945: OUT (C),A
label_4946: OUT (C),B
label_4947: OUT (C),C
label_4948: OUT (C),D
label_4949: OUT (C),E
label_4950: OUT (C),H
label_4951: OUT (C),L
label_4952: OUTI
label_4953: OTIR
label_4954: OUTD
label_4955: OTDR
label_4956: LD A,A
label_4957: LD A,B
label_4958: LD A,C
label_4959: LD A,D
label_4960: LD A,E
label_4961: LD A,H
label_4962: LD A,L
label_4963: LD B,A
label_4964: LD B,B
label_4965: LD B,C
label_4966: LD B,D
label_4967: LD B,E
label_4968: LD B,H
label_4969: LD B,L
label_4970: LD C,A
label_4971: LD C,B
label_4972: LD C,C
label_4973: LD C,D
label_4974: LD C,E
label_4975: LD C,H
label_4976: LD C,L
label_4977: LD D,A
label_4978: LD D,B
label_4979: LD D,C
label_4980: LD D,D
label_4981: LD D,E
label_4982: LD D,H
label_4983: LD D,L
label_4984: LD E,A
label_4985: LD E,B
label_4986: LD E,C
label_4987: LD E,D
label_4988: LD E,E
label_4989: LD E,H
label_4990: LD E,L
label_4991: LD H,A
label_4992: LD H,B
label_4993: LD H,C
label_4994: LD H,D
label_4995: LD H,E
label_4996: LD H,H
label_4997: LD H,L
label_4998: LD L,A
label_4999: LD L,B
label_5000: LD L,C
label_5001: LD L,D
label_5002: LD L,E
label_5003: LD L,H
label_5004: LD L,L
label_5005: LD A,$12
label_5006: LD B,$12
label_5007: LD C,$12
label_5008: LD D,$12
label_5009: LD E,$12
label_5010: LD H,$12
label_5011: LD L,$12
label_5012: LD A,(HL)
label_5013: LD B,(HL)
label_5014: LD C,(HL)
label_5015: LD D,(HL)
label_5016: LD E,(HL)
label_5017: LD H,(HL)
label_5018: LD L,(HL)
label_5019: LD A,(IX + 127)
label_5020: LD B,(IX + 127)
label_5021: LD C,(IX + 127)
label_5022: LD D,(IX + 127)
label_5023: LD E,(IX + 127)
label_5024: LD H,(IX + 127)
label_5025: LD L,(IX + 127)
label_5026: LD A,(IY - 128)
label_5027: LD B,(IY - 128)
label_5028: LD C,(IY - 128)
label_5029: LD D,(IY - 128)
label_5030: LD E,(IY - 128)
label_5031: LD H,(IY - 128)
label_5032: LD L,(IY - 128)
label_5033: LD (HL),A
label_5034: LD (HL),B
label_5035: LD (HL),C
label_5036: LD (HL),D
label_5037: LD (HL),E
label_5038: LD (HL),H
label_5039: LD (HL),L
label_5040: LD (IX + 127),A
label_5041: LD (IX + 127),B
label_5042: LD (IX + 127),C
label_5043: LD (IX + 127),D
label_5044: LD (IX + 127),E
label_5045: LD (IX + 127),H
label_5046: LD (IX + 127),L
label_5047: LD (IY - 128),A
label_5048: LD (IY - 128),B
label_5049: LD (IY - 128),C
label_5050: LD (IY - 128),D
label_5051: LD (IY - 128),E
label_5052: LD (IY - 128),H
label_5053: LD (IY - 128),L
label_5054: LD (HL),$12
label_5055: LD (IX + 127),$12
label_5056: LD (IY - 128),$12
label_5057: LD A,(BC)
label_5058: LD A,(DE)
label_5059: LD A,($5678)
label_5060: LD (BC),A
label_5061: LD (DE),A
label_5062: LD ($5678),A
label_5063: LD A,I
label_5064: LD A,R
label_5065: LD I,A
label_5066: LD R,A
label_5067: LD BC,$5678
label_5068: LD DE,$5678
label_5069: LD HL,$5678
label_5070: LD SP,$5678
label_5071: LD IX,$5678
label_5072: LD IY,$5678
label_5073: LD HL,($5678)
label_5074: LD BC,($5678)
label_5075: LD DE,($5678)
label_5076: LD HL,($5678)
label_5077: LD SP,($5678)
label_5078: LD IX,($5678)
label_5079: LD IY,($5678)
label_5080: LD ($5678),HL
label_5081: LD ($5678),BC
label_5082: LD ($5678),DE
label_5083: LD ($5678),HL
label_5084: LD ($5678),SP
label_5085: LD ($5678),IX
label_5086: LD ($5678),IY
label_5087: LD SP,HL
label_5088: LD SP,IX
label_5089: LD SP,IY
label_5090: PUSH BC
label_5091: PUSH DE
label_5092: PUSH HL
label_5093: PUSH AF
label_5094: PUSH IX
label_5095: PUSH IY
label_5096: POP BC
label_5097: POP DE
label_5098: POP HL
label_5099: POP AF
label_5100: POP IX
label_5101: POP IY
label_5102: EX DE,HL
label_5103: EX AF,AF'
label_5104: EXX
label_5105: EX (SP),HL
label_5106: EX (SP),IX
label_5107: EX (SP),IY
label_5108: LDI
label_5109: LDIR
label_5110: LDD
label_5111: LDDR
label_5112: CPI
label_5113: CPIR
label_5114: CPD
label_5115: CPDR
label_5116: ADD A,A
label_5117: ADD A,B
label_5118: ADD A,C
label_5119: ADD A,D
label_5120: ADD A,E
label_5121: ADD A,H
label_5122: ADD A,L
label_5123: ADD A,$12
label_5124: ADD A,(HL)
label_5125: ADD A,(IX + 127)
label_5126: ADD A,(IY - 128)
label_5127: ADC A,A
label_5128: ADC A,B
label_5129: ADC A,C
label_5130: ADC A,D
label_5131: ADC A,E
label_5132: ADC A,H
label_5133: ADC A,L
label_5134: ADC A,$12
label_5135: ADC A,(HL)
label_5136: ADC A,(IX + 127)
label_5137: ADC A,(IY - 128)
label_5138: SUB A
label_5139: SUB B
label_5140: SUB C
label_5141: SUB D
label_5142: SUB E
label_5143: SUB H
label_5144: SUB L
label_5145: SUB $12
label_5146: SUB (HL)
label_5147: SUB (IX + 127)
label_5148: SUB (IY - 128)
label_5149: SBC A,A
label_5150: SBC A,B
label_5151: SBC A,C
label_5152: SBC A,D
label_5153: SBC A,E
label_5154: SBC A,H
label_5155: SBC A,L
label_5156: SBC A,$12
label_5157: SBC A,(HL)
label_5158: SBC A,(IX + 127)
label_5159: SBC A,(IY - 128)
label_5160: AND A
label_5161: AND B
label_5162: AND C
label_5163: AND D
label_5164: AND E
label_5165: AND H
label_5166: AND L
label_5167: AND $12
label_5168: AND (HL)
label_5169: AND (IX + 127)
label_5170: AND (IY - 128)
label_5171: AND A
label_5172: AND B
label_5173: AND C
label_5174: AND D
label_5175: AND E
label_5176: AND H
label_5177: AND L
label_5178: AND $12
label_5179: AND (HL)
label_5180: AND (IX + 127)
label_5181: AND (IY - 128)
label_5182: OR A
label_5183: OR B
label_5184: OR C
label_5185: OR D
label_5186: OR E
label_5187: OR H
label_5188: OR L
label_5189: OR $12
label_5190: OR (HL)
label_5191: OR (IX + 127)
label_5192: OR (IY - 128)
label_5193: XOR A
label_5194: XOR B
label_5195: XOR C
label_5196: XOR D
label_5197: XOR E
label_5198: XOR H
label_5199: XOR L
label_5200: XOR $12
label_5201: XOR (HL)
label_5202: XOR (IX + 127)
label_5203: XOR (IY - 128)
label_5204: CP A
label_5205: CP B
label_5206: CP C
label_5207: CP D
label_5208: CP E
label_5209: CP H
label_5210: CP L
label_5211: CP $12
label_5212: CP (HL)
label_5213: CP (IX + 127)
label_5214: CP (IY - 128)
label_5215: INC A
label_5216: INC B
label_5217: INC C
label_5218: INC D
label_5219: INC E
label_5220: INC H
label_5221: INC L
label_5222: INC (HL)
label_5223: INC (IX + 127)
label_5224: INC (IY - 128)
label_5225: DEC A
label_5226: DEC B
label_5227: DEC C
label_5228: DEC D
label_5229: DEC E
label_5230: DEC H
label_5231: DEC L
label_5232: DEC (HL)
label_5233: DEC (IX + 127)
label_5234: DEC (IY - 128)
label_5235: DAA
label_5236: CPL
label_5237: NEG
label_5238: CCF
label_5239: SCF
label_5240: NOP
label_5241: HALT
label_5242: DI
label_5243: EI
label_5244: IM 0
label_5245: IM 1
label_5246: IM 2
label_5247: ADD HL,BC
label_5248: ADD HL,DE
label_5249: ADD HL,HL
label_5250: ADD HL,SP
label_5251: ADC HL,BC
label_5252: ADC HL,DE
label_5253: ADC HL,HL
label_5254: ADC HL,SP
label_5255: SBC HL,BC
label_5256: SBC HL,DE
label_5257: SBC HL,HL
label_5258: SBC HL,SP
label_5259: ADD IX,BC
label_5260: ADD IX,DE
label_5261: ADD IX,SP
label_5262: ADD IY,BC
label_5263: ADD IY,DE
label_5264: ADD IY,SP
label_5265: INC BC
label_5266: INC DE
label_5267: INC HL
label_5268: INC SP
label_5269: INC IX
label_5270: INC IY
label_5271: DEC BC
label_5272: DEC DE
label_5273: DEC HL
label_5274: DEC SP
label_5275: DEC IX
label_5276: DEC IY
label_5277: RLCA
label_5278: RLA
label_5279: RRCA
label_5280: RRA
label_5281: RLC A
label_5282: RLC B
label_5283: RLC C
label_5284: RLC D
label_5285: RLC E
label_5286: RLC H
label_5287: RLC L
label_5288: RLC (HL)
label_5289: RLC (IX + 127)
label_5290: RLC (IY - 128)
label_5291: RL A
label_5292: RL B
label_5293: RL C
label_5294: RL D
label_5295: RL E
label_5296: RL H
label_5297: RL L
label_5298: RL (HL)
label_5299: RL (IX + 127)
label_5300: RL (IY - 128)
label_5301: RRC A
label_5302: RRC B
label_5303: RRC C
label_5304: RRC D
label_5305: RRC E
label_5306: RRC H
label_5307: RRC L
label_5308: RRC (HL)
label_5309: RRC (IX + 127)
label_5310: RRC (IY - 128)
label_5311: RR A
label_5312: RR B
label_5313: RR C
label_5314: RR D
label_5315: RR E
label_5316: RR H
label_5317: RR L
label_5318: RR (HL)
label_5319: RR (IX + 127)
label_5320: RR (IY - 128)
label_5321: SLA A
label_5322: SLA B
label_5323: SLA C
label_5324: SLA D
label_5325: SLA E
label_5326: SLA H
label_5327: SLA L
label_5328: SLA (HL)
label_5329: SLA (IX + 127)
label_5330: SLA (IY - 128)
label_5331: SRA A
label_5332: SRA B
label_5333: SRA C
label_5334: SRA D
label_5335: SRA E
label_5336: SRA H
label_5337: SRA L
label_5338: SRA (HL)
label_5339: SRA (IX + 127)
label_5340: SRA (IY - 128)
label_5341: SRL A
label_5342: SRL B
label_5343: SRL C
label_5344: SRL D
label_5345: SRL E
label_5346: SRL H
label_5347: SRL L
label_5348: SRL (HL)
label_5349: SRL (IX + 127)
label_5350: SRL (IY - 128)
label_5351: RLD
label_5352: RRD
label_5353: BIT 0,A
label_5354: BIT 1,A
label_5355: BIT 2,A
label_5356: BIT 3,A
label_5357: BIT 4,A
label_5358: BIT 5,A
label_5359: BIT 6,A
label_5360: BIT 7,A
label_5361: BIT 0,B
label_5362: BIT 1,B
label_5363: BIT 2,B
label_5364: BIT 3,B
label_5365: BIT 4,B
label_5366: BIT 5,B
label_5367: BIT 6,B
label_5368: BIT 7,B
label_5369: BIT 0,C
label_5370: BIT 1,C
label_5371: BIT 2,C
label_5372: BIT 3,C
label_5373: BIT 4,C
label_5374: BIT 5,C
label_5375: BIT 6,C
label_5376: BIT 7,C
label_5377: BIT 0,D
label_5378: BIT 1,D
label_5379: BIT 2,D
label_5380: BIT 3,D
label_5381: BIT 4,D
label_5382: BIT 5,D
label_5383: BIT 6,D
label_5384: BIT 7,D
label_5385: BIT 0,E
label_5386: BIT 1,E
label_5387: BIT 2,E
label_5388: BIT 3,E
label_5389: BIT 4,E
label_5390: BIT 5,E
label_5391: BIT 6,E
label_5392: BIT 7,E
label_5393: BIT 0,H
label_5394: BIT 1,H
label_5395: BIT 2,H
label_5396: BIT 3,H
label_5397: BIT 4,H
label_5398: BIT 5,H
label_5399: BIT 6,H
label_5400: BIT 7,H
label_5401: BIT 0,L
label_5402: BIT 1,L
label_5403: BIT 2,L
label_5404: BIT 3,L
label_5405: BIT 4,L
label_5406: BIT 5,L
label_5407: BIT 6,L
label_5408: BIT 7,L
label_5409: BIT 0,(HL)
label_5410: BIT 1,(HL)
label_5411: BIT 2,(HL)
label_5412: BIT 3,(HL)
label_5413: BIT 4,(HL)
label_5414: BIT 5,(HL)
label_5415: BIT 6,(HL)
label_5416: BIT 7,(HL)
label_5417: BIT 0,(IX + 127)
label_5418: BIT 1,(IX + 127)
label_5419: BIT 2,(IX + 127)
label_5420: BIT 3,(IX + 127)
label_5421: BIT 4,(IX + 127)
label_5422: BIT 5,(IX + 127)
label_5423: BIT 6,(IX + 127)
label_5424: BIT 7,(IX + 127)
label_5425: BIT 0,(IY - 128)
label_5426: BIT 1,(IY - 128)
label_5427: BIT 2,(IY - 128)
label_5428: BIT 3,(IY - 128)
label_5429: BIT 4,(IY - 128)
label_5430: BIT 5,(IY - 128)
label_5431: BIT 6,(IY - 128)
label_5432: BIT 7,(IY - 128)
label_5433: SET 0,A
label_5434: SET 1,A
label_5435: SET 2,A
label_5436: SET 3,A
label_5437: SET 4,A
label_5438: SET 5,A
label_5439: SET 6,A
label_5440: SET 7,A
label_5441: SET 0,B
label_5442: SET 1,B
label_5443: SET 2,B
label_5444: SET 3,B
label_5445: SET 4,B
label_5446: SET 5,B
label_5447: SET 6,B
label_5448: SET 7,B
label_5449: SET 0,C
label_5450: SET 1,C
label_5451: SET 2,C
label_5452: SET 3,C
label_5453: SET 4,C
label_5454: SET 5,C
label_5455: SET 6,C
label_5456: SET 7,C
label_5457: SET 0,D
label_5458: SET 1,D
label_5459: SET 2,D
label_5460: SET 3,D
label_5461: SET 4,D
label_5462: SET 5,D
label_5463: SET 6,D
label_5464: SET 7,D
label_5465: SET 0,E
label_5466: SET 1,E
label_5467: SET 2,E
label_5468: SET 3,E
label_5469: SET 4,E
label_5470: SET 5,E
label_5471: SET 6,E
label_5472: SET 7,E
label_5473: SET 0,H
label_5474: SET 1,H
label_5475: SET 2,H
label_5476: SET 3,H
label_5477: SET 4,H
label_5478: SET 5,H
label_5479: SET 6,H
label_5480: SET 7,H
label_5481: SET 0,L
label_5482: SET 1,L
label_5483: SET 2,L
label_5484: SET 3,L
label_5485: SET 4,L
label_5486: SET 5,L
label_5487: SET 6,L
label_5488: SET 7,L
label_5489: SET 0,(HL)
label_5490: SET 1,(HL)
label_5491: SET 2,(HL)
label_5492: SET 3,(HL)
label_5493: SET 4,(HL)
label_5494: SET 5,(HL)
label_5495: SET 6,(HL)
label_5496: SET 7,(HL)
label_5497: SET 0,(IX + 127)
label_5498: SET 1,(IX + 127)
label_5499: SET 2,(IX + 127)
label_5500: SET 3,(IX + 127)
label_5501: SET 4,(IX + 127)
label_5502: SET 5,(IX + 127)
label_5503: SET 6,(IX + 127)
label_5504: SET 7,(IX + 127)
label_5505: SET 0,(IY - 128)
label_5506: SET 1,(IY - 128)
label_5507: SET 2,(IY - 128)
label_5508: SET 3,(IY - 128)
label_5509: SET 4,(IY - 128)
label_5510: SET 5,(IY - 128)
label_5511: SET 6,(IY - 128)
label_5512: SET 7,(IY - 128)
label_5513: RES 0,A
label_5514: RES 1,A
label_5515: RES 2,A
label_5516: RES 3,A
label_5517: RES 4,A
label_5518: RES 5,A
label_5519: RES 6,A
label_5520: RES 7,A
label_5521: RES 0,B
label_5522: RES 1,B
label_5523: RES 2,B
label_5524: RES 3,B
label_5525: RES 4,B
label_5526: RES 5,B
label_5527: RES 6,B
label_5528: RES 7,B
label_5529: RES 0,C
label_5530: RES 1,C
label_5531: RES 2,C
label_5532: RES 3,C
label_5533: RES 4,C
label_5534: RES 5,C
label_5535: RES 6,C
label_5536: RES 7,C
label_5537: RES 0,D
label_5538: RES 1,D
label_5539: RES 2,D
label_5540: RES 3,D
label_5541: RES 4,D
label_5542: RES 5,D
label_5543: RES 6,D
label_5544: RES 7,D
label_5545: RES 0,E
label_5546: RES 1,E
label_5547: RES 2,E
label_5548: RES 3,E
label_5549: RES 4,E
label_5550: RES 5,E
label_5551: RES 6,E
label_5552: RES 7,E
label_5553: RES 0,H
label_5554: RES 1,H
label_5555: RES 2,H
label_5556: RES 3,H
label_5557: RES 4,H
label_5558: RES 5,H
label_5559: RES 6,H
label_5560: RES 7,H
label_5561: RES 0,L
label_5562: RES 1,L
label_5563: RES 2,L
label_5564: RES 3,L
label_5565: RES 4,L
label_5566: RES 5,L
label_5567: RES 6,L
label_5568: RES 7,L
label_5569: RES 0,(HL)
label_5570: RES 1,(HL)
label_5571: RES 2,(HL)
label_5572: RES 3,(HL)
label_5573: RES 4,(HL)
label_5574: RES 5,(HL)
label_5575: RES 6,(HL)
label_5576: RES 7,(HL)
label_5577: RES 0,(IX + 127)
label_5578: RES 1,(IX + 127)
label_5579: RES 2,(IX + 127)
label_5580: RES 3,(IX + 127)
label_5581: RES 4,(IX + 127)
label_5582: RES 5,(IX + 127)
label_5583: RES 6,(IX + 127)
label_5584: RES 7,(IX + 127)
label_5585: RES 0,(IY - 128)
label_5586: RES 1,(IY - 128)
label_5587: RES 2,(IY - 128)
label_5588: RES 3,(IY - 128)
label_5589: RES 4,(IY - 128)
label_5590: RES 5,(IY - 128)
label_5591: RES 6,(IY - 128)
label_5592: RES 7,(IY - 128)
label_5593: JP $5678
label_5594: JP NZ,$5678
label_5595: JP Z,$5678
label_5596: JP NC,$5678
label_5597: JP C,$5678
label_5598: JP PO,$5678
label_5599: JP PE,$5678
label_5600: JP P,$5678
label_5601: JP M,$5678
label_5602: JR $ + 2
label_5603: JR NZ,$ + 2
label_5604: JR Z,$ + 2
label_5605: JR NC,$ + 2
label_5606: JR C,$ + 2
label_5607: JP (HL)
label_5608: JP (IX)
label_5609: JP (IY)
label_5610: DJNZ $ + 2
label_5611: CALL $5678
label_5612: CALL NZ,$5678
label_5613: CALL Z,$5678
label_5614: CALL NC,$5678
label_5615: CALL C,$5678
label_5616: CALL PO,$5678
label_5617: CALL PE,$5678
label_5618: CALL P,$5678
label_5619: CALL M,$5678
label_5620: RET
label_5621: RET NZ
label_5622: RET Z
label_5623: RET NC
label_5624: RET C
label_5625: RET PO
label_5626: RET PE
label_5627: RET P
label_5628: RET M
label_5629: RETI
label_5630: RETN
label_5631: RST $00
label_5632: RST $08
label_5633: RST $10
label_5634: RST $18
label_5635: RST $20
label_5636: RST $28
label_5637: RST $30
label_5638: RST $38
label_5639: IN A,($12)
label_5640: IN A,(C)
label_5641: IN B,(C)
label_5642: IN C,(C)
label_5643: IN D,(C)
label_5644: IN E,(C)
label_5645: IN H,(C)
label_5646: IN L,(C)
label_5647: IN F,(C)
label_5648: INI
label_5649: INIR
label_5650: IND
label_5651: INDR
label_5652: OUT ($12),A
label_5653: OUT (C),A
label_5654: OUT (C),B
label_5655: OUT (C),C
label_5656: OUT (C),D
label_5657: OUT (C),E
label_5658: OUT (C),H
label_5659: OUT (C),L
label_5660: OUTI
label_5661: OTIR
label_5662: OUTD
label_5663: OTDR
label_5664: LD A,A
label_5665: LD A,B
label_5666: LD A,C
label_5667: LD A,D
label_5668: LD A,E
label_5669: LD A,H
label_5670: LD A,L
label_5671: LD B,A
label_5672: LD B,B
label_5673: LD B,C
label_5674: LD B,D
label_5675: LD B,E
label_5676: LD B,H
label_5677: LD B,L
label_5678: LD C,A
label_5679: LD C,B
label_5680: LD C,C
label_5681: LD C,D
label_5682: LD C,E
label_5683: LD C,H
label_5684: LD C,L
label_5685: LD D,A
label_5686: LD D,B
label_5687: LD D,C
label_5688: LD D,D
label_5689: LD D,E
label_5690: LD D,H
label_5691: LD D,L
label_5692: LD E,A
label_5693: LD E,B
label_5694: LD E,C
label_5695: LD E,D
label_5696: LD E,E
label_5697: LD E,H
label_5698: LD E,L
label_5699: LD H,A
label_5700: LD H,B
label_5701: LD H,C
label_5702: LD H,D
label_5703: LD H,E
label_5704: LD H,H
label_5705: LD H,L
label_5706: LD L,A
label_5707: LD L,B
label_5708: LD L,C
label_5709: LD L,D
label_5710: LD L,E
label_5711: LD L,H
label_5712: LD L,L
label_5713: LD A,$12
label_5714: LD B,$12
label_5715: LD C,$12
label_5716: LD D,$12
label_5717: LD E,$12
label_5718: LD H,$12
label_5719: LD L,$12
label_5720: LD A,(HL)
label_5721: LD B,(HL)
label_5722: LD C,(HL)
label_5723: LD D,(HL)
label_5724: LD E,(HL)
label_5725: LD H,(HL)
label_5726: LD L,(HL)
label_5727: LD A,(IX + 127)
label_5728: LD B,(IX + 127)
label_5729: LD C,(IX + 127)
label_5730: LD D,(IX + 127)
label_5731: LD E,(IX + 127)
label_5732: LD H,(IX + 127)
label_5733: LD L,(IX + 127)
label_5734: LD A,(IY - 128)
label_5735: LD B,(IY - 128)
label_5736: LD C,(IY - 128)
label_5737: LD D,(IY - 128)
label_5738: LD E,(IY - 128)
label_5739: LD H,(IY - 128)
label_5740: LD L,(IY - 128)
label_5741: LD (HL),A
label_5742: LD (HL),B
label_5743: LD (HL),C
label_5744: LD (HL),D
label_5745: LD (HL),E
label_5746: LD (HL),H
label_5747: LD (HL),L
label_5748: LD (IX + 127),A
label_5749: LD (IX + 127),B
label_5750: LD (IX + 127),C
label_5751: LD (IX + 127),D
label_5752: LD (IX + 127),E
label_5753: LD (IX + 127),H
label_5754: LD (IX + 127),L
label_5755: LD (IY - 128),A
label_5756: LD (IY - 128),B
label_5757: LD (IY - 128),C
label_5758: LD (IY - 128),D
label_5759: LD (IY - 128),E
label_5760: LD (IY - 128),H
label_5761: LD (IY - 128),L
label_5762: LD (HL),$12
label_5763: LD (IX + 127),$12
label_5764: LD (IY - 128),$12
label_5765: LD A,(BC)
label_5766: LD A,(DE)
label_5767: LD A,($5678)
label_5768: LD (BC),A
label_5769: LD (DE),A
label_5770: LD ($5678),A
label_5771: LD A,I
label_5772: LD A,R
label_5773: LD I,A
label_5774: LD R,A
label_5775: LD BC,$5678
label_5776: LD DE,$5678
label_5777: LD HL,$5678
label_5778: LD SP,$5678
label_5779: LD IX,$5678
label_5780: LD IY,$5678
label_5781: LD HL,($5678)
label_5782: LD BC,($5678)
label_5783: LD DE,($5678)
label_5784: LD HL,($5678)
label_5785: LD SP,($5678)
label_5786: LD IX,($5678)
label_5787: LD IY,($5678)
label_5788: LD ($5678),HL
label_5789: LD ($5678),BC
label_5790: LD ($5678),DE
label_5791: LD ($5678),HL
label_5792: LD ($5678),SP
label_5793: LD ($5678),IX
label_5794: LD ($5678),IY
label_5795: LD SP,HL
label_5796: LD SP,IX
label_5797: LD SP,IY
label_5798: PUSH BC
label_5799: PUSH DE
label_5800: PUSH HL
label_5801: PUSH AF
label_5802: PUSH IX
label_5803: PUSH IY
label_5804: POP BC
label_5805: POP DE
label_5806: POP HL
label_5807: POP AF
label_5808: POP IX
label_5809: POP IY
label_5810: EX DE,HL
label_5811: EX AF,AF'
label_5812: EXX
label_5813: EX (SP),HL
label_5814: EX (SP),IX
label_5815: EX (SP),IY
label_5816: LDI
label_5817: LDIR
label_5818: LDD
label_5819: LDDR
label_5820: CPI
label_5821: CPIR
label_5822: CPD
label_5823: CPDR
label_5824: ADD A,A
label_5825: ADD A,B
label_5826: ADD A,C
label_5827: ADD A,D
label_5828: ADD A,E
label_5829: ADD A,H
label_5830: ADD A,L
label_5831: ADD A,$12
label_5832: ADD A,(HL)
label_5833: ADD A,(IX + 127)
label_5834: ADD A,(IY - 128)
label_5835: ADC A,A
label_5836: ADC A,B
label_5837: ADC A,C
label_5838: ADC A,D
label_5839: ADC A,E
label_5840: ADC A,H
label_5841: ADC A,L
label_5842: ADC A,$12
label_5843: ADC A,(HL)
label_5844: ADC A,(IX + 127)
label_5845: ADC A,(IY - 128)
label_5846: SUB A
label_5847: SUB B
label_5848: SUB C
label_5849: SUB D
label_5850: SUB E
label_5851: SUB H
label_5852: SUB L
label_5853: SUB $12
label_5854: SUB (HL)
label_5855: SUB (IX + 127)
label_5856: SUB (IY - 128)
label_5857: SBC A,A
label_5858: SBC A,B
label_5859: SBC A,C
label_5860: SBC A,D
label_5861: SBC A,E
label_5862: SBC A,H
label_5863: SBC A,L
label_5864: SBC A,$12
label_5865: SBC A,(HL)
label_5866: SBC A,(IX + 127)
label_5867: SBC A,(IY - 128)
label_5868: AND A
label_5869: AND B
label_5870: AND C
label_5871: AND D
label_5872: AND E
label_5873: AND H
label_5874: AND L
label_5875: AND $12
label_5876: AND (HL)
label_5877: AND (IX + 127)
label_5878: AND (IY - 128)
label_5879: AND A
label_5880: AND B
label_5881: AND C
label_5882: AND D
label_5883: AND E
label_5884: AND H
label_5885: AND L
label_5886: AND $12
label_5887: AND (HL)
label_5888: AND (IX + 127)
label_5889: AND (IY - 128)
label_5890: OR A
label_5891: OR B
label_5892: OR C
label_5893: OR D
label_5894: OR E
label_5895: OR H
label_5896: OR L
label_5897: OR $12
label_5898: OR (HL)
label_5899: OR (IX + 127)
label_5900: OR (IY - 128)
label_5901: XOR A
label_5902: XOR B
label_5903: XOR C
label_5904: XOR D
label_5905: XOR E
label_5906: XOR H
label_5907: XOR L
label_5908: XOR $12
label_5909: XOR (HL)
label_5910: XOR (IX + 127)
label_5911: XOR (IY - 128)
label_5912: CP A
label_5913: CP B
label_5914: CP C
label_5915: CP D
label_5916: CP E
label_5917: CP H
label_5918: CP L
label_5919: CP $12
label_5920: CP (HL)
label_5921: CP (IX + 127)
label_5922: CP (IY - 128)
label_5923: INC A
label_5924: INC B
label_5925: INC C
label_5926: INC D
label_5927: INC E
label_5928: INC H
label_5929: INC L
label_5930: INC (HL)
label_5931: INC (IX + 127)
label_5932: INC (IY - 128)
label_5933: DEC A
label_5934: DEC B
label_5935: DEC C
label_5936: DEC D
label_5937: DEC E
label_5938: DEC H
label_5939: DEC L
label_5940: DEC (HL)
label_5941: DEC (IX + 127)
label_5942: DEC (IY - 128)
label_5943: DAA
label_5944: CPL
label_5945: NEG
label_5946: CCF
label_5947: SCF
label_5948: NOP
label_5949: HALT
label_5950: DI
label_5951: EI
label_5952: IM 0
label_5953: IM 1
label_5954: IM 2
label_5955: ADD HL,BC
label_5956: ADD HL,DE
label_5957: ADD HL,HL
label_5958: ADD HL,SP
label_5959: ADC HL,BC
label_5960: ADC HL,DE
label_5961: ADC HL,HL
label_5962: ADC HL,SP
label_5963: SBC HL,BC
label_5964: SBC HL,DE
label_5965: SBC HL,HL
label_5966: SBC HL,SP
label_5967: ADD IX,BC
label_5968: ADD IX,DE
label_5969: ADD IX,SP
label_5970: ADD IY,BC
label_5971: ADD IY,DE
label_5972: ADD IY,SP
label_5973: INC BC
label_5974: INC DE
label_5975: INC HL
label_5976: INC SP
label_5977: INC IX
label_5978: INC IY
label_5979: DEC BC
label_5980: DEC DE
label_5981: DEC HL
label_5982: DEC SP
label_5983: DEC IX
label_5984: DEC IY
label_5985: RLCA
label_5986: RLA
label_5987: RRCA
label_5988: RRA
label_5989: RLC A
label_5990: RLC B
label_5991: RLC C
label_5992: RLC D
label_5993: RLC E
label_5994: RLC H
label_5995: RLC L
label_5996: RLC (HL)
label_5997: RLC (IX + 127)
label_5998: RLC (IY - 128)
label_5999: RL A
label_6000: RL B
label_6001: RL C
label_6002: RL D
label_6003: RL E
label_6004: RL H
label_6005: RL L
label_6006: RL (HL)
label_6007: RL (IX + 127)
label_6008: RL (IY - 128)
label_6009: RRC A
label_6010: RRC B
label_6011: RRC C
label_6012: RRC D
label_6013: RRC E
label_6014: RRC H
label_6015: RRC L
label_6016: RRC (HL)
label_6017: RRC (IX + 127)
label_6018: RRC (IY - 128)
label_6019: RR A
label_6020: RR B
label_6021: RR C
label_6022: RR D
label_6023: RR E
label_6024: RR H
label_6025: RR L
label_6026: RR (HL)
label_6027: RR (IX + 127)
label_6028: RR (IY - 128)
label_6029: SLA A
label_6030: SLA B
label_6031: SLA C
label_6032: SLA D
label_6033: SLA E
label_6034: SLA H
label_6035: SLA L
label_6036: SLA (HL)
label_6037: SLA (IX + 127)
label_6038: SLA (IY - 128)
label_6039: SRA A
label_6040: SRA B
label_6041: SRA C
label_6042: SRA D
label_6043: SRA E
label_6044: SRA H
label_6045: SRA L
label_6046: SRA (HL)
label_6047: SRA (IX + 127)
label_6048: SRA (IY - 128)
label_6049: SRL A
label_6050: SRL B
label_6051: SRL C
label_6052: SRL D
label_6053: SRL E
label_6054: SRL H
label_6055: SRL L
label_6056: SRL (HL)
label_6057: SRL (IX + 127)
label_6058: SRL (IY - 128)
label_6059: RLD
label_6060: RRD
label_6061: BIT 0,A
label_6062: BIT 1,A
label_6063: BIT 2,A
label_6064: BIT 3,A
label_6065: BIT 4,A
label_6066: BIT 5,A
label_6067: BIT 6,A
label_6068: BIT 7,A
label_6069: BIT 0,B
label_6070: BIT 1,B
label_6071: BIT 2,B
label_6072: BIT 3,B
label_6073: BIT 4,B
label_6074: BIT 5,B
label_6075: BIT 6,B
label_6076: BIT 7,B
label_6077: BIT 0,C
label_6078: BIT 1,C
label_6079: BIT 2,C
label_6080: BIT 3,C
label_6081: BIT 4,C
label_6082: BIT 5,C
label_6083: BIT 6,C
label_6084: BIT 7,C
label_6085: BIT 0,D
label_6086: BIT 1,D
label_6087: BIT 2,D
label_6088: BIT 3,D
label_6089: BIT 4,D
label_6090: BIT 5,D
label_6091: BIT 6,D
label_6092: BIT 7,D
label_6093: BIT 0,E
label_6094: BIT 1,E
label_6095: BIT 2,E
label_6096: BIT 3,E
label_6097: BIT 4,E
label_6098: BIT 5,E
label_6099: BIT 6,E
label_6100: BIT 7,E
label_6101: BIT 0,H
label_6102: BIT 1,H
label_6103: BIT 2,H
label_6104: BIT 3,H
label_6105: BIT 4,H
label_6106: BIT 5,H
label_6107: BIT 6,H
label_6108: BIT 7,H
label_6109: BIT 0,L
label_6110: BIT 1,L
label_6111: BIT 2,L
label_6112: BIT 3,L
label_6113: BIT 4,L
label_6114: BIT 5,L
label_6115: BIT 6,L
label_6116: BIT 7,L
label_6117: BIT 0,(HL)
label_6118: BIT 1,(HL)
label_6119: BIT 2,(HL)
label_6120: BIT 3,(HL)
label_6121: BIT 4,(HL)
label_6122: BIT 5,(HL)
label_6123: BIT 6,(HL)
label_6124: BIT 7,(HL)
label_6125: BIT 0,(IX + 127)
label_6126: BIT 1,(IX + 127)
label_6127: BIT 2,(IX + 127)
label_6128: BIT 3,(IX + 127)
label_6129: BIT 4,(IX + 127)
label_6130: BIT 5,(IX + 127)
label_6131: BIT 6,(IX + 127)
label_6132: BIT 7,(IX + 127)
label_6133: BIT 0,(IY - 128)
label_6134: BIT 1,(IY - 128)
label_6135: BIT 2,(IY - 128)
label_6136: BIT 3,(IY - 128)
label_6137: BIT 4,(IY - 128)
label_6138: BIT 5,(IY - 128)
label_6139: BIT 6,(IY - 128)
label_6140: BIT 7,(IY - 128)
label_6141: SET 0,A
label_6142: SET 1,A
label_6143: SET 2,A
label_6144: SET 3,A
label_6145: SET 4,A
label_6146: SET 5,A
label_6147: SET 6,A
label_6148: SET 7,A
label_6149: SET 0,B
label_6150: SET 1,B
label_6151: SET 2,B
label_6152: SET 3,B
label_6153: SET 4,B
label_6154: SET 5,B
label_6155: SET 6,B
label_6156: SET 7,B
label_6157: SET 0,C
label_6158: SET 1,C
label_6159: SET 2,C
label_6160: SET 3,C
label_6161: SET 4,C
label_6162: SET 5,C
label_6163: SET 6,C
label_6164: SET 7,C
label_6165: SET 0,D
label_6166: SET 1,D
label_6167: SET 2,D
label_6168: SET 3,D
label_6169: SET 4,D
label_6170: SET 5,D
label_6171: SET 6,D
label_6172: SET 7,D
label_6173: SET 0,E
label_6174: SET 1,E
label_6175: SET 2,E
label_6176: SET 3,E
label_6177: SET 4,E
label_6178: SET 5,E
label_6179: SET 6,E
label_6180: SET 7,E
label_6181: SET 0,H
label_6182: SET 1,H
label_6183: SET 2,H
label_6184: SET 3,H
label_6185: SET 4,H
label_6186: SET 5,H
label_6187: SET 6,H
label_6188: SET 7,H
label_6189: SET 0,L
label_6190: SET 1,L
label_6191: SET 2,L
label_6192: SET 3,L
label_6193: SET 4,L
label_6194: SET 5,L
label_6195: SET 6,L
label_6196: SET 7,L
label_6197: SET 0,(HL)
label_6198: SET 1,(HL)
label_6199: SET 2,(HL)
label_6200: SET 3,(HL)
label_6201: SET 4,(HL)
label_6202: SET 5,(HL)
label_6203: SET 6,(HL)
label_6204: SET 7,(HL)
label_6205: SET 0,(IX + 127)
label_6206: SET 1,(IX + 127)
label_6207: SET 2,(IX + 127)
label_6208: SET 3,(IX + 127)
label_6209: SET 4,(IX + 127)
label_6210: SET 5,(IX + 127)
label_6211: SET 6,(IX + 127)
label_6212: SET 7,(IX + 127)
label_6213: SET 0,(IY - 128)
label_6214: SET 1,(IY - 128)
label_6215: SET 2,(IY - 128)
label_6216: SET 3,(IY - 128)
label_6217: SET 4,(IY - 128)
label_6218: SET 5,(IY - 128)
label_6219: SET 6,(IY - 128)
label_6220: SET 7,(IY - 128)
label_6221: RES 0,A
label_6222: RES 1,A
label_6223: RES 2,A
label_6224: RES 3,A
label_6225: RES 4,A
label_6226: RES 5,A
label_6227: RES 6,A
label_6228: RES 7,A
label_6229: RES 0,B
label_6230: RES 1,B
label_6231: RES 2,B
label_6232: RES 3,B
label_6233: RES 4,B
label_6234: RES 5,B
label_6235: RES 6,B
label_6236: RES 7,B
label_6237: RES 0,C
label_6238: RES 1,C
label_6239: RES 2,C
label_6240: RES 3,C
label_6241: RES 4,C
label_6242: RES 5,C
label_6243: RES 6,C
label_6244: RES 7,C
label_6245: RES 0,D
label_6246: RES 1,D
label_6247: RES 2,D
label_6248: RES 3,D
label_6249: RES 4,D
label_6250: RES 5,D
label_6251: RES 6,D
label_6252: RES 7,D
label_6253: RES 0,E
label_6254: RES 1,E
label_6255: RES 2,E
label_6256: RES 3,E
label_6257: RES 4,E
label_6258: RES 5,E
label_6259: RES 6,E
label_6260: RES 7,E
label_6261: RES 0,H
label_6262: RES 1,H
label_6263: RES 2,H
label_6264: RES 3,H
label_6265: RES 4,H
label_6266: RES 5,H
label_6267: RES 6,H
label_6268: RES 7,H
label_6269: RES 0,L
label_6270: RES 1,L
label_6271: RES 2,L
label_6272: RES 3,L
label_6273: RES 4,L
label_6274: RES 5,L
label_6275: RES 6,L
label_6276: RES 7,L
label_6277: RES 0,(HL)
label_6278: RES 1,(HL)
label_6279: RES 2,(HL)
label_6280: RES 3,(HL)
label_6281: RES 4,(HL)
label_6282: RES 5,(HL)
label_6283: RES 6,(HL)
label_6284: RES 7,(HL)
label_6285: RES 0,(IX + 127)
label_6286: RES 1,(IX + 127)
label_6287: RES 2,(IX + 127)
label_6288: RES 3,(IX + 127)
label_6289: RES 4,(IX + 127)
label_6290: RES 5,(IX + 127)
label_6291: RES 6,(IX + 127)
label_6292: RES 7,(IX + 127)
label_6293: RES 0,(IY - 128)
label_6294: RES 1,(IY - 128)
label_6295: RES 2,(IY - 128)
label_6296: RES 3,(IY - 128)
label_6297: RES 4,(IY - 128)
label_6298: RES 5,(IY - 128)
label_6299: RES 6,(IY - 128)
label_6300: RES 7,(IY - 128)
label_6301: JP $5678
label_6302: JP NZ,$5678
label_6303: JP Z,$5678
label_6304: JP NC,$5678
label_6305: JP C,$5678
label_6306: JP PO,$5678
label_6307: JP PE,$5678
label_6308: JP P,$5678
label_6309: JP M,$5678
label_6310: JR $ + 2
label_6311: JR NZ,$ + 2
label_6312: JR Z,$ + 2
label_6313: JR NC,$ + 2
label_6314: JR C,$ + 2
label_6315: JP (HL)
label_6316: JP (IX)
label_6317: JP (IY)
label_6318: DJNZ $ + 2
label_6319: CALL $5678
label_6320: CALL NZ,$5678
label_6321: CALL Z,$5678
label_6322: CALL NC,$5678
label_6323: CALL C,$5678
label_6324: CALL PO,$5678
label_6325: CALL PE,$5678
label_6326: CALL P,$5678
label_6327: CALL M,$5678
label_6328: RET
label_6329: RET NZ
label_6330: RET Z
label_6331: RET NC
label_6332: RET C
label_6333: RET PO
label_6334: RET PE
label_6335: RET P
label_6336: RET M
label_6337: RETI
label_6338: RETN
label_6339: RST $00
label_6340: RST $08
label_6341: RST $10
label_6342: RST $18
label_6343: RST $20
label_6344: RST $28
label_6345: RST $30
label_6346: RST $38
label_6347: IN A,($12)
label_6348: IN A,(C)
label_6349: IN B,(C)
label_6350: IN C,(C)
label_6351: IN D,(C)
label_6352: IN E,(C)
label_6353: IN H,(C)
label_6354: IN L,(C)
label_6355: IN F,(C)
label_6356: INI
label_6357: INIR
label_6358: IND
label_6359: INDR
label_6360: OUT ($12),A
label_6361: OUT (C),A
label_6362: OUT (C),B
label_6363: OUT (C),C
label_6364: OUT (C),D
label_6365: OUT (C),E
label_6366: OUT (C),H
label_6367: OUT (C),L
label_6368: OUTI
label_6369: OTIR
label_6370: OUTD
label_6371: OTDR
label_6372: LD A,A
label_6373: LD A,B
label_6374: LD A,C
label_6375: LD A,D
label_6376: LD A,E
label_6377: LD A,H
label_6378: LD A,L
label_6379: LD B,A
label_6380: LD B,B
label_6381: LD B,C
label_6382: LD B,D
label_6383: LD B,E
label_6384: LD B,H
label_6385: LD B,L
label_6386: LD C,A
label_6387: LD C,B
label_6388: LD C,C
label_6389: LD C,D
label_6390: LD C,E
label_6391: LD C,H
label_6392: LD C,L
label_6393: LD D,A
label_6394: LD D,B
label_6395: LD D,C
label_6396: LD D,D
label_6397: LD D,E
label_6398: LD D,H
label_6399: LD D,L
label_6400: LD E,A
label_6401: LD E,B
label_6402: LD E,C
label_6403: LD E,D
label_6404: LD E,E
label_6405: LD E,H
label_6406: LD E,L
label_6407: LD H,A
label_6408: LD H,B
label_6409: LD H,C
label_6410: LD H,D
label_6411: LD H,E
label_6412: LD H,H
label_6413: LD H,L
label_6414: LD L,A
label_6415: LD L,B
label_6416: LD L,C
label_6417: LD L,D
label_6418: LD L,E
label_6419: LD L,H
label_6420: LD L,L
label_6421: LD A,$12
label_6422: LD B,$12
label_6423: LD C,$12
label_6424: LD D,$12
label_6425: LD E,$12
label_6426: LD H,$12
label_6427: LD L,$12
label_6428: LD A,(HL)
label_6429: LD B,(HL)
label_6430: LD C,(HL)
label_6431: LD D,(HL)
label_6432: LD E,(HL)
label_6433: LD H,(HL)
label_6434: LD L,(HL)
label_6435: LD A,(IX + 127)
label_6436: LD B,(IX + 127)
label_6437: LD C,(IX + 127)
label_6438: LD D,(IX + 127)
label_6439: LD E,(IX + 127)
label_6440: LD H,(IX + 127)
label_6441: LD L,(IX + 127)
label_6442: LD A,(IY - 128)
label_6443: LD B,(IY - 128)
label_6444: LD C,(IY - 128)
label_6445: LD D,(IY - 128)
label_6446: LD E,(IY - 128)
label_6447: LD H,(IY - 128)
label_6448: LD L,(IY - 128)
label_6449: LD (HL),A
label_6450: LD (HL),B
label_6451: LD (HL),C
label_6452: LD (HL),D
label_6453: LD (HL),E
label_6454: LD (HL),H
label_6455: LD (HL),L
label_6456: LD (IX + 127),A
label_6457: LD (IX + 127),B
label_6458: LD (IX + 127),C
label_6459: LD (IX + 127),D
label_6460: LD (IX + 127),E
label_6461: LD (IX + 127),H
label_6462: LD (IX + 127),L
label_6463: LD (IY - 128),A
label_6464: LD (IY - 128),B
label_6465: LD (IY - 128),C
label_6466: LD (IY - 128),D
label_6467: LD (IY - 128),E
label_6468: LD (IY - 128),H
label_6469: LD (IY - 128),L
label_6470: LD (HL),$12
label_6471: LD (IX + 127),$12
label_6472: LD (IY - 128),$12
label_6473: LD A,(BC)
label_6474: LD A,(DE)
label_6475: LD A,($5678)
label_6476: LD (BC),A
label_6477: LD (DE),A
label_6478: LD ($5678),A
label_6479: LD A,I
label_6480: LD A,R
label_6481: LD I,A
label_6482: LD R,A
label_6483: LD BC,$5678
label_6484: LD DE,$5678
label_6485: LD HL,$5678
label_6486: LD SP,$5678
label_6487: LD IX,$5678
label_6488: LD IY,$5678
label_6489: LD HL,($5678)
label_6490: LD BC,($5678)
label_6491: LD DE,($5678)
label_6492: LD HL,($5678)
label_6493: LD SP,($5678)
label_6494: LD IX,($5678)
label_6495: LD IY,($5678)
label_6496: LD ($5678),HL
label_6497: LD ($5678),BC
label_6498: LD ($5678),DE
label_6499: LD ($5678),HL
label_6500: LD ($5678),SP
label_6501: LD ($5678),IX
label_6502: LD ($5678),IY
label_6503: LD SP,HL
label_6504: LD SP,IX
label_6505: LD SP,IY
label_6506: PUSH BC
label_6507: PUSH DE
label_6508: PUSH HL
label_6509: PUSH AF
label_6510: PUSH IX
label_6511: PUSH IY
label_6512: POP BC
label_6513: POP DE
label_6514: POP HL
label_6515: POP AF
label_6516: POP IX
label_6517: POP IY
label_6518: EX DE,HL
label_6519: EX AF,AF'
label_6520: EXX
label_6521: EX (SP),HL
label_6522: EX (SP),IX
label_6523: EX (SP),IY
label_6524: LDI
label_6525: LDIR
label_6526: LDD
label_6527: LDDR
label_6528: CPI
label_6529: CPIR
label_6530: CPD
label_6531: CPDR
label_6532: ADD A,A
label_6533: ADD A,B
label_6534: ADD A,C
label_6535: ADD A,D
label_6536: ADD A,E
label_6537: ADD A,H
label_6538: ADD A,L
label_6539: ADD A,$12
label_6540: ADD A,(HL)
label_6541: ADD A,(IX + 127)
label_6542: ADD A,(IY - 128)
label_6543: ADC A,A
label_6544: ADC A,B
label_6545: ADC A,C
label_6546: ADC A,D
label_6547: ADC A,E
label_6548: ADC A,H
label_6549: ADC A,L
label_6550: ADC A,$12
label_6551: ADC A,(HL)
label_6552: ADC A,(IX + 127)
label_6553: ADC A,(IY - 128)
label_6554: SUB A
label_6555: SUB B
label_6556: SUB C
label_6557: SUB D
label_6558: SUB E
label_6559: SUB H
label_6560: SUB L
label_6561: SUB $12
label_6562: SUB (HL)
label_6563: SUB (IX + 127)
label_6564: SUB (IY - 128)
label_6565: SBC A,A
label_6566: SBC A,B
label_6567: SBC A,C
label_6568: SBC A,D
label_6569: SBC A,E
label_6570: SBC A,H
label_6571: SBC A,L
label_6572: SBC A,$12
label_6573: SBC A,(HL)
label_6574: SBC A,(IX + 127)
label_6575: SBC A,(IY - 128)
label_6576: AND A
label_6577: AND B
label_6578: AND C
label_6579: AND D
label_6580: AND E
label_6581: AND H
label_6582: AND L
label_6583: AND $12
label_6584: AND (HL)
label_6585: AND (IX + 127)
label_6586: AND (IY - 128)
label_6587: AND A
label_6588: AND B
label_6589: AND C
label_6590: AND D
label_6591: AND E
label_6592: AND H
label_6593: AND L
label_6594: AND $12
label_6595: AND (HL)
label_6596: AND (IX + 127)
label_6597: AND (IY - 128)
label_6598: OR A
label_6599: OR B
label_6600: OR C
label_6601: OR D
label_6602: OR E
label_6603: OR H
label_6604: OR L
label_6605: OR $12
label_6606: OR (HL)
label_6607: OR (IX + 127)
label_6608: OR (IY - 128)
label_6609: XOR A
label_6610: XOR B
label_6611: XOR C
label_6612: XOR D
label_6613: XOR E
label_6614: XOR H
label_6615: XOR L
label_6616: XOR $12
label_6617: XOR (HL)
label_6618: XOR (IX + 127)
label_6619: XOR (IY - 128)
label_6620: CP A
label_6621: CP B
label_6622: CP C
label_6623: CP D
label_6624: CP E
label_6625: CP H
label_6626: CP L
label_6627: CP $12
label_6628: CP (HL)
label_6629: CP (IX + 127)
label_6630: CP (IY - 128)
label_6631: INC A
label_6632: INC B
label_6633: INC C
label_6634: INC D
label_6635: INC E
label_6636: INC H
label_6637: INC L
label_6638: INC (HL)
label_6639: INC (IX + 127)
label_6640: INC (IY - 128)
label_6641: DEC A
label_6642: DEC B
label_6643: DEC C
label_6644: DEC D
label_6645: DEC E
label_6646: DEC H
label_6647: DEC L
label_6648: DEC (HL)
label_6649: DEC (IX + 127)
label_6650: DEC (IY - 128)
label_6651: DAA
label_6652: CPL
label_6653: NEG
label_6654: CCF
label_6655: SCF
label_6656: NOP
label_6657: HALT
label_6658: DI
label_6659: EI
label_6660: IM 0
label_6661: IM 1
label_6662: IM 2
label_6663: ADD HL,BC
label_6664: ADD HL,DE
label_6665: ADD HL,HL
label_6666: ADD HL,SP
label_6667: ADC HL,BC
label_6668: ADC HL,DE
label_6669: ADC HL,HL
label_6670: ADC HL,SP
label_6671: SBC HL,BC
label_6672: SBC HL,DE
label_6673: SBC HL,HL
label_6674: SBC HL,SP
label_6675: ADD IX,BC
label_6676: ADD IX,DE
label_6677: ADD IX,SP
label_6678: ADD IY,BC
label_6679: ADD IY,DE
label_6680: ADD IY,SP
label_6681: INC BC
label_6682: INC DE
label_6683: INC HL
label_6684: INC SP
label_6685: INC IX
label_6686: INC IY
label_6687: DEC BC
label_6688: DEC DE
label_6689: DEC HL
label_6690: DEC SP
label_6691: DEC IX
label_6692: DEC IY
label_6693: RLCA
label_6694: RLA
label_6695: RRCA
label_6696: RRA
label_6697: RLC A
label_6698: RLC B
label_6699: RLC C
label_6700: RLC D
label_6701: RLC E
label_6702: RLC H
label_6703: RLC L
label_6704: RLC (HL)
label_6705: RLC (IX + 127)
label_6706: RLC (IY - 128)
label_6707: RL A
label_6708: RL B
label_6709: RL C
label_6710: RL D
label_6711: RL E
label_6712: RL H
label_6713: RL L
label_6714: RL (HL)
label_6715: RL (IX + 127)
label_6716: RL (IY - 128)
label_6717: RRC A
label_6718: RRC B
label_6719: RRC C
label_6720: RRC D
label_6721: RRC E
label_6722: RRC H
label_6723: RRC L
label_6724: RRC (HL)
label_6725: RRC (IX + 127)
label_6726: RRC (IY - 128)
label_6727: RR A
label_6728: RR B
label_6729: RR C
label_6730: RR D
label_6731: RR E
label_6732: RR H
label_6733: RR L
label_6734: RR (HL)
label_6735: RR (IX + 127)
label_6736: RR (IY - 128)
label_6737: SLA A
label_6738: SLA B
label_6739: SLA C
label_6740: SLA D
label_6741: SLA E
label_6742: SLA H
label_6743: SLA L
label_6744: SLA (HL)
label_6745: SLA (IX + 127)
label_6746: SLA (IY - 128)
label_6747: SRA A
label_6748: SRA B
label_6749: SRA C
label_6750: SRA D
label_6751: SRA E
label_6752: SRA H
label_6753: SRA L
label_6754: SRA (HL)
label_6755: SRA (IX + 127)
label_6756: SRA (IY - 128)
label_6757: SRL A
label_6758: SRL B
label_6759: SRL C
label_6760: SRL D
label_6761: SRL E
label_6762: SRL H
label_6763: SRL L
label_6764: SRL (HL)
label_6765: SRL (IX + 127)
label_6766: SRL (IY - 128)
label_6767: RLD
label_6768: RRD
label_6769: BIT 0,A
label_6770: BIT 1,A
label_6771: BIT 2,A
label_6772: BIT 3,A
label_6773: BIT 4,A
label_6774: BIT 5,A
label_6775: BIT 6,A
label_6776: BIT 7,A
label_6777: BIT 0,B
label_6778: BIT 1,B
label_6779: BIT 2,B
label_6780: BIT 3,B
label_6781: BIT 4,B
label_6782: BIT 5,B
label_6783: BIT 6,B
label_6784: BIT 7,B
label_6785: BIT 0,C
label_6786: BIT 1,C
label_6787: BIT 2,C
label_6788: BIT 3,C
label_6789: BIT 4,C
label_6790: BIT 5,C
label_6791: BIT 6,C
label_6792: BIT 7,C
label_6793: BIT 0,D
label_6794: BIT 1,D
label_6795: BIT 2,D
label_6796: BIT 3,D
label_6797: BIT 4,D
label_6798: BIT 5,D
label_6799: BIT 6,D
label_6800: BIT 7,D
label_6801: BIT 0,E
label_6802: BIT 1,E
label_6803: BIT 2,E
label_6804: BIT 3,E
label_6805: BIT 4,E
label_6806: BIT 5,E
label_6807: BIT 6,E
label_6808: BIT 7,E
label_6809: BIT 0,H
label_6810: BIT 1,H
label_6811: BIT 2,H
label_6812: BIT 3,H
label_6813: BIT 4,H
label_6814: BIT 5,H
label_6815: BIT 6,H
label_6816: BIT 7,H
label_6817: BIT 0,L
label_6818: BIT 1,L
label_6819: BIT 2,L
label_6820: BIT 3,L
label_6821: BIT 4,L
label_6822: BIT 5,L
label_6823: BIT 6,L
label_6824: BIT 7,L
label_6825: BIT 0,(HL)
label_6826: BIT 1,(HL)
label_6827: BIT 2,(HL)
label_6828: BIT 3,(HL)
label_6829: BIT 4,(HL)
label_6830: BIT 5,(HL)
label_6831: BIT 6,(HL)
label_6832: BIT 7,(HL)
label_6833: BIT 0,(IX + 127)
label_6834: BIT 1,(IX + 127)
label_6835: BIT 2,(IX + 127)
label_6836: BIT 3,(IX + 127)
label_6837: BIT 4,(IX + 127)
label_6838: BIT 5,(IX + 127)
label_6839: BIT 6,(IX + 127)
label_6840: BIT 7,(IX + 127)
label_6841: BIT 0,(IY - 128)
label_6842: BIT 1,(IY - 128)
label_6843: BIT 2,(IY - 128)
label_6844: BIT 3,(IY - 128)
label_6845: BIT 4,(IY - 128)
label_6846: BIT 5,(IY - 128)
label_6847: BIT 6,(IY - 128)
label_6848: BIT 7,(IY - 128)
label_6849: SET 0,A
label_6850: SET 1,A
label_6851: SET 2,A
label_6852: SET 3,A
label_6853: SET 4,A
label_6854: SET 5,A
label_6855: SET 6,A
label_6856: SET 7,A
label_6857: SET 0,B
label_6858: SET 1,B
label_6859: SET 2,B
label_6860: SET 3,B
label_6861: SET 4,B
label_6862: SET 5,B
label_6863: SET 6,B
label_6864: SET 7,B
label_6865: SET 0,C
label_6866: SET 1,C
label_6867: SET 2,C
label_6868: SET 3,C
label_6869: SET 4,C
label_6870: SET 5,C
label_6871: SET 6,C
label_6872: SET 7,C
label_6873: SET 0,D
label_6874: SET 1,D
label_6875: SET 2,D
label_6876: SET 3,D
label_6877: SET 4,D
label_6878: SET 5,D
label_6879: SET 6,D
label_6880: SET 7,D
label_6881: SET 0,E
label_6882: SET 1,E
label_6883: SET 2,E
label_6884: SET 3,E
label_6885: SET 4,E
label_6886: SET 5,E
label_6887: SET 6,E
label_6888: SET 7,E
label_6889: SET 0,H
label_6890: SET 1,H
label_6891: SET 2,H
label_6892: SET 3,H
label_6893: SET 4,H
label_6894: SET 5,H
label_6895: SET 6,H
label_6896: SET 7,H
label_6897: SET 0,L
label_6898: SET 1,L
label_6899: SET 2,L
label_6900: SET 3,L
label_6901: SET 4,L
label_6902: SET 5,L
label_6903: SET 6,L
label_6904: SET 7,L
label_6905: SET 0,(HL)
label_6906: SET 1,(HL)
label_6907: SET 2,(HL)
label_6908: SET 3,(HL)
label_6909: SET 4,(HL)
label_6910: SET 5,(HL)
label_6911: SET 6,(HL)
label_6912: SET 7,(HL)
label_6913: SET 0,(IX + 127)
label_6914: SET 1,(IX + 127)
label_6915: SET 2,(IX + 127)
label_6916: SET 3,(IX + 127)
label_6917: SET 4,(IX + 127)
label_6918: SET 5,(IX + 127)
label_6919: SET 6,(IX + 127)
label_6920: SET 7,(IX + 127)
label_6921: SET 0,(IY - 128)
label_6922: SET 1,(IY - 128)
label_6923: SET 2,(IY - 128)
label_6924: SET 3,(IY - 128)
label_6925: SET 4,(IY - 128)
label_6926: SET 5,(IY - 128)
label_6927: SET 6,(IY - 128)
label_6928: SET 7,(IY - 128)
label_6929: RES 0,A
label_6930: RES 1,A
label_6931: RES 2,A
label_6932: RES 3,A
label_6933: RES 4,A
label_6934: RES 5,A
label_6935: RES 6,A
label_6936: RES 7,A
label_6937: RES 0,B
label_6938: RES 1,B
label_6939: RES 2,B
label_6940: RES 3,B
label_6941: RES 4,B
label_6942: RES 5,B
label_6943: RES 6,B
label_6944: RES 7,B
label_6945: RES 0,C
label_6946: RES 1,C
label_6947: RES 2,C
label_6948: RES 3,C
label_6949: RES 4,C
label_6950: RES 5,C
label_6951: RES 6,C
label_6952: RES 7,C
label_6953: RES 0,D
label_6954: RES 1,D
label_6955: RES 2,D
label_6956: RES 3,D
label_6957: RES 4,D
label_6958: RES 5,D
label_6959: RES 6,D
label_6960: RES 7,D
label_6961: RES 0,E
label_6962: RES 1,E
label_6963: RES 2,E
label_6964: RES 3,E
label_6965: RES 4,E
label_6966: RES 5,E
label_6967: RES 6,E
label_6968: RES 7,E
label_6969: RES 0,H
label_6970: RES 1,H
label_6971: RES 2,H
label_6972: RES 3,H
label_6973: RES 4,H
label_6974: RES 5,H
label_6975: RES 6,H
label_6976: RES 7,H
label_6977: RES 0,L
label_6978: RES 1,L
label_6979: RES 2,L
label_6980: RES 3,L
label_6981: RES 4,L
label_6982: RES 5,L
label_6983: RES 6,L
label_6984: RES 7,L
label_6985: RES 0,(HL)
label_6986: RES 1,(HL)
label_6987: RES 2,(HL)
label_6988: RES 3,(HL)
label_6989: RES 4,(HL)
label_6990: RES 5,(HL)
label_6991: RES 6,(HL)
label_6992: RES 7,(HL)
label_6993: RES 0,(IX + 127)
label_6994: RES 1,(IX + 127)
label_6995: RES 2,(IX + 127)
label_6996: RES 3,(IX + 127)
label_6997: RES 4,(IX + 127)
label_6998: RES 5,(IX + 127)
label_6999: RES 6,(IX + 127)
label_7000: RES 7,(IX + 127)
label_7001: RES 0,(IY - 128)
label_7002: RES 1,(IY - 128)
label_7003: RES 2,(IY - 128)
label_7004: RES 3,(IY - 128)
label_7005: RES 4,(IY - 128)
label_7006: RES 5,(IY - 128)
label_7007: RES 6,(IY - 128)
label_7008: RES 7,(IY - 128)
label_7009: JP $5678
label_7010: JP NZ,$5678
label_7011: JP Z,$5678
label_7012: JP NC,$5678
label_7013: JP C,$5678
label_7014: JP PO,$5678
label_7015: JP PE,$5678
label_7016: JP P,$5678
label_7017: JP M,$5678
label_7018: JR $ + 2
label_7019: JR NZ,$ + 2
label_7020: JR Z,$ + 2
label_7021: JR NC,$ + 2
label_7022: JR C,$ + 2
label_7023: JP (HL)
label_7024: JP (IX)
label_7025: JP (IY)
label_7026: DJNZ $ + 2
label_7027: CALL $5678
label_7028: CALL NZ,$5678
label_7029: CALL Z,$5678
label_7030: CALL NC,$5678
label_7031: CALL C,$5678
label_7032: CALL PO,$5678
label_7033: CALL PE,$5678
label_7034: CALL P,$5678
label_7035: CALL M,$5678
label_7036: RET
label_7037: RET NZ
label_7038: RET Z
label_7039: RET NC
label_7040: RET C
label_7041: RET PO
label_7042: RET PE
label_7043: RET P
label_7044: RET M
label_7045: RETI
label_7046: RETN
label_7047: RST $00
label_7048: RST $08
label_7049: RST $10
label_7050: RST $18
label_7051: RST $20
label_7052: RST $28
label_7053: RST $30
label_7054: RST $38
label_7055: IN A,($12)
label_7056: IN A,(C)
label_7057: IN B,(C)
label_7058: IN C,(C)
label_7059: IN D,(C)
label_7060: IN E,(C)
label_7061: IN H,(C)
label_7062: IN L,(C)
label_7063: IN F,(C)
label_7064: INI
label_7065: INIR
label_7066: IND
label_7067: INDR
label_7068: OUT ($12),A
label_7069: OUT (C),A
label_7070: OUT (C),B
label_7071: OUT (C),C
label_7072: OUT (C),D
label_7073: OUT (C),E
label_7074: OUT (C),H
label_7075: OUT (C),L
label_7076: OUTI
label_7077: OTIR
label_7078: OUTD
label_7079: OTDR
label_7080: LD A,A
label_7081: LD A,B
label_7082: LD A,C
label_7083: LD A,D
label_7084: LD A,E
label_7085: LD A,H
label_7086: LD A,L
label_7087: LD B,A
label_7088: LD B,B
label_7089: LD B,C
label_7090: LD B,D
label_7091: LD B,E
label_7092: LD B,H
label_7093: LD B,L
label_7094: LD C,A
label_7095: LD C,B
label_7096: LD C,C
label_7097: LD C,D
label_7098: LD C,E
label_7099: LD C,H
label_7100: LD C,L
label_7101: LD D,A
label_7102: LD D,B
label_7103: LD D,C
label_7104: LD D,D
label_7105: LD D,E
label_7106: LD D,H
label_7107: LD D,L
label_7108: LD E,A
label_7109: LD E,B
label_7110: LD E,C
label_7111: LD E,D
label_7112: LD E,E
label_7113: LD E,H
label_7114: LD E,L
label_7115: LD H,A
label_7116: LD H,B
label_7117: LD H,C
label_7118: LD H,D
label_7119: LD H,E
label_7120: LD H,H
label_7121: LD H,L
label_7122: LD L,A
label_7123: LD L,B
label_7124: LD L,C
label_7125: LD L,D
label_7126: LD L,E
label_7127: LD L,H
label_7128: LD L,L
label_7129: LD A,$12
label_7130: LD B,$12
label_7131: LD C,$12
label_7132: LD D,$12
label_7133: LD E,$12
label_7134: LD H,$12
label_7135: LD L,$12
label_7136: LD A,(HL)
label_7137: LD B,(HL)
label_7138: LD C,(HL)
label_7139: LD D,(HL)
label_7140: LD E,(HL)
label_7141: LD H,(HL)
label_7142: LD L,(HL)
label_7143: LD A,(IX + 127)
label_7144: LD B,(IX + 127)
label_7145: LD C,(IX + 127)
label_7146: LD D,(IX + 127)
label_7147: LD E,(IX + 127)
label_7148: LD H,(IX + 127)
label_7149: LD L,(IX + 127)
label_7150: LD A,(IY - 128)
label_7151: LD B,(IY - 128)
label_7152: LD C,(IY - 128)
label_7153: LD D,(IY - 128)
label_7154: LD E,(IY - 128)
label_7155: LD H,(IY - 128)
label_7156: LD L,(IY - 128)
label_7157: LD (HL),A
label_7158: LD (HL),B
label_7159: LD (HL),C
label_7160: LD (HL),D
label_7161: LD (HL),E
label_7162: LD (HL),H
label_7163: LD (HL),L
label_7164: LD (IX + 127),A
label_7165: LD (IX + 127),B
label_7166: LD (IX + 127),C
label_7167: LD (IX + 127),D
label_7168: LD (IX + 127),E
label_7169: LD (IX + 127),H
label_7170: LD (IX + 127),L
label_7171: LD (IY - 128),A
label_7172: LD (IY - 128),B
label_7173: LD (IY - 128),C
label_7174: LD (IY - 128),D
label_7175: LD (IY - 128),E
label_7176: LD (IY - 128),H
label_7177: LD (IY - 128),L
label_7178: LD (HL),$12
label_7179: LD (IX + 127),$12
label_7180: LD (IY - 128),$12
label_7181: LD A,(BC)
label_7182: LD A,(DE)
label_7183: LD A,($5678)
label_7184: LD (BC),A
label_7185: LD (DE),A
label_7186: LD ($5678),A
label_7187: LD A,I
label_7188: LD A,R
label_7189: LD I,A
label_7190: LD R,A
label_7191: LD BC,$5678
label_7192: LD DE,$5678
label_7193: LD HL,$5678
label_7194: LD SP,$5678
label_7195: LD IX,$5678
label_7196: LD IY,$5678
label_7197: LD HL,($5678)
label_7198: LD BC,($5678)
label_7199: LD DE,($5678)
label_7200: LD HL,($5678)
label_7201: LD SP,($5678)
label_7202: LD IX,($5678)
label_7203: LD IY,($5678)
label_7204: LD ($5678),HL
label_7205: LD ($5678),BC
label_7206: LD ($5678),DE
label_7207: LD ($5678),HL
label_7208: LD ($5678),SP
label_7209: LD ($5678),IX
label_7210: LD ($5678),IY
label_7211: LD SP,HL
label_7212: LD SP,IX
label_7213: LD SP,IY
label_7214: PUSH BC
label_7215: PUSH DE
label_7216: PUSH HL
label_7217: PUSH AF
label_7218: PUSH IX
label_7219: PUSH IY
label_7220: POP BC
label_7221: POP DE
label_7222: POP HL
label_7223: POP AF
label_7224: POP IX
label_7225: POP IY
label_7226: EX DE,HL
label_7227: EX AF,AF'
label_7228: EXX
label_7229: EX (SP),HL
label_7230: EX (SP),IX
label_7231: EX (SP),IY
label_7232: LDI
label_7233: LDIR
label_7234: LDD
label_7235: LDDR
label_7236: CPI
label_7237: CPIR
label_7238: CPD
label_7239: CPDR
label_7240: ADD A,A
label_7241: ADD A,B
label_7242: ADD A,C
label_7243: ADD A,D
label_7244: ADD A,E
label_7245: ADD A,H
label_7246: ADD A,L
label_7247: ADD A,$12
label_7248: ADD A,(HL)
label_7249: ADD A,(IX + 127)
label_7250: ADD A,(IY - 128)
label_7251: ADC A,A
label_7252: ADC A,B
label_7253: ADC A,C
label_7254: ADC A,D
label_7255: ADC A,E
label_7256: ADC A,H
label_7257: ADC A,L
label_7258: ADC A,$12
label_7259: ADC A,(HL)
label_7260: ADC A,(IX + 127)
label_7261: ADC A,(IY - 128)
label_7262: SUB A
label_7263: SUB B
label_7264: SUB C
label_7265: SUB D
label_7266: SUB E
label_7267: SUB H
label_7268: SUB L
label_7269: SUB $12
label_7270: SUB (HL)
label_7271: SUB (IX + 127)
label_7272: SUB (IY - 128)
label_7273: SBC A,A
label_7274: SBC A,B
label_7275: SBC A,C
label_7276: SBC A,D
label_7277: SBC A,E
label_7278: SBC A,H
label_7279: SBC A,L
label_7280: SBC A,$12
label_7281: SBC A,(HL)
label_7282: SBC A,(IX + 127)
label_7283: SBC A,(IY - 128)
label_7284: AND A
label_7285: AND B
label_7286: AND C
label_7287: AND D
label_7288: AND E
label_7289: AND H
label_7290: AND L
label_7291: AND $12
label_7292: AND (HL)
label_7293: AND (IX + 127)
label_7294: AND (IY - 128)
label_7295: AND A
label_7296: AND B
label_7297: AND C
label_7298: AND D
label_7299: AND E
label_7300: AND H
label_7301: AND L
label_7302: AND $12
label_7303: AND (HL)
label_7304: AND (IX + 127)
label_7305: AND (IY - 128)
label_7306: OR A
label_7307: OR B
label_7308: OR C
label_7309: OR D
label_7310: OR E
label_7311: OR H
label_7312: OR L
label_7313: OR $12
label_7314: OR (HL)
label_7315: OR (IX + 127)
label_7316: OR (IY - 128)
label_7317: XOR A
label_7318: XOR B
label_7319: XOR C
label_7320: XOR D
label_7321: XOR E
label_7322: XOR H
label_7323: XOR L
label_7324: XOR $12
label_7325: XOR (HL)
label_7326: XOR (IX + 127)
label_7327: XOR (IY - 128)
label_7328: CP A
label_7329: CP B
label_7330: CP C
label_7331: CP D
label_7332: CP E
label_7333: CP H
label_7334: CP L
label_7335: CP $12
label_7336: CP (HL)
label_7337: CP (IX + 127)
label_7338: CP (IY - 128)
label_7339: INC A
label_7340: INC B
label_7341: INC C
label_7342: INC D
label_7343: INC E
label_7344: INC H
label_7345: INC L
label_7346: INC (HL)
label_7347: INC (IX + 127)
label_7348: INC (IY - 128)
label_7349: DEC A
label_7350: DEC B
label_7351: DEC C
label_7352: DEC D
label_7353: DEC E
label_7354: DEC H
label_7355: DEC L
label_7356: DEC (HL)
label_7357: DEC (IX + 127)
label_7358: DEC (IY - 128)
label_7359: DAA
label_7360: CPL
label_7361: NEG
label_7362: CCF
label_7363: SCF
label_7364: NOP
label_7365: HALT
label_7366: DI
label_7367: EI
label_7368: IM 0
label_7369: IM 1
label_7370: IM 2
label_7371: ADD HL,BC
label_7372: ADD HL,DE
label_7373: ADD HL,HL
label_7374: ADD HL,SP
label_7375: ADC HL,BC
label_7376: ADC HL,DE
label_7377: ADC HL,HL
label_7378: ADC HL,SP
label_7379: SBC HL,BC
label_7380: SBC HL,DE
label_7381: SBC HL,HL
label_7382: SBC HL,SP
label_7383: ADD IX,BC
label_7384: ADD IX,DE
label_7385: ADD IX,SP
label_7386: ADD IY,BC
label_7387: ADD IY,DE
label_7388: ADD IY,SP
label_7389: INC BC
label_7390: INC DE
label_7391: INC HL
label_7392: INC SP
label_7393: INC IX
label_7394: INC IY
label_7395: DEC BC
label_7396: DEC DE
label_7397: DEC HL
label_7398: DEC SP
label_7399: DEC IX
label_7400: DEC IY
label_7401: RLCA
label_7402: RLA
label_7403: RRCA
label_7404: RRA
label_7405: RLC A
label_7406: RLC B
label_7407: RLC C
label_7408: RLC D
label_7409: RLC E
label_7410: RLC H
label_7411: RLC L
label_7412: RLC (HL)
label_7413: RLC (IX + 127)
label_7414: RLC (IY - 128)
label_7415: RL A
label_7416: RL B
label_7417: RL C
label_7418: RL D
label_7419: RL E
label_7420: RL H
label_7421: RL L
label_7422: RL (HL)
label_7423: RL (IX + 127)
label_7424: RL (IY - 128)
label_7425: RRC A
label_7426: RRC B
label_7427: RRC C
label_7428: RRC D
label_7429: RRC E
label_7430: RRC H
label_7431: RRC L
label_7432: RRC (HL)
label_7433: RRC (IX + 127)
label_7434: RRC (IY - 128)
label_7435: RR A
label_7436: RR B
label_7437: RR C
label_7438: RR D
label_7439: RR E
label_7440: RR H
label_7441: RR L
label_7442: RR (HL)
label_7443: RR (IX + 127)
label_7444: RR (IY - 128)
label_7445: SLA A
label_7446: SLA B
label_7447: SLA C
label_7448: SLA D
label_7449: SLA E
label_7450: SLA H
label_7451: SLA L
label_7452: SLA (HL)
label_7453: SLA (IX + 127)
label_7454: SLA (IY - 128)
label_7455: SRA A
label_7456: SRA B
label_7457: SRA C
label_7458: SRA D
label_7459: SRA E
label_7460: SRA H
label_7461: SRA L
label_7462: SRA (HL)
label_7463: SRA (IX + 127)
label_7464: SRA (IY - 128)
label_7465: SRL A
label_7466: SRL B
label_7467: SRL C
label_7468: SRL D
label_7469: SRL E
label_7470: SRL H
label_7471: SRL L
label_7472: SRL (HL)
label_7473: SRL (IX + 127)
label_7474: SRL (IY - 128)
label_7475: RLD
label_7476: RRD
label_7477: BIT 0,A
label_7478: BIT 1,A
label_7479: BIT 2,A
label_7480: BIT 3,A
label_7481: BIT 4,A
label_7482: BIT 5,A
label_7483: BIT 6,A
label_7484: BIT 7,A
label_7485: BIT 0,B
label_7486: BIT 1,B
label_7487: BIT 2,B
label_7488: BIT 3,B
label_7489: BIT 4,B
label_7490: BIT 5,B
label_7491: BIT 6,B
label_7492: BIT 7,B
label_7493: BIT 0,C
label_7494: BIT 1,C
label_7495: BIT 2,C
label_7496: BIT 3,C
label_7497: BIT 4,C
label_7498: BIT 5,C
label_7499: BIT 6,C
label_7500: BIT 7,C
label_7501: BIT 0,D
label_7502: BIT 1,D
label_7503: BIT 2,D
label_7504: BIT 3,D
label_7505: BIT 4,D
label_7506: BIT 5,D
label_7507: BIT 6,D
label_7508: BIT 7,D
label_7509: BIT 0,E
label_7510: BIT 1,E
label_7511: BIT 2,E
label_7512: BIT 3,E
label_7513: BIT 4,E
label_7514: BIT 5,E
label_7515: BIT 6,E
label_7516: BIT 7,E
label_7517: BIT 0,H
label_7518: BIT 1,H
label_7519: BIT 2,H
label_7520: BIT 3,H
label_7521: BIT 4,H
label_7522: BIT 5,H
label_7523: BIT 6,H
label_7524: BIT 7,H
label_7525: BIT 0,L
label_7526: BIT 1,L
label_7527: BIT 2,L
label_7528: BIT 3,L
label_7529: BIT 4,L
label_7530: BIT 5,L
label_7531: BIT 6,L
label_7532: BIT 7,L
label_7533: BIT 0,(HL)
label_7534: BIT 1,(HL)
label_7535: BIT 2,(HL)
label_7536: BIT 3,(HL)
label_7537: BIT 4,(HL)
label_7538: BIT 5,(HL)
label_7539: BIT 6,(HL)
label_7540: BIT 7,(HL)
label_7541: BIT 0,(IX + 127)
label_7542: BIT 1,(IX + 127)
label_7543: BIT 2,(IX + 127)
label_7544: BIT 3,(IX + 127)
label_7545: BIT 4,(IX + 127)
label_7546: BIT 5,(IX + 127)
label_7547: BIT 6,(IX + 127)
label_7548: BIT 7,(IX + 127)
label_7549: BIT 0,(IY - 128)
label_7550: BIT 1,(IY - 128)
label_7551: BIT 2,(IY - 128)
label_7552: BIT 3,(IY - 128)
label_7553: BIT 4,(IY - 128)
label_7554: BIT 5,(IY - 128)
label_7555: BIT 6,(IY - 128)
label_7556: BIT 7,(IY - 128)
label_7557: SET 0,A
label_7558: SET 1,A
label_7559: SET 2,A
label_7560: SET 3,A
label_7561: SET 4,A
label_7562: SET 5,A
label_7563: SET 6,A
label_7564: SET 7,A
label_7565: SET 0,B
label_7566: SET 1,B
label_7567: SET 2,B
label_7568: SET 3,B
label_7569: SET 4,B
label_7570: SET 5,B
label_7571: SET 6,B
label_7572: SET 7,B
label_7573: SET 0,C
label_7574: SET 1,C
label_7575: SET 2,C
label_7576: SET 3,C
label_7577: SET 4,C
label_7578: SET 5,C
label_7579: SET 6,C
label_7580: SET 7,C
label_7581: SET 0,D
label_7582: SET 1,D
label_7583: SET 2,D
label_7584: SET 3,D
label_7585: SET 4,D
label_7586: SET 5,D
label_7587: SET 6,D
label_7588: SET 7,D
label_7589: SET 0,E
label_7590: SET 1,E
label_7591: SET 2,E
label_7592: SET 3,E
label_7593: SET 4,E
label_7594: SET 5,E
label_7595: SET 6,E
label_7596: SET 7,E
label_7597: SET 0,H
label_7598: SET 1,H
label_7599: SET 2,H
label_7600: SET 3,H
label_7601: SET 4,H
label_7602: SET 5,H
label_7603: SET 6,H
label_7604: SET 7,H
label_7605: SET 0,L
label_7606: SET 1,L
label_7607: SET 2,L
label_7608: SET 3,L
label_7609: SET 4,L
label_7610: SET 5,L
label_7611: SET 6,L
label_7612: SET 7,L
label_7613: SET 0,(HL)
label_7614: SET 1,(HL)
label_7615: SET 2,(HL)
label_7616: SET 3,(HL)
label_7617: SET 4,(HL)
label_7618: SET 5,(HL)
label_7619: SET 6,(HL)
label_7620: SET 7,(HL)
label_7621: SET 0,(IX + 127)
label_7622: SET 1,(IX + 127)
label_7623: SET 2,(IX + 127)
label_7624: SET 3,(IX + 127)
label_7625: SET 4,(IX + 127)
label_7626: SET 5,(IX + 127)
label_7627: SET 6,(IX + 127)
label_7628: SET 7,(IX + 127)
label_7629: SET 0,(IY - 128)
label_7630: SET 1,(IY - 128)
label_7631: SET 2,(IY - 128)
label_7632: SET 3,(IY - 128)
label_7633: SET 4,(IY - 128)
label_7634: SET 5,(IY - 128)
label_7635: SET 6,(IY - 128)
label_7636: SET 7,(IY - 128)
label_7637: RES 0,A
label_7638: RES 1,A
label_7639: RES 2,A
label_7640: RES 3,A
label_7641: RES 4,A
label_7642: RES 5,A
label_7643: RES 6,A
label_7644: RES 7,A
label_7645: RES 0,B
label_7646: RES 1,B
label_7647: RES 2,B
label_7648: RES 3,B
label_7649: RES 4,B
label_7650: RES 5,B
label_7651: RES 6,B
label_7652: RES 7,B
label_7653: RES 0,C
label_7654: RES 1,C
label_7655: RES 2,C
label_7656: RES 3,C
label_7657: RES 4,C
label_7658: RES 5,C
label_7659: RES 6,C
label_7660: RES 7,C
label_7661: RES 0,D
label_7662: RES 1,D
label_7663: RES 2,D
label_7664: RES 3,D
label_7665: RES 4,D
label_7666: RES 5,D
label_7667: RES 6,D
label_7668: RES 7,D
label_7669: RES 0,E
label_7670: RES 1,E
label_7671: RES 2,E
label_7672: RES 3,E
label_7673: RES 4,E
label_7674: RES 5,E
label_7675: RES 6,E
label_7676: RES 7,E
label_7677: RES 0,H
label_7678: RES 1,H
label_7679: RES 2,H
label_7680: RES 3,H
label_7681: RES 4,H
label_7682: RES 5,H
label_7683: RES 6,H
label_7684: RES 7,H
label_7685: RES 0,L
label_7686: RES 1,L
label_7687: RES 2,L
label_7688: RES 3,L
label_7689: RES 4,L
label_7690: RES 5,L
label_7691: RES 6,L
label_7692: RES 7,L
label_7693: RES 0,(HL)
label_7694: RES 1,(HL)
label_7695: RES 2,(HL)
label_7696: RES 3,(HL)
label_7697: RES 4,(HL)
label_7698: RES 5,(HL)
label_7699: RES 6,(HL)
label_7700: RES 7,(HL)
label_7701: RES 0,(IX + 127)
label_7702: RES 1,(IX + 127)
label_7703: RES 2,(IX + 127)
label_7704: RES 3,(IX + 127)
label_7705: RES 4,(IX + 127)
label_7706: RES 5,(IX + 127)
label_7707: RES 6,(IX + 127)
label_7708: RES 7,(IX + 127)
label_7709: RES 0,(IY - 128)
label_7710: RES 1,(IY - 128)
label_7711: RES 2,(IY - 128)
label_7712: RES 3,(IY - 128)
label_7713: RES 4,(IY - 128)
label_7714: RES 5,(IY - 128)
label_7715: RES 6,(IY - 128)
label_7716: RES 7,(IY - 128)
label_7717: JP $5678
label_7718: JP NZ,$5678
label_7719: JP Z,$5678
label_7720: JP NC,$5678
label_7721: JP C,$5678
label_7722: JP PO,$5678
label_7723: JP PE,$5678
label_7724: JP P,$5678
label_7725: JP M,$5678
label_7726: JR $ + 2
label_7727: JR NZ,$ + 2
label_7728: JR Z,$ + 2
label_7729: JR NC,$ + 2
label_7730: JR C,$ + 2
label_7731: JP (HL)
label_7732: JP (IX)
label_7733: JP (IY)
label_7734: DJNZ $ + 2
label_7735: CALL $5678
label_7736: CALL NZ,$5678
label_7737: CALL Z,$5678
label_7738: CALL NC,$5678
label_7739: CALL C,$5678
label_7740: CALL PO,$5678
label_7741: CALL PE,$5678
label_7742: CALL P,$5678
label_7743: CALL M,$5678
label_7744: RET
label_7745: RET NZ
label_7746: RET Z
label_7747: RET NC
label_7748: RET C
label_7749: RET PO
label_7750: RET PE
label_7751: RET P
label_7752: RET M
label_7753: RETI
label_7754: RETN
label_7755: RST $00
label_7756: RST $08
label_7757: RST $10
label_7758: RST $18
label_7759: RST $20
label_7760: RST $28
label_7761: RST $30
label_7762: RST $38
label_7763: IN A,($12)
label_7764: IN A,(C)
label_7765: IN B,(C)
label_7766: IN C,(C)
label_7767: IN D,(C)
label_7768: IN E,(C)
label_7769: IN H,(C)
label_7770: IN L,(C)
label_7771: IN F,(C)
label_7772: INI
label_7773: INIR
label_7774: IND
label_7775: INDR
label_7776: OUT ($12),A
label_7777: OUT (C),A
label_7778: OUT (C),B
label_7779: OUT (C),C
label_7780: OUT (C),D
label_7781: OUT (C),E
label_7782: OUT (C),H
label_7783: OUT (C),L
label_7784: OUTI
label_7785: OTIR
label_7786: OUTD
label_7787: OTDR
label_7788: LD A,A
label_7789: LD A,B
label_7790: LD A,C
label_7791: LD A,D
label_7792: LD A,E
label_7793: LD A,H
label_7794: LD A,L
label_7795: LD B,A
label_7796: LD B,B
label_7797: LD B,C
label_7798: LD B,D
label_7799: LD B,E
label_7800: LD B,H
label_7801: LD B,L
label_7802: LD C,A
label_7803: LD C,B
label_7804: LD C,C
label_7805: LD C,D
label_7806: LD C,E
label_7807: LD C,H
label_7808: LD C,L
label_7809: LD D,A
label_7810: LD D,B
label_7811: LD D,C
label_7812: LD D,D
label_7813: LD D,E
label_7814: LD D,H
label_7815: LD D,L
label_7816: LD E,A
label_7817: LD E,B
label_7818: LD E,C
label_7819: LD E,D
label_7820: LD E,E
label_7821: LD E,H
label_7822: LD E,L
label_7823: LD H,A
label_7824: LD H,B
label_7825: LD H,C
label_7826: LD H,D
label_7827: LD H,E
label_7828: LD H,H
label_7829: LD H,L
label_7830: LD L,A
label_7831: LD L,B
label_7832: LD L,C
label_7833: LD L,D
label_7834: LD L,E
label_7835: LD L,H
label_7836: LD L,L
label_7837: LD A,$12
label_7838: LD B,$12
label_7839: LD C,$12
label_7840: LD D,$12
label_7841: LD E,$12
label_7842: LD H,$12
label_7843: LD L,$12
label_7844: LD A,(HL)
label_7845: LD B,(HL)
label_7846: LD C,(HL)
label_7847: LD D,(HL)
label_7848: LD E,(HL)
label_7849: LD H,(HL)
label_7850: LD L,(HL)
label_7851: LD A,(IX + 127)
label_7852: LD B,(IX + 127)
label_7853: LD C,(IX + 127)
label_7854: LD D,(IX + 127)
label_7855: LD E,(IX + 127)
label_7856: LD H,(IX + 127)
label_7857: LD L,(IX + 127)
label_7858: LD A,(IY - 128)
label_7859: LD B,(IY - 128)
label_7860: LD C,(IY - 128)
label_7861: LD D,(IY - 128)
label_7862: LD E,(IY - 128)
label_7863: LD H,(IY - 128)
label_7864: LD L,(IY - 128)
label_7865: LD (HL),A
label_7866: LD (HL),B
label_7867: LD (HL),C
label_7868: LD (HL),D
label_7869: LD (HL),E
label_7870: LD (HL),H
label_7871: LD (HL),L
label_7872: LD (IX + 127),A
label_7873: LD (IX + 127),B
label_7874: LD (IX + 127),C
label_7875: LD (IX + 127),D
label_7876: LD (IX + 127),E
label_7877: LD (IX + 127),H
label_7878: LD (IX + 127),L
label_7879: LD (IY - 128),A
label_7880: LD (IY - 128),B
label_7881: LD (IY - 128),C
label_7882: LD (IY - 128),D
label_7883: LD (IY - 128),E
label_7884: LD (IY - 128),H
label_7885: LD (IY - 128),L
label_7886: LD (HL),$12
label_7887: LD (IX + 127),$12
label_7888: LD (IY - 128),$12
label_7889: LD A,(BC)
label_7890: LD A,(DE)
label_7891: LD A,($5678)
label_7892: LD (BC),A
label_7893: LD (DE),A
label_7894: LD ($5678),A
label_7895: LD A,I
label_7896: LD A,R
label_7897: LD I,A
label_7898: LD R,A
label_7899: LD BC,$5678
label_7900: LD DE,$5678
label_7901: LD HL,$5678
label_7902: LD SP,$5678
label_7903: LD IX,$5678
label_7904: LD IY,$5678
label_7905: LD HL,($5678)
label_7906: LD BC,($5678)
label_7907: LD DE,($5678)
label_7908: LD HL,($5678)
label_7909: LD SP,($5678)
label_7910: LD IX,($5678)
label_7911: LD IY,($5678)
label_7912: LD ($5678),HL
label_7913: LD ($5678),BC
label_7914: LD ($5678),DE
label_7915: LD ($5678),HL
label_7916: LD ($5678),SP
label_7917: LD ($5678),IX
label_7918: LD ($5678),IY
label_7919: LD SP,HL
label_7920: LD SP,IX
label_7921: LD SP,IY
label_7922: PUSH BC
label_7923: PUSH DE
label_7924: PUSH HL
label_7925: PUSH AF
label_7926: PUSH IX
label_7927: PUSH IY
label_7928: POP BC
label_7929: POP DE
label_7930: POP HL
label_7931: POP AF
label_7932: POP IX
label_7933: POP IY
label_7934: EX DE,HL
label_7935: EX AF,AF'
label_7936: EXX
label_7937: EX (SP),HL
label_7938: EX (SP),IX
label_7939: EX (SP),IY
label_7940: LDI
label_7941: LDIR
label_7942: LDD
label_7943: LDDR
label_7944: CPI
label_7945: CPIR
label_7946: CPD
label_7947: CPDR
label_7948: ADD A,A
label_7949: ADD A,B
label_7950: ADD A,C
label_7951: ADD A,D
label_7952: ADD A,E
label_7953: ADD A,H
label_7954: ADD A,L
label_7955: ADD A,$12
label_7956: ADD A,(HL)
label_7957: ADD A,(IX + 127)
label_7958: ADD A,(IY - 128)
label_7959: ADC A,A
label_7960: ADC A,B
label_7961: ADC A,C
label_7962: ADC A,D
label_7963: ADC A,E
label_7964: ADC A,H
label_7965: ADC A,L
label_7966: ADC A,$12
label_7967: ADC A,(HL)
label_7968: ADC A,(IX + 127)
label_7969: ADC A,(IY - 128)
label_7970: SUB A
label_7971: SUB B
label_7972: SUB C
label_7973: SUB D
label_7974: SUB E
label_7975: SUB H
label_7976: SUB L
label_7977: SUB $12
label_7978: SUB (HL)
label_7979: SUB (IX + 127)
label_7980: SUB (IY - 128)
label_7981: SBC A,A
label_7982: SBC A,B
label_7983: SBC A,C
label_7984: SBC A,D
label_7985: SBC A,E
label_7986: SBC A,H
label_7987: SBC A,L
label_7988: SBC A,$12
label_7989: SBC A,(HL)
label_7990: SBC A,(IX + 127)
label_7991: SBC A,(IY - 128)
label_7992: AND A
label_7993: AND B
label_7994: AND C
label_7995: AND D
label_7996: AND E
label_7997: AND H
label_7998: AND L
label_7999: AND $12
label_8000: AND (HL)
label_8001: AND (IX + 127)
label_8002: AND (IY - 128)
label_8003: AND A
label_8004: AND B
label_8005: AND C
label_8006: AND D
label_8007: AND E
label_8008: AND H
label_8009: AND L
label_8010: AND $12
label_8011: AND (HL)
label_8012: AND (IX + 127)
label_8013: AND (IY - 128)
label_8014: OR A
label_8015: OR B
label_8016: OR C
label_8017: OR D
label_8018: OR E
label_8019: OR H
label_8020: OR L
label_8021: OR $12
label_8022: OR (HL)
label_8023: OR (IX + 127)
label_8024: OR (IY - 128)
label_8025: XOR A
label_8026: XOR B
label_8027: XOR C
label_8028: XOR D
label_8029: XOR E
label_8030: XOR H
label_8031: XOR L
label_8032: XOR $12
label_8033: XOR (HL)
label_8034: XOR (IX + 127)
label_8035: XOR (IY - 128)
label_8036: CP A
label_8037: CP B
label_8038: CP C
label_8039: CP D
label_8040: CP E
label_8041: CP H
label_8042: CP L
label_8043: CP $12
label_8044: CP (HL)
label_8045: CP (IX + 127)
label_8046: CP (IY - 128)
label_8047: INC A
label_8048: INC B
label_8049: INC C
label_8050: INC D
label_8051: INC E
label_8052: INC H
label_8053: INC L
label_8054: INC (HL)
label_8055: INC (IX + 127)
label_8056: INC (IY - 128)
label_8057: DEC A
label_8058: DEC B
label_8059: DEC C
label_8060: DEC D
label_8061: DEC E
label_8062: DEC H
label_8063: DEC L
label_8064: DEC (HL)
label_8065: DEC (IX + 127)
label_8066: DEC (IY - 128)
label_8067: DAA
label_8068: CPL
label_8069: NEG
label_8070: CCF
label_8071: SCF
label_8072: NOP
label_8073: HALT
label_8074: DI
label_8075: EI
label_8076: IM 0
label_8077: IM 1
label_8078: IM 2
label_8079: ADD HL,BC
label_8080: ADD HL,DE
label_8081: ADD HL,HL
label_8082: ADD HL,SP
label_8083: ADC HL,BC
label_8084: ADC HL,DE
label_8085: ADC HL,HL
label_8086: ADC HL,SP
label_8087: SBC HL,BC
label_8088: SBC HL,DE
label_8089: SBC HL,HL
label_8090: SBC HL,SP
label_8091: ADD IX,BC
label_8092: ADD IX,DE
label_8093: ADD IX,SP
label_8094: ADD IY,BC
label_8095: ADD IY,DE
label_8096: ADD IY,SP
label_8097: INC BC
label_8098: INC DE
label_8099: INC HL
label_8100: INC SP
label_8101: INC IX
label_8102: INC IY
label_8103: DEC BC
label_8104: DEC DE
label_8105: DEC HL
label_8106: DEC SP
label_8107: DEC IX
label_8108: DEC IY
label_8109: RLCA
label_8110: RLA
label_8111: RRCA
label_8112: RRA
label_8113: RLC A
label_8114: RLC B
label_8115: RLC C
label_8116: RLC D
label_8117: RLC E
label_8118: RLC H
label_8119: RLC L
label_8120: RLC (HL)
label_8121: RLC (IX + 127)
label_8122: RLC (IY - 128)
label_8123: RL A
label_8124: RL B
label_8125: RL C
label_8126: RL D
label_8127: RL E
label_8128: RL H
label_8129: RL L
label_8130: RL (HL)
label_8131: RL (IX + 127)
label_8132: RL (IY - 128)
label_8133: RRC A
label_8134: RRC B
label_8135: RRC C
label_8136: RRC D
label_8137: RRC E
label_8138: RRC H
label_8139: RRC L
label_8140: RRC (HL)
label_8141: RRC (IX + 127)
label_8142: RRC (IY - 128)
label_8143: RR A
label_8144: RR B
label_8145: RR C
label_8146: RR D
label_8147: RR E
label_8148: RR H
label_8149: RR L
label_8150: RR (HL)
label_8151: RR (IX + 127)
label_8152: RR (IY - 128)
label_8153: SLA A
label_8154: SLA B
label_8155: SLA C
label_8156: SLA D
label_8157: SLA E
label_8158: SLA H
label_8159: SLA L
label_8160: SLA (HL)
label_8161: SLA (IX + 127)
label_8162: SLA (IY - 128)
label_8163: SRA A
label_8164: SRA B
label_8165: SRA C
label_8166: SRA D
label_8167: SRA E
label_8168: SRA H
label_8169: SRA L
label_8170: SRA (HL)
label_8171: SRA (IX + 127)
label_8172: SRA (IY - 128)
label_8173: SRL A
label_8174: SRL B
label_8175: SRL C
label_8176: SRL D
label_8177: SRL E
label_8178: SRL H
label_8179: SRL L
label_8180: SRL (HL)
label_8181: SRL (IX + 127)
label_8182: SRL (IY - 128)
label_8183: RLD
label_8184: RRD
label_8185: BIT 0,A
label_8186: BIT 1,A
label_8187: BIT 2,A
label_8188: BIT 3,A
label_8189: BIT 4,A
label_8190: BIT 5,A
label_8191: BIT 6,A
label_8192: BIT 7,A
label_8193: BIT 0,B
label_8194: BIT 1,B
label_8195: BIT 2,B
label_8196: BIT 3,B
label_8197: BIT 4,B
label_8198: BIT 5,B
label_8199: BIT 6,B
label_8200: BIT 7,B
label_8201: BIT 0,C
label_8202: BIT 1,C
label_8203: BIT 2,C
label_8204: BIT 3,C
label_8205: BIT 4,C
label_8206: BIT 5,C
label_8207: BIT 6,C
label_8208: BIT 7,C
label_8209: BIT 0,D
label_8210: BIT 1,D
label_8211: BIT 2,D
label_8212: BIT 3,D
label_8213: BIT 4,D
label_8214: BIT 5,D
label_8215: BIT 6,D
label_8216: BIT 7,D
label_8217: BIT 0,E
label_8218: BIT 1,E
label_8219: BIT 2,E
label_8220: BIT 3,E
label_8221: BIT 4,E
label_8222: BIT 5,E
label_8223: BIT 6,E
label_8224: BIT 7,E
label_8225: BIT 0,H
label_8226: BIT 1,H
label_8227: BIT 2,H
label_8228: BIT 3,H
label_8229: BIT 4,H
label_8230: BIT 5,H
label_8231: BIT 6,H
label_8232: BIT 7,H
label_8233: BIT 0,L
label_8234: BIT 1,L
label_8235: BIT 2,L
label_8236: BIT 3,L
label_8237: BIT 4,L
label_8238: BIT 5,L
label_8239: BIT 6,L
label_8240: BIT 7,L
label_8241: BIT 0,(HL)
label_8242: BIT 1,(HL)
label_8243: BIT 2,(HL)
label_8244: BIT 3,(HL)
label_8245: BIT 4,(HL)
label_8246: BIT 5,(HL)
label_8247: BIT 6,(HL)
label_8248: BIT 7,(HL)
label_8249: BIT 0,(IX + 127)
label_8250: BIT 1,(IX + 127)
label_8251: BIT 2,(IX + 127)
label_8252: BIT 3,(IX + 127)
label_8253: BIT 4,(IX + 127)
label_8254: BIT 5,(IX + 127)
label_8255: BIT 6,(IX + 127)
label_8256: BIT 7,(IX + 127)
label_8257: BIT 0,(IY - 128)
label_8258: BIT 1,(IY - 128)
label_8259: BIT 2,(IY - 128)
label_8260: BIT 3,(IY - 128)
label_8261: BIT 4,(IY - 128)
label_8262: BIT 5,(IY - 128)
label_8263: BIT 6,(IY - 128)
label_8264: BIT 7,(IY - 128)
label_8265: SET 0,A
label_8266: SET 1,A
label_8267: SET 2,A
label_8268: SET 3,A
label_8269: SET 4,A
label_8270: SET 5,A
label_8271: SET 6,A
label_8272: SET 7,A
label_8273: SET 0,B
label_8274: SET 1,B
label_8275: SET 2,B
label_8276: SET 3,B
label_8277: SET 4,B
label_8278: SET 5,B
label_8279: SET 6,B
label_8280: SET 7,B
label_8281: SET 0,C
label_8282: SET 1,C
label_8283: SET 2,C
label_8284: SET 3,C
label_8285: SET 4,C
label_8286: SET 5,C
label_8287: SET 6,C
label_8288: SET 7,C
label_8289: SET 0,D
label_8290: SET 1,D
label_8291: SET 2,D
label_8292: SET 3,D
label_8293: SET 4,D
label_8294: SET 5,D
label_8295: SET 6,D
label_8296: SET 7,D
label_8297: SET 0,E
label_8298: SET 1,E
label_8299: SET 2,E
label_8300: SET 3,E
label_8301: SET 4,E
label_8302: SET 5,E
label_8303: SET 6,E
label_8304: SET 7,E
label_8305: SET 0,H
label_8306: SET 1,H
label_8307: SET 2,H
label_8308: SET 3,H
label_8309: SET 4,H
label_8310: SET 5,H
label_8311: SET 6,H
label_8312: SET 7,H
label_8313: SET 0,L
label_8314: SET 1,L
label_8315: SET 2,L
label_8316: SET 3,L
label_8317: SET 4,L
label_8318: SET 5,L
label_8319: SET 6,L
label_8320: SET 7,L
label_8321: SET 0,(HL)
label_8322: SET 1,(HL)
label_8323: SET 2,(HL)
label_8324: SET 3,(HL)
label_8325: SET 4,(HL)
label_8326: SET 5,(HL)
label_8327: SET 6,(HL)
label_8328: SET 7,(HL)
label_8329: SET 0,(IX + 127)
label_8330: SET 1,(IX + 127)
label_8331: SET 2,(IX + 127)
label_8332: SET 3,(IX + 127)
label_8333: SET 4,(IX + 127)
label_8334: SET 5,(IX + 127)
label_8335: SET 6,(IX + 127)
label_8336: SET 7,(IX + 127)
label_8337: SET 0,(IY - 128)
label_8338: SET 1,(IY - 128)
label_8339: SET 2,(IY - 128)
label_8340: SET 3,(IY - 128)
label_8341: SET 4,(IY - 128)
label_8342: SET 5,(IY - 128)
label_8343: SET 6,(IY - 128)
label_8344: SET 7,(IY - 128)
label_8345: RES 0,A
label_8346: RES 1,A
label_8347: RES 2,A
label_8348: RES 3,A
label_8349: RES 4,A
label_8350: RES 5,A
label_8351: RES 6,A
label_8352: RES 7,A
label_8353: RES 0,B
label_8354: RES 1,B
label_8355: RES 2,B
label_8356: RES 3,B
label_8357: RES 4,B
label_8358: RES 5,B
label_8359: RES 6,B
label_8360: RES 7,B
label_8361: RES 0,C
label_8362: RES 1,C
label_8363: RES 2,C
label_8364: RES 3,C
label_8365: RES 4,C
label_8366: RES 5,C
label_8367: RES 6,C
label_8368: RES 7,C
label_8369: RES 0,D
label_8370: RES 1,D
label_8371: RES 2,D
label_8372: RES 3,D
label_8373: RES 4,D
label_8374: RES 5,D
label_8375: RES 6,D
label_8376: RES 7,D
label_8377: RES 0,E
label_8378: RES 1,E
label_8379: RES 2,E
label_8380: RES 3,E
label_8381: RES 4,E
label_8382: RES 5,E
label_8383: RES 6,E
label_8384: RES 7,E
label_8385: RES 0,H
label_8386: RES 1,H
label_8387: RES 2,H
label_8388: RES 3,H
label_8389: RES 4,H
label_8390: RES 5,H
label_8391: RES 6,H
label_8392: RES 7,H
label_8393: RES 0,L
label_8394: RES 1,L
label_8395: RES 2,L
label_8396: RES 3,L
label_8397: RES 4,L
label_8398: RES 5,L
label_8399: RES 6,L
label_8400: RES 7,L
label_8401: RES 0,(HL)
label_8402: RES 1,(HL)
label_8403: RES 2,(HL)
label_8404: RES 3,(HL)
label_8405: RES 4,(HL)
label_8406: RES 5,(HL)
label_8407: RES 6,(HL)
label_8408: RES 7,(HL)
label_8409: RES 0,(IX + 127)
label_8410: RES 1,(IX + 127)
label_8411: RES 2,(IX + 127)
label_8412: RES 3,(IX + 127)
label_8413: RES 4,(IX + 127)
label_8414: RES 5,(IX + 127)
label_8415: RES 6,(IX + 127)
label_8416: RES 7,(IX + 127)
label_8417: RES 0,(IY - 128)
label_8418: RES 1,(IY - 128)
label_8419: RES 2,(IY - 128)
label_8420: RES 3,(IY - 128)
label_8421: RES 4,(IY - 128)
label_8422: RES 5,(IY - 128)
label_8423: RES 6,(IY - 128)
label_8424: RES 7,(IY - 128)
label_8425: JP $5678
label_8426: JP NZ,$5678
label_8427: JP Z,$5678
label_8428: JP NC,$5678
label_8429: JP C,$5678
label_8430: JP PO,$5678
label_8431: JP PE,$5678
label_8432: JP P,$5678
label_8433: JP M,$5678
label_8434: JR $ + 2
label_8435: JR NZ,$ + 2
label_8436: JR Z,$ + 2
label_8437: JR NC,$ + 2
label_8438: JR C,$ + 2
label_8439: JP (HL)
label_8440: JP (IX)
label_8441: JP (IY)
label_8442: DJNZ $ + 2
label_8443: CALL $5678
label_8444: CALL NZ,$5678
label_8445: CALL Z,$5678
label_8446: CALL NC,$5678
label_8447: CALL C,$5678
label_8448: CALL PO,$5678
label_8449: CALL PE,$5678
label_8450: CALL P,$5678
label_8451: CALL M,$5678
label_8452: RET
label_8453: RET NZ
label_8454: RET Z
label_8455: RET NC
label_8456: RET C
label_8457: RET PO
label_8458: RET PE
label_8459: RET P
label_8460: RET M
label_8461: RETI
label_8462: RETN
label_8463: RST $00
label_8464: RST $08
label_8465: RST $10
label_8466: RST $18
label_8467: RST $20
label_8468: RST $28
label_8469: RST $30
label_8470: RST $38
label_8471: IN A,($12)
label_8472: IN A,(C)
label_8473: IN B,(C)
label_8474: IN C,(C)
label_8475: IN D,(C)
label_8476: IN E,(C)
label_8477: IN H,(C)
label_8478: IN L,(C)
label_8479: IN F,(C)
label_8480: INI
label_8481: INIR
label_8482: IND
label_8483: INDR
label_8484: OUT ($12),A
label_8485: OUT (C),A
label_8486: OUT (C),B
label_8487: OUT (C),C
label_8488: OUT (C),D
label_8489: OUT (C),E
label_8490: OUT (C),H
label_8491: OUT (C),L
label_8492: OUTI
label_8493: OTIR
label_8494: OUTD
label_8495: OTDR
label_8496: LD A,A
label_8497: LD A,B
label_8498: LD A,C
label_8499: LD A,D
label_8500: LD A,E
label_8501: LD A,H
label_8502: LD A,L
label_8503: LD B,A
label_8504: LD B,B
label_8505: LD B,C
label_8506: LD B,D
label_8507: LD B,E
label_8508: LD B,H
label_8509: LD B,L
label_8510: LD C,A
label_8511: LD C,B
label_8512: LD C,C
label_8513: LD C,D
label_8514: LD C,E
label_8515: LD C,H
label_8516: LD C,L
label_8517: LD D,A
label_8518: LD D,B
label_8519: LD D,C
label_8520: LD D,D
label_8521: LD D,E
label_8522: LD D,H
label_8523: LD D,L
label_8524: LD E,A
label_8525: LD E,B
label_8526: LD E,C
label_8527: LD E,D
label_8528: LD E,E
label_8529: LD E,H
label_8530: LD E,L
label_8531: LD H,A
label_8532: LD H,B
label_8533: LD H,C
label_8534: LD H,D
label_8535: LD H,E
label_8536: LD H,H
label_8537: LD H,L
label_8538: LD L,A
label_8539: LD L,B
label_8540: LD L,C
label_8541: LD L,D
label_8542: LD L,E
label_8543: LD L,H
label_8544: LD L,L
label_8545: LD A,$12
label_8546: LD B,$12
label_8547: LD C,$12
label_8548: LD D,$12
label_8549: LD E,$12
label_8550: LD H,$12
label_8551: LD L,$12
label_8552: LD A,(HL)
label_8553: LD B,(HL)
label_8554: LD C,(HL)
label_8555: LD D,(HL)
label_8556: LD E,(HL)
label_8557: LD H,(HL)
label_8558: LD L,(HL)
label_8559: LD A,(IX + 127)
label_8560: LD B,(IX + 127)
label_8561: LD C,(IX + 127)
label_8562: LD D,(IX + 127)
label_8563: LD E,(IX + 127)
label_8564: LD H,(IX + 127)
label_8565: LD L,(IX + 127)
label_8566: LD A,(IY - 128)
label_8567: LD B,(IY - 128)
label_8568: LD C,(IY - 128)
label_8569: LD D,(IY - 128)
label_8570: LD E,(IY - 128)
label_8571: LD H,(IY - 128)
label_8572: LD L,(IY - 128)
label_8573: LD (HL),A
label_8574: LD (HL),B
label_8575: LD (HL),C
label_8576: LD (HL),D
label_8577: LD (HL),E
label_8578: LD (HL),H
label_8579: LD (HL),L
label_8580: LD (IX + 127),A
label_8581: LD (IX + 127),B
label_8582: LD (IX + 127),C
label_8583: LD (IX + 127),D
label_8584: LD (IX + 127),E
label_8585: LD (IX + 127),H
label_8586: LD (IX + 127),L
label_8587: LD (IY - 128),A
label_8588: LD (IY - 128),B
label_8589: LD (IY - 128),C
label_8590: LD (IY - 128),D
label_8591: LD (IY - 128),E
label_8592: LD (IY - 128),H
label_8593: LD (IY - 128),L
label_8594: LD (HL),$12
label_8595: LD (IX + 127),$12
label_8596: LD (IY - 128),$12
label_8597: LD A,(BC)
label_8598: LD A,(DE)
label_8599: LD A,($5678)
label_8600: LD (BC),A
label_8601: LD (DE),A
label_8602: LD ($5678),A
label_8603: LD A,I
label_8604: LD A,R
label_8605: LD I,A
label_8606: LD R,A
label_8607: LD BC,$5678
label_8608: LD DE,$5678
label_8609: LD HL,$5678
label_8610: LD SP,$5678
label_8611: LD IX,$5678
label_8612: LD IY,$5678
label_8613: LD HL,($5678)
label_8614: LD BC,($5678)
label_8615: LD DE,($5678)
label_8616: LD HL,($5678)
label_8617: LD SP,($5678)
label_8618: LD IX,($5678)
label_8619: LD IY,($5678)
label_8620: LD ($5678),HL
label_8621: LD ($5678),BC
label_8622: LD ($5678),DE
label_8623: LD ($5678),HL
label_8624: LD ($5678),SP
label_8625: LD ($5678),IX
label_8626: LD ($5678),IY
label_8627: LD SP,HL
label_8628: LD SP,IX
label_8629: LD SP,IY
label_8630: PUSH BC
label_8631: PUSH DE
label_8632: PUSH HL
label_8633: PUSH AF
label_8634: PUSH IX
label_8635: PUSH IY
label_8636: POP BC
label_8637: POP DE
label_8638: POP HL
label_8639: POP AF
label_8640: POP IX
label_8641: POP IY
label_8642: EX DE,HL
label_8643: EX AF,AF'
label_8644: EXX
label_8645: EX (SP),HL
label_8646: EX (SP),IX
label_8647: EX (SP),IY
label_8648: LDI
label_8649: LDIR
label_8650: LDD
label_8651: LDDR
label_8652: CPI
label_8653: CPIR
label_8654: CPD
label_8655: CPDR
label_8656: ADD A,A
label_8657: ADD A,B
label_8658: ADD A,C
label_8659: ADD A,D
label_8660: ADD A,E
label_8661: ADD A,H
label_8662: ADD A,L
label_8663: ADD A,$12
label_8664: ADD A,(HL)
label_8665: ADD A,(IX + 127)
label_8666: ADD A,(IY - 128)
label_8667: ADC A,A
label_8668: ADC A,B
label_8669: ADC A,C
label_8670: ADC A,D
label_8671: ADC A,E
label_8672: ADC A,H
label_8673: ADC A,L
label_8674: ADC A,$12
label_8675: ADC A,(HL)
label_8676: ADC A,(IX + 127)
label_8677: ADC A,(IY - 128)
label_8678: SUB A
label_8679: SUB B
label_8680: SUB C
label_8681: SUB D
label_8682: SUB E
label_8683: SUB H
label_8684: SUB L
label_8685: SUB $12
label_8686: SUB (HL)
label_8687: SUB (IX + 127)
label_8688: SUB (IY - 128)
label_8689: SBC A,A
label_8690: SBC A,B
label_8691: SBC A,C
label_8692: SBC A,D
label_8693: SBC A,E
label_8694: SBC A,H
label_8695: SBC A,L
label_8696: SBC A,$12
label_8697: SBC A,(HL)
label_8698: SBC A,(IX + 127)
label_8699: SBC A,(IY - 128)
label_8700: AND A
label_8701: AND B
label_8702: AND C
label_8703: AND D
label_8704: AND E
label_8705: AND H
label_8706: AND L
label_8707: AND $12
label_8708: AND (HL)
label_8709: AND (IX + 127)
label_8710: AND (IY - 128)
label_8711: AND A
label_8712: AND B
label_8713: AND C
label_8714: AND D
label_8715: AND E
label_8716: AND H
label_8717: AND L
label_8718: AND $12
label_8719: AND (HL)
label_8720: AND (IX + 127)
label_8721: AND (IY - 128)
label_8722: OR A
label_8723: OR B
label_8724: OR C
label_8725: OR D
label_8726: OR E
label_8727: OR H
label_8728: OR L
label_8729: OR $12
label_8730: OR (HL)
label_8731: OR (IX + 127)
label_8732: OR (IY - 128)
label_8733: XOR A
label_8734: XOR B
label_8735: XOR C
label_8736: XOR D
label_8737: XOR E
label_8738: XOR H
label_8739: XOR L
label_8740: XOR $12
label_8741: XOR (HL)
label_8742: XOR (IX + 127)
label_8743: XOR (IY - 128)
label_8744: CP A
label_8745: CP B
label_8746: CP C
label_8747: CP D
label_8748: CP E
label_8749: CP H
label_8750: CP L
label_8751: CP $12
label_8752: CP (HL)
label_8753: CP (IX + 127)
label_8754: CP (IY - 128)
label_8755: INC A
label_8756: INC B
label_8757: INC C
label_8758: INC D
label_8759: INC E
label_8760: INC H
label_8761: INC L
label_8762: INC (HL)
label_8763: INC (IX + 127)
label_8764: INC (IY - 128)
label_8765: DEC A
label_8766: DEC B
label_8767: DEC C
label_8768: DEC D
label_8769: DEC E
label_8770: DEC H
label_8771: DEC L
label_8772: DEC (HL)
label_8773: DEC (IX + 127)
label_8774: DEC (IY - 128)
label_8775: DAA
label_8776: CPL
label_8777: NEG
label_8778: CCF
label_8779: SCF
label_8780: NOP
label_8781: HALT
label_8782: DI
label_8783: EI
label_8784: IM 0
label_8785: IM 1
label_8786: IM 2
label_8787: ADD HL,BC
label_8788: ADD HL,DE
label_8789: ADD HL,HL
label_8790: ADD HL,SP
label_8791: ADC HL,BC
label_8792: ADC HL,DE
label_8793: ADC HL,HL
label_8794: ADC HL,SP
label_8795: SBC HL,BC
label_8796: SBC HL,DE
label_8797: SBC HL,HL
label_8798: SBC HL,SP
label_8799: ADD IX,BC
label_8800: ADD IX,DE
label_8801: ADD IX,SP
label_8802: ADD IY,BC
label_8803: ADD IY,DE
label_8804: ADD IY,SP
label_8805: INC BC
label_8806: INC DE
label_8807: INC HL
label_8808: INC SP
label_8809: INC IX
label_8810: INC IY
label_8811: DEC BC
label_8812: DEC DE
label_8813: DEC HL
label_8814: DEC SP
label_8815: DEC IX
label_8816: DEC IY
label_8817: RLCA
label_8818: RLA
label_8819: RRCA
label_8820: RRA
label_8821: RLC A
label_8822: RLC B
label_8823: RLC C
label_8824: RLC D
label_8825: RLC E
label_8826: RLC H
label_8827: RLC L
label_8828: RLC (HL)
label_8829: RLC (IX + 127)
label_8830: RLC (IY - 128)
label_8831: RL A
label_8832: RL B
label_8833: RL C
label_8834: RL D
label_8835: RL E
label_8836: RL H
label_8837: RL L
label_8838: RL (HL)
label_8839: RL (IX + 127)
label_8840: RL (IY - 128)
label_8841: RRC A
label_8842: RRC B
label_8843: RRC C
label_8844: RRC D
label_8845: RRC E
label_8846: RRC H
label_8847: RRC L
label_8848: RRC (HL)
label_8849: RRC (IX + 127)
label_8850: RRC (IY - 128)
label_8851: RR A
label_8852: RR B
label_8853: RR C
label_8854: RR D
label_8855: RR E
label_8856: RR H
label_8857: RR L
label_8858: RR (HL)
label_8859: RR (IX + 127)
label_8860: RR (IY - 128)
label_8861: SLA A
label_8862: SLA B
label_8863: SLA C
label_8864: SLA D
label_8865: SLA E
label_8866: SLA H
label_8867: SLA L
label_8868: SLA (HL)
label_8869: SLA (IX + 127)
label_8870: SLA (IY - 128)
label_8871: SRA A
label_8872: SRA B
label_8873: SRA C
label_8874: SRA D
label_8875: SRA E
label_8876: SRA H
label_8877: SRA L
label_8878: SRA (HL)
label_8879: SRA (IX + 127)
label_8880: SRA (IY - 128)
label_8881: SRL A
label_8882: SRL B
label_8883: SRL C
label_8884: SRL D
label_8885: SRL E
label_8886: SRL H
label_8887: SRL L
label_8888: SRL (HL)
label_8889: SRL (IX + 127)
label_8890: SRL (IY - 128)
label_8891: RLD
label_8892: RRD
label_8893: BIT 0,A
label_8894: BIT 1,A
label_8895: BIT 2,A
label_8896: BIT 3,A
label_8897: BIT 4,A
label_8898: BIT 5,A
label_8899: BIT 6,A
label_8900: BIT 7,A
label_8901: BIT 0,B
label_8902: BIT 1,B
label_8903: BIT 2,B
label_8904: BIT 3,B
label_8905: BIT 4,B
label_8906: BIT 5,B
label_8907: BIT 6,B
label_8908: BIT 7,B
label_8909: BIT 0,C
label_8910: BIT 1,C
label_8911: BIT 2,C
label_8912: BIT 3,C
label_8913: BIT 4,C
label_8914: BIT 5,C
label_8915: BIT 6,C
label_8916: BIT 7,C
label_8917: BIT 0,D
label_8918: BIT 1,D
label_8919: BIT 2,D
label_8920: BIT 3,D
label_8921: BIT 4,D
label_8922: BIT 5,D
label_8923: BIT 6,D
label_8924: BIT 7,D
label_8925: BIT 0,E
label_8926: BIT 1,E
label_8927: BIT 2,E
label_8928: BIT 3,E
label_8929: BIT 4,E
label_8930: BIT 5,E
label_8931: BIT 6,E
label_8932: BIT 7,E
label_8933: BIT 0,H
label_8934: BIT 1,H
label_8935: BIT 2,H
label_8936: BIT 3,H
label_8937: BIT 4,H
label_8938: BIT 5,H
label_8939: BIT 6,H
label_8940: BIT 7,H
label_8941: BIT 0,L
label_8942: BIT 1,L
label_8943: BIT 2,L
label_8944: BIT 3,L
label_8945: BIT 4,L
label_8946: BIT 5,L
label_8947: BIT 6,L
label_8948: BIT 7,L
label_8949: BIT 0,(HL)
label_8950: BIT 1,(HL)
label_8951: BIT 2,(HL)
label_8952: BIT 3,(HL)
label_8953: BIT 4,(HL)
label_8954: BIT 5,(HL)
label_8955: BIT 6,(HL)
label_8956: BIT 7,(HL)
label_8957: BIT 0,(IX + 127)
label_8958: BIT 1,(IX + 127)
label_8959: BIT 2,(IX + 127)
label_8960: BIT 3,(IX + 127)
label_8961: BIT 4,(IX + 127)
label_8962: BIT 5,(IX + 127)
label_8963: BIT 6,(IX + 127)
label_8964: BIT 7,(IX + 127)
label_8965: BIT 0,(IY - 128)
label_8966: BIT 1,(IY - 128)
label_8967: BIT 2,(IY - 128)
label_8968: BIT 3,(IY - 128)
label_8969: BIT 4,(IY - 128)
label_8970: BIT 5,(IY - 128)
label_8971: BIT 6,(IY - 128)
label_8972: BIT 7,(IY - 128)
label_8973: SET 0,A
label_8974: SET 1,A
label_8975: SET 2,A
label_8976: SET 3,A
label_8977: SET 4,A
label_8978: SET 5,A
label_8979: SET 6,A
label_8980: SET 7,A
label_8981: SET 0,B
label_8982: SET 1,B
label_8983: SET 2,B
label_8984: SET 3,B
label_8985: SET 4,B
label_8986: SET 5,B
label_8987: SET 6,B
label_8988: SET 7,B
label_8989: SET 0,C
label_8990: SET 1,C
label_8991: SET 2,C
label_8992: SET 3,C
label_8993: SET 4,C
label_8994: SET 5,C
label_8995: SET 6,C
label_8996: SET 7,C
label_8997: SET 0,D
label_8998: SET 1,D
label_8999: SET 2,D
label_9000: SET 3,D
label_9001: SET 4,D
label_9002: SET 5,D
label_9003: SET 6,D
label_9004: SET 7,D
label_9005: SET 0,E
label_9006: SET 1,E
label_9007: SET 2,E
label_9008: SET 3,E
label_9009: SET 4,E
label_9010: SET 5,E
label_9011: SET 6,E
label_9012: SET 7,E
label_9013: SET 0,H
label_9014: SET 1,H
label_9015: SET 2,H
label_9016: SET 3,H
label_9017: SET 4,H
label_9018: SET 5,H
label_9019: SET 6,H
label_9020: SET 7,H
label_9021: SET 0,L
label_9022: SET 1,L
label_9023: SET 2,L
label_9024: SET 3,L
label_9025: SET 4,L
label_9026: SET 5,L
label_9027: SET 6,L
label_9028: SET 7,L
label_9029: SET 0,(HL)
label_9030: SET 1,(HL)
label_9031: SET 2,(HL)
label_9032: SET 3,(HL)
label_9033: SET 4,(HL)
label_9034: SET 5,(HL)
label_9035: SET 6,(HL)
label_9036: SET 7,(HL)
label_9037: SET 0,(IX + 127)
label_9038: SET 1,(IX + 127)
label_9039: SET 2,(IX + 127)
label_9040: SET 3,(IX + 127)
label_9041: SET 4,(IX + 127)
label_9042: SET 5,(IX + 127)
label_9043: SET 6,(IX + 127)
label_9044: SET 7,(IX + 127)
label_9045: SET 0,(IY - 128)
label_9046: SET 1,(IY - 128)
label_9047: SET 2,(IY - 128)
label_9048: SET 3,(IY - 128)
label_9049: SET 4,(IY - 128)
label_9050: SET 5,(IY - 128)
label_9051: SET 6,(IY - 128)
label_9052: SET 7,(IY - 128)
label_9053: RES 0,A
label_9054: RES 1,A
label_9055: RES 2,A
label_9056: RES 3,A
label_9057: RES 4,A
label_9058: RES 5,A
label_9059: RES 6,A
label_9060: RES 7,A
label_9061: RES 0,B
label_9062: RES 1,B
label_9063: RES 2,B
label_9064: RES 3,B
label_9065: RES 4,B
label_9066: RES 5,B
label_9067: RES 6,B
label_9068: RES 7,B
label_9069: RES 0,C
label_9070: RES 1,C
label_9071: RES 2,C
label_9072: RES 3,C
label_9073: RES 4,C
label_9074: RES 5,C
label_9075: RES 6,C
label_9076: RES 7,C
label_9077: RES 0,D
label_9078: RES 1,D
label_9079: RES 2,D
label_9080: RES 3,D
label_9081: RES 4,D
label_9082: RES 5,D
label_9083: RES 6,D
label_9084: RES 7,D
label_9085: RES 0,E
label_9086: RES 1,E
label_9087: RES 2,E
label_9088: RES 3,E
label_9089: RES 4,E
label_9090: RES 5,E
label_9091: RES 6,E
label_9092: RES 7,E
label_9093: RES 0,H
label_9094: RES 1,H
label_9095: RES 2,H
label_9096: RES 3,H
label_9097: RES 4,H
label_9098: RES 5,H
label_9099: RES 6,H
label_9100: RES 7,H
label_9101: RES 0,L
label_9102: RES 1,L
label_9103: RES 2,L
label_9104: RES 3,L
label_9105: RES 4,L
label_9106: RES 5,L
label_9107: RES 6,L
label_9108: RES 7,L
label_9109: RES 0,(HL)
label_9110: RES 1,(HL)
label_9111: RES 2,(HL)
label_9112: RES 3,(HL)
label_9113: RES 4,(HL)
label_9114: RES 5,(HL)
label_9115: RES 6,(HL)
label_9116: RES 7,(HL)
label_9117: RES 0,(IX + 127)
label_9118: RES 1,(IX + 127)
label_9119: RES 2,(IX + 127)
label_9120: RES 3,(IX + 127)
label_9121: RES 4,(IX + 127)
label_9122: RES 5,(IX + 127)
label_9123: RES 6,(IX + 127)
label_9124: RES 7,(IX + 127)
label_9125: RES 0,(IY - 128)
label_9126: RES 1,(IY - 128)
label_9127: RES 2,(IY - 128)
label_9128: RES 3,(IY - 128)
label_9129: RES 4,(IY - 128)
label_9130: RES 5,(IY - 128)
label_9131: RES 6,(IY - 128)
label_9132: RES 7,(IY - 128)
label_9133: JP $5678
label_9134: JP NZ,$5678
label_9135: JP Z,$5678
label_9136: JP NC,$5678
label_9137: JP C,$5678
label_9138: JP PO,$5678
label_9139: JP PE,$5678
label_9140: JP P,$5678
label_9141: JP M,$5678
label_9142: JR $ + 2
label_9143: JR NZ,$ + 2
label_9144: JR Z,$ + 2
label_9145: JR NC,$ + 2
label_9146: JR C,$ + 2
label_9147: JP (HL)
label_9148: JP (IX)
label_9149: JP (IY)
label_9150: DJNZ $ + 2
label_9151: CALL $5678
label_9152: CALL NZ,$5678
label_9153: CALL Z,$5678
label_9154: CALL NC,$5678
label_9155: CALL C,$5678
label_9156: CALL PO,$5678
label_9157: CALL PE,$5678
label_9158: CALL P,$5678
label_9159: CALL M,$5678
label_9160: RET
label_9161: RET NZ
label_9162: RET Z
label_9163: RET NC
label_9164: RET C
label_9165: RET PO
label_9166: RET PE
label_9167: RET P
label_9168: RET M
label_9169: RETI
label_9170: RETN
label_9171: RST $00
label_9172: RST $08
label_9173: RST $10
label_9174: RST $18
label_9175: RST $20
label_9176: RST $28
label_9177: RST $30
label_9178: RST $38
label_9179: IN A,($12)
label_9180: IN A,(C)
label_9181: IN B,(C)
label_9182: IN C,(C)
label_9183: IN D,(C)
label_9184: IN E,(C)
label_9185: IN H,(C)
label_9186: IN L,(C)
label_9187: IN F,(C)
label_9188: INI
label_9189: INIR
label_9190: IND
label_9191: INDR
label_9192: OUT ($12),A
label_9193: OUT (C),A
label_9194: OUT (C),B
label_9195: OUT (C),C
label_9196: OUT (C),D
label_9197: OUT (C),E
label_9198: OUT (C),H
label_9199: OUT (C),L
label_9200: OUTI
label_9201: OTIR
label_9202: OUTD
label_9203: OTDR
label_9204: LD A,A
label_9205: LD A,B
label_9206: LD A,C
label_9207: LD A,D
label_9208: LD A,E
label_9209: LD A,H
label_9210: LD A,L
label_9211: LD B,A
label_9212: LD B,B
label_9213: LD B,C
label_9214: LD B,D
label_9215: LD B,E
label_9216: LD B,H
label_9217: LD B,L
label_9218: LD C,A
label_9219: LD C,B
label_9220: LD C,C
label_9221: LD C,D
label_9222: LD C,E
label_9223: LD C,H
label_9224: LD C,L
label_9225: LD D,A
label_9226: LD D,B
label_9227: LD D,C
label_9228: LD D,D
label_9229: LD D,E
label_9230: LD D,H
label_9231: LD D,L
label_9232: LD E,A
label_9233: LD E,B
label_9234: LD E,C
label_9235: LD E,D
label_9236: LD E,E
label_9237: LD E,H
label_9238: LD E,L
label_9239: LD H,A
label_9240: LD H,B
label_9241: LD H,C
label_9242: LD H,D
label_9243: LD H,E
label_9244: LD H,H
label_9245: LD H,L
label_9246: LD L,A
label_9247: LD L,B
label_9248: LD L,C
label_9249: LD L,D
label_9250: LD L,E
label_9251: LD L,H
label_9252: LD L,L
label_9253: LD A,$12
label_9254: LD B,$12
label_9255: LD C,$12
label_9256: LD D,$12
label_9257: LD E,$12
label_9258: LD H,$12
label_9259: LD L,$12
label_9260: LD A,(HL)
label_9261: LD B,(HL)
label_9262: LD C,(HL)
label_9263: LD D,(HL)
label_9264: LD E,(HL)
label_9265: LD H,(HL)
label_9266: LD L,(HL)
label_9267: LD A,(IX + 127)
label_9268: LD B,(IX + 127)
label_9269: LD C,(IX + 127)
label_9270: LD D,(IX + 127)
label_9271: LD E,(IX + 127)
label_9272: LD H,(IX + 127)
label_9273: LD L,(IX + 127)
label_9274: LD A,(IY - 128)
label_9275: LD B,(IY - 128)
label_9276: LD C,(IY - 128)
label_9277: LD D,(IY - 128)
label_9278: LD E,(IY - 128)
label_9279: LD H,(IY - 128)
label_9280: LD L,(IY - 128)
label_9281: LD (HL),A
label_9282: LD (HL),B
label_9283: LD (HL),C
label_9284: LD (HL),D
label_9285: LD (HL),E
label_9286: LD (HL),H
label_9287: LD (HL),L
label_9288: LD (IX + 127),A
label_9289: LD (IX + 127),B
label_9290: LD (IX + 127),C
label_9291: LD (IX + 127),D
label_9292: LD (IX + 127),E
label_9293: LD (IX + 127),H
label_9294: LD (IX + 127),L
label_9295: LD (IY - 128),A
label_9296: LD (IY - 128),B
label_9297: LD (IY - 128),C
label_9298: LD (IY - 128),D
label_9299: LD (IY - 128),E
label_9300: LD (IY - 128),H
label_9301: LD (IY - 128),L
label_9302: LD (HL),$12
label_9303: LD (IX + 127),$12
label_9304: LD (IY - 128),$12
label_9305: LD A,(BC)
label_9306: LD A,(DE)
label_9307: LD A,($5678)
label_9308: LD (BC),A
label_9309: LD (DE),A
label_9310: LD ($5678),A
label_9311: LD A,I
label_9312: LD A,R
label_9313: LD I,A
label_9314: LD R,A
label_9315: LD BC,$5678
label_9316: LD DE,$5678
label_9317: LD HL,$5678
label_9318: LD SP,$5678
label_9319: LD IX,$5678
label_9320: LD IY,$5678
label_9321: LD HL,($5678)
label_9322: LD BC,($5678)
label_9323: LD DE,($5678)
label_9324: LD HL,($5678)
label_9325: LD SP,($5678)
label_9326: LD IX,($5678)
label_9327: LD IY,($5678)
label_9328: LD ($5678),HL
label_9329: LD ($5678),BC
label_9330: LD ($5678),DE
label_9331: LD ($5678),HL
label_9332: LD ($5678),SP
label_9333: LD ($5678),IX
label_9334: LD ($5678),IY
label_9335: LD SP,HL
label_9336: LD SP,IX
label_9337: LD SP,IY
label_9338: PUSH BC
label_9339: PUSH DE
label_9340: PUSH HL
label_9341: PUSH AF
label_9342: PUSH IX
label_9343: PUSH IY
label_9344: POP BC
label_9345: POP DE
label_9346: POP HL
label_9347: POP AF
label_9348: POP IX
label_9349: POP IY
label_9350: EX DE,HL
label_9351: EX AF,AF'
label_9352: EXX
label_9353: EX (SP),HL
label_9354: EX (SP),IX
label_9355: EX (SP),IY
label_9356: LDI
label_9357: LDIR
label_9358: LDD
label_9359: LDDR
label_9360: CPI
label_9361: CPIR
label_9362: CPD
label_9363: CPDR
label_9364: ADD A,A
label_9365: ADD A,B
label_9366: ADD A,C
label_9367: ADD A,D
label_9368: ADD A,E
label_9369: ADD A,H
label_9370: ADD A,L
label_9371: ADD A,$12
label_9372: ADD A,(HL)
label_9373: ADD A,(IX + 127)
label_9374: ADD A,(IY - 128)
label_9375: ADC A,A
label_9376: ADC A,B
label_9377: ADC A,C
label_9378: ADC A,D
label_9379: ADC A,E
label_9380: ADC A,H
label_9381: ADC A,L
label_9382: ADC A,$12
label_9383: ADC A,(HL)
label_9384: ADC A,(IX + 127)
label_9385: ADC A,(IY - 128)
label_9386: SUB A
label_9387: SUB B
label_9388: SUB C
label_9389: SUB D
label_9390: SUB E
label_9391: SUB H
label_9392: SUB L
label_9393: SUB $12
label_9394: SUB (HL)
label_9395: SUB (IX + 127)
label_9396: SUB (IY - 128)
label_9397: SBC A,A
label_9398: SBC A,B
label_9399: SBC A,C
label_9400: SBC A,D
label_9401: SBC A,E
label_9402: SBC A,H
label_9403: SBC A,L
label_9404: SBC A,$12
label_9405: SBC A,(HL)
label_9406: SBC A,(IX + 127)
label_9407: SBC A,(IY - 128)
label_9408: AND A
label_9409: AND B
label_9410: AND C
label_9411: AND D
label_9412: AND E
label_9413: AND H
label_9414: AND L
label_9415: AND $12
label_9416: AND (HL)
label_9417: AND (IX + 127)
label_9418: AND (IY - 128)
label_9419: AND A
label_9420: AND B
label_9421: AND C
label_9422: AND D
label_9423: AND E
label_9424: AND H
label_9425: AND L
label_9426: AND $12
label_9427: AND (HL)
label_9428: AND (IX + 127)
label_9429: AND (IY - 128)
label_9430: OR A
label_9431: OR B
label_9432: OR C
label_9433: OR D
label_9434: OR E
label_9435: OR H
label_9436: OR L
label_9437: OR $12
label_9438: OR (HL)
label_9439: OR (IX + 127)
label_9440: OR (IY - 128)
label_9441: XOR A
label_9442: XOR B
label_9443: XOR C
label_9444: XOR D
label_9445: XOR E
label_9446: XOR H
label_9447: XOR L
label_9448: XOR $12
label_9449: XOR (HL)
label_9450: XOR (IX + 127)
label_9451: XOR (IY - 128)
label_9452: CP A
label_9453: CP B
label_9454: CP C
label_9455: CP D
label_9456: CP E
label_9457: CP H
label_9458: CP L
label_9459: CP $12
label_9460: CP (HL)
label_9461: CP (IX + 127)
label_9462: CP (IY - 128)
label_9463: INC A
label_9464: INC B
label_9465: INC C
label_9466: INC D
label_9467: INC E
label_9468: INC H
label_9469: INC L
label_9470: INC (HL)
label_9471: INC (IX + 127)
label_9472: INC (IY - 128)
label_9473: DEC A
label_9474: DEC B
label_9475: DEC C
label_9476: DEC D
label_9477: DEC E
label_9478: DEC H
label_9479: DEC L
label_9480: DEC (HL)
label_9481: DEC (IX + 127)
label_9482: DEC (IY - 128)
label_9483: DAA
label_9484: CPL
label_9485: NEG
label_9486: CCF
label_9487: SCF
label_9488: NOP
label_9489: HALT
label_9490: DI
label_9491: EI
label_9492: IM 0
label_9493: IM 1
label_9494: IM 2
label_9495: ADD HL,BC
label_9496: ADD HL,DE
label_9497: ADD HL,HL
label_9498: ADD HL,SP
label_9499: ADC HL,BC
label_9500: ADC HL,DE
label_9501: ADC HL,HL
label_9502: ADC HL,SP
label_9503: SBC HL,BC
label_9504: SBC HL,DE
label_9505: SBC HL,HL
label_9506: SBC HL,SP
label_9507: ADD IX,BC
label_9508: ADD IX,DE
label_9509: ADD IX,SP
label_9510: ADD IY,BC
label_9511: ADD IY,DE
label_9512: ADD IY,SP
label_9513: INC BC
label_9514: INC DE
label_9515: INC HL
label_9516: INC SP
label_9517: INC IX
label_9518: INC IY
label_9519: DEC BC
label_9520: DEC DE
label_9521: DEC HL
label_9522: DEC SP
label_9523: DEC IX
label_9524: DEC IY
label_9525: RLCA
label_9526: RLA
label_9527: RRCA
label_9528: RRA
label_9529: RLC A
label_9530: RLC B
label_9531: RLC C
label_9532: RLC D
label_9533: RLC E
label_9534: RLC H
label_9535: RLC L
label_9536: RLC (HL)
label_9537: RLC (IX + 127)
label_9538: RLC (IY - 128)
label_9539: RL A
label_9540: RL B
label_9541: RL C
label_9542: RL D
label_9543: RL E
label_9544: RL H
label_9545: RL L
label_9546: RL (HL)
label_9547: RL (IX + 127)
label_9548: RL (IY - 128)
label_9549: RRC A
label_9550: RRC B
label_9551: RRC C
label_9552: RRC D
label_9553: RRC E
label_9554: RRC H
label_9555: RRC L
label_9556: RRC (HL)
label_9557: RRC (IX + 127)
label_9558: RRC (IY - 128)
label_9559: RR A
label_9560: RR B
label_9561: RR C
label_9562: RR D
label_9563: RR E
label_9564: RR H
label_9565: RR L
label_9566: RR (HL)
label_9567: RR (IX + 127)
label_9568: RR (IY - 128)
label_9569: SLA A
label_9570: SLA B
label_9571: SLA C
label_9572: SLA D
label_9573: SLA E
label_9574: SLA H
label_9575: SLA L
label_9576: SLA (HL)
label_9577: SLA (IX + 127)
label_9578: SLA (IY - 128)
label_9579: SRA A
label_9580: SRA B
label_9581: SRA C
label_9582: SRA D
label_9583: SRA E
label_9584: SRA H
label_9585: SRA L
label_9586: SRA (HL)
label_9587: SRA (IX + 127)
label_9588: SRA (IY - 128)
label_9589: SRL A
label_9590: SRL B
label_9591: SRL C
label_9592: SRL D
label_9593: SRL E
label_9594: SRL H
label_9595: SRL L
label_9596: SRL (HL)
label_9597: SRL (IX + 127)
label_9598: SRL (IY - 128)
label_9599: RLD
label_9600: RRD
label_9601: BIT 0,A
label_9602: BIT 1,A
label_9603: BIT 2,A
label_9604: BIT 3,A
label_9605: BIT 4,A
label_9606: BIT 5,A
label_9607: BIT 6,A
label_9608: BIT 7,A
label_9609: BIT 0,B
label_9610: BIT 1,B
label_9611: BIT 2,B
label_9612: BIT 3,B
label_9613: BIT 4,B
label_9614: BIT 5,B
label_9615: BIT 6,B
label_9616: BIT 7,B
label_9617: BIT 0,C
label_9618: BIT 1,C
label_9619: BIT 2,C
label_9620: BIT 3,C
label_9621: BIT 4,C
label_9622: BIT 5,C
label_9623: BIT 6,C
label_9624: BIT 7,C
label_9625: BIT 0,D
label_9626: BIT 1,D
label_9627: BIT 2,D
label_9628: BIT 3,D
label_9629: BIT 4,D
label_9630: BIT 5,D
label_9631: BIT 6,D
label_9632: BIT 7,D
label_9633: BIT 0,E
label_9634: BIT 1,E
label_9635: BIT 2,E
label_9636: BIT 3,E
label_9637: BIT 4,E
label_9638: BIT 5,E
label_9639: BIT 6,E
label_9640: BIT 7,E
label_9641: BIT 0,H
label_9642: BIT 1,H
label_9643: BIT 2,H
label_9644: BIT 3,H
label_9645: BIT 4,H
label_9646: BIT 5,H
label_9647: BIT 6,H
label_9648: BIT 7,H
label_9649: BIT 0,L
label_9650: BIT 1,L
label_9651: BIT 2,L
label_9652: BIT 3,L
label_9653: BIT 4,L
label_9654: BIT 5,L
label_9655: BIT 6,L
label_9656: BIT 7,L
label_9657: BIT 0,(HL)
label_9658: BIT 1,(HL)
label_9659: BIT 2,(HL)
label_9660: BIT 3,(HL)
label_9661: BIT 4,(HL)
label_9662: BIT 5,(HL)
label_9663: BIT 6,(HL)
label_9664: BIT 7,(HL)
label_9665: BIT 0,(IX + 127)
label_9666: BIT 1,(IX + 127)
label_9667: BIT 2,(IX + 127)
label_9668: BIT 3,(IX + 127)
label_9669: BIT 4,(IX + 127)
label_9670: BIT 5,(IX + 127)
label_9671: BIT 6,(IX + 127)
label_9672: BIT 7,(IX + 127)
label_9673: BIT 0,(IY - 128)
label_9674: BIT 1,(IY - 128)
label_9675: BIT 2,(IY - 128)
label_9676: BIT 3,(IY - 128)
label_9677: BIT 4,(IY - 128)
label_9678: BIT 5,(IY - 128)
label_9679: BIT 6,(IY - 128)
label_9680: BIT 7,(IY - 128)
label_9681: SET 0,A
label_9682: SET 1,A
label_9683: SET 2,A
label_9684: SET 3,A
label_9685: SET 4,A
label_9686: SET 5,A
label_9687: SET 6,A
label_9688: SET 7,A
label_9689: SET 0,B
label_9690: SET 1,B
label_9691: SET 2,B
label_9692: SET 3,B
label_9693: SET 4,B
label_9694: SET 5,B
label_9695: SET 6,B
label_9696: SET 7,B
label_9697: SET 0,C
label_9698: SET 1,C
label_9699: SET 2,C
label_9700: SET 3,C
label_9701: SET 4,C
label_9702: SET 5,C
label_9703: SET 6,C
label_9704: SET 7,C
label_9705: SET 0,D
label_9706: SET 1,D
label_9707: SET 2,D
label_9708: SET 3,D
label_9709: SET 4,D
label_9710: SET 5,D
label_9711: SET 6,D
label_9712: SET 7,D
label_9713: SET 0,E
label_9714: SET 1,E
label_9715: SET 2,E
label_9716: SET 3,E
label_9717: SET 4,E
label_9718: SET 5,E
label_9719: SET 6,E
label_9720: SET 7,E
label_9721: SET 0,H
label_9722: SET 1,H
label_9723: SET 2,H
label_9724: SET 3,H
label_9725: SET 4,H
label_9726: SET 5,H
label_9727: SET 6,H
label_9728: SET 7,H
label_9729: SET 0,L
label_9730: SET 1,L
label_9731: SET 2,L
label_9732: SET 3,L
label_9733: SET 4,L
label_9734: SET 5,L
label_9735: SET 6,L
label_9736: SET 7,L
label_9737: SET 0,(HL)
label_9738: SET 1,(HL)
label_9739: SET 2,(HL)
label_9740: SET 3,(HL)
label_9741: SET 4,(HL)
label_9742: SET 5,(HL)
label_9743: SET 6,(HL)
label_9744: SET 7,(HL)
label_9745: SET 0,(IX + 127)
label_9746: SET 1,(IX + 127)
label_9747: SET 2,(IX + 127)
label_9748: SET 3,(IX + 127)
label_9749: SET 4,(IX + 127)
label_9750: SET 5,(IX + 127)
label_9751: SET 6,(IX + 127)
label_9752: SET 7,(IX + 127)
label_9753: SET 0,(IY - 128)
label_9754: SET 1,(IY - 128)
label_9755: SET 2,(IY - 128)
label_9756: SET 3,(IY - 128)
label_9757: SET 4,(IY - 128)
label_9758: SET 5,(IY - 128)
label_9759: SET 6,(IY - 128)
label_9760: SET 7,(IY - 128)
label_9761: RES 0,A
label_9762: RES 1,A
label_9763: RES 2,A
label_9764: RES 3,A
label_9765: RES 4,A
label_9766: RES 5,A
label_9767: RES 6,A
label_9768: RES 7,A
label_9769: RES 0,B
label_9770: RES 1,B
label_9771: RES 2,B
label_9772: RES 3,B
label_9773: RES 4,B
label_9774: RES 5,B
label_9775: RES 6,B
label_9776: RES 7,B
label_9777: RES 0,C
label_9778: RES 1,C
label_9779: RES 2,C
label_9780: RES 3,C
label_9781: RES 4,C
label_9782: RES 5,C
label_9783: RES 6,C
label_9784: RES 7,C
label_9785: RES 0,D
label_9786: RES 1,D
label_9787: RES 2,D
label_9788: RES 3,D
label_9789: RES 4,D
label_9790: RES 5,D
label_9791: RES 6,D
label_9792: RES 7,D
label_9793: RES 0,E
label_9794: RES 1,E
label_9795: RES 2,E
label_9796: RES 3,E
label_9797: RES 4,E
label_9798: RES 5,E
label_9799: RES 6,E
label_9800: RES 7,E
label_9801: RES 0,H
label_9802: RES 1,H
label_9803: RES 2,H
label_9804: RES 3,H
label_9805: RES 4,H
label_9806: RES 5,H
label_9807: RES 6,H
label_9808: RES 7,H
label_9809: RES 0,L
label_9810: RES 1,L
label_9811: RES 2,L
label_9812: RES 3,L
label_9813: RES 4,L
label_9814: RES 5,L
label_9815: RES 6,L
label_9816: RES 7,L
label_9817: RES 0,(HL)
label_9818: RES 1,(HL)
label_9819: RES 2,(HL)
label_9820: RES 3,(HL)
label_9821: RES 4,(HL)
label_9822: RES 5,(HL)
label_9823: RES 6,(HL)
label_9824: RES 7,(HL)
label_9825: RES 0,(IX + 127)
label_9826: RES 1,(IX + 127)
label_9827: RES 2,(IX + 127)
label_9828: RES 3,(IX + 127)
label_9829: RES 4,(IX + 127)
label_9830: RES 5,(IX + 127)
label_9831: RES 6,(IX + 127)
label_9832: RES 7,(IX + 127)
label_9833: RES 0,(IY - 128)
label_9834: RES 1,(IY - 128)
label_9835: RES 2,(IY - 128)
label_9836: RES 3,(IY - 128)
label_9837: RES 4,(IY - 128)
label_9838: RES 5,(IY - 128)
label_9839: RES 6,(IY - 128)
label_9840: RES 7,(IY - 128)
label_9841: JP $5678
label_9842: JP NZ,$5678
label_9843: JP Z,$5678
label_9844: JP NC,$5678
label_9845: JP C,$5678
label_9846: JP PO,$5678
label_9847: JP PE,$5678
label_9848: JP P,$5678
label_9849: JP M,$5678
label_9850: JR $ + 2
label_9851: JR NZ,$ + 2
label_9852: JR Z,$ + 2
label_9853: JR NC,$ + 2
label_9854: JR C,$ + 2
label_9855: JP (HL)
label_9856: JP (IX)
label_9857: JP (IY)
label_9858: DJNZ $ + 2
label_9859: CALL $5678
label_9860: CALL NZ,$5678
label_9861: CALL Z,$5678
label_9862: CALL NC,$5678
label_9863: CALL C,$5678
label_9864: CALL PO,$5678
label_9865: CALL PE,$5678
label_9866: CALL P,$5678
label_9867: CALL M,$5678
label_9868: RET
label_9869: RET NZ
label_9870: RET Z
label_9871: RET NC
label_9872: RET C
label_9873: RET PO
label_9874: RET PE
label_9875: RET P
label_9876: RET M
label_9877: RETI
label_9878: RETN
label_9879: RST $00
label_9880: RST $08
label_9881: RST $10
label_9882: RST $18
label_9883: RST $20
label_9884: RST $28
label_9885: RST $30
label_9886: RST $38
label_9887: IN A,($12)
label_9888: IN A,(C)
label_9889: IN B,(C)
label_9890: IN C,(C)
label_9891: IN D,(C)
label_9892: IN E,(C)
label_9893: IN H,(C)
label_9894: IN L,(C)
label_9895: IN F,(C)
label_9896: INI
label_9897: INIR
label_9898: IND
label_9899: INDR
label_9900: OUT ($12),A
label_9901: OUT (C),A
label_9902: OUT (C),B
label_9903: OUT (C),C
label_9904: OUT (C),D
label_9905: OUT (C),E
label_9906: OUT (C),H
label_9907: OUT (C),L
label_9908: OUTI
label_9909: OTIR
label_9910: OUTD
label_9911: OTDR
label_9912: LD A,A
label_9913: LD A,B
label_9914: LD A,C
label_9915: LD A,D
label_9916: LD A,E
label_9917: LD A,H
label_9918: LD A,L
label_9919: LD B,A
label_9920: LD B,B
label_9921: LD B,C
label_9922: LD B,D
label_9923: LD B,E
label_9924: LD B,H
label_9925: LD B,L
label_9926: LD C,A
label_9927: LD C,B
label_9928: LD C,C
label_9929: LD C,D
label_9930: LD C,E
label_9931: LD C,H
label_9932: LD C,L
label_9933: LD D,A
label_9934: LD D,B
label_9935: LD D,C
label_9936: LD D,D
label_9937: LD D,E
label_9938: LD D,H
label_9939: LD D,L
label_9940: LD E,A
label_9941: LD E,B
label_9942: LD E,C
label_9943: LD E,D
label_9944: LD E,E
label_9945: LD E,H
label_9946: LD E,L
label_9947: LD H,A
label_9948: LD H,B
label_9949: LD H,C
label_9950: LD H,D
label_9951: LD H,E
label_9952: LD H,H
label_9953: LD H,L
label_9954: LD L,A
label_9955: LD L,B
label_9956: LD L,C
label_9957: LD L,D
label_9958: LD L,E
label_9959: LD L,H
label_9960: LD L,L
label_9961: LD A,$12
label_9962: LD B,$12
label_9963: LD C,$12
label_9964: LD D,$12
label_9965: LD E,$12
label_9966: LD H,$12
label_9967: LD L,$12
label_9968: LD A,(HL)
label_9969: LD B,(HL)
label_9970: LD C,(HL)
label_9971: LD D,(HL)
label_9972: LD E,(HL)
label_9973: LD H,(HL)
label_9974: LD L,(HL)
label_9975: LD A,(IX + 127)
label_9976: LD B,(IX + 127)
label_9977: LD C,(IX + 127)
label_9978: LD D,(IX + 127)
label_9979: LD E,(IX + 127)
label_9980: LD H,(IX + 127)
label_9981: LD L,(IX + 127)
label_9982: LD A,(IY - 128)
label_9983: LD B,(IY - 128)
label_9984: LD C,(IY - 128)
label_9985: LD D,(IY - 128)
label_9986: LD E,(IY - 128)
label_9987: LD H,(IY - 128)
label_9988: LD L,(IY - 128)
label_9989: LD (HL),A
label_9990: LD (HL),B
label_9991: LD (HL),C
label_9992: LD (HL),D
label_9993: LD (HL),E
label_9994: LD (HL),H
label_9995: LD (HL),L
label_9996: LD (IX + 127),A
label_9997: LD (IX + 127),B
label_9998: LD (IX + 127),C
label_9999: LD (IX + 127),D
label_10000: LD (IX + 127),E
label_10001: LD (IX + 127),H
label_10002: LD (IX + 127),L
label_10003: LD (IY - 128),A
label_10004: LD (IY - 128),B
label_10005: LD (IY - 128),C
label_10006: LD (IY - 128),D
label_10007: LD (IY - 128),E
label_10008: LD (IY - 128),H
label_10009: LD (IY - 128),L
label_10010: LD (HL),$12
label_10011: LD (IX + 127),$12
label_10012: LD (IY - 128),$12
label_10013: LD A,(BC)
label_10014: LD A,(DE)
label_10015: LD A,($5678)
label_10016: LD (BC),A
label_10017: LD (DE),A
label_10018: LD ($5678),A
label_10019: LD A,I
label_10020: LD A,R
label_10021: LD I,A
label_10022: LD R,A
label_10023: LD BC,$5678
label_10024: LD DE,$5678
label_10025: LD HL,$5678
label_10026: LD SP,$5678
label_10027: LD IX,$5678
label_10028: LD IY,$5678
label_10029: LD HL,($5678)
label_10030: LD BC,($5678)
label_10031: LD DE,($5678)
label_10032: LD HL,($5678)
label_10033: LD SP,($5678)
label_10034: LD IX,($5678)
label_10035: LD IY,($5678)
label_10036: LD ($5678),HL
label_10037: LD ($5678),BC
label_10038: LD ($5678),DE
label_10039: LD ($5678),HL
label_10040: LD ($5678),SP
label_10041: LD ($5678),IX
label_10042: LD ($5678),IY
label_10043: LD SP,HL
label_10044: LD SP,IX
label_10045: LD SP,IY
label_10046: PUSH BC
label_10047: PUSH DE
label_10048: PUSH HL
label_10049: PUSH AF
label_10050: PUSH IX
label_10051: PUSH IY
label_10052: POP BC
label_10053: POP DE
label_10054: POP HL
label_10055: POP AF
label_10056: POP IX
label_10057: POP IY
label_10058: EX DE,HL
label_10059: EX AF,AF'
label_10060: EXX
label_10061: EX (SP),HL
label_10062: EX (SP),IX
label_10063: EX (SP),IY
label_10064: LDI
label_10065: LDIR
label_10066: LDD
label_10067: LDDR
label_10068: CPI
label_10069: CPIR
label_10070: CPD
label_10071: CPDR
label_10072: ADD A,A
label_10073: ADD A,B
label_10074: ADD A,C
label_10075: ADD A,D
label_10076: ADD A,E
label_10077: ADD A,H
label_10078: ADD A,L
label_10079: ADD A,$12
label_10080: ADD A,(HL)
label_10081: ADD A,(IX + 127)
label_10082: ADD A,(IY - 128)
label_10083: ADC A,A
label_10084: ADC A,B
label_10085: ADC A,C
label_10086: ADC A,D
label_10087: ADC A,E
label_10088: ADC A,H
label_10089: ADC A,L
label_10090: ADC A,$12
label_10091: ADC A,(HL)
label_10092: ADC A,(IX + 127)
label_10093: ADC A,(IY - 128)
label_10094: SUB A
label_10095: SUB B
label_10096: SUB C
label_10097: SUB D
label_10098: SUB E
label_10099: SUB H
label_10100: SUB L
label_10101: SUB $12
label_10102: SUB (HL)
label_10103: SUB (IX + 127)
label_10104: SUB (IY - 128)
label_10105: SBC A,A
label_10106: SBC A,B
label_10107: SBC A,C
label_10108: SBC A,D
label_10109: SBC A,E
label_10110: SBC A,H
label_10111: SBC A,L
label_10112: SBC A,$12
label_10113: SBC A,(HL)
label_10114: SBC A,(IX + 127)
label_10115: SBC A,(IY - 128)
label_10116: AND A
label_10117: AND B
label_10118: AND C
label_10119: AND D
label_10120: AND E
label_10121: AND H
label_10122: AND L
label_10123: AND $12
label_10124: AND (HL)
label_10125: AND (IX + 127)
label_10126: AND (IY - 128)
label_10127: AND A
label_10128: AND B
label_10129: AND C
label_10130: AND D
label_10131: AND E
label_10132: AND H
label_10133: AND L
label_10134: AND $12
label_10135: AND (HL)
label_10136: AND (IX + 127)
label_10137: AND (IY - 128)
label_10138: OR A
label_10139: OR B
label_10140: OR C
label_10141: OR D
label_10142: OR E
label_10143: OR H
label_10144: OR L
label_10145: OR $12
label_10146: OR (HL)
label_10147: OR (IX + 127)
label_10148: OR (IY - 128)
label_10149: XOR A
label_10150: XOR B
label_10151: XOR C
label_10152: XOR D
label_10153: XOR E
label_10154: XOR H
label_10155: XOR L
label_10156: XOR $12
label_10157: XOR (HL)
label_10158: XOR (IX + 127)
label_10159: XOR (IY - 128)
label_10160: CP A
label_10161: CP B
label_10162: CP C
label_10163: CP D
label_10164: CP E
label_10165: CP H
label_10166: CP L
label_10167: CP $12
label_10168: CP (HL)
label_10169: CP (IX + 127)
label_10170: CP (IY - 128)
label_10171: INC A
label_10172: INC B
label_10173: INC C
label_10174: INC D
label_10175: INC E
label_10176: INC H
label_10177: INC L
label_10178: INC (HL)
label_10179: INC (IX + 127)
label_10180: INC (IY - 128)
label_10181: DEC A
label_10182: DEC B
label_10183: DEC C
label_10184: DEC D
label_10185: DEC E
label_10186: DEC H
label_10187: DEC L
label_10188: DEC (HL)
label_10189: DEC (IX + 127)
label_10190: DEC (IY - 128)
label_10191: DAA
label_10192: CPL
label_10193: NEG
label_10194: CCF
label_10195: SCF
label_10196: NOP
label_10197: HALT
label_10198: DI
label_10199: EI
label_10200: IM 0
label_10201: IM 1
label_10202: IM 2
label_10203: ADD HL,BC
label_10204: ADD HL,DE
label_10205: ADD HL,HL
label_10206: ADD HL,SP
label_10207: ADC HL,BC
label_10208: ADC HL,DE
label_10209: ADC HL,HL
label_10210: ADC HL,SP
label_10211: SBC HL,BC
label_10212: SBC HL,DE
label_10213: SBC HL,HL
label_10214: SBC HL,SP
label_10215: ADD IX,BC
label_10216: ADD IX,DE
label_10217: ADD IX,SP
label_10218: ADD IY,BC
label_10219: ADD IY,DE
label_10220: ADD IY,SP
label_10221: INC BC
label_10222: INC DE
label_10223: INC HL
label_10224: INC SP
label_10225: INC IX
label_10226: INC IY
label_10227: DEC BC
label_10228: DEC DE
label_10229: DEC HL
label_10230: DEC SP
label_10231: DEC IX
label_10232: DEC IY
label_10233: RLCA
label_10234: RLA
label_10235: RRCA
label_10236: RRA
label_10237: RLC A
label_10238: RLC B
label_10239: RLC C
label_10240: RLC D
label_10241: RLC E
label_10242: RLC H
label_10243: RLC L
label_10244: RLC (HL)
label_10245: RLC (IX + 127)
label_10246: RLC (IY - 128)
label_10247: RL A
label_10248: RL B
label_10249: RL C
label_10250: RL D
label_10251: RL E
label_10252: RL H
label_10253: RL L
label_10254: RL (HL)
label_10255: RL (IX + 127)
label_10256: RL (IY - 128)
label_10257: RRC A
label_10258: RRC B
label_10259: RRC C
label_10260: RRC D
label_10261: RRC E
label_10262: RRC H
label_10263: RRC L
label_10264: RRC (HL)
label_10265: RRC (IX + 127)
label_10266: RRC (IY - 128)
label_10267: RR A
label_10268: RR B
label_10269: RR C
label_10270: RR D
label_10271: RR E
label_10272: RR H
label_10273: RR L
label_10274: RR (HL)
label_10275: RR (IX + 127)
label_10276: RR (IY - 128)
label_10277: SLA A
label_10278: SLA B
label_10279: SLA C
label_10280: SLA D
label_10281: SLA E
label_10282: SLA H
label_10283: SLA L
label_10284: SLA (HL)
label_10285: SLA (IX + 127)
label_10286: SLA (IY - 128)
label_10287: SRA A
label_10288: SRA B
label_10289: SRA C
label_10290: SRA D
label_10291: SRA E
label_10292: SRA H
label_10293: SRA L
label_10294: SRA (HL)
label_10295: SRA (IX + 127)
label_10296: SRA (IY - 128)
label_10297: SRL A
label_10298: SRL B
label_10299: SRL C
label_10300: SRL D
label_10301: SRL E
label_10302: SRL H
label_10303: SRL L
label_10304: SRL (HL)
label_10305: SRL (IX + 127)
label_10306: SRL (IY - 128)
label_10307: RLD
label_10308: RRD
label_10309: BIT 0,A
label_10310: BIT 1,A
label_10311: BIT 2,A
label_10312: BIT 3,A
label_10313: BIT 4,A
label_10314: BIT 5,A
label_10315: BIT 6,A
label_10316: BIT 7,A
label_10317: BIT 0,B
label_10318: BIT 1,B
label_10319: BIT 2,B
label_10320: BIT 3,B
label_10321: BIT 4,B
label_10322: BIT 5,B
label_10323: BIT 6,B
label_10324: BIT 7,B
label_10325: BIT 0,C
label_10326: BIT 1,C
label_10327: BIT 2,C
label_10328: BIT 3,C
label_10329: BIT 4,C
label_10330: BIT 5,C
label_10331: BIT 6,C
label_10332: BIT 7,C
label_10333: BIT 0,D
label_10334: BIT 1,D
label_10335: BIT 2,D
label_10336: BIT 3,D
label_10337: BIT 4,D
label_10338: BIT 5,D
label_10339: BIT 6,D
label_10340: BIT 7,D
label_10341: BIT 0,E
label_10342: BIT 1,E
label_10343: BIT 2,E
label_10344: BIT 3,E
label_10345: BIT 4,E
label_10346: BIT 5,E
label_10347: BIT 6,E
label_10348: BIT 7,E
label_10349: BIT 0,H
label_10350: BIT 1,H
label_10351: BIT 2,H
label_10352: BIT 3,H
label_10353: BIT 4,H
label_10354: BIT 5,H
label_10355: BIT 6,H
label_10356: BIT 7,H
label_10357: BIT 0,L
label_10358: BIT 1,L
label_10359: BIT 2,L
label_10360: BIT 3,L
label_10361: BIT 4,L
label_10362: BIT 5,L
label_10363: BIT 6,L
label_10364: BIT 7,L
label_10365: BIT 0,(HL)
label_10366: BIT 1,(HL)
label_10367: BIT 2,(HL)
label_10368: BIT 3,(HL)
label_10369: BIT 4,(HL)
label_10370: BIT 5,(HL)
label_10371: BIT 6,(HL)
label_10372: BIT 7,(HL)
label_10373: BIT 0,(IX + 127)
label_10374: BIT 1,(IX + 127)
label_10375: BIT 2,(IX + 127)
label_10376: BIT 3,(IX + 127)
label_10377: BIT 4,(IX + 127)
label_10378: BIT 5,(IX + 127)
label_10379: BIT 6,(IX + 127)
label_10380: BIT 7,(IX + 127)
label_10381: BIT 0,(IY - 128)
label_10382: BIT 1,(IY - 128)
label_10383: BIT 2,(IY - 128)
label_10384: BIT 3,(IY - 128)
label_10385: BIT 4,(IY - 128)
label_10386: BIT 5,(IY - 128)
label_10387: BIT 6,(IY - 128)
label_10388: BIT 7,(IY - 128)
label_10389: SET 0,A
label_10390: SET 1,A
label_10391: SET 2,A
label_10392: SET 3,A
label_10393: SET 4,A
label_10394: SET 5,A
label_10395: SET 6,A
label_10396: SET 7,A
label_10397: SET 0,B
label_10398: SET 1,B
label_10399: SET 2,B
label_10400: SET 3,B
label_10401: SET 4,B
label_10402: SET 5,B
label_10403: SET 6,B
label_10404: SET 7,B
label_10405: SET 0,C
label_10406: SET 1,C
label_10407: SET 2,C
label_10408: SET 3,C
label_10409: SET 4,C
label_10410: SET 5,C
label_10411: SET 6,C
label_10412: SET 7,C
label_10413: SET 0,D
label_10414: SET 1,D
label_10415: SET 2,D
label_10416: SET 3,D
label_10417: SET 4,D
label_10418: SET 5,D
label_10419: SET 6,D
label_10420: SET 7,D
label_10421: SET 0,E
label_10422: SET 1,E
label_10423: SET 2,E
label_10424: SET 3,E
label_10425: SET 4,E
label_10426: SET 5,E
label_10427: SET 6,E
label_10428: SET 7,E
label_10429: SET 0,H
label_10430: SET 1,H
label_10431: SET 2,H
label_10432: SET 3,H
label_10433: SET 4,H
label_10434: SET 5,H
label_10435: SET 6,H
label_10436: SET 7,H
label_10437: SET 0,L
label_10438: SET 1,L
label_10439: SET 2,L
label_10440: SET 3,L
label_10441: SET 4,L
label_10442: SET 5,L
label_10443: SET 6,L
label_10444: SET 7,L
label_10445: SET 0,(HL)
label_10446: SET 1,(HL)
label_10447: SET 2,(HL)
label_10448: SET 3,(HL)
label_10449: SET 4,(HL)
label_10450: SET 5,(HL)
label_10451: SET 6,(HL)
label_10452: SET 7,(HL)
label_10453: SET 0,(IX + 127)
label_10454: SET 1,(IX + 127)
label_10455: SET 2,(IX + 127)
label_10456: SET 3,(IX + 127)
label_10457: SET 4,(IX + 127)
label_10458: SET 5,(IX + 127)
label_10459: SET 6,(IX + 127)
label_10460: SET 7,(IX + 127)
label_10461: SET 0,(IY - 128)
label_10462: SET 1,(IY - 128)
label_10463: SET 2,(IY - 128)
label_10464: SET 3,(IY - 128)
label_10465: SET 4,(IY - 128)
label_10466: SET 5,(IY - 128)
label_10467: SET 6,(IY - 128)
label_10468: SET 7,(IY - 128)
label_10469: RES 0,A
label_10470: RES 1,A
label_10471: RES 2,A
label_10472: RES 3,A
label_10473: RES 4,A
label_10474: RES 5,A
label_10475: RES 6,A
label_10476: RES 7,A
label_10477: RES 0,B
label_10478: RES 1,B
label_10479: RES 2,B
label_10480: RES 3,B
label_10481: RES 4,B
label_10482: RES 5,B
label_10483: RES 6,B
label_10484: RES 7,B
label_10485: RES 0,C
label_10486: RES 1,C
label_10487: RES 2,C
label_10488: RES 3,C
label_10489: RES 4,C
label_10490: RES 5,C
label_10491: RES 6,C
label_10492: RES 7,C
label_10493: RES 0,D
label_10494: RES 1,D
label_10495: RES 2,D
label_10496: RES 3,D
label_10497: RES 4,D
label_10498: RES 5,D
label_10499: RES 6,D
label_10500: RES 7,D
label_10501: RES 0,E
label_10502: RES 1,E
label_10503: RES 2,E
label_10504: RES 3,E
label_10505: RES 4,E
label_10506: RES 5,E
label_10507: RES 6,E
label_10508: RES 7,E
label_10509: RES 0,H
label_10510: RES 1,H
label_10511: RES 2,H
label_10512: RES 3,H
label_10513: RES 4,H
label_10514: RES 5,H
label_10515: RES 6,H
label_10516: RES 7,H
label_10517: RES 0,L
label_10518: RES 1,L
label_10519: RES 2,L
label_10520: RES 3,L
label_10521: RES 4,L
label_10522: RES 5,L
label_10523: RES 6,L
label_10524: RES 7,L
label_10525: RES 0,(HL)
label_10526: RES 1,(HL)
label_10527: RES 2,(HL)
label_10528: RES 3,(HL)
label_10529: RES 4,(HL)
label_10530: RES 5,(HL)
label_10531: RES 6,(HL)
label_10532: RES 7,(HL)
label_10533: RES 0,(IX + 127)
label_10534: RES 1,(IX + 127)
label_10535: RES 2,(IX + 127)
label_10536: RES 3,(IX + 127)
label_10537: RES 4,(IX + 127)
label_10538: RES 5,(IX + 127)
label_10539: RES 6,(IX + 127)
label_10540: RES 7,(IX + 127)
label_10541: RES 0,(IY - 128)
label_10542: RES 1,(IY - 128)
label_10543: RES 2,(IY - 128)
label_10544: RES 3,(IY - 128)
label_10545: RES 4,(IY - 128)
label_10546: RES 5,(IY - 128)
label_10547: RES 6,(IY - 128)
label_10548: RES 7,(IY - 128)
label_10549: JP $5678
label_10550: JP NZ,$5678
label_10551: JP Z,$5678
label_10552: JP NC,$5678
label_10553: JP C,$5678
label_10554: JP PO,$5678
label_10555: JP PE,$5678
label_10556: JP P,$5678
label_10557: JP M,$5678
label_10558: JR $ + 2
label_10559: JR NZ,$ + 2
label_10560: JR Z,$ + 2
label_10561: JR NC,$ + 2
label_10562: JR C,$ + 2
label_10563: JP (HL)
label_10564: JP (IX)
label_10565: JP (IY)
label_10566: DJNZ $ + 2
label_10567: CALL $5678
label_10568: CALL NZ,$5678
label_10569: CALL Z,$5678
label_10570: CALL NC,$5678
label_10571: CALL C,$5678
label_10572: CALL PO,$5678
label_10573: CALL PE,$5678
label_10574: CALL P,$5678
label_10575: CALL M,$5678
label_10576: RET
label_10577: RET NZ
label_10578: RET Z
label_10579: RET NC
label_10580: RET C
label_10581: RET PO
label_10582: RET PE
label_10583: RET P
label_10584: RET M
label_10585: RETI
label_10586: RETN
label_10587: RST $00
label_10588: RST $08
label_10589: RST $10
label_10590: RST $18
label_10591: RST $20
label_10592: RST $28
label_10593: RST $30
label_10594: RST $38
label_10595: IN A,($12)
label_10596: IN A,(C)
label_10597: IN B,(C)
label_10598: IN C,(C)
label_10599: IN D,(C)
label_10600: IN E,(C)
label_10601: IN H,(C)
label_10602: IN L,(C)
label_10603: IN F,(C)
label_10604: INI
label_10605: INIR
label_10606: IND
label_10607: INDR
label_10608: OUT ($12),A
label_10609: OUT (C),A
label_10610: OUT (C),B
label_10611: OUT (C),C
label_10612: OUT (C),D
label_10613: OUT (C),E
label_10614: OUT (C),H
label_10615: OUT (C),L
label_10616: OUTI
label_10617: OTIR
label_10618: OUTD
label_10619: OTDR
label_10620: LD A,A
label_10621: LD A,B
label_10622: LD A,C
label_10623: LD A,D
label_10624: LD A,E
label_10625: LD A,H
label_10626: LD A,L
label_10627: LD B,A
label_10628: LD B,B
label_10629: LD B,C
label_10630: LD B,D
label_10631: LD B,E
label_10632: LD B,H
label_10633: LD B,L
label_10634: LD C,A
label_10635: LD C,B
label_10636: LD C,C
label_10637: LD C,D
label_10638: LD C,E
label_10639: LD C,H
label_10640: LD C,L
label_10641: LD D,A
label_10642: LD D,B
label_10643: LD D,C
label_10644: LD D,D
label_10645: LD D,E
label_10646: LD D,H
label_10647: LD D,L
label_10648: LD E,A
label_10649: LD E,B
label_10650: LD E,C
label_10651: LD E,D
label_10652: LD E,E
label_10653: LD E,H
label_10654: LD E,L
label_10655: LD H,A
label_10656: LD H,B
label_10657: LD H,C
label_10658: LD H,D
label_10659: LD H,E
label_10660: LD H,H
label_10661: LD H,L
label_10662: LD L,A
label_10663: LD L,B
label_10664: LD L,C
label_10665: LD L,D
label_10666: LD L,E
label_10667: LD L,H
label_10668: LD L,L
label_10669: LD A,$12
label_10670: LD B,$12
label_10671: LD C,$12
label_10672: LD D,$12
label_10673: LD E,$12
label_10674: LD H,$12
label_10675: LD L,$12
label_10676: LD A,(HL)
label_10677: LD B,(HL)
label_10678: LD C,(HL)
label_10679: LD D,(HL)
label_10680: LD E,(HL)
label_10681: LD H,(HL)
label_10682: LD L,(HL)
label_10683: LD A,(IX + 127)
label_10684: LD B,(IX + 127)
label_10685: LD C,(IX + 127)
label_10686: LD D,(IX + 127)
label_10687: LD E,(IX + 127)
label_10688: LD H,(IX + 127)
label_10689: LD L,(IX + 127)
label_10690: LD A,(IY - 128)
label_10691: LD B,(IY - 128)
label_10692: LD C,(IY - 128)
label_10693: LD D,(IY - 128)
label_10694: LD E,(IY - 128)
label_10695: LD H,(IY - 128)
label_10696: LD L,(IY - 128)
label_10697: LD (HL),A
label_10698: LD (HL),B
label_10699: LD (HL),C
label_10700: LD (HL),D
label_10701: LD (HL),E
label_10702: LD (HL),H
label_10703: LD (HL),L
label_10704: LD (IX + 127),A
label_10705: LD (IX + 127),B
label_10706: LD (IX + 127),C
label_10707: LD (IX + 127),D
label_10708: LD (IX + 127),E
label_10709: LD (IX + 127),H
label_10710: LD (IX + 127),L
label_10711: LD (IY - 128),A
label_10712: LD (IY - 128),B
label_10713: LD (IY - 128),C
label_10714: LD (IY - 128),D
label_10715: LD (IY - 128),E
label_10716: LD (IY - 128),H
label_10717: LD (IY - 128),L
label_10718: LD (HL),$12
label_10719: LD (IX + 127),$12
label_10720: LD (IY - 128),$12
label_10721: LD A,(BC)
label_10722: LD A,(DE)
label_10723: LD A,($5678)
label_10724: LD (BC),A
label_10725: LD (DE),A
label_10726: LD ($5678),A
label_10727: LD A,I
label_10728: LD A,R
label_10729: LD I,A
label_10730: LD R,A
label_10731: LD BC,$5678
label_10732: LD DE,$5678
label_10733: LD HL,$5678
label_10734: LD SP,$5678
label_10735: LD IX,$5678
label_10736: LD IY,$5678
label_10737: LD HL,($5678)
label_10738: LD BC,($5678)
label_10739: LD DE,($5678)
label_10740: LD HL,($5678)
label_10741: LD SP,($5678)
label_10742: LD IX,($5678)
label_10743: LD IY,($5678)
label_10744: LD ($5678),HL
label_10745: LD ($5678),BC
label_10746: LD ($5678),DE
label_10747: LD ($5678),HL
label_10748: LD ($5678),SP
label_10749: LD ($5678),IX
label_10750: LD ($5678),IY
label_10751: LD SP,HL
label_10752: LD SP,IX
label_10753: LD SP,IY
label_10754: PUSH BC
label_10755: PUSH DE
label_10756: PUSH HL
label_10757: PUSH AF
label_10758: PUSH IX
label_10759: PUSH IY
label_10760: POP BC
label_10761: POP DE
label_10762: POP HL
label_10763: POP AF
label_10764: POP IX
label_10765: POP IY
label_10766: EX DE,HL
label_10767: EX AF,AF'
label_10768: EXX
label_10769: EX (SP),HL
label_10770: EX (SP),IX
label_10771: EX (SP),IY
label_10772: LDI
label_10773: LDIR
label_10774: LDD
label_10775: LDDR
label_10776: CPI
label_10777: CPIR
label_10778: CPD
label_10779: CPDR
label_10780: ADD A,A
label_10781: ADD A,B
label_10782: ADD A,C
label_10783: ADD A,D
label_10784: ADD A,E
label_10785: ADD A,H
label_10786: ADD A,L
label_10787: ADD A,$12
label_10788: ADD A,(HL)
label_10789: ADD A,(IX + 127)
label_10790: ADD A,(IY - 128)
label_10791: ADC A,A
label_10792: ADC A,B
label_10793: ADC A,C
label_10794: ADC A,D
label_10795: ADC A,E
label_10796: ADC A,H
label_10797: ADC A,L
label_10798: ADC A,$12
label_10799: ADC A,(HL)
label_10800: ADC A,(IX + 127)
label_10801: ADC A,(IY - 128)
label_10802: SUB A
label_10803: SUB B
label_10804: SUB C
label_10805: SUB D
label_10806: SUB E
label_10807: SUB H
label_10808: SUB L
label_10809: SUB $12
label_10810: SUB (HL)
label_10811: SUB (IX + 127)
label_10812: SUB (IY - 128)
label_10813: SBC A,A
label_10814: SBC A,B
label_10815: SBC A,C
label_10816: SBC A,D
label_10817: SBC A,E
label_10818: SBC A,H
label_10819: SBC A,L
label_10820: SBC A,$12
label_10821: SBC A,(HL)
label_10822: SBC A,(IX + 127)
label_10823: SBC A,(IY - 128)
label_10824: AND A
label_10825: AND B
label_10826: AND C
label_10827: AND D
label_10828: AND E
label_10829: AND H
label_10830: AND L
label_10831: AND $12
label_10832: AND (HL)
label_10833: AND (IX + 127)
label_10834: AND (IY - 128)
label_10835: AND A
label_10836: AND B
label_10837: AND C
label_10838: AND D
label_10839: AND E
label_10840: AND H
label_10841: AND L
label_10842: AND $12
label_10843: AND (HL)
label_10844: AND (IX + 127)
label_10845: AND (IY - 128)
label_10846: OR A
label_10847: OR B
label_10848: OR C
label_10849: OR D
label_10850: OR E
label_10851: OR H
label_10852: OR L
label_10853: OR $12
label_10854: OR (HL)
label_10855: OR (IX + 127)
label_10856: OR (IY - 128)
label_10857: XOR A
label_10858: XOR B
label_10859: XOR C
label_10860: XOR D
label_10861: XOR E
label_10862: XOR H
label_10863: XOR L
label_10864: XOR $12
label_10865: XOR (HL)
label_10866: XOR (IX + 127)
label_10867: XOR (IY - 128)
label_10868: CP A
label_10869: CP B
label_10870: CP C
label_10871: CP D
label_10872: CP E
label_10873: CP H
label_10874: CP L
label_10875: CP $12
label_10876: CP (HL)
label_10877: CP (IX + 127)
label_10878: CP (IY - 128)
label_10879: INC A
label_10880: INC B
label_10881: INC C
label_10882: INC D
label_10883: INC E
label_10884: INC H
label_10885: INC L
label_10886: INC (HL)
label_10887: INC (IX + 127)
label_10888: INC (IY - 128)
label_10889: DEC A
label_10890: DEC B
label_10891: DEC C
label_10892: DEC D
label_10893: DEC E
label_10894: DEC H
label_10895: DEC L
label_10896: DEC (HL)
label_10897: DEC (IX + 127)
label_10898: DEC (IY - 128)
label_10899: DAA
label_10900: CPL
label_10901: NEG
label_10902: CCF
label_10903: SCF
label_10904: NOP
label_10905: HALT
label_10906: DI
label_10907: EI
label_10908: IM 0
label_10909: IM 1
label_10910: IM 2
label_10911: ADD HL,BC
label_10912: ADD HL,DE
label_10913: ADD HL,HL
label_10914: ADD HL,SP
label_10915: ADC HL,BC
label_10916: ADC HL,DE
label_10917: ADC HL,HL
label_10918: ADC HL,SP
label_10919: SBC HL,BC
label_10920: SBC HL,DE
label_10921: SBC HL,HL
label_10922: SBC HL,SP
label_10923: ADD IX,BC
label_10924: ADD IX,DE
label_10925: ADD IX,SP
label_10926: ADD IY,BC
label_10927: ADD IY,DE
label_10928: ADD IY,SP
label_10929: INC BC
label_10930: INC DE
label_10931: INC HL
label_10932: INC SP
label_10933: INC IX
label_10934: INC IY
label_10935: DEC BC
label_10936: DEC DE
label_10937: DEC HL
label_10938: DEC SP
label_10939: DEC IX
label_10940: DEC IY
label_10941: RLCA
label_10942: RLA
label_10943: RRCA
label_10944: RRA
label_10945: RLC A
label_10946: RLC B
label_10947: RLC C
label_10948: RLC D
label_10949: RLC E
label_10950: RLC H
label_10951: RLC L
label_10952: RLC (HL)
label_10953: RLC (IX + 127)
label_10954: RLC (IY - 128)
label_10955: RL A
label_10956: RL B
label_10957: RL C
label_10958: RL D
label_10959: RL E
label_10960: RL H
label_10961: RL L
label_10962: RL (HL)
label_10963: RL (IX + 127)
label_10964: RL (IY - 128)
label_10965: RRC A
label_10966: RRC B
label_10967: RRC C
label_10968: RRC D
label_10969: RRC E
label_10970: RRC H
label_10971: RRC L
label_10972: RRC (HL)
label_10973: RRC (IX + 127)
label_10974: RRC (IY - 128)
label_10975: RR A
label_10976: RR B
label_10977: RR C
label_10978: RR D
label_10979: RR E
label_10980: RR H
label_10981: RR L
label_10982: RR (HL)
label_10983: RR (IX + 127)
label_10984: RR (IY - 128)
label_10985: SLA A
label_10986: SLA B
label_10987: SLA C
label_10988: SLA D
label_10989: SLA E
label_10990: SLA H
label_10991: SLA L
label_10992: SLA (HL)
label_10993: SLA (IX + 127)
label_10994: SLA (IY - 128)
label_10995: SRA A
label_10996: SRA B
label_10997: SRA C
label_10998: SRA D
label_10999: SRA E
label_11000: SRA H
label_11001: SRA L
label_11002: SRA (HL)
label_11003: SRA (IX + 127)
label_11004: SRA (IY - 128)
label_11005: SRL A
label_11006: SRL B
label_11007: SRL C
label_11008: SRL D
label_11009: SRL E
label_11010: SRL H
label_11011: SRL L
label_11012: SRL (HL)
label_11013: SRL (IX + 127)
label_11014: SRL (IY - 128)
label_11015: RLD
label_11016: RRD
label_11017: BIT 0,A
label_11018: BIT 1,A
label_11019: BIT 2,A
label_11020: BIT 3,A
label_11021: BIT 4,A
label_11022: BIT 5,A
label_11023: BIT 6,A
label_11024: BIT 7,A
label_11025: BIT 0,B
label_11026: BIT 1,B
label_11027: BIT 2,B
label_11028: BIT 3,B
label_11029: BIT 4,B
label_11030: BIT 5,B
label_11031: BIT 6,B
label_11032: BIT 7,B
label_11033: BIT 0,C
label_11034: BIT 1,C
label_11035: BIT 2,C
label_11036: BIT 3,C
label_11037: BIT 4,C
label_11038: BIT 5,C
label_11039: BIT 6,C
label_11040: BIT 7,C
label_11041: BIT 0,D
label_11042: BIT 1,D
label_11043: BIT 2,D
label_11044: BIT 3,D
label_11045: BIT 4,D
label_11046: BIT 5,D
label_11047: BIT 6,D
label_11048: BIT 7,D
label_11049: BIT 0,E
label_11050: BIT 1,E
label_11051: BIT 2,E
label_11052: BIT 3,E
label_11053: BIT 4,E
label_11054: BIT 5,E
label_11055: BIT 6,E
label_11056: BIT 7,E
label_11057: BIT 0,H
label_11058: BIT 1,H
label_11059: BIT 2,H
label_11060: BIT 3,H
label_11061: BIT 4,H
label_11062: BIT 5,H
label_11063: BIT 6,H
label_11064: BIT 7,H
label_11065: BIT 0,L
label_11066: BIT 1,L
label_11067: BIT 2,L
label_11068: BIT 3,L
label_11069: BIT 4,L
label_11070: BIT 5,L
label_11071: BIT 6,L
label_11072: BIT 7,L
label_11073: BIT 0,(HL)
label_11074: BIT 1,(HL)
label_11075: BIT 2,(HL)
label_11076: BIT 3,(HL)
label_11077: BIT 4,(HL)
label_11078: BIT 5,(HL)
label_11079: BIT 6,(HL)
label_11080: BIT 7,(HL)
label_11081: BIT 0,(IX + 127)
label_11082: BIT 1,(IX + 127)
label_11083: BIT 2,(IX + 127)
label_11084: BIT 3,(IX + 127)
label_11085: BIT 4,(IX + 127)
label_11086: BIT 5,(IX + 127)
label_11087: BIT 6,(IX + 127)
label_11088: BIT 7,(IX + 127)
label_11089: BIT 0,(IY - 128)
label_11090: BIT 1,(IY - 128)
label_11091: BIT 2,(IY - 128)
label_11092: BIT 3,(IY - 128)
label_11093: BIT 4,(IY - 128)
label_11094: BIT 5,(IY - 128)
label_11095: BIT 6,(IY - 128)
label_11096: BIT 7,(IY - 128)
label_11097: SET 0,A
label_11098: SET 1,A
label_11099: SET 2,A
label_11100: SET 3,A
label_11101: SET 4,A
label_11102: SET 5,A
label_11103: SET 6,A
label_11104: SET 7,A
label_11105: SET 0,B
label_11106: SET 1,B
label_11107: SET 2,B
label_11108: SET 3,B
label_11109: SET 4,B
label_11110: SET 5,B
label_11111: SET 6,B
label_11112: SET 7,B
label_11113: SET 0,C
label_11114: SET 1,C
label_11115: SET 2,C
label_11116: SET 3,C
label_11117: SET 4,C
label_11118: SET 5,C
label_11119: SET 6,C
label_11120: SET 7,C
label_11121: SET 0,D
label_11122: SET 1,D
label_11123: SET 2,D
label_11124: SET 3,D
label_11125: SET 4,D
label_11126: SET 5,D
label_11127: SET 6,D
label_11128: SET 7,D
label_11129: SET 0,E
label_11130: SET 1,E
label_11131: SET 2,E
label_11132: SET 3,E
label_11133: SET 4,E
label_11134: SET 5,E
label_11135: SET 6,E
label_11136: SET 7,E
label_11137: SET 0,H
label_11138: SET 1,H
label_11139: SET 2,H
label_11140: SET 3,H
label_11141: SET 4,H
label_11142: SET 5,H
label_11143: SET 6,H
label_11144: SET 7,H
label_11145: SET 0,L
label_11146: SET 1,L
label_11147: SET 2,L
label_11148: SET 3,L
label_11149: SET 4,L
label_11150: SET 5,L
label_11151: SET 6,L
label_11152: SET 7,L
label_11153: SET 0,(HL)
label_11154: SET 1,(HL)
label_11155: SET 2,(HL)
label_11156: SET 3,(HL)
label_11157: SET 4,(HL)
label_11158: SET 5,(HL)
label_11159: SET 6,(HL)
label_11160: SET 7,(HL)
label_11161: SET 0,(IX + 127)
label_11162: SET 1,(IX + 127)
label_11163: SET 2,(IX + 127)
label_11164: SET 3,(IX + 127)
label_11165: SET 4,(IX + 127)
label_11166: SET 5,(IX + 127)
label_11167: SET 6,(IX + 127)
label_11168: SET 7,(IX + 127)
label_11169: SET 0,(IY - 128)
label_11170: SET 1,(IY - 128)
label_11171: SET 2,(IY - 128)
label_11172: SET 3,(IY - 128)
label_11173: SET 4,(IY - 128)
label_11174: SET 5,(IY - 128)
label_11175: SET 6,(IY - 128)
label_11176: SET 7,(IY - 128)
label_11177: RES 0,A
label_11178: RES 1,A
label_11179: RES 2,A
label_11180: RES 3,A
label_11181: RES 4,A
label_11182: RES 5,A
label_11183: RES 6,A
label_11184: RES 7,A
label_11185: RES 0,B
label_11186: RES 1,B
label_11187: RES 2,B
label_11188: RES 3,B
label_11189: RES 4,B
label_11190: RES 5,B
label_11191: RES 6,B
label_11192: RES 7,B
label_11193: RES 0,C
label_11194: RES 1,C
label_11195: RES 2,C
label_11196: RES 3,C
label_11197: RES 4,C
label_11198: RES 5,C
label_11199: RES 6,C
label_11200: RES 7,C
label_11201: RES 0,D
label_11202: RES 1,D
label_11203: RES 2,D
label_11204: RES 3,D
label_11205: RES 4,D
label_11206: RES 5,D
label_11207: RES 6,D
label_11208: RES 7,D
label_11209: RES 0,E
label_11210: RES 1,E
label_11211: RES 2,E
label_11212: RES 3,E
label_11213: RES 4,E
label_11214: RES 5,E
label_11215: RES 6,E
label_11216: RES 7,E
label_11217: RES 0,H
label_11218: RES 1,H
label_11219: RES 2,H
label_11220: RES 3,H
label_11221: RES 4,H
label_11222: RES 5,H
label_11223: RES 6,H
label_11224: RES 7,H
label_11225: RES 0,L
label_11226: RES 1,L
label_11227: RES 2,L
label_11228: RES 3,L
label_11229: RES 4,L
label_11230: RES 5,L
label_11231: RES 6,L
label_11232: RES 7,L
label_11233: RES 0,(HL)
label_11234: RES 1,(HL)
label_11235: RES 2,(HL)
label_11236: RES 3,(HL)
label_11237: RES 4,(HL)
label_11238: RES 5,(HL)
label_11239: RES 6,(HL)
label_11240: RES 7,(HL)
label_11241: RES 0,(IX + 127)
label_11242: RES 1,(IX + 127)
label_11243: RES 2,(IX + 127)
label_11244: RES 3,(IX + 127)
label_11245: RES 4,(IX + 127)
label_11246: RES 5,(IX + 127)
label_11247: RES 6,(IX + 127)
label_11248: RES 7,(IX + 127)
label_11249: RES 0,(IY - 128)
label_11250: RES 1,(IY - 128)
label_11251: RES 2,(IY - 128)
label_11252: RES 3,(IY - 128)
label_11253: RES 4,(IY - 128)
label_11254: RES 5,(IY - 128)
label_11255: RES 6,(IY - 128)
label_11256: RES 7,(IY - 128)
label_11257: JP $5678
label_11258: JP NZ,$5678
label_11259: JP Z,$5678
label_11260: JP NC,$5678
label_11261: JP C,$5678
label_11262: JP PO,$5678
label_11263: JP PE,$5678
label_11264: JP P,$5678
label_11265: JP M,$5678
label_11266: JR $ + 2
label_11267: JR NZ,$ + 2
label_11268: JR Z,$ + 2
label_11269: JR NC,$ + 2
label_11270: JR C,$ + 2
label_11271: JP (HL)
label_11272: JP (IX)
label_11273: JP (IY)
label_11274: DJNZ $ + 2
label_11275: CALL $5678
label_11276: CALL NZ,$5678
label_11277: CALL Z,$5678
label_11278: CALL NC,$5678
label_11279: CALL C,$5678
label_11280: CALL PO,$5678
label_11281: CALL PE,$5678
label_11282: CALL P,$5678
label_11283: CALL M,$5678
label_11284: RET
label_11285: RET NZ
label_11286: RET Z
label_11287: RET NC
label_11288: RET C
label_11289: RET PO
label_11290: RET PE
label_11291: RET P
label_11292: RET M
label_11293: RETI
label_11294: RETN
label_11295: RST $00
label_11296: RST $08
label_11297: RST $10
label_11298: RST $18
label_11299: RST $20
label_11300: RST $28
label_11301: RST $30
label_11302: RST $38
label_11303: IN A,($12)
label_11304: IN A,(C)
label_11305: IN B,(C)
label_11306: IN C,(C)
label_11307: IN D,(C)
label_11308: IN E,(C)
label_11309: IN H,(C)
label_11310: IN L,(C)
label_11311: IN F,(C)
label_11312: INI
label_11313: INIR
label_11314: IND
label_11315: INDR
label_11316: OUT ($12),A
label_11317: OUT (C),A
label_11318: OUT (C),B
label_11319: OUT (C),C
label_11320: OUT (C),D
label_11321: OUT (C),E
label_11322: OUT (C),H
label_11323: OUT (C),L
label_11324: OUTI
label_11325: OTIR
label_11326: OUTD
label_11327: OTDR
label_11328: LD A,A
label_11329: LD A,B
label_11330: LD A,C
label_11331: LD A,D
label_11332: LD A,E
label_11333: LD A,H
label_11334: LD A,L
label_11335: LD B,A
label_11336: LD B,B
label_11337: LD B,C
label_11338: LD B,D
label_11339: LD B,E
label_11340: LD B,H
label_11341: LD B,L
label_11342: LD C,A
label_11343: LD C,B
label_11344: LD C,C
label_11345: LD C,D
label_11346: LD C,E
label_11347: LD C,H
label_11348: LD C,L
label_11349: LD D,A
label_11350: LD D,B
label_11351: LD D,C
label_11352: LD D,D
label_11353: LD D,E
label_11354: LD D,H
label_11355: LD D,L
label_11356: LD E,A
label_11357: LD E,B
label_11358: LD E,C
label_11359: LD E,D
label_11360: LD E,E
label_11361: LD E,H
label_11362: LD E,L
label_11363: LD H,A
label_11364: LD H,B
label_11365: LD H,C
label_11366: LD H,D
label_11367: LD H,E
label_11368: LD H,H
label_11369: LD H,L
label_11370: LD L,A
label_11371: LD L,B
label_11372: LD L,C
label_11373: LD L,D
label_11374: LD L,E
label_11375: LD L,H
label_11376: LD L,L
label_11377: LD A,$12
label_11378: LD B,$12
label_11379: LD C,$12
label_11380: LD D,$12
label_11381: LD E,$12
label_11382: LD H,$12
label_11383: LD L,$12
label_11384: LD A,(HL)
label_11385: LD B,(HL)
label_11386: LD C,(HL)
label_11387: LD D,(HL)
label_11388: LD E,(HL)
label_11389: LD H,(HL)
label_11390: LD L,(HL)
label_11391: LD A,(IX + 127)
label_11392: LD B,(IX + 127)
label_11393: LD C,(IX + 127)
label_11394: LD D,(IX + 127)
label_11395: LD E,(IX + 127)
label_11396: LD H,(IX + 127)
label_11397: LD L,(IX + 127)
label_11398: LD A,(IY - 128)
label_11399: LD B,(IY - 128)
label_11400: LD C,(IY - 128)
label_11401: LD D,(IY - 128)
label_11402: LD E,(IY - 128)
label_11403: LD H,(IY - 128)
label_11404: LD L,(IY - 128)
label_11405: LD (HL),A
label_11406: LD (HL),B
label_11407: LD (HL),C
label_11408: LD (HL),D
label_11409: LD (HL),E
label_11410: LD (HL),H
label_11411: LD (HL),L
label_11412: LD (IX + 127),A
label_11413: LD (IX + 127),B
label_11414: LD (IX + 127),C
label_11415: LD (IX + 127),D
label_11416: LD (IX + 127),E
label_11417: LD (IX + 127),H
label_11418: LD (IX + 127),L
label_11419: LD (IY - 128),A
label_11420: LD (IY - 128),B
label_11421: LD (IY - 128),C
label_11422: LD (IY - 128),D
label_11423: LD (IY - 128),E
label_11424: LD (IY - 128),H
label_11425: LD (IY - 128),L
label_11426: LD (HL),$12
label_11427: LD (IX + 127),$12
label_11428: LD (IY - 128),$12
label_11429: LD A,(BC)
label_11430: LD A,(DE)
label_11431: LD A,($5678)
label_11432: LD (BC),A
label_11433: LD (DE),A
label_11434: LD ($5678),A
label_11435: LD A,I
label_11436: LD A,R
label_11437: LD I,A
label_11438: LD R,A
label_11439: LD BC,$5678
label_11440: LD DE,$5678
label_11441: LD HL,$5678
label_11442: LD SP,$5678
label_11443: LD IX,$5678
label_11444: LD IY,$5678
label_11445: LD HL,($5678)
label_11446: LD BC,($5678)
label_11447: LD DE,($5678)
label_11448: LD HL,($5678)
label_11449: LD SP,($5678)
label_11450: LD IX,($5678)
label_11451: LD IY,($5678)
label_11452: LD ($5678),HL
label_11453: LD ($5678),BC
label_11454: LD ($5678),DE
label_11455: LD ($5678),HL
label_11456: LD ($5678),SP
label_11457: LD ($5678),IX
label_11458: LD ($5678),IY
label_11459: LD SP,HL
label_11460: LD SP,IX
label_11461: LD SP,IY
label_11462: PUSH BC
label_11463: PUSH DE
label_11464: PUSH HL
label_11465: PUSH AF
label_11466: PUSH IX
label_11467: PUSH IY
label_11468: POP BC
label_11469: POP DE
label_11470: POP HL
label_11471: POP AF
label_11472: POP IX
label_11473: POP IY
label_11474: EX DE,HL
label_11475: EX AF,AF'
label_11476: EXX
label_11477: EX (SP),HL
label_11478: EX (SP),IX
label_11479: EX (SP),IY
label_11480: LDI
label_11481: LDIR
label_11482: LDD
label_11483: LDDR
label_11484: CPI
label_11485: CPIR
label_11486: CPD
label_11487: CPDR
label_11488: ADD A,A
label_11489: ADD A,B
label_11490: ADD A,C
label_11491: ADD A,D
label_11492: ADD A,E
label_11493: ADD A,H
label_11494: ADD A,L
label_11495: ADD A,$12
label_11496: ADD A,(HL)
label_11497: ADD A,(IX + 127)
label_11498: ADD A,(IY - 128)
label_11499: ADC A,A
label_11500: ADC A,B
label_11501: ADC A,C
label_11502: ADC A,D
label_11503: ADC A,E
label_11504: ADC A,H
label_11505: ADC A,L
label_11506: ADC A,$12
label_11507: ADC A,(HL)
label_11508: ADC A,(IX + 127)
label_11509: ADC A,(IY - 128)
label_11510: SUB A
label_11511: SUB B
label_11512: SUB C
label_11513: SUB D
label_11514: SUB E
label_11515: SUB H
label_11516: SUB L
label_11517: SUB $12
label_11518: SUB (HL)
label_11519: SUB (IX + 127)
label_11520: SUB (IY - 128)
label_11521: SBC A,A
label_11522: SBC A,B
label_11523: SBC A,C
label_11524: SBC A,D
label_11525: SBC A,E
label_11526: SBC A,H
label_11527: SBC A,L
label_11528: SBC A,$12
label_11529: SBC A,(HL)
label_11530: SBC A,(IX + 127)
label_11531: SBC A,(IY - 128)
label_11532: AND A
label_11533: AND B
label_11534: AND C
label_11535: AND D
label_11536: AND E
label_11537: AND H
label_11538: AND L
label_11539: AND $12
label_11540: AND (HL)
label_11541: AND (IX + 127)
label_11542: AND (IY - 128)
label_11543: AND A
label_11544: AND B
label_11545: AND C
label_11546: AND D
label_11547: AND E
label_11548: AND H
label_11549: AND L
label_11550: AND $12
label_11551: AND (HL)
label_11552: AND (IX + 127)
label_11553: AND (IY - 128)
label_11554: OR A
label_11555: OR B
label_11556: OR C
label_11557: OR D
label_11558: OR E
label_11559: OR H
label_11560: OR L
label_11561: OR $12
label_11562: OR (HL)
label_11563: OR (IX + 127)
label_11564: OR (IY - 128)
label_11565: XOR A
label_11566: XOR B
label_11567: XOR C
label_11568: XOR D
label_11569: XOR E
label_11570: XOR H
label_11571: XOR L
label_11572: XOR $12
label_11573: XOR (HL)
label_11574: XOR (IX + 127)
label_11575: XOR (IY - 128)
label_11576: CP A
label_11577: CP B
label_11578: CP C
label_11579: CP D
label_11580: CP E
label_11581: CP H
label_11582: CP L
label_11583: CP $12
label_11584: CP (HL)
label_11585: CP (IX + 127)
label_11586: CP (IY - 128)
label_11587: INC A
label_11588: INC B
label_11589: INC C
label_11590: INC D
label_11591: INC E
label_11592: INC H
label_11593: INC L
label_11594: INC (HL)
label_11595: INC (IX + 127)
label_11596: INC (IY - 128)
label_11597: DEC A
label_11598: DEC B
label_11599: DEC C
label_11600: DEC D
label_11601: DEC E
label_11602: DEC H
label_11603: DEC L
label_11604: DEC (HL)
label_11605: DEC (IX + 127)
label_11606: DEC (IY - 128)
label_11607: DAA
label_11608: CPL
label_11609: NEG
label_11610: CCF
label_11611: SCF
label_11612: NOP
label_11613: HALT
label_11614: DI
label_11615: EI
label_11616: IM 0
label_11617: IM 1
label_11618: IM 2
label_11619: ADD HL,BC
label_11620: ADD HL,DE
label_11621: ADD HL,HL
label_11622: ADD HL,SP
label_11623: ADC HL,BC
label_11624: ADC HL,DE
label_11625: ADC HL,HL
label_11626: ADC HL,SP
label_11627: SBC HL,BC
label_11628: SBC HL,DE
label_11629: SBC HL,HL
label_11630: SBC HL,SP
label_11631: ADD IX,BC
label_11632: ADD IX,DE
label_11633: ADD IX,SP
label_11634: ADD IY,BC
label_11635: ADD IY,DE
label_11636: ADD IY,SP
label_11637: INC BC
label_11638: INC DE
label_11639: INC HL
label_11640: INC SP
label_11641: INC IX
label_11642: INC IY
label_11643: DEC BC
label_11644: DEC DE
label_11645: DEC HL
label_11646: DEC SP
label_11647: DEC IX
label_11648: DEC IY
label_11649: RLCA
label_11650: RLA
label_11651: RRCA
label_11652: RRA
label_11653: RLC A
label_11654: RLC B
label_11655: RLC C
label_11656: RLC D
label_11657: RLC E
label_11658: RLC H
label_11659: RLC L
label_11660: RLC (HL)
label_11661: RLC (IX + 127)
label_11662: RLC (IY - 128)
label_11663: RL A
label_11664: RL B
label_11665: RL C
label_11666: RL D
label_11667: RL E
label_11668: RL H
label_11669: RL L
label_11670: RL (HL)
label_11671: RL (IX + 127)
label_11672: RL (IY - 128)
label_11673: RRC A
label_11674: RRC B
label_11675: RRC C
label_11676: RRC D
label_11677: RRC E
label_11678: RRC H
label_11679: RRC L
label_11680: RRC (HL)
label_11681: RRC (IX + 127)
label_11682: RRC (IY - 128)
label_11683: RR A
label_11684: RR B
label_11685: RR C
label_11686: RR D
label_11687: RR E
label_11688: RR H
label_11689: RR L
label_11690: RR (HL)
label_11691: RR (IX + 127)
label_11692: RR (IY - 128)
label_11693: SLA A
label_11694: SLA B
label_11695: SLA C
label_11696: SLA D
label_11697: SLA E
label_11698: SLA H
label_11699: SLA L
label_11700: SLA (HL)
label_11701: SLA (IX + 127)
label_11702: SLA (IY - 128)
label_11703: SRA A
label_11704: SRA B
label_11705: SRA C
label_11706: SRA D
label_11707: SRA E
label_11708: SRA H
label_11709: SRA L
label_11710: SRA (HL)
label_11711: SRA (IX + 127)
label_11712: SRA (IY - 128)
label_11713: SRL A
label_11714: SRL B
label_11715: SRL C
label_11716: SRL D
label_11717: SRL E
label_11718: SRL H
label_11719: SRL L
label_11720: SRL (HL)
label_11721: SRL (IX + 127)
label_11722: SRL (IY - 128)
label_11723: RLD
label_11724: RRD
label_11725: BIT 0,A
label_11726: BIT 1,A
label_11727: BIT 2,A
label_11728: BIT 3,A
label_11729: BIT 4,A
label_11730: BIT 5,A
label_11731: BIT 6,A
label_11732: BIT 7,A
label_11733: BIT 0,B
label_11734: BIT 1,B
label_11735: BIT 2,B
label_11736: BIT 3,B
label_11737: BIT 4,B
label_11738: BIT 5,B
label_11739: BIT 6,B
label_11740: BIT 7,B
label_11741: BIT 0,C
label_11742: BIT 1,C
label_11743: BIT 2,C
label_11744: BIT 3,C
label_11745: BIT 4,C
label_11746: BIT 5,C
label_11747: BIT 6,C
label_11748: BIT 7,C
label_11749: BIT 0,D
label_11750: BIT 1,D
label_11751: BIT 2,D
label_11752: BIT 3,D
label_11753: BIT 4,D
label_11754: BIT 5,D
label_11755: BIT 6,D
label_11756: BIT 7,D
label_11757: BIT 0,E
label_11758: BIT 1,E
label_11759: BIT 2,E
label_11760: BIT 3,E
label_11761: BIT 4,E
label_11762: BIT 5,E
label_11763: BIT 6,E
label_11764: BIT 7,E
label_11765: BIT 0,H
label_11766: BIT 1,H
label_11767: BIT 2,H
label_11768: BIT 3,H
label_11769: BIT 4,H
label_11770: BIT 5,H
label_11771: BIT 6,H
label_11772: BIT 7,H
label_11773: BIT 0,L
label_11774: BIT 1,L
label_11775: BIT 2,L
label_11776: BIT 3,L
label_11777: BIT 4,L
label_11778: BIT 5,L
label_11779: BIT 6,L
label_11780: BIT 7,L
label_11781: BIT 0,(HL)
label_11782: BIT 1,(HL)
label_11783: BIT 2,(HL)
label_11784: BIT 3,(HL)
label_11785: BIT 4,(HL)
label_11786: BIT 5,(HL)
label_11787: BIT 6,(HL)
label_11788: BIT 7,(HL)
label_11789: BIT 0,(IX + 127)
label_11790: BIT 1,(IX + 127)
label_11791: BIT 2,(IX + 127)
label_11792: BIT 3,(IX + 127)
label_11793: BIT 4,(IX + 127)
label_11794: BIT 5,(IX + 127)
label_11795: BIT 6,(IX + 127)
label_11796: BIT 7,(IX + 127)
label_11797: BIT 0,(IY - 128)
label_11798: BIT 1,(IY - 128)
label_11799: BIT 2,(IY - 128)
label_11800: BIT 3,(IY - 128)
label_11801: BIT 4,(IY - 128)
label_11802: BIT 5,(IY - 128)
label_11803: BIT 6,(IY - 128)
label_11804: BIT 7,(IY - 128)
label_11805: SET 0,A
label_11806: SET 1,A
label_11807: SET 2,A
label_11808: SET 3,A
label_11809: SET 4,A
label_11810: SET 5,A
label_11811: SET 6,A
label_11812: SET 7,A
label_11813: SET 0,B
label_11814: SET 1,B
label_11815: SET 2,B
label_11816: SET 3,B
label_11817: SET 4,B
label_11818: SET 5,B
label_11819: SET 6,B
label_11820: SET 7,B
label_11821: SET 0,C
label_11822: SET 1,C
label_11823: SET 2,C
label_11824: SET 3,C
label_11825: SET 4,C
label_11826: SET 5,C
label_11827: SET 6,C
label_11828: SET 7,C
label_11829: SET 0,D
label_11830: SET 1,D
label_11831: SET 2,D
label_11832: SET 3,D
label_11833: SET 4,D
label_11834: SET 5,D
label_11835: SET 6,D
label_11836: SET 7,D
label_11837: SET 0,E
label_11838: SET 1,E
label_11839: SET 2,E
label_11840: SET 3,E
label_11841: SET 4,E
label_11842: SET 5,E
label_11843: SET 6,E
label_11844: SET 7,E
label_11845: SET 0,H
label_11846: SET 1,H
label_11847: SET 2,H
label_11848: SET 3,H
label_11849: SET 4,H
label_11850: SET 5,H
label_11851: SET 6,H
label_11852: SET 7,H
label_11853: SET 0,L
label_11854: SET 1,L
label_11855: SET 2,L
label_11856: SET 3,L
label_11857: SET 4,L
label_11858: SET 5,L
label_11859: SET 6,L
label_11860: SET 7,L
label_11861: SET 0,(HL)
label_11862: SET 1,(HL)
label_11863: SET 2,(HL)
label_11864: SET 3,(HL)
label_11865: SET 4,(HL)
label_11866: SET 5,(HL)
label_11867: SET 6,(HL)
label_11868: SET 7,(HL)
label_11869: SET 0,(IX + 127)
label_11870: SET 1,(IX + 127)
label_11871: SET 2,(IX + 127)
label_11872: SET 3,(IX + 127)
label_11873: SET 4,(IX + 127)
label_11874: SET 5,(IX + 127)
label_11875: SET 6,(IX + 127)
label_11876: SET 7,(IX + 127)
label_11877: SET 0,(IY - 128)
label_11878: SET 1,(IY - 128)
label_11879: SET 2,(IY - 128)
label_11880: SET 3,(IY - 128)
label_11881: SET 4,(IY - 128)
label_11882: SET 5,(IY - 128)
label_11883: SET 6,(IY - 128)
label_11884: SET 7,(IY - 128)
label_11885: RES 0,A
label_11886: RES 1,A
label_11887: RES 2,A
label_11888: RES 3,A
label_11889: RES 4,A
label_11890: RES 5,A
label_11891: RES 6,A
label_11892: RES 7,A
label_11893: RES 0,B
label_11894: RES 1,B
label_11895: RES 2,B
label_11896: RES 3,B
label_11897: RES 4,B
label_11898: RES 5,B
label_11899: RES 6,B
label_11900: RES 7,B
label_11901: RES 0,C
label_11902: RES 1,C
label_11903: RES 2,C
label_11904: RES 3,C
label_11905: RES 4,C
label_11906: RES 5,C
label_11907: RES 6,C
label_11908: RES 7,C
label_11909: RES 0,D
label_11910: RES 1,D
label_11911: RES 2,D
label_11912: RES 3,D
label_11913: RES 4,D
label_11914: RES 5,D
label_11915: RES 6,D
label_11916: RES 7,D
label_11917: RES 0,E
label_11918: RES 1,E
label_11919: RES 2,E
label_11920: RES 3,E
label_11921: RES 4,E
label_11922: RES 5,E
label_11923: RES 6,E
label_11924: RES 7,E
label_11925: RES 0,H
label_11926: RES 1,H
label_11927: RES 2,H
label_11928: RES 3,H
label_11929: RES 4,H
label_11930: RES 5,H
label_11931: RES 6,H
label_11932: RES 7,H
label_11933: RES 0,L
label_11934: RES 1,L
label_11935: RES 2,L
label_11936: RES 3,L
label_11937: RES 4,L
label_11938: RES 5,L
label_11939: RES 6,L
label_11940: RES 7,L
label_11941: RES 0,(HL)
label_11942: RES 1,(HL)
label_11943: RES 2,(HL)
label_11944: RES 3,(HL)
label_11945: RES 4,(HL)
label_11946: RES 5,(HL)
label_11947: RES 6,(HL)
label_11948: RES 7,(HL)
label_11949: RES 0,(IX + 127)
label_11950: RES 1,(IX + 127)
label_11951: RES 2,(IX + 127)
label_11952: RES 3,(IX + 127)
label_11953: RES 4,(IX + 127)
label_11954: RES 5,(IX + 127)
label_11955: RES 6,(IX + 127)
label_11956: RES 7,(IX + 127)
label_11957: RES 0,(IY - 128)
label_11958: RES 1,(IY - 128)
label_11959: RES 2,(IY - 128)
label_11960: RES 3,(IY - 128)
label_11961: RES 4,(IY - 128)
label_11962: RES 5,(IY - 128)
label_11963: RES 6,(IY - 128)
label_11964: RES 7,(IY - 128)
label_11965: JP $5678
label_11966: JP NZ,$5678
label_11967: JP Z,$5678
label_11968: JP NC,$5678
label_11969: JP C,$5678
label_11970: JP PO,$5678
label_11971: JP PE,$5678
label_11972: JP P,$5678
label_11973: JP M,$5678
label_11974: JR $ + 2
label_11975: JR NZ,$ + 2
label_11976: JR Z,$ + 2
label_11977: JR NC,$ + 2
label_11978: JR C,$ + 2
label_11979: JP (HL)
label_11980: JP (IX)
label_11981: JP (IY)
label_11982: DJNZ $ + 2
label_11983: CALL $5678
label_11984: CALL NZ,$5678
label_11985: CALL Z,$5678
label_11986: CALL NC,$5678
label_11987: CALL C,$5678
label_11988: CALL PO,$5678
label_11989: CALL PE,$5678
label_11990: CALL P,$5678
label_11991: CALL M,$5678
label_11992: RET
label_11993: RET NZ
label_11994: RET Z
label_11995: RET NC
label_11996: RET C
label_11997: RET PO
label_11998: RET PE
label_11999: RET P
label_12000: RET M
label_12001: RETI
label_12002: RETN
label_12003: RST $00
label_12004: RST $08
label_12005: RST $10
label_12006: RST $18
label_12007: RST $20
label_12008: RST $28
label_12009: RST $30
label_12010: RST $38
label_12011: IN A,($12)
label_12012: IN A,(C)
label_12013: IN B,(C)
label_12014: IN C,(C)
label_12015: IN D,(C)
label_12016: IN E,(C)
label_12017: IN H,(C)
label_12018: IN L,(C)
label_12019: IN F,(C)
label_12020: INI
label_12021: INIR
label_12022: IND
label_12023: INDR
label_12024: OUT ($12),A
label_12025: OUT (C),A
label_12026: OUT (C),B
label_12027: OUT (C),C
label_12028: OUT (C),D
label_12029: OUT (C),E
label_12030: OUT (C),H
label_12031: OUT (C),L
label_12032: OUTI
label_12033: OTIR
label_12034: OUTD
label_12035: OTDR
label_12036: LD A,A
label_12037: LD A,B
label_12038: LD A,C
label_12039: LD A,D
label_12040: LD A,E
label_12041: LD A,H
label_12042: LD A,L
label_12043: LD B,A
label_12044: LD B,B
label_12045: LD B,C
label_12046: LD B,D
label_12047: LD B,E
label_12048: LD B,H
label_12049: LD B,L
label_12050: LD C,A
label_12051: LD C,B
label_12052: LD C,C
label_12053: LD C,D
label_12054: LD C,E
label_12055: LD C,H
label_12056: LD C,L
label_12057: LD D,A
label_12058: LD D,B
label_12059: LD D,C
label_12060: LD D,D
label_12061: LD D,E
label_12062: LD D,H
label_12063: LD D,L
label_12064: LD E,A
label_12065: LD E,B
label_12066: LD E,C
label_12067: LD E,D
label_12068: LD E,E
label_12069: LD E,H
label_12070: LD E,L
label_12071: LD H,A
label_12072: LD H,B
label_12073: LD H,C
label_12074: LD H,D
label_12075: LD H,E
label_12076: LD H,H
label_12077: LD H,L
label_12078: LD L,A
label_12079: LD L,B
label_12080: LD L,C
label_12081: LD L,D
label_12082: LD L,E
label_12083: LD L,H
label_12084: LD L,L
label_12085: LD A,$12
label_12086: LD B,$12
label_12087: LD C,$12
label_12088: LD D,$12
label_12089: LD E,$12
label_12090: LD H,$12
label_12091: LD L,$12
label_12092: LD A,(HL)
label_12093: LD B,(HL)
label_12094: LD C,(HL)
label_12095: LD D,(HL)
label_12096: LD E,(HL)
label_12097: LD H,(HL)
label_12098: LD L,(HL)
label_12099: LD A,(IX + 127)
label_12100: LD B,(IX + 127)
label_12101: LD C,(IX + 127)
label_12102: LD D,(IX + 127)
label_12103: LD E,(IX + 127)
label_12104: LD H,(IX + 127)
label_12105: LD L,(IX + 127)
label_12106: LD A,(IY - 128)
label_12107: LD B,(IY - 128)
label_12108: LD C,(IY - 128)
label_12109: LD D,(IY - 128)
label_12110: LD E,(IY - 128)
label_12111: LD H,(IY - 128)
label_12112: LD L,(IY - 128)
label_12113: LD (HL),A
label_12114: LD (HL),B
label_12115: LD (HL),C
label_12116: LD (HL),D
label_12117: LD (HL),E
label_12118: LD (HL),H
label_12119: LD (HL),L
label_12120: LD (IX + 127),A
label_12121: LD (IX + 127),B
label_12122: LD (IX + 127),C
label_12123: LD (IX + 127),D
label_12124: LD (IX + 127),E
label_12125: LD (IX + 127),H
label_12126: LD (IX + 127),L
label_12127: LD (IY - 128),A
label_12128: LD (IY - 128),B
label_12129: LD (IY - 128),C
label_12130: LD (IY - 128),D
label_12131: LD (IY - 128),E
label_12132: LD (IY - 128),H
label_12133: LD (IY - 128),L
label_12134: LD (HL),$12
label_12135: LD (IX + 127),$12
label_12136: LD (IY - 128),$12
label_12137: LD A,(BC)
label_12138: LD A,(DE)
label_12139: LD A,($5678)
label_12140: LD (BC),A
label_12141: LD (DE),A
label_12142: LD ($5678),A
label_12143: LD A,I
label_12144: LD A,R
label_12145: LD I,A
label_12146: LD R,A
label_12147: LD BC,$5678
label_12148: LD DE,$5678
label_12149: LD HL,$5678
label_12150: LD SP,$5678
label_12151: LD IX,$5678
label_12152: LD IY,$5678
label_12153: LD HL,($5678)
label_12154: LD BC,($5678)
label_12155: LD DE,($5678)
label_12156: LD HL,($5678)
label_12157: LD SP,($5678)
label_12158: LD IX,($5678)
label_12159: LD IY,($5678)
label_12160: LD ($5678),HL
label_12161: LD ($5678),BC
label_12162: LD ($5678),DE
label_12163: LD ($5678),HL
label_12164: LD ($5678),SP
label_12165: LD ($5678),IX
label_12166: LD ($5678),IY
label_12167: LD SP,HL
label_12168: LD SP,IX
label_12169: LD SP,IY
label_12170: PUSH BC
label_12171: PUSH DE
label_12172: PUSH HL
label_12173: PUSH AF
label_12174: PUSH IX
label_12175: PUSH IY
label_12176: POP BC
label_12177: POP DE
label_12178: POP HL
label_12179: POP AF
label_12180: POP IX
label_12181: POP IY
label_12182: EX DE,HL
label_12183: EX AF,AF'
label_12184: EXX
label_12185: EX (SP),HL
label_12186: EX (SP),IX
label_12187: EX (SP),IY
label_12188: LDI
label_12189: LDIR
label_12190: LDD
label_12191: LDDR
label_12192: CPI
label_12193: CPIR
label_12194: CPD
label_12195: CPDR
label_12196: ADD A,A
label_12197: ADD A,B
label_12198: ADD A,C
label_12199: ADD A,D
label_12200: ADD A,E
label_12201: ADD A,H
label_12202: ADD A,L
label_12203: ADD A,$12
label_12204: ADD A,(HL)
label_12205: ADD A,(IX + 127)
label_12206: ADD A,(IY - 128)
label_12207: ADC A,A
label_12208: ADC A,B
label_12209: ADC A,C
label_12210: ADC A,D
label_12211: ADC A,E
label_12212: ADC A,H
label_12213: ADC A,L
label_12214: ADC A,$12
label_12215: ADC A,(HL)
label_12216: ADC A,(IX + 127)
label_12217: ADC A,(IY - 128)
label_12218: SUB A
label_12219: SUB B
label_12220: SUB C
label_12221: SUB D
label_12222: SUB E
label_12223: SUB H
label_12224: SUB L
label_12225: SUB $12
label_12226: SUB (HL)
label_12227: SUB (IX + 127)
label_12228: SUB (IY - 128)
label_12229: SBC A,A
label_12230: SBC A,B
label_12231: SBC A,C
label_12232: SBC A,D
label_12233: SBC A,E
label_12234: SBC A,H
label_12235: SBC A,L
label_12236: SBC A,$12
label_12237: SBC A,(HL)
label_12238: SBC A,(IX + 127)
label_12239: SBC A,(IY - 128)
label_12240: AND A
label_12241: AND B
label_12242: AND C
label_12243: AND D
label_12244: AND E
label_12245: AND H
label_12246: AND L
label_12247: AND $12
label_12248: AND (HL)
label_12249: AND (IX + 127)
label_12250: AND (IY - 128)
label_12251: AND A
label_12252: AND B
label_12253: AND C
label_12254: AND D
label_12255: AND E
label_12256: AND H
label_12257: AND L
label_12258: AND $12
label_12259: AND (HL)
label_12260: AND (IX + 127)
label_12261: AND (IY - 128)
label_12262: OR A
label_12263: OR B
label_12264: OR C
label_12265: OR D
label_12266: OR E
label_12267: OR H
label_12268: OR L
label_12269: OR $12
label_12270: OR (HL)
label_12271: OR (IX + 127)
label_12272: OR (IY - 128)
label_12273: XOR A
label_12274: XOR B
label_12275: XOR C
label_12276: XOR D
label_12277: XOR E
label_12278: XOR H
label_12279: XOR L
label_12280: XOR $12
label_12281: XOR (HL)
label_12282: XOR (IX + 127)
label_12283: XOR (IY - 128)
label_12284: CP A
label_12285: CP B
label_12286: CP C
label_12287: CP D
label_12288: CP E
label_12289: CP H
label_12290: CP L
label_12291: CP $12
label_12292: CP (HL)
label_12293: CP (IX + 127)
label_12294: CP (IY - 128)
label_12295: INC A
label_12296: INC B
label_12297: INC C
label_12298: INC D
label_12299: INC E
label_12300: INC H
label_12301: INC L
label_12302: INC (HL)
label_12303: INC (IX + 127)
label_12304: INC (IY - 128)
label_12305: DEC A
label_12306: DEC B
label_12307: DEC C
label_12308: DEC D
label_12309: DEC E
label_12310: DEC H
label_12311: DEC L
label_12312: DEC (HL)
label_12313: DEC (IX + 127)
label_12314: DEC (IY - 128)
label_12315: DAA
label_12316: CPL
label_12317: NEG
label_12318: CCF
label_12319: SCF
label_12320: NOP
label_12321: HALT
label_12322: DI
label_12323: EI
label_12324: IM 0
label_12325: IM 1
label_12326: IM 2
label_12327: ADD HL,BC
label_12328: ADD HL,DE
label_12329: ADD HL,HL
label_12330: ADD HL,SP
label_12331: ADC HL,BC
label_12332: ADC HL,DE
label_12333: ADC HL,HL
label_12334: ADC HL,SP
label_12335: SBC HL,BC
label_12336: SBC HL,DE
label_12337: SBC HL,HL
label_12338: SBC HL,SP
label_12339: ADD IX,BC
label_12340: ADD IX,DE
label_12341: ADD IX,SP
label_12342: ADD IY,BC
label_12343: ADD IY,DE
label_12344: ADD IY,SP
label_12345: INC BC
label_12346: INC DE
label_12347: INC HL
label_12348: INC SP
label_12349: INC IX
label_12350: INC IY
label_12351: DEC BC
label_12352: DEC DE
label_12353: DEC HL
label_12354: DEC SP
label_12355: DEC IX
label_12356: DEC IY
label_12357: RLCA
label_12358: RLA
label_12359: RRCA
label_12360: RRA
label_12361: RLC A
label_12362: RLC B
label_12363: RLC C
label_12364: RLC D
label_12365: RLC E
label_12366: RLC H
label_12367: RLC L
label_12368: RLC (HL)
label_12369: RLC (IX + 127)
label_12370: RLC (IY - 128)
label_12371: RL A
label_12372: RL B
label_12373: RL C
label_12374: RL D
label_12375: RL E
label_12376: RL H
label_12377: RL L
label_12378: RL (HL)
label_12379: RL (IX + 127)
label_12380: RL (IY - 128)
label_12381: RRC A
label_12382: RRC B
label_12383: RRC C
label_12384: RRC D
label_12385: RRC E
label_12386: RRC H
label_12387: RRC L
label_12388: RRC (HL)
label_12389: RRC (IX + 127)
label_12390: RRC (IY - 128)
label_12391: RR A
label_12392: RR B
label_12393: RR C
label_12394: RR D
label_12395: RR E
label_12396: RR H
label_12397: RR L
label_12398: RR (HL)
label_12399: RR (IX + 127)
label_12400: RR (IY - 128)
label_12401: SLA A
label_12402: SLA B
label_12403: SLA C
label_12404: SLA D
label_12405: SLA E
label_12406: SLA H
label_12407: SLA L
label_12408: SLA (HL)
label_12409: SLA (IX + 127)
label_12410: SLA (IY - 128)
label_12411: SRA A
label_12412: SRA B
label_12413: SRA C
label_12414: SRA D
label_12415: SRA E
label_12416: SRA H
label_12417: SRA L
label_12418: SRA (HL)
label_12419: SRA (IX + 127)
label_12420: SRA (IY - 128)
label_12421: SRL A
label_12422: SRL B
label_12423: SRL C
label_12424: SRL D
label_12425: SRL E
label_12426: SRL H
label_12427: SRL L
label_12428: SRL (HL)
label_12429: SRL (IX + 127)
label_12430: SRL (IY - 128)
label_12431: RLD
label_12432: RRD
label_12433: BIT 0,A
label_12434: BIT 1,A
label_12435: BIT 2,A
label_12436: BIT 3,A
label_12437: BIT 4,A
label_12438: BIT 5,A
label_12439: BIT 6,A
label_12440: BIT 7,A
label_12441: BIT 0,B
label_12442: BIT 1,B
label_12443: BIT 2,B
label_12444: BIT 3,B
label_12445: BIT 4,B
label_12446: BIT 5,B
label_12447: BIT 6,B
label_12448: BIT 7,B
label_12449: BIT 0,C
label_12450: BIT 1,C
label_12451: BIT 2,C
label_12452: BIT 3,C
label_12453: BIT 4,C
label_12454: BIT 5,C
label_12455: BIT 6,C
label_12456: BIT 7,C
label_12457: BIT 0,D
label_12458: BIT 1,D
label_12459: BIT 2,D
label_12460: BIT 3,D
label_12461: BIT 4,D
label_12462: BIT 5,D
label_12463: BIT 6,D
label_12464: BIT 7,D
label_12465: BIT 0,E
label_12466: BIT 1,E
label_12467: BIT 2,E
label_12468: BIT 3,E
label_12469: BIT 4,E
label_12470: BIT 5,E
label_12471: BIT 6,E
label_12472: BIT 7,E
label_12473: BIT 0,H
label_12474: BIT 1,H
label_12475: BIT 2,H
label_12476: BIT 3,H
label_12477: BIT 4,H
label_12478: BIT 5,H
label_12479: BIT 6,H
label_12480: BIT 7,H
label_12481: BIT 0,L
label_12482: BIT 1,L
label_12483: BIT 2,L
label_12484: BIT 3,L
label_12485: BIT 4,L
label_12486: BIT 5,L
label_12487: BIT 6,L
label_12488: BIT 7,L
label_12489: BIT 0,(HL)
label_12490: BIT 1,(HL)
label_12491: BIT 2,(HL)
label_12492: BIT 3,(HL)
label_12493: BIT 4,(HL)
label_12494: BIT 5,(HL)
label_12495: BIT 6,(HL)
label_12496: BIT 7,(HL)
label_12497: BIT 0,(IX + 127)
label_12498: BIT 1,(IX + 127)
label_12499: BIT 2,(IX + 127)
label_12500: BIT 3,(IX + 127)
label_12501: BIT 4,(IX + 127)
label_12502: BIT 5,(IX + 127)
label_12503: BIT 6,(IX + 127)
label_12504: BIT 7,(IX + 127)
label_12505: BIT 0,(IY - 128)
label_12506: BIT 1,(IY - 128)
label_12507: BIT 2,(IY - 128)
label_12508: BIT 3,(IY - 128)
label_12509: BIT 4,(IY - 128)
label_12510: BIT 5,(IY - 128)
label_12511: BIT 6,(IY - 128)
label_12512: BIT 7,(IY - 128)
label_12513: SET 0,A
label_12514: SET 1,A
label_12515: SET 2,A
label_12516: SET 3,A
label_12517: SET 4,A
label_12518: SET 5,A
label_12519: SET 6,A
label_12520: SET 7,A
label_12521: SET 0,B
label_12522: SET 1,B
label_12523: SET 2,B
label_12524: SET 3,B
label_12525: SET 4,B
label_12526: SET 5,B
label_12527: SET 6,B
label_12528: SET 7,B
label_12529: SET 0,C
label_12530: SET 1,C
label_12531: SET 2,C
label_12532: SET 3,C
label_12533: SET 4,C
label_12534: SET 5,C
label_12535: SET 6,C
label_12536: SET 7,C
label_12537: SET 0,D
label_12538: SET 1,D
label_12539: SET 2,D
label_12540: SET 3,D
label_12541: SET 4,D
label_12542: SET 5,D
label_12543: SET 6,D
label_12544: SET 7,D
label_12545: SET 0,E
label_12546: SET 1,E
label_12547: SET 2,E
label_12548: SET 3,E
label_12549: SET 4,E
label_12550: SET 5,E
label_12551: SET 6,E
label_12552: SET 7,E
label_12553: SET 0,H
label_12554: SET 1,H
label_12555: SET 2,H
label_12556: SET 3,H
label_12557: SET 4,H
label_12558: SET 5,H
label_12559: SET 6,H
label_12560: SET 7,H
label_12561: SET 0,L
label_12562: SET 1,L
label_12563: SET 2,L
label_12564: SET 3,L
label_12565: SET 4,L
label_12566: SET 5,L
label_12567: SET 6,L
label_12568: SET 7,L
label_12569: SET 0,(HL)
label_12570: SET 1,(HL)
label_12571: SET 2,(HL)
label_12572: SET 3,(HL)
label_12573: SET 4,(HL)
label_12574: SET 5,(HL)
label_12575: SET 6,(HL)
label_12576: SET 7,(HL)
label_12577: SET 0,(IX + 127)
label_12578: SET 1,(IX + 127)
label_12579: SET 2,(IX + 127)
label_12580: SET 3,(IX + 127)
label_12581: SET 4,(IX + 127)
label_12582: SET 5,(IX + 127)
label_12583: SET 6,(IX + 127)
label_12584: SET 7,(IX + 127)
label_12585: SET 0,(IY - 128)
label_12586: SET 1,(IY - 128)
label_12587: SET 2,(IY - 128)
label_12588: SET 3,(IY - 128)
label_12589: SET 4,(IY - 128)
label_12590: SET 5,(IY - 128)
label_12591: SET 6,(IY - 128)
label_12592: SET 7,(IY - 128)
label_12593: RES 0,A
label_12594: RES 1,A
label_12595: RES 2,A
label_12596: RES 3,A
label_12597: RES 4,A
label_12598: RES 5,A
label_12599: RES 6,A
label_12600: RES 7,A
label_12601: RES 0,B
label_12602: RES 1,B
label_12603: RES 2,B
label_12604: RES 3,B
label_12605: RES 4,B
label_12606: RES 5,B
label_12607: RES 6,B
label_12608: RES 7,B
label_12609: RES 0,C
label_12610: RES 1,C
label_12611: RES 2,C
label_12612: RES 3,C
label_12613: RES 4,C
label_12614: RES 5,C
label_12615: RES 6,C
label_12616: RES 7,C
label_12617: RES 0,D
label_12618: RES 1,D
label_12619: RES 2,D
label_12620: RES 3,D
label_12621: RES 4,D
label_12622: RES 5,D
label_12623: RES 6,D
label_12624: RES 7,D
label_12625: RES 0,E
label_12626: RES 1,E
label_12627: RES 2,E
label_12628: RES 3,E
label_12629: RES 4,E
label_12630: RES 5,E
label_12631: RES 6,E
label_12632: RES 7,E
label_12633: RES 0,H
label_12634: RES 1,H
label_12635: RES 2,H
label_12636: RES 3,H
label_12637: RES 4,H
label_12638: RES 5,H
label_12639: RES 6,H
label_12640: RES 7,H
label_12641: RES 0,L
label_12642: RES 1,L
label_12643: RES 2,L
label_12644: RES 3,L
label_12645: RES 4,L
label_12646: RES 5,L
label_12647: RES 6,L
label_12648: RES 7,L
label_12649: RES 0,(HL)
label_12650: RES 1,(HL)
label_12651: RES 2,(HL)
label_12652: RES 3,(HL)
label_12653: RES 4,(HL)
label_12654: RES 5,(HL)
label_12655: RES 6,(HL)
label_12656: RES 7,(HL)
label_12657: RES 0,(IX + 127)
label_12658: RES 1,(IX + 127)
label_12659: RES 2,(IX + 127)
label_12660: RES 3,(IX + 127)
label_12661: RES 4,(IX + 127)
label_12662: RES 5,(IX + 127)
label_12663: RES 6,(IX + 127)
label_12664: RES 7,(IX + 127)
label_12665: RES 0,(IY - 128)
label_12666: RES 1,(IY - 128)
label_12667: RES 2,(IY - 128)
label_12668: RES 3,(IY - 128)
label_12669: RES 4,(IY - 128)
label_12670: RES 5,(IY - 128)
label_12671: RES 6,(IY - 128)
label_12672: RES 7,(IY - 128)
label_12673: JP $5678
label_12674: JP NZ,$5678
label_12675: JP Z,$5678
label_12676: JP NC,$5678
label_12677: JP C,$5678
label_12678: JP PO,$5678
label_12679: JP PE,$5678
label_12680: JP P,$5678
label_12681: JP M,$5678
label_12682: JR $ + 2
label_12683: JR NZ,$ + 2
label_12684: JR Z,$ + 2
label_12685: JR NC,$ + 2
label_12686: JR C,$ + 2
label_12687: JP (HL)
label_12688: JP (IX)
label_12689: JP (IY)
label_12690: DJNZ $ + 2
label_12691: CALL $5678
label_12692: CALL NZ,$5678
label_12693: CALL Z,$5678
label_12694: CALL NC,$5678
label_12695: CALL C,$5678
label_12696: CALL PO,$5678
label_12697: CALL PE,$5678
label_12698: CALL P,$5678
label_12699: CALL M,$5678
label_12700: RET
label_12701: RET NZ
label_12702: RET Z
label_12703: RET NC
label_12704: RET C
label_12705: RET PO
label_12706: RET PE
label_12707: RET P
label_12708: RET M
label_12709: RETI
label_12710: RETN
label_12711: RST $00
label_12712: RST $08
label_12713: RST $10
label_12714: RST $18
label_12715: RST $20
label_12716: RST $28
label_12717: RST $30
label_12718: RST $38
label_12719: IN A,($12)
label_12720: IN A,(C)
label_12721: IN B,(C)
label_12722: IN C,(C)
label_12723: IN D,(C)
label_12724: IN E,(C)
label_12725: IN H,(C)
label_12726: IN L,(C)
label_12727: IN F,(C)
label_12728: INI
label_12729: INIR
label_12730: IND
label_12731: INDR
label_12732: OUT ($12),A
label_12733: OUT (C),A
label_12734: OUT (C),B
label_12735: OUT (C),C
label_12736: OUT (C),D
label_12737: OUT (C),E
label_12738: OUT (C),H
label_12739: OUT (C),L
label_12740: OUTI
label_12741: OTIR
label_12742: OUTD
label_12743: OTDR
label_12744: LD A,A
label_12745: LD A,B
label_12746: LD A,C
label_12747: LD A,D
label_12748: LD A,E
label_12749: LD A,H
label_12750: LD A,L
label_12751: LD B,A
label_12752: LD B,B
label_12753: LD B,C
label_12754: LD B,D
label_12755: LD B,E
label_12756: LD B,H
label_12757: LD B,L
label_12758: LD C,A
label_12759: LD C,B
label_12760: LD C,C
label_12761: LD C,D
label_12762: LD C,E
label_12763: LD C,H
label_12764: LD C,L
label_12765: LD D,A
label_12766: LD D,B
label_12767: LD D,C
label_12768: LD D,D
label_12769: LD D,E
label_12770: LD D,H
label_12771: LD D,L
label_12772: LD E,A
label_12773: LD E,B
label_12774: LD E,C
label_12775: LD E,D
label_12776: LD E,E
label_12777: LD E,H
label_12778: LD E,L
label_12779: LD H,A
label_12780: LD H,B
label_12781: LD H,C
label_12782: LD H,D
label_12783: LD H,E
label_12784: LD H,H
label_12785: LD H,L
label_12786: LD L,A
label_12787: LD L,B
label_12788: LD L,C
label_12789: LD L,D
label_12790: LD L,E
label_12791: LD L,H
label_12792: LD L,L
label_12793: LD A,$12
label_12794: LD B,$12
label_12795: LD C,$12
label_12796: LD D,$12
label_12797: LD E,$12
label_12798: LD H,$12
label_12799: LD L,$12
label_12800: LD A,(HL)
label_12801: LD B,(HL)
label_12802: LD C,(HL)
label_12803: LD D,(HL)
label_12804: LD E,(HL)
label_12805: LD H,(HL)
label_12806: LD L,(HL)
label_12807: LD A,(IX + 127)
label_12808: LD B,(IX + 127)
label_12809: LD C,(IX + 127)
label_12810: LD D,(IX + 127)
label_12811: LD E,(IX + 127)
label_12812: LD H,(IX + 127)
label_12813: LD L,(IX + 127)
label_12814: LD A,(IY - 128)
label_12815: LD B,(IY - 128)
label_12816: LD C,(IY - 128)
label_12817: LD D,(IY - 128)
label_12818: LD E,(IY - 128)
label_12819: LD H,(IY - 128)
label_12820: LD L,(IY - 128)
label_12821: LD (HL),A
label_12822: LD (HL),B
label_12823: LD (HL),C
label_12824: LD (HL),D
label_12825: LD (HL),E
label_12826: LD (HL),H
label_12827: LD (HL),L
label_12828: LD (IX + 127),A
label_12829: LD (IX + 127),B
label_12830: LD (IX + 127),C
label_12831: LD (IX + 127),D
label_12832: LD (IX + 127),E
label_12833: LD (IX + 127),H
label_12834: LD (IX + 127),L
label_12835: LD (IY - 128),A
label_12836: LD (IY - 128),B
label_12837: LD (IY - 128),C
label_12838: LD (IY - 128),D
label_12839: LD (IY - 128),E
label_12840: LD (IY - 128),H
label_12841: LD (IY - 128),L
label_12842: LD (HL),$12
label_12843: LD (IX + 127),$12
label_12844: LD (IY - 128),$12
label_12845: LD A,(BC)
label_12846: LD A,(DE)
label_12847: LD A,($5678)
label_12848: LD (BC),A
label_12849: LD (DE),A
label_12850: LD ($5678),A
label_12851: LD A,I
label_12852: LD A,R
label_12853: LD I,A
label_12854: LD R,A
label_12855: LD BC,$5678
label_12856: LD DE,$5678
label_12857: LD HL,$5678
label_12858: LD SP,$5678
label_12859: LD IX,$5678
label_12860: LD IY,$5678
label_12861: LD HL,($5678)
label_12862: LD BC,($5678)
label_12863: LD DE,($5678)
label_12864: LD HL,($5678)
label_12865: LD SP,($5678)
label_12866: LD IX,($5678)
label_12867: LD IY,($5678)
label_12868: LD ($5678),HL
label_12869: LD ($5678),BC
label_12870: LD ($5678),DE
label_12871: LD ($5678),HL
label_12872: LD ($5678),SP
label_12873: LD ($5678),IX
label_12874: LD ($5678),IY
label_12875: LD SP,HL
label_12876: LD SP,IX
label_12877: LD SP,IY
label_12878: PUSH BC
label_12879: PUSH DE
label_12880: PUSH HL
label_12881: PUSH AF
label_12882: PUSH IX
label_12883: PUSH IY
label_12884: POP BC
label_12885: POP DE
label_12886: POP HL
label_12887: POP AF
label_12888: POP IX
label_12889: POP IY
label_12890: EX DE,HL
label_12891: EX AF,AF'
label_12892: EXX
label_12893: EX (SP),HL
label_12894: EX (SP),IX
label_12895: EX (SP),IY
label_12896: LDI
label_12897: LDIR
label_12898: LDD
label_12899: LDDR
label_12900: CPI
label_12901: CPIR
label_12902: CPD
label_12903: CPDR
label_12904: ADD A,A
label_12905: ADD A,B
label_12906: ADD A,C
label_12907: ADD A,D
label_12908: ADD A,E
label_12909: ADD A,H
label_12910: ADD A,L
label_12911: ADD A,$12
label_12912: ADD A,(HL)
label_12913: ADD A,(IX + 127)
label_12914: ADD A,(IY - 128)
label_12915: ADC A,A
label_12916: ADC A,B
label_12917: ADC A,C
label_12918: ADC A,D
label_12919: ADC A,E
label_12920: ADC A,H
label_12921: ADC A,L
label_12922: ADC A,$12
label_12923: ADC A,(HL)
label_12924: ADC A,(IX + 127)
label_12925: ADC A,(IY - 128)
label_12926: SUB A
label_12927: SUB B
label_12928: SUB C
label_12929: SUB D
label_12930: SUB E
label_12931: SUB H
label_12932: SUB L
label_12933: SUB $12
label_12934: SUB (HL)
label_12935: SUB (IX + 127)
label_12936: SUB (IY - 128)
label_12937: SBC A,A
label_12938: SBC A,B
label_12939: SBC A,C
label_12940: SBC A,D
label_12941: SBC A,E
label_12942: SBC A,H
label_12943: SBC A,L
label_12944: SBC A,$12
label_12945: SBC A,(HL)
label_12946: SBC A,(IX + 127)
label_12947: SBC A,(IY - 128)
label_12948: AND A
label_12949: AND B
label_12950: AND C
label_12951: AND D
label_12952: AND E
label_12953: AND H
label_12954: AND L
label_12955: AND $12
label_12956: AND (HL)
label_12957: AND (IX + 127)
label_12958: AND (IY - 128)
label_12959: AND A
label_12960: AND B
label_12961: AND C
label_12962: AND D
label_12963: AND E
label_12964: AND H
label_12965: AND L
label_12966: AND $12
label_12967: AND (HL)
label_12968: AND (IX + 127)
label_12969: AND (IY - 128)
label_12970: OR A
label_12971: OR B
label_12972: OR C
label_12973: OR D
label_12974: OR E
label_12975: OR H
label_12976: OR L
label_12977: OR $12
label_12978: OR (HL)
label_12979: OR (IX + 127)
label_12980: OR (IY - 128)
label_12981: XOR A
label_12982: XOR B
label_12983: XOR C
label_12984: XOR D
label_12985: XOR E
label_12986: XOR H
label_12987: XOR L
label_12988: XOR $12
label_12989: XOR (HL)
label_12990: XOR (IX + 127)
label_12991: XOR (IY - 128)
label_12992: CP A
label_12993: CP B
label_12994: CP C
label_12995: CP D
label_12996: CP E
label_12997: CP H
label_12998: CP L
label_12999: CP $12
label_13000: CP (HL)
label_13001: CP (IX + 127)
label_13002: CP (IY - 128)
label_13003: INC A
label_13004: INC B
label_13005: INC C
label_13006: INC D
label_13007: INC E
label_13008: INC H
label_13009: INC L
label_13010: INC (HL)
label_13011: INC (IX + 127)
label_13012: INC (IY - 128)
label_13013: DEC A
label_13014: DEC B
label_13015: DEC C
label_13016: DEC D
label_13017: DEC E
label_13018: DEC H
label_13019: DEC L
label_13020: DEC (HL)
label_13021: DEC (IX + 127)
label_13022: DEC (IY - 128)
label_13023: DAA
label_13024: CPL
label_13025: NEG
label_13026: CCF
label_13027: SCF
label_13028: NOP
label_13029: HALT
label_13030: DI
label_13031: EI
label_13032: IM 0
label_13033: IM 1
label_13034: IM 2
label_13035: ADD HL,BC
label_13036: ADD HL,DE
label_13037: ADD HL,HL
label_13038: ADD HL,SP
label_13039: ADC HL,BC
label_13040: ADC HL,DE
label_13041: ADC HL,HL
label_13042: ADC HL,SP
label_13043: SBC HL,BC
label_13044: SBC HL,DE
label_13045: SBC HL,HL
label_13046: SBC HL,SP
label_13047: ADD IX,BC
label_13048: ADD IX,DE
label_13049: ADD IX,SP
label_13050: ADD IY,BC
label_13051: ADD IY,DE
label_13052: ADD IY,SP
label_13053: INC BC
label_13054: INC DE
label_13055: INC HL
label_13056: INC SP
label_13057: INC IX
label_13058: INC IY
label_13059: DEC BC
label_13060: DEC DE
label_13061: DEC HL
label_13062: DEC SP
label_13063: DEC IX
label_13064: DEC IY
label_13065: RLCA
label_13066: RLA
label_13067: RRCA
label_13068: RRA
label_13069: RLC A
label_13070: RLC B
label_13071: RLC C
label_13072: RLC D
label_13073: RLC E
label_13074: RLC H
label_13075: RLC L
label_13076: RLC (HL)
label_13077: RLC (IX + 127)
label_13078: RLC (IY - 128)
label_13079: RL A
label_13080: RL B
label_13081: RL C
label_13082: RL D
label_13083: RL E
label_13084: RL H
label_13085: RL L
label_13086: RL (HL)
label_13087: RL (IX + 127)
label_13088: RL (IY - 128)
label_13089: RRC A
label_13090: RRC B
label_13091: RRC C
label_13092: RRC D
label_13093: RRC E
label_13094: RRC H
label_13095: RRC L
label_13096: RRC (HL)
label_13097: RRC (IX + 127)
label_13098: RRC (IY - 128)
label_13099: RR A
label_13100: RR B
label_13101: RR C
label_13102: RR D
label_13103: RR E
label_13104: RR H
label_13105: RR L
label_13106: RR (HL)
label_13107: RR (IX + 127)
label_13108: RR (IY - 128)
label_13109: SLA A
label_13110: SLA B
label_13111: SLA C
label_13112: SLA D
label_13113: SLA E
label_13114: SLA H
label_13115: SLA L
label_13116: SLA (HL)
label_13117: SLA (IX + 127)
label_13118: SLA (IY - 128)
label_13119: SRA A
label_13120: SRA B
label_13121: SRA C
label_13122: SRA D
label_13123: SRA E
label_13124: SRA H
label_13125: SRA L
label_13126: SRA (HL)
label_13127: SRA (IX + 127)
label_13128: SRA (IY - 128)
label_13129: SRL A
label_13130: SRL B
label_13131: SRL C
label_13132: SRL D
label_13133: SRL E
label_13134: SRL H
label_13135: SRL L
label_13136: SRL (HL)
label_13137: SRL (IX + 127)
label_13138: SRL (IY - 128)
label_13139: RLD
label_13140: RRD
label_13141: BIT 0,A
label_13142: BIT 1,A
label_13143: BIT 2,A
label_13144: BIT 3,A
label_13145: BIT 4,A
label_13146: BIT 5,A
label_13147: BIT 6,A
label_13148: BIT 7,A
label_13149: BIT 0,B
label_13150: BIT 1,B
label_13151: BIT 2,B
label_13152: BIT 3,B
label_13153: BIT 4,B
label_13154: BIT 5,B
label_13155: BIT 6,B
label_13156: BIT 7,B
label_13157: BIT 0,C
label_13158: BIT 1,C
label_13159: BIT 2,C
label_13160: BIT 3,C
label_13161: BIT 4,C
label_13162: BIT 5,C
label_13163: BIT 6,C
label_13164: BIT 7,C
label_13165: BIT 0,D
label_13166: BIT 1,D
label_13167: BIT 2,D
label_13168: BIT 3,D
label_13169: BIT 4,D
label_13170: BIT 5,D
label_13171: BIT 6,D
label_13172: BIT 7,D
label_13173: BIT 0,E
label_13174: BIT 1,E
label_13175: BIT 2,E
label_13176: BIT 3,E
label_13177: BIT 4,E
label_13178: BIT 5,E
label_13179: BIT 6,E
label_13180: BIT 7,E
label_13181: BIT 0,H
label_13182: BIT 1,H
label_13183: BIT 2,H
label_13184: BIT 3,H
label_13185: BIT 4,H
label_13186: BIT 5,H
label_13187: BIT 6,H
label_13188: BIT 7,H
label_13189: BIT 0,L
label_13190: BIT 1,L
label_13191: BIT 2,L
label_13192: BIT 3,L
label_13193: BIT 4,L
label_13194: BIT 5,L
label_13195: BIT 6,L
label_13196: BIT 7,L
label_13197: BIT 0,(HL)
label_13198: BIT 1,(HL)
label_13199: BIT 2,(HL)
label_13200: BIT 3,(HL)
label_13201: BIT 4,(HL)
label_13202: BIT 5,(HL)
label_13203: BIT 6,(HL)
label_13204: BIT 7,(HL)
label_13205: BIT 0,(IX + 127)
label_13206: BIT 1,(IX + 127)
label_13207: BIT 2,(IX + 127)
label_13208: BIT 3,(IX + 127)
label_13209: BIT 4,(IX + 127)
label_13210: BIT 5,(IX + 127)
label_13211: BIT 6,(IX + 127)
label_13212: BIT 7,(IX + 127)
label_13213: BIT 0,(IY - 128)
label_13214: BIT 1,(IY - 128)
label_13215: BIT 2,(IY - 128)
label_13216: BIT 3,(IY - 128)
label_13217: BIT 4,(IY - 128)
label_13218: BIT 5,(IY - 128)
label_13219: BIT 6,(IY - 128)
label_13220: BIT 7,(IY - 128)
label_13221: SET 0,A
label_13222: SET 1,A
label_13223: SET 2,A
label_13224: SET 3,A
label_13225: SET 4,A
label_13226: SET 5,A
label_13227: SET 6,A
label_13228: SET 7,A
label_13229: SET 0,B
label_13230: SET 1,B
label_13231: SET 2,B
label_13232: SET 3,B
label_13233: SET 4,B
label_13234: SET 5,B
label_13235: SET 6,B
label_13236: SET 7,B
label_13237: SET 0,C
label_13238: SET 1,C
label_13239: SET 2,C
label_13240: SET 3,C
label_13241: SET 4,C
label_13242: SET 5,C
label_13243: SET 6,C
label_13244: SET 7,C
label_13245: SET 0,D
label_13246: SET 1,D
label_13247: SET 2,D
label_13248: SET 3,D
label_13249: SET 4,D
label_13250: SET 5,D
label_13251: SET 6,D
label_13252: SET 7,D
label_13253: SET 0,E
label_13254: SET 1,E
label_13255: SET 2,E
label_13256: SET 3,E
label_13257: SET 4,E
label_13258: SET 5,E
label_13259: SET 6,E
label_13260: SET 7,E
label_13261: SET 0,H
label_13262: SET 1,H
label_13263: SET 2,H
label_13264: SET 3,H
label_13265: SET 4,H
label_13266: SET 5,H
label_13267: SET 6,H
label_13268: SET 7,H
label_13269: SET 0,L
label_13270: SET 1,L
label_13271: SET 2,L
label_13272: SET 3,L
label_13273: SET 4,L
label_13274: SET 5,L
label_13275: SET 6,L
label_13276: SET 7,L
label_13277: SET 0,(HL)
label_13278: SET 1,(HL)
label_13279: SET 2,(HL)
label_13280: SET 3,(HL)
label_13281: SET 4,(HL)
label_13282: SET 5,(HL)
label_13283: SET 6,(HL)
label_13284: SET 7,(HL)
label_13285: SET 0,(IX + 127)
label_13286: SET 1,(IX + 127)
label_13287: SET 2,(IX + 127)
label_13288: SET 3,(IX + 127)
label_13289: SET 4,(IX + 127)
label_13290: SET 5,(IX + 127)
label_13291: SET 6,(IX + 127)
label_13292: SET 7,(IX + 127)
label_13293: SET 0,(IY - 128)
label_13294: SET 1,(IY - 128)
label_13295: SET 2,(IY - 128)
label_13296: SET 3,(IY - 128)
label_13297: SET 4,(IY - 128)
label_13298: SET 5,(IY - 128)
label_13299: SET 6,(IY - 128)
label_13300: SET 7,(IY - 128)
label_13301: RES 0,A
label_13302: RES 1,A
label_13303: RES 2,A
label_13304: RES 3,A
label_13305: RES 4,A
label_13306: RES 5,A
label_13307: RES 6,A
label_13308: RES 7,A
label_13309: RES 0,B
label_13310: RES 1,B
label_13311: RES 2,B
label_13312: RES 3,B
label_13313: RES 4,B
label_13314: RES 5,B
label_13315: RES 6,B
label_13316: RES 7,B
label_13317: RES 0,C
label_13318: RES 1,C
label_13319: RES 2,C
label_13320: RES 3,C
label_13321: RES 4,C
label_13322: RES 5,C
label_13323: RES 6,C
label_13324: RES 7,C
label_13325: RES 0,D
label_13326: RES 1,D
label_13327: RES 2,D
label_13328: RES 3,D
label_13329: RES 4,D
label_13330: RES 5,D
label_13331: RES 6,D
label_13332: RES 7,D
label_13333: RES 0,E
label_13334: RES 1,E
label_13335: RES 2,E
label_13336: RES 3,E
label_13337: RES 4,E
label_13338: RES 5,E
label_13339: RES 6,E
label_13340: RES 7,E
label_13341: RES 0,H
label_13342: RES 1,H
label_13343: RES 2,H
label_13344: RES 3,H
label_13345: RES 4,H
label_13346: RES 5,H
label_13347: RES 6,H
label_13348: RES 7,H
label_13349: RES 0,L
label_13350: RES 1,L
label_13351: RES 2,L
label_13352: RES 3,L
label_13353: RES 4,L
label_13354: RES 5,L
label_13355: RES 6,L
label_13356: RES 7,L
label_13357: RES 0,(HL)
label_13358: RES 1,(HL)
label_13359: RES 2,(HL)
label_13360: RES 3,(HL)
label_13361: RES 4,(HL)
label_13362: RES 5,(HL)
label_13363: RES 6,(HL)
label_13364: RES 7,(HL)
label_13365: RES 0,(IX + 127)
label_13366: RES 1,(IX + 127)
label_13367: RES 2,(IX + 127)
label_13368: RES 3,(IX + 127)
label_13369: RES 4,(IX + 127)
label_13370: RES 5,(IX + 127)
label_13371: RES 6,(IX + 127)
label_13372: RES 7,(IX + 127)
label_13373: RES 0,(IY - 128)
label_13374: RES 1,(IY - 128)
label_13375: RES 2,(IY - 128)
label_13376: RES 3,(IY - 128)
label_13377: RES 4,(IY - 128)
label_13378: RES 5,(IY - 128)
label_13379: RES 6,(IY - 128)
label_13380: RES 7,(IY - 128)
label_13381: JP $5678
label_13382: JP NZ,$5678
label_13383: JP Z,$5678
label_13384: JP NC,$5678
label_13385: JP C,$5678
label_13386: JP PO,$5678
label_13387: JP PE,$5678
label_13388: JP P,$5678
label_13389: JP M,$5678
label_13390: JR $ + 2
label_13391: JR NZ,$ + 2
label_13392: JR Z,$ + 2
label_13393: JR NC,$ + 2
label_13394: JR C,$ + 2
label_13395: JP (HL)
label_13396: JP (IX)
label_13397: JP (IY)
label_13398: DJNZ $ + 2
label_13399: CALL $5678
label_13400: CALL NZ,$5678
label_13401: CALL Z,$5678
label_13402: CALL NC,$5678
label_13403: CALL C,$5678
label_13404: CALL PO,$5678
label_13405: CALL PE,$5678
label_13406: CALL P,$5678
label_13407: CALL M,$5678
label_13408: RET
label_13409: RET NZ
label_13410: RET Z
label_13411: RET NC
label_13412: RET C
label_13413: RET PO
label_13414: RET PE
label_13415: RET P
label_13416: RET M
label_13417: RETI
label_13418: RETN
label_13419: RST $00
label_13420: RST $08
label_13421: RST $10
label_13422: RST $18
label_13423: RST $20
label_13424: RST $28
label_13425: RST $30
label_13426: RST $38
label_13427: IN A,($12)
label_13428: IN A,(C)
label_13429: IN B,(C)
label_13430: IN C,(C)
label_13431: IN D,(C)
label_13432: IN E,(C)
label_13433: IN H,(C)
label_13434: IN L,(C)
label_13435: IN F,(C)
label_13436: INI
label_13437: INIR
label_13438: IND
label_13439: INDR
label_13440: OUT ($12),A
label_13441: OUT (C),A
label_13442: OUT (C),B
label_13443: OUT (C),C
label_13444: OUT (C),D
label_13445: OUT (C),E
label_13446: OUT (C),H
label_13447: OUT (C),L
label_13448: OUTI
label_13449: OTIR
label_13450: OUTD
label_13451: OTDR
label_13452: LD A,A
label_13453: LD A,B
label_13454: LD A,C
label_13455: LD A,D
label_13456: LD A,E
label_13457: LD A,H
label_13458: LD A,L
label_13459: LD B,A
label_13460: LD B,B
label_13461: LD B,C
label_13462: LD B,D
label_13463: LD B,E
label_13464: LD B,H
label_13465: LD B,L
label_13466: LD C,A
label_13467: LD C,B
label_13468: LD C,C
label_13469: LD C,D
label_13470: LD C,E
label_13471: LD C,H
label_13472: LD C,L
label_13473: LD D,A
label_13474: LD D,B
label_13475: LD D,C
label_13476: LD D,D
label_13477: LD D,E
label_13478: LD D,H
label_13479: LD D,L
label_13480: LD E,A
label_13481: LD E,B
label_13482: LD E,C
label_13483: LD E,D
label_13484: LD E,E
label_13485: LD E,H
label_13486: LD E,L
label_13487: LD H,A
label_13488: LD H,B
label_13489: LD H,C
label_13490: LD H,D
label_13491: LD H,E
label_13492: LD H,H
label_13493: LD H,L
label_13494: LD L,A
label_13495: LD L,B
label_13496: LD L,C
label_13497: LD L,D
label_13498: LD L,E
label_13499: LD L,H
label_13500: LD L,L
label_13501: LD A,$12
label_13502: LD B,$12
label_13503: LD C,$12
label_13504: LD D,$12
label_13505: LD E,$12
label_13506: LD H,$12
label_13507: LD L,$12
label_13508: LD A,(HL)
label_13509: LD B,(HL)
label_13510: LD C,(HL)
label_13511: LD D,(HL)
label_13512: LD E,(HL)
label_13513: LD H,(HL)
label_13514: LD L,(HL)
label_13515: LD A,(IX + 127)
label_13516: LD B,(IX + 127)
label_13517: LD C,(IX + 127)
label_13518: LD D,(IX + 127)
label_13519: LD E,(IX + 127)
label_13520: LD H,(IX + 127)
label_13521: LD L,(IX + 127)
label_13522: LD A,(IY - 128)
label_13523: LD B,(IY - 128)
label_13524: LD C,(IY - 128)
label_13525: LD D,(IY - 128)
label_13526: LD E,(IY - 128)
label_13527: LD H,(IY - 128)
label_13528: LD L,(IY - 128)
label_13529: LD (HL),A
label_13530: LD (HL),B
label_13531: LD (HL),C
label_13532: LD (HL),D
label_13533: LD (HL),E
label_13534: LD (HL),H
label_13535: LD (HL),L
label_13536: LD (IX + 127),A
label_13537: LD (IX + 127),B
label_13538: LD (IX + 127),C
label_13539: LD (IX + 127),D
label_13540: LD (IX + 127),E
label_13541: LD (IX + 127),H
label_13542: LD (IX + 127),L
label_13543: LD (IY - 128),A
label_13544: LD (IY - 128),B
label_13545: LD (IY - 128),C
label_13546: LD (IY - 128),D
label_13547: LD (IY - 128),E
label_13548: LD (IY - 128),H
label_13549: LD (IY - 128),L
label_13550: LD (HL),$12
label_13551: LD (IX + 127),$12
label_13552: LD (IY - 128),$12
label_13553: LD A,(BC)
label_13554: LD A,(DE)
label_13555: LD A,($5678)
label_13556: LD (BC),A
label_13557: LD (DE),A
label_13558: LD ($5678),A
label_13559: LD A,I
label_13560: LD A,R
label_13561: LD I,A
label_13562: LD R,A
label_13563: LD BC,$5678
label_13564: LD DE,$5678
label_13565: LD HL,$5678
label_13566: LD SP,$5678
label_13567: LD IX,$5678
label_13568: LD IY,$5678
label_13569: LD HL,($5678)
label_13570: LD BC,($5678)
label_13571: LD DE,($5678)
label_13572: LD HL,($5678)
label_13573: LD SP,($5678)
label_13574: LD IX,($5678)
label_13575: LD IY,($5678)
label_13576: LD ($5678),HL
label_13577: LD ($5678),BC
label_13578: LD ($5678),DE
label_13579: LD ($5678),HL
label_13580: LD ($5678),SP
label_13581: LD ($5678),IX
label_13582: LD ($5678),IY
label_13583: LD SP,HL
label_13584: LD SP,IX
label_13585: LD SP,IY
label_13586: PUSH BC
label_13587: PUSH DE
label_13588: PUSH HL
label_13589: PUSH AF
label_13590: PUSH IX
label_13591: PUSH IY
label_13592: POP BC
label_13593: POP DE
label_13594: POP HL
label_13595: POP AF
label_13596: POP IX
label_13597: POP IY
label_13598: EX DE,HL
label_13599: EX AF,AF'
label_13600: EXX
label_13601: EX (SP),HL
label_13602: EX (SP),IX
label_13603: EX (SP),IY
label_13604: LDI
label_13605: LDIR
label_13606: LDD
label_13607: LDDR
label_13608: CPI
label_13609: CPIR
label_13610: CPD
label_13611: CPDR
label_13612: ADD A,A
label_13613: ADD A,B
label_13614: ADD A,C
label_13615: ADD A,D
label_13616: ADD A,E
label_13617: ADD A,H
label_13618: ADD A,L
label_13619: ADD A,$12
label_13620: ADD A,(HL)
label_13621: ADD A,(IX + 127)
label_13622: ADD A,(IY - 128)
label_13623: ADC A,A
label_13624: ADC A,B
label_13625: ADC A,C
label_13626: ADC A,D
label_13627: ADC A,E
label_13628: ADC A,H
label_13629: ADC A,L
label_13630: ADC A,$12
label_13631: ADC A,(HL)
label_13632: ADC A,(IX + 127)
label_13633: ADC A,(IY - 128)
label_13634: SUB A
label_13635: SUB B
label_13636: SUB C
label_13637: SUB D
label_13638: SUB E
label_13639: SUB H
label_13640: SUB L
label_13641: SUB $12
label_13642: SUB (HL)
label_13643: SUB (IX + 127)
label_13644: SUB (IY - 128)
label_13645: SBC A,A
label_13646: SBC A,B
label_13647: SBC A,C
label_13648: SBC A,D
label_13649: SBC A,E
label_13650: SBC A,H
label_13651: SBC A,L
label_13652: SBC A,$12
label_13653: SBC A,(HL)
label_13654: SBC A,(IX + 127)
label_13655: SBC A,(IY - 128)
label_13656: AND A
label_13657: AND B
label_13658: AND C
label_13659: AND D
label_13660: AND E
label_13661: AND H
label_13662: AND L
label_13663: AND $12
label_13664: AND (HL)
label_13665: AND (IX + 127)
label_13666: AND (IY - 128)
label_13667: AND A
label_13668: AND B
label_13669: AND C
label_13670: AND D
label_13671: AND E
label_13672: AND H
label_13673: AND L
label_13674: AND $12
label_13675: AND (HL)
label_13676: AND (IX + 127)
label_13677: AND (IY - 128)
label_13678: OR A
label_13679: OR B
label_13680: OR C
label_13681: OR D
label_13682: OR E
label_13683: OR H
label_13684: OR L
label_13685: OR $12
label_13686: OR (HL)
label_13687: OR (IX + 127)
label_13688: OR (IY - 128)
label_13689: XOR A
label_13690: XOR B
label_13691: XOR C
label_13692: XOR D
label_13693: XOR E
label_13694: XOR H
label_13695: XOR L
label_13696: XOR $12
label_13697: XOR (HL)
label_13698: XOR (IX + 127)
label_13699: XOR (IY - 128)
label_13700: CP A
label_13701: CP B
label_13702: CP C
label_13703: CP D
label_13704: CP E
label_13705: CP H
label_13706: CP L
label_13707: CP $12
label_13708: CP (HL)
label_13709: CP (IX + 127)
label_13710: CP (IY - 128)
label_13711: INC A
label_13712: INC B
label_13713: INC C
label_13714: INC D
label_13715: INC E
label_13716: INC H
label_13717: INC L
label_13718: INC (HL)
label_13719: INC (IX + 127)
label_13720: INC (IY - 128)
label_13721: DEC A
label_13722: DEC B
label_13723: DEC C
label_13724: DEC D
label_13725: DEC E
label_13726: DEC H
label_13727: DEC L
label_13728: DEC (HL)
label_13729: DEC (IX + 127)
label_13730: DEC (IY - 128)
label_13731: DAA
label_13732: CPL
label_13733: NEG
label_13734: CCF
label_13735: SCF
label_13736: NOP
label_13737: HALT
label_13738: DI
label_13739: EI
label_13740: IM 0
label_13741: IM 1
label_13742: IM 2
label_13743: ADD HL,BC
label_13744: ADD HL,DE
label_13745: ADD HL,HL
label_13746: ADD HL,SP
label_13747: ADC HL,BC
label_13748: ADC HL,DE
label_13749: ADC HL,HL
label_13750: ADC HL,SP
label_13751: SBC HL,BC
label_13752: SBC HL,DE
label_13753: SBC HL,HL
label_13754: SBC HL,SP
label_13755: ADD IX,BC
label_13756: ADD IX,DE
label_13757: ADD IX,SP
label_13758: ADD IY,BC
label_13759: ADD IY,DE
label_13760: ADD IY,SP
label_13761: INC BC
label_13762: INC DE
label_13763: INC HL
label_13764: INC SP
label_13765: INC IX
label_13766: INC IY
label_13767: DEC BC
label_13768: DEC DE
label_13769: DEC HL
label_13770: DEC SP
label_13771: DEC IX
label_13772: DEC IY
label_13773: RLCA
label_13774: RLA
label_13775: RRCA
label_13776: RRA
label_13777: RLC A
label_13778: RLC B
label_13779: RLC C
label_13780: RLC D
label_13781: RLC E
label_13782: RLC H
label_13783: RLC L
label_13784: RLC (HL)
label_13785: RLC (IX + 127)
label_13786: RLC (IY - 128)
label_13787: RL A
label_13788: RL B
label_13789: RL C
label_13790: RL D
label_13791: RL E
label_13792: RL H
label_13793: RL L
label_13794: RL (HL)
label_13795: RL (IX + 127)
label_13796: RL (IY - 128)
label_13797: RRC A
label_13798: RRC B
label_13799: RRC C
label_13800: RRC D
label_13801: RRC E
label_13802: RRC H
label_13803: RRC L
label_13804: RRC (HL)
label_13805: RRC (IX + 127)
label_13806: RRC (IY - 128)
label_13807: RR A
label_13808: RR B
label_13809: RR C
label_13810: RR D
label_13811: RR E
label_13812: RR H
label_13813: RR L
label_13814: RR (HL)
label_13815: RR (IX + 127)
label_13816: RR (IY - 128)
label_13817: SLA A
label_13818: SLA B
label_13819: SLA C
label_13820: SLA D
label_13821: SLA E
label_13822: SLA H
label_13823: SLA L
label_13824: SLA (HL)
label_13825: SLA (IX + 127)
label_13826: SLA (IY - 128)
label_13827: SRA A
label_13828: SRA B
label_13829: SRA C
label_13830: SRA D
label_13831: SRA E
label_13832: SRA H
label_13833: SRA L
label_13834: SRA (HL)
label_13835: SRA (IX + 127)
label_13836: SRA (IY - 128)
label_13837: SRL A
label_13838: SRL B
label_13839: SRL C
label_13840: SRL D
label_13841: SRL E
label_13842: SRL H
label_13843: SRL L
label_13844: SRL (HL)
label_13845: SRL (IX + 127)
label_13846: SRL (IY - 128)
label_13847: RLD
label_13848: RRD
label_13849: BIT 0,A
label_13850: BIT 1,A
label_13851: BIT 2,A
label_13852: BIT 3,A
label_13853: BIT 4,A
label_13854: BIT 5,A
label_13855: BIT 6,A
label_13856: BIT 7,A
label_13857: BIT 0,B
label_13858: BIT 1,B
label_13859: BIT 2,B
label_13860: BIT 3,B
label_13861: BIT 4,B
label_13862: BIT 5,B
label_13863: BIT 6,B
label_13864: BIT 7,B
label_13865: BIT 0,C
label_13866: BIT 1,C
label_13867: BIT 2,C
label_13868: BIT 3,C
label_13869: BIT 4,C
label_13870: BIT 5,C
label_13871: BIT 6,C
label_13872: BIT 7,C
label_13873: BIT 0,D
label_13874: BIT 1,D
label_13875: BIT 2,D
label_13876: BIT 3,D
label_13877: BIT 4,D
label_13878: BIT 5,D
label_13879: BIT 6,D
label_13880: BIT 7,D
label_13881: BIT 0,E
label_13882: BIT 1,E
label_13883: BIT 2,E
label_13884: BIT 3,E
label_13885: BIT 4,E
label_13886: BIT 5,E
label_13887: BIT 6,E
label_13888: BIT 7,E
label_13889: BIT 0,H
label_13890: BIT 1,H
label_13891: BIT 2,H
label_13892: BIT 3,H
label_13893: BIT 4,H
label_13894: BIT 5,H
label_13895: BIT 6,H
label_13896: BIT 7,H
label_13897: BIT 0,L
label_13898: BIT 1,L
label_13899: BIT 2,L
label_13900: BIT 3,L
label_13901: BIT 4,L
label_13902: BIT 5,L
label_13903: BIT 6,L
label_13904: BIT 7,L
label_13905: BIT 0,(HL)
label_13906: BIT 1,(HL)
label_13907: BIT 2,(HL)
label_13908: BIT 3,(HL)
label_13909: BIT 4,(HL)
label_13910: BIT 5,(HL)
label_13911: BIT 6,(HL)
label_13912: BIT 7,(HL)
label_13913: BIT 0,(IX + 127)
label_13914: BIT 1,(IX + 127)
label_13915: BIT 2,(IX + 127)
label_13916: BIT 3,(IX + 127)
label_13917: BIT 4,(IX + 127)
label_13918: BIT 5,(IX + 127)
label_13919: BIT 6,(IX + 127)
label_13920: BIT 7,(IX + 127)
label_13921: BIT 0,(IY - 128)
label_13922: BIT 1,(IY - 128)
label_13923: BIT 2,(IY - 128)
label_13924: BIT 3,(IY - 128)
label_13925: BIT 4,(IY - 128)
label_13926: BIT 5,(IY - 128)
label_13927: BIT 6,(IY - 128)
label_13928: BIT 7,(IY - 128)
label_13929: SET 0,A
label_13930: SET 1,A
label_13931: SET 2,A
label_13932: SET 3,A
label_13933: SET 4,A
label_13934: SET 5,A
label_13935: SET 6,A
label_13936: SET 7,A
label_13937: SET 0,B
label_13938: SET 1,B
label_13939: SET 2,B
label_13940: SET 3,B
label_13941: SET 4,B
label_13942: SET 5,B
label_13943: SET 6,B
label_13944: SET 7,B
label_13945: SET 0,C
label_13946: SET 1,C
label_13947: SET 2,C
label_13948: SET 3,C
label_13949: SET 4,C
label_13950: SET 5,C
label_13951: SET 6,C
label_13952: SET 7,C
label_13953: SET 0,D
label_13954: SET 1,D
label_13955: SET 2,D
label_13956: SET 3,D
label_13957: SET 4,D
label_13958: SET 5,D
label_13959: SET 6,D
label_13960: SET 7,D
label_13961: SET 0,E
label_13962: SET 1,E
label_13963: SET 2,E
label_13964: SET 3,E
label_13965: SET 4,E
label_13966: SET 5,E
label_13967: SET 6,E
label_13968: SET 7,E
label_13969: SET 0,H
label_13970: SET 1,H
label_13971: SET 2,H
label_13972: SET 3,H
label_13973: SET 4,H
label_13974: SET 5,H
label_13975: SET 6,H
label_13976: SET 7,H
label_13977: SET 0,L
label_13978: SET 1,L
label_13979: SET 2,L
label_13980: SET 3,L
label_13981: SET 4,L
label_13982: SET 5,L
label_13983: SET 6,L
label_13984: SET 7,L
label_13985: SET 0,(HL)
label_13986: SET 1,(HL)
label_13987: SET 2,(HL)
label_13988: SET 3,(HL)
label_13989: SET 4,(HL)
label_13990: SET 5,(HL)
label_13991: SET 6,(HL)
label_13992: SET 7,(HL)
label_13993: SET 0,(IX + 127)
label_13994: SET 1,(IX + 127)
label_13995: SET 2,(IX + 127)
label_13996: SET 3,(IX + 127)
label_13997: SET 4,(IX + 127)
label_13998: SET 5,(IX + 127)
label_13999: SET 6,(IX + 127)
label_14000: SET 7,(IX + 127)
label_14001: SET 0,(IY - 128)
label_14002: SET 1,(IY - 128)
label_14003: SET 2,(IY - 128)
label_14004: SET 3,(IY - 128)
label_14005: SET 4,(IY - 128)
label_14006: SET 5,(IY - 128)
label_14007: SET 6,(IY - 128)
label_14008: SET 7,(IY - 128)
label_14009: RES 0,A
label_14010: RES 1,A
label_14011: RES 2,A
label_14012: RES 3,A
label_14013: RES 4,A
label_14014: RES 5,A
label_14015: RES 6,A
label_14016: RES 7,A
label_14017: RES 0,B
label_14018: RES 1,B
label_14019: RES 2,B
label_14020: RES 3,B
label_14021: RES 4,B
label_14022: RES 5,B
label_14023: RES 6,B
label_14024: RES 7,B
label_14025: RES 0,C
label_14026: RES 1,C
label_14027: RES 2,C
label_14028: RES 3,C
label_14029: RES 4,C
label_14030: RES 5,C
label_14031: RES 6,C
label_14032: RES 7,C
label_14033: RES 0,D
label_14034: RES 1,D
label_14035: RES 2,D
label_14036: RES 3,D
label_14037: RES 4,D
label_14038: RES 5,D
label_14039: RES 6,D
label_14040: RES 7,D
label_14041: RES 0,E
label_14042: RES 1,E
label_14043: RES 2,E
label_14044: RES 3,E
label_14045: RES 4,E
label_14046: RES 5,E
label_14047: RES 6,E
label_14048: RES 7,E
label_14049: RES 0,H
label_14050: RES 1,H
label_14051: RES 2,H
label_14052: RES 3,H
label_14053: RES 4,H
label_14054: RES 5,H
label_14055: RES 6,H
label_14056: RES 7,H
label_14057: RES 0,L
label_14058: RES 1,L
label_14059: RES 2,L
label_14060: RES 3,L
label_14061: RES 4,L
label_14062: RES 5,L
label_14063: RES 6,L
label_14064: RES 7,L
label_14065: RES 0,(HL)
label_14066: RES 1,(HL)
label_14067: RES 2,(HL)
label_14068: RES 3,(HL)
label_14069: RES 4,(HL)
label_14070: RES 5,(HL)
label_14071: RES 6,(HL)
label_14072: RES 7,(HL)
label_14073: RES 0,(IX + 127)
label_14074: RES 1,(IX + 127)
label_14075: RES 2,(IX + 127)
label_14076: RES 3,(IX + 127)
label_14077: RES 4,(IX + 127)
label_14078: RES 5,(IX + 127)
label_14079: RES 6,(IX + 127)
label_14080: RES 7,(IX + 127)
label_14081: RES 0,(IY - 128)
label_14082: RES 1,(IY - 128)
label_14083: RES 2,(IY - 128)
label_14084: RES 3,(IY - 128)
label_14085: RES 4,(IY - 128)
label_14086: RES 5,(IY - 128)
label_14087: RES 6,(IY - 128)
label_14088: RES 7,(IY - 128)
label_14089: JP $5678
label_14090: JP NZ,$5678
label_14091: JP Z,$5678
label_14092: JP NC,$5678
label_14093: JP C,$5678
label_14094: JP PO,$5678
label_14095: JP PE,$5678
label_14096: JP P,$5678
label_14097: JP M,$5678
label_14098: JR $ + 2
label_14099: JR NZ,$ + 2
label_14100: JR Z,$ + 2
label_14101: JR NC,$ + 2
label_14102: JR C,$ + 2
label_14103: JP (HL)
label_14104: JP (IX)
label_14105: JP (IY)
label_14106: DJNZ $ + 2
label_14107: CALL $5678
label_14108: CALL NZ,$5678
label_14109: CALL Z,$5678
label_14110: CALL NC,$5678
label_14111: CALL C,$5678
label_14112: CALL PO,$5678
label_14113: CALL PE,$5678
label_14114: CALL P,$5678
label_14115: CALL M,$5678
label_14116: RET
label_14117: RET NZ
label_14118: RET Z
label_14119: RET NC
label_14120: RET C
label_14121: RET PO
label_14122: RET PE
label_14123: RET P
label_14124: RET M
label_14125: RETI
label_14126: RETN
label_14127: RST $00
label_14128: RST $08
label_14129: RST $10
label_14130: RST $18
label_14131: RST $20
label_14132: RST $28
label_14133: RST $30
label_14134: RST $38
label_14135: IN A,($12)
label_14136: IN A,(C)
label_14137: IN B,(C)
label_14138: IN C,(C)
label_14139: IN D,(C)
label_14140: IN E,(C)
label_14141: IN H,(C)
label_14142: IN L,(C)
label_14143: IN F,(C)
label_14144: INI
label_14145: INIR
label_14146: IND
label_14147: INDR
label_14148: OUT ($12),A
label_14149: OUT (C),A
label_14150: OUT (C),B
label_14151: OUT (C),C
label_14152: OUT (C),D
label_14153: OUT (C),E
label_14154: OUT (C),H
label_14155: OUT (C),L
label_14156: OUTI
label_14157: OTIR
label_14158: OUTD
label_14159: OTDR
label_14160: LD A,A
label_14161: LD A,B
label_14162: LD A,C
label_14163: LD A,D
label_14164: LD A,E
label_14165: LD A,H
label_14166: LD A,L
label_14167: LD B,A
label_14168: LD B,B
label_14169: LD B,C
label_14170: LD B,D
label_14171: LD B,E
label_14172: LD B,H
label_14173: LD B,L
label_14174: LD C,A
label_14175: LD C,B
label_14176: LD C,C
label_14177: LD C,D
label_14178: LD C,E
label_14179: LD C,H
label_14180: LD C,L
label_14181: LD D,A
label_14182: LD D,B
label_14183: LD D,C
label_14184: LD D,D
label_14185: LD D,E
label_14186: LD D,H
label_14187: LD D,L
label_14188: LD E,A
label_14189: LD E,B
label_14190: LD E,C
label_14191: LD E,D
label_14192: LD E,E
label_14193: LD E,H
label_14194: LD E,L
label_14195: LD H,A
label_14196: LD H,B
label_14197: LD H,C
label_14198: LD H,D
label_14199: LD H,E
label_14200: LD H,H
label_14201: LD H,L
label_14202: LD L,A
label_14203: LD L,B
label_14204: LD L,C
label_14205: LD L,D
label_14206: LD L,E
label_14207: LD L,H
label_14208: LD L,L
label_14209: LD A,$12
label_14210: LD B,$12
label_14211: LD C,$12
label_14212: LD D,$12
label_14213: LD E,$12
label_14214: LD H,$12
label_14215: LD L,$12
label_14216: LD A,(HL)
label_14217: LD B,(HL)
label_14218: LD C,(HL)
label_14219: LD D,(HL)
label_14220: LD E,(HL)
label_14221: LD H,(HL)
label_14222: LD L,(HL)
label_14223: LD A,(IX + 127)
label_14224: LD B,(IX + 127)
label_14225: LD C,(IX + 127)
label_14226: LD D,(IX + 127)
label_14227: LD E,(IX + 127)
label_14228: LD H,(IX + 127)
label_14229: LD L,(IX + 127)
label_14230: LD A,(IY - 128)
label_14231: LD B,(IY - 128)
label_14232: LD C,(IY - 128)
label_14233: LD D,(IY - 128)
label_14234: LD E,(IY - 128)
label_14235: LD H,(IY - 128)
label_14236: LD L,(IY - 128)
label_14237: LD (HL),A
label_14238: LD (HL),B
label_14239: LD (HL),C
label_14240: LD (HL),D
label_14241: LD (HL),E
label_14242: LD (HL),H
label_14243: LD (HL),L
label_14244: LD (IX + 127),A
label_14245: LD (IX + 127),B
label_14246: LD (IX + 127),C
label_14247: LD (IX + 127),D
label_14248: LD (IX + 127),E
label_14249: LD (IX + 127),H
label_14250: LD (IX + 127),L
label_14251: LD (IY - 128),A
label_14252: LD (IY - 128),B
label_14253: LD (IY - 128),C
label_14254: LD (IY - 128),D
label_14255: LD (IY - 128),E
label_14256: LD (IY - 128),H
label_14257: LD (IY - 128),L
label_14258: LD (HL),$12
label_14259: LD (IX + 127),$12
label_14260: LD (IY - 128),$12
label_14261: LD A,(BC)
label_14262: LD A,(DE)
label_14263: LD A,($5678)
label_14264: LD (BC),A
label_14265: LD (DE),A
label_14266: LD ($5678),A
label_14267: LD A,I
label_14268: LD A,R
label_14269: LD I,A
label_14270: LD R,A
label_14271: LD BC,$5678
label_14272: LD DE,$5678
label_14273: LD HL,$5678
label_14274: LD SP,$5678
label_14275: LD IX,$5678
label_14276: LD IY,$5678
label_14277: LD HL,($5678)
label_14278: LD BC,($5678)
label_14279: LD DE,($5678)
label_14280: LD HL,($5678)
label_14281: LD SP,($5678)
label_14282: LD IX,($5678)
label_14283: LD IY,($5678)
label_14284: LD ($5678),HL
label_14285: LD ($5678),BC
label_14286: LD ($5678),DE
label_14287: LD ($5678),HL
label_14288: LD ($5678),SP
label_14289: LD ($5678),IX
label_14290: LD ($5678),IY
label_14291: LD SP,HL
label_14292: LD SP,IX
label_14293: LD SP,IY
label_14294: PUSH BC
label_14295: PUSH DE
label_14296: PUSH HL
label_14297: PUSH AF
label_14298: PUSH IX
label_14299: PUSH IY
label_14300: POP BC
label_14301: POP DE
label_14302: POP HL
label_14303: POP AF
label_14304: POP IX
label_14305: POP IY
label_14306: EX DE,HL
label_14307: EX AF,AF'
label_14308: EXX
label_14309: EX (SP),HL
label_14310: EX (SP),IX
label_14311: EX (SP),IY
label_14312: LDI
label_14313: LDIR
label_14314: LDD
label_14315: LDDR
label_14316: CPI
label_14317: CPIR
label_14318: CPD
label_14319: CPDR
label_14320: ADD A,A
label_14321: ADD A,B
label_14322: ADD A,C
label_14323: ADD A,D
label_14324: ADD A,E
label_14325: ADD A,H
label_14326: ADD A,L
label_14327: ADD A,$12
label_14328: ADD A,(HL)
label_14329: ADD A,(IX + 127)
label_14330: ADD A,(IY - 128)
label_14331: ADC A,A
label_14332: ADC A,B
label_14333: ADC A,C
label_14334: ADC A,D
label_14335: ADC A,E
label_14336: ADC A,H
label_14337: ADC A,L
label_14338: ADC A,$12
label_14339: ADC A,(HL)
label_14340: ADC A,(IX + 127)
label_14341: ADC A,(IY - 128)
label_14342: SUB A
label_14343: SUB B
label_14344: SUB C
label_14345: SUB D
label_14346: SUB E
label_14347: SUB H
label_14348: SUB L
label_14349: SUB $12
label_14350: SUB (HL)
label_14351: SUB (IX + 127)
label_14352: SUB (IY - 128)
label_14353: SBC A,A
label_14354: SBC A,B
label_14355: SBC A,C
label_14356: SBC A,D
label_14357: SBC A,E
label_14358: SBC A,H
label_14359: SBC A,L
label_14360: SBC A,$12
label_14361: SBC A,(HL)
label_14362: SBC A,(IX + 127)
label_14363: SBC A,(IY - 128)
label_14364: AND A
label_14365: AND B
label_14366: AND C
label_14367: AND D
label_14368: AND E
label_14369: AND H
label_14370: AND L
label_14371: AND $12
label_14372: AND (HL)
label_14373: AND (IX + 127)
label_14374: AND (IY - 128)
label_14375: AND A
label_14376: AND B
label_14377: AND C
label_14378: AND D
label_14379: AND E
label_14380: AND H
label_14381: AND L
label_14382: AND $12
label_14383: AND (HL)
label_14384: AND (IX + 127)
label_14385: AND (IY - 128)
label_14386: OR A
label_14387: OR B
label_14388: OR C
label_14389: OR D
label_14390: OR E
label_14391: OR H
label_14392: OR L
label_14393: OR $12
label_14394: OR (HL)
label_14395: OR (IX + 127)
label_14396: OR (IY - 128)
label_14397: XOR A
label_14398: XOR B
label_14399: XOR C
label_14400: XOR D
label_14401: XOR E
label_14402: XOR H
label_14403: XOR L
label_14404: XOR $12
label_14405: XOR (HL)
label_14406: XOR (IX + 127)
label_14407: XOR (IY - 128)
label_14408: CP A
label_14409: CP B
label_14410: CP C
label_14411: CP D
label_14412: CP E
label_14413: CP H
label_14414: CP L
label_14415: CP $12
label_14416: CP (HL)
label_14417: CP (IX + 127)
label_14418: CP (IY - 128)
label_14419: INC A
label_14420: INC B
label_14421: INC C
label_14422: INC D
label_14423: INC E
label_14424: INC H
label_14425: INC L
label_14426: INC (HL)
label_14427: INC (IX + 127)
label_14428: INC (IY - 128)
label_14429: DEC A
label_14430: DEC B
label_14431: DEC C
label_14432: DEC D
label_14433: DEC E
label_14434: DEC H
label_14435: DEC L
label_14436: DEC (HL)
label_14437: DEC (IX + 127)
label_14438: DEC (IY - 128)
label_14439: DAA
label_14440: CPL
label_14441: NEG
label_14442: CCF
label_14443: SCF
label_14444: NOP
label_14445: HALT
label_14446: DI
label_14447: EI
label_14448: IM 0
label_14449: IM 1
label_14450: IM 2
label_14451: ADD HL,BC
label_14452: ADD HL,DE
label_14453: ADD HL,HL
label_14454: ADD HL,SP
label_14455: ADC HL,BC
label_14456: ADC HL,DE
label_14457: ADC HL,HL
label_14458: ADC HL,SP
label_14459: SBC HL,BC
label_14460: SBC HL,DE
label_14461: SBC HL,HL
label_14462: SBC HL,SP
label_14463: ADD IX,BC
label_14464: ADD IX,DE
label_14465: ADD IX,SP
label_14466: ADD IY,BC
label_14467: ADD IY,DE
label_14468: ADD IY,SP
label_14469: INC BC
label_14470: INC DE
label_14471: INC HL
label_14472: INC SP
label_14473: INC IX
label_14474: INC IY
label_14475: DEC BC
label_14476: DEC DE
label_14477: DEC HL
label_14478: DEC SP
label_14479: DEC IX
label_14480: DEC IY
label_14481: RLCA
label_14482: RLA
label_14483: RRCA
label_14484: RRA
label_14485: RLC A
label_14486: RLC B
label_14487: RLC C
label_14488: RLC D
label_14489: RLC E
label_14490: RLC H
label_14491: RLC L
label_14492: RLC (HL)
label_14493: RLC (IX + 127)
label_14494: RLC (IY - 128)
label_14495: RL A
label_14496: RL B
label_14497: RL C
label_14498: RL D
label_14499: RL E
label_14500: RL H
label_14501: RL L
label_14502: RL (HL)
label_14503: RL (IX + 127)
label_14504: RL (IY - 128)
label_14505: RRC A
label_14506: RRC B
label_14507: RRC C
label_14508: RRC D
label_14509: RRC E
label_14510: RRC H
label_14511: RRC L
label_14512: RRC (HL)
label_14513: RRC (IX + 127)
label_14514: RRC (IY - 128)
label_14515: RR A
label_14516: RR B
label_14517: RR C
label_14518: RR D
label_14519: RR E
label_14520: RR H
label_14521: RR L
label_14522: RR (HL)
label_14523: RR (IX + 127)
label_14524: RR (IY - 128)
label_14525: SLA A
label_14526: SLA B
label_14527: SLA C
label_14528: SLA D
label_14529: SLA E
label_14530: SLA H
label_14531: SLA L
label_14532: SLA (HL)
label_14533: SLA (IX + 127)
label_14534: SLA (IY - 128)
label_14535: SRA A
label_14536: SRA B
label_14537: SRA C
label_14538: SRA D
label_14539: SRA E
label_14540: SRA H
label_14541: SRA L
label_14542: SRA (HL)
label_14543: SRA (IX + 127)
label_14544: SRA (IY - 128)
label_14545: SRL A
label_14546: SRL B
label_14547: SRL C
label_14548: SRL D
label_14549: SRL E
label_14550: SRL H
label_14551: SRL L
label_14552: SRL (HL)
label_14553: SRL (IX + 127)
label_14554: SRL (IY - 128)
label_14555: RLD
label_14556: RRD
label_14557: BIT 0,A
label_14558: BIT 1,A
label_14559: BIT 2,A
label_14560: BIT 3,A
label_14561: BIT 4,A
label_14562: BIT 5,A
label_14563: BIT 6,A
label_14564: BIT 7,A
label_14565: BIT 0,B
label_14566: BIT 1,B
label_14567: BIT 2,B
label_14568: BIT 3,B
label_14569: BIT 4,B
label_14570: BIT 5,B
label_14571: BIT 6,B
label_14572: BIT 7,B
label_14573: BIT 0,C
label_14574: BIT 1,C
label_14575: BIT 2,C
label_14576: BIT 3,C
label_14577: BIT 4,C
label_14578: BIT 5,C
label_14579: BIT 6,C
label_14580: BIT 7,C
label_14581: BIT 0,D
label_14582: BIT 1,D
label_14583: BIT 2,D
label_14584: BIT 3,D
label_14585: BIT 4,D
label_14586: BIT 5,D
label_14587: BIT 6,D
label_14588: BIT 7,D
label_14589: BIT 0,E
label_14590: BIT 1,E
label_14591: BIT 2,E
label_14592: BIT 3,E
label_14593: BIT 4,E
label_14594: BIT 5,E
label_14595: BIT 6,E
label_14596: BIT 7,E
label_14597: BIT 0,H
label_14598: BIT 1,H
label_14599: BIT 2,H
label_14600: BIT 3,H
label_14601: BIT 4,H
label_14602: BIT 5,H
label_14603: BIT 6,H
label_14604: BIT 7,H
label_14605: BIT 0,L
label_14606: BIT 1,L
label_14607: BIT 2,L
label_14608: BIT 3,L
label_14609: BIT 4,L
label_14610: BIT 5,L
label_14611: BIT 6,L
label_14612: BIT 7,L
label_14613: BIT 0,(HL)
label_14614: BIT 1,(HL)
label_14615: BIT 2,(HL)
label_14616: BIT 3,(HL)
label_14617: BIT 4,(HL)
label_14618: BIT 5,(HL)
label_14619: BIT 6,(HL)
label_14620: BIT 7,(HL)
label_14621: BIT 0,(IX + 127)
label_14622: BIT 1,(IX + 127)
label_14623: BIT 2,(IX + 127)
label_14624: BIT 3,(IX + 127)
label_14625: BIT 4,(IX + 127)
label_14626: BIT 5,(IX + 127)
label_14627: BIT 6,(IX + 127)
label_14628: BIT 7,(IX + 127)
label_14629: BIT 0,(IY - 128)
label_14630: BIT 1,(IY - 128)
label_14631: BIT 2,(IY - 128)
label_14632: BIT 3,(IY - 128)
label_14633: BIT 4,(IY - 128)
label_14634: BIT 5,(IY - 128)
label_14635: BIT 6,(IY - 128)
label_14636: BIT 7,(IY - 128)
label_14637: SET 0,A
label_14638: SET 1,A
label_14639: SET 2,A
label_14640: SET 3,A
label_14641: SET 4,A
label_14642: SET 5,A
label_14643: SET 6,A
label_14644: SET 7,A
label_14645: SET 0,B
label_14646: SET 1,B
label_14647: SET 2,B
label_14648: SET 3,B
label_14649: SET 4,B
label_14650: SET 5,B
label_14651: SET 6,B
label_14652: SET 7,B
label_14653: SET 0,C
label_14654: SET 1,C
label_14655: SET 2,C
label_14656: SET 3,C
label_14657: SET 4,C
label_14658: SET 5,C
label_14659: SET 6,C
label_14660: SET 7,C
label_14661: SET 0,D
label_14662: SET 1,D
label_14663: SET 2,D
label_14664: SET 3,D
label_14665: SET 4,D
label_14666: SET 5,D
label_14667: SET 6,D
label_14668: SET 7,D
label_14669: SET 0,E
label_14670: SET 1,E
label_14671: SET 2,E
label_14672: SET 3,E
label_14673: SET 4,E
label_14674: SET 5,E
label_14675: SET 6,E
label_14676: SET 7,E
label_14677: SET 0,H
label_14678: SET 1,H
label_14679: SET 2,H
label_14680: SET 3,H
label_14681: SET 4,H
label_14682: SET 5,H
label_14683: SET 6,H
label_14684: SET 7,H
label_14685: SET 0,L
label_14686: SET 1,L
label_14687: SET 2,L
label_14688: SET 3,L
label_14689: SET 4,L
label_14690: SET 5,L
label_14691: SET 6,L
label_14692: SET 7,L
label_14693: SET 0,(HL)
label_14694: SET 1,(HL)
label_14695: SET 2,(HL)
label_14696: SET 3,(HL)
label_14697: SET 4,(HL)
label_14698: SET 5,(HL)
label_14699: SET 6,(HL)
label_14700: SET 7,(HL)
label_14701: SET 0,(IX + 127)
label_14702: SET 1,(IX + 127)
label_14703: SET 2,(IX + 127)
label_14704: SET 3,(IX + 127)
label_14705: SET 4,(IX + 127)
label_14706: SET 5,(IX + 127)
label_14707: SET 6,(IX + 127)
label_14708: SET 7,(IX + 127)
label_14709: SET 0,(IY - 128)
label_14710: SET 1,(IY - 128)
label_14711: SET 2,(IY - 128)
label_14712: SET 3,(IY - 128)
label_14713: SET 4,(IY - 128)
label_14714: SET 5,(IY - 128)
label_14715: SET 6,(IY - 128)
label_14716: SET 7,(IY - 128)
label_14717: RES 0,A
label_14718: RES 1,A
label_14719: RES 2,A
label_14720: RES 3,A
label_14721: RES 4,A
label_14722: RES 5,A
label_14723: RES 6,A
label_14724: RES 7,A
label_14725: RES 0,B
label_14726: RES 1,B
label_14727: RES 2,B
label_14728: RES 3,B
label_14729: RES 4,B
label_14730: RES 5,B
label_14731: RES 6,B
label_14732: RES 7,B
label_14733: RES 0,C
label_14734: RES 1,C
label_14735: RES 2,C
label_14736: RES 3,C
label_14737: RES 4,C
label_14738: RES 5,C
label_14739: RES 6,C
label_14740: RES 7,C
label_14741: RES 0,D
label_14742: RES 1,D
label_14743: RES 2,D
label_14744: RES 3,D
label_14745: RES 4,D
label_14746: RES 5,D
label_14747: RES 6,D
label_14748: RES 7,D
label_14749: RES 0,E
label_14750: RES 1,E
label_14751: RES 2,E
label_14752: RES 3,E
label_14753: RES 4,E
label_14754: RES 5,E
label_14755: RES 6,E
label_14756: RES 7,E
label_14757: RES 0,H
label_14758: RES 1,H
label_14759: RES 2,H
label_14760: RES 3,H
label_14761: RES 4,H
label_14762: RES 5,H
label_14763: RES 6,H
label_14764: RES 7,H
label_14765: RES 0,L
label_14766: RES 1,L
label_14767: RES 2,L
label_14768: RES 3,L
label_14769: RES 4,L
label_14770: RES 5,L
label_14771: RES 6,L
label_14772: RES 7,L
label_14773: RES 0,(HL)
label_14774: RES 1,(HL)
label_14775: RES 2,(HL)
label_14776: RES 3,(HL)
label_14777: RES 4,(HL)
label_14778: RES 5,(HL)
label_14779: RES 6,(HL)
label_14780: RES 7,(HL)
label_14781: RES 0,(IX + 127)
label_14782: RES 1,(IX + 127)
label_14783: RES 2,(IX + 127)
label_14784: RES 3,(IX + 127)
label_14785: RES 4,(IX + 127)
label_14786: RES 5,(IX + 127)
label_14787: RES 6,(IX + 127)
label_14788: RES 7,(IX + 127)
label_14789: RES 0,(IY - 128)
label_14790: RES 1,(IY - 128)
label_14791: RES 2,(IY - 128)
label_14792: RES 3,(IY - 128)
label_14793: RES 4,(IY - 128)
label_14794: RES 5,(IY - 128)
label_14795: RES 6,(IY - 128)
label_14796: RES 7,(IY - 128)
label_14797: JP $5678
label_14798: JP NZ,$5678
label_14799: JP Z,$5678
label_14800: JP NC,$5678
label_14801: JP C,$5678
label_14802: JP PO,$5678
label_14803: JP PE,$5678
label_14804: JP P,$5678
label_14805: JP M,$5678
label_14806: JR $ + 2
label_14807: JR NZ,$ + 2
label_14808: JR Z,$ + 2
label_14809: JR NC,$ + 2
label_14810: JR C,$ + 2
label_14811: JP (HL)
label_14812: JP (IX)
label_14813: JP (IY)
label_14814: DJNZ $ + 2
label_14815: CALL $5678
label_14816: CALL NZ,$5678
label_14817: CALL Z,$5678
label_14818: CALL NC,$5678
label_14819: CALL C,$5678
label_14820: CALL PO,$5678
label_14821: CALL PE,$5678
label_14822: CALL P,$5678
label_14823: CALL M,$5678
label_14824: RET
label_14825: RET NZ
label_14826: RET Z
label_14827: RET NC
label_14828: RET C
label_14829: RET PO
label_14830: RET PE
label_14831: RET P
label_14832: RET M
label_14833: RETI
label_14834: RETN
label_14835: RST $00
label_14836: RST $08
label_14837: RST $10
label_14838: RST $18
label_14839: RST $20
label_14840: RST $28
label_14841: RST $30
label_14842: RST $38
label_14843: IN A,($12)
label_14844: IN A,(C)
label_14845: IN B,(C)
label_14846: IN C,(C)
label_14847: IN D,(C)
label_14848: IN E,(C)
label_14849: IN H,(C)
label_14850: IN L,(C)
label_14851: IN F,(C)
label_14852: INI
label_14853: INIR
label_14854: IND
label_14855: INDR
label_14856: OUT ($12),A
label_14857: OUT (C),A
label_14858: OUT (C),B
label_14859: OUT (C),C
label_14860: OUT (C),D
label_14861: OUT (C),E
label_14862: OUT (C),H
label_14863: OUT (C),L
label_14864: OUTI
label_14865: OTIR
label_14866: OUTD
label_14867: OTDR
label_14868: LD A,A
label_14869: LD A,B
label_14870: LD A,C
label_14871: LD A,D
label_14872: LD A,E
label_14873: LD A,H
label_14874: LD A,L
label_14875: LD B,A
label_14876: LD B,B
label_14877: LD B,C
label_14878: LD B,D
label_14879: LD B,E
label_14880: LD B,H
label_14881: LD B,L
label_14882: LD C,A
label_14883: LD C,B
label_14884: LD C,C
label_14885: LD C,D
label_14886: LD C,E
label_14887: LD C,H
label_14888: LD C,L
label_14889: LD D,A
label_14890: LD D,B
label_14891: LD D,C
label_14892: LD D,D
label_14893: LD D,E
label_14894: LD D,H
label_14895: LD D,L
label_14896: LD E,A
label_14897: LD E,B
label_14898: LD E,C
label_14899: LD E,D
label_14900: LD E,E
label_14901: LD E,H
label_14902: LD E,L
label_14903: LD H,A
label_14904: LD H,B
label_14905: LD H,C
label_14906: LD H,D
label_14907: LD H,E
label_14908: LD H,H
label_14909: LD H,L
label_14910: LD L,A
label_14911: LD L,B
label_14912: LD L,C
label_14913: LD L,D
label_14914: LD L,E
label_14915: LD L,H
label_14916: LD L,L
label_14917: LD A,$12
label_14918: LD B,$12
label_14919: LD C,$12
label_14920: LD D,$12
label_14921: LD E,$12
label_14922: LD H,$12
label_14923: LD L,$12
label_14924: LD A,(HL)
label_14925: LD B,(HL)
label_14926: LD C,(HL)
label_14927: LD D,(HL)
label_14928: LD E,(HL)
label_14929: LD H,(HL)
label_14930: LD L,(HL)
label_14931: LD A,(IX + 127)
label_14932: LD B,(IX + 127)
label_14933: LD C,(IX + 127)
label_14934: LD D,(IX + 127)
label_14935: LD E,(IX + 127)
label_14936: LD H,(IX + 127)
label_14937: LD L,(IX + 127)
label_14938: LD A,(IY - 128)
label_14939: LD B,(IY - 128)
label_14940: LD C,(IY - 128)
label_14941: LD D,(IY - 128)
label_14942: LD E,(IY - 128)
label_14943: LD H,(IY - 128)
label_14944: LD L,(IY - 128)
label_14945: LD (HL),A
label_14946: LD (HL),B
label_14947: LD (HL),C
label_14948: LD (HL),D
label_14949: LD (HL),E
label_14950: LD (HL),H
label_14951: LD (HL),L
label_14952: LD (IX + 127),A
label_14953: LD (IX + 127),B
label_14954: LD (IX + 127),C
label_14955: LD (IX + 127),D
label_14956: LD (IX + 127),E
label_14957: LD (IX + 127),H
label_14958: LD (IX + 127),L
label_14959: LD (IY - 128),A
label_14960: LD (IY - 128),B
label_14961: LD (IY - 128),C
label_14962: LD (IY - 128),D
label_14963: LD (IY - 128),E
label_14964: LD (IY - 128),H
label_14965: LD (IY - 128),L
label_14966: LD (HL),$12
label_14967: LD (IX + 127),$12
label_14968: LD (IY - 128),$12
label_14969: LD A,(BC)
label_14970: LD A,(DE)
label_14971: LD A,($5678)
label_14972: LD (BC),A
label_14973: LD (DE),A
label_14974: LD ($5678),A
label_14975: LD A,I
label_14976: LD A,R
label_14977: LD I,A
label_14978: LD R,A
label_14979: LD BC,$5678
label_14980: LD DE,$5678
label_14981: LD HL,$5678
label_14982: LD SP,$5678
label_14983: LD IX,$5678
label_14984: LD IY,$5678
label_14985: LD HL,($5678)
label_14986: LD BC,($5678)
label_14987: LD DE,($5678)
label_14988: LD HL,($5678)
label_14989: LD SP,($5678)
label_14990: LD IX,($5678)
label_14991: LD IY,($5678)
label_14992: LD ($5678),HL
label_14993: LD ($5678),BC
label_14994: LD ($5678),DE
label_14995: LD ($5678),HL
label_14996: LD ($5678),SP
label_14997: LD ($5678),IX
label_14998: LD ($5678),IY
label_14999: LD SP,HL
label_15000: LD SP,IX
label_15001: LD SP,IY
label_15002: PUSH BC
label_15003: PUSH DE
label_15004: PUSH HL
label_15005: PUSH AF
label_15006: PUSH IX
label_15007: PUSH IY
label_15008: POP BC
label_15009: POP DE
label_15010: POP HL
label_15011: POP AF
label_15012: POP IX
label_15013: POP IY
label_15014: EX DE,HL
label_15015: EX AF,AF'
label_15016: EXX
label_15017: EX (SP),HL
label_15018: EX (SP),IX
label_15019: EX (SP),IY
label_15020: LDI
label_15021: LDIR
label_15022: LDD
label_15023: LDDR
label_15024: CPI
label_15025: CPIR
label_15026: CPD
label_15027: CPDR
label_15028: ADD A,A
label_15029: ADD A,B
label_15030: ADD A,C
label_15031: ADD A,D
label_15032: ADD A,E
label_15033: ADD A,H
label_15034: ADD A,L
label_15035: ADD A,$12
label_15036: ADD A,(HL)
label_15037: ADD A,(IX + 127)
label_15038: ADD A,(IY - 128)
label_15039: ADC A,A
label_15040: ADC A,B
label_15041: ADC A,C
label_15042: ADC A,D
label_15043: ADC A,E
label_15044: ADC A,H
label_15045: ADC A,L
label_15046: ADC A,$12
label_15047: ADC A,(HL)
label_15048: ADC A,(IX + 127)
label_15049: ADC A,(IY - 128)
label_15050: SUB A
label_15051: SUB B
label_15052: SUB C
label_15053: SUB D
label_15054: SUB E
label_15055: SUB H
label_15056: SUB L
label_15057: SUB $12
label_15058: SUB (HL)
label_15059: SUB (IX + 127)
label_15060: SUB (IY - 128)
label_15061: SBC A,A
label_15062: SBC A,B
label_15063: SBC A,C
label_15064: SBC A,D
label_15065: SBC A,E
label_15066: SBC A,H
label_15067: SBC A,L
label_15068: SBC A,$12
label_15069: SBC A,(HL)
label_15070: SBC A,(IX + 127)
label_15071: SBC A,(IY - 128)
label_15072: AND A
label_15073: AND B
label_15074: AND C
label_15075: AND D
label_15076: AND E
label_15077: AND H
label_15078: AND L
label_15079: AND $12
label_15080: AND (HL)
label_15081: AND (IX + 127)
label_15082: AND (IY - 128)
label_15083: AND A
label_15084: AND B
label_15085: AND C
label_15086: AND D
label_15087: AND E
label_15088: AND H
label_15089: AND L
label_15090: AND $12
label_15091: AND (HL)
label_15092: AND (IX + 127)
label_15093: AND (IY - 128)
label_15094: OR A
label_15095: OR B
label_15096: OR C
label_15097: OR D
label_15098: OR E
label_15099: OR H
label_15100: OR L
label_15101: OR $12
label_15102: OR (HL)
label_15103: OR (IX + 127)
label_15104: OR (IY - 128)
label_15105: XOR A
label_15106: XOR B
label_15107: XOR C
label_15108: XOR D
label_15109: XOR E
label_15110: XOR H
label_15111: XOR L
label_15112: XOR $12
label_15113: XOR (HL)
label_15114: XOR (IX + 127)
label_15115: XOR (IY - 128)
label_15116: CP A
label_15117: CP B
label_15118: CP C
label_15119: CP D
label_15120: CP E
label_15121: CP H
label_15122: CP L
label_15123: CP $12
label_15124: CP (HL)
label_15125: CP (IX + 127)
label_15126: CP (IY - 128)
label_15127: INC A
label_15128: INC B
label_15129: INC C
label_15130: INC D
label_15131: INC E
label_15132: INC H
label_15133: INC L
label_15134: INC (HL)
label_15135: INC (IX + 127)
label_15136: INC (IY - 128)
label_15137: DEC A
label_15138: DEC B
label_15139: DEC C
label_15140: DEC D
label_15141: DEC E
label_15142: DEC H
label_15143: DEC L
label_15144: DEC (HL)
label_15145: DEC (IX + 127)
label_15146: DEC (IY - 128)
label_15147: DAA
label_15148: CPL
label_15149: NEG
label_15150: CCF
label_15151: SCF
label_15152: NOP
label_15153: HALT
label_15154: DI
label_15155: EI
label_15156: IM 0
label_15157: IM 1
label_15158: IM 2
label_15159: ADD HL,BC
label_15160: ADD HL,DE
label_15161: ADD HL,HL
label_15162: ADD HL,SP
label_15163: ADC HL,BC
label_15164: ADC HL,DE
label_15165: ADC HL,HL
label_15166: ADC HL,SP
label_15167: SBC HL,BC
label_15168: SBC HL,DE
label_15169: SBC HL,HL
label_15170: SBC HL,SP
label_15171: ADD IX,BC
label_15172: ADD IX,DE
label_15173: ADD IX,SP
label_15174: ADD IY,BC
label_15175: ADD IY,DE
label_15176: ADD IY,SP
label_15177: INC BC
label_15178: INC DE
label_15179: INC HL
label_15180: INC SP
label_15181: INC IX
label_15182: INC IY
label_15183: DEC BC
label_15184: DEC DE
label_15185: DEC HL
label_15186: DEC SP
label_15187: DEC IX
label_15188: DEC IY
label_15189: RLCA
label_15190: RLA
label_15191: RRCA
label_15192: RRA
label_15193: RLC A
label_15194: RLC B
label_15195: RLC C
label_15196: RLC D
label_15197: RLC E
label_15198: RLC H
label_15199: RLC L
label_15200: RLC (HL)
label_15201: RLC (IX + 127)
label_15202: RLC (IY - 128)
label_15203: RL A
label_15204: RL B
label_15205: RL C
label_15206: RL D
label_15207: RL E
label_15208: RL H
label_15209: RL L
label_15210: RL (HL)
label_15211: RL (IX + 127)
label_15212: RL (IY - 128)
label_15213: RRC A
label_15214: RRC B
label_15215: RRC C
label_15216: RRC D
label_15217: RRC E
label_15218: RRC H
label_15219: RRC L
label_15220: RRC (HL)
label_15221: RRC (IX + 127)
label_15222: RRC (IY - 128)
label_15223: RR A
label_15224: RR B
label_15225: RR C
label_15226: RR D
label_15227: RR E
label_15228: RR H
label_15229: RR L
label_15230: RR (HL)
label_15231: RR (IX + 127)
label_15232: RR (IY - 128)
label_15233: SLA A
label_15234: SLA B
label_15235: SLA C
label_15236: SLA D
label_15237: SLA E
label_15238: SLA H
label_15239: SLA L
label_15240: SLA (HL)
label_15241: SLA (IX + 127)
label_15242: SLA (IY - 128)
label_15243: SRA A
label_15244: SRA B
label_15245: SRA C
label_15246: SRA D
label_15247: SRA E
label_15248: SRA H
label_15249: SRA L
label_15250: SRA (HL)
label_15251: SRA (IX + 127)
label_15252: SRA (IY - 128)
label_15253: SRL A
label_15254: SRL B
label_15255: SRL C
label_15256: SRL D
label_15257: SRL E
label_15258: SRL H
label_15259: SRL L
label_15260: SRL (HL)
label_15261: SRL (IX + 127)
label_15262: SRL (IY - 128)
label_15263: RLD
label_15264: RRD
label_15265: BIT 0,A
label_15266: BIT 1,A
label_15267: BIT 2,A
label_15268: BIT 3,A
label_15269: BIT 4,A
label_15270: BIT 5,A
label_15271: BIT 6,A
label_15272: BIT 7,A
label_15273: BIT 0,B
label_15274: BIT 1,B
label_15275: BIT 2,B
label_15276: BIT 3,B
label_15277: BIT 4,B
label_15278: BIT 5,B
label_15279: BIT 6,B
label_15280: BIT 7,B
label_15281: BIT 0,C
label_15282: BIT 1,C
label_15283: BIT 2,C
label_15284: BIT 3,C
label_15285: BIT 4,C
label_15286: BIT 5,C
label_15287: BIT 6,C
label_15288: BIT 7,C
label_15289: BIT 0,D
label_15290: BIT 1,D
label_15291: BIT 2,D
label_15292: BIT 3,D
label_15293: BIT 4,D
label_15294: BIT 5,D
label_15295: BIT 6,D
label_15296: BIT 7,D
label_15297: BIT 0,E
label_15298: BIT 1,E
label_15299: BIT 2,E
label_15300: BIT 3,E
label_15301: BIT 4,E
label_15302: BIT 5,E
label_15303: BIT 6,E
label_15304: BIT 7,E
label_15305: BIT 0,H
label_15306: BIT 1,H
label_15307: BIT 2,H
label_15308: BIT 3,H
label_15309: BIT 4,H
label_15310: BIT 5,H
label_15311: BIT 6,H
label_15312: BIT 7,H
label_15313: BIT 0,L
label_15314: BIT 1,L
label_15315: BIT 2,L
label_15316: BIT 3,L
label_15317: BIT 4,L
label_15318: BIT 5,L
label_15319: BIT 6,L
label_15320: BIT 7,L
label_15321: BIT 0,(HL)
label_15322: BIT 1,(HL)
label_15323: BIT 2,(HL)
label_15324: BIT 3,(HL)
label_15325: BIT 4,(HL)
label_15326: BIT 5,(HL)
label_15327: BIT 6,(HL)
label_15328: BIT 7,(HL)
label_15329: BIT 0,(IX + 127)
label_15330: BIT 1,(IX + 127)
label_15331: BIT 2,(IX + 127)
label_15332: BIT 3,(IX + 127)
label_15333: BIT 4,(IX + 127)
label_15334: BIT 5,(IX + 127)
label_15335: BIT 6,(IX + 127)
label_15336: BIT 7,(IX + 127)
label_15337: BIT 0,(IY - 128)
label_15338: BIT 1,(IY - 128)
label_15339: BIT 2,(IY - 128)
label_15340: BIT 3,(IY - 128)
label_15341: BIT 4,(IY - 128)
label_15342: BIT 5,(IY - 128)
label_15343: BIT 6,(IY - 128)
label_15344: BIT 7,(IY - 128)
label_15345: SET 0,A
label_15346: SET 1,A
label_15347: SET 2,A
label_15348: SET 3,A
label_15349: SET 4,A
label_15350: SET 5,A
label_15351: SET 6,A
label_15352: SET 7,A
label_15353: SET 0,B
label_15354: SET 1,B
label_15355: SET 2,B
label_15356: SET 3,B
label_15357: SET 4,B
label_15358: SET 5,B
label_15359: SET 6,B
label_15360: SET 7,B
label_15361: SET 0,C
label_15362: SET 1,C
label_15363: SET 2,C
label_15364: SET 3,C
label_15365: SET 4,C
label_15366: SET 5,C
label_15367: SET 6,C
label_15368: SET 7,C
label_15369: SET 0,D
label_15370: SET 1,D
label_15371: SET 2,D
label_15372: SET 3,D
label_15373: SET 4,D
label_15374: SET 5,D
label_15375: SET 6,D
label_15376: SET 7,D
label_15377: SET 0,E
label_15378: SET 1,E
label_15379: SET 2,E
label_15380: SET 3,E
label_15381: SET 4,E
label_15382: SET 5,E
label_15383: SET 6,E
label_15384: SET 7,E
label_15385: SET 0,H
label_15386: SET 1,H
label_15387: SET 2,H
label_15388: SET 3,H
label_15389: SET 4,H
label_15390: SET 5,H
label_15391: SET 6,H
label_15392: SET 7,H
label_15393: SET 0,L
label_15394: SET 1,L
label_15395: SET 2,L
label_15396: SET 3,L
label_15397: SET 4,L
label_15398: SET 5,L
label_15399: SET 6,L
label_15400: SET 7,L
label_15401: SET 0,(HL)
label_15402: SET 1,(HL)
label_15403: SET 2,(HL)
label_15404: SET 3,(HL)
label_15405: SET 4,(HL)
label_15406: SET 5,(HL)
label_15407: SET 6,(HL)
label_15408: SET 7,(HL)
label_15409: SET 0,(IX + 127)
label_15410: SET 1,(IX + 127)
label_15411: SET 2,(IX + 127)
label_15412: SET 3,(IX + 127)
label_15413: SET 4,(IX + 127)
label_15414: SET 5,(IX + 127)
label_15415: SET 6,(IX + 127)
label_15416: SET 7,(IX + 127)
label_15417: SET 0,(IY - 128)
label_15418: SET 1,(IY - 128)
label_15419: SET 2,(IY - 128)
label_15420: SET 3,(IY - 128)
label_15421: SET 4,(IY - 128)
label_15422: SET 5,(IY - 128)
label_15423: SET 6,(IY - 128)
label_15424: SET 7,(IY - 128)
label_15425: RES 0,A
label_15426: RES 1,A
label_15427: RES 2,A
label_15428: RES 3,A
label_15429: RES 4,A
label_15430: RES 5,A
label_15431: RES 6,A
label_15432: RES 7,A
label_15433: RES 0,B
label_15434: RES 1,B
label_15435: RES 2,B
label_15436: RES 3,B
label_15437: RES 4,B
label_15438: RES 5,B
label_15439: RES 6,B
label_15440: RES 7,B
label_15441: RES 0,C
label_15442: RES 1,C
label_15443: RES 2,C
label_15444: RES 3,C
label_15445: RES 4,C
label_15446: RES 5,C
label_15447: RES 6,C
label_15448: RES 7,C
label_15449: RES 0,D
label_15450: RES 1,D
label_15451: RES 2,D
label_15452: RES 3,D
label_15453: RES 4,D
label_15454: RES 5,D
label_15455: RES 6,D
label_15456: RES 7,D
label_15457: RES 0,E
label_15458: RES 1,E
label_15459: RES 2,E
label_15460: RES 3,E
label_15461: RES 4,E
label_15462: RES 5,E
label_15463: RES 6,E
label_15464: RES 7,E
label_15465: RES 0,H
label_15466: RES 1,H
label_15467: RES 2,H
label_15468: RES 3,H
label_15469: RES 4,H
label_15470: RES 5,H
label_15471: RES 6,H
label_15472: RES 7,H
label_15473: RES 0,L
label_15474: RES 1,L
label_15475: RES 2,L
label_15476: RES 3,L
label_15477: RES 4,L
label_15478: RES 5,L
label_15479: RES 6,L
label_15480: RES 7,L
label_15481: RES 0,(HL)
label_15482: RES 1,(HL)
label_15483: RES 2,(HL)
label_15484: RES 3,(HL)
label_15485: RES 4,(HL)
label_15486: RES 5,(HL)
label_15487: RES 6,(HL)
label_15488: RES 7,(HL)
label_15489: RES 0,(IX + 127)
label_15490: RES 1,(IX + 127)
label_15491: RES 2,(IX + 127)
label_15492: RES 3,(IX + 127)
label_15493: RES 4,(IX + 127)
label_15494: RES 5,(IX + 127)
label_15495: RES 6,(IX + 127)
label_15496: RES 7,(IX + 127)
label_15497: RES 0,(IY - 128)
label_15498: RES 1,(IY - 128)
label_15499: RES 2,(IY - 128)
label_15500: RES 3,(IY - 128)
label_15501: RES 4,(IY - 128)
label_15502: RES 5,(IY - 128)
label_15503: RES 6,(IY - 128)
label_15504: RES 7,(IY - 128)
label_15505: JP $5678
label_15506: JP NZ,$5678
label_15507: JP Z,$5678
label_15508: JP NC,$5678
label_15509: JP C,$5678
label_15510: JP PO,$5678
label_15511: JP PE,$5678
label_15512: JP P,$5678
label_15513: JP M,$5678
label_15514: JR $ + 2
label_15515: JR NZ,$ + 2
label_15516: JR Z,$ + 2
label_15517: JR NC,$ + 2
label_15518: JR C,$ + 2
label_15519: JP (HL)
label_15520: JP (IX)
label_15521: JP (IY)
label_15522: DJNZ $ + 2
label_15523: CALL $5678
label_15524: CALL NZ,$5678
label_15525: CALL Z,$5678
label_15526: CALL NC,$5678
label_15527: CALL C,$5678
label_15528: CALL PO,$5678
label_15529: CALL PE,$5678
label_15530: CALL P,$5678
label_15531: CALL M,$5678
label_15532: RET
label_15533: RET NZ
label_15534: RET Z
label_15535: RET NC
label_15536: RET C
label_15537: RET PO
label_15538: RET PE
label_15539: RET P
label_15540: RET M
label_15541: RETI
label_15542: RETN
label_15543: RST $00
label_15544: RST $08
label_15545: RST $10
label_15546: RST $18
label_15547: RST $20
label_15548: RST $28
label_15549: RST $30
label_15550: RST $38
label_15551: IN A,($12)
label_15552: IN A,(C)
label_15553: IN B,(C)
label_15554: IN C,(C)
label_15555: IN D,(C)
label_15556: IN E,(C)
label_15557: IN H,(C)
label_15558: IN L,(C)
label_15559: IN F,(C)
label_15560: INI
label_15561: INIR
label_15562: IND
label_15563: INDR
label_15564: OUT ($12),A
label_15565: OUT (C),A
label_15566: OUT (C),B
label_15567: OUT (C),C
label_15568: OUT (C),D
label_15569: OUT (C),E
label_15570: OUT (C),H
label_15571: OUT (C),L
label_15572: OUTI
label_15573: OTIR
label_15574: OUTD
label_15575: OTDR
label_15576: LD A,A
label_15577: LD A,B
label_15578: LD A,C
label_15579: LD A,D
label_15580: LD A,E
label_15581: LD A,H
label_15582: LD A,L
label_15583: LD B,A
label_15584: LD B,B
label_15585: LD B,C
label_15586: LD B,D
label_15587: LD B,E
label_15588: LD B,H
label_15589: LD B,L
label_15590: LD C,A
label_15591: LD C,B
label_15592: LD C,C
label_15593: LD C,D
label_15594: LD C,E
label_15595: LD C,H
label_15596: LD C,L
label_15597: LD D,A
label_15598: LD D,B
label_15599: LD D,C
label_15600: LD D,D
label_15601: LD D,E
label_15602: LD D,H
label_15603: LD D,L
label_15604: LD E,A
label_15605: LD E,B
label_15606: LD E,C
label_15607: LD E,D
label_15608: LD E,E
label_15609: LD E,H
label_15610: LD E,L
label_15611: LD H,A
label_15612: LD H,B
label_15613: LD H,C
label_15614: LD H,D
label_15615: LD H,E
label_15616: LD H,H
label_15617: LD H,L
label_15618: LD L,A
label_15619: LD L,B
label_15620: LD L,C
label_15621: LD L,D
label_15622: LD L,E
label_15623: LD L,H
label_15624: LD L,L
label_15625: LD A,$12
label_15626: LD B,$12
label_15627: LD C,$12
label_15628: LD D,$12
label_15629: LD E,$12
label_15630: LD H,$12
label_15631: LD L,$12
label_15632: LD A,(HL)
label_15633: LD B,(HL)
label_15634: LD C,(HL)
label_15635: LD D,(HL)
label_15636: LD E,(HL)
label_15637: LD H,(HL)
label_15638: LD L,(HL)
label_15639: LD A,(IX + 127)
label_15640: LD B,(IX + 127)
label_15641: LD C,(IX + 127)
label_15642: LD D,(IX + 127)
label_15643: LD E,(IX + 127)
label_15644: LD H,(IX + 127)
label_15645: LD L,(IX + 127)
label_15646: LD A,(IY - 128)
label_15647: LD B,(IY - 128)
label_15648: LD C,(IY - 128)
label_15649: LD D,(IY - 128)
label_15650: LD E,(IY - 128)
label_15651: LD H,(IY - 128)
label_15652: LD L,(IY - 128)
label_15653: LD (HL),A
label_15654: LD (HL),B
label_15655: LD (HL),C
label_15656: LD (HL),D
label_15657: LD (HL),E
label_15658: LD (HL),H
label_15659: LD (HL),L
label_15660: LD (IX + 127),A
label_15661: LD (IX + 127),B
label_15662: LD (IX + 127),C
label_15663: LD (IX + 127),D
label_15664: LD (IX + 127),E
label_15665: LD (IX + 127),H
label_15666: LD (IX + 127),L
label_15667: LD (IY - 128),A
label_15668: LD (IY - 128),B
label_15669: LD (IY - 128),C
label_15670: LD (IY - 128),D
label_15671: LD (IY - 128),E
label_15672: LD (IY - 128),H
label_15673: LD (IY - 128),L
label_15674: LD (HL),$12
label_15675: LD (IX + 127),$12
label_15676: LD (IY - 128),$12
label_15677: LD A,(BC)
label_15678: LD A,(DE)
label_15679: LD A,($5678)
label_15680: LD (BC),A
label_15681: LD (DE),A
label_15682: LD ($5678),A
label_15683: LD A,I
label_15684: LD A,R
label_15685: LD I,A
label_15686: LD R,A
label_15687: LD BC,$5678
label_15688: LD DE,$5678
label_15689: LD HL,$5678
label_15690: LD SP,$5678
label_15691: LD IX,$5678
label_15692: LD IY,$5678
label_15693: LD HL,($5678)
label_15694: LD BC,($5678)
label_15695: LD DE,($5678)
label_15696: LD HL,($5678)
label_15697: LD SP,($5678)
label_15698: LD IX,($5678)
label_15699: LD IY,($5678)
label_15700: LD ($5678),HL
label_15701: LD ($5678),BC
label_15702: LD ($5678),DE
label_15703: LD ($5678),HL
label_15704: LD ($5678),SP
label_15705: LD ($5678),IX
label_15706: LD ($5678),IY
label_15707: LD SP,HL
label_15708: LD SP,IX
label_15709: LD SP,IY
label_15710: PUSH BC
label_15711: PUSH DE
label_15712: PUSH HL
label_15713: PUSH AF
label_15714: PUSH IX
label_15715: PUSH IY
label_15716: POP BC
label_15717: POP DE
label_15718: POP HL
label_15719: POP AF
label_15720: POP IX
label_15721: POP IY
label_15722: EX DE,HL
label_15723: EX AF,AF'
label_15724: EXX
label_15725: EX (SP),HL
label_15726: EX (SP),IX
label_15727: EX (SP),IY
label_15728: LDI
label_15729: LDIR
label_15730: LDD
label_15731: LDDR
label_15732: CPI
label_15733: CPIR
label_15734: CPD
label_15735: CPDR
label_15736: ADD A,A
label_15737: ADD A,B
label_15738: ADD A,C
label_15739: ADD A,D
label_15740: ADD A,E
label_15741: ADD A,H
label_15742: ADD A,L
label_15743: ADD A,$12
label_15744: ADD A,(HL)
label_15745: ADD A,(IX + 127)
label_15746: ADD A,(IY - 128)
label_15747: ADC A,A
label_15748: ADC A,B
label_15749: ADC A,C
label_15750: ADC A,D
label_15751: ADC A,E
label_15752: ADC A,H
label_15753: ADC A,L
label_15754: ADC A,$12
label_15755: ADC A,(HL)
label_15756: ADC A,(IX + 127)
label_15757: ADC A,(IY - 128)
label_15758: SUB A
label_15759: SUB B
label_15760: SUB C
label_15761: SUB D
label_15762: SUB E
label_15763: SUB H
label_15764: SUB L
label_15765: SUB $12
label_15766: SUB (HL)
label_15767: SUB (IX + 127)
label_15768: SUB (IY - 128)
label_15769: SBC A,A
label_15770: SBC A,B
label_15771: SBC A,C
label_15772: SBC A,D
label_15773: SBC A,E
label_15774: SBC A,H
label_15775: SBC A,L
label_15776: SBC A,$12
label_15777: SBC A,(HL)
label_15778: SBC A,(IX + 127)
label_15779: SBC A,(IY - 128)
label_15780: AND A
label_15781: AND B
label_15782: AND C
label_15783: AND D
label_15784: AND E
label_15785: AND H
label_15786: AND L
label_15787: AND $12
label_15788: AND (HL)
label_15789: AND (IX + 127)
label_15790: AND (IY - 128)
label_15791: AND A
label_15792: AND B
label_15793: AND C
label_15794: AND D
label_15795: AND E
label_15796: AND H
label_15797: AND L
label_15798: AND $12
label_15799: AND (HL)
label_15800: AND (IX + 127)
label_15801: AND (IY - 128)
label_15802: OR A
label_15803: OR B
label_15804: OR C
label_15805: OR D
label_15806: OR E
label_15807: OR H
label_15808: OR L
label_15809: OR $12
label_15810: OR (HL)
label_15811: OR (IX + 127)
label_15812: OR (IY - 128)
label_15813: XOR A
label_15814: XOR B
label_15815: XOR C
label_15816: XOR D
label_15817: XOR E
label_15818: XOR H
label_15819: XOR L
label_15820: XOR $12
label_15821: XOR (HL)
label_15822: XOR (IX + 127)
label_15823: XOR (IY - 128)
label_15824: CP A
label_15825: CP B
label_15826: CP C
label_15827: CP D
label_15828: CP E
label_15829: CP H
label_15830: CP L
label_15831: CP $12
label_15832: CP (HL)
label_15833: CP (IX + 127)
label_15834: CP (IY - 128)
label_15835: INC A
label_15836: INC B
label_15837: INC C
label_15838: INC D
label_15839: INC E
label_15840: INC H
label_15841: INC L
label_15842: INC (HL)
label_15843: INC (IX + 127)
label_15844: INC (IY - 128)
label_15845: DEC A
label_15846: DEC B
label_15847: DEC C
label_15848: DEC D
label_15849: DEC E
label_15850: DEC H
label_15851: DEC L
label_15852: DEC (HL)
label_15853: DEC (IX + 127)
label_15854: DEC (IY - 128)
label_15855: DAA
label_15856: CPL
label_15857: NEG
label_15858: CCF
label_15859: SCF
label_15860: NOP
label_15861: HALT
label_15862: DI
label_15863: EI
label_15864: IM 0
label_15865: IM 1
label_15866: IM 2
label_15867: ADD HL,BC
label_15868: ADD HL,DE
label_15869: ADD HL,HL
label_15870: ADD HL,SP
label_15871: ADC HL,BC
label_15872: ADC HL,DE
label_15873: ADC HL,HL
label_15874: ADC HL,SP
label_15875: SBC HL,BC
label_15876: SBC HL,DE
label_15877: SBC HL,HL
label_15878: SBC HL,SP
label_15879: ADD IX,BC
label_15880: ADD IX,DE
label_15881: ADD IX,SP
label_15882: ADD IY,BC
label_15883: ADD IY,DE
label_15884: ADD IY,SP
label_15885: INC BC
label_15886: INC DE
label_15887: INC HL
label_15888: INC SP
label_15889: INC IX
label_15890: INC IY
label_15891: DEC BC
label_15892: DEC DE
label_15893: DEC HL
label_15894: DEC SP
label_15895: DEC IX
label_15896: DEC IY
label_15897: RLCA
label_15898: RLA
label_15899: RRCA
label_15900: RRA
label_15901: RLC A
label_15902: RLC B
label_15903: RLC C
label_15904: RLC D
label_15905: RLC E
label_15906: RLC H
label_15907: RLC L
label_15908: RLC (HL)
label_15909: RLC (IX + 127)
label_15910: RLC (IY - 128)
label_15911: RL A
label_15912: RL B
label_15913: RL C
label_15914: RL D
label_15915: RL E
label_15916: RL H
label_15917: RL L
label_15918: RL (HL)
label_15919: RL (IX + 127)
label_15920: RL (IY - 128)
label_15921: RRC A
label_15922: RRC B
label_15923: RRC C
label_15924: RRC D
label_15925: RRC E
label_15926: RRC H
label_15927: RRC L
label_15928: RRC (HL)
label_15929: RRC (IX + 127)
label_15930: RRC (IY - 128)
label_15931: RR A
label_15932: RR B
label_15933: RR C
label_15934: RR D
label_15935: RR E
label_15936: RR H
label_15937: RR L
label_15938: RR (HL)
label_15939: RR (IX + 127)
label_15940: RR (IY - 128)
label_15941: SLA A
label_15942: SLA B
label_15943: SLA C
label_15944: SLA D
label_15945: SLA E
label_15946: SLA H
label_15947: SLA L
label_15948: SLA (HL)
label_15949: SLA (IX + 127)
label_15950: SLA (IY - 128)
label_15951: SRA A
label_15952: SRA B
label_15953: SRA C
label_15954: SRA D
label_15955: SRA E
label_15956: SRA H
label_15957: SRA L
label_15958: SRA (HL)
label_15959: SRA (IX + 127)
label_15960: SRA (IY - 128)
label_15961: SRL A
label_15962: SRL B
label_15963: SRL C
label_15964: SRL D
label_15965: SRL E
label_15966: SRL H
label_15967: SRL L
label_15968: SRL (HL)
label_15969: SRL (IX + 127)
label_15970: SRL (IY - 128)
label_15971: RLD
label_15972: RRD
label_15973: BIT 0,A
label_15974: BIT 1,A
label_15975: BIT 2,A
label_15976: BIT 3,A
label_15977: BIT 4,A
label_15978: BIT 5,A
label_15979: BIT 6,A
label_15980: BIT 7,A
label_15981: BIT 0,B
label_15982: BIT 1,B
label_15983: BIT 2,B
label_15984: BIT 3,B
label_15985: BIT 4,B
label_15986: BIT 5,B
label_15987: BIT 6,B
label_15988: BIT 7,B
label_15989: BIT 0,C
label_15990: BIT 1,C
label_15991: BIT 2,C
label_15992: BIT 3,C
label_15993: BIT 4,C
label_15994: BIT 5,C
label_15995: BIT 6,C
label_15996: BIT 7,C
label_15997: BIT 0,D
label_15998: BIT 1,D
label_15999: BIT 2,D
label_16000: BIT 3,D
label_16001: BIT 4,D
label_16002: BIT 5,D
label_16003: BIT 6,D
label_16004: BIT 7,D
label_16005: BIT 0,E
label_16006: BIT 1,E
label_16007: BIT 2,E
label_16008: BIT 3,E
label_16009: BIT 4,E
label_16010: BIT 5,E
label_16011: BIT 6,E
label_16012: BIT 7,E
label_16013: BIT 0,H
label_16014: BIT 1,H
label_16015: BIT 2,H
label_16016: BIT 3,H
label_16017: BIT 4,H
label_16018: BIT 5,H
label_16019: BIT 6,H
label_16020: BIT 7,H
label_16021: BIT 0,L
label_16022: BIT 1,L
label_16023: BIT 2,L
label_16024: BIT 3,L
label_16025: BIT 4,L
label_16026: BIT 5,L
label_16027: BIT 6,L
label_16028: BIT 7,L
label_16029: BIT 0,(HL)
label_16030: BIT 1,(HL)
label_16031: BIT 2,(HL)
label_16032: BIT 3,(HL)
label_16033: BIT 4,(HL)
label_16034: BIT 5,(HL)
label_16035: BIT 6,(HL)
label_16036: BIT 7,(HL)
label_16037: BIT 0,(IX + 127)
label_16038: BIT 1,(IX + 127)
label_16039: BIT 2,(IX + 127)
label_16040: BIT 3,(IX + 127)
label_16041: BIT 4,(IX + 127)
label_16042: BIT 5,(IX + 127)
label_16043: BIT 6,(IX + 127)
label_16044: BIT 7,(IX + 127)
label_16045: BIT 0,(IY - 128)
label_16046: BIT 1,(IY - 128)
label_16047: BIT 2,(IY - 128)
label_16048: BIT 3,(IY - 128)
label_16049: BIT 4,(IY - 128)
label_16050: BIT 5,(IY - 128)
label_16051: BIT 6,(IY - 128)
label_16052: BIT 7,(IY - 128)
label_16053: SET 0,A
label_16054: SET 1,A
label_16055: SET 2,A
label_16056: SET 3,A
label_16057: SET 4,A
label_16058: SET 5,A
label_16059: SET 6,A
label_16060: SET 7,A
label_16061: SET 0,B
label_16062: SET 1,B
label_16063: SET 2,B
label_16064: SET 3,B
label_16065: SET 4,B
label_16066: SET 5,B
label_16067: SET 6,B
label_16068: SET 7,B
label_16069: SET 0,C
label_16070: SET 1,C
label_16071: SET 2,C
label_16072: SET 3,C
label_16073: SET 4,C
label_16074: SET 5,C
label_16075: SET 6,C
label_16076: SET 7,C
label_16077: SET 0,D
label_16078: SET 1,D
label_16079: SET 2,D
label_16080: SET 3,D
label_16081: SET 4,D
label_16082: SET 5,D
label_16083: SET 6,D
label_16084: SET 7,D
label_16085: SET 0,E
label_16086: SET 1,E
label_16087: SET 2,E
label_16088: SET 3,E
label_16089: SET 4,E
label_16090: SET 5,E
label_16091: SET 6,E
label_16092: SET 7,E
label_16093: SET 0,H
label_16094: SET 1,H
label_16095: SET 2,H
label_16096: SET 3,H
label_16097: SET 4,H
label_16098: SET 5,H
label_16099: SET 6,H
label_16100: SET 7,H
label_16101: SET 0,L
label_16102: SET 1,L
label_16103: SET 2,L
label_16104: SET 3,L
label_16105: SET 4,L
label_16106: SET 5,L
label_16107: SET 6,L
label_16108: SET 7,L
label_16109: SET 0,(HL)
label_16110: SET 1,(HL)
label_16111: SET 2,(HL)
label_16112: SET 3,(HL)
label_16113: SET 4,(HL)
label_16114: SET 5,(HL)
label_16115: SET 6,(HL)
label_16116: SET 7,(HL)
label_16117: SET 0,(IX + 127)
label_16118: SET 1,(IX + 127)
label_16119: SET 2,(IX + 127)
label_16120: SET 3,(IX + 127)
label_16121: SET 4,(IX + 127)
label_16122: SET 5,(IX + 127)
label_16123: SET 6,(IX + 127)
label_16124: SET 7,(IX + 127)
label_16125: SET 0,(IY - 128)
label_16126: SET 1,(IY - 128)
label_16127: SET 2,(IY - 128)
label_16128: SET 3,(IY - 128)
label_16129: SET 4,(IY - 128)
label_16130: SET 5,(IY - 128)
label_16131: SET 6,(IY - 128)
label_16132: SET 7,(IY - 128)
label_16133: RES 0,A
label_16134: RES 1,A
label_16135: RES 2,A
label_16136: RES 3,A
label_16137: RES 4,A
label_16138: RES 5,A
label_16139: RES 6,A
label_16140: RES 7,A
label_16141: RES 0,B
label_16142: RES 1,B
label_16143: RES 2,B
label_16144: RES 3,B
label_16145: RES 4,B
label_16146: RES 5,B
label_16147: RES 6,B
label_16148: RES 7,B
label_16149: RES 0,C
label_16150: RES 1,C
label_16151: RES 2,C
label_16152: RES 3,C
label_16153: RES 4,C
label_16154: RES 5,C
label_16155: RES 6,C
label_16156: RES 7,C
label_16157: RES 0,D
label_16158: RES 1,D
label_16159: RES 2,D
label_16160: RES 3,D
label_16161: RES 4,D
label_16162: RES 5,D
label_16163: RES 6,D
label_16164: RES 7,D
label_16165: RES 0,E
label_16166: RES 1,E
label_16167: RES 2,E
label_16168: RES 3,E
label_16169: RES 4,E
label_16170: RES 5,E
label_16171: RES 6,E
label_16172: RES 7,E
label_16173: RES 0,H
label_16174: RES 1,H
label_16175: RES 2,H
label_16176: RES 3,H
label_16177: RES 4,H
label_16178: RES 5,H
label_16179: RES 6,H
label_16180: RES 7,H
label_16181: RES 0,L
label_16182: RES 1,L
label_16183: RES 2,L
label_16184: RES 3,L
label_16185: RES 4,L
label_16186: RES 5,L
label_16187: RES 6,L
label_16188: RES 7,L
label_16189: RES 0,(HL)
label_16190: RES 1,(HL)
label_16191: RES 2,(HL)
label_16192: RES 3,(HL)
label_16193: RES 4,(HL)
label_16194: RES 5,(HL)
label_16195: RES 6,(HL)
label_16196: RES 7,(HL)
label_16197: RES 0,(IX + 127)
label_16198: RES 1,(IX + 127)
label_16199: RES 2,(IX + 127)
label_16200: RES 3,(IX + 127)
label_16201: RES 4,(IX + 127)
label_16202: RES 5,(IX + 127)
label_16203: RES 6,(IX + 127)
label_16204: RES 7,(IX + 127)
label_16205: RES 0,(IY - 128)
label_16206: RES 1,(IY - 128)
label_16207: RES 2,(IY - 128)
label_16208: RES 3,(IY - 128)
label_16209: RES 4,(IY - 128)
label_16210: RES 5,(IY - 128)
label_16211: RES 6,(IY - 128)
label_16212: RES 7,(IY - 128)
label_16213: JP $5678
label_16214: JP NZ,$5678
label_16215: JP Z,$5678
label_16216: JP NC,$5678
label_16217: JP C,$5678
label_16218: JP PO,$5678
label_16219: JP PE,$5678
label_16220: JP P,$5678
label_16221: JP M,$5678
label_16222: JR $ + 2
label_16223: JR NZ,$ + 2
label_16224: JR Z,$ + 2
label_16225: JR NC,$ + 2
label_16226: JR C,$ + 2
label_16227: JP (HL)
label_16228: JP (IX)
label_16229: JP (IY)
label_16230: DJNZ $ + 2
label_16231: CALL $5678
label_16232: CALL NZ,$5678
label_16233: CALL Z,$5678
label_16234: CALL NC,$5678
label_16235: CALL C,$5678
label_16236: CALL PO,$5678
label_16237: CALL PE,$5678
label_16238: CALL P,$5678
label_16239: CALL M,$5678
label_16240: RET
label_16241: RET NZ
label_16242: RET Z
label_16243: RET NC
label_16244: RET C
label_16245: RET PO
label_16246: RET PE
label_16247: RET P
label_16248: RET M
label_16249: RETI
label_16250: RETN
label_16251: RST $00
label_16252: RST $08
label_16253: RST $10
label_16254: RST $18
label_16255: RST $20
label_16256: RST $28
label_16257: RST $30
label_16258: RST $38
label_16259: IN A,($12)
label_16260: IN A,(C)
label_16261: IN B,(C)
label_16262: IN C,(C)
label_16263: IN D,(C)
label_16264: IN E,(C)
label_16265: IN H,(C)
label_16266: IN L,(C)
label_16267: IN F,(C)
label_16268: INI
label_16269: INIR
label_16270: IND
label_16271: INDR
label_16272: OUT ($12),A
label_16273: OUT (C),A
label_16274: OUT (C),B
label_16275: OUT (C),C
label_16276: OUT (C),D
label_16277: OUT (C),E
label_16278: OUT (C),H
label_16279: OUT (C),L
label_16280: OUTI
label_16281: OTIR
label_16282: OUTD
label_16283: OTDR
label_16284: LD A,A
label_16285: LD A,B
label_16286: LD A,C
label_16287: LD A,D
label_16288: LD A,E
label_16289: LD A,H
label_16290: LD A,L
label_16291: LD B,A
label_16292: LD B,B
label_16293: LD B,C
label_16294: LD B,D
label_16295: LD B,E
label_16296: LD B,H
label_16297: LD B,L
label_16298: LD C,A
label_16299: LD C,B
label_16300: LD C,C
label_16301: LD C,D
label_16302: LD C,E
label_16303: LD C,H
label_16304: LD C,L
label_16305: LD D,A
label_16306: LD D,B
label_16307: LD D,C
label_16308: LD D,D
label_16309: LD D,E
label_16310: LD D,H
label_16311: LD D,L
label_16312: LD E,A
label_16313: LD E,B
label_16314: LD E,C
label_16315: LD E,D
label_16316: LD E,E
label_16317: LD E,H
label_16318: LD E,L
label_16319: LD H,A
label_16320: LD H,B
label_16321: LD H,C
label_16322: LD H,D
label_16323: LD H,E
label_16324: LD H,H
label_16325: LD H,L
label_16326: LD L,A
label_16327: LD L,B
label_16328: LD L,C
label_16329: LD L,D
label_16330: LD L,E
label_16331: LD L,H
label_16332: LD L,L
label_16333: LD A,$12
label_16334: LD B,$12
label_16335: LD C,$12
label_16336: LD D,$12
label_16337: LD E,$12
label_16338: LD H,$12
label_16339: LD L,$12
label_16340: LD A,(HL)
label_16341: LD B,(HL)
label_16342: LD C,(HL)
label_16343: LD D,(HL)
label_16344: LD E,(HL)
label_16345: LD H,(HL)
label_16346: LD L,(HL)
label_16347: LD A,(IX + 127)
label_16348: LD B,(IX + 127)
label_16349: LD C,(IX + 127)
label_16350: LD D,(IX + 127)
label_16351: LD E,(IX + 127)
label_16352: LD H,(IX + 127)
label_16353: LD L,(IX + 127)
label_16354: LD A,(IY - 128)
label_16355: LD B,(IY - 128)
label_16356: LD C,(IY - 128)
label_16357: LD D,(IY - 128)
label_16358: LD E,(IY - 128)
label_16359: LD H,(IY - 128)
label_16360: LD L,(IY - 128)
label_16361: LD (HL),A
label_16362: LD (HL),B
label_16363: LD (HL),C
label_16364: LD (HL),D
label_16365: LD (HL),E
label_16366: LD (HL),H
label_16367: LD (HL),L
label_16368: LD (IX + 127),A
label_16369: LD (IX + 127),B
label_16370: LD (IX + 127),C
label_16371: LD (IX + 127),D
label_16372: LD (IX + 127),E
label_16373: LD (IX + 127),H
label_16374: LD (IX + 127),L
label_16375: LD (IY - 128),A
label_16376: LD (IY - 128),B
label_16377: LD (IY - 128),C
label_16378: LD (IY - 128),D
label_16379: LD (IY - 128),E
label_16380: LD (IY - 128),H
label_16381: LD (IY - 128),L
label_16382: LD (HL),$12
label_16383: LD (IX + 127),$12
label_16384: LD (IY - 128),$12
label_16385: LD A,(BC)
label_16386: LD A,(DE)
label_16387: LD A,($5678)
label_16388: LD (BC),A
label_16389: LD (DE),A
label_16390: LD ($5678),A
label_16391: LD A,I
label_16392: LD A,R
label_16393: LD I,A
label_16394: LD R,A
label_16395: LD BC,$5678
label_16396: LD DE,$5678
label_16397: LD HL,$5678
label_16398: LD SP,$5678
label_16399: LD IX,$5678
label_16400: LD IY,$5678
label_16401: LD HL,($5678)
label_16402: LD BC,($5678)
label_16403: LD DE,($5678)
label_16404: LD HL,($5678)
label_16405: LD SP,($5678)
label_16406: LD IX,($5678)
label_16407: LD IY,($5678)
label_16408: LD ($5678),HL
label_16409: LD ($5678),BC
label_16410: LD ($5678),DE
label_16411: LD ($5678),HL
label_16412: LD ($5678),SP
label_16413: LD ($5678),IX
label_16414: LD ($5678),IY
label_16415: LD SP,HL
label_16416: LD SP,IX
label_16417: LD SP,IY
label_16418: PUSH BC
label_16419: PUSH DE
label_16420: PUSH HL
label_16421: PUSH AF
label_16422: PUSH IX
label_16423: PUSH IY
label_16424: POP BC
label_16425: POP DE
label_16426: POP HL
label_16427: POP AF
label_16428: POP IX
label_16429: POP IY
label_16430: EX DE,HL
label_16431: EX AF,AF'
label_16432: EXX
label_16433: EX (SP),HL
label_16434: EX (SP),IX
label_16435: EX (SP),IY
label_16436: LDI
label_16437: LDIR
label_16438: LDD
label_16439: LDDR
label_16440: CPI
label_16441: CPIR
label_16442: CPD
label_16443: CPDR
label_16444: ADD A,A
label_16445: ADD A,B
label_16446: ADD A,C
label_16447: ADD A,D
label_16448: ADD A,E
label_16449: ADD A,H
label_16450: ADD A,L
label_16451: ADD A,$12
label_16452: ADD A,(HL)
label_16453: ADD A,(IX + 127)
label_16454: ADD A,(IY - 128)
label_16455: ADC A,A
label_16456: ADC A,B
label_16457: ADC A,C
label_16458: ADC A,D
label_16459: ADC A,E
label_16460: ADC A,H
label_16461: ADC A,L
label_16462: ADC A,$12
label_16463: ADC A,(HL)
label_16464: ADC A,(IX + 127)
label_16465: ADC A,(IY - 128)
label_16466: SUB A
label_16467: SUB B
label_16468: SUB C
label_16469: SUB D
label_16470: SUB E
label_16471: SUB H
label_16472: SUB L
label_16473: SUB $12
label_16474: SUB (HL)
label_16475: SUB (IX + 127)
label_16476: SUB (IY - 128)
label_16477: SBC A,A
label_16478: SBC A,B
label_16479: SBC A,C
label_16480: SBC A,D
label_16481: SBC A,E
label_16482: SBC A,H
label_16483: SBC A,L
label_16484: SBC A,$12
label_16485: SBC A,(HL)
label_16486: SBC A,(IX + 127)
label_16487: SBC A,(IY - 128)
label_16488: AND A
label_16489: AND B
label_16490: AND C
label_16491: AND D
label_16492: AND E
label_16493: AND H
label_16494: AND L
label_16495: AND $12
label_16496: AND (HL)
label_16497: AND (IX + 127)
label_16498: AND (IY - 128)
label_16499: AND A
label_16500: AND B
label_16501: AND C
label_16502: AND D
label_16503: AND E
label_16504: AND H
label_16505: AND L
label_16506: AND $12
label_16507: AND (HL)
label_16508: AND (IX + 127)
label_16509: AND (IY - 128)
label_16510: OR A
label_16511: OR B
label_16512: OR C
label_16513: OR D
label_16514: OR E
label_16515: OR H
label_16516: OR L
label_16517: OR $12
label_16518: OR (HL)
label_16519: OR (IX + 127)
label_16520: OR (IY - 128)
label_16521: XOR A
label_16522: XOR B
label_16523: XOR C
label_16524: XOR D
label_16525: XOR E
label_16526: XOR H
label_16527: XOR L
label_16528: XOR $12
label_16529: XOR (HL)
label_16530: XOR (IX + 127)
label_16531: XOR (IY - 128)
label_16532: CP A
label_16533: CP B
label_16534: CP C
label_16535: CP D
label_16536: CP E
label_16537: CP H
label_16538: CP L
label_16539: CP $12
label_16540: CP (HL)
label_16541: CP (IX + 127)
label_16542: CP (IY - 128)
label_16543: INC A
label_16544: INC B
label_16545: INC C
label_16546: INC D
label_16547: INC E
label_16548: INC H
label_16549: INC L
label_16550: INC (HL)
label_16551: INC (IX + 127)
label_16552: INC (IY - 128)
label_16553: DEC A
label_16554: DEC B
label_16555: DEC C
label_16556: DEC D
label_16557: DEC E
label_16558: DEC H
label_16559: DEC L
label_16560: DEC (HL)
label_16561: DEC (IX + 127)
label_16562: DEC (IY - 128)
label_16563: DAA
label_16564: CPL
label_16565: NEG
label_16566: CCF
label_16567: SCF
label_16568: NOP
label_16569: HALT
label_16570: DI
label_16571: EI
label_16572: IM 0
label_16573: IM 1
label_16574: IM 2
label_16575: ADD HL,BC
label_16576: ADD HL,DE
label_16577: ADD HL,HL
label_16578: ADD HL,SP
label_16579: ADC HL,BC
label_16580: ADC HL,DE
label_16581: ADC HL,HL
label_16582: ADC HL,SP
label_16583: SBC HL,BC
label_16584: SBC HL,DE
label_16585: SBC HL,HL
label_16586: SBC HL,SP
label_16587: ADD IX,BC
label_16588: ADD IX,DE
label_16589: ADD IX,SP
label_16590: ADD IY,BC
label_16591: ADD IY,DE
label_16592: ADD IY,SP
label_16593: INC BC
label_16594: INC DE
label_16595: INC HL
label_16596: INC SP
label_16597: INC IX
label_16598: INC IY
label_16599: DEC BC
label_16600: DEC DE
label_16601: DEC HL
label_16602: DEC SP
label_16603: DEC IX
label_16604: DEC IY
label_16605: RLCA
label_16606: RLA
label_16607: RRCA
label_16608: RRA
label_16609: RLC A
label_16610: RLC B
label_16611: RLC C
label_16612: RLC D
label_16613: RLC E
label_16614: RLC H
label_16615: RLC L
label_16616: RLC (HL)
label_16617: RLC (IX + 127)
label_16618: RLC (IY - 128)
label_16619: RL A
label_16620: RL B
label_16621: RL C
label_16622: RL D
label_16623: RL E
label_16624: RL H
label_16625: RL L
label_16626: RL (HL)
label_16627: RL (IX + 127)
label_16628: RL (IY - 128)
label_16629: RRC A
label_16630: RRC B
label_16631: RRC C
label_16632: RRC D
label_16633: RRC E
label_16634: RRC H
label_16635: RRC L
label_16636: RRC (HL)
label_16637: RRC (IX + 127)
label_16638: RRC (IY - 128)
label_16639: RR A
label_16640: RR B
label_16641: RR C
label_16642: RR D
label_16643: RR E
label_16644: RR H
label_16645: RR L
label_16646: RR (HL)
label_16647: RR (IX + 127)
label_16648: RR (IY - 128)
label_16649: SLA A
label_16650: SLA B
label_16651: SLA C
label_16652: SLA D
label_16653: SLA E
label_16654: SLA H
label_16655: SLA L
label_16656: SLA (HL)
label_16657: SLA (IX + 127)
label_16658: SLA (IY - 128)
label_16659: SRA A
label_16660: SRA B
label_16661: SRA C
label_16662: SRA D
label_16663: SRA E
label_16664: SRA H
label_16665: SRA L
label_16666: SRA (HL)
label_16667: SRA (IX + 127)
label_16668: SRA (IY - 128)
label_16669: SRL A
label_16670: SRL B
label_16671: SRL C
label_16672: SRL D
label_16673: SRL E
label_16674: SRL H
label_16675: SRL L
label_16676: SRL (HL)
label_16677: SRL (IX + 127)
label_16678: SRL (IY - 128)
label_16679: RLD
label_16680: RRD
label_16681: BIT 0,A
label_16682: BIT 1,A
label_16683: BIT 2,A
label_16684: BIT 3,A
label_16685: BIT 4,A
label_16686: BIT 5,A
label_16687: BIT 6,A
label_16688: BIT 7,A
label_16689: BIT 0,B
label_16690: BIT 1,B
label_16691: BIT 2,B
label_16692: BIT 3,B
label_16693: BIT 4,B
label_16694: BIT 5,B
label_16695: BIT 6,B
label_16696: BIT 7,B
label_16697: BIT 0,C
label_16698: BIT 1,C
label_16699: BIT 2,C
label_16700: BIT 3,C
label_16701: BIT 4,C
label_16702: BIT 5,C
label_16703: BIT 6,C
label_16704: BIT 7,C
label_16705: BIT 0,D
label_16706: BIT 1,D
label_16707: BIT 2,D
label_16708: BIT 3,D
label_16709: BIT 4,D
label_16710: BIT 5,D
label_16711: BIT 6,D
label_16712: BIT 7,D
label_16713: BIT 0,E
label_16714: BIT 1,E
label_16715: BIT 2,E
label_16716: BIT 3,E
label_16717: BIT 4,E
label_16718: BIT 5,E
label_16719: BIT 6,E
label_16720: BIT 7,E
label_16721: BIT 0,H
label_16722: BIT 1,H
label_16723: BIT 2,H
label_16724: BIT 3,H
label_16725: BIT 4,H
label_16726: BIT 5,H
label_16727: BIT 6,H
label_16728: BIT 7,H
label_16729: BIT 0,L
label_16730: BIT 1,L
label_16731: BIT 2,L
label_16732: BIT 3,L
label_16733: BIT 4,L
label_16734: BIT 5,L
label_16735: BIT 6,L
label_16736: BIT 7,L
label_16737: BIT 0,(HL)
label_16738: BIT 1,(HL)
label_16739: BIT 2,(HL)
label_16740: BIT 3,(HL)
label_16741: BIT 4,(HL)
label_16742: BIT 5,(HL)
label_16743: BIT 6,(HL)
label_16744: BIT 7,(HL)
label_16745: BIT 0,(IX + 127)
label_16746: BIT 1,(IX + 127)
label_16747: BIT 2,(IX + 127)
label_16748: BIT 3,(IX + 127)
label_16749: BIT 4,(IX + 127)
label_16750: BIT 5,(IX + 127)
label_16751: BIT 6,(IX + 127)
label_16752: BIT 7,(IX + 127)
label_16753: BIT 0,(IY - 128)
label_16754: BIT 1,(IY - 128)
label_16755: BIT 2,(IY - 128)
label_16756: BIT 3,(IY - 128)
label_16757: BIT 4,(IY - 128)
label_16758: BIT 5,(IY - 128)
label_16759: BIT 6,(IY - 128)
label_16760: BIT 7,(IY - 128)
label_16761: SET 0,A
label_16762: SET 1,A
label_16763: SET 2,A
label_16764: SET 3,A
label_16765: SET 4,A
label_16766: SET 5,A
label_16767: SET 6,A
label_16768: SET 7,A
label_16769: SET 0,B
label_16770: SET 1,B
label_16771: SET 2,B
label_16772: SET 3,B
label_16773: SET 4,B
label_16774: SET 5,B
label_16775: SET 6,B
label_16776: SET 7,B
label_16777: SET 0,C
label_16778: SET 1,C
label_16779: SET 2,C
label_16780: SET 3,C
label_16781: SET 4,C
label_16782: SET 5,C
label_16783: SET 6,C
label_16784: SET 7,C
label_16785: SET 0,D
label_16786: SET 1,D
label_16787: SET 2,D
label_16788: SET 3,D
label_16789: SET 4,D
label_16790: SET 5,D
label_16791: SET 6,D
label_16792: SET 7,D
label_16793: SET 0,E
label_16794: SET 1,E
label_16795: SET 2,E
label_16796: SET 3,E
label_16797: SET 4,E
label_16798: SET 5,E
label_16799: SET 6,E
label_16800: SET 7,E
label_16801: SET 0,H
label_16802: SET 1,H
label_16803: SET 2,H
label_16804: SET 3,H
label_16805: SET 4,H
label_16806: SET 5,H
label_16807: SET 6,H
label_16808: SET 7,H
label_16809: SET 0,L
label_16810: SET 1,L
label_16811: SET 2,L
label_16812: SET 3,L
label_16813: SET 4,L
label_16814: SET 5,L
label_16815: SET 6,L
label_16816: SET 7,L
label_16817: SET 0,(HL)
label_16818: SET 1,(HL)
label_16819: SET 2,(HL)
label_16820: SET 3,(HL)
label_16821: SET 4,(HL)
label_16822: SET 5,(HL)
label_16823: SET 6,(HL)
label_16824: SET 7,(HL)
label_16825: SET 0,(IX + 127)
label_16826: SET 1,(IX + 127)
label_16827: SET 2,(IX + 127)
label_16828: SET 3,(IX + 127)
label_16829: SET 4,(IX + 127)
label_16830: SET 5,(IX + 127)
label_16831: SET 6,(IX + 127)
label_16832: SET 7,(IX + 127)
label_16833: SET 0,(IY - 128)
label_16834: SET 1,(IY - 128)
label_16835: SET 2,(IY - 128)
label_16836: SET 3,(IY - 128)
label_16837: SET 4,(IY - 128)
label_16838: SET 5,(IY - 128)
label_16839: SET 6,(IY - 128)
label_16840: SET 7,(IY - 128)
label_16841: RES 0,A
label_16842: RES 1,A
label_16843: RES 2,A
label_16844: RES 3,A
label_16845: RES 4,A
label_16846: RES 5,A
label_16847: RES 6,A
label_16848: RES 7,A
label_16849: RES 0,B
label_16850: RES 1,B
label_16851: RES 2,B
label_16852: RES 3,B
label_16853: RES 4,B
label_16854: RES 5,B
label_16855: RES 6,B
label_16856: RES 7,B
label_16857: RES 0,C
label_16858: RES 1,C
label_16859: RES 2,C
label_16860: RES 3,C
label_16861: RES 4,C
label_16862: RES 5,C
label_16863: RES 6,C
label_16864: RES 7,C
label_16865: RES 0,D
label_16866: RES 1,D
label_16867: RES 2,D
label_16868: RES 3,D
label_16869: RES 4,D
label_16870: RES 5,D
label_16871: RES 6,D
label_16872: RES 7,D
label_16873: RES 0,E
label_16874: RES 1,E
label_16875: RES 2,E
label_16876: RES 3,E
label_16877: RES 4,E
label_16878: RES 5,E
label_16879: RES 6,E
label_16880: RES 7,E
label_16881: RES 0,H
label_16882: RES 1,H
label_16883: RES 2,H
label_16884: RES 3,H
label_16885: RES 4,H
label_16886: RES 5,H
label_16887: RES 6,H
label_16888: RES 7,H
label_16889: RES 0,L
label_16890: RES 1,L
label_16891: RES 2,L
label_16892: RES 3,L
label_16893: RES 4,L
label_16894: RES 5,L
label_16895: RES 6,L
label_16896: RES 7,L
label_16897: RES 0,(HL)
label_16898: RES 1,(HL)
label_16899: RES 2,(HL)
label_16900: RES 3,(HL)
label_16901: RES 4,(HL)
label_16902: RES 5,(HL)
label_16903: RES 6,(HL)
label_16904: RES 7,(HL)
label_16905: RES 0,(IX + 127)
label_16906: RES 1,(IX + 127)
label_16907: RES 2,(IX + 127)
label_16908: RES 3,(IX + 127)
label_16909: RES 4,(IX + 127)
label_16910: RES 5,(IX + 127)
label_16911: RES 6,(IX + 127)
label_16912: RES 7,(IX + 127)
label_16913: RES 0,(IY - 128)
label_16914: RES 1,(IY - 128)
label_16915: RES 2,(IY - 128)
label_16916: RES 3,(IY - 128)
label_16917: RES 4,(IY - 128)
label_16918: RES 5,(IY - 128)
label_16919: RES 6,(IY - 128)
label_16920: RES 7,(IY - 128)
label_16921: JP $5678
label_16922: JP NZ,$5678
label_16923: JP Z,$5678
label_16924: JP NC,$5678
label_16925: JP C,$5678
label_16926: JP PO,$5678
label_16927: JP PE,$5678
label_16928: JP P,$5678
label_16929: JP M,$5678
label_16930: JR $ + 2
label_16931: JR NZ,$ + 2
label_16932: JR Z,$ + 2
label_16933: JR NC,$ + 2
label_16934: JR C,$ + 2
label_16935: JP (HL)
label_16936: JP (IX)
label_16937: JP (IY)
label_16938: DJNZ $ + 2
label_16939: CALL $5678
label_16940: CALL NZ,$5678
label_16941: CALL Z,$5678
label_16942: CALL NC,$5678
label_16943: CALL C,$5678
label_16944: CALL PO,$5678
label_16945: CALL PE,$5678
label_16946: CALL P,$5678
label_16947: CALL M,$5678
label_16948: RET
label_16949: RET NZ
label_16950: RET Z
label_16951: RET NC
label_16952: RET C
label_16953: RET PO
label_16954: RET PE
label_16955: RET P
label_16956: RET M
label_16957: RETI
label_16958: RETN
label_16959: RST $00
label_16960: RST $08
label_16961: RST $10
label_16962: RST $18
label_16963: RST $20
label_16964: RST $28
label_16965: RST $30
label_16966: RST $38
label_16967: IN A,($12)
label_16968: IN A,(C)
label_16969: IN B,(C)
label_16970: IN C,(C)
label_16971: IN D,(C)
label_16972: IN E,(C)
label_16973: IN H,(C)
label_16974: IN L,(C)
label_16975: IN F,(C)
label_16976: INI
label_16977: INIR
label_16978: IND
label_16979: INDR
label_16980: OUT ($12),A
label_16981: OUT (C),A
label_16982: OUT (C),B
label_16983: OUT (C),C
label_16984: OUT (C),D
label_16985: OUT (C),E
label_16986: OUT (C),H
label_16987: OUT (C),L
label_16988: OUTI
label_16989: OTIR
label_16990: OUTD
label_16991: OTDR
label_16992: LD A,A
label_16993: LD A,B
label_16994: LD A,C
label_16995: LD A,D
label_16996: LD A,E
label_16997: LD A,H
label_16998: LD A,L
label_16999: LD B,A
label_17000: LD B,B
label_17001: LD B,C
label_17002: LD B,D
label_17003: LD B,E
label_17004: LD B,H
label_17005: LD B,L
label_17006: LD C,A
label_17007: LD C,B
label_17008: LD C,C
label_17009: LD C,D
label_17010: LD C,E
label_17011: LD C,H
label_17012: LD C,L
label_17013: LD D,A
label_17014: LD D,B
label_17015: LD D,C
label_17016: LD D,D
label_17017: LD D,E
label_17018: LD D,H
label_17019: LD D,L
label_17020: LD E,A
label_17021: LD E,B
label_17022: LD E,C
label_17023: LD E,D
label_17024: LD E,E
label_17025: LD E,H
label_17026: LD E,L
label_17027: LD H,A
label_17028: LD H,B
label_17029: LD H,C
label_17030: LD H,D
label_17031: LD H,E
label_17032: LD H,H
label_17033: LD H,L
label_17034: LD L,A
label_17035: LD L,B
label_17036: LD L,C
label_17037: LD L,D
label_17038: LD L,E
label_17039: LD L,H
label_17040: LD L,L
label_17041: LD A,$12
label_17042: LD B,$12
label_17043: LD C,$12
label_17044: LD D,$12
label_17045: LD E,$12
label_17046: LD H,$12
label_17047: LD L,$12
label_17048: LD A,(HL)
label_17049: LD B,(HL)
label_17050: LD C,(HL)
label_17051: LD D,(HL)
label_17052: LD E,(HL)
label_17053: LD H,(HL)
label_17054: LD L,(HL)
label_17055: LD A,(IX + 127)
label_17056: LD B,(IX + 127)
label_17057: LD C,(IX + 127)
label_17058: LD D,(IX + 127)
label_17059: LD E,(IX + 127)
label_17060: LD H,(IX + 127)
label_17061: LD L,(IX + 127)
label_17062: LD A,(IY - 128)
label_17063: LD B,(IY - 128)
label_17064: LD C,(IY - 128)
label_17065: LD D,(IY - 128)
label_17066: LD E,(IY - 128)
label_17067: LD H,(IY - 128)
label_17068: LD L,(IY - 128)
label_17069: LD (HL),A
label_17070: LD (HL),B
label_17071: LD (HL),C
label_17072: LD (HL),D
label_17073: LD (HL),E
label_17074: LD (HL),H
label_17075: LD (HL),L
label_17076: LD (IX + 127),A
label_17077: LD (IX + 127),B
label_17078: LD (IX + 127),C
label_17079: LD (IX + 127),D
label_17080: LD (IX + 127),E
label_17081: LD (IX + 127),H
label_17082: LD (IX + 127),L
label_17083: LD (IY - 128),A
label_17084: LD (IY - 128),B
label_17085: LD (IY - 128),C
label_17086: LD (IY - 128),D
label_17087: LD (IY - 128),E
label_17088: LD (IY - 128),H
label_17089: LD (IY - 128),L
label_17090: LD (HL),$12
label_17091: LD (IX + 127),$12
label_17092: LD (IY - 128),$12
label_17093: LD A,(BC)
label_17094: LD A,(DE)
label_17095: LD A,($5678)
label_17096: LD (BC),A
label_17097: LD (DE),A
label_17098: LD ($5678),A
label_17099: LD A,I
label_17100: LD A,R
label_17101: LD I,A
label_17102: LD R,A
label_17103: LD BC,$5678
label_17104: LD DE,$5678
label_17105: LD HL,$5678
label_17106: LD SP,$5678
label_17107: LD IX,$5678
label_17108: LD IY,$5678
label_17109: LD HL,($5678)
label_17110: LD BC,($5678)
label_17111: LD DE,($5678)
label_17112: LD HL,($5678)
label_17113: LD SP,($5678)
label_17114: LD IX,($5678)
label_17115: LD IY,($5678)
label_17116: LD ($5678),HL
label_17117: LD ($5678),BC
label_17118: LD ($5678),DE
label_17119: LD ($5678),HL
label_17120: LD ($5678),SP
label_17121: LD ($5678),IX
label_17122: LD ($5678),IY
label_17123: LD SP,HL
label_17124: LD SP,IX
label_17125: LD SP,IY
label_17126: PUSH BC
label_17127: PUSH DE
label_17128: PUSH HL
label_17129: PUSH AF
label_17130: PUSH IX
label_17131: PUSH IY
label_17132: POP BC
label_17133: POP DE
label_17134: POP HL
label_17135: POP AF
label_17136: POP IX
label_17137: POP IY
label_17138: EX DE,HL
label_17139: EX AF,AF'
label_17140: EXX
label_17141: EX (SP),HL
label_17142: EX (SP),IX
label_17143: EX (SP),IY
label_17144: LDI
label_17145: LDIR
label_17146: LDD
label_17147: LDDR
label_17148: CPI
label_17149: CPIR
label_17150: CPD
label_17151: CPDR
label_17152: ADD A,A
label_17153: ADD A,B
label_17154: ADD A,C
label_17155: ADD A,D
label_17156: ADD A,E
label_17157: ADD A,H
label_17158: ADD A,L
label_17159: ADD A,$12
label_17160: ADD A,(HL)
label_17161: ADD A,(IX + 127)
label_17162: ADD A,(IY - 128)
label_17163: ADC A,A
label_17164: ADC A,B
label_17165: ADC A,C
label_17166: ADC A,D
label_17167: ADC A,E
label_17168: ADC A,H
label_17169: ADC A,L
label_17170: ADC A,$12
label_17171: ADC A,(HL)
label_17172: ADC A,(IX + 127)
label_17173: ADC A,(IY - 128)
label_17174: SUB A
label_17175: SUB B
label_17176: SUB C
label_17177: SUB D
label_17178: SUB E
label_17179: SUB H
label_17180: SUB L
label_17181: SUB $12
label_17182: SUB (HL)
label_17183: SUB (IX + 127)
label_17184: SUB (IY - 128)
label_17185: SBC A,A
label_17186: SBC A,B
label_17187: SBC A,C
label_17188: SBC A,D
label_17189: SBC A,E
label_17190: SBC A,H
label_17191: SBC A,L
label_17192: SBC A,$12
label_17193: SBC A,(HL)
label_17194: SBC A,(IX + 127)
label_17195: SBC A,(IY - 128)
label_17196: AND A
label_17197: AND B
label_17198: AND C
label_17199: AND D
label_17200: AND E
label_17201: AND H
label_17202: AND L
label_17203: AND $12
label_17204: AND (HL)
label_17205: AND (IX + 127)
label_17206: AND (IY - 128)
label_17207: AND A
label_17208: AND B
label_17209: AND C
label_17210: AND D
label_17211: AND E
label_17212: AND H
label_17213: AND L
label_17214: AND $12
label_17215: AND (HL)
label_17216: AND (IX + 127)
label_17217: AND (IY - 128)
label_17218: OR A
label_17219: OR B
label_17220: OR C
label_17221: OR D
label_17222: OR E
label_17223: OR H
label_17224: OR L
label_17225: OR $12
label_17226: OR (HL)
label_17227: OR (IX + 127)
label_17228: OR (IY - 128)
label_17229: XOR A
label_17230: XOR B
label_17231: XOR C
label_17232: XOR D
label_17233: XOR E
label_17234: XOR H
label_17235: XOR L
label_17236: XOR $12
label_17237: XOR (HL)
label_17238: XOR (IX + 127)
label_17239: XOR (IY - 128)
label_17240: CP A
label_17241: CP B
label_17242: CP C
label_17243: CP D
label_17244: CP E
label_17245: CP H
label_17246: CP L
label_17247: CP $12
label_17248: CP (HL)
label_17249: CP (IX + 127)
label_17250: CP (IY - 128)
label_17251: INC A
label_17252: INC B
label_17253: INC C
label_17254: INC D
label_17255: INC E
label_17256: INC H
label_17257: INC L
label_17258: INC (HL)
label_17259: INC (IX + 127)
label_17260: INC (IY - 128)
label_17261: DEC A
label_17262: DEC B
label_17263: DEC C
label_17264: DEC D
label_17265: DEC E
label_17266: DEC H
label_17267: DEC L
label_17268: DEC (HL)
label_17269: DEC (IX + 127)
label_17270: DEC (IY - 128)
label_17271: DAA
label_17272: CPL
label_17273: NEG
label_17274: CCF
label_17275: SCF
label_17276: NOP
label_17277: HALT
label_17278: DI
label_17279: EI
label_17280: IM 0
label_17281: IM 1
label_17282: IM 2
label_17283: ADD HL,BC
label_17284: ADD HL,DE
label_17285: ADD HL,HL
label_17286: ADD HL,SP
label_17287: ADC HL,BC
label_17288: ADC HL,DE
label_17289: ADC HL,HL
label_17290: ADC HL,SP
label_17291: SBC HL,BC
label_17292: SBC HL,DE
label_17293: SBC HL,HL
label_17294: SBC HL,SP
label_17295: ADD IX,BC
label_17296: ADD IX,DE
label_17297: ADD IX,SP
label_17298: ADD IY,BC
label_17299: ADD IY,DE
label_17300: ADD IY,SP
label_17301: INC BC
label_17302: INC DE
label_17303: INC HL
label_17304: INC SP
label_17305: INC IX
label_17306: INC IY
label_17307: DEC BC
label_17308: DEC DE
label_17309: DEC HL
label_17310: DEC SP
label_17311: DEC IX
label_17312: DEC IY
label_17313: RLCA
label_17314: RLA
label_17315: RRCA
label_17316: RRA
label_17317: RLC A
label_17318: RLC B
label_17319: RLC C
label_17320: RLC D
label_17321: RLC E
label_17322: RLC H
label_17323: RLC L
label_17324: RLC (HL)
label_17325: RLC (IX + 127)
label_17326: RLC (IY - 128)
label_17327: RL A
label_17328: RL B
label_17329: RL C
label_17330: RL D
label_17331: RL E
label_17332: RL H
label_17333: RL L
label_17334: RL (HL)
label_17335: RL (IX + 127)
label_17336: RL (IY - 128)
label_17337: RRC A
label_17338: RRC B
label_17339: RRC C
label_17340: RRC D
label_17341: RRC E
label_17342: RRC H
label_17343: RRC L
label_17344: RRC (HL)
label_17345: RRC (IX + 127)
label_17346: RRC (IY - 128)
label_17347: RR A
label_17348: RR B
label_17349: RR C
label_17350: RR D
label_17351: RR E
label_17352: RR H
label_17353: RR L
label_17354: RR (HL)
label_17355: RR (IX + 127)
label_17356: RR (IY - 128)
label_17357: SLA A
label_17358: SLA B
label_17359: SLA C
label_17360: SLA D
label_17361: SLA E
label_17362: SLA H
label_17363: SLA L
label_17364: SLA (HL)
label_17365: SLA (IX + 127)
label_17366: SLA (IY - 128)
label_17367: SRA A
label_17368: SRA B
label_17369: SRA C
label_17370: SRA D
label_17371: SRA E
label_17372: SRA H
label_17373: SRA L
label_17374: SRA (HL)
label_17375: SRA (IX + 127)
label_17376: SRA (IY - 128)
label_17377: SRL A
label_17378: SRL B
label_17379: SRL C
label_17380: SRL D
label_17381: SRL E
label_17382: SRL H
label_17383: SRL L
label_17384: SRL (HL)
label_17385: SRL (IX + 127)
label_17386: SRL (IY - 128)
label_17387: RLD
label_17388: RRD
label_17389: BIT 0,A
label_17390: BIT 1,A
label_17391: BIT 2,A
label_17392: BIT 3,A
label_17393: BIT 4,A
label_17394: BIT 5,A
label_17395: BIT 6,A
label_17396: BIT 7,A
label_17397: BIT 0,B
label_17398: BIT 1,B
label_17399: BIT 2,B
label_17400: BIT 3,B
label_17401: BIT 4,B
label_17402: BIT 5,B
label_17403: BIT 6,B
label_17404: BIT 7,B
label_17405: BIT 0,C
label_17406: BIT 1,C
label_17407: BIT 2,C
label_17408: BIT 3,C
label_17409: BIT 4,C
label_17410: BIT 5,C
label_17411: BIT 6,C
label_17412: BIT 7,C
label_17413: BIT 0,D
label_17414: BIT 1,D
label_17415: BIT 2,D
label_17416: BIT 3,D
label_17417: BIT 4,D
label_17418: BIT 5,D
label_17419: BIT 6,D
label_17420: BIT 7,D
label_17421: BIT 0,E
label_17422: BIT 1,E
label_17423: BIT 2,E
label_17424: BIT 3,E
label_17425: BIT 4,E
label_17426: BIT 5,E
label_17427: BIT 6,E
label_17428: BIT 7,E
label_17429: BIT 0,H
label_17430: BIT 1,H
label_17431: BIT 2,H
label_17432: BIT 3,H
label_17433: BIT 4,H
label_17434: BIT 5,H
label_17435: BIT 6,H
label_17436: BIT 7,H
label_17437: BIT 0,L
label_17438: BIT 1,L
label_17439: BIT 2,L
label_17440: BIT 3,L
label_17441: BIT 4,L
label_17442: BIT 5,L
label_17443: BIT 6,L
label_17444: BIT 7,L
label_17445: BIT 0,(HL)
label_17446: BIT 1,(HL)
label_17447: BIT 2,(HL)
label_17448: BIT 3,(HL)
label_17449: BIT 4,(HL)
label_17450: BIT 5,(HL)
label_17451: BIT 6,(HL)
label_17452: BIT 7,(HL)
label_17453: BIT 0,(IX + 127)
label_17454: BIT 1,(IX + 127)
label_17455: BIT 2,(IX + 127)
label_17456: BIT 3,(IX + 127)
label_17457: BIT 4,(IX + 127)
label_17458: BIT 5,(IX + 127)
label_17459: BIT 6,(IX + 127)
label_17460: BIT 7,(IX + 127)
label_17461: BIT 0,(IY - 128)
label_17462: BIT 1,(IY - 128)
label_17463: BIT 2,(IY - 128)
label_17464: BIT 3,(IY - 128)
label_17465: BIT 4,(IY - 128)
label_17466: BIT 5,(IY - 128)
label_17467: BIT 6,(IY - 128)
label_17468: BIT 7,(IY - 128)
label_17469: SET 0,A
label_17470: SET 1,A
label_17471: SET 2,A
label_17472: SET 3,A
label_17473: SET 4,A
label_17474: SET 5,A
label_17475: SET 6,A
label_17476: SET 7,A
label_17477: SET 0,B
label_17478: SET 1,B
label_17479: SET 2,B
label_17480: SET 3,B
label_17481: SET 4,B
label_17482: SET 5,B
label_17483: SET 6,B
label_17484: SET 7,B
label_17485: SET 0,C
label_17486: SET 1,C
label_17487: SET 2,C
label_17488: SET 3,C
label_17489: SET 4,C
label_17490: SET 5,C
label_17491: SET 6,C
label_17492: SET 7,C
label_17493: SET 0,D
label_17494: SET 1,D
label_17495: SET 2,D
label_17496: SET 3,D
label_17497: SET 4,D
label_17498: SET 5,D
label_17499: SET 6,D
label_17500: SET 7,D
label_17501: SET 0,E
label_17502: SET 1,E
label_17503: SET 2,E
label_17504: SET 3,E
label_17505: SET 4,E
label_17506: SET 5,E
label_17507: SET 6,E
label_17508: SET 7,E
label_17509: SET 0,H
label_17510: SET 1,H
label_17511: SET 2,H
label_17512: SET 3,H
label_17513: SET 4,H
label_17514: SET 5,H
label_17515: SET 6,H
label_17516: SET 7,H
label_17517: SET 0,L
label_17518: SET 1,L
label_17519: SET 2,L
label_17520: SET 3,L
label_17521: SET 4,L
label_17522: SET 5,L
label_17523: SET 6,L
label_17524: SET 7,L
label_17525: SET 0,(HL)
label_17526: SET 1,(HL)
label_17527: SET 2,(HL)
label_17528: SET 3,(HL)
label_17529: SET 4,(HL)
label_17530: SET 5,(HL)
label_17531: SET 6,(HL)
label_17532: SET 7,(HL)
label_17533: SET 0,(IX + 127)
label_17534: SET 1,(IX + 127)
label_17535: SET 2,(IX + 127)
label_17536: SET 3,(IX + 127)
label_17537: SET 4,(IX + 127)
label_17538: SET 5,(IX + 127)
label_17539: SET 6,(IX + 127)
label_17540: SET 7,(IX + 127)
label_17541: SET 0,(IY - 128)
label_17542: SET 1,(IY - 128)
label_17543: SET 2,(IY - 128)
label_17544: SET 3,(IY - 128)
label_17545: SET 4,(IY - 128)
label_17546: SET 5,(IY - 128)
label_17547: SET 6,(IY - 128)
label_17548: SET 7,(IY - 128)
label_17549: RES 0,A
label_17550: RES 1,A
label_17551: RES 2,A
label_17552: RES 3,A
label_17553: RES 4,A
label_17554: RES 5,A
label_17555: RES 6,A
label_17556: RES 7,A
label_17557: RES 0,B
label_17558: RES 1,B
label_17559: RES 2,B
label_17560: RES 3,B
label_17561: RES 4,B
label_17562: RES 5,B
label_17563: RES 6,B
label_17564: RES 7,B
label_17565: RES 0,C
label_17566: RES 1,C
label_17567: RES 2,C
label_17568: RES 3,C
label_17569: RES 4,C
label_17570: RES 5,C
label_17571: RES 6,C
label_17572: RES 7,C
label_17573: RES 0,D
label_17574: RES 1,D
label_17575: RES 2,D
label_17576: RES 3,D
label_17577: RES 4,D
label_17578: RES 5,D
label_17579: RES 6,D
label_17580: RES 7,D
label_17581: RES 0,E
label_17582: RES 1,E
label_17583: RES 2,E
label_17584: RES 3,E
label_17585: RES 4,E
label_17586: RES 5,E
label_17587: RES 6,E
label_17588: RES 7,E
label_17589: RES 0,H
label_17590: RES 1,H
label_17591: RES 2,H
label_17592: RES 3,H
label_17593: RES 4,H
label_17594: RES 5,H
label_17595: RES 6,H
label_17596: RES 7,H
label_17597: RES 0,L
label_17598: RES 1,L
label_17599: RES 2,L
label_17600: RES 3,L
label_17601: RES 4,L
label_17602: RES 5,L
label_17603: RES 6,L
label_17604: RES 7,L
label_17605: RES 0,(HL)
label_17606: RES 1,(HL)
label_17607: RES 2,(HL)
label_17608: RES 3,(HL)
label_17609: RES 4,(HL)
label_17610: RES 5,(HL)
label_17611: RES 6,(HL)
label_17612: RES 7,(HL)
label_17613: RES 0,(IX + 127)
label_17614: RES 1,(IX + 127)
label_17615: RES 2,(IX + 127)
label_17616: RES 3,(IX + 127)
label_17617: RES 4,(IX + 127)
label_17618: RES 5,(IX + 127)
label_17619: RES 6,(IX + 127)
label_17620: RES 7,(IX + 127)
label_17621: RES 0,(IY - 128)
label_17622: RES 1,(IY - 128)
label_17623: RES 2,(IY - 128)
label_17624: RES 3,(IY - 128)
label_17625: RES 4,(IY - 128)
label_17626: RES 5,(IY - 128)
label_17627: RES 6,(IY - 128)
label_17628: RES 7,(IY - 128)
label_17629: JP $5678
label_17630: JP NZ,$5678
label_17631: JP Z,$5678
label_17632: JP NC,$5678
label_17633: JP C,$5678
label_17634: JP PO,$5678
label_17635: JP PE,$5678
label_17636: JP P,$5678
label_17637: JP M,$5678
label_17638: JR $ + 2
label_17639: JR NZ,$ + 2
label_17640: JR Z,$ + 2
label_17641: JR NC,$ + 2
label_17642: JR C,$ + 2
label_17643: JP (HL)
label_17644: JP (IX)
label_17645: JP (IY)
label_17646: DJNZ $ + 2
label_17647: CALL $5678
label_17648: CALL NZ,$5678
label_17649: CALL Z,$5678
label_17650: CALL NC,$5678
label_17651: CALL C,$5678
label_17652: CALL PO,$5678
label_17653: CALL PE,$5678
label_17654: CALL P,$5678
label_17655: CALL M,$5678
label_17656: RET
label_17657: RET NZ
label_17658: RET Z
label_17659: RET NC
label_17660: RET C
label_17661: RET PO
label_17662: RET PE
label_17663: RET P
label_17664: RET M
label_17665: RETI
label_17666: RETN
label_17667: RST $00
label_17668: RST $08
label_17669: RST $10
label_17670: RST $18
label_17671: RST $20
label_17672: RST $28
label_17673: RST $30
label_17674: RST $38
label_17675: IN A,($12)
label_17676: IN A,(C)
label_17677: IN B,(C)
label_17678: IN C,(C)
label_17679: IN D,(C)
label_17680: IN E,(C)
label_17681: IN H,(C)
label_17682: IN L,(C)
label_17683: IN F,(C)
label_17684: INI
label_17685: INIR
label_17686: IND
label_17687: INDR
label_17688: OUT ($12),A
label_17689: OUT (C),A
label_17690: OUT (C),B
label_17691: OUT (C),C
label_17692: OUT (C),D
label_17693: OUT (C),E
label_17694: OUT (C),H
label_17695: OUT (C),L
label_17696: OUTI
label_17697: OTIR
label_17698: OUTD
label_17699: OTDR
label_17700: LD A,A
label_17701: LD A,B
label_17702: LD A,C
label_17703: LD A,D
label_17704: LD A,E
label_17705: LD A,H
label_17706: LD A,L
label_17707: LD B,A
label_17708: LD B,B
label_17709: LD B,C
label_17710: LD B,D
label_17711: LD B,E
label_17712: LD B,H
label_17713: LD B,L
label_17714: LD C,A
label_17715: LD C,B
label_17716: LD C,C
label_17717: LD C,D
label_17718: LD C,E
label_17719: LD C,H
label_17720: LD C,L
label_17721: LD D,A
label_17722: LD D,B
label_17723: LD D,C
label_17724: LD D,D
label_17725: LD D,E
label_17726: LD D,H
label_17727: LD D,L
label_17728: LD E,A
label_17729: LD E,B
label_17730: LD E,C
label_17731: LD E,D
label_17732: LD E,E
label_17733: LD E,H
label_17734: LD E,L
label_17735: LD H,A
label_17736: LD H,B
label_17737: LD H,C
label_17738: LD H,D
label_17739: LD H,E
label_17740: LD H,H
label_17741: LD H,L
label_17742: LD L,A
label_17743: LD L,B
label_17744: LD L,C
label_17745: LD L,D
label_17746: LD L,E
label_17747: LD L,H
label_17748: LD L,L
label_17749: LD A,$12
label_17750: LD B,$12
label_17751: LD C,$12
label_17752: LD D,$12
label_17753: LD E,$12
label_17754: LD H,$12
label_17755: LD L,$12
label_17756: LD A,(HL)
label_17757: LD B,(HL)
label_17758: LD C,(HL)
label_17759: LD D,(HL)
label_17760: LD E,(HL)
label_17761: LD H,(HL)
label_17762: LD L,(HL)
label_17763: LD A,(IX + 127)
label_17764: LD B,(IX + 127)
label_17765: LD C,(IX + 127)
label_17766: LD D,(IX + 127)
label_17767: LD E,(IX + 127)
label_17768: LD H,(IX + 127)
label_17769: LD L,(IX + 127)
label_17770: LD A,(IY - 128)
label_17771: LD B,(IY - 128)
label_17772: LD C,(IY - 128)
label_17773: LD D,(IY - 128)
label_17774: LD E,(IY - 128)
label_17775: LD H,(IY - 128)
label_17776: LD L,(IY - 128)
label_17777: LD (HL),A
label_17778: LD (HL),B
label_17779: LD (HL),C
label_17780: LD (HL),D
label_17781: LD (HL),E
label_17782: LD (HL),H
label_17783: LD (HL),L
label_17784: LD (IX + 127),A
label_17785: LD (IX + 127),B
label_17786: LD (IX + 127),C
label_17787: LD (IX + 127),D
label_17788: LD (IX + 127),E
label_17789: LD (IX + 127),H
label_17790: LD (IX + 127),L
label_17791: LD (IY - 128),A
label_17792: LD (IY - 128),B
label_17793: LD (IY - 128),C
label_17794: LD (IY - 128),D
label_17795: LD (IY - 128),E
label_17796: LD (IY - 128),H
label_17797: LD (IY - 128),L
label_17798: LD (HL),$12
label_17799: LD (IX + 127),$12
label_17800: LD (IY - 128),$12
label_17801: LD A,(BC)
label_17802: LD A,(DE)
label_17803: LD A,($5678)
label_17804: LD (BC),A
label_17805: LD (DE),A
label_17806: LD ($5678),A
label_17807: LD A,I
label_17808: LD A,R
label_17809: LD I,A
label_17810: LD R,A
label_17811: LD BC,$5678
label_17812: LD DE,$5678
label_17813: LD HL,$5678
label_17814: LD SP,$5678
label_17815: LD IX,$5678
label_17816: LD IY,$5678
label_17817: LD HL,($5678)
label_17818: LD BC,($5678)
label_17819: LD DE,($5678)
label_17820: LD HL,($5678)
label_17821: LD SP,($5678)
label_17822: LD IX,($5678)
label_17823: LD IY,($5678)
label_17824: LD ($5678),HL
label_17825: LD ($5678),BC
label_17826: LD ($5678),DE
label_17827: LD ($5678),HL
label_17828: LD ($5678),SP
label_17829: LD ($5678),IX
label_17830: LD ($5678),IY
label_17831: LD SP,HL
label_17832: LD SP,IX
label_17833: LD SP,IY
label_17834: PUSH BC
label_17835: PUSH DE
label_17836: PUSH HL
label_17837: PUSH AF
label_17838: PUSH IX
label_17839: PUSH IY
label_17840: POP BC
label_17841: POP DE
label_17842: POP HL
label_17843: POP AF
label_17844: POP IX
label_17845: POP IY
label_17846: EX DE,HL
label_17847: EX AF,AF'
label_17848: EXX
label_17849: EX (SP),HL
label_17850: EX (SP),IX
label_17851: EX (SP),IY
label_17852: LDI
label_17853: LDIR
label_17854: LDD
label_17855: LDDR
label_17856: CPI
label_17857: CPIR
label_17858: CPD
label_17859: CPDR
label_17860: ADD A,A
label_17861: ADD A,B
label_17862: ADD A,C
label_17863: ADD A,D
label_17864: ADD A,E
label_17865: ADD A,H
label_17866: ADD A,L
label_17867: ADD A,$12
label_17868: ADD A,(HL)
label_17869: ADD A,(IX + 127)
label_17870: ADD A,(IY - 128)
label_17871: ADC A,A
label_17872: ADC A,B
label_17873: ADC A,C
label_17874: ADC A,D
label_17875: ADC A,E
label_17876: ADC A,H
label_17877: ADC A,L
label_17878: ADC A,$12
label_17879: ADC A,(HL)
label_17880: ADC A,(IX + 127)
label_17881: ADC A,(IY - 128)
label_17882: SUB A
label_17883: SUB B
label_17884: SUB C
label_17885: SUB D
label_17886: SUB E
label_17887: SUB H
label_17888: SUB L
label_17889: SUB $12
label_17890: SUB (HL)
label_17891: SUB (IX + 127)
label_17892: SUB (IY - 128)
label_17893: SBC A,A
label_17894: SBC A,B
label_17895: SBC A,C
label_17896: SBC A,D
label_17897: SBC A,E
label_17898: SBC A,H
label_17899: SBC A,L
label_17900: SBC A,$12
label_17901: SBC A,(HL)
label_17902: SBC A,(IX + 127)
label_17903: SBC A,(IY - 128)
label_17904: AND A
label_17905: AND B
label_17906: AND C
label_17907: AND D
label_17908: AND E
label_17909: AND H
label_17910: AND L
label_17911: AND $12
label_17912: AND (HL)
label_17913: AND (IX + 127)
label_17914: AND (IY - 128)
label_17915: AND A
label_17916: AND B
label_17917: AND C
label_17918: AND D
label_17919: AND E
label_17920: AND H
label_17921: AND L
label_17922: AND $12
label_17923: AND (HL)
label_17924: AND (IX + 127)
label_17925: AND (IY - 128)
label_17926: OR A
label_17927: OR B
label_17928: OR C
label_17929: OR D
label_17930: OR E
label_17931: OR H
label_17932: OR L
label_17933: OR $12
label_17934: OR (HL)
label_17935: OR (IX + 127)
label_17936: OR (IY - 128)
label_17937: XOR A
label_17938: XOR B
label_17939: XOR C
label_17940: XOR D
label_17941: XOR E
label_17942: XOR H
label_17943: XOR L
label_17944: XOR $12
label_17945: XOR (HL)
label_17946: XOR (IX + 127)
label_17947: XOR (IY - 128)
label_17948: CP A
label_17949: CP B
label_17950: CP C
label_17951: CP D
label_17952: CP E
label_17953: CP H
label_17954: CP L
label_17955: CP $12
label_17956: CP (HL)
label_17957: CP (IX + 127)
label_17958: CP (IY - 128)
label_17959: INC A
label_17960: INC B
label_17961: INC C
label_17962: INC D
label_17963: INC E
label_17964: INC H
label_17965: INC L
label_17966: INC (HL)
label_17967: INC (IX + 127)
label_17968: INC (IY - 128)
label_17969: DEC A
label_17970: DEC B
label_17971: DEC C
label_17972: DEC D
label_17973: DEC E
label_17974: DEC H
label_17975: DEC L
label_17976: DEC (HL)
label_17977: DEC (IX + 127)
label_17978: DEC (IY - 128)
label_17979: DAA
label_17980: CPL
label_17981: NEG
label_17982: CCF
label_17983: SCF
label_17984: NOP
label_17985: HALT
label_17986: DI
label_17987: EI
label_17988: IM 0
label_17989: IM 1
label_17990: IM 2
label_17991: ADD HL,BC
label_17992: ADD HL,DE
label_17993: ADD HL,HL
label_17994: ADD HL,SP
label_17995: ADC HL,BC
label_17996: ADC HL,DE
label_17997: ADC HL,HL
label_17998: ADC HL,SP
label_17999: SBC HL,BC
label_18000: SBC HL,DE
label_18001: SBC HL,HL
label_18002: SBC HL,SP
label_18003: ADD IX,BC
label_18004: ADD IX,DE
label_18005: ADD IX,SP
label_18006: ADD IY,BC
label_18007: ADD IY,DE
label_18008: ADD IY,SP
label_18009: INC BC
label_18010: INC DE
label_18011: INC HL
label_18012: INC SP
label_18013: INC IX
label_18014: INC IY
label_18015: DEC BC
label_18016: DEC DE
label_18017: DEC HL
label_18018: DEC SP
label_18019: DEC IX
label_18020: DEC IY
label_18021: RLCA
label_18022: RLA
label_18023: RRCA
label_18024: RRA
label_18025: RLC A
label_18026: RLC B
label_18027: RLC C
label_18028: RLC D
label_18029: RLC E
label_18030: RLC H
label_18031: RLC L
label_18032: RLC (HL)
label_18033: RLC (IX + 127)
label_18034: RLC (IY - 128)
label_18035: RL A
label_18036: RL B
label_18037: RL C
label_18038: RL D
label_18039: RL E
label_18040: RL H
label_18041: RL L
label_18042: RL (HL)
label_18043: RL (IX + 127)
label_18044: RL (IY - 128)
label_18045: RRC A
label_18046: RRC B
label_18047: RRC C
label_18048: RRC D
label_18049: RRC E
label_18050: RRC H
label_18051: RRC L
label_18052: RRC (HL)
label_18053: RRC (IX + 127)
label_18054: RRC (IY - 128)
label_18055: RR A
label_18056: RR B
label_18057: RR C
label_18058: RR D
label_18059: RR E
label_18060: RR H
label_18061: RR L
label_18062: RR (HL)
label_18063: RR (IX + 127)
label_18064: RR (IY - 128)
label_18065: SLA A
label_18066: SLA B
label_18067: SLA C
label_18068: SLA D
label_18069: SLA E
label_18070: SLA H
label_18071: SLA L
label_18072: SLA (HL)
label_18073: SLA (IX + 127)
label_18074: SLA (IY - 128)
label_18075: SRA A
label_18076: SRA B
label_18077: SRA C
label_18078: SRA D
label_18079: SRA E
label_18080: SRA H
label_18081: SRA L
label_18082: SRA (HL)
label_18083: SRA (IX + 127)
label_18084: SRA (IY - 128)
label_18085: SRL A
label_18086: SRL B
label_18087: SRL C
label_18088: SRL D
label_18089: SRL E
label_18090: SRL H
label_18091: SRL L
label_18092: SRL (HL)
label_18093: SRL (IX + 127)
label_18094: SRL (IY - 128)
label_18095: RLD
label_18096: RRD
label_18097: BIT 0,A
label_18098: BIT 1,A
label_18099: BIT 2,A
label_18100: BIT 3,A
label_18101: BIT 4,A
label_18102: BIT 5,A
label_18103: BIT 6,A
label_18104: BIT 7,A
label_18105: BIT 0,B
label_18106: BIT 1,B
label_18107: BIT 2,B
label_18108: BIT 3,B
label_18109: BIT 4,B
label_18110: BIT 5,B
label_18111: BIT 6,B
label_18112: BIT 7,B
label_18113: BIT 0,C
label_18114: BIT 1,C
label_18115: BIT 2,C
label_18116: BIT 3,C
label_18117: BIT 4,C
label_18118: BIT 5,C
label_18119: BIT 6,C
label_18120: BIT 7,C
label_18121: BIT 0,D
label_18122: BIT 1,D
label_18123: BIT 2,D
label_18124: BIT 3,D
label_18125: BIT 4,D
label_18126: BIT 5,D
label_18127: BIT 6,D
label_18128: BIT 7,D
label_18129: BIT 0,E
label_18130: BIT 1,E
label_18131: BIT 2,E
label_18132: BIT 3,E
label_18133: BIT 4,E
label_18134: BIT 5,E
label_18135: BIT 6,E
label_18136: BIT 7,E
label_18137: BIT 0,H
label_18138: BIT 1,H
label_18139: BIT 2,H
label_18140: BIT 3,H
label_18141: BIT 4,H
label_18142: BIT 5,H
label_18143: BIT 6,H
label_18144: BIT 7,H
label_18145: BIT 0,L
label_18146: BIT 1,L
label_18147: BIT 2,L
label_18148: BIT 3,L
label_18149: BIT 4,L
label_18150: BIT 5,L
label_18151: BIT 6,L
label_18152: BIT 7,L
label_18153: BIT 0,(HL)
label_18154: BIT 1,(HL)
label_18155: BIT 2,(HL)
label_18156: BIT 3,(HL)
label_18157: BIT 4,(HL)
label_18158: BIT 5,(HL)
label_18159: BIT 6,(HL)
label_18160: BIT 7,(HL)
label_18161: BIT 0,(IX + 127)
label_18162: BIT 1,(IX + 127)
label_18163: BIT 2,(IX + 127)
label_18164: BIT 3,(IX + 127)
label_18165: BIT 4,(IX + 127)
label_18166: BIT 5,(IX + 127)
label_18167: BIT 6,(IX + 127)
label_18168: BIT 7,(IX + 127)
label_18169: BIT 0,(IY - 128)
label_18170: BIT 1,(IY - 128)
label_18171: BIT 2,(IY - 128)
label_18172: BIT 3,(IY - 128)
label_18173: BIT 4,(IY - 128)
label_18174: BIT 5,(IY - 128)
label_18175: BIT 6,(IY - 128)
label_18176: BIT 7,(IY - 128)
label_18177: SET 0,A
label_18178: SET 1,A
label_18179: SET 2,A
label_18180: SET 3,A
label_18181: SET 4,A
label_18182: SET 5,A
label_18183: SET 6,A
label_18184: SET 7,A
label_18185: SET 0,B
label_18186: SET 1,B
label_18187: SET 2,B
label_18188: SET 3,B
label_18189: SET 4,B
label_18190: SET 5,B
label_18191: SET 6,B
label_18192: SET 7,B
label_18193: SET 0,C
label_18194: SET 1,C
label_18195: SET 2,C
label_18196: SET 3,C
label_18197: SET 4,C
label_18198: SET 5,C
label_18199: SET 6,C
label_18200: SET 7,C
label_18201: SET 0,D
label_18202: SET 1,D
label_18203: SET 2,D
label_18204: SET 3,D
label_18205: SET 4,D
label_18206: SET 5,D
label_18207: SET 6,D
label_18208: SET 7,D
label_18209: SET 0,E
label_18210: SET 1,E
label_18211: SET 2,E
label_18212: SET 3,E
label_18213: SET 4,E
label_18214: SET 5,E
label_18215: SET 6,E
label_18216: SET 7,E
label_18217: SET 0,H
label_18218: SET 1,H
label_18219: SET 2,H
label_18220: SET 3,H
label_18221: SET 4,H
label_18222: SET 5,H
label_18223: SET 6,H
label_18224: SET 7,H
label_18225: SET 0,L
label_18226: SET 1,L
label_18227: SET 2,L
label_18228: SET 3,L
label_18229: SET 4,L
label_18230: SET 5,L
label_18231: SET 6,L
label_18232: SET 7,L
label_18233: SET 0,(HL)
label_18234: SET 1,(HL)
label_18235: SET 2,(HL)
label_18236: SET 3,(HL)
label_18237: SET 4,(HL)
label_18238: SET 5,(HL)
label_18239: SET 6,(HL)
label_18240: SET 7,(HL)
label_18241: SET 0,(IX + 127)
label_18242: SET 1,(IX + 127)
label_18243: SET 2,(IX + 127)
label_18244: SET 3,(IX + 127)
label_18245: SET 4,(IX + 127)
label_18246: SET 5,(IX + 127)
label_18247: SET 6,(IX + 127)
label_18248: SET 7,(IX + 127)
label_18249: SET 0,(IY - 128)
label_18250: SET 1,(IY - 128)
label_18251: SET 2,(IY - 128)
label_18252: SET 3,(IY - 128)
label_18253: SET 4,(IY - 128)
label_18254: SET 5,(IY - 128)
label_18255: SET 6,(IY - 128)
label_18256: SET 7,(IY - 128)
label_18257: RES 0,A
label_18258: RES 1,A
label_18259: RES 2,A
label_18260: RES 3,A
label_18261: RES 4,A
label_18262: RES 5,A
label_18263: RES 6,A
label_18264: RES 7,A
label_18265: RES 0,B
label_18266: RES 1,B
label_18267: RES 2,B
label_18268: RES 3,B
label_18269: RES 4,B
label_18270: RES 5,B
label_18271: RES 6,B
label_18272: RES 7,B
label_18273: RES 0,C
label_18274: RES 1,C
label_18275: RES 2,C
label_18276: RES 3,C
label_18277: RES 4,C
label_18278: RES 5,C
label_18279: RES 6,C
label_18280: RES 7,C
label_18281: RES 0,D
label_18282: RES 1,D
label_18283: RES 2,D
label_18284: RES 3,D
label_18285: RES 4,D
label_18286: RES 5,D
label_18287: RES 6,D
label_18288: RES 7,D
label_18289: RES 0,E
label_18290: RES 1,E
label_18291: RES 2,E
label_18292: RES 3,E
label_18293: RES 4,E
label_18294: RES 5,E
label_18295: RES 6,E
label_18296: RES 7,E
label_18297: RES 0,H
label_18298: RES 1,H
label_18299: RES 2,H
label_18300: RES 3,H
label_18301: RES 4,H
label_18302: RES 5,H
label_18303: RES 6,H
label_18304: RES 7,H
label_18305: RES 0,L
label_18306: RES 1,L
label_18307: RES 2,L
label_18308: RES 3,L
label_18309: RES 4,L
label_18310: RES 5,L
label_18311: RES 6,L
label_18312: RES 7,L
label_18313: RES 0,(HL)
label_18314: RES 1,(HL)
label_18315: RES 2,(HL)
label_18316: RES 3,(HL)
label_18317: RES 4,(HL)
label_18318: RES 5,(HL)
label_18319: RES 6,(HL)
label_18320: RES 7,(HL)
label_18321: RES 0,(IX + 127)
label_18322: RES 1,(IX + 127)
label_18323: RES 2,(IX + 127)
label_18324: RES 3,(IX + 127)
label_18325: RES 4,(IX + 127)
label_18326: RES 5,(IX + 127)
label_18327: RES 6,(IX + 127)
label_18328: RES 7,(IX + 127)
label_18329: RES 0,(IY - 128)
label_18330: RES 1,(IY - 128)
label_18331: RES 2,(IY - 128)
label_18332: RES 3,(IY - 128)
label_18333: RES 4,(IY - 128)
label_18334: RES 5,(IY - 128)
label_18335: RES 6,(IY - 128)
label_18336: RES 7,(IY - 128)
label_18337: JP $5678
label_18338: JP NZ,$5678
label_18339: JP Z,$5678
label_18340: JP NC,$5678
label_18341: JP C,$5678
label_18342: JP PO,$5678
label_18343: JP PE,$5678
label_18344: JP P,$5678
label_18345: JP M,$5678
label_18346: JR $ + 2
label_18347: JR NZ,$ + 2
label_18348: JR Z,$ + 2
label_18349: JR NC,$ + 2
label_18350: JR C,$ + 2
label_18351: JP (HL)
label_18352: JP (IX)
label_18353: JP (IY)
label_18354: DJNZ $ + 2
label_18355: CALL $5678
label_18356: CALL NZ,$5678
label_18357: CALL Z,$5678
label_18358: CALL NC,$5678
label_18359: CALL C,$5678
label_18360: CALL PO,$5678
label_18361: CALL PE,$5678
label_18362: CALL P,$5678
label_18363: CALL M,$5678
label_18364: RET
label_18365: RET NZ
label_18366: RET Z
label_18367: RET NC
label_18368: RET C
label_18369: RET PO
label_18370: RET PE
label_18371: RET P
label_18372: RET M
label_18373: RETI
label_18374: RETN
label_18375: RST $00
label_18376: RST $08
label_18377: RST $10
label_18378: RST $18
label_18379: RST $20
label_18380: RST $28
label_18381: RST $30
label_18382: RST $38
label_18383: IN A,($12)
label_18384: IN A,(C)
label_18385: IN B,(C)
label_18386: IN C,(C)
label_18387: IN D,(C)
label_18388: IN E,(C)
label_18389: IN H,(C)
label_18390: IN L,(C)
label_18391: IN F,(C)
label_18392: INI
label_18393: INIR
label_18394: IND
label_18395: INDR
label_18396: OUT ($12),A
label_18397: OUT (C),A
label_18398: OUT (C),B
label_18399: OUT (C),C
label_18400: OUT (C),D
label_18401: OUT (C),E
label_18402: OUT (C),H
label_18403: OUT (C),L
label_18404: OUTI
label_18405: OTIR
label_18406: OUTD
label_18407: OTDR
label_18408: LD A,A
label_18409: LD A,B
label_18410: LD A,C
label_18411: LD A,D
label_18412: LD A,E
label_18413: LD A,H
label_18414: LD A,L
label_18415: LD B,A
label_18416: LD B,B
label_18417: LD B,C
label_18418: LD B,D
label_18419: LD B,E
label_18420: LD B,H
label_18421: LD B,L
label_18422: LD C,A
label_18423: LD C,B
label_18424: LD C,C
label_18425: LD C,D
label_18426: LD C,E
label_18427: LD C,H
label_18428: LD C,L
label_18429: LD D,A
label_18430: LD D,B
label_18431: LD D,C
label_18432: LD D,D
label_18433: LD D,E
label_18434: LD D,H
label_18435: LD D,L
label_18436: LD E,A
label_18437: LD E,B
label_18438: LD E,C
label_18439: LD E,D
label_18440: LD E,E
label_18441: LD E,H
label_18442: LD E,L
label_18443: LD H,A
label_18444: LD H,B
label_18445: LD H,C
label_18446: LD H,D
label_18447: LD H,E
label_18448: LD H,H
label_18449: LD H,L
label_18450: LD L,A
label_18451: LD L,B
label_18452: LD L,C
label_18453: LD L,D
label_18454: LD L,E
label_18455: LD L,H
label_18456: LD L,L
label_18457: LD A,$12
label_18458: LD B,$12
label_18459: LD C,$12
label_18460: LD D,$12
label_18461: LD E,$12
label_18462: LD H,$12
label_18463: LD L,$12
label_18464: LD A,(HL)
label_18465: LD B,(HL)
label_18466: LD C,(HL)
label_18467: LD D,(HL)
label_18468: LD E,(HL)
label_18469: LD H,(HL)
label_18470: LD L,(HL)
label_18471: LD A,(IX + 127)
label_18472: LD B,(IX + 127)
label_18473: LD C,(IX + 127)
label_18474: LD D,(IX + 127)
label_18475: LD E,(IX + 127)
label_18476: LD H,(IX + 127)
label_18477: LD L,(IX + 127)
label_18478: LD A,(IY - 128)
label_18479: LD B,(IY - 128)
label_18480: LD C,(IY - 128)
label_18481: LD D,(IY - 128)
label_18482: LD E,(IY - 128)
label_18483: LD H,(IY - 128)
label_18484: LD L,(IY - 128)
label_18485: LD (HL),A
label_18486: LD (HL),B
label_18487: LD (HL),C
label_18488: LD (HL),D
label_18489: LD (HL),E
label_18490: LD (HL),H
label_18491: LD (HL),L
label_18492: LD (IX + 127),A
label_18493: LD (IX + 127),B
label_18494: LD (IX + 127),C
label_18495: LD (IX + 127),D
label_18496: LD (IX + 127),E
label_18497: LD (IX + 127),H
label_18498: LD (IX + 127),L
label_18499: LD (IY - 128),A
label_18500: LD (IY - 128),B
label_18501: LD (IY - 128),C
label_18502: LD (IY - 128),D
label_18503: LD (IY - 128),E
label_18504: LD (IY - 128),H
label_18505: LD (IY - 128),L
label_18506: LD (HL),$12
label_18507: LD (IX + 127),$12
label_18508: LD (IY - 128),$12
label_18509: LD A,(BC)
label_18510: LD A,(DE)
label_18511: LD A,($5678)
label_18512: LD (BC),A
label_18513: LD (DE),A
label_18514: LD ($5678),A
label_18515: LD A,I
label_18516: LD A,R
label_18517: LD I,A
label_18518: LD R,A
label_18519: LD BC,$5678
label_18520: LD DE,$5678
label_18521: LD HL,$5678
label_18522: LD SP,$5678
label_18523: LD IX,$5678
label_18524: LD IY,$5678
label_18525: LD HL,($5678)
label_18526: LD BC,($5678)
label_18527: LD DE,($5678)
label_18528: LD HL,($5678)
label_18529: LD SP,($5678)
label_18530: LD IX,($5678)
label_18531: LD IY,($5678)
label_18532: LD ($5678),HL
label_18533: LD ($5678),BC
label_18534: LD ($5678),DE
label_18535: LD ($5678),HL
label_18536: LD ($5678),SP
label_18537: LD ($5678),IX
label_18538: LD ($5678),IY
label_18539: LD SP,HL
label_18540: LD SP,IX
label_18541: LD SP,IY
label_18542: PUSH BC
label_18543: PUSH DE
label_18544: PUSH HL
label_18545: PUSH AF
label_18546: PUSH IX
label_18547: PUSH IY
label_18548: POP BC
label_18549: POP DE
label_18550: POP HL
label_18551: POP AF
label_18552: POP IX
label_18553: POP IY
label_18554: EX DE,HL
label_18555: EX AF,AF'
label_18556: EXX
label_18557: EX (SP),HL
label_18558: EX (SP),IX
label_18559: EX (SP),IY
label_18560: LDI
label_18561: LDIR
label_18562: LDD
label_18563: LDDR
label_18564: CPI
label_18565: CPIR
label_18566: CPD
label_18567: CPDR
label_18568: ADD A,A
label_18569: ADD A,B
label_18570: ADD A,C
label_18571: ADD A,D
label_18572: ADD A,E
label_18573: ADD A,H
label_18574: ADD A,L
label_18575: ADD A,$12
label_18576: ADD A,(HL)
label_18577: ADD A,(IX + 127)
label_18578: ADD A,(IY - 128)
label_18579: ADC A,A
label_18580: ADC A,B
label_18581: ADC A,C
label_18582: ADC A,D
label_18583: ADC A,E
label_18584: ADC A,H
label_18585: ADC A,L
label_18586: ADC A,$12
label_18587: ADC A,(HL)
label_18588: ADC A,(IX + 127)
label_18589: ADC A,(IY - 128)
label_18590: SUB A
label_18591: SUB B
label_18592: SUB C
label_18593: SUB D
label_18594: SUB E
label_18595: SUB H
label_18596: SUB L
label_18597: SUB $12
label_18598: SUB (HL)
label_18599: SUB (IX + 127)
label_18600: SUB (IY - 128)
label_18601: SBC A,A
label_18602: SBC A,B
label_18603: SBC A,C
label_18604: SBC A,D
label_18605: SBC A,E
label_18606: SBC A,H
label_18607: SBC A,L
label_18608: SBC A,$12
label_18609: SBC A,(HL)
label_18610: SBC A,(IX + 127)
label_18611: SBC A,(IY - 128)
label_18612: AND A
label_18613: AND B
label_18614: AND C
label_18615: AND D
label_18616: AND E
label_18617: AND H
label_18618: AND L
label_18619: AND $12
label_18620: AND (HL)
label_18621: AND (IX + 127)
label_18622: AND (IY - 128)
label_18623: AND A
label_18624: AND B
label_18625: AND C
label_18626: AND D
label_18627: AND E
label_18628: AND H
label_18629: AND L
label_18630: AND $12
label_18631: AND (HL)
label_18632: AND (IX + 127)
label_18633: AND (IY - 128)
label_18634: OR A
label_18635: OR B
label_18636: OR C
label_18637: OR D
label_18638: OR E
label_18639: OR H
label_18640: OR L
label_18641: OR $12
label_18642: OR (HL)
label_18643: OR (IX + 127)
label_18644: OR (IY - 128)
label_18645: XOR A
label_18646: XOR B
label_18647: XOR C
label_18648: XOR D
label_18649: XOR E
label_18650: XOR H
label_18651: XOR L
label_18652: XOR $12
label_18653: XOR (HL)
label_18654: XOR (IX + 127)
label_18655: XOR (IY - 128)
label_18656: CP A
label_18657: CP B
label_18658: CP C
label_18659: CP D
label_18660: CP E
label_18661: CP H
label_18662: CP L
label_18663: CP $12
label_18664: CP (HL)
label_18665: CP (IX + 127)
label_18666: CP (IY - 128)
label_18667: INC A
label_18668: INC B
label_18669: INC C
label_18670: INC D
label_18671: INC E
label_18672: INC H
label_18673: INC L
label_18674: INC (HL)
label_18675: INC (IX + 127)
label_18676: INC (IY - 128)
label_18677: DEC A
label_18678: DEC B
label_18679: DEC C
label_18680: DEC D
label_18681: DEC E
label_18682: DEC H
label_18683: DEC L
label_18684: DEC (HL)
label_18685: DEC (IX + 127)
label_18686: DEC (IY - 128)
label_18687: DAA
label_18688: CPL
label_18689: NEG
label_18690: CCF
label_18691: SCF
label_18692: NOP
label_18693: HALT
label_18694: DI
label_18695: EI
label_18696: IM 0
label_18697: IM 1
label_18698: IM 2
label_18699: ADD HL,BC
label_18700: ADD HL,DE
label_18701: ADD HL,HL
label_18702: ADD HL,SP
label_18703: ADC HL,BC
label_18704: ADC HL,DE
label_18705: ADC HL,HL
label_18706: ADC HL,SP
label_18707: SBC HL,BC
label_18708: SBC HL,DE
label_18709: SBC HL,HL
label_18710: SBC HL,SP
label_18711: ADD IX,BC
label_18712: ADD IX,DE
label_18713: ADD IX,SP
label_18714: ADD IY,BC
label_18715: ADD IY,DE
label_18716: ADD IY,SP
label_18717: INC BC
label_18718: INC DE
label_18719: INC HL
label_18720: INC SP
label_18721: INC IX
label_18722: INC IY
label_18723: DEC BC
label_18724: DEC DE
label_18725: DEC HL
label_18726: DEC SP
label_18727: DEC IX
label_18728: DEC IY
label_18729: RLCA
label_18730: RLA
label_18731: RRCA
label_18732: RRA
label_18733: RLC A
label_18734: RLC B
label_18735: RLC C
label_18736: RLC D
label_18737: RLC E
label_18738: RLC H
label_18739: RLC L
label_18740: RLC (HL)
label_18741: RLC (IX + 127)
label_18742: RLC (IY - 128)
label_18743: RL A
label_18744: RL B
label_18745: RL C
label_18746: RL D
label_18747: RL E
label_18748: RL H
label_18749: RL L
label_18750: RL (HL)
label_18751: RL (IX + 127)
label_18752: RL (IY - 128)
label_18753: RRC A
label_18754: RRC B
label_18755: RRC C
label_18756: RRC D
label_18757: RRC E
label_18758: RRC H
label_18759: RRC L
label_18760: RRC (HL)
label_18761: RRC (IX + 127)
label_18762: RRC (IY - 128)
label_18763: RR A
label_18764: RR B
label_18765: RR C
label_18766: RR D
label_18767: RR E
label_18768: RR H
label_18769: RR L
label_18770: RR (HL)
label_18771: RR (IX + 127)
label_18772: RR (IY - 128)
label_18773: SLA A
label_18774: SLA B
label_18775: SLA C
label_18776: SLA D
label_18777: SLA E
label_18778: SLA H
label_18779: SLA L
label_18780: SLA (HL)
label_18781: SLA (IX + 127)
label_18782: SLA (IY - 128)
label_18783: SRA A
label_18784: SRA B
label_18785: SRA C
label_18786: SRA D
label_18787: SRA E
label_18788: SRA H
label_18789: SRA L
label_18790: SRA (HL)
label_18791: SRA (IX + 127)
label_18792: SRA (IY - 128)
label_18793: SRL A
label_18794: SRL B
label_18795: SRL C
label_18796: SRL D
label_18797: SRL E
label_18798: SRL H
label_18799: SRL L
label_18800: SRL (HL)
label_18801: SRL (IX + 127)
label_18802: SRL (IY - 128)
label_18803: RLD
label_18804: RRD
label_18805: BIT 0,A
label_18806: BIT 1,A
label_18807: BIT 2,A
label_18808: BIT 3,A
label_18809: BIT 4,A
label_18810: BIT 5,A
label_18811: BIT 6,A
label_18812: BIT 7,A
label_18813: BIT 0,B
label_18814: BIT 1,B
label_18815: BIT 2,B
label_18816: BIT 3,B
label_18817: BIT 4,B
label_18818: BIT 5,B
label_18819: BIT 6,B
label_18820: BIT 7,B
label_18821: BIT 0,C
label_18822: BIT 1,C
label_18823: BIT 2,C
label_18824: BIT 3,C
label_18825: BIT 4,C
label_18826: BIT 5,C
label_18827: BIT 6,C
label_18828: BIT 7,C
label_18829: BIT 0,D
label_18830: BIT 1,D
label_18831: BIT 2,D
label_18832: BIT 3,D
label_18833: BIT 4,D
label_18834: BIT 5,D
label_18835: BIT 6,D
label_18836: BIT 7,D
label_18837: BIT 0,E
label_18838: BIT 1,E
label_18839: BIT 2,E
label_18840: BIT 3,E
label_18841: BIT 4,E
label_18842: BIT 5,E
label_18843: BIT 6,E
label_18844: BIT 7,E
label_18845: BIT 0,H
label_18846: BIT 1,H
label_18847: BIT 2,H
label_18848: BIT 3,H
label_18849: BIT 4,H
label_18850: BIT 5,H
label_18851: BIT 6,H
label_18852: BIT 7,H
label_18853: BIT 0,L
label_18854: BIT 1,L
label_18855: BIT 2,L
label_18856: BIT 3,L
label_18857: BIT 4,L
label_18858: BIT 5,L
label_18859: BIT 6,L
label_18860: BIT 7,L
label_18861: BIT 0,(HL)
label_18862: BIT 1,(HL)
label_18863: BIT 2,(HL)
label_18864: BIT 3,(HL)
label_18865: BIT 4,(HL)
label_18866: BIT 5,(HL)
label_18867: BIT 6,(HL)
label_18868: BIT 7,(HL)
label_18869: BIT 0,(IX + 127)
label_18870: BIT 1,(IX + 127)
label_18871: BIT 2,(IX + 127)
label_18872: BIT 3,(IX + 127)
label_18873: BIT 4,(IX + 127)
label_18874: BIT 5,(IX + 127)
label_18875: BIT 6,(IX + 127)
label_18876: BIT 7,(IX + 127)
label_18877: BIT 0,(IY - 128)
label_18878: BIT 1,(IY - 128)
label_18879: BIT 2,(IY - 128)
label_18880: BIT 3,(IY - 128)
label_18881: BIT 4,(IY - 128)
label_18882: BIT 5,(IY - 128)
label_18883: BIT 6,(IY - 128)
label_18884: BIT 7,(IY - 128)
label_18885: SET 0,A
label_18886: SET 1,A
label_18887: SET 2,A
label_18888: SET 3,A
label_18889: SET 4,A
label_18890: SET 5,A
label_18891: SET 6,A
label_18892: SET 7,A
label_18893: SET 0,B
label_18894: SET 1,B
label_18895: SET 2,B
label_18896: SET 3,B
label_18897: SET 4,B
label_18898: SET 5,B
label_18899: SET 6,B
label_18900: SET 7,B
label_18901: SET 0,C
label_18902: SET 1,C
label_18903: SET 2,C
label_18904: SET 3,C
label_18905: SET 4,C
label_18906: SET 5,C
label_18907: SET 6,C
label_18908: SET 7,C
label_18909: SET 0,D
label_18910: SET 1,D
label_18911: SET 2,D
label_18912: SET 3,D
label_18913: SET 4,D
label_18914: SET 5,D
label_18915: SET 6,D
label_18916: SET 7,D
label_18917: SET 0,E
label_18918: SET 1,E
label_18919: SET 2,E
label_18920: SET 3,E
label_18921: SET 4,E
label_18922: SET 5,E
label_18923: SET 6,E
label_18924: SET 7,E
label_18925: SET 0,H
label_18926: SET 1,H
label_18927: SET 2,H
label_18928: SET 3,H
label_18929: SET 4,H
label_18930: SET 5,H
label_18931: SET 6,H
label_18932: SET 7,H
label_18933: SET 0,L
label_18934: SET 1,L
label_18935: SET 2,L
label_18936: SET 3,L
label_18937: SET 4,L
label_18938: SET 5,L
label_18939: SET 6,L
label_18940: SET 7,L
label_18941: SET 0,(HL)
label_18942: SET 1,(HL)
label_18943: SET 2,(HL)
label_18944: SET 3,(HL)
label_18945: SET 4,(HL)
label_18946: SET 5,(HL)
label_18947: SET 6,(HL)
label_18948: SET 7,(HL)
label_18949: SET 0,(IX + 127)
label_18950: SET 1,(IX + 127)
label_18951: SET 2,(IX + 127)
label_18952: SET 3,(IX + 127)
label_18953: SET 4,(IX + 127)
label_18954: SET 5,(IX + 127)
label_18955: SET 6,(IX + 127)
label_18956: SET 7,(IX + 127)
label_18957: SET 0,(IY - 128)
label_18958: SET 1,(IY - 128)
label_18959: SET 2,(IY - 128)
label_18960: SET 3,(IY - 128)
label_18961: SET 4,(IY - 128)
label_18962: SET 5,(IY - 128)
label_18963: SET 6,(IY - 128)
label_18964: SET 7,(IY - 128)
label_18965: RES 0,A
label_18966: RES 1,A
label_18967: RES 2,A
label_18968: RES 3,A
label_18969: RES 4,A
label_18970: RES 5,A
label_18971: RES 6,A
label_18972: RES 7,A
label_18973: RES 0,B
label_18974: RES 1,B
label_18975: RES 2,B
label_18976: RES 3,B
label_18977: RES 4,B
label_18978: RES 5,B
label_18979: RES 6,B
label_18980: RES 7,B
label_18981: RES 0,C
label_18982: RES 1,C
label_18983: RES 2,C
label_18984: RES 3,C
label_18985: RES 4,C
label_18986: RES 5,C
label_18987: RES 6,C
label_18988: RES 7,C
label_18989: RES 0,D
label_18990: RES 1,D
label_18991: RES 2,D
label_18992: RES 3,D
label_18993: RES 4,D
label_18994: RES 5,D
label_18995: RES 6,D
label_18996: RES 7,D
label_18997: RES 0,E
label_18998: RES 1,E
label_18999: RES 2,E
label_19000: RES 3,E
label_19001: RES 4,E
label_19002: RES 5,E
label_19003: RES 6,E
label_19004: RES 7,E
label_19005: RES 0,H
label_19006: RES 1,H
label_19007: RES 2,H
label_19008: RES 3,H
label_19009: RES 4,H
label_19010: RES 5,H
label_19011: RES 6,H
label_19012: RES 7,H
label_19013: RES 0,L
label_19014: RES 1,L
label_19015: RES 2,L
label_19016: RES 3,L
label_19017: RES 4,L
label_19018: RES 5,L
label_19019: RES 6,L
label_19020: RES 7,L
label_19021: RES 0,(HL)
label_19022: RES 1,(HL)
label_19023: RES 2,(HL)
label_19024: RES 3,(HL)
label_19025: RES 4,(HL)
label_19026: RES 5,(HL)
label_19027: RES 6,(HL)
label_19028: RES 7,(HL)
label_19029: RES 0,(IX + 127)
label_19030: RES 1,(IX + 127)
label_19031: RES 2,(IX + 127)
label_19032: RES 3,(IX + 127)
label_19033: RES 4,(IX + 127)
label_19034: RES 5,(IX + 127)
label_19035: RES 6,(IX + 127)
label_19036: RES 7,(IX + 127)
label_19037: RES 0,(IY - 128)
label_19038: RES 1,(IY - 128)
label_19039: RES 2,(IY - 128)
label_19040: RES 3,(IY - 128)
label_19041: RES 4,(IY - 128)
label_19042: RES 5,(IY - 128)
label_19043: RES 6,(IY - 128)
label_19044: RES 7,(IY - 128)
label_19045: JP $5678
label_19046: JP NZ,$5678
label_19047: JP Z,$5678
label_19048: JP NC,$5678
label_19049: JP C,$5678
label_19050: JP PO,$5678
label_19051: JP PE,$5678
label_19052: JP P,$5678
label_19053: JP M,$5678
label_19054: JR $ + 2
label_19055: JR NZ,$ + 2
label_19056: JR Z,$ + 2
label_19057: JR NC,$ + 2
label_19058: JR C,$ + 2
label_19059: JP (HL)
label_19060: JP (IX)
label_19061: JP (IY)
label_19062: DJNZ $ + 2
label_19063: CALL $5678
label_19064: CALL NZ,$5678
label_19065: CALL Z,$5678
label_19066: CALL NC,$5678
label_19067: CALL C,$5678
label_19068: CALL PO,$5678
label_19069: CALL PE,$5678
label_19070: CALL P,$5678
label_19071: CALL M,$5678
label_19072: RET
label_19073: RET NZ
label_19074: RET Z
label_19075: RET NC
label_19076: RET C
label_19077: RET PO
label_19078: RET PE
label_19079: RET P
label_19080: RET M
label_19081: RETI
label_19082: RETN
label_19083: RST $00
label_19084: RST $08
label_19085: RST $10
label_19086: RST $18
label_19087: RST $20
label_19088: RST $28
label_19089: RST $30
label_19090: RST $38
label_19091: IN A,($12)
label_19092: IN A,(C)
label_19093: IN B,(C)
label_19094: IN C,(C)
label_19095: IN D,(C)
label_19096: IN E,(C)
label_19097: IN H,(C)
label_19098: IN L,(C)
label_19099: IN F,(C)
label_19100: INI
label_19101: INIR
label_19102: IND
label_19103: INDR
label_19104: OUT ($12),A
label_19105: OUT (C),A
label_19106: OUT (C),B
label_19107: OUT (C),C
label_19108: OUT (C),D
label_19109: OUT (C),E
label_19110: OUT (C),H
label_19111: OUT (C),L
label_19112: OUTI
label_19113: OTIR
label_19114: OUTD
label_19115: OTDR
label_19116: LD A,A
label_19117: LD A,B
label_19118: LD A,C
label_19119: LD A,D
label_19120: LD A,E
label_19121: LD A,H
label_19122: LD A,L
label_19123: LD B,A
label_19124: LD B,B
label_19125: LD B,C
label_19126: LD B,D
label_19127: LD B,E
label_19128: LD B,H
label_19129: LD B,L
label_19130: LD C,A
label_19131: LD C,B
label_19132: LD C,C
label_19133: LD C,D
label_19134: LD C,E
label_19135: LD C,H
label_19136: LD C,L
label_19137: LD D,A
label_19138: LD D,B
label_19139: LD D,C
label_19140: LD D,D
label_19141: LD D,E
label_19142: LD D,H
label_19143: LD D,L
label_19144: LD E,A
label_19145: LD E,B
label_19146: LD E,C
label_19147: LD E,D
label_19148: LD E,E
label_19149: LD E,H
label_19150: LD E,L
label_19151: LD H,A
label_19152: LD H,B
label_19153: LD H,C
label_19154: LD H,D
label_19155: LD H,E
label_19156: LD H,H
label_19157: LD H,L
label_19158: LD L,A
label_19159: LD L,B
label_19160: LD L,C
label_19161: LD L,D
label_19162: LD L,E
label_19163: LD L,H
label_19164: LD L,L
label_19165: LD A,$12
label_19166: LD B,$12
label_19167: LD C,$12
label_19168: LD D,$12
label_19169: LD E,$12
label_19170: LD H,$12
label_19171: LD L,$12
label_19172: LD A,(HL)
label_19173: LD B,(HL)
label_19174: LD C,(HL)
label_19175: LD D,(HL)
label_19176: LD E,(HL)
label_19177: LD H,(HL)
label_19178: LD L,(HL)
label_19179: LD A,(IX + 127)
label_19180: LD B,(IX + 127)
label_19181: LD C,(IX + 127)
label_19182: LD D,(IX + 127)
label_19183: LD E,(IX + 127)
label_19184: LD H,(IX + 127)
label_19185: LD L,(IX + 127)
label_19186: LD A,(IY - 128)
label_19187: LD B,(IY - 128)
label_19188: LD C,(IY - 128)
label_19189: LD D,(IY - 128)
label_19190: LD E,(IY - 128)
label_19191: LD H,(IY - 128)
label_19192: LD L,(IY - 128)
label_19193: LD (HL),A
label_19194: LD (HL),B
label_19195: LD (HL),C
label_19196: LD (HL),D
label_19197: LD (HL),E
label_19198: LD (HL),H
label_19199: LD (HL),L
label_19200: LD (IX + 127),A
label_19201: LD (IX + 127),B
label_19202: LD (IX + 127),C
label_19203: LD (IX + 127),D
label_19204: LD (IX + 127),E
label_19205: LD (IX + 127),H
label_19206: LD (IX + 127),L
label_19207: LD (IY - 128),A
label_19208: LD (IY - 128),B
label_19209: LD (IY - 128),C
label_19210: LD (IY - 128),D
label_19211: LD (IY - 128),E
label_19212: LD (IY - 128),H
label_19213: LD (IY - 128),L
label_19214: LD (HL),$12
label_19215: LD (IX + 127),$12
label_19216: LD (IY - 128),$12
label_19217: LD A,(BC)
label_19218: LD A,(DE)
label_19219: LD A,($5678)
label_19220: LD (BC),A
label_19221: LD (DE),A
label_19222: LD ($5678),A
label_19223: LD A,I
label_19224: LD A,R
label_19225: LD I,A
label_19226: LD R,A
label_19227: LD BC,$5678
label_19228: LD DE,$5678
label_19229: LD HL,$5678
label_19230: LD SP,$5678
label_19231: LD IX,$5678
label_19232: LD IY,$5678
label_19233: LD HL,($5678)
label_19234: LD BC,($5678)
label_19235: LD DE,($5678)
label_19236: LD HL,($5678)
label_19237: LD SP,($5678)
label_19238: LD IX,($5678)
label_19239: LD IY,($5678)
label_19240: LD ($5678),HL
label_19241: LD ($5678),BC
label_19242: LD ($5678),DE
label_19243: LD ($5678),HL
label_19244: LD ($5678),SP
label_19245: LD ($5678),IX
label_19246: LD ($5678),IY
label_19247: LD SP,HL
label_19248: LD SP,IX
label_19249: LD SP,IY
label_19250: PUSH BC
label_19251: PUSH DE
label_19252: PUSH HL
label_19253: PUSH AF
label_19254: PUSH IX
label_19255: PUSH IY
label_19256: POP BC
label_19257: POP DE
label_19258: POP HL
label_19259: POP AF
label_19260: POP IX
label_19261: POP IY
label_19262: EX DE,HL
label_19263: EX AF,AF'
label_19264: EXX
label_19265: EX (SP),HL
label_19266: EX (SP),IX
label_19267: EX (SP),IY
label_19268: LDI
label_19269: LDIR
label_19270: LDD
label_19271: LDDR
label_19272: CPI
label_19273: CPIR
label_19274: CPD
label_19275: CPDR
label_19276: ADD A,A
label_19277: ADD A,B
label_19278: ADD A,C
label_19279: ADD A,D
label_19280: ADD A,E
label_19281: ADD A,H
label_19282: ADD A,L
label_19283: ADD A,$12
label_19284: ADD A,(HL)
label_19285: ADD A,(IX + 127)
label_19286: ADD A,(IY - 128)
label_19287: ADC A,A
label_19288: ADC A,B
label_19289: ADC A,C
label_19290: ADC A,D
label_19291: ADC A,E
label_19292: ADC A,H
label_19293: ADC A,L
label_19294: ADC A,$12
label_19295: ADC A,(HL)
label_19296: ADC A,(IX + 127)
label_19297: ADC A,(IY - 128)
label_19298: SUB A
label_19299: SUB B
label_19300: SUB C
label_19301: SUB D
label_19302: SUB E
label_19303: SUB H
label_19304: SUB L
label_19305: SUB $12
label_19306: SUB (HL)
label_19307: SUB (IX + 127)
label_19308: SUB (IY - 128)
label_19309: SBC A,A
label_19310: SBC A,B
label_19311: SBC A,C
label_19312: SBC A,D
label_19313: SBC A,E
label_19314: SBC A,H
label_19315: SBC A,L
label_19316: SBC A,$12
label_19317: SBC A,(HL)
label_19318: SBC A,(IX + 127)
label_19319: SBC A,(IY - 128)
label_19320: AND A
label_19321: AND B
label_19322: AND C
label_19323: AND D
label_19324: AND E
label_19325: AND H
label_19326: AND L
label_19327: AND $12
label_19328: AND (HL)
label_19329: AND (IX + 127)
label_19330: AND (IY - 128)
label_19331: AND A
label_19332: AND B
label_19333: AND C
label_19334: AND D
label_19335: AND E
label_19336: AND H
label_19337: AND L
label_19338: AND $12
label_19339: AND (HL)
label_19340: AND (IX + 127)
label_19341: AND (IY - 128)
label_19342: OR A
label_19343: OR B
label_19344: OR C
label_19345: OR D
label_19346: OR E
label_19347: OR H
label_19348: OR L
label_19349: OR $12
label_19350: OR (HL)
label_19351: OR (IX + 127)
label_19352: OR (IY - 128)
label_19353: XOR A
label_19354: XOR B
label_19355: XOR C
label_19356: XOR D
label_19357: XOR E
label_19358: XOR H
label_19359: XOR L
label_19360: XOR $12
label_19361: XOR (HL)
label_19362: XOR (IX + 127)
label_19363: XOR (IY - 128)
label_19364: CP A
label_19365: CP B
label_19366: CP C
label_19367: CP D
label_19368: CP E
label_19369: CP H
label_19370: CP L
label_19371: CP $12
label_19372: CP (HL)
label_19373: CP (IX + 127)
label_19374: CP (IY - 128)
label_19375: INC A
label_19376: INC B
label_19377: INC C
label_19378: INC D
label_19379: INC E
label_19380: INC H
label_19381: INC L
label_19382: INC (HL)
label_19383: INC (IX + 127)
label_19384: INC (IY - 128)
label_19385: DEC A
label_19386: DEC B
label_19387: DEC C
label_19388: DEC D
label_19389: DEC E
label_19390: DEC H
label_19391: DEC L
label_19392: DEC (HL)
label_19393: DEC (IX + 127)
label_19394: DEC (IY - 128)
label_19395: DAA
label_19396: CPL
label_19397: NEG
label_19398: CCF
label_19399: SCF
label_19400: NOP
label_19401: HALT
label_19402: DI
label_19403: EI
label_19404: IM 0
label_19405: IM 1
label_19406: IM 2
label_19407: ADD HL,BC
label_19408: ADD HL,DE
label_19409: ADD HL,HL
label_19410: ADD HL,SP
label_19411: ADC HL,BC
label_19412: ADC HL,DE
label_19413: ADC HL,HL
label_19414: ADC HL,SP
label_19415: SBC HL,BC
label_19416: SBC HL,DE
label_19417: SBC HL,HL
label_19418: SBC HL,SP
label_19419: ADD IX,BC
label_19420: ADD IX,DE
label_19421: ADD IX,SP
label_19422: ADD IY,BC
label_19423: ADD IY,DE
label_19424: ADD IY,SP
label_19425: INC BC
label_19426: INC DE
label_19427: INC HL
label_19428: INC SP
label_19429: INC IX
label_19430: INC IY
label_19431: DEC BC
label_19432: DEC DE
label_19433: DEC HL
label_19434: DEC SP
label_19435: DEC IX
label_19436: DEC IY
label_19437: RLCA
label_19438: RLA
label_19439: RRCA
label_19440: RRA
label_19441: RLC A
label_19442: RLC B
label_19443: RLC C
label_19444: RLC D
label_19445: RLC E
label_19446: RLC H
label_19447: RLC L
label_19448: RLC (HL)
label_19449: RLC (IX + 127)
label_19450: RLC (IY - 128)
label_19451: RL A
label_19452: RL B
label_19453: RL C
label_19454: RL D
label_19455: RL E
label_19456: RL H
label_19457: RL L
label_19458: RL (HL)
label_19459: RL (IX + 127)
label_19460: RL (IY - 128)
label_19461: RRC A
label_19462: RRC B
label_19463: RRC C
label_19464: RRC D
label_19465: RRC E
label_19466: RRC H
label_19467: RRC L
label_19468: RRC (HL)
label_19469: RRC (IX + 127)
label_19470: RRC (IY - 128)
label_19471: RR A
label_19472: RR B
label_19473: RR C
label_19474: RR D
label_19475: RR E
label_19476: RR H
label_19477: RR L
label_19478: RR (HL)
label_19479: RR (IX + 127)
label_19480: RR (IY - 128)
label_19481: SLA A
label_19482: SLA B
label_19483: SLA C
label_19484: SLA D
label_19485: SLA E
label_19486: SLA H
label_19487: SLA L
label_19488: SLA (HL)
label_19489: SLA (IX + 127)
label_19490: SLA (IY - 128)
label_19491: SRA A
label_19492: SRA B
label_19493: SRA C
label_19494: SRA D
label_19495: SRA E
label_19496: SRA H
label_19497: SRA L
label_19498: SRA (HL)
label_19499: SRA (IX + 127)
label_19500: SRA (IY - 128)
label_19501: SRL A
label_19502: SRL B
label_19503: SRL C
label_19504: SRL D
label_19505: SRL E
label_19506: SRL H
label_19507: SRL L
label_19508: SRL (HL)
label_19509: SRL (IX + 127)
label_19510: SRL (IY - 128)
label_19511: RLD
label_19512: RRD
label_19513: BIT 0,A
label_19514: BIT 1,A
label_19515: BIT 2,A
label_19516: BIT 3,A
label_19517: BIT 4,A
label_19518: BIT 5,A
label_19519: BIT 6,A
label_19520: BIT 7,A
label_19521: BIT 0,B
label_19522: BIT 1,B
label_19523: BIT 2,B
label_19524: BIT 3,B
label_19525: BIT 4,B
label_19526: BIT 5,B
label_19527: BIT 6,B
label_19528: BIT 7,B
label_19529: BIT 0,C
label_19530: BIT 1,C
label_19531: BIT 2,C
label_19532: BIT 3,C
label_19533: BIT 4,C
label_19534: BIT 5,C
label_19535: BIT 6,C
label_19536: BIT 7,C
label_19537: BIT 0,D
label_19538: BIT 1,D
label_19539: BIT 2,D
label_19540: BIT 3,D
label_19541: BIT 4,D
label_19542: BIT 5,D
label_19543: BIT 6,D
label_19544: BIT 7,D
label_19545: BIT 0,E
label_19546: BIT 1,E
label_19547: BIT 2,E
label_19548: BIT 3,E
label_19549: BIT 4,E
label_19550: BIT 5,E
label_19551: BIT 6,E
label_19552: BIT 7,E
label_19553: BIT 0,H
label_19554: BIT 1,H
label_19555: BIT 2,H
label_19556: BIT 3,H
label_19557: BIT 4,H
label_19558: BIT 5,H
label_19559: BIT 6,H
label_19560: BIT 7,H
label_19561: BIT 0,L
label_19562: BIT 1,L
label_19563: BIT 2,L
label_19564: BIT 3,L
label_19565: BIT 4,L
label_19566: BIT 5,L
label_19567: BIT 6,L
label_19568: BIT 7,L
label_19569: BIT 0,(HL)
label_19570: BIT 1,(HL)
label_19571: BIT 2,(HL)
label_19572: BIT 3,(HL)
label_19573: BIT 4,(HL)
label_19574: BIT 5,(HL)
label_19575: BIT 6,(HL)
label_19576: BIT 7,(HL)
label_19577: BIT 0,(IX + 127)
label_19578: BIT 1,(IX + 127)
label_19579: BIT 2,(IX + 127)
label_19580: BIT 3,(IX + 127)
label_19581: BIT 4,(IX + 127)
label_19582: BIT 5,(IX + 127)
label_19583: BIT 6,(IX + 127)
label_19584: BIT 7,(IX + 127)
label_19585: BIT 0,(IY - 128)
label_19586: BIT 1,(IY - 128)
label_19587: BIT 2,(IY - 128)
label_19588: BIT 3,(IY - 128)
label_19589: BIT 4,(IY - 128)
label_19590: BIT 5,(IY - 128)
label_19591: BIT 6,(IY - 128)
label_19592: BIT 7,(IY - 128)
label_19593: SET 0,A
label_19594: SET 1,A
label_19595: SET 2,A
label_19596: SET 3,A
label_19597: SET 4,A
label_19598: SET 5,A
label_19599: SET 6,A
label_19600: SET 7,A
label_19601: SET 0,B
label_19602: SET 1,B
label_19603: SET 2,B
label_19604: SET 3,B
label_19605: SET 4,B
label_19606: SET 5,B
label_19607: SET 6,B
label_19608: SET 7,B
label_19609: SET 0,C
label_19610: SET 1,C
label_19611: SET 2,C
label_19612: SET 3,C
label_19613: SET 4,C
label_19614: SET 5,C
label_19615: SET 6,C
label_19616: SET 7,C
label_19617: SET 0,D
label_19618: SET 1,D
label_19619: SET 2,D
label_19620: SET 3,D
label_19621: SET 4,D
label_19622: SET 5,D
label_19623: SET 6,D
label_19624: SET 7,D
label_19625: SET 0,E
label_19626: SET 1,E
label_19627: SET 2,E
label_19628: SET 3,E
label_19629: SET 4,E
label_19630: SET 5,E
label_19631: SET 6,E
label_19632: SET 7,E
label_19633: SET 0,H
label_19634: SET 1,H
label_19635: SET 2,H
label_19636: SET 3,H
label_19637: SET 4,H
label_19638: SET 5,H
label_19639: SET 6,H
label_19640: SET 7,H
label_19641: SET 0,L
label_19642: SET 1,L
label_19643: SET 2,L
label_19644: SET 3,L
label_19645: SET 4,L
label_19646: SET 5,L
label_19647: SET 6,L
label_19648: SET 7,L
label_19649: SET 0,(HL)
label_19650: SET 1,(HL)
label_19651: SET 2,(HL)
label_19652: SET 3,(HL)
label_19653: SET 4,(HL)
label_19654: SET 5,(HL)
label_19655: SET 6,(HL)
label_19656: SET 7,(HL)
label_19657: SET 0,(IX + 127)
label_19658: SET 1,(IX + 127)
label_19659: SET 2,(IX + 127)
label_19660: SET 3,(IX + 127)
label_19661: SET 4,(IX + 127)
label_19662: SET 5,(IX + 127)
label_19663: SET 6,(IX + 127)
label_19664: SET 7,(IX + 127)
label_19665: SET 0,(IY - 128)
label_19666: SET 1,(IY - 128)
label_19667: SET 2,(IY - 128)
label_19668: SET 3,(IY - 128)
label_19669: SET 4,(IY - 128)
label_19670: SET 5,(IY - 128)
label_19671: SET 6,(IY - 128)
label_19672: SET 7,(IY - 128)
label_19673: RES 0,A
label_19674: RES 1,A
label_19675: RES 2,A
label_19676: RES 3,A
label_19677: RES 4,A
label_19678: RES 5,A
label_19679: RES 6,A
label_19680: RES 7,A
label_19681: RES 0,B
label_19682: RES 1,B
label_19683: RES 2,B
label_19684: RES 3,B
label_19685: RES 4,B
label_19686: RES 5,B
label_19687: RES 6,B
label_19688: RES 7,B
label_19689: RES 0,C
label_19690: RES 1,C
label_19691: RES 2,C
label_19692: RES 3,C
label_19693: RES 4,C
label_19694: RES 5,C
label_19695: RES 6,C
label_19696: RES 7,C
label_19697: RES 0,D
label_19698: RES 1,D
label_19699: RES 2,D
label_19700: RES 3,D
label_19701: RES 4,D
label_19702: RES 5,D
label_19703: RES 6,D
label_19704: RES 7,D
label_19705: RES 0,E
label_19706: RES 1,E
label_19707: RES 2,E
label_19708: RES 3,E
label_19709: RES 4,E
label_19710: RES 5,E
label_19711: RES 6,E
label_19712: RES 7,E
label_19713: RES 0,H
label_19714: RES 1,H
label_19715: RES 2,H
label_19716: RES 3,H
label_19717: RES 4,H
label_19718: RES 5,H
label_19719: RES 6,H
label_19720: RES 7,H
label_19721: RES 0,L
label_19722: RES 1,L
label_19723: RES 2,L
label_19724: RES 3,L
label_19725: RES 4,L
label_19726: RES 5,L
label_19727: RES 6,L
label_19728: RES 7,L
label_19729: RES 0,(HL)
label_19730: RES 1,(HL)
label_19731: RES 2,(HL)
label_19732: RES 3,(HL)
label_19733: RES 4,(HL)
label_19734: RES 5,(HL)
label_19735: RES 6,(HL)
label_19736: RES 7,(HL)
label_19737: RES 0,(IX + 127)
label_19738: RES 1,(IX + 127)
label_19739: RES 2,(IX + 127)
label_19740: RES 3,(IX + 127)
label_19741: RES 4,(IX + 127)
label_19742: RES 5,(IX + 127)
label_19743: RES 6,(IX + 127)
label_19744: RES 7,(IX + 127)
label_19745: RES 0,(IY - 128)
label_19746: RES 1,(IY - 128)
label_19747: RES 2,(IY - 128)
label_19748: RES 3,(IY - 128)
label_19749: RES 4,(IY - 128)
label_19750: RES 5,(IY - 128)
label_19751: RES 6,(IY - 128)
label_19752: RES 7,(IY - 128)
label_19753: JP $5678
label_19754: JP NZ,$5678
label_19755: JP Z,$5678
label_19756: JP NC,$5678
label_19757: JP C,$5678
label_19758: JP PO,$5678
label_19759: JP PE,$5678
label_19760: JP P,$5678
label_19761: JP M,$5678
label_19762: JR $ + 2
label_19763: JR NZ,$ + 2
label_19764: JR Z,$ + 2
label_19765: JR NC,$ + 2
label_19766: JR C,$ + 2
label_19767: JP (HL)
label_19768: JP (IX)
label_19769: JP (IY)
label_19770: DJNZ $ + 2
label_19771: CALL $5678
label_19772: CALL NZ,$5678
label_19773: CALL Z,$5678
label_19774: CALL NC,$5678
label_19775: CALL C,$5678
label_19776: CALL PO,$5678
label_19777: CALL PE,$5678
label_19778: CALL P,$5678
label_19779: CALL M,$5678
label_19780: RET
label_19781: RET NZ
label_19782: RET Z
label_19783: RET NC
label_19784: RET C
label_19785: RET PO
label_19786: RET PE
label_19787: RET P
label_19788: RET M
label_19789: RETI
label_19790: RETN
label_19791: RST $00
label_19792: RST $08
label_19793: RST $10
label_19794: RST $18
label_19795: RST $20
label_19796: RST $28
label_19797: RST $30
label_19798: RST $38
label_19799: IN A,($12)
label_19800: IN A,(C)
label_19801: IN B,(C)
label_19802: IN C,(C)
label_19803: IN D,(C)
label_19804: IN E,(C)
label_19805: IN H,(C)
label_19806: IN L,(C)
label_19807: IN F,(C)
label_19808: INI
label_19809: INIR
label_19810: IND
label_19811: INDR
label_19812: OUT ($12),A
label_19813: OUT (C),A
label_19814: OUT (C),B
label_19815: OUT (C),C
label_19816: OUT (C),D
label_19817: OUT (C),E
label_19818: OUT (C),H
label_19819: OUT (C),L
label_19820: OUTI
label_19821: OTIR
label_19822: OUTD
label_19823: OTDR
label_19824: LD A,A
label_19825: LD A,B
label_19826: LD A,C
label_19827: LD A,D
label_19828: LD A,E
label_19829: LD A,H
label_19830: LD A,L
label_19831: LD B,A
label_19832: LD B,B
label_19833: LD B,C
label_19834: LD B,D
label_19835: LD B,E
label_19836: LD B,H
label_19837: LD B,L
label_19838: LD C,A
label_19839: LD C,B
label_19840: LD C,C
label_19841: LD C,D
label_19842: LD C,E
label_19843: LD C,H
label_19844: LD C,L
label_19845: LD D,A
label_19846: LD D,B
label_19847: LD D,C
label_19848: LD D,D
label_19849: LD D,E
label_19850: LD D,H
label_19851: LD D,L
label_19852: LD E,A
label_19853: LD E,B
label_19854: LD E,C
label_19855: LD E,D
label_19856: LD E,E
label_19857: LD E,H
label_19858: LD E,L
label_19859: LD H,A
label_19860: LD H,B
label_19861: LD H,C
label_19862: LD H,D
label_19863: LD H,E
label_19864: LD H,H
label_19865: LD H,L
label_19866: LD L,A
label_19867: LD L,B
label_19868: LD L,C
label_19869: LD L,D
label_19870: LD L,E
label_19871: LD L,H
label_19872: LD L,L
label_19873: LD A,$12
label_19874: LD B,$12
label_19875: LD C,$12
label_19876: LD D,$12
label_19877: LD E,$12
label_19878: LD H,$12
label_19879: LD L,$12
label_19880: LD A,(HL)
label_19881: LD B,(HL)
label_19882: LD C,(HL)
label_19883: LD D,(HL)
label_19884: LD E,(HL)
label_19885: LD H,(HL)
label_19886: LD L,(HL)
label_19887: LD A,(IX + 127)
label_19888: LD B,(IX + 127)
label_19889: LD C,(IX + 127)
label_19890: LD D,(IX + 127)
label_19891: LD E,(IX + 127)
label_19892: LD H,(IX + 127)
label_19893: LD L,(IX + 127)
label_19894: LD A,(IY - 128)
label_19895: LD B,(IY - 128)
label_19896: LD C,(IY - 128)
label_19897: LD D,(IY - 128)
label_19898: LD E,(IY - 128)
label_19899: LD H,(IY - 128)
label_19900: LD L,(IY - 128)
label_19901: LD (HL),A
label_19902: LD (HL),B
label_19903: LD (HL),C
label_19904: LD (HL),D
label_19905: LD (HL),E
label_19906: LD (HL),H
label_19907: LD (HL),L
label_19908: LD (IX + 127),A
label_19909: LD (IX + 127),B
label_19910: LD (IX + 127),C
label_19911: LD (IX + 127),D
label_19912: LD (IX + 127),E
label_19913: LD (IX + 127),H
label_19914: LD (IX + 127),L
label_19915: LD (IY - 128),A
label_19916: LD (IY - 128),B
label_19917: LD (IY - 128),C
label_19918: LD (IY - 128),D
label_19919: LD (IY - 128),E
label_19920: LD (IY - 128),H
label_19921: LD (IY - 128),L
label_19922: LD (HL),$12
label_19923: LD (IX + 127),$12
label_19924: LD (IY - 128),$12
label_19925: LD A,(BC)
label_19926: LD A,(DE)
label_19927: LD A,($5678)
label_19928: LD (BC),A
label_19929: LD (DE),A
label_19930: LD ($5678),A
label_19931: LD A,I
label_19932: LD A,R
label_19933: LD I,A
label_19934: LD R,A
label_19935: LD BC,$5678
label_19936: LD DE,$5678
label_19937: LD HL,$5678
label_19938: LD SP,$5678
label_19939: LD IX,$5678
label_19940: LD IY,$5678
label_19941: LD HL,($5678)
label_19942: LD BC,($5678)
label_19943: LD DE,($5678)
label_19944: LD HL,($5678)
label_19945: LD SP,($5678)
label_19946: LD IX,($5678)
label_19947: LD IY,($5678)
label_19948: LD ($5678),HL
label_19949: LD ($5678),BC
label_19950: LD ($5678),DE
label_19951: LD ($5678),HL
label_19952: LD ($5678),SP
label_19953: LD ($5678),IX
label_19954: LD ($5678),IY
label_19955: LD SP,HL
label_19956: LD SP,IX
label_19957: LD SP,IY
label_19958: PUSH BC
label_19959: PUSH DE
label_19960: PUSH HL
label_19961: PUSH AF
label_19962: PUSH IX
label_19963: PUSH IY
label_19964: POP BC
label_19965: POP DE
label_19966: POP HL
label_19967: POP AF
label_19968: POP IX
label_19969: POP IY
label_19970: EX DE,HL
label_19971: EX AF,AF'
label_19972: EXX
label_19973: EX (SP),HL
label_19974: EX (SP),IX
label_19975: EX (SP),IY
label_19976: LDI
label_19977: LDIR
label_19978: LDD
label_19979: LDDR
label_19980: CPI
label_19981: CPIR
label_19982: CPD
label_19983: CPDR
label_19984: ADD A,A
label_19985: ADD A,B
label_19986: ADD A,C
label_19987: ADD A,D
label_19988: ADD A,E
label_19989: ADD A,H
label_19990: ADD A,L
label_19991: ADD A,$12
label_19992: ADD A,(HL)
label_19993: ADD A,(IX + 127)
label_19994: ADD A,(IY - 128)
label_19995: ADC A,A
label_19996: ADC A,B
label_19997: ADC A,C
label_19998: ADC A,D
label_19999: ADC A,E
label_20000: ADC A,H
label_20001: ADC A,L
label_20002: ADC A,$12
label_20003: ADC A,(HL)
label_20004: ADC A,(IX + 127)
label_20005: ADC A,(IY - 128)
label_20006: SUB A
label_20007: SUB B
label_20008: SUB C
label_20009: SUB D
label_20010: SUB E
label_20011: SUB H
label_20012: SUB L
label_20013: SUB $12
label_20014: SUB (HL)
label_20015: SUB (IX + 127)
label_20016: SUB (IY - 128)
label_20017: SBC A,A
label_20018: SBC A,B
label_20019: SBC A,C
label_20020: SBC A,D
label_20021: SBC A,E
label_20022: SBC A,H
label_20023: SBC A,L
label_20024: SBC A,$12
label_20025: SBC A,(HL)
label_20026: SBC A,(IX + 127)
label_20027: SBC A,(IY - 128)
label_20028: AND A
label_20029: AND B
label_20030: AND C
label_20031: AND D
label_20032: AND E
label_20033: AND H
label_20034: AND L
label_20035: AND $12
label_20036: AND (HL)
label_20037: AND (IX + 127)
label_20038: AND (IY - 128)
label_20039: AND A
label_20040: AND B
label_20041: AND C
label_20042: AND D
label_20043: AND E
label_20044: AND H
label_20045: AND L
label_20046: AND $12
label_20047: AND (HL)
label_20048: AND (IX + 127)
label_20049: AND (IY - 128)
label_20050: OR A
label_20051: OR B
label_20052: OR C
label_20053: OR D
label_20054: OR E
label_20055: OR H
label_20056: OR L
label_20057: OR $12
label_20058: OR (HL)
label_20059: OR (IX + 127)
label_20060: OR (IY - 128)
label_20061: XOR A
label_20062: XOR B
label_20063: XOR C
label_20064: XOR D
label_20065: XOR E
label_20066: XOR H
label_20067: XOR L
label_20068: XOR $12
label_20069: XOR (HL)
label_20070: XOR (IX + 127)
label_20071: XOR (IY - 128)
label_20072: CP A
label_20073: CP B
label_20074: CP C
label_20075: CP D
label_20076: CP E
label_20077: CP H
label_20078: CP L
label_20079: CP $12
label_20080: CP (HL)
label_20081: CP (IX + 127)
label_20082: CP (IY - 128)
label_20083: INC A
label_20084: INC B
label_20085: INC C
label_20086: INC D
label_20087: INC E
label_20088: INC H
label_20089: INC L
label_20090: INC (HL)
label_20091: INC (IX + 127)
label_20092: INC (IY - 128)
label_20093: DEC A
label_20094: DEC B
label_20095: DEC C
label_20096: DEC D
label_20097: DEC E
label_20098: DEC H
label_20099: DEC L
label_20100: DEC (HL)
label_20101: DEC (IX + 127)
label_20102: DEC (IY - 128)
label_20103: DAA
label_20104: CPL
label_20105: NEG
label_20106: CCF
label_20107: SCF
label_20108: NOP
label_20109: HALT
label_20110: DI
label_20111: EI
label_20112: IM 0
label_20113: IM 1
label_20114: IM 2
label_20115: ADD HL,BC
label_20116: ADD HL,DE
label_20117: ADD HL,HL
label_20118: ADD HL,SP
label_20119: ADC HL,BC
label_20120: ADC HL,DE
label_20121: ADC HL,HL
label_20122: ADC HL,SP
label_20123: SBC HL,BC
label_20124: SBC HL,DE
label_20125: SBC HL,HL
label_20126: SBC HL,SP
label_20127: ADD IX,BC
label_20128: ADD IX,DE
label_20129: ADD IX,SP
label_20130: ADD IY,BC
label_20131: ADD IY,DE
label_20132: ADD IY,SP
label_20133: INC BC
label_20134: INC DE
label_20135: INC HL
label_20136: INC SP
label_20137: INC IX
label_20138: INC IY
label_20139: DEC BC
label_20140: DEC DE
label_20141: DEC HL
label_20142: DEC SP
label_20143: DEC IX
label_20144: DEC IY
label_20145: RLCA
label_20146: RLA
label_20147: RRCA
label_20148: RRA
label_20149: RLC A
label_20150: RLC B
label_20151: RLC C
label_20152: RLC D
label_20153: RLC E
label_20154: RLC H
label_20155: RLC L
label_20156: RLC (HL)
label_20157: RLC (IX + 127)
label_20158: RLC (IY - 128)
label_20159: RL A
label_20160: RL B
label_20161: RL C
label_20162: RL D
label_20163: RL E
label_20164: RL H
label_20165: RL L
label_20166: RL (HL)
label_20167: RL (IX + 127)
label_20168: RL (IY - 128)
label_20169: RRC A
label_20170: RRC B
label_20171: RRC C
label_20172: RRC D
label_20173: RRC E
label_20174: RRC H
label_20175: RRC L
label_20176: RRC (HL)
label_20177: RRC (IX + 127)
label_20178: RRC (IY - 128)
label_20179: RR A
label_20180: RR B
label_20181: RR C
label_20182: RR D
label_20183: RR E
label_20184: RR H
label_20185: RR L
label_20186: RR (HL)
label_20187: RR (IX + 127)
label_20188: RR (IY - 128)
label_20189: SLA A
label_20190: SLA B
label_20191: SLA C
label_20192: SLA D
label_20193: SLA E
label_20194: SLA H
label_20195: SLA L
label_20196: SLA (HL)
label_20197: SLA (IX + 127)
label_20198: SLA (IY - 128)
label_20199: SRA A
label_20200: SRA B
label_20201: SRA C
label_20202: SRA D
label_20203: SRA E
label_20204: SRA H
label_20205: SRA L
label_20206: SRA (HL)
label_20207: SRA (IX + 127)
label_20208: SRA (IY - 128)
label_20209: SRL A
label_20210: SRL B
label_20211: SRL C
label_20212: SRL D
label_20213: SRL E
label_20214: SRL H
label_20215: SRL L
label_20216: SRL (HL)
label_20217: SRL (IX + 127)
label_20218: SRL (IY - 128)
label_20219: RLD
label_20220: RRD
label_20221: BIT 0,A
label_20222: BIT 1,A
label_20223: BIT 2,A
label_20224: BIT 3,A
label_20225: BIT 4,A
label_20226: BIT 5,A
label_20227: BIT 6,A
label_20228: BIT 7,A
label_20229: BIT 0,B
label_20230: BIT 1,B
label_20231: BIT 2,B
label_20232: BIT 3,B
label_20233: BIT 4,B
label_20234: BIT 5,B
label_20235: BIT 6,B
label_20236: BIT 7,B
label_20237: BIT 0,C
label_20238: BIT 1,C
label_20239: BIT 2,C
label_20240: BIT 3,C
label_20241: BIT 4,C
label_20242: BIT 5,C
label_20243: BIT 6,C
label_20244: BIT 7,C
label_20245: BIT 0,D
label_20246: BIT 1,D
label_20247: BIT 2,D
label_20248: BIT 3,D
label_20249: BIT 4,D
label_20250: BIT 5,D
label_20251: BIT 6,D
label_20252: BIT 7,D
label_20253: BIT 0,E
label_20254: BIT 1,E
label_20255: BIT 2,E
label_20256: BIT 3,E
label_20257: BIT 4,E
label_20258: BIT 5,E
label_20259: BIT 6,E
label_20260: BIT 7,E
label_20261: BIT 0,H
label_20262: BIT 1,H
label_20263: BIT 2,H
label_20264: BIT 3,H
label_20265: BIT 4,H
label_20266: BIT 5,H
label_20267: BIT 6,H
label_20268: BIT 7,H
label_20269: BIT 0,L
label_20270: BIT 1,L
label_20271: BIT 2,L
label_20272: BIT 3,L
label_20273: BIT 4,L
label_20274: BIT 5,L
label_20275: BIT 6,L
label_20276: BIT 7,L
label_20277: BIT 0,(HL)
label_20278: BIT 1,(HL)
label_20279: BIT 2,(HL)
label_20280: BIT 3,(HL)
label_20281: BIT 4,(HL)
label_20282: BIT 5,(HL)
label_20283: BIT 6,(HL)
label_20284: BIT 7,(HL)
label_20285: BIT 0,(IX + 127)
label_20286: BIT 1,(IX + 127)
label_20287: BIT 2,(IX + 127)
label_20288: BIT 3,(IX + 127)
label_20289: BIT 4,(IX + 127)
label_20290: BIT 5,(IX + 127)
label_20291: BIT 6,(IX + 127)
label_20292: BIT 7,(IX + 127)
label_20293: BIT 0,(IY - 128)
label_20294: BIT 1,(IY - 128)
label_20295: BIT 2,(IY - 128)
label_20296: BIT 3,(IY - 128)
label_20297: BIT 4,(IY - 128)
label_20298: BIT 5,(IY - 128)
label_20299: BIT 6,(IY - 128)
label_20300: BIT 7,(IY - 128)
label_20301: SET 0,A
label_20302: SET 1,A
label_20303: SET 2,A
label_20304: SET 3,A
label_20305: SET 4,A
label_20306: SET 5,A
label_20307: SET 6,A
label_20308: SET 7,A
label_20309: SET 0,B
label_20310: SET 1,B
label_20311: SET 2,B
label_20312: SET 3,B
label_20313: SET 4,B
label_20314: SET 5,B
label_20315: SET 6,B
label_20316: SET 7,B
label_20317: SET 0,C
label_20318: SET 1,C
label_20319: SET 2,C
label_20320: SET 3,C
label_20321: SET 4,C
label_20322: SET 5,C
label_20323: SET 6,C
label_20324: SET 7,C
label_20325: SET 0,D
label_20326: SET 1,D
label_20327: SET 2,D
label_20328: SET 3,D
label_20329: SET 4,D
label_20330: SET 5,D
label_20331: SET 6,D
label_20332: SET 7,D
label_20333: SET 0,E
label_20334: SET 1,E
label_20335: SET 2,E
label_20336: SET 3,E
label_20337: SET 4,E
label_20338: SET 5,E
label_20339: SET 6,E
label_20340: SET 7,E
label_20341: SET 0,H
label_20342: SET 1,H
label_20343: SET 2,H
label_20344: SET 3,H
label_20345: SET 4,H
label_20346: SET 5,H
label_20347: SET 6,H
label_20348: SET 7,H
label_20349: SET 0,L
label_20350: SET 1,L
label_20351: SET 2,L
label_20352: SET 3,L
label_20353: SET 4,L
label_20354: SET 5,L
label_20355: SET 6,L
label_20356: SET 7,L
label_20357: SET 0,(HL)
label_20358: SET 1,(HL)
label_20359: SET 2,(HL)
label_20360: SET 3,(HL)
label_20361: SET 4,(HL)
label_20362: SET 5,(HL)
label_20363: SET 6,(HL)
label_20364: SET 7,(HL)
label_20365: SET 0,(IX + 127)
label_20366: SET 1,(IX + 127)
label_20367: SET 2,(IX + 127)
label_20368: SET 3,(IX + 127)
label_20369: SET 4,(IX + 127)
label_20370: SET 5,(IX + 127)
label_20371: SET 6,(IX + 127)
label_20372: SET 7,(IX + 127)
label_20373: SET 0,(IY - 128)
label_20374: SET 1,(IY - 128)
label_20375: SET 2,(IY - 128)
label_20376: SET 3,(IY - 128)
label_20377: SET 4,(IY - 128)
label_20378: SET 5,(IY - 128)
label_20379: SET 6,(IY - 128)
label_20380: SET 7,(IY - 128)
label_20381: RES 0,A
label_20382: RES 1,A
label_20383: RES 2,A
label_20384: RES 3,A
label_20385: RES 4,A
label_20386: RES 5,A
label_20387: RES 6,A
label_20388: RES 7,A
label_20389: RES 0,B
label_20390: RES 1,B
label_20391: RES 2,B
label_20392: RES 3,B
label_20393: RES 4,B
label_20394: RES 5,B
label_20395: RES 6,B
label_20396: RES 7,B
label_20397: RES 0,C
label_20398: RES 1,C
label_20399: RES 2,C
label_20400: RES 3,C
label_20401: RES 4,C
label_20402: RES 5,C
label_20403: RES 6,C
label_20404: RES 7,C
label_20405: RES 0,D
label_20406: RES 1,D
label_20407: RES 2,D
label_20408: RES 3,D
label_20409: RES 4,D
label_20410: RES 5,D
label_20411: RES 6,D
label_20412: RES 7,D
label_20413: RES 0,E
label_20414: RES 1,E
label_20415: RES 2,E
label_20416: RES 3,E
label_20417: RES 4,E
label_20418: RES 5,E
label_20419: RES 6,E
label_20420: RES 7,E
label_20421: RES 0,H
label_20422: RES 1,H
label_20423: RES 2,H
label_20424: RES 3,H
label_20425: RES 4,H
label_20426: RES 5,H
label_20427: RES 6,H
label_20428: RES 7,H
label_20429: RES 0,L
label_20430: RES 1,L
label_20431: RES 2,L
label_20432: RES 3,L
label_20433: RES 4,L
label_20434: RES 5,L
label_20435: RES 6,L
label_20436: RES 7,L
label_20437: RES 0,(HL)
label_20438: RES 1,(HL)
label_20439: RES 2,(HL)
label_20440: RES 3,(HL)
label_20441: RES 4,(HL)
label_20442: RES 5,(HL)
label_20443: RES 6,(HL)
label_20444: RES 7,(HL)
label_20445: RES 0,(IX + 127)
label_20446: RES 1,(IX + 127)
label_20447: RES 2,(IX + 127)
label_20448: RES 3,(IX + 127)
label_20449: RES 4,(IX + 127)
label_20450: RES 5,(IX + 127)
label_20451: RES 6,(IX + 127)
label_20452: RES 7,(IX + 127)
label_20453: RES 0,(IY - 128)
label_20454: RES 1,(IY - 128)
label_20455: RES 2,(IY - 128)
label_20456: RES 3,(IY - 128)
label_20457: RES 4,(IY - 128)
label_20458: RES 5,(IY - 128)
label_20459: RES 6,(IY - 128)
label_20460: RES 7,(IY - 128)
label_20461: JP $5678
label_20462: JP NZ,$5678
label_20463: JP Z,$5678
label_20464: JP NC,$5678
label_20465: JP C,$5678
label_20466: JP PO,$5678
label_20467: JP PE,$5678
label_20468: JP P,$5678
label_20469: JP M,$5678
label_20470: JR $ + 2
label_20471: JR NZ,$ + 2
label_20472: JR Z,$ + 2
label_20473: JR NC,$ + 2
label_20474: JR C,$ + 2
label_20475: JP (HL)
label_20476: JP (IX)
label_20477: JP (IY)
label_20478: DJNZ $ + 2
label_20479: CALL $5678
label_20480: CALL NZ,$5678
label_20481: CALL Z,$5678
label_20482: CALL NC,$5678
label_20483: CALL C,$5678
label_20484: CALL PO,$5678
label_20485: CALL PE,$5678
label_20486: CALL P,$5678
label_20487: CALL M,$5678
label_20488: RET
label_20489: RET NZ
label_20490: RET Z
label_20491: RET NC
label_20492: RET C
label_20493: RET PO
label_20494: RET PE
label_20495: RET P
label_20496: RET M
label_20497: RETI
label_20498: RETN
label_20499: RST $00
label_20500: RST $08
label_20501: RST $10
label_20502: RST $18
label_20503: RST $20
label_20504: RST $28
label_20505: RST $30
label_20506: RST $38
label_20507: IN A,($12)
label_20508: IN A,(C)
label_20509: IN B,(C)
label_20510: IN C,(C)
label_20511: IN D,(C)
label_20512: IN E,(C)
label_20513: IN H,(C)
label_20514: IN L,(C)
label_20515: IN F,(C)
label_20516: INI
label_20517: INIR
label_20518: IND
label_20519: INDR
label_20520: OUT ($12),A
label_20521: OUT (C),A
label_20522: OUT (C),B
label_20523: OUT (C),C
label_20524: OUT (C),D
label_20525: OUT (C),E
label_20526: OUT (C),H
label_20527: OUT (C),L
label_20528: OUTI
label_20529: OTIR
label_20530: OUTD
label_20531: OTDR
label_20532: LD A,A
label_20533: LD A,B
label_20534: LD A,C
label_20535: LD A,D
label_20536: LD A,E
label_20537: LD A,H
label_20538: LD A,L
label_20539: LD B,A
label_20540: LD B,B
label_20541: LD B,C
label_20542: LD B,D
label_20543: LD B,E
label_20544: LD B,H
label_20545: LD B,L
label_20546: LD C,A
label_20547: LD C,B
label_20548: LD C,C
label_20549: LD C,D
label_20550: LD C,E
label_20551: LD C,H
label_20552: LD C,L
label_20553: LD D,A
label_20554: LD D,B
label_20555: LD D,C
label_20556: LD D,D
label_20557: LD D,E
label_20558: LD D,H
label_20559: LD D,L
label_20560: LD E,A
label_20561: LD E,B
label_20562: LD E,C
label_20563: LD E,D
label_20564: LD E,E
label_20565: LD E,H
label_20566: LD E,L
label_20567: LD H,A
label_20568: LD H,B
label_20569: LD H,C
label_20570: LD H,D
label_20571: LD H,E
label_20572: LD H,H
label_20573: LD H,L
label_20574: LD L,A
label_20575: LD L,B
label_20576: LD L,C
label_20577: LD L,D
label_20578: LD L,E
label_20579: LD L,H
label_20580: LD L,L
label_20581: LD A,$12
label_20582: LD B,$12
label_20583: LD C,$12
label_20584: LD D,$12
label_20585: LD E,$12
label_20586: LD H,$12
label_20587: LD L,$12
label_20588: LD A,(HL)
label_20589: LD B,(HL)
label_20590: LD C,(HL)
label_20591: LD D,(HL)
label_20592: LD E,(HL)
label_20593: LD H,(HL)
label_20594: LD L,(HL)
label_20595: LD A,(IX + 127)
label_20596: LD B,(IX + 127)
label_20597: LD C,(IX + 127)
label_20598: LD D,(IX + 127)
label_20599: LD E,(IX + 127)
label_20600: LD H,(IX + 127)
label_20601: LD L,(IX + 127)
label_20602: LD A,(IY - 128)
label_20603: LD B,(IY - 128)
label_20604: LD C,(IY - 128)
label_20605: LD D,(IY - 128)
label_20606: LD E,(IY - 128)
label_20607: LD H,(IY - 128)
label_20608: LD L,(IY - 128)
label_20609: LD (HL),A
label_20610: LD (HL),B
label_20611: LD (HL),C
label_20612: LD (HL),D
label_20613: LD (HL),E
label_20614: LD (HL),H
label_20615: LD (HL),L
label_20616: LD (IX + 127),A
label_20617: LD (IX + 127),B
label_20618: LD (IX + 127),C
label_20619: LD (IX + 127),D
label_20620: LD (IX + 127),E
label_20621: LD (IX + 127),H
label_20622: LD (IX + 127),L
label_20623: LD (IY - 128),A
label_20624: LD (IY - 128),B
label_20625: LD (IY - 128),C
label_20626: LD (IY - 128),D
label_20627: LD (IY - 128),E
label_20628: LD (IY - 128),H
label_20629: LD (IY - 128),L
label_20630: LD (HL),$12
label_20631: LD (IX + 127),$12
label_20632: LD (IY - 128),$12
label_20633: LD A,(BC)
label_20634: LD A,(DE)
label_20635: LD A,($5678)
label_20636: LD (BC),A
label_20637: LD (DE),A
label_20638: LD ($5678),A
label_20639: LD A,I
label_20640: LD A,R
label_20641: LD I,A
label_20642: LD R,A
label_20643: LD BC,$5678
label_20644: LD DE,$5678
label_20645: LD HL,$5678
label_20646: LD SP,$5678
label_20647: LD IX,$5678
label_20648: LD IY,$5678
label_20649: LD HL,($5678)
label_20650: LD BC,($5678)
label_20651: LD DE,($5678)
label_20652: LD HL,($5678)
label_20653: LD SP,($5678)
label_20654: LD IX,($5678)
label_20655: LD IY,($5678)
label_20656: LD ($5678),HL
label_20657: LD ($5678),BC
label_20658: LD ($5678),DE
label_20659: LD ($5678),HL
label_20660: LD ($5678),SP
label_20661: LD ($5678),IX
label_20662: LD ($5678),IY
label_20663: LD SP,HL
label_20664: LD SP,IX
label_20665: LD SP,IY
label_20666: PUSH BC
label_20667: PUSH DE
label_20668: PUSH HL
label_20669: PUSH AF
label_20670: PUSH IX
label_20671: PUSH IY
label_20672: POP BC
label_20673: POP DE
label_20674: POP HL
label_20675: POP AF
label_20676: POP IX
label_20677: POP IY
label_20678: EX DE,HL
label_20679: EX AF,AF'
label_20680: EXX
label_20681: EX (SP),HL
label_20682: EX (SP),IX
label_20683: EX (SP),IY
label_20684: LDI
label_20685: LDIR
label_20686: LDD
label_20687: LDDR
label_20688: CPI
label_20689: CPIR
label_20690: CPD
label_20691: CPDR
label_20692: ADD A,A
label_20693: ADD A,B
label_20694: ADD A,C
label_20695: ADD A,D
label_20696: ADD A,E
label_20697: ADD A,H
label_20698: ADD A,L
label_20699: ADD A,$12
label_20700: ADD A,(HL)
label_20701: ADD A,(IX + 127)
label_20702: ADD A,(IY - 128)
label_20703: ADC A,A
label_20704: ADC A,B
label_20705: ADC A,C
label_20706: ADC A,D
label_20707: ADC A,E
label_20708: ADC A,H
label_20709: ADC A,L
label_20710: ADC A,$12
label_20711: ADC A,(HL)
label_20712: ADC A,(IX + 127)
label_20713: ADC A,(IY - 128)
label_20714: SUB A
label_20715: SUB B
label_20716: SUB C
label_20717: SUB D
label_20718: SUB E
label_20719: SUB H
label_20720: SUB L
label_20721: SUB $12
label_20722: SUB (HL)
label_20723: SUB (IX + 127)
label_20724: SUB (IY - 128)
label_20725: SBC A,A
label_20726: SBC A,B
label_20727: SBC A,C
label_20728: SBC A,D
label_20729: SBC A,E
label_20730: SBC A,H
label_20731: SBC A,L
label_20732: SBC A,$12
label_20733: SBC A,(HL)
label_20734: SBC A,(IX + 127)
label_20735: SBC A,(IY - 128)
label_20736: AND A
label_20737: AND B
label_20738: AND C
label_20739: AND D
label_20740: AND E
label_20741: AND H
label_20742: AND L
label_20743: AND $12
label_20744: AND (HL)
label_20745: AND (IX + 127)
label_20746: AND (IY - 128)
label_20747: AND A
label_20748: AND B
label_20749: AND C
label_20750: AND D
label_20751: AND E
label_20752: AND H
label_20753: AND L
label_20754: AND $12
label_20755: AND (HL)
label_20756: AND (IX + 127)
label_20757: AND (IY - 128)
label_20758: OR A
label_20759: OR B
label_20760: OR C
label_20761: OR D
label_20762: OR E
label_20763: OR H
label_20764: OR L
label_20765: OR $12
label_20766: OR (HL)
label_20767: OR (IX + 127)
label_20768: OR (IY - 128)
label_20769: XOR A
label_20770: XOR B
label_20771: XOR C
label_20772: XOR D
label_20773: XOR E
label_20774: XOR H
label_20775: XOR L
label_20776: XOR $12
label_20777: XOR (HL)
label_20778: XOR (IX + 127)
label_20779: XOR (IY - 128)
label_20780: CP A
label_20781: CP B
label_20782: CP C
label_20783: CP D
label_20784: CP E
label_20785: CP H
label_20786: CP L
label_20787: CP $12
label_20788: CP (HL)
label_20789: CP (IX + 127)
label_20790: CP (IY - 128)
label_20791: INC A
label_20792: INC B
label_20793: INC C
label_20794: INC D
label_20795: INC E
label_20796: INC H
label_20797: INC L
label_20798: INC (HL)
label_20799: INC (IX + 127)
label_20800: INC (IY - 128)
label_20801: DEC A
label_20802: DEC B
label_20803: DEC C
label_20804: DEC D
label_20805: DEC E
label_20806: DEC H
label_20807: DEC L
label_20808: DEC (HL)
label_20809: DEC (IX + 127)
label_20810: DEC (IY - 128)
label_20811: DAA
label_20812: CPL
label_20813: NEG
label_20814: CCF
label_20815: SCF
label_20816: NOP
label_20817: HALT
label_20818: DI
label_20819: EI
label_20820: IM 0
label_20821: IM 1
label_20822: IM 2
label_20823: ADD HL,BC
label_20824: ADD HL,DE
label_20825: ADD HL,HL
label_20826: ADD HL,SP
label_20827: ADC HL,BC
label_20828: ADC HL,DE
label_20829: ADC HL,HL
label_20830: ADC HL,SP
label_20831: SBC HL,BC
label_20832: SBC HL,DE
label_20833: SBC HL,HL
label_20834: SBC HL,SP
label_20835: ADD IX,BC
label_20836: ADD IX,DE
label_20837: ADD IX,SP
label_20838: ADD IY,BC
label_20839: ADD IY,DE
label_20840: ADD IY,SP
label_20841: INC BC
label_20842: INC DE
label_20843: INC HL
label_20844: INC SP
label_20845: INC IX
label_20846: INC IY
label_20847: DEC BC
label_20848: DEC DE
label_20849: DEC HL
label_20850: DEC SP
label_20851: DEC IX
label_20852: DEC IY
label_20853: RLCA
label_20854: RLA
label_20855: RRCA
label_20856: RRA
label_20857: RLC A
label_20858: RLC B
label_20859: RLC C
label_20860: RLC D
label_20861: RLC E
label_20862: RLC H
label_20863: RLC L
label_20864: RLC (HL)
label_20865: RLC (IX + 127)
label_20866: RLC (IY - 128)
label_20867: RL A
label_20868: RL B
label_20869: RL C
label_20870: RL D
label_20871: RL E
label_20872: RL H
label_20873: RL L
label_20874: RL (HL)
label_20875: RL (IX + 127)
label_20876: RL (IY - 128)
label_20877: RRC A
label_20878: RRC B
label_20879: RRC C
label_20880: RRC D
label_20881: RRC E
label_20882: RRC H
label_20883: RRC L
label_20884: RRC (HL)
label_20885: RRC (IX + 127)
label_20886: RRC (IY - 128)
label_20887: RR A
label_20888: RR B
label_20889: RR C
label_20890: RR D
label_20891: RR E
label_20892: RR H
label_20893: RR L
label_20894: RR (HL)
label_20895: RR (IX + 127)
label_20896: RR (IY - 128)
label_20897: SLA A
label_20898: SLA B
label_20899: SLA C
label_20900: SLA D
label_20901: SLA E
label_20902: SLA H
label_20903: SLA L
label_20904: SLA (HL)
label_20905: SLA (IX + 127)
label_20906: SLA (IY - 128)
label_20907: SRA A
label_20908: SRA B
label_20909: SRA C
label_20910: SRA D
label_20911: SRA E
label_20912: SRA H
label_20913: SRA L
label_20914: SRA (HL)
label_20915: SRA (IX + 127)
label_20916: SRA (IY - 128)
label_20917: SRL A
label_20918: SRL B
label_20919: SRL C
label_20920: SRL D
label_20921: SRL E
label_20922: SRL H
label_20923: SRL L
label_20924: SRL (HL)
label_20925: SRL (IX + 127)
label_20926: SRL (IY - 128)
label_20927: RLD
label_20928: RRD
label_20929: BIT 0,A
label_20930: BIT 1,A
label_20931: BIT 2,A
label_20932: BIT 3,A
label_20933: BIT 4,A
label_20934: BIT 5,A
label_20935: BIT 6,A
label_20936: BIT 7,A
label_20937: BIT 0,B
label_20938: BIT 1,B
label_20939: BIT 2,B
label_20940: BIT 3,B
label_20941: BIT 4,B
label_20942: BIT 5,B
label_20943: BIT 6,B
label_20944: BIT 7,B
label_20945: BIT 0,C
label_20946: BIT 1,C
label_20947: BIT 2,C
label_20948: BIT 3,C
label_20949: BIT 4,C
label_20950: BIT 5,C
label_20951: BIT 6,C
label_20952: BIT 7,C
label_20953: BIT 0,D
label_20954: BIT 1,D
label_20955: BIT 2,D
label_20956: BIT 3,D
label_20957: BIT 4,D
label_20958: BIT 5,D
label_20959: BIT 6,D
label_20960: BIT 7,D
label_20961: BIT 0,E
label_20962: BIT 1,E
label_20963: BIT 2,E
label_20964: BIT 3,E
label_20965: BIT 4,E
label_20966: BIT 5,E
label_20967: BIT 6,E
label_20968: BIT 7,E
label_20969: BIT 0,H
label_20970: BIT 1,H
label_20971: BIT 2,H
label_20972: BIT 3,H
label_20973: BIT 4,H
label_20974: BIT 5,H
label_20975: BIT 6,H
label_20976: BIT 7,H
label_20977: BIT 0,L
label_20978: BIT 1,L
label_20979: BIT 2,L
label_20980: BIT 3,L
label_20981: BIT 4,L
label_20982: BIT 5,L
label_20983: BIT 6,L
label_20984: BIT 7,L
label_20985: BIT 0,(HL)
label_20986: BIT 1,(HL)
label_20987: BIT 2,(HL)
label_20988: BIT 3,(HL)
label_20989: BIT 4,(HL)
label_20990: BIT 5,(HL)
label_20991: BIT 6,(HL)
label_20992: BIT 7,(HL)
label_20993: BIT 0,(IX + 127)
label_20994: BIT 1,(IX + 127)
label_20995: BIT 2,(IX + 127)
label_20996: BIT 3,(IX + 127)
label_20997: BIT 4,(IX + 127)
label_20998: BIT 5,(IX + 127)
label_20999: BIT 6,(IX + 127)
label_21000: BIT 7,(IX + 127)
label_21001: BIT 0,(IY - 128)
label_21002: BIT 1,(IY - 128)
label_21003: BIT 2,(IY - 128)
label_21004: BIT 3,(IY - 128)
label_21005: BIT 4,(IY - 128)
label_21006: BIT 5,(IY - 128)
label_21007: BIT 6,(IY - 128)
label_21008: BIT 7,(IY - 128)
label_21009: SET 0,A
label_21010: SET 1,A
label_21011: SET 2,A
label_21012: SET 3,A
label_21013: SET 4,A
label_21014: SET 5,A
label_21015: SET 6,A
label_21016: SET 7,A
label_21017: SET 0,B
label_21018: SET 1,B
label_21019: SET 2,B
label_21020: SET 3,B
label_21021: SET 4,B
label_21022: SET 5,B
label_21023: SET 6,B
label_21024: SET 7,B
label_21025: SET 0,C
label_21026: SET 1,C
label_21027: SET 2,C
label_21028: SET 3,C
label_21029: SET 4,C
label_21030: SET 5,C
label_21031: SET 6,C
label_21032: SET 7,C
label_21033: SET 0,D
label_21034: SET 1,D
label_21035: SET 2,D
label_21036: SET 3,D
label_21037: SET 4,D
label_21038: SET 5,D
label_21039: SET 6,D
label_21040: SET 7,D
label_21041: SET 0,E
label_21042: SET 1,E
label_21043: SET 2,E
label_21044: SET 3,E
label_21045: SET 4,E
label_21046: SET 5,E
label_21047: SET 6,E
label_21048: SET 7,E
label_21049: SET 0,H
label_21050: SET 1,H
label_21051: SET 2,H
label_21052: SET 3,H
label_21053: SET 4,H
label_21054: SET 5,H
label_21055: SET 6,H
label_21056: SET 7,H
label_21057: SET 0,L
label_21058: SET 1,L
label_21059: SET 2,L
label_21060: SET 3,L
label_21061: SET 4,L
label_21062: SET 5,L
label_21063: SET 6,L
label_21064: SET 7,L
label_21065: SET 0,(HL)
label_21066: SET 1,(HL)
label_21067: SET 2,(HL)
label_21068: SET 3,(HL)
label_21069: SET 4,(HL)
label_21070: SET 5,(HL)
label_21071: SET 6,(HL)
label_21072: SET 7,(HL)
label_21073: SET 0,(IX + 127)
label_21074: SET 1,(IX + 127)
label_21075: SET 2,(IX + 127)
label_21076: SET 3,(IX + 127)
label_21077: SET 4,(IX + 127)
label_21078: SET 5,(IX + 127)
label_21079: SET 6,(IX + 127)
label_21080: SET 7,(IX + 127)
label_21081: SET 0,(IY - 128)
label_21082: SET 1,(IY - 128)
label_21083: SET 2,(IY - 128)
label_21084: SET 3,(IY - 128)
label_21085: SET 4,(IY - 128)
label_21086: SET 5,(IY - 128)
label_21087: SET 6,(IY - 128)
label_21088: SET 7,(IY - 128)
label_21089: RES 0,A
label_21090: RES 1,A
label_21091: RES 2,A
label_21092: RES 3,A
label_21093: RES 4,A
label_21094: RES 5,A
label_21095: RES 6,A
label_21096: RES 7,A
label_21097: RES 0,B
label_21098: RES 1,B
label_21099: RES 2,B
label_21100: RES 3,B
label_21101: RES 4,B
label_21102: RES 5,B
label_21103: RES 6,B
label_21104: RES 7,B
label_21105: RES 0,C
label_21106: RES 1,C
label_21107: RES 2,C
label_21108: RES 3,C
label_21109: RES 4,C
label_21110: RES 5,C
label_21111: RES 6,C
label_21112: RES 7,C
label_21113: RES 0,D
label_21114: RES 1,D
label_21115: RES 2,D
label_21116: RES 3,D
label_21117: RES 4,D
label_21118: RES 5,D
label_21119: RES 6,D
label_21120: RES 7,D
label_21121: RES 0,E
label_21122: RES 1,E
label_21123: RES 2,E
label_21124: RES 3,E
label_21125: RES 4,E
label_21126: RES 5,E
label_21127: RES 6,E
label_21128: RES 7,E
label_21129: RES 0,H
label_21130: RES 1,H
label_21131: RES 2,H
label_21132: RES 3,H
label_21133: RES 4,H
label_21134: RES 5,H
label_21135: RES 6,H
label_21136: RES 7,H
label_21137: RES 0,L
label_21138: RES 1,L
label_21139: RES 2,L
label_21140: RES 3,L
label_21141: RES 4,L
label_21142: RES 5,L
label_21143: RES 6,L
label_21144: RES 7,L
label_21145: RES 0,(HL)
label_21146: RES 1,(HL)
label_21147: RES 2,(HL)
label_21148: RES 3,(HL)
label_21149: RES 4,(HL)
label_21150: RES 5,(HL)
label_21151: RES 6,(HL)
label_21152: RES 7,(HL)
label_21153: RES 0,(IX + 127)
label_21154: RES 1,(IX + 127)
label_21155: RES 2,(IX + 127)
label_21156: RES 3,(IX + 127)
label_21157: RES 4,(IX + 127)
label_21158: RES 5,(IX + 127)
label_21159: RES 6,(IX + 127)
label_21160: RES 7,(IX + 127)
label_21161: RES 0,(IY - 128)
label_21162: RES 1,(IY - 128)
label_21163: RES 2,(IY - 128)
label_21164: RES 3,(IY - 128)
label_21165: RES 4,(IY - 128)
label_21166: RES 5,(IY - 128)
label_21167: RES 6,(IY - 128)
label_21168: RES 7,(IY - 128)
label_21169: JP $5678
label_21170: JP NZ,$5678
label_21171: JP Z,$5678
label_21172: JP NC,$5678
label_21173: JP C,$5678
label_21174: JP PO,$5678
label_21175: JP PE,$5678
label_21176: JP P,$5678
label_21177: JP M,$5678
label_21178: JR $ + 2
label_21179: JR NZ,$ + 2
label_21180: JR Z,$ + 2
label_21181: JR NC,$ + 2
label_21182: JR C,$ + 2
label_21183: JP (HL)
label_21184: JP (IX)
label_21185: JP (IY)
label_21186: DJNZ $ + 2
label_21187: CALL $5678
label_21188: CALL NZ,$5678
label_21189: CALL Z,$5678
label_21190: CALL NC,$5678
label_21191: CALL C,$5678
label_21192: CALL PO,$5678
label_21193: CALL PE,$5678
label_21194: CALL P,$5678
label_21195: CALL M,$5678
label_21196: RET
label_21197: RET NZ
label_21198: RET Z
label_21199: RET NC
label_21200: RET C
label_21201: RET PO
label_21202: RET PE
label_21203: RET P
label_21204: RET M
label_21205: RETI
label_21206: RETN
label_21207: RST $00
label_21208: RST $08
label_21209: RST $10
label_21210: RST $18
label_21211: RST $20
label_21212: RST $28
label_21213: RST $30
label_21214: RST $38
label_21215: IN A,($12)
label_21216: IN A,(C)
label_21217: IN B,(C)
label_21218: IN C,(C)
label_21219: IN D,(C)
label_21220: IN E,(C)
label_21221: IN H,(C)
label_21222: IN L,(C)
label_21223: IN F,(C)
label_21224: INI
label_21225: INIR
label_21226: IND
label_21227: INDR
label_21228: OUT ($12),A
label_21229: OUT (C),A
label_21230: OUT (C),B
label_21231: OUT (C),C
label_21232: OUT (C),D
label_21233: OUT (C),E
label_21234: OUT (C),H
label_21235: OUT (C),L
label_21236: OUTI
label_21237: OTIR
label_21238: OUTD
label_21239: OTDR
label_21240: LD A,A
label_21241: LD A,B
label_21242: LD A,C
label_21243: LD A,D
label_21244: LD A,E
label_21245: LD A,H
label_21246: LD A,L
label_21247: LD B,A
label_21248: LD B,B
label_21249: LD B,C
label_21250: LD B,D
label_21251: LD B,E
label_21252: LD B,H
label_21253: LD B,L
label_21254: LD C,A
label_21255: LD C,B
label_21256: LD C,C
label_21257: LD C,D
label_21258: LD C,E
label_21259: LD C,H
label_21260: LD C,L
label_21261: LD D,A
label_21262: LD D,B
label_21263: LD D,C
label_21264: LD D,D
label_21265: LD D,E
label_21266: LD D,H
label_21267: LD D,L
label_21268: LD E,A
label_21269: LD E,B
label_21270: LD E,C
label_21271: LD E,D
label_21272: LD E,E
label_21273: LD E,H
label_21274: LD E,L
label_21275: LD H,A
label_21276: LD H,B
label_21277: LD H,C
label_21278: LD H,D
label_21279: LD H,E
label_21280: LD H,H
label_21281: LD H,L
label_21282: LD L,A
label_21283: LD L,B
label_21284: LD L,C
label_21285: LD L,D
label_21286: LD L,E
label_21287: LD L,H
label_21288: LD L,L
label_21289: LD A,$12
label_21290: LD B,$12
label_21291: LD C,$12
label_21292: LD D,$12
label_21293: LD E,$12
label_21294: LD H,$12
label_21295: LD L,$12
label_21296: LD A,(HL)
label_21297: LD B,(HL)
label_21298: LD C,(HL)
label_21299: LD D,(HL)
label_21300: LD E,(HL)
label_21301: LD H,(HL)
label_21302: LD L,(HL)
label_21303: LD A,(IX + 127)
label_21304: LD B,(IX + 127)
label_21305: LD C,(IX + 127)
label_21306: LD D,(IX + 127)
label_21307: LD E,(IX + 127)
label_21308: LD H,(IX + 127)
label_21309: LD L,(IX + 127)
label_21310: LD A,(IY - 128)
label_21311: LD B,(IY - 128)
label_21312: LD C,(IY - 128)
label_21313: LD D,(IY - 128)
label_21314: LD E,(IY - 128)
label_21315: LD H,(IY - 128)
label_21316: LD L,(IY - 128)
label_21317: LD (HL),A
label_21318: LD (HL),B
label_21319: LD (HL),C
label_21320: LD (HL),D
label_21321: LD (HL),E
label_21322: LD (HL),H
label_21323: LD (HL),L
label_21324: LD (IX + 127),A
label_21325: LD (IX + 127),B
label_21326: LD (IX + 127),C
label_21327: LD (IX + 127),D
label_21328: LD (IX + 127),E
label_21329: LD (IX + 127),H
label_21330: LD (IX + 127),L
label_21331: LD (IY - 128),A
label_21332: LD (IY - 128),B
label_21333: LD (IY - 128),C
label_21334: LD (IY - 128),D
label_21335: LD (IY - 128),E
label_21336: LD (IY - 128),H
label_21337: LD (IY - 128),L
label_21338: LD (HL),$12
label_21339: LD (IX + 127),$12
label_21340: LD (IY - 128),$12
label_21341: LD A,(BC)
label_21342: LD A,(DE)
label_21343: LD A,($5678)
label_21344: LD (BC),A
label_21345: LD (DE),A
label_21346: LD ($5678),A
label_21347: LD A,I
label_21348: LD A,R
label_21349: LD I,A
label_21350: LD R,A
label_21351: LD BC,$5678
label_21352: LD DE,$5678
label_21353: LD HL,$5678
label_21354: LD SP,$5678
label_21355: LD IX,$5678
label_21356: LD IY,$5678
label_21357: LD HL,($5678)
label_21358: LD BC,($5678)
label_21359: LD DE,($5678)
label_21360: LD HL,($5678)
label_21361: LD SP,($5678)
label_21362: LD IX,($5678)
label_21363: LD IY,($5678)
label_21364: LD ($5678),HL
label_21365: LD ($5678),BC
label_21366: LD ($5678),DE
label_21367: LD ($5678),HL
label_21368: LD ($5678),SP
label_21369: LD ($5678),IX
label_21370: LD ($5678),IY
label_21371: LD SP,HL
label_21372: LD SP,IX
label_21373: LD SP,IY
label_21374: PUSH BC
label_21375: PUSH DE
label_21376: PUSH HL
label_21377: PUSH AF
label_21378: PUSH IX
label_21379: PUSH IY
label_21380: POP BC
label_21381: POP DE
label_21382: POP HL
label_21383: POP AF
label_21384: POP IX
label_21385: POP IY
label_21386: EX DE,HL
label_21387: EX AF,AF'
label_21388: EXX
label_21389: EX (SP),HL
label_21390: EX (SP),IX
label_21391: EX (SP),IY
label_21392: LDI
label_21393: LDIR
label_21394: LDD
label_21395: LDDR
label_21396: CPI
label_21397: CPIR
label_21398: CPD
label_21399: CPDR
label_21400: ADD A,A
label_21401: ADD A,B
label_21402: ADD A,C
label_21403: ADD A,D
label_21404: ADD A,E
label_21405: ADD A,H
label_21406: ADD A,L
label_21407: ADD A,$12
label_21408: ADD A,(HL)
label_21409: ADD A,(IX + 127)
label_21410: ADD A,(IY - 128)
label_21411: ADC A,A
label_21412: ADC A,B
label_21413: ADC A,C
label_21414: ADC A,D
label_21415: ADC A,E
label_21416: ADC A,H
label_21417: ADC A,L
label_21418: ADC A,$12
label_21419: ADC A,(HL)
label_21420: ADC A,(IX + 127)
label_21421: ADC A,(IY - 128)
label_21422: SUB A
label_21423: SUB B
label_21424: SUB C
label_21425: SUB D
label_21426: SUB E
label_21427: SUB H
label_21428: SUB L
label_21429: SUB $12
label_21430: SUB (HL)
label_21431: SUB (IX + 127)
label_21432: SUB (IY - 128)
label_21433: SBC A,A
label_21434: SBC A,B
label_21435: SBC A,C
label_21436: SBC A,D
label_21437: SBC A,E
label_21438: SBC A,H
label_21439: SBC A,L
label_21440: SBC A,$12
label_21441: SBC A,(HL)
label_21442: SBC A,(IX + 127)
label_21443: SBC A,(IY - 128)
label_21444: AND A
label_21445: AND B
label_21446: AND C
label_21447: AND D
label_21448: AND E
label_21449: AND H
label_21450: AND L
label_21451: AND $12
label_21452: AND (HL)
label_21453: AND (IX + 127)
label_21454: AND (IY - 128)
label_21455: AND A
label_21456: AND B
label_21457: AND C
label_21458: AND D
label_21459: AND E
label_21460: AND H
label_21461: AND L
label_21462: AND $12
label_21463: AND (HL)
label_21464: AND (IX + 127)
label_21465: AND (IY - 128)
label_21466: OR A
label_21467: OR B
label_21468: OR C
label_21469: OR D
label_21470: OR E
label_21471: OR H
label_21472: OR L
label_21473: OR $12
label_21474: OR (HL)
label_21475: OR (IX + 127)
label_21476: OR (IY - 128)
label_21477: XOR A
label_21478: XOR B
label_21479: XOR C
label_21480: XOR D
label_21481: XOR E
label_21482: XOR H
label_21483: XOR L
label_21484: XOR $12
label_21485: XOR (HL)
label_21486: XOR (IX + 127)
label_21487: XOR (IY - 128)
label_21488: CP A
label_21489: CP B
label_21490: CP C
label_21491: CP D
label_21492: CP E
label_21493: CP H
label_21494: CP L
label_21495: CP $12
label_21496: CP (HL)
label_21497: CP (IX + 127)
label_21498: CP (IY - 128)
label_21499: INC A
label_21500: INC B
label_21501: INC C
label_21502: INC D
label_21503: INC E
label_21504: INC H
label_21505: INC L
label_21506: INC (HL)
label_21507: INC (IX + 127)
label_21508: INC (IY - 128)
label_21509: DEC A
label_21510: DEC B
label_21511: DEC C
label_21512: DEC D
label_21513: DEC E
label_21514: DEC H
label_21515: DEC L
label_21516: DEC (HL)
label_21517: DEC (IX + 127)
label_21518: DEC (IY - 128)
label_21519: DAA
label_21520: CPL
label_21521: NEG
label_21522: CCF
label_21523: SCF
label_21524: NOP
label_21525: HALT
label_21526: DI
label_21527: EI
label_21528: IM 0
label_21529: IM 1
label_21530: IM 2
label_21531: ADD HL,BC
label_21532: ADD HL,DE
label_21533: ADD HL,HL
label_21534: ADD HL,SP
label_21535: ADC HL,BC
label_21536: ADC HL,DE
label_21537: ADC HL,HL
label_21538: ADC HL,SP
label_21539: SBC HL,BC
label_21540: SBC HL,DE
label_21541: SBC HL,HL
label_21542: SBC HL,SP
label_21543: ADD IX,BC
label_21544: ADD IX,DE
label_21545: ADD IX,SP
label_21546: ADD IY,BC
label_21547: ADD IY,DE
label_21548: ADD IY,SP
label_21549: INC BC
label_21550: INC DE
label_21551: INC HL
label_21552: INC SP
label_21553: INC IX
label_21554: INC IY
label_21555: DEC BC
label_21556: DEC DE
label_21557: DEC HL
label_21558: DEC SP
label_21559: DEC IX
label_21560: DEC IY
label_21561: RLCA
label_21562: RLA
label_21563: RRCA
label_21564: RRA
label_21565: RLC A
label_21566: RLC B
label_21567: RLC C
label_21568: RLC D
label_21569: RLC E
label_21570: RLC H
label_21571: RLC L
label_21572: RLC (HL)
label_21573: RLC (IX + 127)
label_21574: RLC (IY - 128)
label_21575: RL A
label_21576: RL B
label_21577: RL C
label_21578: RL D
label_21579: RL E
label_21580: RL H
label_21581: RL L
label_21582: RL (HL)
label_21583: RL (IX + 127)
label_21584: RL (IY - 128)
label_21585: RRC A
label_21586: RRC B
label_21587: RRC C
label_21588: RRC D
label_21589: RRC E
label_21590: RRC H
label_21591: RRC L
label_21592: RRC (HL)
label_21593: RRC (IX + 127)
label_21594: RRC (IY - 128)
label_21595: RR A
label_21596: RR B
label_21597: RR C
label_21598: RR D
label_21599: RR E
label_21600: RR H
label_21601: RR L
label_21602: RR (HL)
label_21603: RR (IX + 127)
label_21604: RR (IY - 128)
label_21605: SLA A
label_21606: SLA B
label_21607: SLA C
label_21608: SLA D
label_21609: SLA E
label_21610: SLA H
label_21611: SLA L
label_21612: SLA (HL)
label_21613: SLA (IX + 127)
label_21614: SLA (IY - 128)
label_21615: SRA A
label_21616: SRA B
label_21617: SRA C
label_21618: SRA D
label_21619: SRA E
label_21620: SRA H
label_21621: SRA L
label_21622: SRA (HL)
label_21623: SRA (IX + 127)
label_21624: SRA (IY - 128)
label_21625: SRL A
label_21626: SRL B
label_21627: SRL C
label_21628: SRL D
label_21629: SRL E
label_21630: SRL H
label_21631: SRL L
label_21632: SRL (HL)
label_21633: SRL (IX + 127)
label_21634: SRL (IY - 128)
label_21635: RLD
label_21636: RRD
label_21637: BIT 0,A
label_21638: BIT 1,A
label_21639: BIT 2,A
label_21640: BIT 3,A
label_21641: BIT 4,A
label_21642: BIT 5,A
label_21643: BIT 6,A
label_21644: BIT 7,A
label_21645: BIT 0,B
label_21646: BIT 1,B
label_21647: BIT 2,B
label_21648: BIT 3,B
label_21649: BIT 4,B
label_21650: BIT 5,B
label_21651: BIT 6,B
label_21652: BIT 7,B
label_21653: BIT 0,C
label_21654: BIT 1,C
label_21655: BIT 2,C
label_21656: BIT 3,C
label_21657: BIT 4,C
label_21658: BIT 5,C
label_21659: BIT 6,C
label_21660: BIT 7,C
label_21661: BIT 0,D
label_21662: BIT 1,D
label_21663: BIT 2,D
label_21664: BIT 3,D
label_21665: BIT 4,D
label_21666: BIT 5,D
label_21667: BIT 6,D
label_21668: BIT 7,D
label_21669: BIT 0,E
label_21670: BIT 1,E
label_21671: BIT 2,E
label_21672: BIT 3,E
label_21673: BIT 4,E
label_21674: BIT 5,E
label_21675: BIT 6,E
label_21676: BIT 7,E
label_21677: BIT 0,H
label_21678: BIT 1,H
label_21679: BIT 2,H
label_21680: BIT 3,H
label_21681: BIT 4,H
label_21682: BIT 5,H
label_21683: BIT 6,H
label_21684: BIT 7,H
label_21685: BIT 0,L
label_21686: BIT 1,L
label_21687: BIT 2,L
label_21688: BIT 3,L
label_21689: BIT 4,L
label_21690: BIT 5,L
label_21691: BIT 6,L
label_21692: BIT 7,L
label_21693: BIT 0,(HL)
label_21694: BIT 1,(HL)
label_21695: BIT 2,(HL)
label_21696: BIT 3,(HL)
label_21697: BIT 4,(HL)
label_21698: BIT 5,(HL)
label_21699: BIT 6,(HL)
label_21700: BIT 7,(HL)
label_21701: BIT 0,(IX + 127)
label_21702: BIT 1,(IX + 127)
label_21703: BIT 2,(IX + 127)
label_21704: BIT 3,(IX + 127)
label_21705: BIT 4,(IX + 127)
label_21706: BIT 5,(IX + 127)
label_21707: BIT 6,(IX + 127)
label_21708: BIT 7,(IX + 127)
label_21709: BIT 0,(IY - 128)
label_21710: BIT 1,(IY - 128)
label_21711: BIT 2,(IY - 128)
label_21712: BIT 3,(IY - 128)
label_21713: BIT 4,(IY - 128)
label_21714: BIT 5,(IY - 128)
label_21715: BIT 6,(IY - 128)
label_21716: BIT 7,(IY - 128)
label_21717: SET 0,A
label_21718: SET 1,A
label_21719: SET 2,A
label_21720: SET 3,A
label_21721: SET 4,A
label_21722: SET 5,A
label_21723: SET 6,A
label_21724: SET 7,A
label_21725: SET 0,B
label_21726: SET 1,B
label_21727: SET 2,B
label_21728: SET 3,B
label_21729: SET 4,B
label_21730: SET 5,B
label_21731: SET 6,B
label_21732: SET 7,B
label_21733: SET 0,C
label_21734: SET 1,C
label_21735: SET 2,C
label_21736: SET 3,C
label_21737: SET 4,C
label_21738: SET 5,C
label_21739: SET 6,C
label_21740: SET 7,C
label_21741: SET 0,D
label_21742: SET 1,D
label_21743: SET 2,D
label_21744: SET 3,D
label_21745: SET 4,D
label_21746: SET 5,D
label_21747: SET 6,D
label_21748: SET 7,D
label_21749: SET 0,E
label_21750: SET 1,E
label_21751: SET 2,E
label_21752: SET 3,E
label_21753: SET 4,E
label_21754: SET 5,E
label_21755: SET 6,E
label_21756: SET 7,E
label_21757: SET 0,H
label_21758: SET 1,H
label_21759: SET 2,H
label_21760: SET 3,H
label_21761: SET 4,H
label_21762: SET 5,H
label_21763: SET 6,H
label_21764: SET 7,H
label_21765: SET 0,L
label_21766: SET 1,L
label_21767: SET 2,L
label_21768: SET 3,L
label_21769: SET 4,L
label_21770: SET 5,L
label_21771: SET 6,L
label_21772: SET 7,L
label_21773: SET 0,(HL)
label_21774: SET 1,(HL)
label_21775: SET 2,(HL)
label_21776: SET 3,(HL)
label_21777: SET 4,(HL)
label_21778: SET 5,(HL)
label_21779: SET 6,(HL)
label_21780: SET 7,(HL)
label_21781: SET 0,(IX + 127)
label_21782: SET 1,(IX + 127)
label_21783: SET 2,(IX + 127)
label_21784: SET 3,(IX + 127)
label_21785: SET 4,(IX + 127)
label_21786: SET 5,(IX + 127)
label_21787: SET 6,(IX + 127)
label_21788: SET 7,(IX + 127)
label_21789: SET 0,(IY - 128)
label_21790: SET 1,(IY - 128)
label_21791: SET 2,(IY - 128)
label_21792: SET 3,(IY - 128)
label_21793: SET 4,(IY - 128)
label_21794: SET 5,(IY - 128)
label_21795: SET 6,(IY - 128)
label_21796: SET 7,(IY - 128)
label_21797: RES 0,A
label_21798: RES 1,A
label_21799: RES 2,A
label_21800: RES 3,A
label_21801: RES 4,A
label_21802: RES 5,A
label_21803: RES 6,A
label_21804: RES 7,A
label_21805: RES 0,B
label_21806: RES 1,B
label_21807: RES 2,B
label_21808: RES 3,B
label_21809: RES 4,B
label_21810: RES 5,B
label_21811: RES 6,B
label_21812: RES 7,B
label_21813: RES 0,C
label_21814: RES 1,C
label_21815: RES 2,C
label_21816: RES 3,C
label_21817: RES 4,C
label_21818: RES 5,C
label_21819: RES 6,C
label_21820: RES 7,C
label_21821: RES 0,D
label_21822: RES 1,D
label_21823: RES 2,D
label_21824: RES 3,D
label_21825: RES 4,D
label_21826: RES 5,D
label_21827: RES 6,D
label_21828: RES 7,D
label_21829: RES 0,E
label_21830: RES 1,E
label_21831: RES 2,E
label_21832: RES 3,E
label_21833: RES 4,E
label_21834: RES 5,E
label_21835: RES 6,E
label_21836: RES 7,E
label_21837: RES 0,H
label_21838: RES 1,H
label_21839: RES 2,H
label_21840: RES 3,H
label_21841: RES 4,H
label_21842: RES 5,H
label_21843: RES 6,H
label_21844: RES 7,H
label_21845: RES 0,L
label_21846: RES 1,L
label_21847: RES 2,L
label_21848: RES 3,L
label_21849: RES 4,L
label_21850: RES 5,L
label_21851: RES 6,L
label_21852: RES 7,L
label_21853: RES 0,(HL)
label_21854: RES 1,(HL)
label_21855: RES 2,(HL)
label_21856: RES 3,(HL)
label_21857: RES 4,(HL)
label_21858: RES 5,(HL)
label_21859: RES 6,(HL)
label_21860: RES 7,(HL)
label_21861: RES 0,(IX + 127)
label_21862: RES 1,(IX + 127)
label_21863: RES 2,(IX + 127)
label_21864: RES 3,(IX + 127)
label_21865: RES 4,(IX + 127)
label_21866: RES 5,(IX + 127)
label_21867: RES 6,(IX + 127)
label_21868: RES 7,(IX + 127)
label_21869: RES 0,(IY - 128)
label_21870: RES 1,(IY - 128)
label_21871: RES 2,(IY - 128)
label_21872: RES 3,(IY - 128)
label_21873: RES 4,(IY - 128)
label_21874: RES 5,(IY - 128)
label_21875: RES 6,(IY - 128)
label_21876: RES 7,(IY - 128)
label_21877: JP $5678
label_21878: JP NZ,$5678
label_21879: JP Z,$5678
label_21880: JP NC,$5678
label_21881: JP C,$5678
label_21882: JP PO,$5678
label_21883: JP PE,$5678
label_21884: JP P,$5678
label_21885: JP M,$5678
label_21886: JR $ + 2
label_21887: JR NZ,$ + 2
label_21888: JR Z,$ + 2
label_21889: JR NC,$ + 2
label_21890: JR C,$ + 2
label_21891: JP (HL)
label_21892: JP (IX)
label_21893: JP (IY)
label_21894: DJNZ $ + 2
label_21895: CALL $5678
label_21896: CALL NZ,$5678
label_21897: CALL Z,$5678
label_21898: CALL NC,$5678
label_21899: CALL C,$5678
label_21900: CALL PO,$5678
label_21901: CALL PE,$5678
label_21902: CALL P,$5678
label_21903: CALL M,$5678
label_21904: RET
label_21905: RET NZ
label_21906: RET Z
label_21907: RET NC
label_21908: RET C
label_21909: RET PO
label_21910: RET PE
label_21911: RET P
label_21912: RET M
label_21913: RETI
label_21914: RETN
label_21915: RST $00
label_21916: RST $08
label_21917: RST $10
label_21918: RST $18
label_21919: RST $20
label_21920: RST $28
label_21921: RST $30
label_21922: RST $38
label_21923: IN A,($12)
label_21924: IN A,(C)
label_21925: IN B,(C)
label_21926: IN C,(C)
label_21927: IN D,(C)
label_21928: IN E,(C)
label_21929: IN H,(C)
label_21930: IN L,(C)
label_21931: IN F,(C)
label_21932: INI
label_21933: INIR
label_21934: IND
label_21935: INDR
label_21936: OUT ($12),A
label_21937: OUT (C),A
label_21938: OUT (C),B
label_21939: OUT (C),C
label_21940: OUT (C),D
label_21941: OUT (C),E
label_21942: OUT (C),H
label_21943: OUT (C),L
label_21944: OUTI
label_21945: OTIR
label_21946: OUTD
label_21947: OTDR
label_21948: LD A,A
label_21949: LD A,B
label_21950: LD A,C
label_21951: LD A,D
label_21952: LD A,E
label_21953: LD A,H
label_21954: LD A,L
label_21955: LD B,A
label_21956: LD B,B
label_21957: LD B,C
label_21958: LD B,D
label_21959: LD B,E
label_21960: LD B,H
label_21961: LD B,L
label_21962: LD C,A
label_21963: LD C,B
label_21964: LD C,C
label_21965: LD C,D
label_21966: LD C,E
label_21967: LD C,H
label_21968: LD C,L
label_21969: LD D,A
label_21970: LD D,B
label_21971: LD D,C
label_21972: LD D,D
label_21973: LD D,E
label_21974: LD D,H
label_21975: LD D,L
label_21976: LD E,A
label_21977: LD E,B
label_21978: LD E,C
label_21979: LD E,D
label_21980: LD E,E
label_21981: LD E,H
label_21982: LD E,L
label_21983: LD H,A
label_21984: LD H,B
label_21985: LD H,C
label_21986: LD H,D
label_21987: LD H,E
label_21988: LD H,H
label_21989: LD H,L
label_21990: LD L,A
label_21991: LD L,B
label_21992: LD L,C
label_21993: LD L,D
label_21994: LD L,E
label_21995: LD L,H
label_21996: LD L,L
label_21997: LD A,$12
label_21998: LD B,$12
label_21999: LD C,$12
label_22000: LD D,$12
label_22001: LD E,$12
label_22002: LD H,$12
label_22003: LD L,$12
label_22004: LD A,(HL)
label_22005: LD B,(HL)
label_22006: LD C,(HL)
label_22007: LD D,(HL)
label_22008: LD E,(HL)
label_22009: LD H,(HL)
label_22010: LD L,(HL)
label_22011: LD A,(IX + 127)
label_22012: LD B,(IX + 127)
label_22013: LD C,(IX + 127)
label_22014: LD D,(IX + 127)
label_22015: LD E,(IX + 127)
label_22016: LD H,(IX + 127)
label_22017: LD L,(IX + 127)
label_22018: LD A,(IY - 128)
label_22019: LD B,(IY - 128)
label_22020: LD C,(IY - 128)
label_22021: LD D,(IY - 128)
label_22022: LD E,(IY - 128)
label_22023: LD H,(IY - 128)
label_22024: LD L,(IY - 128)
label_22025: LD (HL),A
label_22026: LD (HL),B
label_22027: LD (HL),C
label_22028: LD (HL),D
label_22029: LD (HL),E
label_22030: LD (HL),H
label_22031: LD (HL),L
label_22032: LD (IX + 127),A
label_22033: LD (IX + 127),B
label_22034: LD (IX + 127),C
label_22035: LD (IX + 127),D
label_22036: LD (IX + 127),E
label_22037: LD (IX + 127),H
label_22038: LD (IX + 127),L
label_22039: LD (IY - 128),A
label_22040: LD (IY - 128),B
label_22041: LD (IY - 128),C
label_22042: LD (IY - 128),D
label_22043: LD (IY - 128),E
label_22044: LD (IY - 128),H
label_22045: LD (IY - 128),L
label_22046: LD (HL),$12
label_22047: LD (IX + 127),$12
label_22048: LD (IY - 128),$12
label_22049: LD A,(BC)
label_22050: LD A,(DE)
label_22051: LD A,($5678)
label_22052: LD (BC),A
label_22053: LD (DE),A
label_22054: LD ($5678),A
label_22055: LD A,I
label_22056: LD A,R
label_22057: LD I,A
label_22058: LD R,A
label_22059: LD BC,$5678
label_22060: LD DE,$5678
label_22061: LD HL,$5678
label_22062: LD SP,$5678
label_22063: LD IX,$5678
label_22064: LD IY,$5678
label_22065: LD HL,($5678)
label_22066: LD BC,($5678)
label_22067: LD DE,($5678)
label_22068: LD HL,($5678)
label_22069: LD SP,($5678)
label_22070: LD IX,($5678)
label_22071: LD IY,($5678)
label_22072: LD ($5678),HL
label_22073: LD ($5678),BC
label_22074: LD ($5678),DE
label_22075: LD ($5678),HL
label_22076: LD ($5678),SP
label_22077: LD ($5678),IX
label_22078: LD ($5678),IY
label_22079: LD SP,HL
label_22080: LD SP,IX
label_22081: LD SP,IY
label_22082: PUSH BC
label_22083: PUSH DE
label_22084: PUSH HL
label_22085: PUSH AF
label_22086: PUSH IX
label_22087: PUSH IY
label_22088: POP BC
label_22089: POP DE
label_22090: POP HL
label_22091: POP AF
label_22092: POP IX
label_22093: POP IY
label_22094: EX DE,HL
label_22095: EX AF,AF'
label_22096: EXX
label_22097: EX (SP),HL
label_22098: EX (SP),IX
label_22099: EX (SP),IY
label_22100: LDI
label_22101: LDIR
label_22102: LDD
label_22103: LDDR
label_22104: CPI
label_22105: CPIR
label_22106: CPD
label_22107: CPDR
label_22108: ADD A,A
label_22109: ADD A,B
label_22110: ADD A,C
label_22111: ADD A,D
label_22112: ADD A,E
label_22113: ADD A,H
label_22114: ADD A,L
label_22115: ADD A,$12
label_22116: ADD A,(HL)
label_22117: ADD A,(IX + 127)
label_22118: ADD A,(IY - 128)
label_22119: ADC A,A
label_22120: ADC A,B
label_22121: ADC A,C
label_22122: ADC A,D
label_22123: ADC A,E
label_22124: ADC A,H
label_22125: ADC A,L
label_22126: ADC A,$12
label_22127: ADC A,(HL)
label_22128: ADC A,(IX + 127)
label_22129: ADC A,(IY - 128)
label_22130: SUB A
label_22131: SUB B
label_22132: SUB C
label_22133: SUB D
label_22134: SUB E
label_22135: SUB H
label_22136: SUB L
label_22137: SUB $12
label_22138: SUB (HL)
label_22139: SUB (IX + 127)
label_22140: SUB (IY - 128)
label_22141: SBC A,A
label_22142: SBC A,B
label_22143: SBC A,C
label_22144: SBC A,D
label_22145: SBC A,E
label_22146: SBC A,H
label_22147: SBC A,L
label_22148: SBC A,$12
label_22149: SBC A,(HL)
label_22150: SBC A,(IX + 127)
label_22151: SBC A,(IY - 128)
label_22152: AND A
label_22153: AND B
label_22154: AND C
label_22155: AND D
label_22156: AND E
label_22157: AND H
label_22158: AND L
label_22159: AND $12
label_22160: AND (HL)
label_22161: AND (IX + 127)
label_22162: AND (IY - 128)
label_22163: AND A
label_22164: AND B
label_22165: AND C
label_22166: AND D
label_22167: AND E
label_22168: AND H
label_22169: AND L
label_22170: AND $12
label_22171: AND (HL)
label_22172: AND (IX + 127)
label_22173: AND (IY - 128)
label_22174: OR A
label_22175: OR B
label_22176: OR C
label_22177: OR D
label_22178: OR E
label_22179: OR H
label_22180: OR L
label_22181: OR $12
label_22182: OR (HL)
label_22183: OR (IX + 127)
label_22184: OR (IY - 128)
label_22185: XOR A
label_22186: XOR B
label_22187: XOR C
label_22188: XOR D
label_22189: XOR E
label_22190: XOR H
label_22191: XOR L
label_22192: XOR $12
label_22193: XOR (HL)
label_22194: XOR (IX + 127)
label_22195: XOR (IY - 128)
label_22196: CP A
label_22197: CP B
label_22198: CP C
label_22199: CP D
label_22200: CP E
label_22201: CP H
label_22202: CP L
label_22203: CP $12
label_22204: CP (HL)
label_22205: CP (IX + 127)
label_22206: CP (IY - 128)
label_22207: INC A
label_22208: INC B
label_22209: INC C
label_22210: INC D
label_22211: INC E
label_22212: INC H
label_22213: INC L
label_22214: INC (HL)
label_22215: INC (IX + 127)
label_22216: INC (IY - 128)
label_22217: DEC A
label_22218: DEC B
label_22219: DEC C
label_22220: DEC D
label_22221: DEC E
label_22222: DEC H
label_22223: DEC L
label_22224: DEC (HL)
label_22225: DEC (IX + 127)
label_22226: DEC (IY - 128)
label_22227: DAA
label_22228: CPL
label_22229: NEG
label_22230: CCF
label_22231: SCF
label_22232: NOP
label_22233: HALT
label_22234: DI
label_22235: EI
label_22236: IM 0
label_22237: IM 1
label_22238: IM 2
label_22239: ADD HL,BC
label_22240: ADD HL,DE
label_22241: ADD HL,HL
label_22242: ADD HL,SP
label_22243: ADC HL,BC
label_22244: ADC HL,DE
label_22245: ADC HL,HL
label_22246: ADC HL,SP
label_22247: SBC HL,BC
label_22248: SBC HL,DE
label_22249: SBC HL,HL
label_22250: SBC HL,SP
label_22251: ADD IX,BC
label_22252: ADD IX,DE
label_22253: ADD IX,SP
label_22254: ADD IY,BC
label_22255: ADD IY,DE
label_22256: ADD IY,SP
label_22257: INC BC
label_22258: INC DE
label_22259: INC HL
label_22260: INC SP
label_22261: INC IX
label_22262: INC IY
label_22263: DEC BC
label_22264: DEC DE
label_22265: DEC HL
label_22266: DEC SP
label_22267: DEC IX
label_22268: DEC IY
label_22269: RLCA
label_22270: RLA
label_22271: RRCA
label_22272: RRA
label_22273: RLC A
label_22274: RLC B
label_22275: RLC C
label_22276: RLC D
label_22277: RLC E
label_22278: RLC H
label_22279: RLC L
label_22280: RLC (HL)
label_22281: RLC (IX + 127)
label_22282: RLC (IY - 128)
label_22283: RL A
label_22284: RL B
label_22285: RL C
label_22286: RL D
label_22287: RL E
label_22288: RL H
label_22289: RL L
label_22290: RL (HL)
label_22291: RL (IX + 127)
label_22292: RL (IY - 128)
label_22293: RRC A
label_22294: RRC B
label_22295: RRC C
label_22296: RRC D
label_22297: RRC E
label_22298: RRC H
label_22299: RRC L
label_22300: RRC (HL)
label_22301: RRC (IX + 127)
label_22302: RRC (IY - 128)
label_22303: RR A
label_22304: RR B
label_22305: RR C
label_22306: RR D
label_22307: RR E
label_22308: RR H
label_22309: RR L
label_22310: RR (HL)
label_22311: RR (IX + 127)
label_22312: RR (IY - 128)
label_22313: SLA A
label_22314: SLA B
label_22315: SLA C
label_22316: SLA D
label_22317: SLA E
label_22318: SLA H
label_22319: SLA L
label_22320: SLA (HL)
label_22321: SLA (IX + 127)
label_22322: SLA (IY - 128)
label_22323: SRA A
label_22324: SRA B
label_22325: SRA C
label_22326: SRA D
label_22327: SRA E
label_22328: SRA H
label_22329: SRA L
label_22330: SRA (HL)
label_22331: SRA (IX + 127)
label_22332: SRA (IY - 128)
label_22333: SRL A
label_22334: SRL B
label_22335: SRL C
label_22336: SRL D
label_22337: SRL E
label_22338: SRL H
label_22339: SRL L
label_22340: SRL (HL)
label_22341: SRL (IX + 127)
label_22342: SRL (IY - 128)
label_22343: RLD
label_22344: RRD
label_22345: BIT 0,A
label_22346: BIT 1,A
label_22347: BIT 2,A
label_22348: BIT 3,A
label_22349: BIT 4,A
label_22350: BIT 5,A
label_22351: BIT 6,A
label_22352: BIT 7,A
label_22353: BIT 0,B
label_22354: BIT 1,B
label_22355: BIT 2,B
label_22356: BIT 3,B
label_22357: BIT 4,B
label_22358: BIT 5,B
label_22359: BIT 6,B
label_22360: BIT 7,B
label_22361: BIT 0,C
label_22362: BIT 1,C
label_22363: BIT 2,C
label_22364: BIT 3,C
label_22365: BIT 4,C
label_22366: BIT 5,C
label_22367: BIT 6,C
label_22368: BIT 7,C
label_22369: BIT 0,D
label_22370: BIT 1,D
label_22371: BIT 2,D
label_22372: BIT 3,D
label_22373: BIT 4,D
label_22374: BIT 5,D
label_22375: BIT 6,D
label_22376: BIT 7,D
label_22377: BIT 0,E
label_22378: BIT 1,E
label_22379: BIT 2,E
label_22380: BIT 3,E
label_22381: BIT 4,E
label_22382: BIT 5,E
label_22383: BIT 6,E
label_22384: BIT 7,E
label_22385: BIT 0,H
label_22386: BIT 1,H
label_22387: BIT 2,H
label_22388: BIT 3,H
label_22389: BIT 4,H
label_22390: BIT 5,H
label_22391: BIT 6,H
label_22392: BIT 7,H
label_22393: BIT 0,L
label_22394: BIT 1,L
label_22395: BIT 2,L
label_22396: BIT 3,L
label_22397: BIT 4,L
label_22398: BIT 5,L
label_22399: BIT 6,L
label_22400: BIT 7,L
label_22401: BIT 0,(HL)
label_22402: BIT 1,(HL)
label_22403: BIT 2,(HL)
label_22404: BIT 3,(HL)
label_22405: BIT 4,(HL)
label_22406: BIT 5,(HL)
label_22407: BIT 6,(HL)
label_22408: BIT 7,(HL)
label_22409: BIT 0,(IX + 127)
label_22410: BIT 1,(IX + 127)
label_22411: BIT 2,(IX + 127)
label_22412: BIT 3,(IX + 127)
label_22413: BIT 4,(IX + 127)
label_22414: BIT 5,(IX + 127)
label_22415: BIT 6,(IX + 127)
label_22416: BIT 7,(IX + 127)
label_22417: BIT 0,(IY - 128)
label_22418: BIT 1,(IY - 128)
label_22419: BIT 2,(IY - 128)
label_22420: BIT 3,(IY - 128)
label_22421: BIT 4,(IY - 128)
label_22422: BIT 5,(IY - 128)
label_22423: BIT 6,(IY - 128)
label_22424: BIT 7,(IY - 128)
label_22425: SET 0,A
label_22426: SET 1,A
label_22427: SET 2,A
label_22428: SET 3,A
label_22429: SET 4,A
label_22430: SET 5,A
label_22431: SET 6,A
label_22432: SET 7,A
label_22433: SET 0,B
label_22434: SET 1,B
label_22435: SET 2,B
label_22436: SET 3,B
label_22437: SET 4,B
label_22438: SET 5,B
label_22439: SET 6,B
label_22440: SET 7,B
label_22441: SET 0,C
label_22442: SET 1,C
label_22443: SET 2,C
label_22444: SET 3,C
label_22445: SET 4,C
label_22446: SET 5,C
label_22447: SET 6,C
label_22448: SET 7,C
label_22449: SET 0,D
label_22450: SET 1,D
label_22451: SET 2,D
label_22452: SET 3,D
label_22453: SET 4,D
label_22454: SET 5,D
label_22455: SET 6,D
label_22456: SET 7,D
label_22457: SET 0,E
label_22458: SET 1,E
label_22459: SET 2,E
label_22460: SET 3,E
label_22461: SET 4,E
label_22462: SET 5,E
label_22463: SET 6,E
label_22464: SET 7,E
label_22465: SET 0,H
label_22466: SET 1,H
label_22467: SET 2,H
label_22468: SET 3,H
label_22469: SET 4,H
label_22470: SET 5,H
label_22471: SET 6,H
label_22472: SET 7,H
label_22473: SET 0,L
label_22474: SET 1,L
label_22475: SET 2,L
label_22476: SET 3,L
label_22477: SET 4,L
label_22478: SET 5,L
label_22479: SET 6,L
label_22480: SET 7,L
label_22481: SET 0,(HL)
label_22482: SET 1,(HL)
label_22483: SET 2,(HL)
label_22484: SET 3,(HL)
label_22485: SET 4,(HL)
label_22486: SET 5,(HL)
label_22487: SET 6,(HL)
label_22488: SET 7,(HL)
label_22489: SET 0,(IX + 127)
label_22490: SET 1,(IX + 127)
label_22491: SET 2,(IX + 127)
label_22492: SET 3,(IX + 127)
label_22493: SET 4,(IX + 127)
label_22494: SET 5,(IX + 127)
label_22495: SET 6,(IX + 127)
label_22496: SET 7,(IX + 127)
label_22497: SET 0,(IY - 128)
label_22498: SET 1,(IY - 128)
label_22499: SET 2,(IY - 128)
label_22500: SET 3,(IY - 128)
label_22501: SET 4,(IY - 128)
label_22502: SET 5,(IY - 128)
label_22503: SET 6,(IY - 128)
label_22504: SET 7,(IY - 128)
label_22505: RES 0,A
label_22506: RES 1,A
label_22507: RES 2,A
label_22508: RES 3,A
label_22509: RES 4,A
label_22510: RES 5,A
label_22511: RES 6,A
label_22512: RES 7,A
label_22513: RES 0,B
label_22514: RES 1,B
label_22515: RES 2,B
label_22516: RES 3,B
label_22517: RES 4,B
label_22518: RES 5,B
label_22519: RES 6,B
label_22520: RES 7,B
label_22521: RES 0,C
label_22522: RES 1,C
label_22523: RES 2,C
label_22524: RES 3,C
label_22525: RES 4,C
label_22526: RES 5,C
label_22527: RES 6,C
label_22528: RES 7,C
label_22529: RES 0,D
label_22530: RES 1,D
label_22531: RES 2,D
label_22532: RES 3,D
label_22533: RES 4,D
label_22534: RES 5,D
label_22535: RES 6,D
label_22536: RES 7,D
label_22537: RES 0,E
label_22538: RES 1,E
label_22539: RES 2,E
label_22540: RES 3,E
label_22541: RES 4,E
label_22542: RES 5,E
label_22543: RES 6,E
label_22544: RES 7,E
label_22545: RES 0,H
label_22546: RES 1,H
label_22547: RES 2,H
label_22548: RES 3,H
label_22549: RES 4,H
label_22550: RES 5,H
label_22551: RES 6,H
label_22552: RES 7,H
label_22553: RES 0,L
label_22554: RES 1,L
label_22555: RES 2,L
label_22556: RES 3,L
label_22557: RES 4,L
label_22558: RES 5,L
label_22559: RES 6,L
label_22560: RES 7,L
label_22561: RES 0,(HL)
label_22562: RES 1,(HL)
label_22563: RES 2,(HL)
label_22564: RES 3,(HL)
label_22565: RES 4,(HL)
label_22566: RES 5,(HL)
label_22567: RES 6,(HL)
label_22568: RES 7,(HL)
label_22569: RES 0,(IX + 127)
label_22570: RES 1,(IX + 127)
label_22571: RES 2,(IX + 127)
label_22572: RES 3,(IX + 127)
label_22573: RES 4,(IX + 127)
label_22574: RES 5,(IX + 127)
label_22575: RES 6,(IX + 127)
label_22576: RES 7,(IX + 127)
label_22577: RES 0,(IY - 128)
label_22578: RES 1,(IY - 128)
label_22579: RES 2,(IY - 128)
label_22580: RES 3,(IY - 128)
label_22581: RES 4,(IY - 128)
label_22582: RES 5,(IY - 128)
label_22583: RES 6,(IY - 128)
label_22584: RES 7,(IY - 128)
label_22585: JP $5678
label_22586: JP NZ,$5678
label_22587: JP Z,$5678
label_22588: JP NC,$5678
label_22589: JP C,$5678
label_22590: JP PO,$5678
label_22591: JP PE,$5678
label_22592: JP P,$5678
label_22593: JP M,$5678
label_22594: JR $ + 2
label_22595: JR NZ,$ + 2
label_22596: JR Z,$ + 2
label_22597: JR NC,$ + 2
label_22598: JR C,$ + 2
label_22599: JP (HL)
label_22600: JP (IX)
label_22601: JP (IY)
label_22602: DJNZ $ + 2
label_22603: CALL $5678
label_22604: CALL NZ,$5678
label_22605: CALL Z,$5678
label_22606: CALL NC,$5678
label_22607: CALL C,$5678
label_22608: CALL PO,$5678
label_22609: CALL PE,$5678
label_22610: CALL P,$5678
label_22611: CALL M,$5678
label_22612: RET
label_22613: RET NZ
label_22614: RET Z
label_22615: RET NC
label_22616: RET C
label_22617: RET PO
label_22618: RET PE
label_22619: RET P
label_22620: RET M
label_22621: RETI
label_22622: RETN
label_22623: RST $00
label_22624: RST $08
label_22625: RST $10
label_22626: RST $18
label_22627: RST $20
label_22628: RST $28
label_22629: RST $30
label_22630: RST $38
label_22631: IN A,($12)
label_22632: IN A,(C)
label_22633: IN B,(C)
label_22634: IN C,(C)
label_22635: IN D,(C)
label_22636: IN E,(C)
label_22637: IN H,(C)
label_22638: IN L,(C)
label_22639: IN F,(C)
label_22640: INI
label_22641: INIR
label_22642: IND
label_22643: INDR
label_22644: OUT ($12),A
label_22645: OUT (C),A
label_22646: OUT (C),B
label_22647: OUT (C),C
label_22648: OUT (C),D
label_22649: OUT (C),E
label_22650: OUT (C),H
label_22651: OUT (C),L
label_22652: OUTI
label_22653: OTIR
label_22654: OUTD
label_22655: OTDR
label_22656: LD A,A
label_22657: LD A,B
label_22658: LD A,C
label_22659: LD A,D
label_22660: LD A,E
label_22661: LD A,H
label_22662: LD A,L
label_22663: LD B,A
label_22664: LD B,B
label_22665: LD B,C
label_22666: LD B,D
label_22667: LD B,E
label_22668: LD B,H
label_22669: LD B,L
label_22670: LD C,A
label_22671: LD C,B
label_22672: LD C,C
label_22673: LD C,D
label_22674: LD C,E
label_22675: LD C,H
label_22676: LD C,L
label_22677: LD D,A
label_22678: LD D,B
label_22679: LD D,C
label_22680: LD D,D
label_22681: LD D,E
label_22682: LD D,H
label_22683: LD D,L
label_22684: LD E,A
label_22685: LD E,B
label_22686: LD E,C
label_22687: LD E,D
label_22688: LD E,E
label_22689: LD E,H
label_22690: LD E,L
label_22691: LD H,A
label_22692: LD H,B
label_22693: LD H,C
label_22694: LD H,D
label_22695: LD H,E
label_22696: LD H,H
label_22697: LD H,L
label_22698: LD L,A
label_22699: LD L,B
label_22700: LD L,C
label_22701: LD L,D
label_22702: LD L,E
label_22703: LD L,H
label_22704: LD L,L
label_22705: LD A,$12
label_22706: LD B,$12
label_22707: LD C,$12
label_22708: LD D,$12
label_22709: LD E,$12
label_22710: LD H,$12
label_22711: LD L,$12
label_22712: LD A,(HL)
label_22713: LD B,(HL)
label_22714: LD C,(HL)
label_22715: LD D,(HL)
label_22716: LD E,(HL)
label_22717: LD H,(HL)
label_22718: LD L,(HL)
label_22719: LD A,(IX + 127)
label_22720: LD B,(IX + 127)
label_22721: LD C,(IX + 127)
label_22722: LD D,(IX + 127)
label_22723: LD E,(IX + 127)
label_22724: LD H,(IX + 127)
label_22725: LD L,(IX + 127)
label_22726: LD A,(IY - 128)
label_22727: LD B,(IY - 128)
label_22728: LD C,(IY - 128)
label_22729: LD D,(IY - 128)
label_22730: LD E,(IY - 128)
label_22731: LD H,(IY - 128)
label_22732: LD L,(IY - 128)
label_22733: LD (HL),A
label_22734: LD (HL),B
label_22735: LD (HL),C
label_22736: LD (HL),D
label_22737: LD (HL),E
label_22738: LD (HL),H
label_22739: LD (HL),L
label_22740: LD (IX + 127),A
label_22741: LD (IX + 127),B
label_22742: LD (IX + 127),C
label_22743: LD (IX + 127),D
label_22744: LD (IX + 127),E
label_22745: LD (IX + 127),H
label_22746: LD (IX + 127),L
label_22747: LD (IY - 128),A
label_22748: LD (IY - 128),B
label_22749: LD (IY - 128),C
label_22750: LD (IY - 128),D
label_22751: LD (IY - 128),E
label_22752: LD (IY - 128),H
label_22753: LD (IY - 128),L
label_22754: LD (HL),$12
label_22755: LD (IX + 127),$12
label_22756: LD (IY - 128),$12
label_22757: LD A,(BC)
label_22758: LD A,(DE)
label_22759: LD A,($5678)
label_22760: LD (BC),A
label_22761: LD (DE),A
label_22762: LD ($5678),A
label_22763: LD A,I
label_22764: LD A,R
label_22765: LD I,A
label_22766: LD R,A
label_22767: LD BC,$5678
label_22768: LD DE,$5678
label_22769: LD HL,$5678
label_22770: LD SP,$5678
label_22771: LD IX,$5678
label_22772: LD IY,$5678
label_22773: LD HL,($5678)
label_22774: LD BC,($5678)
label_22775: LD DE,($5678)
label_22776: LD HL,($5678)
label_22777: LD SP,($5678)
label_22778: LD IX,($5678)
label_22779: LD IY,($5678)
label_22780: LD ($5678),HL
label_22781: LD ($5678),BC
label_22782: LD ($5678),DE
label_22783: LD ($5678),HL
label_22784: LD ($5678),SP
label_22785: LD ($5678),IX
label_22786: LD ($5678),IY
label_22787: LD SP,HL
label_22788: LD SP,IX
label_22789: LD SP,IY
label_22790: PUSH BC
label_22791: PUSH DE
label_22792: PUSH HL
label_22793: PUSH AF
label_22794: PUSH IX
label_22795: PUSH IY
label_22796: POP BC
label_22797: POP DE
label_22798: POP HL
label_22799: POP AF
label_22800: POP IX
label_22801: POP IY
label_22802: EX DE,HL
label_22803: EX AF,AF'
label_22804: EXX
label_22805: EX (SP),HL
label_22806: EX (SP),IX
label_22807: EX (SP),IY
label_22808: LDI
label_22809: LDIR
label_22810: LDD
label_22811: LDDR
label_22812: CPI
label_22813: CPIR
label_22814: CPD
label_22815: CPDR
label_22816: ADD A,A
label_22817: ADD A,B
label_22818: ADD A,C
label_22819: ADD A,D
label_22820: ADD A,E
label_22821: ADD A,H
label_22822: ADD A,L
label_22823: ADD A,$12
label_22824: ADD A,(HL)
label_22825: ADD A,(IX + 127)
label_22826: ADD A,(IY - 128)
label_22827: ADC A,A
label_22828: ADC A,B
label_22829: ADC A,C
label_22830: ADC A,D
label_22831: ADC A,E
label_22832: ADC A,H
label_22833: ADC A,L
label_22834: ADC A,$12
label_22835: ADC A,(HL)
label_22836: ADC A,(IX + 127)
label_22837: ADC A,(IY - 128)
label_22838: SUB A
label_22839: SUB B
label_22840: SUB C
label_22841: SUB D
label_22842: SUB E
label_22843: SUB H
label_22844: SUB L
label_22845: SUB $12
label_22846: SUB (HL)
label_22847: SUB (IX + 127)
label_22848: SUB (IY - 128)
label_22849: SBC A,A
label_22850: SBC A,B
label_22851: SBC A,C
label_22852: SBC A,D
label_22853: SBC A,E
label_22854: SBC A,H
label_22855: SBC A,L
label_22856: SBC A,$12
label_22857: SBC A,(HL)
label_22858: SBC A,(IX + 127)
label_22859: SBC A,(IY - 128)
label_22860: AND A
label_22861: AND B
label_22862: AND C
label_22863: AND D
label_22864: AND E
label_22865: AND H
label_22866: AND L
label_22867: AND $12
label_22868: AND (HL)
label_22869: AND (IX + 127)
label_22870: AND (IY - 128)
label_22871: AND A
label_22872: AND B
label_22873: AND C
label_22874: AND D
label_22875: AND E
label_22876: AND H
label_22877: AND L
label_22878: AND $12
label_22879: AND (HL)
label_22880: AND (IX + 127)
label_22881: AND (IY - 128)
label_22882: OR A
label_22883: OR B
label_22884: OR C
label_22885: OR D
label_22886: OR E
label_22887: OR H
label_22888: OR L
label_22889: OR $12
label_22890: OR (HL)
label_22891: OR (IX + 127)
label_22892: OR (IY - 128)
label_22893: XOR A
label_22894: XOR B
label_22895: XOR C
label_22896: XOR D
label_22897: XOR E
label_22898: XOR H
label_22899: XOR L
label_22900: XOR $12
label_22901: XOR (HL)
label_22902: XOR (IX + 127)
label_22903: XOR (IY - 128)
label_22904: CP A
label_22905: CP B
label_22906: CP C
label_22907: CP D
label_22908: CP E
label_22909: CP H
label_22910: CP L
label_22911: CP $12
label_22912: CP (HL)
label_22913: CP (IX + 127)
label_22914: CP (IY - 128)
label_22915: INC A
label_22916: INC B
label_22917: INC C
label_22918: INC D
label_22919: INC E
label_22920: INC H
label_22921: INC L
label_22922: INC (HL)
label_22923: INC (IX + 127)
label_22924: INC (IY - 128)
label_22925: DEC A
label_22926: DEC B
label_22927: DEC C
label_22928: DEC D
label_22929: DEC E
label_22930: DEC H
label_22931: DEC L
label_22932: DEC (HL)
label_22933: DEC (IX + 127)
label_22934: DEC (IY - 128)
label_22935: DAA
label_22936: CPL
label_22937: NEG
label_22938: CCF
label_22939: SCF
label_22940: NOP
label_22941: HALT
label_22942: DI
label_22943: EI
label_22944: IM 0
label_22945: IM 1
label_22946: IM 2
label_22947: ADD HL,BC
label_22948: ADD HL,DE
label_22949: ADD HL,HL
label_22950: ADD HL,SP
label_22951: ADC HL,BC
label_22952: ADC HL,DE
label_22953: ADC HL,HL
label_22954: ADC HL,SP
label_22955: SBC HL,BC
label_22956: SBC HL,DE
label_22957: SBC HL,HL
label_22958: SBC HL,SP
label_22959: ADD IX,BC
label_22960: ADD IX,DE
label_22961: ADD IX,SP
label_22962: ADD IY,BC
label_22963: ADD IY,DE
label_22964: ADD IY,SP
label_22965: INC BC
label_22966: INC DE
label_22967: INC HL
label_22968: INC SP
label_22969: INC IX
label_22970: INC IY
label_22971: DEC BC
label_22972: DEC DE
label_22973: DEC HL
label_22974: DEC SP
label_22975: DEC IX
label_22976: DEC IY
label_22977: RLCA
label_22978: RLA
label_22979: RRCA
label_22980: RRA
label_22981: RLC A
label_22982: RLC B
label_22983: RLC C
label_22984: RLC D
label_22985: RLC E
label_22986: RLC H
label_22987: RLC L
label_22988: RLC (HL)
label_22989: RLC (IX + 127)
label_22990: RLC (IY - 128)
label_22991: RL A
label_22992: RL B
label_22993: RL C
label_22994: RL D
label_22995: RL E
label_22996: RL H
label_22997: RL L
label_22998: RL (HL)
label_22999: RL (IX + 127)
label_23000: RL (IY - 128)
label_23001: RRC A
label_23002: RRC B
label_23003: RRC C
label_23004: RRC D
label_23005: RRC E
label_23006: RRC H
label_23007: RRC L
label_23008: RRC (HL)
label_23009: RRC (IX + 127)
label_23010: RRC (IY - 128)
label_23011: RR A
label_23012: RR B
label_23013: RR C
label_23014: RR D
label_23015: RR E
label_23016: RR H
label_23017: RR L
label_23018: RR (HL)
label_23019: RR (IX + 127)
label_23020: RR (IY - 128)
label_23021: SLA A
label_23022: SLA B
label_23023: SLA C
label_23024: SLA D
label_23025: SLA E
label_23026: SLA H
label_23027: SLA L
label_23028: SLA (HL)
label_23029: SLA (IX + 127)
label_23030: SLA (IY - 128)
label_23031: SRA A
label_23032: SRA B
label_23033: SRA C
label_23034: SRA D
label_23035: SRA E
label_23036: SRA H
label_23037: SRA L
label_23038: SRA (HL)
label_23039: SRA (IX + 127)
label_23040: SRA (IY - 128)
label_23041: SRL A
label_23042: SRL B
label_23043: SRL C
label_23044: SRL D
label_23045: SRL E
label_23046: SRL H
label_23047: SRL L
label_23048: SRL (HL)
label_23049: SRL (IX + 127)
label_23050: SRL (IY - 128)
label_23051: RLD
label_23052: RRD
label_23053: BIT 0,A
label_23054: BIT 1,A
label_23055: BIT 2,A
label_23056: BIT 3,A
label_23057: BIT 4,A
label_23058: BIT 5,A
label_23059: BIT 6,A
label_23060: BIT 7,A
label_23061: BIT 0,B
label_23062: BIT 1,B
label_23063: BIT 2,B
label_23064: BIT 3,B
label_23065: BIT 4,B
label_23066: BIT 5,B
label_23067: BIT 6,B
label_23068: BIT 7,B
label_23069: BIT 0,C
label_23070: BIT 1,C
label_23071: BIT 2,C
label_23072: BIT 3,C
label_23073: BIT 4,C
label_23074: BIT 5,C
label_23075: BIT 6,C
label_23076: BIT 7,C
label_23077: BIT 0,D
label_23078: BIT 1,D
label_23079: BIT 2,D
label_23080: BIT 3,D
label_23081: BIT 4,D
label_23082: BIT 5,D
label_23083: BIT 6,D
label_23084: BIT 7,D
label_23085: BIT 0,E
label_23086: BIT 1,E
label_23087: BIT 2,E
label_23088: BIT 3,E
label_23089: BIT 4,E
label_23090: BIT 5,E
label_23091: BIT 6,E
label_23092: BIT 7,E
label_23093: BIT 0,H
label_23094: BIT 1,H
label_23095: BIT 2,H
label_23096: BIT 3,H
label_23097: BIT 4,H
label_23098: BIT 5,H
label_23099: BIT 6,H
label_23100: BIT 7,H
label_23101: BIT 0,L
label_23102: BIT 1,L
label_23103: BIT 2,L
label_23104: BIT 3,L
label_23105: BIT 4,L
label_23106: BIT 5,L
label_23107: BIT 6,L
label_23108: BIT 7,L
label_23109: BIT 0,(HL)
label_23110: BIT 1,(HL)
label_23111: BIT 2,(HL)
label_23112: BIT 3,(HL)
label_23113: BIT 4,(HL)
label_23114: BIT 5,(HL)
label_23115: BIT 6,(HL)
label_23116: BIT 7,(HL)
label_23117: BIT 0,(IX + 127)
label_23118: BIT 1,(IX + 127)
label_23119: BIT 2,(IX + 127)
label_23120: BIT 3,(IX + 127)
label_23121: BIT 4,(IX + 127)
label_23122: BIT 5,(IX + 127)
label_23123: BIT 6,(IX + 127)
label_23124: BIT 7,(IX + 127)
label_23125: BIT 0,(IY - 128)
label_23126: BIT 1,(IY - 128)
label_23127: BIT 2,(IY - 128)
label_23128: BIT 3,(IY - 128)
label_23129: BIT 4,(IY - 128)
label_23130: BIT 5,(IY - 128)
label_23131: BIT 6,(IY - 128)
label_23132: BIT 7,(IY - 128)
label_23133: SET 0,A
label_23134: SET 1,A
label_23135: SET 2,A
label_23136: SET 3,A
label_23137: SET 4,A
label_23138: SET 5,A
label_23139: SET 6,A
label_23140: SET 7,A
label_23141: SET 0,B
label_23142: SET 1,B
label_23143: SET 2,B
label_23144: SET 3,B
label_23145: SET 4,B
label_23146: SET 5,B
label_23147: SET 6,B
label_23148: SET 7,B
label_23149: SET 0,C
label_23150: SET 1,C
label_23151: SET 2,C
label_23152: SET 3,C
label_23153: SET 4,C
label_23154: SET 5,C
label_23155: SET 6,C
label_23156: SET 7,C
label_23157: SET 0,D
label_23158: SET 1,D
label_23159: SET 2,D
label_23160: SET 3,D
label_23161: SET 4,D
label_23162: SET 5,D
label_23163: SET 6,D
label_23164: SET 7,D
label_23165: SET 0,E
label_23166: SET 1,E
label_23167: SET 2,E
label_23168: SET 3,E
label_23169: SET 4,E
label_23170: SET 5,E
label_23171: SET 6,E
label_23172: SET 7,E
label_23173: SET 0,H
label_23174: SET 1,H
label_23175: SET 2,H
label_23176: SET 3,H
label_23177: SET 4,H
label_23178: SET 5,H
label_23179: SET 6,H
label_23180: SET 7,H
label_23181: SET 0,L
label_23182: SET 1,L
label_23183: SET 2,L
label_23184: SET 3,L
label_23185: SET 4,L
label_23186: SET 5,L
label_23187: SET 6,L
label_23188: SET 7,L
label_23189: SET 0,(HL)
label_23190: SET 1,(HL)
label_23191: SET 2,(HL)
label_23192: SET 3,(HL)
label_23193: SET 4,(HL)
label_23194: SET 5,(HL)
label_23195: SET 6,(HL)
label_23196: SET 7,(HL)
label_23197: SET 0,(IX + 127)
label_23198: SET 1,(IX + 127)
label_23199: SET 2,(IX + 127)
label_23200: SET 3,(IX + 127)
label_23201: SET 4,(IX + 127)
label_23202: SET 5,(IX + 127)
label_23203: SET 6,(IX + 127)
label_23204: SET 7,(IX + 127)
label_23205: SET 0,(IY - 128)
label_23206: SET 1,(IY - 128)
label_23207: SET 2,(IY - 128)
label_23208: SET 3,(IY - 128)
label_23209: SET 4,(IY - 128)
label_23210: SET 5,(IY - 128)
label_23211: SET 6,(IY - 128)
label_23212: SET 7,(IY - 128)
label_23213: RES 0,A
label_23214: RES 1,A
label_23215: RES 2,A
label_23216: RES 3,A
label_23217: RES 4,A
label_23218: RES 5,A
label_23219: RES 6,A
label_23220: RES 7,A
label_23221: RES 0,B
label_23222: RES 1,B
label_23223: RES 2,B
label_23224: RES 3,B
label_23225: RES 4,B
label_23226: RES 5,B
label_23227: RES 6,B
label_23228: RES 7,B
label_23229: RES 0,C
label_23230: RES 1,C
label_23231: RES 2,C
label_23232: RES 3,C
label_23233: RES 4,C
label_23234: RES 5,C
label_23235: RES 6,C
label_23236: RES 7,C
label_23237: RES 0,D
label_23238: RES 1,D
label_23239: RES 2,D
label_23240: RES 3,D
label_23241: RES 4,D
label_23242: RES 5,D
label_23243: RES 6,D
label_23244: RES 7,D
label_23245: RES 0,E
label_23246: RES 1,E
label_23247: RES 2,E
label_23248: RES 3,E
label_23249: RES 4,E
label_23250: RES 5,E
label_23251: RES 6,E
label_23252: RES 7,E
label_23253: RES 0,H
label_23254: RES 1,H
label_23255: RES 2,H
label_23256: RES 3,H
label_23257: RES 4,H
label_23258: RES 5,H
label_23259: RES 6,H
label_23260: RES 7,H
label_23261: RES 0,L
label_23262: RES 1,L
label_23263: RES 2,L
label_23264: RES 3,L
label_23265: RES 4,L
label_23266: RES 5,L
label_23267: RES 6,L
label_23268: RES 7,L
label_23269: RES 0,(HL)
label_23270: RES 1,(HL)
label_23271: RES 2,(HL)
label_23272: RES 3,(HL)
label_23273: RES 4,(HL)
label_23274: RES 5,(HL)
label_23275: RES 6,(HL)
label_23276: RES 7,(HL)
label_23277: RES 0,(IX + 127)
label_23278: RES 1,(IX + 127)
label_23279: RES 2,(IX + 127)
label_23280: RES 3,(IX + 127)
label_23281: RES 4,(IX + 127)
label_23282: RES 5,(IX + 127)
label_23283: RES 6,(IX + 127)
label_23284: RES 7,(IX + 127)
label_23285: RES 0,(IY - 128)
label_23286: RES 1,(IY - 128)
label_23287: RES 2,(IY - 128)
label_23288: RES 3,(IY - 128)
label_23289: RES 4,(IY - 128)
label_23290: RES 5,(IY - 128)
label_23291: RES 6,(IY - 128)
label_23292: RES 7,(IY - 128)
label_23293: JP $5678
label_23294: JP NZ,$5678
label_23295: JP Z,$5678
label_23296: JP NC,$5678
label_23297: JP C,$5678
label_23298: JP PO,$5678
label_23299: JP PE,$5678
label_23300: JP P,$5678
label_23301: JP M,$5678
label_23302: JR $ + 2
label_23303: JR NZ,$ + 2
label_23304: JR Z,$ + 2
label_23305: JR NC,$ + 2
label_23306: JR C,$ + 2
label_23307: JP (HL)
label_23308: JP (IX)
label_23309: JP (IY)
label_23310: DJNZ $ + 2
label_23311: CALL $5678
label_23312: CALL NZ,$5678
label_23313: CALL Z,$5678
label_23314: CALL NC,$5678
label_23315: CALL C,$5678
label_23316: CALL PO,$5678
label_23317: CALL PE,$5678
label_23318: CALL P,$5678
label_23319: CALL M,$5678
label_23320: RET
label_23321: RET NZ
label_23322: RET Z
label_23323: RET NC
label_23324: RET C
label_23325: RET PO
label_23326: RET PE
label_23327: RET P
label_23328: RET M
label_23329: RETI
label_23330: RETN
label_23331: RST $00
label_23332: RST $08
label_23333: RST $10
label_23334: RST $18
label_23335: RST $20
label_23336: RST $28
label_23337: RST $30
label_23338: RST $38
label_23339: IN A,($12)
label_23340: IN A,(C)
label_23341: IN B,(C)
label_23342: IN C,(C)
label_23343: IN D,(C)
label_23344: IN E,(C)
label_23345: IN H,(C)
label_23346: IN L,(C)
label_23347: IN F,(C)
label_23348: INI
label_23349: INIR
label_23350: IND
label_23351: INDR
label_23352: OUT ($12),A
label_23353: OUT (C),A
label_23354: OUT (C),B
label_23355: OUT (C),C
label_23356: OUT (C),D
label_23357: OUT (C),E
label_23358: OUT (C),H
label_23359: OUT (C),L
label_23360: OUTI
label_23361: OTIR
label_23362: OUTD
label_23363: OTDR
label_23364: LD A,A
label_23365: LD A,B
label_23366: LD A,C
label_23367: LD A,D
label_23368: LD A,E
label_23369: LD A,H
label_23370: LD A,L
label_23371: LD B,A
label_23372: LD B,B
label_23373: LD B,C
label_23374: LD B,D
label_23375: LD B,E
label_23376: LD B,H
label_23377: LD B,L
label_23378: LD C,A
label_23379: LD C,B
label_23380: LD C,C
label_23381: LD C,D
label_23382: LD C,E
label_23383: LD C,H
label_23384: LD C,L
label_23385: LD D,A
label_23386: LD D,B
label_23387: LD D,C
label_23388: LD D,D
label_23389: LD D,E
label_23390: LD D,H
label_23391: LD D,L
label_23392: LD E,A
label_23393: LD E,B
label_23394: LD E,C
label_23395: LD E,D
label_23396: LD E,E
label_23397: LD E,H
label_23398: LD E,L
label_23399: LD H,A
label_23400: LD H,B
label_23401: LD H,C
label_23402: LD H,D
label_23403: LD H,E
label_23404: LD H,H
label_23405: LD H,L
label_23406: LD L,A
label_23407: LD L,B
label_23408: LD L,C
label_23409: LD L,D
label_23410: LD L,E
label_23411: LD L,H
label_23412: LD L,L
label_23413: LD A,$12
label_23414: LD B,$12
label_23415: LD C,$12
label_23416: LD D,$12
label_23417: LD E,$12
label_23418: LD H,$12
label_23419: LD L,$12
label_23420: LD A,(HL)
label_23421: LD B,(HL)
label_23422: LD C,(HL)
label_23423: LD D,(HL)
label_23424: LD E,(HL)
label_23425: LD H,(HL)
label_23426: LD L,(HL)
label_23427: LD A,(IX + 127)
label_23428: LD B,(IX + 127)
label_23429: LD C,(IX + 127)
label_23430: LD D,(IX + 127)
label_23431: LD E,(IX + 127)
label_23432: LD H,(IX + 127)
label_23433: LD L,(IX + 127)
label_23434: LD A,(IY - 128)
label_23435: LD B,(IY - 128)
label_23436: LD C,(IY - 128)
label_23437: LD D,(IY - 128)
label_23438: LD E,(IY - 128)
label_23439: LD H,(IY - 128)
label_23440: LD L,(IY - 128)
label_23441: LD (HL),A
label_23442: LD (HL),B
label_23443: LD (HL),C
label_23444: LD (HL),D
label_23445: LD (HL),E
label_23446: LD (HL),H
label_23447: LD (HL),L
label_23448: LD (IX + 127),A
label_23449: LD (IX + 127),B
label_23450: LD (IX + 127),C
label_23451: LD (IX + 127),D
label_23452: LD (IX + 127),E
label_23453: LD (IX + 127),H
label_23454: LD (IX + 127),L
label_23455: LD (IY - 128),A
label_23456: LD (IY - 128),B
label_23457: LD (IY - 128),C
label_23458: LD (IY - 128),D
label_23459: LD (IY - 128),E
label_23460: LD (IY - 128),H
label_23461: LD (IY - 128),L
label_23462: LD (HL),$12
label_23463: LD (IX + 127),$12
label_23464: LD (IY - 128),$12
label_23465: LD A,(BC)
label_23466: LD A,(DE)
label_23467: LD A,($5678)
label_23468: LD (BC),A
label_23469: LD (DE),A
label_23470: LD ($5678),A
label_23471: LD A,I
label_23472: LD A,R
label_23473: LD I,A
label_23474: LD R,A
label_23475: LD BC,$5678
label_23476: LD DE,$5678
label_23477: LD HL,$5678
label_23478: LD SP,$5678
label_23479: LD IX,$5678
label_23480: LD IY,$5678
label_23481: LD HL,($5678)
label_23482: LD BC,($5678)
label_23483: LD DE,($5678)
label_23484: LD HL,($5678)
label_23485: LD SP,($5678)
label_23486: LD IX,($5678)
label_23487: LD IY,($5678)
label_23488: LD ($5678),HL
label_23489: LD ($5678),BC
label_23490: LD ($5678),DE
label_23491: LD ($5678),HL
label_23492: LD ($5678),SP
label_23493: LD ($5678),IX
label_23494: LD ($5678),IY
label_23495: LD SP,HL
label_23496: LD SP,IX
label_23497: LD SP,IY
label_23498: PUSH BC
label_23499: PUSH DE
label_23500: PUSH HL
label_23501: PUSH AF
label_23502: PUSH IX
label_23503: PUSH IY
label_23504: POP BC
label_23505: POP DE
label_23506: POP HL
label_23507: POP AF
label_23508: POP IX
label_23509: POP IY
label_23510: EX DE,HL
label_23511: EX AF,AF'
label_23512: EXX
label_23513: EX (SP),HL
label_23514: EX (SP),IX
label_23515: EX (SP),IY
label_23516: LDI
label_23517: LDIR
label_23518: LDD
label_23519: LDDR
label_23520: CPI
label_23521: CPIR
label_23522: CPD
label_23523: CPDR
label_23524: ADD A,A
label_23525: ADD A,B
label_23526: ADD A,C
label_23527: ADD A,D
label_23528: ADD A,E
label_23529: ADD A,H
label_23530: ADD A,L
label_23531: ADD A,$12
label_23532: ADD A,(HL)
label_23533: ADD A,(IX + 127)
label_23534: ADD A,(IY - 128)
label_23535: ADC A,A
label_23536: ADC A,B
label_23537: ADC A,C
label_23538: ADC A,D
label_23539: ADC A,E
label_23540: ADC A,H
label_23541: ADC A,L
label_23542: ADC A,$12
label_23543: ADC A,(HL)
label_23544: ADC A,(IX + 127)
label_23545: ADC A,(IY - 128)
label_23546: SUB A
label_23547: SUB B
label_23548: SUB C
label_23549: SUB D
label_23550: SUB E
label_23551: SUB H
label_23552: SUB L
label_23553: SUB $12
label_23554: SUB (HL)
label_23555: SUB (IX + 127)
label_23556: SUB (IY - 128)
label_23557: SBC A,A
label_23558: SBC A,B
label_23559: SBC A,C
label_23560: SBC A,D
label_23561: SBC A,E
label_23562: SBC A,H
label_23563: SBC A,L
label_23564: SBC A,$12
label_23565: SBC A,(HL)
label_23566: SBC A,(IX + 127)
label_23567: SBC A,(IY - 128)
label_23568: AND A
label_23569: AND B
label_23570: AND C
label_23571: AND D
label_23572: AND E
label_23573: AND H
label_23574: AND L
label_23575: AND $12
label_23576: AND (HL)
label_23577: AND (IX + 127)
label_23578: AND (IY - 128)
label_23579: AND A
label_23580: AND B
label_23581: AND C
label_23582: AND D
label_23583: AND E
label_23584: AND H
label_23585: AND L
label_23586: AND $12
label_23587: AND (HL)
label_23588: AND (IX + 127)
label_23589: AND (IY - 128)
label_23590: OR A
label_23591: OR B
label_23592: OR C
label_23593: OR D
label_23594: OR E
label_23595: OR H
label_23596: OR L
label_23597: OR $12
label_23598: OR (HL)
label_23599: OR (IX + 127)
label_23600: OR (IY - 128)
label_23601: XOR A
label_23602: XOR B
label_23603: XOR C
label_23604: XOR D
label_23605: XOR E
label_23606: XOR H
label_23607: XOR L
label_23608: XOR $12
label_23609: XOR (HL)
label_23610: XOR (IX + 127)
label_23611: XOR (IY - 128)
label_23612: CP A
label_23613: CP B
label_23614: CP C
label_23615: CP D
label_23616: CP E
label_23617: CP H
label_23618: CP L
label_23619: CP $12
label_23620: CP (HL)
label_23621: CP (IX + 127)
label_23622: CP (IY - 128)
label_23623: INC A
label_23624: INC B
label_23625: INC C
label_23626: INC D
label_23627: INC E
label_23628: INC H
label_23629: INC L
label_23630: INC (HL)
label_23631: INC (IX + 127)
label_23632: INC (IY - 128)
label_23633: DEC A
label_23634: DEC B
label_23635: DEC C
label_23636: DEC D
label_23637: DEC E
label_23638: DEC H
label_23639: DEC L
label_23640: DEC (HL)
label_23641: DEC (IX + 127)
label_23642: DEC (IY - 128)
label_23643: DAA
label_23644: CPL
label_23645: NEG
label_23646: CCF
label_23647: SCF
label_23648: NOP
label_23649: HALT
label_23650: DI
label_23651: EI
label_23652: IM 0
label_23653: IM 1
label_23654: IM 2
label_23655: ADD HL,BC
label_23656: ADD HL,DE
label_23657: ADD HL,HL
label_23658: ADD HL,SP
label_23659: ADC HL,BC
label_23660: ADC HL,DE
label_23661: ADC HL,HL
label_23662: ADC HL,SP
label_23663: SBC HL,BC
label_23664: SBC HL,DE
label_23665: SBC HL,HL
label_23666: SBC HL,SP
label_23667: ADD IX,BC
label_23668: ADD IX,DE
label_23669: ADD IX,SP
label_23670: ADD IY,BC
label_23671: ADD IY,DE
label_23672: ADD IY,SP
label_23673: INC BC
label_23674: INC DE
label_23675: INC HL
label_23676: INC SP
label_23677: INC IX
label_23678: INC IY
label_23679: DEC BC
label_23680: DEC DE
label_23681: DEC HL
label_23682: DEC SP
label_23683: DEC IX
label_23684: DEC IY
label_23685: RLCA
label_23686: RLA
label_23687: RRCA
label_23688: RRA
label_23689: RLC A
label_23690: RLC B
label_23691: RLC C
label_23692: RLC D
label_23693: RLC E
label_23694: RLC H
label_23695: RLC L
label_23696: RLC (HL)
label_23697: RLC (IX + 127)
label_23698: RLC (IY - 128)
label_23699: RL A
label_23700: RL B
label_23701: RL C
label_23702: RL D
label_23703: RL E
label_23704: RL H
label_23705: RL L
label_23706: RL (HL)
label_23707: RL (IX + 127)
label_23708: RL (IY - 128)
label_23709: RRC A
label_23710: RRC B
label_23711: RRC C
label_23712: RRC D
label_23713: RRC E
label_23714: RRC H
label_23715: RRC L
label_23716: RRC (HL)
label_23717: RRC (IX + 127)
label_23718: RRC (IY - 128)
label_23719: RR A
label_23720: RR B
label_23721: RR C
label_23722: RR D
label_23723: RR E
label_23724: RR H
label_23725: RR L
label_23726: RR (HL)
label_23727: RR (IX + 127)
label_23728: RR (IY - 128)
label_23729: SLA A
label_23730: SLA B
label_23731: SLA C
label_23732: SLA D
label_23733: SLA E
label_23734: SLA H
label_23735: SLA L
label_23736: SLA (HL)
label_23737: SLA (IX + 127)
label_23738: SLA (IY - 128)
label_23739: SRA A
label_23740: SRA B
label_23741: SRA C
label_23742: SRA D
label_23743: SRA E
label_23744: SRA H
label_23745: SRA L
label_23746: SRA (HL)
label_23747: SRA (IX + 127)
label_23748: SRA (IY - 128)
label_23749: SRL A
label_23750: SRL B
label_23751: SRL C
label_23752: SRL D
label_23753: SRL E
label_23754: SRL H
label_23755: SRL L
label_23756: SRL (HL)
label_23757: SRL (IX + 127)
label_23758: SRL (IY - 128)
label_23759: RLD
label_23760: RRD
label_23761: BIT 0,A
label_23762: BIT 1,A
label_23763: BIT 2,A
label_23764: BIT 3,A
label_23765: BIT 4,A
label_23766: BIT 5,A
label_23767: BIT 6,A
label_23768: BIT 7,A
label_23769: BIT 0,B
label_23770: BIT 1,B
label_23771: BIT 2,B
label_23772: BIT 3,B
label_23773: BIT 4,B
label_23774: BIT 5,B
label_23775: BIT 6,B
label_23776: BIT 7,B
label_23777: BIT 0,C
label_23778: BIT 1,C
label_23779: BIT 2,C
label_23780: BIT 3,C
label_23781: BIT 4,C
label_23782: BIT 5,C
label_23783: BIT 6,C
label_23784: BIT 7,C
label_23785: BIT 0,D
label_23786: BIT 1,D
label_23787: BIT 2,D
label_23788: BIT 3,D
label_23789: BIT 4,D
label_23790: BIT 5,D
label_23791: BIT 6,D
label_23792: BIT 7,D
label_23793: BIT 0,E
label_23794: BIT 1,E
label_23795: BIT 2,E
label_23796: BIT 3,E
label_23797: BIT 4,E
label_23798: BIT 5,E
label_23799: BIT 6,E
label_23800: BIT 7,E
label_23801: BIT 0,H
label_23802: BIT 1,H
label_23803: BIT 2,H
label_23804: BIT 3,H
label_23805: BIT 4,H
label_23806: BIT 5,H
label_23807: BIT 6,H
label_23808: BIT 7,H
label_23809: BIT 0,L
label_23810: BIT 1,L
label_23811: BIT 2,L
label_23812: BIT 3,L
label_23813: BIT 4,L
label_23814: BIT 5,L
label_23815: BIT 6,L
label_23816: BIT 7,L
label_23817: BIT 0,(HL)
label_23818: BIT 1,(HL)
label_23819: BIT 2,(HL)
label_23820: BIT 3,(HL)
label_23821: BIT 4,(HL)
label_23822: BIT 5,(HL)
label_23823: BIT 6,(HL)
label_23824: BIT 7,(HL)
label_23825: BIT 0,(IX + 127)
label_23826: BIT 1,(IX + 127)
label_23827: BIT 2,(IX + 127)
label_23828: BIT 3,(IX + 127)
label_23829: BIT 4,(IX + 127)
label_23830: BIT 5,(IX + 127)
label_23831: BIT 6,(IX + 127)
label_23832: BIT 7,(IX + 127)
label_23833: BIT 0,(IY - 128)
label_23834: BIT 1,(IY - 128)
label_23835: BIT 2,(IY - 128)
label_23836: BIT 3,(IY - 128)
label_23837: BIT 4,(IY - 128)
label_23838: BIT 5,(IY - 128)
label_23839: BIT 6,(IY - 128)
label_23840: BIT 7,(IY - 128)
label_23841: SET 0,A
label_23842: SET 1,A
label_23843: SET 2,A
label_23844: SET 3,A
label_23845: SET 4,A
label_23846: SET 5,A
label_23847: SET 6,A
label_23848: SET 7,A
label_23849: SET 0,B
label_23850: SET 1,B
label_23851: SET 2,B
label_23852: SET 3,B
label_23853: SET 4,B
label_23854: SET 5,B
label_23855: SET 6,B
label_23856: SET 7,B
label_23857: SET 0,C
label_23858: SET 1,C
label_23859: SET 2,C
label_23860: SET 3,C
label_23861: SET 4,C
label_23862: SET 5,C
label_23863: SET 6,C
label_23864: SET 7,C
label_23865: SET 0,D
label_23866: SET 1,D
label_23867: SET 2,D
label_23868: SET 3,D
label_23869: SET 4,D
label_23870: SET 5,D
label_23871: SET 6,D
label_23872: SET 7,D
label_23873: SET 0,E
label_23874: SET 1,E
label_23875: SET 2,E
label_23876: SET 3,E
label_23877: SET 4,E
label_23878: SET 5,E
label_23879: SET 6,E
label_23880: SET 7,E
label_23881: SET 0,H
label_23882: SET 1,H
label_23883: SET 2,H
label_23884: SET 3,H
label_23885: SET 4,H
label_23886: SET 5,H
label_23887: SET 6,H
label_23888: SET 7,H
label_23889: SET 0,L
label_23890: SET 1,L
label_23891: SET 2,L
label_23892: SET 3,L
label_23893: SET 4,L
label_23894: SET 5,L
label_23895: SET 6,L
label_23896: SET 7,L
label_23897: SET 0,(HL)
label_23898: SET 1,(HL)
label_23899: SET 2,(HL)
label_23900: SET 3,(HL)
label_23901: SET 4,(HL)
label_23902: SET 5,(HL)
label_23903: SET 6,(HL)
label_23904: SET 7,(HL)
label_23905: SET 0,(IX + 127)
label_23906: SET 1,(IX + 127)
label_23907: SET 2,(IX + 127)
label_23908: SET 3,(IX + 127)
label_23909: SET 4,(IX + 127)
label_23910: SET 5,(IX + 127)
label_23911: SET 6,(IX + 127)
label_23912: SET 7,(IX + 127)
label_23913: SET 0,(IY - 128)
label_23914: SET 1,(IY - 128)
label_23915: SET 2,(IY - 128)
label_23916: SET 3,(IY - 128)
label_23917: SET 4,(IY - 128)
label_23918: SET 5,(IY - 128)
label_23919: SET 6,(IY - 128)
label_23920: SET 7,(IY - 128)
label_23921: RES 0,A
label_23922: RES 1,A
label_23923: RES 2,A
label_23924: RES 3,A
label_23925: RES 4,A
label_23926: RES 5,A
label_23927: RES 6,A
label_23928: RES 7,A
label_23929: RES 0,B
label_23930: RES 1,B
label_23931: RES 2,B
label_23932: RES 3,B
label_23933: RES 4,B
label_23934: RES 5,B
label_23935: RES 6,B
label_23936: RES 7,B
label_23937: RES 0,C
label_23938: RES 1,C
label_23939: RES 2,C
label_23940: RES 3,C
label_23941: RES 4,C
label_23942: RES 5,C
label_23943: RES 6,C
label_23944: RES 7,C
label_23945: RES 0,D
label_23946: RES 1,D
label_23947: RES 2,D
label_23948: RES 3,D
label_23949: RES 4,D
label_23950: RES 5,D
label_23951: RES 6,D
label_23952: RES 7,D
label_23953: RES 0,E
label_23954: RES 1,E
label_23955: RES 2,E
label_23956: RES 3,E
label_23957: RES 4,E
label_23958: RES 5,E
label_23959: RES 6,E
label_23960: RES 7,E
label_23961: RES 0,H
label_23962: RES 1,H
label_23963: RES 2,H
label_23964: RES 3,H
label_23965: RES 4,H
label_23966: RES 5,H
label_23967: RES 6,H
label_23968: RES 7,H
label_23969: RES 0,L
label_23970: RES 1,L
label_23971: RES 2,L
label_23972: RES 3,L
label_23973: RES 4,L
label_23974: RES 5,L
label_23975: RES 6,L
label_23976: RES 7,L
label_23977: RES 0,(HL)
label_23978: RES 1,(HL)
label_23979: RES 2,(HL)
label_23980: RES 3,(HL)
label_23981: RES 4,(HL)
label_23982: RES 5,(HL)
label_23983: RES 6,(HL)
label_23984: RES 7,(HL)
label_23985: RES 0,(IX + 127)
label_23986: RES 1,(IX + 127)
label_23987: RES 2,(IX + 127)
label_23988: RES 3,(IX + 127)
label_23989: RES 4,(IX + 127)
label_23990: RES 5,(IX + 127)
label_23991: RES 6,(IX + 127)
label_23992: RES 7,(IX + 127)
label_23993: RES 0,(IY - 128)
label_23994: RES 1,(IY - 128)
label_23995: RES 2,(IY - 128)
label_23996: RES 3,(IY - 128)
label_23997: RES 4,(IY - 128)
label_23998: RES 5,(IY - 128)
label_23999: RES 6,(IY - 128)
label_24000: RES 7,(IY - 128)
label_24001: JP $5678
label_24002: JP NZ,$5678
label_24003: JP Z,$5678
label_24004: JP NC,$5678
label_24005: JP C,$5678
label_24006: JP PO,$5678
label_24007: JP PE,$5678
label_24008: JP P,$5678
label_24009: JP M,$5678
label_24010: JR $ + 2
label_24011: JR NZ,$ + 2
label_24012: JR Z,$ + 2
label_24013: JR NC,$ + 2
label_24014: JR C,$ + 2
label_24015: JP (HL)
label_24016: JP (IX)
label_24017: JP (IY)
label_24018: DJNZ $ + 2
label_24019: CALL $5678
label_24020: CALL NZ,$5678
label_24021: CALL Z,$5678
label_24022: CALL NC,$5678
label_24023: CALL C,$5678
label_24024: CALL PO,$5678
label_24025: CALL PE,$5678
label_24026: CALL P,$5678
label_24027: CALL M,$5678
label_24028: RET
label_24029: RET NZ
label_24030: RET Z
label_24031: RET NC
label_24032: RET C
label_24033: RET PO
label_24034: RET PE
label_24035: RET P
label_24036: RET M
label_24037: RETI
label_24038: RETN
label_24039: RST $00
label_24040: RST $08
label_24041: RST $10
label_24042: RST $18
label_24043: RST $20
label_24044: RST $28
label_24045: RST $30
label_24046: RST $38
label_24047: IN A,($12)
label_24048: IN A,(C)
label_24049: IN B,(C)
label_24050: IN C,(C)
label_24051: IN D,(C)
label_24052: IN E,(C)
label_24053: IN H,(C)
label_24054: IN L,(C)
label_24055: IN F,(C)
label_24056: INI
label_24057: INIR
label_24058: IND
label_24059: INDR
label_24060: OUT ($12),A
label_24061: OUT (C),A
label_24062: OUT (C),B
label_24063: OUT (C),C
label_24064: OUT (C),D
label_24065: OUT (C),E
label_24066: OUT (C),H
label_24067: OUT (C),L
label_24068: OUTI
label_24069: OTIR
label_24070: OUTD
label_24071: OTDR
label_24072: LD A,A
label_24073: LD A,B
label_24074: LD A,C
label_24075: LD A,D
label_24076: LD A,E
label_24077: LD A,H
label_24078: LD A,L
label_24079: LD B,A
label_24080: LD B,B
label_24081: LD B,C
label_24082: LD B,D
label_24083: LD B,E
label_24084: LD B,H
label_24085: LD B,L
label_24086: LD C,A
label_24087: LD C,B
label_24088: LD C,C
label_24089: LD C,D
label_24090: LD C,E
label_24091: LD C,H
label_24092: LD C,L
label_24093: LD D,A
label_24094: LD D,B
label_24095: LD D,C
label_24096: LD D,D
label_24097: LD D,E
label_24098: LD D,H
label_24099: LD D,L
label_24100: LD E,A
label_24101: LD E,B
label_24102: LD E,C
label_24103: LD E,D
label_24104: LD E,E
label_24105: LD E,H
label_24106: LD E,L
label_24107: LD H,A
label_24108: LD H,B
label_24109: LD H,C
label_24110: LD H,D
label_24111: LD H,E
label_24112: LD H,H
label_24113: LD H,L
label_24114: LD L,A
label_24115: LD L,B
label_24116: LD L,C
label_24117: LD L,D
label_24118: LD L,E
label_24119: LD L,H
label_24120: LD L,L
label_24121: LD A,$12
label_24122: LD B,$12
label_24123: LD C,$12
label_24124: LD D,$12
label_24125: LD E,$12
label_24126: LD H,$12
label_24127: LD L,$12
label_24128: LD A,(HL)
label_24129: LD B,(HL)
label_24130: LD C,(HL)
label_24131: LD D,(HL)
label_24132: LD E,(HL)
label_24133: LD H,(HL)
label_24134: LD L,(HL)
label_24135: LD A,(IX + 127)
label_24136: LD B,(IX + 127)
label_24137: LD C,(IX + 127)
label_24138: LD D,(IX + 127)
label_24139: LD E,(IX + 127)
label_24140: LD H,(IX + 127)
label_24141: LD L,(IX + 127)
label_24142: LD A,(IY - 128)
label_24143: LD B,(IY - 128)
label_24144: LD C,(IY - 128)
label_24145: LD D,(IY - 128)
label_24146: LD E,(IY - 128)
label_24147: LD H,(IY - 128)
label_24148: LD L,(IY - 128)
label_24149: LD (HL),A
label_24150: LD (HL),B
label_24151: LD (HL),C
label_24152: LD (HL),D
label_24153: LD (HL),E
label_24154: LD (HL),H
label_24155: LD (HL),L
label_24156: LD (IX + 127),A
label_24157: LD (IX + 127),B
label_24158: LD (IX + 127),C
label_24159: LD (IX + 127),D
label_24160: LD (IX + 127),E
label_24161: LD (IX + 127),H
label_24162: LD (IX + 127),L
label_24163: LD (IY - 128),A
label_24164: LD (IY - 128),B
label_24165: LD (IY - 128),C
label_24166: LD (IY - 128),D
label_24167: LD (IY - 128),E
label_24168: LD (IY - 128),H
label_24169: LD (IY - 128),L
label_24170: LD (HL),$12
label_24171: LD (IX + 127),$12
label_24172: LD (IY - 128),$12
label_24173: LD A,(BC)
label_24174: LD A,(DE)
label_24175: LD A,($5678)
label_24176: LD (BC),A
label_24177: LD (DE),A
label_24178: LD ($5678),A
label_24179: LD A,I
label_24180: LD A,R
label_24181: LD I,A
label_24182: LD R,A
label_24183: LD BC,$5678
label_24184: LD DE,$5678
label_24185: LD HL,$5678
label_24186: LD SP,$5678
label_24187: LD IX,$5678
label_24188: LD IY,$5678
label_24189: LD HL,($5678)
label_24190: LD BC,($5678)
label_24191: LD DE,($5678)
label_24192: LD HL,($5678)
label_24193: LD SP,($5678)
label_24194: LD IX,($5678)
label_24195: LD IY,($5678)
label_24196: LD ($5678),HL
label_24197: LD ($5678),BC
label_24198: LD ($5678),DE
label_24199: LD ($5678),HL
label_24200: LD ($5678),SP
label_24201: LD ($5678),IX
label_24202: LD ($5678),IY
label_24203: LD SP,HL
label_24204: LD SP,IX
label_24205: LD SP,IY
label_24206: PUSH BC
label_24207: PUSH DE
label_24208: PUSH HL
label_24209: PUSH AF
label_24210: PUSH IX
label_24211: PUSH IY
label_24212: POP BC
label_24213: POP DE
label_24214: POP HL
label_24215: POP AF
label_24216: POP IX
label_24217: POP IY
label_24218: EX DE,HL
label_24219: EX AF,AF'
label_24220: EXX
label_24221: EX (SP),HL
label_24222: EX (SP),IX
label_24223: EX (SP),IY
label_24224: LDI
label_24225: LDIR
label_24226: LDD
label_24227: LDDR
label_24228: CPI
label_24229: CPIR
label_24230: CPD
label_24231: CPDR
label_24232: ADD A,A
label_24233: ADD A,B
label_24234: ADD A,C
label_24235: ADD A,D
label_24236: ADD A,E
label_24237: ADD A,H
label_24238: ADD A,L
label_24239: ADD A,$12
label_24240: ADD A,(HL)
label_24241: ADD A,(IX + 127)
label_24242: ADD A,(IY - 128)
label_24243: ADC A,A
label_24244: ADC A,B
label_24245: ADC A,C
label_24246: ADC A,D
label_24247: ADC A,E
label_24248: ADC A,H
label_24249: ADC A,L
label_24250: ADC A,$12
label_24251: ADC A,(HL)
label_24252: ADC A,(IX + 127)
label_24253: ADC A,(IY - 128)
label_24254: SUB A
label_24255: SUB B
label_24256: SUB C
label_24257: SUB D
label_24258: SUB E
label_24259: SUB H
label_24260: SUB L
label_24261: SUB $12
label_24262: SUB (HL)
label_24263: SUB (IX + 127)
label_24264: SUB (IY - 128)
label_24265: SBC A,A
label_24266: SBC A,B
label_24267: SBC A,C
label_24268: SBC A,D
label_24269: SBC A,E
label_24270: SBC A,H
label_24271: SBC A,L
label_24272: SBC A,$12
label_24273: SBC A,(HL)
label_24274: SBC A,(IX + 127)
label_24275: SBC A,(IY - 128)
label_24276: AND A
label_24277: AND B
label_24278: AND C
label_24279: AND D
label_24280: AND E
label_24281: AND H
label_24282: AND L
label_24283: AND $12
label_24284: AND (HL)
label_24285: AND (IX + 127)
label_24286: AND (IY - 128)
label_24287: AND A
label_24288: AND B
label_24289: AND C
label_24290: AND D
label_24291: AND E
label_24292: AND H
label_24293: AND L
label_24294: AND $12
label_24295: AND (HL)
label_24296: AND (IX + 127)
label_24297: AND (IY - 128)
label_24298: OR A
label_24299: OR B
label_24300: OR C
label_24301: OR D
label_24302: OR E
label_24303: OR H
label_24304: OR L
label_24305: OR $12
label_24306: OR (HL)
label_24307: OR (IX + 127)
label_24308: OR (IY - 128)
label_24309: XOR A
label_24310: XOR B
label_24311: XOR C
label_24312: XOR D
label_24313: XOR E
label_24314: XOR H
label_24315: XOR L
label_24316: XOR $12
label_24317: XOR (HL)
label_24318: XOR (IX + 127)
label_24319: XOR (IY - 128)
label_24320: CP A
label_24321: CP B
label_24322: CP C
label_24323: CP D
label_24324: CP E
label_24325: CP H
label_24326: CP L
label_24327: CP $12
label_24328: CP (HL)
label_24329: CP (IX + 127)
label_24330: CP (IY - 128)
label_24331: INC A
label_24332: INC B
label_24333: INC C
label_24334: INC D
label_24335: INC E
label_24336: INC H
label_24337: INC L
label_24338: INC (HL)
label_24339: INC (IX + 127)
label_24340: INC (IY - 128)
label_24341: DEC A
label_24342: DEC B
label_24343: DEC C
label_24344: DEC D
label_24345: DEC E
label_24346: DEC H
label_24347: DEC L
label_24348: DEC (HL)
label_24349: DEC (IX + 127)
label_24350: DEC (IY - 128)
label_24351: DAA
label_24352: CPL
label_24353: NEG
label_24354: CCF
label_24355: SCF
label_24356: NOP
label_24357: HALT
label_24358: DI
label_24359: EI
label_24360: IM 0
label_24361: IM 1
label_24362: IM 2
label_24363: ADD HL,BC
label_24364: ADD HL,DE
label_24365: ADD HL,HL
label_24366: ADD HL,SP
label_24367: ADC HL,BC
label_24368: ADC HL,DE
label_24369: ADC HL,HL
label_24370: ADC HL,SP
label_24371: SBC HL,BC
label_24372: SBC HL,DE
label_24373: SBC HL,HL
label_24374: SBC HL,SP
label_24375: ADD IX,BC
label_24376: ADD IX,DE
label_24377: ADD IX,SP
label_24378: ADD IY,BC
label_24379: ADD IY,DE
label_24380: ADD IY,SP
label_24381: INC BC
label_24382: INC DE
label_24383: INC HL
label_24384: INC SP
label_24385: INC IX
label_24386: INC IY
label_24387: DEC BC
label_24388: DEC DE
label_24389: DEC HL
label_24390: DEC SP
label_24391: DEC IX
label_24392: DEC IY
label_24393: RLCA
label_24394: RLA
label_24395: RRCA
label_24396: RRA
label_24397: RLC A
label_24398: RLC B
label_24399: RLC C
label_24400: RLC D
label_24401: RLC E
label_24402: RLC H
label_24403: RLC L
label_24404: RLC (HL)
label_24405: RLC (IX + 127)
label_24406: RLC (IY - 128)
label_24407: RL A
label_24408: RL B
label_24409: RL C
label_24410: RL D
label_24411: RL E
label_24412: RL H
label_24413: RL L
label_24414: RL (HL)
label_24415: RL (IX + 127)
label_24416: RL (IY - 128)
label_24417: RRC A
label_24418: RRC B
label_24419: RRC C
label_24420: RRC D
label_24421: RRC E
label_24422: RRC H
label_24423: RRC L
label_24424: RRC (HL)
label_24425: RRC (IX + 127)
label_24426: RRC (IY - 128)
label_24427: RR A
label_24428: RR B
label_24429: RR C
label_24430: RR D
label_24431: RR E
label_24432: RR H
label_24433: RR L
label_24434: RR (HL)
label_24435: RR (IX + 127)
label_24436: RR (IY - 128)
label_24437: SLA A
label_24438: SLA B
label_24439: SLA C
label_24440: SLA D
label_24441: SLA E
label_24442: SLA H
label_24443: SLA L
label_24444: SLA (HL)
label_24445: SLA (IX + 127)
label_24446: SLA (IY - 128)
label_24447: SRA A
label_24448: SRA B
label_24449: SRA C
label_24450: SRA D
label_24451: SRA E
label_24452: SRA H
label_24453: SRA L
label_24454: SRA (HL)
label_24455: SRA (IX + 127)
label_24456: SRA (IY - 128)
label_24457: SRL A
label_24458: SRL B
label_24459: SRL C
label_24460: SRL D
label_24461: SRL E
label_24462: SRL H
label_24463: SRL L
label_24464: SRL (HL)
label_24465: SRL (IX + 127)
label_24466: SRL (IY - 128)
label_24467: RLD
label_24468: RRD
label_24469: BIT 0,A
label_24470: BIT 1,A
label_24471: BIT 2,A
label_24472: BIT 3,A
label_24473: BIT 4,A
label_24474: BIT 5,A
label_24475: BIT 6,A
label_24476: BIT 7,A
label_24477: BIT 0,B
label_24478: BIT 1,B
label_24479: BIT 2,B
label_24480: BIT 3,B
label_24481: BIT 4,B
label_24482: BIT 5,B
label_24483: BIT 6,B
label_24484: BIT 7,B
label_24485: BIT 0,C
label_24486: BIT 1,C
label_24487: BIT 2,C
label_24488: BIT 3,C
label_24489: BIT 4,C
label_24490: BIT 5,C
label_24491: BIT 6,C
label_24492: BIT 7,C
label_24493: BIT 0,D
label_24494: BIT 1,D
label_24495: BIT 2,D
label_24496: BIT 3,D
label_24497: BIT 4,D
label_24498: BIT 5,D
label_24499: BIT 6,D
label_24500: BIT 7,D
label_24501: BIT 0,E
label_24502: BIT 1,E
label_24503: BIT 2,E
label_24504: BIT 3,E
label_24505: BIT 4,E
label_24506: BIT 5,E
label_24507: BIT 6,E
label_24508: BIT 7,E
label_24509: BIT 0,H
label_24510: BIT 1,H
label_24511: BIT 2,H
label_24512: BIT 3,H
label_24513: BIT 4,H
label_24514: BIT 5,H
label_24515: BIT 6,H
label_24516: BIT 7,H
label_24517: BIT 0,L
label_24518: BIT 1,L
label_24519: BIT 2,L
label_24520: BIT 3,L
label_24521: BIT 4,L
label_24522: BIT 5,L
label_24523: BIT 6,L
label_24524: BIT 7,L
label_24525: BIT 0,(HL)
label_24526: BIT 1,(HL)
label_24527: BIT 2,(HL)
label_24528: BIT 3,(HL)
label_24529: BIT 4,(HL)
label_24530: BIT 5,(HL)
label_24531: BIT 6,(HL)
label_24532: BIT 7,(HL)
label_24533: BIT 0,(IX + 127)
label_24534: BIT 1,(IX + 127)
label_24535: BIT 2,(IX + 127)
label_24536: BIT 3,(IX + 127)
label_24537: BIT 4,(IX + 127)
label_24538: BIT 5,(IX + 127)
label_24539: BIT 6,(IX + 127)
label_24540: BIT 7,(IX + 127)
label_24541: BIT 0,(IY - 128)
label_24542: BIT 1,(IY - 128)
label_24543: BIT 2,(IY - 128)
label_24544: BIT 3,(IY - 128)
label_24545: BIT 4,(IY - 128)
label_24546: BIT 5,(IY - 128)
label_24547: BIT 6,(IY - 128)
label_24548: BIT 7,(IY - 128)
label_24549: SET 0,A
label_24550: SET 1,A
label_24551: SET 2,A
label_24552: SET 3,A
label_24553: SET 4,A
label_24554: SET 5,A
label_24555: SET 6,A
label_24556: SET 7,A
label_24557: SET 0,B
label_24558: SET 1,B
label_24559: SET 2,B
label_24560: SET 3,B
label_24561: SET 4,B
label_24562: SET 5,B
label_24563: SET 6,B
label_24564: SET 7,B
label_24565: SET 0,C
label_24566: SET 1,C
label_24567: SET 2,C
label_24568: SET 3,C
label_24569: SET 4,C
label_24570: SET 5,C
label_24571: SET 6,C
label_24572: SET 7,C
label_24573: SET 0,D
label_24574: SET 1,D
label_24575: SET 2,D
label_24576: SET 3,D
label_24577: SET 4,D
label_24578: SET 5,D
label_24579: SET 6,D
label_24580: SET 7,D
label_24581: SET 0,E
label_24582: SET 1,E
label_24583: SET 2,E
label_24584: SET 3,E
label_24585: SET 4,E
label_24586: SET 5,E
label_24587: SET 6,E
label_24588: SET 7,E
label_24589: SET 0,H
label_24590: SET 1,H
label_24591: SET 2,H
label_24592: SET 3,H
label_24593: SET 4,H
label_24594: SET 5,H
label_24595: SET 6,H
label_24596: SET 7,H
label_24597: SET 0,L
label_24598: SET 1,L
label_24599: SET 2,L
label_24600: SET 3,L
label_24601: SET 4,L
label_24602: SET 5,L
label_24603: SET 6,L
label_24604: SET 7,L
label_24605: SET 0,(HL)
label_24606: SET 1,(HL)
label_24607: SET 2,(HL)
label_24608: SET 3,(HL)
label_24609: SET 4,(HL)
label_24610: SET 5,(HL)
label_24611: SET 6,(HL)
label_24612: SET 7,(HL)
label_24613: SET 0,(IX + 127)
label_24614: SET 1,(IX + 127)
label_24615: SET 2,(IX + 127)
label_24616: SET 3,(IX + 127)
label_24617: SET 4,(IX + 127)
label_24618: SET 5,(IX + 127)
label_24619: SET 6,(IX + 127)
label_24620: SET 7,(IX + 127)
label_24621: SET 0,(IY - 128)
label_24622: SET 1,(IY - 128)
label_24623: SET 2,(IY - 128)
label_24624: SET 3,(IY - 128)
label_24625: SET 4,(IY - 128)
label_24626: SET 5,(IY - 128)
label_24627: SET 6,(IY - 128)
label_24628: SET 7,(IY - 128)
label_24629: RES 0,A
label_24630: RES 1,A
label_24631: RES 2,A
label_24632: RES 3,A
label_24633: RES 4,A
label_24634: RES 5,A
label_24635: RES 6,A
label_24636: RES 7,A
label_24637: RES 0,B
label_24638: RES 1,B
label_24639: RES 2,B
label_24640: RES 3,B
label_24641: RES 4,B
label_24642: RES 5,B
label_24643: RES 6,B
label_24644: RES 7,B
label_24645: RES 0,C
label_24646: RES 1,C
label_24647: RES 2,C
label_24648: RES 3,C
label_24649: RES 4,C
label_24650: RES 5,C
label_24651: RES 6,C
label_24652: RES 7,C
label_24653: RES 0,D
label_24654: RES 1,D
label_24655: RES 2,D
label_24656: RES 3,D
label_24657: RES 4,D
label_24658: RES 5,D
label_24659: RES 6,D
label_24660: RES 7,D
label_24661: RES 0,E
label_24662: RES 1,E
label_24663: RES 2,E
label_24664: RES 3,E
label_24665: RES 4,E
label_24666: RES 5,E
label_24667: RES 6,E
label_24668: RES 7,E
label_24669: RES 0,H
label_24670: RES 1,H
label_24671: RES 2,H
label_24672: RES 3,H
label_24673: RES 4,H
label_24674: RES 5,H
label_24675: RES 6,H
label_24676: RES 7,H
label_24677: RES 0,L
label_24678: RES 1,L
label_24679: RES 2,L
label_24680: RES 3,L
label_24681: RES 4,L
label_24682: RES 5,L
label_24683: RES 6,L
label_24684: RES 7,L
label_24685: RES 0,(HL)
label_24686: RES 1,(HL)
label_24687: RES 2,(HL)
label_24688: RES 3,(HL)
label_24689: RES 4,(HL)
label_24690: RES 5,(HL)
label_24691: RES 6,(HL)
label_24692: RES 7,(HL)
label_24693: RES 0,(IX + 127)
label_24694: RES 1,(IX + 127)
label_24695: RES 2,(IX + 127)
label_24696: RES 3,(IX + 127)
label_24697: RES 4,(IX + 127)
label_24698: RES 5,(IX + 127)
label_24699: RES 6,(IX + 127)
label_24700: RES 7,(IX + 127)
label_24701: RES 0,(IY - 128)
label_24702: RES 1,(IY - 128)
label_24703: RES 2,(IY - 128)
label_24704: RES 3,(IY - 128)
label_24705: RES 4,(IY - 128)
label_24706: RES 5,(IY - 128)
label_24707: RES 6,(IY - 128)
label_24708: RES 7,(IY - 128)
label_24709: JP $5678
label_24710: JP NZ,$5678
label_24711: JP Z,$5678
label_24712: JP NC,$5678
label_24713: JP C,$5678
label_24714: JP PO,$5678
label_24715: JP PE,$5678
label_24716: JP P,$5678
label_24717: JP M,$5678
label_24718: JR $ + 2
label_24719: JR NZ,$ + 2
label_24720: JR Z,$ + 2
label_24721: JR NC,$ + 2
label_24722: JR C,$ + 2
label_24723: JP (HL)
label_24724: JP (IX)
label_24725: JP (IY)
label_24726: DJNZ $ + 2
label_24727: CALL $5678
label_24728: CALL NZ,$5678
label_24729: CALL Z,$5678
label_24730: CALL NC,$5678
label_24731: CALL C,$5678
label_24732: CALL PO,$5678
label_24733: CALL PE,$5678
label_24734: CALL P,$5678
label_24735: CALL M,$5678
label_24736: RET
label_24737: RET NZ
label_24738: RET Z
label_24739: RET NC
label_24740: RET C
label_24741: RET PO
label_24742: RET PE
label_24743: RET P
label_24744: RET M
label_24745: RETI
label_24746: RETN
label_24747: RST $00
label_24748: RST $08
label_24749: RST $10
label_24750: RST $18
label_24751: RST $20
label_24752: RST $28
label_24753: RST $30
label_24754: RST $38
label_24755: IN A,($12)
label_24756: IN A,(C)
label_24757: IN B,(C)
label_24758: IN C,(C)
label_24759: IN D,(C)
label_24760: IN E,(C)
label_24761: IN H,(C)
label_24762: IN L,(C)
label_24763: IN F,(C)
label_24764: INI
label_24765: INIR
label_24766: IND
label_24767: INDR
label_24768: OUT ($12),A
label_24769: OUT (C),A
label_24770: OUT (C),B
label_24771: OUT (C),C
label_24772: OUT (C),D
label_24773: OUT (C),E
label_24774: OUT (C),H
label_24775: OUT (C),L
label_24776: OUTI
label_24777: OTIR
label_24778: OUTD
label_24779: OTDR
label_24780: LD A,A
label_24781: LD A,B
label_24782: LD A,C
label_24783: LD A,D
label_24784: LD A,E
label_24785: LD A,H
label_24786: LD A,L
label_24787: LD B,A
label_24788: LD B,B
label_24789: LD B,C
label_24790: LD B,D
label_24791: LD B,E
label_24792: LD B,H
label_24793: LD B,L
label_24794: LD C,A
label_24795: LD C,B
label_24796: LD C,C
label_24797: LD C,D
label_24798: LD C,E
label_24799: LD C,H
label_24800: LD C,L
label_24801: LD D,A
label_24802: LD D,B
label_24803: LD D,C
label_24804: LD D,D
label_24805: LD D,E
label_24806: LD D,H
label_24807: LD D,L
label_24808: LD E,A
label_24809: LD E,B
label_24810: LD E,C
label_24811: LD E,D
label_24812: LD E,E
label_24813: LD E,H
label_24814: LD E,L
label_24815: LD H,A
label_24816: LD H,B
label_24817: LD H,C
label_24818: LD H,D
label_24819: LD H,E
label_24820: LD H,H
label_24821: LD H,L
label_24822: LD L,A
label_24823: LD L,B
label_24824: LD L,C
label_24825: LD L,D
label_24826: LD L,E
label_24827: LD L,H
label_24828: LD L,L
label_24829: LD A,$12
label_24830: LD B,$12
label_24831: LD C,$12
label_24832: LD D,$12
label_24833: LD E,$12
label_24834: LD H,$12
label_24835: LD L,$12
label_24836: LD A,(HL)
label_24837: LD B,(HL)
label_24838: LD C,(HL)
label_24839: LD D,(HL)
label_24840: LD E,(HL)
label_24841: LD H,(HL)
label_24842: LD L,(HL)
label_24843: LD A,(IX + 127)
label_24844: LD B,(IX + 127)
label_24845: LD C,(IX + 127)
label_24846: LD D,(IX + 127)
label_24847: LD E,(IX + 127)
label_24848: LD H,(IX + 127)
label_24849: LD L,(IX + 127)
label_24850: LD A,(IY - 128)
label_24851: LD B,(IY - 128)
label_24852: LD C,(IY - 128)
label_24853: LD D,(IY - 128)
label_24854: LD E,(IY - 128)
label_24855: LD H,(IY - 128)
label_24856: LD L,(IY - 128)
label_24857: LD (HL),A
label_24858: LD (HL),B
label_24859: LD (HL),C
label_24860: LD (HL),D
label_24861: LD (HL),E
label_24862: LD (HL),H
label_24863: LD (HL),L
label_24864: LD (IX + 127),A
label_24865: LD (IX + 127),B
label_24866: LD (IX + 127),C
label_24867: LD (IX + 127),D
label_24868: LD (IX + 127),E
label_24869: LD (IX + 127),H
label_24870: LD (IX + 127),L
label_24871: LD (IY - 128),A
label_24872: LD (IY - 128),B
label_24873: LD (IY - 128),C
label_24874: LD (IY - 128),D
label_24875: LD (IY - 128),E
label_24876: LD (IY - 128),H
label_24877: LD (IY - 128),L
label_24878: LD (HL),$12
label_24879: LD (IX + 127),$12
label_24880: LD (IY - 128),$12
label_24881: LD A,(BC)
label_24882: LD A,(DE)
label_24883: LD A,($5678)
label_24884: LD (BC),A
label_24885: LD (DE),A
label_24886: LD ($5678),A
label_24887: LD A,I
label_24888: LD A,R
label_24889: LD I,A
label_24890: LD R,A
label_24891: LD BC,$5678
label_24892: LD DE,$5678
label_24893: LD HL,$5678
label_24894: LD SP,$5678
label_24895: LD IX,$5678
label_24896: LD IY,$5678
label_24897: LD HL,($5678)
label_24898: LD BC,($5678)
label_24899: LD DE,($5678)
label_24900: LD HL,($5678)
label_24901: LD SP,($5678)
label_24902: LD IX,($5678)
label_24903: LD IY,($5678)
label_24904: LD ($5678),HL
label_24905: LD ($5678),BC
label_24906: LD ($5678),DE
label_24907: LD ($5678),HL
label_24908: LD ($5678),SP
label_24909: LD ($5678),IX
label_24910: LD ($5678),IY
label_24911: LD SP,HL
label_24912: LD SP,IX
label_24913: LD SP,IY
label_24914: PUSH BC
label_24915: PUSH DE
label_24916: PUSH HL
label_24917: PUSH AF
label_24918: PUSH IX
label_24919: PUSH IY
label_24920: POP BC
label_24921: POP DE
label_24922: POP HL
label_24923: POP AF
label_24924: POP IX
label_24925: POP IY
label_24926: EX DE,HL
label_24927: EX AF,AF'
label_24928: EXX
label_24929: EX (SP),HL
label_24930: EX (SP),IX
label_24931: EX (SP),IY
label_24932: LDI
label_24933: LDIR
label_24934: LDD
label_24935: LDDR
label_24936: CPI
label_24937: CPIR
label_24938: CPD
label_24939: CPDR
label_24940: ADD A,A
label_24941: ADD A,B
label_24942: ADD A,C
label_24943: ADD A,D
label_24944: ADD A,E
label_24945: ADD A,H
label_24946: ADD A,L
label_24947: ADD A,$12
label_24948: ADD A,(HL)
label_24949: ADD A,(IX + 127)
label_24950: ADD A,(IY - 128)
label_24951: ADC A,A
label_24952: ADC A,B
label_24953: ADC A,C
label_24954: ADC A,D
label_24955: ADC A,E
label_24956: ADC A,H
label_24957: ADC A,L
label_24958: ADC A,$12
label_24959: ADC A,(HL)
label_24960: ADC A,(IX + 127)
label_24961: ADC A,(IY - 128)
label_24962: SUB A
label_24963: SUB B
label_24964: SUB C
label_24965: SUB D
label_24966: SUB E
label_24967: SUB H
label_24968: SUB L
label_24969: SUB $12
label_24970: SUB (HL)
label_24971: SUB (IX + 127)
label_24972: SUB (IY - 128)
label_24973: SBC A,A
label_24974: SBC A,B
label_24975: SBC A,C
label_24976: SBC A,D
label_24977: SBC A,E
label_24978: SBC A,H
label_24979: SBC A,L
label_24980: SBC A,$12
label_24981: SBC A,(HL)
label_24982: SBC A,(IX + 127)
label_24983: SBC A,(IY - 128)
label_24984: AND A
label_24985: AND B
label_24986: AND C
label_24987: AND D
label_24988: AND E
label_24989: AND H
label_24990: AND L
label_24991: AND $12
label_24992: AND (HL)
label_24993: AND (IX + 127)
label_24994: AND (IY - 128)
label_24995: AND A
label_24996: AND B
label_24997: AND C
label_24998: AND D
label_24999: AND E
label_25000: AND H
label_25001: AND L
label_25002: AND $12
label_25003: AND (HL)
label_25004: AND (IX + 127)
label_25005: AND (IY - 128)
label_25006: OR A
label_25007: OR B
label_25008: OR C
label_25009: OR D
label_25010: OR E
label_25011: OR H
label_25012: OR L
label_25013: OR $12
label_25014: OR (HL)
label_25015: OR (IX + 127)
label_25016: OR (IY - 128)
label_25017: XOR A
label_25018: XOR B
label_25019: XOR C
label_25020: XOR D
label_25021: XOR E
label_25022: XOR H
label_25023: XOR L
label_25024: XOR $12
label_25025: XOR (HL)
label_25026: XOR (IX + 127)
label_25027: XOR (IY - 128)
label_25028: CP A
label_25029: CP B
label_25030: CP C
label_25031: CP D
label_25032: CP E
label_25033: CP H
label_25034: CP L
label_25035: CP $12
label_25036: CP (HL)
label_25037: CP (IX + 127)
label_25038: CP (IY - 128)
label_25039: INC A
label_25040: INC B
label_25041: INC C
label_25042: INC D
label_25043: INC E
label_25044: INC H
label_25045: INC L
label_25046: INC (HL)
label_25047: INC (IX + 127)
label_25048: INC (IY - 128)
label_25049: DEC A
label_25050: DEC B
label_25051: DEC C
label_25052: DEC D
label_25053: DEC E
label_25054: DEC H
label_25055: DEC L
label_25056: DEC (HL)
label_25057: DEC (IX + 127)
label_25058: DEC (IY - 128)
label_25059: DAA
label_25060: CPL
label_25061: NEG
label_25062: CCF
label_25063: SCF
label_25064: NOP
label_25065: HALT
label_25066: DI
label_25067: EI
label_25068: IM 0
label_25069: IM 1
label_25070: IM 2
label_25071: ADD HL,BC
label_25072: ADD HL,DE
label_25073: ADD HL,HL
label_25074: ADD HL,SP
label_25075: ADC HL,BC
label_25076: ADC HL,DE
label_25077: ADC HL,HL
label_25078: ADC HL,SP
label_25079: SBC HL,BC
label_25080: SBC HL,DE
label_25081: SBC HL,HL
label_25082: SBC HL,SP
label_25083: ADD IX,BC
label_25084: ADD IX,DE
label_25085: ADD IX,SP
label_25086: ADD IY,BC
label_25087: ADD IY,DE
label_25088: ADD IY,SP
label_25089: INC BC
label_25090: INC DE
label_25091: INC HL
label_25092: INC SP
label_25093: INC IX
label_25094: INC IY
label_25095: DEC BC
label_25096: DEC DE
label_25097: DEC HL
label_25098: DEC SP
label_25099: DEC IX
label_25100: DEC IY
label_25101: RLCA
label_25102: RLA
label_25103: RRCA
label_25104: RRA
label_25105: RLC A
label_25106: RLC B
label_25107: RLC C
label_25108: RLC D
label_25109: RLC E
label_25110: RLC H
label_25111: RLC L
label_25112: RLC (HL)
label_25113: RLC (IX + 127)
label_25114: RLC (IY - 128)
label_25115: RL A
label_25116: RL B
label_25117: RL C
label_25118: RL D
label_25119: RL E
label_25120: RL H
label_25121: RL L
label_25122: RL (HL)
label_25123: RL (IX + 127)
label_25124: RL (IY - 128)
label_25125: RRC A
label_25126: RRC B
label_25127: RRC C
label_25128: RRC D
label_25129: RRC E
label_25130: RRC H
label_25131: RRC L
label_25132: RRC (HL)
label_25133: RRC (IX + 127)
label_25134: RRC (IY - 128)
label_25135: RR A
label_25136: RR B
label_25137: RR C
label_25138: RR D
label_25139: RR E
label_25140: RR H
label_25141: RR L
label_25142: RR (HL)
label_25143: RR (IX + 127)
label_25144: RR (IY - 128)
label_25145: SLA A
label_25146: SLA B
label_25147: SLA C
label_25148: SLA D
label_25149: SLA E
label_25150: SLA H
label_25151: SLA L
label_25152: SLA (HL)
label_25153: SLA (IX + 127)
label_25154: SLA (IY - 128)
label_25155: SRA A
label_25156: SRA B
label_25157: SRA C
label_25158: SRA D
label_25159: SRA E
label_25160: SRA H
label_25161: SRA L
label_25162: SRA (HL)
label_25163: SRA (IX + 127)
label_25164: SRA (IY - 128)
label_25165: SRL A
label_25166: SRL B
label_25167: SRL C
label_25168: SRL D
label_25169: SRL E
label_25170: SRL H
label_25171: SRL L
label_25172: SRL (HL)
label_25173: SRL (IX + 127)
label_25174: SRL (IY - 128)
label_25175: RLD
label_25176: RRD
label_25177: BIT 0,A
label_25178: BIT 1,A
label_25179: BIT 2,A
label_25180: BIT 3,A
label_25181: BIT 4,A
label_25182: BIT 5,A
label_25183: BIT 6,A
label_25184: BIT 7,A
label_25185: BIT 0,B
label_25186: BIT 1,B
label_25187: BIT 2,B
label_25188: BIT 3,B
label_25189: BIT 4,B
label_25190: BIT 5,B
label_25191: BIT 6,B
label_25192: BIT 7,B
label_25193: BIT 0,C
label_25194: BIT 1,C
label_25195: BIT 2,C
label_25196: BIT 3,C
label_25197: BIT 4,C
label_25198: BIT 5,C
label_25199: BIT 6,C
label_25200: BIT 7,C
label_25201: BIT 0,D
label_25202: BIT 1,D
label_25203: BIT 2,D
label_25204: BIT 3,D
label_25205: BIT 4,D
label_25206: BIT 5,D
label_25207: BIT 6,D
label_25208: BIT 7,D
label_25209: BIT 0,E
label_25210: BIT 1,E
label_25211: BIT 2,E
label_25212: BIT 3,E
label_25213: BIT 4,E
label_25214: BIT 5,E
label_25215: BIT 6,E
label_25216: BIT 7,E
label_25217: BIT 0,H
label_25218: BIT 1,H
label_25219: BIT 2,H
label_25220: BIT 3,H
label_25221: BIT 4,H
label_25222: BIT 5,H
label_25223: BIT 6,H
label_25224: BIT 7,H
label_25225: BIT 0,L
label_25226: BIT 1,L
label_25227: BIT 2,L
label_25228: BIT 3,L
label_25229: BIT 4,L
label_25230: BIT 5,L
label_25231: BIT 6,L
label_25232: BIT 7,L
label_25233: BIT 0,(HL)
label_25234: BIT 1,(HL)
label_25235: BIT 2,(HL)
label_25236: BIT 3,(HL)
label_25237: BIT 4,(HL)
label_25238: BIT 5,(HL)
label_25239: BIT 6,(HL)
label_25240: BIT 7,(HL)
label_25241: BIT 0,(IX + 127)
label_25242: BIT 1,(IX + 127)
label_25243: BIT 2,(IX + 127)
label_25244: BIT 3,(IX + 127)
label_25245: BIT 4,(IX + 127)
label_25246: BIT 5,(IX + 127)
label_25247: BIT 6,(IX + 127)
label_25248: BIT 7,(IX + 127)
label_25249: BIT 0,(IY - 128)
label_25250: BIT 1,(IY - 128)
label_25251: BIT 2,(IY - 128)
label_25252: BIT 3,(IY - 128)
label_25253: BIT 4,(IY - 128)
label_25254: BIT 5,(IY - 128)
label_25255: BIT 6,(IY - 128)
label_25256: BIT 7,(IY - 128)
label_25257: SET 0,A
label_25258: SET 1,A
label_25259: SET 2,A
label_25260: SET 3,A
label_25261: SET 4,A
label_25262: SET 5,A
label_25263: SET 6,A
label_25264: SET 7,A
label_25265: SET 0,B
label_25266: SET 1,B
label_25267: SET 2,B
label_25268: SET 3,B
label_25269: SET 4,B
label_25270: SET 5,B
label_25271: SET 6,B
label_25272: SET 7,B
label_25273: SET 0,C
label_25274: SET 1,C
label_25275: SET 2,C
label_25276: SET 3,C
label_25277: SET 4,C
label_25278: SET 5,C
label_25279: SET 6,C
label_25280: SET 7,C
label_25281: SET 0,D
label_25282: SET 1,D
label_25283: SET 2,D
label_25284: SET 3,D
label_25285: SET 4,D
label_25286: SET 5,D
label_25287: SET 6,D
label_25288: SET 7,D
label_25289: SET 0,E
label_25290: SET 1,E
label_25291: SET 2,E
label_25292: SET 3,E
label_25293: SET 4,E
label_25294: SET 5,E
label_25295: SET 6,E
label_25296: SET 7,E
label_25297: SET 0,H
label_25298: SET 1,H
label_25299: SET 2,H
label_25300: SET 3,H
label_25301: SET 4,H
label_25302: SET 5,H
label_25303: SET 6,H
label_25304: SET 7,H
label_25305: SET 0,L
label_25306: SET 1,L
label_25307: SET 2,L
label_25308: SET 3,L
label_25309: SET 4,L
label_25310: SET 5,L
label_25311: SET 6,L
label_25312: SET 7,L
label_25313: SET 0,(HL)
label_25314: SET 1,(HL)
label_25315: SET 2,(HL)
label_25316: SET 3,(HL)
label_25317: SET 4,(HL)
label_25318: SET 5,(HL)
label_25319: SET 6,(HL)
label_25320: SET 7,(HL)
label_25321: SET 0,(IX + 127)
label_25322: SET 1,(IX + 127)
label_25323: SET 2,(IX + 127)
label_25324: SET 3,(IX + 127)
label_25325: SET 4,(IX + 127)
label_25326: SET 5,(IX + 127)
label_25327: SET 6,(IX + 127)
label_25328: SET 7,(IX + 127)
label_25329: SET 0,(IY - 128)
label_25330: SET 1,(IY - 128)
label_25331: SET 2,(IY - 128)
label_25332: SET 3,(IY - 128)
label_25333: SET 4,(IY - 128)
label_25334: SET 5,(IY - 128)
label_25335: SET 6,(IY - 128)
label_25336: SET 7,(IY - 128)
label_25337: RES 0,A
label_25338: RES 1,A
label_25339: RES 2,A
label_25340: RES 3,A
label_25341: RES 4,A
label_25342: RES 5,A
label_25343: RES 6,A
label_25344: RES 7,A
label_25345: RES 0,B
label_25346: RES 1,B
label_25347: RES 2,B
label_25348: RES 3,B
label_25349: RES 4,B
label_25350: RES 5,B
label_25351: RES 6,B
label_25352: RES 7,B
label_25353: RES 0,C
label_25354: RES 1,C
label_25355: RES 2,C
label_25356: RES 3,C
label_25357: RES 4,C
label_25358: RES 5,C
label_25359: RES 6,C
label_25360: RES 7,C
label_25361: RES 0,D
label_25362: RES 1,D
label_25363: RES 2,D
label_25364: RES 3,D
label_25365: RES 4,D
label_25366: RES 5,D
label_25367: RES 6,D
label_25368: RES 7,D
label_25369: RES 0,E
label_25370: RES 1,E
label_25371: RES 2,E
label_25372: RES 3,E
label_25373: RES 4,E
label_25374: RES 5,E
label_25375: RES 6,E
label_25376: RES 7,E
label_25377: RES 0,H
label_25378: RES 1,H
label_25379: RES 2,H
label_25380: RES 3,H
label_25381: RES 4,H
label_25382: RES 5,H
label_25383: RES 6,H
label_25384: RES 7,H
label_25385: RES 0,L
label_25386: RES 1,L
label_25387: RES 2,L
label_25388: RES 3,L
label_25389: RES 4,L
label_25390: RES 5,L
label_25391: RES 6,L
label_25392: RES 7,L
label_25393: RES 0,(HL)
label_25394: RES 1,(HL)
label_25395: RES 2,(HL)
label_25396: RES 3,(HL)
label_25397: RES 4,(HL)
label_25398: RES 5,(HL)
label_25399: RES 6,(HL)
label_25400: RES 7,(HL)
label_25401: RES 0,(IX + 127)
label_25402: RES 1,(IX + 127)
label_25403: RES 2,(IX + 127)
label_25404: RES 3,(IX + 127)
label_25405: RES 4,(IX + 127)
label_25406: RES 5,(IX + 127)
label_25407: RES 6,(IX + 127)
label_25408: RES 7,(IX + 127)
label_25409: RES 0,(IY - 128)
label_25410: RES 1,(IY - 128)
label_25411: RES 2,(IY - 128)
label_25412: RES 3,(IY - 128)
label_25413: RES 4,(IY - 128)
label_25414: RES 5,(IY - 128)
label_25415: RES 6,(IY - 128)
label_25416: RES 7,(IY - 128)
label_25417: JP $5678
label_25418: JP NZ,$5678
label_25419: JP Z,$5678
label_25420: JP NC,$5678
label_25421: JP C,$5678
label_25422: JP PO,$5678
label_25423: JP PE,$5678
label_25424: JP P,$5678
label_25425: JP M,$5678
label_25426: JR $ + 2
label_25427: JR NZ,$ + 2
label_25428: JR Z,$ + 2
label_25429: JR NC,$ + 2
label_25430: JR C,$ + 2
label_25431: JP (HL)
label_25432: JP (IX)
label_25433: JP (IY)
label_25434: DJNZ $ + 2
label_25435: CALL $5678
label_25436: CALL NZ,$5678
label_25437: CALL Z,$5678
label_25438: CALL NC,$5678
label_25439: CALL C,$5678
label_25440: CALL PO,$5678
label_25441: CALL PE,$5678
label_25442: CALL P,$5678
label_25443: CALL M,$5678
label_25444: RET
label_25445: RET NZ
label_25446: RET Z
label_25447: RET NC
label_25448: RET C
label_25449: RET PO
label_25450: RET PE
label_25451: RET P
label_25452: RET M
label_25453: RETI
label_25454: RETN
label_25455: RST $00
label_25456: RST $08
label_25457: RST $10
label_25458: RST $18
label_25459: RST $20
label_25460: RST $28
label_25461: RST $30
label_25462: RST $38
label_25463: IN A,($12)
label_25464: IN A,(C)
label_25465: IN B,(C)
label_25466: IN C,(C)
label_25467: IN D,(C)
label_25468: IN E,(C)
label_25469: IN H,(C)
label_25470: IN L,(C)
label_25471: IN F,(C)
label_25472: INI
label_25473: INIR
label_25474: IND
label_25475: INDR
label_25476: OUT ($12),A
label_25477: OUT (C),A
label_25478: OUT (C),B
label_25479: OUT (C),C
label_25480: OUT (C),D
label_25481: OUT (C),E
label_25482: OUT (C),H
label_25483: OUT (C),L
label_25484: OUTI
label_25485: OTIR
label_25486: OUTD
label_25487: OTDR
label_25488: LD A,A
label_25489: LD A,B
label_25490: LD A,C
label_25491: LD A,D
label_25492: LD A,E
label_25493: LD A,H
label_25494: LD A,L
label_25495: LD B,A
label_25496: LD B,B
label_25497: LD B,C
label_25498: LD B,D
label_25499: LD B,E
label_25500: LD B,H
label_25501: LD B,L
label_25502: LD C,A
label_25503: LD C,B
label_25504: LD C,C
label_25505: LD C,D
label_25506: LD C,E
label_25507: LD C,H
label_25508: LD C,L
label_25509: LD D,A
label_25510: LD D,B
label_25511: LD D,C
label_25512: LD D,D
label_25513: LD D,E
label_25514: LD D,H
label_25515: LD D,L
label_25516: LD E,A
label_25517: LD E,B
label_25518: LD E,C
label_25519: LD E,D
label_25520: LD E,E
label_25521: LD E,H
label_25522: LD E,L
label_25523: LD H,A
label_25524: LD H,B
label_25525: LD H,C
label_25526: LD H,D
label_25527: LD H,E
label_25528: LD H,H
label_25529: LD H,L
label_25530: LD L,A
label_25531: LD L,B
label_25532: LD L,C
label_25533: LD L,D
label_25534: LD L,E
label_25535: LD L,H
label_25536: LD L,L
label_25537: LD A,$12
label_25538: LD B,$12
label_25539: LD C,$12
label_25540: LD D,$12
label_25541: LD E,$12
label_25542: LD H,$12
label_25543: LD L,$12
label_25544: LD A,(HL)
label_25545: LD B,(HL)
label_25546: LD C,(HL)
label_25547: LD D,(HL)
label_25548: LD E,(HL)
label_25549: LD H,(HL)
label_25550: LD L,(HL)
label_25551: LD A,(IX + 127)
label_25552: LD B,(IX + 127)
label_25553: LD C,(IX + 127)
label_25554: LD D,(IX + 127)
label_25555: LD E,(IX + 127)
label_25556: LD H,(IX + 127)
label_25557: LD L,(IX + 127)
label_25558: LD A,(IY - 128)
label_25559: LD B,(IY - 128)
label_25560: LD C,(IY - 128)
label_25561: LD D,(IY - 128)
label_25562: LD E,(IY - 128)
label_25563: LD H,(IY - 128)
label_25564: LD L,(IY - 128)
label_25565: LD (HL),A
label_25566: LD (HL),B
label_25567: LD (HL),C
label_25568: LD (HL),D
label_25569: LD (HL),E
label_25570: LD (HL),H
label_25571: LD (HL),L
label_25572: LD (IX + 127),A
label_25573: LD (IX + 127),B
label_25574: LD (IX + 127),C
label_25575: LD (IX + 127),D
label_25576: LD (IX + 127),E
label_25577: LD (IX + 127),H
label_25578: LD (IX + 127),L
label_25579: LD (IY - 128),A
label_25580: LD (IY - 128),B
label_25581: LD (IY - 128),C
label_25582: LD (IY - 128),D
label_25583: LD (IY - 128),E
label_25584: LD (IY - 128),H
label_25585: LD (IY - 128),L
label_25586: LD (HL),$12
label_25587: LD (IX + 127),$12
label_25588: LD (IY - 128),$12
label_25589: LD A,(BC)
label_25590: LD A,(DE)
label_25591: LD A,($5678)
label_25592: LD (BC),A
label_25593: LD (DE),A
label_25594: LD ($5678),A
label_25595: LD A,I
label_25596: LD A,R
label_25597: LD I,A
label_25598: LD R,A
label_25599: LD BC,$5678
label_25600: LD DE,$5678
label_25601: LD HL,$5678
label_25602: LD SP,$5678
label_25603: LD IX,$5678
label_25604: LD IY,$5678
label_25605: LD HL,($5678)
label_25606: LD BC,($5678)
label_25607: LD DE,($5678)
label_25608: LD HL,($5678)
label_25609: LD SP,($5678)
label_25610: LD IX,($5678)
label_25611: LD IY,($5678)
label_25612: LD ($5678),HL
label_25613: LD ($5678),BC
label_25614: LD ($5678),DE
label_25615: LD ($5678),HL
label_25616: LD ($5678),SP
label_25617: LD ($5678),IX
label_25618: LD ($5678),IY
label_25619: LD SP,HL
label_25620: LD SP,IX
label_25621: LD SP,IY
label_25622: PUSH BC
label_25623: PUSH DE
label_25624: PUSH HL
label_25625: PUSH AF
label_25626: PUSH IX
label_25627: PUSH IY
label_25628: POP BC
label_25629: POP DE
label_25630: POP HL
label_25631: POP AF
label_25632: POP IX
label_25633: POP IY
label_25634: EX DE,HL
label_25635: EX AF,AF'
label_25636: EXX
label_25637: EX (SP),HL
label_25638: EX (SP),IX
label_25639: EX (SP),IY
label_25640: LDI
label_25641: LDIR
label_25642: LDD
label_25643: LDDR
label_25644: CPI
label_25645: CPIR
label_25646: CPD
label_25647: CPDR
label_25648: ADD A,A
label_25649: ADD A,B
label_25650: ADD A,C
label_25651: ADD A,D
label_25652: ADD A,E
label_25653: ADD A,H
label_25654: ADD A,L
label_25655: ADD A,$12
label_25656: ADD A,(HL)
label_25657: ADD A,(IX + 127)
label_25658: ADD A,(IY - 128)
label_25659: ADC A,A
label_25660: ADC A,B
label_25661: ADC A,C
label_25662: ADC A,D
label_25663: ADC A,E
label_25664: ADC A,H
label_25665: ADC A,L
label_25666: ADC A,$12
label_25667: ADC A,(HL)
label_25668: ADC A,(IX + 127)
label_25669: ADC A,(IY - 128)
label_25670: SUB A
label_25671: SUB B
label_25672: SUB C
label_25673: SUB D
label_25674: SUB E
label_25675: SUB H
label_25676: SUB L
label_25677: SUB $12
label_25678: SUB (HL)
label_25679: SUB (IX + 127)
label_25680: SUB (IY - 128)
label_25681: SBC A,A
label_25682: SBC A,B
label_25683: SBC A,C
label_25684: SBC A,D
label_25685: SBC A,E
label_25686: SBC A,H
label_25687: SBC A,L
label_25688: SBC A,$12
label_25689: SBC A,(HL)
label_25690: SBC A,(IX + 127)
label_25691: SBC A,(IY - 128)
label_25692: AND A
label_25693: AND B
label_25694: AND C
label_25695: AND D
label_25696: AND E
label_25697: AND H
label_25698: AND L
label_25699: AND $12
label_25700: AND (HL)
label_25701: AND (IX + 127)
label_25702: AND (IY - 128)
label_25703: AND A
label_25704: AND B
label_25705: AND C
label_25706: AND D
label_25707: AND E
label_25708: AND H
label_25709: AND L
label_25710: AND $12
label_25711: AND (HL)
label_25712: AND (IX + 127)
label_25713: AND (IY - 128)
label_25714: OR A
label_25715: OR B
label_25716: OR C
label_25717: OR D
label_25718: OR E
label_25719: OR H
label_25720: OR L
label_25721: OR $12
label_25722: OR (HL)
label_25723: OR (IX + 127)
label_25724: OR (IY - 128)
label_25725: XOR A
label_25726: XOR B
label_25727: XOR C
label_25728: XOR D
label_25729: XOR E
label_25730: XOR H
label_25731: XOR L
label_25732: XOR $12
label_25733: XOR (HL)
label_25734: XOR (IX + 127)
label_25735: XOR (IY - 128)
label_25736: CP A
label_25737: CP B
label_25738: CP C
label_25739: CP D
label_25740: CP E
label_25741: CP H
label_25742: CP L
label_25743: CP $12
label_25744: CP (HL)
label_25745: CP (IX + 127)
label_25746: CP (IY - 128)
label_25747: INC A
label_25748: INC B
label_25749: INC C
label_25750: INC D
label_25751: INC E
label_25752: INC H
label_25753: INC L
label_25754: INC (HL)
label_25755: INC (IX + 127)
label_25756: INC (IY - 128)
label_25757: DEC A
label_25758: DEC B
label_25759: DEC C
label_25760: DEC D
label_25761: DEC E
label_25762: DEC H
label_25763: DEC L
label_25764: DEC (HL)
label_25765: DEC (IX + 127)
label_25766: DEC (IY - 128)
label_25767: DAA
label_25768: CPL
label_25769: NEG
label_25770: CCF
label_25771: SCF
label_25772: NOP
label_25773: HALT
label_25774: DI
label_25775: EI
label_25776: IM 0
label_25777: IM 1
label_25778: IM 2
label_25779: ADD HL,BC
label_25780: ADD HL,DE
label_25781: ADD HL,HL
label_25782: ADD HL,SP
label_25783: ADC HL,BC
label_25784: ADC HL,DE
label_25785: ADC HL,HL
label_25786: ADC HL,SP
label_25787: SBC HL,BC
label_25788: SBC HL,DE
label_25789: SBC HL,HL
label_25790: SBC HL,SP
label_25791: ADD IX,BC
label_25792: ADD IX,DE
label_25793: ADD IX,SP
label_25794: ADD IY,BC
label_25795: ADD IY,DE
label_25796: ADD IY,SP
label_25797: INC BC
label_25798: INC DE
label_25799: INC HL
label_25800: INC SP
label_25801: INC IX
label_25802: INC IY
label_25803: DEC BC
label_25804: DEC DE
label_25805: DEC HL
label_25806: DEC SP
label_25807: DEC IX
label_25808: DEC IY
label_25809: RLCA
label_25810: RLA
label_25811: RRCA
label_25812: RRA
label_25813: RLC A
label_25814: RLC B
label_25815: RLC C
label_25816: RLC D
label_25817: RLC E
label_25818: RLC H
label_25819: RLC L
label_25820: RLC (HL)
label_25821: RLC (IX + 127)
label_25822: RLC (IY - 128)
label_25823: RL A
label_25824: RL B
label_25825: RL C
label_25826: RL D
label_25827: RL E
label_25828: RL H
label_25829: RL L
label_25830: RL (HL)
label_25831: RL (IX + 127)
label_25832: RL (IY - 128)
label_25833: RRC A
label_25834: RRC B
label_25835: RRC C
label_25836: RRC D
label_25837: RRC E
label_25838: RRC H
label_25839: RRC L
label_25840: RRC (HL)
label_25841: RRC (IX + 127)
label_25842: RRC (IY - 128)
label_25843: RR A
label_25844: RR B
label_25845: RR C
label_25846: RR D
label_25847: RR E
label_25848: RR H
label_25849: RR L
label_25850: RR (HL)
label_25851: RR (IX + 127)
label_25852: RR (IY - 128)
label_25853: SLA A
label_25854: SLA B
label_25855: SLA C
label_25856: SLA D
label_25857: SLA E
label_25858: SLA H
label_25859: SLA L
label_25860: SLA (HL)
label_25861: SLA (IX + 127)
label_25862: SLA (IY - 128)
label_25863: SRA A
label_25864: SRA B
label_25865: SRA C
label_25866: SRA D
label_25867: SRA E
label_25868: SRA H
label_25869: SRA L
label_25870: SRA (HL)
label_25871: SRA (IX + 127)
label_25872: SRA (IY - 128)
label_25873: SRL A
label_25874: SRL B
label_25875: SRL C
label_25876: SRL D
label_25877: SRL E
label_25878: SRL H
label_25879: SRL L
label_25880: SRL (HL)
label_25881: SRL (IX + 127)
label_25882: SRL (IY - 128)
label_25883: RLD
label_25884: RRD
label_25885: BIT 0,A
label_25886: BIT 1,A
label_25887: BIT 2,A
label_25888: BIT 3,A
label_25889: BIT 4,A
label_25890: BIT 5,A
label_25891: BIT 6,A
label_25892: BIT 7,A
label_25893: BIT 0,B
label_25894: BIT 1,B
label_25895: BIT 2,B
label_25896: BIT 3,B
label_25897: BIT 4,B
label_25898: BIT 5,B
label_25899: BIT 6,B
label_25900: BIT 7,B
label_25901: BIT 0,C
label_25902: BIT 1,C
label_25903: BIT 2,C
label_25904: BIT 3,C
label_25905: BIT 4,C
label_25906: BIT 5,C
label_25907: BIT 6,C
label_25908: BIT 7,C
label_25909: BIT 0,D
label_25910: BIT 1,D
label_25911: BIT 2,D
label_25912: BIT 3,D
label_25913: BIT 4,D
label_25914: BIT 5,D
label_25915: BIT 6,D
label_25916: BIT 7,D
label_25917: BIT 0,E
label_25918: BIT 1,E
label_25919: BIT 2,E
label_25920: BIT 3,E
label_25921: BIT 4,E
label_25922: BIT 5,E
label_25923: BIT 6,E
label_25924: BIT 7,E
label_25925: BIT 0,H
label_25926: BIT 1,H
label_25927: BIT 2,H
label_25928: BIT 3,H
label_25929: BIT 4,H
label_25930: BIT 5,H
label_25931: BIT 6,H
label_25932: BIT 7,H
label_25933: BIT 0,L
label_25934: BIT 1,L
label_25935: BIT 2,L
label_25936: BIT 3,L
label_25937: BIT 4,L
label_25938: BIT 5,L
label_25939: BIT 6,L
label_25940: BIT 7,L
label_25941: BIT 0,(HL)
label_25942: BIT 1,(HL)
label_25943: BIT 2,(HL)
label_25944: BIT 3,(HL)
label_25945: BIT 4,(HL)
label_25946: BIT 5,(HL)
label_25947: BIT 6,(HL)
label_25948: BIT 7,(HL)
label_25949: BIT 0,(IX + 127)
label_25950: BIT 1,(IX + 127)
label_25951: BIT 2,(IX + 127)
label_25952: BIT 3,(IX + 127)
label_25953: BIT 4,(IX + 127)
label_25954: BIT 5,(IX + 127)
label_25955: BIT 6,(IX + 127)
label_25956: BIT 7,(IX + 127)
label_25957: BIT 0,(IY - 128)
label_25958: BIT 1,(IY - 128)
label_25959: BIT 2,(IY - 128)
label_25960: BIT 3,(IY - 128)
label_25961: BIT 4,(IY - 128)
label_25962: BIT 5,(IY - 128)
label_25963: BIT 6,(IY - 128)
label_25964: BIT 7,(IY - 128)
label_25965: SET 0,A
label_25966: SET 1,A
label_25967: SET 2,A
label_25968: SET 3,A
label_25969: SET 4,A
label_25970: SET 5,A
label_25971: SET 6,A
label_25972: SET 7,A
label_25973: SET 0,B
label_25974: SET 1,B
label_25975: SET 2,B
label_25976: SET 3,B
label_25977: SET 4,B
label_25978: SET 5,B
label_25979: SET 6,B
label_25980: SET 7,B
label_25981: SET 0,C
label_25982: SET 1,C
label_25983: SET 2,C
label_25984: SET 3,C
label_25985: SET 4,C
label_25986: SET 5,C
label_25987: SET 6,C
label_25988: SET 7,C
label_25989: SET 0,D
label_25990: SET 1,D
label_25991: SET 2,D
label_25992: SET 3,D
label_25993: SET 4,D
label_25994: SET 5,D
label_25995: SET 6,D
label_25996: SET 7,D
label_25997: SET 0,E
label_25998: SET 1,E
label_25999: SET 2,E
label_26000: SET 3,E
label_26001: SET 4,E
label_26002: SET 5,E
label_26003: SET 6,E
label_26004: SET 7,E
label_26005: SET 0,H
label_26006: SET 1,H
label_26007: SET 2,H
label_26008: SET 3,H
label_26009: SET 4,H
label_26010: SET 5,H
label_26011: SET 6,H
label_26012: SET 7,H
label_26013: SET 0,L
label_26014: SET 1,L
label_26015: SET 2,L
label_26016: SET 3,L
label_26017: SET 4,L
label_26018: SET 5,L
label_26019: SET 6,L
label_26020: SET 7,L
label_26021: SET 0,(HL)
label_26022: SET 1,(HL)
label_26023: SET 2,(HL)
label_26024: SET 3,(HL)
label_26025: SET 4,(HL)
label_26026: SET 5,(HL)
label_26027: SET 6,(HL)
label_26028: SET 7,(HL)
label_26029: SET 0,(IX + 127)
label_26030: SET 1,(IX + 127)
label_26031: SET 2,(IX + 127)
label_26032: SET 3,(IX + 127)
label_26033: SET 4,(IX + 127)
label_26034: SET 5,(IX + 127)
label_26035: SET 6,(IX + 127)
label_26036: SET 7,(IX + 127)
label_26037: SET 0,(IY - 128)
label_26038: SET 1,(IY - 128)
label_26039: SET 2,(IY - 128)
label_26040: SET 3,(IY - 128)
label_26041: SET 4,(IY - 128)
label_26042: SET 5,(IY - 128)
label_26043: SET 6,(IY - 128)
label_26044: SET 7,(IY - 128)
label_26045: RES 0,A
label_26046: RES 1,A
label_26047: RES 2,A
label_26048: RES 3,A
label_26049: RES 4,A
label_26050: RES 5,A
label_26051: RES 6,A
label_26052: RES 7,A
label_26053: RES 0,B
label_26054: RES 1,B
label_26055: RES 2,B
label_26056: RES 3,B
label_26057: RES 4,B
label_26058: RES 5,B
label_26059: RES 6,B
label_26060: RES 7,B
label_26061: RES 0,C
label_26062: RES 1,C
label_26063: RES 2,C
label_26064: RES 3,C
label_26065: RES 4,C
label_26066: RES 5,C
label_26067: RES 6,C
label_26068: RES 7,C
label_26069: RES 0,D
label_26070: RES 1,D
label_26071: RES 2,D
label_26072: RES 3,D
label_26073: RES 4,D
label_26074: RES 5,D
label_26075: RES 6,D
label_26076: RES 7,D
label_26077: RES 0,E
label_26078: RES 1,E
label_26079: RES 2,E
label_26080: RES 3,E
label_26081: RES 4,E
label_26082: RES 5,E
label_26083: RES 6,E
label_26084: RES 7,E
label_26085: RES 0,H
label_26086: RES 1,H
label_26087: RES 2,H
label_26088: RES 3,H
label_26089: RES 4,H
label_26090: RES 5,H
label_26091: RES 6,H
label_26092: RES 7,H
label_26093: RES 0,L
label_26094: RES 1,L
label_26095: RES 2,L
label_26096: RES 3,L
label_26097: RES 4,L
label_26098: RES 5,L
label_26099: RES 6,L
label_26100: RES 7,L
label_26101: RES 0,(HL)
label_26102: RES 1,(HL)
label_26103: RES 2,(HL)
label_26104: RES 3,(HL)
label_26105: RES 4,(HL)
label_26106: RES 5,(HL)
label_26107: RES 6,(HL)
label_26108: RES 7,(HL)
label_26109: RES 0,(IX + 127)
label_26110: RES 1,(IX + 127)
label_26111: RES 2,(IX + 127)
label_26112: RES 3,(IX + 127)
label_26113: RES 4,(IX + 127)
label_26114: RES 5,(IX + 127)
label_26115: RES 6,(IX + 127)
label_26116: RES 7,(IX + 127)
label_26117: RES 0,(IY - 128)
label_26118: RES 1,(IY - 128)
label_26119: RES 2,(IY - 128)
label_26120: RES 3,(IY - 128)
label_26121: RES 4,(IY - 128)
label_26122: RES 5,(IY - 128)
label_26123: RES 6,(IY - 128)
label_26124: RES 7,(IY - 128)
label_26125: JP $5678
label_26126: JP NZ,$5678
label_26127: JP Z,$5678
label_26128: JP NC,$5678
label_26129: JP C,$5678
label_26130: JP PO,$5678
label_26131: JP PE,$5678
label_26132: JP P,$5678
label_26133: JP M,$5678
label_26134: JR $ + 2
label_26135: JR NZ,$ + 2
label_26136: JR Z,$ + 2
label_26137: JR NC,$ + 2
label_26138: JR C,$ + 2
label_26139: JP (HL)
label_26140: JP (IX)
label_26141: JP (IY)
label_26142: DJNZ $ + 2
label_26143: CALL $5678
label_26144: CALL NZ,$5678
label_26145: CALL Z,$5678
label_26146: CALL NC,$5678
label_26147: CALL C,$5678
label_26148: CALL PO,$5678
label_26149: CALL PE,$5678
label_26150: CALL P,$5678
label_26151: CALL M,$5678
label_26152: RET
label_26153: RET NZ
label_26154: RET Z
label_26155: RET NC
label_26156: RET C
label_26157: RET PO
label_26158: RET PE
label_26159: RET P
label_26160: RET M
label_26161: RETI
label_26162: RETN
label_26163: RST $00
label_26164: RST $08
label_26165: RST $10
label_26166: RST $18
label_26167: RST $20
label_26168: RST $28
label_26169: RST $30
label_26170: RST $38
label_26171: IN A,($12)
label_26172: IN A,(C)
label_26173: IN B,(C)
label_26174: IN C,(C)
label_26175: IN D,(C)
label_26176: IN E,(C)
label_26177: IN H,(C)
label_26178: IN L,(C)
label_26179: IN F,(C)
label_26180: INI
label_26181: INIR
label_26182: IND
label_26183: INDR
label_26184: OUT ($12),A
label_26185: OUT (C),A
label_26186: OUT (C),B
label_26187: OUT (C),C
label_26188: OUT (C),D
label_26189: OUT (C),E
label_26190: OUT (C),H
label_26191: OUT (C),L
label_26192: OUTI
label_26193: OTIR
label_26194: OUTD
label_26195: OTDR
label_26196: LD A,A
label_26197: LD A,B
label_26198: LD A,C
label_26199: LD A,D
label_26200: LD A,E
label_26201: LD A,H
label_26202: LD A,L
label_26203: LD B,A
label_26204: LD B,B
label_26205: LD B,C
label_26206: LD B,D
label_26207: LD B,E
label_26208: LD B,H
label_26209: LD B,L
label_26210: LD C,A
label_26211: LD C,B
label_26212: LD C,C
label_26213: LD C,D
label_26214: LD C,E
label_26215: LD C,H
label_26216: LD C,L
label_26217: LD D,A
label_26218: LD D,B
label_26219: LD D,C
label_26220: LD D,D
label_26221: LD D,E
label_26222: LD D,H
label_26223: LD D,L
label_26224: LD E,A
label_26225: LD E,B
label_26226: LD E,C
label_26227: LD E,D
label_26228: LD E,E
label_26229: LD E,H
label_26230: LD E,L
label_26231: LD H,A
label_26232: LD H,B
label_26233: LD H,C
label_26234: LD H,D
label_26235: LD H,E
label_26236: LD H,H
label_26237: LD H,L
label_26238: LD L,A
label_26239: LD L,B
label_26240: LD L,C
label_26241: LD L,D
label_26242: LD L,E
label_26243: LD L,H
label_26244: LD L,L
label_26245: LD A,$12
label_26246: LD B,$12
label_26247: LD C,$12
label_26248: LD D,$12
label_26249: LD E,$12
label_26250: LD H,$12
label_26251: LD L,$12
label_26252: LD A,(HL)
label_26253: LD B,(HL)
label_26254: LD C,(HL)
label_26255: LD D,(HL)
label_26256: LD E,(HL)
label_26257: LD H,(HL)
label_26258: LD L,(HL)
label_26259: LD A,(IX + 127)
label_26260: LD B,(IX + 127)
label_26261: LD C,(IX + 127)
label_26262: LD D,(IX + 127)
label_26263: LD E,(IX + 127)
label_26264: LD H,(IX + 127)
label_26265: LD L,(IX + 127)
label_26266: LD A,(IY - 128)
label_26267: LD B,(IY - 128)
label_26268: LD C,(IY - 128)
label_26269: LD D,(IY - 128)
label_26270: LD E,(IY - 128)
label_26271: LD H,(IY - 128)
label_26272: LD L,(IY - 128)
label_26273: LD (HL),A
label_26274: LD (HL),B
label_26275: LD (HL),C
label_26276: LD (HL),D
label_26277: LD (HL),E
label_26278: LD (HL),H
label_26279: LD (HL),L
label_26280: LD (IX + 127),A
label_26281: LD (IX + 127),B
label_26282: LD (IX + 127),C
label_26283: LD (IX + 127),D
label_26284: LD (IX + 127),E
label_26285: LD (IX + 127),H
label_26286: LD (IX + 127),L
label_26287: LD (IY - 128),A
label_26288: LD (IY - 128),B
label_26289: LD (IY - 128),C
label_26290: LD (IY - 128),D
label_26291: LD (IY - 128),E
label_26292: LD (IY - 128),H
label_26293: LD (IY - 128),L
label_26294: LD (HL),$12
label_26295: LD (IX + 127),$12
label_26296: LD (IY - 128),$12
label_26297: LD A,(BC)
label_26298: LD A,(DE)
label_26299: LD A,($5678)
label_26300: LD (BC),A
label_26301: LD (DE),A
label_26302: LD ($5678),A
label_26303: LD A,I
label_26304: LD A,R
label_26305: LD I,A
label_26306: LD R,A
label_26307: LD BC,$5678
label_26308: LD DE,$5678
label_26309: LD HL,$5678
label_26310: LD SP,$5678
label_26311: LD IX,$5678
label_26312: LD IY,$5678
label_26313: LD HL,($5678)
label_26314: LD BC,($5678)
label_26315: LD DE,($5678)
label_26316: LD HL,($5678)
label_26317: LD SP,($5678)
label_26318: LD IX,($5678)
label_26319: LD IY,($5678)
label_26320: LD ($5678),HL
label_26321: LD ($5678),BC
label_26322: LD ($5678),DE
label_26323: LD ($5678),HL
label_26324: LD ($5678),SP
label_26325: LD ($5678),IX
label_26326: LD ($5678),IY
label_26327: LD SP,HL
label_26328: LD SP,IX
label_26329: LD SP,IY
label_26330: PUSH BC
label_26331: PUSH DE
label_26332: PUSH HL
label_26333: PUSH AF
label_26334: PUSH IX
label_26335: PUSH IY
label_26336: POP BC
label_26337: POP DE
label_26338: POP HL
label_26339: POP AF
label_26340: POP IX
label_26341: POP IY
label_26342: EX DE,HL
label_26343: EX AF,AF'
label_26344: EXX
label_26345: EX (SP),HL
label_26346: EX (SP),IX
label_26347: EX (SP),IY
label_26348: LDI
label_26349: LDIR
label_26350: LDD
label_26351: LDDR
label_26352: CPI
label_26353: CPIR
label_26354: CPD
label_26355: CPDR
label_26356: ADD A,A
label_26357: ADD A,B
label_26358: ADD A,C
label_26359: ADD A,D
label_26360: ADD A,E
label_26361: ADD A,H
label_26362: ADD A,L
label_26363: ADD A,$12
label_26364: ADD A,(HL)
label_26365: ADD A,(IX + 127)
label_26366: ADD A,(IY - 128)
label_26367: ADC A,A
label_26368: ADC A,B
label_26369: ADC A,C
label_26370: ADC A,D
label_26371: ADC A,E
label_26372: ADC A,H
label_26373: ADC A,L
label_26374: ADC A,$12
label_26375: ADC A,(HL)
label_26376: ADC A,(IX + 127)
label_26377: ADC A,(IY - 128)
label_26378: SUB A
label_26379: SUB B
label_26380: SUB C
label_26381: SUB D
label_26382: SUB E
label_26383: SUB H
label_26384: SUB L
label_26385: SUB $12
label_26386: SUB (HL)
label_26387: SUB (IX + 127)
label_26388: SUB (IY - 128)
label_26389: SBC A,A
label_26390: SBC A,B
label_26391: SBC A,C
label_26392: SBC A,D
label_26393: SBC A,E
label_26394: SBC A,H
label_26395: SBC A,L
label_26396: SBC A,$12
label_26397: SBC A,(HL)
label_26398: SBC A,(IX + 127)
label_26399: SBC A,(IY - 128)
label_26400: AND A
label_26401: AND B
label_26402: AND C
label_26403: AND D
label_26404: AND E
label_26405: AND H
label_26406: AND L
label_26407: AND $12
label_26408: AND (HL)
label_26409: AND (IX + 127)
label_26410: AND (IY - 128)
label_26411: AND A
label_26412: AND B
label_26413: AND C
label_26414: AND D
label_26415: AND E
label_26416: AND H
label_26417: AND L
label_26418: AND $12
label_26419: AND (HL)
label_26420: AND (IX + 127)
label_26421: AND (IY - 128)
label_26422: OR A
label_26423: OR B
label_26424: OR C
label_26425: OR D
label_26426: OR E
label_26427: OR H
label_26428: OR L
label_26429: OR $12
label_26430: OR (HL)
label_26431: OR (IX + 127)
label_26432: OR (IY - 128)
label_26433: XOR A
label_26434: XOR B
label_26435: XOR C
label_26436: XOR D
label_26437: XOR E
label_26438: XOR H
label_26439: XOR L
label_26440: XOR $12
label_26441: XOR (HL)
label_26442: XOR (IX + 127)
label_26443: XOR (IY - 128)
label_26444: CP A
label_26445: CP B
label_26446: CP C
label_26447: CP D
label_26448: CP E
label_26449: CP H
label_26450: CP L
label_26451: CP $12
label_26452: CP (HL)
label_26453: CP (IX + 127)
label_26454: CP (IY - 128)
label_26455: INC A
label_26456: INC B
label_26457: INC C
label_26458: INC D
label_26459: INC E
label_26460: INC H
label_26461: INC L
label_26462: INC (HL)
label_26463: INC (IX + 127)
label_26464: INC (IY - 128)
label_26465: DEC A
label_26466: DEC B
label_26467: DEC C
label_26468: DEC D
label_26469: DEC E
label_26470: DEC H
label_26471: DEC L
label_26472: DEC (HL)
label_26473: DEC (IX + 127)
label_26474: DEC (IY - 128)
label_26475: DAA
label_26476: CPL
label_26477: NEG
label_26478: CCF
label_26479: SCF
label_26480: NOP
label_26481: HALT
label_26482: DI
label_26483: EI
label_26484: IM 0
label_26485: IM 1
label_26486: IM 2
label_26487: ADD HL,BC
label_26488: ADD HL,DE
label_26489: ADD HL,HL
label_26490: ADD HL,SP
label_26491: ADC HL,BC
label_26492: ADC HL,DE
label_26493: ADC HL,HL
label_26494: ADC HL,SP
label_26495: SBC HL,BC
label_26496: SBC HL,DE
label_26497: SBC HL,HL
label_26498: SBC HL,SP
label_26499: ADD IX,BC
label_26500: ADD IX,DE
label_26501: ADD IX,SP
label_26502: ADD IY,BC
label_26503: ADD IY,DE
label_26504: ADD IY,SP
label_26505: INC BC
label_26506: INC DE
label_26507: INC HL
label_26508: INC SP
label_26509: INC IX
label_26510: INC IY
label_26511: DEC BC
label_26512: DEC DE
label_26513: DEC HL
label_26514: DEC SP
label_26515: DEC IX
label_26516: DEC IY
label_26517: RLCA
label_26518: RLA
label_26519: RRCA
label_26520: RRA
label_26521: RLC A
label_26522: RLC B
label_26523: RLC C
label_26524: RLC D
label_26525: RLC E
label_26526: RLC H
label_26527: RLC L
label_26528: RLC (HL)
label_26529: RLC (IX + 127)
label_26530: RLC (IY - 128)
label_26531: RL A
label_26532: RL B
label_26533: RL C
label_26534: RL D
label_26535: RL E
label_26536: RL H
label_26537: RL L
label_26538: RL (HL)
label_26539: RL (IX + 127)
label_26540: RL (IY - 128)
label_26541: RRC A
label_26542: RRC B
label_26543: RRC C
label_26544: RRC D
label_26545: RRC E
label_26546: RRC H
label_26547: RRC L
label_26548: RRC (HL)
label_26549: RRC (IX + 127)
label_26550: RRC (IY - 128)
label_26551: RR A
label_26552: RR B
label_26553: RR C
label_26554: RR D
label_26555: RR E
label_26556: RR H
label_26557: RR L
label_26558: RR (HL)
label_26559: RR (IX + 127)
label_26560: RR (IY - 128)
label_26561: SLA A
label_26562: SLA B
label_26563: SLA C
label_26564: SLA D
label_26565: SLA E
label_26566: SLA H
label_26567: SLA L
label_26568: SLA (HL)
label_26569: SLA (IX + 127)
label_26570: SLA (IY - 128)
label_26571: SRA A
label_26572: SRA B
label_26573: SRA C
label_26574: SRA D
label_26575: SRA E
label_26576: SRA H
label_26577: SRA L
label_26578: SRA (HL)
label_26579: SRA (IX + 127)
label_26580: SRA (IY - 128)
label_26581: SRL A
label_26582: SRL B
label_26583: SRL C
label_26584: SRL D
label_26585: SRL E
label_26586: SRL H
label_26587: SRL L
label_26588: SRL (HL)
label_26589: SRL (IX + 127)
label_26590: SRL (IY - 128)
label_26591: RLD
label_26592: RRD
label_26593: BIT 0,A
label_26594: BIT 1,A
label_26595: BIT 2,A
label_26596: BIT 3,A
label_26597: BIT 4,A
label_26598: BIT 5,A
label_26599: BIT 6,A
label_26600: BIT 7,A
label_26601: BIT 0,B
label_26602: BIT 1,B
label_26603: BIT 2,B
label_26604: BIT 3,B
label_26605: BIT 4,B
label_26606: BIT 5,B
label_26607: BIT 6,B
label_26608: BIT 7,B
label_26609: BIT 0,C
label_26610: BIT 1,C
label_26611: BIT 2,C
label_26612: BIT 3,C
label_26613: BIT 4,C
label_26614: BIT 5,C
label_26615: BIT 6,C
label_26616: BIT 7,C
label_26617: BIT 0,D
label_26618: BIT 1,D
label_26619: BIT 2,D
label_26620: BIT 3,D
label_26621: BIT 4,D
label_26622: BIT 5,D
label_26623: BIT 6,D
label_26624: BIT 7,D
label_26625: BIT 0,E
label_26626: BIT 1,E
label_26627: BIT 2,E
label_26628: BIT 3,E
label_26629: BIT 4,E
label_26630: BIT 5,E
label_26631: BIT 6,E
label_26632: BIT 7,E
label_26633: BIT 0,H
label_26634: BIT 1,H
label_26635: BIT 2,H
label_26636: BIT 3,H
label_26637: BIT 4,H
label_26638: BIT 5,H
label_26639: BIT 6,H
label_26640: BIT 7,H
label_26641: BIT 0,L
label_26642: BIT 1,L
label_26643: BIT 2,L
label_26644: BIT 3,L
label_26645: BIT 4,L
label_26646: BIT 5,L
label_26647: BIT 6,L
label_26648: BIT 7,L
label_26649: BIT 0,(HL)
label_26650: BIT 1,(HL)
label_26651: BIT 2,(HL)
label_26652: BIT 3,(HL)
label_26653: BIT 4,(HL)
label_26654: BIT 5,(HL)
label_26655: BIT 6,(HL)
label_26656: BIT 7,(HL)
label_26657: BIT 0,(IX + 127)
label_26658: BIT 1,(IX + 127)
label_26659: BIT 2,(IX + 127)
label_26660: BIT 3,(IX + 127)
label_26661: BIT 4,(IX + 127)
label_26662: BIT 5,(IX + 127)
label_26663: BIT 6,(IX + 127)
label_26664: BIT 7,(IX + 127)
label_26665: BIT 0,(IY - 128)
label_26666: BIT 1,(IY - 128)
label_26667: BIT 2,(IY - 128)
label_26668: BIT 3,(IY - 128)
label_26669: BIT 4,(IY - 128)
label_26670: BIT 5,(IY - 128)
label_26671: BIT 6,(IY - 128)
label_26672: BIT 7,(IY - 128)
label_26673: SET 0,A
label_26674: SET 1,A
label_26675: SET 2,A
label_26676: SET 3,A
label_26677: SET 4,A
label_26678: SET 5,A
label_26679: SET 6,A
label_26680: SET 7,A
label_26681: SET 0,B
label_26682: SET 1,B
label_26683: SET 2,B
label_26684: SET 3,B
label_26685: SET 4,B
label_26686: SET 5,B
label_26687: SET 6,B
label_26688: SET 7,B
label_26689: SET 0,C
label_26690: SET 1,C
label_26691: SET 2,C
label_26692: SET 3,C
label_26693: SET 4,C
label_26694: SET 5,C
label_26695: SET 6,C
label_26696: SET 7,C
label_26697: SET 0,D
label_26698: SET 1,D
label_26699: SET 2,D
label_26700: SET 3,D
label_26701: SET 4,D
label_26702: SET 5,D
label_26703: SET 6,D
label_26704: SET 7,D
label_26705: SET 0,E
label_26706: SET 1,E
label_26707: SET 2,E
label_26708: SET 3,E
label_26709: SET 4,E
label_26710: SET 5,E
label_26711: SET 6,E
label_26712: SET 7,E
label_26713: SET 0,H
label_26714: SET 1,H
label_26715: SET 2,H
label_26716: SET 3,H
label_26717: SET 4,H
label_26718: SET 5,H
label_26719: SET 6,H
label_26720: SET 7,H
label_26721: SET 0,L
label_26722: SET 1,L
label_26723: SET 2,L
label_26724: SET 3,L
label_26725: SET 4,L
label_26726: SET 5,L
label_26727: SET 6,L
label_26728: SET 7,L
label_26729: SET 0,(HL)
label_26730: SET 1,(HL)
label_26731: SET 2,(HL)
label_26732: SET 3,(HL)
label_26733: SET 4,(HL)
label_26734: SET 5,(HL)
label_26735: SET 6,(HL)
label_26736: SET 7,(HL)
label_26737: SET 0,(IX + 127)
label_26738: SET 1,(IX + 127)
label_26739: SET 2,(IX + 127)
label_26740: SET 3,(IX + 127)
label_26741: SET 4,(IX + 127)
label_26742: SET 5,(IX + 127)
label_26743: SET 6,(IX + 127)
label_26744: SET 7,(IX + 127)
label_26745: SET 0,(IY - 128)
label_26746: SET 1,(IY - 128)
label_26747: SET 2,(IY - 128)
label_26748: SET 3,(IY - 128)
label_26749: SET 4,(IY - 128)
label_26750: SET 5,(IY - 128)
label_26751: SET 6,(IY - 128)
label_26752: SET 7,(IY - 128)
label_26753: RES 0,A
label_26754: RES 1,A
label_26755: RES 2,A
label_26756: RES 3,A
label_26757: RES 4,A
label_26758: RES 5,A
label_26759: RES 6,A
label_26760: RES 7,A
label_26761: RES 0,B
label_26762: RES 1,B
label_26763: RES 2,B
label_26764: RES 3,B
label_26765: RES 4,B
label_26766: RES 5,B
label_26767: RES 6,B
label_26768: RES 7,B
label_26769: RES 0,C
label_26770: RES 1,C
label_26771: RES 2,C
label_26772: RES 3,C
label_26773: RES 4,C
label_26774: RES 5,C
label_26775: RES 6,C
label_26776: RES 7,C
label_26777: RES 0,D
label_26778: RES 1,D
label_26779: RES 2,D
label_26780: RES 3,D
label_26781: RES 4,D
label_26782: RES 5,D
label_26783: RES 6,D
label_26784: RES 7,D
label_26785: RES 0,E
label_26786: RES 1,E
label_26787: RES 2,E
label_26788: RES 3,E
label_26789: RES 4,E
label_26790: RES 5,E
label_26791: RES 6,E
label_26792: RES 7,E
label_26793: RES 0,H
label_26794: RES 1,H
label_26795: RES 2,H
label_26796: RES 3,H
label_26797: RES 4,H
label_26798: RES 5,H
label_26799: RES 6,H
label_26800: RES 7,H
label_26801: RES 0,L
label_26802: RES 1,L
label_26803: RES 2,L
label_26804: RES 3,L
label_26805: RES 4,L
label_26806: RES 5,L
label_26807: RES 6,L
label_26808: RES 7,L
label_26809: RES 0,(HL)
label_26810: RES 1,(HL)
label_26811: RES 2,(HL)
label_26812: RES 3,(HL)
label_26813: RES 4,(HL)
label_26814: RES 5,(HL)
label_26815: RES 6,(HL)
label_26816: RES 7,(HL)
label_26817: RES 0,(IX + 127)
label_26818: RES 1,(IX + 127)
label_26819: RES 2,(IX + 127)
label_26820: RES 3,(IX + 127)
label_26821: RES 4,(IX + 127)
label_26822: RES 5,(IX + 127)
label_26823: RES 6,(IX + 127)
label_26824: RES 7,(IX + 127)
label_26825: RES 0,(IY - 128)
label_26826: RES 1,(IY - 128)
label_26827: RES 2,(IY - 128)
label_26828: RES 3,(IY - 128)
label_26829: RES 4,(IY - 128)
label_26830: RES 5,(IY - 128)
label_26831: RES 6,(IY - 128)
label_26832: RES 7,(IY - 128)
label_26833: JP $5678
label_26834: JP NZ,$5678
label_26835: JP Z,$5678
label_26836: JP NC,$5678
label_26837: JP C,$5678
label_26838: JP PO,$5678
label_26839: JP PE,$5678
label_26840: JP P,$5678
label_26841: JP M,$5678
label_26842: JR $ + 2
label_26843: JR NZ,$ + 2
label_26844: JR Z,$ + 2
label_26845: JR NC,$ + 2
label_26846: JR C,$ + 2
label_26847: JP (HL)
label_26848: JP (IX)
label_26849: JP (IY)
label_26850: DJNZ $ + 2
label_26851: CALL $5678
label_26852: CALL NZ,$5678
label_26853: CALL Z,$5678
label_26854: CALL NC,$5678
label_26855: CALL C,$5678
label_26856: CALL PO,$5678
label_26857: CALL PE,$5678
label_26858: CALL P,$5678
label_26859: CALL M,$5678
label_26860: RET
label_26861: RET NZ
label_26862: RET Z
label_26863: RET NC
label_26864: RET C
label_26865: RET PO
label_26866: RET PE
label_26867: RET P
label_26868: RET M
label_26869: RETI
label_26870: RETN
label_26871: RST $00
label_26872: RST $08
label_26873: RST $10
label_26874: RST $18
label_26875: RST $20
label_26876: RST $28
label_26877: RST $30
label_26878: RST $38
label_26879: IN A,($12)
label_26880: IN A,(C)
label_26881: IN B,(C)
label_26882: IN C,(C)
label_26883: IN D,(C)
label_26884: IN E,(C)
label_26885: IN H,(C)
label_26886: IN L,(C)
label_26887: IN F,(C)
label_26888: INI
label_26889: INIR
label_26890: IND
label_26891: INDR
label_26892: OUT ($12),A
label_26893: OUT (C),A
label_26894: OUT (C),B
label_26895: OUT (C),C
label_26896: OUT (C),D
label_26897: OUT (C),E
label_26898: OUT (C),H
label_26899: OUT (C),L
label_26900: OUTI
label_26901: OTIR
label_26902: OUTD
label_26903: OTDR
label_26904: LD A,A
label_26905: LD A,B
label_26906: LD A,C
label_26907: LD A,D
label_26908: LD A,E
label_26909: LD A,H
label_26910: LD A,L
label_26911: LD B,A
label_26912: LD B,B
label_26913: LD B,C
label_26914: LD B,D
label_26915: LD B,E
label_26916: LD B,H
label_26917: LD B,L
label_26918: LD C,A
label_26919: LD C,B
label_26920: LD C,C
label_26921: LD C,D
label_26922: LD C,E
label_26923: LD C,H
label_26924: LD C,L
label_26925: LD D,A
label_26926: LD D,B
label_26927: LD D,C
label_26928: LD D,D
label_26929: LD D,E
label_26930: LD D,H
label_26931: LD D,L
label_26932: LD E,A
label_26933: LD E,B
label_26934: LD E,C
label_26935: LD E,D
label_26936: LD E,E
label_26937: LD E,H
label_26938: LD E,L
label_26939: LD H,A
label_26940: LD H,B
label_26941: LD H,C
label_26942: LD H,D
label_26943: LD H,E
label_26944: LD H,H
label_26945: LD H,L
label_26946: LD L,A
label_26947: LD L,B
label_26948: LD L,C
label_26949: LD L,D
label_26950: LD L,E
label_26951: LD L,H
label_26952: LD L,L
label_26953: LD A,$12
label_26954: LD B,$12
label_26955: LD C,$12
label_26956: LD D,$12
label_26957: LD E,$12
label_26958: LD H,$12
label_26959: LD L,$12
label_26960: LD A,(HL)
label_26961: LD B,(HL)
label_26962: LD C,(HL)
label_26963: LD D,(HL)
label_26964: LD E,(HL)
label_26965: LD H,(HL)
label_26966: LD L,(HL)
label_26967: LD A,(IX + 127)
label_26968: LD B,(IX + 127)
label_26969: LD C,(IX + 127)
label_26970: LD D,(IX + 127)
label_26971: LD E,(IX + 127)
label_26972: LD H,(IX + 127)
label_26973: LD L,(IX + 127)
label_26974: LD A,(IY - 128)
label_26975: LD B,(IY - 128)
label_26976: LD C,(IY - 128)
label_26977: LD D,(IY - 128)
label_26978: LD E,(IY - 128)
label_26979: LD H,(IY - 128)
label_26980: LD L,(IY - 128)
label_26981: LD (HL),A
label_26982: LD (HL),B
label_26983: LD (HL),C
label_26984: LD (HL),D
label_26985: LD (HL),E
label_26986: LD (HL),H
label_26987: LD (HL),L
label_26988: LD (IX + 127),A
label_26989: LD (IX + 127),B
label_26990: LD (IX + 127),C
label_26991: LD (IX + 127),D
label_26992: LD (IX + 127),E
label_26993: LD (IX + 127),H
label_26994: LD (IX + 127),L
label_26995: LD (IY - 128),A
label_26996: LD (IY - 128),B
label_26997: LD (IY - 128),C
label_26998: LD (IY - 128),D
label_26999: LD (IY - 128),E
label_27000: LD (IY - 128),H
label_27001: LD (IY - 128),L
label_27002: LD (HL),$12
label_27003: LD (IX + 127),$12
label_27004: LD (IY - 128),$12
label_27005: LD A,(BC)
label_27006: LD A,(DE)
label_27007: LD A,($5678)
label_27008: LD (BC),A
label_27009: LD (DE),A
label_27010: LD ($5678),A
label_27011: LD A,I
label_27012: LD A,R
label_27013: LD I,A
label_27014: LD R,A
label_27015: LD BC,$5678
label_27016: LD DE,$5678
label_27017: LD HL,$5678
label_27018: LD SP,$5678
label_27019: LD IX,$5678
label_27020: LD IY,$5678
label_27021: LD HL,($5678)
label_27022: LD BC,($5678)
label_27023: LD DE,($5678)
label_27024: LD HL,($5678)
label_27025: LD SP,($5678)
label_27026: LD IX,($5678)
label_27027: LD IY,($5678)
label_27028: LD ($5678),HL
label_27029: LD ($5678),BC
label_27030: LD ($5678),DE
label_27031: LD ($5678),HL
label_27032: LD ($5678),SP
label_27033: LD ($5678),IX
label_27034: LD ($5678),IY
label_27035: LD SP,HL
label_27036: LD SP,IX
label_27037: LD SP,IY
label_27038: PUSH BC
label_27039: PUSH DE
label_27040: PUSH HL
label_27041: PUSH AF
label_27042: PUSH IX
label_27043: PUSH IY
label_27044: POP BC
label_27045: POP DE
label_27046: POP HL
label_27047: POP AF
label_27048: POP IX
label_27049: POP IY
label_27050: EX DE,HL
label_27051: EX AF,AF'
label_27052: EXX
label_27053: EX (SP),HL
label_27054: EX (SP),IX
label_27055: EX (SP),IY
label_27056: LDI
label_27057: LDIR
label_27058: LDD
label_27059: LDDR
label_27060: CPI
label_27061: CPIR
label_27062: CPD
label_27063: CPDR
label_27064: ADD A,A
label_27065: ADD A,B
label_27066: ADD A,C
label_27067: ADD A,D
label_27068: ADD A,E
label_27069: ADD A,H
label_27070: ADD A,L
label_27071: ADD A,$12
label_27072: ADD A,(HL)
label_27073: ADD A,(IX + 127)
label_27074: ADD A,(IY - 128)
label_27075: ADC A,A
label_27076: ADC A,B
label_27077: ADC A,C
label_27078: ADC A,D
label_27079: ADC A,E
label_27080: ADC A,H
label_27081: ADC A,L
label_27082: ADC A,$12
label_27083: ADC A,(HL)
label_27084: ADC A,(IX + 127)
label_27085: ADC A,(IY - 128)
label_27086: SUB A
label_27087: SUB B
label_27088: SUB C
label_27089: SUB D
label_27090: SUB E
label_27091: SUB H
label_27092: SUB L
label_27093: SUB $12
label_27094: SUB (HL)
label_27095: SUB (IX + 127)
label_27096: SUB (IY - 128)
label_27097: SBC A,A
label_27098: SBC A,B
label_27099: SBC A,C
label_27100: SBC A,D
label_27101: SBC A,E
label_27102: SBC A,H
label_27103: SBC A,L
label_27104: SBC A,$12
label_27105: SBC A,(HL)
label_27106: SBC A,(IX + 127)
label_27107: SBC A,(IY - 128)
label_27108: AND A
label_27109: AND B
label_27110: AND C
label_27111: AND D
label_27112: AND E
label_27113: AND H
label_27114: AND L
label_27115: AND $12
label_27116: AND (HL)
label_27117: AND (IX + 127)
label_27118: AND (IY - 128)
label_27119: AND A
label_27120: AND B
label_27121: AND C
label_27122: AND D
label_27123: AND E
label_27124: AND H
label_27125: AND L
label_27126: AND $12
label_27127: AND (HL)
label_27128: AND (IX + 127)
label_27129: AND (IY - 128)
label_27130: OR A
label_27131: OR B
label_27132: OR C
label_27133: OR D
label_27134: OR E
label_27135: OR H
label_27136: OR L
label_27137: OR $12
label_27138: OR (HL)
label_27139: OR (IX + 127)
label_27140: OR (IY - 128)
label_27141: XOR A
label_27142: XOR B
label_27143: XOR C
label_27144: XOR D
label_27145: XOR E
label_27146: XOR H
label_27147: XOR L
label_27148: XOR $12
label_27149: XOR (HL)
label_27150: XOR (IX + 127)
label_27151: XOR (IY - 128)
label_27152: CP A
label_27153: CP B
label_27154: CP C
label_27155: CP D
label_27156: CP E
label_27157: CP H
label_27158: CP L
label_27159: CP $12
label_27160: CP (HL)
label_27161: CP (IX + 127)
label_27162: CP (IY - 128)
label_27163: INC A
label_27164: INC B
label_27165: INC C
label_27166: INC D
label_27167: INC E
label_27168: INC H
label_27169: INC L
label_27170: INC (HL)
label_27171: INC (IX + 127)
label_27172: INC (IY - 128)
label_27173: DEC A
label_27174: DEC B
label_27175: DEC C
label_27176: DEC D
label_27177: DEC E
label_27178: DEC H
label_27179: DEC L
label_27180: DEC (HL)
label_27181: DEC (IX + 127)
label_27182: DEC (IY - 128)
label_27183: DAA
label_27184: CPL
label_27185: NEG
label_27186: CCF
label_27187: SCF
label_27188: NOP
label_27189: HALT
label_27190: DI
label_27191: EI
label_27192: IM 0
label_27193: IM 1
label_27194: IM 2
label_27195: ADD HL,BC
label_27196: ADD HL,DE
label_27197: ADD HL,HL
label_27198: ADD HL,SP
label_27199: ADC HL,BC
label_27200: ADC HL,DE
label_27201: ADC HL,HL
label_27202: ADC HL,SP
label_27203: SBC HL,BC
label_27204: SBC HL,DE
label_27205: SBC HL,HL
label_27206: SBC HL,SP
label_27207: ADD IX,BC
label_27208: ADD IX,DE
label_27209: ADD IX,SP
label_27210: ADD IY,BC
label_27211: ADD IY,DE
label_27212: ADD IY,SP
label_27213: INC BC
label_27214: INC DE
label_27215: INC HL
label_27216: INC SP
label_27217: INC IX
label_27218: INC IY
label_27219: DEC BC
label_27220: DEC DE
label_27221: DEC HL
label_27222: DEC SP
label_27223: DEC IX
label_27224: DEC IY
label_27225: RLCA
label_27226: RLA
label_27227: RRCA
label_27228: RRA
label_27229: RLC A
label_27230: RLC B
label_27231: RLC C
label_27232: RLC D
label_27233: RLC E
label_27234: RLC H
label_27235: RLC L
label_27236: RLC (HL)
label_27237: RLC (IX + 127)
label_27238: RLC (IY - 128)
label_27239: RL A
label_27240: RL B
label_27241: RL C
label_27242: RL D
label_27243: RL E
label_27244: RL H
label_27245: RL L
label_27246: RL (HL)
label_27247: RL (IX + 127)
label_27248: RL (IY - 128)
label_27249: RRC A
label_27250: RRC B
label_27251: RRC C
label_27252: RRC D
label_27253: RRC E
label_27254: RRC H
label_27255: RRC L
label_27256: RRC (HL)
label_27257: RRC (IX + 127)
label_27258: RRC (IY - 128)
label_27259: RR A
label_27260: RR B
label_27261: RR C
label_27262: RR D
label_27263: RR E
label_27264: RR H
label_27265: RR L
label_27266: RR (HL)
label_27267: RR (IX + 127)
label_27268: RR (IY - 128)
label_27269: SLA A
label_27270: SLA B
label_27271: SLA C
label_27272: SLA D
label_27273: SLA E
label_27274: SLA H
label_27275: SLA L
label_27276: SLA (HL)
label_27277: SLA (IX + 127)
label_27278: SLA (IY - 128)
label_27279: SRA A
label_27280: SRA B
label_27281: SRA C
label_27282: SRA D
label_27283: SRA E
label_27284: SRA H
label_27285: SRA L
label_27286: SRA (HL)
label_27287: SRA (IX + 127)
label_27288: SRA (IY - 128)
label_27289: SRL A
label_27290: SRL B
label_27291: SRL C
label_27292: SRL D
label_27293: SRL E
label_27294: SRL H
label_27295: SRL L
label_27296: SRL (HL)
label_27297: SRL (IX + 127)
label_27298: SRL (IY - 128)
label_27299: RLD
label_27300: RRD
label_27301: BIT 0,A
label_27302: BIT 1,A
label_27303: BIT 2,A
label_27304: BIT 3,A
label_27305: BIT 4,A
label_27306: BIT 5,A
label_27307: BIT 6,A
label_27308: BIT 7,A
label_27309: BIT 0,B
label_27310: BIT 1,B
label_27311: BIT 2,B
label_27312: BIT 3,B
label_27313: BIT 4,B
label_27314: BIT 5,B
label_27315: BIT 6,B
label_27316: BIT 7,B
label_27317: BIT 0,C
label_27318: BIT 1,C
label_27319: BIT 2,C
label_27320: BIT 3,C
label_27321: BIT 4,C
label_27322: BIT 5,C
label_27323: BIT 6,C
label_27324: BIT 7,C
label_27325: BIT 0,D
label_27326: BIT 1,D
label_27327: BIT 2,D
label_27328: BIT 3,D
label_27329: BIT 4,D
label_27330: BIT 5,D
label_27331: BIT 6,D
label_27332: BIT 7,D
label_27333: BIT 0,E
label_27334: BIT 1,E
label_27335: BIT 2,E
label_27336: BIT 3,E
label_27337: BIT 4,E
label_27338: BIT 5,E
label_27339: BIT 6,E
label_27340: BIT 7,E
label_27341: BIT 0,H
label_27342: BIT 1,H
label_27343: BIT 2,H
label_27344: BIT 3,H
label_27345: BIT 4,H
label_27346: BIT 5,H
label_27347: BIT 6,H
label_27348: BIT 7,H
label_27349: BIT 0,L
label_27350: BIT 1,L
label_27351: BIT 2,L
label_27352: BIT 3,L
label_27353: BIT 4,L
label_27354: BIT 5,L
label_27355: BIT 6,L
label_27356: BIT 7,L
label_27357: BIT 0,(HL)
label_27358: BIT 1,(HL)
label_27359: BIT 2,(HL)
label_27360: BIT 3,(HL)
label_27361: BIT 4,(HL)
label_27362: BIT 5,(HL)
label_27363: BIT 6,(HL)
label_27364: BIT 7,(HL)
label_27365: BIT 0,(IX + 127)
label_27366: BIT 1,(IX + 127)
label_27367: BIT 2,(IX + 127)
label_27368: BIT 3,(IX + 127)
label_27369: BIT 4,(IX + 127)
label_27370: BIT 5,(IX + 127)
label_27371: BIT 6,(IX + 127)
label_27372: BIT 7,(IX + 127)
label_27373: BIT 0,(IY - 128)
label_27374: BIT 1,(IY - 128)
label_27375: BIT 2,(IY - 128)
label_27376: BIT 3,(IY - 128)
label_27377: BIT 4,(IY - 128)
label_27378: BIT 5,(IY - 128)
label_27379: BIT 6,(IY - 128)
label_27380: BIT 7,(IY - 128)
label_27381: SET 0,A
label_27382: SET 1,A
label_27383: SET 2,A
label_27384: SET 3,A
label_27385: SET 4,A
label_27386: SET 5,A
label_27387: SET 6,A
label_27388: SET 7,A
label_27389: SET 0,B
label_27390: SET 1,B
label_27391: SET 2,B
label_27392: SET 3,B
label_27393: SET 4,B
label_27394: SET 5,B
label_27395: SET 6,B
label_27396: SET 7,B
label_27397: SET 0,C
label_27398: SET 1,C
label_27399: SET 2,C
label_27400: SET 3,C
label_27401: SET 4,C
label_27402: SET 5,C
label_27403: SET 6,C
label_27404: SET 7,C
label_27405: SET 0,D
label_27406: SET 1,D
label_27407: SET 2,D
label_27408: SET 3,D
label_27409: SET 4,D
label_27410: SET 5,D
label_27411: SET 6,D
label_27412: SET 7,D
label_27413: SET 0,E
label_27414: SET 1,E
label_27415: SET 2,E
label_27416: SET 3,E
label_27417: SET 4,E
label_27418: SET 5,E
label_27419: SET 6,E
label_27420: SET 7,E
label_27421: SET 0,H
label_27422: SET 1,H
label_27423: SET 2,H
label_27424: SET 3,H
label_27425: SET 4,H
label_27426: SET 5,H
label_27427: SET 6,H
label_27428: SET 7,H
label_27429: SET 0,L
label_27430: SET 1,L
label_27431: SET 2,L
label_27432: SET 3,L
label_27433: SET 4,L
label_27434: SET 5,L
label_27435: SET 6,L
label_27436: SET 7,L
label_27437: SET 0,(HL)
label_27438: SET 1,(HL)
label_27439: SET 2,(HL)
label_27440: SET 3,(HL)
label_27441: SET 4,(HL)
label_27442: SET 5,(HL)
label_27443: SET 6,(HL)
label_27444: SET 7,(HL)
label_27445: SET 0,(IX + 127)
label_27446: SET 1,(IX + 127)
label_27447: SET 2,(IX + 127)
label_27448: SET 3,(IX + 127)
label_27449: SET 4,(IX + 127)
label_27450: SET 5,(IX + 127)
label_27451: SET 6,(IX + 127)
label_27452: SET 7,(IX + 127)
label_27453: SET 0,(IY - 128)
label_27454: SET 1,(IY - 128)
label_27455: SET 2,(IY - 128)
label_27456: SET 3,(IY - 128)
label_27457: SET 4,(IY - 128)
label_27458: SET 5,(IY - 128)
label_27459: SET 6,(IY - 128)
label_27460: SET 7,(IY - 128)
label_27461: RES 0,A
label_27462: RES 1,A
label_27463: RES 2,A
label_27464: RES 3,A
label_27465: RES 4,A
label_27466: RES 5,A
label_27467: RES 6,A
label_27468: RES 7,A
label_27469: RES 0,B
label_27470: RES 1,B
label_27471: RES 2,B
label_27472: RES 3,B
label_27473: RES 4,B
label_27474: RES 5,B
label_27475: RES 6,B
label_27476: RES 7,B
label_27477: RES 0,C
label_27478: RES 1,C
label_27479: RES 2,C
label_27480: RES 3,C
label_27481: RES 4,C
label_27482: RES 5,C
label_27483: RES 6,C
label_27484: RES 7,C
label_27485: RES 0,D
label_27486: RES 1,D
label_27487: RES 2,D
label_27488: RES 3,D
label_27489: RES 4,D
label_27490: RES 5,D
label_27491: RES 6,D
label_27492: RES 7,D
label_27493: RES 0,E
label_27494: RES 1,E
label_27495: RES 2,E
label_27496: RES 3,E
label_27497: RES 4,E
label_27498: RES 5,E
label_27499: RES 6,E
label_27500: RES 7,E
label_27501: RES 0,H
label_27502: RES 1,H
label_27503: RES 2,H
label_27504: RES 3,H
label_27505: RES 4,H
label_27506: RES 5,H
label_27507: RES 6,H
label_27508: RES 7,H
label_27509: RES 0,L
label_27510: RES 1,L
label_27511: RES 2,L
label_27512: RES 3,L
label_27513: RES 4,L
label_27514: RES 5,L
label_27515: RES 6,L
label_27516: RES 7,L
label_27517: RES 0,(HL)
label_27518: RES 1,(HL)
label_27519: RES 2,(HL)
label_27520: RES 3,(HL)
label_27521: RES 4,(HL)
label_27522: RES 5,(HL)
label_27523: RES 6,(HL)
label_27524: RES 7,(HL)
label_27525: RES 0,(IX + 127)
label_27526: RES 1,(IX + 127)
label_27527: RES 2,(IX + 127)
label_27528: RES 3,(IX + 127)
label_27529: RES 4,(IX + 127)
label_27530: RES 5,(IX + 127)
label_27531: RES 6,(IX + 127)
label_27532: RES 7,(IX + 127)
label_27533: RES 0,(IY - 128)
label_27534: RES 1,(IY - 128)
label_27535: RES 2,(IY - 128)
label_27536: RES 3,(IY - 128)
label_27537: RES 4,(IY - 128)
label_27538: RES 5,(IY - 128)
label_27539: RES 6,(IY - 128)
label_27540: RES 7,(IY - 128)
label_27541: JP $5678
label_27542: JP NZ,$5678
label_27543: JP Z,$5678
label_27544: JP NC,$5678
label_27545: JP C,$5678
label_27546: JP PO,$5678
label_27547: JP PE,$5678
label_27548: JP P,$5678
label_27549: JP M,$5678
label_27550: JR $ + 2
label_27551: JR NZ,$ + 2
label_27552: JR Z,$ + 2
label_27553: JR NC,$ + 2
label_27554: JR C,$ + 2
label_27555: JP (HL)
label_27556: JP (IX)
label_27557: JP (IY)
label_27558: DJNZ $ + 2
label_27559: CALL $5678
label_27560: CALL NZ,$5678
label_27561: CALL Z,$5678
label_27562: CALL NC,$5678
label_27563: CALL C,$5678
label_27564: CALL PO,$5678
label_27565: CALL PE,$5678
label_27566: CALL P,$5678
label_27567: CALL M,$5678
label_27568: RET
label_27569: RET NZ
label_27570: RET Z
label_27571: RET NC
label_27572: RET C
label_27573: RET PO
label_27574: RET PE
label_27575: RET P
label_27576: RET M
label_27577: RETI
label_27578: RETN
label_27579: RST $00
label_27580: RST $08
label_27581: RST $10
label_27582: RST $18
label_27583: RST $20
label_27584: RST $28
label_27585: RST $30
label_27586: RST $38
label_27587: IN A,($12)
label_27588: IN A,(C)
label_27589: IN B,(C)
label_27590: IN C,(C)
label_27591: IN D,(C)
label_27592: IN E,(C)
label_27593: IN H,(C)
label_27594: IN L,(C)
label_27595: IN F,(C)
label_27596: INI
label_27597: INIR
label_27598: IND
label_27599: INDR
label_27600: OUT ($12),A
label_27601: OUT (C),A
label_27602: OUT (C),B
label_27603: OUT (C),C
label_27604: OUT (C),D
label_27605: OUT (C),E
label_27606: OUT (C),H
label_27607: OUT (C),L
label_27608: OUTI
label_27609: OTIR
label_27610: OUTD
label_27611: OTDR
label_27612: LD A,A
label_27613: LD A,B
label_27614: LD A,C
label_27615: LD A,D
label_27616: LD A,E
label_27617: LD A,H
label_27618: LD A,L
label_27619: LD B,A
label_27620: LD B,B
label_27621: LD B,C
label_27622: LD B,D
label_27623: LD B,E
label_27624: LD B,H
label_27625: LD B,L
label_27626: LD C,A
label_27627: LD C,B
label_27628: LD C,C
label_27629: LD C,D
label_27630: LD C,E
label_27631: LD C,H
label_27632: LD C,L
label_27633: LD D,A
label_27634: LD D,B
label_27635: LD D,C
label_27636: LD D,D
label_27637: LD D,E
label_27638: LD D,H
label_27639: LD D,L
label_27640: LD E,A
label_27641: LD E,B
label_27642: LD E,C
label_27643: LD E,D
label_27644: LD E,E
label_27645: LD E,H
label_27646: LD E,L
label_27647: LD H,A
label_27648: LD H,B
label_27649: LD H,C
label_27650: LD H,D
label_27651: LD H,E
label_27652: LD H,H
label_27653: LD H,L
label_27654: LD L,A
label_27655: LD L,B
label_27656: LD L,C
label_27657: LD L,D
label_27658: LD L,E
label_27659: LD L,H
label_27660: LD L,L
label_27661: LD A,$12
label_27662: LD B,$12
label_27663: LD C,$12
label_27664: LD D,$12
label_27665: LD E,$12
label_27666: LD H,$12
label_27667: LD L,$12
label_27668: LD A,(HL)
label_27669: LD B,(HL)
label_27670: LD C,(HL)
label_27671: LD D,(HL)
label_27672: LD E,(HL)
label_27673: LD H,(HL)
label_27674: LD L,(HL)
label_27675: LD A,(IX + 127)
label_27676: LD B,(IX + 127)
label_27677: LD C,(IX + 127)
label_27678: LD D,(IX + 127)
label_27679: LD E,(IX + 127)
label_27680: LD H,(IX + 127)
label_27681: LD L,(IX + 127)
label_27682: LD A,(IY - 128)
label_27683: LD B,(IY - 128)
label_27684: LD C,(IY - 128)
label_27685: LD D,(IY - 128)
label_27686: LD E,(IY - 128)
label_27687: LD H,(IY - 128)
label_27688: LD L,(IY - 128)
label_27689: LD (HL),A
label_27690: LD (HL),B
label_27691: LD (HL),C
label_27692: LD (HL),D
label_27693: LD (HL),E
label_27694: LD (HL),H
label_27695: LD (HL),L
label_27696: LD (IX + 127),A
label_27697: LD (IX + 127),B
label_27698: LD (IX + 127),C
label_27699: LD (IX + 127),D
label_27700: LD (IX + 127),E
label_27701: LD (IX + 127),H
label_27702: LD (IX + 127),L
label_27703: LD (IY - 128),A
label_27704: LD (IY - 128),B
label_27705: LD (IY - 128),C
label_27706: LD (IY - 128),D
label_27707: LD (IY - 128),E
label_27708: LD (IY - 128),H
label_27709: LD (IY - 128),L
label_27710: LD (HL),$12
label_27711: LD (IX + 127),$12
label_27712: LD (IY - 128),$12
label_27713: LD A,(BC)
label_27714: LD A,(DE)
label_27715: LD A,($5678)
label_27716: LD (BC),A
label_27717: LD (DE),A
label_27718: LD ($5678),A
label_27719: LD A,I
label_27720: LD A,R
label_27721: LD I,A
label_27722: LD R,A
label_27723: LD BC,$5678
label_27724: LD DE,$5678
label_27725: LD HL,$5678
label_27726: LD SP,$5678
label_27727: LD IX,$5678
label_27728: LD IY,$5678
label_27729: LD HL,($5678)
label_27730: LD BC,($5678)
label_27731: LD DE,($5678)
label_27732: LD HL,($5678)
label_27733: LD SP,($5678)
label_27734: LD IX,($5678)
label_27735: LD IY,($5678)
label_27736: LD ($5678),HL
label_27737: LD ($5678),BC
label_27738: LD ($5678),DE
label_27739: LD ($5678),HL
label_27740: LD ($5678),SP
label_27741: LD ($5678),IX
label_27742: LD ($5678),IY
label_27743: LD SP,HL
label_27744: LD SP,IX
label_27745: LD SP,IY
label_27746: PUSH BC
label_27747: PUSH DE
label_27748: PUSH HL
label_27749: PUSH AF
label_27750: PUSH IX
label_27751: PUSH IY
label_27752: POP BC
label_27753: POP DE
label_27754: POP HL
label_27755: POP AF
label_27756: POP IX
label_27757: POP IY
label_27758: EX DE,HL
label_27759: EX AF,AF'
label_27760: EXX
label_27761: EX (SP),HL
label_27762: EX (SP),IX
label_27763: EX (SP),IY
label_27764: LDI
label_27765: LDIR
label_27766: LDD
label_27767: LDDR
label_27768: CPI
label_27769: CPIR
label_27770: CPD
label_27771: CPDR
label_27772: ADD A,A
label_27773: ADD A,B
label_27774: ADD A,C
label_27775: ADD A,D
label_27776: ADD A,E
label_27777: ADD A,H
label_27778: ADD A,L
label_27779: ADD A,$12
label_27780: ADD A,(HL)
label_27781: ADD A,(IX + 127)
label_27782: ADD A,(IY - 128)
label_27783: ADC A,A
label_27784: ADC A,B
label_27785: ADC A,C
label_27786: ADC A,D
label_27787: ADC A,E
label_27788: ADC A,H
label_27789: ADC A,L
label_27790: ADC A,$12
label_27791: ADC A,(HL)
label_27792: ADC A,(IX + 127)
label_27793: ADC A,(IY - 128)
label_27794: SUB A
label_27795: SUB B
label_27796: SUB C
label_27797: SUB D
label_27798: SUB E
label_27799: SUB H
label_27800: SUB L
label_27801: SUB $12
label_27802: SUB (HL)
label_27803: SUB (IX + 127)
label_27804: SUB (IY - 128)
label_27805: SBC A,A
label_27806: SBC A,B
label_27807: SBC A,C
label_27808: SBC A,D
label_27809: SBC A,E
label_27810: SBC A,H
label_27811: SBC A,L
label_27812: SBC A,$12
label_27813: SBC A,(HL)
label_27814: SBC A,(IX + 127)
label_27815: SBC A,(IY - 128)
label_27816: AND A
label_27817: AND B
label_27818: AND C
label_27819: AND D
label_27820: AND E
label_27821: AND H
label_27822: AND L
label_27823: AND $12
label_27824: AND (HL)
label_27825: AND (IX + 127)
label_27826: AND (IY - 128)
label_27827: AND A
label_27828: AND B
label_27829: AND C
label_27830: AND D
label_27831: AND E
label_27832: AND H
label_27833: AND L
label_27834: AND $12
label_27835: AND (HL)
label_27836: AND (IX + 127)
label_27837: AND (IY - 128)
label_27838: OR A
label_27839: OR B
label_27840: OR C
label_27841: OR D
label_27842: OR E
label_27843: OR H
label_27844: OR L
label_27845: OR $12
label_27846: OR (HL)
label_27847: OR (IX + 127)
label_27848: OR (IY - 128)
label_27849: XOR A
label_27850: XOR B
label_27851: XOR C
label_27852: XOR D
label_27853: XOR E
label_27854: XOR H
label_27855: XOR L
label_27856: XOR $12
label_27857: XOR (HL)
label_27858: XOR (IX + 127)
label_27859: XOR (IY - 128)
label_27860: CP A
label_27861: CP B
label_27862: CP C
label_27863: CP D
label_27864: CP E
label_27865: CP H
label_27866: CP L
label_27867: CP $12
label_27868: CP (HL)
label_27869: CP (IX + 127)
label_27870: CP (IY - 128)
label_27871: INC A
label_27872: INC B
label_27873: INC C
label_27874: INC D
label_27875: INC E
label_27876: INC H
label_27877: INC L
label_27878: INC (HL)
label_27879: INC (IX + 127)
label_27880: INC (IY - 128)
label_27881: DEC A
label_27882: DEC B
label_27883: DEC C
label_27884: DEC D
label_27885: DEC E
label_27886: DEC H
label_27887: DEC L
label_27888: DEC (HL)
label_27889: DEC (IX + 127)
label_27890: DEC (IY - 128)
label_27891: DAA
label_27892: CPL
label_27893: NEG
label_27894: CCF
label_27895: SCF
label_27896: NOP
label_27897: HALT
label_27898: DI
label_27899: EI
label_27900: IM 0
label_27901: IM 1
label_27902: IM 2
label_27903: ADD HL,BC
label_27904: ADD HL,DE
label_27905: ADD HL,HL
label_27906: ADD HL,SP
label_27907: ADC HL,BC
label_27908: ADC HL,DE
label_27909: ADC HL,HL
label_27910: ADC HL,SP
label_27911: SBC HL,BC
label_27912: SBC HL,DE
label_27913: SBC HL,HL
label_27914: SBC HL,SP
label_27915: ADD IX,BC
label_27916: ADD IX,DE
label_27917: ADD IX,SP
label_27918: ADD IY,BC
label_27919: ADD IY,DE
label_27920: ADD IY,SP
label_27921: INC BC
label_27922: INC DE
label_27923: INC HL
label_27924: INC SP
label_27925: INC IX
label_27926: INC IY
label_27927: DEC BC
label_27928: DEC DE
label_27929: DEC HL
label_27930: DEC SP
label_27931: DEC IX
label_27932: DEC IY
label_27933: RLCA
label_27934: RLA
label_27935: RRCA
label_27936: RRA
label_27937: RLC A
label_27938: RLC B
label_27939: RLC C
label_27940: RLC D
label_27941: RLC E
label_27942: RLC H
label_27943: RLC L
label_27944: RLC (HL)
label_27945: RLC (IX + 127)
label_27946: RLC (IY - 128)
label_27947: RL A
label_27948: RL B
label_27949: RL C
label_27950: RL D
label_27951: RL E
label_27952: RL H
label_27953: RL L
label_27954: RL (HL)
label_27955: RL (IX + 127)
label_27956: RL (IY - 128)
label_27957: RRC A
label_27958: RRC B
label_27959: RRC C
label_27960: RRC D
label_27961: RRC E
label_27962: RRC H
label_27963: RRC L
label_27964: RRC (HL)
label_27965: RRC (IX + 127)
label_27966: RRC (IY - 128)
label_27967: RR A
label_27968: RR B
label_27969: RR C
label_27970: RR D
label_27971: RR E
label_27972: RR H
label_27973: RR L
label_27974: RR (HL)
label_27975: RR (IX + 127)
label_27976: RR (IY - 128)
label_27977: SLA A
label_27978: SLA B
label_27979: SLA C
label_27980: SLA D
label_27981: SLA E
label_27982: SLA H
label_27983: SLA L
label_27984: SLA (HL)
label_27985: SLA (IX + 127)
label_27986: SLA (IY - 128)
label_27987: SRA A
label_27988: SRA B
label_27989: SRA C
label_27990: SRA D
label_27991: SRA E
label_27992: SRA H
label_27993: SRA L
label_27994: SRA (HL)
label_27995: SRA (IX + 127)
label_27996: SRA (IY - 128)
label_27997: SRL A
label_27998: SRL B
label_27999: SRL C
label_28000: SRL D
label_28001: SRL E
label_28002: SRL H
label_28003: SRL L
label_28004: SRL (HL)
label_28005: SRL (IX + 127)
label_28006: SRL (IY - 128)
label_28007: RLD
label_28008: RRD
label_28009: BIT 0,A
label_28010: BIT 1,A
label_28011: BIT 2,A
label_28012: BIT 3,A
label_28013: BIT 4,A
label_28014: BIT 5,A
label_28015: BIT 6,A
label_28016: BIT 7,A
label_28017: BIT 0,B
label_28018: BIT 1,B
label_28019: BIT 2,B
label_28020: BIT 3,B
label_28021: BIT 4,B
label_28022: BIT 5,B
label_28023: BIT 6,B
label_28024: BIT 7,B
label_28025: BIT 0,C
label_28026: BIT 1,C
label_28027: BIT 2,C
label_28028: BIT 3,C
label_28029: BIT 4,C
label_28030: BIT 5,C
label_28031: BIT 6,C
label_28032: BIT 7,C
label_28033: BIT 0,D
label_28034: BIT 1,D
label_28035: BIT 2,D
label_28036: BIT 3,D
label_28037: BIT 4,D
label_28038: BIT 5,D
label_28039: BIT 6,D
label_28040: BIT 7,D
label_28041: BIT 0,E
label_28042: BIT 1,E
label_28043: BIT 2,E
label_28044: BIT 3,E
label_28045: BIT 4,E
label_28046: BIT 5,E
label_28047: BIT 6,E
label_28048: BIT 7,E
label_28049: BIT 0,H
label_28050: BIT 1,H
label_28051: BIT 2,H
label_28052: BIT 3,H
label_28053: BIT 4,H
label_28054: BIT 5,H
label_28055: BIT 6,H
label_28056: BIT 7,H
label_28057: BIT 0,L
label_28058: BIT 1,L
label_28059: BIT 2,L
label_28060: BIT 3,L
label_28061: BIT 4,L
label_28062: BIT 5,L
label_28063: BIT 6,L
label_28064: BIT 7,L
label_28065: BIT 0,(HL)
label_28066: BIT 1,(HL)
label_28067: BIT 2,(HL)
label_28068: BIT 3,(HL)
label_28069: BIT 4,(HL)
label_28070: BIT 5,(HL)
label_28071: BIT 6,(HL)
label_28072: BIT 7,(HL)
label_28073: BIT 0,(IX + 127)
label_28074: BIT 1,(IX + 127)
label_28075: BIT 2,(IX + 127)
label_28076: BIT 3,(IX + 127)
label_28077: BIT 4,(IX + 127)
label_28078: BIT 5,(IX + 127)
label_28079: BIT 6,(IX + 127)
label_28080: BIT 7,(IX + 127)
label_28081: BIT 0,(IY - 128)
label_28082: BIT 1,(IY - 128)
label_28083: BIT 2,(IY - 128)
label_28084: BIT 3,(IY - 128)
label_28085: BIT 4,(IY - 128)
label_28086: BIT 5,(IY - 128)
label_28087: BIT 6,(IY - 128)
label_28088: BIT 7,(IY - 128)
label_28089: SET 0,A
label_28090: SET 1,A
label_28091: SET 2,A
label_28092: SET 3,A
label_28093: SET 4,A
label_28094: SET 5,A
label_28095: SET 6,A
label_28096: SET 7,A
label_28097: SET 0,B
label_28098: SET 1,B
label_28099: SET 2,B
label_28100: SET 3,B
label_28101: SET 4,B
label_28102: SET 5,B
label_28103: SET 6,B
label_28104: SET 7,B
label_28105: SET 0,C
label_28106: SET 1,C
label_28107: SET 2,C
label_28108: SET 3,C
label_28109: SET 4,C
label_28110: SET 5,C
label_28111: SET 6,C
label_28112: SET 7,C
label_28113: SET 0,D
label_28114: SET 1,D
label_28115: SET 2,D
label_28116: SET 3,D
label_28117: SET 4,D
label_28118: SET 5,D
label_28119: SET 6,D
label_28120: SET 7,D
label_28121: SET 0,E
label_28122: SET 1,E
label_28123: SET 2,E
label_28124: SET 3,E
label_28125: SET 4,E
label_28126: SET 5,E
label_28127: SET 6,E
label_28128: SET 7,E
label_28129: SET 0,H
label_28130: SET 1,H
label_28131: SET 2,H
label_28132: SET 3,H
label_28133: SET 4,H
label_28134: SET 5,H
label_28135: SET 6,H
label_28136: SET 7,H
label_28137: SET 0,L
label_28138: SET 1,L
label_28139: SET 2,L
label_28140: SET 3,L
label_28141: SET 4,L
label_28142: SET 5,L
label_28143: SET 6,L
label_28144: SET 7,L
label_28145: SET 0,(HL)
label_28146: SET 1,(HL)
label_28147: SET 2,(HL)
label_28148: SET 3,(HL)
label_28149: SET 4,(HL)
label_28150: SET 5,(HL)
label_28151: SET 6,(HL)
label_28152: SET 7,(HL)
label_28153: SET 0,(IX + 127)
label_28154: SET 1,(IX + 127)
label_28155: SET 2,(IX + 127)
label_28156: SET 3,(IX + 127)
label_28157: SET 4,(IX + 127)
label_28158: SET 5,(IX + 127)
label_28159: SET 6,(IX + 127)
label_28160: SET 7,(IX + 127)
label_28161: SET 0,(IY - 128)
label_28162: SET 1,(IY - 128)
label_28163: SET 2,(IY - 128)
label_28164: SET 3,(IY - 128)
label_28165: SET 4,(IY - 128)
label_28166: SET 5,(IY - 128)
label_28167: SET 6,(IY - 128)
label_28168: SET 7,(IY - 128)
label_28169: RES 0,A
label_28170: RES 1,A
label_28171: RES 2,A
label_28172: RES 3,A
label_28173: RES 4,A
label_28174: RES 5,A
label_28175: RES 6,A
label_28176: RES 7,A
label_28177: RES 0,B
label_28178: RES 1,B
label_28179: RES 2,B
label_28180: RES 3,B
label_28181: RES 4,B
label_28182: RES 5,B
label_28183: RES 6,B
label_28184: RES 7,B
label_28185: RES 0,C
label_28186: RES 1,C
label_28187: RES 2,C
label_28188: RES 3,C
label_28189: RES 4,C
label_28190: RES 5,C
label_28191: RES 6,C
label_28192: RES 7,C
label_28193: RES 0,D
label_28194: RES 1,D
label_28195: RES 2,D
label_28196: RES 3,D
label_28197: RES 4,D
label_28198: RES 5,D
label_28199: RES 6,D
label_28200: RES 7,D
label_28201: RES 0,E
label_28202: RES 1,E
label_28203: RES 2,E
label_28204: RES 3,E
label_28205: RES 4,E
label_28206: RES 5,E
label_28207: RES 6,E
label_28208: RES 7,E
label_28209: RES 0,H
label_28210: RES 1,H
label_28211: RES 2,H
label_28212: RES 3,H
label_28213: RES 4,H
label_28214: RES 5,H
label_28215: RES 6,H
label_28216: RES 7,H
label_28217: RES 0,L
label_28218: RES 1,L
label_28219: RES 2,L
label_28220: RES 3,L
label_28221: RES 4,L
label_28222: RES 5,L
label_28223: RES 6,L
label_28224: RES 7,L
label_28225: RES 0,(HL)
label_28226: RES 1,(HL)
label_28227: RES 2,(HL)
label_28228: RES 3,(HL)
label_28229: RES 4,(HL)
label_28230: RES 5,(HL)
label_28231: RES 6,(HL)
label_28232: RES 7,(HL)
label_28233: RES 0,(IX + 127)
label_28234: RES 1,(IX + 127)
label_28235: RES 2,(IX + 127)
label_28236: RES 3,(IX + 127)
label_28237: RES 4,(IX + 127)
label_28238: RES 5,(IX + 127)
label_28239: RES 6,(IX + 127)
label_28240: RES 7,(IX + 127)
label_28241: RES 0,(IY - 128)
label_28242: RES 1,(IY - 128)
label_28243: RES 2,(IY - 128)
label_28244: RES 3,(IY - 128)
label_28245: RES 4,(IY - 128)
label_28246: RES 5,(IY - 128)
label_28247: RES 6,(IY - 128)
label_28248: RES 7,(IY - 128)
label_28249: JP $5678
label_28250: JP NZ,$5678
label_28251: JP Z,$5678
label_28252: JP NC,$5678
label_28253: JP C,$5678
label_28254: JP PO,$5678
label_28255: JP PE,$5678
label_28256: JP P,$5678
label_28257: JP M,$5678
label_28258: JR $ + 2
label_28259: JR NZ,$ + 2
label_28260: JR Z,$ + 2
label_28261: JR NC,$ + 2
label_28262: JR C,$ + 2
label_28263: JP (HL)
label_28264: JP (IX)
label_28265: JP (IY)
label_28266: DJNZ $ + 2
label_28267: CALL $5678
label_28268: CALL NZ,$5678
label_28269: CALL Z,$5678
label_28270: CALL NC,$5678
label_28271: CALL C,$5678
label_28272: CALL PO,$5678
label_28273: CALL PE,$5678
label_28274: CALL P,$5678
label_28275: CALL M,$5678
label_28276: RET
label_28277: RET NZ
label_28278: RET Z
label_28279: RET NC
label_28280: RET C
label_28281: RET PO
label_28282: RET PE
label_28283: RET P
label_28284: RET M
label_28285: RETI
label_28286: RETN
label_28287: RST $00
label_28288: RST $08
label_28289: RST $10
label_28290: RST $18
label_28291: RST $20
label_28292: RST $28
label_28293: RST $30
label_28294: RST $38
label_28295: IN A,($12)
label_28296: IN A,(C)
label_28297: IN B,(C)
label_28298: IN C,(C)
label_28299: IN D,(C)
label_28300: IN E,(C)
label_28301: IN H,(C)
label_28302: IN L,(C)
label_28303: IN F,(C)
label_28304: INI
label_28305: INIR
label_28306: IND
label_28307: INDR
label_28308: OUT ($12),A
label_28309: OUT (C),A
label_28310: OUT (C),B
label_28311: OUT (C),C
label_28312: OUT (C),D
label_28313: OUT (C),E
label_28314: OUT (C),H
label_28315: OUT (C),L
label_28316: OUTI
label_28317: OTIR
label_28318: OUTD
label_28319: OTDR
label_28320: LD A,A
label_28321: LD A,B
label_28322: LD A,C
label_28323: LD A,D
label_28324: LD A,E
label_28325: LD A,H
label_28326: LD A,L
label_28327: LD B,A
label_28328: LD B,B
label_28329: LD B,C
label_28330: LD B,D
label_28331: LD B,E
label_28332: LD B,H
label_28333: LD B,L
label_28334: LD C,A
label_28335: LD C,B
label_28336: LD C,C
label_28337: LD C,D
label_28338: LD C,E
label_28339: LD C,H
label_28340: LD C,L
label_28341: LD D,A
label_28342: LD D,B
label_28343: LD D,C
label_28344: LD D,D
label_28345: LD D,E
label_28346: LD D,H
label_28347: LD D,L
label_28348: LD E,A
label_28349: LD E,B
label_28350: LD E,C
label_28351: LD E,D
label_28352: LD E,E
label_28353: LD E,H
label_28354: LD E,L
label_28355: LD H,A
label_28356: LD H,B
label_28357: LD H,C
label_28358: LD H,D
label_28359: LD H,E
label_28360: LD H,H
label_28361: LD H,L
label_28362: LD L,A
label_28363: LD L,B
label_28364: LD L,C
label_28365: LD L,D
label_28366: LD L,E
label_28367: LD L,H
label_28368: LD L,L
label_28369: LD A,$12
label_28370: LD B,$12
label_28371: LD C,$12
label_28372: LD D,$12
label_28373: LD E,$12
label_28374: LD H,$12
label_28375: LD L,$12
label_28376: LD A,(HL)
label_28377: LD B,(HL)
label_28378: LD C,(HL)
label_28379: LD D,(HL)
label_28380: LD E,(HL)
label_28381: LD H,(HL)
label_28382: LD L,(HL)
label_28383: LD A,(IX + 127)
label_28384: LD B,(IX + 127)
label_28385: LD C,(IX + 127)
label_28386: LD D,(IX + 127)
label_28387: LD E,(IX + 127)
label_28388: LD H,(IX + 127)
label_28389: LD L,(IX + 127)
label_28390: LD A,(IY - 128)
label_28391: LD B,(IY - 128)
label_28392: LD C,(IY - 128)
label_28393: LD D,(IY - 128)
label_28394: LD E,(IY - 128)
label_28395: LD H,(IY - 128)
label_28396: LD L,(IY - 128)
label_28397: LD (HL),A
label_28398: LD (HL),B
label_28399: LD (HL),C
label_28400: LD (HL),D
label_28401: LD (HL),E
label_28402: LD (HL),H
label_28403: LD (HL),L
label_28404: LD (IX + 127),A
label_28405: LD (IX + 127),B
label_28406: LD (IX + 127),C
label_28407: LD (IX + 127),D
label_28408: LD (IX + 127),E
label_28409: LD (IX + 127),H
label_28410: LD (IX + 127),L
label_28411: LD (IY - 128),A
label_28412: LD (IY - 128),B
label_28413: LD (IY - 128),C
label_28414: LD (IY - 128),D
label_28415: LD (IY - 128),E
label_28416: LD (IY - 128),H
label_28417: LD (IY - 128),L
label_28418: LD (HL),$12
label_28419: LD (IX + 127),$12
label_28420: LD (IY - 128),$12
label_28421: LD A,(BC)
label_28422: LD A,(DE)
label_28423: LD A,($5678)
label_28424: LD (BC),A
label_28425: LD (DE),A
label_28426: LD ($5678),A
label_28427: LD A,I
label_28428: LD A,R
label_28429: LD I,A
label_28430: LD R,A
label_28431: LD BC,$5678
label_28432: LD DE,$5678
label_28433: LD HL,$5678
label_28434: LD SP,$5678
label_28435: LD IX,$5678
label_28436: LD IY,$5678
label_28437: LD HL,($5678)
label_28438: LD BC,($5678)
label_28439: LD DE,($5678)
label_28440: LD HL,($5678)
label_28441: LD SP,($5678)
label_28442: LD IX,($5678)
label_28443: LD IY,($5678)
label_28444: LD ($5678),HL
label_28445: LD ($5678),BC
label_28446: LD ($5678),DE
label_28447: LD ($5678),HL
label_28448: LD ($5678),SP
label_28449: LD ($5678),IX
label_28450: LD ($5678),IY
label_28451: LD SP,HL
label_28452: LD SP,IX
label_28453: LD SP,IY
label_28454: PUSH BC
label_28455: PUSH DE
label_28456: PUSH HL
label_28457: PUSH AF
label_28458: PUSH IX
label_28459: PUSH IY
label_28460: POP BC
label_28461: POP DE
label_28462: POP HL
label_28463: POP AF
label_28464: POP IX
label_28465: POP IY
label_28466: EX DE,HL
label_28467: EX AF,AF'
label_28468: EXX
label_28469: EX (SP),HL
label_28470: EX (SP),IX
label_28471: EX (SP),IY
label_28472: LDI
label_28473: LDIR
label_28474: LDD
label_28475: LDDR
label_28476: CPI
label_28477: CPIR
label_28478: CPD
label_28479: CPDR
label_28480: ADD A,A
label_28481: ADD A,B
label_28482: ADD A,C
label_28483: ADD A,D
label_28484: ADD A,E
label_28485: ADD A,H
label_28486: ADD A,L
label_28487: ADD A,$12
label_28488: ADD A,(HL)
label_28489: ADD A,(IX + 127)
label_28490: ADD A,(IY - 128)
label_28491: ADC A,A
label_28492: ADC A,B
label_28493: ADC A,C
label_28494: ADC A,D
label_28495: ADC A,E
label_28496: ADC A,H
label_28497: ADC A,L
label_28498: ADC A,$12
label_28499: ADC A,(HL)
label_28500: ADC A,(IX + 127)
label_28501: ADC A,(IY - 128)
label_28502: SUB A
label_28503: SUB B
label_28504: SUB C
label_28505: SUB D
label_28506: SUB E
label_28507: SUB H
label_28508: SUB L
label_28509: SUB $12
label_28510: SUB (HL)
label_28511: SUB (IX + 127)
label_28512: SUB (IY - 128)
label_28513: SBC A,A
label_28514: SBC A,B
label_28515: SBC A,C
label_28516: SBC A,D
label_28517: SBC A,E
label_28518: SBC A,H
label_28519: SBC A,L
label_28520: SBC A,$12
label_28521: SBC A,(HL)
label_28522: SBC A,(IX + 127)
label_28523: SBC A,(IY - 128)
label_28524: AND A
label_28525: AND B
label_28526: AND C
label_28527: AND D
label_28528: AND E
label_28529: AND H
label_28530: AND L
label_28531: AND $12
label_28532: AND (HL)
label_28533: AND (IX + 127)
label_28534: AND (IY - 128)
label_28535: AND A
label_28536: AND B
label_28537: AND C
label_28538: AND D
label_28539: AND E
label_28540: AND H
label_28541: AND L
label_28542: AND $12
label_28543: AND (HL)
label_28544: AND (IX + 127)
label_28545: AND (IY - 128)
label_28546: OR A
label_28547: OR B
label_28548: OR C
label_28549: OR D
label_28550: OR E
label_28551: OR H
label_28552: OR L
label_28553: OR $12
label_28554: OR (HL)
label_28555: OR (IX + 127)
label_28556: OR (IY - 128)
label_28557: XOR A
label_28558: XOR B
label_28559: XOR C
label_28560: XOR D
label_28561: XOR E
label_28562: XOR H
label_28563: XOR L
label_28564: XOR $12
label_28565: XOR (HL)
label_28566: XOR (IX + 127)
label_28567: XOR (IY - 128)
label_28568: CP A
label_28569: CP B
label_28570: CP C
label_28571: CP D
label_28572: CP E
label_28573: CP H
label_28574: CP L
label_28575: CP $12
label_28576: CP (HL)
label_28577: CP (IX + 127)
label_28578: CP (IY - 128)
label_28579: INC A
label_28580: INC B
label_28581: INC C
label_28582: INC D
label_28583: INC E
label_28584: INC H
label_28585: INC L
label_28586: INC (HL)
label_28587: INC (IX + 127)
label_28588: INC (IY - 128)
label_28589: DEC A
label_28590: DEC B
label_28591: DEC C
label_28592: DEC D
label_28593: DEC E
label_28594: DEC H
label_28595: DEC L
label_28596: DEC (HL)
label_28597: DEC (IX + 127)
label_28598: DEC (IY - 128)
label_28599: DAA
label_28600: CPL
label_28601: NEG
label_28602: CCF
label_28603: SCF
label_28604: NOP
label_28605: HALT
label_28606: DI
label_28607: EI
label_28608: IM 0
label_28609: IM 1
label_28610: IM 2
label_28611: ADD HL,BC
label_28612: ADD HL,DE
label_28613: ADD HL,HL
label_28614: ADD HL,SP
label_28615: ADC HL,BC
label_28616: ADC HL,DE
label_28617: ADC HL,HL
label_28618: ADC HL,SP
label_28619: SBC HL,BC
label_28620: SBC HL,DE
label_28621: SBC HL,HL
label_28622: SBC HL,SP
label_28623: ADD IX,BC
label_28624: ADD IX,DE
label_28625: ADD IX,SP
label_28626: ADD IY,BC
label_28627: ADD IY,DE
label_28628: ADD IY,SP
label_28629: INC BC
label_28630: INC DE
label_28631: INC HL
label_28632: INC SP
label_28633: INC IX
label_28634: INC IY
label_28635: DEC BC
label_28636: DEC DE
label_28637: DEC HL
label_28638: DEC SP
label_28639: DEC IX
label_28640: DEC IY
label_28641: RLCA
label_28642: RLA
label_28643: RRCA
label_28644: RRA
label_28645: RLC A
label_28646: RLC B
label_28647: RLC C
label_28648: RLC D
label_28649: RLC E
label_28650: RLC H
label_28651: RLC L
label_28652: RLC (HL)
label_28653: RLC (IX + 127)
label_28654: RLC (IY - 128)
label_28655: RL A
label_28656: RL B
label_28657: RL C
label_28658: RL D
label_28659: RL E
label_28660: RL H
label_28661: RL L
label_28662: RL (HL)
label_28663: RL (IX + 127)
label_28664: RL (IY - 128)
label_28665: RRC A
label_28666: RRC B
label_28667: RRC C
label_28668: RRC D
label_28669: RRC E
label_28670: RRC H
label_28671: RRC L
label_28672: RRC (HL)
label_28673: RRC (IX + 127)
label_28674: RRC (IY - 128)
label_28675: RR A
label_28676: RR B
label_28677: RR C
label_28678: RR D
label_28679: RR E
label_28680: RR H
label_28681: RR L
label_28682: RR (HL)
label_28683: RR (IX + 127)
label_28684: RR (IY - 128)
label_28685: SLA A
label_28686: SLA B
label_28687: SLA C
label_28688: SLA D
label_28689: SLA E
label_28690: SLA H
label_28691: SLA L
label_28692: SLA (HL)
label_28693: SLA (IX + 127)
label_28694: SLA (IY - 128)
label_28695: SRA A
label_28696: SRA B
label_28697: SRA C
label_28698: SRA D
label_28699: SRA E
label_28700: SRA H
label_28701: SRA L
label_28702: SRA (HL)
label_28703: SRA (IX + 127)
label_28704: SRA (IY - 128)
label_28705: SRL A
label_28706: SRL B
label_28707: SRL C
label_28708: SRL D
label_28709: SRL E
label_28710: SRL H
label_28711: SRL L
label_28712: SRL (HL)
label_28713: SRL (IX + 127)
label_28714: SRL (IY - 128)
label_28715: RLD
label_28716: RRD
label_28717: BIT 0,A
label_28718: BIT 1,A
label_28719: BIT 2,A
label_28720: BIT 3,A
label_28721: BIT 4,A
label_28722: BIT 5,A
label_28723: BIT 6,A
label_28724: BIT 7,A
label_28725: BIT 0,B
label_28726: BIT 1,B
label_28727: BIT 2,B
label_28728: BIT 3,B
label_28729: BIT 4,B
label_28730: BIT 5,B
label_28731: BIT 6,B
label_28732: BIT 7,B
label_28733: BIT 0,C
label_28734: BIT 1,C
label_28735: BIT 2,C
label_28736: BIT 3,C
label_28737: BIT 4,C
label_28738: BIT 5,C
label_28739: BIT 6,C
label_28740: BIT 7,C
label_28741: BIT 0,D
label_28742: BIT 1,D
label_28743: BIT 2,D
label_28744: BIT 3,D
label_28745: BIT 4,D
label_28746: BIT 5,D
label_28747: BIT 6,D
label_28748: BIT 7,D
label_28749: BIT 0,E
label_28750: BIT 1,E
label_28751: BIT 2,E
label_28752: BIT 3,E
label_28753: BIT 4,E
label_28754: BIT 5,E
label_28755: BIT 6,E
label_28756: BIT 7,E
label_28757: BIT 0,H
label_28758: BIT 1,H
label_28759: BIT 2,H
label_28760: BIT 3,H
label_28761: BIT 4,H
label_28762: BIT 5,H
label_28763: BIT 6,H
label_28764: BIT 7,H
label_28765: BIT 0,L
label_28766: BIT 1,L
label_28767: BIT 2,L
label_28768: BIT 3,L
label_28769: BIT 4,L
label_28770: BIT 5,L
label_28771: BIT 6,L
label_28772: BIT 7,L
label_28773: BIT 0,(HL)
label_28774: BIT 1,(HL)
label_28775: BIT 2,(HL)
label_28776: BIT 3,(HL)
label_28777: BIT 4,(HL)
label_28778: BIT 5,(HL)
label_28779: BIT 6,(HL)
label_28780: BIT 7,(HL)
label_28781: BIT 0,(IX + 127)
label_28782: BIT 1,(IX + 127)
label_28783: BIT 2,(IX + 127)
label_28784: BIT 3,(IX + 127)
label_28785: BIT 4,(IX + 127)
label_28786: BIT 5,(IX + 127)
label_28787: BIT 6,(IX + 127)
label_28788: BIT 7,(IX + 127)
label_28789: BIT 0,(IY - 128)
label_28790: BIT 1,(IY - 128)
label_28791: BIT 2,(IY - 128)
label_28792: BIT 3,(IY - 128)
label_28793: BIT 4,(IY - 128)
label_28794: BIT 5,(IY - 128)
label_28795: BIT 6,(IY - 128)
label_28796: BIT 7,(IY - 128)
label_28797: SET 0,A
label_28798: SET 1,A
label_28799: SET 2,A
label_28800: SET 3,A
label_28801: SET 4,A
label_28802: SET 5,A
label_28803: SET 6,A
label_28804: SET 7,A
label_28805: SET 0,B
label_28806: SET 1,B
label_28807: SET 2,B
label_28808: SET 3,B
label_28809: SET 4,B
label_28810: SET 5,B
label_28811: SET 6,B
label_28812: SET 7,B
label_28813: SET 0,C
label_28814: SET 1,C
label_28815: SET 2,C
label_28816: SET 3,C
label_28817: SET 4,C
label_28818: SET 5,C
label_28819: SET 6,C
label_28820: SET 7,C
label_28821: SET 0,D
label_28822: SET 1,D
label_28823: SET 2,D
label_28824: SET 3,D
label_28825: SET 4,D
label_28826: SET 5,D
label_28827: SET 6,D
label_28828: SET 7,D
label_28829: SET 0,E
label_28830: SET 1,E
label_28831: SET 2,E
label_28832: SET 3,E
label_28833: SET 4,E
label_28834: SET 5,E
label_28835: SET 6,E
label_28836: SET 7,E
label_28837: SET 0,H
label_28838: SET 1,H
label_28839: SET 2,H
label_28840: SET 3,H
label_28841: SET 4,H
label_28842: SET 5,H
label_28843: SET 6,H
label_28844: SET 7,H
label_28845: SET 0,L
label_28846: SET 1,L
label_28847: SET 2,L
label_28848: SET 3,L
label_28849: SET 4,L
label_28850: SET 5,L
label_28851: SET 6,L
label_28852: SET 7,L
label_28853: SET 0,(HL)
label_28854: SET 1,(HL)
label_28855: SET 2,(HL)
label_28856: SET 3,(HL)
label_28857: SET 4,(HL)
label_28858: SET 5,(HL)
label_28859: SET 6,(HL)
label_28860: SET 7,(HL)
label_28861: SET 0,(IX + 127)
label_28862: SET 1,(IX + 127)
label_28863: SET 2,(IX + 127)
label_28864: SET 3,(IX + 127)
label_28865: SET 4,(IX + 127)
label_28866: SET 5,(IX + 127)
label_28867: SET 6,(IX + 127)
label_28868: SET 7,(IX + 127)
label_28869: SET 0,(IY - 128)
label_28870: SET 1,(IY - 128)
label_28871: SET 2,(IY - 128)
label_28872: SET 3,(IY - 128)
label_28873: SET 4,(IY - 128)
label_28874: SET 5,(IY - 128)
label_28875: SET 6,(IY - 128)
label_28876: SET 7,(IY - 128)
label_28877: RES 0,A
label_28878: RES 1,A
label_28879: RES 2,A
label_28880: RES 3,A
label_28881: RES 4,A
label_28882: RES 5,A
label_28883: RES 6,A
label_28884: RES 7,A
label_28885: RES 0,B
label_28886: RES 1,B
label_28887: RES 2,B
label_28888: RES 3,B
label_28889: RES 4,B
label_28890: RES 5,B
label_28891: RES 6,B
label_28892: RES 7,B
label_28893: RES 0,C
label_28894: RES 1,C
label_28895: RES 2,C
label_28896: RES 3,C
label_28897: RES 4,C
label_28898: RES 5,C
label_28899: RES 6,C
label_28900: RES 7,C
label_28901: RES 0,D
label_28902: RES 1,D
label_28903: RES 2,D
label_28904: RES 3,D
label_28905: RES 4,D
label_28906: RES 5,D
label_28907: RES 6,D
label_28908: RES 7,D
label_28909: RES 0,E
label_28910: RES 1,E
label_28911: RES 2,E
label_28912: RES 3,E
label_28913: RES 4,E
label_28914: RES 5,E
label_28915: RES 6,E
label_28916: RES 7,E
label_28917: RES 0,H
label_28918: RES 1,H
label_28919: RES 2,H
label_28920: RES 3,H
label_28921: RES 4,H
label_28922: RES 5,H
label_28923: RES 6,H
label_28924: RES 7,H
label_28925: RES 0,L
label_28926: RES 1,L
label_28927: RES 2,L
label_28928: RES 3,L
label_28929: RES 4,L
label_28930: RES 5,L
label_28931: RES 6,L
label_28932: RES 7,L
label_28933: RES 0,(HL)
label_28934: RES 1,(HL)
label_28935: RES 2,(HL)
label_28936: RES 3,(HL)
label_28937: RES 4,(HL)
label_28938: RES 5,(HL)
label_28939: RES 6,(HL)
label_28940: RES 7,(HL)
label_28941: RES 0,(IX + 127)
label_28942: RES 1,(IX + 127)
label_28943: RES 2,(IX + 127)
label_28944: RES 3,(IX + 127)
label_28945: RES 4,(IX + 127)
label_28946: RES 5,(IX + 127)
label_28947: RES 6,(IX + 127)
label_28948: RES 7,(IX + 127)
label_28949: RES 0,(IY - 128)
label_28950: RES 1,(IY - 128)
label_28951: RES 2,(IY - 128)
label_28952: RES 3,(IY - 128)
label_28953: RES 4,(IY - 128)
label_28954: RES 5,(IY - 128)
label_28955: RES 6,(IY - 128)
label_28956: RES 7,(IY - 128)
label_28957: JP $5678
label_28958: JP NZ,$5678
label_28959: JP Z,$5678
label_28960: JP NC,$5678
label_28961: JP C,$5678
label_28962: JP PO,$5678
label_28963: JP PE,$5678
label_28964: JP P,$5678
label_28965: JP M,$5678
label_28966: JR $ + 2
label_28967: JR NZ,$ + 2
label_28968: JR Z,$ + 2
label_28969: JR NC,$ + 2
label_28970: JR C,$ + 2
label_28971: JP (HL)
label_28972: JP (IX)
label_28973: JP (IY)
label_28974: DJNZ $ + 2
label_28975: CALL $5678
label_28976: CALL NZ,$5678
label_28977: CALL Z,$5678
label_28978: CALL NC,$5678
label_28979: CALL C,$5678
label_28980: CALL PO,$5678
label_28981: CALL PE,$5678
label_28982: CALL P,$5678
label_28983: CALL M,$5678
label_28984: RET
label_28985: RET NZ
label_28986: RET Z
label_28987: RET NC
label_28988: RET C
label_28989: RET PO
label_28990: RET PE
label_28991: RET P
label_28992: RET M
label_28993: RETI
label_28994: RETN
label_28995: RST $00
label_28996: RST $08
label_28997: RST $10
label_28998: RST $18
label_28999: RST $20
label_29000: RST $28
label_29001: RST $30
label_29002: RST $38
label_29003: IN A,($12)
label_29004: IN A,(C)
label_29005: IN B,(C)
label_29006: IN C,(C)
label_29007: IN D,(C)
label_29008: IN E,(C)
label_29009: IN H,(C)
label_29010: IN L,(C)
label_29011: IN F,(C)
label_29012: INI
label_29013: INIR
label_29014: IND
label_29015: INDR
label_29016: OUT ($12),A
label_29017: OUT (C),A
label_29018: OUT (C),B
label_29019: OUT (C),C
label_29020: OUT (C),D
label_29021: OUT (C),E
label_29022: OUT (C),H
label_29023: OUT (C),L
label_29024: OUTI
label_29025: OTIR
label_29026: OUTD
label_29027: OTDR
label_29028: LD A,A
label_29029: LD A,B
label_29030: LD A,C
label_29031: LD A,D
label_29032: LD A,E
label_29033: LD A,H
label_29034: LD A,L
label_29035: LD B,A
label_29036: LD B,B
label_29037: LD B,C
label_29038: LD B,D
label_29039: LD B,E
label_29040: LD B,H
label_29041: LD B,L
label_29042: LD C,A
label_29043: LD C,B
label_29044: LD C,C
label_29045: LD C,D
label_29046: LD C,E
label_29047: LD C,H
label_29048: LD C,L
label_29049: LD D,A
label_29050: LD D,B
label_29051: LD D,C
label_29052: LD D,D
label_29053: LD D,E
label_29054: LD D,H
label_29055: LD D,L
label_29056: LD E,A
label_29057: LD E,B
label_29058: LD E,C
label_29059: LD E,D
label_29060: LD E,E
label_29061: LD E,H
label_29062: LD E,L
label_29063: LD H,A
label_29064: LD H,B
label_29065: LD H,C
label_29066: LD H,D
label_29067: LD H,E
label_29068: LD H,H
label_29069: LD H,L
label_29070: LD L,A
label_29071: LD L,B
label_29072: LD L,C
label_29073: LD L,D
label_29074: LD L,E
label_29075: LD L,H
label_29076: LD L,L
label_29077: LD A,$12
label_29078: LD B,$12
label_29079: LD C,$12
label_29080: LD D,$12
label_29081: LD E,$12
label_29082: LD H,$12
label_29083: LD L,$12
label_29084: LD A,(HL)
label_29085: LD B,(HL)
label_29086: LD C,(HL)
label_29087: LD D,(HL)
label_29088: LD E,(HL)
label_29089: LD H,(HL)
label_29090: LD L,(HL)
label_29091: LD A,(IX + 127)
label_29092: LD B,(IX + 127)
label_29093: LD C,(IX + 127)
label_29094: LD D,(IX + 127)
label_29095: LD E,(IX + 127)
label_29096: LD H,(IX + 127)
label_29097: LD L,(IX + 127)
label_29098: LD A,(IY - 128)
label_29099: LD B,(IY - 128)
label_29100: LD C,(IY - 128)
label_29101: LD D,(IY - 128)
label_29102: LD E,(IY - 128)
label_29103: LD H,(IY - 128)
label_29104: LD L,(IY - 128)
label_29105: LD (HL),A
label_29106: LD (HL),B
label_29107: LD (HL),C
label_29108: LD (HL),D
label_29109: LD (HL),E
label_29110: LD (HL),H
label_29111: LD (HL),L
label_29112: LD (IX + 127),A
label_29113: LD (IX + 127),B
label_29114: LD (IX + 127),C
label_29115: LD (IX + 127),D
label_29116: LD (IX + 127),E
label_29117: LD (IX + 127),H
label_29118: LD (IX + 127),L
label_29119: LD (IY - 128),A
label_29120: LD (IY - 128),B
label_29121: LD (IY - 128),C
label_29122: LD (IY - 128),D
label_29123: LD (IY - 128),E
label_29124: LD (IY - 128),H
label_29125: LD (IY - 128),L
label_29126: LD (HL),$12
label_29127: LD (IX + 127),$12
label_29128: LD (IY - 128),$12
label_29129: LD A,(BC)
label_29130: LD A,(DE)
label_29131: LD A,($5678)
label_29132: LD (BC),A
label_29133: LD (DE),A
label_29134: LD ($5678),A
label_29135: LD A,I
label_29136: LD A,R
label_29137: LD I,A
label_29138: LD R,A
label_29139: LD BC,$5678
label_29140: LD DE,$5678
label_29141: LD HL,$5678
label_29142: LD SP,$5678
label_29143: LD IX,$5678
label_29144: LD IY,$5678
label_29145: LD HL,($5678)
label_29146: LD BC,($5678)
label_29147: LD DE,($5678)
label_29148: LD HL,($5678)
label_29149: LD SP,($5678)
label_29150: LD IX,($5678)
label_29151: LD IY,($5678)
label_29152: LD ($5678),HL
label_29153: LD ($5678),BC
label_29154: LD ($5678),DE
label_29155: LD ($5678),HL
label_29156: LD ($5678),SP
label_29157: LD ($5678),IX
label_29158: LD ($5678),IY
label_29159: LD SP,HL
label_29160: LD SP,IX
label_29161: LD SP,IY
label_29162: PUSH BC
label_29163: PUSH DE
label_29164: PUSH HL
label_29165: PUSH AF
label_29166: PUSH IX
label_29167: PUSH IY
label_29168: POP BC
label_29169: POP DE
label_29170: POP HL
label_29171: POP AF
label_29172: POP IX
label_29173: POP IY
label_29174: EX DE,HL
label_29175: EX AF,AF'
label_29176: EXX
label_29177: EX (SP),HL
label_29178: EX (SP),IX
label_29179: EX (SP),IY
label_29180: LDI
label_29181: LDIR
label_29182: LDD
label_29183: LDDR
label_29184: CPI
label_29185: CPIR
label_29186: CPD
label_29187: CPDR
label_29188: ADD A,A
label_29189: ADD A,B
label_29190: ADD A,C
label_29191: ADD A,D
label_29192: ADD A,E
label_29193: ADD A,H
label_29194: ADD A,L
label_29195: ADD A,$12
label_29196: ADD A,(HL)
label_29197: ADD A,(IX + 127)
label_29198: ADD A,(IY - 128)
label_29199: ADC A,A
label_29200: ADC A,B
label_29201: ADC A,C
label_29202: ADC A,D
label_29203: ADC A,E
label_29204: ADC A,H
label_29205: ADC A,L
label_29206: ADC A,$12
label_29207: ADC A,(HL)
label_29208: ADC A,(IX + 127)
label_29209: ADC A,(IY - 128)
label_29210: SUB A
label_29211: SUB B
label_29212: SUB C
label_29213: SUB D
label_29214: SUB E
label_29215: SUB H
label_29216: SUB L
label_29217: SUB $12
label_29218: SUB (HL)
label_29219: SUB (IX + 127)
label_29220: SUB (IY - 128)
label_29221: SBC A,A
label_29222: SBC A,B
label_29223: SBC A,C
label_29224: SBC A,D
label_29225: SBC A,E
label_29226: SBC A,H
label_29227: SBC A,L
label_29228: SBC A,$12
label_29229: SBC A,(HL)
label_29230: SBC A,(IX + 127)
label_29231: SBC A,(IY - 128)
label_29232: AND A
label_29233: AND B
label_29234: AND C
label_29235: AND D
label_29236: AND E
label_29237: AND H
label_29238: AND L
label_29239: AND $12
label_29240: AND (HL)
label_29241: AND (IX + 127)
label_29242: AND (IY - 128)
label_29243: AND A
label_29244: AND B
label_29245: AND C
label_29246: AND D
label_29247: AND E
label_29248: AND H
label_29249: AND L
label_29250: AND $12
label_29251: AND (HL)
label_29252: AND (IX + 127)
label_29253: AND (IY - 128)
label_29254: OR A
label_29255: OR B
label_29256: OR C
label_29257: OR D
label_29258: OR E
label_29259: OR H
label_29260: OR L
label_29261: OR $12
label_29262: OR (HL)
label_29263: OR (IX + 127)
label_29264: OR (IY - 128)
label_29265: XOR A
label_29266: XOR B
label_29267: XOR C
label_29268: XOR D
label_29269: XOR E
label_29270: XOR H
label_29271: XOR L
label_29272: XOR $12
label_29273: XOR (HL)
label_29274: XOR (IX + 127)
label_29275: XOR (IY - 128)
label_29276: CP A
label_29277: CP B
label_29278: CP C
label_29279: CP D
label_29280: CP E
label_29281: CP H
label_29282: CP L
label_29283: CP $12
label_29284: CP (HL)
label_29285: CP (IX + 127)
label_29286: CP (IY - 128)
label_29287: INC A
label_29288: INC B
label_29289: INC C
label_29290: INC D
label_29291: INC E
label_29292: INC H
label_29293: INC L
label_29294: INC (HL)
label_29295: INC (IX + 127)
label_29296: INC (IY - 128)
label_29297: DEC A
label_29298: DEC B
label_29299: DEC C
label_29300: DEC D
label_29301: DEC E
label_29302: DEC H
label_29303: DEC L
label_29304: DEC (HL)
label_29305: DEC (IX + 127)
label_29306: DEC (IY - 128)
label_29307: DAA
label_29308: CPL
label_29309: NEG
label_29310: CCF
label_29311: SCF
label_29312: NOP
label_29313: HALT
label_29314: DI
label_29315: EI
label_29316: IM 0
label_29317: IM 1
label_29318: IM 2
label_29319: ADD HL,BC
label_29320: ADD HL,DE
label_29321: ADD HL,HL
label_29322: ADD HL,SP
label_29323: ADC HL,BC
label_29324: ADC HL,DE
label_29325: ADC HL,HL
label_29326: ADC HL,SP
label_29327: SBC HL,BC
label_29328: SBC HL,DE
label_29329: SBC HL,HL
label_29330: SBC HL,SP
label_29331: ADD IX,BC
label_29332: ADD IX,DE
label_29333: ADD IX,SP
label_29334: ADD IY,BC
label_29335: ADD IY,DE
label_29336: ADD IY,SP
label_29337: INC BC
label_29338: INC DE
label_29339: INC HL
label_29340: INC SP
label_29341: INC IX
label_29342: INC IY
label_29343: DEC BC
label_29344: DEC DE
label_29345: DEC HL
label_29346: DEC SP
label_29347: DEC IX
label_29348: DEC IY
label_29349: RLCA
label_29350: RLA
label_29351: RRCA
label_29352: RRA
label_29353: RLC A
label_29354: RLC B
label_29355: RLC C
label_29356: RLC D
label_29357: RLC E
label_29358: RLC H
label_29359: RLC L
label_29360: RLC (HL)
label_29361: RLC (IX + 127)
label_29362: RLC (IY - 128)
label_29363: RL A
label_29364: RL B
label_29365: RL C
label_29366: RL D
label_29367: RL E
label_29368: RL H
label_29369: RL L
label_29370: RL (HL)
label_29371: RL (IX + 127)
label_29372: RL (IY - 128)
label_29373: RRC A
label_29374: RRC B
label_29375: RRC C
label_29376: RRC D
label_29377: RRC E
label_29378: RRC H
label_29379: RRC L
label_29380: RRC (HL)
label_29381: RRC (IX + 127)
label_29382: RRC (IY - 128)
label_29383: RR A
label_29384: RR B
label_29385: RR C
label_29386: RR D
label_29387: RR E
label_29388: RR H
label_29389: RR L
label_29390: RR (HL)
label_29391: RR (IX + 127)
label_29392: RR (IY - 128)
label_29393: SLA A
label_29394: SLA B
label_29395: SLA C
label_29396: SLA D
label_29397: SLA E
label_29398: SLA H
label_29399: SLA L
label_29400: SLA (HL)
label_29401: SLA (IX + 127)
label_29402: SLA (IY - 128)
label_29403: SRA A
label_29404: SRA B
label_29405: SRA C
label_29406: SRA D
label_29407: SRA E
label_29408: SRA H
label_29409: SRA L
label_29410: SRA (HL)
label_29411: SRA (IX + 127)
label_29412: SRA (IY - 128)
label_29413: SRL A
label_29414: SRL B
label_29415: SRL C
label_29416: SRL D
label_29417: SRL E
label_29418: SRL H
label_29419: SRL L
label_29420: SRL (HL)
label_29421: SRL (IX + 127)
label_29422: SRL (IY - 128)
label_29423: RLD
label_29424: RRD
label_29425: BIT 0,A
label_29426: BIT 1,A
label_29427: BIT 2,A
label_29428: BIT 3,A
label_29429: BIT 4,A
label_29430: BIT 5,A
label_29431: BIT 6,A
label_29432: BIT 7,A
label_29433: BIT 0,B
label_29434: BIT 1,B
label_29435: BIT 2,B
label_29436: BIT 3,B
label_29437: BIT 4,B
label_29438: BIT 5,B
label_29439: BIT 6,B
label_29440: BIT 7,B
label_29441: BIT 0,C
label_29442: BIT 1,C
label_29443: BIT 2,C
label_29444: BIT 3,C
label_29445: BIT 4,C
label_29446: BIT 5,C
label_29447: BIT 6,C
label_29448: BIT 7,C
label_29449: BIT 0,D
label_29450: BIT 1,D
label_29451: BIT 2,D
label_29452: BIT 3,D
label_29453: BIT 4,D
label_29454: BIT 5,D
label_29455: BIT 6,D
label_29456: BIT 7,D
label_29457: BIT 0,E
label_29458: BIT 1,E
label_29459: BIT 2,E
label_29460: BIT 3,E
label_29461: BIT 4,E
label_29462: BIT 5,E
label_29463: BIT 6,E
label_29464: BIT 7,E
label_29465: BIT 0,H
label_29466: BIT 1,H
label_29467: BIT 2,H
label_29468: BIT 3,H
label_29469: BIT 4,H
label_29470: BIT 5,H
label_29471: BIT 6,H
label_29472: BIT 7,H
label_29473: BIT 0,L
label_29474: BIT 1,L
label_29475: BIT 2,L
label_29476: BIT 3,L
label_29477: BIT 4,L
label_29478: BIT 5,L
label_29479: BIT 6,L
label_29480: BIT 7,L
label_29481: BIT 0,(HL)
label_29482: BIT 1,(HL)
label_29483: BIT 2,(HL)
label_29484: BIT 3,(HL)
label_29485: BIT 4,(HL)
label_29486: BIT 5,(HL)
label_29487: BIT 6,(HL)
label_29488: BIT 7,(HL)
label_29489: BIT 0,(IX + 127)
label_29490: BIT 1,(IX + 127)
label_29491: BIT 2,(IX + 127)
label_29492: BIT 3,(IX + 127)
label_29493: BIT 4,(IX + 127)
label_29494: BIT 5,(IX + 127)
label_29495: BIT 6,(IX + 127)
label_29496: BIT 7,(IX + 127)
label_29497: BIT 0,(IY - 128)
label_29498: BIT 1,(IY - 128)
label_29499: BIT 2,(IY - 128)
label_29500: BIT 3,(IY - 128)
label_29501: BIT 4,(IY - 128)
label_29502: BIT 5,(IY - 128)
label_29503: BIT 6,(IY - 128)
label_29504: BIT 7,(IY - 128)
label_29505: SET 0,A
label_29506: SET 1,A
label_29507: SET 2,A
label_29508: SET 3,A
label_29509: SET 4,A
label_29510: SET 5,A
label_29511: SET 6,A
label_29512: SET 7,A
label_29513: SET 0,B
label_29514: SET 1,B
label_29515: SET 2,B
label_29516: SET 3,B
label_29517: SET 4,B
label_29518: SET 5,B
label_29519: SET 6,B
label_29520: SET 7,B
label_29521: SET 0,C
label_29522: SET 1,C
label_29523: SET 2,C
label_29524: SET 3,C
label_29525: SET 4,C
label_29526: SET 5,C
label_29527: SET 6,C
label_29528: SET 7,C
label_29529: SET 0,D
label_29530: SET 1,D
label_29531: SET 2,D
label_29532: SET 3,D
label_29533: SET 4,D
label_29534: SET 5,D
label_29535: SET 6,D
label_29536: SET 7,D
label_29537: SET 0,E
label_29538: SET 1,E
label_29539: SET 2,E
label_29540: SET 3,E
label_29541: SET 4,E
label_29542: SET 5,E
label_29543: SET 6,E
label_29544: SET 7,E
label_29545: SET 0,H
label_29546: SET 1,H
label_29547: SET 2,H
label_29548: SET 3,H
label_29549: SET 4,H
label_29550: SET 5,H
label_29551: SET 6,H
label_29552: SET 7,H
label_29553: SET 0,L
label_29554: SET 1,L
label_29555: SET 2,L
label_29556: SET 3,L
label_29557: SET 4,L
label_29558: SET 5,L
label_29559: SET 6,L
label_29560: SET 7,L
label_29561: SET 0,(HL)
label_29562: SET 1,(HL)
label_29563: SET 2,(HL)
label_29564: SET 3,(HL)
label_29565: SET 4,(HL)
label_29566: SET 5,(HL)
label_29567: SET 6,(HL)
label_29568: SET 7,(HL)
label_29569: SET 0,(IX + 127)
label_29570: SET 1,(IX + 127)
label_29571: SET 2,(IX + 127)
label_29572: SET 3,(IX + 127)
label_29573: SET 4,(IX + 127)
label_29574: SET 5,(IX + 127)
label_29575: SET 6,(IX + 127)
label_29576: SET 7,(IX + 127)
label_29577: SET 0,(IY - 128)
label_29578: SET 1,(IY - 128)
label_29579: SET 2,(IY - 128)
label_29580: SET 3,(IY - 128)
label_29581: SET 4,(IY - 128)
label_29582: SET 5,(IY - 128)
label_29583: SET 6,(IY - 128)
label_29584: SET 7,(IY - 128)
label_29585: RES 0,A
label_29586: RES 1,A
label_29587: RES 2,A
label_29588: RES 3,A
label_29589: RES 4,A
label_29590: RES 5,A
label_29591: RES 6,A
label_29592: RES 7,A
label_29593: RES 0,B
label_29594: RES 1,B
label_29595: RES 2,B
label_29596: RES 3,B
label_29597: RES 4,B
label_29598: RES 5,B
label_29599: RES 6,B
label_29600: RES 7,B
label_29601: RES 0,C
label_29602: RES 1,C
label_29603: RES 2,C
label_29604: RES 3,C
label_29605: RES 4,C
label_29606: RES 5,C
label_29607: RES 6,C
label_29608: RES 7,C
label_29609: RES 0,D
label_29610: RES 1,D
label_29611: RES 2,D
label_29612: RES 3,D
label_29613: RES 4,D
label_29614: RES 5,D
label_29615: RES 6,D
label_29616: RES 7,D
label_29617: RES 0,E
label_29618: RES 1,E
label_29619: RES 2,E
label_29620: RES 3,E
label_29621: RES 4,E
label_29622: RES 5,E
label_29623: RES 6,E
label_29624: RES 7,E
label_29625: RES 0,H
label_29626: RES 1,H
label_29627: RES 2,H
label_29628: RES 3,H
label_29629: RES 4,H
label_29630: RES 5,H
label_29631: RES 6,H
label_29632: RES 7,H
label_29633: RES 0,L
label_29634: RES 1,L
label_29635: RES 2,L
label_29636: RES 3,L
label_29637: RES 4,L
label_29638: RES 5,L
label_29639: RES 6,L
label_29640: RES 7,L
label_29641: RES 0,(HL)
label_29642: RES 1,(HL)
label_29643: RES 2,(HL)
label_29644: RES 3,(HL)
label_29645: RES 4,(HL)
label_29646: RES 5,(HL)
label_29647: RES 6,(HL)
label_29648: RES 7,(HL)
label_29649: RES 0,(IX + 127)
label_29650: RES 1,(IX + 127)
label_29651: RES 2,(IX + 127)
label_29652: RES 3,(IX + 127)
label_29653: RES 4,(IX + 127)
label_29654: RES 5,(IX + 127)
label_29655: RES 6,(IX + 127)
label_29656: RES 7,(IX + 127)
label_29657: RES 0,(IY - 128)
label_29658: RES 1,(IY - 128)
label_29659: RES 2,(IY - 128)
label_29660: RES 3,(IY - 128)
label_29661: RES 4,(IY - 128)
label_29662: RES 5,(IY - 128)
label_29663: RES 6,(IY - 128)
label_29664: RES 7,(IY - 128)
label_29665: JP $5678
label_29666: JP NZ,$5678
label_29667: JP Z,$5678
label_29668: JP NC,$5678
label_29669: JP C,$5678
label_29670: JP PO,$5678
label_29671: JP PE,$5678
label_29672: JP P,$5678
label_29673: JP M,$5678
label_29674: JR $ + 2
label_29675: JR NZ,$ + 2
label_29676: JR Z,$ + 2
label_29677: JR NC,$ + 2
label_29678: JR C,$ + 2
label_29679: JP (HL)
label_29680: JP (IX)
label_29681: JP (IY)
label_29682: DJNZ $ + 2
label_29683: CALL $5678
label_29684: CALL NZ,$5678
label_29685: CALL Z,$5678
label_29686: CALL NC,$5678
label_29687: CALL C,$5678
label_29688: CALL PO,$5678
label_29689: CALL PE,$5678
label_29690: CALL P,$5678
label_29691: CALL M,$5678
label_29692: RET
label_29693: RET NZ
label_29694: RET Z
label_29695: RET NC
label_29696: RET C
label_29697: RET PO
label_29698: RET PE
label_29699: RET P
label_29700: RET M
label_29701: RETI
label_29702: RETN
label_29703: RST $00
label_29704: RST $08
label_29705: RST $10
label_29706: RST $18
label_29707: RST $20
label_29708: RST $28
label_29709: RST $30
label_29710: RST $38
label_29711: IN A,($12)
label_29712: IN A,(C)
label_29713: IN B,(C)
label_29714: IN C,(C)
label_29715: IN D,(C)
label_29716: IN E,(C)
label_29717: IN H,(C)
label_29718: IN L,(C)
label_29719: IN F,(C)
label_29720: INI
label_29721: INIR
label_29722: IND
label_29723: INDR
label_29724: OUT ($12),A
label_29725: OUT (C),A
label_29726: OUT (C),B
label_29727: OUT (C),C
label_29728: OUT (C),D
label_29729: OUT (C),E
label_29730: OUT (C),H
label_29731: OUT (C),L
label_29732: OUTI
label_29733: OTIR
label_29734: OUTD
label_29735: OTDR
label_29736: LD A,A
label_29737: LD A,B
label_29738: LD A,C
label_29739: LD A,D
label_29740: LD A,E
label_29741: LD A,H
label_29742: LD A,L
label_29743: LD B,A
label_29744: LD B,B
label_29745: LD B,C
label_29746: LD B,D
label_29747: LD B,E
label_29748: LD B,H
label_29749: LD B,L
label_29750: LD C,A
label_29751: LD C,B
label_29752: LD C,C
label_29753: LD C,D
label_29754: LD C,E
label_29755: LD C,H
label_29756: LD C,L
label_29757: LD D,A
label_29758: LD D,B
label_29759: LD D,C
label_29760: LD D,D
label_29761: LD D,E
label_29762: LD D,H
label_29763: LD D,L
label_29764: LD E,A
label_29765: LD E,B
label_29766: LD E,C
label_29767: LD E,D
label_29768: LD E,E
label_29769: LD E,H
label_29770: LD E,L
label_29771: LD H,A
label_29772: LD H,B
label_29773: LD H,C
label_29774: LD H,D
label_29775: LD H,E
label_29776: LD H,H
label_29777: LD H,L
label_29778: LD L,A
label_29779: LD L,B
label_29780: LD L,C
label_29781: LD L,D
label_29782: LD L,E
label_29783: LD L,H
label_29784: LD L,L
label_29785: LD A,$12
label_29786: LD B,$12
label_29787: LD C,$12
label_29788: LD D,$12
label_29789: LD E,$12
label_29790: LD H,$12
label_29791: LD L,$12
label_29792: LD A,(HL)
label_29793: LD B,(HL)
label_29794: LD C,(HL)
label_29795: LD D,(HL)
label_29796: LD E,(HL)
label_29797: LD H,(HL)
label_29798: LD L,(HL)
label_29799: LD A,(IX + 127)
label_29800: LD B,(IX + 127)
label_29801: LD C,(IX + 127)
label_29802: LD D,(IX + 127)
label_29803: LD E,(IX + 127)
label_29804: LD H,(IX + 127)
label_29805: LD L,(IX + 127)
label_29806: LD A,(IY - 128)
label_29807: LD B,(IY - 128)
label_29808: LD C,(IY - 128)
label_29809: LD D,(IY - 128)
label_29810: LD E,(IY - 128)
label_29811: LD H,(IY - 128)
label_29812: LD L,(IY - 128)
label_29813: LD (HL),A
label_29814: LD (HL),B
label_29815: LD (HL),C
label_29816: LD (HL),D
label_29817: LD (HL),E
label_29818: LD (HL),H
label_29819: LD (HL),L
label_29820: LD (IX + 127),A
label_29821: LD (IX + 127),B
label_29822: LD (IX + 127),C
label_29823: LD (IX + 127),D
label_29824: LD (IX + 127),E
label_29825: LD (IX + 127),H
label_29826: LD (IX + 127),L
label_29827: LD (IY - 128),A
label_29828: LD (IY - 128),B
label_29829: LD (IY - 128),C
label_29830: LD (IY - 128),D
label_29831: LD (IY - 128),E
label_29832: LD (IY - 128),H
label_29833: LD (IY - 128),L
label_29834: LD (HL),$12
label_29835: LD (IX + 127),$12
label_29836: LD (IY - 128),$12
label_29837: LD A,(BC)
label_29838: LD A,(DE)
label_29839: LD A,($5678)
label_29840: LD (BC),A
label_29841: LD (DE),A
label_29842: LD ($5678),A
label_29843: LD A,I
label_29844: LD A,R
label_29845: LD I,A
label_29846: LD R,A
label_29847: LD BC,$5678
label_29848: LD DE,$5678
label_29849: LD HL,$5678
label_29850: LD SP,$5678
label_29851: LD IX,$5678
label_29852: LD IY,$5678
label_29853: LD HL,($5678)
label_29854: LD BC,($5678)
label_29855: LD DE,($5678)
label_29856: LD HL,($5678)
label_29857: LD SP,($5678)
label_29858: LD IX,($5678)
label_29859: LD IY,($5678)
label_29860: LD ($5678),HL
label_29861: LD ($5678),BC
label_29862: LD ($5678),DE
label_29863: LD ($5678),HL
label_29864: LD ($5678),SP
label_29865: LD ($5678),IX
label_29866: LD ($5678),IY
label_29867: LD SP,HL
label_29868: LD SP,IX
label_29869: LD SP,IY
label_29870: PUSH BC
label_29871: PUSH DE
label_29872: PUSH HL
label_29873: PUSH AF
label_29874: PUSH IX
label_29875: PUSH IY
label_29876: POP BC
label_29877: POP DE
label_29878: POP HL
label_29879: POP AF
label_29880: POP IX
label_29881: POP IY
label_29882: EX DE,HL
label_29883: EX AF,AF'
label_29884: EXX
label_29885: EX (SP),HL
label_29886: EX (SP),IX
label_29887: EX (SP),IY
label_29888: LDI
label_29889: LDIR
label_29890: LDD
label_29891: LDDR
label_29892: CPI
label_29893: CPIR
label_29894: CPD
label_29895: CPDR
label_29896: ADD A,A
label_29897: ADD A,B
label_29898: ADD A,C
label_29899: ADD A,D
label_29900: ADD A,E
label_29901: ADD A,H
label_29902: ADD A,L
label_29903: ADD A,$12
label_29904: ADD A,(HL)
label_29905: ADD A,(IX + 127)
label_29906: ADD A,(IY - 128)
label_29907: ADC A,A
label_29908: ADC A,B
label_29909: ADC A,C
label_29910: ADC A,D
label_29911: ADC A,E
label_29912: ADC A,H
label_29913: ADC A,L
label_29914: ADC A,$12
label_29915: ADC A,(HL)
label_29916: ADC A,(IX + 127)
label_29917: ADC A,(IY - 128)
label_29918: SUB A
label_29919: SUB B
label_29920: SUB C
label_29921: SUB D
label_29922: SUB E
label_29923: SUB H
label_29924: SUB L
label_29925: SUB $12
label_29926: SUB (HL)
label_29927: SUB (IX + 127)
label_29928: SUB (IY - 128)
label_29929: SBC A,A
label_29930: SBC A,B
label_29931: SBC A,C
label_29932: SBC A,D
label_29933: SBC A,E
label_29934: SBC A,H
label_29935: SBC A,L
label_29936: SBC A,$12
label_29937: SBC A,(HL)
label_29938: SBC A,(IX + 127)
label_29939: SBC A,(IY - 128)
label_29940: AND A
label_29941: AND B
label_29942: AND C
label_29943: AND D
label_29944: AND E
label_29945: AND H
label_29946: AND L
label_29947: AND $12
label_29948: AND (HL)
label_29949: AND (IX + 127)
label_29950: AND (IY - 128)
label_29951: AND A
label_29952: AND B
label_29953: AND C
label_29954: AND D
label_29955: AND E
label_29956: AND H
label_29957: AND L
label_29958: AND $12
label_29959: AND (HL)
label_29960: AND (IX + 127)
label_29961: AND (IY - 128)
label_29962: OR A
label_29963: OR B
label_29964: OR C
label_29965: OR D
label_29966: OR E
label_29967: OR H
label_29968: OR L
label_29969: OR $12
label_29970: OR (HL)
label_29971: OR (IX + 127)
label_29972: OR (IY - 128)
label_29973: XOR A
label_29974: XOR B
label_29975: XOR C
label_29976: XOR D
label_29977: XOR E
label_29978: XOR H
label_29979: XOR L
label_29980: XOR $12
label_29981: XOR (HL)
label_29982: XOR (IX + 127)
label_29983: XOR (IY - 128)
label_29984: CP A
label_29985: CP B
label_29986: CP C
label_29987: CP D
label_29988: CP E
label_29989: CP H
label_29990: CP L
label_29991: CP $12
label_29992: CP (HL)
label_29993: CP (IX + 127)
label_29994: CP (IY - 128)
label_29995: INC A
label_29996: INC B
label_29997: INC C
label_29998: INC D
label_29999: INC E
label_30000: INC H
label_30001: INC L
label_30002: INC (HL)
label_30003: INC (IX + 127)
label_30004: INC (IY - 128)
label_30005: DEC A
label_30006: DEC B
label_30007: DEC C
label_30008: DEC D
label_30009: DEC E
label_30010: DEC H
label_30011: DEC L
label_30012: DEC (HL)
label_30013: DEC (IX + 127)
label_30014: DEC (IY - 128)
label_30015: DAA
label_30016: CPL
label_30017: NEG
label_30018: CCF
label_30019: SCF
label_30020: NOP
label_30021: HALT
label_30022: DI
label_30023: EI
label_30024: IM 0
label_30025: IM 1
label_30026: IM 2
label_30027: ADD HL,BC
label_30028: ADD HL,DE
label_30029: ADD HL,HL
label_30030: ADD HL,SP
label_30031: ADC HL,BC
label_30032: ADC HL,DE
label_30033: ADC HL,HL
label_30034: ADC HL,SP
label_30035: SBC HL,BC
label_30036: SBC HL,DE
label_30037: SBC HL,HL
label_30038: SBC HL,SP
label_30039: ADD IX,BC
label_30040: ADD IX,DE
label_30041: ADD IX,SP
label_30042: ADD IY,BC
label_30043: ADD IY,DE
label_30044: ADD IY,SP
label_30045: INC BC
label_30046: INC DE
label_30047: INC HL
label_30048: INC SP
label_30049: INC IX
label_30050: INC IY
label_30051: DEC BC
label_30052: DEC DE
label_30053: DEC HL
label_30054: DEC SP
label_30055: DEC IX
label_30056: DEC IY
label_30057: RLCA
label_30058: RLA
label_30059: RRCA
label_30060: RRA
label_30061: RLC A
label_30062: RLC B
label_30063: RLC C
label_30064: RLC D
label_30065: RLC E
label_30066: RLC H
label_30067: RLC L
label_30068: RLC (HL)
label_30069: RLC (IX + 127)
label_30070: RLC (IY - 128)
label_30071: RL A
label_30072: RL B
label_30073: RL C
label_30074: RL D
label_30075: RL E
label_30076: RL H
label_30077: RL L
label_30078: RL (HL)
label_30079: RL (IX + 127)
label_30080: RL (IY - 128)
label_30081: RRC A
label_30082: RRC B
label_30083: RRC C
label_30084: RRC D
label_30085: RRC E
label_30086: RRC H
label_30087: RRC L
label_30088: RRC (HL)
label_30089: RRC (IX + 127)
label_30090: RRC (IY - 128)
label_30091: RR A
label_30092: RR B
label_30093: RR C
label_30094: RR D
label_30095: RR E
label_30096: RR H
label_30097: RR L
label_30098: RR (HL)
label_30099: RR (IX + 127)
label_30100: RR (IY - 128)
label_30101: SLA A
label_30102: SLA B
label_30103: SLA C
label_30104: SLA D
label_30105: SLA E
label_30106: SLA H
label_30107: SLA L
label_30108: SLA (HL)
label_30109: SLA (IX + 127)
label_30110: SLA (IY - 128)
label_30111: SRA A
label_30112: SRA B
label_30113: SRA C
label_30114: SRA D
label_30115: SRA E
label_30116: SRA H
label_30117: SRA L
label_30118: SRA (HL)
label_30119: SRA (IX + 127)
label_30120: SRA (IY - 128)
label_30121: SRL A
label_30122: SRL B
label_30123: SRL C
label_30124: SRL D
label_30125: SRL E
label_30126: SRL H
label_30127: SRL L
label_30128: SRL (HL)
label_30129: SRL (IX + 127)
label_30130: SRL (IY - 128)
label_30131: RLD
label_30132: RRD
label_30133: BIT 0,A
label_30134: BIT 1,A
label_30135: BIT 2,A
label_30136: BIT 3,A
label_30137: BIT 4,A
label_30138: BIT 5,A
label_30139: BIT 6,A
label_30140: BIT 7,A
label_30141: BIT 0,B
label_30142: BIT 1,B
label_30143: BIT 2,B
label_30144: BIT 3,B
label_30145: BIT 4,B
label_30146: BIT 5,B
label_30147: BIT 6,B
label_30148: BIT 7,B
label_30149: BIT 0,C
label_30150: BIT 1,C
label_30151: BIT 2,C
label_30152: BIT 3,C
label_30153: BIT 4,C
label_30154: BIT 5,C
label_30155: BIT 6,C
label_30156: BIT 7,C
label_30157: BIT 0,D
label_30158: BIT 1,D
label_30159: BIT 2,D
label_30160: BIT 3,D
label_30161: BIT 4,D
label_30162: BIT 5,D
label_30163: BIT 6,D
label_30164: BIT 7,D
label_30165: BIT 0,E
label_30166: BIT 1,E
label_30167: BIT 2,E
label_30168: BIT 3,E
label_30169: BIT 4,E
label_30170: BIT 5,E
label_30171: BIT 6,E
label_30172: BIT 7,E
label_30173: BIT 0,H
label_30174: BIT 1,H
label_30175: BIT 2,H
label_30176: BIT 3,H
label_30177: BIT 4,H
label_30178: BIT 5,H
label_30179: BIT 6,H
label_30180: BIT 7,H
label_30181: BIT 0,L
label_30182: BIT 1,L
label_30183: BIT 2,L
label_30184: BIT 3,L
label_30185: BIT 4,L
label_30186: BIT 5,L
label_30187: BIT 6,L
label_30188: BIT 7,L
label_30189: BIT 0,(HL)
label_30190: BIT 1,(HL)
label_30191: BIT 2,(HL)
label_30192: BIT 3,(HL)
label_30193: BIT 4,(HL)
label_30194: BIT 5,(HL)
label_30195: BIT 6,(HL)
label_30196: BIT 7,(HL)
label_30197: BIT 0,(IX + 127)
label_30198: BIT 1,(IX + 127)
label_30199: BIT 2,(IX + 127)
label_30200: BIT 3,(IX + 127)
label_30201: BIT 4,(IX + 127)
label_30202: BIT 5,(IX + 127)
label_30203: BIT 6,(IX + 127)
label_30204: BIT 7,(IX + 127)
label_30205: BIT 0,(IY - 128)
label_30206: BIT 1,(IY - 128)
label_30207: BIT 2,(IY - 128)
label_30208: BIT 3,(IY - 128)
label_30209: BIT 4,(IY - 128)
label_30210: BIT 5,(IY - 128)
label_30211: BIT 6,(IY - 128)
label_30212: BIT 7,(IY - 128)
label_30213: SET 0,A
label_30214: SET 1,A
label_30215: SET 2,A
label_30216: SET 3,A
label_30217: SET 4,A
label_30218: SET 5,A
label_30219: SET 6,A
label_30220: SET 7,A
label_30221: SET 0,B
label_30222: SET 1,B
label_30223: SET 2,B
label_30224: SET 3,B
label_30225: SET 4,B
label_30226: SET 5,B
label_30227: SET 6,B
label_30228: SET 7,B
label_30229: SET 0,C
label_30230: SET 1,C
label_30231: SET 2,C
label_30232: SET 3,C
label_30233: SET 4,C
label_30234: SET 5,C
label_30235: SET 6,C
label_30236: SET 7,C
label_30237: SET 0,D
label_30238: SET 1,D
label_30239: SET 2,D
label_30240: SET 3,D
label_30241: SET 4,D
label_30242: SET 5,D
label_30243: SET 6,D
label_30244: SET 7,D
label_30245: SET 0,E
label_30246: SET 1,E
label_30247: SET 2,E
label_30248: SET 3,E
label_30249: SET 4,E
label_30250: SET 5,E
label_30251: SET 6,E
label_30252: SET 7,E
label_30253: SET 0,H
label_30254: SET 1,H
label_30255: SET 2,H
label_30256: SET 3,H
label_30257: SET 4,H
label_30258: SET 5,H
label_30259: SET 6,H
label_30260: SET 7,H
label_30261: SET 0,L
label_30262: SET 1,L
label_30263: SET 2,L
label_30264: SET 3,L
label_30265: SET 4,L
label_30266: SET 5,L
label_30267: SET 6,L
label_30268: SET 7,L
label_30269: SET 0,(HL)
label_30270: SET 1,(HL)
label_30271: SET 2,(HL)
label_30272: SET 3,(HL)
label_30273: SET 4,(HL)
label_30274: SET 5,(HL)
label_30275: SET 6,(HL)
label_30276: SET 7,(HL)
label_30277: SET 0,(IX + 127)
label_30278: SET 1,(IX + 127)
label_30279: SET 2,(IX + 127)
label_30280: SET 3,(IX + 127)
label_30281: SET 4,(IX + 127)
label_30282: SET 5,(IX + 127)
label_30283: SET 6,(IX + 127)
label_30284: SET 7,(IX + 127)
label_30285: SET 0,(IY - 128)
label_30286: SET 1,(IY - 128)
label_30287: SET 2,(IY - 128)
label_30288: SET 3,(IY - 128)
label_30289: SET 4,(IY - 128)
label_30290: SET 5,(IY - 128)
label_30291: SET 6,(IY - 128)
label_30292: SET 7,(IY - 128)
label_30293: RES 0,A
label_30294: RES 1,A
label_30295: RES 2,A
label_30296: RES 3,A
label_30297: RES 4,A
label_30298: RES 5,A
label_30299: RES 6,A
label_30300: RES 7,A
label_30301: RES 0,B
label_30302: RES 1,B
label_30303: RES 2,B
label_30304: RES 3,B
label_30305: RES 4,B
label_30306: RES 5,B
label_30307: RES 6,B
label_30308: RES 7,B
label_30309: RES 0,C
label_30310: RES 1,C
label_30311: RES 2,C
label_30312: RES 3,C
label_30313: RES 4,C
label_30314: RES 5,C
label_30315: RES 6,C
label_30316: RES 7,C
label_30317: RES 0,D
label_30318: RES 1,D
label_30319: RES 2,D
label_30320: RES 3,D
label_30321: RES 4,D
label_30322: RES 5,D
label_30323: RES 6,D
label_30324: RES 7,D
label_30325: RES 0,E
label_30326: RES 1,E
label_30327: RES 2,E
label_30328: RES 3,E
label_30329: RES 4,E
label_30330: RES 5,E
label_30331: RES 6,E
label_30332: RES 7,E
label_30333: RES 0,H
label_30334: RES 1,H
label_30335: RES 2,H
label_30336: RES 3,H
label_30337: RES 4,H
label_30338: RES 5,H
label_30339: RES 6,H
label_30340: RES 7,H
label_30341: RES 0,L
label_30342: RES 1,L
label_30343: RES 2,L
label_30344: RES 3,L
label_30345: RES 4,L
label_30346: RES 5,L
label_30347: RES 6,L
label_30348: RES 7,L
label_30349: RES 0,(HL)
label_30350: RES 1,(HL)
label_30351: RES 2,(HL)
label_30352: RES 3,(HL)
label_30353: RES 4,(HL)
label_30354: RES 5,(HL)
label_30355: RES 6,(HL)
label_30356: RES 7,(HL)
label_30357: RES 0,(IX + 127)
label_30358: RES 1,(IX + 127)
label_30359: RES 2,(IX + 127)
label_30360: RES 3,(IX + 127)
label_30361: RES 4,(IX + 127)
label_30362: RES 5,(IX + 127)
label_30363: RES 6,(IX + 127)
label_30364: RES 7,(IX + 127)
label_30365: RES 0,(IY - 128)
label_30366: RES 1,(IY - 128)
label_30367: RES 2,(IY - 128)
label_30368: RES 3,(IY - 128)
label_30369: RES 4,(IY - 128)
label_30370: RES 5,(IY - 128)
label_30371: RES 6,(IY - 128)
label_30372: RES 7,(IY - 128)
label_30373: JP $5678
label_30374: JP NZ,$5678
label_30375: JP Z,$5678
label_30376: JP NC,$5678
label_30377: JP C,$5678
label_30378: JP PO,$5678
label_30379: JP PE,$5678
label_30380: JP P,$5678
label_30381: JP M,$5678
label_30382: JR $ + 2
label_30383: JR NZ,$ + 2
label_30384: JR Z,$ + 2
label_30385: JR NC,$ + 2
label_30386: JR C,$ + 2
label_30387: JP (HL)
label_30388: JP (IX)
label_30389: JP (IY)
label_30390: DJNZ $ + 2
label_30391: CALL $5678
label_30392: CALL NZ,$5678
label_30393: CALL Z,$5678
label_30394: CALL NC,$5678
label_30395: CALL C,$5678
label_30396: CALL PO,$5678
label_30397: CALL PE,$5678
label_30398: CALL P,$5678
label_30399: CALL M,$5678
label_30400: RET
label_30401: RET NZ
label_30402: RET Z
label_30403: RET NC
label_30404: RET C
label_30405: RET PO
label_30406: RET PE
label_30407: RET P
label_30408: RET M
label_30409: RETI
label_30410: RETN
label_30411: RST $00
label_30412: RST $08
label_30413: RST $10
label_30414: RST $18
label_30415: RST $20
label_30416: RST $28
label_30417: RST $30
label_30418: RST $38
label_30419: IN A,($12)
label_30420: IN A,(C)
label_30421: IN B,(C)
label_30422: IN C,(C)
label_30423: IN D,(C)
label_30424: IN E,(C)
label_30425: IN H,(C)
label_30426: IN L,(C)
label_30427: IN F,(C)
label_30428: INI
label_30429: INIR
label_30430: IND
label_30431: INDR
label_30432: OUT ($12),A
label_30433: OUT (C),A
label_30434: OUT (C),B
label_30435: OUT (C),C
label_30436: OUT (C),D
label_30437: OUT (C),E
label_30438: OUT (C),H
label_30439: OUT (C),L
label_30440: OUTI
label_30441: OTIR
label_30442: OUTD
label_30443: OTDR
label_30444: LD A,A
label_30445: LD A,B
label_30446: LD A,C
label_30447: LD A,D
label_30448: LD A,E
label_30449: LD A,H
label_30450: LD A,L
label_30451: LD B,A
label_30452: LD B,B
label_30453: LD B,C
label_30454: LD B,D
label_30455: LD B,E
label_30456: LD B,H
label_30457: LD B,L
label_30458: LD C,A
label_30459: LD C,B
label_30460: LD C,C
label_30461: LD C,D
label_30462: LD C,E
label_30463: LD C,H
label_30464: LD C,L
label_30465: LD D,A
label_30466: LD D,B
label_30467: LD D,C
label_30468: LD D,D
label_30469: LD D,E
label_30470: LD D,H
label_30471: LD D,L
label_30472: LD E,A
label_30473: LD E,B
label_30474: LD E,C
label_30475: LD E,D
label_30476: LD E,E
label_30477: LD E,H
label_30478: LD E,L
label_30479: LD H,A
label_30480: LD H,B
label_30481: LD H,C
label_30482: LD H,D
label_30483: LD H,E
label_30484: LD H,H
label_30485: LD H,L
label_30486: LD L,A
label_30487: LD L,B
label_30488: LD L,C
label_30489: LD L,D
label_30490: LD L,E
label_30491: LD L,H
label_30492: LD L,L
label_30493: LD A,$12
label_30494: LD B,$12
label_30495: LD C,$12
label_30496: LD D,$12
label_30497: LD E,$12
label_30498: LD H,$12
label_30499: LD L,$12
label_30500: LD A,(HL)
label_30501: LD B,(HL)
label_30502: LD C,(HL)
label_30503: LD D,(HL)
label_30504: LD E,(HL)
label_30505: LD H,(HL)
label_30506: LD L,(HL)
label_30507: LD A,(IX + 127)
label_30508: LD B,(IX + 127)
label_30509: LD C,(IX + 127)
label_30510: LD D,(IX + 127)
label_30511: LD E,(IX + 127)
label_30512: LD H,(IX + 127)
label_30513: LD L,(IX + 127)
label_30514: LD A,(IY - 128)
label_30515: LD B,(IY - 128)
label_30516: LD C,(IY - 128)
label_30517: LD D,(IY - 128)
label_30518: LD E,(IY - 128)
label_30519: LD H,(IY - 128)
label_30520: LD L,(IY - 128)
label_30521: LD (HL),A
label_30522: LD (HL),B
label_30523: LD (HL),C
label_30524: LD (HL),D
label_30525: LD (HL),E
label_30526: LD (HL),H
label_30527: LD (HL),L
label_30528: LD (IX + 127),A
label_30529: LD (IX + 127),B
label_30530: LD (IX + 127),C
label_30531: LD (IX + 127),D
label_30532: LD (IX + 127),E
label_30533: LD (IX + 127),H
label_30534: LD (IX + 127),L
label_30535: LD (IY - 128),A
label_30536: LD (IY - 128),B
label_30537: LD (IY - 128),C
label_30538: LD (IY - 128),D
label_30539: LD (IY - 128),E
label_30540: LD (IY - 128),H
label_30541: LD (IY - 128),L
label_30542: LD (HL),$12
label_30543: LD (IX + 127),$12
label_30544: LD (IY - 128),$12
label_30545: LD A,(BC)
label_30546: LD A,(DE)
label_30547: LD A,($5678)
label_30548: LD (BC),A
label_30549: LD (DE),A
label_30550: LD ($5678),A
label_30551: LD A,I
label_30552: LD A,R
label_30553: LD I,A
label_30554: LD R,A
label_30555: LD BC,$5678
label_30556: LD DE,$5678
label_30557: LD HL,$5678
label_30558: LD SP,$5678
label_30559: LD IX,$5678
label_30560: LD IY,$5678
label_30561: LD HL,($5678)
label_30562: LD BC,($5678)
label_30563: LD DE,($5678)
label_30564: LD HL,($5678)
label_30565: LD SP,($5678)
label_30566: LD IX,($5678)
label_30567: LD IY,($5678)
label_30568: LD ($5678),HL
label_30569: LD ($5678),BC
label_30570: LD ($5678),DE
label_30571: LD ($5678),HL
label_30572: LD ($5678),SP
label_30573: LD ($5678),IX
label_30574: LD ($5678),IY
label_30575: LD SP,HL
label_30576: LD SP,IX
label_30577: LD SP,IY
label_30578: PUSH BC
label_30579: PUSH DE
label_30580: PUSH HL
label_30581: PUSH AF
label_30582: PUSH IX
label_30583: PUSH IY
label_30584: POP BC
label_30585: POP DE
label_30586: POP HL
label_30587: POP AF
label_30588: POP IX
label_30589: POP IY
label_30590: EX DE,HL
label_30591: EX AF,AF'
label_30592: EXX
label_30593: EX (SP),HL
label_30594: EX (SP),IX
label_30595: EX (SP),IY
label_30596: LDI
label_30597: LDIR
label_30598: LDD
label_30599: LDDR
label_30600: CPI
label_30601: CPIR
label_30602: CPD
label_30603: CPDR
label_30604: ADD A,A
label_30605: ADD A,B
label_30606: ADD A,C
label_30607: ADD A,D
label_30608: ADD A,E
label_30609: ADD A,H
label_30610: ADD A,L
label_30611: ADD A,$12
label_30612: ADD A,(HL)
label_30613: ADD A,(IX + 127)
label_30614: ADD A,(IY - 128)
label_30615: ADC A,A
label_30616: ADC A,B
label_30617: ADC A,C
label_30618: ADC A,D
label_30619: ADC A,E
label_30620: ADC A,H
label_30621: ADC A,L
label_30622: ADC A,$12
label_30623: ADC A,(HL)
label_30624: ADC A,(IX + 127)
label_30625: ADC A,(IY - 128)
label_30626: SUB A
label_30627: SUB B
label_30628: SUB C
label_30629: SUB D
label_30630: SUB E
label_30631: SUB H
label_30632: SUB L
label_30633: SUB $12
label_30634: SUB (HL)
label_30635: SUB (IX + 127)
label_30636: SUB (IY - 128)
label_30637: SBC A,A
label_30638: SBC A,B
label_30639: SBC A,C
label_30640: SBC A,D
label_30641: SBC A,E
label_30642: SBC A,H
label_30643: SBC A,L
label_30644: SBC A,$12
label_30645: SBC A,(HL)
label_30646: SBC A,(IX + 127)
label_30647: SBC A,(IY - 128)
label_30648: AND A
label_30649: AND B
label_30650: AND C
label_30651: AND D
label_30652: AND E
label_30653: AND H
label_30654: AND L
label_30655: AND $12
label_30656: AND (HL)
label_30657: AND (IX + 127)
label_30658: AND (IY - 128)
label_30659: AND A
label_30660: AND B
label_30661: AND C
label_30662: AND D
label_30663: AND E
label_30664: AND H
label_30665: AND L
label_30666: AND $12
label_30667: AND (HL)
label_30668: AND (IX + 127)
label_30669: AND (IY - 128)
label_30670: OR A
label_30671: OR B
label_30672: OR C
label_30673: OR D
label_30674: OR E
label_30675: OR H
label_30676: OR L
label_30677: OR $12
label_30678: OR (HL)
label_30679: OR (IX + 127)
label_30680: OR (IY - 128)
label_30681: XOR A
label_30682: XOR B
label_30683: XOR C
label_30684: XOR D
label_30685: XOR E
label_30686: XOR H
label_30687: XOR L
label_30688: XOR $12
label_30689: XOR (HL)
label_30690: XOR (IX + 127)
label_30691: XOR (IY - 128)
label_30692: CP A
label_30693: CP B
label_30694: CP C
label_30695: CP D
label_30696: CP E
label_30697: CP H
label_30698: CP L
label_30699: CP $12
label_30700: CP (HL)
label_30701: CP (IX + 127)
label_30702: CP (IY - 128)
label_30703: INC A
label_30704: INC B
label_30705: INC C
label_30706: INC D
label_30707: INC E
label_30708: INC H
label_30709: INC L
label_30710: INC (HL)
label_30711: INC (IX + 127)
label_30712: INC (IY - 128)
label_30713: DEC A
label_30714: DEC B
label_30715: DEC C
label_30716: DEC D
label_30717: DEC E
label_30718: DEC H
label_30719: DEC L
label_30720: DEC (HL)
label_30721: DEC (IX + 127)
label_30722: DEC (IY - 128)
label_30723: DAA
label_30724: CPL
label_30725: NEG
label_30726: CCF
label_30727: SCF
label_30728: NOP
label_30729: HALT
label_30730: DI
label_30731: EI
label_30732: IM 0
label_30733: IM 1
label_30734: IM 2
label_30735: ADD HL,BC
label_30736: ADD HL,DE
label_30737: ADD HL,HL
label_30738: ADD HL,SP
label_30739: ADC HL,BC
label_30740: ADC HL,DE
label_30741: ADC HL,HL
label_30742: ADC HL,SP
label_30743: SBC HL,BC
label_30744: SBC HL,DE
label_30745: SBC HL,HL
label_30746: SBC HL,SP
label_30747: ADD IX,BC
label_30748: ADD IX,DE
label_30749: ADD IX,SP
label_30750: ADD IY,BC
label_30751: ADD IY,DE
label_30752: ADD IY,SP
label_30753: INC BC
label_30754: INC DE
label_30755: INC HL
label_30756: INC SP
label_30757: INC IX
label_30758: INC IY
label_30759: DEC BC
label_30760: DEC DE
label_30761: DEC HL
label_30762: DEC SP
label_30763: DEC IX
label_30764: DEC IY
label_30765: RLCA
label_30766: RLA
label_30767: RRCA
label_30768: RRA
label_30769: RLC A
label_30770: RLC B
label_30771: RLC C
label_30772: RLC D
label_30773: RLC E
label_30774: RLC H
label_30775: RLC L
label_30776: RLC (HL)
label_30777: RLC (IX + 127)
label_30778: RLC (IY - 128)
label_30779: RL A
label_30780: RL B
label_30781: RL C
label_30782: RL D
label_30783: RL E
label_30784: RL H
label_30785: RL L
label_30786: RL (HL)
label_30787: RL (IX + 127)
label_30788: RL (IY - 128)
label_30789: RRC A
label_30790: RRC B
label_30791: RRC C
label_30792: RRC D
label_30793: RRC E
label_30794: RRC H
label_30795: RRC L
label_30796: RRC (HL)
label_30797: RRC (IX + 127)
label_30798: RRC (IY - 128)
label_30799: RR A
label_30800: RR B
label_30801: RR C
label_30802: RR D
label_30803: RR E
label_30804: RR H
label_30805: RR L
label_30806: RR (HL)
label_30807: RR (IX + 127)
label_30808: RR (IY - 128)
label_30809: SLA A
label_30810: SLA B
label_30811: SLA C
label_30812: SLA D
label_30813: SLA E
label_30814: SLA H
label_30815: SLA L
label_30816: SLA (HL)
label_30817: SLA (IX + 127)
label_30818: SLA (IY - 128)
label_30819: SRA A
label_30820: SRA B
label_30821: SRA C
label_30822: SRA D
label_30823: SRA E
label_30824: SRA H
label_30825: SRA L
label_30826: SRA (HL)
label_30827: SRA (IX + 127)
label_30828: SRA (IY - 128)
label_30829: SRL A
label_30830: SRL B
label_30831: SRL C
label_30832: SRL D
label_30833: SRL E
label_30834: SRL H
label_30835: SRL L
label_30836: SRL (HL)
label_30837: SRL (IX + 127)
label_30838: SRL (IY - 128)
label_30839: RLD
label_30840: RRD
label_30841: BIT 0,A
label_30842: BIT 1,A
label_30843: BIT 2,A
label_30844: BIT 3,A
label_30845: BIT 4,A
label_30846: BIT 5,A
label_30847: BIT 6,A
label_30848: BIT 7,A
label_30849: BIT 0,B
label_30850: BIT 1,B
label_30851: BIT 2,B
label_30852: BIT 3,B
label_30853: BIT 4,B
label_30854: BIT 5,B
label_30855: BIT 6,B
label_30856: BIT 7,B
label_30857: BIT 0,C
label_30858: BIT 1,C
label_30859: BIT 2,C
label_30860: BIT 3,C
label_30861: BIT 4,C
label_30862: BIT 5,C
label_30863: BIT 6,C
label_30864: BIT 7,C
label_30865: BIT 0,D
label_30866: BIT 1,D
label_30867: BIT 2,D
label_30868: BIT 3,D
label_30869: BIT 4,D
label_30870: BIT 5,D
label_30871: BIT 6,D
label_30872: BIT 7,D
label_30873: BIT 0,E
label_30874: BIT 1,E
label_30875: BIT 2,E
label_30876: BIT 3,E
label_30877: BIT 4,E
label_30878: BIT 5,E
label_30879: BIT 6,E
label_30880: BIT 7,E
label_30881: BIT 0,H
label_30882: BIT 1,H
label_30883: BIT 2,H
label_30884: BIT 3,H
label_30885: BIT 4,H
label_30886: BIT 5,H
label_30887: BIT 6,H
label_30888: BIT 7,H
label_30889: BIT 0,L
label_30890: BIT 1,L
label_30891: BIT 2,L
label_30892: BIT 3,L
label_30893: BIT 4,L
label_30894: BIT 5,L
label_30895: BIT 6,L
label_30896: BIT 7,L
label_30897: BIT 0,(HL)
label_30898: BIT 1,(HL)
label_30899: BIT 2,(HL)
label_30900: BIT 3,(HL)
label_30901: BIT 4,(HL)
label_30902: BIT 5,(HL)
label_30903: BIT 6,(HL)
label_30904: BIT 7,(HL)
label_30905: BIT 0,(IX + 127)
label_30906: BIT 1,(IX + 127)
label_30907: BIT 2,(IX + 127)
label_30908: BIT 3,(IX + 127)
label_30909: BIT 4,(IX + 127)
label_30910: BIT 5,(IX + 127)
label_30911: BIT 6,(IX + 127)
label_30912: BIT 7,(IX + 127)
label_30913: BIT 0,(IY - 128)
label_30914: BIT 1,(IY - 128)
label_30915: BIT 2,(IY - 128)
label_30916: BIT 3,(IY - 128)
label_30917: BIT 4,(IY - 128)
label_30918: BIT 5,(IY - 128)
label_30919: BIT 6,(IY - 128)
label_30920: BIT 7,(IY - 128)
label_30921: SET 0,A
label_30922: SET 1,A
label_30923: SET 2,A
label_30924: SET 3,A
label_30925: SET 4,A
label_30926: SET 5,A
label_30927: SET 6,A
label_30928: SET 7,A
label_30929: SET 0,B
label_30930: SET 1,B
label_30931: SET 2,B
label_30932: SET 3,B
label_30933: SET 4,B
label_30934: SET 5,B
label_30935: SET 6,B
label_30936: SET 7,B
label_30937: SET 0,C
label_30938: SET 1,C
label_30939: SET 2,C
label_30940: SET 3,C
label_30941: SET 4,C
label_30942: SET 5,C
label_30943: SET 6,C
label_30944: SET 7,C
label_30945: SET 0,D
label_30946: SET 1,D
label_30947: SET 2,D
label_30948: SET 3,D
label_30949: SET 4,D
label_30950: SET 5,D
label_30951: SET 6,D
label_30952: SET 7,D
label_30953: SET 0,E
label_30954: SET 1,E
label_30955: SET 2,E
label_30956: SET 3,E
label_30957: SET 4,E
label_30958: SET 5,E
label_30959: SET 6,E
label_30960: SET 7,E
label_30961: SET 0,H
label_30962: SET 1,H
label_30963: SET 2,H
label_30964: SET 3,H
label_30965: SET 4,H
label_30966: SET 5,H
label_30967: SET 6,H
label_30968: SET 7,H
label_30969: SET 0,L
label_30970: SET 1,L
label_30971: SET 2,L
label_30972: SET 3,L
label_30973: SET 4,L
label_30974: SET 5,L
label_30975: SET 6,L
label_30976: SET 7,L
label_30977: SET 0,(HL)
label_30978: SET 1,(HL)
label_30979: SET 2,(HL)
label_30980: SET 3,(HL)
label_30981: SET 4,(HL)
label_30982: SET 5,(HL)
label_30983: SET 6,(HL)
label_30984: SET 7,(HL)
label_30985: SET 0,(IX + 127)
label_30986: SET 1,(IX + 127)
label_30987: SET 2,(IX + 127)
label_30988: SET 3,(IX + 127)
label_30989: SET 4,(IX + 127)
label_30990: SET 5,(IX + 127)
label_30991: SET 6,(IX + 127)
label_30992: SET 7,(IX + 127)
label_30993: SET 0,(IY - 128)
label_30994: SET 1,(IY - 128)
label_30995: SET 2,(IY - 128)
label_30996: SET 3,(IY - 128)
label_30997: SET 4,(IY - 128)
label_30998: SET 5,(IY - 128)
label_30999: SET 6,(IY - 128)
label_31000: SET 7,(IY - 128)
label_31001: RES 0,A
label_31002: RES 1,A
label_31003: RES 2,A
label_31004: RES 3,A
label_31005: RES 4,A
label_31006: RES 5,A
label_31007: RES 6,A
label_31008: RES 7,A
label_31009: RES 0,B
label_31010: RES 1,B
label_31011: RES 2,B
label_31012: RES 3,B
label_31013: RES 4,B
label_31014: RES 5,B
label_31015: RES 6,B
label_31016: RES 7,B
label_31017: RES 0,C
label_31018: RES 1,C
label_31019: RES 2,C
label_31020: RES 3,C
label_31021: RES 4,C
label_31022: RES 5,C
label_31023: RES 6,C
label_31024: RES 7,C
label_31025: RES 0,D
label_31026: RES 1,D
label_31027: RES 2,D
label_31028: RES 3,D
label_31029: RES 4,D
label_31030: RES 5,D
label_31031: RES 6,D
label_31032: RES 7,D
label_31033: RES 0,E
label_31034: RES 1,E
label_31035: RES 2,E
label_31036: RES 3,E
label_31037: RES 4,E
label_31038: RES 5,E
label_31039: RES 6,E
label_31040: RES 7,E
label_31041: RES 0,H
label_31042: RES 1,H
label_31043: RES 2,H
label_31044: RES 3,H
label_31045: RES 4,H
label_31046: RES 5,H
label_31047: RES 6,H
label_31048: RES 7,H
label_31049: RES 0,L
label_31050: RES 1,L
label_31051: RES 2,L
label_31052: RES 3,L
label_31053: RES 4,L
label_31054: RES 5,L
label_31055: RES 6,L
label_31056: RES 7,L
label_31057: RES 0,(HL)
label_31058: RES 1,(HL)
label_31059: RES 2,(HL)
label_31060: RES 3,(HL)
label_31061: RES 4,(HL)
label_31062: RES 5,(HL)
label_31063: RES 6,(HL)
label_31064: RES 7,(HL)
label_31065: RES 0,(IX + 127)
label_31066: RES 1,(IX + 127)
label_31067: RES 2,(IX + 127)
label_31068: RES 3,(IX + 127)
label_31069: RES 4,(IX + 127)
label_31070: RES 5,(IX + 127)
label_31071: RES 6,(IX + 127)
label_31072: RES 7,(IX + 127)
label_31073: RES 0,(IY - 128)
label_31074: RES 1,(IY - 128)
label_31075: RES 2,(IY - 128)
label_31076: RES 3,(IY - 128)
label_31077: RES 4,(IY - 128)
label_31078: RES 5,(IY - 128)
label_31079: RES 6,(IY - 128)
label_31080: RES 7,(IY - 128)
label_31081: JP $5678
label_31082: JP NZ,$5678
label_31083: JP Z,$5678
label_31084: JP NC,$5678
label_31085: JP C,$5678
label_31086: JP PO,$5678
label_31087: JP PE,$5678
label_31088: JP P,$5678
label_31089: JP M,$5678
label_31090: JR $ + 2
label_31091: JR NZ,$ + 2
label_31092: JR Z,$ + 2
label_31093: JR NC,$ + 2
label_31094: JR C,$ + 2
label_31095: JP (HL)
label_31096: JP (IX)
label_31097: JP (IY)
label_31098: DJNZ $ + 2
label_31099: CALL $5678
label_31100: CALL NZ,$5678
label_31101: CALL Z,$5678
label_31102: CALL NC,$5678
label_31103: CALL C,$5678
label_31104: CALL PO,$5678
label_31105: CALL PE,$5678
label_31106: CALL P,$5678
label_31107: CALL M,$5678
label_31108: RET
label_31109: RET NZ
label_31110: RET Z
label_31111: RET NC
label_31112: RET C
label_31113: RET PO
label_31114: RET PE
label_31115: RET P
label_31116: RET M
label_31117: RETI
label_31118: RETN
label_31119: RST $00
label_31120: RST $08
label_31121: RST $10
label_31122: RST $18
label_31123: RST $20
label_31124: RST $28
label_31125: RST $30
label_31126: RST $38
label_31127: IN A,($12)
label_31128: IN A,(C)
label_31129: IN B,(C)
label_31130: IN C,(C)
label_31131: IN D,(C)
label_31132: IN E,(C)
label_31133: IN H,(C)
label_31134: IN L,(C)
label_31135: IN F,(C)
label_31136: INI
label_31137: INIR
label_31138: IND
label_31139: INDR
label_31140: OUT ($12),A
label_31141: OUT (C),A
label_31142: OUT (C),B
label_31143: OUT (C),C
label_31144: OUT (C),D
label_31145: OUT (C),E
label_31146: OUT (C),H
label_31147: OUT (C),L
label_31148: OUTI
label_31149: OTIR
label_31150: OUTD
label_31151: OTDR
label_31152: LD A,A
label_31153: LD A,B
label_31154: LD A,C
label_31155: LD A,D
label_31156: LD A,E
label_31157: LD A,H
label_31158: LD A,L
label_31159: LD B,A
label_31160: LD B,B
label_31161: LD B,C
label_31162: LD B,D
label_31163: LD B,E
label_31164: LD B,H
label_31165: LD B,L
label_31166: LD C,A
label_31167: LD C,B
label_31168: LD C,C
label_31169: LD C,D
label_31170: LD C,E
label_31171: LD C,H
label_31172: LD C,L
label_31173: LD D,A
label_31174: LD D,B
label_31175: LD D,C
label_31176: LD D,D
label_31177: LD D,E
label_31178: LD D,H
label_31179: LD D,L
label_31180: LD E,A
label_31181: LD E,B
label_31182: LD E,C
label_31183: LD E,D
label_31184: LD E,E
label_31185: LD E,H
label_31186: LD E,L
label_31187: LD H,A
label_31188: LD H,B
label_31189: LD H,C
label_31190: LD H,D
label_31191: LD H,E
label_31192: LD H,H
label_31193: LD H,L
label_31194: LD L,A
label_31195: LD L,B
label_31196: LD L,C
label_31197: LD L,D
label_31198: LD L,E
label_31199: LD L,H
label_31200: LD L,L
label_31201: LD A,$12
label_31202: LD B,$12
label_31203: LD C,$12
label_31204: LD D,$12
label_31205: LD E,$12
label_31206: LD H,$12
label_31207: LD L,$12
label_31208: LD A,(HL)
label_31209: LD B,(HL)
label_31210: LD C,(HL)
label_31211: LD D,(HL)
label_31212: LD E,(HL)
label_31213: LD H,(HL)
label_31214: LD L,(HL)
label_31215: LD A,(IX + 127)
label_31216: LD B,(IX + 127)
label_31217: LD C,(IX + 127)
label_31218: LD D,(IX + 127)
label_31219: LD E,(IX + 127)
label_31220: LD H,(IX + 127)
label_31221: LD L,(IX + 127)
label_31222: LD A,(IY - 128)
label_31223: LD B,(IY - 128)
label_31224: LD C,(IY - 128)
label_31225: LD D,(IY - 128)
label_31226: LD E,(IY - 128)
label_31227: LD H,(IY - 128)
label_31228: LD L,(IY - 128)
label_31229: LD (HL),A
label_31230: LD (HL),B
label_31231: LD (HL),C
label_31232: LD (HL),D
label_31233: LD (HL),E
label_31234: LD (HL),H
label_31235: LD (HL),L
label_31236: LD (IX + 127),A
label_31237: LD (IX + 127),B
label_31238: LD (IX + 127),C
label_31239: LD (IX + 127),D
label_31240: LD (IX + 127),E
label_31241: LD (IX + 127),H
label_31242: LD (IX + 127),L
label_31243: LD (IY - 128),A
label_31244: LD (IY - 128),B
label_31245: LD (IY - 128),C
label_31246: LD (IY - 128),D
label_31247: LD (IY - 128),E
label_31248: LD (IY - 128),H
label_31249: LD (IY - 128),L
label_31250: LD (HL),$12
label_31251: LD (IX + 127),$12
label_31252: LD (IY - 128),$12
label_31253: LD A,(BC)
label_31254: LD A,(DE)
label_31255: LD A,($5678)
label_31256: LD (BC),A
label_31257: LD (DE),A
label_31258: LD ($5678),A
label_31259: LD A,I
label_31260: LD A,R
label_31261: LD I,A
label_31262: LD R,A
label_31263: LD BC,$5678
label_31264: LD DE,$5678
label_31265: LD HL,$5678
label_31266: LD SP,$5678
label_31267: LD IX,$5678
label_31268: LD IY,$5678
label_31269: LD HL,($5678)
label_31270: LD BC,($5678)
label_31271: LD DE,($5678)
label_31272: LD HL,($5678)
label_31273: LD SP,($5678)
label_31274: LD IX,($5678)
label_31275: LD IY,($5678)
label_31276: LD ($5678),HL
label_31277: LD ($5678),BC
label_31278: LD ($5678),DE
label_31279: LD ($5678),HL
label_31280: LD ($5678),SP
label_31281: LD ($5678),IX
label_31282: LD ($5678),IY
label_31283: LD SP,HL
label_31284: LD SP,IX
label_31285: LD SP,IY
label_31286: PUSH BC
label_31287: PUSH DE
label_31288: PUSH HL
label_31289: PUSH AF
label_31290: PUSH IX
label_31291: PUSH IY
label_31292: POP BC
label_31293: POP DE
label_31294: POP HL
label_31295: POP AF
label_31296: POP IX
label_31297: POP IY
label_31298: EX DE,HL
label_31299: EX AF,AF'
label_31300: EXX
label_31301: EX (SP),HL
label_31302: EX (SP),IX
label_31303: EX (SP),IY
label_31304: LDI
label_31305: LDIR
label_31306: LDD
label_31307: LDDR
label_31308: CPI
label_31309: CPIR
label_31310: CPD
label_31311: CPDR
label_31312: ADD A,A
label_31313: ADD A,B
label_31314: ADD A,C
label_31315: ADD A,D
label_31316: ADD A,E
label_31317: ADD A,H
label_31318: ADD A,L
label_31319: ADD A,$12
label_31320: ADD A,(HL)
label_31321: ADD A,(IX + 127)
label_31322: ADD A,(IY - 128)
label_31323: ADC A,A
label_31324: ADC A,B
label_31325: ADC A,C
label_31326: ADC A,D
label_31327: ADC A,E
label_31328: ADC A,H
label_31329: ADC A,L
label_31330: ADC A,$12
label_31331: ADC A,(HL)
label_31332: ADC A,(IX + 127)
label_31333: ADC A,(IY - 128)
label_31334: SUB A
label_31335: SUB B
label_31336: SUB C
label_31337: SUB D
label_31338: SUB E
label_31339: SUB H
label_31340: SUB L
label_31341: SUB $12
label_31342: SUB (HL)
label_31343: SUB (IX + 127)
label_31344: SUB (IY - 128)
label_31345: SBC A,A
label_31346: SBC A,B
label_31347: SBC A,C
label_31348: SBC A,D
label_31349: SBC A,E
label_31350: SBC A,H
label_31351: SBC A,L
label_31352: SBC A,$12
label_31353: SBC A,(HL)
label_31354: SBC A,(IX + 127)
label_31355: SBC A,(IY - 128)
label_31356: AND A
label_31357: AND B
label_31358: AND C
label_31359: AND D
label_31360: AND E
label_31361: AND H
label_31362: AND L
label_31363: AND $12
label_31364: AND (HL)
label_31365: AND (IX + 127)
label_31366: AND (IY - 128)
label_31367: AND A
label_31368: AND B
label_31369: AND C
label_31370: AND D
label_31371: AND E
label_31372: AND H
label_31373: AND L
label_31374: AND $12
label_31375: AND (HL)
label_31376: AND (IX + 127)
label_31377: AND (IY - 128)
label_31378: OR A
label_31379: OR B
label_31380: OR C
label_31381: OR D
label_31382: OR E
label_31383: OR H
label_31384: OR L
label_31385: OR $12
label_31386: OR (HL)
label_31387: OR (IX + 127)
label_31388: OR (IY - 128)
label_31389: XOR A
label_31390: XOR B
label_31391: XOR C
label_31392: XOR D
label_31393: XOR E
label_31394: XOR H
label_31395: XOR L
label_31396: XOR $12
label_31397: XOR (HL)
label_31398: XOR (IX + 127)
label_31399: XOR (IY - 128)
label_31400: CP A
label_31401: CP B
label_31402: CP C
label_31403: CP D
label_31404: CP E
label_31405: CP H
label_31406: CP L
label_31407: CP $12
label_31408: CP (HL)
label_31409: CP (IX + 127)
label_31410: CP (IY - 128)
label_31411: INC A
label_31412: INC B
label_31413: INC C
label_31414: INC D
label_31415: INC E
label_31416: INC H
label_31417: INC L
label_31418: INC (HL)
label_31419: INC (IX + 127)
label_31420: INC (IY - 128)
label_31421: DEC A
label_31422: DEC B
label_31423: DEC C
label_31424: DEC D
label_31425: DEC E
label_31426: DEC H
label_31427: DEC L
label_31428: DEC (HL)
label_31429: DEC (IX + 127)
label_31430: DEC (IY - 128)
label_31431: DAA
label_31432: CPL
label_31433: NEG
label_31434: CCF
label_31435: SCF
label_31436: NOP
label_31437: HALT
label_31438: DI
label_31439: EI
label_31440: IM 0
label_31441: IM 1
label_31442: IM 2
label_31443: ADD HL,BC
label_31444: ADD HL,DE
label_31445: ADD HL,HL
label_31446: ADD HL,SP
label_31447: ADC HL,BC
label_31448: ADC HL,DE
label_31449: ADC HL,HL
label_31450: ADC HL,SP
label_31451: SBC HL,BC
label_31452: SBC HL,DE
label_31453: SBC HL,HL
label_31454: SBC HL,SP
label_31455: ADD IX,BC
label_31456: ADD IX,DE
label_31457: ADD IX,SP
label_31458: ADD IY,BC
label_31459: ADD IY,DE
label_31460: ADD IY,SP
label_31461: INC BC
label_31462: INC DE
label_31463: INC HL
label_31464: INC SP
label_31465: INC IX
label_31466: INC IY
label_31467: DEC BC
label_31468: DEC DE
label_31469: DEC HL
label_31470: DEC SP
label_31471: DEC IX
label_31472: DEC IY
label_31473: RLCA
label_31474: RLA
label_31475: RRCA
label_31476: RRA
label_31477: RLC A
label_31478: RLC B
label_31479: RLC C
label_31480: RLC D
label_31481: RLC E
label_31482: RLC H
label_31483: RLC L
label_31484: RLC (HL)
label_31485: RLC (IX + 127)
label_31486: RLC (IY - 128)
label_31487: RL A
label_31488: RL B
label_31489: RL C
label_31490: RL D
label_31491: RL E
label_31492: RL H
label_31493: RL L
label_31494: RL (HL)
label_31495: RL (IX + 127)
label_31496: RL (IY - 128)
label_31497: RRC A
label_31498: RRC B
label_31499: RRC C
label_31500: RRC D
label_31501: RRC E
label_31502: RRC H
label_31503: RRC L
label_31504: RRC (HL)
label_31505: RRC (IX + 127)
label_31506: RRC (IY - 128)
label_31507: RR A
label_31508: RR B
label_31509: RR C
label_31510: RR D
label_31511: RR E
label_31512: RR H
label_31513: RR L
label_31514: RR (HL)
label_31515: RR (IX + 127)
label_31516: RR (IY - 128)
label_31517: SLA A
label_31518: SLA B
label_31519: SLA C
label_31520: SLA D
label_31521: SLA E
label_31522: SLA H
label_31523: SLA L
label_31524: SLA (HL)
label_31525: SLA (IX + 127)
label_31526: SLA (IY - 128)
label_31527: SRA A
label_31528: SRA B
label_31529: SRA C
label_31530: SRA D
label_31531: SRA E
label_31532: SRA H
label_31533: SRA L
label_31534: SRA (HL)
label_31535: SRA (IX + 127)
label_31536: SRA (IY - 128)
label_31537: SRL A
label_31538: SRL B
label_31539: SRL C
label_31540: SRL D
label_31541: SRL E
label_31542: SRL H
label_31543: SRL L
label_31544: SRL (HL)
label_31545: SRL (IX + 127)
label_31546: SRL (IY - 128)
label_31547: RLD
label_31548: RRD
label_31549: BIT 0,A
label_31550: BIT 1,A
label_31551: BIT 2,A
label_31552: BIT 3,A
label_31553: BIT 4,A
label_31554: BIT 5,A
label_31555: BIT 6,A
label_31556: BIT 7,A
label_31557: BIT 0,B
label_31558: BIT 1,B
label_31559: BIT 2,B
label_31560: BIT 3,B
label_31561: BIT 4,B
label_31562: BIT 5,B
label_31563: BIT 6,B
label_31564: BIT 7,B
label_31565: BIT 0,C
label_31566: BIT 1,C
label_31567: BIT 2,C
label_31568: BIT 3,C
label_31569: BIT 4,C
label_31570: BIT 5,C
label_31571: BIT 6,C
label_31572: BIT 7,C
label_31573: BIT 0,D
label_31574: BIT 1,D
label_31575: BIT 2,D
label_31576: BIT 3,D
label_31577: BIT 4,D
label_31578: BIT 5,D
label_31579: BIT 6,D
label_31580: BIT 7,D
label_31581: BIT 0,E
label_31582: BIT 1,E
label_31583: BIT 2,E
label_31584: BIT 3,E
label_31585: BIT 4,E
label_31586: BIT 5,E
label_31587: BIT 6,E
label_31588: BIT 7,E
label_31589: BIT 0,H
label_31590: BIT 1,H
label_31591: BIT 2,H
label_31592: BIT 3,H
label_31593: BIT 4,H
label_31594: BIT 5,H
label_31595: BIT 6,H
label_31596: BIT 7,H
label_31597: BIT 0,L
label_31598: BIT 1,L
label_31599: BIT 2,L
label_31600: BIT 3,L
label_31601: BIT 4,L
label_31602: BIT 5,L
label_31603: BIT 6,L
label_31604: BIT 7,L
label_31605: BIT 0,(HL)
label_31606: BIT 1,(HL)
label_31607: BIT 2,(HL)
label_31608: BIT 3,(HL)
label_31609: BIT 4,(HL)
label_31610: BIT 5,(HL)
label_31611: BIT 6,(HL)
label_31612: BIT 7,(HL)
label_31613: BIT 0,(IX + 127)
label_31614: BIT 1,(IX + 127)
label_31615: BIT 2,(IX + 127)
label_31616: BIT 3,(IX + 127)
label_31617: BIT 4,(IX + 127)
label_31618: BIT 5,(IX + 127)
label_31619: BIT 6,(IX + 127)
label_31620: BIT 7,(IX + 127)
label_31621: BIT 0,(IY - 128)
label_31622: BIT 1,(IY - 128)
label_31623: BIT 2,(IY - 128)
label_31624: BIT 3,(IY - 128)
label_31625: BIT 4,(IY - 128)
label_31626: BIT 5,(IY - 128)
label_31627: BIT 6,(IY - 128)
label_31628: BIT 7,(IY - 128)
label_31629: SET 0,A
label_31630: SET 1,A
label_31631: SET 2,A
label_31632: SET 3,A
label_31633: SET 4,A
label_31634: SET 5,A
label_31635: SET 6,A
label_31636: SET 7,A
label_31637: SET 0,B
label_31638: SET 1,B
label_31639: SET 2,B
label_31640: SET 3,B
label_31641: SET 4,B
label_31642: SET 5,B
label_31643: SET 6,B
label_31644: SET 7,B
label_31645: SET 0,C
label_31646: SET 1,C
label_31647: SET 2,C
label_31648: SET 3,C
label_31649: SET 4,C
label_31650: SET 5,C
label_31651: SET 6,C
label_31652: SET 7,C
label_31653: SET 0,D
label_31654: SET 1,D
label_31655: SET 2,D
label_31656: SET 3,D
label_31657: SET 4,D
label_31658: SET 5,D
label_31659: SET 6,D
label_31660: SET 7,D
label_31661: SET 0,E
label_31662: SET 1,E
label_31663: SET 2,E
label_31664: SET 3,E
label_31665: SET 4,E
label_31666: SET 5,E
label_31667: SET 6,E
label_31668: SET 7,E
label_31669: SET 0,H
label_31670: SET 1,H
label_31671: SET 2,H
label_31672: SET 3,H
label_31673: SET 4,H
label_31674: SET 5,H
label_31675: SET 6,H
label_31676: SET 7,H
label_31677: SET 0,L
label_31678: SET 1,L
label_31679: SET 2,L
label_31680: SET 3,L
label_31681: SET 4,L
label_31682: SET 5,L
label_31683: SET 6,L
label_31684: SET 7,L
label_31685: SET 0,(HL)
label_31686: SET 1,(HL)
label_31687: SET 2,(HL)
label_31688: SET 3,(HL)
label_31689: SET 4,(HL)
label_31690: SET 5,(HL)
label_31691: SET 6,(HL)
label_31692: SET 7,(HL)
label_31693: SET 0,(IX + 127)
label_31694: SET 1,(IX + 127)
label_31695: SET 2,(IX + 127)
label_31696: SET 3,(IX + 127)
label_31697: SET 4,(IX + 127)
label_31698: SET 5,(IX + 127)
label_31699: SET 6,(IX + 127)
label_31700: SET 7,(IX + 127)
label_31701: SET 0,(IY - 128)
label_31702: SET 1,(IY - 128)
label_31703: SET 2,(IY - 128)
label_31704: SET 3,(IY - 128)
label_31705: SET 4,(IY - 128)
label_31706: SET 5,(IY - 128)
label_31707: SET 6,(IY - 128)
label_31708: SET 7,(IY - 128)
label_31709: RES 0,A
label_31710: RES 1,A
label_31711: RES 2,A
label_31712: RES 3,A
label_31713: RES 4,A
label_31714: RES 5,A
label_31715: RES 6,A
label_31716: RES 7,A
label_31717: RES 0,B
label_31718: RES 1,B
label_31719: RES 2,B
label_31720: RES 3,B
label_31721: RES 4,B
label_31722: RES 5,B
label_31723: RES 6,B
label_31724: RES 7,B
label_31725: RES 0,C
label_31726: RES 1,C
label_31727: RES 2,C
label_31728: RES 3,C
label_31729: RES 4,C
label_31730: RES 5,C
label_31731: RES 6,C
label_31732: RES 7,C
label_31733: RES 0,D
label_31734: RES 1,D
label_31735: RES 2,D
label_31736: RES 3,D
label_31737: RES 4,D
label_31738: RES 5,D
label_31739: RES 6,D
label_31740: RES 7,D
label_31741: RES 0,E
label_31742: RES 1,E
label_31743: RES 2,E
label_31744: RES 3,E
label_31745: RES 4,E
label_31746: RES 5,E
label_31747: RES 6,E
label_31748: RES 7,E
label_31749: RES 0,H
label_31750: RES 1,H
label_31751: RES 2,H
label_31752: RES 3,H
label_31753: RES 4,H
label_31754: RES 5,H
label_31755: RES 6,H
label_31756: RES 7,H
label_31757: RES 0,L
label_31758: RES 1,L
label_31759: RES 2,L
label_31760: RES 3,L
label_31761: RES 4,L
label_31762: RES 5,L
label_31763: RES 6,L
label_31764: RES 7,L
label_31765: RES 0,(HL)
label_31766: RES 1,(HL)
label_31767: RES 2,(HL)
label_31768: RES 3,(HL)
label_31769: RES 4,(HL)
label_31770: RES 5,(HL)
label_31771: RES 6,(HL)
label_31772: RES 7,(HL)
label_31773: RES 0,(IX + 127)
label_31774: RES 1,(IX + 127)
label_31775: RES 2,(IX + 127)
label_31776: RES 3,(IX + 127)
label_31777: RES 4,(IX + 127)
label_31778: RES 5,(IX + 127)
label_31779: RES 6,(IX + 127)
label_31780: RES 7,(IX + 127)
label_31781: RES 0,(IY - 128)
label_31782: RES 1,(IY - 128)
label_31783: RES 2,(IY - 128)
label_31784: RES 3,(IY - 128)
label_31785: RES 4,(IY - 128)
label_31786: RES 5,(IY - 128)
label_31787: RES 6,(IY - 128)
label_31788: RES 7,(IY - 128)
label_31789: JP $5678
label_31790: JP NZ,$5678
label_31791: JP Z,$5678
label_31792: JP NC,$5678
label_31793: JP C,$5678
label_31794: JP PO,$5678
label_31795: JP PE,$5678
label_31796: JP P,$5678
label_31797: JP M,$5678
label_31798: JR $ + 2
label_31799: JR NZ,$ + 2
label_31800: JR Z,$ + 2
label_31801: JR NC,$ + 2
label_31802: JR C,$ + 2
label_31803: JP (HL)
label_31804: JP (IX)
label_31805: JP (IY)
label_31806: DJNZ $ + 2
label_31807: CALL $5678
label_31808: CALL NZ,$5678
label_31809: CALL Z,$5678
label_31810: CALL NC,$5678
label_31811: CALL C,$5678
label_31812: CALL PO,$5678
label_31813: CALL PE,$5678
label_31814: CALL P,$5678
label_31815: CALL M,$5678
label_31816: RET
label_31817: RET NZ
label_31818: RET Z
label_31819: RET NC
label_31820: RET C
label_31821: RET PO
label_31822: RET PE
label_31823: RET P
label_31824: RET M
label_31825: RETI
label_31826: RETN
label_31827: RST $00
label_31828: RST $08
label_31829: RST $10
label_31830: RST $18
label_31831: RST $20
label_31832: RST $28
label_31833: RST $30
label_31834: RST $38
label_31835: IN A,($12)
label_31836: IN A,(C)
label_31837: IN B,(C)
label_31838: IN C,(C)
label_31839: IN D,(C)
label_31840: IN E,(C)
label_31841: IN H,(C)
label_31842: IN L,(C)
label_31843: IN F,(C)
label_31844: INI
label_31845: INIR
label_31846: IND
label_31847: INDR
label_31848: OUT ($12),A
label_31849: OUT (C),A
label_31850: OUT (C),B
label_31851: OUT (C),C
label_31852: OUT (C),D
label_31853: OUT (C),E
label_31854: OUT (C),H
label_31855: OUT (C),L
label_31856: OUTI
label_31857: OTIR
label_31858: OUTD
label_31859: OTDR
label_31860: LD A,A
label_31861: LD A,B
label_31862: LD A,C
label_31863: LD A,D
label_31864: LD A,E
label_31865: LD A,H
label_31866: LD A,L
label_31867: LD B,A
label_31868: LD B,B
label_31869: LD B,C
label_31870: LD B,D
label_31871: LD B,E
label_31872: LD B,H
label_31873: LD B,L
label_31874: LD C,A
label_31875: LD C,B
label_31876: LD C,C
label_31877: LD C,D
label_31878: LD C,E
label_31879: LD C,H
label_31880: LD C,L
label_31881: LD D,A
label_31882: LD D,B
label_31883: LD D,C
label_31884: LD D,D
label_31885: LD D,E
label_31886: LD D,H
label_31887: LD D,L
label_31888: LD E,A
label_31889: LD E,B
label_31890: LD E,C
label_31891: LD E,D
label_31892: LD E,E
label_31893: LD E,H
label_31894: LD E,L
label_31895: LD H,A
label_31896: LD H,B
label_31897: LD H,C
label_31898: LD H,D
label_31899: LD H,E
label_31900: LD H,H
label_31901: LD H,L
label_31902: LD L,A
label_31903: LD L,B
label_31904: LD L,C
label_31905: LD L,D
label_31906: LD L,E
label_31907: LD L,H
label_31908: LD L,L
label_31909: LD A,$12
label_31910: LD B,$12
label_31911: LD C,$12
label_31912: LD D,$12
label_31913: LD E,$12
label_31914: LD H,$12
label_31915: LD L,$12
label_31916: LD A,(HL)
label_31917: LD B,(HL)
label_31918: LD C,(HL)
label_31919: LD D,(HL)
label_31920: LD E,(HL)
label_31921: LD H,(HL)
label_31922: LD L,(HL)
label_31923: LD A,(IX + 127)
label_31924: LD B,(IX + 127)
label_31925: LD C,(IX + 127)
label_31926: LD D,(IX + 127)
label_31927: LD E,(IX + 127)
label_31928: LD H,(IX + 127)
label_31929: LD L,(IX + 127)
label_31930: LD A,(IY - 128)
label_31931: LD B,(IY - 128)
label_31932: LD C,(IY - 128)
label_31933: LD D,(IY - 128)
label_31934: LD E,(IY - 128)
label_31935: LD H,(IY - 128)
label_31936: LD L,(IY - 128)
label_31937: LD (HL),A
label_31938: LD (HL),B
label_31939: LD (HL),C
label_31940: LD (HL),D
label_31941: LD (HL),E
label_31942: LD (HL),H
label_31943: LD (HL),L
label_31944: LD (IX + 127),A
label_31945: LD (IX + 127),B
label_31946: LD (IX + 127),C
label_31947: LD (IX + 127),D
label_31948: LD (IX + 127),E
label_31949: LD (IX + 127),H
label_31950: LD (IX + 127),L
label_31951: LD (IY - 128),A
label_31952: LD (IY - 128),B
label_31953: LD (IY - 128),C
label_31954: LD (IY - 128),D
label_31955: LD (IY - 128),E
label_31956: LD (IY - 128),H
label_31957: LD (IY - 128),L
label_31958: LD (HL),$12
label_31959: LD (IX + 127),$12
label_31960: LD (IY - 128),$12
label_31961: LD A,(BC)
label_31962: LD A,(DE)
label_31963: LD A,($5678)
label_31964: LD (BC),A
label_31965: LD (DE),A
label_31966: LD ($5678),A
label_31967: LD A,I
label_31968: LD A,R
label_31969: LD I,A
label_31970: LD R,A
label_31971: LD BC,$5678
label_31972: LD DE,$5678
label_31973: LD HL,$5678
label_31974: LD SP,$5678
label_31975: LD IX,$5678
label_31976: LD IY,$5678
label_31977: LD HL,($5678)
label_31978: LD BC,($5678)
label_31979: LD DE,($5678)
label_31980: LD HL,($5678)
label_31981: LD SP,($5678)
label_31982: LD IX,($5678)
label_31983: LD IY,($5678)
label_31984: LD ($5678),HL
label_31985: LD ($5678),BC
label_31986: LD ($5678),DE
label_31987: LD ($5678),HL
label_31988: LD ($5678),SP
label_31989: LD ($5678),IX
label_31990: LD ($5678),IY
label_31991: LD SP,HL
label_31992: LD SP,IX
label_31993: LD SP,IY
label_31994: PUSH BC
label_31995: PUSH DE
label_31996: PUSH HL
label_31997: PUSH AF
label_31998: PUSH IX
label_31999: PUSH IY
label_32000: POP BC
label_32001: POP DE
label_32002: POP HL
label_32003: POP AF
label_32004: POP IX
label_32005: POP IY
label_32006: EX DE,HL
label_32007: EX AF,AF'
label_32008: EXX
label_32009: EX (SP),HL
label_32010: EX (SP),IX
label_32011: EX (SP),IY
label_32012: LDI
label_32013: LDIR
label_32014: LDD
label_32015: LDDR
label_32016: CPI
label_32017: CPIR
label_32018: CPD
label_32019: CPDR
label_32020: ADD A,A
label_32021: ADD A,B
label_32022: ADD A,C
label_32023: ADD A,D
label_32024: ADD A,E
label_32025: ADD A,H
label_32026: ADD A,L
label_32027: ADD A,$12
label_32028: ADD A,(HL)
label_32029: ADD A,(IX + 127)
label_32030: ADD A,(IY - 128)
label_32031: ADC A,A
label_32032: ADC A,B
label_32033: ADC A,C
label_32034: ADC A,D
label_32035: ADC A,E
label_32036: ADC A,H
label_32037: ADC A,L
label_32038: ADC A,$12
label_32039: ADC A,(HL)
label_32040: ADC A,(IX + 127)
label_32041: ADC A,(IY - 128)
label_32042: SUB A
label_32043: SUB B
label_32044: SUB C
label_32045: SUB D
label_32046: SUB E
label_32047: SUB H
label_32048: SUB L
label_32049: SUB $12
label_32050: SUB (HL)
label_32051: SUB (IX + 127)
label_32052: SUB (IY - 128)
label_32053: SBC A,A
label_32054: SBC A,B
label_32055: SBC A,C
label_32056: SBC A,D
label_32057: SBC A,E
label_32058: SBC A,H
label_32059: SBC A,L
label_32060: SBC A,$12
label_32061: SBC A,(HL)
label_32062: SBC A,(IX + 127)
label_32063: SBC A,(IY - 128)
label_32064: AND A
label_32065: AND B
label_32066: AND C
label_32067: AND D
label_32068: AND E
label_32069: AND H
label_32070: AND L
label_32071: AND $12
label_32072: AND (HL)
label_32073: AND (IX + 127)
label_32074: AND (IY - 128)
label_32075: AND A
label_32076: AND B
label_32077: AND C
label_32078: AND D
label_32079: AND E
label_32080: AND H
label_32081: AND L
label_32082: AND $12
label_32083: AND (HL)
label_32084: AND (IX + 127)
label_32085: AND (IY - 128)
label_32086: OR A
label_32087: OR B
label_32088: OR C
label_32089: OR D
label_32090: OR E
label_32091: OR H
label_32092: OR L
label_32093: OR $12
label_32094: OR (HL)
label_32095: OR (IX + 127)
label_32096: OR (IY - 128)
label_32097: XOR A
label_32098: XOR B
label_32099: XOR C
label_32100: XOR D
label_32101: XOR E
label_32102: XOR H
label_32103: XOR L
label_32104: XOR $12
label_32105: XOR (HL)
label_32106: XOR (IX + 127)
label_32107: XOR (IY - 128)
label_32108: CP A
label_32109: CP B
label_32110: CP C
label_32111: CP D
label_32112: CP E
label_32113: CP H
label_32114: CP L
label_32115: CP $12
label_32116: CP (HL)
label_32117: CP (IX + 127)
label_32118: CP (IY - 128)
label_32119: INC A
label_32120: INC B
label_32121: INC C
label_32122: INC D
label_32123: INC E
label_32124: INC H
label_32125: INC L
label_32126: INC (HL)
label_32127: INC (IX + 127)
label_32128: INC (IY - 128)
label_32129: DEC A
label_32130: DEC B
label_32131: DEC C
label_32132: DEC D
label_32133: DEC E
label_32134: DEC H
label_32135: DEC L
label_32136: DEC (HL)
label_32137: DEC (IX + 127)
label_32138: DEC (IY - 128)
label_32139: DAA
label_32140: CPL
label_32141: NEG
label_32142: CCF
label_32143: SCF
label_32144: NOP
label_32145: HALT
label_32146: DI
label_32147: EI
label_32148: IM 0
label_32149: IM 1
label_32150: IM 2
label_32151: ADD HL,BC
label_32152: ADD HL,DE
label_32153: ADD HL,HL
label_32154: ADD HL,SP
label_32155: ADC HL,BC
label_32156: ADC HL,DE
label_32157: ADC HL,HL
label_32158: ADC HL,SP
label_32159: SBC HL,BC
label_32160: SBC HL,DE
label_32161: SBC HL,HL
label_32162: SBC HL,SP
label_32163: ADD IX,BC
label_32164: ADD IX,DE
label_32165: ADD IX,SP
label_32166: ADD IY,BC
label_32167: ADD IY,DE
label_32168: ADD IY,SP
label_32169: INC BC
label_32170: INC DE
label_32171: INC HL
label_32172: INC SP
label_32173: INC IX
label_32174: INC IY
label_32175: DEC BC
label_32176: DEC DE
label_32177: DEC HL
label_32178: DEC SP
label_32179: DEC IX
label_32180: DEC IY
label_32181: RLCA
label_32182: RLA
label_32183: RRCA
label_32184: RRA
label_32185: RLC A
label_32186: RLC B
label_32187: RLC C
label_32188: RLC D
label_32189: RLC E
label_32190: RLC H
label_32191: RLC L
label_32192: RLC (HL)
label_32193: RLC (IX + 127)
label_32194: RLC (IY - 128)
label_32195: RL A
label_32196: RL B
label_32197: RL C
label_32198: RL D
label_32199: RL E
label_32200: RL H
label_32201: RL L
label_32202: RL (HL)
label_32203: RL (IX + 127)
label_32204: RL (IY - 128)
label_32205: RRC A
label_32206: RRC B
label_32207: RRC C
label_32208: RRC D
label_32209: RRC E
label_32210: RRC H
label_32211: RRC L
label_32212: RRC (HL)
label_32213: RRC (IX + 127)
label_32214: RRC (IY - 128)
label_32215: RR A
label_32216: RR B
label_32217: RR C
label_32218: RR D
label_32219: RR E
label_32220: RR H
label_32221: RR L
label_32222: RR (HL)
label_32223: RR (IX + 127)
label_32224: RR (IY - 128)
label_32225: SLA A
label_32226: SLA B
label_32227: SLA C
label_32228: SLA D
label_32229: SLA E
label_32230: SLA H
label_32231: SLA L
label_32232: SLA (HL)
label_32233: SLA (IX + 127)
label_32234: SLA (IY - 128)
label_32235: SRA A
label_32236: SRA B
label_32237: SRA C
label_32238: SRA D
label_32239: SRA E
label_32240: SRA H
label_32241: SRA L
label_32242: SRA (HL)
label_32243: SRA (IX + 127)
label_32244: SRA (IY - 128)
label_32245: SRL A
label_32246: SRL B
label_32247: SRL C
label_32248: SRL D
label_32249: SRL E
label_32250: SRL H
label_32251: SRL L
label_32252: SRL (HL)
label_32253: SRL (IX + 127)
label_32254: SRL (IY - 128)
label_32255: RLD
label_32256: RRD
label_32257: BIT 0,A
label_32258: BIT 1,A
label_32259: BIT 2,A
label_32260: BIT 3,A
label_32261: BIT 4,A
label_32262: BIT 5,A
label_32263: BIT 6,A
label_32264: BIT 7,A
label_32265: BIT 0,B
label_32266: BIT 1,B
label_32267: BIT 2,B
label_32268: BIT 3,B
label_32269: BIT 4,B
label_32270: BIT 5,B
label_32271: BIT 6,B
label_32272: BIT 7,B
label_32273: BIT 0,C
label_32274: BIT 1,C
label_32275: BIT 2,C
label_32276: BIT 3,C
label_32277: BIT 4,C
label_32278: BIT 5,C
label_32279: BIT 6,C
label_32280: BIT 7,C
label_32281: BIT 0,D
label_32282: BIT 1,D
label_32283: BIT 2,D
label_32284: BIT 3,D
label_32285: BIT 4,D
label_32286: BIT 5,D
label_32287: BIT 6,D
label_32288: BIT 7,D
label_32289: BIT 0,E
label_32290: BIT 1,E
label_32291: BIT 2,E
label_32292: BIT 3,E
label_32293: BIT 4,E
label_32294: BIT 5,E
label_32295: BIT 6,E
label_32296: BIT 7,E
label_32297: BIT 0,H
label_32298: BIT 1,H
label_32299: BIT 2,H
label_32300: BIT 3,H
label_32301: BIT 4,H
label_32302: BIT 5,H
label_32303: BIT 6,H
label_32304: BIT 7,H
label_32305: BIT 0,L
label_32306: BIT 1,L
label_32307: BIT 2,L
label_32308: BIT 3,L
label_32309: BIT 4,L
label_32310: BIT 5,L
label_32311: BIT 6,L
label_32312: BIT 7,L
label_32313: BIT 0,(HL)
label_32314: BIT 1,(HL)
label_32315: BIT 2,(HL)
label_32316: BIT 3,(HL)
label_32317: BIT 4,(HL)
label_32318: BIT 5,(HL)
label_32319: BIT 6,(HL)
label_32320: BIT 7,(HL)
label_32321: BIT 0,(IX + 127)
label_32322: BIT 1,(IX + 127)
label_32323: BIT 2,(IX + 127)
label_32324: BIT 3,(IX + 127)
label_32325: BIT 4,(IX + 127)
label_32326: BIT 5,(IX + 127)
label_32327: BIT 6,(IX + 127)
label_32328: BIT 7,(IX + 127)
label_32329: BIT 0,(IY - 128)
label_32330: BIT 1,(IY - 128)
label_32331: BIT 2,(IY - 128)
label_32332: BIT 3,(IY - 128)
label_32333: BIT 4,(IY - 128)
label_32334: BIT 5,(IY - 128)
label_32335: BIT 6,(IY - 128)
label_32336: BIT 7,(IY - 128)
label_32337: SET 0,A
label_32338: SET 1,A
label_32339: SET 2,A
