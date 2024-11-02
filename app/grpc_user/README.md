# 说明文档

## ent使用
ent的命令需要再应用目录中执行，本项目中在`app/grpc_user`目录下执行
```shell
# 初始化ent
make entinit
# 导入已存在的数据库的表
make entimport
# 生成ent代码
make ent
```

## ent软删除的实现，使用的是ent interceptor
1、升级ent到最新版
```shell
go get -u entgo.io/ent
```
2、添加interceptor

调用`go generate ./ent`生成代码时加入`--feature intercept`
`make ent`中已经加入了`--feature intercept`，所以直接执行`make ent`即可

3、实现interceptor
拷贝`ent/schema`目录下的`soft_delete_mixin.go`文件到指定目录下，更改包的引用路径为项目路径。

4、使用interceptor
在`ent/schema`目录下的`user.go`文件中添加`Mixin`.以user表为例：
```go
// Mixin of the UserInfo.
func (UserInfo) Mixin() []ent.Mixin {
	return []ent.Mixin{
		SoftDeleteMixin{},
	}
}
```

5、最后在初始化ent的客户端时，引入`ent/runtime`,此处示例：
```go
_ "kratos-project/app/grpc_user/internal/data/ent/runtime"
```

## 缓存使用
缓存驱动使用了`jetcache-go`，`jetcache-go`是基于`go-redis/cache`拓展的通用缓存访问框架。 实现了类似Java版`JetCache`的核心功能。
[github地址](github.com/mgtv-tech/jetcache-go)