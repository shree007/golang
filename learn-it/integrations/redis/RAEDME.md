
$ docker run --name redis-container -d -p 6379:6379 redis

$ go run main.go

$ curl -X POST http://localhost:8080/db0 \
     -H "Content-Type: application/json" \
     -d '{
           "Key": "Prometheus",
           "Value": "Grafana"
         }'

$ curl -X GET http://localhost:8080/db0/Prometheus

$ curl -X DELETE http://localhost:8080/db0/Prometheus

$ curl -X PUT http://localhost:8080/db0/Prometheus \
     -H "Content-Type: application/json" \
     -d '{
           "Key": "Prometheus",
           "Value": "node-exporter"
         }'
