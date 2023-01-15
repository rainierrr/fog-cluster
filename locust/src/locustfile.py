from locust import HttpUser, task, constant_throughput
import fog_client
import time
import random

# import locust_plugins
TIMEOUT = 3
post_data = {"temperature": 20, "humidity": 0.5,
             "pressure": 1000, "location": "Tokyo"}


class QuickstartUser(HttpUser):
    wait_time = constant_throughput(10)

    # ランダムにクラスタを選択する
    # def on_start(self):
    #     node_a = {
    #         "name": "cluster-a",
    #         "ip": fog_client.MG_NODE_A_IP
    #     }
    #     node_b = {
    #         "name": "cluster-b",
    #         "ip": fog_client.MG_NODE_B_IP
    #     }
    #     cluster_list = [node_a, node_b]
    #     self.target_node = random.choice(cluster_list)

    def on_start(self):
        self.mg_node_list = fog_client.setup()
        selected_node = fog_client.find_node(self.mg_node_list)
        self.target_node = selected_node
        self.start_time = time.time()

    # @task
    # def post_to_node_a(self):
    #     url = f'http://{fog_client.MG_NODE_A_IP}:{fog_client.FOG_APP_PORT}/post'
    #     self.client.post(url, json=post_data,
    #                      name='cluster-a', timeout=TIMEOUT)

    @task
    def select_post(self):
        if time.time() - self.start_time > 15:
            selected_node = fog_client.find_node(self.mg_node_list)
            self.target_node_ip = selected_node['ip']
            self.start_time = time.time()

        url = f'http://{self.target_node["ip"]}:{fog_client.FOG_APP_PORT}/post'
        self.client.post(
            url, json=post_data, name=self.target_node['name'], timeout=TIMEOUT)
