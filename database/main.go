package database

import (
	"database/sql"
	"gorm.io/gorm"
	"gorm.io/driver/sqlite"
	"log"
)

var DB *gorm.DB

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	DB = db
	getDBConnection().SetMaxIdleConns(10)
	return DB
}
func getDBConnection() *sql.DB {
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return sqlDB;
}
func CloseDb()  {
	getDBConnection().Close()
}
