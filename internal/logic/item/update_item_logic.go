package item

import (
	"context"
	"fmt"
	model "go_zero_example/model/mongo"
	"net/http"

	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
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
		return nil, errors.New(http.StatusNotFound, fmt.Sprintf("item not found: %v", err))
	}
	update := model.Item{}
	if err = copier.Copy(&update, &req.Item); err != nil {
		return nil, errors.New(http.StatusInternalServerError, fmt.Sprintf("copy item error: %v", err))
	}
	oid, err := primitive.ObjectIDFromHex(req.ID)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, fmt.Sprintf("invalid id: %v", err))
	}
	update.ID = oid
	updated, err := l.svcCtx.ItemModel.Update(l.ctx, &update)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, fmt.Sprintf("update item error: %v", err))
	}

	return &types.UpdateItemResp{
		Code: 0,
		Msg:  "",
		Data: types.IsOK{
			IsOK: updated.ModifiedCount > 0,
		},
	}, nil
}
