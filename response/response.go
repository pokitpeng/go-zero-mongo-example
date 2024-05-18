package response

import (
	"errors"
	"go_zero_example/pkg/errorx"
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type Body struct {
	Code int64  `json:"Code"`
	Msg  string `json:"Msg"`
	Data any    `json:"Data,omitempty"`
}

func Response(w http.ResponseWriter, resp interface{}, err error) {
	var body Body
	if err != nil {
		var e *errorx.Errorx
		switch {
		case errors.As(err, &e):
			body.Code = e.Code
			body.Msg = e.Msg
		default:
			body.Code = http.StatusInternalServerError
			body.Msg = "Unknown Error"
		}
	} else {
		body.Code = 0
		body.Msg = "OK"
		body.Data = resp
	}
	httpx.OkJson(w, body)
}
