import matplotlib

matplotlib.use("Agg")  # no GUI

import matplotlib.pyplot as plt
import japanize_matplotlib  # noqa: F401

labels = ["yas80", "z80as", "z80asm", "tools80", "ailz80asm"]
data = {
    "min": [
        0.008,
        0.004,
        0.010,
        0.105,
        0.485,
    ],
    "max": [
        0.041,
        0.014,
        0.142,
        0.308,
        0.984,
    ],
    "label": [
        0.062,
        0.025,
        0.275,
        2.496,
        6.853,
    ],
}

colors = [
    "darkorange",
    "cornflowerblue",
    "cornflowerblue",
    "cornflowerblue",
    "cornflowerblue",
]
for src, values in data.items():
    fig = plt.figure(figsize=(9, 3), dpi=100)  # 1000 x 400
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

    plt.savefig(f"{src}.png")

    # グラフを表示
    # plt.show()
