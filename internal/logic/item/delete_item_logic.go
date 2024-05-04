package item

import (
	"context"
	"fmt"
	"net/http"

	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type DeleteItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDeleteItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteItemLogic {
	return &DeleteItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteItemLogic) DeleteItem(req *types.DeleteItemReq) (resp *types.DeleteItemResp, err error) {
	_, err = l.svcCtx.ItemModel.Delete(l.ctx, req.ID, false)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, fmt.Sprintf("delete item error: %v", err))
	}
	return &types.DeleteItemResp{
		Code: 0,
		Msg:  "ok",
		Data: types.IsOK{
			IsOK: true,
		},
	}, nil
}
