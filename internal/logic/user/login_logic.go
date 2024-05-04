package user

import (
	"context"
	model "go_zero_example/model/mongo"
	"net/http"
	"time"

	"go_zero_example/internal/svc"
	"go_zero_example/internal/types"

	"github.com/golang-jwt/jwt/v4"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/x/errors"
	"google.golang.org/protobuf/proto"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginResp, err error) {
	list, _, err := l.svcCtx.UserModel.List(l.ctx, &model.User{Username: proto.String(req.Username)}, 0, 10, "", "")
	if err != nil || len(list) != 1 {
		return nil, errors.New(http.StatusBadRequest, "用户名不存在")
	}
	dbInstance := list[0]
	if *dbInstance.Password != req.Password {
		return nil, errors.New(http.StatusBadRequest, "用户名或密码错误")
	}
	now := time.Now()
	token, err := generateToken(l.svcCtx.Config.Auth.AccessSecret, now.Unix(), l.svcCtx.Config.Auth.AccessExpire)
	if err != nil {
		return nil, errors.New(http.StatusInternalServerError, "生成令牌失败")
	}
	return &types.LoginResp{
		Data: types.LoginRespData{
			Token:    token,
			ExpireAt: now.Add(time.Duration(l.svcCtx.Config.Auth.AccessExpire) * time.Second).Format("2006-01-02 15:04:05"),
			ID:       dbInstance.ID.Hex(),
			Username: *dbInstance.Username,
		},
	}, nil
}

func generateToken(secret string, now, accessExpire int64) (string, error) {
	claims := make(jwt.MapClaims)
	claims["exp"] = now + accessExpire
	claims["iat"] = now
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}
