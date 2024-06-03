package logic

import (
	"context"
	"fmt"
	"qqcc/apps/user/rpc/types/user"
	"qqcc/pkg/jwtx"
	"time"

	"qqcc/apps/bff/api/internal/svc"
	"qqcc/apps/bff/api/internal/types"

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
	// TODO: 一开始应该校验图片验证码 如果有的话
	fmt.Println(req)
	res, err := l.svcCtx.UserRpc.Login(l.ctx, &user.LoginRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}
	// 组装JWT信息
	now := time.Now().Unix()
	accessExpire := l.svcCtx.Config.Auth.AccessExpire
	accessToken, err := jwtx.GetToken(l.svcCtx.Config.Auth.AccessSecret, now, accessExpire, res.Id)
	if err != nil {
		return nil, err
	}
	return &types.LoginResponse{
		AccessToken:  accessToken,
		AccessExpire: accessExpire,
	}, nil
}
