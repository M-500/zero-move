package dao

import "gorm.io/gorm"

// @Description
// @Author 代码小学生王木木

func InitTable(db *gorm.DB) error {
	return db.AutoMigrate(&UserModel{})
}
