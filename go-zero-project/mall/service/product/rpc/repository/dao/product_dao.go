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
	FindById(ctx context.Context, id int64) (model.ProductModel, error)
	Update(ctx context.Context, data model.ProductModel) error
	DeleteById(ctx context.Context, id int64) error
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
func (p *productDao) FindById(ctx context.Context, id int64) (model.ProductModel, error) {
	var product model.ProductModel
	err := p.DB.WithContext(ctx).Model(&model.ProductModel{}).Where("id=?", id).Find(&product).Error
	if err != nil {
		return model.ProductModel{}, err
	}
	return product, nil
}

func (p *productDao) DeleteById(ctx context.Context, id int64) error {
	return p.DB.WithContext(ctx).Delete(&model.ProductModel{}, id).Error
}

func (p *productDao) Update(ctx context.Context, data model.ProductModel) error {
	return nil
}
