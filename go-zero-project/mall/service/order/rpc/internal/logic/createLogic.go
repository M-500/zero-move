package logic

import (
	"context"
	"mall/service/order/rpc/repository/dao/model"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

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

func (l *CreateLogic) Create(in *order.CreateOrderRequest) (*order.CreateResponse, error) {
	id, err := l.svcCtx.OrderDao.Insert(l.ctx, model.OrderModel{
		Uid:    in.Uid,
		Pid:    in.Pid,
		Amount: float64(in.Amount),
		Status: int(in.Status),
	})
	if err != nil {
		return nil, err
	}
	return &order.CreateResponse{
		Id: id,
	}, nil
}
