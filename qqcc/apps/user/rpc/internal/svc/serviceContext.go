package svc

import (
	"qqcc/apps/user/rpc/internal/config"
	"qqcc/apps/user/rpc/internal/dao"
	"qqcc/pkg/gormx"
)

type ServiceContext struct {
	Config  config.Config
	DB      *gormx.DBX
	UserDao dao.UserDao
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := gormx.MustSetupMySQl(&gormx.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	return &ServiceContext{
		Config:  c,
		DB:      db,
		UserDao: dao.NewUserDao(db.DB),
	}
}
