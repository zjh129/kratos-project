package main

import (
	"fmt"

	"kratos-project/app/http_admin/internal/conf"

	consul "github.com/go-kratos/kratos/contrib/registry/consul/v2"
	"github.com/go-kratos/kratos/v2/config"
	"github.com/go-kratos/kratos/v2/config/file"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/hashicorp/consul/api"
)

// initRegistry init the registry
func initRegistry() registry.Registry {
	dirverC := config.New(
		config.WithSource(
			file.NewSource(fmt.Sprintf("%s/registry.yaml", flagconf)),
		),
	)
	defer dirverC.Close()

	if err := dirverC.Load(); err != nil {
		panic(err)
	}
	registryConfig := conf.Registry{}
	if err := dirverC.Scan(&registryConfig); err != nil {
		panic(err)
	}
	switch registryConfig.Type {
	case "consul":
		if registryConfig.Consul == nil {
			panic("consul config is nil")
		}
		// 读取consul配置
		cli, err := api.NewClient(&api.Config{Address: registryConfig.Consul.Address})
		if err != nil {
			panic(err)
		}
		// 创建consul注册中心
		return consul.New(cli, consul.WithHealthCheck(false))
	default:
		panic("unknown registry driver")
	}
}
