package dao

import (
	"context"
	"gorm.io/gorm"
	"mall/service/order/rpc/repository/dao/model"
)

// @Description
// @Author 代码小学生王木木

type OrderDAO interface {
	Insert(ctx context.Context, data model.OrderModel) (int64, error) // 创建订单
	Update(ctx context.Context, data model.OrderModel) error
	Delete(ctx context.Context, id int64) error
	FindById(ctx context.Context, id int64) (model.OrderModel, error)
	FindListByUid(ctx context.Context, id int64) ([]model.OrderModel, error)
}

type orderDAO struct {
	DB *gorm.DB
}

func NewOrderDAO(DB *gorm.DB) OrderDAO {
	return &orderDAO{DB: DB}
}

func (o *orderDAO) Insert(ctx context.Context, data model.OrderModel) (int64, error) {
	err := o.DB.WithContext(ctx).Model(&model.OrderModel{}).Create(&data).Error
	return 0, err
}

func (o *orderDAO) Update(ctx context.Context, data model.OrderModel) error {
	return o.DB.WithContext(ctx).Model(&model.OrderModel{}).Where("id = ?", data.ID).Updates(&data).Error
}

func (o *orderDAO) Delete(ctx context.Context, id int64) error {
	return o.DB.WithContext(ctx).Delete(&model.OrderModel{}, id).Error
}

func (o *orderDAO) FindById(ctx context.Context, id int64) (model.OrderModel, error) {
	var order model.OrderModel
	err := o.DB.WithContext(ctx).Model(&model.OrderModel{}).Where("id = ?", id).Find(&order).Error
	if err != nil {
		return model.OrderModel{}, err
	}
	return order, nil
}

func (o *orderDAO) FindListByUid(ctx context.Context, id int64) ([]model.OrderModel, error) {
	var orders []model.OrderModel
	err := o.DB.WithContext(ctx).Model(&model.OrderModel{}).Where("uid = ?", id).Find(&orders).Error
	if err != nil {
		return nil, err
	}
	return orders, nil
}
