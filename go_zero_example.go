package main

import (
	"flag"
	"fmt"
	"go_zero_example/internal/config"
	"go_zero_example/internal/handler"
	"go_zero_example/internal/svc"
	"go_zero_example/response"
	"go_zero_example/version"
	"os"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
)

var configFile = flag.String("f", "etc/go_zero_example-api.yaml", "the config file")
var showVersion = flag.Bool("v", false, "show version")

func main() {
	flag.Parse()
	if *showVersion {
		fmt.Println(version.New())
		os.Exit(0)
	}

	logx.DisableStat() // 关闭stat日志

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, response.UnauthorizedCallback())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
