package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-project/api/mq_consume"
	"kratos-project/app/cron/internal/biz"
)

type mqPushRepo struct {
	data *Data
	log  *log.Helper
}

// NewMqPushRepo . Create a new mq push repo.
func NewMqPushRepo(data *Data, logger log.Logger) biz.MqPushRepo {
	return &mqPushRepo{data: data, log: log.NewHelper(logger)}
}

// AddUser . Add user.
func (m *mqPushRepo) PushUserInfo(info *mq_consume.MqUserInfo) error {
	_ = m.data.rabbitmqBroker.Publish(context.Background(), "user", info)
	return nil
}
