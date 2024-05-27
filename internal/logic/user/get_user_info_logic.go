package user

import (
	"context"

	"go_zero_example/internal/errorl"
	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"
	"go_zero_example/pkg/errorx"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserInfoLogic) GetUserInfo(req *types.GetUserInfoReq) (resp *types.GetUserInfoResp, err error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, req.ID)
	if err != nil {
		return nil, errorl.GetUserInfoFailed(errorx.WithMeta(map[string]interface{}{
			"ID": req.ID,
		}))
	}

	return &types.GetUserInfoResp{
		User: types.User{
			ID:       user.ID.Hex(),
			Username: *user.Username,
			Password: *user.Password,
			Avatar:   *user.Avatar,
			Role:     *user.Role,
			UpdateAt: *user.UpdateAt,
			CreateAt: *user.CreateAt,
		},
	}, nil
}
