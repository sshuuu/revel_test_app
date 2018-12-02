package models

import (
	"github.com/jinzhu/gorm"
	// for birthdahy
	"time"
)

// Player is test model
type Player struct {
	gorm.Model
	Name     string
	Age      uint
	Birthday time.Time `sql:"not null;type:date"`
}
