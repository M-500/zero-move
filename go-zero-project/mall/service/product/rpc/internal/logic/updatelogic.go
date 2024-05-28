package logic

import (
	"context"
	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateLogic {
	return &UpdateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateLogic) Update(in *product.UpdateProdRequest) (*product.UpdateProdResponse, error) {
	// todo: 会有并发问题的写法
	res, err := l.svcCtx.ProductDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	if in.Name != "" {
		res.Name = in.Name
	}
	if in.Desc != "" {
		res.Desc = in.Desc
	}
	if in.Stock != 0 {
		res.Stock = in.Stock
	}
	if in.Amount != 0 {
		res.Amount = float64(in.Amount)
	}
	if in.Status != 0 {
		res.Status = int(in.Status)
	}
	err = l.svcCtx.ProductDao.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}
	return &product.UpdateProdResponse{}, nil
}
