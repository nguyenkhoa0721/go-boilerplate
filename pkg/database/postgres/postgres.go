package postgres

import (
	"context"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go-boilerplate/config"
	"go-boilerplate/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Sql struct {
	*gorm.DB
}

func NewSql(config *config.Config) (*Sql, error) {
	dns := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d", config.Database.Host, config.Database.Username, config.Database.Password, config.Database.Database, config.Database.Port)

	db, err := gorm.Open(postgres.Open(dns), &gorm.Config{
		Logger: logger.NewGormLogger(),
	})
	if err != nil {
		panic("failed to connect database")
	}

	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Err(err).Msg("failed to use otel plugin")
	}

	return &Sql{db}, nil
}

func (s *Sql) WithCtx(ctx context.Context) *Sql {
	return &Sql{
		DB: s.Session(&gorm.Session{Context: ctx}),
	}
}
