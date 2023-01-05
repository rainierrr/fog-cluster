curl -v -X POST http://localhost:3030/fog_node -H 'Content-Type: application/json' \
  -d '{ "name": "fog-node-1", "token": "fog-node-1-token", "tag": "tag1", "ip": "192.168.1.1"}'
