package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall/pkg/gormx"
	"mall/service/order/rpc/orderClient"
	"mall/service/order/rpc/types/order"
	"mall/service/pay/rpc/internal/config"
	"mall/service/pay/rpc/rpository/dao"
	"mall/service/user/rpc/types/user"
	"mall/service/user/rpc/userclient"
)

type ServiceContext struct {
	Config   config.Config
	DB       *gormx.DBX
	PayDao   dao.PayDao
	UserRpc  user.UserClient   // 用户微服务
	OrderRpc order.OrderClient // 订单微服务
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := gormx.MustSetupMySQl(&gormx.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	err := dao.InitTables(db.DB)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:   c,
		DB:       db,
		PayDao:   dao.NewPayDao(db.DB),
		UserRpc:  userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		OrderRpc: orderClient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
