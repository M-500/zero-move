package dao

import (
	"gorm.io/gorm"
	"mall/service/pay/rpc/rpository/dao/model"
)

// @Description
// @Author 代码小学生王木木

func InitTables(db *gorm.DB) error {
	return db.AutoMigrate(&model.PayModel{})
}
