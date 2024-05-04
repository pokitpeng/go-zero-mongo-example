package base

import (
	"context"
	"go_zero_example/api"
	"go_zero_example/internal/svc"
	"io"

	"github.com/zeromicro/go-zero/core/logx"
)

type SwaggerLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
	writer io.Writer // add writer
}

func NewSwaggerLogic(ctx context.Context, svcCtx *svc.ServiceContext, writer io.Writer) *SwaggerLogic {
	return &SwaggerLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
		writer: writer, // add writer
	}
}

func (l *SwaggerLogic) Swagger() error {
	n, err := l.writer.Write(api.Swagger)
	if err != nil {
		return err
	}
	if n < len(api.Swagger) {
		return io.ErrClosedPipe
	}
	return nil
}
