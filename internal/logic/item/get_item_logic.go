package item

import (
	"context"
	"fmt"
	"net/http"

	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
)

type GetItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetItemLogic {
	return &GetItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetItemLogic) GetItem(req *types.GetItemReq) (resp *types.GetItemResp, err error) {
	instance, err := l.svcCtx.ItemModel.FindOne(l.ctx, req.ID)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, fmt.Sprintf("get item error: %v", err))
	}
	item := types.Item{}
	if err = copier.Copy(&item, &instance); err != nil {
		return nil, errors.New(http.StatusInternalServerError, fmt.Sprintf("copy item error: %v", err))
	}
	item.ID = instance.ID.Hex()
	return &types.GetItemResp{
		Data: item,
	}, nil
}
