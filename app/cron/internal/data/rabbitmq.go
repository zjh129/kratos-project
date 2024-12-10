package data

import (
	"kratos-project/app/cron/internal/conf"

	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/broker/rabbitmq"
)

func NewRabbitMQBroker(cfg *conf.Data) broker.Broker {
	b := rabbitmq.NewBroker(
		broker.WithAddress(cfg.Rabbitmq.Addrs...),
		broker.WithCodec(cfg.Rabbitmq.Codec),
		rabbitmq.WithExchangeName(cfg.Rabbitmq.Exchange),
		rabbitmq.WithDurableExchange(),
	)
	if b == nil {
		return nil
	}

	_ = b.Init()

	if err := b.Connect(); err != nil {
		return nil
	}

	return b
}
