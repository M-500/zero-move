package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"mall/service/product/api/internal/config"
	"mall/service/product/rpc/productclient"
	"mall/service/product/rpc/types/product"
)

type ServiceContext struct {
	Config  config.Config
	ProdRPC product.ProductClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		ProdRPC: productclient.NewProduct(zrpc.MustNewClient(c.ProdRPC)),
	}
}
