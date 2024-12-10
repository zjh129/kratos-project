package data

import (
	"context"

	"kratos-project/api/grpc_user"
	"kratos-project/api/http_admin"
	"kratos-project/app/http_admin/internal/biz"

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
func (r *userRpcRepo) Create(ctx context.Context, req *http_admin.CreateUserRequest) (int64, error) {
	saveData := &grpc_user.UserSaveRequest{
		Account:  req.Account,
		Password: req.Password,
		Name:     req.Name,
		Avatar:   req.Avatar,
		Type:     grpc_user.UserType_USER_TYPE_NORMAL,
	}
	if req.IsEnable {
		saveData.StatusIs = grpc_user.UserStatus_USER_STATUS_ENABLE
	} else {
		saveData.StatusIs = grpc_user.UserStatus_USER_STATUS_DISABLE
	}
	rs, err := r.data.guClient.UserSave(ctx, saveData)
	if err != nil {
		return 0, err
	}
	return rs.Id, nil
}

// Update . Update a user.
func (r *userRpcRepo) Update(ctx context.Context, req *http_admin.UpdateUserRequest) error {
	saveData := &grpc_user.UserSaveRequest{
		Id:       req.Id,
		Account:  req.Account,
		Password: req.Password,
		Name:     req.Name,
		Avatar:   req.Avatar,
		Type:     grpc_user.UserType_USER_TYPE_NORMAL,
		StatusIs: 0,
	}
	if req.IsEnable {
		saveData.StatusIs = grpc_user.UserStatus_USER_STATUS_ENABLE
	} else {
		saveData.StatusIs = grpc_user.UserStatus_USER_STATUS_DISABLE
	}
	_, err := r.data.guClient.UserSave(ctx, saveData)
	if err != nil {
		return err
	}
	return nil
}

// FindByID . Find a user by id.
func (r *userRpcRepo) FindByID(ctx context.Context, id int64) (*http_admin.UserInfo, error) {
	find, err := r.data.guClient.UserFind(ctx, &grpc_user.UserInfoRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return r.convertUserInfo(find), nil
}

// convertUserInfo . Convert grpc_user.UserInfo to http_admin.UserInfo.
func (r *userRpcRepo) convertUserInfo(u *grpc_user.UserInfo) *http_admin.UserInfo {
	return &http_admin.UserInfo{
		Id:        u.Id,
		Account:   u.Account,
		Name:      u.Name,
		Avatar:    u.Avatar,
		IsEnable:  u.StatusIs == grpc_user.UserStatus_USER_STATUS_ENABLE,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}

// FindByAccount . Find a user by account.
func (r *userRpcRepo) FindByAccount(ctx context.Context, account string) (*biz.UserAccountInfo, error) {
	find, err := r.data.guClient.UserFind(ctx, &grpc_user.UserInfoRequest{Account: account})
	if err != nil {
		return nil, err
	}
	return &biz.UserAccountInfo{
		Id:      find.Id,
		Account: find.Account,
		Passwd:  find.Password,
	}, nil
}

// ListByPage . List users by page.
func (r *userRpcRepo) ListByPage(ctx context.Context, req *http_admin.ListUserRequest) (*http_admin.ListUserReply, error) {
	pbReq := &grpc_user.UserListRequest{
		Page:     req.Page,
		PageSize: req.PageSize,
		Name:     req.Name,
		Type:     0,
		StatusIs: 0,
	}
	pbResp, err := r.data.guClient.UserList(ctx, pbReq)
	if err != nil {
		return nil, err
	}
	if pbResp.Total == 0 {
		return &http_admin.ListUserReply{
			Count: 0,
			List:  nil,
		}, nil
	}
	users := make([]*http_admin.UserInfo, 0, len(pbResp.List))
	for _, u := range pbResp.List {
		users = append(users, r.convertUserInfo(u))
	}
	return &http_admin.ListUserReply{
		List:  users,
		Count: pbResp.Total,
	}, nil
}

// Delete . Delete a user.
func (r *userRpcRepo) Delete(ctx context.Context, id int64) error {
	_, err := r.data.guClient.UserDelete(ctx, &grpc_user.UserDeleteRequest{Id: id})
	if err != nil {
		return err
	}
	return nil
}
