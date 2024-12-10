package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"kratos-project/api/mq_consume"
)

type MqPushRepo interface {
	AddUser(info *mq_consume.MqUserInfo) error
}

type MqPushUseCase struct {
	repo MqPushRepo
	log  *log.Helper
}

func NewMqPushUseCase(repo MqPushRepo, logger log.Logger) *MqPushUseCase {
	return &MqPushUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (u *MqPushUseCase) AddUser(info *mq_consume.MqUserInfo) error {
	return u.repo.AddUser(info)
}
