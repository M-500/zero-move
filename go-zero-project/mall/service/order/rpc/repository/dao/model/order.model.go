package model

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木

type OrderModel struct {
	gorm.Model
	Uid    int64   `gorm:"comment:用户ID"`
	Pid    int64   `gorm:"comment:产品ID"`
	Amount float64 `gorm:"comment:订单金额"`
	Status int     `gorm:"comment:订单状态"`
}

func (OrderModel) TableName() string {
	return "order"
}
