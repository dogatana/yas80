package errcode

const (
	// system
	ESYSTEM   = "system error"
	ESYNTAX   = "syntax error"
	EINTERNAL = "internal error %s"

	ESTR_CTRL      = "文字列リテラルに制御文字を含めることは不可"
	ESTR_END_QUOTE = "文字列リテラル終端の引用符なし"

	ENUMBER = "数値リテラル誤り: '%s'"

	EARRAY_NAME  = "配列名誤り"
	EARRAY_INDEX = "配列インデックス誤り"

	E006 = "間接指定オペランド誤り"
	E009 = "'%s' は未定義"
	E030 = "'%s' の定義に循環参照あり"
	E900 = "E900 内部エラー %s"

	EUNI_OP_STRING = "'%c' は単項文字列演算子ではない"
	EUNI_OP_NUMBER = "'%c' は単項数値演算子ではない"
	EUNI_OP_TYPE   = "単項演算子 '%s' が使用できない型"

	EBIN_OP_DIVZERO = "0 除算"
	EBIN_OP_TYPE    = "二項演算子 '%s' は使用不可"

	// データ定義
	EDATA_EMPTY  = "DB/DW/DD に値が指定されていない"
	EDATA_DW_STR = "DW では文字列は指定できない"
	EDATA_ENCODE = "文字列 %q を SHIFT-JIS へ変換できない"

	EDS_COUNT = "DS/DSB/DSW のデータ数が確定できない"
	EDS_FILL  = "DS/DSB/DSW の埋め込みデータが確定できない"

	// Z80
	EZ80_FLAG       = "'%s' はフラグでない"
	EZ80_OP1        = "第1オペランド誤り"
	EZ80_OP2        = "第2オペランド誤り"
	EZ80_OP1_SP     = "第1オペランドは SP 以外指定不可"
	EZ80_OP2_HL_IXY = "第2オペランドは HL, IX, IY 以外指定不可"

	// symbol
	ESYM_UNDEF  = "シンボル %s は未定義"
	ESYM_CYCLIC = "シンボル %s の定義が循環参照を含む"
	ESYM_NULL   = "シンボル %s の値を確定できない"

	// CONST/EQU
	ECONST_DUP  = "CONST/EQU '%s' は定義済み"
	ECONST_USED = "%s を CONST/EQU として再定義不可"

	// ##
	ECONCAT_NOTSYM = "シンボル結合にはシンボルが必要"
	ECONCAT_EXPR   = "シンボル結合式の誤り"
	ECONCAT_TYPE   = "シンボル結合は数値、文字列のみ可能"

	// ラベル
	ELABEL_DUP  = "LABEL %s は定義済み"        // symbol 以外で利用済
	ELABEL_USED = "%s は利用済のため LABEL 定義不可" // symbol 以外で利用済
	ELABEL_EXPR = "LABEL としてシンボルが必要"

	// VAR/ASSIGN
	EVAR_UNDEF    = "変数 %s は未定義"
	EVAR_SYS      = "_ は再定義不可"
	EVAR_USED     = "%s を変数として再定義不可"
	EVAR_VALUE    = "変数 %s の初期値が未確定"
	EASSIGN_LEFT  = "変数/_ 以外へは代入不可"
	EASSIGN_VALUE = "代入文の値が未確定"
	// EASSIGN_UNDEF         = "変数 %s は未定義"

	// PROC 定義
	EPROC_DUP  = "PROC '%s' は定義済み"
	EPROC_USED = "%s を PROC として再定義不可"

	// ENUM 定義
	EENUM_DUP       = "ENUM '%s' は定義済み"
	EENUM_USED      = "'%s' を ENUM として再定義不可"
	EENUM_ELE_DUP   = "ENUM '%s.%s' は定義済みのため無効"
	EENUM_ELE_VALUE = "ENUM 要素に使用できない値"
	EENUM_ELE_FWD   = "ENUM 要素の値が未確定"
	// ENUM 参照
	EENUM_ELE_UNDEF = "ENUM '%s' は未定義"

	// FUNC 定義
	EFUNC_NAME = "%s は関数名として使用不可"
	EFUNC_USED = "%s を FUNC として再定義不可"
	EFUNC_DUP  = "FUNC '%s' は定義済み"
	// FUNC 呼出し
	EFUNC_NOT_FUNC  = "%s は関数名ではない"
	EFUNC_UNDEF     = "FUNC '%s' は未定義"
	EFUNC_ARG_COUNT = "FUNC '%s' の仮引数の数と関数呼出しの実引数の数が不一致"

	// MACRO 定義
	EMACRO_NAME = "%s は MACRO 名として使用不可"
	EMACRO_USED = "%s を MACRO として再定義不可"
	EMACRO_DUP  = "MACRO %s は定義済み" // TODO 最終的に不要
	EMACRO_NEST = "MACRO 定義はネスト不可" // TODO 最終的に不要
	// MACRO 呼出し
	EMACRO_UNDEF     = "MACRO %s は未定義"
	EMACRO_ARG_COUNT = "MACRO %s の仮引数と引数の個数不一致"
	EMACRO_CYCLIC    = "MACRO %s の展開が再帰"

	// REPT
	EREPT_COUNT = "REPT 式の値が数値でない"

	// SCOPE
	ESCOPE_PROC  = "%s は PROC 内部でしか利用できない"
	ESCOPE_MACRO = "%s は MACRO 内部でしか利用できない"

	ENOT_IMPL_EXPR = "未実装エラー: evalExression(%T)"
	ENOT_IMPL_STMT = "未実装エラー: eval(%T')"

	// warning
	WROUND_BYTE  = "数値 '%d(0x%x)' をバイト範囲に丸めました"
	WROUND_WORD  = "数値 '%d(0x%x)' をワード範囲に丸めました"
	WSCOPE_MACRO = "MACRO 内では無効"
	WSCOPE_FUNC  = "FUNC 内では無効 %T"

	// infromation
	ILABEL_LOCAL = "[I]ローカルラベルには ':' は不要"
)
