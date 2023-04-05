package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"sync"
)

var db *gorm.DB

func NewDB() *gorm.DB {
	var once sync.Once
	once.Do(func() {
		// 配置 MySQL 数据库连接信息
		dsn := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
		// 连接 MySQL 数据库
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic(err)
		}
		db = db.Debug()
	})

	return db
}
