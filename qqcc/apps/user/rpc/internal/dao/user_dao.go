package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木
var ErrUserDuplicate = errors.New("用户名冲突")

type UserDao interface {
	FindByUserName(ctx context.Context, username string) (UserModel, error)
	FindById(ctx context.Context, id int64) (UserModel, error)
	Insert(ctx context.Context, data UserModel) (int64, error)
}

type userDao struct {
	DB *gorm.DB
}

func NewUserDao(DB *gorm.DB) UserDao {
	return &userDao{DB: DB}
}
func (u *userDao) FindById(ctx context.Context, id int64) (UserModel, error) {
	var user UserModel
	//err := u.DB.WithContext(ctx).Model(&UserModel{}).Select("").Where("id = ?", id).Find(&user).Error // 如果
	err := u.DB.WithContext(ctx).Model(&UserModel{}).Where("id = ?", id).Find(&user).Error
	if err != nil {
		return UserModel{}, err
	}
	return user, nil
}
func (u *userDao) FindByUserName(ctx context.Context, username string) (UserModel, error) {
	var user = UserModel{}
	err := u.DB.WithContext(ctx).Model(&UserModel{}).Where("username = ?", username).Find(&user).Error
	if err != nil {
		return UserModel{}, err
	}
	return user, nil
}

func (u *userDao) Insert(ctx context.Context, data UserModel) (int64, error) {
	err := u.DB.WithContext(ctx).Model(&UserModel{}).Create(&data).Error
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			const uniqueConflictsErrNo uint16 = 1062
			if mysqlErr.Number == uniqueConflictsErrNo {
				// 邮箱冲突 or 手机号码冲突
				return 0, ErrUserDuplicate
			}
		}
		return 0, err
	}
	return int64(data.ID), nil
}
