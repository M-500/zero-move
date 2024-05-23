package gormx

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

// @Description
// @Author 代码小学生王木木

func SetUpMySQL(conf *Config) (*DBX, error) {
	if conf.MaxIdleConns == 0 {
		conf.MaxIdleConns = 10
	}
	if conf.MaxOpenConns == 0 {
		conf.MaxOpenConns = 100
	}
	if conf.MaxLifetime == 0 {
		conf.MaxLifetime = 3600
	}
	db, err := gorm.Open(mysql.Open(conf.DSN), &gorm.Config{
		Logger: &ormLog{},
	})
	if err != nil {
		return nil, err
	}
	sdb, err := db.DB()
	if err != nil {
		return nil, err
	}
	sdb.SetMaxIdleConns(conf.MaxIdleConns)
	sdb.SetMaxOpenConns(conf.MaxOpenConns)
	sdb.SetConnMaxLifetime(time.Second * time.Duration(conf.MaxLifetime))

	return &DBX{DB: db}, nil
}

func MustSetupMySQl(conf *Config) *DBX {
	sql, err := SetUpMySQL(conf)
	if err != nil {
		panic(err)
	}
	return sql
}
