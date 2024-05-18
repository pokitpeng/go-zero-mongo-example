package item

import (
	"context"
	"go_zero_example/internal/errorl"
	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"
	"go_zero_example/pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
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
		return nil, errorl.ItemModelDeleteFailed(errorx.WithMeta(map[string]any{"ID": req.ID, "err": err}))
	}
	return &types.DeleteItemResp{}, nil
}
