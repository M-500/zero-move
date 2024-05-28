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
	Delete(ctx context.Context, id int64) (model.OrderModel, error)
	FindById(ctx context.Context, id int64) (model.OrderModel, error)
	FindListByUid(ctx context.Context, id int64) ([]model.OrderModel, error)
}

type orderDAO struct {
	DB gorm.DB
}

func (o *orderDAO) Insert(ctx context.Context, data model.OrderModel) (int64, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderDAO) Update(ctx context.Context, data model.OrderModel) error {
	//TODO implement me
	panic("implement me")
}

func (o *orderDAO) Delete(ctx context.Context, id int64) (model.OrderModel, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderDAO) FindById(ctx context.Context, id int64) (model.OrderModel, error) {
	//TODO implement me
	panic("implement me")
}

func (o *orderDAO) FindListByUid(ctx context.Context, id int64) ([]model.OrderModel, error) {
	//TODO implement me
	panic("implement me")
}
