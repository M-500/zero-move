package sharding_stu

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/sharding"
	"testing"
)

// @Description
// @Author 代码小学生王木木

type UserModel struct {
	Id       int64  `gorm:"primarykey"`
	Username string `gorm:"column:username"`
	Password string `gorm:"column:password"`
	Age      int    `gorm:"column:age"`
}

func (UserModel) TableName() string {
	return "user"
}

func TestShadingMigrate(t *testing.T) {
	// 1. 连接MYSQL
	dsn := "root:root@tcp(127.0.0.1:13316)/sharding?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	for i := 0; i < 128; i++ {
		tablename := fmt.Sprintf("user_%0*d", 3, i)
		err := db.Table(tablename).AutoMigrate(&UserModel{})
		if err != nil {

		}
	}
}

func TestShardingInsert(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:13316)/sharding?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	db.Use(sharding.Register(sharding.Config{
		ShardingKey:         "username",
		NumberOfShards:      128,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "user"))
	for i := 0; i < 1000; i++ {
		user := UserModel{
			Id:       int64(i),
			Username: fmt.Sprintf("琳琳_%d", i),
			Password: fmt.Sprintf("1234_%d", i),
			Age:      0,
		}
		db.Create(&user)
	}
}

func TestShardingQuery(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:13316)/sharding?parseTime=true&loc=Local"
	config := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}
	db, err := gorm.Open(mysql.Open(dsn), config)
	if err != nil {
		panic(err)
	}
	db.Use(sharding.Register(sharding.Config{
		ShardingKey:         "username",
		NumberOfShards:      128,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "user"))
	var user = UserModel{}
	db.Where("username = ?", "琳琳_531").Find(&user)
	t.Log(user)
}

func TestShardingQueryMore(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:13316)/sharding?parseTime=true&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		panic(err)
	}
	db.Use(sharding.Register(sharding.Config{
		ShardingKey:         "username",
		NumberOfShards:      128,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}, "user"))
	var user = []UserModel{}
	db.Where("username IN ?", []string{"琳琳_531", "琳琳_517", "琳琳_817"}).Find(&user)
	t.Log(user)
}
