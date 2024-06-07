package svc

import (
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"qqcc/apps/dump/mq/internal/config"
	"qqcc/apps/dump/rpc/dumpClient"
	"qqcc/apps/dump/rpc/types/dump"

	"qqcc/pkg/gormx"
)

type ServiceContext struct {
	Config             config.Config
	DB                 *gormx.DBX
	DumpRpc            dump.DumpClient
	GsBasePusherClient *kq.Pusher
	BizRedis           redis.Cmdable
	//UserRPC      user.User
	//Es           *es.Es
}

func NewServiceContext(c config.Config) *ServiceContext {
	db := gormx.MustSetupMySQl(&gormx.Config{
		DSN:          c.DB.DataSource,
		MaxOpenConns: c.DB.MaxOpenConns,
		MaxIdleConns: c.DB.MaxIdleConns,
		MaxLifetime:  c.DB.MaxLifetime,
	})
	rdb := redis.NewClient(&redis.Options{
		Addr:     c.BizRedis.Host,
		Password: c.BizRedis.Pass,
	})
	// TODO: 是否要初始化表结构
	return &ServiceContext{
		Config:             c,
		DumpRpc:            dumpClient.NewDump(zrpc.MustNewClient(c.DumpRpc)), // 调用Dump服务用到的接口
		DB:                 db,
		GsBasePusherClient: kq.NewPusher(c.KqCompanyPusherConf.Brokers, c.KqCompanyPusherConf.Topic),
		BizRedis:           rdb,
	}
}
