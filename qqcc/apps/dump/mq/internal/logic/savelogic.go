package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/logx"
	"qqcc/apps/dump/mq/internal/svc"
)

// @Description 从kafka中读取工商数据，然后持久化到MySQL中 是否要双写ES?
// @Author 代码小学生王木木

type CompanySaveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCompanySaveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CompanySaveLogic {
	return &CompanySaveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// Consume
//
//	@Description: 消费者程序
//	@receiver c
//	@param _
//	@param val
//	@return error
func (c *CompanySaveLogic) Consume(_, val string) error {
	return nil
}
