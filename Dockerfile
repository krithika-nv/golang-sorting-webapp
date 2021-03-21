FROM golang:1.16.2-buster
ENV GO111MODULE=on
RUN git clone https://github.com/krithika-nv/golang-sorting-webapp.git
WORKDIR /go/golang-sorting-webapp
RUN go mod download
RUN go get github.com/prometheus/client_golang/prometheus
RUN go get github.com/sirupsen/logrus
RUN go build main.go
CMD ./main
