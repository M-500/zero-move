package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/rest/httpx"
	"qqcc/pkg/mdl"
	"qqcc/pkg/xcode"

	"qqcc/apps/bff/api/internal/config"
	"qqcc/apps/bff/api/internal/handler"
	"qqcc/apps/bff/api/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "apps/bff/api/etc/bFF.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithNotAllowedHandler(mdl.NewCorsMiddleware().Handler()))
	defer server.Stop()

	server.Use(mdl.NewCorsMiddleware().Handle)
	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	// 注册自定义错误处理方法
	httpx.SetErrorHandler(xcode.ErrHandler)
	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
