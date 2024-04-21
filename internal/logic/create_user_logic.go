package logic

import (
	"context"
	"fmt"
	"net/http"

	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"
	model "go_zero_example/model/mongo"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
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
	id := primitive.NewObjectID()
	err = l.svcCtx.UserModel.Insert(l.ctx, &model.User{
		ID:       id,
		Username: proto.String(req.Username),
		Password: proto.String(req.Password),
	})
	if err != nil {
		return nil, errors.New(http.StatusBadRequest, fmt.Sprintf("create user error: %v", err))
	}
	return &types.CreateUserResp{
		Data: types.IsOK{
			IsOK: true,
		},
	}, nil
}
