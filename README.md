# golang-sorting-webapp

docker build -t krithikan/golang-sorting-webapp:krithika-nv .

docker run -p 8080:8080 -p 9110:9110 -d krithikan/golang-sorting-webapp:krithika-nv

docker run -p 8081:8081 -p 9090:9090 -e PORT=8081 -e PROM_PORT=9090 -d krithikan/golang-sorting-webapp:krithika-nv

This repo contains dependencies and the Dockerfile used to create the Golang webapp that sorts incoming http headers. Integrated with Prometheus to provide number_of_requests custom metric. 
