package models

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
)

var gGormDB *gorm.DB

func DBInit() (err error) {
	db, err := gorm.Open("mysql", dbConnString())
	if err != nil {
		log.Println("[DB] ", err)
		return
	}
	gGormDB = db

	log.Println("[DB] Start DB Migration ... ")
	db.AutoMigrate(&Users{})
	log.Println("[DB] Start DB ... ")

	return
}

func dbConnString() (dbConnString string) {
	dbHost := viper.GetString(`database.host`)
	dbUser := viper.GetString(`database.user`)
	dbPass := viper.GetString(`database.pass`)
	dbName := viper.GetString(`database.name`)

	if dbHost == "" || dbUser == "" || dbPass == "" || dbName == "" {
		log.Println("Cannot get db env values, check .env .yml file")
		return
	}

	dbConnString = fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPass, dbHost, dbName)

	return
}

func CloseDB() {
	if gGormDB != nil {
		gGormDB.Close()
	}
}
