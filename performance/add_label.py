"""
zilog.asm x 45, ixy-hl * 5 から命令抽出
各命令にラベル追加
結果を label.asm として出力
"""

import re


def main():
    zilog = read_file("zilog.asm")
    ixy = read_file("ixy-hl.asm")

    lines = []
    n = 0
    for _ in range(45):
        for line in zilog:
            lines.append(f"label_{n}: {line}")
            n += 1
    for _ in range(5):
        for line in ixy:
            lines.append(f"label_{n}: {line}")
            n += 1
    open("label.asm", "w", encoding="utf-8").writelines(lines)


def read_file(name):
    out = []
    with open(name, encoding="utf-8") as fp:
        for line in fp:
            line = re.sub(r";.*$", "", line)
            line = re.sub(r"^\s+", "", line)
            if line == "":
                continue
            out.append(line)
    return out


if __name__ == "__main__":
    main()
