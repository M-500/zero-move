package svc

import (
	"mall/pkg/gormx"
	"mall/service/order/rpc/internal/config"
	"mall/service/order/rpc/repository/dao"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gormx.DBX
	OrderDao dao.OrderDAO
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
		Config:   c,
		DB:       db,
		OrderDao: dao.NewOrderDAO(db.DB),
	}
}
