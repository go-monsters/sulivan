package database

import (
	"database/sql"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		fmt.Println("db err: (Init) ", err)
		panic(err)
	}
	DB = db
	getDBConnection().SetMaxIdleConns(10)
	return DB
}
func getDBConnection() *sql.DB {
	sqlDB, err := DB.DB()
	if err != nil {
		fmt.Println("db err: (Init) ", err)
		panic(err)
	}
	return sqlDB;
}
func CloseDb()  {
	sqlDB, _ := DB.DB()
	sqlDB.Close()
}
