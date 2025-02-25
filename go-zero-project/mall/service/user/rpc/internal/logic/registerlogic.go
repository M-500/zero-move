package logic

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"mall/service/user/dao"

	"mall/service/user/rpc/internal/svc"
	"mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type RegisterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RegisterLogic) Register(in *user.RegisterRequest) (*user.RegisterResponse, error) {
	userEntity := dao.UserModel{
		Name:     in.Name,
		Mobile:   in.Mobile,
		Gender:   in.Gender,
		Password: in.Password,
	}
	// 加密
	hash, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	userEntity.Password = string(hash)
	res, err := l.svcCtx.UserDao.Insert(l.ctx, userEntity)
	if err != nil {
		return nil, err
	}
	return &user.RegisterResponse{
		Id:     res,
		Name:   in.Name,
		Gender: in.Gender,
		Mobile: in.Mobile,
	}, nil
}
