package item

import (
	"context"
	stdError "errors"
	"go_zero_example/internal/errorl"
	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"
	model "go_zero_example/model/mongo"
	"go_zero_example/pkg/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
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
		if stdError.Is(err, model.ErrNotFound) {
			return nil, errorl.ItemNotFound(errorx.WithMeta(map[string]interface{}{"ID": req.ID}))
		}
		return nil, errorl.ItemModelFindFailed(errorx.WithMeta(map[string]interface{}{"ID": req.ID}))
	}
	item := types.Item{}
	if err = copier.Copy(&item, &instance); err != nil {
		return nil, errorl.CopyStructFailed(errorx.WithMeta(map[string]interface{}{"err": err}))
	}
	item.ID = instance.ID.Hex()
	return &types.GetItemResp{
		Data: item,
	}, nil
}
