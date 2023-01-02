package provider

import (
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ProvideGorm(config *viper.Viper) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.GetString("dsn")), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
