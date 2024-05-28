package logic

import (
	"context"
	"errors"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type PaidLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPaidLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PaidLogic {
	return &PaidLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *PaidLogic) Paid(in *order.PaidRequest) (*order.PaidResponse, error) {
	// TODO：并发不安全的修改
	res, err := l.svcCtx.OrderDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("订单ID不存在")
	}
	res.Status = 1 // 更新支付状态
	err = l.svcCtx.OrderDao.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}
	return &order.PaidResponse{}, nil
}
