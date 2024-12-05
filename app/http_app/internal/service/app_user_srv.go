package service

import (
	"context"
	"kratos-project/api/http_app"
	"kratos-project/app/http_app/internal/biz"
)

type AppUserSrv struct {
	http_app.UserHTTPServer

	uc *biz.UserUsecase
}

func NewAppUserSrv(uc *biz.UserUsecase) *AppUserSrv {
	return &AppUserSrv{uc: uc}
}

// CreateUser create user.
func (s *AppUserSrv) UserInfo(ctx context.Context, req *http_app.UserInfoRequest) (*http_app.UserInfoReply, error) {
	return s.uc.UserInfo(ctx)
}

// UserLogin user login.
func (s *AppUserSrv) UserLogin(ctx context.Context, req *http_app.UserLoginRequest) (*http_app.UserLoginReply, error) {
	return s.uc.UserLogin(ctx, req)
}
