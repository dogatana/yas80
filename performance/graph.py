import matplotlib.pyplot as plt
import japanize_matplotlib

# labels = ['z80as', 'yas80', 'z80asm(z88dk)', 'tools80', 'ailz80asm']
# men_means = [0.004, 0.064, 0.264, 2.443, 6.740]
# values = [0.016, 0.044, 0.152, 0.395, 1.065]

# データの生成
categories = ['z80as', 'yas80', 'z80asm(z88dk)', 'zma', 'tools80', 'ailz80asm']
values = [0.004, 0.064, 0.264, 1.003, 2.443, 6.740]
# values = [0.016, 0.044, 0.152, 0.395, 1.065]

# 横棒グラフの描画
plt.barh(categories, values, 0.4) # , color='skyblue')


# 各棒の右側に値を表示
for i in range(len(values)):
    plt.text(values[i] * 1.01, i, str(values[i]), va='center')

# グラフのタイトルとラベル
plt.title('Assembler Perfromance')
plt.xlabel('time(sec)')
# plt.ylabel('Categories')

# グリッドを表示
# plt.grid(axis='x')

# グラフを表示
plt.show()