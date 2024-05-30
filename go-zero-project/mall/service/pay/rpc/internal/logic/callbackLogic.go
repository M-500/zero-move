package logic

import (
	"context"
	"errors"
	"google.golang.org/grpc/status"
	"mall/service/order/rpc/types/order"
	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types/pay"
	"mall/service/user/rpc/types/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type CallbackLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCallbackLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CallbackLogic {
	return &CallbackLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CallbackLogic) Callback(in *pay.CallbackRequest) (*pay.CallbackResponse, error) {
	// 1. 查询用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.Uid})
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	// 2. 查询订单是否存在
	_, err = l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{Id: in.Oid})
	if err != nil {
		return nil, errors.New("订单不存在")
	}
	// 3. 查询支付是否存在
	res, err := l.svcCtx.PayDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("支付信息不存在")
	}
	// 4. 判断支付金额与订单金额是否相同
	if res.Amount != float64(in.Amount) {
		return nil, status.Error(100, "支付金额与订单金额不符")
	}
	res.Status = int(in.Status)
	res.Source = int(in.Source)

	// 更新支付状态
	err = l.svcCtx.PayDao.Update(l.ctx, res)
	if err != nil {
		// 支付更新失败
		return nil, status.Error(101, err.Error())
	}
	// 更新订单状态
	_, err = l.svcCtx.OrderRpc.Paid(l.ctx, &order.PaidRequest{Id: in.Oid})
	if err != nil {
		// 更新订单状态失败
		return nil, status.Error(102, err.Error())
	}
	return &pay.CallbackResponse{}, nil
}
