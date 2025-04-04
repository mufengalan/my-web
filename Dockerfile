FROM registry.cn-hangzhou.aliyuncs.com/golang:1.23.0-alpine AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app

COPY . .
RUN go mod download
#RUN go get github.com/githubnemo/CompileDaemon
#ENTRYPOINT CompileDaemon --build="go build -o my-web ./cmd/main.go" --command=./my-web
RUN go build -o my-web ./cmd/main.go

FROM registry.cn-hangzhou.aliyuncs.com/alpine:latest

# 安装 tzdata 包，确保支持时区的配置
#RUN apk add --no-cache tzdata

# 设置工作目录为 /app
WORKDIR /app

# 从编译阶段的镜像中拷贝编译后的二进制文件到运行镜像中
COPY --from=builder /app/my-web /app/my-web


EXPOSE 8080

CMD ["/app/my-web"]
