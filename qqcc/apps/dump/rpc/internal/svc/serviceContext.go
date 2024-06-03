package svc

import (
	"github.com/zeromicro/go-queue/kq"
	"qqcc/apps/dump/rpc/internal/config"
	"qqcc/apps/dump/rpc/internal/repository/dao"
	"qqcc/pkg/gormx"
)

type ServiceContext struct {
	Config         config.Config
	KqPusherClient *kq.Pusher
	DB             *gormx.DBX
	ParserDAO      dao.ParserDao
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
		Config:         c,
		KqPusherClient: kq.NewPusher(c.KqPusherConf.Brokers, c.KqPusherConf.Topic),
		DB:             db,
		ParserDAO:      dao.NewParserDao(db.DB),
	}
}
