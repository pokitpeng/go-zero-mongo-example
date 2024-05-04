package main

import (
	"flag"
	"fmt"
	"go/types"
	"go_zero_example/internal/config"
	"go_zero_example/internal/handler"
	"go_zero_example/internal/svc"
	"net/http"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/rest/httpx"
	"github.com/zeromicro/x/errors"
)

var configFile = flag.String("f", "etc/go_zero_example-api.yaml", "the config file")

type BaseResponse[T any] struct {
	Code int    `json:"Code"`
	Msg  string `json:"Msg"`
	Data T      `json:"Data,omitempty"`
}

func main() {
	flag.Parse()

	logx.DisableStat() // 关闭stat日志

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)
	// 自定义错误返回
	httpx.SetErrorHandler(func(err error) (int, any) {
		switch e := err.(type) {
		case *errors.CodeMsg:
			return http.StatusOK, BaseResponse[types.Nil]{
				Code: e.Code,
				Msg:  e.Msg,
			}
		default:
			return http.StatusInternalServerError, err
		}
	})

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
