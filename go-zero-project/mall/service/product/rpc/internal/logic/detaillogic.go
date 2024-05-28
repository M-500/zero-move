package logic

import (
	"context"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type DetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DetailLogic {
	return &DetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *DetailLogic) Detail(in *product.DetailProdRequest) (*product.DetailProdResponse, error) {
	res, err := l.svcCtx.ProductDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &product.DetailProdResponse{
		Id:     int64(res.ID),
		Name:   res.Name,
		Desc:   res.Desc,
		Stock:  res.Stock,
		Amount: int64(res.Amount),
		Status: int64(res.Status),
	}, nil
}
