package logic

import (
	"context"
	"encoding/json"
	"github.com/zeromicro/go-zero/core/logx"
	"qqcc/apps/dump/domain"
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
func (c *CompanySaveLogic) Consume(_, val string) error {
	var msg *domain.EnterpriseBasicDM
	err := json.Unmarshal([]byte(val), &msg)
	if err != nil {
		logx.Errorf("Consume val: %s error: %v", val, err)
		return err
	}
	return c.LoadMsg(c.ctx, msg)
}

// LoadMsg
//
//	@Description: 持久化到DB中
func (c *CompanySaveLogic) LoadMsg(ctx context.Context, msg *domain.EnterpriseBasicDM) error {

	return nil
}
