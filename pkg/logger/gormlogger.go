package logger

import (
	"context"
	"errors"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
)

type GormLogger struct {
	SlowThreshold         time.Duration
	SourceField           string
	SkipErrRecordNotFound bool
	Logger                zerolog.Logger
}

func NewGormLogger() *GormLogger {
	return &GormLogger{
		Logger:                log.Logger,
		SkipErrRecordNotFound: true,
	}
}

func NewWithLogger(l zerolog.Logger) *GormLogger {
	return &GormLogger{
		Logger: l,
	}
}

func (l *GormLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return l
}

func (l *GormLogger) Info(ctx context.Context, s string, args ...interface{}) {
	Info(ctx).Msgf(s, args)
}

func (l *GormLogger) Warn(ctx context.Context, s string, args ...interface{}) {
	Warn(ctx).Msgf(s, args)
}

func (l *GormLogger) Error(ctx context.Context, s string, args ...interface{}) {
	Error(ctx).Msgf(s, args)
}

func (l *GormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	elapsed := time.Since(begin)
	sql, _ := fc()
	fields := map[string]interface{}{
		"sql":      sql,
		"duration": elapsed,
	}
	if l.SourceField != "" {
		fields[l.SourceField] = utils.FileWithLineNum()
	}
	if err != nil && !(errors.Is(err, gorm.ErrRecordNotFound) && l.SkipErrRecordNotFound) {
		Error(ctx).Err(err).Fields(fields).Msg("[GORM] query error")
		return
	}

	if l.SlowThreshold != 0 && elapsed > l.SlowThreshold {
		Warn(ctx).Fields(fields).Msgf("[GORM] slow query")
		return
	}

	Debug(ctx).Fields(fields).Msgf("[GORM] query")
}
