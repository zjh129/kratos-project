package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"kratos-project/api/grpc_user"
	"kratos-project/app/http_app/internal/biz"
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

func (u *userRpcRepo) FindByAccount(ctx context.Context, account string) (*biz.UserAccountInfo, error) {
	find, err := u.data.gu_client.UserFind(ctx, &grpc_user.UserInfoRequest{Account: account})
	if err != nil {
		return nil, err
	}
	return u.convertUserInfo(find), nil
}

func (u *userRpcRepo) FindByID(ctx context.Context, id int64) (*biz.UserAccountInfo, error) {
	find, err := u.data.gu_client.UserFind(ctx, &grpc_user.UserInfoRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return u.convertUserInfo(find), nil
}

// convertUserInfo . Convert grpc_user.UserInfo to http_admin.UserInfo.
func (u *userRpcRepo) convertUserInfo(info *grpc_user.UserInfo) *biz.UserAccountInfo {
	return &biz.UserAccountInfo{
		Id:      info.Id,
		Account: info.Account,
		Passwd:  info.Password,
		Name:    info.Name,
		Avatar:  info.Avatar,
	}
}
