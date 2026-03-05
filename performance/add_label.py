"""
max.asm の各命令にラベル追加
結果を label.asm として出力
"""

import re


def main():
    with open("max.asm") as fi:
        with open("label.asm", "w") as fo:
            n = 0
            for line in fi:
                line = line.strip()
                if line == "":
                    continue
                print(f"label_{n}: {line}", file=fo)
                n += 1

if __name__ == "__main__":
    main()
