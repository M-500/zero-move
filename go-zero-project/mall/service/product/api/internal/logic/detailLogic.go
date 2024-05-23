package logic

import (
	"context"

	"github.com/zeromicro/go-zero/core/logx"
	"mall/service/product/api/internal/svc"
)

type DetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DetailLogic) Detail() error {
	// todo: add your logic here and delete this line

	return nil
}
