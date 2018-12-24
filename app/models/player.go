package models

import (
	"github.com/jinzhu/gorm"

	// for birthday
	"time"
)

// Player is test model
type Player struct {
	gorm.Model
	Name     string
	Birthday time.Time `sql:"not null;type:date"`
}

// Age return player age from birthday
func (p Player) Age() int {
	return int(time.Now().Sub(p.Birthday).Hours() / 24 / 365)
}
