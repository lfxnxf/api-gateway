FROM golang:alpine AS builder
RUN mkdir /docker

# 为我们的镜像设置必要的环境变量
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOPROXY=https://goproxy.cn,direct

# 移动到工作目录：/docker
WORKDIR /docker
# 将代码复制到容器中
COPY . .
# 编译文件
RUN go mod tidy
RUN go build -o api-gateway ./app/main.go

###################
# 接下来创建一个小镜像
###################
FROM scratch
COPY ./app/config /config
COPY ./app/logs /logs

# 从builder镜像中把/dist/app 拷贝到当前目录
COPY --from=builder /docker/api-gateway /

ENTRYPOINT ["/api-gateway", "-config", "config/producer/config.toml"]