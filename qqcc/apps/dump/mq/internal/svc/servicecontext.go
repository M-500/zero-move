package svc

import (
	"github.com/olivere/elastic/v7"
	"github.com/redis/go-redis/v9"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/zrpc"
	"qqcc/apps/dump/mq/internal/config"
	"qqcc/apps/dump/mq/internal/repository/dao"
	"qqcc/apps/dump/mq/internal/repository/repo"
	"qqcc/apps/dump/rpc/dumpClient"
	"qqcc/apps/dump/rpc/types/dump"
	"qqcc/pkg/es"

	"qqcc/pkg/gormx"
)

type ServiceContext struct {
	Config             config.Config
	DB                 *gormx.DBX
	DumpRpc            dump.DumpClient
	GsBasePusherClient *kq.Pusher
	BizRedis           redis.Cmdable
	EntPriRepo         repo.EnterpriseRepo
	Es                 *elastic.Client
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
	esClient := es.MustNewEs(&es.Config{
		Url:   c.Es.Url,
		Sniff: c.Es.Sniff,
	})
	gormDao := dao.NewEnterpriseDAO(db.DB)
	esDao := dao.NewEsEnterpriseDao(esClient)
	// 初始化表结构
	err := dao.InitTables(db.DB)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{
		Config:             c,
		DumpRpc:            dumpClient.NewDump(zrpc.MustNewClient(c.DumpRpc)), // 调用Dump服务用到的接口
		DB:                 db,
		GsBasePusherClient: kq.NewPusher(c.KqCompanyPusherConf.Brokers, c.KqCompanyPusherConf.Topic),
		BizRedis:           rdb,
		EntPriRepo:         repo.NewEnterpriseRepo(gormDao, esDao),
		Es:                 esClient,
	}
}
