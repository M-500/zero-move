package dao

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

// @Description
// @Author 代码小学生王木木

const CompanyIndexName = "company_info"
const CompanyFastIndexName = "company_fast_search"

type esEnterpriseDao struct {
	client *elastic.Client
}

func NewEsEnterpriseDao(client *elastic.Client) EnterpriseDAO {
	return &esEnterpriseDao{
		client: client,
	}
}

type CompanyFastSearch struct {
	Id          string `json:"id"`
	CompanyName string `json:"company_name"`
	LegalPerson string `json:"legal_person"`
	Logo        string `json:"logo"`
	PassName    string `json:"pass_name"`
}

func (e *esEnterpriseDao) InsertEnterpriseBaseInfo(ctx context.Context, data EnterpriseBasicModel) error {
	// 插入基本信息表
	do, err := e.client.Index().Index(CompanyIndexName).Id(data.SocialCode).BodyJson(data).Do(ctx)
	if err != nil {
		return err
	}
	fmt.Println("插入结果", do)
	response, err := e.client.Index().Index(CompanyFastIndexName).Id(data.SocialCode).BodyJson(e.toCompanyFastSearch(data)).Do(ctx)
	if err != nil {
		return err
	}
	fmt.Println("插入结果", response)
	return nil
}
func (e *esEnterpriseDao) toCompanyFastSearch(data EnterpriseBasicModel) CompanyFastSearch {
	return CompanyFastSearch{
		Id:          data.SocialCode,
		CompanyName: data.CompanyName,
		LegalPerson: data.LegalBoss,
		Logo:        "",
		PassName:    data.FormerName,
	}
}
func (e *esEnterpriseDao) BatchInsertEnterpriseBaseInfo(ctx context.Context, data []EnterpriseBasicModel) error {

	return nil
}
