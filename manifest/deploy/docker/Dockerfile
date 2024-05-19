# 使用官方的 Go 镜像作为基础镜像
FROM golang:1.22.1

# 设置工作目录
WORKDIR /app

COPY . .

# 编译 Go 应用程序
RUN go build -o  main .

# 设置容器启动时执行的命令
CMD ["./main"]
