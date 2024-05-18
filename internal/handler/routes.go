// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	base "go_zero_example/internal/handler/base"
	item "go_zero_example/internal/handler/item"
	user "go_zero_example/internal/handler/user"
	"go_zero_example/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/version",
				Handler: base.VersionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/swagger",
				Handler: base.SwaggerHandler(serverCtx),
			},
		},
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/user",
				Handler: user.CreateUserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/user/login",
				Handler: user.LoginHandler(serverCtx),
			},
		},
		rest.WithPrefix("/api"),
	)

	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodPost,
				Path:    "/item",
				Handler: item.CreateItemHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/item/:id",
				Handler: item.GetItemHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/item",
				Handler: item.ListItemHandler(serverCtx),
			},
			{
				Method:  http.MethodPut,
				Path:    "/item",
				Handler: item.UpdateItemHandler(serverCtx),
			},
			{
				Method:  http.MethodDelete,
				Path:    "/item/:id",
				Handler: item.DeleteItemHandler(serverCtx),
			},
		},
		rest.WithJwt(serverCtx.Config.Auth.AccessSecret),
		rest.WithPrefix("/api"),
	)
}