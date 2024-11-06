package biz

import (
	"context"

	"kratos-project/api/grpc_user"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	// ErrUserNotFound is user not found.
	ErrUserNotFound = errors.NotFound(grpc_user.ErrorReason_USER_NOT_FOUND.String(), "user not found")
)

// User is a User model.
type User struct {
	Id        int64  `json:"id"`       // 用户ID
	Account   string `json:"account"`  // 用户账号
	Password  string `json:"password"` // 用户密码
	Name      string `json:"name"`     // 用户名称
	Avatar    string `json:"avatar"`   // 用户头像地址
	UserType  int    `json:"type"`     // 用户类型(1:OA 用户; 2: 普通账号)
	Status    int    `json:"status"`   // 可用状态(1:启用,2:禁用)
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}

type UserListCondition struct {
	Page     int64  // 页码
	PageSize int64  // 每页数量
	Account  string // 账号
	Name     string // 名称
	UserType int    // 类型
	Status   int    // 状态
}

type UserListResp struct {
	Total int64   `json:"total"`
	List  []*User `json:"list"`
}

// UserRepo is a User repo.
type UserRepo interface {
	Save(context.Context, *User) (*User, error)
	FindByID(context.Context, int64) (*User, error)
	FindByAccount(context.Context, string) (*User, error)
	Count(context.Context, *UserListCondition) (int64, error)
	PageList(context.Context, *UserListCondition) ([]*User, error)
	Delete(context.Context, []int64) error
}

// UserUsecase is a User usecase.
type UserUsecase struct {
	repo UserRepo
	log  *log.Helper
}

// NewGreeterUsecase new a Greeter usecase.
func NewGreeterUsecase(repo UserRepo, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, log: log.NewHelper(logger)}
}

// SaveUser creates a User, and returns the new Greeter.
func (uc *UserUsecase) SaveUser(ctx context.Context, u *User) (*User, error) {
	uc.log.WithContext(ctx).Infof("CreateUser: %v", u.Account)
	return uc.repo.Save(ctx, u)
}

// GetUser gets the specified User.
func (uc *UserUsecase) GetUser(ctx context.Context, id int64) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetUser: %v", id)
	return uc.repo.FindByID(ctx, id)
}

// GetUserByAccount gets the specified User by account.
func (uc *UserUsecase) GetUserByAccount(ctx context.Context, account string) (*User, error) {
	uc.log.WithContext(ctx).Infof("GetUserByAccount: %v", account)
	return uc.repo.FindByAccount(ctx, account)
}

// GetUserList gets the specified User list.
func (uc *UserUsecase) GetUserList(ctx context.Context, cond *UserListCondition) ([]*User, error) {
	uc.log.WithContext(ctx).Infof("GetUserList: %v", cond)
	return uc.repo.PageList(ctx, cond)
}

// CountUser counts the number of users.
func (uc *UserUsecase) CountUser(ctx context.Context, cond *UserListCondition) (int64, error) {
	uc.log.WithContext(ctx).Infof("CountUser: %v", cond)
	return uc.repo.Count(ctx, cond)
}

// DeleteUser deletes the specified User.
func (uc *UserUsecase) DeleteUser(ctx context.Context, ids []int64) error {
	uc.log.WithContext(ctx).Infof("DeleteUser: %v", ids)
	return uc.repo.Delete(ctx, ids)
}
