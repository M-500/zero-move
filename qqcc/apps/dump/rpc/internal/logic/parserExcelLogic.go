package logic

import (
	"context"

	"qqcc/apps/dump/rpc/internal/svc"
	"qqcc/apps/dump/rpc/types/dump"

	"github.com/zeromicro/go-zero/core/logx"
)

type ParserExcelLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewParserExcelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ParserExcelLogic {
	return &ParserExcelLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ParserExcelLogic) ParserExcel(in *dump.ParserRequest) (*dump.ParserResponse, error) {
	// todo: add your logic here and delete this line

	// 1. 解析Excel

	// 2. 写入kafka消息队列

	return &dump.ParserResponse{}, nil
}
