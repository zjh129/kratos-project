# Kratos项目实践

- 项目使用makefile管理操作命令，在linux、mac下运行更佳，windows默认不支持make命令，windows用户可以选择在wsl子系统下开发。
- 该项目是基于Kratos框架的项目实践，主要是为了学习Kratos框架的使用，以及对项目的一些实践。
- 通过kratos框架运行的服务，例如grpc、http、cron、mq消费等，为后续的kratos项目开发提供参考。
- 该项目使用微服务大仓库模式，即一个仓库包含多个服务，每个服务都是一个独立的服务，通过make命令进行统一管理。

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