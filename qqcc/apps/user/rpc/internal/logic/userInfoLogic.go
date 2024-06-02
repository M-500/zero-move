package logic

import (
	"context"
	"google.golang.org/grpc/status"

	"qqcc/apps/user/rpc/internal/svc"
	"qqcc/apps/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UserInfoLogic) UserInfo(in *user.UserInfoRequest) (*user.UserInfoResponse, error) {
	// todo: add your logic here and delete this line
	res, err := l.svcCtx.UserDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, status.Error(500, "用户不存在")
	}
	return &user.UserInfoResponse{
		Id:       int64(res.ID),
		Username: res.Name,
		Gender:   res.Gender,
		Mobile:   res.Mobile,
	}, nil
}
