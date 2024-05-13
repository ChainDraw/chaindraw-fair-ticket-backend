# 使用官方的 Go 镜像作为基础镜像
FROM golang:1.22.1

# 设置工作目录
WORKDIR /app

COPY . .

RUN go env -w GOPROXY=https://goproxy.cn,direct
# 整理依赖项（可选）
RUN go mod tidy
# 编译 Go 应用程序
RUN sudo go build -o  main .

# 设置容器启动时执行的命令
CMD ["./main"]
