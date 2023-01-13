import httpx
import ping3

SUPER_NODE_IP = '52.199.79.207'
MG_NODE_A_IP = '3.113.76.40'
MG_NODE_B_IP = '54.255.41.147'
SUPER_APP_PORT = 3000
MG_APP_PORT = 30002
FOG_APP_PORT = 30003


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


def find_node():
    mg_nodes_list = []
    mg_nodes = get_mg_nodes_info()
    for mg_node in mg_nodes:
        mg_node_info = {}
        cluster_metrics = get_cluster_metrics(mg_node['ip'])

        mg_node_info['name'] = mg_node['name']
        mg_node_info['ip'] = mg_node['ip']
        mg_node_info['cpu'] = cluster_metrics['cpu']
        mg_node_info['rtt'] = ping3.ping(mg_node['ip'])

        mg_nodes_list.append(mg_node_info)

    print(mg_nodes_list)


def main():
    find_node()


if __name__ == '__main__':
    main()
