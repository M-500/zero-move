package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/pay/api/internal/config"
	"mall/service/pay/rpc/payClient"
	"mall/service/pay/rpc/types/pay"
)

type ServiceContext struct {
	Config config.Config
	PayRpc pay.PayClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		PayRpc: payClient.NewPay(zrpc.MustNewClient(c.PayRpc)),
	}
}
