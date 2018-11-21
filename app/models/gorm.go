package models

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	Db *gorm.DB
)

// InitDB initialize mariaDB configure
func InitDB() {
	var err error
	var dsn = os.Getenv("REVELAPP_DBUSER") +
		":" + os.Getenv("REVELAPP_DBPASSWD") +
		"@" + os.Getenv("REVELAPP_DBHOSTNAME") +
		"/" + os.Getenv("REVELAPP_DBNAME") +
		"?parseTime=true&loc=Asia%2FTokyo"

	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	Db.LogMode(true)
	autoMigrate()

	log.Println("Conncted to database.")
}

func autoMigrate() {
	Db.AutoMigrate(&Player{})
}
