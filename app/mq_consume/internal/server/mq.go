package server

import (
	"context"
	"kratos-project/app/mq_consume/internal/service"

	"kratos-project/app/mq_consume/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/transport/rabbitmq"
)

// NewRabbitMQServer create a rabbitmq server.
func NewRabbitMQServer(cfg *conf.Data, _ log.Logger, svc *service.UserService) *rabbitmq.Server {
	ctx := context.Background()

	srv := rabbitmq.NewServer(
		rabbitmq.WithGlobalTracerProvider(),
		rabbitmq.WithGlobalPropagator(),
		rabbitmq.WithCodec("json"),
		rabbitmq.WithAddress(cfg.Rabbitmq.Addrs),
		rabbitmq.WithExchange(cfg.Rabbitmq.Exchange, cfg.Rabbitmq.DurableExchange),
	)
	// 注册订阅者
	registerRabbitMQSubscribers(ctx, srv, svc)

	return srv
}

// registerRabbitMQSubscribers 注册订阅者
func registerRabbitMQSubscribers(ctx context.Context, srv *rabbitmq.Server, svc *service.UserService) {
	// 注册订阅者
	_ = rabbitmq.RegisterSubscriber(srv, ctx, "user", svc.SaveUser)
}
