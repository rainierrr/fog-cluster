import japanize_matplotlib
import ping3
import fog_client
import matplotlib.pyplot as plt
import numpy as np

x = []

cluster_a_cpu_y = []
cluster_b_cpu_y = []
cluster_a_rtt_y = []
cluster_b_rtt_y = []

SLEEP_TIME = 0.25  # グラフを更新する秒数
FONT_SIZE = 10  # フォントサイズ
X_MAX = 50  # x軸の最大値
count = 0
# figure()でグラフを表示する領域をつくり，figというオブジェクトにする．
fig = plt.figure()

# add_subplot()でグラフを描画する領域を追加する．引数は行，列，場所
cpu_ax = fig.add_subplot(2, 2, 1)
rtt_ax = fig.add_subplot(2, 2, 2)

x_label = "秒数(s)"
cpu_ax.set_xlabel(x_label, fontsize=FONT_SIZE)
cpu_ax.set_ylabel("CPU使用率(％)", fontsize=FONT_SIZE)
rtt_ax.set_xlabel(x_label, fontsize=FONT_SIZE)
rtt_ax.set_ylabel("RTT(s)", fontsize=FONT_SIZE)


mg_node_list = fog_client.setup()

while True:
    # 50個以上のデータがある場合は、古いデータを削除

    # y軸のデータを追加
    for mg_node_dict in mg_node_list:
        cluster_metrics = fog_client.get_cluster_metrics(mg_node_dict['ip'])
        if mg_node_dict['name'] == 'mg-node-a':
            cluster_a_cpu_y.append(cluster_metrics['cpu'])
            cluster_a_rtt_y.append(ping3.ping(mg_node_dict['ip']) * 1000)
        else:
            cluster_b_cpu_y.append(cluster_metrics['cpu'])
            cluster_b_rtt_y.append(ping3.ping(mg_node_dict['ip']) * 1000)

    x.append(count)  # 秒数をx軸に追加

    cluster_a_color, cluster_b_color = "blue", "green"      # 各プロットの色
    cluster_a_label, cluster_b_label = "クラスタA", "クラスタB"  # 各プロットのラベル
    # Plot
    x_length = len(x)

    # 表示するx軸の範囲を設定
    start_index = 0 if x_length < X_MAX else x_length - X_MAX
    cpu_ax.set_xlim(start_index, x_length)
    rtt_ax.set_xlim(start_index, x_length)

    cpu_ax.plot(x, cluster_a_cpu_y,
                color=cluster_a_color, label=cluster_a_label)
    cpu_ax.plot(x, cluster_b_cpu_y, color=cluster_b_color,
                label=cluster_b_label)
    rtt_ax.plot(x, cluster_a_rtt_y,
                color=cluster_a_color, label=cluster_a_label)
    rtt_ax.plot(x, cluster_b_rtt_y,
                color=cluster_b_color, label=cluster_b_label)
    fig.tight_layout()  # グラフのレイアウトを調整
    plt.draw()  # グラフを画面に表示開始
    plt.pause(SLEEP_TIME)  # SLEEP_TIME時間だけ表示を継続
    plt.cla()  # プロットした点を消してグラフを初期化

    count += 1
