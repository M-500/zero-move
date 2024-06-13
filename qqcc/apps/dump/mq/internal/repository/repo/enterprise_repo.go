package repo

import (
	"context"
	"qqcc/apps/dump/mq/internal/repository/dao"
)

// @Description
// @Author 代码小学生王木木

type EnterpriseRepo interface {
	InsertDBEntBaseInfo(ctx context.Context, data dao.EnterpriseBasicModel) error
}

type enterpriseRepo struct {
	dao dao.EnterpriseDAO
}

func NewEnterpriseRepo(d dao.EnterpriseDAO) EnterpriseRepo {
	return &enterpriseRepo{
		dao: d,
	}
}

func (r *enterpriseRepo) InsertDBEntBaseInfo(ctx context.Context, data dao.EnterpriseBasicModel) error {
	err := r.dao.InsertEnterpriseBaseInfo(ctx, data)

	return err
}
