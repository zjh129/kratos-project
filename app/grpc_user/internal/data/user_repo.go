package data

import (
	"context"
	"fmt"
	cache "github.com/mgtv-tech/jetcache-go"
	"kratos-project/app/grpc_user/internal/data/ent"
	"kratos-project/app/grpc_user/internal/data/ent/userinfo"
	"time"

	"kratos-project/app/grpc_user/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *userRepo) getCacheKey(id int64) string {
	return fmt.Sprintf("kratos_learn:user:%d", id)
}

func (u *userRepo) Save(ctx context.Context, user *biz.User) (*biz.User, error) {
	var err error
	var get *ent.UserInfo
	if user.Id > 0 {
		get, err = u.data.db.UserInfo.Get(ctx, int(user.Id))
		if err != nil {
			return nil, err
		}
		ur := get.Update().
			SetAccount(user.Account).
			SetPassword(user.Password).
			SetName(user.Name).
			SetAvatar(user.Avatar).
			SetType(int32(user.UserType)).
			SetStatusIs(int32(user.Status))
		get, err = ur.SetUpdatedAt(time.Now()).
			Save(ctx)
		if err != nil {
			return nil, err
		}
	} else {
		cr := u.data.db.UserInfo.Create().
			SetAccount(user.Account).
			SetPassword(user.Password).
			SetName(user.Name).
			SetAvatar(user.Avatar).
			SetType(int32(user.UserType)).
			SetStatusIs(int32(user.Status))
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
	return u.convertUserInfoToBizUser(get), nil
}

func (u *userRepo) FindByID(ctx context.Context, i int64) (*biz.User, error) {
	cacheKey := u.getCacheKey(i)
	getu := &ent.UserInfo{}
	if err := u.data.cache.Once(ctx, cacheKey, cache.Value(getu), cache.TTL(time.Hour), cache.Refresh(true),
		cache.Do(func(ctx context.Context) (any, error) {
			get, err := u.data.db.UserInfo.Get(ctx, int(i))
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

// 转换UserInfo到biz.User
func (u *userRepo) convertUserInfoToBizUser(user *ent.UserInfo) *biz.User {
	return &biz.User{
		Id:        int64(user.ID),
		Account:   user.Account,
		Password:  user.Password,
		Name:      user.Name,
		Avatar:    user.Avatar,
		UserType:  int(user.Type),
		Status:    int(user.StatusIs),
		CreatedAt: user.CreatedAt.Format(time.DateTime),
		UpdatedAt: user.UpdatedAt.Format(time.DateTime),
	}
}

func (u *userRepo) FindByAccount(ctx context.Context, s string) (*biz.User, error) {
	cacheKey := fmt.Sprintf("kratos_learn:user:account:%s", s)
	getu := &ent.UserInfo{}
	if err := u.data.cache.Once(ctx, cacheKey, cache.Value(getu), cache.TTL(time.Hour), cache.Refresh(true),
		cache.Do(func(ctx context.Context) (any, error) {
			get, err := u.data.db.UserInfo.Query().Where(userinfo.Account(s)).First(ctx)
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

// 组装查询条件
func (u *userRepo) getQuery(condition *biz.UserListCondition) *ent.UserInfoQuery {
	query := u.data.db.UserInfo.Query()
	if condition != nil {
		if condition.Name != "" {
			query.Where(userinfo.Name(condition.Name))
		}
		if condition.UserType != 0 {
			query.Where(userinfo.Type(int32(condition.UserType)))
		}
		if condition.Status != 0 {
			query.Where(userinfo.StatusIs(int32(condition.Status)))
		}
	}
	return query
}

func (u *userRepo) Count(ctx context.Context, condition *biz.UserListCondition) (int64, error) {
	count, err := u.getQuery(condition).Count(ctx)
	if err != nil {
		return 0, err
	}
	return int64(count), nil
}

func (u *userRepo) PageList(ctx context.Context, condition *biz.UserListCondition) ([]*biz.User, error) {
	// 组装查询条件
	query := u.getQuery(condition)
	if condition.Page > 0 && condition.PageSize > 0 {
		query.Offset(int((condition.Page - 1) * condition.PageSize)).Limit(int(condition.PageSize))
	}
	users, err := query.All(ctx)
	if err != nil {
		return nil, err
	}
	var result []*biz.User
	for _, user := range users {
		result = append(result, u.convertUserInfoToBizUser(user))
	}
	return result, nil
}

func (u *userRepo) Delete(ctx context.Context, ids []int64) error {
	idsInt := make([]int, 0, len(ids))
	for _, id := range ids {
		idsInt = append(idsInt, int(id))
	}
	_, err := u.data.db.UserInfo.Delete().Where(userinfo.IDIn(idsInt...)).Exec(ctx)
	if err != nil {
		return err
	}
	for _, id := range ids {
		cacheKey := u.getCacheKey(id)
		err = u.data.cache.Delete(ctx, cacheKey)
		if err != nil {
			return err
		}
	}
	return nil
}
