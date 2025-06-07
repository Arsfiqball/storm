package provider

import (
	"context"
	"errors"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm interface {
	DB() *gorm.DB
	Close(context.Context) error
	Ping(context.Context) error
}

type gormState struct {
	db *gorm.DB
}

func ProvideGorm(ctx context.Context) (Gorm, error) {
	dsn := viper.GetString("database.url")
	if dsn == "" {
		return nil, errors.New("database URL is not set in configuration")
	}

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

func (g *gormState) Ping(ctx context.Context) error {
	sqlDB, err := g.db.DB()
	if err != nil {
		return err
	}

	return sqlDB.PingContext(ctx)
}
