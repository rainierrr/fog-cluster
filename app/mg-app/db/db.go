package db

import (
	"log"
	"os"

	fogNode "mg-app/fog_node"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDB() *gorm.DB {
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	HOST := os.Getenv("MYSQL_HOST")
	DBNAME := os.Getenv("MYSQL_DATABASE")

	// GormでMySQLに接続

	dsn := USER + ":" + PASS + "@tcp(" + HOST + ":3306)/" + DBNAME + "?charset=utf8mb4&parseTime=True&loc=Local"
	log.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func AutoMigration() {
	db := GetDB()
	db.AutoMigrate(&fogNode.FogNode{})
}
