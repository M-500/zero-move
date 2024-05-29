package model

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木

type PayModel struct {
	gorm.Model
	Uid    int64   `gorm:"comment:用户ID"`
	Oid    int64   `gorm:"comment:订单ID"`
	Amount float64 `gorm:"comment:产品金额"`
	Source int     `gorm:"comment:支付方式，支付宝,微信"`
	Status int     `gorm:"comment:支付状态,0 预支付，1 支付成功，2 支付失败，3 取消支付"`
}

func (PayModel) TableName() string {
	return "pay"
}
