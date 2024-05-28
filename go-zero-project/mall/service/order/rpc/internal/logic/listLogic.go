package logic

import (
	"context"

	"mall/service/order/rpc/internal/svc"
	"mall/service/order/rpc/types/order"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListLogic {
	return &ListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListLogic) List(in *order.ListRequest) (*order.ListResponse, error) {
	res, err := l.svcCtx.OrderDao.FindListByUid(l.ctx, in.Uid)
	if err != nil {
		return nil, err
	}
	var data []*order.DetailResponse
	for _, item := range res {
		temp := &order.DetailResponse{
			Id:     int64(item.ID),
			Uid:    item.Uid,
			Pid:    item.Pid,
			Amount: int64(item.Amount),
			Status: int64(item.Status),
		}
		data = append(data, temp)
	}
	return &order.ListResponse{
		Data: data,
	}, nil
}
