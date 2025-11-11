package logger

const (
	// system
	E000 = "[E]system error"
	E001 = "[E]syntax error"

	// user
	E002 = "[E]数値リテラル誤り: '%s'"
	E003 = "[E]配列名誤り"
	E004 = "[E]配列インデックス未指定"
	E005 = "[E]配列インデックス誤り"
	E006 = "[E]間接指定オペランド誤り"
	E007 = "[E]単項演算子 '%c' は文字列に利用不可"

	I001 = "[I]ローカルラベルには ':' は不要"
)
