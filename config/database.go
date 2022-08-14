package config

import (
	"github.com/yumekiti/echo-demo/domain/entity"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDB DBと接続する
func NewDB() *gorm.DB {
	dsn := "user:password@tcp(db:3306)/database?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(entity.Task{})

	return db
}
