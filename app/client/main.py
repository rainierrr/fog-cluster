import httpx
import asyncio
import ipaddress

fog_app_port = 3000
fog_app_path = 'local'


def get_request(ip, path):
    url = f'http://{ip}:{fog_app_port}/{path}'
    r = httpx.get(url)

    return r.json()


def post_request(ip, path, data):
    url = f'http://{ip}:{fog_app_port}/{path}'
    r = httpx.post(url, json=data)

    return r.json()


async def get_request_with_client(client, ip):
    url = f'http://{ip}:{fog_app_port}/{fog_app_path}'

    r = await client.get(url, timeout=1)

    if r.status_code == 200:
        print(f'Found server on ip: {ip}')
        return ip


async def find_fog_app(local_network):
    local_ip_lsit = ipaddress.IPv4Network(local_network).hosts()

    async with httpx.AsyncClient() as client:
        tasks = [get_request_with_client(client, str(ip))
                 for ip in local_ip_lsit]
        results = await asyncio.gather(*tasks, return_exceptions=True)

    for result in results:
        if type(result) == str:
            return result


def main():
    local_network = '192.168.11.0/24'

    fog_app_ip = asyncio.run(find_fog_app(local_network))

    get_result = get_request(fog_app_ip, '')
    print(f'get_result: {get_result}')

    post_data = {"temperature": 20, "humidity": 0.5,
                 "pressure": 1000, "location": "Tokyo"}

    post_result = post_request(fog_app_ip, 'post', post_data)
    print(f'post_result: {post_result}')


if __name__ == '__main__':
    main()
