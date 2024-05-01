.PHONY: all build run gotool clean help air

BINARY="chaindraw_fair_ticket"

all: gotool build

build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ${BINARY}

run:
	@go run ./

air:
	@air -c .air.toml	

gotool:
	go fmt ./

clean:
	@if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

help:
	@echo "make - 格式化 Go 代码, 并编译生成二进制文件"
	@echo "make build - 编译 Go 程序, 生成二进制文件"
	@echo "make run - 直接运行 Go 程序"
	@echo "make air - 以热更新模式运行 Go 程序"
	@echo "make clean - 移除二进制文件和 vim swap files"
	@echo "make gotool - 运行 Go 工具 'fmt' and 'vet'"