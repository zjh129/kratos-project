package data

import (
	"context"

	"kratos-project/api/grpc_user"
	"kratos-project/api/mq_consume"
	"kratos-project/app/mq_consume/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRpcRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRpcRepo .
func NewUserRpcRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRpcRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

// Create . Create a new user.
func (r *userRpcRepo) Create(ctx context.Context, req *mq_consume.MqUserInfo) (int64, error) {
	saveData := &grpc_user.UserSaveRequest{
		Account:  req.Account,
		Password: req.Password,
		Name:     req.Name,
		Avatar:   req.Avatar,
		Type:     grpc_user.UserType_USER_TYPE_NORMAL,
		StatusIs: grpc_user.UserStatus_USER_STATUS_ENABLE,
	}
	rs, err := r.data.guClient.UserSave(ctx, saveData)
	if err != nil {
		return 0, err
	}
	return rs.Id, nil
}

// Update . Update a user.
func (r *userRpcRepo) Update(ctx context.Context, req *mq_consume.MqUserInfo) error {
	saveData := &grpc_user.UserSaveRequest{
		Id:       req.Id,
		Account:  req.Account,
		Password: req.Password,
		Name:     req.Name,
		Avatar:   req.Avatar,
		Type:     grpc_user.UserType_USER_TYPE_NORMAL,
		StatusIs: grpc_user.UserStatus_USER_STATUS_ENABLE,
	}
	_, err := r.data.guClient.UserSave(ctx, saveData)
	if err != nil {
		return err
	}
	return nil
}

// convertUserInfo . Convert grpc_user.UserInfo to mq_consume.MqUserInfo.
func (r *userRpcRepo) convertUserInfo(u *grpc_user.UserInfo) *mq_consume.MqUserInfo {
	return &mq_consume.MqUserInfo{
		Id:       u.Id,
		Account:  u.Account,
		Password: u.Password,
		Name:     u.Name,
		Avatar:   u.Avatar,
	}
}

// FindByAccount . Find a user by account.
func (r *userRpcRepo) FindByAccount(ctx context.Context, account string) (*mq_consume.MqUserInfo, error) {
	find, _ := r.data.guClient.UserFind(ctx, &grpc_user.UserInfoRequest{Account: account})
	if find == nil {
		return nil, nil
	}
	return r.convertUserInfo(find), nil
}
