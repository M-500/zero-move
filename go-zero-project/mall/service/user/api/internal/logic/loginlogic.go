package logic

import (
	"context"
	"mall/pkg/jwtx"
	"mall/service/user/rpc/types/user"
	"time"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
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

func (l *LoginLogic) Login(req *types.LoginRequest) (resp *types.LoginResponse, err error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Mobile:   req.Mobile,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	// 组装JWT Token
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.JwtAuth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.JwtAuth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		AccessToken:  accessToken,
		AccessExpire: accessExpire,
	}, nil
}
