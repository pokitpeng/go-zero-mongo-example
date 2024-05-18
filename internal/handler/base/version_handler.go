package base

import (
	"net/http"

	"go_zero_example/internal/logic/base"
	"go_zero_example/internal/svc"
	"go_zero_example/response"
)

func VersionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := base.NewVersionLogic(r.Context(), svcCtx)
		resp, err := l.Version()
		response.Response(w, resp, err)
	}
}
