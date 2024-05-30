package logic

import (
	"context"
	"errors"

	"mall/service/pay/rpc/internal/svc"
	"mall/service/pay/rpc/types/pay"

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

func (l *DetailLogic) Detail(in *pay.DetailRequest) (*pay.DetailResponse, error) {
	res, err := l.svcCtx.PayDao.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, errors.New("支付不存在")
	}
	return &pay.DetailResponse{
		Id:     int64(res.ID),
		Uid:    res.Uid,
		Oid:    res.Oid,
		Amount: int64(res.Amount),
		Source: int64(res.Source),
		Status: int64(res.Status),
	}, nil
}
