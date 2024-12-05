package service

import (
	"context"

	"kratos-project/api/grpc_user"
	"kratos-project/app/grpc_user/internal/biz"
)

// GreeterService is a greeter service.
type UserService struct {
	grpc_user.UnimplementedUserServer

	uc *biz.UserUsecase
}

// NewGreeterService new a greeter service.
func NewUserService(uc *biz.UserUsecase) *UserService {
	return &UserService{uc: uc}
}

// UserFind implements user.UserServer.UserFind
func (s *UserService) UserFind(ctx context.Context, req *grpc_user.UserInfoRequest) (resp *grpc_user.UserInfo, err error) {
	return s.uc.GetUser(ctx, req)
}

// UserList implements user.UserServer.UserList
func (s *UserService) UserList(ctx context.Context, req *grpc_user.UserListRequest) (*grpc_user.UserListReply, error) {
	return s.uc.GetUserList(ctx, req)
}

// UserSave implements user.UserServer.UserSave
func (s *UserService) UserSave(ctx context.Context, req *grpc_user.UserSaveRequest) (*grpc_user.UserSaveReply, error) {
	return s.uc.SaveUser(ctx, req)
}

// UserDelete implements user.UserServer.UserDelete
func (s *UserService) UserDelete(ctx context.Context, req *grpc_user.UserDeleteRequest) (*grpc_user.UserDeleteReply, error) {
	return s.uc.DeleteUser(ctx, req)
}
