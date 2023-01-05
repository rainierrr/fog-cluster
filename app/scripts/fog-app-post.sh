curl -v -X POST http://localhost:3000/post -H 'Content-Type: application/json' \
  -d '{ "temperature": 20, "humidity": 0.5, "pressure": 1000, "location": "Tokyo" }'
