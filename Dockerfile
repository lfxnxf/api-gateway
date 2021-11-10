FROM golang:latest
RUN mkdir /docker
ADD . /docker/
WORKDIR /docker
RUN go env -w GO111MODULE=on
RUN go env -w  GOPROXY=https://goproxy.cn,direct
RUN go mod tidy
RUN go build -o main .
CMD ["/docker/main"]