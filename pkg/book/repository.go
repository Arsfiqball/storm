package book

import (
	"context"
	"fmt"

	"gorm.io/gorm"
)

type BookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) *BookRepository {
	return &BookRepository{
		db: db,
	}
}

func (r *BookRepository) CreateOne(ctx context.Context, payload PayloadBook) (EntityBook, error) {
	entity := payload.ToEntity()
	selects := payload.GetSqlSelectFields()

	result := r.db.WithContext(ctx).
		Model(&entity).
		Select(selects[0], selects...).
		Create(&entity)

	if result.Error != nil {
		return EntityBook{}, fmt.Errorf("failed to insert: %w", result.Error)
	}

	return entity, nil
}
