# バイナリダウンロード＆設置
curl -L https://github.com/rancher/k3s/releases/download/v0.8.0/k3s -o /usr/local/bin/k3s
chown root:root /usr/local/bin/k3s
chmod +x /usr/local/bin/k3s

# symlink
for cmd in kubectl crictl ctr; do
  ln -s /usr/local/bin/k3s /usr/local/bin/$cmd
done

# k3s.service
mv ./k3s.service /etc/systemd/system/k3s.service

# k3s-agent起動
sudo systemctl daemon-reload
sudo systemctl start k3s
