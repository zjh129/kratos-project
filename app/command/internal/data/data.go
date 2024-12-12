package data

import (
	"context"

	"kratos-project/api/grpc_user"
	"kratos-project/app/command/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/circuitbreaker"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"

	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewDiscovery, NewUserRpcClient, NewUserRpcRepo)

// Data .
type Data struct {
	log      *log.Helper
	guClient grpc_user.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, uc grpc_user.UserClient) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{
		log:      l,
		guClient: uc,
	}, cleanup, nil
}

func NewUserRpcClient(r registry.Discovery) grpc_user.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///kratos-project.grpc_user"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			// 异常恢复
			recovery.Recovery(),
			//熔断器
			circuitbreaker.Client(),
			//链路追踪
			tracing.Client(),
		),
	)
	if err != nil {
		panic(err)
	}
	return grpc_user.NewUserClient(conn)
}
