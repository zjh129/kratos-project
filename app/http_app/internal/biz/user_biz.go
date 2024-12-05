package biz

import (
	"context"
	"kratos-project/api/http_app"
	"time"

	"kratos-project/app/http_app/internal/conf"

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
	FindByAccount(context.Context, string) (*UserAccountInfo, error)
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

// UserLogin user login.
func (uc *UserUsecase) UserLogin(ctx context.Context, req *http_app.UserLoginRequest) (*http_app.UserLoginReply, error) {
	uc.log.WithContext(ctx).Infof("UserLogin: %+v", req)
	accountInfo, err := uc.repo.FindByAccount(ctx, req.Account)
	if err != nil {
		return nil, http_app.ErrorUserNotFound("用户不存在")
	}
	if accountInfo.Passwd != req.Password {
		return nil, http_app.ErrorUserLoginFailed("账号或密码错误")
	}
	// generate token
	claims := jwtv5.NewWithClaims(jwtv5.SigningMethodHS256, UserClaims{
		RegisteredClaims: jwtv5.RegisteredClaims{
			ExpiresAt: jwtv5.NewNumericDate(time.Now().Add(uc.authConf.Expire.AsDuration())),
			IssuedAt:  jwtv5.NewNumericDate(time.Now()),
			NotBefore: jwtv5.NewNumericDate(time.Now()),
			Issuer:    "http-app",
			Subject:   "kratos-project http-app",
			ID:        uuid.NewString(),
			Audience:  []string{"http-app api"},
		},
		UserID:  accountInfo.Id,
		Account: accountInfo.Account,
	})
	signedString, err := claims.SignedString([]byte(uc.authConf.Secret))
	if err != nil {
		return nil, http_app.ErrorUserLoginFailed("生成token失败: %s", err.Error())
	}
	return &http_app.UserLoginReply{
		Token: signedString,
	}, nil
}

// UserLogout user logout.
func (uc *UserUsecase) UserInfo(ctx context.Context) (*http_app.UserInfoReply, error) {
	return &http_app.UserInfoReply{}, nil
}
