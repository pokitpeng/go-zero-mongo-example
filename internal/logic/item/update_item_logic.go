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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UpdateItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateItemLogic {
	return &UpdateItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateItemLogic) UpdateItem(req *types.UpdateItemReq) (resp *types.UpdateItemResp, err error) {
	_, err = l.svcCtx.ItemModel.FindOne(l.ctx, req.Item.ID)
	if err != nil {
		if stdError.Is(err, model.ErrNotFound) {
			return nil, errorl.ItemNotFound(errorx.WithMeta(map[string]interface{}{"ID": req.ID}))
		}
		return nil, errorl.ItemModelFindFailed(errorx.WithMeta(map[string]interface{}{"ID": req.ID}))
	}
	update := model.Item{}
	if err = copier.Copy(&update, &req.Item); err != nil {
		return nil, errorl.CopyStructFailed(errorx.WithMeta(map[string]interface{}{"err": err}))
	}
	oid, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, errorl.InvalidObjectId(errorx.WithMeta(map[string]interface{}{"err": err}))
	}
	update.ID = oid
	updated, err := l.svcCtx.ItemModel.Update(l.ctx, &update)
	if err != nil {
		return nil, errorl.ItemModelUpdateFailed(errorx.WithMeta(map[string]interface{}{"ID": req.ID, "err": err}))
	}

	return &types.UpdateItemResp{
		Code: 0,
		Msg:  "",
		Data: types.IsOK{
			IsOK: updated.ModifiedCount > 0,
		},
	}, nil
}
