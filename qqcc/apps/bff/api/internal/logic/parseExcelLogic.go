package logic

import (
	"context"
	"qqcc/apps/dump/rpc/types/dump"

	"qqcc/apps/bff/api/internal/svc"
	"qqcc/apps/bff/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ParseExcelLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewParseExcelLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ParseExcelLogic {
	return &ParseExcelLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ParseExcelLogic) ParseExcel(req *types.ParseExcelRequest) (resp *types.ParseExcelResponse, err error) {
	// todo: add your logic here and delete this line
	_, err = l.svcCtx.DumpRpc.ParserExcel(l.ctx, &dump.ParserRequest{
		Filename: req.FileName,
		Filepath: req.FilePath,
	})
	if err != nil {
		return nil, err
	}
	return &types.ParseExcelResponse{}, nil
}
