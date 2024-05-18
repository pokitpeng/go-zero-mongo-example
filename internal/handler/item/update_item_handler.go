package item

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"go_zero_example/internal/logic/item"
	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"
	"go_zero_example/response"
)

func UpdateItemHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UpdateItemReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := item.NewUpdateItemLogic(r.Context(), svcCtx)
		resp, err := l.UpdateItem(&req)
		response.Response(w, resp, err)
	}
}
