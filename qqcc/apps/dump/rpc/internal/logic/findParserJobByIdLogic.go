package logic

import (
	"context"

	"qqcc/apps/dump/rpc/internal/svc"
	"qqcc/apps/dump/rpc/types/dump"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindParserJobByIdLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindParserJobByIdLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindParserJobByIdLogic {
	return &FindParserJobByIdLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FindParserJobByIdLogic) FindParserJobById(in *dump.FindParserJonRequest) (*dump.FindParserJonResponse, error) {
	res, err := l.svcCtx.ParserDAO.FindById(l.ctx, in.Id)
	if err != nil {
		return nil, err
	}
	return &dump.FindParserJonResponse{
		Id:       int64(res.ID),
		Filepath: res.FilePath,
		Filesize: res.FileSizeMB,
		Mark:     res.Mark,
		Resolved: res.Resolved,
		Rows:     int64(res.Rows),
	}, nil
}
