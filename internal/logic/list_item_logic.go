package logic

import (
	"context"

	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListItemLogic {
	return &ListItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListItemLogic) ListItem(req *types.ListItemReq) (resp *types.ListItemResp, err error) {
	// todo: add your logic here and delete this line

	return
}
