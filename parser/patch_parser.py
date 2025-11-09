import sys
import re
import typing
import os.path


# 特定誤り防止のため、\t \n で囲まれた str で定義する
class PatchData(typing.NamedTuple):
    old: str
    new: str


patch_data: list[PatchData] = [
    # yyLexer.Error() の引数定義を修正
    PatchData(
        "\tError(s string)\n", 
        "\tError(s string, args ...any) // # changed\n"
    ),

    # ステート遷移のデバッグ表示強化
    PatchData(
        '__yyfmt__.Printf("lex %s(%d)\\n", yyTokname(token), uint(char))'
        ,
        '__yyfmt__.Printf("lex %s(%d) %#v\\n", yyTokname(token), uint(char), lval)'
    ),
    PatchData(
        "var yyn int", 
        "var yyn, yySave int // # changed"
    ),
    # PatchData(
    #     '__yyfmt__.Printf("char %v in %v\\n", yyTokname(yytoken), yyStatname(yystate))',
    #     '__yyfmt__.Printf("# %d: char %v (%#v)\\n", yystate, yyTokname(yytoken), yyVAL) // # changed'
    # ),
    PatchData(
        "\t\t" "if Errflag > 0 {\n"
        "\t\t" "\tErrflag--\n"
		"\t\t"  "}\n"
		"\t\t"  "goto yystack\n"
        ,
        "\t\t" "if Errflag > 0 {\n"
        "\t\t" "\tErrflag--\n"
		"\t\t"  "}\n"
		"\t\t"  '__yyfmt__.Printf("# %d: shift %d\\n", yySave, yystate) // # added\n'
		"\t\t"  "goto yystack\n"
    ),
    PatchData(
        "yyn = int(yyAct[yyn])",
        "yyn = int(yyAct[yyn])\n\tyySave = yystate // # added",
    ),
    PatchData(
        '__yyfmt__.Printf("reduce %v in:\\n\\t%v\\n", yyn, yyStatname(yystate))',
        'rule, ok := grammerRules[yyn]\n'
        '\t\tif !ok {\n'
        '\t\t\trule = ""\n'
        '\t\t}\n'
        '\t\t__yyfmt__.Printf("# %d: reduce (%d) %s\\n", yystate, yyn, rule) // # changed',
    ),
    PatchData(
        'yyVAL = yyS[yyp+1]\n'
        ,
        "yyVAL = yyS[yyp+1]\n"
        "\t" '__yyfmt__.Printf("# top %d\\n", yyS[yyp].yys) // # added\n'
        "\t" "yySave = yyS[yyp].yys  // # added"
    ),
    PatchData(
        '\t// dummy call; replaced with literal code\n'
        ,
        '\t__yyfmt__.Printf("# %d: goto %d\\n", yySave, yystate) // # added\n\n'
        '\t// dummy call; replaced with literal code\n',
    ),
    PatchData(
        'return "syntax error"',
        'return __yyfmt__.Sprintf("syntax error(state %d)", state)'
    ),
    PatchData(
        'return "syntax error: " + e.msg',
        'return __yyfmt__.Sprintf("syntax error(state %d): %s", state, e.msg)'
    ),
    PatchData(
        'res := "syntax error: unexpected " + yyTokname(lookAhead)',
        'res := __yyfmt__.Sprintf("syntax error(state %d): unexpected %s(%d)", state,  yyTokname(lookAhead), lookAhead)',
    ),
]


def extract_rules(infile):
    rules = {}
    if infile == "":
        return rules
    with open(infile, encoding="utf-8") as fp:
        for line in fp:
            m = re.match(r"(.+)\s+\((\d+)\)$", line.strip())
            if m is None:
                continue
            num = int(m.group(2))
            rule = m.group(1).strip()
            rules[num] = rule
    if len(rules) == 0:
        print(f"no rulesin {infile}")
        return

    outfile = os.path.join(os.path.dirname(infile), "rules.go")
    with open(outfile, "w", encoding="utf-8") as fp:
        print("package parser", file=fp)
        print("", file=fp)
        print("var grammerRules map[int]string = map[int]string{", file=fp)
        for k, v in rules.items():
            print(f'\t{k}: "{v}",', file=fp)
        print("}", file=fp)

def main(infile, outfile):
    with open(infile, encoding="utf8") as f:
        text = f.read()
    for p in patch_data:
        new_text = text.replace(p.old, p.new, 1)
        if new_text == text:
            print("patch faild")
            print("new")
            print(p.new)
        text = new_text
    with open(outfile, "w", encoding="utf-8", newline="\n") as f:
        f.write(text)

if __name__ == "__main__":
    output = ""
    if len(sys.argv) == 3:
        infile, outfile = sys.argv[1:]
    elif len(sys.argv) == 4:
        infile, outfile, output = sys.argv[1:]
    else:
        exit(1)

    rules = extract_rules(output)
    main(infile, outfile)
