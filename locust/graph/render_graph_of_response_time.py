import csv
import numpy as np
import matplotlib.pyplot as plt
import japanize_matplotlib
import datetime

FONT_SIZE = 10
csv_path = './rps.csv'

rows = []
x = []
cluster_a_y = []
cluster_b_y = []

# csvファイルを読み込む
with open(csv_path) as f:
    reader = csv.reader(f)
    for idx, row in enumerate(reader):
        if idx == 0:
            continue
        if idx == 1:
            start_time = datetime.datetime.strptime(
                row[0], '%Y-%m-%d %H:%M:%S')
        if len(row) < 4:
            raise Exception('Invalid csv file')

        time = datetime.datetime.strptime(
            row[0], '%Y-%m-%d %H:%M:%S') - start_time
        if time.seconds > 120:
            break
        x.append(time.seconds)
        cluster_a_y.append(float(row[1]))
        cluster_b_y.append(float(row[2]))


fig = plt.figure()
fig = fig.add_subplot(111)

# ラベルを設定する
fig.set_xlabel("秒数(s)", size=FONT_SIZE)
fig.set_ylabel("Respons Time(ms)", size=FONT_SIZE)

# 目盛の設定

# 軸の範囲の設定
# plt.xlim(0, max(x))
# plt.ylim(0, max(rps_y)*1.1)


plt.grid(which="major", axis="x", color="black", alpha=0.8,
         linestyle="--", linewidth=1)
plt.grid(which="major", axis="y", color="black", alpha=0.8,
         linestyle="--", linewidth=1)

# グラフを描画する
fig.plot(x, cluster_a_y, color='green', label='クラスターA')
fig.plot(x, cluster_b_y, color='red', label='クラスターB')
# ax2.plot(x, users_y, color='orange', label='ユーザー数')
plt.legend(loc='upper left')
plt.show()
