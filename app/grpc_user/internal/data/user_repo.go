package data

import (
	"context"
	"fmt"
	"time"

	"kratos-project/api/grpc_user"
	"kratos-project/app/grpc_user/internal/biz"
	"kratos-project/app/grpc_user/internal/data/ent"
	"kratos-project/app/grpc_user/internal/data/ent/userinfo"

	"github.com/go-kratos/kratos/v2/log"
	cache "github.com/mgtv-tech/jetcache-go"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "repo/user")),
	}
}

func (u *userRepo) getCacheKey(id int64) string {
	return fmt.Sprintf("kratos_learn:user:%d", id)
}

func (u *userRepo) Save(ctx context.Context, req *grpc_user.UserSaveRequest) (*grpc_user.UserSaveReply, error) {
	var err error
	var get *ent.UserInfo
	if req.Id > 0 {
		get, err = u.data.db.UserInfo.Get(ctx, int(req.Id))
		if err != nil {
			return nil, err
		}
		ur := get.Update().
			SetAccount(req.Account).
			SetPassword(req.Password).
			SetName(req.Name).
			SetAvatar(req.Avatar).
			SetType(int32(req.Type)).
			SetStatusIs(int32(req.StatusIs))
		get, err = ur.SetUpdatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		cr := u.data.db.UserInfo.Create().
			SetAccount(req.Account).
			SetPassword(req.Password).
			SetName(req.Name).
			SetAvatar(req.Avatar).
			SetType(int32(req.Type)).
			SetStatusIs(int32(req.StatusIs))
		get, err = cr.SetCreatedAt(time.Now()).
			SetUpdatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	}
	// cache
	cacheKey := u.getCacheKey(int64(get.ID))
	err = u.data.cache.Set(ctx, cacheKey, cache.Value(get), cache.TTL(time.Hour))
	if err != nil {
		return nil, err
	}
	return &grpc_user.UserSaveReply{
		Id: int64(get.ID),
	}, nil
}

// FindByID 根据ID查找用户
func (u *userRepo) FindById(ctx context.Context, id int64) (*grpc_user.UserInfo, error) {
	cacheKey := u.getCacheKey(id)
	getu := &ent.UserInfo{}
	if err := u.data.cache.Once(ctx, cacheKey, cache.Value(getu), cache.TTL(time.Hour), cache.Refresh(true),
		cache.Do(func(ctx context.Context) (any, error) {
			get, err := u.data.db.UserInfo.Get(ctx, int(id))
			if err != nil {
				if ent.IsNotFound(err) {
					return nil, ErrRecordNotFound
				}
				return nil, err
			}
			return get, nil
		})); err != nil {
		return nil, err
	}
	return u.convertUserInfoToBizUser(getu), nil
}

// FindByAccount 根据账号查找用户
func (u *userRepo) FindByAccount(ctx context.Context, account string) (*grpc_user.UserInfo, error) {
	cacheKey := fmt.Sprintf("kratos_learn:user:account:%s", account)
	// 从缓存中获取用户ID
	uid := 0
	if err := u.data.cache.Once(ctx, cacheKey, cache.Value(&uid), cache.TTL(time.Hour), cache.Refresh(true),
		cache.Do(func(ctx context.Context) (any, error) {
			getU, err := u.data.db.UserInfo.Query().Where(userinfo.Account(account)).First(ctx)
			if err != nil {
				return nil, err
			}
			return getU.ID, nil
		})); err != nil {
		return nil, err
	}
	if uid == 0 {
		return nil, ErrRecordNotFound
	}
	return u.FindById(ctx, int64(uid))
}

// 转换UserInfo到grpc_user.UserInfo
func (u *userRepo) convertUserInfoToBizUser(user *ent.UserInfo) *grpc_user.UserInfo {
	return &grpc_user.UserInfo{
		Id:        int64(user.ID),
		Account:   user.Account,
		Password:  user.Password,
		Name:      user.Name,
		Avatar:    user.Avatar,
		Type:      grpc_user.UserType(user.Type),
		StatusIs:  grpc_user.UserStatus(user.StatusIs),
		CreatedAt: user.CreatedAt.Format(time.DateTime),
		UpdatedAt: user.UpdatedAt.Format(time.DateTime),
	}
}

func (u *userRepo) PageList(ctx context.Context, req *grpc_user.UserListRequest) (*grpc_user.UserListReply, error) {
	// 组装查询条件
	query := u.data.db.UserInfo.Query()
	if req != nil {
		if req.Name != "" {
			query.Where(userinfo.Name(req.Name))
		}
		if req.Type != 0 {
			query.Where(userinfo.Type(int32(req.Type)))
		}
		if req.StatusIs != 0 {
			query.Where(userinfo.StatusIs(int32(req.StatusIs)))
		}
	}
	count, err := query.Count(ctx)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return &grpc_user.UserListReply{Total: 0}, nil
	}
	if req.Page > 0 && req.PageSize > 0 {
		query.Offset(int((req.Page - 1) * req.PageSize)).Limit(int(req.PageSize))
	}
	users, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	var result []*grpc_user.UserInfo
	for _, user := range users {
		result = append(result, u.convertUserInfoToBizUser(user))
	}
	return &grpc_user.UserListReply{
		List:  result,
		Total: int64(count),
	}, nil
}

func (u *userRepo) Delete(ctx context.Context, req *grpc_user.UserDeleteRequest) (*grpc_user.UserDeleteReply, error) {
	// 转换被删除的id类型为int
	idsInt := make([]int, 0)
	if req.Id > 0 {
		idsInt = append(idsInt, int(req.Id))
	}
	if req.Ids != nil {
		for _, id := range req.Ids {
			idsInt = append(idsInt, int(id))
		}
	}
	_, err := u.data.db.UserInfo.Delete().Where(userinfo.IDIn(idsInt...)).Exec(ctx)
	if err != nil {
		return nil, err
	}
	for _, id := range idsInt {
		cacheKey := u.getCacheKey(int64(id))
		err = u.data.cache.Delete(ctx, cacheKey)
		if err != nil {
			return nil, err
		}
	}
	// 转换被删除的id类型为int64
	idsInt64 := make([]int64, 0)
	for _, id := range idsInt {
		idsInt64 = append(idsInt64, int64(id))
	}
	return &grpc_user.UserDeleteReply{
		Ids: idsInt64,
	}, nil
}
