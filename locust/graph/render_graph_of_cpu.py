import csv
import numpy as np
import matplotlib.pyplot as plt
import japanize_matplotlib
import datetime

FONT_SIZE = 10
csv_path = './rtt_and_cpu.csv'

rows = []
x = []
cluster_a_y = []
cluster_b_y = []

# csvファイルを読み込む
with open(csv_path) as f:
    reader = csv.reader(f)
    for idx, row in enumerate(reader):
        x.append(int(row[0]))
        cluster_a_y.append(float(row[1]))
        cluster_b_y.append(float(row[2]))


fig = plt.figure()


# ラベルを設定する
plt.xlabel("秒数(s)", size=FONT_SIZE)
plt.ylabel("CPU使用率(%)", size=FONT_SIZE)

# 目盛の設定
# plt.yticks([i for i in range(0, 3000, 50)], size=FONT_SIZE)

# 軸の範囲の設定
plt.xlim(0, max(x))
plt.ylim(0, max(cluster_a_y)*1.1)


plt.grid(which="major", axis="x", color="black", alpha=0.8,
         linestyle="--", linewidth=1)
plt.grid(which="major", axis="y", color="black", alpha=0.8,
         linestyle="--", linewidth=1)

# グラフを描画する
plt.plot(x, cluster_a_y, color='green', label='クラスタA')
plt.plot(x, cluster_b_y, color='red', label='クラスタB')
plt.legend(loc='lower right')
plt.show()
