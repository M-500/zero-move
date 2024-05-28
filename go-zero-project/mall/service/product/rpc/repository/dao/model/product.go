package model

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木

type ProductModel struct {
	gorm.Model
	Name   string  `gorm:"type:varchar(128);comment:'商品名称'"`
	Desc   string  `gorm:"type:varchar(128);comment:'商品描述'"`
	Amount float64 ``
	Status int
}
