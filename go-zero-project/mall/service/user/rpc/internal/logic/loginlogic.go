package logic

import (
	"context"
	"golang.org/x/crypto/bcrypt"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/types/user"

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
	res, err := l.svcCtx.UserDao.FindByUserName(l.ctx, in.GetMobile())
	if err != nil {
		return nil, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(in.Password), []byte(res.Password))
	//if err != nil {
	//	return nil, errors.New("密码不对")
	//}
	return &user.LoginResponse{
		Id:     int64(res.ID),
		Name:   res.Name,
		Gender: res.Gender,
		Mobile: res.Mobile,
	}, nil
}
