package service

import (
	"context"
	"kratos-project/api/grpc_user"

	"kratos-project/app/grpc_user/internal/biz"
)

// GreeterService is a greeter service.
type UserService struct {
	grpc_user.UnimplementedUserServer

	uc *biz.GreeterUsecase
}

// NewGreeterService new a greeter service.
func NewUserService(uc *biz.GreeterUsecase) *UserService {
	return &UserService{uc: uc}
}

func (s *UserService) UserFind(context.Context, *grpc_user.UserInfoRequest) (*grpc_user.UserInfo, error) {
	panic("implement me")
}

func (s *UserService) UserList(context.Context, *grpc_user.UserListRequest) (*grpc_user.UserListReply, error) {
	panic("implement me")
}

func (s *UserService) UserSave(context.Context, *grpc_user.UserSaveRequest) (*grpc_user.UserSaveReply, error) {
	panic("implement me")
}

func (s *UserService) UserDelete(context.Context, *grpc_user.UserDeleteRequest) (*grpc_user.UserDeleteReply, error) {
	panic("implement me")
}
