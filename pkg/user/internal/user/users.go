package user

import (
	"github.com/Arsfiqball/talkback"
	"gorm.io/gorm"
)

type Users struct {
	Query talkback.Query
	Items []User
}

func (u Users) LoadFrom(db *gorm.DB) error {
	return db.Find(&u.Items).Error
}
