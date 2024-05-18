package base

import (
	"net/http"

	"go_zero_example/internal/logic/base"
	"go_zero_example/internal/svc"
	"go_zero_example/response"
)

func SwaggerHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		l := base.NewSwaggerLogic(r.Context(), svcCtx, w)
		err := l.Swagger()
		response.Response(w, nil, err)
	}
}
