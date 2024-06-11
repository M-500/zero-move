package dao

import (
	"context"
	"errors"
	"github.com/go-sql-driver/mysql"
	"go.opentelemetry.io/otel"
	"gorm.io/gorm"
)

// @Description
// @Author 代码小学生王木木

var ErrEnterpriseDuplicate = errors.New("用户名冲突")

const uniqueConflictsErrNo uint16 = 1062

func NewEnterpriseDAO(db *gorm.DB) EnterpriseDAO {
	return &enterpriseDAO{
		db: db,
	}
}

type enterpriseDAO struct {
	db *gorm.DB
}

func (e *enterpriseDAO) InsertEnterpriseBaseInfo(ctx context.Context, data EnterpriseBasicModel) error {
	_, span := otel.Tracer("AddEnterpriseBaseInfo").Start(ctx, "InsertEnterpriseBaseInfo")
	defer span.End()
	err := e.db.WithContext(ctx).Model(&EnterpriseBasicModel{}).Create(&data).Error
	if mysqlErr, ok := err.(*mysql.MySQLError); ok {
		if mysqlErr.Number == uniqueConflictsErrNo {
			// 企业信息冲突
			return ErrEnterpriseDuplicate
		}
	}
	return err
}

// BatchInsertEnterpriseBaseInfo
//
//	@Description: GORM批量插入 一定要执行原生SQL吗？？？？ 不理解
//	@receiver e
//	@param ctx
//	@param bizId
//	@param biz
//	@param data
//	@return error
func (e *enterpriseDAO) BatchInsertEnterpriseBaseInfo(ctx context.Context, data []EnterpriseBasicModel) error {
	err := e.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		err := tx.Model(&EnterpriseBasicModel{}).CreateInBatches(data, len(data)).Error
		return err
	})
	if err != nil {
		// 如果插入错误，就只能一条一条的插入了，
		for _, datum := range data {
			err2 := e.InsertEnterpriseBaseInfo(ctx, datum)
			if err2 != nil {
				// 记录日志咯，还能怎么办？
				continue
			}
		}
	}
	return nil
}
