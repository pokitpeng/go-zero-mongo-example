package item

import (
	"context"
	"go_zero_example/internal/errorl"
	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"
	model "go_zero_example/model/mongo"
	"go_zero_example/pkg/errorx"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/proto"
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
	find, _ := l.svcCtx.ItemModel.Find(l.ctx, &model.Item{Name: proto.String(req.Name)})
	if len(find) > 0 {
		return nil, errorl.ItemAlreadyExist(errorx.WithMeta(map[string]any{"Name": req.Name}))
	}

	id := primitive.NewObjectID()
	newInstance := model.Item{}
	if err = copier.Copy(&newInstance, req); err != nil {
		return nil, errorl.CopyStructFailed(errorx.WithMeta(map[string]interface{}{"err": err}))
	}
	newInstance.ID = id
	if err = l.svcCtx.ItemModel.Insert(l.ctx, &newInstance); err != nil {
		return nil, errorl.ItemModelInsertFailed(errorx.WithMeta(map[string]interface{}{"err": err}))
	}
	return &types.CreateItemResp{
		Data: types.IsOK{
			IsOK: true,
		},
	}, nil
}
