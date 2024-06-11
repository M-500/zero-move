package logic

import (
	"context"

	"qqcc/apps/dump/rpc/internal/svc"
	"qqcc/apps/dump/rpc/types/dump"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateParserJobLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateParserJobLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateParserJobLogic {
	return &UpdateParserJobLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *UpdateParserJobLogic) UpdateParserJob(in *dump.UpdateParserJonRequest) (*dump.UpdateParserJonResponse, error) {
	// todo: add your logic here and delete this line

	return &dump.UpdateParserJonResponse{}, nil
}
