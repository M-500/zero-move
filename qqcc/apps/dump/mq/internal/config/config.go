package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
)

type Config struct {
	service.ServiceConf

	KqConsumerConf kq.KqConf
	DB             struct {
		DataSource   string
		MaxOpenConns int `json:",default=10"`
		MaxIdleConns int `json:",default=100"`
		MaxLifetime  int `json:",default=3600"`
	}
	//ArticleKqConsumerConf kq.KqConf
	//Datasource            string
	//BizRedis              redis.RedisConf
	//// es config
	//Es struct {
	//	Addresses []string
	//	Username  string
	//	Password  string
	//}
	//UserRPC zrpc.RpcClientConf
}
