package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	DumpRpc               zrpc.RpcClientConf
	KqConsumerConf        kq.KqConf // 消费任务的消费者
	KqCompanyConsumerConf kq.KqConf // 消费信息的消费者
	KqCompanyPusherConf   struct {
		Brokers []string
		Topic   string
	} // 工商基本信息的生产者
	DB struct {
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
