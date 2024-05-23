package svc

import (
	"mall/pkg/gormx"
	"mall/service/user/dao"
	"mall/service/user/rpc/internal/config"
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
	err := dao.InitTable(db.DB)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:  c,
		DB:      db,
		UserDao: dao.NewUserDao(db.DB),
	}
}
