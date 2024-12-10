package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-project/api/mq_consume"
)

// UserRepo is a Greater repo.
type UserRepo interface {
	Create(context.Context, *mq_consume.MqUserInfo) (int64, error)
	Update(context.Context, *mq_consume.MqUserInfo) error
	FindByAccount(context.Context, string) (*mq_consume.MqUserInfo, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// SaveUser save user.
func (uc *UserUsecase) SaveUser(ctx context.Context, msg *mq_consume.MqUserInfo) error {
	if msg.Account == "" {
		return mq_consume.ErrorUserAccountMissing("account is missing")
	}
	find, _ := uc.repo.FindByAccount(ctx, msg.Account)
	if find != nil {
		msg.Id = find.Id
		return uc.repo.Update(ctx, msg)
	} else {
		_, err := uc.repo.Create(ctx, msg)
		if err != nil {
			return err
		}
	}
	return nil
}
