package logic

import (
	"context"
	"mall/service/product/rpc/repository/dao/model"

	"mall/service/product/rpc/internal/svc"
	"mall/service/product/rpc/types/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateLogic) Create(in *product.ProdRequest) (*product.ProdResponse, error) {
	id, err := l.svcCtx.ProductDao.Insert(l.ctx, model.ProductModel{
		Name:   in.Name,
		Desc:   in.Desc,
		Amount: float64(in.Amount),
		Status: int(in.Status),
	})
	if err != nil {
		return nil, err
	}
	return &product.ProdResponse{
		Id: id,
	}, nil
}
