FROM golang:1.23 AS builder

WORKDIR /app

# 设置 Go 模块代理
RUN go env -w GOPROXY=https://goproxy.cn,direct

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# 构建 Linux 可执行文件
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main .

# 第二阶段：构建镜像
FROM alpine:latest

WORKDIR /root/

# 复制可执行文件
COPY --from=builder /app/main .

# 复制配置文件或目录
COPY --from=builder /app/infrastructure /root/infrastructure

EXPOSE 8080
CMD ["./main"]
