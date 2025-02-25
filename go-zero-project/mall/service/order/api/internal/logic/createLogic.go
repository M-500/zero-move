package logic

import (
	"context"
	"mall/service/order/rpc/types/order"

	"mall/service/order/api/internal/svc"
	"mall/service/order/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateLogic {
	return &CreateLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateLogic) Create(req *types.CreateRequest) (resp *types.CreateResponse, err error) {
	// todo: add your logic here and delete this line
	create, err := l.svcCtx.OrderRPC.Create(l.ctx, &order.CreateOrderRequest{
		Uid:    req.Uid,
		Pid:    req.Pid,
		Amount: uint64(req.Amount),
		Status: req.Status,
	})
	if err != nil {
		return nil, err
	}
	return &types.CreateResponse{Id: create.Id}, nil
}
