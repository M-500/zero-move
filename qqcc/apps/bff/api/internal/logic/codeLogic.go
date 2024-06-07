package logic

import (
	"context"
	"errors"

	"qqcc/apps/bff/api/internal/svc"
	"qqcc/apps/bff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CodeLogic {
	return &CodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CodeLogic) Code() (resp *types.CodeRespose, err error) {
	// todo: add your logic here and delete this line
	id, path, err := l.svcCtx.CapSvc.MakeCaptcha()
	if err != nil {
		return nil, errors.New("生成验证码失败")
	}
	res := types.CodeRespose{
		CaptchaId: id,
		PicPath:   path,
	}
	return &res, nil
}
