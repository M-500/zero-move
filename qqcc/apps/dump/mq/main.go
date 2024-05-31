package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"qqcc/apps/dump/mq/internal/config"
	"qqcc/apps/dump/mq/internal/logic"
	"qqcc/apps/dump/mq/internal/svc"
)

// @Description
// @Author 代码小学生王木木
var configFile = flag.String("f", "apps/dump/mq/etc/dump.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	err := c.ServiceConf.SetUp()
	if err != nil {
		panic(err)
	}
	logx.DisableStat()
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()
	serviceGroup := service.NewServiceGroup()
	defer serviceGroup.Stop()

	// 开启消费者
	for _, mq := range logic.Consumers(ctx, svcCtx) {
		serviceGroup.Add(mq)
	}
	serviceGroup.Start()
}
