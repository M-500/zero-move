package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
	"qqcc/apps/dump/domain"
	"qqcc/apps/dump/mq/internal/repository/dao"
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
//	@Description: 从kafka中消费数据，然后 持久化到DB中  1. 双写到ES中  2.只写MySQL 通过Canal来监听 然后写入ES中
func (c *CompanySaveLogic) LoadMsg(ctx context.Context, msg *domain.EnterpriseBasicDM) error {
	fmt.Println("哈哈哈", msg)
	// 去重器   全局的布隆过滤器去重  -> 本地布隆过滤器去重 -> 数据库唯一索引兜底去重
	err := c.svcCtx.EntPriRepo.InsertDBEntBaseInfo(ctx, c.toEntity(msg))
	return err
}

func (c *CompanySaveLogic) toEntity(item *domain.EnterpriseBasicDM) dao.EnterpriseBasicModel {
	return dao.EnterpriseBasicModel{
		Model: gorm.Model{
			ID:        uint(item.Id),
			CreatedAt: item.CreatedAt,
			UpdatedAt: item.UpdatedAt,
		},
		CompanyName:                item.CompanyName,
		RegStatus:                  item.RegStatus,
		LegalBoss:                  item.LegalBoss,
		RegCapital:                 item.RegCapital,
		RegDate:                    item.RegDate,
		SocialCode:                 item.SocialCode,
		CompanyRegAddr:             item.CompanyRegAddr,
		Phone:                      item.Phone,
		Email:                      item.Email,
		Province:                   item.Province,
		City:                       item.City,
		District:                   item.District,
		TaxNumber:                  item.TaxNumber,
		RegNumber:                  item.RegNumber,
		OrgCode:                    item.OrgCode,
		InsuredNumber:              item.InsuredNumber,
		InsuredNumberReport:        item.InsuredNumberReport,
		CompanyType:                item.CompanyType,
		CompanySize:                item.CompanySize,
		ApprovalDate:               item.ApprovalDate,
		BusinessTerm:               item.BusinessTerm,
		GBIndustryCategory:         item.GBIndustryCategory,
		GBIndustryMajorCategory:    item.GBIndustryMajorCategory,
		GBIndustryMediumCategory:   item.GBIndustryMediumCategory,
		GBIndustrySmallCategory:    item.GBIndustrySmallCategory,
		FormerName:                 item.FormerName,
		EnglishName:                item.EnglishName,
		OfficialWebsite:            item.OfficialWebsite,
		CommunicationAdd:           item.CommunicationAdd,
		CompanyProfile:             item.CompanyProfile,
		BusinessScope:              item.BusinessScope,
		EBIPIndustryCategory:       item.EBIPIndustryCategory,
		EBIPIndustryMajorCategory:  item.EBIPIndustryMajorCategory,
		EBIPIndustryMediumCategory: item.EBIPIndustryMediumCategory,
		EBIPIndustrySmallCategory:  item.EBIPIndustrySmallCategory,
	}
}
