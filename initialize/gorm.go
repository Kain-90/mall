package initialize

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"mall/global"
	"mall/model"
)

func SetupDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		global.GVA_CONFIG.DB.User,
		global.GVA_CONFIG.DB.Password,
		global.GVA_CONFIG.DB.Host,
		global.GVA_CONFIG.DB.Database)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		panic("failed to migrate database")
	}
	return db
}
