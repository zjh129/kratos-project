# Kratos项目实践

项目在linux、mac下运行更佳，windows不支持make命令，建议windows用户在wsl子系统下开发。

该项目是基于Kratos框架的项目实践，主要是为了学习Kratos框架的使用，以及对项目的一些实践。

通过kratos框架运行的服务，例如grpc、http、cron、mq消费等，为后续的kratos项目开发提供参考。

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

## 安装环境
```
make init
```
## Create a service
```
# Create a template project
kratos new server

cd server
# Add a proto template
kratos proto add api/server/server.proto
# Generate the proto code
kratos proto client api/server/server.proto
# Generate the source code of service by proto file
kratos proto server api/server/server.proto -t internal/service

go generate ./...
go build -o ./bin/ ./...
./bin/server -conf ./configs
```
## Generate other auxiliary files by Makefile
```
# Download and update dependencies
make init
# Generate API files (include: pb.go, http, grpc, validate, swagger) by proto file
make api
# Generate all files
make all
```

## Automated Initialization (wire)
```
# install wire
go get github.com/google/wire/cmd/wire

# generate wire
cd cmd/server
wire
```

## Docker
```bash
# build
docker build -t <your-docker-image-name> .

# run
docker run --rm -p 8000:8000 -p 9000:9000 -v </path/to/your/configs>:/data/conf <your-docker-image-name>
```