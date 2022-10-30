package user

import (
	"errors"

	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) {
	db.AutoMigrate(&UserEntity{})

	var admin UserEntity
	result := db.Where("email = ?", "admin@example.com").First(&admin)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		db.Create(&UserEntity{
			Email:    "admin@example.com",
			Password: "pass1234",
		})
	}
}
