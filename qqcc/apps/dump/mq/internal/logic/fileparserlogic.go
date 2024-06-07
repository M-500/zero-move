package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"qqcc/apps/dump/mq/internal/svc"
	"qqcc/apps/dump/mq/internal/types"
)

type FileParserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFileParserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FileParserLogic {

	return &FileParserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FileParserLogic) Consume(_, val string) error {
	var msg *types.FileParserMsg
	err := json.Unmarshal([]byte(val), &msg)
	if err != nil {
		logx.Errorf("Consume val: %s error: %v", val, err)
		return err
	}
	return l.ParserFile(l.ctx, msg)
}

func (l *FileParserLogic) ParserFile(ctx context.Context, msg *types.FileParserMsg) error {
	// 1.去数据库查询是否 重复消费了

	// 2. 开始解析，并插入数据库中

	return nil
}

func Consumers(ctx context.Context, svcCtx *svc.ServiceContext) []service.Service {
	return []service.Service{
		kq.MustNewQueue(svcCtx.Config.KqConsumerConf, NewFileParserLogic(ctx, svcCtx)),
	}
}
