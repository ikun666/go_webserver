package dao

import (
	"time"

	"github.com/ikun666/go_webserver/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type BaseDao struct {
	DB *gorm.DB
}

func InitDB() (*gorm.DB, error) {
	//连接数据库
	db, err := gorm.Open(mysql.Open(viper.GetString("db.dns")), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "sys_",
			SingularTable: true,
		},
	})
	if err != nil {
		return db, err
	}
	//设置连接参数
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(viper.GetInt("db.SetMaxIdleConns"))
	sqlDB.SetMaxOpenConns(viper.GetInt("db.SetMaxOpenConns"))
	sqlDB.SetConnMaxLifetime(time.Hour)
	//修改model.User{}内容会自动修改数据库表
	db.AutoMigrate(&model.User{})
	return db, err
}
