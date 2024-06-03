package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"qqcc/apps/bff/api/internal/config"
	"qqcc/apps/user/rpc/types/user"
	"qqcc/apps/user/rpc/userClient"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userClient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
