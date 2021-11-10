FROM golang:latest
RUN mkdir /docker
ADD app /docker/
WORKDIR /docker
RUN go env -w GO111MODULE=on
RUN go env -w  GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN go build -o api-gateway ./app/main.go
CMD ["/docker/main"]