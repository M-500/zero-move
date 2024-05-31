package dao

import (
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木

type UserModel struct {
	gorm.Model
	Name     string `gorm:"uniqueIndex;not null;type:varchar(32);column:username;comment:用户名，唯一索引，不能为空"`
	Mobile   string `gorm:"index;type:varchar(32);comment:手机号;column:mobile"`
	Gender   int64  `gorm:"type:tinyint;comment:性别;column:gender"`
	Password string `gorm:"type:varchar(128);comment:密码;column:password"`
	IsAdmin  int64  `gorm:"int;default:0;comment:是否是admin，默认0,非管理员;column:gender"`
	Avatar   string `gorm:"type:varchar(200);comment:用户头像"`
	CreateId int64  `gorm:"comment:创建人ID;not null;default:1"`
}

func (UserModel) TableName() string {
	return "mall_user"
}
