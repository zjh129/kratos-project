package biz

import (
	"context"
	"time"

	"kratos-project/api/http_admin"
	"kratos-project/app/http_admin/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	jwtv5 "github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
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
	repo     UserRepo
	authConf *conf.Auth
	log      *log.Helper
}

// UserClaims user claims.
type UserClaims struct {
	jwtv5.RegisteredClaims
	UserID  int64  `json:"user_id"`
	Account string `json:"account"`
}

// NewUserUsecase new a User usecase.
func NewUserUsecase(repo UserRepo, authConf *conf.Auth, logger log.Logger) *UserUsecase {
	return &UserUsecase{repo: repo, authConf: authConf, log: log.NewHelper(logger)}
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
	uc.log.WithContext(ctx).Infof("UserLogin: %+v", req)
	accountInfo, err := uc.repo.FindByAccount(ctx, req.Account)
	if err != nil {
		return nil, http_admin.ErrorUserNotFound("用户不存在")
	}
	if accountInfo.Passwd != req.Password {
		return nil, http_admin.ErrorUserLoginFailed("账号或密码错误")
	}
	// generate token
	claims := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, UserClaims{
		RegisteredClaims: jwtv5.RegisteredClaims{
			ExpiresAt: jwtv5.NewNumericDate(time.Now().Add(uc.authConf.Expire.AsDuration())),
			IssuedAt:  jwtv5.NewNumericDate(time.Now()),
			NotBefore: jwtv5.NewNumericDate(time.Now()),
			Issuer:    "http-admin",
			Subject:   "kratos-project http-admin",
			ID:        uuid.NewString(),
			Audience:  []string{"http-admin api"},
		},
		UserID:  accountInfo.Id,
		Account: accountInfo.Account,
	})
	signedString, err := claims.SignedString([]byte(uc.authConf.Secret))
	if err != nil {
		return nil, http_admin.ErrorUserLoginFailed("生成token失败: %s", err.Error())
	}
	return &http_admin.UserLoginReply{
		Token: signedString,
	}, nil
}

// UserLogout user logout.
func (uc *UserUsecase) UserLogout(ctx context.Context, req *http_admin.UserLogoutRequest) (*http_admin.UserLogoutReply, error) {
	uc.log.WithContext(ctx).Infof("UserLogout: %+v", req)
	return &http_admin.UserLogoutReply{}, nil
}
