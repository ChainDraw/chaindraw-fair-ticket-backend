# chaindraw-fair-ticket 后端仓库

## 目录结构

```shell
.
├── Dockerfile
├── Makefile
├── README.md
├── api
├── config
│   ├── app_config.go
│   ├── mysql.go
│   ├── redis.go
│   ├── server.go
│   └── zap.go
├── core
│   ├── server.go
│   ├── viper.go
│   └── zap.go
├── etc
│   └── config.yml
├── global
│   ├── biz
│   └── global.go
├── go.mod
├── go.sum
├── initialize
│   └── gorm.go
├── main.go
├── manifest
│   ├── deploy
│   └── docs
├── middleware
├── model
├── pkg
├── router
│   └── router.go
└── service
```

- api: 可以理解为 Java 中的 controller 层，用于接收参数和返回响应
- config: 配置层，用于存放相关配置的结构体
- core: 初始化，目前只保存服务器、加载配置和日志的初始化文件
- etc: 只存放配置文件，可以是 json、toml 或者 yaml 文件
- global: 全局变量，比如全局日志变量、全局数据库实例等等
- initialize: 初始化层，初始化相关配置
- manifest: 资源清单，里面可以包含部署文件以及 swagger 文档等资源
- middleware: 中间件层
- model: 模型层，存储数据库表对应实体，以及请求响应需要定义的结构体
- pkg: 工具层，包含第三方工具库和自定义的工具库
- router: 路由配置层
- service: 业务层，主要处理业务的模块

## 热更新使用到的 air 库

```shell
go get github.com/cosmtrek/air

go install github.com/cosmtrek/air
```

## 开发规范

### 命名相关规范

接口请使用 RestFul 风格，比如 `GET /api/v1/user/:id`

### 开发注意事项

1. 如果可以，请创建自己要实现模块的分支(不要在main分支提交代码)，名称以模块名定义，比如 `feature-login`
    
    ```shell
    git clone git@github.com:ChainDraw/chaindraw-fair-ticket-backend.git

    # 新建并切换分支为 feature-login
    git checkout -b feature-login

    # 提交时直接提交到远程的相应分支
    git push origin feature-login
    ```

    最后分支的合并可大家一起看看实现

2. 业务逻辑不是非常复杂，可以直接在 service 调用 gorm 操作

3. api 目录下区分前后台 `console` 和 `chaindraw`，其余 service model 等部分不需要像 api 目录下区分 `console` 和 `chaindraw`

4. 每一个功能尽量以单独的 go 文件存在，这样能尽最大可能避免 git 冲突

5. 文件名称以小写开头，多个单词以下划线 `_` 分隔，比如 `get_user_list.go`

6. 测试文件单独放在项目的 `test` 目录下，比如 `test/get_user_list_test.go`