package controllers

import (
	"os"

	// gorm hundle mysql
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"

	"database/sql"
	// short name for revel
	r "github.com/revel/revel"

	"test_app/app/models"
)

// type: revel controller with `*gorm.DB`
// c.Txn will keep `Gdb *gorm.DB`
type GormController struct {
	*r.Controller
	Txn *gorm.DB
}

// Gdb can be used for jobs
var Gdb *gorm.DB

// InitDB initializes the database.
func InitDB() {
	var err error
	var dsn = os.Getenv("REVELAPP_DBUSER") +
		":" + os.Getenv("REVELAPP_DBPASSWD") +
		"@" + os.Getenv("REVELAPP_DBHOSTNAME") +
		"/" + os.Getenv("REVELAPP_DBNAME") +
		"?parseTime=true&loc=Asia%2FTokyo"
	// open db
	Gdb, err = gorm.Open("mysql", dsn)
	if err != nil {
		println("FATAL", err)
		panic(err)
	}
	autoMigrate()
	// unique index if need
	//Gdb.Model(&models.User{}).AddUniqueIndex("idx_user_name", "name")
}

func autoMigrate() {
	Gdb.AutoMigrate(&models.Player{})
}

// transactions

// Begin method fills the c.Txn before each transaction
func (c *GormController) Begin() r.Result {
	txn := Gdb.Begin()
	if txn.Error != nil {
		panic(txn.Error)
	}
	c.Txn = txn
	return nil
}

// Commit method clears the c.Txn after each transaction
func (c *GormController) Commit() r.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Commit()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}

// Rollback method clears the c.Txn after each transaction, too
func (c *GormController) Rollback() r.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Rollback()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	c.Txn = nil
	return nil
}
