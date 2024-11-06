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

// UserFind implements user.UserServer.UserFind
func (s *UserService) UserFind(ctx context.Context, req *grpc_user.UserInfoRequest) (resp *grpc_user.UserInfo, err error) {
	var user *biz.User
	switch {
	case req.Id > 0:
		user, err = s.uc.GetUser(ctx, req.Id)
	case req.Account != "":
		user, err = s.uc.GetUserByAccount(ctx, req.Account)
	default:
		return nil, errors.BadRequest(grpc_user.ErrorReason_INVALID_ARGUMENT.String(), "missing id or account")
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

// UserList implements user.UserServer.UserList
func (s *UserService) UserList(ctx context.Context, req *grpc_user.UserListRequest) (*grpc_user.UserListReply, error) {
	condition := &biz.UserListCondition{
		Page:     req.Page,
		PageSize: req.PageSize,
		Account:  req.Account,
		Name:     req.Name,
		UserType: int(req.Type),
		Status:   int(req.StatusIs),
	}
	count, err := s.uc.CountUser(ctx, condition)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &grpc_user.UserListReply{Total: 0}, nil
	}
	users, err := s.uc.GetUserList(ctx, condition)
	if err != nil {
		return nil, err
	}
	list := make([]*grpc_user.UserInfo, 0, len(users))
	for _, user := range users {
		list = append(list, &grpc_user.UserInfo{
			Id:        user.Id,
			Account:   user.Account,
			Password:  user.Password,
			Name:      user.Name,
			Avatar:    user.Avatar,
			Type:      grpc_user.UserType(user.UserType),
			StatusIs:  grpc_user.UserStatus(user.Status),
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt,
		})
	}
	return &grpc_user.UserListReply{
		Total: count,
		List:  list,
	}, nil
}

// UserSave implements user.UserServer.UserSave
func (s *UserService) UserSave(ctx context.Context, req *grpc_user.UserSaveRequest) (*grpc_user.UserSaveReply, error) {
	user := &biz.User{
		Id:       req.Id,
		Account:  req.Account,
		Password: req.Password,
		Name:     req.Name,
		Avatar:   req.Avatar,
		UserType: int(req.Type),
		Status:   int(req.StatusIs),
	}
	if user.Id > 0 {
		_, err := s.uc.GetUser(ctx, user.Id)
		if err != nil {
			return nil, err
		}
	}
	u, err := s.uc.SaveUser(ctx, user)
	if err != nil {
		return nil, err
	}
	return &grpc_user.UserSaveReply{
		Id: u.Id,
	}, nil
}

// UserDelete implements user.UserServer.UserDelete
func (s *UserService) UserDelete(ctx context.Context, req *grpc_user.UserDeleteRequest) (*grpc_user.UserDeleteReply, error) {
	ids := make([]int64, 0)
	if req.Id > 0 {
		ids = append(ids, req.Id)
	}
	if len(req.Ids) > 0 {
		ids = append(ids, req.Ids...)
	}
	err := s.uc.DeleteUser(ctx, ids)
	if err != nil {
		return nil, err
	}
	return &grpc_user.UserDeleteReply{
		Ids: ids,
	}, nil
}
