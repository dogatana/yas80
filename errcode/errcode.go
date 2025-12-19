package errcode

const (
	// system
	E000 = "[E]system error"
	E001 = "[E]syntax error"
	// user
	E002 = "数値リテラル誤り: '%s'"
	E003 = "配列名誤り"
	E004 = "配列インデックス未指定"
	E005 = "配列インデックス誤り"
	E006 = "間接指定オペランド誤り"

	EUNI_OP_STRING = "'%c' は単項文字列演算子ではない"
	EUNI_OP_NUMBER = "'%c' は単項数値演算子ではない"
	EUNI_OP_TYPE   = "単項演算子 '%s' が使用できない型"

	E009 = "'%s' は未定義"

	EENUM_UNDEF     = "ENUM '%s' は定義されていない"
	EENUM_UNDEF_ELE = "ENUM '%s.%s' は定義されていない"
	EENUM_USED      = "ENUM '%s' は定義済み"
	EENUM_USED_ELE  = "ENUM '%s.%s' は定義済み"
	EENUM_ELE_TYPE  = "ENUM 要素に使用できない型 %T"

	EBIN_OP_DIVZERO = "0 除算"
	EBIN_OP_STRING  = "文字列は '+' 演算子のみ使用可能"
	EBIN_OP_NUMBER  = "不明な整数演算子 '%s'"
	EBIN_OP_TYPE    = "二項演算子 '%s' は使用不可"

	EFUNC_DUP       = "FUNC '%s' は定義済み"
	EFUNC_NAME      = "関数呼出しには関数名が必要"
	EFUNC_UNDEF     = "FUNC '%s' は未定義"
	EFUNC_ARG_COUNT = "FUNC '%s' の仮引数の数と関数呼出しの実引数の数が不一致"

	EZ80_FLAG       = "'%s' はフラグでない"
	EZ80_OP1        = "第1オペランド誤り"
	EZ80_OP2        = "第2オペランド誤り"
	EZ80_OP1_SP     = "第1オペランドは SP 以外指定不可"
	EZ80_OP2_HL_IXY = "第2オペランドは HL, IX, IY 以外指定不可"

	E028 = "'%s' は指定不可"
	E030 = "'%s' の定義に循環参照あり"

	EASSIGN_INVALID_TAGET = "変数/_ 以外へは代入不可"
	EASSIGN_INVALID_VALUE = "代入右辺式の値が未確定"

	ESYM_DUP            = "CONST/EQU '%s' は定義済み"
	ESYM_USED_NAME      = "%s は利用済のため CONST/EQU 定義不可"
	ESYM_UNDEF          = "シンボル %s は未定義"
	ESYM_CYCLIC         = "シンボル %s の定義が循環参照を含む"
	ESYM_NOT_DETERMINED = "シンボル %s の値を確定できない"

	ESYM_CONCAT_NOTSYM = "シンボル結合にはシンボルが必要"
	ESYM_CONCAT_EXPR   = "シンボル結合式の誤り"
	ESYM_CONCAT_TYPE   = "シンボル結合は数値、文字列のみ可能"

	ELABEL_DUP  = "LABEL %s は定義済み"        // symbol 以外で利用済
	ELABEL_USED = "%s は利用済のため LABEL 定義不可" // symbol 以外で利用済

	EMACRO_UNDEF     = "MACRO %s は未定義"
	EMACRO_DUP       = "MACRO %s は定義済み" // TODO 最終的に不要
	EMACRO_NEST      = "MACRO 定義はネスト不可" // TODO 最終的に不要
	EMACRO_ARG_COUNT = "MACRO %s の仮引数と引数の個数不一致"
	EMACRO_REDEF     = "%s を MACRO として再定義不可"
	EMACRO_NOT_MACRO = "%s は MACRO 以外で定義済み"
	EMACRO_CYCLIC    = "MACRO %s の展開が再帰"

	ESCOPE        = "%s はこのスコープでは利用できない"
	ESCOPE_GLOBAL = "%s はグローバルスコープで利用不可"
	WSCOPE_MACRO  = "%s は MACRO スコープでは無視"

	E900 = "E900 内部エラー %s"

	ENOT_IMPL_EXPR = "未実装エラー: evalExression(%T)"
	ENOT_IMPL_STMT = "未実装エラー: eval(%T')"

	// warning
	WEXPR_BYTE = "数値 '%d(0x%x)' をバイト範囲に丸めました"
	WEXPR_WORD = "数値 '%d(0x%x)' をワード範囲に丸めました"

	// infromation
	ILABEL_LOCAL = "[I]ローカルラベルには ':' は不要"
)
