package base

import (
	"net/http"

	"go_zero_example/internal/logic/base"
	"go_zero_example/internal/svc"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func SwaggerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := base.NewSwaggerLogic(r.Context(), svcCtx, w) // add w
		err := l.Swagger()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
