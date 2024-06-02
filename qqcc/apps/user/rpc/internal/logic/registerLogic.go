package logic

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"qqcc/apps/user/rpc/internal/dao"

	"qqcc/apps/user/rpc/internal/svc"
	"qqcc/apps/user/rpc/types/user"

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
	hash, err := bcrypt.GenerateFromPassword([]byte(in.GetPassword()), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	uID, err := l.svcCtx.UserDao.Insert(l.ctx, dao.UserModel{
		Name:     in.GetUsername(),
		Mobile:   in.GetMobile(),
		Gender:   in.GetGender(),
		Password: string(hash),
		IsAdmin:  0,
		Avatar:   "",
		CreateId: 0,
	})
	// 好好 我学到了 errors.Is
	if errors.Is(err, dao.ErrUserDuplicate) {
		return nil, err
	}
	if err != nil {
		l.Logger.Error("[register User RPC]创建用户数据库异常", err)
		return nil, err
	}

	return &user.RegisterResponse{
		Id:       uID,
		Username: in.GetUsername(),
		Gender:   in.GetGender(),
		Mobile:   in.GetMobile(),
	}, nil
}
