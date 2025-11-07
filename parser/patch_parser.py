import sys
import typing


# 特定誤り防止のため、\t \n で囲まれた str で定義する
class PatchData(typing.NamedTuple):
    old: str
    new: str


patch_data: list[PatchData] = [
    # yyLexer.Error() の引数定義を修正
    PatchData("\tError(s string)\n", "\tError(s string, args ...any)\n"),
    # yySymName() で 0x80 以上を数字でなく文字として扱う修正
    PatchData(
        '\treturn __yyfmt__.Sprintf("%d", c)\n',
        '\treturn __yyfmt__.Sprintf("%q", rune(c))\n',
    ),
    # yyParse 内部のエラー発生時にトークン情報を追加
    PatchData(
        '\tyylex.Error(msg)\n',
        '\tyylex.Error(msg, &yyVAL.token, &yylval.token)\n',

    ),
    # PatchData(
    #     '\tmsg = __yyfmt__.Sprintf("unexpected %s", ls)\n',
    #     '\tmsg = __yyfmt__.Sprintf("予期しないトークン %s\\n  yyLVAR.token: %s\\n  yylvar.token: %s", ls, yyVAL.token.String(), yylval.token.String())\n',
    # ),
    # PatchData(
    #     '\tmsg = __yyfmt__.Sprintf("unexpected %s, %s", ls, msg)\n',
    #     '\tmsg = __yyfmt__.Sprintf("予期しないトークン %s, %s\\n  yyLVAR: %s\\n  yylvar: %s", ls, msg, yyVAL.token.String(), yylval.token.String())\n',
    # ),
]


def main(infile, outfile):
    with open(infile, encoding="utf8") as f:
        text = f.read()
    for p in patch_data:
        text = text.replace(p.old, p.new, 1)
    with open(outfile, "w", encoding="utf-8", newline="\n") as f:
        f.write(text)


if __name__ == "__main__":
    if len(sys.argv) != 3:
        exit(1)
    infile, outfile = sys.argv[1:]
    main(infile, outfile)
