package dao

import (
	"context"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木

type ParserDao interface {
	Insert(ctx context.Context, data ParseJob) (int64, error)
	FindById(ctx context.Context, id int64) (ParseJob, error)
}

type parseDao struct {
	DB *gorm.DB
}

func NewParserDao(DB *gorm.DB) ParserDao {
	return &parseDao{
		DB: DB,
	}
}

func (p *parseDao) Insert(ctx context.Context, data ParseJob) (int64, error) {
	err := p.DB.WithContext(ctx).Model(&ParseJob{}).Create(&data).Error
	if err != nil {
		return 0, err
	}
	return int64(data.ID), err
}

func (p *parseDao) FindById(ctx context.Context, id int64) (ParseJob, error) {
	var job ParseJob
	err := p.DB.WithContext(ctx).Model(&ParseJob{}).Where("id = ?", id).Find(&job).Error
	if err != nil {
		return ParseJob{}, err
	}
	return job, nil
}
