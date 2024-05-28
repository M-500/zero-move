package logic

import (
	"context"
	"github.com/pkg/errors"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

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

func (l *UpdateLogic) Update(in *order.UpdateRequest) (*order.UpdateResponse, error) {
	// TODO：并发不安全的修改
	res, err := l.svcCtx.OrderDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("订单ID不存在")
	}
	if in.Uid != 0 {
		res.Uid = in.Uid
	}
	if in.Pid != 0 {
		res.Pid = in.Pid
	}
	if in.Amount != 0 {
		res.Amount = float64(in.Amount)
	}
	if in.Status != 0 {
		res.Status = int(in.Status)
	}
	err = l.svcCtx.OrderDao.Update(l.ctx, res)
	if err != nil {
		return nil, err
	}
	return &order.UpdateResponse{}, nil
}
