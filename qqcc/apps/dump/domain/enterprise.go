package domain

import "time"

// @Description
// @Author 代码小学生王木木

type EnterpriseBasicDM struct {
	Id                         int64     `json:"id"`
	CreatedAt                  time.Time `json:"createdAt"`
	UpdatedAt                  time.Time `json:"updatedAt"`
	CompanyName                string    `json:"companyName"`
	RegStatus                  string    `json:"regStatus"`
	LegalBoss                  string    `json:"legalBoss"`
	RegCapital                 string    `json:"regCapital"`
	RegDate                    time.Time `json:"regDate"`
	SocialCode                 string    `json:"socialCode"`
	CompanyRegAddr             string    `json:"companyRegAddr"`
	Phone                      string    `json:"phone"`
	Email                      string    `json:"email"`
	Province                   string    `json:"province"`
	City                       string    `json:"city"`
	District                   string    `json:"district"`
	TaxNumber                  string    `json:"taxNumber"`
	RegNumber                  string    `json:"regNumber"`
	OrgCode                    string    `json:"orgCode"`
	InsuredNumber              int       `json:"insuredNumber"`
	InsuredNumberReport        string    `json:"insuredNumberReport"`
	CompanyType                string    `json:"companyType"`
	CompanySize                string    `json:"companySize"`
	ApprovalDate               time.Time `json:"approvalDate"`
	BusinessTerm               string    `json:"businessTerm"`
	GBIndustryCategory         string    `json:"GBIndustryCategory"`
	GBIndustryMajorCategory    string    `json:"GBIndustryMajorCategory"`
	GBIndustryMediumCategory   string    `json:"GBIndustryMediumCategory"`
	GBIndustrySmallCategory    string    `json:"GBIndustrySmallCategory"`
	FormerName                 string    `json:"formerName"`
	EnglishName                string    `json:"englishName"`
	OfficialWebsite            string    `json:"officialWebsite"`
	CommunicationAdd           string    `json:"communicationAdd"`
	CompanyProfile             string    `json:"companyProfile"`
	BusinessScope              string    `json:"businessScope"`
	EBIPIndustryCategory       string    `json:"EBIPIndustryCategory"`
	EBIPIndustryMajorCategory  string    `json:"EBIPIndustryMajorCategory"`
	EBIPIndustryMediumCategory string    `json:"EBIPIndustryMediumCategory"`
	EBIPIndustrySmallCategory  string    `json:"EBIPIndustrySmallCategory"`
}
