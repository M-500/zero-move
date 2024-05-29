package logic

import (
	"context"
	"errors"
	"mall/service/order/rpc/types/order"
	"mall/service/pay/rpc/rpository/dao/model"
	"mall/service/user/rpc/types/user"

	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types/pay"

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

func (l *CreateLogic) Create(in *pay.CreateRequest) (*pay.CreateResponse, error) {
	// 1. 调用用户服务，查看用户是否存在
	_, err := l.svcCtx.UserRpc.UserInfo(l.ctx, &user.UserInfoRequest{Id: in.Uid})
	if err != nil {
		return nil, errors.New("用户不存在")
	}
	// 2. 调用订单服务，查看订单是否存在
	_, err2 := l.svcCtx.OrderRpc.Detail(l.ctx, &order.DetailRequest{Id: in.Oid})
	if err2 != nil {
		return nil, errors.New("订单不存在")
	}
	// 3. 查询订单是否已经支付
	res, err := l.svcCtx.PayDao.FindById(l.ctx, in.Oid)
	if err == nil {
		if res.Status == 1 {
			return nil, errors.New("订单已经支付")
		}
		return nil, errors.New("订单已存在，无法重复创建")
	}
	id, err := l.svcCtx.PayDao.Insert(l.ctx, model.PayModel{
		Uid:    in.Uid,
		Oid:    in.Oid,
		Amount: float64(in.Amount),
		Source: 1,
		Status: 0,
	})
	if err != nil {
		return nil, err
	}
	return &pay.CreateResponse{Id: id}, nil
}
