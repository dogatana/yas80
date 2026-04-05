package errcode

const (
	// system
	ESYSTEM   = "system error"
	ESYNTAX   = "syntax error"
	EINTERNAL = "internal error %s"

	EFILE_ENCODING  = "ファイルのエンコーディングが不明です"
	EFILE_NOT_FOUND = "ファイル %s が見つかりません"
	EFILE_ERR       = "ファイル %s エラー: %s"

	EINCLUDE_CYCLIC = "INCLUDE %q で循環 INCLUDE 検出"

	// 定義
	ECHARMAP_DUP             = "CHARMAP %s は定義済み"
	ECHARMAP_USED            = "%s を CHARMAP として再定義不可"
	ECHARMAP_NULL            = "JSON ファイル名が未確定"
	ECHARMAP_NOT_STR         = "JSON ファイル名が文字列でない"
	ECHARMAP_READ            = "%s ファイル読み込みエラー(%s)"
	ECHARMAP_JSON            = "JSON フォーマットが期待どおりでない"
	ECHARMAP_FMT             = "CHARMAP 定義エラー (%q)"
	ECHARMAP_DEFCHAR_NULL    = "デフォルトコードが未確定"
	ECHARMAP_DEFCHAR_NOT_INT = "デフォルトコードが数値でない"
	// 適用
	ECHARMAP_VALUE_NULL = "CHARMAP の適用対象が未確定"
	ECHARMAP_VALUE      = "CHARMAP を適用できない値"
	ECHARMAP_NOT_DEF    = "CHARMAP に文字 '%c' の定義がない"

	ESETMAP_NOT_A_CHAR  = "文字列の文字数が 1 でない"
	ESETMAP_ARRAY_EMPTY = "空の配列が指定された"

	EADDRESS_OVERFLOW = "アドレスオバーフロー $%x"

	ESTR_TO_INT_LEN = "文字数が %d のため数値に変換できない"
	ESTR_CTRL       = "文字列リテラルに制御文字を含めることは不可"
	ESTR_END_QUOTE  = "文字列リテラル終端の引用符なし"

	ENUMBER_LITERAL = "数値リテラル誤り: '%s'"

	EARRAY_TO_INT_LEN   = "配列要素数が %d のため数値に変換できない"
	EARRAY_TO_INT_TYPE  = "配列要素が数値でないため、数値に変換できない"
	EARRAY_EMPTY        = "空の配列"
	EARRAY_NAME         = "配列名誤り"
	EARRAY_INDEX        = "配列インデックス誤り"
	EARRAY_OUT_OF_INDEX = "配列インデックスが範囲外"

	EUNI_OP_TYPE = "単項演算子 '%s' は使用できない"

	EBIN_OP_DIVZERO = "0 除算"
	EBIN_OP_TYPE    = "%s が使用できない値"

	// データ定義
	// EDATA_EMPTY  = "DB/DW/DD に値が指定されていない"
	EDATA_DW_STR = "DW では文字列は指定できない"
	EDATA_ENCODE = "文字列 %q を SHIFT-JIS へ変換できない"
	EDATA_VALUE  = "データ定義(DB/DW/DD)に使用できない値"

	EDS_COUNT = "DS/DSB/DSW のデータ数が未確定"
	EDS_FILL  = "DS/DSB/DSW の埋め込みデータが未確定"

	// レジスタ間接
	EINDIRECT_VALUE      = "(数値) が必要"
	EINDIRECT_NULL       = "(数値) が未確定"
	EINDIRECT_OP         = "レジスタ間接に使用できない演算子"
	EINDIRECT_REG        = "(%s) は使用できない"
	EINDIRECT_DISP_REG   = "オフセットは IX/IY のみ指定可能"
	EINDIRECT_DISP       = "IX/IY のオフセットが非数値"
	EINDIRECT_DISP_RANGE = "IX/IY のオフセットが範囲外 %d(0x%x)"
	EINDIRECT_DISP_NULL  = "IX/IY のオフセットが未確定"

	// Z80
	EZ80_NOT_IMPL         = "評価未実装 %s"
	EZ80_FLAG             = "フラグ未指定"
	EZ80_OP               = "命令オペランドエラー"
	EZ80_OP_NULL          = "命令オペランドが未確定"
	EZ80_OP_REG           = "レジスタ %s は指定不可"
	EZ80_OP1              = "第1オペランドエラー"
	EZ80_OP1_NULL         = "第1オペランドが未確定"
	EZ80_OP1_SP           = "第1オペランドは SP 以外指定不可"
	EZ80_OP1_REG_A        = "第1オペランドは A 以外指定不可（省略可能）"
	EZ80_OP1_REG_HL       = "第1オペランドは HL 以外指定不可"
	EZ80_OP1_REG_HL_IXY   = "第2オペランドは HL/IX/IY 以外指定不可"
	EZ80_OP2              = "第2オペランドエラー"
	EZ80_OP2_NULL         = "第2オペランドが未確定"
	EZ80_OP2_HL_IXY       = "第2オペランドは HL, IX, IY 以外指定不可"
	EZ80_OP2_STR          = "第2オペランドが文字列の場合 1/2 バイトの場合のみ有効"
	EZ80_RST              = "RST のアドレスは 0, 8, 10H, 18H, 20H, 28H, 30H, 38H のいずれかでなければならない (指定値: %d(0x%x))"
	EZ80_JR_RANGE         = "相対ジャンプ先が範囲外 %d(0x%x)"
	EZ80_JR_FLAG          = "JR に '%s' フラグは使用不可"
	EZ80_JP_INDIRECT_DISP = "間接 JP にオフセット指定は不可"
	EZ80_JP_INDIRECT_REG  = "(HL), (IX), (IY) のみ指定可能"
	EZ80_BIT_NUM_RANGE    = "BIT番号(0-7)が範囲外"
	EZ80_PORT_RANGE       = "ポート番号(0-255)が範囲外 %d(0x%x)"
	EZ80_IM_RANGE         = "0/1/2 のみ有効 %d(0x%x)"

	ER800 = "R800 専用命令のため利用不可"

	// 匿名ラベル
	EANON_LABEL_NOT_FOUND = "対応する @@ が定義されていない"
	EANON_LABEL_REF_ONLY  = "%s は参照のみ可能"

	// symbol
	ESYM_UNDEF  = "シンボル %s は未定義"
	ESYM_CYCLIC = "シンボル %s の定義が循環参照を含む"
	ESYM_NULL   = "シンボル %s の値が未確定"

	// ORG
	EORG_ALLOC = "ABS/REL のみ指定可能"
	EORG_NULL  = "アドレスが未確定"
	EORG_VALUE = "アドレス指定誤り"

	// END
	EEND_NULL  = "開始アドレスが未確定"
	EEND_VALUE = "開始アドレス指定誤り"

	// CONST/EQU
	ECONST_DUP  = "CONST/EQU '%s' は定義済み"
	ECONST_USED = "%s を CONST/EQU として再定義不可"
	ECONST_NULL = "%s の値が未確定"

	// ##
	// ECONCAT_NOTSYM = "シンボル結合にはシンボルが必要"
	// ECONCAT_EXPR   = "シンボル結合式の誤り"
	ECONCAT_TYPE = "シンボル結合は数値、文字列のみ可能"

	// ラベル
	ELABEL_DUP  = "LABEL %s は定義済み"        // symbol 以外で利用済
	ELABEL_USED = "%s は利用済のため LABEL 定義不可" // symbol 以外で利用済
	// ELABEL_EXPR = "LABEL としてシンボルが必要"

	// VAR/ASSIGN
	EVAR_UNDEF    = "変数 %s は未定義"
	EVAR_SYS      = "_ は再定義不可"
	EVAR_USED     = "%s を変数として再定義不可"
	EVAR_VALUE    = "変数 %s の初期値が未確定"
	EASSIGN_LEFT  = "変数/_ 以外へは代入不可"
	EASSIGN_VALUE = "代入する値が未確定"
	// EASSIGN_UNDEF         = "変数 %s は未定義"

	// PROC 定義
	EPROC_DUP  = "PROC '%s' は定義済み"
	EPROC_USED = "%s を PROC として再定義不可"
	EPROC_NEST = "PROC 定義はネスト不可"

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
	EFUNC_NOT_FUNC  = "%s は FUNC/CHARMAP でない"
	EFUNC_UNDEF     = "FUNC/CHARMAP %s は未定義"
	EFUNC_ARG_COUNT = "FUNC '%s' の仮引数の数と関数呼出しの実引数の数が不一致"

	// 組み込み関数
	EEBFN_NOT_FOUND = "組み込み関数 %s は未定義"
	EEBFN_ARG_COUNT = "組み込み関数 %s の引数の数に誤り"
	EEBFN_ARG_VALUE = "組み込み関数 %s の引数誤り"
	EEBFN_ARG_NULL  = "組み込み関数 %s の引数が未確定"

	// 組み込みマクロ
	EEBMAC_ARG_COUNT = "%s の引数の数誤り"
	EEBMAC_ARG_VALUE = "%s の引数誤り"
	EEBMAC_ARG_NULL  = "%s の引数が未確定"

	// MACRO 定義
	EMACRO_NAME = "%s は MACRO 名として使用不可"
	EMACRO_USED = "%s を MACRO として再定義不可"
	EMACRO_DUP  = "MACRO %s は定義済み" // TODO 最終的に不要
	EMACRO_NEST = "MACRO 定義はネスト不可" // TODO 最終的に不要
	// MACRO 呼出し
	EMACRO_NOT_MACRO = "%s は MACRO でない"
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

	// BinWriter
	EBW_OVERLAPPED = "ORG $%x と ORG $%x の範囲が重複"
	EBW_WRITE      = "ファイル書込みエラー: %s"

	// warning
	WROUND_BYTE = "数値 '%d(0x%x)' をバイト範囲(-128 - 255)に丸めた"
	WROUND_WORD = "数値 '%d(0x%x)' をワード範囲(-32768 - 65535)に丸めた"
	WROUND_ADDR = "数値 '%d(0x%x)' をアドレス範囲(0 - 5535)に丸めた"

	WSCOPE_MACRO = "MACRO 内では無効"
	WSCOPE_FUNC  = "FUNC 内では無効"
	WSCOPE_PROC  = "PROC では無効"

	WBW_LOAD_NAME = "Load name が ASCII 以外の文字を含んでいるので OUTPUT で代替"
	WBW_NO_CODE   = "コード生成なし"

	WEVAL_CODE_STABLE = "生成コードが不安定。出力コードを確認すること"
)
