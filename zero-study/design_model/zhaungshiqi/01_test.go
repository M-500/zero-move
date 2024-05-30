package zhaungshiqi

import (
	"context"
	"encoding/json"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
type UserModel struct {
	Username string
	Age      int
}

type UserDao interface {
	QueryUserList(ctx context.Context, username string) ([]UserModel, error)
}

type gormUserDao struct {
	DB     *gorm.DB
	Client redis.Cmdable //集成redis的客户端
}

func (dao *gormUserDao) QueryUserList(ctx context.Context, username string) ([]UserModel, error) {
	var users []UserModel

	// 编写缓存的逻辑
	cacheRes, err2 := dao.Client.Get(ctx, "user_cache:+username").Result()
	if err2 != nil {
		//
	}
	err2 = json.Unmarshal([]byte(cacheRes), &users)
	if err2 == nil {
		return users, nil
	}
	// 缓存中查询不到 再去DB中查询

	err := dao.DB.WithContext(ctx).Where("username = ?", username).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}
