import re
import matplotlib
import sys

matplotlib.use("Agg")  # no GUI

import matplotlib.pyplot as plt
import japanize_matplotlib  # noqa: F401


labels = ["yas80", "z80as", "z80asm", "tools80", "ailz80asm"]
colors = [
    "darkorange",
    "cornflowerblue",
    "cornflowerblue",
    "cornflowerblue",
    "cornflowerblue",
]

def main(file:str):
    result = read_data(file)
    result_to_graph(result)
    for file, v in result.items():
        print(file)
        for asm, value in v.items():
            print(f"{asm:20s}{value}")
        print()



def read_data(file):
    result = {}
    with open(file) as f:
        count = 0
        asm = file = ""
        for line in f:
            m = re.search(r'measure.py (\d+) "([^"]+)"\s*$', line)
            if m is not None:
                cols = m.group(2).split()
                count = int(m.group(1))
                asm = cols[0]
                file = cols[-1]
                continue
            m = re.match(r"total (\S+).*?(\S+)$", line)
            if m is not None:
                total = float(m.group(1))
                ave = float(m.group(2))

                v = result.setdefault(file, dict())
                v.setdefault(asm, ave)
    return result

def result_to_graph(result):
    for file, v in result.items():
        values = list(v.values())
        # values.reverse()
        build_graph(file, values)

def build_graph(file, values):
    fig = plt.figure(figsize=(9, 4), dpi=100)  # 1000 x 400
    plt.barh(labels, values, 0.4, color=colors)
    plt.gca().invert_yaxis()

    # 右側に値を表示
    for i in range(len(values)):
        plt.text(values[i] * 1.02, i, str(values[i]), va="center")

    # グラフのタイトルとラベル
    plt.title("平均実行時間")
    plt.xlabel("time(sec)")
    # plt.ylabel("Categories")

    # グリッドを表示
    # plt.grid(axis="x")

    plt.savefig(f"{file}.svg")

if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("usage: python txt2graph.py text-file")
        exit(1)
    main(sys.argv[1])