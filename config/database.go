package config

import (
	"github.com/yumekiti/echo-demo/domain/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewDB DBと接続する
func NewDB() *gorm.DB {
	dsn := "user:password@tcp(sample_db)/sample?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(model.Task{})

	return db
}
