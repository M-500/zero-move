package dao

import (
	"gorm.io/gorm"
	"mall/service/order/rpc/repository/dao/model"
)

// @Description
// @Author 代码小学生王木木
func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&model.OrderModel{})
}
