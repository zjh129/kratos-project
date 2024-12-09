# Kratos项目实践

本项目主要目的是为了学习Kratos框架的使用，以及对项目的一些实践。由于Kratos框架的example较碎片化，未能体现出完整的项目的样子，因此创建了本项目，为其他学习Kratos框架的开发者提供参考。

项目采用微服务大仓库模式，即一个仓库包含多个服务，每个服务都是一个独立的服务，通过make命令进行统一管理。

大仓模式下，Kratos的Layout会在api目录每个模块多出v1样式版本目录，大部分情况下增加了引包的复杂度，造成引包重名，增加了不必要的麻烦，所以将api、rpc的proto目录直接提升到根目录下，方便引包。

## 项目备注
- 项目使用makefile管理操作命令，在linux、mac下运行更佳，windows不支持make命令，windows用户可以选择在wsl2子系统下开发（使用vscode远程开发体验最佳，或使用goland远程连接开发，但使用体验上很差）。
- 通过kratos框架运行的服务，例如grpc、http、cron、mq消费等服务，为其他的kratos项目开发提供参考。
- 该项目使用微服务大仓库模式，即一个仓库包含多个服务，每个服务都是一个独立的服务，通过make命令进行统一管理，make命令参考[beer-shop]项目集中管理，开发时进入对于的`app\模块`目录下通过命令执行make指令，指令在`app_makefile`中查看。

## 新建API或grpc协议文件
```bash
# 新建http服务协议文件
kratos proto add api/http_admin/user.proto
# 新建grpc服务协议文件
kratos proto add api/grpc_user/user.proto
# 新建mq消费数据协议文件
kratos proto add api/mq/user.proto
```

## 新建服务
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

## 运行系统配套服务
```bash
make docker
```

## 软件UI管理界面
- [Consul服务注册/服务发现/配置中心管理界面](http://127.0.0.1:8500/)
- [Jaeger链路追踪管理界面](http://127.0.0.1:16686/)

## 参考项目
- [Kratos Example](https://github.com/go-kratos/examples)
- [beer-shop](https://github.com/go-kratos/beer-shop)
- [Kratos-Cron](https://github.com/igo9go/kratos-cron)