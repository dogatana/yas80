"""
アセンブル時間の計測
count 1 - 100
cmd コマンドライン文字列(単一文字列)
結果として、総実行時間（秒）、1回あたりの平均実行時間（秒）を表示

実行例
> py measure.py 100 "z80asm -b big.asm"
total 15.411 sec, average : 0.154

"""

import subprocess
import sys
import time


def main(count, cmd):
    result = []
    for _ in range(count):
        start = time.perf_counter()
        subprocess.run(cmd, stdout=subprocess.PIPE, stderr=subprocess.PIPE)
        result.append(time.perf_counter() - start)

    total = sum(result)
    print(f"total {total:.3f} sec, average : {total / len(result):.3f}")


if __name__ == "__main__":
    if len(sys.argv) != 3:
        print("usage: python measure.py count cmd")
        exit()
    count = int(sys.argv[1])
    if count <= 0:
        print("count must be positive number")
        exit()
    main(count, sys.argv[2])
