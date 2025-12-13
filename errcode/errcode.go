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

	EUNARY_OP_STRING = "'%c' は単項文字列演算子ではない"
	EUNARY_OP_NUMBER = "'%c' は単項数値演算子ではない"
	EUNARY_OP_TYPE   = "単項演算子 '%s' が使用できない型"

	E009 = "'%s' は未定義"
	E010 = "enum '%s' は定義されていない"
	E011 = "enum '%s.%s' は定義されていない"
	E012 = "enum '%s' は定義済み"
	E013 = "enum '%s.%s' は定義済み"
	E014 = "enum 要素に使用できない型 %T"

	EBIN_OP_DIVZERO = "0 除算"
	EBIN_OP_STRING  = "文字列は '+' 演算子のみ使用可能"
	EBIN_OP_NUMBER  = "不明な整数演算子 '%s'"
	EBIN_OP_TYPE    = "二項演算子 '%s' は使用不可"

	E017 = "'%s' はフラグでない"
	E018 = "func '%s' は定義済み"

	E019 = "関数呼出しには関数名が必要"
	E020 = "func '%s' は未定義"

	EFUNC_ARG_COUNT = "FUNC '%s' の仮引数の数と関数呼出しの実引数の数が不一致"

	E024 = "第1オペランド誤り"
	E025 = "第2オペランド誤り"
	E026 = "第1オペランドは SP 以外指定不可"
	E027 = "第2オペランドは HL, IX, IY 以外指定不可"
	E028 = "'%s' は指定不可"
	E030 = "'%s' の定義に循環参照あり"

	ESYM_DUP            = "CONST/EQU '%s' は定義済み"
	ESYM_USED_NAME      = "%s は利用済のため CONST/EQU 定義不可"
	ESYM_NOT_FOUND      = "シンボル %s は未定義"
	ESYM_CYCLIC         = "シンボル %s の定義が循環参照を含む"
	ESYM_NOT_DETERMINED = "シンボル %s の値を確定できない"

	ELABEL_DUP       = "LABEL %s は定義済み"        // symbol 以外で利用済
	ELABEL_USED_NAME = "%s は利用済のため LABEL 定義不可" // symbol 以外で利用済

	EMACRO_FUNC_NOT_FOUND = "MACRO/FUNCTION %s は未定義"
	EMACRO_NOT_FOUND      = "MACRO %s は未定義"
	EMACRO_DUP            = "MACRO %s は定義済み" // TODO 最終的に不要
	EMACRO_NEST           = "MACRO 定義はネスト不可" // TODO 最終的に不要
	EMACRO_ARG_COUNT      = "MACRO %s の仮引数と引数の個数不一致"
	EMACRO_DEF            = "%s を MACRO として再定義不可"
	EMACRO_NOT_MACRO      = "%s は MACRO 以外で定義済み"

	EGLOBAL_NOT_ALLOWED = "%s はグローバルスコープで利用不可"
	WMACRO_NOT_ALLOWED  = "%s は MACRO スコープでは無視"

	E900           = "E900 内部エラー %s"
	ENOT_IMPL_EXPR = "未実装エラー: evalExression(%T)"
	ENOT_IMPL_STMT = "未実装エラー: eval(%T')"

	// warning
	W001 = "数値 '%d(0x%x)' をバイト範囲に丸めました"
	W002 = "数値 '%d(0x%x)' をワード範囲に丸めました"

	// infromation
	I001 = "[I]ローカルラベルには ':' は不要"
)
