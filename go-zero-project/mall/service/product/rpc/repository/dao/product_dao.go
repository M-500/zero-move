package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/service/product/rpc/repository/dao/model"
)

// @Description
// @Author 代码小学生王木木

type ProductDao interface {
	Insert(ctx context.Context, data model.ProductModel) (int64, error)
}

type productDao struct {
	DB *gorm.DB
}

func NewProductDao(DB *gorm.DB) ProductDao {
	return &productDao{DB: DB}
}

func (p *productDao) Insert(ctx context.Context, data model.ProductModel) (int64, error) {
	err := p.DB.WithContext(ctx).Model(&model.ProductModel{}).Create(&data).Error
	return int64(data.ID), err
}
