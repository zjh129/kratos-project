package service

import (
	"context"
	"github.com/tx7do/kratos-transport/broker"
	"kratos-project/api/mq_consume"
	"kratos-project/app/mq_consume/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type UserService struct {
	logger log.Logger
	uc     *biz.UserUsecase
}

func NewUserService(logger log.Logger, uc *biz.UserUsecase) *UserService {
	return &UserService{
		logger: logger,
		uc:     uc,
	}
}

// SaveUser save user.
func (s *UserService) SaveUser(ctx context.Context, _ string, _ broker.Headers, msg *mq_consume.MqUserInfo) error {
	return s.uc.SaveUser(ctx, msg)
}
