package user

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_example/internal/logic/user"
	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"
	"go_zero_example/response"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CreateUserReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := user.NewCreateUserLogic(r.Context(), svcCtx)
		resp, err := l.CreateUser(&req)
		response.Response(w, resp, err)
	}
}
