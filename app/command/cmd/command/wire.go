//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"kratos-project/app/command/internal/biz"
	"kratos-project/app/command/internal/conf"
	"kratos-project/app/command/internal/data"
	"kratos-project/app/command/internal/server"
	"kratos-project/app/command/internal/service"

	"github.com/go-kratos/kratos/v2/log"

	"github.com/google/wire"
	"github.com/spf13/cobra"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Registry, *cobra.Command, log.Logger) (*server.CommandServer, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet))
}
