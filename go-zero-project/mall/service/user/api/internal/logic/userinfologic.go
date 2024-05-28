package logic

import (
	"context"
	"mall/service/user/rpc/types/user"

	"mall/service/user/api/internal/svc"
	"mall/service/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoResponse, err error) {
	info, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: resp.Id})
	if err != nil {
		return nil, err
	}
	resp.Id = info.Id
	resp.Name = info.Name
	resp.Gender = info.Gender
	resp.Mobile = info.Mobile
	return
}
