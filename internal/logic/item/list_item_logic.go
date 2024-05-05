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
	"google.golang.org/protobuf/proto"
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
	filter := new(model.Item)
	if req.Name != "" {
		filter.Name = proto.String(req.Name)
	}
	offset := (req.Page - 1) * req.Size
	limit := req.Size
	orderBy := req.OrderBy
	order := req.Order
	items, total, err := l.svcCtx.ItemModel.List(l.ctx, filter, offset, limit, orderBy, order)
	if err != nil {
		return nil, errorl.ItemModelListFailed(errorx.WithMeta(map[string]interface{}{"err": err}))
	}
	respItems := []*types.Item{}
	for _, item := range items {
		t := types.Item{}
		if err = copier.Copy(&t, &item); err != nil {
			return nil, errorl.CopyStructFailed(errorx.WithMeta(map[string]interface{}{"err": err}))
		}
		t.ID = item.ID.Hex()
		respItems = append(respItems, &t)
	}

	return &types.ListItemResp{
		Code:  0,
		Msg:   "",
		Data:  respItems,
		Total: total,
	}, nil
}
