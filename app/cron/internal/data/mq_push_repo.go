package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-project/api/mq_consume"
)

type MqPushRepo struct {
	data *Data
	log  *log.Helper
}

// NewMqPushRepo . Create a new mq push repo.
func NewMqPushRepo(data *Data, logger log.Logger) *MqPushRepo {
	return &MqPushRepo{data: data, log: log.NewHelper(logger)}
}

// AddUser . Add user.
func (m *MqPushRepo) AddUser(info *mq_consume.MqUserInfo) error {
	_ = m.data.rabbitmqBroker.Publish(context.Background(), "user", info)
	return nil
}
