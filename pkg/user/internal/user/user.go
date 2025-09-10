package user

import (
	"app/pkg/kernel/typex"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID           uuid.UUID   `gorm:"column:id"`
	Email        typex.Email `gorm:"column:email"` // this is the primary email
	PasswordHash string      `gorm:"column:password_hash"`
}

func NewUser() User {
	return User{
		ID: uuid.New(),
	}
}

func (User) TableName() string {
	return "users"
}

func (u User) SaveTo(db *gorm.DB) error {
	return db.Save(&u).Error
}
