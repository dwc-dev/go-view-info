# 使用轻量级的 Alpine 版 Golang 镜像
FROM golang:1.20-alpine

# 设置容器内的工作目录
WORKDIR /app

# 只复制 main.go 文件
COPY main.go .

# 直接编译 main.go
RUN go build -o server main.go

# 暴露服务器运行的端口
EXPOSE 8080

# 设置 ENTRYPOINT 指令
ENTRYPOINT ["/app/server"]
