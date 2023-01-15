import httpx
import ping3
import time

SUPER_NODE_IP = '52.199.79.207'
MG_NODE_A_IP = '3.113.76.40'
MG_NODE_B_IP = '54.255.41.147'
SUPER_APP_PORT = 3000
MG_APP_PORT = 30002
FOG_APP_PORT = 30003
CPU_THRESHOLD = 0.8


def get_request(url):
    r = httpx.get(url)

    if r.status_code != 200:
        raise Exception(f'Get Request Failed url: {url}')

    return r.json()


def post_request(url, data):
    r = httpx.post(url, json=data)

    if r.status_code != 200:
        raise Exception(f'Post Request Failed url: {url}')

    return r.json()


def get_mg_nodes_info():
    url = f'http://{SUPER_NODE_IP}:{SUPER_APP_PORT}/fog_nodes'
    fog_nodes_info = get_request(url)
    return fog_nodes_info["fog_nodes"]


def get_cluster_metrics(node_ip):
    url = f'http://{node_ip}:{MG_APP_PORT}/cluster_metrics'
    cluster_metrics = get_request(url)
    return cluster_metrics


def setup():
    mg_node_list = []
    mg_nodes = get_mg_nodes_info()

    for mg_node in mg_nodes:
        mg_node_dict = {}
        mg_node_dict['name'] = mg_node['name']
        mg_node_dict['ip'] = mg_node['ip']

        mg_node_list.append(mg_node_dict)

    return mg_node_list


def select_node(mg_node_list):
    min_rtt = min([mg_node_dict['rtt']
                  for mg_node_dict in mg_node_list])
    selected_node = [
        mg_node_dict for mg_node_dict in mg_node_list if mg_node_dict['rtt'] == min_rtt][0]
    return selected_node


def find_node(mg_node_list):
    hight_cpu_node_list = []
    low_cpu_node_list = []
    for mg_node_dict in mg_node_list:
        cluster_metrics = get_cluster_metrics(mg_node_dict['ip'])

        mg_node_dict['cpu'] = cluster_metrics['cpu']
        print(cluster_metrics['cpu'])
        mg_node_dict['rtt'] = ping3.ping(mg_node_dict['ip'])

        if cluster_metrics['cpu'] > CPU_THRESHOLD:
            hight_cpu_node_list.append(mg_node_dict)
        else:
            low_cpu_node_list.append(mg_node_dict)

    # CPUが閾値を超えていないノードがあれば、そのノードを選択
    if len(low_cpu_node_list) > 0:
        return select_node(low_cpu_node_list)

    return select_node(hight_cpu_node_list)


def test(mg_node_list):
    print(mg_node_list[1]['name'])
    while True:
        cluster_metrics = get_cluster_metrics(mg_node_list[1]['ip'])
        print(cluster_metrics['cpu'])
        time.sleep(1)


def main():
    mg_node_list = setup()
    test(mg_node_list)
    # selected_node = find_node(mg_node_list)
    # print(f'selected node name: {selected_node["name"]}')
    # print(f'ip: {selected_node["ip"]}')


if __name__ == '__main__':
    main()
