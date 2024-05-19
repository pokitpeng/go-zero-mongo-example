package response

import (
	"go_zero_example/pkg/errorx"
	"net/http"

	"github.com/zeromicro/go-zero/rest"
)

func UnauthorizedCallback() rest.RunOption {
	return rest.WithUnauthorizedCallback(func(w http.ResponseWriter, r *http.Request, err error) {
		e := errorx.New(401, "Unauthorized")
		Response(w, nil, e)
	})
}
