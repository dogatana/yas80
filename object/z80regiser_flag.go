package object

import "yas80/parser"

// Z80 Regisers and Flags
var (
	Z80_REG_A   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_A}
	Z80_REG_B   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_B}
	Z80_REG_C   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_C}
	Z80_REG_D   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_D}
	Z80_REG_E   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_E}
	Z80_REG_H   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_H}
	Z80_REG_L   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_L}
	Z80_REG_IXH = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_IXH}
	Z80_REG_IXL = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_IXL}
	Z80_REG_IYH = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_IYH}
	Z80_REG_IYL = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_IYL}
	Z80_REG_I   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_I}
	Z80_REG_R   = &RegisterObject{RegisterType: parser.Z80_REG8, Register: parser.Z80_REG_R}

	Z80_REG_SP   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_SP}
	Z80_REG_IX   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_IX}
	Z80_REG_IY   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_IY}
	Z80_REG_AF   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_AF}
	Z80_REG_AFEX = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_AFEX}
	Z80_REG_BC   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_BC}
	Z80_REG_DE   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_DE}
	Z80_REG_HL   = &RegisterObject{RegisterType: parser.Z80_REG16, Register: parser.Z80_REG_HL}

	Z80_FLAG_C  = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_C}
	Z80_FLAG_NC = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_NC}
	Z80_FLAG_Z  = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_Z}
	Z80_FLAG_NZ = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_NZ}
	Z80_FLAG_PO = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_PO}
	Z80_FLAG_PE = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_PE}
	Z80_FLAG_P  = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_P}
	Z80_FLAG_M  = &RegisterObject{RegisterType: parser.Z80_FLAG, Register: parser.Z80_FLAG_M}
)

// TokenSubtype to Object
var Z80RgisterObjects map[int]Object = map[int]Object{
	parser.Z80_REG_A:   Z80_REG_A,
	parser.Z80_REG_B:   Z80_REG_B,
	parser.Z80_REG_C:   Z80_REG_C,
	parser.Z80_REG_D:   Z80_REG_D,
	parser.Z80_REG_E:   Z80_REG_E,
	parser.Z80_REG_H:   Z80_REG_H,
	parser.Z80_REG_L:   Z80_REG_L,
	parser.Z80_REG_IXH: Z80_REG_IXH,
	parser.Z80_REG_IXL: Z80_REG_IXL,
	parser.Z80_REG_IYH: Z80_REG_IYH,
	parser.Z80_REG_IYL: Z80_REG_IYL,
	parser.Z80_REG_I:   Z80_REG_IYL,
	parser.Z80_REG_R:   Z80_REG_IYL,

	parser.Z80_REG_SP:   Z80_REG_SP,
	parser.Z80_REG_IX:   Z80_REG_IX,
	parser.Z80_REG_IY:   Z80_REG_IY,
	parser.Z80_REG_AF:   Z80_REG_AF,
	parser.Z80_REG_AFEX: Z80_REG_AFEX,
	parser.Z80_REG_BC:   Z80_REG_BC,
	parser.Z80_REG_DE:   Z80_REG_DE,
	parser.Z80_REG_HL:   Z80_REG_HL,

	parser.Z80_FLAG_NC: Z80_FLAG_NC,
	parser.Z80_FLAG_Z:  Z80_FLAG_Z,
	parser.Z80_FLAG_NZ: Z80_FLAG_NZ,
	parser.Z80_FLAG_PO: Z80_FLAG_PO,
	parser.Z80_FLAG_PE: Z80_FLAG_PE,
	parser.Z80_FLAG_P:  Z80_FLAG_P,
	parser.Z80_FLAG_M:  Z80_FLAG_M,
}
