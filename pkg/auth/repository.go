package auth

import (
	"context"
	"errors"

	"gorm.io/gorm"
)

var (
	ErrNotFound = gorm.ErrRecordNotFound
)

type IUserRepo interface {
	FindByEmail(context.Context, Email) (*User, error)
}

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{
		db: db,
	}
}

func (ur *UserRepo) FindByEmail(ctx context.Context, email Email) (*User, error) {
	var user User
	result := ur.db.WithContext(ctx).Where("email = ?", email).First(&user)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, ErrNotFound
	}

	return &user, nil
}
