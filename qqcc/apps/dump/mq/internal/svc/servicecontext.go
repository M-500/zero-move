package svc

import (
	"qqcc/apps/dump/mq/internal/config"
	"qqcc/pkg/gormx"
)

type ServiceContext struct {
	Config config.Config
	DB     *gormx.DBX
	//ArticleModel model.ArticleModel
	//BizRedis     *redis.Redis
	//UserRPC      user.User
	//Es           *es.Es
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := gormx.MustSetupMySQl(&gormx.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	// TODO: 是否要初始化表结构
	return &ServiceContext{
		Config: c,
		DB:     db,
	}
}
