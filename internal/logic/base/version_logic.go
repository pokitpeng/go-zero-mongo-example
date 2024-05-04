package base

import (
	"context"

	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type VersionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewVersionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *VersionLogic {
	return &VersionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *VersionLogic) Version() (resp *types.VersionResp, err error) {
	return &types.VersionResp{
		Code: 0,
		Msg:  "",
		Data: types.Version{
			Version: "v0.0.0",
		},
	}, nil
}
