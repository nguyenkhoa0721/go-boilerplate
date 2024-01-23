package postgres

import (
	"fmt"
	"go-boilerplate/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type Sql struct {
	*gorm.DB
}

func NewSql(config *config.Config) (*Sql, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", config.Postgres.Host, config.Postgres.Username, config.Postgres.Password, config.Postgres.Database, config.Postgres.Port)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		panic("failed to connect database")
	}

	return &Sql{db}, nil
}
