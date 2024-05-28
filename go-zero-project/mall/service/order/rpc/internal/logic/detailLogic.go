package logic

import (
	"context"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

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

func (l *DetailLogic) Detail(in *order.DetailRequest) (*order.DetailResponse, error) {
	// todo: add your logic here and delete this line https://juejin.cn/post/7036011643737735198#heading-0
	res, err := l.svcCtx.OrderDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &order.DetailResponse{
		Id:     int64(res.ID),
		Uid:    res.Uid,
		Pid:    res.Pid,
		Amount: int64(res.Amount),
		Status: int64(res.Status),
	}, nil
}
