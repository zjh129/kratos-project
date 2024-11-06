package biz

import (
	"context"

	"kratos-project/api/grpc_user"

	"github.com/go-kratos/kratos/v2/log"
)

// UserRepo is a User repo.
type UserRepo interface {
	Save(context.Context, *grpc_user.UserSaveRequest) (*grpc_user.UserSaveReply, error)
	Find(context.Context, *grpc_user.UserInfoRequest) (*grpc_user.UserInfo, error)
	PageList(context.Context, *grpc_user.UserListRequest) (*grpc_user.UserListReply, error)
	Delete(context.Context, *grpc_user.UserDeleteRequest) (*grpc_user.UserDeleteReply, error)
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewUserUsecase new a Greeter usecase.
func NewUserUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// SaveUser creates a User, and returns the new Greeter.
func (uc *UserUsecase) SaveUser(ctx context.Context, u *grpc_user.UserSaveRequest) (*grpc_user.UserSaveReply, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", u.Account)
	return uc.repo.Save(ctx, u)
}

// GetUser gets the specified User.
func (uc *UserUsecase) GetUser(ctx context.Context, req *grpc_user.UserInfoRequest) (*grpc_user.UserInfo, error) {
	uc.log.WithContext(ctx).Infof("GetUser: %+v", req)
	return uc.repo.Find(ctx, req)
}

// GetUserList gets the specified User list.
func (uc *UserUsecase) GetUserList(ctx context.Context, req *grpc_user.UserListRequest) (*grpc_user.UserListReply, error) {
	uc.log.WithContext(ctx).Infof("GetUserList: %+v", req)
	return uc.repo.PageList(ctx, req)
}

// DeleteUser deletes the specified User.
func (uc *UserUsecase) DeleteUser(ctx context.Context, req *grpc_user.UserDeleteRequest) (*grpc_user.UserDeleteReply, error) {
	uc.log.WithContext(ctx).Infof("DeleteUser: %+v", req)
	return uc.repo.Delete(ctx, req)
}
