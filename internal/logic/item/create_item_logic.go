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

type CreateItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateItemLogic {
	return &CreateItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateItemLogic) CreateItem(req *types.CreateItemReq) (resp *types.CreateItemResp, err error) {
	id := primitive.NewObjectID()
	newInstance := model.Item{}
	if err = copier.Copy(&newInstance, req); err != nil {
		return nil, errors.New(http.StatusInternalServerError, fmt.Sprintf("copy struct error: %v", err))
	}
	newInstance.ID = id
	if err = l.svcCtx.ItemModel.Insert(l.ctx, &newInstance); err != nil {
		return nil, errors.New(http.StatusInternalServerError, fmt.Sprintf("insert error: %v", err))
	}
	return &types.CreateItemResp{
		Data: types.IsOK{
			IsOK: true,
		},
	}, nil
}
