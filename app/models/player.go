package models

import "github.com/jinzhu/gorm"

// Player is test model
type Player struct {
	gorm.Model
	Name string
	Age  uint
}
