package provider

import (
	"context"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type GORM interface {
	DB() *gorm.DB
	Close(context.Context) error
}

type gormState struct {
	db *gorm.DB
}

func ProvideGORM(ctx context.Context) (GORM, error) {
	dsn := os.Getenv("DATABASE_URL")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(10)

	return &gormState{db: db}, nil
}

func (g *gormState) DB() *gorm.DB {
	return g.db
}

func (g *gormState) Close(ctx context.Context) error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
