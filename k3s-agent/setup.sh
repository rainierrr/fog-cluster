# バイナリダウンロード＆設置
curl -L https://github.com/k3s-io/k3s/releases/download/v1.25.4+k3s1/k3s -o /usr/local/bin/k3s
sudo chown root:root /usr/local/bin/k3s
sudo chmod +x /usr/local/bin/k3s

# symlink
for cmd in kubectl crictl ctr; do
  ln -s /usr/local/bin/k3s /usr/local/bin/$cmd
done

# k3s.service
sudo cp ./k3s-agent.service /etc/systemd/system/k3s-agent.service

# k3s-agent起動
sudo systemctl daemon-reload
sudo systemctl start k3s-agent
