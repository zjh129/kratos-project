package biz

import (
	"context"
	"errors"

	"kratos-project/api/grpc_user"

	"github.com/go-kratos/kratos/v2/log"
)

// UserRepo is a User repo.
type UserRepo interface {
	Save(context.Context, *grpc_user.UserSaveRequest) (*grpc_user.UserSaveReply, error)
	FindById(context.Context, int64) (*grpc_user.UserInfo, error)
	FindByAccount(context.Context, string) (*grpc_user.UserInfo, error)
	Count(context.Context, *grpc_user.UserListRequest) (int64, error)
	PageList(context.Context, *grpc_user.UserListRequest) ([]*grpc_user.UserInfo, error)
	DeleteById(context.Context, int64) error
	DeleteByIds(context.Context, []int64) error
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
	switch {
	case req.Id > 0:
		return uc.repo.FindById(ctx, req.Id)
	case req.Account != "":
		return uc.repo.FindByAccount(ctx, req.Account)
	default:
		return nil, errors.New("invalid request")
	}
}

// GetUserList gets the specified User list.
func (uc *UserUsecase) GetUserList(ctx context.Context, req *grpc_user.UserListRequest) (*grpc_user.UserListReply, error) {
	uc.log.WithContext(ctx).Infof("GetUserList: %+v", req)
	count, err := uc.repo.Count(ctx, req)
	if err != nil {
		return nil, err
	}
	list, err := uc.repo.PageList(ctx, req)
	if err != nil {
		return nil, err
	}
	return &grpc_user.UserListReply{Total: count, List: list}, nil
}

// DeleteUser deletes the specified User.
func (uc *UserUsecase) DeleteUser(ctx context.Context, req *grpc_user.UserDeleteRequest) (*grpc_user.UserDeleteReply, error) {
	uc.log.WithContext(ctx).Infof("DeleteUser: %+v", req)
	switch {
	case req.Id > 0:
		err := uc.repo.DeleteById(ctx, req.Id)
		if err != nil {
			return nil, err
		}
		return &grpc_user.UserDeleteReply{Ids: []int64{req.Id}}, nil
	case len(req.Ids) > 0:
		err := uc.repo.DeleteByIds(ctx, req.Ids)
		if err != nil {
			return nil, err
		}
		return &grpc_user.UserDeleteReply{Ids: req.Ids}, nil
	default:
		return nil, errors.New("invalid request")
	}
}
