package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_example/internal/logic/user"
	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"
	"go_zero_example/response"
)

func LoginHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoginReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewLoginLogic(r.Context(), svcCtx)
		resp, err := l.Login(&req)
		response.Response(w, resp, err)
	}
}
