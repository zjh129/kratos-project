package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"kratos-project/api/grpc_user"
	"kratos-project/app/http_admin/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(
	NewData, NewDiscovery,
	NewUserRpcClient,
	NewUserRpcRepo,
)

// Data .
type Data struct {
	log       *log.Helper
	gu_client grpc_user.UserClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, uc grpc_user.UserClient) (*Data, func(), error) {
	l := log.NewHelper(log.With(logger, "module", "data"))
	cleanup := func() {
		l.Info("closing the data resources")
	}
	return &Data{
		log:       l,
		gu_client: uc,
	}, cleanup, nil
}

func NewUserRpcClient(r registry.Discovery) grpc_user.UserClient {
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///kratos-project.grpc_user"),
		grpc.WithDiscovery(r),
		grpc.WithMiddleware(
			recovery.Recovery(),
		),
	)
	if err != nil {
		panic(err)
	}
	return grpc_user.NewUserClient(conn)
}
