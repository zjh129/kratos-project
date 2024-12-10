package server

import (
	"context"
	"github.com/tx7do/kratos-transport/broker"

	"kratos-project/app/mq_consume/internal/conf"
	"kratos-project/app/mq_consume/internal/service"

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
	_ = rabbitmq.RegisterSubscriber(srv, ctx, cfg.Rabbitmq.Routing, svc.SaveUser, broker.WithQueueName("user"))

	return srv
}
