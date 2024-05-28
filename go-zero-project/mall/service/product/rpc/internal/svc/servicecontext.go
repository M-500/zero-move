package svc

import (
	"mall/pkg/gormx"
	"mall/service/product/rpc/internal/config"
	"mall/service/product/rpc/repository/dao"
)

type ServiceContext struct {
	Config     config.Config
	DB         *gormx.DBX
	ProductDao dao.ProductDao
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
		Config:     c,
		DB:         db,
		ProductDao: dao.NewProductDao(db.DB),
	}
}
