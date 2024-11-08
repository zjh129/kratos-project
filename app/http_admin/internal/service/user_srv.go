package service

import (
	"context"
	"kratos-project/api/http_admin"

	"kratos-project/app/http_admin/internal/biz"
)

// UserService is a greeter service.
type UserService struct {
	http_admin.UserHTTPServer

	uc *biz.UserUsecase
}

// NewUserService new a greeter service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) CreateUser(ctx context.Context, req *http_admin.CreateUserRequest) (*http_admin.CreateUserReply, error) {
	return s.uc.CreateUser(ctx, req)
}
func (s *UserService) DeleteUser(ctx context.Context, req *http_admin.DeleteUserRequest) (*http_admin.DeleteUserReply, error) {
	return s.uc.DeleteUser(ctx, req)
}
func (s *UserService) FindUser(ctx context.Context, req *http_admin.FindUserRequest) (*http_admin.FindUserReply, error) {
	return s.uc.FindUser(ctx, req)
}
func (s *UserService) ListUser(ctx context.Context, req *http_admin.ListUserRequest) (*http_admin.ListUserReply, error) {
	return s.uc.ListUser(ctx, req)
}
func (s *UserService) UpdateUser(ctx context.Context, req *http_admin.UpdateUserRequest) (*http_admin.UpdateUserReply, error) {
	return s.uc.UpdateUser(ctx, req)
}
func (s *UserService) UserLogin(ctx context.Context, req *http_admin.UserLoginRequest) (*http_admin.UserLoginReply, error) {
	return s.uc.UserLogin(ctx, req)
}
func (s *UserService) UserLogout(ctx context.Context, req *http_admin.UserLogoutRequest) (*http_admin.UserLogoutReply, error) {
	return s.uc.UserLogout(ctx, req)
}
