import sys
import typing


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
