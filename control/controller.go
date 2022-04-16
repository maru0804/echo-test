package control

import (
	"log"

	"github.com/maru0804/echo-test.git/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// connect DB function
func gormConnect() *gorm.DB {
	CONNECT := "user:password@tcp(echo-db:3306)/echo_db_dev?parseTime=true"
	log.Println("aaaaaaa")
	db, err := gorm.Open(mysql.Open(CONNECT), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	return db
}

// db init
func DbInit() {
	db := gormConnect()
	// コネクション解放
	// defer db.Close()
	defer func() {
		if sqlDB, err := db.DB(); err != nil {
			sqlDB.Close()
		}
	}()

	//構造体に基づいてテーブルを作成
	db.AutoMigrate(&models.User{})
	log.Println("create user table")
}
