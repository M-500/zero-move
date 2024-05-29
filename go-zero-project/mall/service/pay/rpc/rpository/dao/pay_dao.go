package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/service/pay/rpc/rpository/dao/model"
)

// @Description
// @Author 代码小学生王木木

type PayDao interface {
	FindByOrderId(ctx context.Context, oid int64) (model.PayModel, error)
	FindById(ctx context.Context, id int64) (model.PayModel, error)
	Insert(ctx context.Context, data model.PayModel) (int64, error)
	Update(ctx context.Context, data model.PayModel) error
}

type payDao struct {
	DB *gorm.DB
}

func NewPayDao(db *gorm.DB) PayDao {
	return &payDao{
		DB: db,
	}
}

func (p *payDao) FindByOrderId(ctx context.Context, oid int64) (model.PayModel, error) {
	var pay model.PayModel
	err := p.DB.WithContext(ctx).Model(&model.PayModel{}).Where("oid = ?", oid).Find(&pay).Error
	if err != nil {
		return model.PayModel{}, err
	}
	return pay, nil
}
func (p *payDao) FindById(ctx context.Context, id int64) (model.PayModel, error) {
	var pay model.PayModel
	err := p.DB.WithContext(ctx).Model(&model.PayModel{}).Where("id = ?", id).Find(&pay).Error
	if err != nil {
		return model.PayModel{}, err
	}
	return pay, nil
}

func (p *payDao) Insert(ctx context.Context, data model.PayModel) (int64, error) {
	err := p.DB.WithContext(ctx).Model(&model.PayModel{}).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return int64(data.ID), nil
}
func (p *payDao) Update(ctx context.Context, data model.PayModel) error {
	return p.DB.WithContext(ctx).Model(&model.PayModel{}).Updates(&data).Error
}
