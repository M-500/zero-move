package gormx

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木

type CustomerPlugin struct {
}

func NewCustomerPlugin() *CustomerPlugin {
	return &CustomerPlugin{}
}

func (p *CustomerPlugin) Name() string {
	return "CustomerPlugin"
}

func (p *CustomerPlugin) Initialize(db *gorm.DB) {

}
