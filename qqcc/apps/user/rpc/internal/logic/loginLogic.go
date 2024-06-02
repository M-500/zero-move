package logic

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/status"

	"qqcc/apps/user/rpc/internal/svc"
	"qqcc/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type LoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *LoginLogic) Login(in *user.LoginRequest) (*user.LoginResponse, error) {
	res, err := l.svcCtx.UserDao.FindByUserName(l.ctx, in.Username)
	if err != nil {
		return nil, err
	}
	// 2. 校验密码
	err = bcrypt.CompareHashAndPassword([]byte(res.Password), []byte(in.Password))
	if err != nil {
		return nil, status.Error(500, "密码不正确")
	}
	return &user.LoginResponse{
		Id:       int64(res.ID),
		Username: res.Name,
	}, nil
}
