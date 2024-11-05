package service

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
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

func (s *UserService) UserFind(ctx context.Context, req *grpc_user.UserInfoRequest) (resp *grpc_user.UserInfo, err error) {
	var user *biz.User
	switch {
	case req.Id > 0:
		user, err = s.uc.GetUser(ctx, req.Id)
	case req.Account != "":
		user, err = s.uc.GetUserByAccount(ctx, req.Account)
	default:
		return nil, errors.BadRequest("user", "missing id or account")
	}
	if err != nil {
		return nil, err
	}
	return &grpc_user.UserInfo{
		Id:        user.Id,
		Account:   user.Account,
		Password:  user.Password,
		Name:      user.Name,
		Avatar:    user.Avatar,
		Type:      grpc_user.UserType(user.UserType),
		StatusIs:  grpc_user.UserStatus(user.Status),
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}, nil
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
