# Kratos项目实践

本项目主要目的是为了学习Kratos框架的使用，以及对项目的一些实践。由于Kratos框架的example较碎片化，未能体现出完整的项目的样子，因此创建了本项目，为其他学习Kratos框架的开发者提供参考。

项目采用微服务大仓库模式，即一个仓库包含多个服务，每个服务都是一个独立的服务，通过make命令进行统一管理。

大仓模式下，Kratos的Layout会在api目录每个模块多出v1样式版本目录，大部分情况下增加了引包的复杂度，造成引包重名，增加了不必要的麻烦，所以将api、rpc的proto目录直接提升到根目录下，方便引包。

## 项目备注
- 项目使用makefile管理操作命令，在linux、mac下运行更佳，windows不支持make命令，windows用户可以选择在wsl2子系统下开发（使用vscode远程开发体验最佳，或使用goland远程连接开发，但使用体验上很差）。
- 通过kratos框架运行的服务，例如`grpc`、`http`、`cron`、`mq消费`、`cmd命令行`等服务，为其他的kratos项目开发提供参考。
- 该项目使用微服务大仓库模式，即一个仓库包含多个服务，每个服务都是一个独立的服务，通过make命令进行统一管理，make命令参考[beer-shop]项目集中管理，开发时进入对于的`app\模块`目录下通过命令执行make指令，指令在`app_makefile`中查看。

## 新建API或grpc协议文件
> 项目目录下执行
```bash
# 初始化项目需要的软件包
make init
# 启用项目运行需要的docker软件
make docker
# 生成项目需要的wire文件
make wire
# go mod tidy
make tidy
# 一键初始化项目运行需要的所有环境
make all


# 新建http服务协议文件
kratos proto add api/http_admin/user.proto
# 新建grpc服务协议文件
kratos proto add api/grpc_user/user.proto
# 新建mq消费数据协议文件
kratos proto add api/mq_consume/user.proto
```

## 新建服务
> 项目目录下执行
```bash
# 新建http服务
kratos new app/http_admin/ --nomod
# 新建grpc服务
kratos new app/grpc_user/ --nomod
# 新建cron服务
kratos new app/cron/ --nomod
# 新建mq消费服务
kratos new app/mq_consume/ --nomod
```

## 模块开发常用命令
> 进入app下某个模块下执行
```bash
# 模块配置文件proto生成go文件
make config
# 模块微服务proto生成go文件（仅用于微服务开发）
make grpc
# 模块http服务proto生成go文件（仅用于http服务开发）
make http
# 模块错误码proto生成go文件
make errors
# 模块proto文件生成go文件（一般用在根据proto协议生成go数据结构，纯粹用于数据结构定义）
make proto
# ent ORM初始化
make entinit
# ent ORM生成，根据`ent/schema`数据库表生成go文件
make ent
# ent ORM导入，导入数据库表数据schema，要求较高的请通过在线工具手动定制。
make entimport
# 模块wire生成
make wire
# generate代码生成
make generate
# 编译项目可执行程序
make build
# 运行测试
make test
# 运行
make run
```

## 运行系统配套服务
> 项目目录下执行
```bash
make docker
```

## 软件UI管理界面
- [Consul服务注册/服务发现/配置中心管理界面](http://127.0.0.1:8500/)
- [Jaeger链路追踪管理界面](http://127.0.0.1:16686/)
- [RabbitMq管理界面](http://127.0.0.1:15672/)

## 开发常用命令
> 进入app下某个模块下执行
```bash

```

## 参考项目
- [Kratos Example](https://github.com/go-kratos/examples)
- [beer-shop](https://github.com/go-kratos/beer-shop)
- [Kratos-Cron](https://github.com/igo9go/kratos-cron)
- [kratos-transport](https://github.com/tx7do/kratos-transport)