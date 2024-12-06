package server

import (
	"context"
	"strings"

	"kratos-project/api/http_app"
	"kratos-project/app/http_app/internal/conf"
	"kratos-project/app/http_app/internal/service"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/auth/jwt"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/ratelimit"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/middleware/tracing"
	"github.com/go-kratos/kratos/v2/middleware/validate"
	"github.com/go-kratos/kratos/v2/transport/http"

	jwtv5 "github.com/golang-jwt/jwt/v5"
)

// NewHTTPServer new an HTTP server.
func NewHTTPServer(c *conf.Server, authC *conf.Auth, appUser *service.AppUserSrv, logger log.Logger) *http.Server {
	var opts = []http.ServerOption{
		http.Middleware(
			// 异常恢复
			recovery.Recovery(),
			// 默认 bbr limiter
			ratelimit.Server(),
			// 日志
			logging.Server(logger),
			// 参数验证
			validate.Validator(),
			// 链路追踪
			tracing.Server(),
			// jwt 验证
			selector.Server(
				jwt.Server(func(token *jwtv5.Token) (interface{}, error) {
					return []byte(authC.Secret), nil
				}),
			).Match(func(ctx context.Context, operation string) bool {
				// 登录接口不需要验证jwt
				if strings.HasPrefix(operation, "/http_app.User/UserLogin") {
					return false
				}
				return true
			}).Build(),
		),
	}
	if c.Http.Network != "" {
		opts = append(opts, http.Network(c.Http.Network))
	}
	if c.Http.Addr != "" {
		opts = append(opts, http.Address(c.Http.Addr))
	}
	if c.Http.Timeout != nil {
		opts = append(opts, http.Timeout(c.Http.Timeout.AsDuration()))
	}
	// 请求响应封装
	opts = append(opts, http.ResponseEncoder(EncoderResponse()))
	// 错误响应封装`
	opts = append(opts, http.ErrorEncoder(EncoderError()))

	srv := http.NewServer(opts...)
	http_app.RegisterUserHTTPServer(srv, appUser)
	return srv
}
