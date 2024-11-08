package biz

import (
	"context"
	"kratos-project/api/grpc_user"
	"kratos-project/api/http_admin"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(grpc_user.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

type UserAccountInfo struct {
	Id      int64  // 用户ID
	Account string // 用户账号
	Passwd  string // 用户密码
}

// UserRepo is a Greater repo.
type UserRepo interface {
	Create(context.Context, *http_admin.CreateUserRequest) (int64, error)
	Update(context.Context, *http_admin.UpdateUserRequest) error
	FindByID(context.Context, int64) (*http_admin.UserInfo, error)
	FindByAccount(context.Context, string) (*UserAccountInfo, error)
	ListByPage(context.Context, *http_admin.ListUserRequest) (*http_admin.ListUserReply, error)
	Delete(context.Context, int64) error
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

// CreateUser create a new user.
func (uc *UserUsecase) CreateUser(ctx context.Context, req *http_admin.CreateUserRequest) (*http_admin.CreateUserReply, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %+v", req)
	id, err := uc.repo.Create(ctx, req)
	if err != nil {
		return nil, err
	}
	return &http_admin.CreateUserReply{
		Id: id,
	}, nil
}

// DeleteUser delete a user.
func (uc *UserUsecase) DeleteUser(ctx context.Context, req *http_admin.DeleteUserRequest) (*http_admin.DeleteUserReply, error) {
	uc.log.WithContext(ctx).Infof("DeleteUser: %+v", req)
	err := uc.repo.Delete(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &http_admin.DeleteUserReply{}, nil
}

// FindUser find a user.
func (uc *UserUsecase) FindUser(ctx context.Context, req *http_admin.FindUserRequest) (*http_admin.FindUserReply, error) {
	uc.log.WithContext(ctx).Infof("FindUser: %+v", req)
	uInfo, err := uc.repo.FindByID(ctx, req.Id)
	if err != nil {
		return nil, err
	}
	return &http_admin.FindUserReply{
		User: uInfo,
	}, nil
}

// ListUser list user.
func (uc *UserUsecase) ListUser(ctx context.Context, req *http_admin.ListUserRequest) (*http_admin.ListUserReply, error) {
	uc.log.WithContext(ctx).Infof("ListUser: %+v", req)
	return uc.repo.ListByPage(ctx, req)
}

// UpdateUser update a user.
func (uc *UserUsecase) UpdateUser(ctx context.Context, req *http_admin.UpdateUserRequest) (*http_admin.UpdateUserReply, error) {
	uc.log.WithContext(ctx).Infof("UpdateUser: %+v", req)
	err := uc.repo.Update(ctx, req)
	if err != nil {
		return nil, err
	}
	return &http_admin.UpdateUserReply{}, nil
}

// UserLogin user login.
func (uc *UserUsecase) UserLogin(ctx context.Context, req *http_admin.UserLoginRequest) (*http_admin.UserLoginReply, error) {
	panic("implement me")
}

// UserLogout user logout.
func (uc *UserUsecase) UserLogout(ctx context.Context, req *http_admin.UserLogoutRequest) (*http_admin.UserLogoutReply, error) {
	panic("implement me")
}
