from locust import HttpUser, task, constant_throughput
import fog_client
import time
import locust_plugins

TIMEOUT = 10
post_data = {"temperature": 20, "humidity": 0.5,
             "pressure": 1000, "location": "Tokyo"}


class QuickstartUser(HttpUser):
    # wait_time = constant_throughput(10)

    def on_start(self):
        self.mg_node_list = fog_client.setup()
        selected_node = fog_client.find_node(self.mg_node_list)
        self.target_node_ip = selected_node['ip']
        self.start_time = time.time()

    @task
    def get_contents(self):
        if time.time() - self.start_time > 3:
            selected_node = fog_client.find_node(self.mg_node_list)
            self.target_node_ip = selected_node['ip']
            self.start_time = time.time()
        url = f'http://{self.target_node_ip}:{fog_client.FOG_APP_PORT}/post'
        self.client.post(
            url, json=post_data, name=url, timeout=TIMEOUT)
