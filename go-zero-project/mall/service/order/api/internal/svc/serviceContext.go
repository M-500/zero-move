package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/order/api/internal/config"
	"mall/service/order/rpc/orderClient"
	"mall/service/order/rpc/types/order"
)

type ServiceContext struct {
	Config   config.Config
	OrderRPC order.OrderClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:   c,
		OrderRPC: orderClient.NewOrder(zrpc.MustNewClient(c.OrderRpc)),
	}
}
