import csv
import numpy as np
import matplotlib.pyplot as plt
import japanize_matplotlib
import datetime

FONT_SIZE = 10
csv_path = './rps.csv'

rows = []
colums = ['Time', ' Req/s', ' active_users',
          ' throughput mg-node-a', ' throughput mg-node-b']
x = []
rps_y = []
users_y = []
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
        x.append(time.seconds)
        rps_y.append(float(row[1]))
        users_y.append(float(row[2]))
        cluster_a_y.append(float(row[3]))
        cluster_b_y.append(float(row[4]))


fig = plt.figure()
plt2 = plt.twinx()


# ラベルを設定する
plt.xlabel("秒数(s)", size=FONT_SIZE)
plt2.set_ylabel("ユーザー数", size=FONT_SIZE)
plt.ylabel("Req/s", size=FONT_SIZE)

# 目盛の設定
# plt.yticks([i for i in range(0, 3000, 50)], size=FONT_SIZE)

# 軸の範囲の設定
print(rps_y)
plt.xlim(0, max(x))
plt.ylim(0, max(rps_y)*1.1)


plt.grid(which="major", axis="x", color="black", alpha=0.8,
         linestyle="--", linewidth=1)
# plt.grid(which="major", axis="y", color="black", alpha=0.8,
#          linestyle="--", linewidth=1)

# グラフを描画する
plt.plot(x, rps_y, color='blue', label='Req/s')
plt.plot(x, cluster_a_y, color='green', label='cluster a Req/s')
plt.plot(x, cluster_b_y, color='red', label='cluster b Req/s')
plt2.plot(x, users_y, color='orange', label='active users')
plt.show()
