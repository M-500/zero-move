package logic

import (
	"context"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveLogic {
	return &RemoveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *RemoveLogic) Remove(in *product.RemoveProdRequest) (*product.RemoveProdResponse, error) {
	err := l.svcCtx.ProductDao.DeleteById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &product.RemoveProdResponse{}, nil
}
