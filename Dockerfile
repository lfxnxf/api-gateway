FROM golang:latest
RUN mkdir /docker

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/build
WORKDIR /docker
# 将代码复制到容器中
COPY . .
# 编译文件
RUN go build -o api-gateway ./app/main.go
# 将二进制文件从 /build 目录复制到这里
RUN cp /build/api-gateway .
# 移动到用于存放生成的二进制文件的 /dist 目录
WORKDIR /dist
CMD ["/dist/api-gateway"]