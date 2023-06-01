package config

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"rest-api-h8/entity"
)

//buat variable db untuk menyimpan gorm db

var DB *gorm.DB

func DBConnect() {
	//buat dsn untuk database nya
	dsn := "host=localhost user=postgres password=123456 dbname=rest_api_test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//buat variable (db) dan err handler nya.lalu open koneksi gorm nya
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//handle error
	if err != nil {
		fmt.Println("Failed connect to database")
		panic(err)
	} else {
		fmt.Println("Connected to database")
		DB = db
	}
	//migrate
	db.AutoMigrate(&entity.Product{}, &entity.Order{})
}
