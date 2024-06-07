package dao

import "context"

// @Description
// @Author 代码小学生王木木

type EnterpriseDAO interface {
	InsertEnterpriseBaseInfo(ctx context.Context, data EnterpriseBasicModel) error
	BatchInsertEnterpriseBaseInfo(ctx context.Context, data []EnterpriseBasicModel) error
}
