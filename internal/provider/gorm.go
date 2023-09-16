package provider

import (
	"context"

	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	instance *gorm.DB
}

func ProvideGorm(config *viper.Viper) (*Gorm, error) {
	db, err := gorm.Open(postgres.Open(config.GetString("dsn")), &gorm.Config{})

	if err != nil {
		return nil, err
	}

	return &Gorm{
		instance: db,
	}, nil
}

func (g *Gorm) Clean(ctx context.Context) error {
	db, err := g.instance.DB()
	if err != nil {
		return err
	}

	return db.Close()
}

func (g *Gorm) Readiness(ctx context.Context) error {
	db, err := g.instance.DB()
	if err != nil {
		return err
	}

	return db.Ping()
}
