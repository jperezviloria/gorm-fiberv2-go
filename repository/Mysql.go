package repository

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//func MigrationStruct (db *gorm.DB){
//	db.AutoMigrate(&todo.Todo{})
//}

func ConnectMysql() *gorm.DB {
	dsn := "root:example@tcp(127.0.0.1:13306)/university?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		fmt.Println("Failed to connect with Mysql")
	}

	return db
}
