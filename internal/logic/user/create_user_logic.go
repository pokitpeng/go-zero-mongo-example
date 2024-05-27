package user

import (
	"context"
	"go_zero_example/internal/errorl"
	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"
	model "go_zero_example/model/mongo"
	"go_zero_example/pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/proto"
)

type CreateUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateUserLogic {
	return &CreateUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateUserLogic) CreateUser(req *types.CreateUserReq) (resp *types.CreateUserResp, err error) {
	find, _ := l.svcCtx.UserModel.Find(l.ctx, &model.User{Username: proto.String(req.Username)})
	if len(find) > 0 {
		return nil, errorl.UserAlreadyExist(errorx.WithMeta(map[string]any{"Username": req.Username}))
	}
	id := primitive.NewObjectID()
	err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		ID:       id,
		Username: proto.String(req.Username),
		Password: proto.String(req.Password),
		Avatar:   proto.String(req.Avatar),
		Role:     proto.String(req.Role),
	})
	if err != nil {
		return nil, errorl.UserModelInsertFailed(errorx.WithMeta(map[string]any{"ID": id.Hex(), "err": err}))
	}
	return &types.CreateUserResp{}, nil
}
