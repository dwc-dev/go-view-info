# 使用轻量级的 Alpine 版 Golang 镜像
FROM golang:1.20-alpine

# 设置容器内的工作目录
WORKDIR /app

# 将 Go 源代码复制到容器中
COPY . .

# 构建 Go 应用
RUN go build -o server .

# 暴露服务器运行的端口
EXPOSE 8080

# 设置 ENTRYPOINT 指令
ENTRYPOINT ["/app/server"]
